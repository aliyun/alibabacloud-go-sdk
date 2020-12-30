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

type AddGroupMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMemberRequest) SetAccessToken(v string) *AddGroupMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *AddGroupMemberRequest) SetOrganizationId(v string) *AddGroupMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *AddGroupMemberRequest) SetSubUserId(v string) *AddGroupMemberRequest {
	s.SubUserId = &v
	return s
}

func (s *AddGroupMemberRequest) SetClientToken(v string) *AddGroupMemberRequest {
	s.ClientToken = &v
	return s
}

type AddGroupMemberResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*AddGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s AddGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponseBody) SetRequestId(v string) *AddGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetErrorCode(v string) *AddGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetErrorMessage(v string) *AddGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetSuccess(v bool) *AddGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetResult(v []*AddGroupMemberResponseBodyResult) *AddGroupMemberResponseBody {
	s.Result = v
	return s
}

type AddGroupMemberResponseBodyResult struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponseBodyResult) SetExternUserId(v string) *AddGroupMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetEmail(v string) *AddGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetAvatarUrl(v string) *AddGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetState(v string) *AddGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetAccessLevel(v int32) *AddGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetId(v int64) *AddGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

type AddGroupMemberResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponse) SetHeaders(v map[string]*string) *AddGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *AddGroupMemberResponse) SetBody(v *AddGroupMemberResponseBody) *AddGroupMemberResponse {
	s.Body = v
	return s
}

type AddRepositoryMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddRepositoryMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberRequest) SetAccessToken(v string) *AddRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetOrganizationId(v string) *AddRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetSubUserId(v string) *AddRepositoryMemberRequest {
	s.SubUserId = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetClientToken(v string) *AddRepositoryMemberRequest {
	s.ClientToken = &v
	return s
}

type AddRepositoryMemberResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*AddRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s AddRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponseBody) SetRequestId(v string) *AddRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetErrorCode(v string) *AddRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetErrorMessage(v string) *AddRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetSuccess(v bool) *AddRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetResult(v []*AddRepositoryMemberResponseBodyResult) *AddRepositoryMemberResponseBody {
	s.Result = v
	return s
}

type AddRepositoryMemberResponseBodyResult struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponseBodyResult) SetExternUserId(v string) *AddRepositoryMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetEmail(v string) *AddRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *AddRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetState(v string) *AddRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *AddRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetId(v int64) *AddRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

type AddRepositoryMemberResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRepositoryMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponse) SetHeaders(v map[string]*string) *AddRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *AddRepositoryMemberResponse) SetBody(v *AddRepositoryMemberResponseBody) *AddRepositoryMemberResponse {
	s.Body = v
	return s
}

type AddWebhookRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AddWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookRequest) GoString() string {
	return s.String()
}

func (s *AddWebhookRequest) SetAccessToken(v string) *AddWebhookRequest {
	s.AccessToken = &v
	return s
}

func (s *AddWebhookRequest) SetOrganizationId(v string) *AddWebhookRequest {
	s.OrganizationId = &v
	return s
}

type AddWebhookResponseBody struct {
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *AddWebhookResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s AddWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *AddWebhookResponseBody) SetRequestId(v string) *AddWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWebhookResponseBody) SetErrorCode(v string) *AddWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddWebhookResponseBody) SetErrorMessage(v string) *AddWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddWebhookResponseBody) SetSuccess(v bool) *AddWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *AddWebhookResponseBody) SetResult(v *AddWebhookResponseBodyResult) *AddWebhookResponseBody {
	s.Result = v
	return s
}

type AddWebhookResponseBodyResult struct {
	PushEvents            *bool   `json:"PushEvents,omitempty" xml:"PushEvents,omitempty"`
	BuildEvents           *bool   `json:"BuildEvents,omitempty" xml:"BuildEvents,omitempty"`
	ProjectId             *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	CreatedAt             *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	TagPushEvents         *bool   `json:"TagPushEvents,omitempty" xml:"TagPushEvents,omitempty"`
	IssuesEvents          *bool   `json:"IssuesEvents,omitempty" xml:"IssuesEvents,omitempty"`
	Url                   *string `json:"Url,omitempty" xml:"Url,omitempty"`
	LastTestResult        *string `json:"LastTestResult,omitempty" xml:"LastTestResult,omitempty"`
	MergeRequestsEvents   *bool   `json:"MergeRequestsEvents,omitempty" xml:"MergeRequestsEvents,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NoteEvents            *bool   `json:"NoteEvents,omitempty" xml:"NoteEvents,omitempty"`
	SecretToken           *string `json:"SecretToken,omitempty" xml:"SecretToken,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	EnableSslVerification *bool   `json:"EnableSslVerification,omitempty" xml:"EnableSslVerification,omitempty"`
}

func (s AddWebhookResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddWebhookResponseBodyResult) SetPushEvents(v bool) *AddWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetBuildEvents(v bool) *AddWebhookResponseBodyResult {
	s.BuildEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetProjectId(v int64) *AddWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetCreatedAt(v string) *AddWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetTagPushEvents(v bool) *AddWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetIssuesEvents(v bool) *AddWebhookResponseBodyResult {
	s.IssuesEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetUrl(v string) *AddWebhookResponseBodyResult {
	s.Url = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetLastTestResult(v string) *AddWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *AddWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetDescription(v string) *AddWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetNoteEvents(v bool) *AddWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetSecretToken(v string) *AddWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetId(v int64) *AddWebhookResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetEnableSslVerification(v bool) *AddWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

type AddWebhookResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookResponse) GoString() string {
	return s.String()
}

func (s *AddWebhookResponse) SetHeaders(v map[string]*string) *AddWebhookResponse {
	s.Headers = v
	return s
}

func (s *AddWebhookResponse) SetBody(v *AddWebhookResponseBody) *AddWebhookResponse {
	s.Body = v
	return s
}

type CreateBranchRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s CreateBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchRequest) GoString() string {
	return s.String()
}

func (s *CreateBranchRequest) SetAccessToken(v string) *CreateBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateBranchRequest) SetOrganizationId(v string) *CreateBranchRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateBranchRequest) SetSubUserId(v string) *CreateBranchRequest {
	s.SubUserId = &v
	return s
}

type CreateBranchResponseBody struct {
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *CreateBranchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBody) SetRequestId(v string) *CreateBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBranchResponseBody) SetErrorCode(v string) *CreateBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateBranchResponseBody) SetErrorMessage(v string) *CreateBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateBranchResponseBody) SetSuccess(v bool) *CreateBranchResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBranchResponseBody) SetResult(v *CreateBranchResponseBodyResult) *CreateBranchResponseBody {
	s.Result = v
	return s
}

type CreateBranchResponseBodyResult struct {
	ProtectedBranch *bool                                     `json:"ProtectedBranch,omitempty" xml:"ProtectedBranch,omitempty"`
	CommitInfo      *CreateBranchResponseBodyResultCommitInfo `json:"CommitInfo,omitempty" xml:"CommitInfo,omitempty" type:"Struct"`
	BranchName      *string                                   `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
}

func (s CreateBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBodyResult) SetProtectedBranch(v bool) *CreateBranchResponseBodyResult {
	s.ProtectedBranch = &v
	return s
}

func (s *CreateBranchResponseBodyResult) SetCommitInfo(v *CreateBranchResponseBodyResultCommitInfo) *CreateBranchResponseBodyResult {
	s.CommitInfo = v
	return s
}

func (s *CreateBranchResponseBodyResult) SetBranchName(v string) *CreateBranchResponseBodyResult {
	s.BranchName = &v
	return s
}

type CreateBranchResponseBodyResultCommitInfo struct {
	ShortId        *string   `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	AuthorDate     *string   `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	AuthorName     *string   `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommitterName  *string   `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	AuthorEmail    *string   `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	CommitterEmail *string   `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	CommittedDate  *string   `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
}

func (s CreateBranchResponseBodyResultCommitInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBodyResultCommitInfo) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetShortId(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.ShortId = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetAuthorDate(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.AuthorDate = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetCreatedAt(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.CreatedAt = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetMessage(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.Message = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetParentIds(v []*string) *CreateBranchResponseBodyResultCommitInfo {
	s.ParentIds = v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetAuthorName(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.AuthorName = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetCommitterName(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.CommitterName = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetTitle(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.Title = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetAuthorEmail(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.AuthorEmail = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetCommitterEmail(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.CommitterEmail = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetId(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.Id = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetCommittedDate(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.CommittedDate = &v
	return s
}

type CreateBranchResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponse) GoString() string {
	return s.String()
}

func (s *CreateBranchResponse) SetHeaders(v map[string]*string) *CreateBranchResponse {
	s.Headers = v
	return s
}

func (s *CreateBranchResponse) SetBody(v *CreateBranchResponseBody) *CreateBranchResponse {
	s.Body = v
	return s
}

type CreateFileRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) SetAccessToken(v string) *CreateFileRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateFileRequest) SetOrganizationId(v string) *CreateFileRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateFileRequest) SetSubUserId(v string) *CreateFileRequest {
	s.SubUserId = &v
	return s
}

func (s *CreateFileRequest) SetClientToken(v string) *CreateFileRequest {
	s.ClientToken = &v
	return s
}

type CreateFileResponseBody struct {
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *CreateFileResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) SetRequestId(v string) *CreateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileResponseBody) SetErrorCode(v string) *CreateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFileResponseBody) SetErrorMessage(v string) *CreateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFileResponseBody) SetSuccess(v bool) *CreateFileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFileResponseBody) SetResult(v *CreateFileResponseBodyResult) *CreateFileResponseBody {
	s.Result = v
	return s
}

type CreateFileResponseBodyResult struct {
	FilePath   *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
}

func (s CreateFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBodyResult) SetFilePath(v string) *CreateFileResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *CreateFileResponseBodyResult) SetBranchName(v string) *CreateFileResponseBodyResult {
	s.BranchName = &v
	return s
}

type CreateFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponse) GoString() string {
	return s.String()
}

func (s *CreateFileResponse) SetHeaders(v map[string]*string) *CreateFileResponse {
	s.Headers = v
	return s
}

func (s *CreateFileResponse) SetBody(v *CreateFileResponseBody) *CreateFileResponse {
	s.Body = v
	return s
}

type CreateMergeRequestRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s CreateMergeRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestRequest) SetAccessToken(v string) *CreateMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateMergeRequestRequest) SetOrganizationId(v string) *CreateMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateMergeRequestRequest) SetSubUserId(v string) *CreateMergeRequestRequest {
	s.SubUserId = &v
	return s
}

type CreateMergeRequestResponseBody struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *CreateMergeRequestResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateMergeRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBody) SetRequestId(v string) *CreateMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetErrorCode(v string) *CreateMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetErrorMessage(v string) *CreateMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetSuccess(v bool) *CreateMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetResult(v *CreateMergeRequestResponseBodyResult) *CreateMergeRequestResponseBody {
	s.Result = v
	return s
}

type CreateMergeRequestResponseBodyResult struct {
	State              *string                                                 `json:"State,omitempty" xml:"State,omitempty"`
	BehindCommitCount  *int32                                                  `json:"BehindCommitCount,omitempty" xml:"BehindCommitCount,omitempty"`
	ProjectId          *int64                                                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AssigneeList       []*CreateMergeRequestResponseBodyResultAssigneeList     `json:"AssigneeList,omitempty" xml:"AssigneeList,omitempty" type:"Repeated"`
	CreatedAt          *string                                                 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Author             *CreateMergeRequestResponseBodyResultAuthor             `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	AcceptedRevision   *string                                                 `json:"AcceptedRevision,omitempty" xml:"AcceptedRevision,omitempty"`
	ApproveCheckResult *CreateMergeRequestResponseBodyResultApproveCheckResult `json:"ApproveCheckResult,omitempty" xml:"ApproveCheckResult,omitempty" type:"Struct"`
	SourceBranch       *string                                                 `json:"SourceBranch,omitempty" xml:"SourceBranch,omitempty"`
	WebUrl             *string                                                 `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	Description        *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	MergeType          *string                                                 `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	NameWithNamespace  *string                                                 `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	TargetBranch       *string                                                 `json:"TargetBranch,omitempty" xml:"TargetBranch,omitempty"`
	AheadCommitCount   *int32                                                  `json:"AheadCommitCount,omitempty" xml:"AheadCommitCount,omitempty"`
	UpdatedAt          *string                                                 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Title              *string                                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	MergeError         *string                                                 `json:"MergeError,omitempty" xml:"MergeError,omitempty"`
	MergedRevision     *string                                                 `json:"MergedRevision,omitempty" xml:"MergedRevision,omitempty"`
	Id                 *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	MergeStatus        *string                                                 `json:"MergeStatus,omitempty" xml:"MergeStatus,omitempty"`
}

