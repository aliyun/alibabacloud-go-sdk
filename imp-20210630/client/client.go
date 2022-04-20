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

type AddMemberRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 邀请者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	// 被邀请用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
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

func (s *AddMemberRequest) SetFromUserId(v string) *AddMemberRequest {
	s.FromUserId = &v
	return s
}

func (s *AddMemberRequest) SetToUserId(v string) *AddMemberRequest {
	s.ToUserId = &v
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

type AgreeLinkMicRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 同意者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	// 被同意用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
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

func (s *AgreeLinkMicRequest) SetFromUserId(v string) *AgreeLinkMicRequest {
	s.FromUserId = &v
	return s
}

func (s *AgreeLinkMicRequest) SetToUserId(v string) *AgreeLinkMicRequest {
	s.ToUserId = &v
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

type BanCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 禁言时长（秒）
	BanCommentTime *int64 `json:"BanCommentTime,omitempty" xml:"BanCommentTime,omitempty"`
	// 被禁言的用户在房间内的唯一标识
	BanCommentUser *string `json:"BanCommentUser,omitempty" xml:"BanCommentUser,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *BanCommentRequest) SetBanCommentTime(v int64) *BanCommentRequest {
	s.BanCommentTime = &v
	return s
}

func (s *BanCommentRequest) SetBanCommentUser(v string) *BanCommentRequest {
	s.BanCommentUser = &v
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

type CancelBanCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 取消禁言的用户唯一标识
	BanCommentUser *string `json:"BanCommentUser,omitempty" xml:"BanCommentUser,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *CancelBanCommentRequest) SetBanCommentUser(v string) *CancelBanCommentRequest {
	s.BanCommentUser = &v
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

type CancelUserAdminRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelUserAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelUserAdminRequest) GoString() string {
	return s.String()
}

func (s *CancelUserAdminRequest) SetAppId(v string) *CancelUserAdminRequest {
	s.AppId = &v
	return s
}

func (s *CancelUserAdminRequest) SetRoomId(v string) *CancelUserAdminRequest {
	s.RoomId = &v
	return s
}

func (s *CancelUserAdminRequest) SetUserId(v string) *CancelUserAdminRequest {
	s.UserId = &v
	return s
}

type CancelUserAdminResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelUserAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelUserAdminResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUserAdminResponseBody) SetRequestId(v string) *CancelUserAdminResponseBody {
	s.RequestId = &v
	return s
}

type CancelUserAdminResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelUserAdminResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelUserAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelUserAdminResponse) GoString() string {
	return s.String()
}

func (s *CancelUserAdminResponse) SetHeaders(v map[string]*string) *CancelUserAdminResponse {
	s.Headers = v
	return s
}

func (s *CancelUserAdminResponse) SetBody(v *CancelUserAdminResponseBody) *CancelUserAdminResponse {
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

type CreateAppTemplateRequest struct {
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
	// 集成方式（一体化SDK：paasSDK，样板间：standardRoom）
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 应用模板场景，电商business，课堂classroom
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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

func (s *CreateAppTemplateRequest) SetIntegrationMode(v string) *CreateAppTemplateRequest {
	s.IntegrationMode = &v
	return s
}

func (s *CreateAppTemplateRequest) SetScene(v string) *CreateAppTemplateRequest {
	s.Scene = &v
	return s
}

type CreateAppTemplateShrinkRequest struct {
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 组件列表
	ComponentListShrink *string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty"`
	// 集成方式（一体化SDK：paasSDK，样板间：standardRoom）
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 应用模板场景，电商business，课堂classroom
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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

func (s *CreateAppTemplateShrinkRequest) SetIntegrationMode(v string) *CreateAppTemplateShrinkRequest {
	s.IntegrationMode = &v
	return s
}

func (s *CreateAppTemplateShrinkRequest) SetScene(v string) *CreateAppTemplateShrinkRequest {
	s.Scene = &v
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

type CreateClassRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建人用户昵称。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人用户ID。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 课堂标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateClassRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClassRequest) GoString() string {
	return s.String()
}

func (s *CreateClassRequest) SetAppId(v string) *CreateClassRequest {
	s.AppId = &v
	return s
}

func (s *CreateClassRequest) SetCreateNickname(v string) *CreateClassRequest {
	s.CreateNickname = &v
	return s
}

func (s *CreateClassRequest) SetCreateUserId(v string) *CreateClassRequest {
	s.CreateUserId = &v
	return s
}

func (s *CreateClassRequest) SetTitle(v string) *CreateClassRequest {
	s.Title = &v
	return s
}

type CreateClassResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *CreateClassResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClassResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClassResponseBody) SetRequestId(v string) *CreateClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClassResponseBody) SetResult(v *CreateClassResponseBodyResult) *CreateClassResponseBody {
	s.Result = v
	return s
}

type CreateClassResponseBodyResult struct {
	// 课堂唯一标识。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 连麦会议唯一标识。
	ConfId *string `json:"ConfId,omitempty" xml:"ConfId,omitempty"`
	// 创建人昵称。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人ID。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 直播的唯一标识ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 课堂状态，0:未开始 1:上课中 2:已下课。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 课堂标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 白板ID
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
	// 白板录制ID
	WhiteboardRecordId *string `json:"WhiteboardRecordId,omitempty" xml:"WhiteboardRecordId,omitempty"`
}

func (s CreateClassResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateClassResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateClassResponseBodyResult) SetClassId(v string) *CreateClassResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetConfId(v string) *CreateClassResponseBodyResult {
	s.ConfId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetCreateNickname(v string) *CreateClassResponseBodyResult {
	s.CreateNickname = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetCreateUserId(v string) *CreateClassResponseBodyResult {
	s.CreateUserId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetLiveId(v string) *CreateClassResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetRoomId(v string) *CreateClassResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetStatus(v int32) *CreateClassResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetTitle(v string) *CreateClassResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetWhiteboardId(v string) *CreateClassResponseBodyResult {
	s.WhiteboardId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetWhiteboardRecordId(v string) *CreateClassResponseBodyResult {
	s.WhiteboardRecordId = &v
	return s
}

type CreateClassResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClassResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClassResponse) GoString() string {
	return s.String()
}

func (s *CreateClassResponse) SetHeaders(v map[string]*string) *CreateClassResponse {
	s.Headers = v
	return s
}

func (s *CreateClassResponse) SetBody(v *CreateClassResponseBody) *CreateClassResponse {
	s.Body = v
	return s
}

type CreateConferenceRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间ID，最大长度36个字符，传空值，则随机生成一个房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 会议标题，支持中英文，最大长度256位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建会议用户。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *CreateConferenceRequest) SetTitle(v string) *CreateConferenceRequest {
	s.Title = &v
	return s
}

func (s *CreateConferenceRequest) SetUserId(v string) *CreateConferenceRequest {
	s.UserId = &v
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

type CreateLiveRequest struct {
	// 主播ID，支持中英文，最大长度128位，缺省时主播为当前创建直播用户。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播推流码率，缺省时默认为3。取值：  -1：流畅lld。 1：标清lsd。 2：高清lhd。 3：超清lud。
	CodeLevel *int32 `json:"CodeLevel,omitempty" xml:"CodeLevel,omitempty"`
	// 直播简介，支持中英文，最大长度2048位。
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 直播资源的唯一标识ID，缺省时系统自动生成36位随机uuid字符串。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID，最大长度36个字符，传空值，则随机生成一个房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播标题，支持中英文，最大长度256位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建直播用户。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRequest) SetAnchorId(v string) *CreateLiveRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRequest) SetAppId(v string) *CreateLiveRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveRequest) SetCodeLevel(v int32) *CreateLiveRequest {
	s.CodeLevel = &v
	return s
}

func (s *CreateLiveRequest) SetIntroduction(v string) *CreateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *CreateLiveRequest) SetLiveId(v string) *CreateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *CreateLiveRequest) SetRoomId(v string) *CreateLiveRequest {
	s.RoomId = &v
	return s
}

func (s *CreateLiveRequest) SetTitle(v string) *CreateLiveRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRequest) SetUserId(v string) *CreateLiveRequest {
	s.UserId = &v
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

type CreateLiveRoomRequest struct {
	// 主播id，仅支持英文和数字，最大长度36位。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 封面，支持图片地址链接格式
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。最大支持4096个字节。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomRequest) SetAnchorId(v string) *CreateLiveRoomRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRoomRequest) SetAnchorNick(v string) *CreateLiveRoomRequest {
	s.AnchorNick = &v
	return s
}

func (s *CreateLiveRoomRequest) SetAppId(v string) *CreateLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveRoomRequest) SetCoverUrl(v string) *CreateLiveRoomRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveRoomRequest) SetExtension(v map[string]*string) *CreateLiveRoomRequest {
	s.Extension = v
	return s
}

func (s *CreateLiveRoomRequest) SetNotice(v string) *CreateLiveRoomRequest {
	s.Notice = &v
	return s
}

func (s *CreateLiveRoomRequest) SetTitle(v string) *CreateLiveRoomRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRoomRequest) SetUserId(v string) *CreateLiveRoomRequest {
	s.UserId = &v
	return s
}

type CreateLiveRoomShrinkRequest struct {
	// 主播id，仅支持英文和数字，最大长度36位。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 封面，支持图片地址链接格式
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。最大支持4096个字节。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateLiveRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomShrinkRequest) SetAnchorId(v string) *CreateLiveRoomShrinkRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetAnchorNick(v string) *CreateLiveRoomShrinkRequest {
	s.AnchorNick = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetAppId(v string) *CreateLiveRoomShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetCoverUrl(v string) *CreateLiveRoomShrinkRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetExtensionShrink(v string) *CreateLiveRoomShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetNotice(v string) *CreateLiveRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetTitle(v string) *CreateLiveRoomShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetUserId(v string) *CreateLiveRoomShrinkRequest {
	s.UserId = &v
	return s
}

type CreateLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *CreateLiveRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponseBody) SetRequestId(v string) *CreateLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveRoomResponseBody) SetResult(v *CreateLiveRoomResponseBodyResult) *CreateLiveRoomResponseBody {
	s.Result = v
	return s
}

type CreateLiveRoomResponseBodyResult struct {
	// 主播ID。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// RTS低延迟播流信息。
	ArtcInfo *CreateLiveRoomResponseBodyResultArtcInfo `json:"ArtcInfo,omitempty" xml:"ArtcInfo,omitempty" type:"Struct"`
	// 聊天ID。
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 封面。
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 原画HLS播放地址。
	HlsUrl *string `json:"HlsUrl,omitempty" xml:"HlsUrl,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址。
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 直播回放地址。
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*CreateLiveRoomResponseBodyResultPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 直播推流地址。
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateLiveRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponseBodyResult) SetAnchorId(v string) *CreateLiveRoomResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetAnchorNick(v string) *CreateLiveRoomResponseBodyResult {
	s.AnchorNick = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetAppId(v string) *CreateLiveRoomResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetArtcInfo(v *CreateLiveRoomResponseBodyResultArtcInfo) *CreateLiveRoomResponseBodyResult {
	s.ArtcInfo = v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetChatId(v string) *CreateLiveRoomResponseBodyResult {
	s.ChatId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetCoverUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetExtension(v map[string]*string) *CreateLiveRoomResponseBodyResult {
	s.Extension = v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetHlsUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.HlsUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetLiveId(v string) *CreateLiveRoomResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetLiveUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetNotice(v string) *CreateLiveRoomResponseBodyResult {
	s.Notice = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetPlaybackUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetPluginInstanceInfoList(v []*CreateLiveRoomResponseBodyResultPluginInstanceInfoList) *CreateLiveRoomResponseBodyResult {
	s.PluginInstanceInfoList = v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetPushUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetRoomId(v string) *CreateLiveRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetTitle(v string) *CreateLiveRoomResponseBodyResult {
	s.Title = &v
	return s
}

type CreateLiveRoomResponseBodyResultArtcInfo struct {
	// RTS转码流地址，推荐web端使用。
	ArtcH5Url *string `json:"ArtcH5Url,omitempty" xml:"ArtcH5Url,omitempty"`
	// RTS原码流地址，推荐移动端使用。
	ArtcUrl *string `json:"ArtcUrl,omitempty" xml:"ArtcUrl,omitempty"`
}

func (s CreateLiveRoomResponseBodyResultArtcInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponseBodyResultArtcInfo) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponseBodyResultArtcInfo) SetArtcH5Url(v string) *CreateLiveRoomResponseBodyResultArtcInfo {
	s.ArtcH5Url = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResultArtcInfo) SetArtcUrl(v string) *CreateLiveRoomResponseBodyResultArtcInfo {
	s.ArtcUrl = &v
	return s
}

type CreateLiveRoomResponseBodyResultPluginInstanceInfoList struct {
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s CreateLiveRoomResponseBodyResultPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponseBodyResultPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponseBodyResultPluginInstanceInfoList) SetCreateTime(v int64) *CreateLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResultPluginInstanceInfoList) SetExtension(v map[string]*string) *CreateLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.Extension = v
	return s
}

func (s *CreateLiveRoomResponseBodyResultPluginInstanceInfoList) SetPluginId(v string) *CreateLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResultPluginInstanceInfoList) SetPluginType(v string) *CreateLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.PluginType = &v
	return s
}

type CreateLiveRoomResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponse) SetHeaders(v map[string]*string) *CreateLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveRoomResponse) SetBody(v *CreateLiveRoomResponseBody) *CreateLiveRoomResponse {
	s.Body = v
	return s
}

type CreateRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位，传空则随机生成一个房间id。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间模板唯一标识，当前支持的取值：default，传空默认为default。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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

func (s *CreateRoomRequest) SetExtension(v map[string]*string) *CreateRoomRequest {
	s.Extension = v
	return s
}

func (s *CreateRoomRequest) SetNotice(v string) *CreateRoomRequest {
	s.Notice = &v
	return s
}

func (s *CreateRoomRequest) SetRoomId(v string) *CreateRoomRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomRequest) SetRoomOwnerId(v string) *CreateRoomRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *CreateRoomRequest) SetTemplateId(v string) *CreateRoomRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRoomRequest) SetTitle(v string) *CreateRoomRequest {
	s.Title = &v
	return s
}

type CreateRoomShrinkRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位，传空则随机生成一个房间id。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间模板唯一标识，当前支持的取值：default，传空默认为default。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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

func (s *CreateRoomShrinkRequest) SetExtensionShrink(v string) *CreateRoomShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetNotice(v string) *CreateRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetRoomId(v string) *CreateRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetRoomOwnerId(v string) *CreateRoomShrinkRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetTemplateId(v string) *CreateRoomShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetTitle(v string) *CreateRoomShrinkRequest {
	s.Title = &v
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

type CreateSensitiveWordRequest struct {
	// 用户的应用ID，在控制台创建应用时生成。包含小写字母、数字，长度为6个字符。
	AppId    *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WordList []*string `json:"WordList,omitempty" xml:"WordList,omitempty" type:"Repeated"`
}

func (s CreateSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordRequest) SetAppId(v string) *CreateSensitiveWordRequest {
	s.AppId = &v
	return s
}

func (s *CreateSensitiveWordRequest) SetWordList(v []*string) *CreateSensitiveWordRequest {
	s.WordList = v
	return s
}

type CreateSensitiveWordShrinkRequest struct {
	// 用户的应用ID，在控制台创建应用时生成。包含小写字母、数字，长度为6个字符。
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WordListShrink *string `json:"WordList,omitempty" xml:"WordList,omitempty"`
}

func (s CreateSensitiveWordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordShrinkRequest) SetAppId(v string) *CreateSensitiveWordShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateSensitiveWordShrinkRequest) SetWordListShrink(v string) *CreateSensitiveWordShrinkRequest {
	s.WordListShrink = &v
	return s
}

type CreateSensitiveWordResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用发送直播间弹幕的返回结果。
	Result *CreateSensitiveWordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordResponseBody) SetRequestId(v string) *CreateSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSensitiveWordResponseBody) SetResult(v *CreateSensitiveWordResponseBodyResult) *CreateSensitiveWordResponseBody {
	s.Result = v
	return s
}

type CreateSensitiveWordResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSensitiveWordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordResponseBodyResult) SetSuccess(v bool) *CreateSensitiveWordResponseBodyResult {
	s.Success = &v
	return s
}

type CreateSensitiveWordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordResponse) SetHeaders(v map[string]*string) *CreateSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *CreateSensitiveWordResponse) SetBody(v *CreateSensitiveWordResponseBody) *CreateSensitiveWordResponse {
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

type DeleteClassRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 操作人用户ID，仅支持中英文数字，下划线，中划线，1~36个字符。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteClassRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClassRequest) GoString() string {
	return s.String()
}

func (s *DeleteClassRequest) SetAppId(v string) *DeleteClassRequest {
	s.AppId = &v
	return s
}

func (s *DeleteClassRequest) SetClassId(v string) *DeleteClassRequest {
	s.ClassId = &v
	return s
}

func (s *DeleteClassRequest) SetUserId(v string) *DeleteClassRequest {
	s.UserId = &v
	return s
}

type DeleteClassResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClassResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClassResponseBody) SetRequestId(v string) *DeleteClassResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClassResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClassResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClassResponse) GoString() string {
	return s.String()
}

func (s *DeleteClassResponse) SetHeaders(v map[string]*string) *DeleteClassResponse {
	s.Headers = v
	return s
}

func (s *DeleteClassResponse) SetBody(v *DeleteClassResponseBody) *DeleteClassResponse {
	s.Body = v
	return s
}

type DeleteCommentRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 需要删除的弹幕id列表
	CommentIdList []*string `json:"CommentIdList,omitempty" xml:"CommentIdList,omitempty" type:"Repeated"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 删除的操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommentRequest) SetAppId(v string) *DeleteCommentRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCommentRequest) SetCommentIdList(v []*string) *DeleteCommentRequest {
	s.CommentIdList = v
	return s
}

