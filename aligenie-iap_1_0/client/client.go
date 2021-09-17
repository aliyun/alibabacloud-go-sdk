// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type WakeUpAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s WakeUpAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s WakeUpAppHeaders) GoString() string {
	return s.String()
}

func (s *WakeUpAppHeaders) SetCommonHeaders(v map[string]*string) *WakeUpAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WakeUpAppHeaders) SetXAcsAligenieAccessToken(v string) *WakeUpAppHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *WakeUpAppHeaders) SetAuthorization(v string) *WakeUpAppHeaders {
	s.Authorization = &v
	return s
}

type WakeUpAppRequest struct {
	// 应用拉起路径，类似在技能应用控制台中填的唤起链接。
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 猫精应用id，控制台中创建应用后得到的应用id。
	GenieAppId *string `json:"GenieAppId,omitempty" xml:"GenieAppId,omitempty"`
	// 要拉起的目标设备信息。
	TargetInfo *WakeUpAppRequestTargetInfo `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty" type:"Struct"`
	// 是否调试
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
}

func (s WakeUpAppRequest) String() string {
	return tea.Prettify(s)
}

func (s WakeUpAppRequest) GoString() string {
	return s.String()
}

func (s *WakeUpAppRequest) SetPath(v string) *WakeUpAppRequest {
	s.Path = &v
	return s
}

func (s *WakeUpAppRequest) SetGenieAppId(v string) *WakeUpAppRequest {
	s.GenieAppId = &v
	return s
}

func (s *WakeUpAppRequest) SetTargetInfo(v *WakeUpAppRequestTargetInfo) *WakeUpAppRequest {
	s.TargetInfo = v
	return s
}

func (s *WakeUpAppRequest) SetIsDebug(v bool) *WakeUpAppRequest {
	s.IsDebug = &v
	return s
}

type WakeUpAppRequestTargetInfo struct {
	// 推送目标类型，获取到对应设备标识时的类型  DEVICE_UNION_ID：设备unionId； DEVICE_OPEN_ID：设备openId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// 推送目标类型对应的标识值
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
	// 组织标识，推送类型是XX_UNION_XX时才需要配。当存在多种途径获取猫精设备标识且又需要能互通的情况下需要找平台申请组织，申请通过后由平台分配得到。
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// 编码类型，获取猫精的设备标识的途径有多种，不同途径对应不同的编码类型：  PACKAGE_NAME：apk包名 SKILL_ID：技能id
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// 编码类型对应的值，例如：编码类型是SKILLID，其值就为webhook服务中得到的skillId；编码类似是PACKAGENAME，其值就为对应客户端app的packageName。
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
}

func (s WakeUpAppRequestTargetInfo) String() string {
	return tea.Prettify(s)
}

func (s WakeUpAppRequestTargetInfo) GoString() string {
	return s.String()
}

func (s *WakeUpAppRequestTargetInfo) SetTargetType(v string) *WakeUpAppRequestTargetInfo {
	s.TargetType = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetTargetIdentity(v string) *WakeUpAppRequestTargetInfo {
	s.TargetIdentity = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetOrganizationId(v string) *WakeUpAppRequestTargetInfo {
	s.OrganizationId = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetEncodeType(v string) *WakeUpAppRequestTargetInfo {
	s.EncodeType = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetEncodeKey(v string) *WakeUpAppRequestTargetInfo {
	s.EncodeKey = &v
	return s
}

type WakeUpAppResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s WakeUpAppResponse) String() string {
	return tea.Prettify(s)
}

func (s WakeUpAppResponse) GoString() string {
	return s.String()
}

func (s *WakeUpAppResponse) SetHeaders(v map[string]*string) *WakeUpAppResponse {
	s.Headers = v
	return s
}

type PushNotificationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushNotificationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsHeaders) GoString() string {
	return s.String()
}

func (s *PushNotificationsHeaders) SetCommonHeaders(v map[string]*string) *PushNotificationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushNotificationsHeaders) SetXAcsAligenieAccessToken(v string) *PushNotificationsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushNotificationsHeaders) SetAuthorization(v string) *PushNotificationsHeaders {
	s.Authorization = &v
	return s
}

type PushNotificationsRequest struct {
	// 消息推送入参对象。
	NotificationUnicastRequest *PushNotificationsRequestNotificationUnicastRequest `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty" type:"Struct"`
	// 身份信息。
	TenantInfo *PushNotificationsRequestTenantInfo `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty" type:"Struct"`
}

func (s PushNotificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequest) SetNotificationUnicastRequest(v *PushNotificationsRequestNotificationUnicastRequest) *PushNotificationsRequest {
	s.NotificationUnicastRequest = v
	return s
}

func (s *PushNotificationsRequest) SetTenantInfo(v *PushNotificationsRequestTenantInfo) *PushNotificationsRequest {
	s.TenantInfo = v
	return s
}

