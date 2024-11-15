// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Alert struct {
	Body     *string `json:"body,omitempty" xml:"body,omitempty"`
	Subtitle *string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s Alert) String() string {
	return tea.Prettify(s)
}

func (s Alert) GoString() string {
	return s.String()
}

func (s *Alert) SetBody(v string) *Alert {
	s.Body = &v
	return s
}

func (s *Alert) SetSubtitle(v string) *Alert {
	s.Subtitle = &v
	return s
}

func (s *Alert) SetTitle(v string) *Alert {
	s.Title = &v
	return s
}

type AndroidPayload struct {
	Body                 *Body                  `json:"body,omitempty" xml:"body,omitempty"`
	DisplayType          *string                `json:"displayType,omitempty" xml:"displayType,omitempty"`
	Extra                map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
	Message2ThirdChannel *Message2ThirdChannel  `json:"message2ThirdChannel,omitempty" xml:"message2ThirdChannel,omitempty"`
}

func (s AndroidPayload) String() string {
	return tea.Prettify(s)
}

func (s AndroidPayload) GoString() string {
	return s.String()
}

func (s *AndroidPayload) SetBody(v *Body) *AndroidPayload {
	s.Body = v
	return s
}

func (s *AndroidPayload) SetDisplayType(v string) *AndroidPayload {
	s.DisplayType = &v
	return s
}

func (s *AndroidPayload) SetExtra(v map[string]interface{}) *AndroidPayload {
	s.Extra = v
	return s
}

func (s *AndroidPayload) SetMessage2ThirdChannel(v *Message2ThirdChannel) *AndroidPayload {
	s.Message2ThirdChannel = v
	return s
}

type AndroidShortPayload struct {
	Body  *AndroidShortPayloadBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	Extra map[string]interface{}   `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s AndroidShortPayload) String() string {
	return tea.Prettify(s)
}

func (s AndroidShortPayload) GoString() string {
	return s.String()
}

func (s *AndroidShortPayload) SetBody(v *AndroidShortPayloadBody) *AndroidShortPayload {
	s.Body = v
	return s
}

func (s *AndroidShortPayload) SetExtra(v map[string]interface{}) *AndroidShortPayload {
	s.Extra = v
	return s
}

type AndroidShortPayloadBody struct {
	Custom *string `json:"custom,omitempty" xml:"custom,omitempty"`
}

func (s AndroidShortPayloadBody) String() string {
	return tea.Prettify(s)
}

func (s AndroidShortPayloadBody) GoString() string {
	return s.String()
}

func (s *AndroidShortPayloadBody) SetCustom(v string) *AndroidShortPayloadBody {
	s.Custom = &v
	return s
}

type Aps struct {
	Alert *Alert `json:"alert,omitempty" xml:"alert,omitempty"`
	// example:
	//
	// +1(自增)，-1(自减)，4(设置数字)
	Badge             *string `json:"badge,omitempty" xml:"badge,omitempty"`
	Category          *string `json:"category,omitempty" xml:"category,omitempty"`
	ContentAvailable  *int32  `json:"contentAvailable,omitempty" xml:"contentAvailable,omitempty"`
	InterruptionLevel *string `json:"interruptionLevel,omitempty" xml:"interruptionLevel,omitempty"`
	Sound             *string `json:"sound,omitempty" xml:"sound,omitempty"`
	ThreadID          *string `json:"threadID,omitempty" xml:"threadID,omitempty"`
}

func (s Aps) String() string {
	return tea.Prettify(s)
}

func (s Aps) GoString() string {
	return s.String()
}

func (s *Aps) SetAlert(v *Alert) *Aps {
	s.Alert = v
	return s
}

func (s *Aps) SetBadge(v string) *Aps {
	s.Badge = &v
	return s
}

func (s *Aps) SetCategory(v string) *Aps {
	s.Category = &v
	return s
}

func (s *Aps) SetContentAvailable(v int32) *Aps {
	s.ContentAvailable = &v
	return s
}

func (s *Aps) SetInterruptionLevel(v string) *Aps {
	s.InterruptionLevel = &v
	return s
}

func (s *Aps) SetSound(v string) *Aps {
	s.Sound = &v
	return s
}

func (s *Aps) SetThreadID(v string) *Aps {
	s.ThreadID = &v
	return s
}

type Body struct {
	Activity    *string `json:"activity,omitempty" xml:"activity,omitempty"`
	AddBadge    *int32  `json:"addBadge,omitempty" xml:"addBadge,omitempty"`
	AfterOpen   *string `json:"afterOpen,omitempty" xml:"afterOpen,omitempty"`
	BuilderId   *int64  `json:"builderId,omitempty" xml:"builderId,omitempty"`
	Custom      *string `json:"custom,omitempty" xml:"custom,omitempty"`
	ExpandImage *string `json:"expandImage,omitempty" xml:"expandImage,omitempty"`
	Icon        *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Img         *string `json:"img,omitempty" xml:"img,omitempty"`
	PlayLights  *bool   `json:"playLights,omitempty" xml:"playLights,omitempty"`
	PlaySound   *bool   `json:"playSound,omitempty" xml:"playSound,omitempty"`
	PlayVibrate *bool   `json:"playVibrate,omitempty" xml:"playVibrate,omitempty"`
	RePop       *int32  `json:"rePop,omitempty" xml:"rePop,omitempty"`
	SetBadge    *int32  `json:"setBadge,omitempty" xml:"setBadge,omitempty"`
	Sound       *string `json:"sound,omitempty" xml:"sound,omitempty"`
	Text        *string `json:"text,omitempty" xml:"text,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s Body) String() string {
	return tea.Prettify(s)
}

func (s Body) GoString() string {
	return s.String()
}

func (s *Body) SetActivity(v string) *Body {
	s.Activity = &v
	return s
}

func (s *Body) SetAddBadge(v int32) *Body {
	s.AddBadge = &v
	return s
}

func (s *Body) SetAfterOpen(v string) *Body {
	s.AfterOpen = &v
	return s
}

func (s *Body) SetBuilderId(v int64) *Body {
	s.BuilderId = &v
	return s
}

func (s *Body) SetCustom(v string) *Body {
	s.Custom = &v
	return s
}

func (s *Body) SetExpandImage(v string) *Body {
	s.ExpandImage = &v
	return s
}

func (s *Body) SetIcon(v string) *Body {
	s.Icon = &v
	return s
}

func (s *Body) SetImg(v string) *Body {
	s.Img = &v
	return s
}

func (s *Body) SetPlayLights(v bool) *Body {
	s.PlayLights = &v
	return s
}

func (s *Body) SetPlaySound(v bool) *Body {
	s.PlaySound = &v
	return s
}

func (s *Body) SetPlayVibrate(v bool) *Body {
	s.PlayVibrate = &v
	return s
}

func (s *Body) SetRePop(v int32) *Body {
	s.RePop = &v
	return s
}

func (s *Body) SetSetBadge(v int32) *Body {
	s.SetBadge = &v
	return s
}

func (s *Body) SetSound(v string) *Body {
	s.Sound = &v
	return s
}

