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

type CreateLiveRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间ID，最大长度36个字符，传空值，则随机生成一个房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建直播用户。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 直播资源的唯一标识ID，缺省时系统自动生成36位随机uuid字符串。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播标题，支持中英文，最大长度256位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 直播简介，支持中英文，最大长度2048位。
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 主播ID，支持中英文，最大长度128位，缺省时主播为当前创建直播用户。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 直播推流码率，缺省时默认为3。取值：  -1：流畅lld。 1：标清lsd。 2：高清lhd。 3：超清lud。
	CodeLevel *int32 `json:"CodeLevel,omitempty" xml:"CodeLevel,omitempty"`
}

func (s CreateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRequest) SetAppId(v string) *CreateLiveRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveRequest) SetRoomId(v string) *CreateLiveRequest {
	s.RoomId = &v
	return s
}

func (s *CreateLiveRequest) SetUserId(v string) *CreateLiveRequest {
	s.UserId = &v
	return s
}

func (s *CreateLiveRequest) SetLiveId(v string) *CreateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *CreateLiveRequest) SetTitle(v string) *CreateLiveRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRequest) SetIntroduction(v string) *CreateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *CreateLiveRequest) SetAnchorId(v string) *CreateLiveRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRequest) SetCodeLevel(v int32) *CreateLiveRequest {
	s.CodeLevel = &v
	return s
}

type CreateLiveResponseBody struct {
	// 请求ID。
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateLiveResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBody) SetRequestId(v string) *CreateLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveResponseBody) SetResult(v *CreateLiveResponseBodyResult) *CreateLiveResponseBody {
	s.Result = v
	return s
}

type CreateLiveResponseBodyResult struct {
	// 直播的唯一标识ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s CreateLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBodyResult) SetLiveId(v string) *CreateLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

type CreateLiveResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveResponse) SetHeaders(v map[string]*string) *CreateLiveResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveResponse) SetBody(v *CreateLiveResponseBody) *CreateLiveResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	// 应用唯一标识
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetAppId(v string) *DeleteAppRequest {
	s.AppId = &v
	return s
}

type DeleteAppResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type UpdateRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s UpdateRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoomRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoomRequest) SetAppId(v string) *UpdateRoomRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRoomRequest) SetRoomId(v string) *UpdateRoomRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateRoomRequest) SetTitle(v string) *UpdateRoomRequest {
	s.Title = &v
	return s
}

func (s *UpdateRoomRequest) SetNotice(v string) *UpdateRoomRequest {
	s.Notice = &v
	return s
}

func (s *UpdateRoomRequest) SetRoomOwnerId(v string) *UpdateRoomRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *UpdateRoomRequest) SetExtension(v map[string]*string) *UpdateRoomRequest {
	s.Extension = v
	return s
}

type UpdateRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoomResponseBody) SetRequestId(v string) *UpdateRoomResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRoomResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoomResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoomResponse) SetHeaders(v map[string]*string) *UpdateRoomResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoomResponse) SetBody(v *UpdateRoomResponseBody) *UpdateRoomResponse {
	s.Body = v
	return s
}

type GetAppTemplateRequest struct {
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
}

func (s GetAppTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetAppTemplateRequest) SetAppTemplateId(v string) *GetAppTemplateRequest {
	s.AppTemplateId = &v
	return s
}

type GetAppTemplateResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *GetAppTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAppTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponseBody) SetRequestId(v string) *GetAppTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppTemplateResponseBody) SetResult(v *GetAppTemplateResponseBodyResult) *GetAppTemplateResponseBody {
	s.Result = v
	return s
}

type GetAppTemplateResponseBodyResult struct {
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 应用模板创建者
	AppTemplateCreator *string `json:"AppTemplateCreator,omitempty" xml:"AppTemplateCreator,omitempty"`
	// 应用模板使用状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SdkInfo    *string `json:"SdkInfo,omitempty" xml:"SdkInfo,omitempty"`
	// 配置列表
	ConfigList []*GetAppTemplateResponseBodyResultConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s GetAppTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAppTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponseBodyResult) SetAppTemplateName(v string) *GetAppTemplateResponseBodyResult {
	s.AppTemplateName = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetAppTemplateCreator(v string) *GetAppTemplateResponseBodyResult {
	s.AppTemplateCreator = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetStatus(v string) *GetAppTemplateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetComponentList(v []*string) *GetAppTemplateResponseBodyResult {
	s.ComponentList = v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetCreateTime(v string) *GetAppTemplateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetSdkInfo(v string) *GetAppTemplateResponseBodyResult {
	s.SdkInfo = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetConfigList(v []*GetAppTemplateResponseBodyResultConfigList) *GetAppTemplateResponseBodyResult {
	s.ConfigList = v
	return s
}

type GetAppTemplateResponseBodyResultConfigList struct {
	// 配置项
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 配置项内容
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAppTemplateResponseBodyResultConfigList) String() string {
	return tea.Prettify(s)
}

func (s GetAppTemplateResponseBodyResultConfigList) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponseBodyResultConfigList) SetKey(v string) *GetAppTemplateResponseBodyResultConfigList {
	s.Key = &v
	return s
}

func (s *GetAppTemplateResponseBodyResultConfigList) SetValue(v string) *GetAppTemplateResponseBodyResultConfigList {
	s.Value = &v
	return s
}

type GetAppTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponse) SetHeaders(v map[string]*string) *GetAppTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetAppTemplateResponse) SetBody(v *GetAppTemplateResponseBody) *GetAppTemplateResponse {
	s.Body = v
	return s
}

type GetRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s GetRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoomRequest) GoString() string {
	return s.String()
}

func (s *GetRoomRequest) SetAppId(v string) *GetRoomRequest {
	s.AppId = &v
	return s
}

func (s *GetRoomRequest) SetRoomId(v string) *GetRoomRequest {
	s.RoomId = &v
	return s
}

type GetRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 查询房间信息返回结果。
	Result *GetRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBody) SetRequestId(v string) *GetRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoomResponseBody) SetResult(v *GetRoomResponseBodyResult) *GetRoomResponseBody {
	s.Result = v
	return s
}

type GetRoomResponseBodyResult struct {
	// 房间信息。
	RoomInfo *GetRoomResponseBodyResultRoomInfo `json:"RoomInfo,omitempty" xml:"RoomInfo,omitempty" type:"Struct"`
}

func (s GetRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBodyResult) SetRoomInfo(v *GetRoomResponseBodyResultRoomInfo) *GetRoomResponseBodyResult {
	s.RoomInfo = v
	return s
}

type GetRoomResponseBodyResultRoomInfo struct {
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房间标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 房间创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 房间公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房主用户id。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建房间使用的模板id。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s GetRoomResponseBodyResultRoomInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBodyResultRoomInfo) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBodyResultRoomInfo) SetRoomId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.RoomId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetTitle(v string) *GetRoomResponseBodyResultRoomInfo {
	s.Title = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetCreateTime(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.CreateTime = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetNotice(v string) *GetRoomResponseBodyResultRoomInfo {
	s.Notice = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetRoomOwnerId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.RoomOwnerId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetUv(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.Uv = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetOnlineCount(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.OnlineCount = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetPluginInstanceInfoList(v []*GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) *GetRoomResponseBodyResultRoomInfo {
	s.PluginInstanceInfoList = v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetAppId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.AppId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetTemplateId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.TemplateId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetExtension(v map[string]*string) *GetRoomResponseBodyResultRoomInfo {
	s.Extension = v
	return s
}

type GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList struct {
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetPluginType(v string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.PluginType = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetPluginId(v string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetCreateTime(v int64) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetExtension(v map[string]*string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.Extension = v
	return s
}

type GetRoomResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponse) GoString() string {
	return s.String()
}

func (s *GetRoomResponse) SetHeaders(v map[string]*string) *GetRoomResponse {
	s.Headers = v
	return s
}

func (s *GetRoomResponse) SetBody(v *GetRoomResponseBody) *GetRoomResponse {
	s.Body = v
	return s
}

type CreateAppTemplateRequest struct {
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
}

func (s CreateAppTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAppTemplateRequest) SetAppTemplateName(v string) *CreateAppTemplateRequest {
	s.AppTemplateName = &v
	return s
}

func (s *CreateAppTemplateRequest) SetComponentList(v []*string) *CreateAppTemplateRequest {
	s.ComponentList = v
	return s
}

type CreateAppTemplateShrinkRequest struct {
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 组件列表
	ComponentListShrink *string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty"`
}

func (s CreateAppTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppTemplateShrinkRequest) SetAppTemplateName(v string) *CreateAppTemplateShrinkRequest {
	s.AppTemplateName = &v
	return s
}

func (s *CreateAppTemplateShrinkRequest) SetComponentListShrink(v string) *CreateAppTemplateShrinkRequest {
	s.ComponentListShrink = &v
	return s
}

type CreateAppTemplateResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *CreateAppTemplateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAppTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppTemplateResponseBody) SetRequestId(v string) *CreateAppTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppTemplateResponseBody) SetResult(v *CreateAppTemplateResponseBodyResult) *CreateAppTemplateResponseBody {
	s.Result = v
	return s
}

type CreateAppTemplateResponseBodyResult struct {
	// 应用模板唯一标示
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
}

func (s CreateAppTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppTemplateResponseBodyResult) SetAppTemplateId(v string) *CreateAppTemplateResponseBodyResult {
	s.AppTemplateId = &v
	return s
}

type CreateAppTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAppTemplateResponse) SetHeaders(v map[string]*string) *CreateAppTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAppTemplateResponse) SetBody(v *CreateAppTemplateResponseBody) *CreateAppTemplateResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// 查询页码，参数为空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetPageNumber(v int32) *ListAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

type ListAppsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果体
	Result *ListAppsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetResult(v *ListAppsResponseBodyResult) *ListAppsResponseBody {
	s.Result = v
	return s
}

type ListAppsResponseBodyResult struct {
	// 总条目数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 总页数
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// App信息列表
	AppInfoList []*ListAppsResponseBodyResultAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
}

func (s ListAppsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyResult) SetTotalCount(v int32) *ListAppsResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetPageTotal(v int32) *ListAppsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetAppInfoList(v []*ListAppsResponseBodyResultAppInfoList) *ListAppsResponseBodyResult {
	s.AppInfoList = v
	return s
}

type ListAppsResponseBodyResultAppInfoList struct {
	// 应用唯一标识符
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 应用Key
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 应用状态
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// 应用创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
}

func (s ListAppsResponseBodyResultAppInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyResultAppInfoList) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppId(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppId = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppName(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppTemplateId(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppTemplateId = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppTemplateName(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppTemplateName = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppKey(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppKey = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppStatus(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppStatus = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetCreateTime(v string) *ListAppsResponseBodyResultAppInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetComponentList(v []*string) *ListAppsResponseBodyResultAppInfoList {
	s.ComponentList = v
	return s
}

type ListAppsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListRoomsRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRoomsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListRoomsRequest) SetAppId(v string) *ListRoomsRequest {
	s.AppId = &v
	return s
}

func (s *ListRoomsRequest) SetPageNumber(v int32) *ListRoomsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoomsRequest) SetPageSize(v int32) *ListRoomsRequest {
	s.PageSize = &v
	return s
}

type ListRoomsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *ListRoomsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListRoomsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBody) SetRequestId(v string) *ListRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoomsResponseBody) SetResult(v *ListRoomsResponseBodyResult) *ListRoomsResponseBody {
	s.Result = v
	return s
}

type ListRoomsResponseBodyResult struct {
	// 该应用的房间总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 该应用的房间总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 是否还有下一页房间列表。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 房间列表信息。
	RoomInfoList []*ListRoomsResponseBodyResultRoomInfoList `json:"RoomInfoList,omitempty" xml:"RoomInfoList,omitempty" type:"Repeated"`
}

func (s ListRoomsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResult) SetTotalCount(v int32) *ListRoomsResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListRoomsResponseBodyResult) SetPageTotal(v int32) *ListRoomsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListRoomsResponseBodyResult) SetHasMore(v bool) *ListRoomsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListRoomsResponseBodyResult) SetRoomInfoList(v []*ListRoomsResponseBodyResultRoomInfoList) *ListRoomsResponseBodyResult {
	s.RoomInfoList = v
	return s
}

type ListRoomsResponseBodyResultRoomInfoList struct {
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房间标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 房主用户id。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 用户访问数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
	// 用户在线数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 房间创建时间戳，单位：毫秒。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建房间使用的模板id。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s ListRoomsResponseBodyResultRoomInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResultRoomInfoList) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetRoomId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.RoomId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetTitle(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Title = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetRoomOwnerId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.RoomOwnerId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetNotice(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Notice = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetUv(v int64) *ListRoomsResponseBodyResultRoomInfoList {
	s.Uv = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetOnlineCount(v int64) *ListRoomsResponseBodyResultRoomInfoList {
	s.OnlineCount = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetPluginInstanceInfoList(v []*ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) *ListRoomsResponseBodyResultRoomInfoList {
	s.PluginInstanceInfoList = v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetCreateTime(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetAppId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.AppId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetTemplateId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.TemplateId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetExtension(v map[string]*string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Extension = v
	return s
}

type ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList struct {
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetPluginType(v string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.PluginType = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetPluginId(v string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetCreateTime(v int64) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetExtension(v map[string]*string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.Extension = v
	return s
}

type ListRoomsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoomsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListRoomsResponse) SetHeaders(v map[string]*string) *ListRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListRoomsResponse) SetBody(v *ListRoomsResponseBody) *ListRoomsResponse {
	s.Body = v
	return s
}

type DeleteAppTemplateRequest struct {
	// 模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
}

func (s DeleteAppTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppTemplateRequest) SetAppTemplateId(v string) *DeleteAppTemplateRequest {
	s.AppTemplateId = &v
	return s
}

type DeleteAppTemplateResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppTemplateResponseBody) SetRequestId(v string) *DeleteAppTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppTemplateResponse) SetHeaders(v map[string]*string) *DeleteAppTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppTemplateResponse) SetBody(v *DeleteAppTemplateResponseBody) *DeleteAppTemplateResponse {
	s.Body = v
	return s
}

type ListAppTemplatesRequest struct {
	// 查询页码，参数为空默认查询第1页。
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，参数为空默认显示个数为10。
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesRequest) SetPageNumber(v string) *ListAppTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppTemplatesRequest) SetPageSize(v string) *ListAppTemplatesRequest {
	s.PageSize = &v
	return s
}

type ListAppTemplatesResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *ListAppTemplatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAppTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBody) SetRequestId(v string) *ListAppTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppTemplatesResponseBody) SetResult(v *ListAppTemplatesResponseBodyResult) *ListAppTemplatesResponseBody {
	s.Result = v
	return s
}

type ListAppTemplatesResponseBodyResult struct {
	// 总条目数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 总页数
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 应用模板信息列表
	AppTemplateInfoList []*ListAppTemplatesResponseBodyResultAppTemplateInfoList `json:"AppTemplateInfoList,omitempty" xml:"AppTemplateInfoList,omitempty" type:"Repeated"`
}

func (s ListAppTemplatesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyResult) SetTotalCount(v int32) *ListAppTemplatesResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResult) SetPageTotal(v int32) *ListAppTemplatesResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResult) SetAppTemplateInfoList(v []*ListAppTemplatesResponseBodyResultAppTemplateInfoList) *ListAppTemplatesResponseBodyResult {
	s.AppTemplateInfoList = v
	return s
}

