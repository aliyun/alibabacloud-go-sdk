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

type CreateIceProjectRequest struct {
	// appId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 回放地址的地址
	UrlRegionId *string `json:"UrlRegionId,omitempty" xml:"UrlRegionId,omitempty"`
	// 工程标题
	ProjectTitle *string `json:"ProjectTitle,omitempty" xml:"ProjectTitle,omitempty"`
	// 封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 唯一ID，比如直播uuid
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s CreateIceProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIceProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateIceProjectRequest) SetAppId(v string) *CreateIceProjectRequest {
	s.AppId = &v
	return s
}

func (s *CreateIceProjectRequest) SetUrlRegionId(v string) *CreateIceProjectRequest {
	s.UrlRegionId = &v
	return s
}

func (s *CreateIceProjectRequest) SetProjectTitle(v string) *CreateIceProjectRequest {
	s.ProjectTitle = &v
	return s
}

func (s *CreateIceProjectRequest) SetCoverURL(v string) *CreateIceProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *CreateIceProjectRequest) SetLiveId(v string) *CreateIceProjectRequest {
	s.LiveId = &v
	return s
}

type CreateIceProjectResponseBody struct {
	// 请求ID
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateIceProjectResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateIceProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIceProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIceProjectResponseBody) SetRequestId(v string) *CreateIceProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIceProjectResponseBody) SetResult(v *CreateIceProjectResponseBodyResult) *CreateIceProjectResponseBody {
	s.Result = v
	return s
}

type CreateIceProjectResponseBodyResult struct {
	// 工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateIceProjectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateIceProjectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateIceProjectResponseBodyResult) SetProjectId(v string) *CreateIceProjectResponseBodyResult {
	s.ProjectId = &v
	return s
}

type CreateIceProjectResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIceProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIceProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIceProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateIceProjectResponse) SetHeaders(v map[string]*string) *CreateIceProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateIceProjectResponse) SetBody(v *CreateIceProjectResponseBody) *CreateIceProjectResponse {
	s.Body = v
	return s
}

type RemoveMemberRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 被邀请用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
	// 邀请者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
}

func (s RemoveMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveMemberRequest) SetConferenceId(v string) *RemoveMemberRequest {
	s.ConferenceId = &v
	return s
}

func (s *RemoveMemberRequest) SetToUserId(v string) *RemoveMemberRequest {
	s.ToUserId = &v
	return s
}

func (s *RemoveMemberRequest) SetFromUserId(v string) *RemoveMemberRequest {
	s.FromUserId = &v
	return s
}

type RemoveMemberResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMemberResponseBody) SetRequestId(v string) *RemoveMemberResponseBody {
	s.RequestId = &v
	return s
}

type RemoveMemberResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveMemberResponse) SetHeaders(v map[string]*string) *RemoveMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveMemberResponse) SetBody(v *RemoveMemberResponseBody) *RemoveMemberResponse {
	s.Body = v
	return s
}

type ListApplyLinkMicUsersRequest struct {
	// 会议唯一标识符
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 查询页码，从第1页开始。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大显示个数为100。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApplyLinkMicUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplyLinkMicUsersRequest) GoString() string {
	return s.String()
}

func (s *ListApplyLinkMicUsersRequest) SetConferenceId(v string) *ListApplyLinkMicUsersRequest {
	s.ConferenceId = &v
	return s
}

func (s *ListApplyLinkMicUsersRequest) SetPageNumber(v int32) *ListApplyLinkMicUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplyLinkMicUsersRequest) SetPageSize(v int32) *ListApplyLinkMicUsersRequest {
	s.PageSize = &v
	return s
}

type ListApplyLinkMicUsersResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *ListApplyLinkMicUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListApplyLinkMicUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplyLinkMicUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplyLinkMicUsersResponseBody) SetRequestId(v string) *ListApplyLinkMicUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplyLinkMicUsersResponseBody) SetResult(v *ListApplyLinkMicUsersResponseBodyResult) *ListApplyLinkMicUsersResponseBody {
	s.Result = v
	return s
}

type ListApplyLinkMicUsersResponseBodyResult struct {
	// 会议申请连麦用户列表。
	ApplyLinkMicUserList []*string `json:"ApplyLinkMicUserList,omitempty" xml:"ApplyLinkMicUserList,omitempty" type:"Repeated"`
	// 是否还有下一页成员列表。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 该会议的申请连麦成员总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 改会议的申请连麦成员总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
}

func (s ListApplyLinkMicUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListApplyLinkMicUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListApplyLinkMicUsersResponseBodyResult) SetApplyLinkMicUserList(v []*string) *ListApplyLinkMicUsersResponseBodyResult {
	s.ApplyLinkMicUserList = v
	return s
}

func (s *ListApplyLinkMicUsersResponseBodyResult) SetHasMore(v bool) *ListApplyLinkMicUsersResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListApplyLinkMicUsersResponseBodyResult) SetTotalCount(v int32) *ListApplyLinkMicUsersResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListApplyLinkMicUsersResponseBodyResult) SetPageTotal(v int32) *ListApplyLinkMicUsersResponseBodyResult {
	s.PageTotal = &v
	return s
}

type ListApplyLinkMicUsersResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplyLinkMicUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplyLinkMicUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplyLinkMicUsersResponse) GoString() string {
	return s.String()
}

func (s *ListApplyLinkMicUsersResponse) SetHeaders(v map[string]*string) *ListApplyLinkMicUsersResponse {
	s.Headers = v
	return s
}

func (s *ListApplyLinkMicUsersResponse) SetBody(v *ListApplyLinkMicUsersResponseBody) *ListApplyLinkMicUsersResponse {
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
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
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

func (s *GetRoomResponseBodyResultRoomInfo) SetPv(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.Pv = &v
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

type BanCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 被禁言的用户在房间内的唯一标识
	BanCommentUser *string `json:"BanCommentUser,omitempty" xml:"BanCommentUser,omitempty"`
	// 禁言时长（秒）
	BanCommentTime *int64 `json:"BanCommentTime,omitempty" xml:"BanCommentTime,omitempty"`
}

func (s BanCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s BanCommentRequest) GoString() string {
	return s.String()
}

func (s *BanCommentRequest) SetAppId(v string) *BanCommentRequest {
	s.AppId = &v
	return s
}

func (s *BanCommentRequest) SetRoomId(v string) *BanCommentRequest {
	s.RoomId = &v
	return s
}

func (s *BanCommentRequest) SetUserId(v string) *BanCommentRequest {
	s.UserId = &v
	return s
}

func (s *BanCommentRequest) SetBanCommentUser(v string) *BanCommentRequest {
	s.BanCommentUser = &v
	return s
}

func (s *BanCommentRequest) SetBanCommentTime(v int64) *BanCommentRequest {
	s.BanCommentTime = &v
	return s
}

type BanCommentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作是否成功
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BanCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BanCommentResponseBody) GoString() string {
	return s.String()
}

func (s *BanCommentResponseBody) SetRequestId(v string) *BanCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *BanCommentResponseBody) SetResult(v bool) *BanCommentResponseBody {
	s.Result = &v
	return s
}

type BanCommentResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BanCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BanCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s BanCommentResponse) GoString() string {
	return s.String()
}

func (s *BanCommentResponse) SetHeaders(v map[string]*string) *BanCommentResponse {
	s.Headers = v
	return s
}

func (s *BanCommentResponse) SetBody(v *BanCommentResponseBody) *BanCommentResponse {
	s.Body = v
	return s
}

type AddMemberRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 被邀请用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
	// 邀请者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
}

func (s AddMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMemberRequest) GoString() string {
	return s.String()
}

func (s *AddMemberRequest) SetConferenceId(v string) *AddMemberRequest {
	s.ConferenceId = &v
	return s
}

func (s *AddMemberRequest) SetToUserId(v string) *AddMemberRequest {
	s.ToUserId = &v
	return s
}

func (s *AddMemberRequest) SetFromUserId(v string) *AddMemberRequest {
	s.FromUserId = &v
	return s
}

type AddMemberResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberResponseBody) SetRequestId(v string) *AddMemberResponseBody {
	s.RequestId = &v
	return s
}

type AddMemberResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMemberResponse) GoString() string {
	return s.String()
}

func (s *AddMemberResponse) SetHeaders(v map[string]*string) *AddMemberResponse {
	s.Headers = v
	return s
}

func (s *AddMemberResponse) SetBody(v *AddMemberResponseBody) *AddMemberResponse {
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

type RegisterIceOssMediaRequest struct {
	// 工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// appId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 回放cdn域名
	PlaybackUrlDomain *string `json:"PlaybackUrlDomain,omitempty" xml:"PlaybackUrlDomain,omitempty"`
	// oss bucket
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// oss域名
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// 回放地址的区域ID
	UrlRegionId *string `json:"UrlRegionId,omitempty" xml:"UrlRegionId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	MediaUrl *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	// 来源的服务
	FromType *string `json:"FromType,omitempty" xml:"FromType,omitempty"`
	// 媒资标题
	MediaTitle *string `json:"MediaTitle,omitempty" xml:"MediaTitle,omitempty"`
	// 唯一ID，比如直播uuid
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s RegisterIceOssMediaRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterIceOssMediaRequest) GoString() string {
	return s.String()
}

func (s *RegisterIceOssMediaRequest) SetProjectId(v string) *RegisterIceOssMediaRequest {
	s.ProjectId = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetAppId(v string) *RegisterIceOssMediaRequest {
	s.AppId = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetPlaybackUrlDomain(v string) *RegisterIceOssMediaRequest {
	s.PlaybackUrlDomain = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetOssBucket(v string) *RegisterIceOssMediaRequest {
	s.OssBucket = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetOssEndpoint(v string) *RegisterIceOssMediaRequest {
	s.OssEndpoint = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetUrlRegionId(v string) *RegisterIceOssMediaRequest {
	s.UrlRegionId = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetMediaUrl(v string) *RegisterIceOssMediaRequest {
	s.MediaUrl = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetFromType(v string) *RegisterIceOssMediaRequest {
	s.FromType = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetMediaTitle(v string) *RegisterIceOssMediaRequest {
	s.MediaTitle = &v
	return s
}

func (s *RegisterIceOssMediaRequest) SetLiveId(v string) *RegisterIceOssMediaRequest {
	s.LiveId = &v
	return s
}

type RegisterIceOssMediaResponseBody struct {
	// 请求ID
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *RegisterIceOssMediaResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s RegisterIceOssMediaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterIceOssMediaResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterIceOssMediaResponseBody) SetRequestId(v string) *RegisterIceOssMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterIceOssMediaResponseBody) SetResult(v *RegisterIceOssMediaResponseBodyResult) *RegisterIceOssMediaResponseBody {
	s.Result = v
	return s
}

type RegisterIceOssMediaResponseBodyResult struct {
	// 媒体Id
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RegisterIceOssMediaResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RegisterIceOssMediaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RegisterIceOssMediaResponseBodyResult) SetMediaId(v string) *RegisterIceOssMediaResponseBodyResult {
	s.MediaId = &v
	return s
}

type RegisterIceOssMediaResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterIceOssMediaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterIceOssMediaResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterIceOssMediaResponse) GoString() string {
	return s.String()
}

func (s *RegisterIceOssMediaResponse) SetHeaders(v map[string]*string) *RegisterIceOssMediaResponse {
	s.Headers = v
	return s
}

func (s *RegisterIceOssMediaResponse) SetBody(v *RegisterIceOssMediaResponseBody) *RegisterIceOssMediaResponse {
	s.Body = v
	return s
}

type CreateConferenceRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间ID，最大长度36个字符，传空值，则随机生成一个房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建会议用户。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 会议标题，支持中英文，最大长度256位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateConferenceRequest) SetAppId(v string) *CreateConferenceRequest {
	s.AppId = &v
	return s
}

func (s *CreateConferenceRequest) SetRoomId(v string) *CreateConferenceRequest {
	s.RoomId = &v
	return s
}

func (s *CreateConferenceRequest) SetUserId(v string) *CreateConferenceRequest {
	s.UserId = &v
	return s
}

func (s *CreateConferenceRequest) SetTitle(v string) *CreateConferenceRequest {
	s.Title = &v
	return s
}

type CreateConferenceResponseBody struct {
	// 请求ID。
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateConferenceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConferenceResponseBody) SetRequestId(v string) *CreateConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConferenceResponseBody) SetResult(v *CreateConferenceResponseBodyResult) *CreateConferenceResponseBody {
	s.Result = v
	return s
}

type CreateConferenceResponseBodyResult struct {
	// 会议的唯一标识ID。
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s CreateConferenceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateConferenceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateConferenceResponseBodyResult) SetConferenceId(v string) *CreateConferenceResponseBodyResult {
	s.ConferenceId = &v
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
	// 直播域名类型，推流域名: push, 拉流域名: pull, 回放域名: palyback
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty"`
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

func (s *GetLiveDomainStatusRequest) SetLiveDomainType(v string) *GetLiveDomainStatusRequest {
	s.LiveDomainType = &v
	return s
}

type GetLiveDomainStatusShrinkRequest struct {
	// 应用唯一标识
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播域名列表
	LiveDomainListShrink *string `json:"LiveDomainList,omitempty" xml:"LiveDomainList,omitempty"`
	// 直播域名类型，推流域名: push, 拉流域名: pull, 回放域名: palyback
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty"`
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

func (s *GetLiveDomainStatusShrinkRequest) SetLiveDomainType(v string) *GetLiveDomainStatusShrinkRequest {
	s.LiveDomainType = &v
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

type SendCustomMessageToAllRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 消息体内容。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s SendCustomMessageToAllRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToAllRequest) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToAllRequest) SetAppId(v string) *SendCustomMessageToAllRequest {
	s.AppId = &v
	return s
}

func (s *SendCustomMessageToAllRequest) SetRoomId(v string) *SendCustomMessageToAllRequest {
	s.RoomId = &v
	return s
}

func (s *SendCustomMessageToAllRequest) SetBody(v string) *SendCustomMessageToAllRequest {
	s.Body = &v
	return s
}

type SendCustomMessageToAllResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *SendCustomMessageToAllResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendCustomMessageToAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToAllResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToAllResponseBody) SetRequestId(v string) *SendCustomMessageToAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomMessageToAllResponseBody) SetResult(v *SendCustomMessageToAllResponseBodyResult) *SendCustomMessageToAllResponseBody {
	s.Result = v
	return s
}

type SendCustomMessageToAllResponseBodyResult struct {
	// 消息的唯一ID标识。由数字、大小写字母组成，长度不超过20。
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendCustomMessageToAllResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToAllResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToAllResponseBodyResult) SetMessageId(v string) *SendCustomMessageToAllResponseBodyResult {
	s.MessageId = &v
	return s
}

type SendCustomMessageToAllResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCustomMessageToAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCustomMessageToAllResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToAllResponse) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToAllResponse) SetHeaders(v map[string]*string) *SendCustomMessageToAllResponse {
	s.Headers = v
	return s
}

func (s *SendCustomMessageToAllResponse) SetBody(v *SendCustomMessageToAllResponseBody) *SendCustomMessageToAllResponse {
	s.Body = v
	return s
}

type AgreeLinkMicRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 被同意用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
	// 同意者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
}

func (s AgreeLinkMicRequest) String() string {
	return tea.Prettify(s)
}

func (s AgreeLinkMicRequest) GoString() string {
	return s.String()
}

func (s *AgreeLinkMicRequest) SetConferenceId(v string) *AgreeLinkMicRequest {
	s.ConferenceId = &v
	return s
}

func (s *AgreeLinkMicRequest) SetToUserId(v string) *AgreeLinkMicRequest {
	s.ToUserId = &v
	return s
}

func (s *AgreeLinkMicRequest) SetFromUserId(v string) *AgreeLinkMicRequest {
	s.FromUserId = &v
	return s
}

type AgreeLinkMicResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AgreeLinkMicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgreeLinkMicResponseBody) GoString() string {
	return s.String()
}

func (s *AgreeLinkMicResponseBody) SetRequestId(v string) *AgreeLinkMicResponseBody {
	s.RequestId = &v
	return s
}

type AgreeLinkMicResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AgreeLinkMicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AgreeLinkMicResponse) String() string {
	return tea.Prettify(s)
}

func (s AgreeLinkMicResponse) GoString() string {
	return s.String()
}

func (s *AgreeLinkMicResponse) SetHeaders(v map[string]*string) *AgreeLinkMicResponse {
	s.Headers = v
	return s
}

func (s *AgreeLinkMicResponse) SetBody(v *AgreeLinkMicResponseBody) *AgreeLinkMicResponse {
	s.Body = v
	return s
}

type GetDomainOwnerVerifyContentRequest struct {
	// 直播域名
	LiveDomainName *string `json:"LiveDomainName,omitempty" xml:"LiveDomainName,omitempty"`
}

func (s GetDomainOwnerVerifyContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainOwnerVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *GetDomainOwnerVerifyContentRequest) SetLiveDomainName(v string) *GetDomainOwnerVerifyContentRequest {
	s.LiveDomainName = &v
	return s
}

type GetDomainOwnerVerifyContentResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *GetDomainOwnerVerifyContentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDomainOwnerVerifyContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainOwnerVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainOwnerVerifyContentResponseBody) SetRequestId(v string) *GetDomainOwnerVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainOwnerVerifyContentResponseBody) SetResult(v *GetDomainOwnerVerifyContentResponseBodyResult) *GetDomainOwnerVerifyContentResponseBody {
	s.Result = v
	return s
}

type GetDomainOwnerVerifyContentResponseBodyResult struct {
	// 域名归属校验内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetDomainOwnerVerifyContentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetDomainOwnerVerifyContentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDomainOwnerVerifyContentResponseBodyResult) SetContent(v string) *GetDomainOwnerVerifyContentResponseBodyResult {
	s.Content = &v
	return s
}

type GetDomainOwnerVerifyContentResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDomainOwnerVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainOwnerVerifyContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainOwnerVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *GetDomainOwnerVerifyContentResponse) SetHeaders(v map[string]*string) *GetDomainOwnerVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *GetDomainOwnerVerifyContentResponse) SetBody(v *GetDomainOwnerVerifyContentResponseBody) *GetDomainOwnerVerifyContentResponse {
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

type DeleteConferenceRequest struct {
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间ID，最大长度36位
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建会议用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 会议资源的唯一标识ID
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s DeleteConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConferenceRequest) GoString() string {
	return s.String()
}

func (s *DeleteConferenceRequest) SetAppId(v string) *DeleteConferenceRequest {
	s.AppId = &v
	return s
}

func (s *DeleteConferenceRequest) SetRoomId(v string) *DeleteConferenceRequest {
	s.RoomId = &v
	return s
}

func (s *DeleteConferenceRequest) SetUserId(v string) *DeleteConferenceRequest {
	s.UserId = &v
	return s
}

func (s *DeleteConferenceRequest) SetConferenceId(v string) *DeleteConferenceRequest {
	s.ConferenceId = &v
	return s
}

type DeleteConferenceResponseBody struct {
	// 请求ID
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

type VerifyDomainOwnerRequest struct {
	// 直播域名
	LiveDomainName *string `json:"LiveDomainName,omitempty" xml:"LiveDomainName,omitempty"`
}

func (s VerifyDomainOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerRequest) SetLiveDomainName(v string) *VerifyDomainOwnerRequest {
	s.LiveDomainName = &v
	return s
}

type VerifyDomainOwnerResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s VerifyDomainOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerResponseBody) SetRequestId(v string) *VerifyDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyDomainOwnerResponseBody) SetResult(v bool) *VerifyDomainOwnerResponseBody {
	s.Result = &v
	return s
}

type VerifyDomainOwnerResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyDomainOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyDomainOwnerResponse) SetBody(v *VerifyDomainOwnerResponseBody) *VerifyDomainOwnerResponse {
	s.Body = v
	return s
}

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

type GetStandardRoomJumpUrlRequest struct {
	// 用户的应用ID，在控制台创建应用时生成
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 用户UserId,在AppId下单独唯一
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 终端设备类型,通过控制台创建和查询
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 平台：win, mac, android, ios, web
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// 业务类型：互动直播live，互动课堂class
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// 资源ID：根据业务类型来定，比如直播ID，课堂ID等
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// 用户昵称
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s GetStandardRoomJumpUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomJumpUrlRequest) GoString() string {
	return s.String()
}

func (s *GetStandardRoomJumpUrlRequest) SetAppId(v string) *GetStandardRoomJumpUrlRequest {
	s.AppId = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetUserId(v string) *GetStandardRoomJumpUrlRequest {
	s.UserId = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetAppKey(v string) *GetStandardRoomJumpUrlRequest {
	s.AppKey = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetPlatform(v string) *GetStandardRoomJumpUrlRequest {
	s.Platform = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetBizType(v string) *GetStandardRoomJumpUrlRequest {
	s.BizType = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetBizId(v string) *GetStandardRoomJumpUrlRequest {
	s.BizId = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetUserNick(v string) *GetStandardRoomJumpUrlRequest {
	s.UserNick = &v
	return s
}

type GetStandardRoomJumpUrlResponseBody struct {
	// Id of the request
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetStandardRoomJumpUrlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetStandardRoomJumpUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomJumpUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardRoomJumpUrlResponseBody) SetRequestId(v string) *GetStandardRoomJumpUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardRoomJumpUrlResponseBody) SetResult(v *GetStandardRoomJumpUrlResponseBodyResult) *GetStandardRoomJumpUrlResponseBody {
	s.Result = v
	return s
}

type GetStandardRoomJumpUrlResponseBodyResult struct {
	// 样板间跳转协议地址
	StandardRoomJumpUrl *string `json:"StandardRoomJumpUrl,omitempty" xml:"StandardRoomJumpUrl,omitempty"`
}

func (s GetStandardRoomJumpUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomJumpUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetStandardRoomJumpUrlResponseBodyResult) SetStandardRoomJumpUrl(v string) *GetStandardRoomJumpUrlResponseBodyResult {
	s.StandardRoomJumpUrl = &v
	return s
}

type GetStandardRoomJumpUrlResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStandardRoomJumpUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStandardRoomJumpUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomJumpUrlResponse) GoString() string {
	return s.String()
}

func (s *GetStandardRoomJumpUrlResponse) SetHeaders(v map[string]*string) *GetStandardRoomJumpUrlResponse {
	s.Headers = v
	return s
}

func (s *GetStandardRoomJumpUrlResponse) SetBody(v *GetStandardRoomJumpUrlResponseBody) *GetStandardRoomJumpUrlResponse {
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

type ListRoomLivesRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间ID，最大长度36个字符。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播状态筛选条件，0-直播 1-下播，不传则返回全部状态
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 拉取在这个时间戳之前创建的直播，单位毫秒，不传则默认拉取最新创建的。
	QueryTimestamp *int64 `json:"QueryTimestamp,omitempty" xml:"QueryTimestamp,omitempty"`
	// 拉取直播数量。
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// 房间ID列表，可指定多个房间id，过滤优先级高于RoomId。
	RoomIdList []*string `json:"RoomIdList,omitempty" xml:"RoomIdList,omitempty" type:"Repeated"`
}

func (s ListRoomLivesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomLivesRequest) GoString() string {
	return s.String()
}

func (s *ListRoomLivesRequest) SetAppId(v string) *ListRoomLivesRequest {
	s.AppId = &v
	return s
}

func (s *ListRoomLivesRequest) SetRoomId(v string) *ListRoomLivesRequest {
	s.RoomId = &v
	return s
}

func (s *ListRoomLivesRequest) SetStatus(v int32) *ListRoomLivesRequest {
	s.Status = &v
	return s
}

func (s *ListRoomLivesRequest) SetQueryTimestamp(v int64) *ListRoomLivesRequest {
	s.QueryTimestamp = &v
	return s
}

func (s *ListRoomLivesRequest) SetSize(v int32) *ListRoomLivesRequest {
	s.Size = &v
	return s
}

func (s *ListRoomLivesRequest) SetRoomIdList(v []*string) *ListRoomLivesRequest {
	s.RoomIdList = v
	return s
}

type ListRoomLivesShrinkRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间ID，最大长度36个字符。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播状态筛选条件，0-直播 1-下播，不传则返回全部状态
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 拉取在这个时间戳之前创建的直播，单位毫秒，不传则默认拉取最新创建的。
	QueryTimestamp *int64 `json:"QueryTimestamp,omitempty" xml:"QueryTimestamp,omitempty"`
	// 拉取直播数量。
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// 房间ID列表，可指定多个房间id，过滤优先级高于RoomId。
	RoomIdListShrink *string `json:"RoomIdList,omitempty" xml:"RoomIdList,omitempty"`
}

func (s ListRoomLivesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomLivesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRoomLivesShrinkRequest) SetAppId(v string) *ListRoomLivesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListRoomLivesShrinkRequest) SetRoomId(v string) *ListRoomLivesShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *ListRoomLivesShrinkRequest) SetStatus(v int32) *ListRoomLivesShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListRoomLivesShrinkRequest) SetQueryTimestamp(v int64) *ListRoomLivesShrinkRequest {
	s.QueryTimestamp = &v
	return s
}