func (s *DeleteCommentRequest) SetRoomId(v string) *DeleteCommentRequest {
	s.RoomId = &v
	return s
}

func (s *DeleteCommentRequest) SetUserId(v string) *DeleteCommentRequest {
	s.UserId = &v
	return s
}

type DeleteCommentResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用删除直播间弹幕的返回结果。
	Result *DeleteCommentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommentResponseBody) SetRequestId(v string) *DeleteCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCommentResponseBody) SetResult(v *DeleteCommentResponseBodyResult) *DeleteCommentResponseBody {
	s.Result = v
	return s
}

type DeleteCommentResponseBodyResult struct {
	// 删除的结果
	DeleteResult *bool `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
}

func (s DeleteCommentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteCommentResponseBodyResult) SetDeleteResult(v bool) *DeleteCommentResponseBodyResult {
	s.DeleteResult = &v
	return s
}

type DeleteCommentResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommentResponse) SetHeaders(v map[string]*string) *DeleteCommentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommentResponse) SetBody(v *DeleteCommentResponseBody) *DeleteCommentResponse {
	s.Body = v
	return s
}

type DeleteConferenceRequest struct {
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 会议资源的唯一标识ID
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 房间ID，最大长度36位
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建会议用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *DeleteConferenceRequest) SetConferenceId(v string) *DeleteConferenceRequest {
	s.ConferenceId = &v
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

type DeleteLiveRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRoomRequest) SetAppId(v string) *DeleteLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *DeleteLiveRoomRequest) SetLiveId(v string) *DeleteLiveRoomRequest {
	s.LiveId = &v
	return s
}

func (s *DeleteLiveRoomRequest) SetUserId(v string) *DeleteLiveRoomRequest {
	s.UserId = &v
	return s
}

type DeleteLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveRoomResponseBody) SetRequestId(v string) *DeleteLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLiveRoomResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRoomResponse) SetHeaders(v map[string]*string) *DeleteLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveRoomResponse) SetBody(v *DeleteLiveRoomResponseBody) *DeleteLiveRoomResponse {
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

type DeleteSensitiveWordRequest struct {
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	AppId    *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WordList []*string `json:"WordList,omitempty" xml:"WordList,omitempty" type:"Repeated"`
}

func (s DeleteSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordRequest) SetAppId(v string) *DeleteSensitiveWordRequest {
	s.AppId = &v
	return s
}

func (s *DeleteSensitiveWordRequest) SetWordList(v []*string) *DeleteSensitiveWordRequest {
	s.WordList = v
	return s
}

type DeleteSensitiveWordShrinkRequest struct {
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WordListShrink *string `json:"WordList,omitempty" xml:"WordList,omitempty"`
}

func (s DeleteSensitiveWordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordShrinkRequest) SetAppId(v string) *DeleteSensitiveWordShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteSensitiveWordShrinkRequest) SetWordListShrink(v string) *DeleteSensitiveWordShrinkRequest {
	s.WordListShrink = &v
	return s
}

type DeleteSensitiveWordResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用发送直播间弹幕的返回结果。
	Result *DeleteSensitiveWordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordResponseBody) SetRequestId(v string) *DeleteSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSensitiveWordResponseBody) SetResult(v *DeleteSensitiveWordResponseBodyResult) *DeleteSensitiveWordResponseBody {
	s.Result = v
	return s
}

type DeleteSensitiveWordResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSensitiveWordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordResponseBodyResult) SetSuccess(v bool) *DeleteSensitiveWordResponseBodyResult {
	s.Success = &v
	return s
}

type DeleteSensitiveWordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordResponse) SetHeaders(v map[string]*string) *DeleteSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *DeleteSensitiveWordResponse) SetBody(v *DeleteSensitiveWordResponseBody) *DeleteSensitiveWordResponse {
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
	// 应用配置状态
	AppConfigStatus *string `json:"AppConfigStatus,omitempty" xml:"AppConfigStatus,omitempty"`
	// 应用Key
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用状态
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// 模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 组件列表。
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 样板间信息
	StandardRoomInfo *string `json:"StandardRoomInfo,omitempty" xml:"StandardRoomInfo,omitempty"`
}

func (s GetAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResult) SetAppConfigStatus(v string) *GetAppResponseBodyResult {
	s.AppConfigStatus = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppKey(v string) *GetAppResponseBodyResult {
	s.AppKey = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppName(v string) *GetAppResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppStatus(v string) *GetAppResponseBodyResult {
	s.AppStatus = &v
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

func (s *GetAppResponseBodyResult) SetComponentList(v []*string) *GetAppResponseBodyResult {
	s.ComponentList = v
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
	// 应用模板创建者
	AppTemplateCreator *string `json:"AppTemplateCreator,omitempty" xml:"AppTemplateCreator,omitempty"`
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 应用列表信息
	Apps []*GetAppTemplateResponseBodyResultApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// 组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
	// 配置列表
	ConfigList []*GetAppTemplateResponseBodyResultConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 应用模板场景，电商business，课堂classroom
	Scene   *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SdkInfo *string `json:"SdkInfo,omitempty" xml:"SdkInfo,omitempty"`
	// 样板间信息
	StandardRoomInfo *string `json:"StandardRoomInfo,omitempty" xml:"StandardRoomInfo,omitempty"`
	// 应用模板使用状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppTemplateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAppTemplateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponseBodyResult) SetAppTemplateCreator(v string) *GetAppTemplateResponseBodyResult {
	s.AppTemplateCreator = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetAppTemplateName(v string) *GetAppTemplateResponseBodyResult {
	s.AppTemplateName = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetApps(v []*GetAppTemplateResponseBodyResultApps) *GetAppTemplateResponseBodyResult {
	s.Apps = v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetComponentList(v []*string) *GetAppTemplateResponseBodyResult {
	s.ComponentList = v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetConfigList(v []*GetAppTemplateResponseBodyResultConfigList) *GetAppTemplateResponseBodyResult {
	s.ConfigList = v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetCreateTime(v string) *GetAppTemplateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetIntegrationMode(v string) *GetAppTemplateResponseBodyResult {
	s.IntegrationMode = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetScene(v string) *GetAppTemplateResponseBodyResult {
	s.Scene = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetSdkInfo(v string) *GetAppTemplateResponseBodyResult {
	s.SdkInfo = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetStandardRoomInfo(v string) *GetAppTemplateResponseBodyResult {
	s.StandardRoomInfo = &v
	return s
}

func (s *GetAppTemplateResponseBodyResult) SetStatus(v string) *GetAppTemplateResponseBodyResult {
	s.Status = &v
	return s
}

type GetAppTemplateResponseBodyResultApps struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用的Key
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用状态
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
}

func (s GetAppTemplateResponseBodyResultApps) String() string {
	return tea.Prettify(s)
}

func (s GetAppTemplateResponseBodyResultApps) GoString() string {
	return s.String()
}

func (s *GetAppTemplateResponseBodyResultApps) SetAppId(v string) *GetAppTemplateResponseBodyResultApps {
	s.AppId = &v
	return s
}

func (s *GetAppTemplateResponseBodyResultApps) SetAppKey(v string) *GetAppTemplateResponseBodyResultApps {
	s.AppKey = &v
	return s
}

func (s *GetAppTemplateResponseBodyResultApps) SetAppName(v string) *GetAppTemplateResponseBodyResultApps {
	s.AppName = &v
	return s
}

func (s *GetAppTemplateResponseBodyResultApps) SetAppStatus(v string) *GetAppTemplateResponseBodyResultApps {
	s.AppStatus = &v
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

type GetAuthTokenRequest struct {
	// 用户的应用ID，在控制台创建应用时生成
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 终端设备类型,通过控制台创建和查询
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 终端设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 用户UserId,在AppId下单独唯一
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *GetAuthTokenRequest) SetAppKey(v string) *GetAuthTokenRequest {
	s.AppKey = &v
	return s
}

func (s *GetAuthTokenRequest) SetDeviceId(v string) *GetAuthTokenRequest {
	s.DeviceId = &v
	return s
}

func (s *GetAuthTokenRequest) SetUserId(v string) *GetAuthTokenRequest {
	s.UserId = &v
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
	// 用于长连接建连的token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 登录token过期时间(毫秒)
	AccessTokenExpiredTime *int64 `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	// 更新Token，若AccessToken过期，则可以使用RefreshToken再次获取新Token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
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

func (s *GetAuthTokenResponseBodyResult) SetAccessTokenExpiredTime(v int64) *GetAuthTokenResponseBodyResult {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetRefreshToken(v string) *GetAuthTokenResponseBodyResult {
	s.RefreshToken = &v
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

type GetClassDetailRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 操作人用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetClassDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClassDetailRequest) GoString() string {
	return s.String()
}

func (s *GetClassDetailRequest) SetAppId(v string) *GetClassDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetClassDetailRequest) SetClassId(v string) *GetClassDetailRequest {
	s.ClassId = &v
	return s
}

func (s *GetClassDetailRequest) SetUserId(v string) *GetClassDetailRequest {
	s.UserId = &v
	return s
}

type GetClassDetailResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *GetClassDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetClassDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClassDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetClassDetailResponseBody) SetRequestId(v string) *GetClassDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClassDetailResponseBody) SetResult(v *GetClassDetailResponseBodyResult) *GetClassDetailResponseBody {
	s.Result = v
	return s
}

type GetClassDetailResponseBodyResult struct {
	// 课堂唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 连麦会议唯一标识。
	ConfId *string `json:"ConfId,omitempty" xml:"ConfId,omitempty"`
	// 创建人昵称。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人ID。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 下课时间戳，毫秒。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 直播的唯一标识ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 开始上课时间戳，毫秒。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 课堂状态，0:未开始 1:上课中 2:已下课。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 课堂标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 白板ID
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
	// 白板录制ID
	WhiteboardRecordId *string `json:"WhiteboardRecordId,omitempty" xml:"WhiteboardRecordId,omitempty"`
}

