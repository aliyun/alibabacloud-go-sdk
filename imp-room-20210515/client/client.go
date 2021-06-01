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

type GetLoginTokenRequest struct {
	// AppId
	AppId         *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *GetLoginTokenRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s GetLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenRequest) SetAppId(v string) *GetLoginTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetLoginTokenRequest) SetRequestParams(v *GetLoginTokenRequestRequestParams) *GetLoginTokenRequest {
	s.RequestParams = v
	return s
}

type GetLoginTokenRequestRequestParams struct {
	// 用户ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// AppKey
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s GetLoginTokenRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenRequestRequestParams) GoString() string {
	return s.String()
}

func (s *GetLoginTokenRequestRequestParams) SetAppUid(v string) *GetLoginTokenRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *GetLoginTokenRequestRequestParams) SetAppKey(v string) *GetLoginTokenRequestRequestParams {
	s.AppKey = &v
	return s
}

func (s *GetLoginTokenRequestRequestParams) SetDeviceId(v string) *GetLoginTokenRequestRequestParams {
	s.DeviceId = &v
	return s
}

type GetLoginTokenShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s GetLoginTokenShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenShrinkRequest) SetAppId(v string) *GetLoginTokenShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetLoginTokenShrinkRequest) SetRequestParamsShrink(v string) *GetLoginTokenShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type GetLoginTokenResponseBody struct {
	// Id of the request
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *GetLoginTokenResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBody) SetRequestId(v string) *GetLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetCode(v string) *GetLoginTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetMessage(v string) *GetLoginTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetResult(v *GetLoginTokenResponseBodyResult) *GetLoginTokenResponseBody {
	s.Result = v
	return s
}