func (s CreateMergeRequestResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResult) SetState(v string) *CreateMergeRequestResponseBodyResult {
	s.State = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetBehindCommitCount(v int32) *CreateMergeRequestResponseBodyResult {
	s.BehindCommitCount = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetProjectId(v int64) *CreateMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAssigneeList(v []*CreateMergeRequestResponseBodyResultAssigneeList) *CreateMergeRequestResponseBodyResult {
	s.AssigneeList = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetCreatedAt(v string) *CreateMergeRequestResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAuthor(v *CreateMergeRequestResponseBodyResultAuthor) *CreateMergeRequestResponseBodyResult {
	s.Author = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAcceptedRevision(v string) *CreateMergeRequestResponseBodyResult {
	s.AcceptedRevision = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetApproveCheckResult(v *CreateMergeRequestResponseBodyResultApproveCheckResult) *CreateMergeRequestResponseBodyResult {
	s.ApproveCheckResult = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetSourceBranch(v string) *CreateMergeRequestResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetWebUrl(v string) *CreateMergeRequestResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetDescription(v string) *CreateMergeRequestResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMergeType(v string) *CreateMergeRequestResponseBodyResult {
	s.MergeType = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetNameWithNamespace(v string) *CreateMergeRequestResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetTargetBranch(v string) *CreateMergeRequestResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAheadCommitCount(v int32) *CreateMergeRequestResponseBodyResult {
	s.AheadCommitCount = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetUpdatedAt(v string) *CreateMergeRequestResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetTitle(v string) *CreateMergeRequestResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMergeError(v string) *CreateMergeRequestResponseBodyResult {
	s.MergeError = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMergedRevision(v string) *CreateMergeRequestResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetId(v int64) *CreateMergeRequestResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMergeStatus(v string) *CreateMergeRequestResponseBodyResult {
	s.MergeStatus = &v
	return s
}

type CreateMergeRequestResponseBodyResultAssigneeList struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultAssigneeList) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultAssigneeList) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultAssigneeList) SetExternUserId(v string) *CreateMergeRequestResponseBodyResultAssigneeList {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAssigneeList) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultAssigneeList {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAssigneeList) SetName(v string) *CreateMergeRequestResponseBodyResultAssigneeList {
	s.Name = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAssigneeList) SetId(v string) *CreateMergeRequestResponseBodyResultAssigneeList {
	s.Id = &v
	return s
}

type CreateMergeRequestResponseBodyResultAuthor struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetExternUserId(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetName(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetId(v int64) *CreateMergeRequestResponseBodyResultAuthor {
	s.Id = &v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResult struct {
	TotalCheckResult        *string                                                                          `json:"TotalCheckResult,omitempty" xml:"TotalCheckResult,omitempty"`
	UnsatisfiedCheckResults []*CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults `json:"UnsatisfiedCheckResults,omitempty" xml:"UnsatisfiedCheckResults,omitempty" type:"Repeated"`
	SatisfiedCheckResults   []*CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults   `json:"SatisfiedCheckResults,omitempty" xml:"SatisfiedCheckResults,omitempty" type:"Repeated"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResult) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResult) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResult) SetTotalCheckResult(v string) *CreateMergeRequestResponseBodyResultApproveCheckResult {
	s.TotalCheckResult = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResult) SetUnsatisfiedCheckResults(v []*CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) *CreateMergeRequestResponseBodyResultApproveCheckResult {
	s.UnsatisfiedCheckResults = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResult) SetSatisfiedCheckResults(v []*CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) *CreateMergeRequestResponseBodyResultApproveCheckResult {
	s.SatisfiedCheckResults = v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults struct {
	SatisfiedItems   []*string                                                                                  `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	CheckStatus      *string                                                                                    `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                    `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	UnsatisfiedItems []*string                                                                                  `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
	ExtraUsers       []*CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	CheckName        *string                                                                                    `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetSatisfiedItems(v []*string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckStatus(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckType(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetExtraUsers(v []*CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckName(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckName = &v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetName(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetId(v int64) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults struct {
	SatisfiedItems   []*string                                                                                `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	CheckStatus      *string                                                                                  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	UnsatisfiedItems []*string                                                                                `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
	ExtraUsers       []*CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	CheckName        *string                                                                                  `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetSatisfiedItems(v []*string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckStatus(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckType(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetExtraUsers(v []*CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckName(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckName = &v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetName(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetId(v int64) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

type CreateMergeRequestResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMergeRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponse) SetHeaders(v map[string]*string) *CreateMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateMergeRequestResponse) SetBody(v *CreateMergeRequestResponseBody) *CreateMergeRequestResponse {
	s.Body = v
	return s
}

type CreateRepositoryRequest struct {
	AccessToken      *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Sync             *bool   `json:"Sync,omitempty" xml:"Sync,omitempty"`
	CreateParentPath *bool   `json:"CreateParentPath,omitempty" xml:"CreateParentPath,omitempty"`
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId        *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryRequest) SetAccessToken(v string) *CreateRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateRepositoryRequest) SetSync(v bool) *CreateRepositoryRequest {
	s.Sync = &v
	return s
}

func (s *CreateRepositoryRequest) SetCreateParentPath(v bool) *CreateRepositoryRequest {
	s.CreateParentPath = &v
	return s
}

func (s *CreateRepositoryRequest) SetOrganizationId(v string) *CreateRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateRepositoryRequest) SetSubUserId(v string) *CreateRepositoryRequest {
	s.SubUserId = &v
	return s
}

func (s *CreateRepositoryRequest) SetClientToken(v string) *CreateRepositoryRequest {
	s.ClientToken = &v
	return s
}

type CreateRepositoryResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *int32                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *CreateRepositoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBody) SetRequestId(v string) *CreateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetErrorCode(v int32) *CreateRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetErrorMessage(v string) *CreateRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetSuccess(v bool) *CreateRepositoryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetResult(v *CreateRepositoryResponseBodyResult) *CreateRepositoryResponseBody {
	s.Result = v
	return s
}

type CreateRepositoryResponseBodyResult struct {
	DefaultBranch            *string                                      `json:"DefaultBranch,omitempty" xml:"DefaultBranch,omitempty"`
	SnippetsEnableStatus     *bool                                        `json:"SnippetsEnableStatus,omitempty" xml:"SnippetsEnableStatus,omitempty"`
	CreatedAt                *string                                      `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	BuildsEnableStatus       *bool                                        `json:"BuildsEnableStatus,omitempty" xml:"BuildsEnableStatus,omitempty"`
	WebUrl                   *string                                      `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	Description              *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	TagList                  []*string                                    `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	Public                   *bool                                        `json:"Public,omitempty" xml:"Public,omitempty"`
	PathWithNamespace        *string                                      `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	VisibilityLevel          *string                                      `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WikiEnableStatus         *bool                                        `json:"WikiEnableStatus,omitempty" xml:"WikiEnableStatus,omitempty"`
	Name                     *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	LastActivityAt           *string                                      `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	AvatarUrl                *string                                      `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Archive                  *bool                                        `json:"Archive,omitempty" xml:"Archive,omitempty"`
	Namespace                *CreateRepositoryResponseBodyResultNamespace `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Struct"`
	IssuesEnableStatus       *bool                                        `json:"IssuesEnableStatus,omitempty" xml:"IssuesEnableStatus,omitempty"`
	DemoProjectStatus        *bool                                        `json:"DemoProjectStatus,omitempty" xml:"DemoProjectStatus,omitempty"`
	CreatorId                *int64                                       `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	HttpUrlToRepo            *string                                      `json:"HttpUrlToRepo,omitempty" xml:"HttpUrlToRepo,omitempty"`
	NameWithNamespace        *string                                      `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	MergeRequestEnableStatus *bool                                        `json:"MergeRequestEnableStatus,omitempty" xml:"MergeRequestEnableStatus,omitempty"`
	Path                     *string                                      `json:"Path,omitempty" xml:"Path,omitempty"`
	SshUrlToRepo             *string                                      `json:"SshUrlToRepo,omitempty" xml:"SshUrlToRepo,omitempty"`
	Id                       *int64                                       `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateRepositoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBodyResult) SetDefaultBranch(v string) *CreateRepositoryResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetSnippetsEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.SnippetsEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetCreatedAt(v string) *CreateRepositoryResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetBuildsEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.BuildsEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetWebUrl(v string) *CreateRepositoryResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDescription(v string) *CreateRepositoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetTagList(v []*string) *CreateRepositoryResponseBodyResult {
	s.TagList = v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPublic(v bool) *CreateRepositoryResponseBodyResult {
	s.Public = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPathWithNamespace(v string) *CreateRepositoryResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetVisibilityLevel(v string) *CreateRepositoryResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetWikiEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.WikiEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetName(v string) *CreateRepositoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetLastActivityAt(v string) *CreateRepositoryResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetAvatarUrl(v string) *CreateRepositoryResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetArchive(v bool) *CreateRepositoryResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetNamespace(v *CreateRepositoryResponseBodyResultNamespace) *CreateRepositoryResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetIssuesEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.IssuesEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDemoProjectStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.DemoProjectStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetCreatorId(v int64) *CreateRepositoryResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetHttpUrlToRepo(v string) *CreateRepositoryResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetNameWithNamespace(v string) *CreateRepositoryResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetMergeRequestEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.MergeRequestEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPath(v string) *CreateRepositoryResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetSshUrlToRepo(v string) *CreateRepositoryResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetId(v int64) *CreateRepositoryResponseBodyResult {
	s.Id = &v
	return s
}

type CreateRepositoryResponseBodyResultNamespace struct {
	Avatar          *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Public          *bool   `json:"Public,omitempty" xml:"Public,omitempty"`
	VisibilityLevel *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	CreatedAt       *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdatedAt       *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateRepositoryResponseBodyResultNamespace) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBodyResultNamespace) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetAvatar(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Avatar = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetDescription(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetState(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.State = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetPublic(v bool) *CreateRepositoryResponseBodyResultNamespace {
	s.Public = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetVisibilityLevel(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetCreatedAt(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetPath(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetUpdatedAt(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetName(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Name = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetOwnerId(v int64) *CreateRepositoryResponseBodyResultNamespace {
	s.OwnerId = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetId(v int64) *CreateRepositoryResponseBodyResultNamespace {
	s.Id = &v
	return s
}

type CreateRepositoryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponse) SetHeaders(v map[string]*string) *CreateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryResponse) SetBody(v *CreateRepositoryResponseBody) *CreateRepositoryResponse {
	s.Body = v
	return s
}

type CreateRepositoryGroupRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateRepositoryGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupRequest) SetAccessToken(v string) *CreateRepositoryGroupRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetOrganizationId(v string) *CreateRepositoryGroupRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetSubUserId(v string) *CreateRepositoryGroupRequest {
	s.SubUserId = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetClientToken(v string) *CreateRepositoryGroupRequest {
	s.ClientToken = &v
	return s
}

type CreateRepositoryGroupResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *int32                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *CreateRepositoryGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateRepositoryGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupResponseBody) SetRequestId(v string) *CreateRepositoryGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetErrorCode(v int32) *CreateRepositoryGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetErrorMessage(v string) *CreateRepositoryGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetSuccess(v bool) *CreateRepositoryGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetResult(v *CreateRepositoryGroupResponseBodyResult) *CreateRepositoryGroupResponseBody {
	s.Result = v
	return s
}

type CreateRepositoryGroupResponseBodyResult struct {
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	AvatarUrl         *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentId          *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateRepositoryGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupResponseBodyResult) SetType(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetAvatarUrl(v string) *CreateRepositoryGroupResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetOwnerId(v int64) *CreateRepositoryGroupResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetParentId(v int64) *CreateRepositoryGroupResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetWebUrl(v string) *CreateRepositoryGroupResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetDescription(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetNameWithNamespace(v string) *CreateRepositoryGroupResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetPathWithNamespace(v string) *CreateRepositoryGroupResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetPath(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetVisibilityLevel(v string) *CreateRepositoryGroupResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetName(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetId(v int64) *CreateRepositoryGroupResponseBodyResult {
	s.Id = &v
	return s
}

type CreateRepositoryGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRepositoryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRepositoryGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupResponse) SetHeaders(v map[string]*string) *CreateRepositoryGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryGroupResponse) SetBody(v *CreateRepositoryGroupResponseBody) *CreateRepositoryGroupResponse {
	s.Body = v
	return s
}

type CreateTagRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s CreateTagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTagRequest) GoString() string {
	return s.String()
}

func (s *CreateTagRequest) SetAccessToken(v string) *CreateTagRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateTagRequest) SetOrganizationId(v string) *CreateTagRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateTagRequest) SetSubUserId(v string) *CreateTagRequest {
	s.SubUserId = &v
	return s
}

type CreateTagResponseBody struct {
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *CreateTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) SetErrorCode(v string) *CreateTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTagResponseBody) SetErrorMessage(v string) *CreateTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTagResponseBody) SetSuccess(v bool) *CreateTagResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTagResponseBody) SetResult(v *CreateTagResponseBodyResult) *CreateTagResponseBody {
	s.Result = v
	return s
}

type CreateTagResponseBodyResult struct {
	Release    *CreateTagResponseBodyResultRelease    `json:"Release,omitempty" xml:"Release,omitempty" type:"Struct"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	CommitInfo *CreateTagResponseBodyResultCommitInfo `json:"CommitInfo,omitempty" xml:"CommitInfo,omitempty" type:"Struct"`
	Name       *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyResult) SetRelease(v *CreateTagResponseBodyResultRelease) *CreateTagResponseBodyResult {
	s.Release = v
	return s
}

func (s *CreateTagResponseBodyResult) SetMessage(v string) *CreateTagResponseBodyResult {
	s.Message = &v
	return s
}

func (s *CreateTagResponseBodyResult) SetCommitInfo(v *CreateTagResponseBodyResultCommitInfo) *CreateTagResponseBodyResult {
	s.CommitInfo = v
	return s
}

func (s *CreateTagResponseBodyResult) SetName(v string) *CreateTagResponseBodyResult {
	s.Name = &v
	return s
}

type CreateTagResponseBodyResultRelease struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TagName     *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s CreateTagResponseBodyResultRelease) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBodyResultRelease) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyResultRelease) SetDescription(v string) *CreateTagResponseBodyResultRelease {
	s.Description = &v
	return s
}

func (s *CreateTagResponseBodyResultRelease) SetTagName(v string) *CreateTagResponseBodyResultRelease {
	s.TagName = &v
	return s
}

type CreateTagResponseBodyResultCommitInfo struct {
	ShortId        *string   `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	AuthoredDate   *string   `json:"AuthoredDate,omitempty" xml:"AuthoredDate,omitempty"`
	ParentIds      []*string `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	AuthorName     *string   `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommitterName  *string   `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	AuthorEmail    *string   `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	CommitterEmail *string   `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	CommittedDate  *string   `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
}

func (s CreateTagResponseBodyResultCommitInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBodyResultCommitInfo) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyResultCommitInfo) SetShortId(v string) *CreateTagResponseBodyResultCommitInfo {
	s.ShortId = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetCreatedAt(v string) *CreateTagResponseBodyResultCommitInfo {
	s.CreatedAt = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetMessage(v string) *CreateTagResponseBodyResultCommitInfo {
	s.Message = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetAuthoredDate(v string) *CreateTagResponseBodyResultCommitInfo {
	s.AuthoredDate = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetParentIds(v []*string) *CreateTagResponseBodyResultCommitInfo {
	s.ParentIds = v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetAuthorName(v string) *CreateTagResponseBodyResultCommitInfo {
	s.AuthorName = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetCommitterName(v string) *CreateTagResponseBodyResultCommitInfo {
	s.CommitterName = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetTitle(v string) *CreateTagResponseBodyResultCommitInfo {
	s.Title = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetAuthorEmail(v string) *CreateTagResponseBodyResultCommitInfo {
	s.AuthorEmail = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetCommitterEmail(v string) *CreateTagResponseBodyResultCommitInfo {
	s.CommitterEmail = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetId(v string) *CreateTagResponseBodyResultCommitInfo {
	s.Id = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetCommittedDate(v string) *CreateTagResponseBodyResultCommitInfo {
	s.CommittedDate = &v
	return s
}

type CreateTagResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponse) GoString() string {
	return s.String()
}

func (s *CreateTagResponse) SetHeaders(v map[string]*string) *CreateTagResponse {
	s.Headers = v
	return s
}

func (s *CreateTagResponse) SetBody(v *CreateTagResponseBody) *CreateTagResponse {
	s.Body = v
	return s
}

type DeleteBranchRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	BranchName     *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DeleteBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchRequest) GoString() string {
	return s.String()
}

func (s *DeleteBranchRequest) SetAccessToken(v string) *DeleteBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteBranchRequest) SetBranchName(v string) *DeleteBranchRequest {
	s.BranchName = &v
	return s
}

func (s *DeleteBranchRequest) SetOrganizationId(v string) *DeleteBranchRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteBranchRequest) SetSubUserId(v string) *DeleteBranchRequest {
	s.SubUserId = &v
	return s
}

type DeleteBranchResponseBody struct {
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *DeleteBranchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponseBody) SetRequestId(v string) *DeleteBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBranchResponseBody) SetErrorCode(v string) *DeleteBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteBranchResponseBody) SetErrorMessage(v string) *DeleteBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteBranchResponseBody) SetSuccess(v bool) *DeleteBranchResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBranchResponseBody) SetResult(v *DeleteBranchResponseBodyResult) *DeleteBranchResponseBody {
	s.Result = v
	return s
}

type DeleteBranchResponseBodyResult struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
}

func (s DeleteBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponseBodyResult) SetBranchName(v string) *DeleteBranchResponseBodyResult {
	s.BranchName = &v
	return s
}

type DeleteBranchResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchResponse) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponse) SetHeaders(v map[string]*string) *DeleteBranchResponse {
	s.Headers = v
	return s
}

func (s *DeleteBranchResponse) SetBody(v *DeleteBranchResponseBody) *DeleteBranchResponse {
	s.Body = v
	return s
}

type DeleteFileRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	BranchName     *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	FilePath       *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	CommitMessage  *string `json:"CommitMessage,omitempty" xml:"CommitMessage,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetAccessToken(v string) *DeleteFileRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteFileRequest) SetOrganizationId(v string) *DeleteFileRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteFileRequest) SetBranchName(v string) *DeleteFileRequest {
	s.BranchName = &v
	return s
}

func (s *DeleteFileRequest) SetFilePath(v string) *DeleteFileRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteFileRequest) SetCommitMessage(v string) *DeleteFileRequest {
	s.CommitMessage = &v
	return s
}

func (s *DeleteFileRequest) SetSubUserId(v string) *DeleteFileRequest {
	s.SubUserId = &v
	return s
}

type DeleteFileResponseBody struct {
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *DeleteFileResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrorCode(v string) *DeleteFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrorMessage(v string) *DeleteFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFileResponseBody) SetResult(v *DeleteFileResponseBodyResult) *DeleteFileResponseBody {
	s.Result = v
	return s
}

type DeleteFileResponseBodyResult struct {
	FilePath   *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
}

func (s DeleteFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBodyResult) SetFilePath(v string) *DeleteFileResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *DeleteFileResponseBodyResult) SetBranchName(v string) *DeleteFileResponseBodyResult {
	s.BranchName = &v
	return s
}

type DeleteFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type DeleteGroupMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DeleteGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberRequest) SetAccessToken(v string) *DeleteGroupMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteGroupMemberRequest) SetOrganizationId(v string) *DeleteGroupMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteGroupMemberRequest) SetSubUserId(v string) *DeleteGroupMemberRequest {
	s.SubUserId = &v
	return s
}

type DeleteGroupMemberResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *DeleteGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberResponseBody) SetRequestId(v string) *DeleteGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetErrorCode(v string) *DeleteGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetErrorMessage(v string) *DeleteGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetSuccess(v bool) *DeleteGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetResult(v *DeleteGroupMemberResponseBodyResult) *DeleteGroupMemberResponseBody {
	s.Result = v
	return s
}

type DeleteGroupMemberResponseBodyResult struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberResponseBodyResult) SetExternUserId(v string) *DeleteGroupMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetEmail(v string) *DeleteGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetAvatarUrl(v string) *DeleteGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetState(v string) *DeleteGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetAccessLevel(v int32) *DeleteGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetId(v int64) *DeleteGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

type DeleteGroupMemberResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberResponse) SetHeaders(v map[string]*string) *DeleteGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupMemberResponse) SetBody(v *DeleteGroupMemberResponseBody) *DeleteGroupMemberResponse {
	s.Body = v
	return s
}

type DeleteRepositoryRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DeleteRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryRequest) SetAccessToken(v string) *DeleteRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryRequest) SetOrganizationId(v string) *DeleteRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteRepositoryRequest) SetSubUserId(v string) *DeleteRepositoryRequest {
	s.SubUserId = &v
	return s
}

type DeleteRepositoryResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *DeleteRepositoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBody) SetRequestId(v string) *DeleteRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetErrorCode(v string) *DeleteRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetErrorMessage(v string) *DeleteRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetSuccess(v bool) *DeleteRepositoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetResult(v *DeleteRepositoryResponseBodyResult) *DeleteRepositoryResponseBody {
	s.Result = v
	return s
}

