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

type ListAppInfosRequest struct {
	RequestParams *ListAppInfosRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s ListAppInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfosRequest) GoString() string {
	return s.String()
}

func (s *ListAppInfosRequest) SetRequestParams(v *ListAppInfosRequestRequestParams) *ListAppInfosRequest {
	s.RequestParams = v
	return s
}

type ListAppInfosRequestRequestParams struct {
	// 关键字类型，包含appName、appId两类
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 关键字
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// 分页大小，大于0的任意数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 页码，从1开始
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListAppInfosRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfosRequestRequestParams) GoString() string {
	return s.String()
}

func (s *ListAppInfosRequestRequestParams) SetType(v string) *ListAppInfosRequestRequestParams {
	s.Type = &v
	return s
}

func (s *ListAppInfosRequestRequestParams) SetKeyword(v string) *ListAppInfosRequestRequestParams {
	s.Keyword = &v
	return s
}

func (s *ListAppInfosRequestRequestParams) SetPageSize(v int32) *ListAppInfosRequestRequestParams {
	s.PageSize = &v
	return s
}

func (s *ListAppInfosRequestRequestParams) SetPageNumber(v int32) *ListAppInfosRequestRequestParams {
	s.PageNumber = &v
	return s
}

type ListAppInfosShrinkRequest struct {
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s ListAppInfosShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfosShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAppInfosShrinkRequest) SetRequestParamsShrink(v string) *ListAppInfosShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type ListAppInfosResponseBody struct {
	// desc
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// result
	Result *ListAppInfosResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAppInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInfosResponseBody) SetMessage(v string) *ListAppInfosResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppInfosResponseBody) SetRequestId(v string) *ListAppInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInfosResponseBody) SetHttpStatusCode(v int32) *ListAppInfosResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAppInfosResponseBody) SetCode(v string) *ListAppInfosResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppInfosResponseBody) SetSuccess(v bool) *ListAppInfosResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppInfosResponseBody) SetResult(v *ListAppInfosResponseBodyResult) *ListAppInfosResponseBody {
	s.Result = v
	return s
}

type ListAppInfosResponseBodyResult struct {
	// 总数，用于分页
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 应用信息列表
	AppInfos []*ListAppInfosResponseBodyResultAppInfos `json:"AppInfos,omitempty" xml:"AppInfos,omitempty" type:"Repeated"`
}

func (s ListAppInfosResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfosResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppInfosResponseBodyResult) SetTotalCount(v int32) *ListAppInfosResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListAppInfosResponseBodyResult) SetAppInfos(v []*ListAppInfosResponseBodyResultAppInfos) *ListAppInfosResponseBodyResult {
	s.AppInfos = v
	return s
}

type ListAppInfosResponseBodyResultAppInfos struct {
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用名
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用状态
	AppStatus *int32 `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// 产品版本
	ProdVersion *string `json:"ProdVersion,omitempty" xml:"ProdVersion,omitempty"`
	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAppInfosResponseBodyResultAppInfos) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfosResponseBodyResultAppInfos) GoString() string {
	return s.String()
}

func (s *ListAppInfosResponseBodyResultAppInfos) SetAppId(v string) *ListAppInfosResponseBodyResultAppInfos {
	s.AppId = &v
	return s
}

func (s *ListAppInfosResponseBodyResultAppInfos) SetAppName(v string) *ListAppInfosResponseBodyResultAppInfos {
	s.AppName = &v
	return s
}

func (s *ListAppInfosResponseBodyResultAppInfos) SetCreateTime(v string) *ListAppInfosResponseBodyResultAppInfos {
	s.CreateTime = &v
	return s
}

func (s *ListAppInfosResponseBodyResultAppInfos) SetAppStatus(v int32) *ListAppInfosResponseBodyResultAppInfos {
	s.AppStatus = &v
	return s
}

func (s *ListAppInfosResponseBodyResultAppInfos) SetProdVersion(v string) *ListAppInfosResponseBodyResultAppInfos {
	s.ProdVersion = &v
	return s
}

func (s *ListAppInfosResponseBodyResultAppInfos) SetInstanceId(v string) *ListAppInfosResponseBodyResultAppInfos {
	s.InstanceId = &v
	return s
}

type ListAppInfosResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfosResponse) GoString() string {
	return s.String()
}

func (s *ListAppInfosResponse) SetHeaders(v map[string]*string) *ListAppInfosResponse {
	s.Headers = v
	return s
}

func (s *ListAppInfosResponse) SetBody(v *ListAppInfosResponseBody) *ListAppInfosResponse {
	s.Body = v
	return s
}

type RemoveSingleChatExtensionByKeysRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 单聊移除拓展字段请求实体
	RequestParams *RemoveSingleChatExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RemoveSingleChatExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSingleChatExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *RemoveSingleChatExtensionByKeysRequest) SetAppId(v string) *RemoveSingleChatExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *RemoveSingleChatExtensionByKeysRequest) SetRequestParams(v *RemoveSingleChatExtensionByKeysRequestRequestParams) *RemoveSingleChatExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type RemoveSingleChatExtensionByKeysRequestRequestParams struct {
	// 用户id
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 会话id
	AppCid *string   `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	Keys   []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s RemoveSingleChatExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RemoveSingleChatExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RemoveSingleChatExtensionByKeysRequestRequestParams) SetAppUid(v string) *RemoveSingleChatExtensionByKeysRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *RemoveSingleChatExtensionByKeysRequestRequestParams) SetAppCid(v string) *RemoveSingleChatExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RemoveSingleChatExtensionByKeysRequestRequestParams) SetKeys(v []*string) *RemoveSingleChatExtensionByKeysRequestRequestParams {
	s.Keys = v
	return s
}

type RemoveSingleChatExtensionByKeysShrinkRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 单聊移除拓展字段请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RemoveSingleChatExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSingleChatExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveSingleChatExtensionByKeysShrinkRequest) SetAppId(v string) *RemoveSingleChatExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveSingleChatExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *RemoveSingleChatExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RemoveSingleChatExtensionByKeysResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveSingleChatExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSingleChatExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSingleChatExtensionByKeysResponseBody) SetRequestId(v string) *RemoveSingleChatExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSingleChatExtensionByKeysResponseBody) SetCode(v string) *RemoveSingleChatExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSingleChatExtensionByKeysResponseBody) SetMessage(v string) *RemoveSingleChatExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type RemoveSingleChatExtensionByKeysResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveSingleChatExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSingleChatExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSingleChatExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *RemoveSingleChatExtensionByKeysResponse) SetHeaders(v map[string]*string) *RemoveSingleChatExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *RemoveSingleChatExtensionByKeysResponse) SetBody(v *RemoveSingleChatExtensionByKeysResponseBody) *RemoveSingleChatExtensionByKeysResponse {
	s.Body = v
	return s
}

type ImportMessageRequest struct {
	// AppId
	AppId         *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *ImportMessageRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s ImportMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageRequest) GoString() string {
	return s.String()
}

func (s *ImportMessageRequest) SetAppId(v string) *ImportMessageRequest {
	s.AppId = &v
	return s
}

func (s *ImportMessageRequest) SetRequestParams(v *ImportMessageRequestRequestParams) *ImportMessageRequest {
	s.RequestParams = v
	return s
}

type ImportMessageRequestRequestParams struct {
	Messages []*ImportMessageRequestRequestParamsMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
}

func (s ImportMessageRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageRequestRequestParams) GoString() string {
	return s.String()
}

func (s *ImportMessageRequestRequestParams) SetMessages(v []*ImportMessageRequestRequestParamsMessages) *ImportMessageRequestRequestParams {
	s.Messages = v
	return s
}

type ImportMessageRequestRequestParamsMessages struct {
	// 唯一标识，用于重入
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 会话类型1 单聊 2 群聊
	ConversationType *int64 `json:"ConversationType,omitempty" xml:"ConversationType,omitempty"`
	// 发送者ID
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 接受者列表, 群聊如果列表为空者全员接收
	ReceiverIds []*string `json:"ReceiverIds,omitempty" xml:"ReceiverIds,omitempty" type:"Repeated"`
	// 消息类型
	ContentType *int64 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// 消息内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 消息发送时间戳
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 自定义信息
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s ImportMessageRequestRequestParamsMessages) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageRequestRequestParamsMessages) GoString() string {
	return s.String()
}

func (s *ImportMessageRequestRequestParamsMessages) SetUuid(v string) *ImportMessageRequestRequestParamsMessages {
	s.Uuid = &v
	return s
}

func (s *ImportMessageRequestRequestParamsMessages) SetAppCid(v string) *ImportMessageRequestRequestParamsMessages {
	s.AppCid = &v
	return s
}

func (s *ImportMessageRequestRequestParamsMessages) SetConversationType(v int64) *ImportMessageRequestRequestParamsMessages {
	s.ConversationType = &v
	return s
}

func (s *ImportMessageRequestRequestParamsMessages) SetSenderId(v string) *ImportMessageRequestRequestParamsMessages {
	s.SenderId = &v
	return s
}

func (s *ImportMessageRequestRequestParamsMessages) SetReceiverIds(v []*string) *ImportMessageRequestRequestParamsMessages {
	s.ReceiverIds = v
	return s
}

func (s *ImportMessageRequestRequestParamsMessages) SetContentType(v int64) *ImportMessageRequestRequestParamsMessages {
	s.ContentType = &v
	return s
}

func (s *ImportMessageRequestRequestParamsMessages) SetContent(v string) *ImportMessageRequestRequestParamsMessages {
	s.Content = &v
	return s
}

func (s *ImportMessageRequestRequestParamsMessages) SetCreateTime(v int64) *ImportMessageRequestRequestParamsMessages {
	s.CreateTime = &v
	return s
}

func (s *ImportMessageRequestRequestParamsMessages) SetExtensions(v map[string]*string) *ImportMessageRequestRequestParamsMessages {
	s.Extensions = v
	return s
}

type ImportMessageShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s ImportMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportMessageShrinkRequest) SetAppId(v string) *ImportMessageShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ImportMessageShrinkRequest) SetRequestParamsShrink(v string) *ImportMessageShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type ImportMessageResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *ImportMessageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ImportMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ImportMessageResponseBody) SetRequestId(v string) *ImportMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportMessageResponseBody) SetCode(v string) *ImportMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ImportMessageResponseBody) SetMessage(v string) *ImportMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ImportMessageResponseBody) SetResult(v *ImportMessageResponseBodyResult) *ImportMessageResponseBody {
	s.Result = v
	return s
}

type ImportMessageResponseBodyResult struct {
	ImportMessageResult map[string]*ResultImportMessageResultValue `json:"ImportMessageResult,omitempty" xml:"ImportMessageResult,omitempty"`
}

func (s ImportMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ImportMessageResponseBodyResult) SetImportMessageResult(v map[string]*ResultImportMessageResultValue) *ImportMessageResponseBodyResult {
	s.ImportMessageResult = v
	return s
}

type ImportMessageResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageResponse) GoString() string {
	return s.String()
}

func (s *ImportMessageResponse) SetHeaders(v map[string]*string) *ImportMessageResponse {
	s.Headers = v
	return s
}

func (s *ImportMessageResponse) SetBody(v *ImportMessageResponseBody) *ImportMessageResponse {
	s.Body = v
	return s
}

type UnbindInterconnectionUidRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 解绑用户请求体
	RequestParams *UnbindInterconnectionUidRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s UnbindInterconnectionUidRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindInterconnectionUidRequest) GoString() string {
	return s.String()
}

func (s *UnbindInterconnectionUidRequest) SetAppId(v string) *UnbindInterconnectionUidRequest {
	s.AppId = &v
	return s
}

func (s *UnbindInterconnectionUidRequest) SetRequestParams(v *UnbindInterconnectionUidRequestRequestParams) *UnbindInterconnectionUidRequest {
	s.RequestParams = v
	return s
}

type UnbindInterconnectionUidRequestRequestParams struct {
	// AIM 用户ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 钉钉 用户ID
	DingTalkUid *string `json:"DingTalkUid,omitempty" xml:"DingTalkUid,omitempty"`
}

func (s UnbindInterconnectionUidRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UnbindInterconnectionUidRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UnbindInterconnectionUidRequestRequestParams) SetAppUid(v string) *UnbindInterconnectionUidRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *UnbindInterconnectionUidRequestRequestParams) SetDingTalkUid(v string) *UnbindInterconnectionUidRequestRequestParams {
	s.DingTalkUid = &v
	return s
}

type UnbindInterconnectionUidShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 解绑用户请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s UnbindInterconnectionUidShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindInterconnectionUidShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnbindInterconnectionUidShrinkRequest) SetAppId(v string) *UnbindInterconnectionUidShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UnbindInterconnectionUidShrinkRequest) SetRequestParamsShrink(v string) *UnbindInterconnectionUidShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type UnbindInterconnectionUidResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UnbindInterconnectionUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindInterconnectionUidResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindInterconnectionUidResponseBody) SetRequestId(v string) *UnbindInterconnectionUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindInterconnectionUidResponseBody) SetCode(v string) *UnbindInterconnectionUidResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindInterconnectionUidResponseBody) SetMessage(v string) *UnbindInterconnectionUidResponseBody {
	s.Message = &v
	return s
}

type UnbindInterconnectionUidResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindInterconnectionUidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindInterconnectionUidResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindInterconnectionUidResponse) GoString() string {
	return s.String()
}

func (s *UnbindInterconnectionUidResponse) SetHeaders(v map[string]*string) *UnbindInterconnectionUidResponse {
	s.Headers = v
	return s
}

func (s *UnbindInterconnectionUidResponse) SetBody(v *UnbindInterconnectionUidResponseBody) *UnbindInterconnectionUidResponse {
	s.Body = v
	return s
}

type SilenceAllGroupMembersRequest struct {
	// AppId
	AppId         *string                                     `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *SilenceAllGroupMembersRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SilenceAllGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s SilenceAllGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *SilenceAllGroupMembersRequest) SetAppId(v string) *SilenceAllGroupMembersRequest {
	s.AppId = &v
	return s
}

func (s *SilenceAllGroupMembersRequest) SetRequestParams(v *SilenceAllGroupMembersRequestRequestParams) *SilenceAllGroupMembersRequest {
	s.RequestParams = v
	return s
}

type SilenceAllGroupMembersRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 操作者uid
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
}

func (s SilenceAllGroupMembersRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SilenceAllGroupMembersRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SilenceAllGroupMembersRequestRequestParams) SetAppCid(v string) *SilenceAllGroupMembersRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *SilenceAllGroupMembersRequestRequestParams) SetOperatorAppUid(v string) *SilenceAllGroupMembersRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

type SilenceAllGroupMembersShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SilenceAllGroupMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SilenceAllGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *SilenceAllGroupMembersShrinkRequest) SetAppId(v string) *SilenceAllGroupMembersShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SilenceAllGroupMembersShrinkRequest) SetRequestParamsShrink(v string) *SilenceAllGroupMembersShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SilenceAllGroupMembersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SilenceAllGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SilenceAllGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *SilenceAllGroupMembersResponseBody) SetRequestId(v string) *SilenceAllGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SilenceAllGroupMembersResponseBody) SetCode(v string) *SilenceAllGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *SilenceAllGroupMembersResponseBody) SetMessage(v string) *SilenceAllGroupMembersResponseBody {
	s.Message = &v
	return s
}

type SilenceAllGroupMembersResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SilenceAllGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SilenceAllGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s SilenceAllGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *SilenceAllGroupMembersResponse) SetHeaders(v map[string]*string) *SilenceAllGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *SilenceAllGroupMembersResponse) SetBody(v *SilenceAllGroupMembersResponseBody) *SilenceAllGroupMembersResponse {
	s.Body = v
	return s
}

type ListRoomMessagesRequest struct {
	// 请求参数的结构体。
	Request *ListRoomMessagesRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s ListRoomMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListRoomMessagesRequest) SetRequest(v *ListRoomMessagesRequestRequest) *ListRoomMessagesRequest {
	s.Request = v
	return s
}

type ListRoomMessagesRequestRequest struct {
	// 应用的appKey。
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 房间ID，由调用CreateRoom时返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 要查询的消息的类型,请传递100000以上的整数，如果不传，则默认拉取全部类型的消息。
	SubType *int32 `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// 分页查询时的页数，从1开始，每次分页查询时加1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时的请求大小，要求大于0，且最大不得超过100。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRoomMessagesRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomMessagesRequestRequest) GoString() string {
	return s.String()
}

func (s *ListRoomMessagesRequestRequest) SetDomain(v string) *ListRoomMessagesRequestRequest {
	s.Domain = &v
	return s
}

func (s *ListRoomMessagesRequestRequest) SetRoomId(v string) *ListRoomMessagesRequestRequest {
	s.RoomId = &v
	return s
}

func (s *ListRoomMessagesRequestRequest) SetSubType(v int32) *ListRoomMessagesRequestRequest {
	s.SubType = &v
	return s
}

func (s *ListRoomMessagesRequestRequest) SetPageNumber(v int32) *ListRoomMessagesRequestRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoomMessagesRequestRequest) SetPageSize(v int32) *ListRoomMessagesRequestRequest {
	s.PageSize = &v
	return s
}

type ListRoomMessagesResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功。
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码，请求异常时返回。
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息，请求异常时返回。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求的返回结果。
	Result *ListRoomMessagesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListRoomMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoomMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoomMessagesResponseBody) SetRequestId(v string) *ListRoomMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoomMessagesResponseBody) SetResponseSuccess(v bool) *ListRoomMessagesResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *ListRoomMessagesResponseBody) SetErrorCode(v string) *ListRoomMessagesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRoomMessagesResponseBody) SetErrorMessage(v string) *ListRoomMessagesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRoomMessagesResponseBody) SetResult(v *ListRoomMessagesResponseBodyResult) *ListRoomMessagesResponseBody {
	s.Result = v
	return s
}

type ListRoomMessagesResponseBodyResult struct {
	// 互动消息的总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 房间的互动消息列表，按照发送时间戳由大到小排序。
	RoomMessageList []*ListRoomMessagesResponseBodyResultRoomMessageList `json:"RoomMessageList,omitempty" xml:"RoomMessageList,omitempty" type:"Repeated"`
	// 是否还有下一页查询的数据。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
}

func (s ListRoomMessagesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoomMessagesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoomMessagesResponseBodyResult) SetTotalCount(v int32) *ListRoomMessagesResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListRoomMessagesResponseBodyResult) SetRoomMessageList(v []*ListRoomMessagesResponseBodyResultRoomMessageList) *ListRoomMessagesResponseBodyResult {
	s.RoomMessageList = v
	return s
}

func (s *ListRoomMessagesResponseBodyResult) SetHasMore(v bool) *ListRoomMessagesResponseBodyResult {
	s.HasMore = &v
	return s
}

type ListRoomMessagesResponseBodyResultRoomMessageList struct {
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 消息的唯一ID标识。由数字、大小写字母组成，长度不超过20。
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// 消息的类型。
	SubType *int32 `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// 消息的发送者ID。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 消息的发送时间,毫秒unix时间戳。
	SendTimeMillis *int64 `json:"SendTimeMillis,omitempty" xml:"SendTimeMillis,omitempty"`
	// 消息体。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s ListRoomMessagesResponseBodyResultRoomMessageList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomMessagesResponseBodyResultRoomMessageList) GoString() string {
	return s.String()
}

func (s *ListRoomMessagesResponseBodyResultRoomMessageList) SetRoomId(v string) *ListRoomMessagesResponseBodyResultRoomMessageList {
	s.RoomId = &v
	return s
}

func (s *ListRoomMessagesResponseBodyResultRoomMessageList) SetMessageId(v string) *ListRoomMessagesResponseBodyResultRoomMessageList {
	s.MessageId = &v
	return s
}

func (s *ListRoomMessagesResponseBodyResultRoomMessageList) SetSubType(v int32) *ListRoomMessagesResponseBodyResultRoomMessageList {
	s.SubType = &v
	return s
}

func (s *ListRoomMessagesResponseBodyResultRoomMessageList) SetSenderId(v string) *ListRoomMessagesResponseBodyResultRoomMessageList {
	s.SenderId = &v
	return s
}

func (s *ListRoomMessagesResponseBodyResultRoomMessageList) SetSendTimeMillis(v int64) *ListRoomMessagesResponseBodyResultRoomMessageList {
	s.SendTimeMillis = &v
	return s
}

func (s *ListRoomMessagesResponseBodyResultRoomMessageList) SetBody(v string) *ListRoomMessagesResponseBodyResultRoomMessageList {
	s.Body = &v
	return s
}

type ListRoomMessagesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRoomMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoomMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoomMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListRoomMessagesResponse) SetHeaders(v map[string]*string) *ListRoomMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListRoomMessagesResponse) SetBody(v *ListRoomMessagesResponseBody) *ListRoomMessagesResponse {
	s.Body = v
	return s
}

type SetGroupExtensionByKeysRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群聊设置扩展字段请求实体
	RequestParams *SetGroupExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SetGroupExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGroupExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *SetGroupExtensionByKeysRequest) SetAppId(v string) *SetGroupExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *SetGroupExtensionByKeysRequest) SetRequestParams(v *SetGroupExtensionByKeysRequestRequestParams) *SetGroupExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type SetGroupExtensionByKeysRequestRequestParams struct {
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 扩展字段
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s SetGroupExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SetGroupExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SetGroupExtensionByKeysRequestRequestParams) SetAppCid(v string) *SetGroupExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *SetGroupExtensionByKeysRequestRequestParams) SetExtensions(v map[string]*string) *SetGroupExtensionByKeysRequestRequestParams {
	s.Extensions = v
	return s
}

type SetGroupExtensionByKeysShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群聊设置扩展字段请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SetGroupExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGroupExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetGroupExtensionByKeysShrinkRequest) SetAppId(v string) *SetGroupExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SetGroupExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *SetGroupExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SetGroupExtensionByKeysResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SetGroupExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGroupExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *SetGroupExtensionByKeysResponseBody) SetRequestId(v string) *SetGroupExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGroupExtensionByKeysResponseBody) SetCode(v string) *SetGroupExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *SetGroupExtensionByKeysResponseBody) SetMessage(v string) *SetGroupExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type SetGroupExtensionByKeysResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGroupExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGroupExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGroupExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *SetGroupExtensionByKeysResponse) SetHeaders(v map[string]*string) *SetGroupExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *SetGroupExtensionByKeysResponse) SetBody(v *SetGroupExtensionByKeysResponseBody) *SetGroupExtensionByKeysResponse {
	s.Body = v
	return s
}

type RemoveGroupMemberExtensionByKeysRequest struct {
	// App ID, IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 删除群成员扩展信息的请求体
	RequestParams *RemoveGroupMemberExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RemoveGroupMemberExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberExtensionByKeysRequest) SetAppId(v string) *RemoveGroupMemberExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupMemberExtensionByKeysRequest) SetRequestParams(v *RemoveGroupMemberExtensionByKeysRequestRequestParams) *RemoveGroupMemberExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type RemoveGroupMemberExtensionByKeysRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 用户ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 扩展信息中需要删除的key列表
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s RemoveGroupMemberExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberExtensionByKeysRequestRequestParams) SetAppCid(v string) *RemoveGroupMemberExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RemoveGroupMemberExtensionByKeysRequestRequestParams) SetAppUid(v string) *RemoveGroupMemberExtensionByKeysRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *RemoveGroupMemberExtensionByKeysRequestRequestParams) SetKeys(v []*string) *RemoveGroupMemberExtensionByKeysRequestRequestParams {
	s.Keys = v
	return s
}

type RemoveGroupMemberExtensionByKeysShrinkRequest struct {
	// App ID, IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 删除群成员扩展信息的请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RemoveGroupMemberExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberExtensionByKeysShrinkRequest) SetAppId(v string) *RemoveGroupMemberExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupMemberExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *RemoveGroupMemberExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RemoveGroupMemberExtensionByKeysResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveGroupMemberExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberExtensionByKeysResponseBody) SetRequestId(v string) *RemoveGroupMemberExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGroupMemberExtensionByKeysResponseBody) SetCode(v string) *RemoveGroupMemberExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveGroupMemberExtensionByKeysResponseBody) SetMessage(v string) *RemoveGroupMemberExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type RemoveGroupMemberExtensionByKeysResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveGroupMemberExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveGroupMemberExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberExtensionByKeysResponse) SetHeaders(v map[string]*string) *RemoveGroupMemberExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupMemberExtensionByKeysResponse) SetBody(v *RemoveGroupMemberExtensionByKeysResponseBody) *RemoveGroupMemberExtensionByKeysResponse {
	s.Body = v
	return s
}

type AddGroupSilenceBlacklistRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言添加白名单请求体
	RequestParams *AddGroupSilenceBlacklistRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s AddGroupSilenceBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceBlacklistRequest) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceBlacklistRequest) SetAppId(v string) *AddGroupSilenceBlacklistRequest {
	s.AppId = &v
	return s
}

func (s *AddGroupSilenceBlacklistRequest) SetRequestParams(v *AddGroupSilenceBlacklistRequestRequestParams) *AddGroupSilenceBlacklistRequest {
	s.RequestParams = v
	return s
}

type AddGroupSilenceBlacklistRequestRequestParams struct {
	// 操作者
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 群会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 禁言用户列表
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// 禁言时长
	SilenceDuration *int64 `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
}

func (s AddGroupSilenceBlacklistRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceBlacklistRequestRequestParams) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceBlacklistRequestRequestParams) SetOperatorAppUid(v string) *AddGroupSilenceBlacklistRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *AddGroupSilenceBlacklistRequestRequestParams) SetAppCid(v string) *AddGroupSilenceBlacklistRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *AddGroupSilenceBlacklistRequestRequestParams) SetMembers(v []*string) *AddGroupSilenceBlacklistRequestRequestParams {
	s.Members = v
	return s
}

func (s *AddGroupSilenceBlacklistRequestRequestParams) SetSilenceDuration(v int64) *AddGroupSilenceBlacklistRequestRequestParams {
	s.SilenceDuration = &v
	return s
}

type AddGroupSilenceBlacklistShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言添加白名单请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s AddGroupSilenceBlacklistShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceBlacklistShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceBlacklistShrinkRequest) SetAppId(v string) *AddGroupSilenceBlacklistShrinkRequest {
	s.AppId = &v
	return s
}

func (s *AddGroupSilenceBlacklistShrinkRequest) SetRequestParamsShrink(v string) *AddGroupSilenceBlacklistShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type AddGroupSilenceBlacklistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddGroupSilenceBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceBlacklistResponseBody) SetRequestId(v string) *AddGroupSilenceBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGroupSilenceBlacklistResponseBody) SetCode(v string) *AddGroupSilenceBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *AddGroupSilenceBlacklistResponseBody) SetMessage(v string) *AddGroupSilenceBlacklistResponseBody {
	s.Message = &v
	return s
}

type AddGroupSilenceBlacklistResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGroupSilenceBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGroupSilenceBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceBlacklistResponse) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceBlacklistResponse) SetHeaders(v map[string]*string) *AddGroupSilenceBlacklistResponse {
	s.Headers = v
	return s
}

func (s *AddGroupSilenceBlacklistResponse) SetBody(v *AddGroupSilenceBlacklistResponseBody) *AddGroupSilenceBlacklistResponse {
	s.Body = v
	return s
}

type RemoveGroupSilenceWhitelistRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言添加白名单请求体
	RequestParams *RemoveGroupSilenceWhitelistRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RemoveGroupSilenceWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceWhitelistRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceWhitelistRequest) SetAppId(v string) *RemoveGroupSilenceWhitelistRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupSilenceWhitelistRequest) SetRequestParams(v *RemoveGroupSilenceWhitelistRequestRequestParams) *RemoveGroupSilenceWhitelistRequest {
	s.RequestParams = v
	return s
}

type RemoveGroupSilenceWhitelistRequestRequestParams struct {
	// 操作者
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 群会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 禁言用户列表
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s RemoveGroupSilenceWhitelistRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceWhitelistRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceWhitelistRequestRequestParams) SetOperatorAppUid(v string) *RemoveGroupSilenceWhitelistRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *RemoveGroupSilenceWhitelistRequestRequestParams) SetAppCid(v string) *RemoveGroupSilenceWhitelistRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RemoveGroupSilenceWhitelistRequestRequestParams) SetMembers(v []*string) *RemoveGroupSilenceWhitelistRequestRequestParams {
	s.Members = v
	return s
}

type RemoveGroupSilenceWhitelistShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言添加白名单请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RemoveGroupSilenceWhitelistShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceWhitelistShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceWhitelistShrinkRequest) SetAppId(v string) *RemoveGroupSilenceWhitelistShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupSilenceWhitelistShrinkRequest) SetRequestParamsShrink(v string) *RemoveGroupSilenceWhitelistShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RemoveGroupSilenceWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveGroupSilenceWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceWhitelistResponseBody) SetRequestId(v string) *RemoveGroupSilenceWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGroupSilenceWhitelistResponseBody) SetCode(v string) *RemoveGroupSilenceWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveGroupSilenceWhitelistResponseBody) SetMessage(v string) *RemoveGroupSilenceWhitelistResponseBody {
	s.Message = &v
	return s
}

type RemoveGroupSilenceWhitelistResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveGroupSilenceWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveGroupSilenceWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceWhitelistResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceWhitelistResponse) SetHeaders(v map[string]*string) *RemoveGroupSilenceWhitelistResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupSilenceWhitelistResponse) SetBody(v *RemoveGroupSilenceWhitelistResponseBody) *RemoveGroupSilenceWhitelistResponse {
	s.Body = v
	return s
}

type ListDetailReportStatisticsRequest struct {
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 请求
	RequestParams *ListDetailReportStatisticsRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s ListDetailReportStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDetailReportStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListDetailReportStatisticsRequest) SetAppId(v string) *ListDetailReportStatisticsRequest {
	s.AppId = &v
	return s
}

func (s *ListDetailReportStatisticsRequest) SetRequestParams(v *ListDetailReportStatisticsRequestRequestParams) *ListDetailReportStatisticsRequest {
	s.RequestParams = v
	return s
}

type ListDetailReportStatisticsRequestRequestParams struct {
	// 开始时间，utc
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间，utc
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 报表类型  user、groupChat、message
	ReportStatisticsType *string `json:"ReportStatisticsType,omitempty" xml:"ReportStatisticsType,omitempty"`
}

func (s ListDetailReportStatisticsRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s ListDetailReportStatisticsRequestRequestParams) GoString() string {
	return s.String()
}

func (s *ListDetailReportStatisticsRequestRequestParams) SetStartTime(v string) *ListDetailReportStatisticsRequestRequestParams {
	s.StartTime = &v
	return s
}

func (s *ListDetailReportStatisticsRequestRequestParams) SetEndTime(v string) *ListDetailReportStatisticsRequestRequestParams {
	s.EndTime = &v
	return s
}

func (s *ListDetailReportStatisticsRequestRequestParams) SetReportStatisticsType(v string) *ListDetailReportStatisticsRequestRequestParams {
	s.ReportStatisticsType = &v
	return s
}

type ListDetailReportStatisticsShrinkRequest struct {
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 请求
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s ListDetailReportStatisticsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDetailReportStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDetailReportStatisticsShrinkRequest) SetAppId(v string) *ListDetailReportStatisticsShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListDetailReportStatisticsShrinkRequest) SetRequestParamsShrink(v string) *ListDetailReportStatisticsShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type ListDetailReportStatisticsResponseBody struct {
	// desc
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// result
	Result *ListDetailReportStatisticsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListDetailReportStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDetailReportStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetailReportStatisticsResponseBody) SetMessage(v string) *ListDetailReportStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDetailReportStatisticsResponseBody) SetRequestId(v string) *ListDetailReportStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDetailReportStatisticsResponseBody) SetHttpStatusCode(v int32) *ListDetailReportStatisticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDetailReportStatisticsResponseBody) SetCode(v string) *ListDetailReportStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDetailReportStatisticsResponseBody) SetSuccess(v bool) *ListDetailReportStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDetailReportStatisticsResponseBody) SetResult(v *ListDetailReportStatisticsResponseBodyResult) *ListDetailReportStatisticsResponseBody {
	s.Result = v
	return s
}

type ListDetailReportStatisticsResponseBodyResult struct {
	// 数据
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListDetailReportStatisticsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDetailReportStatisticsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDetailReportStatisticsResponseBodyResult) SetData(v []map[string]interface{}) *ListDetailReportStatisticsResponseBodyResult {
	s.Data = v
	return s
}

type ListDetailReportStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDetailReportStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDetailReportStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDetailReportStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListDetailReportStatisticsResponse) SetHeaders(v map[string]*string) *ListDetailReportStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListDetailReportStatisticsResponse) SetBody(v *ListDetailReportStatisticsResponseBody) *ListDetailReportStatisticsResponse {
	s.Body = v
	return s
}

type SetUserConversationExtensionByKeysRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 设置用户拓展字段请求实体
	RequestParams *SetUserConversationExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SetUserConversationExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserConversationExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *SetUserConversationExtensionByKeysRequest) SetAppId(v string) *SetUserConversationExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *SetUserConversationExtensionByKeysRequest) SetRequestParams(v *SetUserConversationExtensionByKeysRequestRequestParams) *SetUserConversationExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type SetUserConversationExtensionByKeysRequestRequestParams struct {
	// 用户id
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 拓展字段
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s SetUserConversationExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SetUserConversationExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SetUserConversationExtensionByKeysRequestRequestParams) SetAppUid(v string) *SetUserConversationExtensionByKeysRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *SetUserConversationExtensionByKeysRequestRequestParams) SetAppCid(v string) *SetUserConversationExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *SetUserConversationExtensionByKeysRequestRequestParams) SetExtensions(v map[string]*string) *SetUserConversationExtensionByKeysRequestRequestParams {
	s.Extensions = v
	return s
}

type SetUserConversationExtensionByKeysShrinkRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 设置用户拓展字段请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SetUserConversationExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserConversationExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetUserConversationExtensionByKeysShrinkRequest) SetAppId(v string) *SetUserConversationExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SetUserConversationExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *SetUserConversationExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SetUserConversationExtensionByKeysResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SetUserConversationExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetUserConversationExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserConversationExtensionByKeysResponseBody) SetRequestId(v string) *SetUserConversationExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetUserConversationExtensionByKeysResponseBody) SetCode(v string) *SetUserConversationExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *SetUserConversationExtensionByKeysResponseBody) SetMessage(v string) *SetUserConversationExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type SetUserConversationExtensionByKeysResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetUserConversationExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetUserConversationExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserConversationExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *SetUserConversationExtensionByKeysResponse) SetHeaders(v map[string]*string) *SetUserConversationExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *SetUserConversationExtensionByKeysResponse) SetBody(v *SetUserConversationExtensionByKeysResponseBody) *SetUserConversationExtensionByKeysResponse {
	s.Body = v
	return s
}

type GetGroupByIdRequest struct {
	// APP ID, IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群会话信息获取的请求体
	RequestParams *GetGroupByIdRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s GetGroupByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupByIdRequest) GoString() string {
	return s.String()
}

func (s *GetGroupByIdRequest) SetAppId(v string) *GetGroupByIdRequest {
	s.AppId = &v
	return s
}

func (s *GetGroupByIdRequest) SetRequestParams(v *GetGroupByIdRequestRequestParams) *GetGroupByIdRequest {
	s.RequestParams = v
	return s
}

type GetGroupByIdRequestRequestParams struct {
	// 群会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
}

func (s GetGroupByIdRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetGroupByIdRequestRequestParams) GoString() string {
	return s.String()
}

func (s *GetGroupByIdRequestRequestParams) SetAppCid(v string) *GetGroupByIdRequestRequestParams {
	s.AppCid = &v
	return s
}

type GetGroupByIdShrinkRequest struct {
	// APP ID, IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群会话信息获取的请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s GetGroupByIdShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetGroupByIdShrinkRequest) SetAppId(v string) *GetGroupByIdShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetGroupByIdShrinkRequest) SetRequestParamsShrink(v string) *GetGroupByIdShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type GetGroupByIdResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 群信息获取的返回结果
	Result *GetGroupByIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetGroupByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupByIdResponseBody) SetRequestId(v string) *GetGroupByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupByIdResponseBody) SetCode(v string) *GetGroupByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetGroupByIdResponseBody) SetMessage(v string) *GetGroupByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetGroupByIdResponseBody) SetResult(v *GetGroupByIdResponseBodyResult) *GetGroupByIdResponseBody {
	s.Result = v
	return s
}

type GetGroupByIdResponseBodyResult struct {
	// 群会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 群主ID
	OwnerAppUid *string `json:"OwnerAppUid,omitempty" xml:"OwnerAppUid,omitempty"`
	// 群图像
	IconMediaId *string `json:"IconMediaId,omitempty" xml:"IconMediaId,omitempty"`
	// 群名称
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 当前群人数
	MemberCount *int32 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	// 群人数上限
	MemberLimit *int32 `json:"MemberLimit,omitempty" xml:"MemberLimit,omitempty"`
	// 群扩展信息
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// 群创建时间
	CeateTime *int64 `json:"CeateTime,omitempty" xml:"CeateTime,omitempty"`
}

func (s GetGroupByIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetGroupByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupByIdResponseBodyResult) SetAppCid(v string) *GetGroupByIdResponseBodyResult {
	s.AppCid = &v
	return s
}

func (s *GetGroupByIdResponseBodyResult) SetOwnerAppUid(v string) *GetGroupByIdResponseBodyResult {
	s.OwnerAppUid = &v
	return s
}

func (s *GetGroupByIdResponseBodyResult) SetIconMediaId(v string) *GetGroupByIdResponseBodyResult {
	s.IconMediaId = &v
	return s
}

func (s *GetGroupByIdResponseBodyResult) SetTitle(v string) *GetGroupByIdResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetGroupByIdResponseBodyResult) SetMemberCount(v int32) *GetGroupByIdResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *GetGroupByIdResponseBodyResult) SetMemberLimit(v int32) *GetGroupByIdResponseBodyResult {
	s.MemberLimit = &v
	return s
}

func (s *GetGroupByIdResponseBodyResult) SetExtensions(v map[string]*string) *GetGroupByIdResponseBodyResult {
	s.Extensions = v
	return s
}

func (s *GetGroupByIdResponseBodyResult) SetCeateTime(v int64) *GetGroupByIdResponseBodyResult {
	s.CeateTime = &v
	return s
}

type GetGroupByIdResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGroupByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGroupByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupByIdResponse) GoString() string {
	return s.String()
}

func (s *GetGroupByIdResponse) SetHeaders(v map[string]*string) *GetGroupByIdResponse {
	s.Headers = v
	return s
}

func (s *GetGroupByIdResponse) SetBody(v *GetGroupByIdResponseBody) *GetGroupByIdResponse {
	s.Body = v
	return s
}

type UpdateTenantStatusRequest struct {
	Request *UpdateTenantStatusRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s UpdateTenantStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantStatusRequest) SetRequest(v *UpdateTenantStatusRequestRequest) *UpdateTenantStatusRequest {
	s.Request = v
	return s
}

type UpdateTenantStatusRequestRequest struct {
	// 应用appKey
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// 应用状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateTenantStatusRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantStatusRequestRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantStatusRequestRequest) SetDomain(v string) *UpdateTenantStatusRequestRequest {
	s.Domain = &v
	return s
}

func (s *UpdateTenantStatusRequestRequest) SetStatus(v int64) *UpdateTenantStatusRequestRequest {
	s.Status = &v
	return s
}

type UpdateTenantStatusResponseBody struct {
	// Id of the request
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 是否更新成功
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTenantStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTenantStatusResponseBody) SetResponseSuccess(v bool) *UpdateTenantStatusResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *UpdateTenantStatusResponseBody) SetErrorCode(v string) *UpdateTenantStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTenantStatusResponseBody) SetErrorMsg(v string) *UpdateTenantStatusResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateTenantStatusResponseBody) SetResult(v bool) *UpdateTenantStatusResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateTenantStatusResponseBody) SetRequestId(v string) *UpdateTenantStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTenantStatusResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTenantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTenantStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateTenantStatusResponse) SetHeaders(v map[string]*string) *UpdateTenantStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateTenantStatusResponse) SetBody(v *UpdateTenantStatusResponseBody) *UpdateTenantStatusResponse {
	s.Body = v
	return s
}

type GetCommonConfigRequest struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetCommonConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCommonConfigRequest) GoString() string {
	return s.String()
}

func (s *GetCommonConfigRequest) SetAppId(v string) *GetCommonConfigRequest {
	s.AppId = &v
	return s
}

type GetCommonConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回值
	Result *GetCommonConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetCommonConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCommonConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommonConfigResponseBody) SetRequestId(v string) *GetCommonConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommonConfigResponseBody) SetSuccess(v bool) *GetCommonConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetCommonConfigResponseBody) SetCode(v string) *GetCommonConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetCommonConfigResponseBody) SetMessage(v string) *GetCommonConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetCommonConfigResponseBody) SetResult(v *GetCommonConfigResponseBodyResult) *GetCommonConfigResponseBody {
	s.Result = v
	return s
}

type GetCommonConfigResponseBodyResult struct {
	// 通用配置
	CommonConfig *GetCommonConfigResponseBodyResultCommonConfig `json:"CommonConfig,omitempty" xml:"CommonConfig,omitempty" type:"Struct"`
}

func (s GetCommonConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCommonConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCommonConfigResponseBodyResult) SetCommonConfig(v *GetCommonConfigResponseBodyResultCommonConfig) *GetCommonConfigResponseBodyResult {
	s.CommonConfig = v
	return s
}

type GetCommonConfigResponseBodyResultCommonConfig struct {
	// 登录配置
	LoginConfig *GetCommonConfigResponseBodyResultCommonConfigLoginConfig `json:"LoginConfig,omitempty" xml:"LoginConfig,omitempty" type:"Struct"`
	// app配置
	AppConfigs []*GetCommonConfigResponseBodyResultCommonConfigAppConfigs `json:"AppConfigs,omitempty" xml:"AppConfigs,omitempty" type:"Repeated"`
}

func (s GetCommonConfigResponseBodyResultCommonConfig) String() string {
	return tea.Prettify(s)
}

func (s GetCommonConfigResponseBodyResultCommonConfig) GoString() string {
	return s.String()
}

func (s *GetCommonConfigResponseBodyResultCommonConfig) SetLoginConfig(v *GetCommonConfigResponseBodyResultCommonConfigLoginConfig) *GetCommonConfigResponseBodyResultCommonConfig {
	s.LoginConfig = v
	return s
}

func (s *GetCommonConfigResponseBodyResultCommonConfig) SetAppConfigs(v []*GetCommonConfigResponseBodyResultCommonConfigAppConfigs) *GetCommonConfigResponseBodyResultCommonConfig {
	s.AppConfigs = v
	return s
}

type GetCommonConfigResponseBodyResultCommonConfigLoginConfig struct {
	// 登录类型
	LoginType *int32 `json:"LoginType,omitempty" xml:"LoginType,omitempty"`
}

func (s GetCommonConfigResponseBodyResultCommonConfigLoginConfig) String() string {
	return tea.Prettify(s)
}

func (s GetCommonConfigResponseBodyResultCommonConfigLoginConfig) GoString() string {
	return s.String()
}

func (s *GetCommonConfigResponseBodyResultCommonConfigLoginConfig) SetLoginType(v int32) *GetCommonConfigResponseBodyResultCommonConfigLoginConfig {
	s.LoginType = &v
	return s
}

type GetCommonConfigResponseBodyResultCommonConfigAppConfigs struct {
	// appKey
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 平台
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s GetCommonConfigResponseBodyResultCommonConfigAppConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetCommonConfigResponseBodyResultCommonConfigAppConfigs) GoString() string {
	return s.String()
}

func (s *GetCommonConfigResponseBodyResultCommonConfigAppConfigs) SetAppKey(v string) *GetCommonConfigResponseBodyResultCommonConfigAppConfigs {
	s.AppKey = &v
	return s
}

func (s *GetCommonConfigResponseBodyResultCommonConfigAppConfigs) SetPlatform(v string) *GetCommonConfigResponseBodyResultCommonConfigAppConfigs {
	s.Platform = &v
	return s
}

type GetCommonConfigResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCommonConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCommonConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCommonConfigResponse) GoString() string {
	return s.String()
}

func (s *GetCommonConfigResponse) SetHeaders(v map[string]*string) *GetCommonConfigResponse {
	s.Headers = v
	return s
}

func (s *GetCommonConfigResponse) SetBody(v *GetCommonConfigResponseBody) *GetCommonConfigResponse {
	s.Body = v
	return s
}

type SendMessageRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息发送请求体
	RequestParams *SendMessageRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetAppId(v string) *SendMessageRequest {
	s.AppId = &v
	return s
}

func (s *SendMessageRequest) SetRequestParams(v *SendMessageRequestRequestParams) *SendMessageRequest {
	s.RequestParams = v
	return s
}

type SendMessageRequestRequestParams struct {
	// 消息UUID
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 会话类型
	ConversationType *int32 `json:"ConversationType,omitempty" xml:"ConversationType,omitempty"`
	// 发送者UID
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 消息内容类型
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// 消息内容Json
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 消息扩展字段
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// 消息设置
	Options *SendMessageRequestRequestParamsOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
}

func (s SendMessageRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SendMessageRequestRequestParams) SetUuid(v string) *SendMessageRequestRequestParams {
	s.Uuid = &v
	return s
}

func (s *SendMessageRequestRequestParams) SetAppCid(v string) *SendMessageRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *SendMessageRequestRequestParams) SetConversationType(v int32) *SendMessageRequestRequestParams {
	s.ConversationType = &v
	return s
}

func (s *SendMessageRequestRequestParams) SetSenderId(v string) *SendMessageRequestRequestParams {
	s.SenderId = &v
	return s
}

func (s *SendMessageRequestRequestParams) SetContentType(v int32) *SendMessageRequestRequestParams {
	s.ContentType = &v
	return s
}

func (s *SendMessageRequestRequestParams) SetContent(v string) *SendMessageRequestRequestParams {
	s.Content = &v
	return s
}

func (s *SendMessageRequestRequestParams) SetExtensions(v map[string]*string) *SendMessageRequestRequestParams {
	s.Extensions = v
	return s
}

func (s *SendMessageRequestRequestParams) SetOptions(v *SendMessageRequestRequestParamsOptions) *SendMessageRequestRequestParams {
	s.Options = v
	return s
}

type SendMessageRequestRequestParamsOptions struct {
	// 未读消息小红点控制。0:增加小红点; 1:不增加小红点
	RedPointPolicy *int32 `json:"RedPointPolicy,omitempty" xml:"RedPointPolicy,omitempty"`
	// 接受相关设置
	ReceiveScopeOption *SendMessageRequestRequestParamsOptionsReceiveScopeOption `json:"ReceiveScopeOption,omitempty" xml:"ReceiveScopeOption,omitempty" type:"Struct"`
	// 单聊会话不存在时新建自定义单聊请求体
	SingleChatCreateRequest *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest `json:"SingleChatCreateRequest,omitempty" xml:"SingleChatCreateRequest,omitempty" type:"Struct"`
}

func (s SendMessageRequestRequestParamsOptions) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequestRequestParamsOptions) GoString() string {
	return s.String()
}

func (s *SendMessageRequestRequestParamsOptions) SetRedPointPolicy(v int32) *SendMessageRequestRequestParamsOptions {
	s.RedPointPolicy = &v
	return s
}

func (s *SendMessageRequestRequestParamsOptions) SetReceiveScopeOption(v *SendMessageRequestRequestParamsOptionsReceiveScopeOption) *SendMessageRequestRequestParamsOptions {
	s.ReceiveScopeOption = v
	return s
}

func (s *SendMessageRequestRequestParamsOptions) SetSingleChatCreateRequest(v *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest) *SendMessageRequestRequestParamsOptions {
	s.SingleChatCreateRequest = v
	return s
}

type SendMessageRequestRequestParamsOptionsReceiveScopeOption struct {
	// 接受者列表
	ReceiverIds []*string `json:"ReceiverIds,omitempty" xml:"ReceiverIds,omitempty" type:"Repeated"`
	// 不接收者列表
	ExcludeReceiverIds []*string `json:"ExcludeReceiverIds,omitempty" xml:"ExcludeReceiverIds,omitempty" type:"Repeated"`
	// 消息获取控制。0: 会话内除指定ExcludeReceivers均可获取；1: 会话内仅指定ReceiverIds可获取
	ReceiveScope *int32 `json:"ReceiveScope,omitempty" xml:"ReceiveScope,omitempty"`
}

func (s SendMessageRequestRequestParamsOptionsReceiveScopeOption) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequestRequestParamsOptionsReceiveScopeOption) GoString() string {
	return s.String()
}

func (s *SendMessageRequestRequestParamsOptionsReceiveScopeOption) SetReceiverIds(v []*string) *SendMessageRequestRequestParamsOptionsReceiveScopeOption {
	s.ReceiverIds = v
	return s
}

func (s *SendMessageRequestRequestParamsOptionsReceiveScopeOption) SetExcludeReceiverIds(v []*string) *SendMessageRequestRequestParamsOptionsReceiveScopeOption {
	s.ExcludeReceiverIds = v
	return s
}

func (s *SendMessageRequestRequestParamsOptionsReceiveScopeOption) SetReceiveScope(v int32) *SendMessageRequestRequestParamsOptionsReceiveScopeOption {
	s.ReceiveScope = &v
	return s
}

type SendMessageRequestRequestParamsOptionsSingleChatCreateRequest struct {
	// 单聊会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 用户ID列表
	AppUids []*string `json:"AppUids,omitempty" xml:"AppUids,omitempty" type:"Repeated"`
	// 扩展信息
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// 用户会话视图信息
	UserConversation map[string]*RequestParamsOptionsSingleChatCreateRequestUserConversationValue `json:"UserConversation,omitempty" xml:"UserConversation,omitempty"`
}

func (s SendMessageRequestRequestParamsOptionsSingleChatCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequestRequestParamsOptionsSingleChatCreateRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest) SetAppCid(v string) *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest {
	s.AppCid = &v
	return s
}

func (s *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest) SetAppUids(v []*string) *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest {
	s.AppUids = v
	return s
}

func (s *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest) SetExtensions(v map[string]*string) *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest {
	s.Extensions = v
	return s
}

func (s *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest) SetUserConversation(v map[string]*RequestParamsOptionsSingleChatCreateRequestUserConversationValue) *SendMessageRequestRequestParamsOptionsSingleChatCreateRequest {
	s.UserConversation = v
	return s
}

type SendMessageShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息发送请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SendMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendMessageShrinkRequest) SetAppId(v string) *SendMessageShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SendMessageShrinkRequest) SetRequestParamsShrink(v string) *SendMessageShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SendMessageResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *SendMessageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageResponseBody) SetCode(v string) *SendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageResponseBody) SetMessage(v string) *SendMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendMessageResponseBody) SetResult(v *SendMessageResponseBodyResult) *SendMessageResponseBody {
	s.Result = v
	return s
}

type SendMessageResponseBodyResult struct {
	// 消息ID
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// 消息创建时间戳(毫秒)
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s SendMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBodyResult) SetMsgId(v string) *SendMessageResponseBodyResult {
	s.MsgId = &v
	return s
}

func (s *SendMessageResponseBodyResult) SetCreateTime(v int64) *SendMessageResponseBodyResult {
	s.CreateTime = &v
	return s
}

type SendMessageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type UpdateGroupMembersRoleRequest struct {
	// App ID。IMPaaS租户的ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 更新群成员角色请求体。
	RequestParams *UpdateGroupMembersRoleRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s UpdateGroupMembersRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMembersRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupMembersRoleRequest) SetAppId(v string) *UpdateGroupMembersRoleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGroupMembersRoleRequest) SetRequestParams(v *UpdateGroupMembersRoleRequestRequestParams) *UpdateGroupMembersRoleRequest {
	s.RequestParams = v
	return s
}

type UpdateGroupMembersRoleRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 操作用户ID。
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 更新后的成员角色。取值： 2：管理员。 3：普通。 100~127：自定义。 不能为1。
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 需要更改的uids
	AppUids []*string `json:"AppUids,omitempty" xml:"AppUids,omitempty" type:"Repeated"`
}

func (s UpdateGroupMembersRoleRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMembersRoleRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateGroupMembersRoleRequestRequestParams) SetAppCid(v string) *UpdateGroupMembersRoleRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *UpdateGroupMembersRoleRequestRequestParams) SetOperatorAppUid(v string) *UpdateGroupMembersRoleRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *UpdateGroupMembersRoleRequestRequestParams) SetRole(v int32) *UpdateGroupMembersRoleRequestRequestParams {
	s.Role = &v
	return s
}

func (s *UpdateGroupMembersRoleRequestRequestParams) SetAppUids(v []*string) *UpdateGroupMembersRoleRequestRequestParams {
	s.AppUids = v
	return s
}

type UpdateGroupMembersRoleShrinkRequest struct {
	// App ID。IMPaaS租户的ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 更新群成员角色请求体。
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s UpdateGroupMembersRoleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMembersRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupMembersRoleShrinkRequest) SetAppId(v string) *UpdateGroupMembersRoleShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGroupMembersRoleShrinkRequest) SetRequestParamsShrink(v string) *UpdateGroupMembersRoleShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type UpdateGroupMembersRoleResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码。
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息。
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateGroupMembersRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMembersRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupMembersRoleResponseBody) SetRequestId(v string) *UpdateGroupMembersRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupMembersRoleResponseBody) SetCode(v string) *UpdateGroupMembersRoleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupMembersRoleResponseBody) SetMessage(v string) *UpdateGroupMembersRoleResponseBody {
	s.Message = &v
	return s
}

type UpdateGroupMembersRoleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupMembersRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupMembersRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMembersRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupMembersRoleResponse) SetHeaders(v map[string]*string) *UpdateGroupMembersRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupMembersRoleResponse) SetBody(v *UpdateGroupMembersRoleResponseBody) *UpdateGroupMembersRoleResponse {
	s.Body = v
	return s
}

type CancelSilenceAllGroupMembersRequest struct {
	// AppId
	AppId         *string                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *CancelSilenceAllGroupMembersRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s CancelSilenceAllGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelSilenceAllGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *CancelSilenceAllGroupMembersRequest) SetAppId(v string) *CancelSilenceAllGroupMembersRequest {
	s.AppId = &v
	return s
}

func (s *CancelSilenceAllGroupMembersRequest) SetRequestParams(v *CancelSilenceAllGroupMembersRequestRequestParams) *CancelSilenceAllGroupMembersRequest {
	s.RequestParams = v
	return s
}

type CancelSilenceAllGroupMembersRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 操作者uid
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
}

func (s CancelSilenceAllGroupMembersRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s CancelSilenceAllGroupMembersRequestRequestParams) GoString() string {
	return s.String()
}

func (s *CancelSilenceAllGroupMembersRequestRequestParams) SetAppCid(v string) *CancelSilenceAllGroupMembersRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *CancelSilenceAllGroupMembersRequestRequestParams) SetOperatorAppUid(v string) *CancelSilenceAllGroupMembersRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

type CancelSilenceAllGroupMembersShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s CancelSilenceAllGroupMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelSilenceAllGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *CancelSilenceAllGroupMembersShrinkRequest) SetAppId(v string) *CancelSilenceAllGroupMembersShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CancelSilenceAllGroupMembersShrinkRequest) SetRequestParamsShrink(v string) *CancelSilenceAllGroupMembersShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type CancelSilenceAllGroupMembersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CancelSilenceAllGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelSilenceAllGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSilenceAllGroupMembersResponseBody) SetRequestId(v string) *CancelSilenceAllGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSilenceAllGroupMembersResponseBody) SetCode(v string) *CancelSilenceAllGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *CancelSilenceAllGroupMembersResponseBody) SetMessage(v string) *CancelSilenceAllGroupMembersResponseBody {
	s.Message = &v
	return s
}

type CancelSilenceAllGroupMembersResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelSilenceAllGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelSilenceAllGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelSilenceAllGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *CancelSilenceAllGroupMembersResponse) SetHeaders(v map[string]*string) *CancelSilenceAllGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *CancelSilenceAllGroupMembersResponse) SetBody(v *CancelSilenceAllGroupMembersResponseBody) *CancelSilenceAllGroupMembersResponse {
	s.Body = v
	return s
}

type UpdateGroupIconRequest struct {
	// AppId
	AppId         *string                              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *UpdateGroupIconRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s UpdateGroupIconRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupIconRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupIconRequest) SetAppId(v string) *UpdateGroupIconRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGroupIconRequest) SetRequestParams(v *UpdateGroupIconRequestRequestParams) *UpdateGroupIconRequest {
	s.RequestParams = v
	return s
}

type UpdateGroupIconRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 操作者用户ID
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 群聊头像文件MediaID
	IconMediaId *string `json:"IconMediaId,omitempty" xml:"IconMediaId,omitempty"`
}

func (s UpdateGroupIconRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupIconRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateGroupIconRequestRequestParams) SetAppCid(v string) *UpdateGroupIconRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *UpdateGroupIconRequestRequestParams) SetOperatorAppUid(v string) *UpdateGroupIconRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *UpdateGroupIconRequestRequestParams) SetIconMediaId(v string) *UpdateGroupIconRequestRequestParams {
	s.IconMediaId = &v
	return s
}

type UpdateGroupIconShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s UpdateGroupIconShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupIconShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupIconShrinkRequest) SetAppId(v string) *UpdateGroupIconShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGroupIconShrinkRequest) SetRequestParamsShrink(v string) *UpdateGroupIconShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type UpdateGroupIconResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateGroupIconResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupIconResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupIconResponseBody) SetRequestId(v string) *UpdateGroupIconResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupIconResponseBody) SetCode(v string) *UpdateGroupIconResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupIconResponseBody) SetMessage(v string) *UpdateGroupIconResponseBody {
	s.Message = &v
	return s
}

type UpdateGroupIconResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupIconResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupIconResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupIconResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupIconResponse) SetHeaders(v map[string]*string) *UpdateGroupIconResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupIconResponse) SetBody(v *UpdateGroupIconResponseBody) *UpdateGroupIconResponse {
	s.Body = v
	return s
}

type RemoveGroupMembersRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群踢人请求实体
	RequestParams *RemoveGroupMembersRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RemoveGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersRequest) SetAppId(v string) *RemoveGroupMembersRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupMembersRequest) SetRequestParams(v *RemoveGroupMembersRequestRequestParams) *RemoveGroupMembersRequest {
	s.RequestParams = v
	return s
}

type RemoveGroupMembersRequestRequestParams struct {
	OperatorAppUid *string   `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	AppCid         *string   `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	AppUidsRemoved []*string `json:"AppUidsRemoved,omitempty" xml:"AppUidsRemoved,omitempty" type:"Repeated"`
}

func (s RemoveGroupMembersRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersRequestRequestParams) SetOperatorAppUid(v string) *RemoveGroupMembersRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *RemoveGroupMembersRequestRequestParams) SetAppCid(v string) *RemoveGroupMembersRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RemoveGroupMembersRequestRequestParams) SetAppUidsRemoved(v []*string) *RemoveGroupMembersRequestRequestParams {
	s.AppUidsRemoved = v
	return s
}

type RemoveGroupMembersShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群踢人请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RemoveGroupMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersShrinkRequest) SetAppId(v string) *RemoveGroupMembersShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupMembersShrinkRequest) SetRequestParamsShrink(v string) *RemoveGroupMembersShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RemoveGroupMembersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersResponseBody) SetRequestId(v string) *RemoveGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGroupMembersResponseBody) SetCode(v string) *RemoveGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveGroupMembersResponseBody) SetMessage(v string) *RemoveGroupMembersResponseBody {
	s.Message = &v
	return s
}

type RemoveGroupMembersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupMembersResponse) SetHeaders(v map[string]*string) *RemoveGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupMembersResponse) SetBody(v *RemoveGroupMembersResponseBody) *RemoveGroupMembersResponse {
	s.Body = v
	return s
}

type ListGroupAllMembersRequest struct {
	// App ID, IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拉取群成员列表的请求体
	RequestParams *ListGroupAllMembersRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s ListGroupAllMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAllMembersRequest) GoString() string {
	return s.String()
}

func (s *ListGroupAllMembersRequest) SetAppId(v string) *ListGroupAllMembersRequest {
	s.AppId = &v
	return s
}

func (s *ListGroupAllMembersRequest) SetRequestParams(v *ListGroupAllMembersRequestRequestParams) *ListGroupAllMembersRequest {
	s.RequestParams = v
	return s
}

type ListGroupAllMembersRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
}

func (s ListGroupAllMembersRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAllMembersRequestRequestParams) GoString() string {
	return s.String()
}

func (s *ListGroupAllMembersRequestRequestParams) SetAppCid(v string) *ListGroupAllMembersRequestRequestParams {
	s.AppCid = &v
	return s
}

type ListGroupAllMembersShrinkRequest struct {
	// App ID, IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拉取群成员列表的请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s ListGroupAllMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAllMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGroupAllMembersShrinkRequest) SetAppId(v string) *ListGroupAllMembersShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListGroupAllMembersShrinkRequest) SetRequestParamsShrink(v string) *ListGroupAllMembersShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type ListGroupAllMembersResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 拉取群成员列表的结果
	Result *ListGroupAllMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListGroupAllMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAllMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupAllMembersResponseBody) SetRequestId(v string) *ListGroupAllMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupAllMembersResponseBody) SetCode(v string) *ListGroupAllMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupAllMembersResponseBody) SetMessage(v string) *ListGroupAllMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupAllMembersResponseBody) SetResult(v *ListGroupAllMembersResponseBodyResult) *ListGroupAllMembersResponseBody {
	s.Result = v
	return s
}

type ListGroupAllMembersResponseBodyResult struct {
	// 群成员列表
	Members []*ListGroupAllMembersResponseBodyResultMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s ListGroupAllMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAllMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupAllMembersResponseBodyResult) SetMembers(v []*ListGroupAllMembersResponseBodyResultMembers) *ListGroupAllMembersResponseBodyResult {
	s.Members = v
	return s
}

type ListGroupAllMembersResponseBodyResultMembers struct {
	// 群成员ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 群成员角色
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 群成员昵称
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// 群成员入群时间
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// 群成员扩展信息
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s ListGroupAllMembersResponseBodyResultMembers) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAllMembersResponseBodyResultMembers) GoString() string {
	return s.String()
}

func (s *ListGroupAllMembersResponseBodyResultMembers) SetAppUid(v string) *ListGroupAllMembersResponseBodyResultMembers {
	s.AppUid = &v
	return s
}

func (s *ListGroupAllMembersResponseBodyResultMembers) SetRole(v int32) *ListGroupAllMembersResponseBodyResultMembers {
	s.Role = &v
	return s
}

func (s *ListGroupAllMembersResponseBodyResultMembers) SetNick(v string) *ListGroupAllMembersResponseBodyResultMembers {
	s.Nick = &v
	return s
}

func (s *ListGroupAllMembersResponseBodyResultMembers) SetJoinTime(v int64) *ListGroupAllMembersResponseBodyResultMembers {
	s.JoinTime = &v
	return s
}

func (s *ListGroupAllMembersResponseBodyResultMembers) SetExtensions(v map[string]*string) *ListGroupAllMembersResponseBodyResultMembers {
	s.Extensions = v
	return s
}

type ListGroupAllMembersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupAllMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupAllMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupAllMembersResponse) GoString() string {
	return s.String()
}

func (s *ListGroupAllMembersResponse) SetHeaders(v map[string]*string) *ListGroupAllMembersResponse {
	s.Headers = v
	return s
}

func (s *ListGroupAllMembersResponse) SetBody(v *ListGroupAllMembersResponseBody) *ListGroupAllMembersResponse {
	s.Body = v
	return s
}

type GetUserMuteSettingRequest struct {
	// AppId
	AppId         *string                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *GetUserMuteSettingRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s GetUserMuteSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserMuteSettingRequest) GoString() string {
	return s.String()
}

func (s *GetUserMuteSettingRequest) SetAppId(v string) *GetUserMuteSettingRequest {
	s.AppId = &v
	return s
}

func (s *GetUserMuteSettingRequest) SetRequestParams(v *GetUserMuteSettingRequestRequestParams) *GetUserMuteSettingRequest {
	s.RequestParams = v
	return s
}

type GetUserMuteSettingRequestRequestParams struct {
	// 用户列表
	AppUids []*string `json:"AppUids,omitempty" xml:"AppUids,omitempty" type:"Repeated"`
}

func (s GetUserMuteSettingRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetUserMuteSettingRequestRequestParams) GoString() string {
	return s.String()
}

func (s *GetUserMuteSettingRequestRequestParams) SetAppUids(v []*string) *GetUserMuteSettingRequestRequestParams {
	s.AppUids = v
	return s
}

type GetUserMuteSettingShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s GetUserMuteSettingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserMuteSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserMuteSettingShrinkRequest) SetAppId(v string) *GetUserMuteSettingShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetUserMuteSettingShrinkRequest) SetRequestParamsShrink(v string) *GetUserMuteSettingShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type GetUserMuteSettingResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回值
	Result *GetUserMuteSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetUserMuteSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserMuteSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMuteSettingResponseBody) SetRequestId(v string) *GetUserMuteSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserMuteSettingResponseBody) SetCode(v string) *GetUserMuteSettingResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserMuteSettingResponseBody) SetMessage(v string) *GetUserMuteSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserMuteSettingResponseBody) SetResult(v *GetUserMuteSettingResponseBodyResult) *GetUserMuteSettingResponseBody {
	s.Result = v
	return s
}

type GetUserMuteSettingResponseBodyResult struct {
	UserMuteSettings map[string]*ResultUserMuteSettingsValue `json:"UserMuteSettings,omitempty" xml:"UserMuteSettings,omitempty"`
}

func (s GetUserMuteSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserMuteSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserMuteSettingResponseBodyResult) SetUserMuteSettings(v map[string]*ResultUserMuteSettingsValue) *GetUserMuteSettingResponseBodyResult {
	s.UserMuteSettings = v
	return s
}

type GetUserMuteSettingResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserMuteSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserMuteSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserMuteSettingResponse) GoString() string {
	return s.String()
}

func (s *GetUserMuteSettingResponse) SetHeaders(v map[string]*string) *GetUserMuteSettingResponse {
	s.Headers = v
	return s
}

func (s *GetUserMuteSettingResponse) SetBody(v *GetUserMuteSettingResponseBody) *GetUserMuteSettingResponse {
	s.Body = v
	return s
}

type GetRoomStatisticsRequest struct {
	// 请求参数的结构体。
	Request *GetRoomStatisticsRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s GetRoomStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoomStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetRoomStatisticsRequest) SetRequest(v *GetRoomStatisticsRequestRequest) *GetRoomStatisticsRequest {
	s.Request = v
	return s
}

type GetRoomStatisticsRequestRequest struct {
	// 应用的appKey。
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 房间ID，由调用CreateRoom时返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s GetRoomStatisticsRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoomStatisticsRequestRequest) GoString() string {
	return s.String()
}

func (s *GetRoomStatisticsRequestRequest) SetDomain(v string) *GetRoomStatisticsRequestRequest {
	s.Domain = &v
	return s
}

func (s *GetRoomStatisticsRequestRequest) SetRoomId(v string) *GetRoomStatisticsRequestRequest {
	s.RoomId = &v
	return s
}

type GetRoomStatisticsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功。
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码，请求异常时返回。
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息，请求异常时返回。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求的返回结果。
	Result *GetRoomStatisticsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRoomStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoomStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoomStatisticsResponseBody) SetRequestId(v string) *GetRoomStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoomStatisticsResponseBody) SetResponseSuccess(v bool) *GetRoomStatisticsResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *GetRoomStatisticsResponseBody) SetErrorCode(v string) *GetRoomStatisticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRoomStatisticsResponseBody) SetErrorMessage(v string) *GetRoomStatisticsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRoomStatisticsResponseBody) SetResult(v *GetRoomStatisticsResponseBodyResult) *GetRoomStatisticsResponseBody {
	s.Result = v
	return s
}

type GetRoomStatisticsResponseBodyResult struct {
	// 当前房间的在线观众数。
	OnlineCount *int32 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 当前房间的历史用户访问数目（UV）。
	UV *int32 `json:"UV,omitempty" xml:"UV,omitempty"`
	// 当前房间的历史访问数目（PV）。
	PV *int32 `json:"PV,omitempty" xml:"PV,omitempty"`
}

func (s GetRoomStatisticsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRoomStatisticsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRoomStatisticsResponseBodyResult) SetOnlineCount(v int32) *GetRoomStatisticsResponseBodyResult {
	s.OnlineCount = &v
	return s
}

func (s *GetRoomStatisticsResponseBodyResult) SetUV(v int32) *GetRoomStatisticsResponseBodyResult {
	s.UV = &v
	return s
}

func (s *GetRoomStatisticsResponseBodyResult) SetPV(v int32) *GetRoomStatisticsResponseBodyResult {
	s.PV = &v
	return s
}

type GetRoomStatisticsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRoomStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoomStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoomStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetRoomStatisticsResponse) SetHeaders(v map[string]*string) *GetRoomStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetRoomStatisticsResponse) SetBody(v *GetRoomStatisticsResponseBody) *GetRoomStatisticsResponse {
	s.Body = v
	return s
}

type AddGroupMembersRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群加人请求实体
	RequestParams *AddGroupMembersRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s AddGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMembersRequest) SetAppId(v string) *AddGroupMembersRequest {
	s.AppId = &v
	return s
}

func (s *AddGroupMembersRequest) SetRequestParams(v *AddGroupMembersRequestRequestParams) *AddGroupMembersRequest {
	s.RequestParams = v
	return s
}

type AddGroupMembersRequestRequestParams struct {
	// 操作者
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 初始化成员
	InitMembers []*AddGroupMembersRequestRequestParamsInitMembers `json:"InitMembers,omitempty" xml:"InitMembers,omitempty" type:"Repeated"`
}

func (s AddGroupMembersRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersRequestRequestParams) GoString() string {
	return s.String()
}

func (s *AddGroupMembersRequestRequestParams) SetOperatorAppUid(v string) *AddGroupMembersRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *AddGroupMembersRequestRequestParams) SetAppCid(v string) *AddGroupMembersRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *AddGroupMembersRequestRequestParams) SetInitMembers(v []*AddGroupMembersRequestRequestParamsInitMembers) *AddGroupMembersRequestRequestParams {
	s.InitMembers = v
	return s
}

type AddGroupMembersRequestRequestParamsInitMembers struct {
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 1群主，2管理员，3普通
	Role *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// unix毫秒数
	JoinTime   *int64             `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s AddGroupMembersRequestRequestParamsInitMembers) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersRequestRequestParamsInitMembers) GoString() string {
	return s.String()
}

func (s *AddGroupMembersRequestRequestParamsInitMembers) SetAppUid(v string) *AddGroupMembersRequestRequestParamsInitMembers {
	s.AppUid = &v
	return s
}

func (s *AddGroupMembersRequestRequestParamsInitMembers) SetRole(v int32) *AddGroupMembersRequestRequestParamsInitMembers {
	s.Role = &v
	return s
}

func (s *AddGroupMembersRequestRequestParamsInitMembers) SetNick(v string) *AddGroupMembersRequestRequestParamsInitMembers {
	s.Nick = &v
	return s
}

func (s *AddGroupMembersRequestRequestParamsInitMembers) SetJoinTime(v int64) *AddGroupMembersRequestRequestParamsInitMembers {
	s.JoinTime = &v
	return s
}

func (s *AddGroupMembersRequestRequestParamsInitMembers) SetExtensions(v map[string]*string) *AddGroupMembersRequestRequestParamsInitMembers {
	s.Extensions = v
	return s
}

type AddGroupMembersShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群加人请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s AddGroupMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMembersShrinkRequest) SetAppId(v string) *AddGroupMembersShrinkRequest {
	s.AppId = &v
	return s
}

func (s *AddGroupMembersShrinkRequest) SetRequestParamsShrink(v string) *AddGroupMembersShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type AddGroupMembersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupMembersResponseBody) SetRequestId(v string) *AddGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGroupMembersResponseBody) SetCode(v string) *AddGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *AddGroupMembersResponseBody) SetMessage(v string) *AddGroupMembersResponseBody {
	s.Message = &v
	return s
}

type AddGroupMembersResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *AddGroupMembersResponse) SetHeaders(v map[string]*string) *AddGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *AddGroupMembersResponse) SetBody(v *AddGroupMembersResponseBody) *AddGroupMembersResponse {
	s.Body = v
	return s
}

type GetGroupMemberByIdsRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群聊设置扩展字段请求实体
	RequestParams *GetGroupMemberByIdsRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s GetGroupMemberByIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupMemberByIdsRequest) GoString() string {
	return s.String()
}

func (s *GetGroupMemberByIdsRequest) SetAppId(v string) *GetGroupMemberByIdsRequest {
	s.AppId = &v
	return s
}

func (s *GetGroupMemberByIdsRequest) SetRequestParams(v *GetGroupMemberByIdsRequestRequestParams) *GetGroupMemberByIdsRequest {
	s.RequestParams = v
	return s
}

type GetGroupMemberByIdsRequestRequestParams struct {
	// 会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// appUid
	AppUids []*string `json:"AppUids,omitempty" xml:"AppUids,omitempty" type:"Repeated"`
}

func (s GetGroupMemberByIdsRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetGroupMemberByIdsRequestRequestParams) GoString() string {
	return s.String()
}

func (s *GetGroupMemberByIdsRequestRequestParams) SetAppCid(v string) *GetGroupMemberByIdsRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *GetGroupMemberByIdsRequestRequestParams) SetAppUids(v []*string) *GetGroupMemberByIdsRequestRequestParams {
	s.AppUids = v
	return s
}

type GetGroupMemberByIdsShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群聊设置扩展字段请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s GetGroupMemberByIdsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupMemberByIdsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetGroupMemberByIdsShrinkRequest) SetAppId(v string) *GetGroupMemberByIdsShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetGroupMemberByIdsShrinkRequest) SetRequestParamsShrink(v string) *GetGroupMemberByIdsShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type GetGroupMemberByIdsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *GetGroupMemberByIdsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetGroupMemberByIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupMemberByIdsResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupMemberByIdsResponseBody) SetRequestId(v string) *GetGroupMemberByIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupMemberByIdsResponseBody) SetCode(v string) *GetGroupMemberByIdsResponseBody {
	s.Code = &v
	return s
}

func (s *GetGroupMemberByIdsResponseBody) SetMessage(v string) *GetGroupMemberByIdsResponseBody {
	s.Message = &v
	return s
}

func (s *GetGroupMemberByIdsResponseBody) SetResult(v *GetGroupMemberByIdsResponseBodyResult) *GetGroupMemberByIdsResponseBody {
	s.Result = v
	return s
}

type GetGroupMemberByIdsResponseBodyResult struct {
	// 群成员
	Members []*GetGroupMemberByIdsResponseBodyResultMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s GetGroupMemberByIdsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetGroupMemberByIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupMemberByIdsResponseBodyResult) SetMembers(v []*GetGroupMemberByIdsResponseBodyResultMembers) *GetGroupMemberByIdsResponseBodyResult {
	s.Members = v
	return s
}

type GetGroupMemberByIdsResponseBodyResultMembers struct {
	AppUid     *string            `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	Role       *int32             `json:"Role,omitempty" xml:"Role,omitempty"`
	Nick       *string            `json:"Nick,omitempty" xml:"Nick,omitempty"`
	JoinTime   *int64             `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s GetGroupMemberByIdsResponseBodyResultMembers) String() string {
	return tea.Prettify(s)
}

func (s GetGroupMemberByIdsResponseBodyResultMembers) GoString() string {
	return s.String()
}

func (s *GetGroupMemberByIdsResponseBodyResultMembers) SetAppUid(v string) *GetGroupMemberByIdsResponseBodyResultMembers {
	s.AppUid = &v
	return s
}

func (s *GetGroupMemberByIdsResponseBodyResultMembers) SetRole(v int32) *GetGroupMemberByIdsResponseBodyResultMembers {
	s.Role = &v
	return s
}

func (s *GetGroupMemberByIdsResponseBodyResultMembers) SetNick(v string) *GetGroupMemberByIdsResponseBodyResultMembers {
	s.Nick = &v
	return s
}

func (s *GetGroupMemberByIdsResponseBodyResultMembers) SetJoinTime(v int64) *GetGroupMemberByIdsResponseBodyResultMembers {
	s.JoinTime = &v
	return s
}

func (s *GetGroupMemberByIdsResponseBodyResultMembers) SetExtensions(v map[string]*string) *GetGroupMemberByIdsResponseBodyResultMembers {
	s.Extensions = v
	return s
}

type GetGroupMemberByIdsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGroupMemberByIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGroupMemberByIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupMemberByIdsResponse) GoString() string {
	return s.String()
}

func (s *GetGroupMemberByIdsResponse) SetHeaders(v map[string]*string) *GetGroupMemberByIdsResponse {
	s.Headers = v
	return s
}

func (s *GetGroupMemberByIdsResponse) SetBody(v *GetGroupMemberByIdsResponseBody) *GetGroupMemberByIdsResponse {
	s.Body = v
	return s
}

type SendCustomMessageRequest struct {
	// 请求参数的结构体。
	Request *SendCustomMessageRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s SendCustomMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageRequest) GoString() string {
	return s.String()
}

func (s *SendCustomMessageRequest) SetRequest(v *SendCustomMessageRequestRequest) *SendCustomMessageRequest {
	s.Request = v
	return s
}

type SendCustomMessageRequestRequest struct {
	// 应用的appKey。
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 房间ID，由调用CreateRoom时返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 消息的发送者ID。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 消息的类型，由业务自定义，请传递100000以上。
	SubType *int32 `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// 消息体。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s SendCustomMessageRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageRequestRequest) GoString() string {
	return s.String()
}