func (s *ListRoomLivesShrinkRequest) SetSize(v int32) *ListRoomLivesShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListRoomLivesShrinkRequest) SetRoomIdListShrink(v string) *ListRoomLivesShrinkRequest {
	s.RoomIdListShrink = &v
	return s
}

type ListRoomLivesResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *ListRoomLivesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListRoomLivesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoomLivesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoomLivesResponseBody) SetRequestId(v string) *ListRoomLivesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoomLivesResponseBody) SetResult(v *ListRoomLivesResponseBodyResult) *ListRoomLivesResponseBody {
	s.Result = v
	return s
}

type ListRoomLivesResponseBodyResult struct {
	// 是否还有下一页直播列表。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 直播列表信息。
	LiveList []*ListRoomLivesResponseBodyResultLiveList `json:"LiveList,omitempty" xml:"LiveList,omitempty" type:"Repeated"`
	// 下一个拉取的时间戳，单位毫秒。
	NextQueryTimestamp *int64 `json:"NextQueryTimestamp,omitempty" xml:"NextQueryTimestamp,omitempty"`
}

func (s ListRoomLivesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoomLivesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoomLivesResponseBodyResult) SetHasMore(v bool) *ListRoomLivesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListRoomLivesResponseBodyResult) SetLiveList(v []*ListRoomLivesResponseBodyResultLiveList) *ListRoomLivesResponseBodyResult {
	s.LiveList = v
	return s
}

func (s *ListRoomLivesResponseBodyResult) SetNextQueryTimestamp(v int64) *ListRoomLivesResponseBodyResult {
	s.NextQueryTimestamp = &v
	return s
}

type ListRoomLivesResponseBodyResultLiveList struct {
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
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播id。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播状态，0-在播 1-不在播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户访问人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
}