func (s *Body) SetText(v string) *Body {
	s.Text = &v
	return s
}

func (s *Body) SetTitle(v string) *Body {
	s.Title = &v
	return s
}

func (s *Body) SetUrl(v string) *Body {
	s.Url = &v
	return s
}

type ChannelProperties struct {
	ChannelActivity         *string `json:"channelActivity,omitempty" xml:"channelActivity,omitempty"`
	ChannelFcm              *string `json:"channelFcm,omitempty" xml:"channelFcm,omitempty"`
	HuaweiChannelCategory   *string `json:"huaweiChannelCategory,omitempty" xml:"huaweiChannelCategory,omitempty"`
	HuaweiChannelImportance *string `json:"huaweiChannelImportance,omitempty" xml:"huaweiChannelImportance,omitempty"`
	// example:
	//
	// 取值为"NORMAL"和"HIGH",默认为”NORMAL”
	HuaweiMessageUrgency *string `json:"huaweiMessageUrgency,omitempty" xml:"huaweiMessageUrgency,omitempty"`
	MainActivity         *string `json:"mainActivity,omitempty" xml:"mainActivity,omitempty"`
	OppoCategory         *string `json:"oppoCategory,omitempty" xml:"oppoCategory,omitempty"`
	OppoChannelId        *string `json:"oppoChannelId,omitempty" xml:"oppoChannelId,omitempty"`
	OppoNotifyLevel      *string `json:"oppoNotifyLevel,omitempty" xml:"oppoNotifyLevel,omitempty"`
	// example:
	//
	// "true" ,默认为"false"，可不填
	UseHuaweiMessage *string `json:"useHuaweiMessage,omitempty" xml:"useHuaweiMessage,omitempty"`
	// example:
	//
	// true
	UseHuaweiPlainMessage *string `json:"useHuaweiPlainMessage,omitempty" xml:"useHuaweiPlainMessage,omitempty"`
	// example:
	//
	// "true",默认"false"
	VivoAddBadge    *string `json:"vivoAddBadge,omitempty" xml:"vivoAddBadge,omitempty"`
	VivoCategory    *string `json:"vivoCategory,omitempty" xml:"vivoCategory,omitempty"`
	VivoPushMode    *string `json:"vivoPushMode,omitempty" xml:"vivoPushMode,omitempty"`
	XiaomiChannelId *string `json:"xiaomiChannelId,omitempty" xml:"xiaomiChannelId,omitempty"`
}

func (s ChannelProperties) String() string {
	return tea.Prettify(s)
}

func (s ChannelProperties) GoString() string {
	return s.String()
}

func (s *ChannelProperties) SetChannelActivity(v string) *ChannelProperties {
	s.ChannelActivity = &v
	return s
}

func (s *ChannelProperties) SetChannelFcm(v string) *ChannelProperties {
	s.ChannelFcm = &v
	return s
}

func (s *ChannelProperties) SetHuaweiChannelCategory(v string) *ChannelProperties {
	s.HuaweiChannelCategory = &v
	return s
}

func (s *ChannelProperties) SetHuaweiChannelImportance(v string) *ChannelProperties {
	s.HuaweiChannelImportance = &v
	return s
}

func (s *ChannelProperties) SetHuaweiMessageUrgency(v string) *ChannelProperties {
	s.HuaweiMessageUrgency = &v
	return s
}

func (s *ChannelProperties) SetMainActivity(v string) *ChannelProperties {
	s.MainActivity = &v
	return s
}

func (s *ChannelProperties) SetOppoCategory(v string) *ChannelProperties {
	s.OppoCategory = &v
	return s
}

func (s *ChannelProperties) SetOppoChannelId(v string) *ChannelProperties {
	s.OppoChannelId = &v
	return s
}

func (s *ChannelProperties) SetOppoNotifyLevel(v string) *ChannelProperties {
	s.OppoNotifyLevel = &v
	return s
}

func (s *ChannelProperties) SetUseHuaweiMessage(v string) *ChannelProperties {
	s.UseHuaweiMessage = &v
	return s
}

func (s *ChannelProperties) SetUseHuaweiPlainMessage(v string) *ChannelProperties {
	s.UseHuaweiPlainMessage = &v
	return s
}

func (s *ChannelProperties) SetVivoAddBadge(v string) *ChannelProperties {
	s.VivoAddBadge = &v
	return s
}

func (s *ChannelProperties) SetVivoCategory(v string) *ChannelProperties {
	s.VivoCategory = &v
	return s
}

func (s *ChannelProperties) SetVivoPushMode(v string) *ChannelProperties {
	s.VivoPushMode = &v
	return s
}

func (s *ChannelProperties) SetXiaomiChannelId(v string) *ChannelProperties {
	s.XiaomiChannelId = &v
	return s
}

type IosPayload struct {
	Aps   *Aps                   `json:"aps,omitempty" xml:"aps,omitempty"`
	Extra map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s IosPayload) String() string {
	return tea.Prettify(s)
}

func (s IosPayload) GoString() string {
	return s.String()
}

func (s *IosPayload) SetAps(v *Aps) *IosPayload {
	s.Aps = v
	return s
}

func (s *IosPayload) SetExtra(v map[string]interface{}) *IosPayload {
	s.Extra = v
	return s
}

type Message2ThirdChannel struct {
	SetBadge    *int64  `json:"SetBadge,omitempty" xml:"SetBadge,omitempty"`
	AddBadge    *int64  `json:"addBadge,omitempty" xml:"addBadge,omitempty"`
	BigBody     *string `json:"bigBody,omitempty" xml:"bigBody,omitempty"`
	BigTitle    *string `json:"bigTitle,omitempty" xml:"bigTitle,omitempty"`
	ExpandImage *string `json:"expandImage,omitempty" xml:"expandImage,omitempty"`
	Img         *string `json:"img,omitempty" xml:"img,omitempty"`
	Sound       *string `json:"sound,omitempty" xml:"sound,omitempty"`
	Text        *string `json:"text,omitempty" xml:"text,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s Message2ThirdChannel) String() string {
	return tea.Prettify(s)
}

func (s Message2ThirdChannel) GoString() string {
	return s.String()
}

func (s *Message2ThirdChannel) SetSetBadge(v int64) *Message2ThirdChannel {
	s.SetBadge = &v
	return s
}

func (s *Message2ThirdChannel) SetAddBadge(v int64) *Message2ThirdChannel {
	s.AddBadge = &v
	return s
}

func (s *Message2ThirdChannel) SetBigBody(v string) *Message2ThirdChannel {
	s.BigBody = &v
	return s
}

func (s *Message2ThirdChannel) SetBigTitle(v string) *Message2ThirdChannel {
	s.BigTitle = &v
	return s
}

func (s *Message2ThirdChannel) SetExpandImage(v string) *Message2ThirdChannel {
	s.ExpandImage = &v
	return s
}

func (s *Message2ThirdChannel) SetImg(v string) *Message2ThirdChannel {
	s.Img = &v
	return s
}

func (s *Message2ThirdChannel) SetSound(v string) *Message2ThirdChannel {
	s.Sound = &v
	return s
}

func (s *Message2ThirdChannel) SetText(v string) *Message2ThirdChannel {
	s.Text = &v
	return s
}

func (s *Message2ThirdChannel) SetTitle(v string) *Message2ThirdChannel {
	s.Title = &v
	return s
}

type Policy struct {
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	OuterBizNo *string `json:"outerBizNo,omitempty" xml:"outerBizNo,omitempty"`
	// example:
	//
	// 5000
	Speed *int32 `json:"speed,omitempty" xml:"speed,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s Policy) String() string {
	return tea.Prettify(s)
}