func (s *SendCustomMessageRequestRequest) SetDomain(v string) *SendCustomMessageRequestRequest {
	s.Domain = &v
	return s
}

func (s *SendCustomMessageRequestRequest) SetRoomId(v string) *SendCustomMessageRequestRequest {
	s.RoomId = &v
	return s
}

func (s *SendCustomMessageRequestRequest) SetSenderId(v string) *SendCustomMessageRequestRequest {
	s.SenderId = &v
	return s
}

func (s *SendCustomMessageRequestRequest) SetSubType(v int32) *SendCustomMessageRequestRequest {
	s.SubType = &v
	return s
}

func (s *SendCustomMessageRequestRequest) SetBody(v string) *SendCustomMessageRequestRequest {
	s.Body = &v
	return s
}

type SendCustomMessageResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功。
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码，请求异常时返回。
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息，请求异常时返回。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求的返回结果。
	Result *SendCustomMessageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendCustomMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomMessageResponseBody) SetRequestId(v string) *SendCustomMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomMessageResponseBody) SetResponseSuccess(v bool) *SendCustomMessageResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SendCustomMessageResponseBody) SetErrorCode(v string) *SendCustomMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendCustomMessageResponseBody) SetErrorMessage(v string) *SendCustomMessageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendCustomMessageResponseBody) SetResult(v *SendCustomMessageResponseBodyResult) *SendCustomMessageResponseBody {
	s.Result = v
	return s
}

type SendCustomMessageResponseBodyResult struct {
	// 消息的唯一ID标识。由数字、大小写字母组成，长度不超过20。
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendCustomMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendCustomMessageResponseBodyResult) SetMessageId(v string) *SendCustomMessageResponseBodyResult {
	s.MessageId = &v
	return s
}

type SendCustomMessageResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCustomMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCustomMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageResponse) GoString() string {
	return s.String()
}

func (s *SendCustomMessageResponse) SetHeaders(v map[string]*string) *SendCustomMessageResponse {
	s.Headers = v
	return s
}

func (s *SendCustomMessageResponse) SetBody(v *SendCustomMessageResponseBody) *SendCustomMessageResponse {
	s.Body = v
	return s
}

type UpdateAppNameRequest struct {
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 请求
	RequestParams *UpdateAppNameRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s UpdateAppNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppNameRequest) SetAppId(v string) *UpdateAppNameRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppNameRequest) SetRequestParams(v *UpdateAppNameRequestRequestParams) *UpdateAppNameRequest {
	s.RequestParams = v
	return s
}

type UpdateAppNameRequestRequestParams struct {
	// 应用名
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s UpdateAppNameRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppNameRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateAppNameRequestRequestParams) SetAppName(v string) *UpdateAppNameRequestRequestParams {
	s.AppName = &v
	return s
}

type UpdateAppNameShrinkRequest struct {
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 请求
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s UpdateAppNameShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppNameShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppNameShrinkRequest) SetAppId(v string) *UpdateAppNameShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppNameShrinkRequest) SetRequestParamsShrink(v string) *UpdateAppNameShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type UpdateAppNameResponseBody struct {
	// desc
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAppNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppNameResponseBody) SetMessage(v string) *UpdateAppNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppNameResponseBody) SetRequestId(v string) *UpdateAppNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppNameResponseBody) SetHttpStatusCode(v int32) *UpdateAppNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAppNameResponseBody) SetCode(v string) *UpdateAppNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppNameResponseBody) SetSuccess(v bool) *UpdateAppNameResponseBody {
	s.Success = &v
	return s
}

type UpdateAppNameResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppNameResponse) SetHeaders(v map[string]*string) *UpdateAppNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppNameResponse) SetBody(v *UpdateAppNameResponseBody) *UpdateAppNameResponse {
	s.Body = v
	return s
}

type GetIMConfigRequest struct {
	// 应用名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetIMConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIMConfigRequest) GoString() string {
	return s.String()
}

func (s *GetIMConfigRequest) SetAppId(v string) *GetIMConfigRequest {
	s.AppId = &v
	return s
}

type GetIMConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 网络错误码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	Messaage *string `json:"Messaage,omitempty" xml:"Messaage,omitempty"`
	// 返回结果
	Result *GetIMConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetIMConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIMConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetIMConfigResponseBody) SetRequestId(v string) *GetIMConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIMConfigResponseBody) SetSuccess(v bool) *GetIMConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetIMConfigResponseBody) SetCode(v string) *GetIMConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetIMConfigResponseBody) SetHttpStatusCode(v int32) *GetIMConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIMConfigResponseBody) SetMessaage(v string) *GetIMConfigResponseBody {
	s.Messaage = &v
	return s
}

func (s *GetIMConfigResponseBody) SetResult(v *GetIMConfigResponseBodyResult) *GetIMConfigResponseBody {
	s.Result = v
	return s
}

type GetIMConfigResponseBodyResult struct {
	// im相关配置
	ImConfig *GetIMConfigResponseBodyResultImConfig `json:"ImConfig,omitempty" xml:"ImConfig,omitempty" type:"Struct"`
}

func (s GetIMConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetIMConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIMConfigResponseBodyResult) SetImConfig(v *GetIMConfigResponseBodyResultImConfig) *GetIMConfigResponseBodyResult {
	s.ImConfig = v
	return s
}

type GetIMConfigResponseBodyResultImConfig struct {
	// 消息配置
	MsgConfig *GetIMConfigResponseBodyResultImConfigMsgConfig `json:"MsgConfig,omitempty" xml:"MsgConfig,omitempty" type:"Struct"`
	// 回调配置
	CallbackConfig *GetIMConfigResponseBodyResultImConfigCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
}

func (s GetIMConfigResponseBodyResultImConfig) String() string {
	return tea.Prettify(s)
}

func (s GetIMConfigResponseBodyResultImConfig) GoString() string {
	return s.String()
}

func (s *GetIMConfigResponseBodyResultImConfig) SetMsgConfig(v *GetIMConfigResponseBodyResultImConfigMsgConfig) *GetIMConfigResponseBodyResultImConfig {
	s.MsgConfig = v
	return s
}

func (s *GetIMConfigResponseBodyResultImConfig) SetCallbackConfig(v *GetIMConfigResponseBodyResultImConfigCallbackConfig) *GetIMConfigResponseBodyResultImConfig {
	s.CallbackConfig = v
	return s
}

type GetIMConfigResponseBodyResultImConfigMsgConfig struct {
	// 消息撤回时间，分钟
	ClientMsgRecallTimeIntervalMinute *int64 `json:"ClientMsgRecallTimeIntervalMinute,omitempty" xml:"ClientMsgRecallTimeIntervalMinute,omitempty"`
}

func (s GetIMConfigResponseBodyResultImConfigMsgConfig) String() string {
	return tea.Prettify(s)
}

func (s GetIMConfigResponseBodyResultImConfigMsgConfig) GoString() string {
	return s.String()
}

func (s *GetIMConfigResponseBodyResultImConfigMsgConfig) SetClientMsgRecallTimeIntervalMinute(v int64) *GetIMConfigResponseBodyResultImConfigMsgConfig {
	s.ClientMsgRecallTimeIntervalMinute = &v
	return s
}

type GetIMConfigResponseBodyResultImConfigCallbackConfig struct {
	// 回调url，支持外部url
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// 加签密钥-key
	SignatureKey *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	// 加签密钥-value
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
	// 已开通的回调方法Id列表
	Apis map[string]*bool `json:"Apis,omitempty" xml:"Apis,omitempty"`
	// 已开通的回调方法Id列表
	Spis map[string]*bool `json:"Spis,omitempty" xml:"Spis,omitempty"`
	// 已开通的事件输出列表
	Events map[string]*bool `json:"Events,omitempty" xml:"Events,omitempty"`
}

func (s GetIMConfigResponseBodyResultImConfigCallbackConfig) String() string {
	return tea.Prettify(s)
}

func (s GetIMConfigResponseBodyResultImConfigCallbackConfig) GoString() string {
	return s.String()
}

func (s *GetIMConfigResponseBodyResultImConfigCallbackConfig) SetCallbackUrl(v string) *GetIMConfigResponseBodyResultImConfigCallbackConfig {
	s.CallbackUrl = &v
	return s
}

func (s *GetIMConfigResponseBodyResultImConfigCallbackConfig) SetSignatureKey(v string) *GetIMConfigResponseBodyResultImConfigCallbackConfig {
	s.SignatureKey = &v
	return s
}

func (s *GetIMConfigResponseBodyResultImConfigCallbackConfig) SetSignatureValue(v string) *GetIMConfigResponseBodyResultImConfigCallbackConfig {
	s.SignatureValue = &v
	return s
}

func (s *GetIMConfigResponseBodyResultImConfigCallbackConfig) SetApis(v map[string]*bool) *GetIMConfigResponseBodyResultImConfigCallbackConfig {
	s.Apis = v
	return s
}

func (s *GetIMConfigResponseBodyResultImConfigCallbackConfig) SetSpis(v map[string]*bool) *GetIMConfigResponseBodyResultImConfigCallbackConfig {
	s.Spis = v
	return s
}

func (s *GetIMConfigResponseBodyResultImConfigCallbackConfig) SetEvents(v map[string]*bool) *GetIMConfigResponseBodyResultImConfigCallbackConfig {
	s.Events = v
	return s
}

type GetIMConfigResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIMConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIMConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIMConfigResponse) GoString() string {
	return s.String()
}

func (s *GetIMConfigResponse) SetHeaders(v map[string]*string) *GetIMConfigResponse {
	s.Headers = v
	return s
}

func (s *GetIMConfigResponse) SetBody(v *GetIMConfigResponseBody) *GetIMConfigResponse {
	s.Body = v
	return s
}

type SetSingleChatExtensionByKeysRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建群聊请求实体
	RequestParams *SetSingleChatExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SetSingleChatExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSingleChatExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *SetSingleChatExtensionByKeysRequest) SetAppId(v string) *SetSingleChatExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *SetSingleChatExtensionByKeysRequest) SetRequestParams(v *SetSingleChatExtensionByKeysRequestRequestParams) *SetSingleChatExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type SetSingleChatExtensionByKeysRequestRequestParams struct {
	// 用户id
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 拓展字段
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s SetSingleChatExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SetSingleChatExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SetSingleChatExtensionByKeysRequestRequestParams) SetAppUid(v string) *SetSingleChatExtensionByKeysRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *SetSingleChatExtensionByKeysRequestRequestParams) SetAppCid(v string) *SetSingleChatExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *SetSingleChatExtensionByKeysRequestRequestParams) SetExtensions(v map[string]*string) *SetSingleChatExtensionByKeysRequestRequestParams {
	s.Extensions = v
	return s
}

type SetSingleChatExtensionByKeysShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建群聊请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SetSingleChatExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSingleChatExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSingleChatExtensionByKeysShrinkRequest) SetAppId(v string) *SetSingleChatExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SetSingleChatExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *SetSingleChatExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SetSingleChatExtensionByKeysResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SetSingleChatExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSingleChatExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *SetSingleChatExtensionByKeysResponseBody) SetRequestId(v string) *SetSingleChatExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSingleChatExtensionByKeysResponseBody) SetCode(v string) *SetSingleChatExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *SetSingleChatExtensionByKeysResponseBody) SetMessage(v string) *SetSingleChatExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type SetSingleChatExtensionByKeysResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetSingleChatExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetSingleChatExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSingleChatExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *SetSingleChatExtensionByKeysResponse) SetHeaders(v map[string]*string) *SetSingleChatExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *SetSingleChatExtensionByKeysResponse) SetBody(v *SetSingleChatExtensionByKeysResponseBody) *SetSingleChatExtensionByKeysResponse {
	s.Body = v
	return s
}

type UpdateAppStatusRequest struct {
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 请求
	RequestParams *UpdateAppStatusRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s UpdateAppStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppStatusRequest) SetAppId(v string) *UpdateAppStatusRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppStatusRequest) SetRequestParams(v *UpdateAppStatusRequestRequestParams) *UpdateAppStatusRequest {
	s.RequestParams = v
	return s
}

type UpdateAppStatusRequestRequestParams struct {
	// 是否开启
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s UpdateAppStatusRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppStatusRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateAppStatusRequestRequestParams) SetEnable(v bool) *UpdateAppStatusRequestRequestParams {
	s.Enable = &v
	return s
}

type UpdateAppStatusShrinkRequest struct {
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 请求
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s UpdateAppStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppStatusShrinkRequest) SetAppId(v string) *UpdateAppStatusShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppStatusShrinkRequest) SetRequestParamsShrink(v string) *UpdateAppStatusShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type UpdateAppStatusResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateAppStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppStatusResponseBody) SetRequestId(v string) *UpdateAppStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppStatusResponseBody) SetSuccess(v bool) *UpdateAppStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAppStatusResponseBody) SetMessage(v string) *UpdateAppStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppStatusResponseBody) SetCode(v string) *UpdateAppStatusResponseBody {
	s.Code = &v
	return s
}

type UpdateAppStatusResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppStatusResponse) SetHeaders(v map[string]*string) *UpdateAppStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppStatusResponse) SetBody(v *UpdateAppStatusResponseBody) *UpdateAppStatusResponse {
	s.Body = v
	return s
}

type MuteUsersRequest struct {
	// AppId
	AppId         *string                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *MuteUsersRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s MuteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s MuteUsersRequest) GoString() string {
	return s.String()
}

func (s *MuteUsersRequest) SetAppId(v string) *MuteUsersRequest {
	s.AppId = &v
	return s
}

func (s *MuteUsersRequest) SetRequestParams(v *MuteUsersRequestRequestParams) *MuteUsersRequest {
	s.RequestParams = v
	return s
}

type MuteUsersRequestRequestParams struct {
	// 需要禁言的用户
	AppUids []*string `json:"AppUids,omitempty" xml:"AppUids,omitempty" type:"Repeated"`
	// 单位秒
	MuteDuration *int64 `json:"MuteDuration,omitempty" xml:"MuteDuration,omitempty"`
	// true: 禁言, false: 解除禁言
	Mute *bool `json:"Mute,omitempty" xml:"Mute,omitempty"`
}

func (s MuteUsersRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s MuteUsersRequestRequestParams) GoString() string {
	return s.String()
}

func (s *MuteUsersRequestRequestParams) SetAppUids(v []*string) *MuteUsersRequestRequestParams {
	s.AppUids = v
	return s
}

func (s *MuteUsersRequestRequestParams) SetMuteDuration(v int64) *MuteUsersRequestRequestParams {
	s.MuteDuration = &v
	return s
}

func (s *MuteUsersRequestRequestParams) SetMute(v bool) *MuteUsersRequestRequestParams {
	s.Mute = &v
	return s
}

type MuteUsersShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s MuteUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s MuteUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *MuteUsersShrinkRequest) SetAppId(v string) *MuteUsersShrinkRequest {
	s.AppId = &v
	return s
}

func (s *MuteUsersShrinkRequest) SetRequestParamsShrink(v string) *MuteUsersShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type MuteUsersResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s MuteUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MuteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *MuteUsersResponseBody) SetRequestId(v string) *MuteUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteUsersResponseBody) SetCode(v string) *MuteUsersResponseBody {
	s.Code = &v
	return s
}

func (s *MuteUsersResponseBody) SetMessage(v string) *MuteUsersResponseBody {
	s.Message = &v
	return s
}

type MuteUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MuteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MuteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s MuteUsersResponse) GoString() string {
	return s.String()
}

func (s *MuteUsersResponse) SetHeaders(v map[string]*string) *MuteUsersResponse {
	s.Headers = v
	return s
}

func (s *MuteUsersResponse) SetBody(v *MuteUsersResponseBody) *MuteUsersResponse {
	s.Body = v
	return s
}

type RecallMessageRequest struct {
	// AppId
	AppId         *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *RecallMessageRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RecallMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageRequest) GoString() string {
	return s.String()
}

func (s *RecallMessageRequest) SetAppId(v string) *RecallMessageRequest {
	s.AppId = &v
	return s
}

func (s *RecallMessageRequest) SetRequestParams(v *RecallMessageRequestRequestParams) *RecallMessageRequest {
	s.RequestParams = v
	return s
}

type RecallMessageRequestRequestParams struct {
	// 操作者ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 消息ID
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// 撤回显示类型（默认为0)。0：静默撤回，不显示撤回信息，1：普通撤回，显示撤回信息；
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 操作者类型(默认为0)。0: 发送者; 1: 群主; 2: 系统; 3: 安全合规; 101: 业务自定义类型
	OperatorType *int32 `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 业务自定义扩展字段
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s RecallMessageRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RecallMessageRequestRequestParams) SetAppUid(v string) *RecallMessageRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *RecallMessageRequestRequestParams) SetAppCid(v string) *RecallMessageRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RecallMessageRequestRequestParams) SetMsgId(v string) *RecallMessageRequestRequestParams {
	s.MsgId = &v
	return s
}

func (s *RecallMessageRequestRequestParams) SetType(v int32) *RecallMessageRequestRequestParams {
	s.Type = &v
	return s
}

func (s *RecallMessageRequestRequestParams) SetOperatorType(v int32) *RecallMessageRequestRequestParams {
	s.OperatorType = &v
	return s
}

func (s *RecallMessageRequestRequestParams) SetExtensions(v map[string]*string) *RecallMessageRequestRequestParams {
	s.Extensions = v
	return s
}

type RecallMessageShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RecallMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecallMessageShrinkRequest) SetAppId(v string) *RecallMessageShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RecallMessageShrinkRequest) SetRequestParamsShrink(v string) *RecallMessageShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RecallMessageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecallMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RecallMessageResponseBody) SetRequestId(v string) *RecallMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecallMessageResponseBody) SetCode(v string) *RecallMessageResponseBody {
	s.Code = &v
	return s
}

func (s *RecallMessageResponseBody) SetMessage(v string) *RecallMessageResponseBody {
	s.Message = &v
	return s
}

type RecallMessageResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecallMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecallMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallMessageResponse) GoString() string {
	return s.String()
}

func (s *RecallMessageResponse) SetHeaders(v map[string]*string) *RecallMessageResponse {
	s.Headers = v
	return s
}

func (s *RecallMessageResponse) SetBody(v *RecallMessageResponseBody) *RecallMessageResponse {
	s.Body = v
	return s
}

type AddGroupSilenceWhitelistRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言添加白名单请求体
	RequestParams *AddGroupSilenceWhitelistRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s AddGroupSilenceWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceWhitelistRequest) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceWhitelistRequest) SetAppId(v string) *AddGroupSilenceWhitelistRequest {
	s.AppId = &v
	return s
}

func (s *AddGroupSilenceWhitelistRequest) SetRequestParams(v *AddGroupSilenceWhitelistRequestRequestParams) *AddGroupSilenceWhitelistRequest {
	s.RequestParams = v
	return s
}

type AddGroupSilenceWhitelistRequestRequestParams struct {
	// 操作者
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 群会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 禁言用户列表
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s AddGroupSilenceWhitelistRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceWhitelistRequestRequestParams) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceWhitelistRequestRequestParams) SetOperatorAppUid(v string) *AddGroupSilenceWhitelistRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *AddGroupSilenceWhitelistRequestRequestParams) SetAppCid(v string) *AddGroupSilenceWhitelistRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *AddGroupSilenceWhitelistRequestRequestParams) SetMembers(v []*string) *AddGroupSilenceWhitelistRequestRequestParams {
	s.Members = v
	return s
}

type AddGroupSilenceWhitelistShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言添加白名单请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s AddGroupSilenceWhitelistShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceWhitelistShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceWhitelistShrinkRequest) SetAppId(v string) *AddGroupSilenceWhitelistShrinkRequest {
	s.AppId = &v
	return s
}

func (s *AddGroupSilenceWhitelistShrinkRequest) SetRequestParamsShrink(v string) *AddGroupSilenceWhitelistShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type AddGroupSilenceWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddGroupSilenceWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceWhitelistResponseBody) SetRequestId(v string) *AddGroupSilenceWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGroupSilenceWhitelistResponseBody) SetCode(v string) *AddGroupSilenceWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *AddGroupSilenceWhitelistResponseBody) SetMessage(v string) *AddGroupSilenceWhitelistResponseBody {
	s.Message = &v
	return s
}

type AddGroupSilenceWhitelistResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGroupSilenceWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGroupSilenceWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupSilenceWhitelistResponse) GoString() string {
	return s.String()
}

func (s *AddGroupSilenceWhitelistResponse) SetHeaders(v map[string]*string) *AddGroupSilenceWhitelistResponse {
	s.Headers = v
	return s
}

func (s *AddGroupSilenceWhitelistResponse) SetBody(v *AddGroupSilenceWhitelistResponseBody) *AddGroupSilenceWhitelistResponse {
	s.Body = v
	return s
}

type SetGroupOwnerRequest struct {
	// App ID，IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群主转让的请求体
	RequestParams *SetGroupOwnerRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SetGroupOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGroupOwnerRequest) GoString() string {
	return s.String()
}

func (s *SetGroupOwnerRequest) SetAppId(v string) *SetGroupOwnerRequest {
	s.AppId = &v
	return s
}

func (s *SetGroupOwnerRequest) SetRequestParams(v *SetGroupOwnerRequestRequestParams) *SetGroupOwnerRequest {
	s.RequestParams = v
	return s
}

type SetGroupOwnerRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 新群主
	NewOwnerAppUid *string `json:"NewOwnerAppUid,omitempty" xml:"NewOwnerAppUid,omitempty"`
}

func (s SetGroupOwnerRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SetGroupOwnerRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SetGroupOwnerRequestRequestParams) SetAppCid(v string) *SetGroupOwnerRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *SetGroupOwnerRequestRequestParams) SetNewOwnerAppUid(v string) *SetGroupOwnerRequestRequestParams {
	s.NewOwnerAppUid = &v
	return s
}

type SetGroupOwnerShrinkRequest struct {
	// App ID，IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群主转让的请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SetGroupOwnerShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGroupOwnerShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetGroupOwnerShrinkRequest) SetAppId(v string) *SetGroupOwnerShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SetGroupOwnerShrinkRequest) SetRequestParamsShrink(v string) *SetGroupOwnerShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SetGroupOwnerResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码。
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息。
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SetGroupOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGroupOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *SetGroupOwnerResponseBody) SetRequestId(v string) *SetGroupOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGroupOwnerResponseBody) SetCode(v string) *SetGroupOwnerResponseBody {
	s.Code = &v
	return s
}

func (s *SetGroupOwnerResponseBody) SetMessage(v string) *SetGroupOwnerResponseBody {
	s.Message = &v
	return s
}

type SetGroupOwnerResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGroupOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGroupOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGroupOwnerResponse) GoString() string {
	return s.String()
}

func (s *SetGroupOwnerResponse) SetHeaders(v map[string]*string) *SetGroupOwnerResponse {
	s.Headers = v
	return s
}

func (s *SetGroupOwnerResponse) SetBody(v *SetGroupOwnerResponseBody) *SetGroupOwnerResponse {
	s.Body = v
	return s
}

type ListRoomUsersRequest struct {
	// 请求参数的结构体。
	Request *ListRoomUsersRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s ListRoomUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRoomUsersRequest) SetRequest(v *ListRoomUsersRequestRequest) *ListRoomUsersRequest {
	s.Request = v
	return s
}

type ListRoomUsersRequestRequest struct {
	// 应用的appKey。
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 房间ID，由调用CreateRoom时返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 分页查询时的页数，从1开始，每次分页查询时加1。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页查询时的请求大小，要求大于0，最大不得超过100。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRoomUsersRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersRequestRequest) GoString() string {
	return s.String()
}

func (s *ListRoomUsersRequestRequest) SetDomain(v string) *ListRoomUsersRequestRequest {
	s.Domain = &v
	return s
}

func (s *ListRoomUsersRequestRequest) SetRoomId(v string) *ListRoomUsersRequestRequest {
	s.RoomId = &v
	return s
}

func (s *ListRoomUsersRequestRequest) SetPageNumber(v int32) *ListRoomUsersRequestRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoomUsersRequestRequest) SetPageSize(v int32) *ListRoomUsersRequestRequest {
	s.PageSize = &v
	return s
}

type ListRoomUsersResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功。
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码，请求异常时返回。
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息，请求异常时返回。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求的返回结果。
	Result *ListRoomUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListRoomUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponseBody) SetRequestId(v string) *ListRoomUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoomUsersResponseBody) SetResponseSuccess(v bool) *ListRoomUsersResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *ListRoomUsersResponseBody) SetErrorCode(v string) *ListRoomUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRoomUsersResponseBody) SetErrorMessage(v string) *ListRoomUsersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRoomUsersResponseBody) SetResult(v *ListRoomUsersResponseBodyResult) *ListRoomUsersResponseBody {
	s.Result = v
	return s
}

type ListRoomUsersResponseBodyResult struct {
	// 房间的历史观看成员总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 返回的观众列表。
	RoomUserVOList []*ListRoomUsersResponseBodyResultRoomUserVOList `json:"RoomUserVOList,omitempty" xml:"RoomUserVOList,omitempty" type:"Repeated"`
	// 是否还有下一页查询的数据。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
}

func (s ListRoomUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponseBodyResult) SetTotalCount(v int32) *ListRoomUsersResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListRoomUsersResponseBodyResult) SetRoomUserVOList(v []*ListRoomUsersResponseBodyResultRoomUserVOList) *ListRoomUsersResponseBodyResult {
	s.RoomUserVOList = v
	return s
}

func (s *ListRoomUsersResponseBodyResult) SetHasMore(v bool) *ListRoomUsersResponseBodyResult {
	s.HasMore = &v
	return s
}

type ListRoomUsersResponseBodyResultRoomUserVOList struct {
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户的昵称。
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
}

func (s ListRoomUsersResponseBodyResultRoomUserVOList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponseBodyResultRoomUserVOList) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponseBodyResultRoomUserVOList) SetRoomId(v string) *ListRoomUsersResponseBodyResultRoomUserVOList {
	s.RoomId = &v
	return s
}

func (s *ListRoomUsersResponseBodyResultRoomUserVOList) SetUserId(v string) *ListRoomUsersResponseBodyResultRoomUserVOList {
	s.UserId = &v
	return s
}

func (s *ListRoomUsersResponseBodyResultRoomUserVOList) SetNick(v string) *ListRoomUsersResponseBodyResultRoomUserVOList {
	s.Nick = &v
	return s
}

type ListRoomUsersResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRoomUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoomUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponse) SetHeaders(v map[string]*string) *ListRoomUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRoomUsersResponse) SetBody(v *ListRoomUsersResponseBody) *ListRoomUsersResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	// 应用id
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
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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

func (s *DeleteAppResponseBody) SetSuccess(v bool) *DeleteAppResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAppResponseBody) SetMessage(v string) *DeleteAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResponseBody) SetCode(v string) *DeleteAppResponseBody {
	s.Code = &v
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

type RemoveGroupSilenceBlacklistRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言删除黑名单请求体
	RequestParams *RemoveGroupSilenceBlacklistRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RemoveGroupSilenceBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceBlacklistRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceBlacklistRequest) SetAppId(v string) *RemoveGroupSilenceBlacklistRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupSilenceBlacklistRequest) SetRequestParams(v *RemoveGroupSilenceBlacklistRequestRequestParams) *RemoveGroupSilenceBlacklistRequest {
	s.RequestParams = v
	return s
}

type RemoveGroupSilenceBlacklistRequestRequestParams struct {
	// 操作者
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 群会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 禁言用户列表
	Members []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s RemoveGroupSilenceBlacklistRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceBlacklistRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceBlacklistRequestRequestParams) SetOperatorAppUid(v string) *RemoveGroupSilenceBlacklistRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *RemoveGroupSilenceBlacklistRequestRequestParams) SetAppCid(v string) *RemoveGroupSilenceBlacklistRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RemoveGroupSilenceBlacklistRequestRequestParams) SetMembers(v []*string) *RemoveGroupSilenceBlacklistRequestRequestParams {
	s.Members = v
	return s
}

type RemoveGroupSilenceBlacklistShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言删除黑名单请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RemoveGroupSilenceBlacklistShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceBlacklistShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceBlacklistShrinkRequest) SetAppId(v string) *RemoveGroupSilenceBlacklistShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupSilenceBlacklistShrinkRequest) SetRequestParamsShrink(v string) *RemoveGroupSilenceBlacklistShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RemoveGroupSilenceBlacklistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveGroupSilenceBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceBlacklistResponseBody) SetRequestId(v string) *RemoveGroupSilenceBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGroupSilenceBlacklistResponseBody) SetCode(v string) *RemoveGroupSilenceBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveGroupSilenceBlacklistResponseBody) SetMessage(v string) *RemoveGroupSilenceBlacklistResponseBody {
	s.Message = &v
	return s
}

type RemoveGroupSilenceBlacklistResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveGroupSilenceBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveGroupSilenceBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupSilenceBlacklistResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupSilenceBlacklistResponse) SetHeaders(v map[string]*string) *RemoveGroupSilenceBlacklistResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupSilenceBlacklistResponse) SetBody(v *RemoveGroupSilenceBlacklistResponseBody) *RemoveGroupSilenceBlacklistResponse {
	s.Body = v
	return s
}

type RemoveMessageExtensionByKeysRequest struct {
	// AppId
	AppId         *string                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *RemoveMessageExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RemoveMessageExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessageExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *RemoveMessageExtensionByKeysRequest) SetAppId(v string) *RemoveMessageExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *RemoveMessageExtensionByKeysRequest) SetRequestParams(v *RemoveMessageExtensionByKeysRequestRequestParams) *RemoveMessageExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type RemoveMessageExtensionByKeysRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 消息ID
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// 需删除的Key
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s RemoveMessageExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessageExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RemoveMessageExtensionByKeysRequestRequestParams) SetAppCid(v string) *RemoveMessageExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RemoveMessageExtensionByKeysRequestRequestParams) SetMsgId(v string) *RemoveMessageExtensionByKeysRequestRequestParams {
	s.MsgId = &v
	return s
}

func (s *RemoveMessageExtensionByKeysRequestRequestParams) SetKeys(v []*string) *RemoveMessageExtensionByKeysRequestRequestParams {
	s.Keys = v
	return s
}

type RemoveMessageExtensionByKeysShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RemoveMessageExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessageExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveMessageExtensionByKeysShrinkRequest) SetAppId(v string) *RemoveMessageExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveMessageExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *RemoveMessageExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RemoveMessageExtensionByKeysResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码，成功时为0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息，成功时为0	空
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveMessageExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessageExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMessageExtensionByKeysResponseBody) SetRequestId(v string) *RemoveMessageExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveMessageExtensionByKeysResponseBody) SetCode(v string) *RemoveMessageExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveMessageExtensionByKeysResponseBody) SetMessage(v string) *RemoveMessageExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type RemoveMessageExtensionByKeysResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveMessageExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveMessageExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessageExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *RemoveMessageExtensionByKeysResponse) SetHeaders(v map[string]*string) *RemoveMessageExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *RemoveMessageExtensionByKeysResponse) SetBody(v *RemoveMessageExtensionByKeysResponseBody) *RemoveMessageExtensionByKeysResponse {
	s.Body = v
	return s
}

type GetMediaUploadUrlRequest struct {
	// AppId
	AppId         *string                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *GetMediaUploadUrlRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s GetMediaUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMediaUploadUrlRequest) SetAppId(v string) *GetMediaUploadUrlRequest {
	s.AppId = &v
	return s
}

func (s *GetMediaUploadUrlRequest) SetRequestParams(v *GetMediaUploadUrlRequestRequestParams) *GetMediaUploadUrlRequest {
	s.RequestParams = v
	return s
}

type GetMediaUploadUrlRequestRequestParams struct {
	// 多媒体资源类型(文件后缀名)
	MimeType *string `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
}