func (s ListRoomLivesResponseBodyResultLiveList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomLivesResponseBodyResultLiveList) GoString() string {
	return s.String()
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetRoomId(v string) *ListRoomLivesResponseBodyResultLiveList {
	s.RoomId = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetTitle(v string) *ListRoomLivesResponseBodyResultLiveList {
	s.Title = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetRoomOwnerId(v string) *ListRoomLivesResponseBodyResultLiveList {
	s.RoomOwnerId = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetNotice(v string) *ListRoomLivesResponseBodyResultLiveList {
	s.Notice = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetUv(v int64) *ListRoomLivesResponseBodyResultLiveList {
	s.Uv = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetAppId(v string) *ListRoomLivesResponseBodyResultLiveList {
	s.AppId = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetExtension(v map[string]*string) *ListRoomLivesResponseBodyResultLiveList {
	s.Extension = v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetLiveId(v string) *ListRoomLivesResponseBodyResultLiveList {
	s.LiveId = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetStatus(v int32) *ListRoomLivesResponseBodyResultLiveList {
	s.Status = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetPv(v int64) *ListRoomLivesResponseBodyResultLiveList {
	s.Pv = &v
	return s
}

func (s *ListRoomLivesResponseBodyResultLiveList) SetOnlineCount(v int64) *ListRoomLivesResponseBodyResultLiveList {
	s.OnlineCount = &v
	return s
}

type ListRoomLivesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRoomLivesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoomLivesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoomLivesResponse) GoString() string {
	return s.String()
}

func (s *ListRoomLivesResponse) SetHeaders(v map[string]*string) *ListRoomLivesResponse {
	s.Headers = v
	return s
}

func (s *ListRoomLivesResponse) SetBody(v *ListRoomLivesResponseBody) *ListRoomLivesResponse {
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

type UpdateRoomShrinkRequest struct {
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
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s UpdateRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoomShrinkRequest) SetAppId(v string) *UpdateRoomShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetRoomId(v string) *UpdateRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetTitle(v string) *UpdateRoomShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetNotice(v string) *UpdateRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetRoomOwnerId(v string) *UpdateRoomShrinkRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetExtensionShrink(v string) *UpdateRoomShrinkRequest {
	s.ExtensionShrink = &v
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
	// 应用模板场景，电商business，课堂classroom
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 样板间信息
	StandardRoomInfo *string `json:"StandardRoomInfo,omitempty" xml:"StandardRoomInfo,omitempty"`
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

func (s *GetAppTemplateResponseBodyResult) SetScene(v string) *GetAppTemplateResponseBodyResult {
	s.Scene = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetIntegrationMode(v string) *GetAppTemplateResponseBodyResult {
	s.IntegrationMode = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetStandardRoomInfo(v string) *GetAppTemplateResponseBodyResult {
	s.StandardRoomInfo = &v
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

type SendCommentRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 发送的文本内容。最大的长度不超过256个字节。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 扩展字段，服务端仅做透传。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s SendCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCommentRequest) GoString() string {
	return s.String()
}

func (s *SendCommentRequest) SetAppId(v string) *SendCommentRequest {
	s.AppId = &v
	return s
}

func (s *SendCommentRequest) SetRoomId(v string) *SendCommentRequest {
	s.RoomId = &v
	return s
}

func (s *SendCommentRequest) SetSenderId(v string) *SendCommentRequest {
	s.SenderId = &v
	return s
}

func (s *SendCommentRequest) SetContent(v string) *SendCommentRequest {
	s.Content = &v
	return s
}

func (s *SendCommentRequest) SetExtension(v map[string]*string) *SendCommentRequest {
	s.Extension = v
	return s
}

type SendCommentShrinkRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 发送的文本内容。最大的长度不超过256个字节。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 扩展字段，服务端仅做透传。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s SendCommentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCommentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendCommentShrinkRequest) SetAppId(v string) *SendCommentShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SendCommentShrinkRequest) SetRoomId(v string) *SendCommentShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *SendCommentShrinkRequest) SetSenderId(v string) *SendCommentShrinkRequest {
	s.SenderId = &v
	return s
}

func (s *SendCommentShrinkRequest) SetContent(v string) *SendCommentShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendCommentShrinkRequest) SetExtensionShrink(v string) *SendCommentShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

type SendCommentResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用发送直播间弹幕的返回结果。
	Result *SendCommentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCommentResponseBody) GoString() string {
	return s.String()
}

func (s *SendCommentResponseBody) SetRequestId(v string) *SendCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCommentResponseBody) SetResult(v *SendCommentResponseBodyResult) *SendCommentResponseBody {
	s.Result = v
	return s
}

type SendCommentResponseBodyResult struct {
	// 返回的弹幕数据模型。
	CommentVO *SendCommentResponseBodyResultCommentVO `json:"CommentVO,omitempty" xml:"CommentVO,omitempty" type:"Struct"`
}

func (s SendCommentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendCommentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendCommentResponseBodyResult) SetCommentVO(v *SendCommentResponseBodyResultCommentVO) *SendCommentResponseBodyResult {
	s.CommentVO = v
	return s
}

type SendCommentResponseBodyResultCommentVO struct {
	// 弹幕的唯一ID。
	CommentId *string `json:"CommentId,omitempty" xml:"CommentId,omitempty"`
	// 弹幕的发送者ID标识。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
	// 弹幕的创建时间，Unix时间戳，单位：毫秒。
	CreateAt *int64 `json:"CreateAt,omitempty" xml:"CreateAt,omitempty"`
	// 弹幕的内容。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 扩展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s SendCommentResponseBodyResultCommentVO) String() string {
	return tea.Prettify(s)
}

func (s SendCommentResponseBodyResultCommentVO) GoString() string {
	return s.String()
}

func (s *SendCommentResponseBodyResultCommentVO) SetCommentId(v string) *SendCommentResponseBodyResultCommentVO {
	s.CommentId = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetSenderId(v string) *SendCommentResponseBodyResultCommentVO {
	s.SenderId = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetSenderNick(v string) *SendCommentResponseBodyResultCommentVO {
	s.SenderNick = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetCreateAt(v int64) *SendCommentResponseBodyResultCommentVO {
	s.CreateAt = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetContent(v string) *SendCommentResponseBodyResultCommentVO {
	s.Content = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetExtension(v map[string]*string) *SendCommentResponseBodyResultCommentVO {
	s.Extension = v
	return s
}

type SendCommentResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCommentResponse) GoString() string {
	return s.String()
}

func (s *SendCommentResponse) SetHeaders(v map[string]*string) *SendCommentResponse {
	s.Headers = v
	return s
}

func (s *SendCommentResponse) SetBody(v *SendCommentResponseBody) *SendCommentResponse {
	s.Body = v
	return s
}

type CreateAppTemplateRequest struct {
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 应用模板场景，电商business，课堂classroom
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// 集成方式（一体化SDK：paasSDK，样板间：standardRoom）
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
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

func (s *CreateAppTemplateRequest) SetScene(v string) *CreateAppTemplateRequest {
	s.Scene = &v
	return s
}

func (s *CreateAppTemplateRequest) SetIntegrationMode(v string) *CreateAppTemplateRequest {
	s.IntegrationMode = &v
	return s
}

func (s *CreateAppTemplateRequest) SetComponentList(v []*string) *CreateAppTemplateRequest {
	s.ComponentList = v
	return s
}

type CreateAppTemplateShrinkRequest struct {
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 应用模板场景，电商business，课堂classroom
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// 集成方式（一体化SDK：paasSDK，样板间：standardRoom）
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
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

func (s *CreateAppTemplateShrinkRequest) SetScene(v string) *CreateAppTemplateShrinkRequest {
	s.Scene = &v
	return s
}

func (s *CreateAppTemplateShrinkRequest) SetIntegrationMode(v string) *CreateAppTemplateShrinkRequest {
	s.IntegrationMode = &v
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

type GetConferenceRequest struct {
	// 会议资源唯一标识。
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s GetConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceRequest) GoString() string {
	return s.String()
}

func (s *GetConferenceRequest) SetConferenceId(v string) *GetConferenceRequest {
	s.ConferenceId = &v
	return s
}

type GetConferenceResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *GetConferenceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetConferenceResponseBody) SetRequestId(v string) *GetConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConferenceResponseBody) SetResult(v *GetConferenceResponseBodyResult) *GetConferenceResponseBody {
	s.Result = v
	return s
}

type GetConferenceResponseBodyResult struct {
	// 会议资源唯一标识。
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 会议标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 会议状态。
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建会议用户。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 会议创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s GetConferenceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetConferenceResponseBodyResult) SetConferenceId(v string) *GetConferenceResponseBodyResult {
	s.ConferenceId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetTitle(v string) *GetConferenceResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetStatus(v string) *GetConferenceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetRoomId(v string) *GetConferenceResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetUserId(v string) *GetConferenceResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetAppId(v string) *GetConferenceResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetCreateTime(v int64) *GetConferenceResponseBodyResult {
	s.CreateTime = &v
	return s
}

type GetConferenceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceResponse) GoString() string {
	return s.String()
}

func (s *GetConferenceResponse) SetHeaders(v map[string]*string) *GetConferenceResponse {
	s.Headers = v
	return s
}

func (s *GetConferenceResponse) SetBody(v *GetConferenceResponseBody) *GetConferenceResponse {
	s.Body = v
	return s
}

type RejectLinkMicRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 被同意用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
	// 同意者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
}

func (s RejectLinkMicRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectLinkMicRequest) GoString() string {
	return s.String()
}

func (s *RejectLinkMicRequest) SetConferenceId(v string) *RejectLinkMicRequest {
	s.ConferenceId = &v
	return s
}

func (s *RejectLinkMicRequest) SetToUserId(v string) *RejectLinkMicRequest {
	s.ToUserId = &v
	return s
}

func (s *RejectLinkMicRequest) SetFromUserId(v string) *RejectLinkMicRequest {
	s.FromUserId = &v
	return s
}

type RejectLinkMicResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectLinkMicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectLinkMicResponseBody) GoString() string {
	return s.String()
}

func (s *RejectLinkMicResponseBody) SetRequestId(v string) *RejectLinkMicResponseBody {
	s.RequestId = &v
	return s
}

type RejectLinkMicResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RejectLinkMicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RejectLinkMicResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectLinkMicResponse) GoString() string {
	return s.String()
}

func (s *RejectLinkMicResponse) SetHeaders(v map[string]*string) *RejectLinkMicResponse {
	s.Headers = v
	return s
}

func (s *RejectLinkMicResponse) SetBody(v *RejectLinkMicResponseBody) *RejectLinkMicResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// 查询页码，参数为空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 应用状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
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

func (s *ListAppsRequest) SetStatus(v string) *ListAppsRequest {
	s.Status = &v
	return s
}

func (s *ListAppsRequest) SetIntegrationMode(v string) *ListAppsRequest {
	s.IntegrationMode = &v
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
	// 应用配置状态
	AppConfigStatus *string `json:"AppConfigStatus,omitempty" xml:"AppConfigStatus,omitempty"`
	// 应用创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 样板间信息
	StandardRoomInfo *string `json:"StandardRoomInfo,omitempty" xml:"StandardRoomInfo,omitempty"`
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

func (s *ListAppsResponseBodyResultAppInfoList) SetAppConfigStatus(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppConfigStatus = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetCreateTime(v string) *ListAppsResponseBodyResultAppInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetIntegrationMode(v string) *ListAppsResponseBodyResultAppInfoList {
	s.IntegrationMode = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetStandardRoomInfo(v string) *ListAppsResponseBodyResultAppInfoList {
	s.StandardRoomInfo = &v
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

type CancelBanAllCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelBanAllCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelBanAllCommentRequest) GoString() string {
	return s.String()
}

func (s *CancelBanAllCommentRequest) SetAppId(v string) *CancelBanAllCommentRequest {
	s.AppId = &v
	return s
}

func (s *CancelBanAllCommentRequest) SetRoomId(v string) *CancelBanAllCommentRequest {
	s.RoomId = &v
	return s
}

func (s *CancelBanAllCommentRequest) SetUserId(v string) *CancelBanAllCommentRequest {
	s.UserId = &v
	return s
}

type CancelBanAllCommentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作成功标识
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelBanAllCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelBanAllCommentResponseBody) GoString() string {
	return s.String()
}

func (s *CancelBanAllCommentResponseBody) SetRequestId(v string) *CancelBanAllCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelBanAllCommentResponseBody) SetResult(v bool) *CancelBanAllCommentResponseBody {
	s.Result = &v
	return s
}

type CancelBanAllCommentResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelBanAllCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelBanAllCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelBanAllCommentResponse) GoString() string {
	return s.String()
}

func (s *CancelBanAllCommentResponse) SetHeaders(v map[string]*string) *CancelBanAllCommentResponse {
	s.Headers = v
	return s
}

func (s *CancelBanAllCommentResponse) SetBody(v *CancelBanAllCommentResponseBody) *CancelBanAllCommentResponse {
	s.Body = v
	return s
}

type ListConferenceUsersRequest struct {
	// 会议唯一标识符
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 查询页码，从第1页开始。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大显示个数为100。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListConferenceUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersRequest) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersRequest) SetConferenceId(v string) *ListConferenceUsersRequest {
	s.ConferenceId = &v
	return s
}

func (s *ListConferenceUsersRequest) SetPageNumber(v int32) *ListConferenceUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConferenceUsersRequest) SetPageSize(v int32) *ListConferenceUsersRequest {
	s.PageSize = &v
	return s
}

type ListConferenceUsersResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *ListConferenceUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListConferenceUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponseBody) SetRequestId(v string) *ListConferenceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConferenceUsersResponseBody) SetResult(v *ListConferenceUsersResponseBodyResult) *ListConferenceUsersResponseBody {
	s.Result = v
	return s
}

type ListConferenceUsersResponseBodyResult struct {
	// 会议用户列表。
	ConferenceUserList []*ListConferenceUsersResponseBodyResultConferenceUserList `json:"ConferenceUserList,omitempty" xml:"ConferenceUserList,omitempty" type:"Repeated"`
	// 是否还有数据
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 总页数
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
}

func (s ListConferenceUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponseBodyResult) SetConferenceUserList(v []*ListConferenceUsersResponseBodyResultConferenceUserList) *ListConferenceUsersResponseBodyResult {
	s.ConferenceUserList = v
	return s
}

func (s *ListConferenceUsersResponseBodyResult) SetHasMore(v bool) *ListConferenceUsersResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListConferenceUsersResponseBodyResult) SetTotalCount(v int32) *ListConferenceUsersResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListConferenceUsersResponseBodyResult) SetPageTotal(v int32) *ListConferenceUsersResponseBodyResult {
	s.PageTotal = &v
	return s
}

type ListConferenceUsersResponseBodyResultConferenceUserList struct {
	// 用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户状态。
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListConferenceUsersResponseBodyResultConferenceUserList) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponseBodyResultConferenceUserList) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponseBodyResultConferenceUserList) SetUserId(v string) *ListConferenceUsersResponseBodyResultConferenceUserList {
	s.UserId = &v
	return s
}

func (s *ListConferenceUsersResponseBodyResultConferenceUserList) SetStatus(v string) *ListConferenceUsersResponseBodyResultConferenceUserList {
	s.Status = &v
	return s
}

type ListConferenceUsersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConferenceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConferenceUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponse) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponse) SetHeaders(v map[string]*string) *ListConferenceUsersResponse {
	s.Headers = v
	return s
}