func (s Policy) GoString() string {
	return s.String()
}

func (s *Policy) SetExpireTime(v string) *Policy {
	s.ExpireTime = &v
	return s
}

func (s *Policy) SetOuterBizNo(v string) *Policy {
	s.OuterBizNo = &v
	return s
}

func (s *Policy) SetSpeed(v int32) *Policy {
	s.Speed = &v
	return s
}

func (s *Policy) SetStartTime(v string) *Policy {
	s.StartTime = &v
	return s
}

type CancelByMsgIdRequest struct {
	// example:
	//
	// ucj0242167047014687101
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s CancelByMsgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *CancelByMsgIdRequest) SetMsgId(v string) *CancelByMsgIdRequest {
	s.MsgId = &v
	return s
}

type CancelByMsgIdResponseBody struct {
	// example:
	//
	// 0
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CancelByMsgIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelByMsgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *CancelByMsgIdResponseBody) SetCode(v string) *CancelByMsgIdResponseBody {
	s.Code = &v
	return s
}

func (s *CancelByMsgIdResponseBody) SetData(v *CancelByMsgIdResponseBodyData) *CancelByMsgIdResponseBody {
	s.Data = v
	return s
}

func (s *CancelByMsgIdResponseBody) SetHttpStatusCode(v int32) *CancelByMsgIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelByMsgIdResponseBody) SetMessage(v string) *CancelByMsgIdResponseBody {
	s.Message = &v
	return s
}

func (s *CancelByMsgIdResponseBody) SetRequestId(v string) *CancelByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelByMsgIdResponseBody) SetSuccess(v bool) *CancelByMsgIdResponseBody {
	s.Success = &v
	return s
}

type CancelByMsgIdResponseBodyData struct {
	// example:
	//
	// ucj0242167047014687101
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s CancelByMsgIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CancelByMsgIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelByMsgIdResponseBodyData) SetMsgId(v string) *CancelByMsgIdResponseBodyData {
	s.MsgId = &v
	return s
}

type CancelByMsgIdResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelByMsgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelByMsgIdResponse) GoString() string {
	return s.String()
}

func (s *CancelByMsgIdResponse) SetHeaders(v map[string]*string) *CancelByMsgIdResponse {
	s.Headers = v
	return s
}

func (s *CancelByMsgIdResponse) SetStatusCode(v int32) *CancelByMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelByMsgIdResponse) SetBody(v *CancelByMsgIdResponseBody) *CancelByMsgIdResponse {
	s.Body = v
	return s
}

type QueryMsgStatRequest struct {
	// example:
	//
	// ufe29y2167046828041801
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s QueryMsgStatRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMsgStatRequest) GoString() string {
	return s.String()
}

func (s *QueryMsgStatRequest) SetMsgId(v string) *QueryMsgStatRequest {
	s.MsgId = &v
	return s
}

type QueryMsgStatResponseBody struct {
	// example:
	//
	// 0
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryMsgStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMsgStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMsgStatResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMsgStatResponseBody) SetCode(v string) *QueryMsgStatResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMsgStatResponseBody) SetData(v *QueryMsgStatResponseBodyData) *QueryMsgStatResponseBody {
	s.Data = v
	return s
}

func (s *QueryMsgStatResponseBody) SetHttpStatusCode(v int32) *QueryMsgStatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryMsgStatResponseBody) SetMessage(v string) *QueryMsgStatResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMsgStatResponseBody) SetRequestId(v string) *QueryMsgStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMsgStatResponseBody) SetSuccess(v bool) *QueryMsgStatResponseBody {
	s.Success = &v
	return s
}

type QueryMsgStatResponseBodyData struct {
	// example:
	//
	// 1
	Accept *int64 `json:"Accept,omitempty" xml:"Accept,omitempty"`
	// example:
	//
	// 1
	Arrive *int64 `json:"Arrive,omitempty" xml:"Arrive,omitempty"`
	// example:
	//
	// 0
	ClosePush *int64 `json:"ClosePush,omitempty" xml:"ClosePush,omitempty"`
	// example:
	//
	// 0
	Dismiss *int64 `json:"Dismiss,omitempty" xml:"Dismiss,omitempty"`
	// example:
	//
	// ufe29y2167046828041801
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// example:
	//
	// 1
	Open *int64 `json:"Open,omitempty" xml:"Open,omitempty"`
	// example:
	//
	// 1
	Sent *int64 `json:"Sent,omitempty" xml:"Sent,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryMsgStatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMsgStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMsgStatResponseBodyData) SetAccept(v int64) *QueryMsgStatResponseBodyData {
	s.Accept = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetArrive(v int64) *QueryMsgStatResponseBodyData {
	s.Arrive = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetClosePush(v int64) *QueryMsgStatResponseBodyData {
	s.ClosePush = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetDismiss(v int64) *QueryMsgStatResponseBodyData {
	s.Dismiss = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetMsgId(v string) *QueryMsgStatResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetOpen(v int64) *QueryMsgStatResponseBodyData {
	s.Open = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetSent(v int64) *QueryMsgStatResponseBodyData {
	s.Sent = &v
	return s
}

func (s *QueryMsgStatResponseBodyData) SetStatus(v int32) *QueryMsgStatResponseBodyData {
	s.Status = &v
	return s
}

type QueryMsgStatResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMsgStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMsgStatResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMsgStatResponse) GoString() string {
	return s.String()
}

func (s *QueryMsgStatResponse) SetHeaders(v map[string]*string) *QueryMsgStatResponse {
	s.Headers = v
	return s
}

func (s *QueryMsgStatResponse) SetStatusCode(v int32) *QueryMsgStatResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMsgStatResponse) SetBody(v *QueryMsgStatResponseBody) *QueryMsgStatResponse {
	s.Body = v
	return s
}

type SendByAliasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Alias               *string              `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AliasType           *string              `json:"AliasType,omitempty" xml:"AliasType,omitempty"`
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	IosPayload          *IosPayload          `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy              *Policy              `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasRequest) GoString() string {
	return s.String()
}

func (s *SendByAliasRequest) SetAlias(v string) *SendByAliasRequest {
	s.Alias = &v
	return s
}

func (s *SendByAliasRequest) SetAliasType(v string) *SendByAliasRequest {
	s.AliasType = &v
	return s
}

func (s *SendByAliasRequest) SetAndroidPayload(v *AndroidPayload) *SendByAliasRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByAliasRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByAliasRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByAliasRequest) SetChannelProperties(v *ChannelProperties) *SendByAliasRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByAliasRequest) SetDescription(v string) *SendByAliasRequest {
	s.Description = &v
	return s
}