type GetLoginTokenResponseBodyResult struct {
	// 登录Tokon
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 更新Token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// 登录Token过期时间
	AccessTokenExpiredTime *int64 `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
}

func (s GetLoginTokenResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyResult) SetAccessToken(v string) *GetLoginTokenResponseBodyResult {
	s.AccessToken = &v
	return s
}

func (s *GetLoginTokenResponseBodyResult) SetRefreshToken(v string) *GetLoginTokenResponseBodyResult {
	s.RefreshToken = &v
	return s
}

func (s *GetLoginTokenResponseBodyResult) SetAccessTokenExpiredTime(v int64) *GetLoginTokenResponseBodyResult {
	s.AccessTokenExpiredTime = &v
	return s
}

type GetLoginTokenResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponse) SetHeaders(v map[string]*string) *GetLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLoginTokenResponse) SetBody(v *GetLoginTokenResponseBody) *GetLoginTokenResponse {
	s.Body = v
	return s
}

type CreateRoomRequest struct {
	Request *CreateRoomRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s CreateRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomRequest) SetRequest(v *CreateRoomRequestRequest) *CreateRoomRequest {
	s.Request = v
	return s
}

type CreateRoomRequestRequest struct {
	// 租户名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 业务类型
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// 模板id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房间标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 房间公告
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房主id
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 拓展字段
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateRoomRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomRequestRequest) SetDomain(v string) *CreateRoomRequestRequest {
	s.Domain = &v
	return s
}

func (s *CreateRoomRequestRequest) SetBizType(v string) *CreateRoomRequestRequest {
	s.BizType = &v
	return s
}

func (s *CreateRoomRequestRequest) SetTemplateId(v string) *CreateRoomRequestRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRoomRequestRequest) SetRoomId(v string) *CreateRoomRequestRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomRequestRequest) SetTitle(v string) *CreateRoomRequestRequest {
	s.Title = &v
	return s
}

func (s *CreateRoomRequestRequest) SetNotice(v string) *CreateRoomRequestRequest {
	s.Notice = &v
	return s
}

func (s *CreateRoomRequestRequest) SetOwnerId(v string) *CreateRoomRequestRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRoomRequestRequest) SetExtension(v map[string]*string) *CreateRoomRequestRequest {
	s.Extension = v
	return s
}

type CreateRoomResponseBody struct {
	// Id of the request
	RequestId       *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result          *CreateRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	ResponseSuccess *bool                         `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
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

func (s *CreateRoomResponseBody) SetResponseSuccess(v bool) *CreateRoomResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *CreateRoomResponseBody) SetErrorCode(v string) *CreateRoomResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRoomResponseBody) SetErrorMsg(v string) *CreateRoomResponseBody {
	s.ErrorMsg = &v
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

type DestroyRoomRequest struct {
	Request *DestroyRoomRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s DestroyRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DestroyRoomRequest) GoString() string {
	return s.String()
}

func (s *DestroyRoomRequest) SetRequest(v *DestroyRoomRequestRequest) *DestroyRoomRequest {
	s.Request = v
	return s
}

type DestroyRoomRequestRequest struct {
	// 租户名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s DestroyRoomRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s DestroyRoomRequestRequest) GoString() string {
	return s.String()
}

func (s *DestroyRoomRequestRequest) SetDomain(v string) *DestroyRoomRequestRequest {
	s.Domain = &v
	return s
}

func (s *DestroyRoomRequestRequest) SetRoomId(v string) *DestroyRoomRequestRequest {
	s.RoomId = &v
	return s
}

type DestroyRoomResponseBody struct {
	// Id of the request
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseSuccess *bool   `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s DestroyRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DestroyRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DestroyRoomResponseBody) SetRequestId(v string) *DestroyRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *DestroyRoomResponseBody) SetResponseSuccess(v bool) *DestroyRoomResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *DestroyRoomResponseBody) SetErrorCode(v string) *DestroyRoomResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DestroyRoomResponseBody) SetErrorMsg(v string) *DestroyRoomResponseBody {
	s.ErrorMsg = &v
	return s
}

type DestroyRoomResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DestroyRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DestroyRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DestroyRoomResponse) GoString() string {
	return s.String()
}

func (s *DestroyRoomResponse) SetHeaders(v map[string]*string) *DestroyRoomResponse {
	s.Headers = v
	return s
}

func (s *DestroyRoomResponse) SetBody(v *DestroyRoomResponseBody) *DestroyRoomResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	Request *CreateInstanceRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetRequest(v *CreateInstanceRequestRequest) *CreateInstanceRequest {
	s.Request = v
	return s
}

type CreateInstanceRequestRequest struct {
	// 租户名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 插件ID
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 拓展字段
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateInstanceRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestRequest) SetDomain(v string) *CreateInstanceRequestRequest {
	s.Domain = &v
	return s
}

func (s *CreateInstanceRequestRequest) SetRoomId(v string) *CreateInstanceRequestRequest {
	s.RoomId = &v
	return s
}

func (s *CreateInstanceRequestRequest) SetPluginId(v string) *CreateInstanceRequestRequest {
	s.PluginId = &v
	return s
}

func (s *CreateInstanceRequestRequest) SetExtension(v map[string]*string) *CreateInstanceRequestRequest {
	s.Extension = v
	return s
}

type CreateInstanceResponseBody struct {
	// Id of the request
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result          *CreateInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	ResponseSuccess *bool                             `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetResult(v *CreateInstanceResponseBodyResult) *CreateInstanceResponseBody {
	s.Result = v
	return s
}

func (s *CreateInstanceResponseBody) SetResponseSuccess(v bool) *CreateInstanceResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorCode(v string) *CreateInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorMsg(v string) *CreateInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

type CreateInstanceResponseBodyResult struct {
	// 插件实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 扩展信息
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyResult) SetInstanceId(v string) *CreateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyResult) SetExtension(v map[string]*string) *CreateInstanceResponseBodyResult {
	s.Extension = v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type GetRoomDetailRequest struct {
	Request *GetRoomDetailRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s GetRoomDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoomDetailRequest) GoString() string {
	return s.String()
}

func (s *GetRoomDetailRequest) SetRequest(v *GetRoomDetailRequestRequest) *GetRoomDetailRequest {
	s.Request = v
	return s
}

type GetRoomDetailRequestRequest struct {
	// 租户名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s GetRoomDetailRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoomDetailRequestRequest) GoString() string {
	return s.String()
}

func (s *GetRoomDetailRequestRequest) SetDomain(v string) *GetRoomDetailRequestRequest {
	s.Domain = &v
	return s
}

func (s *GetRoomDetailRequestRequest) SetRoomId(v string) *GetRoomDetailRequestRequest {
	s.RoomId = &v
	return s
}

type GetRoomDetailResponseBody struct {
	// Id of the request
	RequestId       *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result          *GetRoomDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	ResponseSuccess *bool                            `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s GetRoomDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoomDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoomDetailResponseBody) SetRequestId(v string) *GetRoomDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoomDetailResponseBody) SetResult(v *GetRoomDetailResponseBodyResult) *GetRoomDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetRoomDetailResponseBody) SetResponseSuccess(v bool) *GetRoomDetailResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *GetRoomDetailResponseBody) SetErrorCode(v string) *GetRoomDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRoomDetailResponseBody) SetErrorMsg(v string) *GetRoomDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

type GetRoomDetailResponseBodyResult struct {
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房间标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 房间公告
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房主id
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// uv
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
	// 在线数
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 活跃插件列表
	PluginInstanceInfoList []*GetRoomDetailResponseBodyResultPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
}

func (s GetRoomDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRoomDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRoomDetailResponseBodyResult) SetRoomId(v string) *GetRoomDetailResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetRoomDetailResponseBodyResult) SetTitle(v string) *GetRoomDetailResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetRoomDetailResponseBodyResult) SetNotice(v string) *GetRoomDetailResponseBodyResult {
	s.Notice = &v
	return s
}

func (s *GetRoomDetailResponseBodyResult) SetOwnerId(v string) *GetRoomDetailResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *GetRoomDetailResponseBodyResult) SetUv(v int64) *GetRoomDetailResponseBodyResult {
	s.Uv = &v
	return s
}

func (s *GetRoomDetailResponseBodyResult) SetOnlineCount(v int64) *GetRoomDetailResponseBodyResult {
	s.OnlineCount = &v
	return s
}

func (s *GetRoomDetailResponseBodyResult) SetPluginInstanceInfoList(v []*GetRoomDetailResponseBodyResultPluginInstanceInfoList) *GetRoomDetailResponseBodyResult {
	s.PluginInstanceInfoList = v
	return s
}

type GetRoomDetailResponseBodyResultPluginInstanceInfoList struct {
	// 插件id
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 对应实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 创建时间戳
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 拓展字段
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s GetRoomDetailResponseBodyResultPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetRoomDetailResponseBodyResultPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *GetRoomDetailResponseBodyResultPluginInstanceInfoList) SetPluginId(v string) *GetRoomDetailResponseBodyResultPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *GetRoomDetailResponseBodyResultPluginInstanceInfoList) SetInstanceId(v string) *GetRoomDetailResponseBodyResultPluginInstanceInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetRoomDetailResponseBodyResultPluginInstanceInfoList) SetCreateTime(v int64) *GetRoomDetailResponseBodyResultPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetRoomDetailResponseBodyResultPluginInstanceInfoList) SetExtension(v map[string]*string) *GetRoomDetailResponseBodyResultPluginInstanceInfoList {
	s.Extension = v
	return s
}

type GetRoomDetailResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRoomDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoomDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoomDetailResponse) GoString() string {
	return s.String()
}

func (s *GetRoomDetailResponse) SetHeaders(v map[string]*string) *GetRoomDetailResponse {
	s.Headers = v
	return s
}

func (s *GetRoomDetailResponse) SetBody(v *GetRoomDetailResponseBody) *GetRoomDetailResponse {
	s.Body = v
	return s
}

type GetRoomListRequest struct {
	Request *GetRoomListRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s GetRoomListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoomListRequest) GoString() string {
	return s.String()
}

func (s *GetRoomListRequest) SetRequest(v *GetRoomListRequestRequest) *GetRoomListRequest {
	s.Request = v
	return s
}

type GetRoomListRequestRequest struct {
	// 租户名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 业务类型
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// 查询第几页的房间列表
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 该页面房间数量(最大支持50)
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetRoomListRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoomListRequestRequest) GoString() string {
	return s.String()
}

func (s *GetRoomListRequestRequest) SetDomain(v string) *GetRoomListRequestRequest {
	s.Domain = &v
	return s
}

func (s *GetRoomListRequestRequest) SetBizType(v string) *GetRoomListRequestRequest {
	s.BizType = &v
	return s
}

func (s *GetRoomListRequestRequest) SetPageNum(v int64) *GetRoomListRequestRequest {
	s.PageNum = &v
	return s
}

func (s *GetRoomListRequestRequest) SetPageSize(v int64) *GetRoomListRequestRequest {
	s.PageSize = &v
	return s
}

type GetRoomListResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务结果
	Result *GetRoomListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// 请求是否成功
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s GetRoomListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoomListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoomListResponseBody) SetRequestId(v string) *GetRoomListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoomListResponseBody) SetResult(v *GetRoomListResponseBodyResult) *GetRoomListResponseBody {
	s.Result = v
	return s
}

func (s *GetRoomListResponseBody) SetResponseSuccess(v bool) *GetRoomListResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *GetRoomListResponseBody) SetErrorCode(v string) *GetRoomListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRoomListResponseBody) SetErrorMsg(v string) *GetRoomListResponseBody {
	s.ErrorMsg = &v
	return s
}

type GetRoomListResponseBodyResult struct {
	// 租户下的房间列表总数
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// 是否还有下一页房间列表
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 租户下的房间列表基础信息
	RoomInfoList []*GetRoomListResponseBodyResultRoomInfoList `json:"RoomInfoList,omitempty" xml:"RoomInfoList,omitempty" type:"Repeated"`
}

func (s GetRoomListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRoomListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRoomListResponseBodyResult) SetTotal(v int64) *GetRoomListResponseBodyResult {
	s.Total = &v
	return s
}

func (s *GetRoomListResponseBodyResult) SetHasMore(v bool) *GetRoomListResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetRoomListResponseBodyResult) SetRoomInfoList(v []*GetRoomListResponseBodyResultRoomInfoList) *GetRoomListResponseBodyResult {
	s.RoomInfoList = v
	return s
}

type GetRoomListResponseBodyResultRoomInfoList struct {
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房间标题名字
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 房主的用户id
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 业务类型
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// 应用id，同appId
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetRoomListResponseBodyResultRoomInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetRoomListResponseBodyResultRoomInfoList) GoString() string {
	return s.String()
}

func (s *GetRoomListResponseBodyResultRoomInfoList) SetRoomId(v string) *GetRoomListResponseBodyResultRoomInfoList {
	s.RoomId = &v
	return s
}

func (s *GetRoomListResponseBodyResultRoomInfoList) SetTitle(v string) *GetRoomListResponseBodyResultRoomInfoList {
	s.Title = &v
	return s
}

func (s *GetRoomListResponseBodyResultRoomInfoList) SetOwnerId(v string) *GetRoomListResponseBodyResultRoomInfoList {
	s.OwnerId = &v
	return s
}

func (s *GetRoomListResponseBodyResultRoomInfoList) SetBizType(v string) *GetRoomListResponseBodyResultRoomInfoList {
	s.BizType = &v
	return s
}

func (s *GetRoomListResponseBodyResultRoomInfoList) SetDomain(v string) *GetRoomListResponseBodyResultRoomInfoList {
	s.Domain = &v
	return s
}

type GetRoomListResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRoomListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoomListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoomListResponse) GoString() string {
	return s.String()
}

func (s *GetRoomListResponse) SetHeaders(v map[string]*string) *GetRoomListResponse {
	s.Headers = v
	return s
}

func (s *GetRoomListResponse) SetBody(v *GetRoomListResponseBody) *GetRoomListResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imp-room"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetLoginTokenWithOptions(tmpReq *GetLoginTokenRequest, runtime *util.RuntimeOptions) (_result *GetLoginTokenResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetLoginTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLoginToken"), tea.String("2021-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoginToken(request *GetLoginTokenRequest) (_result *GetLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.GetLoginTokenWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateRoom"), tea.String("2021-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DestroyRoomWithOptions(request *DestroyRoomRequest, runtime *util.RuntimeOptions) (_result *DestroyRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DestroyRoomResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DestroyRoom"), tea.String("2021-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DestroyRoom(request *DestroyRoomRequest) (_result *DestroyRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DestroyRoomResponse{}
	_body, _err := client.DestroyRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2021-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoomDetailWithOptions(request *GetRoomDetailRequest, runtime *util.RuntimeOptions) (_result *GetRoomDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRoomDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRoomDetail"), tea.String("2021-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoomDetail(request *GetRoomDetailRequest) (_result *GetRoomDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRoomDetailResponse{}
	_body, _err := client.GetRoomDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoomListWithOptions(request *GetRoomListRequest, runtime *util.RuntimeOptions) (_result *GetRoomListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRoomListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRoomList"), tea.String("2021-05-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoomList(request *GetRoomListRequest) (_result *GetRoomListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRoomListResponse{}
	_body, _err := client.GetRoomListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