type ListAppTemplatesResponseBodyResultAppTemplateInfoList struct {
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 应用模板创建者
	AppTemplateCreator *string `json:"AppTemplateCreator,omitempty" xml:"AppTemplateCreator,omitempty"`
	// 应用模板使用状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// SDK信息
	SdkInfo *string `json:"SdkInfo,omitempty" xml:"SdkInfo,omitempty"`
	// 配置列表
	ConfigList []*ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s ListAppTemplatesResponseBodyResultAppTemplateInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBodyResultAppTemplateInfoList) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetAppTemplateId(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.AppTemplateId = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetAppTemplateName(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.AppTemplateName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetAppTemplateCreator(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.AppTemplateCreator = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetStatus(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.Status = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetComponentList(v []*string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.ComponentList = v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetCreateTime(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetSdkInfo(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.SdkInfo = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetConfigList(v []*ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.ConfigList = v
	return s
}

type ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList struct {
	// 配置项
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 配置项内容
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList) SetKey(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList {
	s.Key = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList) SetValue(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList {
	s.Value = &v
	return s
}

type ListAppTemplatesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponse) SetHeaders(v map[string]*string) *ListAppTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListAppTemplatesResponse) SetBody(v *ListAppTemplatesResponseBody) *ListAppTemplatesResponse {
	s.Body = v
	return s
}

type ListComponentsRequest struct {
	// 应用唯一标识
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) SetAppId(v string) *ListComponentsRequest {
	s.AppId = &v
	return s
}

func (s *ListComponentsRequest) SetAppTemplateId(v string) *ListComponentsRequest {
	s.AppTemplateId = &v
	return s
}

type ListComponentsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果体
	Result *ListComponentsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBody) SetRequestId(v string) *ListComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentsResponseBody) SetResult(v *ListComponentsResponseBodyResult) *ListComponentsResponseBody {
	s.Result = v
	return s
}

type ListComponentsResponseBodyResult struct {
	// 组件信息
	ComponentCategory []*ListComponentsResponseBodyResultComponentCategory `json:"ComponentCategory,omitempty" xml:"ComponentCategory,omitempty" type:"Repeated"`
	// 配置信息
	ConfigGroup []*ListComponentsResponseBodyResultConfigGroup `json:"ConfigGroup,omitempty" xml:"ConfigGroup,omitempty" type:"Repeated"`
}

func (s ListComponentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResult) SetComponentCategory(v []*ListComponentsResponseBodyResultComponentCategory) *ListComponentsResponseBodyResult {
	s.ComponentCategory = v
	return s
}

func (s *ListComponentsResponseBodyResult) SetConfigGroup(v []*ListComponentsResponseBodyResultConfigGroup) *ListComponentsResponseBodyResult {
	s.ConfigGroup = v
	return s
}

type ListComponentsResponseBodyResultComponentCategory struct {
	// 组件类别
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 类别下的组件列表
	List []*ListComponentsResponseBodyResultComponentCategoryList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListComponentsResponseBodyResultComponentCategory) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultComponentCategory) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultComponentCategory) SetType(v string) *ListComponentsResponseBodyResultComponentCategory {
	s.Type = &v
	return s
}

func (s *ListComponentsResponseBodyResultComponentCategory) SetList(v []*ListComponentsResponseBodyResultComponentCategoryList) *ListComponentsResponseBodyResultComponentCategory {
	s.List = v
	return s
}

type ListComponentsResponseBodyResultComponentCategoryList struct {
	// 组件类型
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// 组件名称
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// 是否使用
	InUse *string `json:"InUse,omitempty" xml:"InUse,omitempty"`
}

func (s ListComponentsResponseBodyResultComponentCategoryList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultComponentCategoryList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultComponentCategoryList) SetComponentType(v string) *ListComponentsResponseBodyResultComponentCategoryList {
	s.ComponentType = &v
	return s
}

func (s *ListComponentsResponseBodyResultComponentCategoryList) SetComponentName(v string) *ListComponentsResponseBodyResultComponentCategoryList {
	s.ComponentName = &v
	return s
}

func (s *ListComponentsResponseBodyResultComponentCategoryList) SetInUse(v string) *ListComponentsResponseBodyResultComponentCategoryList {
	s.InUse = &v
	return s
}

type ListComponentsResponseBodyResultConfigGroup struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s ListComponentsResponseBodyResultConfigGroup) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultConfigGroup) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultConfigGroup) SetKey(v string) *ListComponentsResponseBodyResultConfigGroup {
	s.Key = &v
	return s
}