func (s *SendByAliasRequest) SetIosPayload(v *IosPayload) *SendByAliasRequest {
	s.IosPayload = v
	return s
}

func (s *SendByAliasRequest) SetPolicy(v *Policy) *SendByAliasRequest {
	s.Policy = v
	return s
}

func (s *SendByAliasRequest) SetProductionMode(v bool) *SendByAliasRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAliasRequest) SetReceiptType(v int32) *SendByAliasRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAliasRequest) SetReceiptUrl(v string) *SendByAliasRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAliasRequest) SetThirdPartyId(v string) *SendByAliasRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAliasRequest) SetCallbackParams(v string) *SendByAliasRequest {
	s.CallbackParams = &v
	return s
}

type SendByAliasShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Alias                     *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AliasType                 *string `json:"AliasType,omitempty" xml:"AliasType,omitempty"`
	AndroidPayloadShrink      *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayloadShrink *string `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink   *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IosPayloadShrink          *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink              *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByAliasShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByAliasShrinkRequest) SetAlias(v string) *SendByAliasShrinkRequest {
	s.Alias = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetAliasType(v string) *SendByAliasShrinkRequest {
	s.AliasType = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetAndroidPayloadShrink(v string) *SendByAliasShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByAliasShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetChannelPropertiesShrink(v string) *SendByAliasShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetDescription(v string) *SendByAliasShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetIosPayloadShrink(v string) *SendByAliasShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetPolicyShrink(v string) *SendByAliasShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetProductionMode(v bool) *SendByAliasShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetReceiptType(v int32) *SendByAliasShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetReceiptUrl(v string) *SendByAliasShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetThirdPartyId(v string) *SendByAliasShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAliasShrinkRequest) SetCallbackParams(v string) *SendByAliasShrinkRequest {
	s.CallbackParams = &v
	return s
}

type SendByAliasResponseBody struct {
	// example:
	//
	// 0
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByAliasResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasResponseBody) GoString() string {
	return s.String()
}

func (s *SendByAliasResponseBody) SetCode(v string) *SendByAliasResponseBody {
	s.Code = &v
	return s
}

func (s *SendByAliasResponseBody) SetData(v *SendByAliasResponseBodyData) *SendByAliasResponseBody {
	s.Data = v
	return s
}

func (s *SendByAliasResponseBody) SetHttpStatusCode(v int32) *SendByAliasResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByAliasResponseBody) SetMessage(v string) *SendByAliasResponseBody {
	s.Message = &v
	return s
}

func (s *SendByAliasResponseBody) SetRequestId(v string) *SendByAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByAliasResponseBody) SetSuccess(v bool) *SendByAliasResponseBody {
	s.Success = &v
	return s
}

type SendByAliasResponseBodyData struct {
	// example:
	//
	// uacxo27167041814609201
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByAliasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByAliasResponseBodyData) SetMsgId(v string) *SendByAliasResponseBodyData {
	s.MsgId = &v
	return s
}

type SendByAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasResponse) GoString() string {
	return s.String()
}

func (s *SendByAliasResponse) SetHeaders(v map[string]*string) *SendByAliasResponse {
	s.Headers = v
	return s
}

func (s *SendByAliasResponse) SetStatusCode(v int32) *SendByAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByAliasResponse) SetBody(v *SendByAliasResponseBody) *SendByAliasResponse {
	s.Body = v
	return s
}

type SendByAliasFileIdRequest struct {
	AliasType           *string              `json:"AliasType,omitempty" xml:"AliasType,omitempty"`
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PF835431668603208261
	FileId     *string     `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IosPayload *IosPayload `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy     *Policy     `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByAliasFileIdRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasFileIdRequest) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdRequest) SetAliasType(v string) *SendByAliasFileIdRequest {
	s.AliasType = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetAndroidPayload(v *AndroidPayload) *SendByAliasFileIdRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByAliasFileIdRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByAliasFileIdRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByAliasFileIdRequest) SetChannelProperties(v *ChannelProperties) *SendByAliasFileIdRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByAliasFileIdRequest) SetDescription(v string) *SendByAliasFileIdRequest {
	s.Description = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetFileId(v string) *SendByAliasFileIdRequest {
	s.FileId = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetIosPayload(v *IosPayload) *SendByAliasFileIdRequest {
	s.IosPayload = v
	return s
}

func (s *SendByAliasFileIdRequest) SetPolicy(v *Policy) *SendByAliasFileIdRequest {
	s.Policy = v
	return s
}

func (s *SendByAliasFileIdRequest) SetProductionMode(v bool) *SendByAliasFileIdRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetReceiptType(v int32) *SendByAliasFileIdRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetReceiptUrl(v string) *SendByAliasFileIdRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetThirdPartyId(v string) *SendByAliasFileIdRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAliasFileIdRequest) SetCallbackParams(v string) *SendByAliasFileIdRequest {
	s.CallbackParams = &v
	return s
}

type SendByAliasFileIdShrinkRequest struct {
	AliasType                 *string `json:"AliasType,omitempty" xml:"AliasType,omitempty"`
	AndroidPayloadShrink      *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayloadShrink *string `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink   *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PF835431668603208261
	FileId           *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IosPayloadShrink *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink     *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByAliasFileIdShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasFileIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdShrinkRequest) SetAliasType(v string) *SendByAliasFileIdShrinkRequest {
	s.AliasType = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetAndroidPayloadShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetChannelPropertiesShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetDescription(v string) *SendByAliasFileIdShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetFileId(v string) *SendByAliasFileIdShrinkRequest {
	s.FileId = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetIosPayloadShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetPolicyShrink(v string) *SendByAliasFileIdShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetProductionMode(v bool) *SendByAliasFileIdShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetReceiptType(v int32) *SendByAliasFileIdShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetReceiptUrl(v string) *SendByAliasFileIdShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetThirdPartyId(v string) *SendByAliasFileIdShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAliasFileIdShrinkRequest) SetCallbackParams(v string) *SendByAliasFileIdShrinkRequest {
	s.CallbackParams = &v
	return s
}

type SendByAliasFileIdResponseBody struct {
	// example:
	//
	// 0
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByAliasFileIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByAliasFileIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasFileIdResponseBody) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdResponseBody) SetCode(v string) *SendByAliasFileIdResponseBody {
	s.Code = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetData(v *SendByAliasFileIdResponseBodyData) *SendByAliasFileIdResponseBody {
	s.Data = v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetHttpStatusCode(v int32) *SendByAliasFileIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetMessage(v string) *SendByAliasFileIdResponseBody {
	s.Message = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetRequestId(v string) *SendByAliasFileIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetSuccess(v bool) *SendByAliasFileIdResponseBody {
	s.Success = &v
	return s
}

type SendByAliasFileIdResponseBodyData struct {
	// example:
	//
	// ucj0242167047014687101
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByAliasFileIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasFileIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdResponseBodyData) SetMsgId(v string) *SendByAliasFileIdResponseBodyData {
	s.MsgId = &v
	return s
}

type SendByAliasFileIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByAliasFileIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByAliasFileIdResponse) String() string {
	return tea.Prettify(s)
}

func (s SendByAliasFileIdResponse) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdResponse) SetHeaders(v map[string]*string) *SendByAliasFileIdResponse {
	s.Headers = v
	return s
}

func (s *SendByAliasFileIdResponse) SetStatusCode(v int32) *SendByAliasFileIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByAliasFileIdResponse) SetBody(v *SendByAliasFileIdResponseBody) *SendByAliasFileIdResponse {
	s.Body = v
	return s
}

type SendByAppRequest struct {
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	IosPayload          *IosPayload          `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy              *Policy              `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByAppRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByAppRequest) GoString() string {
	return s.String()
}

func (s *SendByAppRequest) SetAndroidPayload(v *AndroidPayload) *SendByAppRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByAppRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByAppRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByAppRequest) SetChannelProperties(v *ChannelProperties) *SendByAppRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByAppRequest) SetDescription(v string) *SendByAppRequest {
	s.Description = &v
	return s
}

func (s *SendByAppRequest) SetIosPayload(v *IosPayload) *SendByAppRequest {
	s.IosPayload = v
	return s
}

func (s *SendByAppRequest) SetPolicy(v *Policy) *SendByAppRequest {
	s.Policy = v
	return s
}

func (s *SendByAppRequest) SetProductionMode(v bool) *SendByAppRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAppRequest) SetReceiptType(v int32) *SendByAppRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAppRequest) SetReceiptUrl(v string) *SendByAppRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAppRequest) SetThirdPartyId(v string) *SendByAppRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAppRequest) SetCallbackParams(v string) *SendByAppRequest {
	s.CallbackParams = &v
	return s
}

type SendByAppShrinkRequest struct {
	AndroidPayloadShrink      *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayloadShrink *string `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink   *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IosPayloadShrink          *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink              *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByAppShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByAppShrinkRequest) SetAndroidPayloadShrink(v string) *SendByAppShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByAppShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetChannelPropertiesShrink(v string) *SendByAppShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetDescription(v string) *SendByAppShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByAppShrinkRequest) SetIosPayloadShrink(v string) *SendByAppShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetPolicyShrink(v string) *SendByAppShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByAppShrinkRequest) SetProductionMode(v bool) *SendByAppShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByAppShrinkRequest) SetReceiptType(v int32) *SendByAppShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByAppShrinkRequest) SetReceiptUrl(v string) *SendByAppShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByAppShrinkRequest) SetThirdPartyId(v string) *SendByAppShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByAppShrinkRequest) SetCallbackParams(v string) *SendByAppShrinkRequest {
	s.CallbackParams = &v
	return s
}