type DeleteRepositoryResponseBodyResult struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteRepositoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBodyResult) SetResult(v bool) *DeleteRepositoryResponseBodyResult {
	s.Result = &v
	return s
}

type DeleteRepositoryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponse) SetHeaders(v map[string]*string) *DeleteRepositoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryResponse) SetBody(v *DeleteRepositoryResponseBody) *DeleteRepositoryResponse {
	s.Body = v
	return s
}

type DeleteRepositoryGroupRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DeleteRepositoryGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupRequest) SetAccessToken(v string) *DeleteRepositoryGroupRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryGroupRequest) SetOrganizationId(v string) *DeleteRepositoryGroupRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteRepositoryGroupRequest) SetSubUserId(v string) *DeleteRepositoryGroupRequest {
	s.SubUserId = &v
	return s
}

type DeleteRepositoryGroupResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *DeleteRepositoryGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteRepositoryGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupResponseBody) SetRequestId(v string) *DeleteRepositoryGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetErrorCode(v string) *DeleteRepositoryGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetErrorMessage(v string) *DeleteRepositoryGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetSuccess(v bool) *DeleteRepositoryGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetResult(v *DeleteRepositoryGroupResponseBodyResult) *DeleteRepositoryGroupResponseBody {
	s.Result = v
	return s
}

type DeleteRepositoryGroupResponseBodyResult struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteRepositoryGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupResponseBodyResult) SetResult(v bool) *DeleteRepositoryGroupResponseBodyResult {
	s.Result = &v
	return s
}

type DeleteRepositoryGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRepositoryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupResponse) SetHeaders(v map[string]*string) *DeleteRepositoryGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryGroupResponse) SetBody(v *DeleteRepositoryGroupResponseBody) *DeleteRepositoryGroupResponse {
	s.Body = v
	return s
}

type DeleteRepositoryMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DeleteRepositoryMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberRequest) SetAccessToken(v string) *DeleteRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryMemberRequest) SetOrganizationId(v string) *DeleteRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteRepositoryMemberRequest) SetSubUserId(v string) *DeleteRepositoryMemberRequest {
	s.SubUserId = &v
	return s
}

type DeleteRepositoryMemberResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *DeleteRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberResponseBody) SetRequestId(v string) *DeleteRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetErrorCode(v string) *DeleteRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetErrorMessage(v string) *DeleteRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetSuccess(v bool) *DeleteRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetResult(v *DeleteRepositoryMemberResponseBodyResult) *DeleteRepositoryMemberResponseBody {
	s.Result = v
	return s
}

type DeleteRepositoryMemberResponseBodyResult struct {
	UserId            *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SourceType        *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	AccessLevel       *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	UpdatedAt         *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	SourceId          *int64  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	NotificationLevel *int32  `json:"NotificationLevel,omitempty" xml:"NotificationLevel,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetUserId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetSourceType(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetCreatedAt(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetMessage(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *DeleteRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetUpdatedAt(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetSourceId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetNotificationLevel(v int32) *DeleteRepositoryMemberResponseBodyResult {
	s.NotificationLevel = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

type DeleteRepositoryMemberResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberResponse) SetHeaders(v map[string]*string) *DeleteRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryMemberResponse) SetBody(v *DeleteRepositoryMemberResponseBody) *DeleteRepositoryMemberResponse {
	s.Body = v
	return s
}

type DeleteRepositoryTagRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteRepositoryTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagRequest) SetAccessToken(v string) *DeleteRepositoryTagRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryTagRequest) SetOrganizationId(v string) *DeleteRepositoryTagRequest {
	s.OrganizationId = &v
	return s
}

type DeleteRepositoryTagResponseBody struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *DeleteRepositoryTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteRepositoryTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagResponseBody) SetRequestId(v string) *DeleteRepositoryTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryTagResponseBody) SetErrorCode(v string) *DeleteRepositoryTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryTagResponseBody) SetErrorMessage(v string) *DeleteRepositoryTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryTagResponseBody) SetSuccess(v bool) *DeleteRepositoryTagResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryTagResponseBody) SetResult(v *DeleteRepositoryTagResponseBodyResult) *DeleteRepositoryTagResponseBody {
	s.Result = v
	return s
}

type DeleteRepositoryTagResponseBodyResult struct {
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DeleteRepositoryTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagResponseBodyResult) SetTagName(v string) *DeleteRepositoryTagResponseBodyResult {
	s.TagName = &v
	return s
}

type DeleteRepositoryTagResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRepositoryTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagResponse) SetHeaders(v map[string]*string) *DeleteRepositoryTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryTagResponse) SetBody(v *DeleteRepositoryTagResponseBody) *DeleteRepositoryTagResponse {
	s.Body = v
	return s
}

type DeleteRepositoryWebhookRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteRepositoryWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryWebhookRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookRequest) SetAccessToken(v string) *DeleteRepositoryWebhookRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryWebhookRequest) SetOrganizationId(v string) *DeleteRepositoryWebhookRequest {
	s.OrganizationId = &v
	return s
}

type DeleteRepositoryWebhookResponseBody struct {
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *DeleteRepositoryWebhookResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteRepositoryWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookResponseBody) SetRequestId(v string) *DeleteRepositoryWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetErrorCode(v string) *DeleteRepositoryWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetErrorMessage(v string) *DeleteRepositoryWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetSuccess(v bool) *DeleteRepositoryWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetResult(v *DeleteRepositoryWebhookResponseBodyResult) *DeleteRepositoryWebhookResponseBody {
	s.Result = v
	return s
}

type DeleteRepositoryWebhookResponseBodyResult struct {
	PushEvents            *bool   `json:"PushEvents,omitempty" xml:"PushEvents,omitempty"`
	ProjectId             *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	CreatedAt             *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	TagPushEvents         *bool   `json:"TagPushEvents,omitempty" xml:"TagPushEvents,omitempty"`
	Url                   *string `json:"Url,omitempty" xml:"Url,omitempty"`
	LastTestResult        *string `json:"LastTestResult,omitempty" xml:"LastTestResult,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MergeRequestsEvents   *bool   `json:"MergeRequestsEvents,omitempty" xml:"MergeRequestsEvents,omitempty"`
	SecretToken           *string `json:"SecretToken,omitempty" xml:"SecretToken,omitempty"`
	NoteEvents            *bool   `json:"NoteEvents,omitempty" xml:"NoteEvents,omitempty"`
	EnableSslVerification *bool   `json:"EnableSslVerification,omitempty" xml:"EnableSslVerification,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRepositoryWebhookResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetPushEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetProjectId(v int64) *DeleteRepositoryWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetCreatedAt(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetTagPushEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetUrl(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.Url = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetLastTestResult(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetDescription(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetSecretToken(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetNoteEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetEnableSslVerification(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetId(v int64) *DeleteRepositoryWebhookResponseBodyResult {
	s.Id = &v
	return s
}

type DeleteRepositoryWebhookResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRepositoryWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryWebhookResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookResponse) SetHeaders(v map[string]*string) *DeleteRepositoryWebhookResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryWebhookResponse) SetBody(v *DeleteRepositoryWebhookResponseBody) *DeleteRepositoryWebhookResponse {
	s.Body = v
	return s
}

type GetBranchInfoRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	BranchName     *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
}

func (s GetBranchInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBranchInfoRequest) SetAccessToken(v string) *GetBranchInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *GetBranchInfoRequest) SetOrganizationId(v string) *GetBranchInfoRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetBranchInfoRequest) SetSubUserId(v string) *GetBranchInfoRequest {
	s.SubUserId = &v
	return s
}

func (s *GetBranchInfoRequest) SetBranchName(v string) *GetBranchInfoRequest {
	s.BranchName = &v
	return s
}

type GetBranchInfoResponseBody struct {
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *GetBranchInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetBranchInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBody) SetRequestId(v string) *GetBranchInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetErrorCode(v string) *GetBranchInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetErrorMessage(v string) *GetBranchInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetSuccess(v bool) *GetBranchInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetResult(v *GetBranchInfoResponseBodyResult) *GetBranchInfoResponseBody {
	s.Result = v
	return s
}

type GetBranchInfoResponseBodyResult struct {
	ProtectedBranch *bool                                      `json:"ProtectedBranch,omitempty" xml:"ProtectedBranch,omitempty"`
	CommitInfo      *GetBranchInfoResponseBodyResultCommitInfo `json:"CommitInfo,omitempty" xml:"CommitInfo,omitempty" type:"Struct"`
	BranchName      *string                                    `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
}

func (s GetBranchInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResult) SetProtectedBranch(v bool) *GetBranchInfoResponseBodyResult {
	s.ProtectedBranch = &v
	return s
}

func (s *GetBranchInfoResponseBodyResult) SetCommitInfo(v *GetBranchInfoResponseBodyResultCommitInfo) *GetBranchInfoResponseBodyResult {
	s.CommitInfo = v
	return s
}