func (s *ListConferenceUsersResponse) SetBody(v *ListConferenceUsersResponseBody) *ListConferenceUsersResponse {
	s.Body = v
	return s
}

type CancelBanCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 取消禁言的用户唯一标识
	BanCommentUser *string `json:"BanCommentUser,omitempty" xml:"BanCommentUser,omitempty"`
}

func (s CancelBanCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelBanCommentRequest) GoString() string {
	return s.String()
}

func (s *CancelBanCommentRequest) SetAppId(v string) *CancelBanCommentRequest {
	s.AppId = &v
	return s
}

func (s *CancelBanCommentRequest) SetRoomId(v string) *CancelBanCommentRequest {
	s.RoomId = &v
	return s
}

func (s *CancelBanCommentRequest) SetUserId(v string) *CancelBanCommentRequest {
	s.UserId = &v
	return s
}

func (s *CancelBanCommentRequest) SetBanCommentUser(v string) *CancelBanCommentRequest {
	s.BanCommentUser = &v
	return s
}

type CancelBanCommentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作成功标识
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelBanCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelBanCommentResponseBody) GoString() string {
	return s.String()
}

func (s *CancelBanCommentResponseBody) SetRequestId(v string) *CancelBanCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelBanCommentResponseBody) SetResult(v bool) *CancelBanCommentResponseBody {
	s.Result = &v
	return s
}

type CancelBanCommentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelBanCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelBanCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelBanCommentResponse) GoString() string {
	return s.String()
}

func (s *CancelBanCommentResponse) SetHeaders(v map[string]*string) *CancelBanCommentResponse {
	s.Headers = v
	return s
}

func (s *CancelBanCommentResponse) SetBody(v *CancelBanCommentResponseBody) *CancelBanCommentResponse {
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
	// 应用模板场景，电商business，课堂classroom
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 样板间信息
	StandardRoomInfo *string `json:"StandardRoomInfo,omitempty" xml:"StandardRoomInfo,omitempty"`
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

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetScene(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.Scene = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetIntegrationMode(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.IntegrationMode = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetStandardRoomInfo(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.StandardRoomInfo = &v
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
	// 配置信息
	ConfigGroup []*ListComponentsResponseBodyResultConfigGroup `json:"ConfigGroup,omitempty" xml:"ConfigGroup,omitempty" type:"Repeated"`
	// 场景列表
	SceneList []*ListComponentsResponseBodyResultSceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
	// 组件信息
	ComponentCategory []*ListComponentsResponseBodyResultComponentCategory `json:"ComponentCategory,omitempty" xml:"ComponentCategory,omitempty" type:"Repeated"`
}

func (s ListComponentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResult) SetConfigGroup(v []*ListComponentsResponseBodyResultConfigGroup) *ListComponentsResponseBodyResult {
	s.ConfigGroup = v
	return s
}

func (s *ListComponentsResponseBodyResult) SetSceneList(v []*ListComponentsResponseBodyResultSceneList) *ListComponentsResponseBodyResult {
	s.SceneList = v
	return s
}

func (s *ListComponentsResponseBodyResult) SetComponentCategory(v []*ListComponentsResponseBodyResultComponentCategory) *ListComponentsResponseBodyResult {
	s.ComponentCategory = v
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

type ListComponentsResponseBodyResultSceneList struct {
	// 场景类别
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// 组件信息
	ComponentCategory []*ListComponentsResponseBodyResultSceneListComponentCategory `json:"ComponentCategory,omitempty" xml:"ComponentCategory,omitempty" type:"Repeated"`
}

func (s ListComponentsResponseBodyResultSceneList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultSceneList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultSceneList) SetScene(v string) *ListComponentsResponseBodyResultSceneList {
	s.Scene = &v
	return s
}

func (s *ListComponentsResponseBodyResultSceneList) SetComponentCategory(v []*ListComponentsResponseBodyResultSceneListComponentCategory) *ListComponentsResponseBodyResultSceneList {
	s.ComponentCategory = v
	return s
}

type ListComponentsResponseBodyResultSceneListComponentCategory struct {
	// 组件类别
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 类别下的组件列表
	List []*ListComponentsResponseBodyResultSceneListComponentCategoryList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListComponentsResponseBodyResultSceneListComponentCategory) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultSceneListComponentCategory) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategory) SetType(v string) *ListComponentsResponseBodyResultSceneListComponentCategory {
	s.Type = &v
	return s
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategory) SetList(v []*ListComponentsResponseBodyResultSceneListComponentCategoryList) *ListComponentsResponseBodyResultSceneListComponentCategory {
	s.List = v
	return s
}

type ListComponentsResponseBodyResultSceneListComponentCategoryList struct {
	// 组件类型
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// 组件名称
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// 是否使用
	InUse *string `json:"InUse,omitempty" xml:"InUse,omitempty"`
}

func (s ListComponentsResponseBodyResultSceneListComponentCategoryList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultSceneListComponentCategoryList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategoryList) SetComponentType(v string) *ListComponentsResponseBodyResultSceneListComponentCategoryList {
	s.ComponentType = &v
	return s
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategoryList) SetComponentName(v string) *ListComponentsResponseBodyResultSceneListComponentCategoryList {
	s.ComponentName = &v
	return s
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategoryList) SetInUse(v string) *ListComponentsResponseBodyResultSceneListComponentCategoryList {
	s.InUse = &v
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

type ApplyLinkMicRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 申请连麦用户
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ApplyLinkMicRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyLinkMicRequest) GoString() string {
	return s.String()
}

func (s *ApplyLinkMicRequest) SetConferenceId(v string) *ApplyLinkMicRequest {
	s.ConferenceId = &v
	return s
}

func (s *ApplyLinkMicRequest) SetUserId(v string) *ApplyLinkMicRequest {
	s.UserId = &v
	return s
}

type ApplyLinkMicResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyLinkMicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyLinkMicResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyLinkMicResponseBody) SetRequestId(v string) *ApplyLinkMicResponseBody {
	s.RequestId = &v
	return s
}

type ApplyLinkMicResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyLinkMicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyLinkMicResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyLinkMicResponse) GoString() string {
	return s.String()
}