func (s GetClassDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClassDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClassDetailResponseBodyResult) SetClassId(v string) *GetClassDetailResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetConfId(v string) *GetClassDetailResponseBodyResult {
	s.ConfId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetCreateNickname(v string) *GetClassDetailResponseBodyResult {
	s.CreateNickname = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetCreateUserId(v string) *GetClassDetailResponseBodyResult {
	s.CreateUserId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetEndTime(v int64) *GetClassDetailResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetLiveId(v string) *GetClassDetailResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetRoomId(v string) *GetClassDetailResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetStartTime(v int64) *GetClassDetailResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetStatus(v int32) *GetClassDetailResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetTitle(v string) *GetClassDetailResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetWhiteboardId(v string) *GetClassDetailResponseBodyResult {
	s.WhiteboardId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetWhiteboardRecordId(v string) *GetClassDetailResponseBodyResult {
	s.WhiteboardRecordId = &v
	return s
}

type GetClassDetailResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClassDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClassDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClassDetailResponse) GoString() string {
	return s.String()
}

func (s *GetClassDetailResponse) SetHeaders(v map[string]*string) *GetClassDetailResponse {
	s.Headers = v
	return s
}

func (s *GetClassDetailResponse) SetBody(v *GetClassDetailResponseBody) *GetClassDetailResponse {
	s.Body = v
	return s
}

type GetClassRecordRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课程唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 操作人用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetClassRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClassRecordRequest) GoString() string {
	return s.String()
}

func (s *GetClassRecordRequest) SetAppId(v string) *GetClassRecordRequest {
	s.AppId = &v
	return s
}

func (s *GetClassRecordRequest) SetClassId(v string) *GetClassRecordRequest {
	s.ClassId = &v
	return s
}

func (s *GetClassRecordRequest) SetUserId(v string) *GetClassRecordRequest {
	s.UserId = &v
	return s
}

type GetClassRecordResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *GetClassRecordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetClassRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClassRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetClassRecordResponseBody) SetRequestId(v string) *GetClassRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClassRecordResponseBody) SetResult(v *GetClassRecordResponseBodyResult) *GetClassRecordResponseBody {
	s.Result = v
	return s
}

type GetClassRecordResponseBodyResult struct {
	PlaybackUrlMap map[string][]*string `json:"PlaybackUrlMap,omitempty" xml:"PlaybackUrlMap,omitempty"`
}

func (s GetClassRecordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClassRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClassRecordResponseBodyResult) SetPlaybackUrlMap(v map[string][]*string) *GetClassRecordResponseBodyResult {
	s.PlaybackUrlMap = v
	return s
}

type GetClassRecordResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClassRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClassRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClassRecordResponse) GoString() string {
	return s.String()
}

func (s *GetClassRecordResponse) SetHeaders(v map[string]*string) *GetClassRecordResponse {
	s.Headers = v
	return s
}

func (s *GetClassRecordResponse) SetBody(v *GetClassRecordResponseBody) *GetClassRecordResponse {
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
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 会议资源唯一标识。
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 会议创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 录制回放地址，m3u8格式，会议结束后10秒才会生成。
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 会议状态。
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 会议标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建会议用户。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetConferenceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetConferenceResponseBodyResult) SetAppId(v string) *GetConferenceResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetConferenceId(v string) *GetConferenceResponseBodyResult {
	s.ConferenceId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetCreateTime(v int64) *GetConferenceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetPlaybackUrl(v string) *GetConferenceResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetRoomId(v string) *GetConferenceResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetStatus(v string) *GetConferenceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetTitle(v string) *GetConferenceResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetUserId(v string) *GetConferenceResponseBodyResult {
	s.UserId = &v
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
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// rts播流信息
	ArtcInfo *GetLiveResponseBodyResultArtcInfo `json:"ArtcInfo,omitempty" xml:"ArtcInfo,omitempty" type:"Struct"`
	// 直播推送分辨率 -1:lld 1:lsd 2:lhd 3:lud
	CodeLevel *int32 `json:"CodeLevel,omitempty" xml:"CodeLevel,omitempty"`
	// 封面图片
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播创建时间（毫秒ms）
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 直播持续时间（毫秒ms）
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 直播结束时间（毫秒ms）
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// hls播放地址
	HlsUrl *string `json:"HlsUrl,omitempty" xml:"HlsUrl,omitempty"`
	// 直播简介
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 多分辨率多协议播放信息
	PlayUrlInfoList []*GetLiveResponseBodyResultPlayUrlInfoList `json:"PlayUrlInfoList,omitempty" xml:"PlayUrlInfoList,omitempty" type:"Repeated"`
	// 直播回放地址
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// 直播推流地址
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播状态：Created-已创建，未开播，Living-直播中，End-直播结束
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 直播标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 用户自定义数据存储
	UserDefineField *string `json:"UserDefineField,omitempty" xml:"UserDefineField,omitempty"`
	// 创建直播用户
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *GetLiveResponseBodyResult) SetAppId(v string) *GetLiveResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetArtcInfo(v *GetLiveResponseBodyResultArtcInfo) *GetLiveResponseBodyResult {
	s.ArtcInfo = v
	return s
}

func (s *GetLiveResponseBodyResult) SetCodeLevel(v int32) *GetLiveResponseBodyResult {
	s.CodeLevel = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetCoverUrl(v string) *GetLiveResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetCreateTime(v int64) *GetLiveResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetDuration(v int64) *GetLiveResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetEndTime(v int64) *GetLiveResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetHlsUrl(v string) *GetLiveResponseBodyResult {
	s.HlsUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetIntroduction(v string) *GetLiveResponseBodyResult {
	s.Introduction = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetLiveId(v string) *GetLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetLiveUrl(v string) *GetLiveResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetPlayUrlInfoList(v []*GetLiveResponseBodyResultPlayUrlInfoList) *GetLiveResponseBodyResult {
	s.PlayUrlInfoList = v
	return s
}

func (s *GetLiveResponseBodyResult) SetPlaybackUrl(v string) *GetLiveResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetPushUrl(v string) *GetLiveResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetRoomId(v string) *GetLiveResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetStatus(v string) *GetLiveResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetTitle(v string) *GetLiveResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetUserDefineField(v string) *GetLiveResponseBodyResult {
	s.UserDefineField = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetUserId(v string) *GetLiveResponseBodyResult {
	s.UserId = &v
	return s
}

type GetLiveResponseBodyResultArtcInfo struct {
	// 原画转码地址
	ArtcH5Url *string `json:"ArtcH5Url,omitempty" xml:"ArtcH5Url,omitempty"`
	// 源码地址
	ArtcUrl *string `json:"ArtcUrl,omitempty" xml:"ArtcUrl,omitempty"`
}

func (s GetLiveResponseBodyResultArtcInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponseBodyResultArtcInfo) GoString() string {
	return s.String()
}

func (s *GetLiveResponseBodyResultArtcInfo) SetArtcH5Url(v string) *GetLiveResponseBodyResultArtcInfo {
	s.ArtcH5Url = &v
	return s
}

func (s *GetLiveResponseBodyResultArtcInfo) SetArtcUrl(v string) *GetLiveResponseBodyResultArtcInfo {
	s.ArtcUrl = &v
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
	// 直播域名CNAME
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// 直播域名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 直播域名状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) GoString() string {
	return s.String()
}

func (s *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) SetCname(v string) *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList {
	s.Cname = &v
	return s
}

func (s *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList) SetDomain(v string) *GetLiveDomainStatusResponseBodyResultLiveDomainInfoList {
	s.Domain = &v
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

type GetLiveRecordRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播唯一标识，由调用CreateLiveRoom返回。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 操作人用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetLiveRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRecordRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRecordRequest) SetAppId(v string) *GetLiveRecordRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveRecordRequest) SetLiveId(v string) *GetLiveRecordRequest {
	s.LiveId = &v
	return s
}

func (s *GetLiveRecordRequest) SetUserId(v string) *GetLiveRecordRequest {
	s.UserId = &v
	return s
}

type GetLiveRecordResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *GetLiveRecordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRecordResponseBody) SetRequestId(v string) *GetLiveRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRecordResponseBody) SetResult(v *GetLiveRecordResponseBodyResult) *GetLiveRecordResponseBody {
	s.Result = v
	return s
}

type GetLiveRecordResponseBodyResult struct {
	PlaybackUrlMap map[string][]*string `json:"PlaybackUrlMap,omitempty" xml:"PlaybackUrlMap,omitempty"`
}

func (s GetLiveRecordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveRecordResponseBodyResult) SetPlaybackUrlMap(v map[string][]*string) *GetLiveRecordResponseBodyResult {
	s.PlaybackUrlMap = v
	return s
}

type GetLiveRecordResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLiveRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRecordResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRecordResponse) SetHeaders(v map[string]*string) *GetLiveRecordResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRecordResponse) SetBody(v *GetLiveRecordResponseBody) *GetLiveRecordResponse {
	s.Body = v
	return s
}

type GetLiveRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s GetLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRoomRequest) SetAppId(v string) *GetLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveRoomRequest) SetLiveId(v string) *GetLiveRoomRequest {
	s.LiveId = &v
	return s
}

type GetLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *GetLiveRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponseBody) SetRequestId(v string) *GetLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRoomResponseBody) SetResult(v *GetLiveRoomResponseBodyResult) *GetLiveRoomResponseBody {
	s.Result = v
	return s
}