func (s *GetBranchInfoResponseBodyResult) SetBranchName(v string) *GetBranchInfoResponseBodyResult {
	s.BranchName = &v
	return s
}

type GetBranchInfoResponseBodyResultCommitInfo struct {
	ShortId        *string   `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	AuthorDate     *string   `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	AuthorName     *string   `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommitterName  *string   `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	AuthorEmail    *string   `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	CommitterEmail *string   `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	CommittedDate  *string   `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommitInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommitInfo) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetShortId(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.ShortId = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetAuthorDate(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.AuthorDate = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetCreatedAt(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.CreatedAt = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetMessage(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.Message = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetParentIds(v []*string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.ParentIds = v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetAuthorName(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.AuthorName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetCommitterName(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.CommitterName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetTitle(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.Title = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetAuthorEmail(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.AuthorEmail = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetCommitterEmail(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.CommitterEmail = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetId(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.Id = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetCommittedDate(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.CommittedDate = &v
	return s
}

type GetBranchInfoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBranchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBranchInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponse) SetHeaders(v map[string]*string) *GetBranchInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBranchInfoResponse) SetBody(v *GetBranchInfoResponseBody) *GetBranchInfoResponse {
	s.Body = v
	return s
}

type GetCodeupOrganizationRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s GetCodeupOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationRequest) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationRequest) SetAccessToken(v string) *GetCodeupOrganizationRequest {
	s.AccessToken = &v
	return s
}

func (s *GetCodeupOrganizationRequest) SetOrganizationId(v string) *GetCodeupOrganizationRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetCodeupOrganizationRequest) SetSubUserId(v string) *GetCodeupOrganizationRequest {
	s.SubUserId = &v
	return s
}

type GetCodeupOrganizationResponseBody struct {
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *GetCodeupOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetCodeupOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponseBody) SetRequestId(v string) *GetCodeupOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetErrorCode(v string) *GetCodeupOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetErrorMessage(v string) *GetCodeupOrganizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetSuccess(v bool) *GetCodeupOrganizationResponseBody {
	s.Success = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetResult(v *GetCodeupOrganizationResponseBodyResult) *GetCodeupOrganizationResponseBody {
	s.Result = v
	return s
}

type GetCodeupOrganizationResponseBodyResult struct {
	NamespaceId    *int64  `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	UserRole       *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	CreatedAt      *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt      *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCodeupOrganizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponseBodyResult) SetNamespaceId(v int64) *GetCodeupOrganizationResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetUserRole(v string) *GetCodeupOrganizationResponseBodyResult {
	s.UserRole = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetPath(v string) *GetCodeupOrganizationResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetCreatedAt(v string) *GetCodeupOrganizationResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetUpdatedAt(v string) *GetCodeupOrganizationResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetId(v int64) *GetCodeupOrganizationResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetOrganizationId(v string) *GetCodeupOrganizationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type GetCodeupOrganizationResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCodeupOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCodeupOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationResponse) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponse) SetHeaders(v map[string]*string) *GetCodeupOrganizationResponse {
	s.Headers = v
	return s
}

func (s *GetCodeupOrganizationResponse) SetBody(v *GetCodeupOrganizationResponseBody) *GetCodeupOrganizationResponse {
	s.Body = v
	return s
}

type GetFileBlobsRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	FilePath       *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Ref            *string `json:"Ref,omitempty" xml:"Ref,omitempty"`
	From           *int64  `json:"From,omitempty" xml:"From,omitempty"`
	To             *int64  `json:"To,omitempty" xml:"To,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s GetFileBlobsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsRequest) GoString() string {
	return s.String()
}

func (s *GetFileBlobsRequest) SetAccessToken(v string) *GetFileBlobsRequest {
	s.AccessToken = &v
	return s
}

func (s *GetFileBlobsRequest) SetOrganizationId(v string) *GetFileBlobsRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetFileBlobsRequest) SetFilePath(v string) *GetFileBlobsRequest {
	s.FilePath = &v
	return s
}

func (s *GetFileBlobsRequest) SetRef(v string) *GetFileBlobsRequest {
	s.Ref = &v
	return s
}

func (s *GetFileBlobsRequest) SetFrom(v int64) *GetFileBlobsRequest {
	s.From = &v
	return s
}

func (s *GetFileBlobsRequest) SetTo(v int64) *GetFileBlobsRequest {
	s.To = &v
	return s
}

func (s *GetFileBlobsRequest) SetSubUserId(v string) *GetFileBlobsRequest {
	s.SubUserId = &v
	return s
}

type GetFileBlobsResponseBody struct {
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *GetFileBlobsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetFileBlobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponseBody) SetRequestId(v string) *GetFileBlobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetErrorCode(v string) *GetFileBlobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetErrorMessage(v string) *GetFileBlobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetSuccess(v bool) *GetFileBlobsResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetResult(v *GetFileBlobsResponseBodyResult) *GetFileBlobsResponseBody {
	s.Result = v
	return s
}

type GetFileBlobsResponseBodyResult struct {
	TotalLines *int32  `json:"TotalLines,omitempty" xml:"TotalLines,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetFileBlobsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponseBodyResult) SetTotalLines(v int32) *GetFileBlobsResponseBodyResult {
	s.TotalLines = &v
	return s
}

func (s *GetFileBlobsResponseBodyResult) SetContent(v string) *GetFileBlobsResponseBodyResult {
	s.Content = &v
	return s
}

type GetFileBlobsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFileBlobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileBlobsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsResponse) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponse) SetHeaders(v map[string]*string) *GetFileBlobsResponse {
	s.Headers = v
	return s
}

func (s *GetFileBlobsResponse) SetBody(v *GetFileBlobsResponseBody) *GetFileBlobsResponse {
	s.Body = v
	return s
}

type GetGroupDetailRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	GroupId        *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s GetGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGroupDetailRequest) SetAccessToken(v string) *GetGroupDetailRequest {
	s.AccessToken = &v
	return s
}

func (s *GetGroupDetailRequest) SetGroupId(v int64) *GetGroupDetailRequest {
	s.GroupId = &v
	return s
}

func (s *GetGroupDetailRequest) SetOrganizationId(v string) *GetGroupDetailRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetGroupDetailRequest) SetSubUserId(v string) *GetGroupDetailRequest {
	s.SubUserId = &v
	return s
}

type GetGroupDetailResponseBody struct {
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *GetGroupDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponseBody) SetRequestId(v string) *GetGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetErrorCode(v string) *GetGroupDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetErrorMessage(v string) *GetGroupDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetSuccess(v bool) *GetGroupDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetResult(v *GetGroupDetailResponseBodyResult) *GetGroupDetailResponseBody {
	s.Result = v
	return s
}

type GetGroupDetailResponseBodyResult struct {
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	AvatarUrl         *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentId          *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetGroupDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponseBodyResult) SetType(v string) *GetGroupDetailResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetAvatarUrl(v string) *GetGroupDetailResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetOwnerId(v int64) *GetGroupDetailResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetParentId(v int64) *GetGroupDetailResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetWebUrl(v string) *GetGroupDetailResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetDescription(v string) *GetGroupDetailResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetNameWithNamespace(v string) *GetGroupDetailResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetPathWithNamespace(v string) *GetGroupDetailResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetVisibilityLevel(v string) *GetGroupDetailResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetPath(v string) *GetGroupDetailResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetName(v string) *GetGroupDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetId(v int64) *GetGroupDetailResponseBodyResult {
	s.Id = &v
	return s
}

type GetGroupDetailResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponse) SetHeaders(v map[string]*string) *GetGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGroupDetailResponse) SetBody(v *GetGroupDetailResponseBody) *GetGroupDetailResponse {
	s.Body = v
	return s
}

type GetProjectMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s GetProjectMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *GetProjectMemberRequest) SetAccessToken(v string) *GetProjectMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *GetProjectMemberRequest) SetOrganizationId(v string) *GetProjectMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetProjectMemberRequest) SetSubUserId(v string) *GetProjectMemberRequest {
	s.SubUserId = &v
	return s
}

type GetProjectMemberResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *GetProjectMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBody) SetRequestId(v string) *GetProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetErrorCode(v string) *GetProjectMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetErrorMessage(v string) *GetProjectMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetSuccess(v bool) *GetProjectMemberResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetResult(v *GetProjectMemberResponseBodyResult) *GetProjectMemberResponseBody {
	s.Result = v
	return s
}

type GetProjectMemberResponseBodyResult struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProjectMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBodyResult) SetExternUserId(v string) *GetProjectMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetAvatarUrl(v string) *GetProjectMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetAccessLevel(v int32) *GetProjectMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetName(v string) *GetProjectMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetId(v int64) *GetProjectMemberResponseBodyResult {
	s.Id = &v
	return s
}

type GetProjectMemberResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponse) SetHeaders(v map[string]*string) *GetProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *GetProjectMemberResponse) SetBody(v *GetProjectMemberResponseBody) *GetProjectMemberResponse {
	s.Body = v
	return s
}

type GetRepositoryInfoRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Identity       *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetRepositoryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoRequest) SetAccessToken(v string) *GetRepositoryInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *GetRepositoryInfoRequest) SetIdentity(v string) *GetRepositoryInfoRequest {
	s.Identity = &v
	return s
}

func (s *GetRepositoryInfoRequest) SetOrganizationId(v string) *GetRepositoryInfoRequest {
	s.OrganizationId = &v
	return s
}

type GetRepositoryInfoResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *int32                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *GetRepositoryInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRepositoryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBody) SetRequestId(v string) *GetRepositoryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryInfoResponseBody) SetErrorCode(v int32) *GetRepositoryInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryInfoResponseBody) SetErrorMessage(v string) *GetRepositoryInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryInfoResponseBody) SetSuccess(v bool) *GetRepositoryInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetRepositoryInfoResponseBody) SetResult(v *GetRepositoryInfoResponseBodyResult) *GetRepositoryInfoResponseBody {
	s.Result = v
	return s
}

type GetRepositoryInfoResponseBodyResult struct {
	DefaultBranch        *string                                         `json:"DefaultBranch,omitempty" xml:"DefaultBranch,omitempty"`
	ImportUrl            *string                                         `json:"ImportUrl,omitempty" xml:"ImportUrl,omitempty"`
	CreatedAt            *string                                         `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	WebUrl               *string                                         `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	TagList              []*string                                       `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	Description          *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Public               *bool                                           `json:"Public,omitempty" xml:"Public,omitempty"`
	PathWithNamespace    *string                                         `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	VisibilityLevel      *string                                         `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	Name                 *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	LastActivityAt       *string                                         `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	AvatarUrl            *string                                         `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Archive              *bool                                           `json:"Archive,omitempty" xml:"Archive,omitempty"`
	Namespace            *GetRepositoryInfoResponseBodyResultNamespace   `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Struct"`
	DemoProjectStatus    *bool                                           `json:"DemoProjectStatus,omitempty" xml:"DemoProjectStatus,omitempty"`
	CreatorId            *int64                                          `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	ImportStatus         *string                                         `json:"ImportStatus,omitempty" xml:"ImportStatus,omitempty"`
	HttpUrlToRepo        *string                                         `json:"HttpUrlToRepo,omitempty" xml:"HttpUrlToRepo,omitempty"`
	Permissions          *GetRepositoryInfoResponseBodyResultPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Struct"`
	NameWithNamespace    *string                                         `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	Path                 *string                                         `json:"Path,omitempty" xml:"Path,omitempty"`
	AccessLevel          *int32                                          `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	ImportFromSubversion *bool                                           `json:"ImportFromSubversion,omitempty" xml:"ImportFromSubversion,omitempty"`
	SshUrlToRepo         *string                                         `json:"SshUrlToRepo,omitempty" xml:"SshUrlToRepo,omitempty"`
	Id                   *int64                                          `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRepositoryInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBodyResult) SetDefaultBranch(v string) *GetRepositoryInfoResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetImportUrl(v string) *GetRepositoryInfoResponseBodyResult {
	s.ImportUrl = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetCreatedAt(v string) *GetRepositoryInfoResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetWebUrl(v string) *GetRepositoryInfoResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetTagList(v []*string) *GetRepositoryInfoResponseBodyResult {
	s.TagList = v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetDescription(v string) *GetRepositoryInfoResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetPublic(v bool) *GetRepositoryInfoResponseBodyResult {
	s.Public = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetPathWithNamespace(v string) *GetRepositoryInfoResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetVisibilityLevel(v string) *GetRepositoryInfoResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetName(v string) *GetRepositoryInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetLastActivityAt(v string) *GetRepositoryInfoResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetAvatarUrl(v string) *GetRepositoryInfoResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetArchive(v bool) *GetRepositoryInfoResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetNamespace(v *GetRepositoryInfoResponseBodyResultNamespace) *GetRepositoryInfoResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetDemoProjectStatus(v bool) *GetRepositoryInfoResponseBodyResult {
	s.DemoProjectStatus = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetCreatorId(v int64) *GetRepositoryInfoResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetImportStatus(v string) *GetRepositoryInfoResponseBodyResult {
	s.ImportStatus = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetHttpUrlToRepo(v string) *GetRepositoryInfoResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetPermissions(v *GetRepositoryInfoResponseBodyResultPermissions) *GetRepositoryInfoResponseBodyResult {
	s.Permissions = v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetNameWithNamespace(v string) *GetRepositoryInfoResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetPath(v string) *GetRepositoryInfoResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetAccessLevel(v int32) *GetRepositoryInfoResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetImportFromSubversion(v bool) *GetRepositoryInfoResponseBodyResult {
	s.ImportFromSubversion = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetSshUrlToRepo(v string) *GetRepositoryInfoResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetId(v int64) *GetRepositoryInfoResponseBodyResult {
	s.Id = &v
	return s
}

type GetRepositoryInfoResponseBodyResultNamespace struct {
	Avatar          *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Public          *bool   `json:"Public,omitempty" xml:"Public,omitempty"`
	VisibilityLevel *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	CreatedAt       *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdatedAt       *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRepositoryInfoResponseBodyResultNamespace) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBodyResultNamespace) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetAvatar(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Avatar = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetDescription(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetState(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.State = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetPublic(v bool) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Public = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetVisibilityLevel(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetCreatedAt(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetPath(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetUpdatedAt(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetName(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Name = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetOwnerId(v int64) *GetRepositoryInfoResponseBodyResultNamespace {
	s.OwnerId = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetId(v int64) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Id = &v
	return s
}

type GetRepositoryInfoResponseBodyResultPermissions struct {
	ProjectAccess *GetRepositoryInfoResponseBodyResultPermissionsProjectAccess `json:"ProjectAccess,omitempty" xml:"ProjectAccess,omitempty" type:"Struct"`
	GroupAccess   *GetRepositoryInfoResponseBodyResultPermissionsGroupAccess   `json:"GroupAccess,omitempty" xml:"GroupAccess,omitempty" type:"Struct"`
}

func (s GetRepositoryInfoResponseBodyResultPermissions) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBodyResultPermissions) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBodyResultPermissions) SetProjectAccess(v *GetRepositoryInfoResponseBodyResultPermissionsProjectAccess) *GetRepositoryInfoResponseBodyResultPermissions {
	s.ProjectAccess = v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultPermissions) SetGroupAccess(v *GetRepositoryInfoResponseBodyResultPermissionsGroupAccess) *GetRepositoryInfoResponseBodyResultPermissions {
	s.GroupAccess = v
	return s
}

type GetRepositoryInfoResponseBodyResultPermissionsProjectAccess struct {
	AccessLevel *int32 `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
}