func (s *ListComponentsResponseBodyResultConfigGroup) SetValue(v string) *ListComponentsResponseBodyResultConfigGroup {
	s.Value = &v
	return s
}

func (s *ListComponentsResponseBodyResultConfigGroup) SetCategory(v string) *ListComponentsResponseBodyResultConfigGroup {
	s.Category = &v
	return s
}

type ListComponentsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponse) GoString() string {
	return s.String()
}

func (s *ListComponentsResponse) SetHeaders(v map[string]*string) *ListComponentsResponse {
	s.Headers = v
	return s
}

func (s *ListComponentsResponse) SetBody(v *ListComponentsResponseBody) *ListComponentsResponse {
	s.Body = v
	return s
}

type UpdateLiveRequest struct {
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播标题，支持中英文，最大长度256位
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 直播简介，支持中英文，最大长度2048位
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
}

func (s UpdateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRequest) SetLiveId(v string) *UpdateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRequest) SetTitle(v string) *UpdateLiveRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveRequest) SetIntroduction(v string) *UpdateLiveRequest {
	s.Introduction = &v
	return s
}

type UpdateLiveResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponseBody) SetRequestId(v string) *UpdateLiveResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLiveResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponse) SetHeaders(v map[string]*string) *UpdateLiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveResponse) SetBody(v *UpdateLiveResponseBody) *UpdateLiveResponse {
	s.Body = v
	return s
}

type UpdateAppTemplateConfigRequest struct {
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 更新配置
	ConfigList []*UpdateAppTemplateConfigRequestConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s UpdateAppTemplateConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateConfigRequest) SetAppTemplateId(v string) *UpdateAppTemplateConfigRequest {
	s.AppTemplateId = &v
	return s
}

func (s *UpdateAppTemplateConfigRequest) SetConfigList(v []*UpdateAppTemplateConfigRequestConfigList) *UpdateAppTemplateConfigRequest {
	s.ConfigList = v
	return s
}

type UpdateAppTemplateConfigRequestConfigList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateAppTemplateConfigRequestConfigList) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateConfigRequestConfigList) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateConfigRequestConfigList) SetKey(v string) *UpdateAppTemplateConfigRequestConfigList {
	s.Key = &v
	return s
}

func (s *UpdateAppTemplateConfigRequestConfigList) SetValue(v string) *UpdateAppTemplateConfigRequestConfigList {
	s.Value = &v
	return s
}

type UpdateAppTemplateConfigShrinkRequest struct {
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 更新配置
	ConfigListShrink *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
}

func (s UpdateAppTemplateConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateConfigShrinkRequest) SetAppTemplateId(v string) *UpdateAppTemplateConfigShrinkRequest {
	s.AppTemplateId = &v
	return s
}

func (s *UpdateAppTemplateConfigShrinkRequest) SetConfigListShrink(v string) *UpdateAppTemplateConfigShrinkRequest {
	s.ConfigListShrink = &v
	return s
}

type UpdateAppTemplateConfigResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppTemplateConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateConfigResponseBody) SetRequestId(v string) *UpdateAppTemplateConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppTemplateConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppTemplateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppTemplateConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateConfigResponse) SetHeaders(v map[string]*string) *UpdateAppTemplateConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppTemplateConfigResponse) SetBody(v *UpdateAppTemplateConfigResponseBody) *UpdateAppTemplateConfigResponse {
	s.Body = v
	return s
}