func (s GetMediaUploadUrlRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUploadUrlRequestRequestParams) GoString() string {
	return s.String()
}

func (s *GetMediaUploadUrlRequestRequestParams) SetMimeType(v string) *GetMediaUploadUrlRequestRequestParams {
	s.MimeType = &v
	return s
}

type GetMediaUploadUrlShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s GetMediaUploadUrlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUploadUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMediaUploadUrlShrinkRequest) SetAppId(v string) *GetMediaUploadUrlShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetMediaUploadUrlShrinkRequest) SetRequestParamsShrink(v string) *GetMediaUploadUrlShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type GetMediaUploadUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 调用返回值
	Result *GetMediaUploadUrlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetMediaUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaUploadUrlResponseBody) SetRequestId(v string) *GetMediaUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaUploadUrlResponseBody) SetCode(v string) *GetMediaUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetMediaUploadUrlResponseBody) SetMessage(v string) *GetMediaUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetMediaUploadUrlResponseBody) SetResult(v *GetMediaUploadUrlResponseBodyResult) *GetMediaUploadUrlResponseBody {
	s.Result = v
	return s
}

type GetMediaUploadUrlResponseBodyResult struct {
	// 上传Url
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
	// 多媒体文件ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetMediaUploadUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUploadUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMediaUploadUrlResponseBodyResult) SetUploadUrl(v string) *GetMediaUploadUrlResponseBodyResult {
	s.UploadUrl = &v
	return s
}

func (s *GetMediaUploadUrlResponseBodyResult) SetMediaId(v string) *GetMediaUploadUrlResponseBodyResult {
	s.MediaId = &v
	return s
}

type GetMediaUploadUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMediaUploadUrlResponse) SetHeaders(v map[string]*string) *GetMediaUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMediaUploadUrlResponse) SetBody(v *GetMediaUploadUrlResponseBody) *GetMediaUploadUrlResponse {
	s.Body = v
	return s
}

type BindInterconnectionUidRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 绑定用户请求体
	RequestParams *BindInterconnectionUidRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s BindInterconnectionUidRequest) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionUidRequest) GoString() string {
	return s.String()
}

func (s *BindInterconnectionUidRequest) SetAppId(v string) *BindInterconnectionUidRequest {
	s.AppId = &v
	return s
}

func (s *BindInterconnectionUidRequest) SetRequestParams(v *BindInterconnectionUidRequestRequestParams) *BindInterconnectionUidRequest {
	s.RequestParams = v
	return s
}

type BindInterconnectionUidRequestRequestParams struct {
	// AIM用户ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 钉钉用户ID
	DingTalkUid *string `json:"DingTalkUid,omitempty" xml:"DingTalkUid,omitempty"`
}

func (s BindInterconnectionUidRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionUidRequestRequestParams) GoString() string {
	return s.String()
}

func (s *BindInterconnectionUidRequestRequestParams) SetAppUid(v string) *BindInterconnectionUidRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *BindInterconnectionUidRequestRequestParams) SetDingTalkUid(v string) *BindInterconnectionUidRequestRequestParams {
	s.DingTalkUid = &v
	return s
}

type BindInterconnectionUidShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 绑定用户请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s BindInterconnectionUidShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionUidShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindInterconnectionUidShrinkRequest) SetAppId(v string) *BindInterconnectionUidShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindInterconnectionUidShrinkRequest) SetRequestParamsShrink(v string) *BindInterconnectionUidShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type BindInterconnectionUidResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s BindInterconnectionUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionUidResponseBody) GoString() string {
	return s.String()
}

func (s *BindInterconnectionUidResponseBody) SetRequestId(v string) *BindInterconnectionUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindInterconnectionUidResponseBody) SetCode(v string) *BindInterconnectionUidResponseBody {
	s.Code = &v
	return s
}

func (s *BindInterconnectionUidResponseBody) SetMessage(v string) *BindInterconnectionUidResponseBody {
	s.Message = &v
	return s
}

type BindInterconnectionUidResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindInterconnectionUidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindInterconnectionUidResponse) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionUidResponse) GoString() string {
	return s.String()
}

func (s *BindInterconnectionUidResponse) SetHeaders(v map[string]*string) *BindInterconnectionUidResponse {
	s.Headers = v
	return s
}

func (s *BindInterconnectionUidResponse) SetBody(v *BindInterconnectionUidResponseBody) *BindInterconnectionUidResponse {
	s.Body = v
	return s
}

type GetMediaUrlRequest struct {
	// AppId
	AppId         *string                          `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *GetMediaUrlRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s GetMediaUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMediaUrlRequest) SetAppId(v string) *GetMediaUrlRequest {
	s.AppId = &v
	return s
}

func (s *GetMediaUrlRequest) SetRequestParams(v *GetMediaUrlRequestRequestParams) *GetMediaUrlRequest {
	s.RequestParams = v
	return s
}

type GetMediaUrlRequestRequestParams struct {
	// 多媒体资源ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// URL过期时间(秒，最大86400)
	UrlExpireTime *int64 `json:"UrlExpireTime,omitempty" xml:"UrlExpireTime,omitempty"`
}

func (s GetMediaUrlRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlRequestRequestParams) GoString() string {
	return s.String()
}

func (s *GetMediaUrlRequestRequestParams) SetMediaId(v string) *GetMediaUrlRequestRequestParams {
	s.MediaId = &v
	return s
}

func (s *GetMediaUrlRequestRequestParams) SetUrlExpireTime(v int64) *GetMediaUrlRequestRequestParams {
	s.UrlExpireTime = &v
	return s
}

type GetMediaUrlShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s GetMediaUrlShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMediaUrlShrinkRequest) SetAppId(v string) *GetMediaUrlShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetMediaUrlShrinkRequest) SetRequestParamsShrink(v string) *GetMediaUrlShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type GetMediaUrlResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *GetMediaUrlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetMediaUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaUrlResponseBody) SetRequestId(v string) *GetMediaUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaUrlResponseBody) SetCode(v string) *GetMediaUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetMediaUrlResponseBody) SetMessage(v string) *GetMediaUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetMediaUrlResponseBody) SetResult(v *GetMediaUrlResponseBodyResult) *GetMediaUrlResponseBody {
	s.Result = v
	return s
}

type GetMediaUrlResponseBodyResult struct {
	// 文件Url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMediaUrlResponseBodyResult) SetUrl(v string) *GetMediaUrlResponseBodyResult {
	s.Url = &v
	return s
}

type GetMediaUrlResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMediaUrlResponse) SetHeaders(v map[string]*string) *GetMediaUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMediaUrlResponse) SetBody(v *GetMediaUrlResponseBody) *GetMediaUrlResponse {
	s.Body = v
	return s
}

type ImportSingleConversationRequest struct {
	// AppId
	AppId         *string                                       `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *ImportSingleConversationRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s ImportSingleConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportSingleConversationRequest) GoString() string {
	return s.String()
}

func (s *ImportSingleConversationRequest) SetAppId(v string) *ImportSingleConversationRequest {
	s.AppId = &v
	return s
}

func (s *ImportSingleConversationRequest) SetRequestParams(v *ImportSingleConversationRequestRequestParams) *ImportSingleConversationRequest {
	s.RequestParams = v
	return s
}

type ImportSingleConversationRequestRequestParams struct {
	// 会话基础信息
	Conversation *ImportSingleConversationRequestRequestParamsConversation `json:"Conversation,omitempty" xml:"Conversation,omitempty" type:"Struct"`
	// 用户会话视图
	UserConversations map[string]*RequestParamsUserConversationsValue `json:"UserConversations,omitempty" xml:"UserConversations,omitempty"`
}

func (s ImportSingleConversationRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s ImportSingleConversationRequestRequestParams) GoString() string {
	return s.String()
}

func (s *ImportSingleConversationRequestRequestParams) SetConversation(v *ImportSingleConversationRequestRequestParamsConversation) *ImportSingleConversationRequestRequestParams {
	s.Conversation = v
	return s
}

func (s *ImportSingleConversationRequestRequestParams) SetUserConversations(v map[string]*RequestParamsUserConversationsValue) *ImportSingleConversationRequestRequestParams {
	s.UserConversations = v
	return s
}

type ImportSingleConversationRequestRequestParamsConversation struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 用户ID列表
	AppUids []*string `json:"AppUids,omitempty" xml:"AppUids,omitempty" type:"Repeated"`
	// 扩展字段
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	CreateTime *int64             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ImportSingleConversationRequestRequestParamsConversation) String() string {
	return tea.Prettify(s)
}

func (s ImportSingleConversationRequestRequestParamsConversation) GoString() string {
	return s.String()
}

func (s *ImportSingleConversationRequestRequestParamsConversation) SetAppCid(v string) *ImportSingleConversationRequestRequestParamsConversation {
	s.AppCid = &v
	return s
}

func (s *ImportSingleConversationRequestRequestParamsConversation) SetAppUids(v []*string) *ImportSingleConversationRequestRequestParamsConversation {
	s.AppUids = v
	return s
}

func (s *ImportSingleConversationRequestRequestParamsConversation) SetExtensions(v map[string]*string) *ImportSingleConversationRequestRequestParamsConversation {
	s.Extensions = v
	return s
}

func (s *ImportSingleConversationRequestRequestParamsConversation) SetCreateTime(v int64) *ImportSingleConversationRequestRequestParamsConversation {
	s.CreateTime = &v
	return s
}

type ImportSingleConversationShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s ImportSingleConversationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportSingleConversationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportSingleConversationShrinkRequest) SetAppId(v string) *ImportSingleConversationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ImportSingleConversationShrinkRequest) SetRequestParamsShrink(v string) *ImportSingleConversationShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type ImportSingleConversationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ImportSingleConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportSingleConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ImportSingleConversationResponseBody) SetRequestId(v string) *ImportSingleConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportSingleConversationResponseBody) SetCode(v string) *ImportSingleConversationResponseBody {
	s.Code = &v
	return s
}

func (s *ImportSingleConversationResponseBody) SetMessage(v string) *ImportSingleConversationResponseBody {
	s.Message = &v
	return s
}

type ImportSingleConversationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportSingleConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportSingleConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportSingleConversationResponse) GoString() string {
	return s.String()
}

func (s *ImportSingleConversationResponse) SetHeaders(v map[string]*string) *ImportSingleConversationResponse {
	s.Headers = v
	return s
}

func (s *ImportSingleConversationResponse) SetBody(v *ImportSingleConversationResponseBody) *ImportSingleConversationResponse {
	s.Body = v
	return s
}

type UpdateCallbackConfigRequest struct {
	// 应用Id
	AppId         *string                                   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *UpdateCallbackConfigRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s UpdateCallbackConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigRequest) SetAppId(v string) *UpdateCallbackConfigRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCallbackConfigRequest) SetRequestParams(v *UpdateCallbackConfigRequestRequestParams) *UpdateCallbackConfigRequest {
	s.RequestParams = v
	return s
}

type UpdateCallbackConfigRequestRequestParams struct {
	// 回调url
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// 加签密钥-key
	SignatureKey *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	// 加签密钥-value
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
	// 回调api列表
	Apis map[string]*bool `json:"Apis,omitempty" xml:"Apis,omitempty"`
	// 回调列表
	Spis map[string]*bool `json:"Spis,omitempty" xml:"Spis,omitempty"`
	// 事件输出列表
	Events map[string]*bool `json:"Events,omitempty" xml:"Events,omitempty"`
}

func (s UpdateCallbackConfigRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigRequestRequestParams) SetCallbackUrl(v string) *UpdateCallbackConfigRequestRequestParams {
	s.CallbackUrl = &v
	return s
}

func (s *UpdateCallbackConfigRequestRequestParams) SetSignatureKey(v string) *UpdateCallbackConfigRequestRequestParams {
	s.SignatureKey = &v
	return s
}

func (s *UpdateCallbackConfigRequestRequestParams) SetSignatureValue(v string) *UpdateCallbackConfigRequestRequestParams {
	s.SignatureValue = &v
	return s
}

func (s *UpdateCallbackConfigRequestRequestParams) SetApis(v map[string]*bool) *UpdateCallbackConfigRequestRequestParams {
	s.Apis = v
	return s
}

func (s *UpdateCallbackConfigRequestRequestParams) SetSpis(v map[string]*bool) *UpdateCallbackConfigRequestRequestParams {
	s.Spis = v
	return s
}

func (s *UpdateCallbackConfigRequestRequestParams) SetEvents(v map[string]*bool) *UpdateCallbackConfigRequestRequestParams {
	s.Events = v
	return s
}

type UpdateCallbackConfigShrinkRequest struct {
	// 应用Id
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s UpdateCallbackConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigShrinkRequest) SetAppId(v string) *UpdateCallbackConfigShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCallbackConfigShrinkRequest) SetRequestParamsShrink(v string) *UpdateCallbackConfigShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type UpdateCallbackConfigResponseBody struct {
	// desc
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// result
	Result *UpdateCallbackConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateCallbackConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigResponseBody) SetMessage(v string) *UpdateCallbackConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCallbackConfigResponseBody) SetRequestId(v string) *UpdateCallbackConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCallbackConfigResponseBody) SetHttpStatusCode(v int32) *UpdateCallbackConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCallbackConfigResponseBody) SetCode(v string) *UpdateCallbackConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCallbackConfigResponseBody) SetSuccess(v bool) *UpdateCallbackConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCallbackConfigResponseBody) SetResult(v *UpdateCallbackConfigResponseBodyResult) *UpdateCallbackConfigResponseBody {
	s.Result = v
	return s
}

type UpdateCallbackConfigResponseBodyResult struct {
	// im相关配置
	ImConfig *UpdateCallbackConfigResponseBodyResultImConfig `json:"ImConfig,omitempty" xml:"ImConfig,omitempty" type:"Struct"`
}

func (s UpdateCallbackConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigResponseBodyResult) SetImConfig(v *UpdateCallbackConfigResponseBodyResultImConfig) *UpdateCallbackConfigResponseBodyResult {
	s.ImConfig = v
	return s
}

type UpdateCallbackConfigResponseBodyResultImConfig struct {
	// 消息配置
	MsgConfig *UpdateCallbackConfigResponseBodyResultImConfigMsgConfig `json:"MsgConfig,omitempty" xml:"MsgConfig,omitempty" type:"Struct"`
	// 回调配置
	CallbackConfig *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
}

func (s UpdateCallbackConfigResponseBodyResultImConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigResponseBodyResultImConfig) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigResponseBodyResultImConfig) SetMsgConfig(v *UpdateCallbackConfigResponseBodyResultImConfigMsgConfig) *UpdateCallbackConfigResponseBodyResultImConfig {
	s.MsgConfig = v
	return s
}

func (s *UpdateCallbackConfigResponseBodyResultImConfig) SetCallbackConfig(v *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig) *UpdateCallbackConfigResponseBodyResultImConfig {
	s.CallbackConfig = v
	return s
}

type UpdateCallbackConfigResponseBodyResultImConfigMsgConfig struct {
	// 消息撤回时间间隔，毫秒
	MsgRecallTimeInterval *int64 `json:"MsgRecallTimeInterval,omitempty" xml:"MsgRecallTimeInterval,omitempty"`
}

func (s UpdateCallbackConfigResponseBodyResultImConfigMsgConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigResponseBodyResultImConfigMsgConfig) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigResponseBodyResultImConfigMsgConfig) SetMsgRecallTimeInterval(v int64) *UpdateCallbackConfigResponseBodyResultImConfigMsgConfig {
	s.MsgRecallTimeInterval = &v
	return s
}

type UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig struct {
	// 回调url
	BackendUrl *string `json:"BackendUrl,omitempty" xml:"BackendUrl,omitempty"`
	// 加签密钥-key
	SignatureKey *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	// 加签密钥-value
	SignatureValue *string `json:"SignatureValue,omitempty" xml:"SignatureValue,omitempty"`
	// 回调方法列表
	ApiIds []*string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty" type:"Repeated"`
}

func (s UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig) SetBackendUrl(v string) *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig {
	s.BackendUrl = &v
	return s
}

func (s *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig) SetSignatureKey(v string) *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig {
	s.SignatureKey = &v
	return s
}

func (s *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig) SetSignatureValue(v string) *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig {
	s.SignatureValue = &v
	return s
}

func (s *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig) SetApiIds(v []*string) *UpdateCallbackConfigResponseBodyResultImConfigCallbackConfig {
	s.ApiIds = v
	return s
}

type UpdateCallbackConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCallbackConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCallbackConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCallbackConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateCallbackConfigResponse) SetHeaders(v map[string]*string) *UpdateCallbackConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateCallbackConfigResponse) SetBody(v *UpdateCallbackConfigResponseBody) *UpdateCallbackConfigResponse {
	s.Body = v
	return s
}

type BindInterconnectionCidRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 绑定会话ID请求体
	RequestParams *BindInterconnectionCidRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s BindInterconnectionCidRequest) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionCidRequest) GoString() string {
	return s.String()
}

func (s *BindInterconnectionCidRequest) SetAppId(v string) *BindInterconnectionCidRequest {
	s.AppId = &v
	return s
}

func (s *BindInterconnectionCidRequest) SetRequestParams(v *BindInterconnectionCidRequestRequestParams) *BindInterconnectionCidRequest {
	s.RequestParams = v
	return s
}

type BindInterconnectionCidRequestRequestParams struct {
	// AIM 群会话ID
	AimAppCid *string `json:"AimAppCid,omitempty" xml:"AimAppCid,omitempty"`
	// 钉钉 群会话ID
	DingTalkCid *string `json:"DingTalkCid,omitempty" xml:"DingTalkCid,omitempty"`
}

func (s BindInterconnectionCidRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionCidRequestRequestParams) GoString() string {
	return s.String()
}

func (s *BindInterconnectionCidRequestRequestParams) SetAimAppCid(v string) *BindInterconnectionCidRequestRequestParams {
	s.AimAppCid = &v
	return s
}

func (s *BindInterconnectionCidRequestRequestParams) SetDingTalkCid(v string) *BindInterconnectionCidRequestRequestParams {
	s.DingTalkCid = &v
	return s
}

type BindInterconnectionCidShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 绑定会话ID请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s BindInterconnectionCidShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionCidShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindInterconnectionCidShrinkRequest) SetAppId(v string) *BindInterconnectionCidShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BindInterconnectionCidShrinkRequest) SetRequestParamsShrink(v string) *BindInterconnectionCidShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type BindInterconnectionCidResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s BindInterconnectionCidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionCidResponseBody) GoString() string {
	return s.String()
}

func (s *BindInterconnectionCidResponseBody) SetRequestId(v string) *BindInterconnectionCidResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindInterconnectionCidResponseBody) SetCode(v string) *BindInterconnectionCidResponseBody {
	s.Code = &v
	return s
}

func (s *BindInterconnectionCidResponseBody) SetMessage(v string) *BindInterconnectionCidResponseBody {
	s.Message = &v
	return s
}

type BindInterconnectionCidResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindInterconnectionCidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindInterconnectionCidResponse) String() string {
	return tea.Prettify(s)
}

func (s BindInterconnectionCidResponse) GoString() string {
	return s.String()
}

func (s *BindInterconnectionCidResponse) SetHeaders(v map[string]*string) *BindInterconnectionCidResponse {
	s.Headers = v
	return s
}

func (s *BindInterconnectionCidResponse) SetBody(v *BindInterconnectionCidResponseBody) *BindInterconnectionCidResponse {
	s.Body = v
	return s
}

type InitTenantRequest struct {
	Request *InitTenantRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
}

func (s InitTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s InitTenantRequest) GoString() string {
	return s.String()
}

func (s *InitTenantRequest) SetRequest(v *InitTenantRequestRequest) *InitTenantRequest {
	s.Request = v
	return s
}

type InitTenantRequestRequest struct {
	// 应用appKey
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
}

func (s InitTenantRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s InitTenantRequestRequest) GoString() string {
	return s.String()
}

func (s *InitTenantRequestRequest) SetDomain(v string) *InitTenantRequestRequest {
	s.Domain = &v
	return s
}