func (s GetRepositoryInfoResponseBodyResultPermissionsProjectAccess) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBodyResultPermissionsProjectAccess) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBodyResultPermissionsProjectAccess) SetAccessLevel(v int32) *GetRepositoryInfoResponseBodyResultPermissionsProjectAccess {
	s.AccessLevel = &v
	return s
}

type GetRepositoryInfoResponseBodyResultPermissionsGroupAccess struct {
	AccessLevel *int32 `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
}

func (s GetRepositoryInfoResponseBodyResultPermissionsGroupAccess) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBodyResultPermissionsGroupAccess) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBodyResultPermissionsGroupAccess) SetAccessLevel(v int32) *GetRepositoryInfoResponseBodyResultPermissionsGroupAccess {
	s.AccessLevel = &v
	return s
}

type GetRepositoryInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRepositoryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRepositoryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponse) SetHeaders(v map[string]*string) *GetRepositoryInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryInfoResponse) SetBody(v *GetRepositoryInfoResponseBody) *GetRepositoryInfoResponse {
	s.Body = v
	return s
}

type GetRepositoryTagRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetRepositoryTagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagRequest) SetAccessToken(v string) *GetRepositoryTagRequest {
	s.AccessToken = &v
	return s
}

func (s *GetRepositoryTagRequest) SetOrganizationId(v string) *GetRepositoryTagRequest {
	s.OrganizationId = &v
	return s
}

type GetRepositoryTagResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *GetRepositoryTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRepositoryTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBody) SetRequestId(v string) *GetRepositoryTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetErrorCode(v string) *GetRepositoryTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetErrorMessage(v string) *GetRepositoryTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetSuccess(v bool) *GetRepositoryTagResponseBody {
	s.Success = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetResult(v *GetRepositoryTagResponseBodyResult) *GetRepositoryTagResponseBody {
	s.Result = v
	return s
}

type GetRepositoryTagResponseBodyResult struct {
	Signature *GetRepositoryTagResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Commit    *GetRepositoryTagResponseBodyResultCommit    `json:"Commit,omitempty" xml:"Commit,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Name      *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string                                      `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRepositoryTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResult) SetSignature(v *GetRepositoryTagResponseBodyResultSignature) *GetRepositoryTagResponseBodyResult {
	s.Signature = v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) SetCommit(v *GetRepositoryTagResponseBodyResultCommit) *GetRepositoryTagResponseBodyResult {
	s.Commit = v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) SetMessage(v string) *GetRepositoryTagResponseBodyResult {
	s.Message = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) SetName(v string) *GetRepositoryTagResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) SetId(v string) *GetRepositoryTagResponseBodyResult {
	s.Id = &v
	return s
}

type GetRepositoryTagResponseBodyResultSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s GetRepositoryTagResponseBodyResultSignature) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResultSignature) SetGpgKeyId(v string) *GetRepositoryTagResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultSignature) SetVerificationStatus(v string) *GetRepositoryTagResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

type GetRepositoryTagResponseBodyResultCommit struct {
	ShortId        *string                                            `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	CreatedAt      *string                                            `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	AuthoredDate   *string                                            `json:"AuthoredDate,omitempty" xml:"AuthoredDate,omitempty"`
	Signature      *GetRepositoryTagResponseBodyResultCommitSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	ParentIds      []*string                                          `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	AuthorName     *string                                            `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommitterName  *string                                            `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	Title          *string                                            `json:"Title,omitempty" xml:"Title,omitempty"`
	AuthorEmail    *string                                            `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	CommitterEmail *string                                            `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	Id             *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	CommittedDate  *string                                            `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
}

func (s GetRepositoryTagResponseBodyResultCommit) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetShortId(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCreatedAt(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetMessage(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthoredDate(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetSignature(v *GetRepositoryTagResponseBodyResultCommitSignature) *GetRepositoryTagResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetParentIds(v []*string) *GetRepositoryTagResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthorName(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommitterName(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetTitle(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Title = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthorEmail(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommitterEmail(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetId(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommittedDate(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

type GetRepositoryTagResponseBodyResultCommitSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s GetRepositoryTagResponseBodyResultCommitSignature) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResultCommitSignature) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResultCommitSignature) SetGpgKeyId(v string) *GetRepositoryTagResponseBodyResultCommitSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommitSignature) SetVerificationStatus(v string) *GetRepositoryTagResponseBodyResultCommitSignature {
	s.VerificationStatus = &v
	return s
}

type GetRepositoryTagResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRepositoryTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRepositoryTagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponse) SetHeaders(v map[string]*string) *GetRepositoryTagResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryTagResponse) SetBody(v *GetRepositoryTagResponseBody) *GetRepositoryTagResponse {
	s.Body = v
	return s
}

type GetUserInfoRequest struct {
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
}

func (s GetUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserInfoRequest) SetAccessToken(v string) *GetUserInfoRequest {
	s.AccessToken = &v
	return s
}

type GetUserInfoResponseBody struct {
	RequestId    *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *GetUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBody) SetRequestId(v string) *GetUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInfoResponseBody) SetErrorCode(v string) *GetUserInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserInfoResponseBody) SetErrorMessage(v string) *GetUserInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserInfoResponseBody) SetSuccess(v bool) *GetUserInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserInfoResponseBody) SetResult(v *GetUserInfoResponseBodyResult) *GetUserInfoResponseBody {
	s.Result = v
	return s
}

type GetUserInfoResponseBodyResult struct {
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl      *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBodyResult) SetEmail(v string) *GetUserInfoResponseBodyResult {
	s.Email = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetAvatarUrl(v string) *GetUserInfoResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetExternalUserId(v string) *GetUserInfoResponseBodyResult {
	s.ExternalUserId = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetName(v string) *GetUserInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetId(v int64) *GetUserInfoResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetUsername(v string) *GetUserInfoResponseBodyResult {
	s.Username = &v
	return s
}

type GetUserInfoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponse) SetHeaders(v map[string]*string) *GetUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserInfoResponse) SetBody(v *GetUserInfoResponseBody) *GetUserInfoResponse {
	s.Body = v
	return s
}

type ListGroupMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *ListGroupMemberRequest) SetAccessToken(v string) *ListGroupMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *ListGroupMemberRequest) SetOrganizationId(v string) *ListGroupMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListGroupMemberRequest) SetPage(v int64) *ListGroupMemberRequest {
	s.Page = &v
	return s
}

func (s *ListGroupMemberRequest) SetPageSize(v int64) *ListGroupMemberRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupMemberRequest) SetSubUserId(v string) *ListGroupMemberRequest {
	s.SubUserId = &v
	return s
}

type ListGroupMemberResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                               `json:"Total,omitempty" xml:"Total,omitempty"`
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponseBody) SetRequestId(v string) *ListGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetTotal(v int64) *ListGroupMemberResponseBody {
	s.Total = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetErrorCode(v string) *ListGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetErrorMessage(v string) *ListGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetSuccess(v bool) *ListGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetResult(v []*ListGroupMemberResponseBodyResult) *ListGroupMemberResponseBody {
	s.Result = v
	return s
}

type ListGroupMemberResponseBodyResult struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponseBodyResult) SetExternUserId(v string) *ListGroupMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetEmail(v string) *ListGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetAvatarUrl(v string) *ListGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetState(v string) *ListGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetAccessLevel(v int32) *ListGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetId(v int64) *ListGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

type ListGroupMemberResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponse) SetHeaders(v map[string]*string) *ListGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *ListGroupMemberResponse) SetBody(v *ListGroupMemberResponseBody) *ListGroupMemberResponse {
	s.Body = v
	return s
}

type ListGroupRepositoriesRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	IsMember       *bool   `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGroupRepositoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesRequest) SetAccessToken(v string) *ListGroupRepositoriesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetOrganizationId(v string) *ListGroupRepositoriesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetIsMember(v bool) *ListGroupRepositoriesRequest {
	s.IsMember = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetSubUserId(v string) *ListGroupRepositoriesRequest {
	s.SubUserId = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetSearch(v string) *ListGroupRepositoriesRequest {
	s.Search = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetPage(v int64) *ListGroupRepositoriesRequest {
	s.Page = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetPageSize(v int64) *ListGroupRepositoriesRequest {
	s.PageSize = &v
	return s
}

type ListGroupRepositoriesResponseBody struct {
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                     `json:"Total,omitempty" xml:"Total,omitempty"`
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListGroupRepositoriesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListGroupRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesResponseBody) SetRequestId(v string) *ListGroupRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetTotal(v int64) *ListGroupRepositoriesResponseBody {
	s.Total = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetErrorCode(v string) *ListGroupRepositoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetErrorMessage(v string) *ListGroupRepositoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetSuccess(v bool) *ListGroupRepositoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetResult(v []*ListGroupRepositoriesResponseBodyResult) *ListGroupRepositoriesResponseBody {
	s.Result = v
	return s
}

type ListGroupRepositoriesResponseBodyResult struct {
	LastActivityAt    *string `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	NamespaceId       *int64  `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	HttpCloneUrl      *string `json:"HttpCloneUrl,omitempty" xml:"HttpCloneUrl,omitempty"`
	Archive           *bool   `json:"Archive,omitempty" xml:"Archive,omitempty"`
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	SshCloneUrl       *string `json:"SshCloneUrl,omitempty" xml:"SshCloneUrl,omitempty"`
	CreatorId         *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	ImportStatus      *string `json:"ImportStatus,omitempty" xml:"ImportStatus,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	VisibilityLevel   *int32  `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdatedAt         *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListGroupRepositoriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRepositoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesResponseBodyResult) SetLastActivityAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetNamespaceId(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetHttpCloneUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.HttpCloneUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetArchive(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetCreatedAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetSshCloneUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.SshCloneUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetCreatorId(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetImportStatus(v string) *ListGroupRepositoriesResponseBodyResult {
	s.ImportStatus = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetWebUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetNameWithNamespace(v string) *ListGroupRepositoriesResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetPathWithNamespace(v string) *ListGroupRepositoriesResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetVisibilityLevel(v int32) *ListGroupRepositoriesResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetPath(v string) *ListGroupRepositoriesResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetUpdatedAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetName(v string) *ListGroupRepositoriesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetId(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.Id = &v
	return s
}

type ListGroupRepositoriesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesResponse) SetHeaders(v map[string]*string) *ListGroupRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListGroupRepositoriesResponse) SetBody(v *ListGroupRepositoriesResponseBody) *ListGroupRepositoriesResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	AccessToken     *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Search          *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrganizationId  *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page            *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SubUserId       *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	IncludePersonal *bool   `json:"IncludePersonal,omitempty" xml:"IncludePersonal,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetAccessToken(v string) *ListGroupsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListGroupsRequest) SetSearch(v string) *ListGroupsRequest {
	s.Search = &v
	return s
}

func (s *ListGroupsRequest) SetOrganizationId(v string) *ListGroupsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListGroupsRequest) SetPage(v int64) *ListGroupsRequest {
	s.Page = &v
	return s
}

func (s *ListGroupsRequest) SetPageSize(v int64) *ListGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsRequest) SetSubUserId(v string) *ListGroupsRequest {
	s.SubUserId = &v
	return s
}

func (s *ListGroupsRequest) SetIncludePersonal(v bool) *ListGroupsRequest {
	s.IncludePersonal = &v
	return s
}

type ListGroupsResponseBody struct {
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                          `json:"Total,omitempty" xml:"Total,omitempty"`
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListGroupsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotal(v int64) *ListGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *ListGroupsResponseBody) SetErrorCode(v string) *ListGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupsResponseBody) SetErrorMessage(v string) *ListGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupsResponseBody) SetSuccess(v bool) *ListGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupsResponseBody) SetResult(v []*ListGroupsResponseBodyResult) *ListGroupsResponseBody {
	s.Result = v
	return s
}

type ListGroupsResponseBodyResult struct {
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentId          *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	AccessLevel       *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	UpdatedAt         *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListGroupsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyResult) SetType(v string) *ListGroupsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetCreatedAt(v string) *ListGroupsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetOwnerId(v int64) *ListGroupsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetParentId(v int64) *ListGroupsResponseBodyResult {
	s.ParentId = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetWebUrl(v string) *ListGroupsResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetDescription(v string) *ListGroupsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetNameWithNamespace(v string) *ListGroupsResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetPathWithNamespace(v string) *ListGroupsResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetPath(v string) *ListGroupsResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetVisibilityLevel(v string) *ListGroupsResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetAccessLevel(v int32) *ListGroupsResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetUpdatedAt(v string) *ListGroupsResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetName(v string) *ListGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetId(v int64) *ListGroupsResponseBodyResult {
	s.Id = &v
	return s
}

type ListGroupsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetHeaders(v map[string]*string) *ListGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListOrganizationsRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AccessLevel    *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	MinAccessLevel *int32  `json:"MinAccessLevel,omitempty" xml:"MinAccessLevel,omitempty"`
}

func (s ListOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationsRequest) SetAccessToken(v string) *ListOrganizationsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListOrganizationsRequest) SetAccessLevel(v int32) *ListOrganizationsRequest {
	s.AccessLevel = &v
	return s
}

func (s *ListOrganizationsRequest) SetMinAccessLevel(v int32) *ListOrganizationsRequest {
	s.MinAccessLevel = &v
	return s
}

type ListOrganizationsResponseBody struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListOrganizationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationsResponseBody) SetRequestId(v string) *ListOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetErrorCode(v string) *ListOrganizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetErrorMessage(v string) *ListOrganizationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetSuccess(v bool) *ListOrganizationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetResult(v []*ListOrganizationsResponseBodyResult) *ListOrganizationsResponseBody {
	s.Result = v
	return s
}

type ListOrganizationsResponseBodyResult struct {
	OrganizationRole *string `json:"OrganizationRole,omitempty" xml:"OrganizationRole,omitempty"`
	AccessLevel      *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListOrganizationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationRole(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationRole = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetAccessLevel(v int32) *ListOrganizationsResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationName(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationId(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationId = &v
	return s
}

type ListOrganizationsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrganizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationsResponse) SetHeaders(v map[string]*string) *ListOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationsResponse) SetBody(v *ListOrganizationsResponseBody) *ListOrganizationsResponse {
	s.Body = v
	return s
}

type ListRepositoriesRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	Sort           *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	Archive        *bool   `json:"Archive,omitempty" xml:"Archive,omitempty"`
}

func (s ListRepositoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoriesRequest) SetAccessToken(v string) *ListRepositoriesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoriesRequest) SetOrganizationId(v string) *ListRepositoriesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoriesRequest) SetPage(v int64) *ListRepositoriesRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoriesRequest) SetPageSize(v int64) *ListRepositoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoriesRequest) SetOrder(v string) *ListRepositoriesRequest {
	s.Order = &v
	return s
}

func (s *ListRepositoriesRequest) SetSort(v string) *ListRepositoriesRequest {
	s.Sort = &v
	return s
}

func (s *ListRepositoriesRequest) SetSearch(v string) *ListRepositoriesRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoriesRequest) SetArchive(v bool) *ListRepositoriesRequest {
	s.Archive = &v
	return s
}

type ListRepositoriesResponseBody struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
	ErrorCode    *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListRepositoriesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBody) SetRequestId(v string) *ListRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetTotal(v int64) *ListRepositoriesResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetErrorCode(v int32) *ListRepositoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetErrorMessage(v string) *ListRepositoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetSuccess(v bool) *ListRepositoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetResult(v []*ListRepositoriesResponseBodyResult) *ListRepositoriesResponseBody {
	s.Result = v
	return s
}

type ListRepositoriesResponseBodyResult struct {
	LastActivityAt    *string `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	NamespaceId       *int64  `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	AvatarUrl         *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	StarCount         *int64  `json:"StarCount,omitempty" xml:"StarCount,omitempty"`
	Archive           *bool   `json:"Archive,omitempty" xml:"Archive,omitempty"`
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DemoProjectStatus *bool   `json:"DemoProjectStatus,omitempty" xml:"DemoProjectStatus,omitempty"`
	Star              *bool   `json:"Star,omitempty" xml:"Star,omitempty"`
	ImportStatus      *string `json:"ImportStatus,omitempty" xml:"ImportStatus,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	AccessLevel       *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	UpdatedAt         *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListRepositoriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBodyResult) SetLastActivityAt(v string) *ListRepositoriesResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetNamespaceId(v int64) *ListRepositoriesResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetAvatarUrl(v string) *ListRepositoriesResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetStarCount(v int64) *ListRepositoriesResponseBodyResult {
	s.StarCount = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetArchive(v bool) *ListRepositoriesResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetCreatedAt(v string) *ListRepositoriesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetDemoProjectStatus(v bool) *ListRepositoriesResponseBodyResult {
	s.DemoProjectStatus = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetStar(v bool) *ListRepositoriesResponseBodyResult {
	s.Star = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetImportStatus(v string) *ListRepositoriesResponseBodyResult {
	s.ImportStatus = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetWebUrl(v string) *ListRepositoriesResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetDescription(v string) *ListRepositoriesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetNameWithNamespace(v string) *ListRepositoriesResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetPathWithNamespace(v string) *ListRepositoriesResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetPath(v string) *ListRepositoriesResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetVisibilityLevel(v string) *ListRepositoriesResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetAccessLevel(v int32) *ListRepositoriesResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetUpdatedAt(v string) *ListRepositoriesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetName(v string) *ListRepositoriesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetId(v int64) *ListRepositoriesResponseBodyResult {
	s.Id = &v
	return s
}

type ListRepositoriesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponse) SetHeaders(v map[string]*string) *ListRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoriesResponse) SetBody(v *ListRepositoriesResponseBody) *ListRepositoriesResponse {
	s.Body = v
	return s
}

type ListRepositoryBranchesRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s ListRepositoryBranchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesRequest) SetAccessToken(v string) *ListRepositoryBranchesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetOrganizationId(v string) *ListRepositoryBranchesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetSubUserId(v string) *ListRepositoryBranchesRequest {
	s.SubUserId = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetPage(v int64) *ListRepositoryBranchesRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetPageSize(v int64) *ListRepositoryBranchesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryBranchesRequest) SetSearch(v string) *ListRepositoryBranchesRequest {
	s.Search = &v
	return s
}

type ListRepositoryBranchesResponseBody struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	ErrorCode    *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListRepositoryBranchesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRepositoryBranchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBody) SetRequestId(v string) *ListRepositoryBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetTotal(v int64) *ListRepositoryBranchesResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetErrorCode(v string) *ListRepositoryBranchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetErrorMessage(v string) *ListRepositoryBranchesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetSuccess(v bool) *ListRepositoryBranchesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetResult(v []*ListRepositoryBranchesResponseBodyResult) *ListRepositoryBranchesResponseBody {
	s.Result = v
	return s
}

type ListRepositoryBranchesResponseBodyResult struct {
	ProtectedBranch *bool                                               `json:"ProtectedBranch,omitempty" xml:"ProtectedBranch,omitempty"`
	CommitInfo      *ListRepositoryBranchesResponseBodyResultCommitInfo `json:"CommitInfo,omitempty" xml:"CommitInfo,omitempty" type:"Struct"`
	BranchName      *string                                             `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
}

func (s ListRepositoryBranchesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBodyResult) SetProtectedBranch(v bool) *ListRepositoryBranchesResponseBodyResult {
	s.ProtectedBranch = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) SetCommitInfo(v *ListRepositoryBranchesResponseBodyResultCommitInfo) *ListRepositoryBranchesResponseBodyResult {
	s.CommitInfo = v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) SetBranchName(v string) *ListRepositoryBranchesResponseBodyResult {
	s.BranchName = &v
	return s
}

type ListRepositoryBranchesResponseBodyResultCommitInfo struct {
	ShortId        *string   `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	AuthorDate     *string   `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	AuthorName     *string   `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommitterName  *string   `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	AuthorEmail    *string   `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	CommitterEmail *string   `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	CommittedDate  *string   `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
}

func (s ListRepositoryBranchesResponseBodyResultCommitInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBodyResultCommitInfo) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetShortId(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetAuthorDate(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.AuthorDate = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetCreatedAt(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetMessage(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.Message = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetParentIds(v []*string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.ParentIds = v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetAuthorName(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetCommitterName(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetTitle(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.Title = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetAuthorEmail(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetCommitterEmail(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetId(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.Id = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetCommittedDate(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.CommittedDate = &v
	return s
}

type ListRepositoryBranchesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRepositoryBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryBranchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponse) SetHeaders(v map[string]*string) *ListRepositoryBranchesResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryBranchesResponse) SetBody(v *ListRepositoryBranchesResponseBody) *ListRepositoryBranchesResponse {
	s.Body = v
	return s
}

type ListRepositoryMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListRepositoryMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberRequest) SetAccessToken(v string) *ListRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryMemberRequest) SetOrganizationId(v string) *ListRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryMemberRequest) SetQuery(v string) *ListRepositoryMemberRequest {
	s.Query = &v
	return s
}

func (s *ListRepositoryMemberRequest) SetPage(v int64) *ListRepositoryMemberRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryMemberRequest) SetPageSize(v int64) *ListRepositoryMemberRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryMemberRequest) SetSubUserId(v string) *ListRepositoryMemberRequest {
	s.SubUserId = &v
	return s
}

type ListRepositoryMemberResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                    `json:"Total,omitempty" xml:"Total,omitempty"`
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberResponseBody) SetRequestId(v string) *ListRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetTotal(v int64) *ListRepositoryMemberResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetErrorCode(v string) *ListRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetErrorMessage(v string) *ListRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetSuccess(v bool) *ListRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetResult(v []*ListRepositoryMemberResponseBodyResult) *ListRepositoryMemberResponseBody {
	s.Result = v
	return s
}

type ListRepositoryMemberResponseBodyResult struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberResponseBodyResult) SetExternUserId(v string) *ListRepositoryMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetEmail(v string) *ListRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *ListRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetState(v string) *ListRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *ListRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetId(v int64) *ListRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

type ListRepositoryMemberResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberResponse) SetHeaders(v map[string]*string) *ListRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryMemberResponse) SetBody(v *ListRepositoryMemberResponseBody) *ListRepositoryMemberResponse {
	s.Body = v
	return s
}

type ListRepositoryTagsRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Sort           *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	ShowSignature  *bool   `json:"ShowSignature,omitempty" xml:"ShowSignature,omitempty"`
}

func (s ListRepositoryTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsRequest) SetAccessToken(v string) *ListRepositoryTagsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetSearch(v string) *ListRepositoryTagsRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetOrganizationId(v string) *ListRepositoryTagsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetPage(v int64) *ListRepositoryTagsRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetPageSize(v int64) *ListRepositoryTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetSort(v string) *ListRepositoryTagsRequest {
	s.Sort = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetShowSignature(v bool) *ListRepositoryTagsRequest {
	s.ShowSignature = &v
	return s
}

type ListRepositoryTagsResponseBody struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                  `json:"Total,omitempty" xml:"Total,omitempty"`
	ErrorCode    *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListRepositoryTagsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRepositoryTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBody) SetRequestId(v string) *ListRepositoryTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetTotal(v int64) *ListRepositoryTagsResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetErrorCode(v string) *ListRepositoryTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetErrorMessage(v string) *ListRepositoryTagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetSuccess(v bool) *ListRepositoryTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetResult(v []*ListRepositoryTagsResponseBodyResult) *ListRepositoryTagsResponseBody {
	s.Result = v
	return s
}

type ListRepositoryTagsResponseBodyResult struct {
	Signature *ListRepositoryTagsResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Commit    *ListRepositoryTagsResponseBodyResultCommit    `json:"Commit,omitempty" xml:"Commit,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Name      *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListRepositoryTagsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResult) SetSignature(v *ListRepositoryTagsResponseBodyResultSignature) *ListRepositoryTagsResponseBodyResult {
	s.Signature = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) SetCommit(v *ListRepositoryTagsResponseBodyResultCommit) *ListRepositoryTagsResponseBodyResult {
	s.Commit = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) SetMessage(v string) *ListRepositoryTagsResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) SetName(v string) *ListRepositoryTagsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) SetId(v string) *ListRepositoryTagsResponseBodyResult {
	s.Id = &v
	return s
}

type ListRepositoryTagsResponseBodyResultSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s ListRepositoryTagsResponseBodyResultSignature) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResultSignature) SetGpgKeyId(v string) *ListRepositoryTagsResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultSignature) SetVerificationStatus(v string) *ListRepositoryTagsResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

type ListRepositoryTagsResponseBodyResultCommit struct {
	ShortId        *string                                              `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	CreatedAt      *string                                              `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	AuthoredDate   *string                                              `json:"AuthoredDate,omitempty" xml:"AuthoredDate,omitempty"`
	Signature      *ListRepositoryTagsResponseBodyResultCommitSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	ParentIds      []*string                                            `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	AuthorName     *string                                              `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommitterName  *string                                              `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	Title          *string                                              `json:"Title,omitempty" xml:"Title,omitempty"`
	AuthorEmail    *string                                              `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	CommitterEmail *string                                              `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	Id             *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	CommittedDate  *string                                              `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
}

func (s ListRepositoryTagsResponseBodyResultCommit) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetShortId(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCreatedAt(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetMessage(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthoredDate(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetSignature(v *ListRepositoryTagsResponseBodyResultCommitSignature) *ListRepositoryTagsResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetParentIds(v []*string) *ListRepositoryTagsResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthorName(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommitterName(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetTitle(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Title = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthorEmail(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommitterEmail(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetId(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommittedDate(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

type ListRepositoryTagsResponseBodyResultCommitSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s ListRepositoryTagsResponseBodyResultCommitSignature) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResultCommitSignature) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResultCommitSignature) SetGpgKeyId(v string) *ListRepositoryTagsResponseBodyResultCommitSignature {
	s.GpgKeyId = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommitSignature) SetVerificationStatus(v string) *ListRepositoryTagsResponseBodyResultCommitSignature {
	s.VerificationStatus = &v
	return s
}

type ListRepositoryTagsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRepositoryTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponse) SetHeaders(v map[string]*string) *ListRepositoryTagsResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryTagsResponse) SetBody(v *ListRepositoryTagsResponseBody) *ListRepositoryTagsResponse {
	s.Body = v
	return s
}

type ListRepositoryTreeRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RefName        *string `json:"RefName,omitempty" xml:"RefName,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListRepositoryTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeRequest) SetAccessToken(v string) *ListRepositoryTreeRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetOrganizationId(v string) *ListRepositoryTreeRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetPath(v string) *ListRepositoryTreeRequest {
	s.Path = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetType(v string) *ListRepositoryTreeRequest {
	s.Type = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetRefName(v string) *ListRepositoryTreeRequest {
	s.RefName = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetSubUserId(v string) *ListRepositoryTreeRequest {
	s.SubUserId = &v
	return s
}

type ListRepositoryTreeResponseBody struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListRepositoryTreeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRepositoryTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponseBody) SetRequestId(v string) *ListRepositoryTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetErrorCode(v string) *ListRepositoryTreeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetErrorMessage(v string) *ListRepositoryTreeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetSuccess(v bool) *ListRepositoryTreeResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetResult(v []*ListRepositoryTreeResponseBodyResult) *ListRepositoryTreeResponseBody {
	s.Result = v
	return s
}

type ListRepositoryTreeResponseBodyResult struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListRepositoryTreeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponseBodyResult) SetType(v string) *ListRepositoryTreeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetPath(v string) *ListRepositoryTreeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetMode(v string) *ListRepositoryTreeResponseBodyResult {
	s.Mode = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetName(v string) *ListRepositoryTreeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetId(v string) *ListRepositoryTreeResponseBodyResult {
	s.Id = &v
	return s
}

type ListRepositoryTreeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRepositoryTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponse) SetHeaders(v map[string]*string) *ListRepositoryTreeResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryTreeResponse) SetBody(v *ListRepositoryTreeResponseBody) *ListRepositoryTreeResponse {
	s.Body = v
	return s
}

type ListRepositoryWebhookRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRepositoryWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookRequest) SetAccessToken(v string) *ListRepositoryWebhookRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetOrganizationId(v string) *ListRepositoryWebhookRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetPage(v int64) *ListRepositoryWebhookRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetPageSize(v int64) *ListRepositoryWebhookRequest {
	s.PageSize = &v
	return s
}

type ListRepositoryWebhookResponseBody struct {
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                     `json:"Total,omitempty" xml:"Total,omitempty"`
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       []*ListRepositoryWebhookResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListRepositoryWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponseBody) SetRequestId(v string) *ListRepositoryWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetTotal(v int64) *ListRepositoryWebhookResponseBody {
	s.Total = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetErrorCode(v string) *ListRepositoryWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetErrorMessage(v string) *ListRepositoryWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetSuccess(v bool) *ListRepositoryWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetResult(v []*ListRepositoryWebhookResponseBodyResult) *ListRepositoryWebhookResponseBody {
	s.Result = v
	return s
}

type ListRepositoryWebhookResponseBodyResult struct {
	PushEvents            *bool   `json:"PushEvents,omitempty" xml:"PushEvents,omitempty"`
	ProjectId             *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	CreatedAt             *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	TagPushEvents         *bool   `json:"TagPushEvents,omitempty" xml:"TagPushEvents,omitempty"`
	Url                   *string `json:"Url,omitempty" xml:"Url,omitempty"`
	LastTestResult        *string `json:"LastTestResult,omitempty" xml:"LastTestResult,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MergeRequestsEvents   *bool   `json:"MergeRequestsEvents,omitempty" xml:"MergeRequestsEvents,omitempty"`
	SecretToken           *string `json:"SecretToken,omitempty" xml:"SecretToken,omitempty"`
	NoteEvents            *bool   `json:"NoteEvents,omitempty" xml:"NoteEvents,omitempty"`
	EnableSslVerification *bool   `json:"EnableSslVerification,omitempty" xml:"EnableSslVerification,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListRepositoryWebhookResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponseBodyResult) SetPushEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetProjectId(v int64) *ListRepositoryWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetCreatedAt(v string) *ListRepositoryWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetTagPushEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetUrl(v string) *ListRepositoryWebhookResponseBodyResult {
	s.Url = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetLastTestResult(v string) *ListRepositoryWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetDescription(v string) *ListRepositoryWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetSecretToken(v string) *ListRepositoryWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetNoteEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetEnableSslVerification(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetId(v int64) *ListRepositoryWebhookResponseBodyResult {
	s.Id = &v
	return s
}

type ListRepositoryWebhookResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRepositoryWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponse) SetHeaders(v map[string]*string) *ListRepositoryWebhookResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryWebhookResponse) SetBody(v *ListRepositoryWebhookResponseBody) *ListRepositoryWebhookResponse {
	s.Body = v
	return s
}

type MergeMergeRequestRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s MergeMergeRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestRequest) SetAccessToken(v string) *MergeMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *MergeMergeRequestRequest) SetOrganizationId(v string) *MergeMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *MergeMergeRequestRequest) SetSubUserId(v string) *MergeMergeRequestRequest {
	s.SubUserId = &v
	return s
}

type MergeMergeRequestResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *MergeMergeRequestResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s MergeMergeRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBody) SetRequestId(v string) *MergeMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetErrorCode(v string) *MergeMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetErrorMessage(v string) *MergeMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetSuccess(v bool) *MergeMergeRequestResponseBody {
	s.Success = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetResult(v *MergeMergeRequestResponseBodyResult) *MergeMergeRequestResponseBody {
	s.Result = v
	return s
}

type MergeMergeRequestResponseBodyResult struct {
	State              *string                                                `json:"State,omitempty" xml:"State,omitempty"`
	BehindCommitCount  *int32                                                 `json:"BehindCommitCount,omitempty" xml:"BehindCommitCount,omitempty"`
	ProjectId          *int64                                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	AssigneeList       []*MergeMergeRequestResponseBodyResultAssigneeList     `json:"AssigneeList,omitempty" xml:"AssigneeList,omitempty" type:"Repeated"`
	CreatedAt          *string                                                `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Author             *MergeMergeRequestResponseBodyResultAuthor             `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	AcceptedRevision   *string                                                `json:"AcceptedRevision,omitempty" xml:"AcceptedRevision,omitempty"`
	ApproveCheckResult *MergeMergeRequestResponseBodyResultApproveCheckResult `json:"ApproveCheckResult,omitempty" xml:"ApproveCheckResult,omitempty" type:"Struct"`
	SourceBranch       *string                                                `json:"SourceBranch,omitempty" xml:"SourceBranch,omitempty"`
	WebUrl             *string                                                `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	Description        *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	MergeType          *string                                                `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	NameWithNamespace  *string                                                `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	TargetBranch       *string                                                `json:"TargetBranch,omitempty" xml:"TargetBranch,omitempty"`
	AheadCommitCount   *int32                                                 `json:"AheadCommitCount,omitempty" xml:"AheadCommitCount,omitempty"`
	UpdatedAt          *string                                                `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Title              *string                                                `json:"Title,omitempty" xml:"Title,omitempty"`
	MergeError         *string                                                `json:"MergeError,omitempty" xml:"MergeError,omitempty"`
	MergedRevision     *string                                                `json:"MergedRevision,omitempty" xml:"MergedRevision,omitempty"`
	Id                 *int64                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	MergeStatus        *string                                                `json:"MergeStatus,omitempty" xml:"MergeStatus,omitempty"`
}

func (s MergeMergeRequestResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResult) SetState(v string) *MergeMergeRequestResponseBodyResult {
	s.State = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetBehindCommitCount(v int32) *MergeMergeRequestResponseBodyResult {
	s.BehindCommitCount = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetProjectId(v int64) *MergeMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetAssigneeList(v []*MergeMergeRequestResponseBodyResultAssigneeList) *MergeMergeRequestResponseBodyResult {
	s.AssigneeList = v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetCreatedAt(v string) *MergeMergeRequestResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetAuthor(v *MergeMergeRequestResponseBodyResultAuthor) *MergeMergeRequestResponseBodyResult {
	s.Author = v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetAcceptedRevision(v string) *MergeMergeRequestResponseBodyResult {
	s.AcceptedRevision = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetApproveCheckResult(v *MergeMergeRequestResponseBodyResultApproveCheckResult) *MergeMergeRequestResponseBodyResult {
	s.ApproveCheckResult = v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetSourceBranch(v string) *MergeMergeRequestResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetWebUrl(v string) *MergeMergeRequestResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetDescription(v string) *MergeMergeRequestResponseBodyResult {
	s.Description = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergeType(v string) *MergeMergeRequestResponseBodyResult {
	s.MergeType = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetNameWithNamespace(v string) *MergeMergeRequestResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetTargetBranch(v string) *MergeMergeRequestResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetAheadCommitCount(v int32) *MergeMergeRequestResponseBodyResult {
	s.AheadCommitCount = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetUpdatedAt(v string) *MergeMergeRequestResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetTitle(v string) *MergeMergeRequestResponseBodyResult {
	s.Title = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergeError(v string) *MergeMergeRequestResponseBodyResult {
	s.MergeError = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergedRevision(v string) *MergeMergeRequestResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetId(v int64) *MergeMergeRequestResponseBodyResult {
	s.Id = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergeStatus(v string) *MergeMergeRequestResponseBodyResult {
	s.MergeStatus = &v
	return s
}

type MergeMergeRequestResponseBodyResultAssigneeList struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultAssigneeList) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultAssigneeList) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultAssigneeList) SetExternUserId(v string) *MergeMergeRequestResponseBodyResultAssigneeList {
	s.ExternUserId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAssigneeList) SetAvatarUrl(v string) *MergeMergeRequestResponseBodyResultAssigneeList {
	s.AvatarUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAssigneeList) SetName(v string) *MergeMergeRequestResponseBodyResultAssigneeList {
	s.Name = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAssigneeList) SetId(v string) *MergeMergeRequestResponseBodyResultAssigneeList {
	s.Id = &v
	return s
}

type MergeMergeRequestResponseBodyResultAuthor struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultAuthor) SetExternUserId(v string) *MergeMergeRequestResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAuthor) SetAvatarUrl(v string) *MergeMergeRequestResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAuthor) SetName(v string) *MergeMergeRequestResponseBodyResultAuthor {
	s.Name = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAuthor) SetId(v int64) *MergeMergeRequestResponseBodyResultAuthor {
	s.Id = &v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResult struct {
	TotalCheckResult        *string                                                                         `json:"TotalCheckResult,omitempty" xml:"TotalCheckResult,omitempty"`
	UnsatisfiedCheckResults []*MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults `json:"UnsatisfiedCheckResults,omitempty" xml:"UnsatisfiedCheckResults,omitempty" type:"Repeated"`
	SatisfiedCheckResults   []*MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults   `json:"SatisfiedCheckResults,omitempty" xml:"SatisfiedCheckResults,omitempty" type:"Repeated"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResult) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResult) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResult) SetTotalCheckResult(v string) *MergeMergeRequestResponseBodyResultApproveCheckResult {
	s.TotalCheckResult = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResult) SetUnsatisfiedCheckResults(v []*MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) *MergeMergeRequestResponseBodyResultApproveCheckResult {
	s.UnsatisfiedCheckResults = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResult) SetSatisfiedCheckResults(v []*MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) *MergeMergeRequestResponseBodyResultApproveCheckResult {
	s.SatisfiedCheckResults = v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults struct {
	SatisfiedItems   []*string                                                                                 `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	CheckStatus      *string                                                                                   `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                   `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	UnsatisfiedItems []*string                                                                                 `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
	ExtraUsers       []*MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	CheckName        *string                                                                                   `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetSatisfiedItems(v []*string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckStatus(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckType(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetExtraUsers(v []*MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckName(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckName = &v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetName(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetId(v int64) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults struct {
	SatisfiedItems   []*string                                                                               `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	CheckStatus      *string                                                                                 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	UnsatisfiedItems []*string                                                                               `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
	ExtraUsers       []*MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	CheckName        *string                                                                                 `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetSatisfiedItems(v []*string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckStatus(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckType(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetExtraUsers(v []*MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckName(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckName = &v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetName(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetId(v int64) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

type MergeMergeRequestResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MergeMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MergeMergeRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponse) SetHeaders(v map[string]*string) *MergeMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *MergeMergeRequestResponse) SetBody(v *MergeMergeRequestResponseBody) *MergeMergeRequestResponse {
	s.Body = v
	return s
}

type UpdateFileRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s UpdateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileRequest) SetAccessToken(v string) *UpdateFileRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateFileRequest) SetOrganizationId(v string) *UpdateFileRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateFileRequest) SetSubUserId(v string) *UpdateFileRequest {
	s.SubUserId = &v
	return s
}

type UpdateFileResponseBody struct {
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *UpdateFileResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBody) SetRequestId(v string) *UpdateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileResponseBody) SetErrorCode(v string) *UpdateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFileResponseBody) SetErrorMessage(v string) *UpdateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFileResponseBody) SetSuccess(v bool) *UpdateFileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFileResponseBody) SetResult(v *UpdateFileResponseBodyResult) *UpdateFileResponseBody {
	s.Result = v
	return s
}

type UpdateFileResponseBodyResult struct {
	FilePath   *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
}

func (s UpdateFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBodyResult) SetFilePath(v string) *UpdateFileResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *UpdateFileResponseBodyResult) SetBranchName(v string) *UpdateFileResponseBodyResult {
	s.BranchName = &v
	return s
}

type UpdateFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileResponse) SetHeaders(v map[string]*string) *UpdateFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileResponse) SetBody(v *UpdateFileResponseBody) *UpdateFileResponse {
	s.Body = v
	return s
}

type UpdateGroupMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s UpdateGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberRequest) SetAccessToken(v string) *UpdateGroupMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateGroupMemberRequest) SetOrganizationId(v string) *UpdateGroupMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateGroupMemberRequest) SetSubUserId(v string) *UpdateGroupMemberRequest {
	s.SubUserId = &v
	return s
}

type UpdateGroupMemberResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *UpdateGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberResponseBody) SetRequestId(v string) *UpdateGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetErrorCode(v string) *UpdateGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetErrorMessage(v string) *UpdateGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetSuccess(v bool) *UpdateGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetResult(v *UpdateGroupMemberResponseBodyResult) *UpdateGroupMemberResponseBody {
	s.Result = v
	return s
}

type UpdateGroupMemberResponseBodyResult struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberResponseBodyResult) SetExternUserId(v string) *UpdateGroupMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetEmail(v string) *UpdateGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetAvatarUrl(v string) *UpdateGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetState(v string) *UpdateGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetAccessLevel(v int32) *UpdateGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetId(v int64) *UpdateGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

type UpdateGroupMemberResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberResponse) SetHeaders(v map[string]*string) *UpdateGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupMemberResponse) SetBody(v *UpdateGroupMemberResponseBody) *UpdateGroupMemberResponse {
	s.Body = v
	return s
}

type UpdateRepositoryMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s UpdateRepositoryMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberRequest) SetAccessToken(v string) *UpdateRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetOrganizationId(v string) *UpdateRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetSubUserId(v string) *UpdateRepositoryMemberRequest {
	s.SubUserId = &v
	return s
}

type UpdateRepositoryMemberResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Result       *UpdateRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponseBody) SetRequestId(v string) *UpdateRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetErrorCode(v string) *UpdateRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetErrorMessage(v string) *UpdateRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetSuccess(v bool) *UpdateRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetResult(v *UpdateRepositoryMemberResponseBodyResult) *UpdateRepositoryMemberResponseBody {
	s.Result = v
	return s
}

type UpdateRepositoryMemberResponseBodyResult struct {
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetExternUserId(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetEmail(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetState(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *UpdateRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetId(v int64) *UpdateRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

type UpdateRepositoryMemberResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRepositoryMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponse) SetHeaders(v map[string]*string) *UpdateRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepositoryMemberResponse) SetBody(v *UpdateRepositoryMemberResponseBody) *UpdateRepositoryMemberResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("codeup"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddGroupMember(GroupId *string, request *AddGroupMemberRequest) (_result *AddGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddGroupMemberResponse{}
	_body, _err := client.AddGroupMemberWithOptions(GroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGroupMemberWithOptions(GroupId *string, request *AddGroupMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("AddGroupMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v4/groups/"+tea.StringValue(GroupId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRepositoryMember(ProjectId *string, request *AddRepositoryMemberRequest) (_result *AddRepositoryMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddRepositoryMemberResponse{}
	_body, _err := client.AddRepositoryMemberWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRepositoryMemberWithOptions(ProjectId *string, request *AddRepositoryMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddRepositoryMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddRepositoryMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("AddRepositoryMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v4/projects/"+tea.StringValue(ProjectId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWebhook(ProjectId *string, request *AddWebhookRequest) (_result *AddWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddWebhookResponse{}
	_body, _err := client.AddWebhookWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWebhookWithOptions(ProjectId *string, request *AddWebhookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddWebhookResponse{}
	_body, _err := client.DoROARequest(tea.String("AddWebhook"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/hooks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBranch(ProjectId *string, request *CreateBranchRequest) (_result *CreateBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBranchResponse{}
	_body, _err := client.CreateBranchWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBranchWithOptions(ProjectId *string, request *CreateBranchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBranchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateBranchResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateBranch"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/branches"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFile(ProjectId *string, request *CreateFileRequest) (_result *CreateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFileResponse{}
	_body, _err := client.CreateFileWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFileWithOptions(ProjectId *string, request *CreateFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateFile"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/files"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMergeRequest(ProjectId *string, request *CreateMergeRequestRequest) (_result *CreateMergeRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMergeRequestResponse{}
	_body, _err := client.CreateMergeRequestWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMergeRequestWithOptions(ProjectId *string, request *CreateMergeRequestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMergeRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateMergeRequestResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateMergeRequest"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v4/projects/"+tea.StringValue(ProjectId)+"/merge_requests"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRepository(request *CreateRepositoryRequest) (_result *CreateRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CreateRepositoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRepositoryWithOptions(request *CreateRepositoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sync)) {
		query["Sync"] = request.Sync
	}

	if !tea.BoolValue(util.IsUnset(request.CreateParentPath)) {
		query["CreateParentPath"] = request.CreateParentPath
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateRepository"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v3/projects"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRepositoryGroup(request *CreateRepositoryGroupRequest) (_result *CreateRepositoryGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRepositoryGroupResponse{}
	_body, _err := client.CreateRepositoryGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRepositoryGroupWithOptions(request *CreateRepositoryGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRepositoryGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateRepositoryGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateRepositoryGroup"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v3/groups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTag(ProjectId *string, request *CreateTagRequest) (_result *CreateTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTagResponse{}
	_body, _err := client.CreateTagWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTagWithOptions(ProjectId *string, request *CreateTagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTag"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/tags"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBranch(ProjectId *string, request *DeleteBranchRequest) (_result *DeleteBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBranchResponse{}
	_body, _err := client.DeleteBranchWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBranchWithOptions(ProjectId *string, request *DeleteBranchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBranchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		query["BranchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteBranchResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteBranch"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/branches/delete"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFile(ProjectId *string, request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileWithOptions(ProjectId *string, request *DeleteFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		query["BranchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.CommitMessage)) {
		query["CommitMessage"] = request.CommitMessage
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteFile"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/files"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroupMember(GroupId *string, UserId *string, request *DeleteGroupMemberRequest) (_result *DeleteGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGroupMemberResponse{}
	_body, _err := client.DeleteGroupMemberWithOptions(GroupId, UserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupMemberWithOptions(GroupId *string, UserId *string, request *DeleteGroupMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteGroupMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v3/groups/"+tea.StringValue(GroupId)+"/members/{UserId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepository(ProjectId *string, request *DeleteRepositoryRequest) (_result *DeleteRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.DeleteRepositoryWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryWithOptions(ProjectId *string, request *DeleteRepositoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRepository"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepositoryGroup(GroupId *string, request *DeleteRepositoryGroupRequest) (_result *DeleteRepositoryGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryGroupResponse{}
	_body, _err := client.DeleteRepositoryGroupWithOptions(GroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryGroupWithOptions(GroupId *string, request *DeleteRepositoryGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteRepositoryGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRepositoryGroup"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v3/groups/"+tea.StringValue(GroupId)+"/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepositoryMember(ProjectId *string, UserId *string, request *DeleteRepositoryMemberRequest) (_result *DeleteRepositoryMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryMemberResponse{}
	_body, _err := client.DeleteRepositoryMemberWithOptions(ProjectId, UserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryMemberWithOptions(ProjectId *string, UserId *string, request *DeleteRepositoryMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteRepositoryMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRepositoryMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/members/{UserId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepositoryTag(ProjectId *string, TagName *string, request *DeleteRepositoryTagRequest) (_result *DeleteRepositoryTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryTagResponse{}
	_body, _err := client.DeleteRepositoryTagWithOptions(ProjectId, TagName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryTagWithOptions(ProjectId *string, TagName *string, request *DeleteRepositoryTagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteRepositoryTagResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRepositoryTag"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/tags/{TagName}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepositoryWebhook(ProjectId *string, WebhookId *string, request *DeleteRepositoryWebhookRequest) (_result *DeleteRepositoryWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryWebhookResponse{}
	_body, _err := client.DeleteRepositoryWebhookWithOptions(ProjectId, WebhookId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryWebhookWithOptions(ProjectId *string, WebhookId *string, request *DeleteRepositoryWebhookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteRepositoryWebhookResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteRepositoryWebhook"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/hooks/{WebhookId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBranchInfo(ProjectId *string, request *GetBranchInfoRequest) (_result *GetBranchInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBranchInfoResponse{}
	_body, _err := client.GetBranchInfoWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBranchInfoWithOptions(ProjectId *string, request *GetBranchInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBranchInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		query["BranchName"] = request.BranchName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetBranchInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetBranchInfo"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/branches/detail"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCodeupOrganization(OrganizationIdentity *string, request *GetCodeupOrganizationRequest) (_result *GetCodeupOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeupOrganizationResponse{}
	_body, _err := client.GetCodeupOrganizationWithOptions(OrganizationIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCodeupOrganizationWithOptions(OrganizationIdentity *string, request *GetCodeupOrganizationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCodeupOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetCodeupOrganizationResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCodeupOrganization"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v4/organization/"+tea.StringValue(OrganizationIdentity)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileBlobs(ProjectId *string, request *GetFileBlobsRequest) (_result *GetFileBlobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileBlobsResponse{}
	_body, _err := client.GetFileBlobsWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileBlobsWithOptions(ProjectId *string, request *GetFileBlobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileBlobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.Ref)) {
		query["Ref"] = request.Ref
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetFileBlobsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFileBlobs"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v4/projects/"+tea.StringValue(ProjectId)+"/repository/blobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGroupDetail(request *GetGroupDetailRequest) (_result *GetGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGroupDetailResponse{}
	_body, _err := client.GetGroupDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupDetailWithOptions(request *GetGroupDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetGroupDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetGroupDetail"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/groups/detail"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectMember(ProjectId *string, UserId *string, request *GetProjectMemberRequest) (_result *GetProjectMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectMemberResponse{}
	_body, _err := client.GetProjectMemberWithOptions(ProjectId, UserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectMemberWithOptions(ProjectId *string, UserId *string, request *GetProjectMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProjectMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/members/{UserId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRepositoryInfo(request *GetRepositoryInfoRequest) (_result *GetRepositoryInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRepositoryInfoResponse{}
	_body, _err := client.GetRepositoryInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRepositoryInfoWithOptions(request *GetRepositoryInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRepositoryInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		query["Identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetRepositoryInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRepositoryInfo"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/projects/info"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRepositoryTag(ProjectId *string, TagName *string, request *GetRepositoryTagRequest) (_result *GetRepositoryTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRepositoryTagResponse{}
	_body, _err := client.GetRepositoryTagWithOptions(ProjectId, TagName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRepositoryTagWithOptions(ProjectId *string, TagName *string, request *GetRepositoryTagRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRepositoryTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetRepositoryTagResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRepositoryTag"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/tags/{TagName}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserInfo(request *GetUserInfoRequest) (_result *GetUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserInfoResponse{}
	_body, _err := client.GetUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserInfoWithOptions(request *GetUserInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserInfo"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/user/current"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupMember(GroupId *string, request *ListGroupMemberRequest) (_result *ListGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupMemberResponse{}
	_body, _err := client.ListGroupMemberWithOptions(GroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupMemberWithOptions(GroupId *string, request *ListGroupMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("ListGroupMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/groups/"+tea.StringValue(GroupId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupRepositories(Identity *string, request *ListGroupRepositoriesRequest) (_result *ListGroupRepositoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupRepositoriesResponse{}
	_body, _err := client.ListGroupRepositoriesWithOptions(Identity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupRepositoriesWithOptions(Identity *string, request *ListGroupRepositoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGroupRepositoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.IsMember)) {
		query["IsMember"] = request.IsMember
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListGroupRepositoriesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListGroupRepositories"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/groups/"+tea.StringValue(Identity)+"/projects"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludePersonal)) {
		query["IncludePersonal"] = request.IncludePersonal
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListGroups"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/groups/all"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrganizations(request *ListOrganizationsRequest) (_result *ListOrganizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOrganizationsResponse{}
	_body, _err := client.ListOrganizationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrganizationsWithOptions(request *ListOrganizationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.AccessLevel)) {
		query["AccessLevel"] = request.AccessLevel
	}

	if !tea.BoolValue(util.IsUnset(request.MinAccessLevel)) {
		query["MinAccessLevel"] = request.MinAccessLevel
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListOrganizationsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListOrganizations"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v4/organization"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositories(request *ListRepositoriesRequest) (_result *ListRepositoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoriesResponse{}
	_body, _err := client.ListRepositoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoriesWithOptions(request *ListRepositoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.Archive)) {
		query["Archive"] = request.Archive
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRepositoriesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRepositories"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/projects/all"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryBranches(ProjectId *string, request *ListRepositoryBranchesRequest) (_result *ListRepositoryBranchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryBranchesResponse{}
	_body, _err := client.ListRepositoryBranchesWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryBranchesWithOptions(ProjectId *string, request *ListRepositoryBranchesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryBranchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRepositoryBranchesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRepositoryBranches"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/repository/branches"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryMember(ProjectId *string, request *ListRepositoryMemberRequest) (_result *ListRepositoryMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryMemberResponse{}
	_body, _err := client.ListRepositoryMemberWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryMemberWithOptions(ProjectId *string, request *ListRepositoryMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRepositoryMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRepositoryMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryTags(ProjectId *string, request *ListRepositoryTagsRequest) (_result *ListRepositoryTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryTagsResponse{}
	_body, _err := client.ListRepositoryTagsWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryTagsWithOptions(ProjectId *string, request *ListRepositoryTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSignature)) {
		query["ShowSignature"] = request.ShowSignature
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRepositoryTagsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRepositoryTags"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v4/projects/"+tea.StringValue(ProjectId)+"/repository/tags"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryTree(ProjectId *string, request *ListRepositoryTreeRequest) (_result *ListRepositoryTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryTreeResponse{}
	_body, _err := client.ListRepositoryTreeWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryTreeWithOptions(ProjectId *string, request *ListRepositoryTreeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.RefName)) {
		query["RefName"] = request.RefName
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRepositoryTreeResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRepositoryTree"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v4/projects/"+tea.StringValue(ProjectId)+"/repository/tree"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryWebhook(ProjectId *string, request *ListRepositoryWebhookRequest) (_result *ListRepositoryWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryWebhookResponse{}
	_body, _err := client.ListRepositoryWebhookWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryWebhookWithOptions(ProjectId *string, request *ListRepositoryWebhookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListRepositoryWebhookResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRepositoryWebhook"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/hooks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MergeMergeRequest(ProjectId *string, MergeRequestId *string, request *MergeMergeRequestRequest) (_result *MergeMergeRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MergeMergeRequestResponse{}
	_body, _err := client.MergeMergeRequestWithOptions(ProjectId, MergeRequestId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeMergeRequestWithOptions(ProjectId *string, MergeRequestId *string, request *MergeMergeRequestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MergeMergeRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &MergeMergeRequestResponse{}
	_body, _err := client.DoROARequest(tea.String("MergeMergeRequest"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/mergeRequest/{MergeRequestId}/merge"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFile(ProjectId *string, request *UpdateFileRequest) (_result *UpdateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFileResponse{}
	_body, _err := client.UpdateFileWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFileWithOptions(ProjectId *string, request *UpdateFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateFile"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v4/projects/"+tea.StringValue(ProjectId)+"/repository/files"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupMember(GroupId *string, UserId *string, request *UpdateGroupMemberRequest) (_result *UpdateGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGroupMemberResponse{}
	_body, _err := client.UpdateGroupMemberWithOptions(GroupId, UserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupMemberWithOptions(GroupId *string, UserId *string, request *UpdateGroupMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &UpdateGroupMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateGroupMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v3/groups/"+tea.StringValue(GroupId)+"/members/{UserId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRepositoryMember(ProjectId *string, UserId *string, request *UpdateRepositoryMemberRequest) (_result *UpdateRepositoryMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRepositoryMemberResponse{}
	_body, _err := client.UpdateRepositoryMemberWithOptions(ProjectId, UserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRepositoryMemberWithOptions(ProjectId *string, UserId *string, request *UpdateRepositoryMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRepositoryMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &UpdateRepositoryMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateRepositoryMember"), tea.String("2020-04-14"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v3/projects/"+tea.StringValue(ProjectId)+"/members/{UserId}"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