func (s *ApplyLinkMicResponse) SetHeaders(v map[string]*string) *ApplyLinkMicResponse {
	s.Headers = v
	return s
}

func (s *ApplyLinkMicResponse) SetBody(v *ApplyLinkMicResponseBody) *ApplyLinkMicResponse {
	s.Body = v
	return s
}

type CancelApplyLinkMicRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 申请连麦用户
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelApplyLinkMicRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelApplyLinkMicRequest) GoString() string {
	return s.String()
}

func (s *CancelApplyLinkMicRequest) SetConferenceId(v string) *CancelApplyLinkMicRequest {
	s.ConferenceId = &v
	return s
}

func (s *CancelApplyLinkMicRequest) SetUserId(v string) *CancelApplyLinkMicRequest {
	s.UserId = &v
	return s
}

type CancelApplyLinkMicResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelApplyLinkMicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelApplyLinkMicResponseBody) GoString() string {
	return s.String()
}

func (s *CancelApplyLinkMicResponseBody) SetRequestId(v string) *CancelApplyLinkMicResponseBody {
	s.RequestId = &v
	return s
}

type CancelApplyLinkMicResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelApplyLinkMicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelApplyLinkMicResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelApplyLinkMicResponse) GoString() string {
	return s.String()
}

func (s *CancelApplyLinkMicResponse) SetHeaders(v map[string]*string) *CancelApplyLinkMicResponse {
	s.Headers = v
	return s
}

func (s *CancelApplyLinkMicResponse) SetBody(v *CancelApplyLinkMicResponseBody) *CancelApplyLinkMicResponse {
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
	// 应用配置状态
	AppConfigStatus *string `json:"AppConfigStatus,omitempty" xml:"AppConfigStatus,omitempty"`
	// 应用Key
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 样板间信息
	StandardRoomInfo *string `json:"StandardRoomInfo,omitempty" xml:"StandardRoomInfo,omitempty"`
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

func (s *GetAppResponseBodyResult) SetAppConfigStatus(v string) *GetAppResponseBodyResult {
	s.AppConfigStatus = &v
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

func (s *GetAppResponseBodyResult) SetIntegrationMode(v string) *GetAppResponseBodyResult {
	s.IntegrationMode = &v
	return s
}

func (s *GetAppResponseBodyResult) SetStandardRoomInfo(v string) *GetAppResponseBodyResult {
	s.StandardRoomInfo = &v
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

type SendCustomMessageToUsersRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 消息体内容。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// 消息指定的接收人ID列表。
	ReceiverList []*string `json:"ReceiverList,omitempty" xml:"ReceiverList,omitempty" type:"Repeated"`
}

func (s SendCustomMessageToUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToUsersRequest) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToUsersRequest) SetAppId(v string) *SendCustomMessageToUsersRequest {
	s.AppId = &v
	return s
}

func (s *SendCustomMessageToUsersRequest) SetRoomId(v string) *SendCustomMessageToUsersRequest {
	s.RoomId = &v
	return s
}

func (s *SendCustomMessageToUsersRequest) SetBody(v string) *SendCustomMessageToUsersRequest {
	s.Body = &v
	return s
}

func (s *SendCustomMessageToUsersRequest) SetReceiverList(v []*string) *SendCustomMessageToUsersRequest {
	s.ReceiverList = v
	return s
}

type SendCustomMessageToUsersResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *SendCustomMessageToUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendCustomMessageToUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToUsersResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToUsersResponseBody) SetRequestId(v string) *SendCustomMessageToUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomMessageToUsersResponseBody) SetResult(v *SendCustomMessageToUsersResponseBodyResult) *SendCustomMessageToUsersResponseBody {
	s.Result = v
	return s
}

type SendCustomMessageToUsersResponseBodyResult struct {
	// 消息的唯一ID标识。由数字、大小写字母组成，长度不超过20。
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendCustomMessageToUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToUsersResponseBodyResult) SetMessageId(v string) *SendCustomMessageToUsersResponseBodyResult {
	s.MessageId = &v
	return s
}

type SendCustomMessageToUsersResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCustomMessageToUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCustomMessageToUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToUsersResponse) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToUsersResponse) SetHeaders(v map[string]*string) *SendCustomMessageToUsersResponse {
	s.Headers = v
	return s
}

func (s *SendCustomMessageToUsersResponse) SetBody(v *SendCustomMessageToUsersResponseBody) *SendCustomMessageToUsersResponse {
	s.Body = v
	return s
}

type BanAllCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BanAllCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s BanAllCommentRequest) GoString() string {
	return s.String()
}

func (s *BanAllCommentRequest) SetAppId(v string) *BanAllCommentRequest {
	s.AppId = &v
	return s
}

func (s *BanAllCommentRequest) SetRoomId(v string) *BanAllCommentRequest {
	s.RoomId = &v
	return s
}

func (s *BanAllCommentRequest) SetUserId(v string) *BanAllCommentRequest {
	s.UserId = &v
	return s
}

type BanAllCommentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作成功标识
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BanAllCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BanAllCommentResponseBody) GoString() string {
	return s.String()
}

func (s *BanAllCommentResponseBody) SetRequestId(v string) *BanAllCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *BanAllCommentResponseBody) SetResult(v bool) *BanAllCommentResponseBody {
	s.Result = &v
	return s
}

type BanAllCommentResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BanAllCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BanAllCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s BanAllCommentResponse) GoString() string {
	return s.String()
}

func (s *BanAllCommentResponse) SetHeaders(v map[string]*string) *BanAllCommentResponse {
	s.Headers = v
	return s
}

func (s *BanAllCommentResponse) SetBody(v *BanAllCommentResponseBody) *BanAllCommentResponse {
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
	// 封面图片
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 用户自定义数据存储
	UserDefineField *string `json:"UserDefineField,omitempty" xml:"UserDefineField,omitempty"`
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

func (s *GetLiveResponseBodyResult) SetCoverUrl(v string) *GetLiveResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetUserDefineField(v string) *GetLiveResponseBodyResult {
	s.UserDefineField = &v
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

type CreateRoomShrinkRequest struct {
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
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomShrinkRequest) SetAppId(v string) *CreateRoomShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetTemplateId(v string) *CreateRoomShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetRoomId(v string) *CreateRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetTitle(v string) *CreateRoomShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetNotice(v string) *CreateRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetRoomOwnerId(v string) *CreateRoomShrinkRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetExtensionShrink(v string) *CreateRoomShrinkRequest {
	s.ExtensionShrink = &v
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

type UpdateConferenceRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 会议标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConferenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateConferenceRequest) SetConferenceId(v string) *UpdateConferenceRequest {
	s.ConferenceId = &v
	return s
}

func (s *UpdateConferenceRequest) SetTitle(v string) *UpdateConferenceRequest {
	s.Title = &v
	return s
}

type UpdateConferenceResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConferenceResponseBody) SetRequestId(v string) *UpdateConferenceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConferenceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConferenceResponse) GoString() string {
	return s.String()
}

func (s *UpdateConferenceResponse) SetHeaders(v map[string]*string) *UpdateConferenceResponse {
	s.Headers = v
	return s
}