type SendByAppResponseBody struct {
	// example:
	//
	// 0
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendByAppResponseBody) GoString() string {
	return s.String()
}

func (s *SendByAppResponseBody) SetCode(v string) *SendByAppResponseBody {
	s.Code = &v
	return s
}

func (s *SendByAppResponseBody) SetData(v *SendByAppResponseBodyData) *SendByAppResponseBody {
	s.Data = v
	return s
}

func (s *SendByAppResponseBody) SetHttpStatusCode(v int32) *SendByAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByAppResponseBody) SetMessage(v string) *SendByAppResponseBody {
	s.Message = &v
	return s
}

func (s *SendByAppResponseBody) SetRequestId(v string) *SendByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByAppResponseBody) SetSuccess(v bool) *SendByAppResponseBody {
	s.Success = &v
	return s
}

type SendByAppResponseBodyData struct {
	// example:
	//
	// um3zlgb166876370784300
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendByAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByAppResponseBodyData) SetMsgId(v string) *SendByAppResponseBodyData {
	s.MsgId = &v
	return s
}

type SendByAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByAppResponse) String() string {
	return tea.Prettify(s)
}

func (s SendByAppResponse) GoString() string {
	return s.String()
}

func (s *SendByAppResponse) SetHeaders(v map[string]*string) *SendByAppResponse {
	s.Headers = v
	return s
}

func (s *SendByAppResponse) SetStatusCode(v int32) *SendByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByAppResponse) SetBody(v *SendByAppResponseBody) *SendByAppResponse {
	s.Body = v
	return s
}

type SendByDeviceRequest struct {
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ArdNyIzFCH2K3szXA8arpu0Y7ywOdA67mCSumtpnMnmf
	DeviceTokens *string     `json:"DeviceTokens,omitempty" xml:"DeviceTokens,omitempty"`
	IosPayload   *IosPayload `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy       *Policy     `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceRequest) GoString() string {
	return s.String()
}

func (s *SendByDeviceRequest) SetAndroidPayload(v *AndroidPayload) *SendByDeviceRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByDeviceRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByDeviceRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByDeviceRequest) SetChannelProperties(v *ChannelProperties) *SendByDeviceRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByDeviceRequest) SetDescription(v string) *SendByDeviceRequest {
	s.Description = &v
	return s
}

func (s *SendByDeviceRequest) SetDeviceTokens(v string) *SendByDeviceRequest {
	s.DeviceTokens = &v
	return s
}

func (s *SendByDeviceRequest) SetIosPayload(v *IosPayload) *SendByDeviceRequest {
	s.IosPayload = v
	return s
}

func (s *SendByDeviceRequest) SetPolicy(v *Policy) *SendByDeviceRequest {
	s.Policy = v
	return s
}

func (s *SendByDeviceRequest) SetProductionMode(v bool) *SendByDeviceRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByDeviceRequest) SetReceiptType(v int32) *SendByDeviceRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByDeviceRequest) SetReceiptUrl(v string) *SendByDeviceRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByDeviceRequest) SetThirdPartyId(v string) *SendByDeviceRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByDeviceRequest) SetCallbackParams(v string) *SendByDeviceRequest {
	s.CallbackParams = &v
	return s
}