type PushNotificationsRequestNotificationUnicastRequest struct {
	// 消息推送的目标信息。
	SendTarget *PushNotificationsRequestNotificationUnicastRequestSendTarget `json:"SendTarget,omitempty" xml:"SendTarget,omitempty" type:"Struct"`
	// 消息模板，在天猫精灵应用平台中申请消息模板时得到的模板id。
	MessageTemplateId *string `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	// 占位符信息，例如：模板是【你好，{nick}！】这里可以是：{"nick":"小甜甜"}
	PlaceHolder map[string]interface{} `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	// 编码类型，获取猫精的设备标识的途径有多种，不同途径对应不同的编码类型： PACKAGE_NAME：apk包名 SKILL_ID：技能id
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// 编码类型对应的值，例如：编码类型是SKILLID，其值就为webhook服务中得到的skillId；编码类似是PACKAGENAME，其值就为对应客户端app的packageName。
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// 组织标识，推送类型是XX_UNION_XX时才需要配。当存在多种途径获取猫精设备或用户标识且又需要能互通的情况下需要找平台申请组织，申请通过后由平台分配得到。
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// 调试标识
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
}

func (s PushNotificationsRequestNotificationUnicastRequest) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsRequestNotificationUnicastRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetSendTarget(v *PushNotificationsRequestNotificationUnicastRequestSendTarget) *PushNotificationsRequestNotificationUnicastRequest {
	s.SendTarget = v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetMessageTemplateId(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.MessageTemplateId = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetPlaceHolder(v map[string]interface{}) *PushNotificationsRequestNotificationUnicastRequest {
	s.PlaceHolder = v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetEncodeType(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.EncodeType = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetEncodeKey(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.EncodeKey = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetOrganizationId(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.OrganizationId = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetIsDebug(v bool) *PushNotificationsRequestNotificationUnicastRequest {
	s.IsDebug = &v
	return s
}

type PushNotificationsRequestNotificationUnicastRequestSendTarget struct {
	// 推送的目标类型，获取到对应设备或用户标识时的类型 - DEVICE_UNION_ID：设备unionId - DEVICE_OPEN_ID：设备openId - USER_UNION_ID：用户unionId - USER_OPEN_ID：用户openId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// 推送目标类型对应的标识值。
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
}

func (s PushNotificationsRequestNotificationUnicastRequestSendTarget) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsRequestNotificationUnicastRequestSendTarget) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) SetTargetType(v string) *PushNotificationsRequestNotificationUnicastRequestSendTarget {
	s.TargetType = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) SetTargetIdentity(v string) *PushNotificationsRequestNotificationUnicastRequestSendTarget {
	s.TargetIdentity = &v
	return s
}

type PushNotificationsRequestTenantInfo struct {
	// 猫精应用id，【开发者平台-技能应用】创建应用后得到的应用id。
	GenieAppId *string `json:"GenieAppId,omitempty" xml:"GenieAppId,omitempty"`
}

func (s PushNotificationsRequestTenantInfo) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsRequestTenantInfo) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequestTenantInfo) SetGenieAppId(v string) *PushNotificationsRequestTenantInfo {
	s.GenieAppId = &v
	return s
}

type PushNotificationsShrinkRequest struct {
	// 消息推送入参对象。
	NotificationUnicastRequestShrink *string `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty"`
	// 身份信息。
	TenantInfoShrink *string `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty"`
}

func (s PushNotificationsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsShrinkRequest) SetNotificationUnicastRequestShrink(v string) *PushNotificationsShrinkRequest {
	s.NotificationUnicastRequestShrink = &v
	return s
}

func (s *PushNotificationsShrinkRequest) SetTenantInfoShrink(v string) *PushNotificationsShrinkRequest {
	s.TenantInfoShrink = &v
	return s
}

type PushNotificationsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PushNotificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsResponse) GoString() string {
	return s.String()
}

func (s *PushNotificationsResponse) SetHeaders(v map[string]*string) *PushNotificationsResponse {
	s.Headers = v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aligenie"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) WakeUpApp(request *WakeUpAppRequest) (_result *WakeUpAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WakeUpAppHeaders{}
	_result = &WakeUpAppResponse{}
	_body, _err := client.WakeUpAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WakeUpAppWithOptions(request *WakeUpAppRequest, headers *WakeUpAppHeaders, runtime *util.RuntimeOptions) (_result *WakeUpAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.GenieAppId)) {
		body["GenieAppId"] = request.GenieAppId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TargetInfo))) {
		body["TargetInfo"] = request.TargetInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IsDebug)) {
		body["IsDebug"] = request.IsDebug
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = headers.XAcsAligenieAccessToken
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = headers.Authorization
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("WakeUpApp"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/wakeup"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &WakeUpAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushNotifications(request *PushNotificationsRequest) (_result *PushNotificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushNotificationsHeaders{}
	_result = &PushNotificationsResponse{}
	_body, _err := client.PushNotificationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushNotificationsWithOptions(tmpReq *PushNotificationsRequest, headers *PushNotificationsHeaders, runtime *util.RuntimeOptions) (_result *PushNotificationsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PushNotificationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.NotificationUnicastRequest))) {
		request.NotificationUnicastRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.NotificationUnicastRequest), tea.String("NotificationUnicastRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.TenantInfo))) {
		request.TenantInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.TenantInfo), tea.String("TenantInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotificationUnicastRequestShrink)) {
		body["NotificationUnicastRequest"] = request.NotificationUnicastRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantInfoShrink)) {
		body["TenantInfo"] = request.TenantInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = headers.XAcsAligenieAccessToken
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = headers.Authorization
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushNotifications"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/notifications"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &PushNotificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