type GetLiveRoomResponseBodyResult struct {
	// 主播ID。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// RTS低延迟播流信息。
	ArtcInfo *GetLiveRoomResponseBodyResultArtcInfo `json:"ArtcInfo,omitempty" xml:"ArtcInfo,omitempty" type:"Struct"`
	// 聊天ID。
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 封面。
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播创建时间，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 直播结束时间，单位：毫秒。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 直播拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 原画HLS播放地址。
	HlsUrl *string `json:"HlsUrl,omitempty" xml:"HlsUrl,omitempty"`
	// https协议的原画HLS播放地址。
	HlsUrlHttps *string `json:"HlsUrlHttps,omitempty" xml:"HlsUrlHttps,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址。
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// https协议的直播拉流地址。
	LiveUrlHttps *string `json:"LiveUrlHttps,omitempty" xml:"LiveUrlHttps,omitempty"`
	// 公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 直播回放地址。
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// https协议的直播回放地址
	PlaybackUrlHttps *string `json:"PlaybackUrlHttps,omitempty" xml:"PlaybackUrlHttps,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*GetLiveRoomResponseBodyResultPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 直播推流地址。
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// rtmp协议的播放地址
	RtmpUrl *string `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty"`
	// 直播开始时间，单位：毫秒。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 直播状态，0-在播 1-下播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s GetLiveRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponseBodyResult) SetAnchorId(v string) *GetLiveRoomResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetAnchorNick(v string) *GetLiveRoomResponseBodyResult {
	s.AnchorNick = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetAppId(v string) *GetLiveRoomResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetArtcInfo(v *GetLiveRoomResponseBodyResultArtcInfo) *GetLiveRoomResponseBodyResult {
	s.ArtcInfo = v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetChatId(v string) *GetLiveRoomResponseBodyResult {
	s.ChatId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetCoverUrl(v string) *GetLiveRoomResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetCreateTime(v int64) *GetLiveRoomResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetEndTime(v int64) *GetLiveRoomResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetExtension(v map[string]*string) *GetLiveRoomResponseBodyResult {
	s.Extension = v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetHlsUrl(v string) *GetLiveRoomResponseBodyResult {
	s.HlsUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetHlsUrlHttps(v string) *GetLiveRoomResponseBodyResult {
	s.HlsUrlHttps = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetLiveId(v string) *GetLiveRoomResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetLiveUrl(v string) *GetLiveRoomResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetLiveUrlHttps(v string) *GetLiveRoomResponseBodyResult {
	s.LiveUrlHttps = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetNotice(v string) *GetLiveRoomResponseBodyResult {
	s.Notice = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetOnlineCount(v int64) *GetLiveRoomResponseBodyResult {
	s.OnlineCount = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPlaybackUrl(v string) *GetLiveRoomResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPlaybackUrlHttps(v string) *GetLiveRoomResponseBodyResult {
	s.PlaybackUrlHttps = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPluginInstanceInfoList(v []*GetLiveRoomResponseBodyResultPluginInstanceInfoList) *GetLiveRoomResponseBodyResult {
	s.PluginInstanceInfoList = v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPushUrl(v string) *GetLiveRoomResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPv(v int64) *GetLiveRoomResponseBodyResult {
	s.Pv = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetRoomId(v string) *GetLiveRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetRtmpUrl(v string) *GetLiveRoomResponseBodyResult {
	s.RtmpUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetStartTime(v int64) *GetLiveRoomResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetStatus(v int32) *GetLiveRoomResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetTitle(v string) *GetLiveRoomResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetUv(v int64) *GetLiveRoomResponseBodyResult {
	s.Uv = &v
	return s
}

type GetLiveRoomResponseBodyResultArtcInfo struct {
	// RTS转码流地址，推荐web端使用。
	ArtcH5Url *string `json:"ArtcH5Url,omitempty" xml:"ArtcH5Url,omitempty"`
	// RTS原码流地址，推荐移动端使用。
	ArtcUrl *string `json:"ArtcUrl,omitempty" xml:"ArtcUrl,omitempty"`
}

func (s GetLiveRoomResponseBodyResultArtcInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponseBodyResultArtcInfo) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponseBodyResultArtcInfo) SetArtcH5Url(v string) *GetLiveRoomResponseBodyResultArtcInfo {
	s.ArtcH5Url = &v
	return s
}

func (s *GetLiveRoomResponseBodyResultArtcInfo) SetArtcUrl(v string) *GetLiveRoomResponseBodyResultArtcInfo {
	s.ArtcUrl = &v
	return s
}

type GetLiveRoomResponseBodyResultPluginInstanceInfoList struct {
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s GetLiveRoomResponseBodyResultPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponseBodyResultPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponseBodyResultPluginInstanceInfoList) SetCreateTime(v int64) *GetLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetLiveRoomResponseBodyResultPluginInstanceInfoList) SetExtension(v map[string]*string) *GetLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.Extension = v
	return s
}

func (s *GetLiveRoomResponseBodyResultPluginInstanceInfoList) SetPluginId(v string) *GetLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResultPluginInstanceInfoList) SetPluginType(v string) *GetLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.PluginType = &v
	return s
}

type GetLiveRoomResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponse) SetHeaders(v map[string]*string) *GetLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRoomResponse) SetBody(v *GetLiveRoomResponseBody) *GetLiveRoomResponse {
	s.Body = v
	return s
}

type GetLiveRoomStatisticsRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s GetLiveRoomStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRoomStatisticsRequest) SetAppId(v string) *GetLiveRoomStatisticsRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveRoomStatisticsRequest) SetLiveId(v string) *GetLiveRoomStatisticsRequest {
	s.LiveId = &v
	return s
}

type GetLiveRoomStatisticsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *GetLiveRoomStatisticsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveRoomStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRoomStatisticsResponseBody) SetRequestId(v string) *GetLiveRoomStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBody) SetResult(v *GetLiveRoomStatisticsResponseBodyResult) *GetLiveRoomStatisticsResponseBody {
	s.Result = v
	return s
}

type GetLiveRoomStatisticsResponseBodyResult struct {
	// 直播结束时间，单位：毫秒。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 点赞数。
	LikeCount *int64 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 互动消息数。
	MessageCount *int64 `json:"MessageCount,omitempty" xml:"MessageCount,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 直播开始时间，单位：毫秒。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 直播状态，0-已创建 1-直播中 2-已关闭。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
	// 总观看时长，单位：毫秒。
	WatchLiveTime *int64 `json:"WatchLiveTime,omitempty" xml:"WatchLiveTime,omitempty"`
}

func (s GetLiveRoomStatisticsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomStatisticsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetEndTime(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetLikeCount(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.LikeCount = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetLiveId(v string) *GetLiveRoomStatisticsResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetMessageCount(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.MessageCount = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetOnlineCount(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.OnlineCount = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetPv(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.Pv = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetStartTime(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetStatus(v int32) *GetLiveRoomStatisticsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetUv(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.Uv = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetWatchLiveTime(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.WatchLiveTime = &v
	return s
}

type GetLiveRoomStatisticsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLiveRoomStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveRoomStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRoomStatisticsResponse) SetHeaders(v map[string]*string) *GetLiveRoomStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRoomStatisticsResponse) SetBody(v *GetLiveRoomStatisticsResponseBody) *GetLiveRoomStatisticsResponse {
	s.Body = v
	return s
}

type GetLiveRoomUserStatisticsRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetLiveRoomUserStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsRequest) SetAppId(v string) *GetLiveRoomUserStatisticsRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsRequest) SetLiveId(v string) *GetLiveRoomUserStatisticsRequest {
	s.LiveId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsRequest) SetPageNumber(v string) *GetLiveRoomUserStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLiveRoomUserStatisticsRequest) SetPageSize(v string) *GetLiveRoomUserStatisticsRequest {
	s.PageSize = &v
	return s
}

type GetLiveRoomUserStatisticsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *GetLiveRoomUserStatisticsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveRoomUserStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsResponseBody) SetRequestId(v string) *GetLiveRoomUserStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBody) SetResult(v *GetLiveRoomUserStatisticsResponseBodyResult) *GetLiveRoomUserStatisticsResponseBody {
	s.Result = v
	return s
}

type GetLiveRoomUserStatisticsResponseBodyResult struct {
	// 是否还有下一页。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 用户总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 用户总数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 用户观看数据列表。
	UserStatisticsList []*GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList `json:"UserStatisticsList,omitempty" xml:"UserStatisticsList,omitempty" type:"Repeated"`
}

func (s GetLiveRoomUserStatisticsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetHasMore(v bool) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetLiveId(v string) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetPageTotal(v int32) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetTotalCount(v int32) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetUserStatisticsList(v []*GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.UserStatisticsList = v
	return s
}

type GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList struct {
	// 用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 观看时长，单位：毫秒。
	WatchLiveTime *int64 `json:"WatchLiveTime,omitempty" xml:"WatchLiveTime,omitempty"`
}

func (s GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) SetUserId(v string) *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList {
	s.UserId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) SetWatchLiveTime(v int64) *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList {
	s.WatchLiveTime = &v
	return s
}

type GetLiveRoomUserStatisticsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLiveRoomUserStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveRoomUserStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsResponse) SetHeaders(v map[string]*string) *GetLiveRoomUserStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRoomUserStatisticsResponse) SetBody(v *GetLiveRoomUserStatisticsResponseBody) *GetLiveRoomUserStatisticsResponse {
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
	// 管理员ID列表。
	AdminIdList []*string `json:"AdminIdList,omitempty" xml:"AdminIdList,omitempty" type:"Repeated"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 房间拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 创建房间使用的模板id。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s GetRoomResponseBodyResultRoomInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBodyResultRoomInfo) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBodyResultRoomInfo) SetAdminIdList(v []*string) *GetRoomResponseBodyResultRoomInfo {
	s.AdminIdList = v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetAppId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.AppId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetCreateTime(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.CreateTime = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetExtension(v map[string]*string) *GetRoomResponseBodyResultRoomInfo {
	s.Extension = v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetNotice(v string) *GetRoomResponseBodyResultRoomInfo {
	s.Notice = &v
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

func (s *GetRoomResponseBodyResultRoomInfo) SetPv(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.Pv = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetRoomId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.RoomId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetRoomOwnerId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.RoomOwnerId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetTemplateId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.TemplateId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetTitle(v string) *GetRoomResponseBodyResultRoomInfo {
	s.Title = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetUv(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.Uv = &v
	return s
}

type GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList struct {
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetCreateTime(v int64) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetExtension(v map[string]*string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.Extension = v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetPluginId(v string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetPluginType(v string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.PluginType = &v
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

type GetStandardRoomHttpsCertificateRequest struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s GetStandardRoomHttpsCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomHttpsCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetStandardRoomHttpsCertificateRequest) SetCertificateId(v string) *GetStandardRoomHttpsCertificateRequest {
	s.CertificateId = &v
	return s
}

type GetStandardRoomHttpsCertificateResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *GetStandardRoomHttpsCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetStandardRoomHttpsCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomHttpsCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardRoomHttpsCertificateResponseBody) SetRequestId(v string) *GetStandardRoomHttpsCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardRoomHttpsCertificateResponseBody) SetResult(v *GetStandardRoomHttpsCertificateResponseBodyResult) *GetStandardRoomHttpsCertificateResponseBody {
	s.Result = v
	return s
}

type GetStandardRoomHttpsCertificateResponseBodyResult struct {
	// 证书名称
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// 证书创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 使用证书的确切域名
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// 证书过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s GetStandardRoomHttpsCertificateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomHttpsCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetStandardRoomHttpsCertificateResponseBodyResult) SetCertificateName(v string) *GetStandardRoomHttpsCertificateResponseBodyResult {
	s.CertificateName = &v
	return s
}

func (s *GetStandardRoomHttpsCertificateResponseBodyResult) SetCreateTime(v string) *GetStandardRoomHttpsCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetStandardRoomHttpsCertificateResponseBodyResult) SetDomainName(v string) *GetStandardRoomHttpsCertificateResponseBodyResult {
	s.DomainName = &v
	return s
}

func (s *GetStandardRoomHttpsCertificateResponseBodyResult) SetExpireTime(v string) *GetStandardRoomHttpsCertificateResponseBodyResult {
	s.ExpireTime = &v
	return s
}

type GetStandardRoomHttpsCertificateResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStandardRoomHttpsCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStandardRoomHttpsCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomHttpsCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetStandardRoomHttpsCertificateResponse) SetHeaders(v map[string]*string) *GetStandardRoomHttpsCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetStandardRoomHttpsCertificateResponse) SetBody(v *GetStandardRoomHttpsCertificateResponseBody) *GetStandardRoomHttpsCertificateResponse {
	s.Body = v
	return s
}

type GetStandardRoomJumpUrlRequest struct {
	// 用户的应用ID，在控制台创建应用时生成
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 终端设备类型,通过控制台创建和查询
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 资源ID：根据业务类型来定，比如直播ID，课堂ID等
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// 业务类型：互动直播live，互动课堂class
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// 平台：win, mac, android, ios, web
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// 用户UserId,在AppId下单独唯一
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *GetStandardRoomJumpUrlRequest) SetAppKey(v string) *GetStandardRoomJumpUrlRequest {
	s.AppKey = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetBizId(v string) *GetStandardRoomJumpUrlRequest {
	s.BizId = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetBizType(v string) *GetStandardRoomJumpUrlRequest {
	s.BizType = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetPlatform(v string) *GetStandardRoomJumpUrlRequest {
	s.Platform = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetUserId(v string) *GetStandardRoomJumpUrlRequest {
	s.UserId = &v
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

type KickRoomUserRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BlockTime *int64  `json:"BlockTime,omitempty" xml:"BlockTime,omitempty"`
	// 被踢出房间的用户ID。
	KickUser *string `json:"KickUser,omitempty" xml:"KickUser,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位，传空则随机生成一个房间id。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 操作人的用户ID，用于表示谁执行了踢人操作。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s KickRoomUserRequest) String() string {
	return tea.Prettify(s)
}

func (s KickRoomUserRequest) GoString() string {
	return s.String()
}

func (s *KickRoomUserRequest) SetAppId(v string) *KickRoomUserRequest {
	s.AppId = &v
	return s
}

func (s *KickRoomUserRequest) SetBlockTime(v int64) *KickRoomUserRequest {
	s.BlockTime = &v
	return s
}

func (s *KickRoomUserRequest) SetKickUser(v string) *KickRoomUserRequest {
	s.KickUser = &v
	return s
}

func (s *KickRoomUserRequest) SetRoomId(v string) *KickRoomUserRequest {
	s.RoomId = &v
	return s
}

func (s *KickRoomUserRequest) SetUserId(v string) *KickRoomUserRequest {
	s.UserId = &v
	return s
}

type KickRoomUserResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KickRoomUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KickRoomUserResponseBody) GoString() string {
	return s.String()
}

func (s *KickRoomUserResponseBody) SetRequestId(v string) *KickRoomUserResponseBody {
	s.RequestId = &v
	return s
}

type KickRoomUserResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KickRoomUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KickRoomUserResponse) String() string {
	return tea.Prettify(s)
}

func (s KickRoomUserResponse) GoString() string {
	return s.String()
}

func (s *KickRoomUserResponse) SetHeaders(v map[string]*string) *KickRoomUserResponse {
	s.Headers = v
	return s
}

func (s *KickRoomUserResponse) SetBody(v *KickRoomUserResponseBody) *KickRoomUserResponse {
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
	// 应用模板信息列表
	AppTemplateInfoList []*ListAppTemplatesResponseBodyResultAppTemplateInfoList `json:"AppTemplateInfoList,omitempty" xml:"AppTemplateInfoList,omitempty" type:"Repeated"`
	// 总页数
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppTemplatesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyResult) SetAppTemplateInfoList(v []*ListAppTemplatesResponseBodyResultAppTemplateInfoList) *ListAppTemplatesResponseBodyResult {
	s.AppTemplateInfoList = v
	return s
}

func (s *ListAppTemplatesResponseBodyResult) SetPageTotal(v int32) *ListAppTemplatesResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResult) SetTotalCount(v int32) *ListAppTemplatesResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListAppTemplatesResponseBodyResultAppTemplateInfoList struct {
	// 应用模板创建者
	AppTemplateCreator *string `json:"AppTemplateCreator,omitempty" xml:"AppTemplateCreator,omitempty"`
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
	// 配置列表
	ConfigList []*ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 应用模板场景，电商business，课堂classroom
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// SDK信息
	SdkInfo *string `json:"SdkInfo,omitempty" xml:"SdkInfo,omitempty"`
	// 样板间信息
	StandardRoomInfo *string `json:"StandardRoomInfo,omitempty" xml:"StandardRoomInfo,omitempty"`
	// 应用模板使用状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppTemplatesResponseBodyResultAppTemplateInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListAppTemplatesResponseBodyResultAppTemplateInfoList) GoString() string {
	return s.String()
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetAppTemplateCreator(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.AppTemplateCreator = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetAppTemplateId(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.AppTemplateId = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetAppTemplateName(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.AppTemplateName = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetComponentList(v []*string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.ComponentList = v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetConfigList(v []*ListAppTemplatesResponseBodyResultAppTemplateInfoListConfigList) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.ConfigList = v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetCreateTime(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetIntegrationMode(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.IntegrationMode = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetScene(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.Scene = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetSdkInfo(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.SdkInfo = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetStandardRoomInfo(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.StandardRoomInfo = &v
	return s
}

func (s *ListAppTemplatesResponseBodyResultAppTemplateInfoList) SetStatus(v string) *ListAppTemplatesResponseBodyResultAppTemplateInfoList {
	s.Status = &v
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
	// 改会议的申请连麦成员总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 该会议的申请连麦成员总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListApplyLinkMicUsersResponseBodyResult) SetPageTotal(v int32) *ListApplyLinkMicUsersResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListApplyLinkMicUsersResponseBodyResult) SetTotalCount(v int32) *ListApplyLinkMicUsersResponseBodyResult {
	s.TotalCount = &v
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

type ListAppsRequest struct {
	// 过滤的应用id列表
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 查询页码，参数为空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 应用状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetAppIds(v string) *ListAppsRequest {
	s.AppIds = &v
	return s
}

func (s *ListAppsRequest) SetIntegrationMode(v string) *ListAppsRequest {
	s.IntegrationMode = &v
	return s
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
	// App信息列表
	AppInfoList []*ListAppsResponseBodyResultAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	// 总页数
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyResult) SetAppInfoList(v []*ListAppsResponseBodyResultAppInfoList) *ListAppsResponseBodyResult {
	s.AppInfoList = v
	return s
}

func (s *ListAppsResponseBodyResult) SetPageTotal(v int32) *ListAppsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetTotalCount(v int32) *ListAppsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListAppsResponseBodyResultAppInfoList struct {
	// 应用配置状态
	AppConfigStatus *string `json:"AppConfigStatus,omitempty" xml:"AppConfigStatus,omitempty"`
	// 应用唯一标识符
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用Key
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用状态
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// 模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 应用组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
	// 应用创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 集成方式：- 一体化SDK：paasSDK - 样板间：standardRoom
	IntegrationMode *string `json:"IntegrationMode,omitempty" xml:"IntegrationMode,omitempty"`
	// 样板间信息
	StandardRoomInfo *string `json:"StandardRoomInfo,omitempty" xml:"StandardRoomInfo,omitempty"`
}

func (s ListAppsResponseBodyResultAppInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyResultAppInfoList) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppConfigStatus(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppConfigStatus = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppId(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppId = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppKey(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppKey = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppName(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyResultAppInfoList) SetAppStatus(v string) *ListAppsResponseBodyResultAppInfoList {
	s.AppStatus = &v
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

func (s *ListAppsResponseBodyResultAppInfoList) SetComponentList(v []*string) *ListAppsResponseBodyResultAppInfoList {
	s.ComponentList = v
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

type ListClassesRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 课程状态，0-未开课 1-上课中 2-已下课，不传则返回所有课程。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListClassesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClassesRequest) GoString() string {
	return s.String()
}

func (s *ListClassesRequest) SetAppId(v string) *ListClassesRequest {
	s.AppId = &v
	return s
}

func (s *ListClassesRequest) SetPageNumber(v int32) *ListClassesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClassesRequest) SetPageSize(v int32) *ListClassesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClassesRequest) SetStatus(v int32) *ListClassesRequest {
	s.Status = &v
	return s
}

type ListClassesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建课程返回的结果。
	Result *ListClassesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListClassesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClassesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClassesResponseBody) SetRequestId(v string) *ListClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClassesResponseBody) SetResult(v *ListClassesResponseBodyResult) *ListClassesResponseBody {
	s.Result = v
	return s
}

type ListClassesResponseBodyResult struct {
	// 课程列表信息。
	ClassList []*ListClassesResponseBodyResultClassList `json:"ClassList,omitempty" xml:"ClassList,omitempty" type:"Repeated"`
	// 是否还有下一页。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 课程总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 课程总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClassesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClassesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClassesResponseBodyResult) SetClassList(v []*ListClassesResponseBodyResultClassList) *ListClassesResponseBodyResult {
	s.ClassList = v
	return s
}

func (s *ListClassesResponseBodyResult) SetHasMore(v bool) *ListClassesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListClassesResponseBodyResult) SetPageTotal(v int32) *ListClassesResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListClassesResponseBodyResult) SetTotalCount(v int32) *ListClassesResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListClassesResponseBodyResultClassList struct {
	// 课堂唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 连麦会议唯一标识。
	ConfId *string `json:"ConfId,omitempty" xml:"ConfId,omitempty"`
	// 创建人昵称。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人ID。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 下课时间戳，毫秒。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 直播的唯一标识ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 开始上课时间戳，毫秒。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 课堂状态，0:未开始 1:上课中 2:已下课。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 课堂标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 白板ID
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
	// 白板录制ID
	WhiteboardRecordId *string `json:"WhiteboardRecordId,omitempty" xml:"WhiteboardRecordId,omitempty"`
}

func (s ListClassesResponseBodyResultClassList) String() string {
	return tea.Prettify(s)
}

func (s ListClassesResponseBodyResultClassList) GoString() string {
	return s.String()
}

func (s *ListClassesResponseBodyResultClassList) SetClassId(v string) *ListClassesResponseBodyResultClassList {
	s.ClassId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetConfId(v string) *ListClassesResponseBodyResultClassList {
	s.ConfId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetCreateNickname(v string) *ListClassesResponseBodyResultClassList {
	s.CreateNickname = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetCreateUserId(v string) *ListClassesResponseBodyResultClassList {
	s.CreateUserId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetEndTime(v int64) *ListClassesResponseBodyResultClassList {
	s.EndTime = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetLiveId(v string) *ListClassesResponseBodyResultClassList {
	s.LiveId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetRoomId(v string) *ListClassesResponseBodyResultClassList {
	s.RoomId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetStartTime(v int64) *ListClassesResponseBodyResultClassList {
	s.StartTime = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetStatus(v int32) *ListClassesResponseBodyResultClassList {
	s.Status = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetTitle(v string) *ListClassesResponseBodyResultClassList {
	s.Title = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetWhiteboardId(v string) *ListClassesResponseBodyResultClassList {
	s.WhiteboardId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetWhiteboardRecordId(v string) *ListClassesResponseBodyResultClassList {
	s.WhiteboardRecordId = &v
	return s
}

type ListClassesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClassesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClassesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClassesResponse) GoString() string {
	return s.String()
}

func (s *ListClassesResponse) SetHeaders(v map[string]*string) *ListClassesResponse {
	s.Headers = v
	return s
}

func (s *ListClassesResponse) SetBody(v *ListClassesResponseBody) *ListClassesResponse {
	s.Body = v
	return s
}

type ListCommentsRequest struct {
	// 用户的应用ID，在控制台创建应用时生成。包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询弹幕消息列表的分页页数。应该从1开始，每次分页拉取时递增。
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 查询弹幕消息列表的分页大小。最小不得小于1，最大不得超过100。如果超过100，会被截断为前100条。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 房间的唯一标识，在调用CreateRoom时返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 查询弹幕消息列表的排序方式。取值仅限0和1，其中0表示按照弹幕消息创建时间递增的顺序拉取，1表示按照弹幕消息创建时间递减的时间拉取。
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// 操作人用户ID，表示谁执行了查询弹幕消息列表的操作。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCommentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsRequest) GoString() string {
	return s.String()
}

func (s *ListCommentsRequest) SetAppId(v string) *ListCommentsRequest {
	s.AppId = &v
	return s
}

func (s *ListCommentsRequest) SetPageNum(v int32) *ListCommentsRequest {
	s.PageNum = &v
	return s
}

func (s *ListCommentsRequest) SetPageSize(v int32) *ListCommentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommentsRequest) SetRoomId(v string) *ListCommentsRequest {
	s.RoomId = &v
	return s
}

func (s *ListCommentsRequest) SetSortType(v int32) *ListCommentsRequest {
	s.SortType = &v
	return s
}

func (s *ListCommentsRequest) SetUserId(v string) *ListCommentsRequest {
	s.UserId = &v
	return s
}

type ListCommentsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用查询弹幕消息列表的返回结果。
	Result *ListCommentsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListCommentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommentsResponseBody) SetRequestId(v string) *ListCommentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommentsResponseBody) SetResult(v *ListCommentsResponseBodyResult) *ListCommentsResponseBody {
	s.Result = v
	return s
}

type ListCommentsResponseBodyResult struct {
	// 弹幕消息列表。
	CommentVOList []*ListCommentsResponseBodyResultCommentVOList `json:"CommentVOList,omitempty" xml:"CommentVOList,omitempty" type:"Repeated"`
	// 是否还有下一页数据。true表示还有，false表示没有。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 分页查询弹幕消息列表的总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 弹幕消息的总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCommentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCommentsResponseBodyResult) SetCommentVOList(v []*ListCommentsResponseBodyResultCommentVOList) *ListCommentsResponseBodyResult {
	s.CommentVOList = v
	return s
}

func (s *ListCommentsResponseBodyResult) SetHasMore(v bool) *ListCommentsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListCommentsResponseBodyResult) SetPageTotal(v int32) *ListCommentsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListCommentsResponseBodyResult) SetTotalCount(v int32) *ListCommentsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListCommentsResponseBodyResultCommentVOList struct {
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 弹幕消息的唯一ID标识。
	CommentId *string `json:"CommentId,omitempty" xml:"CommentId,omitempty"`
	// 弹幕消息的内容。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 弹幕消息的创建时间，Unix时间戳，单位：毫秒。
	CreateAt *int64 `json:"CreateAt,omitempty" xml:"CreateAt,omitempty"`
	// 扩展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 弹幕消息的发送者ID标识。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕消息发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
}

func (s ListCommentsResponseBodyResultCommentVOList) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsResponseBodyResultCommentVOList) GoString() string {
	return s.String()
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetAppId(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.AppId = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetCommentId(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.CommentId = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetContent(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.Content = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetCreateAt(v int64) *ListCommentsResponseBodyResultCommentVOList {
	s.CreateAt = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetExtension(v map[string]*string) *ListCommentsResponseBodyResultCommentVOList {
	s.Extension = v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetRoomId(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.RoomId = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetSenderId(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.SenderId = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetSenderNick(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.SenderNick = &v
	return s
}

type ListCommentsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCommentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsResponse) GoString() string {
	return s.String()
}

func (s *ListCommentsResponse) SetHeaders(v map[string]*string) *ListCommentsResponse {
	s.Headers = v
	return s
}

func (s *ListCommentsResponse) SetBody(v *ListCommentsResponseBody) *ListCommentsResponse {
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
	// 场景列表
	SceneList []*ListComponentsResponseBodyResultSceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
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

func (s *ListComponentsResponseBodyResult) SetSceneList(v []*ListComponentsResponseBodyResultSceneList) *ListComponentsResponseBodyResult {
	s.SceneList = v
	return s
}

type ListComponentsResponseBodyResultComponentCategory struct {
	// 类别下的组件列表
	List []*ListComponentsResponseBodyResultComponentCategoryList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// 组件类别
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListComponentsResponseBodyResultComponentCategory) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultComponentCategory) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultComponentCategory) SetList(v []*ListComponentsResponseBodyResultComponentCategoryList) *ListComponentsResponseBodyResultComponentCategory {
	s.List = v
	return s
}

func (s *ListComponentsResponseBodyResultComponentCategory) SetType(v string) *ListComponentsResponseBodyResultComponentCategory {
	s.Type = &v
	return s
}

type ListComponentsResponseBodyResultComponentCategoryList struct {
	// 组件名称
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// 组件类型
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// 是否使用
	InUse *string `json:"InUse,omitempty" xml:"InUse,omitempty"`
}

func (s ListComponentsResponseBodyResultComponentCategoryList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultComponentCategoryList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultComponentCategoryList) SetComponentName(v string) *ListComponentsResponseBodyResultComponentCategoryList {
	s.ComponentName = &v
	return s
}

func (s *ListComponentsResponseBodyResultComponentCategoryList) SetComponentType(v string) *ListComponentsResponseBodyResultComponentCategoryList {
	s.ComponentType = &v
	return s
}

func (s *ListComponentsResponseBodyResultComponentCategoryList) SetInUse(v string) *ListComponentsResponseBodyResultComponentCategoryList {
	s.InUse = &v
	return s
}

type ListComponentsResponseBodyResultConfigGroup struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListComponentsResponseBodyResultConfigGroup) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultConfigGroup) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultConfigGroup) SetCategory(v string) *ListComponentsResponseBodyResultConfigGroup {
	s.Category = &v
	return s
}

func (s *ListComponentsResponseBodyResultConfigGroup) SetKey(v string) *ListComponentsResponseBodyResultConfigGroup {
	s.Key = &v
	return s
}

func (s *ListComponentsResponseBodyResultConfigGroup) SetValue(v string) *ListComponentsResponseBodyResultConfigGroup {
	s.Value = &v
	return s
}

type ListComponentsResponseBodyResultSceneList struct {
	// 组件信息
	ComponentCategory []*ListComponentsResponseBodyResultSceneListComponentCategory `json:"ComponentCategory,omitempty" xml:"ComponentCategory,omitempty" type:"Repeated"`
	// 场景类别
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListComponentsResponseBodyResultSceneList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultSceneList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultSceneList) SetComponentCategory(v []*ListComponentsResponseBodyResultSceneListComponentCategory) *ListComponentsResponseBodyResultSceneList {
	s.ComponentCategory = v
	return s
}

func (s *ListComponentsResponseBodyResultSceneList) SetScene(v string) *ListComponentsResponseBodyResultSceneList {
	s.Scene = &v
	return s
}

type ListComponentsResponseBodyResultSceneListComponentCategory struct {
	// 类别下的组件列表
	List []*ListComponentsResponseBodyResultSceneListComponentCategoryList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// 组件类别
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListComponentsResponseBodyResultSceneListComponentCategory) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultSceneListComponentCategory) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategory) SetList(v []*ListComponentsResponseBodyResultSceneListComponentCategoryList) *ListComponentsResponseBodyResultSceneListComponentCategory {
	s.List = v
	return s
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategory) SetType(v string) *ListComponentsResponseBodyResultSceneListComponentCategory {
	s.Type = &v
	return s
}

type ListComponentsResponseBodyResultSceneListComponentCategoryList struct {
	// 组件名称
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// 组件类型
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// 是否使用
	InUse *string `json:"InUse,omitempty" xml:"InUse,omitempty"`
}

func (s ListComponentsResponseBodyResultSceneListComponentCategoryList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyResultSceneListComponentCategoryList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategoryList) SetComponentName(v string) *ListComponentsResponseBodyResultSceneListComponentCategoryList {
	s.ComponentName = &v
	return s
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategoryList) SetComponentType(v string) *ListComponentsResponseBodyResultSceneListComponentCategoryList {
	s.ComponentType = &v
	return s
}

func (s *ListComponentsResponseBodyResultSceneListComponentCategoryList) SetInUse(v string) *ListComponentsResponseBodyResultSceneListComponentCategoryList {
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
	// 总页数
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListConferenceUsersResponseBodyResult) SetPageTotal(v int32) *ListConferenceUsersResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListConferenceUsersResponseBodyResult) SetTotalCount(v int32) *ListConferenceUsersResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListConferenceUsersResponseBodyResultConferenceUserList struct {
	// 用户状态。
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListConferenceUsersResponseBodyResultConferenceUserList) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponseBodyResultConferenceUserList) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponseBodyResultConferenceUserList) SetStatus(v string) *ListConferenceUsersResponseBodyResultConferenceUserList {
	s.Status = &v
	return s
}

func (s *ListConferenceUsersResponseBodyResultConferenceUserList) SetUserId(v string) *ListConferenceUsersResponseBodyResultConferenceUserList {
	s.UserId = &v
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

type ListLiveRoomsRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 直播状态，0-在播 1-下播，不传则返回所有直播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLiveRoomsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsRequest) SetAppId(v string) *ListLiveRoomsRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsRequest) SetPageNumber(v int32) *ListLiveRoomsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLiveRoomsRequest) SetPageSize(v int32) *ListLiveRoomsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveRoomsRequest) SetStatus(v int32) *ListLiveRoomsRequest {
	s.Status = &v
	return s
}

type ListLiveRoomsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *ListLiveRoomsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListLiveRoomsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsResponseBody) SetRequestId(v string) *ListLiveRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRoomsResponseBody) SetResult(v *ListLiveRoomsResponseBodyResult) *ListLiveRoomsResponseBody {
	s.Result = v
	return s
}

type ListLiveRoomsResponseBodyResult struct {
	// 是否还有下一页。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 直播列表信息。
	LiveList []*ListLiveRoomsResponseBodyResultLiveList `json:"LiveList,omitempty" xml:"LiveList,omitempty" type:"Repeated"`
	// 直播总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 直播总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveRoomsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsResponseBodyResult) SetHasMore(v bool) *ListLiveRoomsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResult) SetLiveList(v []*ListLiveRoomsResponseBodyResultLiveList) *ListLiveRoomsResponseBodyResult {
	s.LiveList = v
	return s
}

func (s *ListLiveRoomsResponseBodyResult) SetPageTotal(v int32) *ListLiveRoomsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResult) SetTotalCount(v int32) *ListLiveRoomsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListLiveRoomsResponseBodyResultLiveList struct {
	// 主播ID。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 聊天ID。
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 封面。
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播状态，0-在播 1-下播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s ListLiveRoomsResponseBodyResultLiveList) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsResponseBodyResultLiveList) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetAnchorId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.AnchorId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetAnchorNick(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.AnchorNick = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetAppId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetChatId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.ChatId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetCoverUrl(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.CoverUrl = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetExtension(v map[string]*string) *ListLiveRoomsResponseBodyResultLiveList {
	s.Extension = v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetLiveId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.LiveId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetNotice(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.Notice = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetOnlineCount(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.OnlineCount = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetPv(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.Pv = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetRoomId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.RoomId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetStatus(v int32) *ListLiveRoomsResponseBodyResultLiveList {
	s.Status = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetTitle(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.Title = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetUv(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.Uv = &v
	return s
}

type ListLiveRoomsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLiveRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLiveRoomsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsResponse) SetHeaders(v map[string]*string) *ListLiveRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRoomsResponse) SetBody(v *ListLiveRoomsResponseBody) *ListLiveRoomsResponse {
	s.Body = v
	return s
}

type ListLiveRoomsByIdRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID列表。
	LiveIdList []*string `json:"LiveIdList,omitempty" xml:"LiveIdList,omitempty" type:"Repeated"`
}

func (s ListLiveRoomsByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdRequest) SetAppId(v string) *ListLiveRoomsByIdRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsByIdRequest) SetLiveIdList(v []*string) *ListLiveRoomsByIdRequest {
	s.LiveIdList = v
	return s
}

type ListLiveRoomsByIdShrinkRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID列表。
	LiveIdListShrink *string `json:"LiveIdList,omitempty" xml:"LiveIdList,omitempty"`
}

func (s ListLiveRoomsByIdShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdShrinkRequest) SetAppId(v string) *ListLiveRoomsByIdShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsByIdShrinkRequest) SetLiveIdListShrink(v string) *ListLiveRoomsByIdShrinkRequest {
	s.LiveIdListShrink = &v
	return s
}

type ListLiveRoomsByIdResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *ListLiveRoomsByIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListLiveRoomsByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdResponseBody) SetRequestId(v string) *ListLiveRoomsByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBody) SetResult(v *ListLiveRoomsByIdResponseBodyResult) *ListLiveRoomsByIdResponseBody {
	s.Result = v
	return s
}

type ListLiveRoomsByIdResponseBodyResult struct {
	// 直播列表信息。
	LiveList []*ListLiveRoomsByIdResponseBodyResultLiveList `json:"LiveList,omitempty" xml:"LiveList,omitempty" type:"Repeated"`
}

func (s ListLiveRoomsByIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdResponseBodyResult) SetLiveList(v []*ListLiveRoomsByIdResponseBodyResultLiveList) *ListLiveRoomsByIdResponseBodyResult {
	s.LiveList = v
	return s
}

type ListLiveRoomsByIdResponseBodyResultLiveList struct {
	// 主播ID。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 聊天ID。
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 封面。
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播状态，0-在播 1-下播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s ListLiveRoomsByIdResponseBodyResultLiveList) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdResponseBodyResultLiveList) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetAnchorId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.AnchorId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetAnchorNick(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.AnchorNick = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetAppId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetChatId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.ChatId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetCoverUrl(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.CoverUrl = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetExtension(v map[string]*string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Extension = v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetLiveId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.LiveId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetNotice(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Notice = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetOnlineCount(v int64) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.OnlineCount = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetPv(v int64) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Pv = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetRoomId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.RoomId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetStatus(v int32) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Status = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetTitle(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Title = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetUv(v int64) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Uv = &v
	return s
}

type ListLiveRoomsByIdResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLiveRoomsByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLiveRoomsByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdResponse) SetHeaders(v map[string]*string) *ListLiveRoomsByIdResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRoomsByIdResponse) SetBody(v *ListLiveRoomsByIdResponseBody) *ListLiveRoomsByIdResponse {
	s.Body = v
	return s
}

type ListRoomUsersRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 房间ID，最大长度36个字符。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s ListRoomUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRoomUsersRequest) SetAppId(v string) *ListRoomUsersRequest {
	s.AppId = &v
	return s
}

func (s *ListRoomUsersRequest) SetPageNumber(v int32) *ListRoomUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoomUsersRequest) SetPageSize(v int32) *ListRoomUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoomUsersRequest) SetRoomId(v string) *ListRoomUsersRequest {
	s.RoomId = &v
	return s
}

type ListRoomUsersResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
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

func (s *ListRoomUsersResponseBody) SetResult(v *ListRoomUsersResponseBodyResult) *ListRoomUsersResponseBody {
	s.Result = v
	return s
}

type ListRoomUsersResponseBodyResult struct {
	// 是否还有下一页用户列表。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 该房间的用户总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 房间用户列表信息。
	RoomUserList []*ListRoomUsersResponseBodyResultRoomUserList `json:"RoomUserList,omitempty" xml:"RoomUserList,omitempty" type:"Repeated"`
	// 该房间的用户总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoomUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponseBodyResult) SetHasMore(v bool) *ListRoomUsersResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListRoomUsersResponseBodyResult) SetPageTotal(v int32) *ListRoomUsersResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListRoomUsersResponseBodyResult) SetRoomUserList(v []*ListRoomUsersResponseBodyResultRoomUserList) *ListRoomUsersResponseBodyResult {
	s.RoomUserList = v
	return s
}

func (s *ListRoomUsersResponseBodyResult) SetTotalCount(v int32) *ListRoomUsersResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListRoomUsersResponseBodyResultRoomUserList struct {
	// 用户拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 用户昵称。
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// 用户唯一标识。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRoomUsersResponseBodyResultRoomUserList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponseBodyResultRoomUserList) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponseBodyResultRoomUserList) SetExtension(v map[string]*string) *ListRoomUsersResponseBodyResultRoomUserList {
	s.Extension = v
	return s
}

func (s *ListRoomUsersResponseBodyResultRoomUserList) SetNick(v string) *ListRoomUsersResponseBodyResultRoomUserList {
	s.Nick = &v
	return s
}

func (s *ListRoomUsersResponseBodyResultRoomUserList) SetUserId(v string) *ListRoomUsersResponseBodyResultRoomUserList {
	s.UserId = &v
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
	// 是否还有下一页房间列表。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 该应用的房间总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 房间列表信息。
	RoomInfoList []*ListRoomsResponseBodyResultRoomInfoList `json:"RoomInfoList,omitempty" xml:"RoomInfoList,omitempty" type:"Repeated"`
	// 该应用的房间总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoomsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResult) SetHasMore(v bool) *ListRoomsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListRoomsResponseBodyResult) SetPageTotal(v int32) *ListRoomsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListRoomsResponseBodyResult) SetRoomInfoList(v []*ListRoomsResponseBodyResultRoomInfoList) *ListRoomsResponseBodyResult {
	s.RoomInfoList = v
	return s
}

func (s *ListRoomsResponseBodyResult) SetTotalCount(v int32) *ListRoomsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListRoomsResponseBodyResultRoomInfoList struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间创建时间戳，单位：毫秒。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 房间拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 用户在线数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 创建房间使用的模板id。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 用户访问数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s ListRoomsResponseBodyResultRoomInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResultRoomInfoList) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetAppId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.AppId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetCreateTime(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetExtension(v map[string]*string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Extension = v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetNotice(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Notice = &v
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

func (s *ListRoomsResponseBodyResultRoomInfoList) SetRoomId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.RoomId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetRoomOwnerId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.RoomOwnerId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetTemplateId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.TemplateId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetTitle(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Title = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetUv(v int64) *ListRoomsResponseBodyResultRoomInfoList {
	s.Uv = &v
	return s
}

type ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList struct {
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetCreateTime(v int64) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetExtension(v map[string]*string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.Extension = v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetPluginId(v string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetPluginType(v string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.PluginType = &v
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

type ListSensitiveWordRequest struct {
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordRequest) SetAppId(v string) *ListSensitiveWordRequest {
	s.AppId = &v
	return s
}

type ListSensitiveWordResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用发送直播间弹幕的返回结果。
	Result *ListSensitiveWordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponseBody) SetRequestId(v string) *ListSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveWordResponseBody) SetResult(v *ListSensitiveWordResponseBodyResult) *ListSensitiveWordResponseBody {
	s.Result = v
	return s
}

type ListSensitiveWordResponseBodyResult struct {
	WordList []*string `json:"WordList,omitempty" xml:"WordList,omitempty" type:"Repeated"`
}

func (s ListSensitiveWordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponseBodyResult) SetWordList(v []*string) *ListSensitiveWordResponseBodyResult {
	s.WordList = v
	return s
}

type ListSensitiveWordResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponse) SetHeaders(v map[string]*string) *ListSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveWordResponse) SetBody(v *ListSensitiveWordResponseBody) *ListSensitiveWordResponse {
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
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 直播推流地址
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 直播状态：Created-已创建未开播，Living-直播中，End-直播结束
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PublishLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PublishLiveResponseBodyResult) SetAnchorId(v string) *PublishLiveResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetLiveId(v string) *PublishLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetLiveUrl(v string) *PublishLiveResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetPushUrl(v string) *PublishLiveResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetStatus(v string) *PublishLiveResponseBodyResult {
	s.Status = &v
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

type PublishLiveRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PublishLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *PublishLiveRoomRequest) SetAppId(v string) *PublishLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *PublishLiveRoomRequest) SetLiveId(v string) *PublishLiveRoomRequest {
	s.LiveId = &v
	return s
}

func (s *PublishLiveRoomRequest) SetUserId(v string) *PublishLiveRoomRequest {
	s.UserId = &v
	return s
}

type PublishLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *PublishLiveRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s PublishLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *PublishLiveRoomResponseBody) SetRequestId(v string) *PublishLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishLiveRoomResponseBody) SetResult(v *PublishLiveRoomResponseBodyResult) *PublishLiveRoomResponseBody {
	s.Result = v
	return s
}

type PublishLiveRoomResponseBodyResult struct {
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址。
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 直播推流地址。
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
}

func (s PublishLiveRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PublishLiveRoomResponseBodyResult) SetLiveId(v string) *PublishLiveRoomResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *PublishLiveRoomResponseBodyResult) SetLiveUrl(v string) *PublishLiveRoomResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *PublishLiveRoomResponseBodyResult) SetPushUrl(v string) *PublishLiveRoomResponseBodyResult {
	s.PushUrl = &v
	return s
}

type PublishLiveRoomResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *PublishLiveRoomResponse) SetHeaders(v map[string]*string) *PublishLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *PublishLiveRoomResponse) SetBody(v *PublishLiveRoomResponseBody) *PublishLiveRoomResponse {
	s.Body = v
	return s
}

type RejectLinkMicRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 同意者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	// 被同意用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
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

func (s *RejectLinkMicRequest) SetFromUserId(v string) *RejectLinkMicRequest {
	s.FromUserId = &v
	return s
}

func (s *RejectLinkMicRequest) SetToUserId(v string) *RejectLinkMicRequest {
	s.ToUserId = &v
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

type RemoveMemberRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 邀请者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	// 被邀请用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
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

func (s *RemoveMemberRequest) SetFromUserId(v string) *RemoveMemberRequest {
	s.FromUserId = &v
	return s
}

func (s *RemoveMemberRequest) SetToUserId(v string) *RemoveMemberRequest {
	s.ToUserId = &v
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

type SendCommentRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 发送的文本内容。最大的长度不超过256个字节。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 扩展字段，服务端仅做透传。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕消息发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
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

func (s *SendCommentRequest) SetContent(v string) *SendCommentRequest {
	s.Content = &v
	return s
}

func (s *SendCommentRequest) SetExtension(v map[string]*string) *SendCommentRequest {
	s.Extension = v
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

func (s *SendCommentRequest) SetSenderNick(v string) *SendCommentRequest {
	s.SenderNick = &v
	return s
}

type SendCommentShrinkRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 发送的文本内容。最大的长度不超过256个字节。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 扩展字段，服务端仅做透传。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕消息发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
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

func (s *SendCommentShrinkRequest) SetContent(v string) *SendCommentShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendCommentShrinkRequest) SetExtensionShrink(v string) *SendCommentShrinkRequest {
	s.ExtensionShrink = &v
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

func (s *SendCommentShrinkRequest) SetSenderNick(v string) *SendCommentShrinkRequest {
	s.SenderNick = &v
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
	// 弹幕的内容。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 弹幕的创建时间，Unix时间戳，单位：毫秒。
	CreateAt *int64 `json:"CreateAt,omitempty" xml:"CreateAt,omitempty"`
	// 扩展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 弹幕的发送者ID标识。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
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

func (s *SendCommentResponseBodyResultCommentVO) SetContent(v string) *SendCommentResponseBodyResultCommentVO {
	s.Content = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetCreateAt(v int64) *SendCommentResponseBodyResultCommentVO {
	s.CreateAt = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetExtension(v map[string]*string) *SendCommentResponseBodyResultCommentVO {
	s.Extension = v
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

type SendCustomMessageToAllRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息体内容。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
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

func (s *SendCustomMessageToAllRequest) SetBody(v string) *SendCustomMessageToAllRequest {
	s.Body = &v
	return s
}

func (s *SendCustomMessageToAllRequest) SetRoomId(v string) *SendCustomMessageToAllRequest {
	s.RoomId = &v
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

type SendCustomMessageToUsersRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息体内容。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// 消息指定的接收人ID列表。
	ReceiverList []*string `json:"ReceiverList,omitempty" xml:"ReceiverList,omitempty" type:"Repeated"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
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

func (s *SendCustomMessageToUsersRequest) SetBody(v string) *SendCustomMessageToUsersRequest {
	s.Body = &v
	return s
}

func (s *SendCustomMessageToUsersRequest) SetReceiverList(v []*string) *SendCustomMessageToUsersRequest {
	s.ReceiverList = v
	return s
}

func (s *SendCustomMessageToUsersRequest) SetRoomId(v string) *SendCustomMessageToUsersRequest {
	s.RoomId = &v
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

type SetUserAdminRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SetUserAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserAdminRequest) GoString() string {
	return s.String()
}

func (s *SetUserAdminRequest) SetAppId(v string) *SetUserAdminRequest {
	s.AppId = &v
	return s
}

func (s *SetUserAdminRequest) SetRoomId(v string) *SetUserAdminRequest {
	s.RoomId = &v
	return s
}

func (s *SetUserAdminRequest) SetUserId(v string) *SetUserAdminRequest {
	s.UserId = &v
	return s
}

type SetUserAdminResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetUserAdminResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserAdminResponseBody) SetRequestId(v string) *SetUserAdminResponseBody {
	s.RequestId = &v
	return s
}

type SetUserAdminResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetUserAdminResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetUserAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserAdminResponse) GoString() string {
	return s.String()
}

func (s *SetUserAdminResponse) SetHeaders(v map[string]*string) *SetUserAdminResponse {
	s.Headers = v
	return s
}

func (s *SetUserAdminResponse) SetBody(v *SetUserAdminResponseBody) *SetUserAdminResponse {
	s.Body = v
	return s
}

type StopClassRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 操作者用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StopClassRequest) String() string {
	return tea.Prettify(s)
}

func (s StopClassRequest) GoString() string {
	return s.String()
}

func (s *StopClassRequest) SetAppId(v string) *StopClassRequest {
	s.AppId = &v
	return s
}

func (s *StopClassRequest) SetClassId(v string) *StopClassRequest {
	s.ClassId = &v
	return s
}

func (s *StopClassRequest) SetUserId(v string) *StopClassRequest {
	s.UserId = &v
	return s
}

type StopClassResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopClassResponseBody) GoString() string {
	return s.String()
}

func (s *StopClassResponseBody) SetRequestId(v string) *StopClassResponseBody {
	s.RequestId = &v
	return s
}

type StopClassResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopClassResponse) String() string {
	return tea.Prettify(s)
}

func (s StopClassResponse) GoString() string {
	return s.String()
}

func (s *StopClassResponse) SetHeaders(v map[string]*string) *StopClassResponse {
	s.Headers = v
	return s
}

func (s *StopClassResponse) SetBody(v *StopClassResponseBody) *StopClassResponse {
	s.Body = v
	return s
}

type StopLiveRequest struct {
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID，最大长度36位
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建直播用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *StopLiveRequest) SetLiveId(v string) *StopLiveRequest {
	s.LiveId = &v
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

type StopLiveRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StopLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *StopLiveRoomRequest) SetAppId(v string) *StopLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *StopLiveRoomRequest) SetLiveId(v string) *StopLiveRoomRequest {
	s.LiveId = &v
	return s
}

func (s *StopLiveRoomRequest) SetUserId(v string) *StopLiveRoomRequest {
	s.UserId = &v
	return s
}

type StopLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveRoomResponseBody) SetRequestId(v string) *StopLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

type StopLiveRoomResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *StopLiveRoomResponse) SetHeaders(v map[string]*string) *StopLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *StopLiveRoomResponse) SetBody(v *StopLiveRoomResponseBody) *StopLiveRoomResponse {
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

type UpdateAppTemplateRequest struct {
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 组件列表
	ComponentList []*string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Repeated"`
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

func (s *UpdateAppTemplateRequest) SetComponentList(v []*string) *UpdateAppTemplateRequest {
	s.ComponentList = v
	return s
}

type UpdateAppTemplateShrinkRequest struct {
	// 应用模板唯一标识
	AppTemplateId *string `json:"AppTemplateId,omitempty" xml:"AppTemplateId,omitempty"`
	// 应用模板名称
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// 组件列表
	ComponentListShrink *string `json:"ComponentList,omitempty" xml:"ComponentList,omitempty"`
}

func (s UpdateAppTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateShrinkRequest) SetAppTemplateId(v string) *UpdateAppTemplateShrinkRequest {
	s.AppTemplateId = &v
	return s
}

func (s *UpdateAppTemplateShrinkRequest) SetAppTemplateName(v string) *UpdateAppTemplateShrinkRequest {
	s.AppTemplateName = &v
	return s
}

func (s *UpdateAppTemplateShrinkRequest) SetComponentListShrink(v string) *UpdateAppTemplateShrinkRequest {
	s.ComponentListShrink = &v
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
	// 返回结果
	Result *UpdateAppTemplateConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s *UpdateAppTemplateConfigResponseBody) SetResult(v *UpdateAppTemplateConfigResponseBodyResult) *UpdateAppTemplateConfigResponseBody {
	s.Result = v
	return s
}

type UpdateAppTemplateConfigResponseBodyResult struct {
	// 配置日志列表
	ConfigLogs []*UpdateAppTemplateConfigResponseBodyResultConfigLogs `json:"ConfigLogs,omitempty" xml:"ConfigLogs,omitempty" type:"Repeated"`
}

func (s UpdateAppTemplateConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateConfigResponseBodyResult) SetConfigLogs(v []*UpdateAppTemplateConfigResponseBodyResultConfigLogs) *UpdateAppTemplateConfigResponseBodyResult {
	s.ConfigLogs = v
	return s
}

type UpdateAppTemplateConfigResponseBodyResultConfigLogs struct {
	// 日志标示
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 日志内容
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateAppTemplateConfigResponseBodyResultConfigLogs) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppTemplateConfigResponseBodyResultConfigLogs) GoString() string {
	return s.String()
}

func (s *UpdateAppTemplateConfigResponseBodyResultConfigLogs) SetCode(v string) *UpdateAppTemplateConfigResponseBodyResultConfigLogs {
	s.Code = &v
	return s
}

func (s *UpdateAppTemplateConfigResponseBodyResultConfigLogs) SetMessage(v string) *UpdateAppTemplateConfigResponseBodyResultConfigLogs {
	s.Message = &v
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

type UpdateClassRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 创建人用户昵称，1~32个字符。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人用户ID，仅支持中英文数字，下划线，中划线，1~36个字符。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 课堂标题，1~32个字符。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateClassRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClassRequest) GoString() string {
	return s.String()
}

func (s *UpdateClassRequest) SetAppId(v string) *UpdateClassRequest {
	s.AppId = &v
	return s
}

func (s *UpdateClassRequest) SetClassId(v string) *UpdateClassRequest {
	s.ClassId = &v
	return s
}

func (s *UpdateClassRequest) SetCreateNickname(v string) *UpdateClassRequest {
	s.CreateNickname = &v
	return s
}

func (s *UpdateClassRequest) SetCreateUserId(v string) *UpdateClassRequest {
	s.CreateUserId = &v
	return s
}

func (s *UpdateClassRequest) SetTitle(v string) *UpdateClassRequest {
	s.Title = &v
	return s
}

type UpdateClassResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClassResponseBody) SetRequestId(v string) *UpdateClassResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClassResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClassResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClassResponse) GoString() string {
	return s.String()
}