type InitTenantResponseBody struct {
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 是否初始化成功
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitTenantResponseBody) GoString() string {
	return s.String()
}

func (s *InitTenantResponseBody) SetResponseSuccess(v bool) *InitTenantResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *InitTenantResponseBody) SetErrorCode(v string) *InitTenantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InitTenantResponseBody) SetErrorMsg(v string) *InitTenantResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InitTenantResponseBody) SetResult(v bool) *InitTenantResponseBody {
	s.Result = &v
	return s
}

func (s *InitTenantResponseBody) SetRequestId(v string) *InitTenantResponseBody {
	s.RequestId = &v
	return s
}

type InitTenantResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s InitTenantResponse) GoString() string {
	return s.String()
}

func (s *InitTenantResponse) SetHeaders(v map[string]*string) *InitTenantResponse {
	s.Headers = v
	return s
}

func (s *InitTenantResponse) SetBody(v *InitTenantResponseBody) *InitTenantResponse {
	s.Body = v
	return s
}

type ImportGroupChatMemberRequest struct {
	// AppId
	AppId         *string                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *ImportGroupChatMemberRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s ImportGroupChatMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatMemberRequest) GoString() string {
	return s.String()
}

func (s *ImportGroupChatMemberRequest) SetAppId(v string) *ImportGroupChatMemberRequest {
	s.AppId = &v
	return s
}

func (s *ImportGroupChatMemberRequest) SetRequestParams(v *ImportGroupChatMemberRequestRequestParams) *ImportGroupChatMemberRequest {
	s.RequestParams = v
	return s
}

type ImportGroupChatMemberRequestRequestParams struct {
	// 群ID
	AppCid           *string                                                      `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	GroupChatMembers []*ImportGroupChatMemberRequestRequestParamsGroupChatMembers `json:"GroupChatMembers,omitempty" xml:"GroupChatMembers,omitempty" type:"Repeated"`
}

func (s ImportGroupChatMemberRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatMemberRequestRequestParams) GoString() string {
	return s.String()
}

func (s *ImportGroupChatMemberRequestRequestParams) SetAppCid(v string) *ImportGroupChatMemberRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParams) SetGroupChatMembers(v []*ImportGroupChatMemberRequestRequestParamsGroupChatMembers) *ImportGroupChatMemberRequestRequestParams {
	s.GroupChatMembers = v
	return s
}

type ImportGroupChatMemberRequestRequestParamsGroupChatMembers struct {
	// 用户ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 成员权限
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 昵称
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// 是否置顶
	Top *bool `json:"Top,omitempty" xml:"Top,omitempty"`
	// 未读数
	RedPoint *int64 `json:"RedPoint,omitempty" xml:"RedPoint,omitempty"`
	// 是否免打扰
	Mute *bool `json:"Mute,omitempty" xml:"Mute,omitempty"`
	// 是否可见
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
	// 入群时间戳
	JoinTime *int64 `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// 最后修改时间
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 自定义信息
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s ImportGroupChatMemberRequestRequestParamsGroupChatMembers) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatMemberRequestRequestParamsGroupChatMembers) GoString() string {
	return s.String()
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetAppUid(v string) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.AppUid = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetRole(v int64) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.Role = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetNick(v string) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.Nick = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetTop(v bool) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.Top = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetRedPoint(v int64) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.RedPoint = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetMute(v bool) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.Mute = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetVisible(v bool) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.Visible = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetJoinTime(v int64) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.JoinTime = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetModifyTime(v int64) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.ModifyTime = &v
	return s
}

func (s *ImportGroupChatMemberRequestRequestParamsGroupChatMembers) SetExtensions(v map[string]*string) *ImportGroupChatMemberRequestRequestParamsGroupChatMembers {
	s.Extensions = v
	return s
}

type ImportGroupChatMemberShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s ImportGroupChatMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportGroupChatMemberShrinkRequest) SetAppId(v string) *ImportGroupChatMemberShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ImportGroupChatMemberShrinkRequest) SetRequestParamsShrink(v string) *ImportGroupChatMemberShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type ImportGroupChatMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ImportGroupChatMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ImportGroupChatMemberResponseBody) SetRequestId(v string) *ImportGroupChatMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportGroupChatMemberResponseBody) SetCode(v string) *ImportGroupChatMemberResponseBody {
	s.Code = &v
	return s
}

func (s *ImportGroupChatMemberResponseBody) SetMessage(v string) *ImportGroupChatMemberResponseBody {
	s.Message = &v
	return s
}

type ImportGroupChatMemberResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportGroupChatMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportGroupChatMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatMemberResponse) GoString() string {
	return s.String()
}

func (s *ImportGroupChatMemberResponse) SetHeaders(v map[string]*string) *ImportGroupChatMemberResponse {
	s.Headers = v
	return s
}

func (s *ImportGroupChatMemberResponse) SetBody(v *ImportGroupChatMemberResponseBody) *ImportGroupChatMemberResponse {
	s.Body = v
	return s
}

type ListGroupSilenceMembersRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言添加白名单请求体
	RequestParams *ListGroupSilenceMembersRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s ListGroupSilenceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSilenceMembersRequest) GoString() string {
	return s.String()
}

func (s *ListGroupSilenceMembersRequest) SetAppId(v string) *ListGroupSilenceMembersRequest {
	s.AppId = &v
	return s
}

func (s *ListGroupSilenceMembersRequest) SetRequestParams(v *ListGroupSilenceMembersRequestRequestParams) *ListGroupSilenceMembersRequest {
	s.RequestParams = v
	return s
}

type ListGroupSilenceMembersRequestRequestParams struct {
	// 操作者
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 群会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
}

func (s ListGroupSilenceMembersRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSilenceMembersRequestRequestParams) GoString() string {
	return s.String()
}

func (s *ListGroupSilenceMembersRequestRequestParams) SetOperatorAppUid(v string) *ListGroupSilenceMembersRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *ListGroupSilenceMembersRequestRequestParams) SetAppCid(v string) *ListGroupSilenceMembersRequestRequestParams {
	s.AppCid = &v
	return s
}

type ListGroupSilenceMembersShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 群禁言添加白名单请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s ListGroupSilenceMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSilenceMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGroupSilenceMembersShrinkRequest) SetAppId(v string) *ListGroupSilenceMembersShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListGroupSilenceMembersShrinkRequest) SetRequestParamsShrink(v string) *ListGroupSilenceMembersShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type ListGroupSilenceMembersResponseBody struct {
	// Id of the request
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *ListGroupSilenceMembersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListGroupSilenceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSilenceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupSilenceMembersResponseBody) SetRequestId(v string) *ListGroupSilenceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupSilenceMembersResponseBody) SetCode(v string) *ListGroupSilenceMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListGroupSilenceMembersResponseBody) SetMessage(v string) *ListGroupSilenceMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListGroupSilenceMembersResponseBody) SetResult(v *ListGroupSilenceMembersResponseBodyResult) *ListGroupSilenceMembersResponseBody {
	s.Result = v
	return s
}

type ListGroupSilenceMembersResponseBodyResult struct {
	// 群会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 禁言白名单
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
	// 禁言黑名单用户及对应禁言时长
	Blacklist map[string]*int64 `json:"Blacklist,omitempty" xml:"Blacklist,omitempty"`
}

func (s ListGroupSilenceMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSilenceMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupSilenceMembersResponseBodyResult) SetAppCid(v string) *ListGroupSilenceMembersResponseBodyResult {
	s.AppCid = &v
	return s
}

func (s *ListGroupSilenceMembersResponseBodyResult) SetWhitelist(v []*string) *ListGroupSilenceMembersResponseBodyResult {
	s.Whitelist = v
	return s
}

func (s *ListGroupSilenceMembersResponseBodyResult) SetBlacklist(v map[string]*int64) *ListGroupSilenceMembersResponseBodyResult {
	s.Blacklist = v
	return s
}

type ListGroupSilenceMembersResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupSilenceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupSilenceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupSilenceMembersResponse) GoString() string {
	return s.String()
}

func (s *ListGroupSilenceMembersResponse) SetHeaders(v map[string]*string) *ListGroupSilenceMembersResponse {
	s.Headers = v
	return s
}

func (s *ListGroupSilenceMembersResponse) SetBody(v *ListGroupSilenceMembersResponseBody) *ListGroupSilenceMembersResponse {
	s.Body = v
	return s
}

type RemoveGroupExtensionByKeysRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 移除群聊拓展字段请求实体
	RequestParams *RemoveGroupExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RemoveGroupExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupExtensionByKeysRequest) SetAppId(v string) *RemoveGroupExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupExtensionByKeysRequest) SetRequestParams(v *RemoveGroupExtensionByKeysRequestRequestParams) *RemoveGroupExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type RemoveGroupExtensionByKeysRequestRequestParams struct {
	// 会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 拓展字段的key
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s RemoveGroupExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RemoveGroupExtensionByKeysRequestRequestParams) SetAppCid(v string) *RemoveGroupExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RemoveGroupExtensionByKeysRequestRequestParams) SetKeys(v []*string) *RemoveGroupExtensionByKeysRequestRequestParams {
	s.Keys = v
	return s
}

type RemoveGroupExtensionByKeysShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 移除群聊拓展字段请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RemoveGroupExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupExtensionByKeysShrinkRequest) SetAppId(v string) *RemoveGroupExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveGroupExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *RemoveGroupExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RemoveGroupExtensionByKeysResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveGroupExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupExtensionByKeysResponseBody) SetRequestId(v string) *RemoveGroupExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGroupExtensionByKeysResponseBody) SetCode(v string) *RemoveGroupExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveGroupExtensionByKeysResponseBody) SetMessage(v string) *RemoveGroupExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type RemoveGroupExtensionByKeysResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveGroupExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveGroupExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupExtensionByKeysResponse) SetHeaders(v map[string]*string) *RemoveGroupExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupExtensionByKeysResponse) SetBody(v *RemoveGroupExtensionByKeysResponseBody) *RemoveGroupExtensionByKeysResponse {
	s.Body = v
	return s
}

type SetGroupMemberExtensionByKeysRequest struct {
	// App ID, IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 设置群成员扩展信息的请求体
	RequestParams *SetGroupMemberExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SetGroupMemberExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGroupMemberExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *SetGroupMemberExtensionByKeysRequest) SetAppId(v string) *SetGroupMemberExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *SetGroupMemberExtensionByKeysRequest) SetRequestParams(v *SetGroupMemberExtensionByKeysRequestRequestParams) *SetGroupMemberExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type SetGroupMemberExtensionByKeysRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 用户ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 扩展信息
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s SetGroupMemberExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SetGroupMemberExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SetGroupMemberExtensionByKeysRequestRequestParams) SetAppCid(v string) *SetGroupMemberExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *SetGroupMemberExtensionByKeysRequestRequestParams) SetAppUid(v string) *SetGroupMemberExtensionByKeysRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *SetGroupMemberExtensionByKeysRequestRequestParams) SetExtensions(v map[string]*string) *SetGroupMemberExtensionByKeysRequestRequestParams {
	s.Extensions = v
	return s
}

type SetGroupMemberExtensionByKeysShrinkRequest struct {
	// App ID, IMPaaS租户的ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 设置群成员扩展信息的请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SetGroupMemberExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGroupMemberExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetGroupMemberExtensionByKeysShrinkRequest) SetAppId(v string) *SetGroupMemberExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SetGroupMemberExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *SetGroupMemberExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SetGroupMemberExtensionByKeysResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SetGroupMemberExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGroupMemberExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *SetGroupMemberExtensionByKeysResponseBody) SetRequestId(v string) *SetGroupMemberExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGroupMemberExtensionByKeysResponseBody) SetCode(v string) *SetGroupMemberExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *SetGroupMemberExtensionByKeysResponseBody) SetMessage(v string) *SetGroupMemberExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type SetGroupMemberExtensionByKeysResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGroupMemberExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGroupMemberExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGroupMemberExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *SetGroupMemberExtensionByKeysResponse) SetHeaders(v map[string]*string) *SetGroupMemberExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *SetGroupMemberExtensionByKeysResponse) SetBody(v *SetGroupMemberExtensionByKeysResponseBody) *SetGroupMemberExtensionByKeysResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建群聊请求实体
	RequestParams *CreateGroupRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetAppId(v string) *CreateGroupRequest {
	s.AppId = &v
	return s
}

func (s *CreateGroupRequest) SetRequestParams(v *CreateGroupRequestRequestParams) *CreateGroupRequest {
	s.RequestParams = v
	return s
}

type CreateGroupRequestRequestParams struct {
	// UUID(不可重复)
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// 创建者
	CreatorAppUid *string `json:"CreatorAppUid,omitempty" xml:"CreatorAppUid,omitempty"`
	// 群标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 图标的id
	IconMediaId *string `json:"IconMediaId,omitempty" xml:"IconMediaId,omitempty"`
	// 拓展字段
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// 初始化成员
	InitMembers []*CreateGroupRequestRequestParamsInitMembers `json:"InitMembers,omitempty" xml:"InitMembers,omitempty" type:"Repeated"`
}

func (s CreateGroupRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequestRequestParams) GoString() string {
	return s.String()
}

func (s *CreateGroupRequestRequestParams) SetUuid(v string) *CreateGroupRequestRequestParams {
	s.Uuid = &v
	return s
}

func (s *CreateGroupRequestRequestParams) SetCreatorAppUid(v string) *CreateGroupRequestRequestParams {
	s.CreatorAppUid = &v
	return s
}

func (s *CreateGroupRequestRequestParams) SetTitle(v string) *CreateGroupRequestRequestParams {
	s.Title = &v
	return s
}

func (s *CreateGroupRequestRequestParams) SetIconMediaId(v string) *CreateGroupRequestRequestParams {
	s.IconMediaId = &v
	return s
}

func (s *CreateGroupRequestRequestParams) SetExtensions(v map[string]*string) *CreateGroupRequestRequestParams {
	s.Extensions = v
	return s
}

func (s *CreateGroupRequestRequestParams) SetInitMembers(v []*CreateGroupRequestRequestParamsInitMembers) *CreateGroupRequestRequestParams {
	s.InitMembers = v
	return s
}

type CreateGroupRequestRequestParamsInitMembers struct {
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 1群主，2管理员，3普通
	Role *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// unix时间毫秒数
	JoinTime   *int64             `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s CreateGroupRequestRequestParamsInitMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequestRequestParamsInitMembers) GoString() string {
	return s.String()
}

func (s *CreateGroupRequestRequestParamsInitMembers) SetAppUid(v string) *CreateGroupRequestRequestParamsInitMembers {
	s.AppUid = &v
	return s
}

func (s *CreateGroupRequestRequestParamsInitMembers) SetRole(v int32) *CreateGroupRequestRequestParamsInitMembers {
	s.Role = &v
	return s
}

func (s *CreateGroupRequestRequestParamsInitMembers) SetNick(v string) *CreateGroupRequestRequestParamsInitMembers {
	s.Nick = &v
	return s
}

func (s *CreateGroupRequestRequestParamsInitMembers) SetJoinTime(v int64) *CreateGroupRequestRequestParamsInitMembers {
	s.JoinTime = &v
	return s
}

func (s *CreateGroupRequestRequestParamsInitMembers) SetExtensions(v map[string]*string) *CreateGroupRequestRequestParamsInitMembers {
	s.Extensions = v
	return s
}

type CreateGroupShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建群聊请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s CreateGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupShrinkRequest) SetAppId(v string) *CreateGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateGroupShrinkRequest) SetRequestParamsShrink(v string) *CreateGroupShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type CreateGroupResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *CreateGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupResponseBody) SetCode(v string) *CreateGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupResponseBody) SetMessage(v string) *CreateGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupResponseBody) SetResult(v *CreateGroupResponseBodyResult) *CreateGroupResponseBody {
	s.Result = v
	return s
}

type CreateGroupResponseBodyResult struct {
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
}

func (s CreateGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBodyResult) SetAppCid(v string) *CreateGroupResponseBodyResult {
	s.AppCid = &v
	return s
}

type CreateGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type GetMessageByIdRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 请求实体
	RequestParams *GetMessageByIdRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s GetMessageByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMessageByIdRequest) GoString() string {
	return s.String()
}

func (s *GetMessageByIdRequest) SetAppId(v string) *GetMessageByIdRequest {
	s.AppId = &v
	return s
}

func (s *GetMessageByIdRequest) SetRequestParams(v *GetMessageByIdRequestRequestParams) *GetMessageByIdRequest {
	s.RequestParams = v
	return s
}

type GetMessageByIdRequestRequestParams struct {
	// 消息Id
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s GetMessageByIdRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetMessageByIdRequestRequestParams) GoString() string {
	return s.String()
}

func (s *GetMessageByIdRequestRequestParams) SetMsgId(v string) *GetMessageByIdRequestRequestParams {
	s.MsgId = &v
	return s
}

type GetMessageByIdShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s GetMessageByIdShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMessageByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMessageByIdShrinkRequest) SetAppId(v string) *GetMessageByIdShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetMessageByIdShrinkRequest) SetRequestParamsShrink(v string) *GetMessageByIdShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type GetMessageByIdResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *GetMessageByIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetMessageByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMessageByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageByIdResponseBody) SetRequestId(v string) *GetMessageByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageByIdResponseBody) SetCode(v string) *GetMessageByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetMessageByIdResponseBody) SetMessage(v string) *GetMessageByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetMessageByIdResponseBody) SetResult(v *GetMessageByIdResponseBodyResult) *GetMessageByIdResponseBody {
	s.Result = v
	return s
}

type GetMessageByIdResponseBodyResult struct {
	// 消息Id
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// 会话Id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 会话类型
	ConversationType *int32 `json:"ConversationType,omitempty" xml:"ConversationType,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 发送者的用户Id
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 消息类型
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// 消息体
	Content    *string            `json:"Content,omitempty" xml:"Content,omitempty"`
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s GetMessageByIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMessageByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMessageByIdResponseBodyResult) SetMsgId(v string) *GetMessageByIdResponseBodyResult {
	s.MsgId = &v
	return s
}

func (s *GetMessageByIdResponseBodyResult) SetAppCid(v string) *GetMessageByIdResponseBodyResult {
	s.AppCid = &v
	return s
}

func (s *GetMessageByIdResponseBodyResult) SetConversationType(v int32) *GetMessageByIdResponseBodyResult {
	s.ConversationType = &v
	return s
}

func (s *GetMessageByIdResponseBodyResult) SetCreateTime(v int64) *GetMessageByIdResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetMessageByIdResponseBodyResult) SetSenderId(v string) *GetMessageByIdResponseBodyResult {
	s.SenderId = &v
	return s
}

func (s *GetMessageByIdResponseBodyResult) SetContentType(v int32) *GetMessageByIdResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *GetMessageByIdResponseBodyResult) SetContent(v string) *GetMessageByIdResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetMessageByIdResponseBodyResult) SetExtensions(v map[string]*string) *GetMessageByIdResponseBodyResult {
	s.Extensions = v
	return s
}

type GetMessageByIdResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMessageByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMessageByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMessageByIdResponse) GoString() string {
	return s.String()
}

func (s *GetMessageByIdResponse) SetHeaders(v map[string]*string) *GetMessageByIdResponse {
	s.Headers = v
	return s
}

func (s *GetMessageByIdResponse) SetBody(v *GetMessageByIdResponseBody) *GetMessageByIdResponse {
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
	// 应用appKey
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// 房间id
	RoomId *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
	// 操作人id
	OpenId *string `json:"openId,omitempty" xml:"openId,omitempty"`
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

func (s *DestroyRoomRequestRequest) SetOpenId(v string) *DestroyRoomRequestRequest {
	s.OpenId = &v
	return s
}

type DestroyRoomResponseBody struct {
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否销毁成功
	Result          *bool `json:"result,omitempty" xml:"result,omitempty"`
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
}

func (s DestroyRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DestroyRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DestroyRoomResponseBody) SetErrorCode(v string) *DestroyRoomResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DestroyRoomResponseBody) SetErrorMsg(v string) *DestroyRoomResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DestroyRoomResponseBody) SetRequestId(v string) *DestroyRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *DestroyRoomResponseBody) SetResult(v bool) *DestroyRoomResponseBody {
	s.Result = &v
	return s
}

func (s *DestroyRoomResponseBody) SetResponseSuccess(v bool) *DestroyRoomResponseBody {
	s.ResponseSuccess = &v
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

type KickOffRequest struct {
	// AppId
	AppId         *string                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *KickOffRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s KickOffRequest) String() string {
	return tea.Prettify(s)
}

func (s KickOffRequest) GoString() string {
	return s.String()
}

func (s *KickOffRequest) SetAppId(v string) *KickOffRequest {
	s.AppId = &v
	return s
}

func (s *KickOffRequest) SetRequestParams(v *KickOffRequestRequestParams) *KickOffRequest {
	s.RequestParams = v
	return s
}

type KickOffRequestRequestParams struct {
	// 用户ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 被踢下线的App的AppKey列表，为空时全部踢下线
	AppKeys []*string `json:"AppKeys,omitempty" xml:"AppKeys,omitempty" type:"Repeated"`
	// 下线文案
	Information *string `json:"Information,omitempty" xml:"Information,omitempty"`
}

func (s KickOffRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s KickOffRequestRequestParams) GoString() string {
	return s.String()
}

func (s *KickOffRequestRequestParams) SetAppUid(v string) *KickOffRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *KickOffRequestRequestParams) SetAppKeys(v []*string) *KickOffRequestRequestParams {
	s.AppKeys = v
	return s
}

func (s *KickOffRequestRequestParams) SetInformation(v string) *KickOffRequestRequestParams {
	s.Information = &v
	return s
}

type KickOffShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s KickOffShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s KickOffShrinkRequest) GoString() string {
	return s.String()
}

func (s *KickOffShrinkRequest) SetAppId(v string) *KickOffShrinkRequest {
	s.AppId = &v
	return s
}

func (s *KickOffShrinkRequest) SetRequestParamsShrink(v string) *KickOffShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type KickOffResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s KickOffResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KickOffResponseBody) GoString() string {
	return s.String()
}

func (s *KickOffResponseBody) SetRequestId(v string) *KickOffResponseBody {
	s.RequestId = &v
	return s
}

func (s *KickOffResponseBody) SetCode(v string) *KickOffResponseBody {
	s.Code = &v
	return s
}

func (s *KickOffResponseBody) SetMessage(v string) *KickOffResponseBody {
	s.Message = &v
	return s
}

type KickOffResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KickOffResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KickOffResponse) String() string {
	return tea.Prettify(s)
}

func (s KickOffResponse) GoString() string {
	return s.String()
}

func (s *KickOffResponse) SetHeaders(v map[string]*string) *KickOffResponse {
	s.Headers = v
	return s
}

func (s *KickOffResponse) SetBody(v *KickOffResponseBody) *KickOffResponse {
	s.Body = v
	return s
}

type ListCallbackApiIdsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 业务错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 网络错误码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回结果
	Result *ListCallbackApiIdsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListCallbackApiIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCallbackApiIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallbackApiIdsResponseBody) SetRequestId(v string) *ListCallbackApiIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallbackApiIdsResponseBody) SetSuccess(v bool) *ListCallbackApiIdsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCallbackApiIdsResponseBody) SetCode(v string) *ListCallbackApiIdsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallbackApiIdsResponseBody) SetHttpStatusCode(v int32) *ListCallbackApiIdsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCallbackApiIdsResponseBody) SetMessage(v string) *ListCallbackApiIdsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCallbackApiIdsResponseBody) SetResult(v *ListCallbackApiIdsResponseBodyResult) *ListCallbackApiIdsResponseBody {
	s.Result = v
	return s
}

type ListCallbackApiIdsResponseBodyResult struct {
	// 回调列表
	Spis map[string]*bool `json:"Spis,omitempty" xml:"Spis,omitempty"`
	// 事件输出列表
	Events map[string]*bool `json:"Events,omitempty" xml:"Events,omitempty"`
}

func (s ListCallbackApiIdsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCallbackApiIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCallbackApiIdsResponseBodyResult) SetSpis(v map[string]*bool) *ListCallbackApiIdsResponseBodyResult {
	s.Spis = v
	return s
}

func (s *ListCallbackApiIdsResponseBodyResult) SetEvents(v map[string]*bool) *ListCallbackApiIdsResponseBodyResult {
	s.Events = v
	return s
}

type ListCallbackApiIdsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCallbackApiIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCallbackApiIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCallbackApiIdsResponse) GoString() string {
	return s.String()
}

func (s *ListCallbackApiIdsResponse) SetHeaders(v map[string]*string) *ListCallbackApiIdsResponse {
	s.Headers = v
	return s
}

func (s *ListCallbackApiIdsResponse) SetBody(v *ListCallbackApiIdsResponseBody) *ListCallbackApiIdsResponse {
	s.Body = v
	return s
}

type SetMessageReadRequest struct {
	// AppId
	AppId         *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *SetMessageReadRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SetMessageReadRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMessageReadRequest) GoString() string {
	return s.String()
}

func (s *SetMessageReadRequest) SetAppId(v string) *SetMessageReadRequest {
	s.AppId = &v
	return s
}

func (s *SetMessageReadRequest) SetRequestParams(v *SetMessageReadRequestRequestParams) *SetMessageReadRequest {
	s.RequestParams = v
	return s
}

type SetMessageReadRequestRequestParams struct {
	// 操作者ID
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 消息ID
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SetMessageReadRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SetMessageReadRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SetMessageReadRequestRequestParams) SetAppUid(v string) *SetMessageReadRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *SetMessageReadRequestRequestParams) SetMsgId(v string) *SetMessageReadRequestRequestParams {
	s.MsgId = &v
	return s
}

type SetMessageReadShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SetMessageReadShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMessageReadShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetMessageReadShrinkRequest) SetAppId(v string) *SetMessageReadShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SetMessageReadShrinkRequest) SetRequestParamsShrink(v string) *SetMessageReadShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SetMessageReadResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SetMessageReadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMessageReadResponseBody) GoString() string {
	return s.String()
}