type SendByDeviceShrinkRequest struct {
	AndroidPayloadShrink      *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayloadShrink *string `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink   *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ArdNyIzFCH2K3szXA8arpu0Y7ywOdA67mCSumtpnMnmf
	DeviceTokens     *string `json:"DeviceTokens,omitempty" xml:"DeviceTokens,omitempty"`
	IosPayloadShrink *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink     *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByDeviceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByDeviceShrinkRequest) SetAndroidPayloadShrink(v string) *SendByDeviceShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByDeviceShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetChannelPropertiesShrink(v string) *SendByDeviceShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetDescription(v string) *SendByDeviceShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetDeviceTokens(v string) *SendByDeviceShrinkRequest {
	s.DeviceTokens = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetIosPayloadShrink(v string) *SendByDeviceShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetPolicyShrink(v string) *SendByDeviceShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetProductionMode(v bool) *SendByDeviceShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetReceiptType(v int32) *SendByDeviceShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetReceiptUrl(v string) *SendByDeviceShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetThirdPartyId(v string) *SendByDeviceShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByDeviceShrinkRequest) SetCallbackParams(v string) *SendByDeviceShrinkRequest {
	s.CallbackParams = &v
	return s
}

type SendByDeviceResponseBody struct {
	// example:
	//
	// 0
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 内部错误
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 74808AA4-A044-102F-8F5F-AFE4D97A0F26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *SendByDeviceResponseBody) SetCode(v string) *SendByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *SendByDeviceResponseBody) SetData(v *SendByDeviceResponseBodyData) *SendByDeviceResponseBody {
	s.Data = v
	return s
}

func (s *SendByDeviceResponseBody) SetHttpStatusCode(v int32) *SendByDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByDeviceResponseBody) SetMessage(v string) *SendByDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *SendByDeviceResponseBody) SetRequestId(v string) *SendByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByDeviceResponseBody) SetSuccess(v bool) *SendByDeviceResponseBody {
	s.Success = &v
	return s
}

type SendByDeviceResponseBodyData struct {
	// example:
	//
	// ula4wbu166876119986400
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByDeviceResponseBodyData) SetMsgId(v string) *SendByDeviceResponseBodyData {
	s.MsgId = &v
	return s
}

type SendByDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceResponse) GoString() string {
	return s.String()
}

func (s *SendByDeviceResponse) SetHeaders(v map[string]*string) *SendByDeviceResponse {
	s.Headers = v
	return s
}

func (s *SendByDeviceResponse) SetStatusCode(v int32) *SendByDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByDeviceResponse) SetBody(v *SendByDeviceResponseBody) *SendByDeviceResponse {
	s.Body = v
	return s
}

type SendByDeviceFileIdRequest struct {
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PF835431668603208261
	FileId     *string     `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IosPayload *IosPayload `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy     *Policy     `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByDeviceFileIdRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceFileIdRequest) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdRequest) SetAndroidPayload(v *AndroidPayload) *SendByDeviceFileIdRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByDeviceFileIdRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetChannelProperties(v *ChannelProperties) *SendByDeviceFileIdRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetDescription(v string) *SendByDeviceFileIdRequest {
	s.Description = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetFileId(v string) *SendByDeviceFileIdRequest {
	s.FileId = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetIosPayload(v *IosPayload) *SendByDeviceFileIdRequest {
	s.IosPayload = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetPolicy(v *Policy) *SendByDeviceFileIdRequest {
	s.Policy = v
	return s
}

func (s *SendByDeviceFileIdRequest) SetProductionMode(v bool) *SendByDeviceFileIdRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetReceiptType(v int32) *SendByDeviceFileIdRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetReceiptUrl(v string) *SendByDeviceFileIdRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetThirdPartyId(v string) *SendByDeviceFileIdRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByDeviceFileIdRequest) SetCallbackParams(v string) *SendByDeviceFileIdRequest {
	s.CallbackParams = &v
	return s
}

type SendByDeviceFileIdShrinkRequest struct {
	AndroidPayloadShrink      *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayloadShrink *string `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink   *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PF835431668603208261
	FileId           *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IosPayloadShrink *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink     *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByDeviceFileIdShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceFileIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdShrinkRequest) SetAndroidPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetAndroidShortPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.AndroidShortPayloadShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetChannelPropertiesShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetDescription(v string) *SendByDeviceFileIdShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetFileId(v string) *SendByDeviceFileIdShrinkRequest {
	s.FileId = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetIosPayloadShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetPolicyShrink(v string) *SendByDeviceFileIdShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetProductionMode(v bool) *SendByDeviceFileIdShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetReceiptType(v int32) *SendByDeviceFileIdShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetReceiptUrl(v string) *SendByDeviceFileIdShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetThirdPartyId(v string) *SendByDeviceFileIdShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByDeviceFileIdShrinkRequest) SetCallbackParams(v string) *SendByDeviceFileIdShrinkRequest {
	s.CallbackParams = &v
	return s
}

type SendByDeviceFileIdResponseBody struct {
	// example:
	//
	// 0
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByDeviceFileIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByDeviceFileIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceFileIdResponseBody) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdResponseBody) SetCode(v string) *SendByDeviceFileIdResponseBody {
	s.Code = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetData(v *SendByDeviceFileIdResponseBodyData) *SendByDeviceFileIdResponseBody {
	s.Data = v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetHttpStatusCode(v int32) *SendByDeviceFileIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetMessage(v string) *SendByDeviceFileIdResponseBody {
	s.Message = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetRequestId(v string) *SendByDeviceFileIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetSuccess(v bool) *SendByDeviceFileIdResponseBody {
	s.Success = &v
	return s
}

type SendByDeviceFileIdResponseBodyData struct {
	// example:
	//
	// ufe29y2167046828041801
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByDeviceFileIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceFileIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdResponseBodyData) SetMsgId(v string) *SendByDeviceFileIdResponseBodyData {
	s.MsgId = &v
	return s
}

type SendByDeviceFileIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByDeviceFileIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByDeviceFileIdResponse) String() string {
	return tea.Prettify(s)
}

func (s SendByDeviceFileIdResponse) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdResponse) SetHeaders(v map[string]*string) *SendByDeviceFileIdResponse {
	s.Headers = v
	return s
}

func (s *SendByDeviceFileIdResponse) SetStatusCode(v int32) *SendByDeviceFileIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByDeviceFileIdResponse) SetBody(v *SendByDeviceFileIdResponseBody) *SendByDeviceFileIdResponse {
	s.Body = v
	return s
}

type SendByFilterRequest struct {
	AndroidPayload      *AndroidPayload      `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelProperties   *ChannelProperties   `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description         *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// "where":{"and":[{"or":[{"app_version":">=1.0"}]}]}
	Filter     *string     `json:"Filter,omitempty" xml:"Filter,omitempty"`
	IosPayload *IosPayload `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy     *Policy     `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByFilterRequest) GoString() string {
	return s.String()
}

func (s *SendByFilterRequest) SetAndroidPayload(v *AndroidPayload) *SendByFilterRequest {
	s.AndroidPayload = v
	return s
}

func (s *SendByFilterRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByFilterRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByFilterRequest) SetChannelProperties(v *ChannelProperties) *SendByFilterRequest {
	s.ChannelProperties = v
	return s
}

func (s *SendByFilterRequest) SetDescription(v string) *SendByFilterRequest {
	s.Description = &v
	return s
}

func (s *SendByFilterRequest) SetFilter(v string) *SendByFilterRequest {
	s.Filter = &v
	return s
}

func (s *SendByFilterRequest) SetIosPayload(v *IosPayload) *SendByFilterRequest {
	s.IosPayload = v
	return s
}

func (s *SendByFilterRequest) SetPolicy(v *Policy) *SendByFilterRequest {
	s.Policy = v
	return s
}

func (s *SendByFilterRequest) SetProductionMode(v bool) *SendByFilterRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByFilterRequest) SetReceiptType(v int32) *SendByFilterRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByFilterRequest) SetReceiptUrl(v string) *SendByFilterRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByFilterRequest) SetThirdPartyId(v string) *SendByFilterRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByFilterRequest) SetCallbackParams(v string) *SendByFilterRequest {
	s.CallbackParams = &v
	return s
}

type SendByFilterShrinkRequest struct {
	AndroidPayloadShrink    *string              `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	AndroidShortPayload     *AndroidShortPayload `json:"AndroidShortPayload,omitempty" xml:"AndroidShortPayload,omitempty"`
	ChannelPropertiesShrink *string              `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description             *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// "where":{"and":[{"or":[{"app_version":">=1.0"}]}]}
	Filter           *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	IosPayloadShrink *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink     *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// true
	ProductionMode *bool  `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType    *int32 `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	// example:
	//
	// https://msg.umeng.com/upush/receipt
	ReceiptUrl     *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
	ThirdPartyId   *string `json:"ThirdPartyId,omitempty" xml:"ThirdPartyId,omitempty"`
	CallbackParams *string `json:"callbackParams,omitempty" xml:"callbackParams,omitempty"`
}

func (s SendByFilterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendByFilterShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendByFilterShrinkRequest) SetAndroidPayloadShrink(v string) *SendByFilterShrinkRequest {
	s.AndroidPayloadShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetAndroidShortPayload(v *AndroidShortPayload) *SendByFilterShrinkRequest {
	s.AndroidShortPayload = v
	return s
}

func (s *SendByFilterShrinkRequest) SetChannelPropertiesShrink(v string) *SendByFilterShrinkRequest {
	s.ChannelPropertiesShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetDescription(v string) *SendByFilterShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetFilter(v string) *SendByFilterShrinkRequest {
	s.Filter = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetIosPayloadShrink(v string) *SendByFilterShrinkRequest {
	s.IosPayloadShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetPolicyShrink(v string) *SendByFilterShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetProductionMode(v bool) *SendByFilterShrinkRequest {
	s.ProductionMode = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetReceiptType(v int32) *SendByFilterShrinkRequest {
	s.ReceiptType = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetReceiptUrl(v string) *SendByFilterShrinkRequest {
	s.ReceiptUrl = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetThirdPartyId(v string) *SendByFilterShrinkRequest {
	s.ThirdPartyId = &v
	return s
}

func (s *SendByFilterShrinkRequest) SetCallbackParams(v string) *SendByFilterShrinkRequest {
	s.CallbackParams = &v
	return s
}

type SendByFilterResponseBody struct {
	// example:
	//
	// 0
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByFilterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendByFilterResponseBody) GoString() string {
	return s.String()
}

func (s *SendByFilterResponseBody) SetCode(v string) *SendByFilterResponseBody {
	s.Code = &v
	return s
}

func (s *SendByFilterResponseBody) SetData(v *SendByFilterResponseBodyData) *SendByFilterResponseBody {
	s.Data = v
	return s
}

func (s *SendByFilterResponseBody) SetHttpStatusCode(v int32) *SendByFilterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByFilterResponseBody) SetMessage(v string) *SendByFilterResponseBody {
	s.Message = &v
	return s
}

func (s *SendByFilterResponseBody) SetRequestId(v string) *SendByFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByFilterResponseBody) SetSuccess(v bool) *SendByFilterResponseBody {
	s.Success = &v
	return s
}

type SendByFilterResponseBodyData struct {
	// example:
	//
	// usouag1167056659161101
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByFilterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendByFilterResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByFilterResponseBodyData) SetMsgId(v string) *SendByFilterResponseBodyData {
	s.MsgId = &v
	return s
}

type SendByFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendByFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendByFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s SendByFilterResponse) GoString() string {
	return s.String()
}

func (s *SendByFilterResponse) SetHeaders(v map[string]*string) *SendByFilterResponse {
	s.Headers = v
	return s
}

func (s *SendByFilterResponse) SetStatusCode(v int32) *SendByFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *SendByFilterResponse) SetBody(v *SendByFilterResponseBody) *SendByFilterResponse {
	s.Body = v
	return s
}

type UploadDeviceRequest struct {
	// example:
	//
	// device_token_1\\ndevice_token_2\\ndevice_token_3\\n...
	//
	// alias1\\nalias2\\nalias3\\n...
	DeviceTokens *string `json:"DeviceTokens,omitempty" xml:"DeviceTokens,omitempty"`
}

func (s UploadDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceRequest) GoString() string {
	return s.String()
}

func (s *UploadDeviceRequest) SetDeviceTokens(v string) *UploadDeviceRequest {
	s.DeviceTokens = &v
	return s
}

type UploadDeviceResponseBody struct {
	// example:
	//
	// 0
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UploadDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDeviceResponseBody) SetCode(v string) *UploadDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDeviceResponseBody) SetData(v *UploadDeviceResponseBodyData) *UploadDeviceResponseBody {
	s.Data = v
	return s
}

func (s *UploadDeviceResponseBody) SetHttpStatusCode(v int32) *UploadDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadDeviceResponseBody) SetMessage(v string) *UploadDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDeviceResponseBody) SetRequestId(v string) *UploadDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDeviceResponseBody) SetSuccess(v bool) *UploadDeviceResponseBody {
	s.Success = &v
	return s
}

type UploadDeviceResponseBodyData struct {
	// example:
	//
	// PF835431668603208261
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s UploadDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadDeviceResponseBodyData) SetFileId(v string) *UploadDeviceResponseBodyData {
	s.FileId = &v
	return s
}

type UploadDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceResponse) GoString() string {
	return s.String()
}

func (s *UploadDeviceResponse) SetHeaders(v map[string]*string) *UploadDeviceResponse {
	s.Headers = v
	return s
}

func (s *UploadDeviceResponse) SetStatusCode(v int32) *UploadDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDeviceResponse) SetBody(v *UploadDeviceResponseBody) *UploadDeviceResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("umeng-push"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 根据消息ID取消发送
//
// @param request - CancelByMsgIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelByMsgIdResponse
func (client *Client) CancelByMsgIdWithOptions(request *CancelByMsgIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelByMsgIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelByMsgId"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/CancelByMsgId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelByMsgIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据消息ID取消发送
//
// @param request - CancelByMsgIdRequest
//
// @return CancelByMsgIdResponse
func (client *Client) CancelByMsgId(request *CancelByMsgIdRequest) (_result *CancelByMsgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelByMsgIdResponse{}
	_body, _err := client.CancelByMsgIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 消息状态查询
//
// @param request - QueryMsgStatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMsgStatResponse
func (client *Client) QueryMsgStatWithOptions(request *QueryMsgStatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryMsgStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		body["MsgId"] = request.MsgId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMsgStat"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/QueryMsgStat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMsgStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息状态查询
//
// @param request - QueryMsgStatRequest
//
// @return QueryMsgStatResponse
func (client *Client) QueryMsgStat(request *QueryMsgStatRequest) (_result *QueryMsgStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryMsgStatResponse{}
	_body, _err := client.QueryMsgStatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定别名发送
//
// @param tmpReq - SendByAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAliasResponse
func (client *Client) SendByAliasWithOptions(tmpReq *SendByAliasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendByAliasResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendByAliasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidPayload)) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, tea.String("AndroidPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidShortPayload)) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, tea.String("AndroidShortPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ChannelProperties)) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, tea.String("ChannelProperties"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IosPayload)) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, tea.String("IosPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Policy)) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, tea.String("Policy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Alias)) {
		body["Alias"] = request.Alias
	}

	if !tea.BoolValue(util.IsUnset(request.AliasType)) {
		body["AliasType"] = request.AliasType
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidPayloadShrink)) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidShortPayloadShrink)) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelPropertiesShrink)) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IosPayloadShrink)) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyShrink)) {
		body["Policy"] = request.PolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionMode)) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptType)) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptUrl)) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartyId)) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendByAlias"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/SendByAlias"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendByAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定别名发送
//
// @param request - SendByAliasRequest
//
// @return SendByAliasResponse
func (client *Client) SendByAlias(request *SendByAliasRequest) (_result *SendByAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByAliasResponse{}
	_body, _err := client.SendByAliasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定别名文件发送
//
// @param tmpReq - SendByAliasFileIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAliasFileIdResponse
func (client *Client) SendByAliasFileIdWithOptions(tmpReq *SendByAliasFileIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendByAliasFileIdResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendByAliasFileIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidPayload)) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, tea.String("AndroidPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidShortPayload)) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, tea.String("AndroidShortPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ChannelProperties)) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, tea.String("ChannelProperties"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IosPayload)) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, tea.String("IosPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Policy)) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, tea.String("Policy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasType)) {
		body["AliasType"] = request.AliasType
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidPayloadShrink)) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidShortPayloadShrink)) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelPropertiesShrink)) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.IosPayloadShrink)) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyShrink)) {
		body["Policy"] = request.PolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionMode)) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptType)) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptUrl)) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartyId)) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendByAliasFileId"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/SendByAliasFileId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendByAliasFileIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定别名文件发送
//
// @param request - SendByAliasFileIdRequest
//
// @return SendByAliasFileIdResponse
func (client *Client) SendByAliasFileId(request *SendByAliasFileIdRequest) (_result *SendByAliasFileIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByAliasFileIdResponse{}
	_body, _err := client.SendByAliasFileIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 广播
//
// @param tmpReq - SendByAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAppResponse
func (client *Client) SendByAppWithOptions(tmpReq *SendByAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendByAppResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendByAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidPayload)) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, tea.String("AndroidPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidShortPayload)) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, tea.String("AndroidShortPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ChannelProperties)) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, tea.String("ChannelProperties"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IosPayload)) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, tea.String("IosPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Policy)) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, tea.String("Policy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidPayloadShrink)) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidShortPayloadShrink)) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelPropertiesShrink)) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IosPayloadShrink)) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyShrink)) {
		body["Policy"] = request.PolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionMode)) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptType)) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptUrl)) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartyId)) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendByApp"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/SendByApp"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 广播
//
// @param request - SendByAppRequest
//
// @return SendByAppResponse
func (client *Client) SendByApp(request *SendByAppRequest) (_result *SendByAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByAppResponse{}
	_body, _err := client.SendByAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定设备发送
//
// @param tmpReq - SendByDeviceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByDeviceResponse
func (client *Client) SendByDeviceWithOptions(tmpReq *SendByDeviceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendByDeviceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendByDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidPayload)) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, tea.String("AndroidPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidShortPayload)) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, tea.String("AndroidShortPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ChannelProperties)) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, tea.String("ChannelProperties"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IosPayload)) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, tea.String("IosPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Policy)) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, tea.String("Policy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidPayloadShrink)) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidShortPayloadShrink)) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelPropertiesShrink)) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceTokens)) {
		body["DeviceTokens"] = request.DeviceTokens
	}

	if !tea.BoolValue(util.IsUnset(request.IosPayloadShrink)) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyShrink)) {
		body["Policy"] = request.PolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionMode)) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptType)) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptUrl)) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartyId)) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendByDevice"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/SendByDevice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定设备发送
//
// @param request - SendByDeviceRequest
//
// @return SendByDeviceResponse
func (client *Client) SendByDevice(request *SendByDeviceRequest) (_result *SendByDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByDeviceResponse{}
	_body, _err := client.SendByDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定设备文件发送
//
// @param tmpReq - SendByDeviceFileIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByDeviceFileIdResponse
func (client *Client) SendByDeviceFileIdWithOptions(tmpReq *SendByDeviceFileIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendByDeviceFileIdResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendByDeviceFileIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidPayload)) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, tea.String("AndroidPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidShortPayload)) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, tea.String("AndroidShortPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ChannelProperties)) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, tea.String("ChannelProperties"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IosPayload)) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, tea.String("IosPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Policy)) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, tea.String("Policy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidPayloadShrink)) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidShortPayloadShrink)) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelPropertiesShrink)) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.IosPayloadShrink)) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyShrink)) {
		body["Policy"] = request.PolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionMode)) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptType)) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptUrl)) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartyId)) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendByDeviceFileId"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/SendByDeviceFileId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendByDeviceFileIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定设备文件发送
//
// @param request - SendByDeviceFileIdRequest
//
// @return SendByDeviceFileIdResponse
func (client *Client) SendByDeviceFileId(request *SendByDeviceFileIdRequest) (_result *SendByDeviceFileIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByDeviceFileIdResponse{}
	_body, _err := client.SendByDeviceFileIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据筛选条件发送
//
// @param tmpReq - SendByFilterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByFilterResponse
func (client *Client) SendByFilterWithOptions(tmpReq *SendByFilterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendByFilterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendByFilterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AndroidPayload)) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, tea.String("AndroidPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ChannelProperties)) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, tea.String("ChannelProperties"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IosPayload)) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, tea.String("IosPayload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Policy)) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, tea.String("Policy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidPayloadShrink)) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidShortPayload)) {
		body["AndroidShortPayload"] = request.AndroidShortPayload
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelPropertiesShrink)) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.IosPayloadShrink)) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyShrink)) {
		body["Policy"] = request.PolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionMode)) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptType)) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiptUrl)) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartyId)) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendByFilter"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/SendByFilter"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendByFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据筛选条件发送
//
// @param request - SendByFilterRequest
//
// @return SendByFilterResponse
func (client *Client) SendByFilter(request *SendByFilterRequest) (_result *SendByFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByFilterResponse{}
	_body, _err := client.SendByFilterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传设备列表创建设备文件
//
// @param request - UploadDeviceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDeviceResponse
func (client *Client) UploadDeviceWithOptions(request *UploadDeviceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceTokens)) {
		body["DeviceTokens"] = request.DeviceTokens
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadDevice"),
		Version:     tea.String("2022-02-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/UploadDevice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传设备列表创建设备文件
//
// @param request - UploadDeviceRequest
//
// @return UploadDeviceResponse
func (client *Client) UploadDevice(request *UploadDeviceRequest) (_result *UploadDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadDeviceResponse{}
	_body, _err := client.UploadDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