type StopLiveRequest struct {
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间ID，最大长度36位
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建直播用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s StopLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLiveRequest) GoString() string {
	return s.String()
}

func (s *StopLiveRequest) SetAppId(v string) *StopLiveRequest {
	s.AppId = &v
	return s
}

func (s *StopLiveRequest) SetRoomId(v string) *StopLiveRequest {
	s.RoomId = &v
	return s
}

func (s *StopLiveRequest) SetUserId(v string) *StopLiveRequest {
	s.UserId = &v
	return s
}

func (s *StopLiveRequest) SetLiveId(v string) *StopLiveRequest {
	s.LiveId = &v
	return s
}

type StopLiveResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLiveResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveResponseBody) SetRequestId(v string) *StopLiveResponseBody {
	s.RequestId = &v
	return s
}

type StopLiveResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLiveResponse) GoString() string {
	return s.String()
}

func (s *StopLiveResponse) SetHeaders(v map[string]*string) *StopLiveResponse {
	s.Headers = v
	return s
}

func (s *StopLiveResponse) SetBody(v *StopLiveResponseBody) *StopLiveResponse {
	s.Body = v
	return s
}

type GetAppRequest struct {
	// 应用唯一标识
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) SetAppId(v string) *GetAppRequest {
	s.AppId = &v
	return s
}

type GetAppResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *GetAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetResult(v *GetAppResponseBodyResult) *GetAppResponseBody {
	s.Result = v
	return s
}

type GetAppResponseBodyResult struct {
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 应用状态
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// 应用Key
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 组件列表。
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResult) SetAppName(v string) *GetAppResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppTemplateId(v string) *GetAppResponseBodyResult {
	s.AppTemplateId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppTemplateName(v string) *GetAppResponseBodyResult {
	s.AppTemplateName = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppStatus(v string) *GetAppResponseBodyResult {
	s.AppStatus = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppKey(v string) *GetAppResponseBodyResult {
	s.AppKey = &v
	return s
}

func (s *GetAppResponseBodyResult) SetCreateTime(v string) *GetAppResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBodyResult) SetComponentList(v []*string) *GetAppResponseBodyResult {
	s.ComponentList = v
	return s
}

type GetAppResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type DeleteLiveRequest struct {
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s DeleteLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRequest) SetLiveId(v string) *DeleteLiveRequest {
	s.LiveId = &v
	return s
}

type DeleteLiveResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponseBody) SetRequestId(v string) *DeleteLiveResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLiveResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponse) SetHeaders(v map[string]*string) *DeleteLiveResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveResponse) SetBody(v *DeleteLiveResponseBody) *DeleteLiveResponse {
	s.Body = v
	return s
}

type GetLiveDomainStatusRequest struct {
	// 应用唯一标识
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播域名列表
	LiveDomainList []*string `json:"LiveDomainList,omitempty" xml:"LiveDomainList,omitempty" type:"Repeated"`
}

func (s GetLiveDomainStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveDomainStatusRequest) GoString() string {
	return s.String()
}

func (s *GetLiveDomainStatusRequest) SetAppId(v string) *GetLiveDomainStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveDomainStatusRequest) SetLiveDomainList(v []*string) *GetLiveDomainStatusRequest {
	s.LiveDomainList = v
	return s
}

type GetLiveDomainStatusShrinkRequest struct {
	// 应用唯一标识
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播域名列表
	LiveDomainListShrink *string `json:"LiveDomainList,omitempty" xml:"LiveDomainList,omitempty"`
}

func (s GetLiveDomainStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveDomainStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLiveDomainStatusShrinkRequest) SetAppId(v string) *GetLiveDomainStatusShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveDomainStatusShrinkRequest) SetLiveDomainListShrink(v string) *GetLiveDomainStatusShrinkRequest {
	s.LiveDomainListShrink = &v
	return s
}

type GetLiveDomainStatusResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *GetLiveDomainStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveDomainStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveDomainStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveDomainStatusResponseBody) SetRequestId(v string) *GetLiveDomainStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveDomainStatusResponseBody) SetResult(v *GetLiveDomainStatusResponseBodyResult) *GetLiveDomainStatusResponseBody {
	s.Result = v
	return s
}

type GetLiveDomainStatusResponseBodyResult struct {
	// 直播域名信息列表
	LiveDomainInfoList []*GetLiveDomainStatusResponseBodyResultLiveDomainInfoList `json:"LiveDomainInfoList,omitempty" xml:"LiveDomainInfoList,omitempty" type:"Repeated"`
}

func (s GetLiveDomainStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveDomainStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveDomainStatusResponseBodyResult) SetLiveDomainInfoList(v []*GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) *GetLiveDomainStatusResponseBodyResult {
	s.LiveDomainInfoList = v
	return s
}

type GetLiveDomainStatusResponseBodyResultLiveDomainInfoList struct {
	// 直播域名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 直播域名CNAME
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// 直播域名状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) GoString() string {
	return s.String()
}

func (s *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) SetDomain(v string) *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList {
	s.Domain = &v
	return s
}

func (s *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) SetCname(v string) *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList {
	s.Cname = &v
	return s
}

func (s *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) SetStatus(v string) *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList {
	s.Status = &v
	return s
}

type GetLiveDomainStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLiveDomainStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveDomainStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveDomainStatusResponse) GoString() string {
	return s.String()
}

func (s *GetLiveDomainStatusResponse) SetHeaders(v map[string]*string) *GetLiveDomainStatusResponse {
	s.Headers = v
	return s
}

func (s *GetLiveDomainStatusResponse) SetBody(v *GetLiveDomainStatusResponseBody) *GetLiveDomainStatusResponse {
	s.Body = v
	return s
}