func (s *SetMessageReadResponseBody) SetRequestId(v string) *SetMessageReadResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMessageReadResponseBody) SetCode(v string) *SetMessageReadResponseBody {
	s.Code = &v
	return s
}

func (s *SetMessageReadResponseBody) SetMessage(v string) *SetMessageReadResponseBody {
	s.Message = &v
	return s
}

type SetMessageReadResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetMessageReadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetMessageReadResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMessageReadResponse) GoString() string {
	return s.String()
}

func (s *SetMessageReadResponse) SetHeaders(v map[string]*string) *SetMessageReadResponse {
	s.Headers = v
	return s
}

func (s *SetMessageReadResponse) SetBody(v *SetMessageReadResponseBody) *SetMessageReadResponse {
	s.Body = v
	return s
}

type UpdateMsgRecallIntervalRequest struct {
	// 请求
	RequestParams *UpdateMsgRecallIntervalRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s UpdateMsgRecallIntervalRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsgRecallIntervalRequest) GoString() string {
	return s.String()
}

func (s *UpdateMsgRecallIntervalRequest) SetRequestParams(v *UpdateMsgRecallIntervalRequestRequestParams) *UpdateMsgRecallIntervalRequest {
	s.RequestParams = v
	return s
}

func (s *UpdateMsgRecallIntervalRequest) SetAppId(v string) *UpdateMsgRecallIntervalRequest {
	s.AppId = &v
	return s
}

type UpdateMsgRecallIntervalRequestRequestParams struct {
	// 消息撤回时间间隔
	ClientMsgRecallIntervalMinute *int64 `json:"ClientMsgRecallIntervalMinute,omitempty" xml:"ClientMsgRecallIntervalMinute,omitempty"`
}

func (s UpdateMsgRecallIntervalRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsgRecallIntervalRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateMsgRecallIntervalRequestRequestParams) SetClientMsgRecallIntervalMinute(v int64) *UpdateMsgRecallIntervalRequestRequestParams {
	s.ClientMsgRecallIntervalMinute = &v
	return s
}

type UpdateMsgRecallIntervalShrinkRequest struct {
	// 请求
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s UpdateMsgRecallIntervalShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsgRecallIntervalShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMsgRecallIntervalShrinkRequest) SetRequestParamsShrink(v string) *UpdateMsgRecallIntervalShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

func (s *UpdateMsgRecallIntervalShrinkRequest) SetAppId(v string) *UpdateMsgRecallIntervalShrinkRequest {
	s.AppId = &v
	return s
}

type UpdateMsgRecallIntervalResponseBody struct {
	// desc
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// result
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateMsgRecallIntervalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsgRecallIntervalResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMsgRecallIntervalResponseBody) SetMessage(v string) *UpdateMsgRecallIntervalResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMsgRecallIntervalResponseBody) SetRequestId(v string) *UpdateMsgRecallIntervalResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMsgRecallIntervalResponseBody) SetHttpStatusCode(v int32) *UpdateMsgRecallIntervalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMsgRecallIntervalResponseBody) SetCode(v string) *UpdateMsgRecallIntervalResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMsgRecallIntervalResponseBody) SetSuccess(v bool) *UpdateMsgRecallIntervalResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMsgRecallIntervalResponseBody) SetResult(v string) *UpdateMsgRecallIntervalResponseBody {
	s.Result = &v
	return s
}

type UpdateMsgRecallIntervalResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMsgRecallIntervalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMsgRecallIntervalResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMsgRecallIntervalResponse) GoString() string {
	return s.String()
}

func (s *UpdateMsgRecallIntervalResponse) SetHeaders(v map[string]*string) *UpdateMsgRecallIntervalResponse {
	s.Headers = v
	return s
}

func (s *UpdateMsgRecallIntervalResponse) SetBody(v *UpdateMsgRecallIntervalResponseBody) *UpdateMsgRecallIntervalResponse {
	s.Body = v
	return s
}

type SendCustomMessageToRoomUsersRequest struct {
	// 请求参数的结构体。
	Request *SendCustomMessageToRoomUsersRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// 指定的消息接受者的用户ID列表，大小不得超过100。
	Receivers []*string `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
}

func (s SendCustomMessageToRoomUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToRoomUsersRequest) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToRoomUsersRequest) SetRequest(v *SendCustomMessageToRoomUsersRequestRequest) *SendCustomMessageToRoomUsersRequest {
	s.Request = v
	return s
}

func (s *SendCustomMessageToRoomUsersRequest) SetReceivers(v []*string) *SendCustomMessageToRoomUsersRequest {
	s.Receivers = v
	return s
}

type SendCustomMessageToRoomUsersRequestRequest struct {
	// 应用的appKey。
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 房间ID，由调用CreateRoom时返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 消息的发送者ID。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 消息的类型，由业务自定义，请传递100000以上。
	SubType *int32 `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// 消息体。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s SendCustomMessageToRoomUsersRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToRoomUsersRequestRequest) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToRoomUsersRequestRequest) SetDomain(v string) *SendCustomMessageToRoomUsersRequestRequest {
	s.Domain = &v
	return s
}

func (s *SendCustomMessageToRoomUsersRequestRequest) SetRoomId(v string) *SendCustomMessageToRoomUsersRequestRequest {
	s.RoomId = &v
	return s
}

func (s *SendCustomMessageToRoomUsersRequestRequest) SetSenderId(v string) *SendCustomMessageToRoomUsersRequestRequest {
	s.SenderId = &v
	return s
}

func (s *SendCustomMessageToRoomUsersRequestRequest) SetSubType(v int32) *SendCustomMessageToRoomUsersRequestRequest {
	s.SubType = &v
	return s
}

func (s *SendCustomMessageToRoomUsersRequestRequest) SetBody(v string) *SendCustomMessageToRoomUsersRequestRequest {
	s.Body = &v
	return s
}

type SendCustomMessageToRoomUsersResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功。
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码，请求异常时返回。
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 错误信息，请求异常时返回。
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 请求的返回结果。
	Result *SendCustomMessageToRoomUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendCustomMessageToRoomUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToRoomUsersResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToRoomUsersResponseBody) SetRequestId(v string) *SendCustomMessageToRoomUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomMessageToRoomUsersResponseBody) SetResponseSuccess(v bool) *SendCustomMessageToRoomUsersResponseBody {
	s.ResponseSuccess = &v
	return s
}

func (s *SendCustomMessageToRoomUsersResponseBody) SetErrorCode(v string) *SendCustomMessageToRoomUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendCustomMessageToRoomUsersResponseBody) SetErrorMessage(v string) *SendCustomMessageToRoomUsersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendCustomMessageToRoomUsersResponseBody) SetResult(v *SendCustomMessageToRoomUsersResponseBodyResult) *SendCustomMessageToRoomUsersResponseBody {
	s.Result = v
	return s
}

type SendCustomMessageToRoomUsersResponseBodyResult struct {
	// 消息的唯一ID标识。由数字、大小写字母组成，长度不超过20。
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendCustomMessageToRoomUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToRoomUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToRoomUsersResponseBodyResult) SetMessageId(v string) *SendCustomMessageToRoomUsersResponseBodyResult {
	s.MessageId = &v
	return s
}

type SendCustomMessageToRoomUsersResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCustomMessageToRoomUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCustomMessageToRoomUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToRoomUsersResponse) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToRoomUsersResponse) SetHeaders(v map[string]*string) *SendCustomMessageToRoomUsersResponse {
	s.Headers = v
	return s
}

func (s *SendCustomMessageToRoomUsersResponse) SetBody(v *SendCustomMessageToRoomUsersResponseBody) *SendCustomMessageToRoomUsersResponse {
	s.Body = v
	return s
}

type UpdateGroupTitleRequest struct {
	// AppId
	AppId         *string                               `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *UpdateGroupTitleRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s UpdateGroupTitleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupTitleRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupTitleRequest) SetAppId(v string) *UpdateGroupTitleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGroupTitleRequest) SetRequestParams(v *UpdateGroupTitleRequestRequestParams) *UpdateGroupTitleRequest {
	s.RequestParams = v
	return s
}

type UpdateGroupTitleRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 操作者用户ID
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 群聊标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateGroupTitleRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupTitleRequestRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateGroupTitleRequestRequestParams) SetAppCid(v string) *UpdateGroupTitleRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *UpdateGroupTitleRequestRequestParams) SetOperatorAppUid(v string) *UpdateGroupTitleRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *UpdateGroupTitleRequestRequestParams) SetTitle(v string) *UpdateGroupTitleRequestRequestParams {
	s.Title = &v
	return s
}

type UpdateGroupTitleShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s UpdateGroupTitleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupTitleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupTitleShrinkRequest) SetAppId(v string) *UpdateGroupTitleShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateGroupTitleShrinkRequest) SetRequestParamsShrink(v string) *UpdateGroupTitleShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type UpdateGroupTitleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateGroupTitleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupTitleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupTitleResponseBody) SetRequestId(v string) *UpdateGroupTitleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupTitleResponseBody) SetCode(v string) *UpdateGroupTitleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupTitleResponseBody) SetMessage(v string) *UpdateGroupTitleResponseBody {
	s.Message = &v
	return s
}

type UpdateGroupTitleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupTitleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupTitleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupTitleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupTitleResponse) SetHeaders(v map[string]*string) *UpdateGroupTitleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupTitleResponse) SetBody(v *UpdateGroupTitleResponseBody) *UpdateGroupTitleResponse {
	s.Body = v
	return s
}

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

type QueryInterconnectionCidMappingRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询请求体
	RequestParams *QueryInterconnectionCidMappingRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s QueryInterconnectionCidMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInterconnectionCidMappingRequest) GoString() string {
	return s.String()
}

func (s *QueryInterconnectionCidMappingRequest) SetAppId(v string) *QueryInterconnectionCidMappingRequest {
	s.AppId = &v
	return s
}

func (s *QueryInterconnectionCidMappingRequest) SetRequestParams(v *QueryInterconnectionCidMappingRequestRequestParams) *QueryInterconnectionCidMappingRequest {
	s.RequestParams = v
	return s
}

type QueryInterconnectionCidMappingRequestRequestParams struct {
	// 会话ID
	SrcCid *string `json:"SrcCid,omitempty" xml:"SrcCid,omitempty"`
	// 会话ID类型; 1: AIM会话ID 2: 钉钉会话ID
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryInterconnectionCidMappingRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s QueryInterconnectionCidMappingRequestRequestParams) GoString() string {
	return s.String()
}

func (s *QueryInterconnectionCidMappingRequestRequestParams) SetSrcCid(v string) *QueryInterconnectionCidMappingRequestRequestParams {
	s.SrcCid = &v
	return s
}

func (s *QueryInterconnectionCidMappingRequestRequestParams) SetType(v int64) *QueryInterconnectionCidMappingRequestRequestParams {
	s.Type = &v
	return s
}

type QueryInterconnectionCidMappingShrinkRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询请求体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s QueryInterconnectionCidMappingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInterconnectionCidMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInterconnectionCidMappingShrinkRequest) SetAppId(v string) *QueryInterconnectionCidMappingShrinkRequest {
	s.AppId = &v
	return s
}

func (s *QueryInterconnectionCidMappingShrinkRequest) SetRequestParamsShrink(v string) *QueryInterconnectionCidMappingShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type QueryInterconnectionCidMappingResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *QueryInterconnectionCidMappingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s QueryInterconnectionCidMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInterconnectionCidMappingResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInterconnectionCidMappingResponseBody) SetRequestId(v string) *QueryInterconnectionCidMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInterconnectionCidMappingResponseBody) SetCode(v string) *QueryInterconnectionCidMappingResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInterconnectionCidMappingResponseBody) SetMessage(v string) *QueryInterconnectionCidMappingResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInterconnectionCidMappingResponseBody) SetResult(v *QueryInterconnectionCidMappingResponseBodyResult) *QueryInterconnectionCidMappingResponseBody {
	s.Result = v
	return s
}

type QueryInterconnectionCidMappingResponseBodyResult struct {
	DstCid *string `json:"DstCid,omitempty" xml:"DstCid,omitempty"`
}

func (s QueryInterconnectionCidMappingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryInterconnectionCidMappingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryInterconnectionCidMappingResponseBodyResult) SetDstCid(v string) *QueryInterconnectionCidMappingResponseBodyResult {
	s.DstCid = &v
	return s
}

type QueryInterconnectionCidMappingResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInterconnectionCidMappingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInterconnectionCidMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInterconnectionCidMappingResponse) GoString() string {
	return s.String()
}

func (s *QueryInterconnectionCidMappingResponse) SetHeaders(v map[string]*string) *QueryInterconnectionCidMappingResponse {
	s.Headers = v
	return s
}

func (s *QueryInterconnectionCidMappingResponse) SetBody(v *QueryInterconnectionCidMappingResponseBody) *QueryInterconnectionCidMappingResponse {
	s.Body = v
	return s
}

type DismissGroupRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 解散群聊请求实体
	RequestParams *DismissGroupRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s DismissGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupRequest) GoString() string {
	return s.String()
}

func (s *DismissGroupRequest) SetAppId(v string) *DismissGroupRequest {
	s.AppId = &v
	return s
}

func (s *DismissGroupRequest) SetRequestParams(v *DismissGroupRequestRequestParams) *DismissGroupRequest {
	s.RequestParams = v
	return s
}

type DismissGroupRequestRequestParams struct {
	// 操作用户
	OperatorAppUid *string `json:"OperatorAppUid,omitempty" xml:"OperatorAppUid,omitempty"`
	// 会话id
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
}

func (s DismissGroupRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupRequestRequestParams) GoString() string {
	return s.String()
}

func (s *DismissGroupRequestRequestParams) SetOperatorAppUid(v string) *DismissGroupRequestRequestParams {
	s.OperatorAppUid = &v
	return s
}

func (s *DismissGroupRequestRequestParams) SetAppCid(v string) *DismissGroupRequestRequestParams {
	s.AppCid = &v
	return s
}

type DismissGroupShrinkRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 解散群聊请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s DismissGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *DismissGroupShrinkRequest) SetAppId(v string) *DismissGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DismissGroupShrinkRequest) SetRequestParamsShrink(v string) *DismissGroupShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type DismissGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DismissGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DismissGroupResponseBody) SetRequestId(v string) *DismissGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DismissGroupResponseBody) SetCode(v string) *DismissGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DismissGroupResponseBody) SetMessage(v string) *DismissGroupResponseBody {
	s.Message = &v
	return s
}

type DismissGroupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DismissGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DismissGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupResponse) GoString() string {
	return s.String()
}

func (s *DismissGroupResponse) SetHeaders(v map[string]*string) *DismissGroupResponse {
	s.Headers = v
	return s
}

func (s *DismissGroupResponse) SetBody(v *DismissGroupResponseBody) *DismissGroupResponse {
	s.Body = v
	return s
}

type ImportGroupChatConversationRequest struct {
	// AppId
	AppId         *string                                          `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *ImportGroupChatConversationRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s ImportGroupChatConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatConversationRequest) GoString() string {
	return s.String()
}

func (s *ImportGroupChatConversationRequest) SetAppId(v string) *ImportGroupChatConversationRequest {
	s.AppId = &v
	return s
}

func (s *ImportGroupChatConversationRequest) SetRequestParams(v *ImportGroupChatConversationRequestRequestParams) *ImportGroupChatConversationRequest {
	s.RequestParams = v
	return s
}

type ImportGroupChatConversationRequestRequestParams struct {
	// 唯一标识，用于重入
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// 群主uid
	OwnerAppUid *string `json:"OwnerAppUid,omitempty" xml:"OwnerAppUid,omitempty"`
	// 群标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 群头像
	IconMediaId *string `json:"IconMediaId,omitempty" xml:"IconMediaId,omitempty"`
	// 群上限
	MemberLimit *int64 `json:"MemberLimit,omitempty" xml:"MemberLimit,omitempty"`
	// 扩展字段
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 是否全员禁言
	SilenceAll *bool `json:"SilenceAll,omitempty" xml:"SilenceAll,omitempty"`
}

func (s ImportGroupChatConversationRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatConversationRequestRequestParams) GoString() string {
	return s.String()
}

func (s *ImportGroupChatConversationRequestRequestParams) SetUuid(v string) *ImportGroupChatConversationRequestRequestParams {
	s.Uuid = &v
	return s
}

func (s *ImportGroupChatConversationRequestRequestParams) SetOwnerAppUid(v string) *ImportGroupChatConversationRequestRequestParams {
	s.OwnerAppUid = &v
	return s
}

func (s *ImportGroupChatConversationRequestRequestParams) SetTitle(v string) *ImportGroupChatConversationRequestRequestParams {
	s.Title = &v
	return s
}

func (s *ImportGroupChatConversationRequestRequestParams) SetIconMediaId(v string) *ImportGroupChatConversationRequestRequestParams {
	s.IconMediaId = &v
	return s
}

func (s *ImportGroupChatConversationRequestRequestParams) SetMemberLimit(v int64) *ImportGroupChatConversationRequestRequestParams {
	s.MemberLimit = &v
	return s
}

func (s *ImportGroupChatConversationRequestRequestParams) SetExtensions(v map[string]*string) *ImportGroupChatConversationRequestRequestParams {
	s.Extensions = v
	return s
}

func (s *ImportGroupChatConversationRequestRequestParams) SetCreateTime(v int64) *ImportGroupChatConversationRequestRequestParams {
	s.CreateTime = &v
	return s
}

func (s *ImportGroupChatConversationRequestRequestParams) SetSilenceAll(v bool) *ImportGroupChatConversationRequestRequestParams {
	s.SilenceAll = &v
	return s
}

type ImportGroupChatConversationShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s ImportGroupChatConversationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatConversationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportGroupChatConversationShrinkRequest) SetAppId(v string) *ImportGroupChatConversationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ImportGroupChatConversationShrinkRequest) SetRequestParamsShrink(v string) *ImportGroupChatConversationShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type ImportGroupChatConversationResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Result    *ImportGroupChatConversationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ImportGroupChatConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ImportGroupChatConversationResponseBody) SetRequestId(v string) *ImportGroupChatConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportGroupChatConversationResponseBody) SetCode(v string) *ImportGroupChatConversationResponseBody {
	s.Code = &v
	return s
}

func (s *ImportGroupChatConversationResponseBody) SetMessage(v string) *ImportGroupChatConversationResponseBody {
	s.Message = &v
	return s
}

func (s *ImportGroupChatConversationResponseBody) SetResult(v *ImportGroupChatConversationResponseBodyResult) *ImportGroupChatConversationResponseBody {
	s.Result = v
	return s
}

type ImportGroupChatConversationResponseBodyResult struct {
	// 群ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
}

func (s ImportGroupChatConversationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatConversationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ImportGroupChatConversationResponseBodyResult) SetAppCid(v string) *ImportGroupChatConversationResponseBodyResult {
	s.AppCid = &v
	return s
}

type ImportGroupChatConversationResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportGroupChatConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportGroupChatConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatConversationResponse) GoString() string {
	return s.String()
}

func (s *ImportGroupChatConversationResponse) SetHeaders(v map[string]*string) *ImportGroupChatConversationResponse {
	s.Headers = v
	return s
}

func (s *ImportGroupChatConversationResponse) SetBody(v *ImportGroupChatConversationResponseBody) *ImportGroupChatConversationResponse {
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
	// 应用appKey
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// 创建者id
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 创建者昵称
	OwnerNick *string `json:"ownerNick,omitempty" xml:"ownerNick,omitempty"`
	// 创建房间的标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
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

func (s *CreateRoomRequestRequest) SetOwnerId(v string) *CreateRoomRequestRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRoomRequestRequest) SetOwnerNick(v string) *CreateRoomRequestRequest {
	s.OwnerNick = &v
	return s
}

func (s *CreateRoomRequestRequest) SetTitle(v string) *CreateRoomRequestRequest {
	s.Title = &v
	return s
}

type CreateRoomResponseBody struct {
	ResponseSuccess *bool `json:"ResponseSuccess,omitempty" xml:"ResponseSuccess,omitempty"`
	// 错误码
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg  *string                       `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *CreateRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponseBody) GoString() string {
	return s.String()
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

func (s *CreateRoomResponseBody) SetResult(v *CreateRoomResponseBodyResult) *CreateRoomResponseBody {
	s.Result = v
	return s
}

func (s *CreateRoomResponseBody) SetRequestId(v string) *CreateRoomResponseBody {
	s.RequestId = &v
	return s
}

type CreateRoomResponseBodyResult struct {
	// 房间id
	RoomId *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
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

type RemoveUserConversationExtensionByKeysRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 移除用户拓展字段请求实体
	RequestParams *RemoveUserConversationExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s RemoveUserConversationExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserConversationExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserConversationExtensionByKeysRequest) SetAppId(v string) *RemoveUserConversationExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *RemoveUserConversationExtensionByKeysRequest) SetRequestParams(v *RemoveUserConversationExtensionByKeysRequestRequestParams) *RemoveUserConversationExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type RemoveUserConversationExtensionByKeysRequestRequestParams struct {
	// 用户id
	AppUid *string `json:"AppUid,omitempty" xml:"AppUid,omitempty"`
	// 会话id
	AppCid *string   `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	Keys   []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s RemoveUserConversationExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserConversationExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *RemoveUserConversationExtensionByKeysRequestRequestParams) SetAppUid(v string) *RemoveUserConversationExtensionByKeysRequestRequestParams {
	s.AppUid = &v
	return s
}

func (s *RemoveUserConversationExtensionByKeysRequestRequestParams) SetAppCid(v string) *RemoveUserConversationExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *RemoveUserConversationExtensionByKeysRequestRequestParams) SetKeys(v []*string) *RemoveUserConversationExtensionByKeysRequestRequestParams {
	s.Keys = v
	return s
}

type RemoveUserConversationExtensionByKeysShrinkRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 移除用户拓展字段请求实体
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s RemoveUserConversationExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserConversationExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserConversationExtensionByKeysShrinkRequest) SetAppId(v string) *RemoveUserConversationExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveUserConversationExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *RemoveUserConversationExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type RemoveUserConversationExtensionByKeysResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveUserConversationExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserConversationExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserConversationExtensionByKeysResponseBody) SetRequestId(v string) *RemoveUserConversationExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserConversationExtensionByKeysResponseBody) SetCode(v string) *RemoveUserConversationExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUserConversationExtensionByKeysResponseBody) SetMessage(v string) *RemoveUserConversationExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type RemoveUserConversationExtensionByKeysResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUserConversationExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUserConversationExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserConversationExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserConversationExtensionByKeysResponse) SetHeaders(v map[string]*string) *RemoveUserConversationExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserConversationExtensionByKeysResponse) SetBody(v *RemoveUserConversationExtensionByKeysResponseBody) *RemoveUserConversationExtensionByKeysResponse {
	s.Body = v
	return s
}

type SetMessageExtensionByKeysRequest struct {
	// AppId
	AppId         *string                                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParams *SetMessageExtensionByKeysRequestRequestParams `json:"RequestParams,omitempty" xml:"RequestParams,omitempty" type:"Struct"`
}

func (s SetMessageExtensionByKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMessageExtensionByKeysRequest) GoString() string {
	return s.String()
}

func (s *SetMessageExtensionByKeysRequest) SetAppId(v string) *SetMessageExtensionByKeysRequest {
	s.AppId = &v
	return s
}

func (s *SetMessageExtensionByKeysRequest) SetRequestParams(v *SetMessageExtensionByKeysRequestRequestParams) *SetMessageExtensionByKeysRequest {
	s.RequestParams = v
	return s
}

type SetMessageExtensionByKeysRequestRequestParams struct {
	// 会话ID
	AppCid *string `json:"AppCid,omitempty" xml:"AppCid,omitempty"`
	// 消息ID
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// 需设置的K-V对
	Extensions map[string]*string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
}

func (s SetMessageExtensionByKeysRequestRequestParams) String() string {
	return tea.Prettify(s)
}

func (s SetMessageExtensionByKeysRequestRequestParams) GoString() string {
	return s.String()
}

func (s *SetMessageExtensionByKeysRequestRequestParams) SetAppCid(v string) *SetMessageExtensionByKeysRequestRequestParams {
	s.AppCid = &v
	return s
}

func (s *SetMessageExtensionByKeysRequestRequestParams) SetMsgId(v string) *SetMessageExtensionByKeysRequestRequestParams {
	s.MsgId = &v
	return s
}

func (s *SetMessageExtensionByKeysRequestRequestParams) SetExtensions(v map[string]*string) *SetMessageExtensionByKeysRequestRequestParams {
	s.Extensions = v
	return s
}

type SetMessageExtensionByKeysShrinkRequest struct {
	// AppId
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestParamsShrink *string `json:"RequestParams,omitempty" xml:"RequestParams,omitempty"`
}

func (s SetMessageExtensionByKeysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMessageExtensionByKeysShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetMessageExtensionByKeysShrinkRequest) SetAppId(v string) *SetMessageExtensionByKeysShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SetMessageExtensionByKeysShrinkRequest) SetRequestParamsShrink(v string) *SetMessageExtensionByKeysShrinkRequest {
	s.RequestParamsShrink = &v
	return s
}

type SetMessageExtensionByKeysResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码，成功时为0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息，成功时为空
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SetMessageExtensionByKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMessageExtensionByKeysResponseBody) GoString() string {
	return s.String()
}

func (s *SetMessageExtensionByKeysResponseBody) SetRequestId(v string) *SetMessageExtensionByKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMessageExtensionByKeysResponseBody) SetCode(v string) *SetMessageExtensionByKeysResponseBody {
	s.Code = &v
	return s
}

func (s *SetMessageExtensionByKeysResponseBody) SetMessage(v string) *SetMessageExtensionByKeysResponseBody {
	s.Message = &v
	return s
}

type SetMessageExtensionByKeysResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetMessageExtensionByKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetMessageExtensionByKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMessageExtensionByKeysResponse) GoString() string {
	return s.String()
}

func (s *SetMessageExtensionByKeysResponse) SetHeaders(v map[string]*string) *SetMessageExtensionByKeysResponse {
	s.Headers = v
	return s
}