func (s *UpdateConferenceResponse) SetBody(v *UpdateConferenceResponseBody) *UpdateConferenceResponse {
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

func (client *Client) CreateIceProjectWithOptions(request *CreateIceProjectRequest, runtime *util.RuntimeOptions) (_result *CreateIceProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIceProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIceProject"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIceProject(request *CreateIceProjectRequest) (_result *CreateIceProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIceProjectResponse{}
	_body, _err := client.CreateIceProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveMemberWithOptions(request *RemoveMemberRequest, runtime *util.RuntimeOptions) (_result *RemoveMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveMember"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveMember(request *RemoveMemberRequest) (_result *RemoveMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveMemberResponse{}
	_body, _err := client.RemoveMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplyLinkMicUsersWithOptions(request *ListApplyLinkMicUsersRequest, runtime *util.RuntimeOptions) (_result *ListApplyLinkMicUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListApplyLinkMicUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApplyLinkMicUsers"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplyLinkMicUsers(request *ListApplyLinkMicUsersRequest) (_result *ListApplyLinkMicUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplyLinkMicUsersResponse{}
	_body, _err := client.ListApplyLinkMicUsersWithOptions(request, runtime)
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

func (client *Client) BanCommentWithOptions(request *BanCommentRequest, runtime *util.RuntimeOptions) (_result *BanCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BanCommentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BanComment"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BanComment(request *BanCommentRequest) (_result *BanCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BanCommentResponse{}
	_body, _err := client.BanCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddMemberWithOptions(request *AddMemberRequest, runtime *util.RuntimeOptions) (_result *AddMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddMember"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMember(request *AddMemberRequest) (_result *AddMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMemberResponse{}
	_body, _err := client.AddMemberWithOptions(request, runtime)
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

func (client *Client) RegisterIceOssMediaWithOptions(request *RegisterIceOssMediaRequest, runtime *util.RuntimeOptions) (_result *RegisterIceOssMediaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterIceOssMediaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterIceOssMedia"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterIceOssMedia(request *RegisterIceOssMediaRequest) (_result *RegisterIceOssMediaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterIceOssMediaResponse{}
	_body, _err := client.RegisterIceOssMediaWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateConference"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SendCustomMessageToAllWithOptions(request *SendCustomMessageToAllRequest, runtime *util.RuntimeOptions) (_result *SendCustomMessageToAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCustomMessageToAllResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCustomMessageToAll"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCustomMessageToAll(request *SendCustomMessageToAllRequest) (_result *SendCustomMessageToAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomMessageToAllResponse{}
	_body, _err := client.SendCustomMessageToAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AgreeLinkMicWithOptions(request *AgreeLinkMicRequest, runtime *util.RuntimeOptions) (_result *AgreeLinkMicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AgreeLinkMicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AgreeLinkMic"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AgreeLinkMic(request *AgreeLinkMicRequest) (_result *AgreeLinkMicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AgreeLinkMicResponse{}
	_body, _err := client.AgreeLinkMicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainOwnerVerifyContentWithOptions(request *GetDomainOwnerVerifyContentRequest, runtime *util.RuntimeOptions) (_result *GetDomainOwnerVerifyContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDomainOwnerVerifyContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDomainOwnerVerifyContent"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomainOwnerVerifyContent(request *GetDomainOwnerVerifyContentRequest) (_result *GetDomainOwnerVerifyContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainOwnerVerifyContentResponse{}
	_body, _err := client.GetDomainOwnerVerifyContentWithOptions(request, runtime)
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

func (client *Client) DeleteConferenceWithOptions(request *DeleteConferenceRequest, runtime *util.RuntimeOptions) (_result *DeleteConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConferenceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConference"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) VerifyDomainOwnerWithOptions(request *VerifyDomainOwnerRequest, runtime *util.RuntimeOptions) (_result *VerifyDomainOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyDomainOwnerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyDomainOwner"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyDomainOwner(request *VerifyDomainOwnerRequest) (_result *VerifyDomainOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyDomainOwnerResponse{}
	_body, _err := client.VerifyDomainOwnerWithOptions(request, runtime)
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

func (client *Client) GetStandardRoomJumpUrlWithOptions(request *GetStandardRoomJumpUrlRequest, runtime *util.RuntimeOptions) (_result *GetStandardRoomJumpUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStandardRoomJumpUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStandardRoomJumpUrl"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStandardRoomJumpUrl(request *GetStandardRoomJumpUrlRequest) (_result *GetStandardRoomJumpUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStandardRoomJumpUrlResponse{}
	_body, _err := client.GetStandardRoomJumpUrlWithOptions(request, runtime)
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

func (client *Client) ListRoomLivesWithOptions(tmpReq *ListRoomLivesRequest, runtime *util.RuntimeOptions) (_result *ListRoomLivesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListRoomLivesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoomIdList)) {
		request.RoomIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoomIdList, tea.String("RoomIdList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRoomLivesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRoomLives"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoomLives(request *ListRoomLivesRequest) (_result *ListRoomLivesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRoomLivesResponse{}
	_body, _err := client.ListRoomLivesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoomWithOptions(tmpReq *UpdateRoomRequest, runtime *util.RuntimeOptions) (_result *UpdateRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
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

func (client *Client) SendCommentWithOptions(tmpReq *SendCommentRequest, runtime *util.RuntimeOptions) (_result *SendCommentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendCommentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCommentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendComment"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendComment(request *SendCommentRequest) (_result *SendCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCommentResponse{}
	_body, _err := client.SendCommentWithOptions(request, runtime)
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

func (client *Client) GetConferenceWithOptions(request *GetConferenceRequest, runtime *util.RuntimeOptions) (_result *GetConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConferenceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConference"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConference(request *GetConferenceRequest) (_result *GetConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConferenceResponse{}
	_body, _err := client.GetConferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RejectLinkMicWithOptions(request *RejectLinkMicRequest, runtime *util.RuntimeOptions) (_result *RejectLinkMicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RejectLinkMicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RejectLinkMic"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RejectLinkMic(request *RejectLinkMicRequest) (_result *RejectLinkMicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectLinkMicResponse{}
	_body, _err := client.RejectLinkMicWithOptions(request, runtime)
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

func (client *Client) CancelBanAllCommentWithOptions(request *CancelBanAllCommentRequest, runtime *util.RuntimeOptions) (_result *CancelBanAllCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelBanAllCommentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelBanAllComment"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelBanAllComment(request *CancelBanAllCommentRequest) (_result *CancelBanAllCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelBanAllCommentResponse{}
	_body, _err := client.CancelBanAllCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConferenceUsersWithOptions(request *ListConferenceUsersRequest, runtime *util.RuntimeOptions) (_result *ListConferenceUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConferenceUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConferenceUsers"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConferenceUsers(request *ListConferenceUsersRequest) (_result *ListConferenceUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConferenceUsersResponse{}
	_body, _err := client.ListConferenceUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelBanCommentWithOptions(request *CancelBanCommentRequest, runtime *util.RuntimeOptions) (_result *CancelBanCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelBanCommentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelBanComment"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelBanComment(request *CancelBanCommentRequest) (_result *CancelBanCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelBanCommentResponse{}
	_body, _err := client.CancelBanCommentWithOptions(request, runtime)
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

func (client *Client) ApplyLinkMicWithOptions(request *ApplyLinkMicRequest, runtime *util.RuntimeOptions) (_result *ApplyLinkMicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyLinkMicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyLinkMic"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyLinkMic(request *ApplyLinkMicRequest) (_result *ApplyLinkMicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyLinkMicResponse{}
	_body, _err := client.ApplyLinkMicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelApplyLinkMicWithOptions(request *CancelApplyLinkMicRequest, runtime *util.RuntimeOptions) (_result *CancelApplyLinkMicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelApplyLinkMicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelApplyLinkMic"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelApplyLinkMic(request *CancelApplyLinkMicRequest) (_result *CancelApplyLinkMicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelApplyLinkMicResponse{}
	_body, _err := client.CancelApplyLinkMicWithOptions(request, runtime)
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

func (client *Client) SendCustomMessageToUsersWithOptions(request *SendCustomMessageToUsersRequest, runtime *util.RuntimeOptions) (_result *SendCustomMessageToUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCustomMessageToUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCustomMessageToUsers"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCustomMessageToUsers(request *SendCustomMessageToUsersRequest) (_result *SendCustomMessageToUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomMessageToUsersResponse{}
	_body, _err := client.SendCustomMessageToUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BanAllCommentWithOptions(request *BanAllCommentRequest, runtime *util.RuntimeOptions) (_result *BanAllCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BanAllCommentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BanAllComment"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BanAllComment(request *BanAllCommentRequest) (_result *BanAllCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BanAllCommentResponse{}
	_body, _err := client.BanAllCommentWithOptions(request, runtime)
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

func (client *Client) CreateRoomWithOptions(tmpReq *CreateRoomRequest, runtime *util.RuntimeOptions) (_result *CreateRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
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

func (client *Client) UpdateConferenceWithOptions(request *UpdateConferenceRequest, runtime *util.RuntimeOptions) (_result *UpdateConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateConferenceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateConference"), tea.String("2021-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConference(request *UpdateConferenceRequest) (_result *UpdateConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConferenceResponse{}
	_body, _err := client.UpdateConferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