type GetAuthTokenRequest struct {
	// 用户的应用ID，在控制台创建应用时生成
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 用户UserId,在AppId下单独唯一
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 终端设备类型,通过控制台创建和查询
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 终端设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenRequest) SetAppId(v string) *GetAuthTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetAuthTokenRequest) SetUserId(v string) *GetAuthTokenRequest {
	s.UserId = &v
	return s
}

func (s *GetAuthTokenRequest) SetAppKey(v string) *GetAuthTokenRequest {
	s.AppKey = &v
	return s
}

func (s *GetAuthTokenRequest) SetDeviceId(v string) *GetAuthTokenRequest {
	s.DeviceId = &v
	return s
}

type GetAuthTokenResponseBody struct {
	// Id of the request
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAuthTokenResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBody) SetRequestId(v string) *GetAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetResult(v *GetAuthTokenResponseBodyResult) *GetAuthTokenResponseBody {
	s.Result = v
	return s
}

type GetAuthTokenResponseBodyResult struct {
	// 用于长链接建连的token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 更新Token，若AccessToken过期，则可以使用RefreshToken再次获取新Token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// 登录token过期时间(毫秒)
	AccessTokenExpiredTime *int64 `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
}

func (s GetAuthTokenResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBodyResult) SetAccessToken(v string) *GetAuthTokenResponseBodyResult {
	s.AccessToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetRefreshToken(v string) *GetAuthTokenResponseBodyResult {
	s.RefreshToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetAccessTokenExpiredTime(v int64) *GetAuthTokenResponseBodyResult {
	s.AccessTokenExpiredTime = &v
	return s
}

type GetAuthTokenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponse) SetHeaders(v map[string]*string) *GetAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTokenResponse) SetBody(v *GetAuthTokenResponseBody) *GetAuthTokenResponse {
	s.Body = v
	return s
}

type UpdateAppTemplateRequest struct {
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
}

func (s UpdateAppTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateRequest) SetAppTemplateId(v string) *UpdateAppTemplateRequest {
	s.AppTemplateId = &v
	return s
}

func (s *UpdateAppTemplateRequest) SetAppTemplateName(v string) *UpdateAppTemplateRequest {
	s.AppTemplateName = &v
	return s
}

type UpdateAppTemplateResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateResponseBody) SetRequestId(v string) *UpdateAppTemplateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateResponse) SetHeaders(v map[string]*string) *UpdateAppTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppTemplateResponse) SetBody(v *UpdateAppTemplateResponseBody) *UpdateAppTemplateResponse {
	s.Body = v
	return s
}

type GetImpProductStatusResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 开通状态
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetImpProductStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImpProductStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetImpProductStatusResponseBody) SetRequestId(v string) *GetImpProductStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImpProductStatusResponseBody) SetResult(v bool) *GetImpProductStatusResponseBody {
	s.Result = &v
	return s
}

type GetImpProductStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImpProductStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImpProductStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImpProductStatusResponse) GoString() string {
	return s.String()
}

func (s *GetImpProductStatusResponse) SetHeaders(v map[string]*string) *GetImpProductStatusResponse {
	s.Headers = v
	return s
}

func (s *GetImpProductStatusResponse) SetBody(v *GetImpProductStatusResponseBody) *GetImpProductStatusResponse {
	s.Body = v
	return s
}

type PublishLiveRequest struct {
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 当前用户Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PublishLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRequest) GoString() string {
	return s.String()
}

func (s *PublishLiveRequest) SetLiveId(v string) *PublishLiveRequest {
	s.LiveId = &v
	return s
}

func (s *PublishLiveRequest) SetUserId(v string) *PublishLiveRequest {
	s.UserId = &v
	return s
}

type PublishLiveResponseBody struct {
	// Id of the request
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *PublishLiveResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s PublishLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveResponseBody) GoString() string {
	return s.String()
}

func (s *PublishLiveResponseBody) SetRequestId(v string) *PublishLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishLiveResponseBody) SetResult(v *PublishLiveResponseBodyResult) *PublishLiveResponseBody {
	s.Result = v
	return s
}

type PublishLiveResponseBodyResult struct {
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 直播状态：Created-已创建未开播，Living-直播中，End-直播结束
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 直播推流地址
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 直播拉流地址
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
}

func (s PublishLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PublishLiveResponseBodyResult) SetLiveId(v string) *PublishLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetAnchorId(v string) *PublishLiveResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetStatus(v string) *PublishLiveResponseBodyResult {
	s.Status = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetPushUrl(v string) *PublishLiveResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetLiveUrl(v string) *PublishLiveResponseBodyResult {
	s.LiveUrl = &v
	return s
}

type PublishLiveResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveResponse) GoString() string {
	return s.String()
}

func (s *PublishLiveResponse) SetHeaders(v map[string]*string) *PublishLiveResponse {
	s.Headers = v
	return s
}

func (s *PublishLiveResponse) SetBody(v *PublishLiveResponseBody) *PublishLiveResponse {
	s.Body = v
	return s
}

type GetLiveRequest struct {
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s GetLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRequest) SetLiveId(v string) *GetLiveRequest {
	s.LiveId = &v
	return s
}

type GetLiveResponseBody struct {
	// Id of the request
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetLiveResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveResponseBody) SetRequestId(v string) *GetLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveResponseBody) SetResult(v *GetLiveResponseBodyResult) *GetLiveResponseBody {
	s.Result = v
	return s
}

type GetLiveResponseBodyResult struct {
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 直播回放地址
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// 直播创建时间（毫秒ms）
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 直播结束时间（毫秒ms）
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 直播持续时间（毫秒ms）
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 直播推流地址
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 直播拉流地址
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 直播状态：Created-已创建，未开播，Living-直播中，End-直播结束
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 直播简介
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建直播用户
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 直播推送分辨率 -1:lld 1:lsd 2:lhd 3:lud
	CodeLevel *int32 `json:"CodeLevel,omitempty" xml:"CodeLevel,omitempty"`
	// 多分辨率多协议播放信息
	PlayUrlInfoList []*GetLiveResponseBodyResultPlayUrlInfoList `json:"PlayUrlInfoList,omitempty" xml:"PlayUrlInfoList,omitempty" type:"Repeated"`
}

func (s GetLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveResponseBodyResult) SetAnchorId(v string) *GetLiveResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetLiveId(v string) *GetLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetTitle(v string) *GetLiveResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetPlaybackUrl(v string) *GetLiveResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetCreateTime(v int64) *GetLiveResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetEndTime(v int64) *GetLiveResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetDuration(v int64) *GetLiveResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetPushUrl(v string) *GetLiveResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetLiveUrl(v string) *GetLiveResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetStatus(v string) *GetLiveResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetIntroduction(v string) *GetLiveResponseBodyResult {
	s.Introduction = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetRoomId(v string) *GetLiveResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetAppId(v string) *GetLiveResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetUserId(v string) *GetLiveResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetCodeLevel(v int32) *GetLiveResponseBodyResult {
	s.CodeLevel = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetPlayUrlInfoList(v []*GetLiveResponseBodyResultPlayUrlInfoList) *GetLiveResponseBodyResult {
	s.PlayUrlInfoList = v
	return s
}

type GetLiveResponseBodyResultPlayUrlInfoList struct {
	// 直播拉取分辨率 -1:lld 1:lsd 2:lhd 3:lud
	CodeLevel *int32 `json:"CodeLevel,omitempty" xml:"CodeLevel,omitempty"`
	// flv拉流地址
	FlvUrl *string `json:"FlvUrl,omitempty" xml:"FlvUrl,omitempty"`
	// hls拉流地址
	HlsUrl *string `json:"HlsUrl,omitempty" xml:"HlsUrl,omitempty"`
	// rtmp拉流地址
	RtmpUrl *string `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty"`
}

