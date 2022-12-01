// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
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
	Body        *Body                  `json:"body,omitempty" xml:"body,omitempty"`
	DisplayType *string                `json:"displayType,omitempty" xml:"displayType,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
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

type Aps struct {
	Alert             *Alert  `json:"alert,omitempty" xml:"alert,omitempty"`
	Badge             *int32  `json:"badge,omitempty" xml:"badge,omitempty"`
	Category          *string `json:"category,omitempty" xml:"category,omitempty"`
	ContentAvailable  *int32  `json:"contentAvailable,omitempty" xml:"contentAvailable,omitempty"`
	InterruptionLevel *string `json:"interruptionLevel,omitempty" xml:"interruptionLevel,omitempty"`
	Sound             *string `json:"sound,omitempty" xml:"sound,omitempty"`
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

func (s *Aps) SetBadge(v int32) *Aps {
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

type Body struct {
	Activity    *string `json:"activity,omitempty" xml:"activity,omitempty"`
	AfterOpen   *string `json:"afterOpen,omitempty" xml:"afterOpen,omitempty"`
	Badge       *int32  `json:"badge,omitempty" xml:"badge,omitempty"`
	BuilderId   *int64  `json:"builderId,omitempty" xml:"builderId,omitempty"`
	Custom      *string `json:"custom,omitempty" xml:"custom,omitempty"`
	ExpandImage *string `json:"expandImage,omitempty" xml:"expandImage,omitempty"`
	Icon        *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Img         *string `json:"img,omitempty" xml:"img,omitempty"`
	LargeIcon   *string `json:"largeIcon,omitempty" xml:"largeIcon,omitempty"`
	PlayLights  *bool   `json:"playLights,omitempty" xml:"playLights,omitempty"`
	PlaySound   *bool   `json:"playSound,omitempty" xml:"playSound,omitempty"`
	PlayVibrate *bool   `json:"playVibrate,omitempty" xml:"playVibrate,omitempty"`
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

func (s *Body) SetAfterOpen(v string) *Body {
	s.AfterOpen = &v
	return s
}

func (s *Body) SetBadge(v int32) *Body {
	s.Badge = &v
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

func (s *Body) SetLargeIcon(v string) *Body {
	s.LargeIcon = &v
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
	ChannelActivity    *string `json:"channelActivity,omitempty" xml:"channelActivity,omitempty"`
	MainActivity       *string `json:"mainActivity,omitempty" xml:"mainActivity,omitempty"`
	OppoChannelId      *string `json:"oppoChannelId,omitempty" xml:"oppoChannelId,omitempty"`
	VivoClassification *string `json:"vivoClassification,omitempty" xml:"vivoClassification,omitempty"`
	XiaomiChannelId    *string `json:"xiaomiChannelId,omitempty" xml:"xiaomiChannelId,omitempty"`
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

func (s *ChannelProperties) SetMainActivity(v string) *ChannelProperties {
	s.MainActivity = &v
	return s
}

func (s *ChannelProperties) SetOppoChannelId(v string) *ChannelProperties {
	s.OppoChannelId = &v
	return s
}

func (s *ChannelProperties) SetVivoClassification(v string) *ChannelProperties {
	s.VivoClassification = &v
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

type Policy struct {
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	OuterBizNo *string `json:"outerBizNo,omitempty" xml:"outerBizNo,omitempty"`
	Speed      *int32  `json:"speed,omitempty" xml:"speed,omitempty"`
	StartTime  *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
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

type SendByAppRequest struct {
	AndroidPayload    *AndroidPayload    `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	ChannelProperties *ChannelProperties `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description       *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	IosPayload        *IosPayload        `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy            *Policy            `json:"Policy,omitempty" xml:"Policy,omitempty"`
	ProductionMode    *bool              `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType       *int32             `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	ReceiptUrl        *string            `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
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

type SendByAppShrinkRequest struct {
	AndroidPayloadShrink    *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	ChannelPropertiesShrink *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IosPayloadShrink        *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink            *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	ProductionMode          *bool   `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType             *int32  `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	ReceiptUrl              *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
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

type SendByAppResponseBody struct {
	Code           *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *SendByAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendByAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AndroidPayload    *AndroidPayload    `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	ChannelProperties *ChannelProperties `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description       *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceTokens      *string            `json:"DeviceTokens,omitempty" xml:"DeviceTokens,omitempty"`
	IosPayload        *IosPayload        `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	Policy            *Policy            `json:"Policy,omitempty" xml:"Policy,omitempty"`
	ProductionMode    *bool              `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType       *int32             `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	ReceiptUrl        *string            `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
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

type SendByDeviceShrinkRequest struct {
	AndroidPayloadShrink    *string `json:"AndroidPayload,omitempty" xml:"AndroidPayload,omitempty"`
	ChannelPropertiesShrink *string `json:"ChannelProperties,omitempty" xml:"ChannelProperties,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceTokens            *string `json:"DeviceTokens,omitempty" xml:"DeviceTokens,omitempty"`
	IosPayloadShrink        *string `json:"IosPayload,omitempty" xml:"IosPayload,omitempty"`
	PolicyShrink            *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	ProductionMode          *bool   `json:"ProductionMode,omitempty" xml:"ProductionMode,omitempty"`
	ReceiptType             *int32  `json:"ReceiptType,omitempty" xml:"ReceiptType,omitempty"`
	ReceiptUrl              *string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty"`
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

type SendByDeviceResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *SendByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendByDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