func (s *SetMessageExtensionByKeysResponse) SetBody(v *SetMessageExtensionByKeysResponseBody) *SetMessageExtensionByKeysResponse {
	s.Body = v
	return s
}

type ResultImportMessageResultValue struct {
	// 0 成功
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 消息ID
	MsgId *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
}

func (s ResultImportMessageResultValue) String() string {
	return tea.Prettify(s)
}

func (s ResultImportMessageResultValue) GoString() string {
	return s.String()
}

func (s *ResultImportMessageResultValue) SetResult(v int64) *ResultImportMessageResultValue {
	s.Result = &v
	return s
}

func (s *ResultImportMessageResultValue) SetMsgId(v string) *ResultImportMessageResultValue {
	s.MsgId = &v
	return s
}

type RequestParamsOptionsSingleChatCreateRequestUserConversationValue struct {
	// 扩展信息
	UserExtensions map[string]*string `json:"UserExtensions,omitempty" xml:"UserExtensions,omitempty"`
}

func (s RequestParamsOptionsSingleChatCreateRequestUserConversationValue) String() string {
	return tea.Prettify(s)
}

func (s RequestParamsOptionsSingleChatCreateRequestUserConversationValue) GoString() string {
	return s.String()
}

func (s *RequestParamsOptionsSingleChatCreateRequestUserConversationValue) SetUserExtensions(v map[string]*string) *RequestParamsOptionsSingleChatCreateRequestUserConversationValue {
	s.UserExtensions = v
	return s
}

type ResultUserMuteSettingsValue struct {
	Mute       *bool  `json:"Mute,omitempty" xml:"Mute,omitempty"`
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s ResultUserMuteSettingsValue) String() string {
	return tea.Prettify(s)
}

func (s ResultUserMuteSettingsValue) GoString() string {
	return s.String()
}

func (s *ResultUserMuteSettingsValue) SetMute(v bool) *ResultUserMuteSettingsValue {
	s.Mute = &v
	return s
}

func (s *ResultUserMuteSettingsValue) SetExpireTime(v int64) *ResultUserMuteSettingsValue {
	s.ExpireTime = &v
	return s
}

type RequestParamsUserConversationsValue struct {
	// 是否置顶
	Top *bool `json:"Top,omitempty" xml:"Top,omitempty"`
	// 未读数
	RedPoint *int64 `json:"RedPoint,omitempty" xml:"RedPoint,omitempty"`
	// 是否免打扰
	Mute *bool `json:"Mute,omitempty" xml:"Mute,omitempty"`
	// 是否可见
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
	// 创建时间戳
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 修改时间戳
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 自定义信息
	UserExtensions map[string]*string `json:"UserExtensions,omitempty" xml:"UserExtensions,omitempty"`
}

func (s RequestParamsUserConversationsValue) String() string {
	return tea.Prettify(s)
}

func (s RequestParamsUserConversationsValue) GoString() string {
	return s.String()
}

func (s *RequestParamsUserConversationsValue) SetTop(v bool) *RequestParamsUserConversationsValue {
	s.Top = &v
	return s
}

func (s *RequestParamsUserConversationsValue) SetRedPoint(v int64) *RequestParamsUserConversationsValue {
	s.RedPoint = &v
	return s
}

func (s *RequestParamsUserConversationsValue) SetMute(v bool) *RequestParamsUserConversationsValue {
	s.Mute = &v
	return s
}

func (s *RequestParamsUserConversationsValue) SetVisible(v bool) *RequestParamsUserConversationsValue {
	s.Visible = &v
	return s
}

func (s *RequestParamsUserConversationsValue) SetCreateTime(v int64) *RequestParamsUserConversationsValue {
	s.CreateTime = &v
	return s
}

func (s *RequestParamsUserConversationsValue) SetModifyTime(v int64) *RequestParamsUserConversationsValue {
	s.ModifyTime = &v
	return s
}

func (s *RequestParamsUserConversationsValue) SetUserExtensions(v map[string]*string) *RequestParamsUserConversationsValue {
	s.UserExtensions = v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("live-interaction"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ListAppInfosWithOptions(tmpReq *ListAppInfosRequest, runtime *util.RuntimeOptions) (_result *ListAppInfosResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAppInfosShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAppInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAppInfos"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppInfos(request *ListAppInfosRequest) (_result *ListAppInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppInfosResponse{}
	_body, _err := client.ListAppInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSingleChatExtensionByKeysWithOptions(tmpReq *RemoveSingleChatExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *RemoveSingleChatExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveSingleChatExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveSingleChatExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveSingleChatExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSingleChatExtensionByKeys(request *RemoveSingleChatExtensionByKeysRequest) (_result *RemoveSingleChatExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSingleChatExtensionByKeysResponse{}
	_body, _err := client.RemoveSingleChatExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportMessageWithOptions(tmpReq *ImportMessageRequest, runtime *util.RuntimeOptions) (_result *ImportMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportMessage"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportMessage(request *ImportMessageRequest) (_result *ImportMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportMessageResponse{}
	_body, _err := client.ImportMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindInterconnectionUidWithOptions(tmpReq *UnbindInterconnectionUidRequest, runtime *util.RuntimeOptions) (_result *UnbindInterconnectionUidResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UnbindInterconnectionUidShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindInterconnectionUidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindInterconnectionUid"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindInterconnectionUid(request *UnbindInterconnectionUidRequest) (_result *UnbindInterconnectionUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindInterconnectionUidResponse{}
	_body, _err := client.UnbindInterconnectionUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SilenceAllGroupMembersWithOptions(tmpReq *SilenceAllGroupMembersRequest, runtime *util.RuntimeOptions) (_result *SilenceAllGroupMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SilenceAllGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SilenceAllGroupMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SilenceAllGroupMembers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SilenceAllGroupMembers(request *SilenceAllGroupMembersRequest) (_result *SilenceAllGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SilenceAllGroupMembersResponse{}
	_body, _err := client.SilenceAllGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRoomMessagesWithOptions(request *ListRoomMessagesRequest, runtime *util.RuntimeOptions) (_result *ListRoomMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRoomMessagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRoomMessages"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoomMessages(request *ListRoomMessagesRequest) (_result *ListRoomMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRoomMessagesResponse{}
	_body, _err := client.ListRoomMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGroupExtensionByKeysWithOptions(tmpReq *SetGroupExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *SetGroupExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetGroupExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetGroupExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGroupExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGroupExtensionByKeys(request *SetGroupExtensionByKeysRequest) (_result *SetGroupExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGroupExtensionByKeysResponse{}
	_body, _err := client.SetGroupExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveGroupMemberExtensionByKeysWithOptions(tmpReq *RemoveGroupMemberExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *RemoveGroupMemberExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveGroupMemberExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveGroupMemberExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveGroupMemberExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveGroupMemberExtensionByKeys(request *RemoveGroupMemberExtensionByKeysRequest) (_result *RemoveGroupMemberExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveGroupMemberExtensionByKeysResponse{}
	_body, _err := client.RemoveGroupMemberExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGroupSilenceBlacklistWithOptions(tmpReq *AddGroupSilenceBlacklistRequest, runtime *util.RuntimeOptions) (_result *AddGroupSilenceBlacklistResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddGroupSilenceBlacklistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddGroupSilenceBlacklistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddGroupSilenceBlacklist"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGroupSilenceBlacklist(request *AddGroupSilenceBlacklistRequest) (_result *AddGroupSilenceBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGroupSilenceBlacklistResponse{}
	_body, _err := client.AddGroupSilenceBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveGroupSilenceWhitelistWithOptions(tmpReq *RemoveGroupSilenceWhitelistRequest, runtime *util.RuntimeOptions) (_result *RemoveGroupSilenceWhitelistResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveGroupSilenceWhitelistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveGroupSilenceWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveGroupSilenceWhitelist"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveGroupSilenceWhitelist(request *RemoveGroupSilenceWhitelistRequest) (_result *RemoveGroupSilenceWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveGroupSilenceWhitelistResponse{}
	_body, _err := client.RemoveGroupSilenceWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDetailReportStatisticsWithOptions(tmpReq *ListDetailReportStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListDetailReportStatisticsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDetailReportStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDetailReportStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDetailReportStatistics"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDetailReportStatistics(request *ListDetailReportStatisticsRequest) (_result *ListDetailReportStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDetailReportStatisticsResponse{}
	_body, _err := client.ListDetailReportStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetUserConversationExtensionByKeysWithOptions(tmpReq *SetUserConversationExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *SetUserConversationExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetUserConversationExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetUserConversationExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetUserConversationExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetUserConversationExtensionByKeys(request *SetUserConversationExtensionByKeysRequest) (_result *SetUserConversationExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetUserConversationExtensionByKeysResponse{}
	_body, _err := client.SetUserConversationExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupByIdWithOptions(tmpReq *GetGroupByIdRequest, runtime *util.RuntimeOptions) (_result *GetGroupByIdResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetGroupByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetGroupByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetGroupById"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGroupById(request *GetGroupByIdRequest) (_result *GetGroupByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGroupByIdResponse{}
	_body, _err := client.GetGroupByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTenantStatusWithOptions(request *UpdateTenantStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateTenantStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTenantStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTenantStatus"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTenantStatus(request *UpdateTenantStatusRequest) (_result *UpdateTenantStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTenantStatusResponse{}
	_body, _err := client.UpdateTenantStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCommonConfigWithOptions(request *GetCommonConfigRequest, runtime *util.RuntimeOptions) (_result *GetCommonConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCommonConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCommonConfig"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCommonConfig(request *GetCommonConfigRequest) (_result *GetCommonConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCommonConfigResponse{}
	_body, _err := client.GetCommonConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(tmpReq *SendMessageRequest, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendMessage"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupMembersRoleWithOptions(tmpReq *UpdateGroupMembersRoleRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupMembersRoleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateGroupMembersRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGroupMembersRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGroupMembersRole"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupMembersRole(request *UpdateGroupMembersRoleRequest) (_result *UpdateGroupMembersRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupMembersRoleResponse{}
	_body, _err := client.UpdateGroupMembersRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelSilenceAllGroupMembersWithOptions(tmpReq *CancelSilenceAllGroupMembersRequest, runtime *util.RuntimeOptions) (_result *CancelSilenceAllGroupMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CancelSilenceAllGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelSilenceAllGroupMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelSilenceAllGroupMembers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelSilenceAllGroupMembers(request *CancelSilenceAllGroupMembersRequest) (_result *CancelSilenceAllGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelSilenceAllGroupMembersResponse{}
	_body, _err := client.CancelSilenceAllGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupIconWithOptions(tmpReq *UpdateGroupIconRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupIconResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateGroupIconShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGroupIconResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGroupIcon"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupIcon(request *UpdateGroupIconRequest) (_result *UpdateGroupIconResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupIconResponse{}
	_body, _err := client.UpdateGroupIconWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveGroupMembersWithOptions(tmpReq *RemoveGroupMembersRequest, runtime *util.RuntimeOptions) (_result *RemoveGroupMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveGroupMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveGroupMembers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveGroupMembers(request *RemoveGroupMembersRequest) (_result *RemoveGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveGroupMembersResponse{}
	_body, _err := client.RemoveGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupAllMembersWithOptions(tmpReq *ListGroupAllMembersRequest, runtime *util.RuntimeOptions) (_result *ListGroupAllMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListGroupAllMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListGroupAllMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListGroupAllMembers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupAllMembers(request *ListGroupAllMembersRequest) (_result *ListGroupAllMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupAllMembersResponse{}
	_body, _err := client.ListGroupAllMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserMuteSettingWithOptions(tmpReq *GetUserMuteSettingRequest, runtime *util.RuntimeOptions) (_result *GetUserMuteSettingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetUserMuteSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserMuteSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserMuteSetting"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserMuteSetting(request *GetUserMuteSettingRequest) (_result *GetUserMuteSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserMuteSettingResponse{}
	_body, _err := client.GetUserMuteSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoomStatisticsWithOptions(request *GetRoomStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetRoomStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRoomStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRoomStatistics"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoomStatistics(request *GetRoomStatisticsRequest) (_result *GetRoomStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRoomStatisticsResponse{}
	_body, _err := client.GetRoomStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGroupMembersWithOptions(tmpReq *AddGroupMembersRequest, runtime *util.RuntimeOptions) (_result *AddGroupMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddGroupMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddGroupMembers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGroupMembers(request *AddGroupMembersRequest) (_result *AddGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGroupMembersResponse{}
	_body, _err := client.AddGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupMemberByIdsWithOptions(tmpReq *GetGroupMemberByIdsRequest, runtime *util.RuntimeOptions) (_result *GetGroupMemberByIdsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetGroupMemberByIdsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetGroupMemberByIdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetGroupMemberByIds"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGroupMemberByIds(request *GetGroupMemberByIdsRequest) (_result *GetGroupMemberByIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGroupMemberByIdsResponse{}
	_body, _err := client.GetGroupMemberByIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCustomMessageWithOptions(request *SendCustomMessageRequest, runtime *util.RuntimeOptions) (_result *SendCustomMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCustomMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCustomMessage"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCustomMessage(request *SendCustomMessageRequest) (_result *SendCustomMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomMessageResponse{}
	_body, _err := client.SendCustomMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppNameWithOptions(tmpReq *UpdateAppNameRequest, runtime *util.RuntimeOptions) (_result *UpdateAppNameResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAppNameShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppName"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppName(request *UpdateAppNameRequest) (_result *UpdateAppNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppNameResponse{}
	_body, _err := client.UpdateAppNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIMConfigWithOptions(request *GetIMConfigRequest, runtime *util.RuntimeOptions) (_result *GetIMConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetIMConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIMConfig"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIMConfig(request *GetIMConfigRequest) (_result *GetIMConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIMConfigResponse{}
	_body, _err := client.GetIMConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetSingleChatExtensionByKeysWithOptions(tmpReq *SetSingleChatExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *SetSingleChatExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetSingleChatExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetSingleChatExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetSingleChatExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetSingleChatExtensionByKeys(request *SetSingleChatExtensionByKeysRequest) (_result *SetSingleChatExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSingleChatExtensionByKeysResponse{}
	_body, _err := client.SetSingleChatExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppStatusWithOptions(tmpReq *UpdateAppStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateAppStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAppStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppStatus"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppStatus(request *UpdateAppStatusRequest) (_result *UpdateAppStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppStatusResponse{}
	_body, _err := client.UpdateAppStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MuteUsersWithOptions(tmpReq *MuteUsersRequest, runtime *util.RuntimeOptions) (_result *MuteUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &MuteUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MuteUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MuteUsers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MuteUsers(request *MuteUsersRequest) (_result *MuteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MuteUsersResponse{}
	_body, _err := client.MuteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecallMessageWithOptions(tmpReq *RecallMessageRequest, runtime *util.RuntimeOptions) (_result *RecallMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RecallMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecallMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecallMessage"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecallMessage(request *RecallMessageRequest) (_result *RecallMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecallMessageResponse{}
	_body, _err := client.RecallMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGroupSilenceWhitelistWithOptions(tmpReq *AddGroupSilenceWhitelistRequest, runtime *util.RuntimeOptions) (_result *AddGroupSilenceWhitelistResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddGroupSilenceWhitelistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddGroupSilenceWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddGroupSilenceWhitelist"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGroupSilenceWhitelist(request *AddGroupSilenceWhitelistRequest) (_result *AddGroupSilenceWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGroupSilenceWhitelistResponse{}
	_body, _err := client.AddGroupSilenceWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGroupOwnerWithOptions(tmpReq *SetGroupOwnerRequest, runtime *util.RuntimeOptions) (_result *SetGroupOwnerResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetGroupOwnerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetGroupOwnerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGroupOwner"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGroupOwner(request *SetGroupOwnerRequest) (_result *SetGroupOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGroupOwnerResponse{}
	_body, _err := client.SetGroupOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRoomUsersWithOptions(request *ListRoomUsersRequest, runtime *util.RuntimeOptions) (_result *ListRoomUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRoomUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRoomUsers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoomUsers(request *ListRoomUsersRequest) (_result *ListRoomUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRoomUsersResponse{}
	_body, _err := client.ListRoomUsersWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DeleteApp"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RemoveGroupSilenceBlacklistWithOptions(tmpReq *RemoveGroupSilenceBlacklistRequest, runtime *util.RuntimeOptions) (_result *RemoveGroupSilenceBlacklistResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveGroupSilenceBlacklistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveGroupSilenceBlacklistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveGroupSilenceBlacklist"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveGroupSilenceBlacklist(request *RemoveGroupSilenceBlacklistRequest) (_result *RemoveGroupSilenceBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveGroupSilenceBlacklistResponse{}
	_body, _err := client.RemoveGroupSilenceBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveMessageExtensionByKeysWithOptions(tmpReq *RemoveMessageExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *RemoveMessageExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveMessageExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveMessageExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveMessageExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveMessageExtensionByKeys(request *RemoveMessageExtensionByKeysRequest) (_result *RemoveMessageExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveMessageExtensionByKeysResponse{}
	_body, _err := client.RemoveMessageExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaUploadUrlWithOptions(tmpReq *GetMediaUploadUrlRequest, runtime *util.RuntimeOptions) (_result *GetMediaUploadUrlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetMediaUploadUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaUploadUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaUploadUrl"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaUploadUrl(request *GetMediaUploadUrlRequest) (_result *GetMediaUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaUploadUrlResponse{}
	_body, _err := client.GetMediaUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindInterconnectionUidWithOptions(tmpReq *BindInterconnectionUidRequest, runtime *util.RuntimeOptions) (_result *BindInterconnectionUidResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BindInterconnectionUidShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindInterconnectionUidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindInterconnectionUid"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindInterconnectionUid(request *BindInterconnectionUidRequest) (_result *BindInterconnectionUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindInterconnectionUidResponse{}
	_body, _err := client.BindInterconnectionUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaUrlWithOptions(tmpReq *GetMediaUrlRequest, runtime *util.RuntimeOptions) (_result *GetMediaUrlResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetMediaUrlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaUrl"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaUrl(request *GetMediaUrlRequest) (_result *GetMediaUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaUrlResponse{}
	_body, _err := client.GetMediaUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportSingleConversationWithOptions(tmpReq *ImportSingleConversationRequest, runtime *util.RuntimeOptions) (_result *ImportSingleConversationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportSingleConversationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportSingleConversationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportSingleConversation"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportSingleConversation(request *ImportSingleConversationRequest) (_result *ImportSingleConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportSingleConversationResponse{}
	_body, _err := client.ImportSingleConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCallbackConfigWithOptions(tmpReq *UpdateCallbackConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateCallbackConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateCallbackConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCallbackConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCallbackConfig"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCallbackConfig(request *UpdateCallbackConfigRequest) (_result *UpdateCallbackConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCallbackConfigResponse{}
	_body, _err := client.UpdateCallbackConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindInterconnectionCidWithOptions(tmpReq *BindInterconnectionCidRequest, runtime *util.RuntimeOptions) (_result *BindInterconnectionCidResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BindInterconnectionCidShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindInterconnectionCidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindInterconnectionCid"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindInterconnectionCid(request *BindInterconnectionCidRequest) (_result *BindInterconnectionCidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindInterconnectionCidResponse{}
	_body, _err := client.BindInterconnectionCidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitTenantWithOptions(request *InitTenantRequest, runtime *util.RuntimeOptions) (_result *InitTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitTenantResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitTenant"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitTenant(request *InitTenantRequest) (_result *InitTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitTenantResponse{}
	_body, _err := client.InitTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportGroupChatMemberWithOptions(tmpReq *ImportGroupChatMemberRequest, runtime *util.RuntimeOptions) (_result *ImportGroupChatMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportGroupChatMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportGroupChatMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportGroupChatMember"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportGroupChatMember(request *ImportGroupChatMemberRequest) (_result *ImportGroupChatMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportGroupChatMemberResponse{}
	_body, _err := client.ImportGroupChatMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupSilenceMembersWithOptions(tmpReq *ListGroupSilenceMembersRequest, runtime *util.RuntimeOptions) (_result *ListGroupSilenceMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListGroupSilenceMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListGroupSilenceMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListGroupSilenceMembers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupSilenceMembers(request *ListGroupSilenceMembersRequest) (_result *ListGroupSilenceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupSilenceMembersResponse{}
	_body, _err := client.ListGroupSilenceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveGroupExtensionByKeysWithOptions(tmpReq *RemoveGroupExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *RemoveGroupExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveGroupExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveGroupExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveGroupExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveGroupExtensionByKeys(request *RemoveGroupExtensionByKeysRequest) (_result *RemoveGroupExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveGroupExtensionByKeysResponse{}
	_body, _err := client.RemoveGroupExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGroupMemberExtensionByKeysWithOptions(tmpReq *SetGroupMemberExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *SetGroupMemberExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetGroupMemberExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetGroupMemberExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGroupMemberExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGroupMemberExtensionByKeys(request *SetGroupMemberExtensionByKeysRequest) (_result *SetGroupMemberExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGroupMemberExtensionByKeysResponse{}
	_body, _err := client.SetGroupMemberExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupWithOptions(tmpReq *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGroup"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMessageByIdWithOptions(tmpReq *GetMessageByIdRequest, runtime *util.RuntimeOptions) (_result *GetMessageByIdResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetMessageByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMessageByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMessageById"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMessageById(request *GetMessageByIdRequest) (_result *GetMessageByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMessageByIdResponse{}
	_body, _err := client.GetMessageByIdWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DestroyRoom"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) KickOffWithOptions(tmpReq *KickOffRequest, runtime *util.RuntimeOptions) (_result *KickOffResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &KickOffShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &KickOffResponse{}
	_body, _err := client.DoRPCRequest(tea.String("KickOff"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KickOff(request *KickOffRequest) (_result *KickOffResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KickOffResponse{}
	_body, _err := client.KickOffWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCallbackApiIdsWithOptions(runtime *util.RuntimeOptions) (_result *ListCallbackApiIdsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListCallbackApiIdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCallbackApiIds"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCallbackApiIds() (_result *ListCallbackApiIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCallbackApiIdsResponse{}
	_body, _err := client.ListCallbackApiIdsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetMessageReadWithOptions(tmpReq *SetMessageReadRequest, runtime *util.RuntimeOptions) (_result *SetMessageReadResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetMessageReadShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetMessageReadResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetMessageRead"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetMessageRead(request *SetMessageReadRequest) (_result *SetMessageReadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetMessageReadResponse{}
	_body, _err := client.SetMessageReadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMsgRecallIntervalWithOptions(tmpReq *UpdateMsgRecallIntervalRequest, runtime *util.RuntimeOptions) (_result *UpdateMsgRecallIntervalResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateMsgRecallIntervalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMsgRecallIntervalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMsgRecallInterval"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMsgRecallInterval(request *UpdateMsgRecallIntervalRequest) (_result *UpdateMsgRecallIntervalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMsgRecallIntervalResponse{}
	_body, _err := client.UpdateMsgRecallIntervalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCustomMessageToRoomUsersWithOptions(request *SendCustomMessageToRoomUsersRequest, runtime *util.RuntimeOptions) (_result *SendCustomMessageToRoomUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendCustomMessageToRoomUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCustomMessageToRoomUsers"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCustomMessageToRoomUsers(request *SendCustomMessageToRoomUsersRequest) (_result *SendCustomMessageToRoomUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomMessageToRoomUsersResponse{}
	_body, _err := client.SendCustomMessageToRoomUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupTitleWithOptions(tmpReq *UpdateGroupTitleRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupTitleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateGroupTitleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGroupTitleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGroupTitle"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupTitle(request *UpdateGroupTitleRequest) (_result *UpdateGroupTitleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupTitleResponse{}
	_body, _err := client.UpdateGroupTitleWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("GetLoginToken"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) QueryInterconnectionCidMappingWithOptions(tmpReq *QueryInterconnectionCidMappingRequest, runtime *util.RuntimeOptions) (_result *QueryInterconnectionCidMappingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryInterconnectionCidMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryInterconnectionCidMappingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryInterconnectionCidMapping"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInterconnectionCidMapping(request *QueryInterconnectionCidMappingRequest) (_result *QueryInterconnectionCidMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInterconnectionCidMappingResponse{}
	_body, _err := client.QueryInterconnectionCidMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DismissGroupWithOptions(tmpReq *DismissGroupRequest, runtime *util.RuntimeOptions) (_result *DismissGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DismissGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DismissGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DismissGroup"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DismissGroup(request *DismissGroupRequest) (_result *DismissGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DismissGroupResponse{}
	_body, _err := client.DismissGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportGroupChatConversationWithOptions(tmpReq *ImportGroupChatConversationRequest, runtime *util.RuntimeOptions) (_result *ImportGroupChatConversationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportGroupChatConversationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImportGroupChatConversationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImportGroupChatConversation"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportGroupChatConversation(request *ImportGroupChatConversationRequest) (_result *ImportGroupChatConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportGroupChatConversationResponse{}
	_body, _err := client.ImportGroupChatConversationWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateRoom"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RemoveUserConversationExtensionByKeysWithOptions(tmpReq *RemoveUserConversationExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *RemoveUserConversationExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveUserConversationExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveUserConversationExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveUserConversationExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserConversationExtensionByKeys(request *RemoveUserConversationExtensionByKeysRequest) (_result *RemoveUserConversationExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserConversationExtensionByKeysResponse{}
	_body, _err := client.RemoveUserConversationExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetMessageExtensionByKeysWithOptions(tmpReq *SetMessageExtensionByKeysRequest, runtime *util.RuntimeOptions) (_result *SetMessageExtensionByKeysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetMessageExtensionByKeysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.RequestParams))) {
		request.RequestParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.RequestParams), tea.String("RequestParams"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetMessageExtensionByKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetMessageExtensionByKeys"), tea.String("2020-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetMessageExtensionByKeys(request *SetMessageExtensionByKeysRequest) (_result *SetMessageExtensionByKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetMessageExtensionByKeysResponse{}
	_body, _err := client.SetMessageExtensionByKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