func (s GetLiveResponseBodyResultPlayUrlInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponseBodyResultPlayUrlInfoList) GoString() string {
	return s.String()
}

func (s *GetLiveResponseBodyResultPlayUrlInfoList) SetCodeLevel(v int32) *GetLiveResponseBodyResultPlayUrlInfoList {
	s.CodeLevel = &v
	return s
}

func (s *GetLiveResponseBodyResultPlayUrlInfoList) SetFlvUrl(v string) *GetLiveResponseBodyResultPlayUrlInfoList {
	s.FlvUrl = &v
	return s
}

func (s *GetLiveResponseBodyResultPlayUrlInfoList) SetHlsUrl(v string) *GetLiveResponseBodyResultPlayUrlInfoList {
	s.HlsUrl = &v
	return s
}

func (s *GetLiveResponseBodyResultPlayUrlInfoList) SetRtmpUrl(v string) *GetLiveResponseBodyResultPlayUrlInfoList {
	s.RtmpUrl = &v
	return s
}

type GetLiveResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponse) GoString() string {
	return s.String()
}

func (s *GetLiveResponse) SetHeaders(v map[string]*string) *GetLiveResponse {
	s.Headers = v
	return s
}

func (s *GetLiveResponse) SetBody(v *GetLiveResponseBody) *GetLiveResponse {
	s.Body = v
	return s
}

type DeleteRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s DeleteRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoomRequest) SetAppId(v string) *DeleteRoomRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRoomRequest) SetRoomId(v string) *DeleteRoomRequest {
	s.RoomId = &v
	return s
}

type DeleteRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoomResponseBody) SetRequestId(v string) *DeleteRoomResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRoomResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoomResponse) SetHeaders(v map[string]*string) *DeleteRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoomResponse) SetBody(v *DeleteRoomResponseBody) *DeleteRoomResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 模板ID
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAppTemplateId(v string) *CreateAppRequest {
	s.AppTemplateId = &v
	return s
}

type CreateAppResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *CreateAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody {
	s.Result = v
	return s
}

type CreateAppResponseBodyResult struct {
	// 应用唯一标示
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResult) SetAppId(v string) *CreateAppResponseBodyResult {
	s.AppId = &v
	return s
}

type CreateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间模板唯一标识，当前支持的取值：default，传空默认为default。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位，传空则随机生成一个房间id。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomRequest) SetAppId(v string) *CreateRoomRequest {
	s.AppId = &v
	return s
}

func (s *CreateRoomRequest) SetTemplateId(v string) *CreateRoomRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRoomRequest) SetRoomId(v string) *CreateRoomRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomRequest) SetTitle(v string) *CreateRoomRequest {
	s.Title = &v
	return s
}

func (s *CreateRoomRequest) SetNotice(v string) *CreateRoomRequest {
	s.Notice = &v
	return s
}

func (s *CreateRoomRequest) SetRoomOwnerId(v string) *CreateRoomRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *CreateRoomRequest) SetExtension(v map[string]*string) *CreateRoomRequest {
	s.Extension = v
	return s
}

type CreateRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *CreateRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoomResponseBody) SetRequestId(v string) *CreateRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoomResponseBody) SetResult(v *CreateRoomResponseBodyResult) *CreateRoomResponseBody {
	s.Result = v
	return s
}

type CreateRoomResponseBodyResult struct {
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s CreateRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRoomResponseBodyResult) SetRoomId(v string) *CreateRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

type CreateRoomResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateRoomResponse) SetHeaders(v map[string]*string) *CreateRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateRoomResponse) SetBody(v *CreateRoomResponseBody) *CreateRoomResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	// 应用唯一标识
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用状态
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetAppId(v string) *UpdateAppRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppRequest) SetAppName(v string) *UpdateAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateAppRequest) SetAppStatus(v string) *UpdateAppRequest {
	s.AppStatus = &v
	return s
}

type UpdateAppResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateLiveWithOptions(request *CreateLiveRequest, runtime *util.RuntimeOptions) (_result *CreateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLive"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLive(request *CreateLiveRequest) (_result *CreateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveResponse{}
	_body, _err := client.CreateLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteApp"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoomWithOptions(request *UpdateRoomRequest, runtime *util.RuntimeOptions) (_result *UpdateRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRoomResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRoom"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRoom(request *UpdateRoomRequest) (_result *UpdateRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRoomResponse{}
	_body, _err := client.UpdateRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppTemplateWithOptions(request *GetAppTemplateRequest, runtime *util.RuntimeOptions) (_result *GetAppTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAppTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAppTemplate"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppTemplate(request *GetAppTemplateRequest) (_result *GetAppTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppTemplateResponse{}
	_body, _err := client.GetAppTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoomWithOptions(request *GetRoomRequest, runtime *util.RuntimeOptions) (_result *GetRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRoomResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRoom"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoom(request *GetRoomRequest) (_result *GetRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRoomResponse{}
	_body, _err := client.GetRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppTemplateWithOptions(tmpReq *CreateAppTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateAppTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ComponentList)) {
		request.ComponentListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComponentList, tea.String("ComponentList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAppTemplate"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppTemplate(request *CreateAppTemplateRequest) (_result *CreateAppTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppTemplateResponse{}
	_body, _err := client.CreateAppTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApps"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRoomsWithOptions(request *ListRoomsRequest, runtime *util.RuntimeOptions) (_result *ListRoomsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRoomsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRooms"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRooms(request *ListRoomsRequest) (_result *ListRoomsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRoomsResponse{}
	_body, _err := client.ListRoomsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppTemplateWithOptions(request *DeleteAppTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteAppTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAppTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAppTemplate"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppTemplate(request *DeleteAppTemplateRequest) (_result *DeleteAppTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppTemplateResponse{}
	_body, _err := client.DeleteAppTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppTemplatesWithOptions(request *ListAppTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListAppTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAppTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAppTemplates"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppTemplates(request *ListAppTemplatesRequest) (_result *ListAppTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppTemplatesResponse{}
	_body, _err := client.ListAppTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListComponentsWithOptions(request *ListComponentsRequest, runtime *util.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListComponents"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComponents(request *ListComponentsRequest) (_result *ListComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListComponentsResponse{}
	_body, _err := client.ListComponentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveWithOptions(request *UpdateLiveRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateLive"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLive(request *UpdateLiveRequest) (_result *UpdateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveResponse{}
	_body, _err := client.UpdateLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppTemplateConfigWithOptions(tmpReq *UpdateAppTemplateConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateAppTemplateConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAppTemplateConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConfigList)) {
		request.ConfigListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigList, tea.String("ConfigList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppTemplateConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppTemplateConfig"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppTemplateConfig(request *UpdateAppTemplateConfigRequest) (_result *UpdateAppTemplateConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppTemplateConfigResponse{}
	_body, _err := client.UpdateAppTemplateConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopLiveWithOptions(request *StopLiveRequest, runtime *util.RuntimeOptions) (_result *StopLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopLive"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLive(request *StopLiveRequest) (_result *StopLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLiveResponse{}
	_body, _err := client.StopLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppWithOptions(request *GetAppRequest, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetApp"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApp(request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveWithOptions(request *DeleteLiveRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLive"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLive(request *DeleteLiveRequest) (_result *DeleteLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveResponse{}
	_body, _err := client.DeleteLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveDomainStatusWithOptions(tmpReq *GetLiveDomainStatusRequest, runtime *util.RuntimeOptions) (_result *GetLiveDomainStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetLiveDomainStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LiveDomainList)) {
		request.LiveDomainListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LiveDomainList, tea.String("LiveDomainList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLiveDomainStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLiveDomainStatus"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveDomainStatus(request *GetLiveDomainStatusRequest) (_result *GetLiveDomainStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveDomainStatusResponse{}
	_body, _err := client.GetLiveDomainStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthTokenWithOptions(request *GetAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAuthToken"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthToken(request *GetAuthTokenRequest) (_result *GetAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.GetAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppTemplateWithOptions(request *UpdateAppTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateAppTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppTemplate"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppTemplate(request *UpdateAppTemplateRequest) (_result *UpdateAppTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppTemplateResponse{}
	_body, _err := client.UpdateAppTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImpProductStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetImpProductStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &GetImpProductStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetImpProductStatus"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImpProductStatus() (_result *GetImpProductStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImpProductStatusResponse{}
	_body, _err := client.GetImpProductStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishLiveWithOptions(request *PublishLiveRequest, runtime *util.RuntimeOptions) (_result *PublishLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishLive"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishLive(request *PublishLiveRequest) (_result *PublishLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishLiveResponse{}
	_body, _err := client.PublishLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveWithOptions(request *GetLiveRequest, runtime *util.RuntimeOptions) (_result *GetLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLive"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLive(request *GetLiveRequest) (_result *GetLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveResponse{}
	_body, _err := client.GetLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRoomWithOptions(request *DeleteRoomRequest, runtime *util.RuntimeOptions) (_result *DeleteRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRoomResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRoom"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRoom(request *DeleteRoomRequest) (_result *DeleteRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRoomResponse{}
	_body, _err := client.DeleteRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApp"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoomWithOptions(request *CreateRoomRequest, runtime *util.RuntimeOptions) (_result *CreateRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRoomResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRoom"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRoom(request *CreateRoomRequest) (_result *CreateRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRoomResponse{}
	_body, _err := client.CreateRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppWithOptions(request *UpdateAppRequest, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateApp"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApp(request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