func (s *UpdateClassResponse) SetHeaders(v map[string]*string) *UpdateClassResponse {
	s.Headers = v
	return s
}

func (s *UpdateClassResponse) SetBody(v *UpdateClassResponseBody) *UpdateClassResponse {
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

type UpdateLiveRequest struct {
	// 直播简介，支持中英文，最大长度2048位
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播标题，支持中英文，最大长度256位
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRequest) SetIntroduction(v string) *UpdateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *UpdateLiveRequest) SetLiveId(v string) *UpdateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRequest) SetTitle(v string) *UpdateLiveRequest {
	s.Title = &v
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

type UpdateLiveRoomRequest struct {
	// 主播id，仅支持英文和数字，最大长度36位。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 封面，支持图片地址链接格式
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRoomRequest) SetAnchorId(v string) *UpdateLiveRoomRequest {
	s.AnchorId = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetAnchorNick(v string) *UpdateLiveRoomRequest {
	s.AnchorNick = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetAppId(v string) *UpdateLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetCoverUrl(v string) *UpdateLiveRoomRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetExtension(v map[string]*string) *UpdateLiveRoomRequest {
	s.Extension = v
	return s
}

func (s *UpdateLiveRoomRequest) SetLiveId(v string) *UpdateLiveRoomRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetNotice(v string) *UpdateLiveRoomRequest {
	s.Notice = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetTitle(v string) *UpdateLiveRoomRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetUserId(v string) *UpdateLiveRoomRequest {
	s.UserId = &v
	return s
}

type UpdateLiveRoomShrinkRequest struct {
	// 主播id，仅支持英文和数字，最大长度36位。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 封面，支持图片地址链接格式
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLiveRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRoomShrinkRequest) SetAnchorId(v string) *UpdateLiveRoomShrinkRequest {
	s.AnchorId = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetAnchorNick(v string) *UpdateLiveRoomShrinkRequest {
	s.AnchorNick = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetAppId(v string) *UpdateLiveRoomShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetCoverUrl(v string) *UpdateLiveRoomShrinkRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetExtensionShrink(v string) *UpdateLiveRoomShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetLiveId(v string) *UpdateLiveRoomShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetNotice(v string) *UpdateLiveRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetTitle(v string) *UpdateLiveRoomShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetUserId(v string) *UpdateLiveRoomShrinkRequest {
	s.UserId = &v
	return s
}

type UpdateLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveRoomResponseBody) SetRequestId(v string) *UpdateLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLiveRoomResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveRoomResponse) SetHeaders(v map[string]*string) *UpdateLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveRoomResponse) SetBody(v *UpdateLiveRoomResponseBody) *UpdateLiveRoomResponse {
	s.Body = v
	return s
}

type UpdateRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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

func (s *UpdateRoomRequest) SetExtension(v map[string]*string) *UpdateRoomRequest {
	s.Extension = v
	return s
}

func (s *UpdateRoomRequest) SetNotice(v string) *UpdateRoomRequest {
	s.Notice = &v
	return s
}

func (s *UpdateRoomRequest) SetRoomId(v string) *UpdateRoomRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateRoomRequest) SetRoomOwnerId(v string) *UpdateRoomRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *UpdateRoomRequest) SetTitle(v string) *UpdateRoomRequest {
	s.Title = &v
	return s
}

type UpdateRoomShrinkRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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

func (s *UpdateRoomShrinkRequest) SetExtensionShrink(v string) *UpdateRoomShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetNotice(v string) *UpdateRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetRoomId(v string) *UpdateRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetRoomOwnerId(v string) *UpdateRoomShrinkRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetTitle(v string) *UpdateRoomShrinkRequest {
	s.Title = &v
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

func (client *Client) AddMemberWithOptions(request *AddMemberRequest, runtime *util.RuntimeOptions) (_result *AddMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.FromUserId)) {
		body["FromUserId"] = request.FromUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ToUserId)) {
		body["ToUserId"] = request.ToUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMember"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) AgreeLinkMicWithOptions(request *AgreeLinkMicRequest, runtime *util.RuntimeOptions) (_result *AgreeLinkMicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.FromUserId)) {
		body["FromUserId"] = request.FromUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ToUserId)) {
		body["ToUserId"] = request.ToUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AgreeLinkMic"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AgreeLinkMicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ApplyLinkMicWithOptions(request *ApplyLinkMicRequest, runtime *util.RuntimeOptions) (_result *ApplyLinkMicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyLinkMic"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyLinkMicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) BanAllCommentWithOptions(request *BanAllCommentRequest, runtime *util.RuntimeOptions) (_result *BanAllCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BanAllComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BanAllCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) BanCommentWithOptions(request *BanCommentRequest, runtime *util.RuntimeOptions) (_result *BanCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BanCommentTime)) {
		body["BanCommentTime"] = request.BanCommentTime
	}

	if !tea.BoolValue(util.IsUnset(request.BanCommentUser)) {
		body["BanCommentUser"] = request.BanCommentUser
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BanComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BanCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CancelApplyLinkMicWithOptions(request *CancelApplyLinkMicRequest, runtime *util.RuntimeOptions) (_result *CancelApplyLinkMicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelApplyLinkMic"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelApplyLinkMicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CancelBanAllCommentWithOptions(request *CancelBanAllCommentRequest, runtime *util.RuntimeOptions) (_result *CancelBanAllCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelBanAllComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelBanAllCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CancelBanCommentWithOptions(request *CancelBanCommentRequest, runtime *util.RuntimeOptions) (_result *CancelBanCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BanCommentUser)) {
		body["BanCommentUser"] = request.BanCommentUser
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelBanComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelBanCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CancelUserAdminWithOptions(request *CancelUserAdminRequest, runtime *util.RuntimeOptions) (_result *CancelUserAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelUserAdmin"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelUserAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelUserAdmin(request *CancelUserAdminRequest) (_result *CancelUserAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelUserAdminResponse{}
	_body, _err := client.CancelUserAdminWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppTemplateId)) {
		body["AppTemplateId"] = request.AppTemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTemplateName)) {
		body["AppTemplateName"] = request.AppTemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentListShrink)) {
		body["ComponentList"] = request.ComponentListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationMode)) {
		body["IntegrationMode"] = request.IntegrationMode
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppTemplate"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateClassWithOptions(request *CreateClassRequest, runtime *util.RuntimeOptions) (_result *CreateClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateNickname)) {
		body["CreateNickname"] = request.CreateNickname
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClass"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClass(request *CreateClassRequest) (_result *CreateClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClassResponse{}
	_body, _err := client.CreateClassWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConference"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateLiveWithOptions(request *CreateLiveRequest, runtime *util.RuntimeOptions) (_result *CreateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorId)) {
		body["AnchorId"] = request.AnchorId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeLevel)) {
		body["CodeLevel"] = request.CodeLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		body["Introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateLiveRoomWithOptions(tmpReq *CreateLiveRoomRequest, runtime *util.RuntimeOptions) (_result *CreateLiveRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateLiveRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorId)) {
		body["AnchorId"] = request.AnchorId
	}

	if !tea.BoolValue(util.IsUnset(request.AnchorNick)) {
		body["AnchorNick"] = request.AnchorNick
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["CoverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["Notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLiveRoom(request *CreateLiveRoomRequest) (_result *CreateLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveRoomResponse{}
	_body, _err := client.CreateLiveRoomWithOptions(request, runtime)
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["Notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomOwnerId)) {
		body["RoomOwnerId"] = request.RoomOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateSensitiveWordWithOptions(tmpReq *CreateSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *CreateSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSensitiveWordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.WordList)) {
		request.WordListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WordList, tea.String("WordList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.WordListShrink)) {
		body["WordList"] = request.WordListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSensitiveWord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSensitiveWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSensitiveWord(request *CreateSensitiveWordRequest) (_result *CreateSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSensitiveWordResponse{}
	_body, _err := client.CreateSensitiveWordWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteAppTemplateWithOptions(request *DeleteAppTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteAppTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTemplateId)) {
		body["AppTemplateId"] = request.AppTemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppTemplate"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteClassWithOptions(request *DeleteClassRequest, runtime *util.RuntimeOptions) (_result *DeleteClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteClass"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClass(request *DeleteClassRequest) (_result *DeleteClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClassResponse{}
	_body, _err := client.DeleteClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCommentWithOptions(request *DeleteCommentRequest, runtime *util.RuntimeOptions) (_result *DeleteCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommentIdList)) {
		bodyFlat["CommentIdList"] = request.CommentIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteComment(request *DeleteCommentRequest) (_result *DeleteCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCommentResponse{}
	_body, _err := client.DeleteCommentWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConference"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteLiveWithOptions(request *DeleteLiveRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteLiveRoomWithOptions(request *DeleteLiveRoomRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveRoom(request *DeleteLiveRoomRequest) (_result *DeleteLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveRoomResponse{}
	_body, _err := client.DeleteLiveRoomWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteSensitiveWordWithOptions(tmpReq *DeleteSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *DeleteSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteSensitiveWordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.WordList)) {
		request.WordListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WordList, tea.String("WordList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.WordListShrink)) {
		body["WordList"] = request.WordListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSensitiveWord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSensitiveWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSensitiveWord(request *DeleteSensitiveWordRequest) (_result *DeleteSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSensitiveWordResponse{}
	_body, _err := client.DeleteSensitiveWordWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApp"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetAppTemplateWithOptions(request *GetAppTemplateRequest, runtime *util.RuntimeOptions) (_result *GetAppTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTemplateId)) {
		body["AppTemplateId"] = request.AppTemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppTemplate"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetAuthTokenWithOptions(request *GetAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthToken"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetClassDetailWithOptions(request *GetClassDetailRequest, runtime *util.RuntimeOptions) (_result *GetClassDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClassDetail"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClassDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClassDetail(request *GetClassDetailRequest) (_result *GetClassDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClassDetailResponse{}
	_body, _err := client.GetClassDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClassRecordWithOptions(request *GetClassRecordRequest, runtime *util.RuntimeOptions) (_result *GetClassRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClassRecord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClassRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClassRecord(request *GetClassRecordRequest) (_result *GetClassRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClassRecordResponse{}
	_body, _err := client.GetClassRecordWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConference"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetDomainOwnerVerifyContentWithOptions(request *GetDomainOwnerVerifyContentRequest, runtime *util.RuntimeOptions) (_result *GetDomainOwnerVerifyContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveDomainName)) {
		body["LiveDomainName"] = request.LiveDomainName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainOwnerVerifyContent"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainOwnerVerifyContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetLiveWithOptions(request *GetLiveRequest, runtime *util.RuntimeOptions) (_result *GetLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveDomainListShrink)) {
		body["LiveDomainList"] = request.LiveDomainListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LiveDomainType)) {
		body["LiveDomainType"] = request.LiveDomainType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveDomainStatus"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveDomainStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetLiveRecordWithOptions(request *GetLiveRecordRequest, runtime *util.RuntimeOptions) (_result *GetLiveRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveRecord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveRecord(request *GetLiveRecordRequest) (_result *GetLiveRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveRecordResponse{}
	_body, _err := client.GetLiveRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveRoomWithOptions(request *GetLiveRoomRequest, runtime *util.RuntimeOptions) (_result *GetLiveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveRoom(request *GetLiveRoomRequest) (_result *GetLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveRoomResponse{}
	_body, _err := client.GetLiveRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveRoomStatisticsWithOptions(request *GetLiveRoomStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetLiveRoomStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveRoomStatistics"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveRoomStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveRoomStatistics(request *GetLiveRoomStatisticsRequest) (_result *GetLiveRoomStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveRoomStatisticsResponse{}
	_body, _err := client.GetLiveRoomStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveRoomUserStatisticsWithOptions(request *GetLiveRoomUserStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetLiveRoomUserStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveRoomUserStatistics"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveRoomUserStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveRoomUserStatistics(request *GetLiveRoomUserStatisticsRequest) (_result *GetLiveRoomUserStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveRoomUserStatisticsResponse{}
	_body, _err := client.GetLiveRoomUserStatisticsWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetStandardRoomHttpsCertificateWithOptions(request *GetStandardRoomHttpsCertificateRequest, runtime *util.RuntimeOptions) (_result *GetStandardRoomHttpsCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		body["CertificateId"] = request.CertificateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStandardRoomHttpsCertificate"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStandardRoomHttpsCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStandardRoomHttpsCertificate(request *GetStandardRoomHttpsCertificateRequest) (_result *GetStandardRoomHttpsCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStandardRoomHttpsCertificateResponse{}
	_body, _err := client.GetStandardRoomHttpsCertificateWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserNick)) {
		body["UserNick"] = request.UserNick
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStandardRoomJumpUrl"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStandardRoomJumpUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) KickRoomUserWithOptions(request *KickRoomUserRequest, runtime *util.RuntimeOptions) (_result *KickRoomUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BlockTime)) {
		body["BlockTime"] = request.BlockTime
	}

	if !tea.BoolValue(util.IsUnset(request.KickUser)) {
		body["KickUser"] = request.KickUser
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("KickRoomUser"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KickRoomUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KickRoomUser(request *KickRoomUserRequest) (_result *KickRoomUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KickRoomUserResponse{}
	_body, _err := client.KickRoomUserWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppTemplates"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListApplyLinkMicUsersWithOptions(request *ListApplyLinkMicUsersRequest, runtime *util.RuntimeOptions) (_result *ListApplyLinkMicUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplyLinkMicUsers"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplyLinkMicUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationMode)) {
		body["IntegrationMode"] = request.IntegrationMode
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListClassesWithOptions(request *ListClassesRequest, runtime *util.RuntimeOptions) (_result *ListClassesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClasses"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClassesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClasses(request *ListClassesRequest) (_result *ListClassesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClassesResponse{}
	_body, _err := client.ListClassesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommentsWithOptions(request *ListCommentsRequest, runtime *util.RuntimeOptions) (_result *ListCommentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		body["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListComments"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComments(request *ListCommentsRequest) (_result *ListCommentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommentsResponse{}
	_body, _err := client.ListCommentsWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppTemplateId)) {
		body["AppTemplateId"] = request.AppTemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListComponents"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListConferenceUsersWithOptions(request *ListConferenceUsersRequest, runtime *util.RuntimeOptions) (_result *ListConferenceUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConferenceUsers"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConferenceUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListLiveRoomsWithOptions(request *ListLiveRoomsRequest, runtime *util.RuntimeOptions) (_result *ListLiveRoomsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLiveRooms"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLiveRoomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveRooms(request *ListLiveRoomsRequest) (_result *ListLiveRoomsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveRoomsResponse{}
	_body, _err := client.ListLiveRoomsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveRoomsByIdWithOptions(tmpReq *ListLiveRoomsByIdRequest, runtime *util.RuntimeOptions) (_result *ListLiveRoomsByIdResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListLiveRoomsByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LiveIdList)) {
		request.LiveIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LiveIdList, tea.String("LiveIdList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveIdListShrink)) {
		body["LiveIdList"] = request.LiveIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLiveRoomsById"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLiveRoomsByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveRoomsById(request *ListLiveRoomsByIdRequest) (_result *ListLiveRoomsByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveRoomsByIdResponse{}
	_body, _err := client.ListLiveRoomsByIdWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoomUsers"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRoomUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListRoomsWithOptions(request *ListRoomsRequest, runtime *util.RuntimeOptions) (_result *ListRoomsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRooms"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRoomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListSensitiveWordWithOptions(request *ListSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *ListSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSensitiveWord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSensitiveWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSensitiveWord(request *ListSensitiveWordRequest) (_result *ListSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSensitiveWordResponse{}
	_body, _err := client.ListSensitiveWordWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) PublishLiveRoomWithOptions(request *PublishLiveRoomRequest, runtime *util.RuntimeOptions) (_result *PublishLiveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishLiveRoom(request *PublishLiveRoomRequest) (_result *PublishLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishLiveRoomResponse{}
	_body, _err := client.PublishLiveRoomWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.FromUserId)) {
		body["FromUserId"] = request.FromUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ToUserId)) {
		body["ToUserId"] = request.ToUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectLinkMic"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectLinkMicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RemoveMemberWithOptions(request *RemoveMemberRequest, runtime *util.RuntimeOptions) (_result *RemoveMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.FromUserId)) {
		body["FromUserId"] = request.FromUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ToUserId)) {
		body["ToUserId"] = request.ToUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveMember"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		body["SenderId"] = request.SenderId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderNick)) {
		body["SenderNick"] = request.SenderNick
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SendCustomMessageToAllWithOptions(request *SendCustomMessageToAllRequest, runtime *util.RuntimeOptions) (_result *SendCustomMessageToAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCustomMessageToAll"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCustomMessageToAllResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SendCustomMessageToUsersWithOptions(request *SendCustomMessageToUsersRequest, runtime *util.RuntimeOptions) (_result *SendCustomMessageToUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReceiverList)) {
		bodyFlat["ReceiverList"] = request.ReceiverList
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCustomMessageToUsers"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCustomMessageToUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SetUserAdminWithOptions(request *SetUserAdminRequest, runtime *util.RuntimeOptions) (_result *SetUserAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetUserAdmin"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetUserAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetUserAdmin(request *SetUserAdminRequest) (_result *SetUserAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetUserAdminResponse{}
	_body, _err := client.SetUserAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopClassWithOptions(request *StopClassRequest, runtime *util.RuntimeOptions) (_result *StopClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopClass"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopClass(request *StopClassRequest) (_result *StopClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopClassResponse{}
	_body, _err := client.StopClassWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) StopLiveRoomWithOptions(request *StopLiveRoomRequest, runtime *util.RuntimeOptions) (_result *StopLiveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLiveRoom(request *StopLiveRoomRequest) (_result *StopLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLiveRoomResponse{}
	_body, _err := client.StopLiveRoomWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppStatus)) {
		body["AppStatus"] = request.AppStatus
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApp"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateAppTemplateWithOptions(tmpReq *UpdateAppTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateAppTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAppTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ComponentList)) {
		request.ComponentListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComponentList, tea.String("ComponentList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTemplateId)) {
		body["AppTemplateId"] = request.AppTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.AppTemplateName)) {
		body["AppTemplateName"] = request.AppTemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentListShrink)) {
		body["ComponentList"] = request.ComponentListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppTemplate"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTemplateId)) {
		body["AppTemplateId"] = request.AppTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigListShrink)) {
		body["ConfigList"] = request.ConfigListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppTemplateConfig"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppTemplateConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateClassWithOptions(request *UpdateClassRequest, runtime *util.RuntimeOptions) (_result *UpdateClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateNickname)) {
		body["CreateNickname"] = request.CreateNickname
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClass"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateClass(request *UpdateClassRequest) (_result *UpdateClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClassResponse{}
	_body, _err := client.UpdateClassWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConference"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateLiveWithOptions(request *UpdateLiveRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		body["Introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateLiveRoomWithOptions(tmpReq *UpdateLiveRoomRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateLiveRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorId)) {
		body["AnchorId"] = request.AnchorId
	}

	if !tea.BoolValue(util.IsUnset(request.AnchorNick)) {
		body["AnchorNick"] = request.AnchorNick
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["CoverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["Notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveRoom(request *UpdateLiveRoomRequest) (_result *UpdateLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveRoomResponse{}
	_body, _err := client.UpdateLiveRoomWithOptions(request, runtime)
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["Notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomOwnerId)) {
		body["RoomOwnerId"] = request.RoomOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) VerifyDomainOwnerWithOptions(request *VerifyDomainOwnerRequest, runtime *util.RuntimeOptions) (_result *VerifyDomainOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveDomainName)) {
		body["LiveDomainName"] = request.LiveDomainName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyDomainOwner"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyDomainOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
