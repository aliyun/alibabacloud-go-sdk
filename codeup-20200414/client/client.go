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

type AcceptMergeRequestRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AcceptMergeRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestRequest) SetAccessToken(v string) *AcceptMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *AcceptMergeRequestRequest) SetOrganizationId(v string) *AcceptMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

type AcceptMergeRequestResponseBody struct {
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *AcceptMergeRequestResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AcceptMergeRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBody) SetErrorCode(v string) *AcceptMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AcceptMergeRequestResponseBody) SetErrorMessage(v string) *AcceptMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AcceptMergeRequestResponseBody) SetRequestId(v string) *AcceptMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptMergeRequestResponseBody) SetResult(v *AcceptMergeRequestResponseBodyResult) *AcceptMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *AcceptMergeRequestResponseBody) SetSuccess(v bool) *AcceptMergeRequestResponseBody {
	s.Success = &v
	return s
}

type AcceptMergeRequestResponseBodyResult struct {
	AcceptedRevision   *string                                                 `json:"AcceptedRevision,omitempty" xml:"AcceptedRevision,omitempty"`
	AheadCommitCount   *int32                                                  `json:"AheadCommitCount,omitempty" xml:"AheadCommitCount,omitempty"`
	ApproveCheckResult *AcceptMergeRequestResponseBodyResultApproveCheckResult `json:"ApproveCheckResult,omitempty" xml:"ApproveCheckResult,omitempty" type:"Struct"`
	AssigneeList       []*AcceptMergeRequestResponseBodyResultAssigneeList     `json:"AssigneeList,omitempty" xml:"AssigneeList,omitempty" type:"Repeated"`
	Author             *AcceptMergeRequestResponseBodyResultAuthor             `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	BehindCommitCount  *int32                                                  `json:"BehindCommitCount,omitempty" xml:"BehindCommitCount,omitempty"`
	CreatedAt          *string                                                 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description        *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                 *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	MergeError         *string                                                 `json:"MergeError,omitempty" xml:"MergeError,omitempty"`
	MergeStatus        *string                                                 `json:"MergeStatus,omitempty" xml:"MergeStatus,omitempty"`
	MergeType          *string                                                 `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	MergedRevision     *string                                                 `json:"MergedRevision,omitempty" xml:"MergedRevision,omitempty"`
	NameWithNamespace  *string                                                 `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	ProjectId          *int64                                                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SourceBranch       *string                                                 `json:"SourceBranch,omitempty" xml:"SourceBranch,omitempty"`
	State              *string                                                 `json:"State,omitempty" xml:"State,omitempty"`
	TargetBranch       *string                                                 `json:"TargetBranch,omitempty" xml:"TargetBranch,omitempty"`
	Title              *string                                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdatedAt          *string                                                 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	WebUrl             *string                                                 `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s AcceptMergeRequestResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBodyResult) SetAcceptedRevision(v string) *AcceptMergeRequestResponseBodyResult {
	s.AcceptedRevision = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetAheadCommitCount(v int32) *AcceptMergeRequestResponseBodyResult {
	s.AheadCommitCount = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetApproveCheckResult(v *AcceptMergeRequestResponseBodyResultApproveCheckResult) *AcceptMergeRequestResponseBodyResult {
	s.ApproveCheckResult = v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetAssigneeList(v []*AcceptMergeRequestResponseBodyResultAssigneeList) *AcceptMergeRequestResponseBodyResult {
	s.AssigneeList = v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetAuthor(v *AcceptMergeRequestResponseBodyResultAuthor) *AcceptMergeRequestResponseBodyResult {
	s.Author = v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetBehindCommitCount(v int32) *AcceptMergeRequestResponseBodyResult {
	s.BehindCommitCount = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetCreatedAt(v string) *AcceptMergeRequestResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetDescription(v string) *AcceptMergeRequestResponseBodyResult {
	s.Description = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetId(v int64) *AcceptMergeRequestResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetMergeError(v string) *AcceptMergeRequestResponseBodyResult {
	s.MergeError = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetMergeStatus(v string) *AcceptMergeRequestResponseBodyResult {
	s.MergeStatus = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetMergeType(v string) *AcceptMergeRequestResponseBodyResult {
	s.MergeType = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetMergedRevision(v string) *AcceptMergeRequestResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetNameWithNamespace(v string) *AcceptMergeRequestResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetProjectId(v int64) *AcceptMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetSourceBranch(v string) *AcceptMergeRequestResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetState(v string) *AcceptMergeRequestResponseBodyResult {
	s.State = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetTargetBranch(v string) *AcceptMergeRequestResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetTitle(v string) *AcceptMergeRequestResponseBodyResult {
	s.Title = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetUpdatedAt(v string) *AcceptMergeRequestResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResult) SetWebUrl(v string) *AcceptMergeRequestResponseBodyResult {
	s.WebUrl = &v
	return s
}

type AcceptMergeRequestResponseBodyResultApproveCheckResult struct {
	SatisfiedCheckResults   []*AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults   `json:"SatisfiedCheckResults,omitempty" xml:"SatisfiedCheckResults,omitempty" type:"Repeated"`
	TotalCheckResult        *string                                                                          `json:"TotalCheckResult,omitempty" xml:"TotalCheckResult,omitempty"`
	UnsatisfiedCheckResults []*AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults `json:"UnsatisfiedCheckResults,omitempty" xml:"UnsatisfiedCheckResults,omitempty" type:"Repeated"`
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResult) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResult) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResult) SetSatisfiedCheckResults(v []*AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) *AcceptMergeRequestResponseBodyResultApproveCheckResult {
	s.SatisfiedCheckResults = v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResult) SetTotalCheckResult(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResult {
	s.TotalCheckResult = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResult) SetUnsatisfiedCheckResults(v []*AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) *AcceptMergeRequestResponseBodyResultApproveCheckResult {
	s.UnsatisfiedCheckResults = v
	return s
}

type AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults struct {
	CheckName        *string                                                                                  `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckName(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckName = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckStatus(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckType(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetExtraUsers(v []*AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetSatisfiedItems(v []*string) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetId(v int64) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetName(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults struct {
	CheckName        *string                                                                                    `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                    `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                    `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                  `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                  `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckName(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckName = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckStatus(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckType(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetExtraUsers(v []*AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetSatisfiedItems(v []*string) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetId(v int64) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetName(v string) *AcceptMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type AcceptMergeRequestResponseBodyResultAssigneeList struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AcceptMergeRequestResponseBodyResultAssigneeList) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBodyResultAssigneeList) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBodyResultAssigneeList) SetAvatarUrl(v string) *AcceptMergeRequestResponseBodyResultAssigneeList {
	s.AvatarUrl = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultAssigneeList) SetExternUserId(v string) *AcceptMergeRequestResponseBodyResultAssigneeList {
	s.ExternUserId = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultAssigneeList) SetId(v string) *AcceptMergeRequestResponseBodyResultAssigneeList {
	s.Id = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultAssigneeList) SetName(v string) *AcceptMergeRequestResponseBodyResultAssigneeList {
	s.Name = &v
	return s
}

type AcceptMergeRequestResponseBodyResultAuthor struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AcceptMergeRequestResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponseBodyResultAuthor) SetAvatarUrl(v string) *AcceptMergeRequestResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultAuthor) SetExternUserId(v string) *AcceptMergeRequestResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultAuthor) SetId(v int64) *AcceptMergeRequestResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *AcceptMergeRequestResponseBodyResultAuthor) SetName(v string) *AcceptMergeRequestResponseBodyResultAuthor {
	s.Name = &v
	return s
}

type AcceptMergeRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AcceptMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AcceptMergeRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *AcceptMergeRequestResponse) SetHeaders(v map[string]*string) *AcceptMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *AcceptMergeRequestResponse) SetStatusCode(v int32) *AcceptMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptMergeRequestResponse) SetBody(v *AcceptMergeRequestResponseBody) *AcceptMergeRequestResponse {
	s.Body = v
	return s
}

type AddGroupMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
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

type AddGroupMemberResponseBody struct {
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*AddGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponseBody) SetErrorCode(v string) *AddGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetErrorMessage(v string) *AddGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetRequestId(v string) *AddGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGroupMemberResponseBody) SetResult(v []*AddGroupMemberResponseBodyResult) *AddGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *AddGroupMemberResponseBody) SetSuccess(v bool) *AddGroupMemberResponseBody {
	s.Success = &v
	return s
}

type AddGroupMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s AddGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponseBodyResult) SetAccessLevel(v int32) *AddGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetAvatarUrl(v string) *AddGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetEmail(v string) *AddGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetExternUserId(v string) *AddGroupMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetId(v int64) *AddGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddGroupMemberResponseBodyResult) SetState(v string) *AddGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

type AddGroupMemberResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddGroupMemberResponse) SetStatusCode(v int32) *AddGroupMemberResponse {
	s.StatusCode = &v
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

type AddRepositoryMemberResponseBody struct {
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*AddRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponseBody) SetErrorCode(v string) *AddRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetErrorMessage(v string) *AddRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetRequestId(v string) *AddRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetResult(v []*AddRepositoryMemberResponseBodyResult) *AddRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *AddRepositoryMemberResponseBody) SetSuccess(v bool) *AddRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

type AddRepositoryMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s AddRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *AddRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *AddRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetEmail(v string) *AddRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetExternUserId(v string) *AddRepositoryMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetId(v int64) *AddRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddRepositoryMemberResponseBodyResult) SetState(v string) *AddRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

type AddRepositoryMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddRepositoryMemberResponse) SetStatusCode(v int32) *AddRepositoryMemberResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *AddWebhookResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *AddWebhookResponseBody) SetErrorCode(v string) *AddWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddWebhookResponseBody) SetErrorMessage(v string) *AddWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddWebhookResponseBody) SetRequestId(v string) *AddWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWebhookResponseBody) SetResult(v *AddWebhookResponseBodyResult) *AddWebhookResponseBody {
	s.Result = v
	return s
}

func (s *AddWebhookResponseBody) SetSuccess(v bool) *AddWebhookResponseBody {
	s.Success = &v
	return s
}

type AddWebhookResponseBodyResult struct {
	BuildEvents           *bool   `json:"BuildEvents,omitempty" xml:"BuildEvents,omitempty"`
	CreatedAt             *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableSslVerification *bool   `json:"EnableSslVerification,omitempty" xml:"EnableSslVerification,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IssuesEvents          *bool   `json:"IssuesEvents,omitempty" xml:"IssuesEvents,omitempty"`
	LastTestResult        *string `json:"LastTestResult,omitempty" xml:"LastTestResult,omitempty"`
	MergeRequestsEvents   *bool   `json:"MergeRequestsEvents,omitempty" xml:"MergeRequestsEvents,omitempty"`
	NoteEvents            *bool   `json:"NoteEvents,omitempty" xml:"NoteEvents,omitempty"`
	ProjectId             *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PushEvents            *bool   `json:"PushEvents,omitempty" xml:"PushEvents,omitempty"`
	SecretToken           *string `json:"SecretToken,omitempty" xml:"SecretToken,omitempty"`
	TagPushEvents         *bool   `json:"TagPushEvents,omitempty" xml:"TagPushEvents,omitempty"`
	Url                   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s AddWebhookResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddWebhookResponseBodyResult) SetBuildEvents(v bool) *AddWebhookResponseBodyResult {
	s.BuildEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetCreatedAt(v string) *AddWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetDescription(v string) *AddWebhookResponseBodyResult {
	s.Description = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetEnableSslVerification(v bool) *AddWebhookResponseBodyResult {
	s.EnableSslVerification = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetId(v int64) *AddWebhookResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetIssuesEvents(v bool) *AddWebhookResponseBodyResult {
	s.IssuesEvents = &v
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

func (s *AddWebhookResponseBodyResult) SetNoteEvents(v bool) *AddWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetProjectId(v int64) *AddWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetPushEvents(v bool) *AddWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetSecretToken(v string) *AddWebhookResponseBodyResult {
	s.SecretToken = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetTagPushEvents(v bool) *AddWebhookResponseBodyResult {
	s.TagPushEvents = &v
	return s
}

func (s *AddWebhookResponseBodyResult) SetUrl(v string) *AddWebhookResponseBodyResult {
	s.Url = &v
	return s
}

type AddWebhookResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddWebhookResponse) SetStatusCode(v int32) *AddWebhookResponse {
	s.StatusCode = &v
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
	BranchName     *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	Ref            *string `json:"ref,omitempty" xml:"ref,omitempty"`
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

func (s *CreateBranchRequest) SetBranchName(v string) *CreateBranchRequest {
	s.BranchName = &v
	return s
}

func (s *CreateBranchRequest) SetRef(v string) *CreateBranchRequest {
	s.Ref = &v
	return s
}

type CreateBranchResponseBody struct {
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateBranchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBody) SetErrorCode(v string) *CreateBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateBranchResponseBody) SetErrorMessage(v string) *CreateBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateBranchResponseBody) SetRequestId(v string) *CreateBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBranchResponseBody) SetResult(v *CreateBranchResponseBodyResult) *CreateBranchResponseBody {
	s.Result = v
	return s
}

func (s *CreateBranchResponseBody) SetSuccess(v bool) *CreateBranchResponseBody {
	s.Success = &v
	return s
}

type CreateBranchResponseBodyResult struct {
	BranchName      *string                                   `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	CommitInfo      *CreateBranchResponseBodyResultCommitInfo `json:"CommitInfo,omitempty" xml:"CommitInfo,omitempty" type:"Struct"`
	ProtectedBranch *bool                                     `json:"ProtectedBranch,omitempty" xml:"ProtectedBranch,omitempty"`
}

func (s CreateBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBodyResult) SetBranchName(v string) *CreateBranchResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *CreateBranchResponseBodyResult) SetCommitInfo(v *CreateBranchResponseBodyResultCommitInfo) *CreateBranchResponseBodyResult {
	s.CommitInfo = v
	return s
}

func (s *CreateBranchResponseBodyResult) SetProtectedBranch(v bool) *CreateBranchResponseBodyResult {
	s.ProtectedBranch = &v
	return s
}

type CreateBranchResponseBodyResultCommitInfo struct {
	AuthorDate     *string   `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	AuthorEmail    *string   `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string   `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommittedDate  *string   `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string   `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string   `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string   `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateBranchResponseBodyResultCommitInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateBranchResponseBodyResultCommitInfo) GoString() string {
	return s.String()
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetAuthorDate(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.AuthorDate = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetAuthorEmail(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.AuthorEmail = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetAuthorName(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.AuthorName = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetCommittedDate(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.CommittedDate = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetCommitterEmail(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.CommitterEmail = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetCommitterName(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.CommitterName = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetCreatedAt(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.CreatedAt = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetId(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.Id = &v
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

func (s *CreateBranchResponseBodyResultCommitInfo) SetShortId(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.ShortId = &v
	return s
}

func (s *CreateBranchResponseBodyResultCommitInfo) SetTitle(v string) *CreateBranchResponseBodyResultCommitInfo {
	s.Title = &v
	return s
}

type CreateBranchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateBranchResponse) SetStatusCode(v int32) *CreateBranchResponse {
	s.StatusCode = &v
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

type CreateFileResponseBody struct {
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateFileResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) SetErrorCode(v string) *CreateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFileResponseBody) SetErrorMessage(v string) *CreateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFileResponseBody) SetRequestId(v string) *CreateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileResponseBody) SetResult(v *CreateFileResponseBodyResult) *CreateFileResponseBody {
	s.Result = v
	return s
}

func (s *CreateFileResponseBody) SetSuccess(v bool) *CreateFileResponseBody {
	s.Success = &v
	return s
}

type CreateFileResponseBodyResult struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	FilePath   *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s CreateFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBodyResult) SetBranchName(v string) *CreateFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *CreateFileResponseBodyResult) SetFilePath(v string) *CreateFileResponseBodyResult {
	s.FilePath = &v
	return s
}

type CreateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateFileResponse) SetStatusCode(v int32) *CreateFileResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateMergeRequestResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMergeRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBody) SetErrorCode(v string) *CreateMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetErrorMessage(v string) *CreateMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetRequestId(v string) *CreateMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMergeRequestResponseBody) SetResult(v *CreateMergeRequestResponseBodyResult) *CreateMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *CreateMergeRequestResponseBody) SetSuccess(v bool) *CreateMergeRequestResponseBody {
	s.Success = &v
	return s
}

type CreateMergeRequestResponseBodyResult struct {
	AcceptedRevision   *string                                                 `json:"AcceptedRevision,omitempty" xml:"AcceptedRevision,omitempty"`
	AheadCommitCount   *int32                                                  `json:"AheadCommitCount,omitempty" xml:"AheadCommitCount,omitempty"`
	ApproveCheckResult *CreateMergeRequestResponseBodyResultApproveCheckResult `json:"ApproveCheckResult,omitempty" xml:"ApproveCheckResult,omitempty" type:"Struct"`
	AssigneeList       []*CreateMergeRequestResponseBodyResultAssigneeList     `json:"AssigneeList,omitempty" xml:"AssigneeList,omitempty" type:"Repeated"`
	Author             *CreateMergeRequestResponseBodyResultAuthor             `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	BehindCommitCount  *int32                                                  `json:"BehindCommitCount,omitempty" xml:"BehindCommitCount,omitempty"`
	CreatedAt          *string                                                 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description        *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                 *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	MergeError         *string                                                 `json:"MergeError,omitempty" xml:"MergeError,omitempty"`
	MergeStatus        *string                                                 `json:"MergeStatus,omitempty" xml:"MergeStatus,omitempty"`
	MergeType          *string                                                 `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	MergedRevision     *string                                                 `json:"MergedRevision,omitempty" xml:"MergedRevision,omitempty"`
	NameWithNamespace  *string                                                 `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	ProjectId          *int64                                                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SourceBranch       *string                                                 `json:"SourceBranch,omitempty" xml:"SourceBranch,omitempty"`
	State              *string                                                 `json:"State,omitempty" xml:"State,omitempty"`
	TargetBranch       *string                                                 `json:"TargetBranch,omitempty" xml:"TargetBranch,omitempty"`
	Title              *string                                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdatedAt          *string                                                 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	WebUrl             *string                                                 `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s CreateMergeRequestResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResult) SetAcceptedRevision(v string) *CreateMergeRequestResponseBodyResult {
	s.AcceptedRevision = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAheadCommitCount(v int32) *CreateMergeRequestResponseBodyResult {
	s.AheadCommitCount = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetApproveCheckResult(v *CreateMergeRequestResponseBodyResultApproveCheckResult) *CreateMergeRequestResponseBodyResult {
	s.ApproveCheckResult = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAssigneeList(v []*CreateMergeRequestResponseBodyResultAssigneeList) *CreateMergeRequestResponseBodyResult {
	s.AssigneeList = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetAuthor(v *CreateMergeRequestResponseBodyResultAuthor) *CreateMergeRequestResponseBodyResult {
	s.Author = v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetBehindCommitCount(v int32) *CreateMergeRequestResponseBodyResult {
	s.BehindCommitCount = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetCreatedAt(v string) *CreateMergeRequestResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetDescription(v string) *CreateMergeRequestResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetId(v int64) *CreateMergeRequestResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMergeError(v string) *CreateMergeRequestResponseBodyResult {
	s.MergeError = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMergeStatus(v string) *CreateMergeRequestResponseBodyResult {
	s.MergeStatus = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMergeType(v string) *CreateMergeRequestResponseBodyResult {
	s.MergeType = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetMergedRevision(v string) *CreateMergeRequestResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetNameWithNamespace(v string) *CreateMergeRequestResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetProjectId(v int64) *CreateMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetSourceBranch(v string) *CreateMergeRequestResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetState(v string) *CreateMergeRequestResponseBodyResult {
	s.State = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetTargetBranch(v string) *CreateMergeRequestResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetTitle(v string) *CreateMergeRequestResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetUpdatedAt(v string) *CreateMergeRequestResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResult) SetWebUrl(v string) *CreateMergeRequestResponseBodyResult {
	s.WebUrl = &v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResult struct {
	SatisfiedCheckResults   []*CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults   `json:"SatisfiedCheckResults,omitempty" xml:"SatisfiedCheckResults,omitempty" type:"Repeated"`
	TotalCheckResult        *string                                                                          `json:"TotalCheckResult,omitempty" xml:"TotalCheckResult,omitempty"`
	UnsatisfiedCheckResults []*CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults `json:"UnsatisfiedCheckResults,omitempty" xml:"UnsatisfiedCheckResults,omitempty" type:"Repeated"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResult) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResult) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResult) SetSatisfiedCheckResults(v []*CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) *CreateMergeRequestResponseBodyResultApproveCheckResult {
	s.SatisfiedCheckResults = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResult) SetTotalCheckResult(v string) *CreateMergeRequestResponseBodyResultApproveCheckResult {
	s.TotalCheckResult = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResult) SetUnsatisfiedCheckResults(v []*CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) *CreateMergeRequestResponseBodyResultApproveCheckResult {
	s.UnsatisfiedCheckResults = v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults struct {
	CheckName        *string                                                                                  `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckName(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckName = &v
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

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetExtraUsers(v []*CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetSatisfiedItems(v []*string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetId(v int64) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetName(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults struct {
	CheckName        *string                                                                                    `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                    `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                    `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                  `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                  `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckName(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckName = &v
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

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetExtraUsers(v []*CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetSatisfiedItems(v []*string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetId(v int64) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetName(v string) *CreateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type CreateMergeRequestResponseBodyResultAssigneeList struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultAssigneeList) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultAssigneeList) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultAssigneeList) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultAssigneeList {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAssigneeList) SetExternUserId(v string) *CreateMergeRequestResponseBodyResultAssigneeList {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAssigneeList) SetId(v string) *CreateMergeRequestResponseBodyResultAssigneeList {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAssigneeList) SetName(v string) *CreateMergeRequestResponseBodyResultAssigneeList {
	s.Name = &v
	return s
}

type CreateMergeRequestResponseBodyResultAuthor struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMergeRequestResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetAvatarUrl(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetExternUserId(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetId(v int64) *CreateMergeRequestResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestResponseBodyResultAuthor) SetName(v string) *CreateMergeRequestResponseBodyResultAuthor {
	s.Name = &v
	return s
}

type CreateMergeRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateMergeRequestResponse) SetStatusCode(v int32) *CreateMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMergeRequestResponse) SetBody(v *CreateMergeRequestResponseBody) *CreateMergeRequestResponse {
	s.Body = v
	return s
}

type CreateMergeRequestCommentRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateMergeRequestCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestCommentRequest) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestCommentRequest) SetAccessToken(v string) *CreateMergeRequestCommentRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateMergeRequestCommentRequest) SetOrganizationId(v string) *CreateMergeRequestCommentRequest {
	s.OrganizationId = &v
	return s
}

type CreateMergeRequestCommentResponseBody struct {
	ErrorCode    *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateMergeRequestCommentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMergeRequestCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestCommentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestCommentResponseBody) SetErrorCode(v string) *CreateMergeRequestCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBody) SetErrorMessage(v string) *CreateMergeRequestCommentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBody) SetRequestId(v string) *CreateMergeRequestCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBody) SetResult(v *CreateMergeRequestCommentResponseBodyResult) *CreateMergeRequestCommentResponseBody {
	s.Result = v
	return s
}

func (s *CreateMergeRequestCommentResponseBody) SetSuccess(v bool) *CreateMergeRequestCommentResponseBody {
	s.Success = &v
	return s
}

type CreateMergeRequestCommentResponseBodyResult struct {
	Author       *CreateMergeRequestCommentResponseBodyResultAuthor `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	Closed       *int32                                             `json:"Closed,omitempty" xml:"Closed,omitempty"`
	CreatedAt    *string                                            `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id           *int64                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDraft      *bool                                              `json:"IsDraft,omitempty" xml:"IsDraft,omitempty"`
	Line         *int64                                             `json:"Line,omitempty" xml:"Line,omitempty"`
	Note         *string                                            `json:"Note,omitempty" xml:"Note,omitempty"`
	OutDated     *bool                                              `json:"OutDated,omitempty" xml:"OutDated,omitempty"`
	ParentNoteId *int64                                             `json:"ParentNoteId,omitempty" xml:"ParentNoteId,omitempty"`
	Path         *string                                            `json:"Path,omitempty" xml:"Path,omitempty"`
	ProjectId    *int64                                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RangeContext *string                                            `json:"RangeContext,omitempty" xml:"RangeContext,omitempty"`
	Side         *string                                            `json:"Side,omitempty" xml:"Side,omitempty"`
	UpdatedAt    *string                                            `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s CreateMergeRequestCommentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestCommentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetAuthor(v *CreateMergeRequestCommentResponseBodyResultAuthor) *CreateMergeRequestCommentResponseBodyResult {
	s.Author = v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetClosed(v int32) *CreateMergeRequestCommentResponseBodyResult {
	s.Closed = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetCreatedAt(v string) *CreateMergeRequestCommentResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetId(v int64) *CreateMergeRequestCommentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetIsDraft(v bool) *CreateMergeRequestCommentResponseBodyResult {
	s.IsDraft = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetLine(v int64) *CreateMergeRequestCommentResponseBodyResult {
	s.Line = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetNote(v string) *CreateMergeRequestCommentResponseBodyResult {
	s.Note = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetOutDated(v bool) *CreateMergeRequestCommentResponseBodyResult {
	s.OutDated = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetParentNoteId(v int64) *CreateMergeRequestCommentResponseBodyResult {
	s.ParentNoteId = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetPath(v string) *CreateMergeRequestCommentResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetProjectId(v int64) *CreateMergeRequestCommentResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetRangeContext(v string) *CreateMergeRequestCommentResponseBodyResult {
	s.RangeContext = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetSide(v string) *CreateMergeRequestCommentResponseBodyResult {
	s.Side = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResult) SetUpdatedAt(v string) *CreateMergeRequestCommentResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

type CreateMergeRequestCommentResponseBodyResultAuthor struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMergeRequestCommentResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestCommentResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestCommentResponseBodyResultAuthor) SetAvatarUrl(v string) *CreateMergeRequestCommentResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResultAuthor) SetEmail(v string) *CreateMergeRequestCommentResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResultAuthor) SetExternUserId(v string) *CreateMergeRequestCommentResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResultAuthor) SetId(v int64) *CreateMergeRequestCommentResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *CreateMergeRequestCommentResponseBodyResultAuthor) SetName(v string) *CreateMergeRequestCommentResponseBodyResultAuthor {
	s.Name = &v
	return s
}

type CreateMergeRequestCommentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMergeRequestCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMergeRequestCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMergeRequestCommentResponse) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestCommentResponse) SetHeaders(v map[string]*string) *CreateMergeRequestCommentResponse {
	s.Headers = v
	return s
}

func (s *CreateMergeRequestCommentResponse) SetStatusCode(v int32) *CreateMergeRequestCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMergeRequestCommentResponse) SetBody(v *CreateMergeRequestCommentResponseBody) *CreateMergeRequestCommentResponse {
	s.Body = v
	return s
}

type CreateRepositoryRequest struct {
	AccessToken      *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	CreateParentPath *bool   `json:"CreateParentPath,omitempty" xml:"CreateParentPath,omitempty"`
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId        *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	Sync             *bool   `json:"Sync,omitempty" xml:"Sync,omitempty"`
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

func (s *CreateRepositoryRequest) SetSync(v bool) *CreateRepositoryRequest {
	s.Sync = &v
	return s
}

type CreateRepositoryResponseBody struct {
	ErrorCode    *int32                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateRepositoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBody) SetErrorCode(v int32) *CreateRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetErrorMessage(v string) *CreateRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetRequestId(v string) *CreateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetResult(v *CreateRepositoryResponseBodyResult) *CreateRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *CreateRepositoryResponseBody) SetSuccess(v bool) *CreateRepositoryResponseBody {
	s.Success = &v
	return s
}

type CreateRepositoryResponseBodyResult struct {
	Archive                  *bool                                        `json:"Archive,omitempty" xml:"Archive,omitempty"`
	AvatarUrl                *string                                      `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	BuildsEnableStatus       *bool                                        `json:"BuildsEnableStatus,omitempty" xml:"BuildsEnableStatus,omitempty"`
	CreatedAt                *string                                      `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CreatorId                *int64                                       `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DefaultBranch            *string                                      `json:"DefaultBranch,omitempty" xml:"DefaultBranch,omitempty"`
	DemoProjectStatus        *bool                                        `json:"DemoProjectStatus,omitempty" xml:"DemoProjectStatus,omitempty"`
	Description              *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpUrlToRepo            *string                                      `json:"HttpUrlToRepo,omitempty" xml:"HttpUrlToRepo,omitempty"`
	Id                       *int64                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	IssuesEnableStatus       *bool                                        `json:"IssuesEnableStatus,omitempty" xml:"IssuesEnableStatus,omitempty"`
	LastActivityAt           *string                                      `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	MergeRequestEnableStatus *bool                                        `json:"MergeRequestEnableStatus,omitempty" xml:"MergeRequestEnableStatus,omitempty"`
	Name                     *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace        *string                                      `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	Namespace                *CreateRepositoryResponseBodyResultNamespace `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Struct"`
	Path                     *string                                      `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace        *string                                      `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Public                   *bool                                        `json:"Public,omitempty" xml:"Public,omitempty"`
	SnippetsEnableStatus     *bool                                        `json:"SnippetsEnableStatus,omitempty" xml:"SnippetsEnableStatus,omitempty"`
	SshUrlToRepo             *string                                      `json:"SshUrlToRepo,omitempty" xml:"SshUrlToRepo,omitempty"`
	TagList                  []*string                                    `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	VisibilityLevel          *string                                      `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WebUrl                   *string                                      `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
	WikiEnableStatus         *bool                                        `json:"WikiEnableStatus,omitempty" xml:"WikiEnableStatus,omitempty"`
}

func (s CreateRepositoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBodyResult) SetArchive(v bool) *CreateRepositoryResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetAvatarUrl(v string) *CreateRepositoryResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetBuildsEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.BuildsEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetCreatedAt(v string) *CreateRepositoryResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetCreatorId(v int64) *CreateRepositoryResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDefaultBranch(v string) *CreateRepositoryResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDemoProjectStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.DemoProjectStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDescription(v string) *CreateRepositoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetHttpUrlToRepo(v string) *CreateRepositoryResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetId(v int64) *CreateRepositoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetIssuesEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.IssuesEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetLastActivityAt(v string) *CreateRepositoryResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetMergeRequestEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.MergeRequestEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetName(v string) *CreateRepositoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetNameWithNamespace(v string) *CreateRepositoryResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetNamespace(v *CreateRepositoryResponseBodyResultNamespace) *CreateRepositoryResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPath(v string) *CreateRepositoryResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPathWithNamespace(v string) *CreateRepositoryResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPublic(v bool) *CreateRepositoryResponseBodyResult {
	s.Public = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetSnippetsEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.SnippetsEnableStatus = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetSshUrlToRepo(v string) *CreateRepositoryResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetTagList(v []*string) *CreateRepositoryResponseBodyResult {
	s.TagList = v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetVisibilityLevel(v string) *CreateRepositoryResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetWebUrl(v string) *CreateRepositoryResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetWikiEnableStatus(v bool) *CreateRepositoryResponseBodyResult {
	s.WikiEnableStatus = &v
	return s
}

type CreateRepositoryResponseBodyResultNamespace struct {
	Avatar          *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	CreatedAt       *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Public          *bool   `json:"Public,omitempty" xml:"Public,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdatedAt       *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	VisibilityLevel *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
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

func (s *CreateRepositoryResponseBodyResultNamespace) SetCreatedAt(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetDescription(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetId(v int64) *CreateRepositoryResponseBodyResultNamespace {
	s.Id = &v
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

func (s *CreateRepositoryResponseBodyResultNamespace) SetPath(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetPublic(v bool) *CreateRepositoryResponseBodyResultNamespace {
	s.Public = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetState(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.State = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetUpdatedAt(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetVisibilityLevel(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

type CreateRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRepositoryResponse) SetStatusCode(v int32) *CreateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryResponse) SetBody(v *CreateRepositoryResponseBody) *CreateRepositoryResponse {
	s.Body = v
	return s
}

type CreateRepositoryDeployKeyRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s CreateRepositoryDeployKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryDeployKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryDeployKeyRequest) SetAccessToken(v string) *CreateRepositoryDeployKeyRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateRepositoryDeployKeyRequest) SetOrganizationId(v string) *CreateRepositoryDeployKeyRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateRepositoryDeployKeyRequest) SetSubUserId(v string) *CreateRepositoryDeployKeyRequest {
	s.SubUserId = &v
	return s
}

type CreateRepositoryDeployKeyResponseBody struct {
	ErrorCode    *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateRepositoryDeployKeyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRepositoryDeployKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryDeployKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryDeployKeyResponseBody) SetErrorCode(v string) *CreateRepositoryDeployKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryDeployKeyResponseBody) SetErrorMessage(v string) *CreateRepositoryDeployKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryDeployKeyResponseBody) SetRequestId(v string) *CreateRepositoryDeployKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryDeployKeyResponseBody) SetResult(v *CreateRepositoryDeployKeyResponseBodyResult) *CreateRepositoryDeployKeyResponseBody {
	s.Result = v
	return s
}

func (s *CreateRepositoryDeployKeyResponseBody) SetSuccess(v bool) *CreateRepositoryDeployKeyResponseBody {
	s.Success = &v
	return s
}

type CreateRepositoryDeployKeyResponseBodyResult struct {
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	FingerPrint *string `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateRepositoryDeployKeyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryDeployKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryDeployKeyResponseBodyResult) SetCreatedAt(v string) *CreateRepositoryDeployKeyResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryDeployKeyResponseBodyResult) SetFingerPrint(v string) *CreateRepositoryDeployKeyResponseBodyResult {
	s.FingerPrint = &v
	return s
}

func (s *CreateRepositoryDeployKeyResponseBodyResult) SetId(v int64) *CreateRepositoryDeployKeyResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateRepositoryDeployKeyResponseBodyResult) SetKey(v string) *CreateRepositoryDeployKeyResponseBodyResult {
	s.Key = &v
	return s
}

func (s *CreateRepositoryDeployKeyResponseBodyResult) SetTitle(v string) *CreateRepositoryDeployKeyResponseBodyResult {
	s.Title = &v
	return s
}

type CreateRepositoryDeployKeyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRepositoryDeployKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRepositoryDeployKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryDeployKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryDeployKeyResponse) SetHeaders(v map[string]*string) *CreateRepositoryDeployKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryDeployKeyResponse) SetStatusCode(v int32) *CreateRepositoryDeployKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryDeployKeyResponse) SetBody(v *CreateRepositoryDeployKeyResponseBody) *CreateRepositoryDeployKeyResponse {
	s.Body = v
	return s
}

type CreateRepositoryGroupRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
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

type CreateRepositoryGroupResponseBody struct {
	ErrorCode    *int32                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateRepositoryGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRepositoryGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupResponseBody) SetErrorCode(v int32) *CreateRepositoryGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetErrorMessage(v string) *CreateRepositoryGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetRequestId(v string) *CreateRepositoryGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetResult(v *CreateRepositoryGroupResponseBodyResult) *CreateRepositoryGroupResponseBody {
	s.Result = v
	return s
}

func (s *CreateRepositoryGroupResponseBody) SetSuccess(v bool) *CreateRepositoryGroupResponseBody {
	s.Success = &v
	return s
}

type CreateRepositoryGroupResponseBodyResult struct {
	AvatarUrl         *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentId          *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s CreateRepositoryGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupResponseBodyResult) SetAvatarUrl(v string) *CreateRepositoryGroupResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetDescription(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetId(v int64) *CreateRepositoryGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetName(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetNameWithNamespace(v string) *CreateRepositoryGroupResponseBodyResult {
	s.NameWithNamespace = &v
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

func (s *CreateRepositoryGroupResponseBodyResult) SetPath(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetPathWithNamespace(v string) *CreateRepositoryGroupResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetType(v string) *CreateRepositoryGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetVisibilityLevel(v string) *CreateRepositoryGroupResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryGroupResponseBodyResult) SetWebUrl(v string) *CreateRepositoryGroupResponseBodyResult {
	s.WebUrl = &v
	return s
}

type CreateRepositoryGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRepositoryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRepositoryGroupResponse) SetStatusCode(v int32) *CreateRepositoryGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryGroupResponse) SetBody(v *CreateRepositoryGroupResponseBody) *CreateRepositoryGroupResponse {
	s.Body = v
	return s
}

type CreateRepositoryProtectedBranchRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateRepositoryProtectedBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchRequest) SetAccessToken(v string) *CreateRepositoryProtectedBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateRepositoryProtectedBranchRequest) SetOrganizationId(v string) *CreateRepositoryProtectedBranchRequest {
	s.OrganizationId = &v
	return s
}

type CreateRepositoryProtectedBranchResponseBody struct {
	ErrorCode    *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateRepositoryProtectedBranchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRepositoryProtectedBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponseBody) SetErrorCode(v string) *CreateRepositoryProtectedBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBody) SetErrorMessage(v string) *CreateRepositoryProtectedBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBody) SetRequestId(v string) *CreateRepositoryProtectedBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBody) SetResult(v *CreateRepositoryProtectedBranchResponseBodyResult) *CreateRepositoryProtectedBranchResponseBody {
	s.Result = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBody) SetSuccess(v bool) *CreateRepositoryProtectedBranchResponseBody {
	s.Success = &v
	return s
}

type CreateRepositoryProtectedBranchResponseBodyResult struct {
	AllowMergeRoles     []*int32                                                              `json:"AllowMergeRoles,omitempty" xml:"AllowMergeRoles,omitempty" type:"Repeated"`
	AllowPushRoles      []*int32                                                              `json:"AllowPushRoles,omitempty" xml:"AllowPushRoles,omitempty" type:"Repeated"`
	Branch              *string                                                               `json:"Branch,omitempty" xml:"Branch,omitempty"`
	Id                  *int64                                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	MergeRequestSetting *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting `json:"MergeRequestSetting,omitempty" xml:"MergeRequestSetting,omitempty" type:"Struct"`
	TestSetting         *CreateRepositoryProtectedBranchResponseBodyResultTestSetting         `json:"TestSetting,omitempty" xml:"TestSetting,omitempty" type:"Struct"`
}

func (s CreateRepositoryProtectedBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponseBodyResult) SetAllowMergeRoles(v []*int32) *CreateRepositoryProtectedBranchResponseBodyResult {
	s.AllowMergeRoles = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResult) SetAllowPushRoles(v []*int32) *CreateRepositoryProtectedBranchResponseBodyResult {
	s.AllowPushRoles = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResult) SetBranch(v string) *CreateRepositoryProtectedBranchResponseBodyResult {
	s.Branch = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResult) SetId(v int64) *CreateRepositoryProtectedBranchResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResult) SetMergeRequestSetting(v *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) *CreateRepositoryProtectedBranchResponseBodyResult {
	s.MergeRequestSetting = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResult) SetTestSetting(v *CreateRepositoryProtectedBranchResponseBodyResultTestSetting) *CreateRepositoryProtectedBranchResponseBodyResult {
	s.TestSetting = v
	return s
}

type CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting struct {
	AllowMergeRequestRoles       []*int32  `json:"AllowMergeRequestRoles,omitempty" xml:"AllowMergeRequestRoles,omitempty" type:"Repeated"`
	AllowSelfApproval            *bool     `json:"AllowSelfApproval,omitempty" xml:"AllowSelfApproval,omitempty"`
	DefaultAssignees             []*string `json:"DefaultAssignees,omitempty" xml:"DefaultAssignees,omitempty" type:"Repeated"`
	IsRequireDiscussionProcessed *bool     `json:"IsRequireDiscussionProcessed,omitempty" xml:"IsRequireDiscussionProcessed,omitempty"`
	IsResetApprovalWhenNewPush   *bool     `json:"IsResetApprovalWhenNewPush,omitempty" xml:"IsResetApprovalWhenNewPush,omitempty"`
	MergeRequestMode             *string   `json:"MergeRequestMode,omitempty" xml:"MergeRequestMode,omitempty"`
	MinimualApproval             *int32    `json:"MinimualApproval,omitempty" xml:"MinimualApproval,omitempty"`
	Required                     *bool     `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetAllowSelfApproval(v bool) *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.AllowSelfApproval = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetDefaultAssignees(v []*string) *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetIsResetApprovalWhenNewPush(v bool) *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.IsResetApprovalWhenNewPush = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetMergeRequestMode(v string) *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.MergeRequestMode = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetMinimualApproval(v int32) *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.MinimualApproval = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetRequired(v bool) *CreateRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.Required = &v
	return s
}

type CreateRepositoryProtectedBranchResponseBodyResultTestSetting struct {
	CheckConfig               *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig               `json:"CheckConfig,omitempty" xml:"CheckConfig,omitempty" type:"Struct"`
	CodingGuidelinesDetection *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection `json:"CodingGuidelinesDetection,omitempty" xml:"CodingGuidelinesDetection,omitempty" type:"Struct"`
	Required                  *bool                                                                                  `json:"Required,omitempty" xml:"Required,omitempty"`
	SensitiveInfoDetection    *CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection    `json:"SensitiveInfoDetection,omitempty" xml:"SensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSetting) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSetting) SetCheckConfig(v *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig) *CreateRepositoryProtectedBranchResponseBodyResultTestSetting {
	s.CheckConfig = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSetting) SetCodingGuidelinesDetection(v *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) *CreateRepositoryProtectedBranchResponseBodyResultTestSetting {
	s.CodingGuidelinesDetection = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSetting) SetRequired(v bool) *CreateRepositoryProtectedBranchResponseBodyResultTestSetting {
	s.Required = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSetting) SetSensitiveInfoDetection(v *CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) *CreateRepositoryProtectedBranchResponseBodyResultTestSetting {
	s.SensitiveInfoDetection = v
	return s
}

type CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig struct {
	CheckItems []*CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems `json:"CheckItems,omitempty" xml:"CheckItems,omitempty" type:"Repeated"`
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig) SetCheckItems(v []*CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig {
	s.CheckItems = v
	return s
}

type CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) SetName(v string) *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems {
	s.Name = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) SetRequired(v bool) *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems {
	s.Required = &v
	return s
}

type CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection struct {
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) SetEnabled(v bool) *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) SetMessage(v string) *CreateRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection {
	s.Message = &v
	return s
}

type CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection struct {
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) SetEnabled(v bool) *CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) SetMessage(v string) *CreateRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection {
	s.Message = &v
	return s
}

type CreateRepositoryProtectedBranchResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRepositoryProtectedBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRepositoryProtectedBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryProtectedBranchResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryProtectedBranchResponse) SetHeaders(v map[string]*string) *CreateRepositoryProtectedBranchResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryProtectedBranchResponse) SetStatusCode(v int32) *CreateRepositoryProtectedBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryProtectedBranchResponse) SetBody(v *CreateRepositoryProtectedBranchResponseBody) *CreateRepositoryProtectedBranchResponse {
	s.Body = v
	return s
}

type CreateSshKeyRequest struct {
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
}

func (s CreateSshKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSshKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateSshKeyRequest) SetAccessToken(v string) *CreateSshKeyRequest {
	s.AccessToken = &v
	return s
}

type CreateSshKeyResponseBody struct {
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateSshKeyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSshKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSshKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponseBody) SetErrorCode(v string) *CreateSshKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetErrorMessage(v string) *CreateSshKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetRequestId(v string) *CreateSshKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSshKeyResponseBody) SetResult(v *CreateSshKeyResponseBodyResult) *CreateSshKeyResponseBody {
	s.Result = v
	return s
}

func (s *CreateSshKeyResponseBody) SetSuccess(v bool) *CreateSshKeyResponseBody {
	s.Success = &v
	return s
}

type CreateSshKeyResponseBodyResult struct {
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	FingerPrint *string `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyScope    *string `json:"KeyScope,omitempty" xml:"KeyScope,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateSshKeyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateSshKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponseBodyResult) SetCreatedAt(v string) *CreateSshKeyResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateSshKeyResponseBodyResult) SetFingerPrint(v string) *CreateSshKeyResponseBodyResult {
	s.FingerPrint = &v
	return s
}

func (s *CreateSshKeyResponseBodyResult) SetId(v int64) *CreateSshKeyResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateSshKeyResponseBodyResult) SetKey(v string) *CreateSshKeyResponseBodyResult {
	s.Key = &v
	return s
}

func (s *CreateSshKeyResponseBodyResult) SetKeyScope(v string) *CreateSshKeyResponseBodyResult {
	s.KeyScope = &v
	return s
}

func (s *CreateSshKeyResponseBodyResult) SetTitle(v string) *CreateSshKeyResponseBodyResult {
	s.Title = &v
	return s
}

type CreateSshKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSshKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSshKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSshKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponse) SetHeaders(v map[string]*string) *CreateSshKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateSshKeyResponse) SetStatusCode(v int32) *CreateSshKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSshKeyResponse) SetBody(v *CreateSshKeyResponseBody) *CreateSshKeyResponse {
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
	ErrorCode    *string                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *CreateTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) SetErrorCode(v string) *CreateTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTagResponseBody) SetErrorMessage(v string) *CreateTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagResponseBody) SetResult(v *CreateTagResponseBodyResult) *CreateTagResponseBody {
	s.Result = v
	return s
}

func (s *CreateTagResponseBody) SetSuccess(v bool) *CreateTagResponseBody {
	s.Success = &v
	return s
}

type CreateTagResponseBodyResult struct {
	CommitInfo *CreateTagResponseBodyResultCommitInfo `json:"CommitInfo,omitempty" xml:"CommitInfo,omitempty" type:"Struct"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Name       *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Release    *CreateTagResponseBodyResultRelease    `json:"Release,omitempty" xml:"Release,omitempty" type:"Struct"`
}

func (s CreateTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyResult) SetCommitInfo(v *CreateTagResponseBodyResultCommitInfo) *CreateTagResponseBodyResult {
	s.CommitInfo = v
	return s
}

func (s *CreateTagResponseBodyResult) SetMessage(v string) *CreateTagResponseBodyResult {
	s.Message = &v
	return s
}

func (s *CreateTagResponseBodyResult) SetName(v string) *CreateTagResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateTagResponseBodyResult) SetRelease(v *CreateTagResponseBodyResultRelease) *CreateTagResponseBodyResult {
	s.Release = v
	return s
}

type CreateTagResponseBodyResultCommitInfo struct {
	AuthorEmail    *string   `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string   `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	AuthoredDate   *string   `json:"AuthoredDate,omitempty" xml:"AuthoredDate,omitempty"`
	CommittedDate  *string   `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string   `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string   `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string   `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTagResponseBodyResultCommitInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResponseBodyResultCommitInfo) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyResultCommitInfo) SetAuthorEmail(v string) *CreateTagResponseBodyResultCommitInfo {
	s.AuthorEmail = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetAuthorName(v string) *CreateTagResponseBodyResultCommitInfo {
	s.AuthorName = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetAuthoredDate(v string) *CreateTagResponseBodyResultCommitInfo {
	s.AuthoredDate = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetCommittedDate(v string) *CreateTagResponseBodyResultCommitInfo {
	s.CommittedDate = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetCommitterEmail(v string) *CreateTagResponseBodyResultCommitInfo {
	s.CommitterEmail = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetCommitterName(v string) *CreateTagResponseBodyResultCommitInfo {
	s.CommitterName = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetCreatedAt(v string) *CreateTagResponseBodyResultCommitInfo {
	s.CreatedAt = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetId(v string) *CreateTagResponseBodyResultCommitInfo {
	s.Id = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetMessage(v string) *CreateTagResponseBodyResultCommitInfo {
	s.Message = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetParentIds(v []*string) *CreateTagResponseBodyResultCommitInfo {
	s.ParentIds = v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetShortId(v string) *CreateTagResponseBodyResultCommitInfo {
	s.ShortId = &v
	return s
}

func (s *CreateTagResponseBodyResultCommitInfo) SetTitle(v string) *CreateTagResponseBodyResultCommitInfo {
	s.Title = &v
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

type CreateTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTagResponse) SetStatusCode(v int32) *CreateTagResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteBranchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBranchResponseBody) SetErrorCode(v string) *DeleteBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteBranchResponseBody) SetErrorMessage(v string) *DeleteBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteBranchResponseBody) SetRequestId(v string) *DeleteBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBranchResponseBody) SetResult(v *DeleteBranchResponseBodyResult) *DeleteBranchResponseBody {
	s.Result = v
	return s
}

func (s *DeleteBranchResponseBody) SetSuccess(v bool) *DeleteBranchResponseBody {
	s.Success = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteBranchResponse) SetStatusCode(v int32) *DeleteBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBranchResponse) SetBody(v *DeleteBranchResponseBody) *DeleteBranchResponse {
	s.Body = v
	return s
}

type DeleteFileRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	BranchName     *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	CommitMessage  *string `json:"CommitMessage,omitempty" xml:"CommitMessage,omitempty"`
	FilePath       *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
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

func (s *DeleteFileRequest) SetBranchName(v string) *DeleteFileRequest {
	s.BranchName = &v
	return s
}

func (s *DeleteFileRequest) SetCommitMessage(v string) *DeleteFileRequest {
	s.CommitMessage = &v
	return s
}

func (s *DeleteFileRequest) SetFilePath(v string) *DeleteFileRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteFileRequest) SetOrganizationId(v string) *DeleteFileRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteFileRequest) SetSubUserId(v string) *DeleteFileRequest {
	s.SubUserId = &v
	return s
}

type DeleteFileResponseBody struct {
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteFileResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetErrorCode(v string) *DeleteFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFileResponseBody) SetErrorMessage(v string) *DeleteFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetResult(v *DeleteFileResponseBodyResult) *DeleteFileResponseBody {
	s.Result = v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

type DeleteFileResponseBodyResult struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	FilePath   *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s DeleteFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBodyResult) SetBranchName(v string) *DeleteFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *DeleteFileResponseBodyResult) SetFilePath(v string) *DeleteFileResponseBodyResult {
	s.FilePath = &v
	return s
}

type DeleteFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteFileResponse) SetStatusCode(v int32) *DeleteFileResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberResponseBody) SetErrorCode(v string) *DeleteGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetErrorMessage(v string) *DeleteGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetRequestId(v string) *DeleteGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetResult(v *DeleteGroupMemberResponseBodyResult) *DeleteGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *DeleteGroupMemberResponseBody) SetSuccess(v bool) *DeleteGroupMemberResponseBody {
	s.Success = &v
	return s
}

type DeleteGroupMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DeleteGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberResponseBodyResult) SetAccessLevel(v int32) *DeleteGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetAvatarUrl(v string) *DeleteGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetEmail(v string) *DeleteGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetExternUserId(v string) *DeleteGroupMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetId(v int64) *DeleteGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DeleteGroupMemberResponseBodyResult) SetState(v string) *DeleteGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

type DeleteGroupMemberResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteGroupMemberResponse) SetStatusCode(v int32) *DeleteGroupMemberResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteRepositoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBody) SetErrorCode(v string) *DeleteRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetErrorMessage(v string) *DeleteRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetRequestId(v string) *DeleteRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetResult(v *DeleteRepositoryResponseBodyResult) *DeleteRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryResponseBody) SetSuccess(v bool) *DeleteRepositoryResponseBody {
	s.Success = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRepositoryResponse) SetStatusCode(v int32) *DeleteRepositoryResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteRepositoryGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRepositoryGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryGroupResponseBody) SetErrorCode(v string) *DeleteRepositoryGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetErrorMessage(v string) *DeleteRepositoryGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetRequestId(v string) *DeleteRepositoryGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetResult(v *DeleteRepositoryGroupResponseBodyResult) *DeleteRepositoryGroupResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryGroupResponseBody) SetSuccess(v bool) *DeleteRepositoryGroupResponseBody {
	s.Success = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRepositoryGroupResponse) SetStatusCode(v int32) *DeleteRepositoryGroupResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberResponseBody) SetErrorCode(v string) *DeleteRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetErrorMessage(v string) *DeleteRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetRequestId(v string) *DeleteRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetResult(v *DeleteRepositoryMemberResponseBodyResult) *DeleteRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryMemberResponseBody) SetSuccess(v bool) *DeleteRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

type DeleteRepositoryMemberResponseBodyResult struct {
	AccessLevel       *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NotificationLevel *int32  `json:"NotificationLevel,omitempty" xml:"NotificationLevel,omitempty"`
	SourceId          *int64  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType        *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	UpdatedAt         *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	UserId            *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *DeleteRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetCreatedAt(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetMessage(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetNotificationLevel(v int32) *DeleteRepositoryMemberResponseBodyResult {
	s.NotificationLevel = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetSourceId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetSourceType(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetUpdatedAt(v string) *DeleteRepositoryMemberResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *DeleteRepositoryMemberResponseBodyResult) SetUserId(v int64) *DeleteRepositoryMemberResponseBodyResult {
	s.UserId = &v
	return s
}

type DeleteRepositoryMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRepositoryMemberResponse) SetStatusCode(v int32) *DeleteRepositoryMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryMemberResponse) SetBody(v *DeleteRepositoryMemberResponseBody) *DeleteRepositoryMemberResponse {
	s.Body = v
	return s
}

type DeleteRepositoryMemberWithExternUidRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	ExternUserId   *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteRepositoryMemberWithExternUidRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberWithExternUidRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberWithExternUidRequest) SetAccessToken(v string) *DeleteRepositoryMemberWithExternUidRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidRequest) SetExternUserId(v string) *DeleteRepositoryMemberWithExternUidRequest {
	s.ExternUserId = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidRequest) SetOrganizationId(v string) *DeleteRepositoryMemberWithExternUidRequest {
	s.OrganizationId = &v
	return s
}

type DeleteRepositoryMemberWithExternUidResponseBody struct {
	ErrorCode    *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteRepositoryMemberWithExternUidResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRepositoryMemberWithExternUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberWithExternUidResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberWithExternUidResponseBody) SetErrorCode(v string) *DeleteRepositoryMemberWithExternUidResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBody) SetErrorMessage(v string) *DeleteRepositoryMemberWithExternUidResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBody) SetRequestId(v string) *DeleteRepositoryMemberWithExternUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBody) SetResult(v *DeleteRepositoryMemberWithExternUidResponseBodyResult) *DeleteRepositoryMemberWithExternUidResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBody) SetSuccess(v bool) *DeleteRepositoryMemberWithExternUidResponseBody {
	s.Success = &v
	return s
}

type DeleteRepositoryMemberWithExternUidResponseBodyResult struct {
	AccessLevel *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SourceId    *int64  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType  *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	UpdatedAt   *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	UserId      *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteRepositoryMemberWithExternUidResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberWithExternUidResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberWithExternUidResponseBodyResult) SetAccessLevel(v int32) *DeleteRepositoryMemberWithExternUidResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBodyResult) SetCreatedAt(v string) *DeleteRepositoryMemberWithExternUidResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBodyResult) SetId(v int64) *DeleteRepositoryMemberWithExternUidResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBodyResult) SetSourceId(v int64) *DeleteRepositoryMemberWithExternUidResponseBodyResult {
	s.SourceId = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBodyResult) SetSourceType(v string) *DeleteRepositoryMemberWithExternUidResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBodyResult) SetUpdatedAt(v string) *DeleteRepositoryMemberWithExternUidResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponseBodyResult) SetUserId(v int64) *DeleteRepositoryMemberWithExternUidResponseBodyResult {
	s.UserId = &v
	return s
}

type DeleteRepositoryMemberWithExternUidResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryMemberWithExternUidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryMemberWithExternUidResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryMemberWithExternUidResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberWithExternUidResponse) SetHeaders(v map[string]*string) *DeleteRepositoryMemberWithExternUidResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponse) SetStatusCode(v int32) *DeleteRepositoryMemberWithExternUidResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryMemberWithExternUidResponse) SetBody(v *DeleteRepositoryMemberWithExternUidResponseBody) *DeleteRepositoryMemberWithExternUidResponse {
	s.Body = v
	return s
}

type DeleteRepositoryProtectedBranchRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteRepositoryProtectedBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryProtectedBranchRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryProtectedBranchRequest) SetAccessToken(v string) *DeleteRepositoryProtectedBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryProtectedBranchRequest) SetOrganizationId(v string) *DeleteRepositoryProtectedBranchRequest {
	s.OrganizationId = &v
	return s
}

type DeleteRepositoryProtectedBranchResponseBody struct {
	ErrorCode    *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteRepositoryProtectedBranchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRepositoryProtectedBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryProtectedBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryProtectedBranchResponseBody) SetErrorCode(v string) *DeleteRepositoryProtectedBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryProtectedBranchResponseBody) SetErrorMessage(v string) *DeleteRepositoryProtectedBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryProtectedBranchResponseBody) SetRequestId(v string) *DeleteRepositoryProtectedBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryProtectedBranchResponseBody) SetResult(v *DeleteRepositoryProtectedBranchResponseBodyResult) *DeleteRepositoryProtectedBranchResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryProtectedBranchResponseBody) SetSuccess(v bool) *DeleteRepositoryProtectedBranchResponseBody {
	s.Success = &v
	return s
}

type DeleteRepositoryProtectedBranchResponseBodyResult struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteRepositoryProtectedBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryProtectedBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryProtectedBranchResponseBodyResult) SetResult(v bool) *DeleteRepositoryProtectedBranchResponseBodyResult {
	s.Result = &v
	return s
}

type DeleteRepositoryProtectedBranchResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryProtectedBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryProtectedBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryProtectedBranchResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryProtectedBranchResponse) SetHeaders(v map[string]*string) *DeleteRepositoryProtectedBranchResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryProtectedBranchResponse) SetStatusCode(v int32) *DeleteRepositoryProtectedBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryProtectedBranchResponse) SetBody(v *DeleteRepositoryProtectedBranchResponseBody) *DeleteRepositoryProtectedBranchResponse {
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
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteRepositoryTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRepositoryTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagResponseBody) SetErrorCode(v string) *DeleteRepositoryTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryTagResponseBody) SetErrorMessage(v string) *DeleteRepositoryTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryTagResponseBody) SetRequestId(v string) *DeleteRepositoryTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryTagResponseBody) SetResult(v *DeleteRepositoryTagResponseBodyResult) *DeleteRepositoryTagResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryTagResponseBody) SetSuccess(v bool) *DeleteRepositoryTagResponseBody {
	s.Success = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRepositoryTagResponse) SetStatusCode(v int32) *DeleteRepositoryTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryTagResponse) SetBody(v *DeleteRepositoryTagResponseBody) *DeleteRepositoryTagResponse {
	s.Body = v
	return s
}

type DeleteRepositoryTagV2Request struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	TagName        *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DeleteRepositoryTagV2Request) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagV2Request) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagV2Request) SetAccessToken(v string) *DeleteRepositoryTagV2Request {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryTagV2Request) SetOrganizationId(v string) *DeleteRepositoryTagV2Request {
	s.OrganizationId = &v
	return s
}

func (s *DeleteRepositoryTagV2Request) SetTagName(v string) *DeleteRepositoryTagV2Request {
	s.TagName = &v
	return s
}

type DeleteRepositoryTagV2ResponseBody struct {
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteRepositoryTagV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRepositoryTagV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagV2ResponseBody) SetErrorCode(v string) *DeleteRepositoryTagV2ResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryTagV2ResponseBody) SetErrorMessage(v string) *DeleteRepositoryTagV2ResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryTagV2ResponseBody) SetRequestId(v string) *DeleteRepositoryTagV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryTagV2ResponseBody) SetResult(v *DeleteRepositoryTagV2ResponseBodyResult) *DeleteRepositoryTagV2ResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryTagV2ResponseBody) SetSuccess(v bool) *DeleteRepositoryTagV2ResponseBody {
	s.Success = &v
	return s
}

type DeleteRepositoryTagV2ResponseBodyResult struct {
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DeleteRepositoryTagV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagV2ResponseBodyResult) SetTagName(v string) *DeleteRepositoryTagV2ResponseBodyResult {
	s.TagName = &v
	return s
}

type DeleteRepositoryTagV2Response struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryTagV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRepositoryTagV2Response) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryTagV2Response) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryTagV2Response) SetHeaders(v map[string]*string) *DeleteRepositoryTagV2Response {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryTagV2Response) SetStatusCode(v int32) *DeleteRepositoryTagV2Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryTagV2Response) SetBody(v *DeleteRepositoryTagV2ResponseBody) *DeleteRepositoryTagV2Response {
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
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *DeleteRepositoryWebhookResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRepositoryWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookResponseBody) SetErrorCode(v string) *DeleteRepositoryWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetErrorMessage(v string) *DeleteRepositoryWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetRequestId(v string) *DeleteRepositoryWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetResult(v *DeleteRepositoryWebhookResponseBodyResult) *DeleteRepositoryWebhookResponseBody {
	s.Result = v
	return s
}

func (s *DeleteRepositoryWebhookResponseBody) SetSuccess(v bool) *DeleteRepositoryWebhookResponseBody {
	s.Success = &v
	return s
}

type DeleteRepositoryWebhookResponseBodyResult struct {
	CreatedAt             *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableSslVerification *bool   `json:"EnableSslVerification,omitempty" xml:"EnableSslVerification,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LastTestResult        *string `json:"LastTestResult,omitempty" xml:"LastTestResult,omitempty"`
	MergeRequestsEvents   *bool   `json:"MergeRequestsEvents,omitempty" xml:"MergeRequestsEvents,omitempty"`
	NoteEvents            *bool   `json:"NoteEvents,omitempty" xml:"NoteEvents,omitempty"`
	ProjectId             *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PushEvents            *bool   `json:"PushEvents,omitempty" xml:"PushEvents,omitempty"`
	SecretToken           *string `json:"SecretToken,omitempty" xml:"SecretToken,omitempty"`
	TagPushEvents         *bool   `json:"TagPushEvents,omitempty" xml:"TagPushEvents,omitempty"`
	Url                   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DeleteRepositoryWebhookResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetCreatedAt(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetDescription(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.Description = &v
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

func (s *DeleteRepositoryWebhookResponseBodyResult) SetLastTestResult(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetNoteEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetProjectId(v int64) *DeleteRepositoryWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetPushEvents(v bool) *DeleteRepositoryWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *DeleteRepositoryWebhookResponseBodyResult) SetSecretToken(v string) *DeleteRepositoryWebhookResponseBodyResult {
	s.SecretToken = &v
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

type DeleteRepositoryWebhookResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRepositoryWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRepositoryWebhookResponse) SetStatusCode(v int32) *DeleteRepositoryWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryWebhookResponse) SetBody(v *DeleteRepositoryWebhookResponseBody) *DeleteRepositoryWebhookResponse {
	s.Body = v
	return s
}

type EnableRepositoryDeployKeyRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s EnableRepositoryDeployKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRepositoryDeployKeyRequest) GoString() string {
	return s.String()
}

func (s *EnableRepositoryDeployKeyRequest) SetAccessToken(v string) *EnableRepositoryDeployKeyRequest {
	s.AccessToken = &v
	return s
}

func (s *EnableRepositoryDeployKeyRequest) SetOrganizationId(v string) *EnableRepositoryDeployKeyRequest {
	s.OrganizationId = &v
	return s
}

func (s *EnableRepositoryDeployKeyRequest) SetSubUserId(v string) *EnableRepositoryDeployKeyRequest {
	s.SubUserId = &v
	return s
}

type EnableRepositoryDeployKeyResponseBody struct {
	ErrorCode    *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *EnableRepositoryDeployKeyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableRepositoryDeployKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRepositoryDeployKeyResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRepositoryDeployKeyResponseBody) SetErrorCode(v string) *EnableRepositoryDeployKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *EnableRepositoryDeployKeyResponseBody) SetErrorMessage(v string) *EnableRepositoryDeployKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *EnableRepositoryDeployKeyResponseBody) SetRequestId(v string) *EnableRepositoryDeployKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableRepositoryDeployKeyResponseBody) SetResult(v *EnableRepositoryDeployKeyResponseBodyResult) *EnableRepositoryDeployKeyResponseBody {
	s.Result = v
	return s
}

func (s *EnableRepositoryDeployKeyResponseBody) SetSuccess(v bool) *EnableRepositoryDeployKeyResponseBody {
	s.Success = &v
	return s
}

type EnableRepositoryDeployKeyResponseBodyResult struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s EnableRepositoryDeployKeyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EnableRepositoryDeployKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EnableRepositoryDeployKeyResponseBodyResult) SetResult(v bool) *EnableRepositoryDeployKeyResponseBodyResult {
	s.Result = &v
	return s
}

type EnableRepositoryDeployKeyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableRepositoryDeployKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableRepositoryDeployKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRepositoryDeployKeyResponse) GoString() string {
	return s.String()
}

func (s *EnableRepositoryDeployKeyResponse) SetHeaders(v map[string]*string) *EnableRepositoryDeployKeyResponse {
	s.Headers = v
	return s
}

func (s *EnableRepositoryDeployKeyResponse) SetStatusCode(v int32) *EnableRepositoryDeployKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableRepositoryDeployKeyResponse) SetBody(v *EnableRepositoryDeployKeyResponseBody) *EnableRepositoryDeployKeyResponse {
	s.Body = v
	return s
}

type GetBranchInfoRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	BranchName     *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
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

func (s *GetBranchInfoRequest) SetBranchName(v string) *GetBranchInfoRequest {
	s.BranchName = &v
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

type GetBranchInfoResponseBody struct {
	ErrorCode    *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetBranchInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBranchInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBody) SetErrorCode(v string) *GetBranchInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetErrorMessage(v string) *GetBranchInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetRequestId(v string) *GetBranchInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBranchInfoResponseBody) SetResult(v *GetBranchInfoResponseBodyResult) *GetBranchInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetBranchInfoResponseBody) SetSuccess(v bool) *GetBranchInfoResponseBody {
	s.Success = &v
	return s
}

type GetBranchInfoResponseBodyResult struct {
	BranchName      *string                                    `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	CommitInfo      *GetBranchInfoResponseBodyResultCommitInfo `json:"CommitInfo,omitempty" xml:"CommitInfo,omitempty" type:"Struct"`
	ProtectedBranch *bool                                      `json:"ProtectedBranch,omitempty" xml:"ProtectedBranch,omitempty"`
}

func (s GetBranchInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResult) SetBranchName(v string) *GetBranchInfoResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResult) SetCommitInfo(v *GetBranchInfoResponseBodyResultCommitInfo) *GetBranchInfoResponseBodyResult {
	s.CommitInfo = v
	return s
}

func (s *GetBranchInfoResponseBodyResult) SetProtectedBranch(v bool) *GetBranchInfoResponseBodyResult {
	s.ProtectedBranch = &v
	return s
}

type GetBranchInfoResponseBodyResultCommitInfo struct {
	AuthorDate     *string   `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	AuthorEmail    *string   `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string   `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommittedDate  *string   `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string   `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string   `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string   `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetBranchInfoResponseBodyResultCommitInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBranchInfoResponseBodyResultCommitInfo) GoString() string {
	return s.String()
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetAuthorDate(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.AuthorDate = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetAuthorEmail(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.AuthorEmail = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetAuthorName(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.AuthorName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetCommittedDate(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.CommittedDate = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetCommitterEmail(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.CommitterEmail = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetCommitterName(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.CommitterName = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetCreatedAt(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.CreatedAt = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetId(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.Id = &v
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

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetShortId(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.ShortId = &v
	return s
}

func (s *GetBranchInfoResponseBodyResultCommitInfo) SetTitle(v string) *GetBranchInfoResponseBodyResultCommitInfo {
	s.Title = &v
	return s
}

type GetBranchInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBranchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetBranchInfoResponse) SetStatusCode(v int32) *GetBranchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBranchInfoResponse) SetBody(v *GetBranchInfoResponseBody) *GetBranchInfoResponse {
	s.Body = v
	return s
}

type GetCodeCompletionRequest struct {
	FetchKeys   *string `json:"FetchKeys,omitempty" xml:"FetchKeys,omitempty"`
	IsEncrypted *bool   `json:"IsEncrypted,omitempty" xml:"IsEncrypted,omitempty"`
}

func (s GetCodeCompletionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCodeCompletionRequest) GoString() string {
	return s.String()
}

func (s *GetCodeCompletionRequest) SetFetchKeys(v string) *GetCodeCompletionRequest {
	s.FetchKeys = &v
	return s
}

func (s *GetCodeCompletionRequest) SetIsEncrypted(v bool) *GetCodeCompletionRequest {
	s.IsEncrypted = &v
	return s
}

type GetCodeCompletionResponseBody struct {
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetCodeCompletionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCodeCompletionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeCompletionResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeCompletionResponseBody) SetErrorCode(v string) *GetCodeCompletionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCodeCompletionResponseBody) SetErrorMessage(v string) *GetCodeCompletionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCodeCompletionResponseBody) SetRequestId(v string) *GetCodeCompletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeCompletionResponseBody) SetResult(v *GetCodeCompletionResponseBodyResult) *GetCodeCompletionResponseBody {
	s.Result = v
	return s
}

func (s *GetCodeCompletionResponseBody) SetSuccess(v bool) *GetCodeCompletionResponseBody {
	s.Success = &v
	return s
}

type GetCodeCompletionResponseBodyResult struct {
	Body             *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ClientTimestamp  *string `json:"ClientTimestamp,omitempty" xml:"ClientTimestamp,omitempty"`
	FetchTimestamp   *string `json:"FetchTimestamp,omitempty" xml:"FetchTimestamp,omitempty"`
	InvokeTimestamp  *string `json:"InvokeTimestamp,omitempty" xml:"InvokeTimestamp,omitempty"`
	ReceiveTimestamp *string `json:"ReceiveTimestamp,omitempty" xml:"ReceiveTimestamp,omitempty"`
	RspTimestamp     *string `json:"RspTimestamp,omitempty" xml:"RspTimestamp,omitempty"`
}

func (s GetCodeCompletionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCodeCompletionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCodeCompletionResponseBodyResult) SetBody(v string) *GetCodeCompletionResponseBodyResult {
	s.Body = &v
	return s
}

func (s *GetCodeCompletionResponseBodyResult) SetClientTimestamp(v string) *GetCodeCompletionResponseBodyResult {
	s.ClientTimestamp = &v
	return s
}

func (s *GetCodeCompletionResponseBodyResult) SetFetchTimestamp(v string) *GetCodeCompletionResponseBodyResult {
	s.FetchTimestamp = &v
	return s
}

func (s *GetCodeCompletionResponseBodyResult) SetInvokeTimestamp(v string) *GetCodeCompletionResponseBodyResult {
	s.InvokeTimestamp = &v
	return s
}

func (s *GetCodeCompletionResponseBodyResult) SetReceiveTimestamp(v string) *GetCodeCompletionResponseBodyResult {
	s.ReceiveTimestamp = &v
	return s
}

func (s *GetCodeCompletionResponseBodyResult) SetRspTimestamp(v string) *GetCodeCompletionResponseBodyResult {
	s.RspTimestamp = &v
	return s
}

type GetCodeCompletionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCodeCompletionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCodeCompletionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCodeCompletionResponse) GoString() string {
	return s.String()
}

func (s *GetCodeCompletionResponse) SetHeaders(v map[string]*string) *GetCodeCompletionResponse {
	s.Headers = v
	return s
}

func (s *GetCodeCompletionResponse) SetStatusCode(v int32) *GetCodeCompletionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeCompletionResponse) SetBody(v *GetCodeCompletionResponseBody) *GetCodeCompletionResponse {
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
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetCodeupOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCodeupOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponseBody) SetErrorCode(v string) *GetCodeupOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetErrorMessage(v string) *GetCodeupOrganizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetRequestId(v string) *GetCodeupOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetResult(v *GetCodeupOrganizationResponseBodyResult) *GetCodeupOrganizationResponseBody {
	s.Result = v
	return s
}

func (s *GetCodeupOrganizationResponseBody) SetSuccess(v bool) *GetCodeupOrganizationResponseBody {
	s.Success = &v
	return s
}

type GetCodeupOrganizationResponseBodyResult struct {
	CreatedAt      *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	NamespaceId    *int64  `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdatedAt      *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	UserRole       *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
}

func (s GetCodeupOrganizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCodeupOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationResponseBodyResult) SetCreatedAt(v string) *GetCodeupOrganizationResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetId(v int64) *GetCodeupOrganizationResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetNamespaceId(v int64) *GetCodeupOrganizationResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetOrganizationId(v string) *GetCodeupOrganizationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetPath(v string) *GetCodeupOrganizationResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetUpdatedAt(v string) *GetCodeupOrganizationResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *GetCodeupOrganizationResponseBodyResult) SetUserRole(v string) *GetCodeupOrganizationResponseBodyResult {
	s.UserRole = &v
	return s
}

type GetCodeupOrganizationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCodeupOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCodeupOrganizationResponse) SetStatusCode(v int32) *GetCodeupOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeupOrganizationResponse) SetBody(v *GetCodeupOrganizationResponseBody) *GetCodeupOrganizationResponse {
	s.Body = v
	return s
}

type GetFileBlobsRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	FilePath       *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	From           *int64  `json:"From,omitempty" xml:"From,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Ref            *string `json:"Ref,omitempty" xml:"Ref,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	To             *int64  `json:"To,omitempty" xml:"To,omitempty"`
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

func (s *GetFileBlobsRequest) SetFilePath(v string) *GetFileBlobsRequest {
	s.FilePath = &v
	return s
}

func (s *GetFileBlobsRequest) SetFrom(v int64) *GetFileBlobsRequest {
	s.From = &v
	return s
}

func (s *GetFileBlobsRequest) SetOrganizationId(v string) *GetFileBlobsRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetFileBlobsRequest) SetRef(v string) *GetFileBlobsRequest {
	s.Ref = &v
	return s
}

func (s *GetFileBlobsRequest) SetSubUserId(v string) *GetFileBlobsRequest {
	s.SubUserId = &v
	return s
}

func (s *GetFileBlobsRequest) SetTo(v int64) *GetFileBlobsRequest {
	s.To = &v
	return s
}

type GetFileBlobsResponseBody struct {
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetFileBlobsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileBlobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponseBody) SetErrorCode(v string) *GetFileBlobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetErrorMessage(v string) *GetFileBlobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetRequestId(v string) *GetFileBlobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileBlobsResponseBody) SetResult(v *GetFileBlobsResponseBodyResult) *GetFileBlobsResponseBody {
	s.Result = v
	return s
}

func (s *GetFileBlobsResponseBody) SetSuccess(v bool) *GetFileBlobsResponseBody {
	s.Success = &v
	return s
}

type GetFileBlobsResponseBodyResult struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	TotalLines *int32  `json:"TotalLines,omitempty" xml:"TotalLines,omitempty"`
}

func (s GetFileBlobsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFileBlobsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileBlobsResponseBodyResult) SetContent(v string) *GetFileBlobsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetFileBlobsResponseBodyResult) SetTotalLines(v int32) *GetFileBlobsResponseBodyResult {
	s.TotalLines = &v
	return s
}

type GetFileBlobsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileBlobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetFileBlobsResponse) SetStatusCode(v int32) *GetFileBlobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileBlobsResponse) SetBody(v *GetFileBlobsResponseBody) *GetFileBlobsResponse {
	s.Body = v
	return s
}

type GetFileLastCommitRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	FilePath       *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Sha            *string `json:"Sha,omitempty" xml:"Sha,omitempty"`
}

func (s GetFileLastCommitRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitRequest) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitRequest) SetAccessToken(v string) *GetFileLastCommitRequest {
	s.AccessToken = &v
	return s
}

func (s *GetFileLastCommitRequest) SetFilePath(v string) *GetFileLastCommitRequest {
	s.FilePath = &v
	return s
}

func (s *GetFileLastCommitRequest) SetOrganizationId(v string) *GetFileLastCommitRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetFileLastCommitRequest) SetSha(v string) *GetFileLastCommitRequest {
	s.Sha = &v
	return s
}

type GetFileLastCommitResponseBody struct {
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetFileLastCommitResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileLastCommitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBody) SetErrorCode(v string) *GetFileLastCommitResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetErrorMessage(v string) *GetFileLastCommitResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetRequestId(v string) *GetFileLastCommitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileLastCommitResponseBody) SetResult(v *GetFileLastCommitResponseBodyResult) *GetFileLastCommitResponseBody {
	s.Result = v
	return s
}

func (s *GetFileLastCommitResponseBody) SetSuccess(v bool) *GetFileLastCommitResponseBody {
	s.Success = &v
	return s
}

type GetFileLastCommitResponseBodyResult struct {
	AuthorDate     *string                                       `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	AuthorEmail    *string                                       `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string                                       `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommittedDate  *string                                       `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string                                       `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string                                       `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string                                       `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string                                     `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string                                       `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Signature      *GetFileLastCommitResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Title          *string                                       `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetFileLastCommitResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorDate(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorDate = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorEmail(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorEmail = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetAuthorName(v string) *GetFileLastCommitResponseBodyResult {
	s.AuthorName = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommittedDate(v string) *GetFileLastCommitResponseBodyResult {
	s.CommittedDate = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommitterEmail(v string) *GetFileLastCommitResponseBodyResult {
	s.CommitterEmail = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCommitterName(v string) *GetFileLastCommitResponseBodyResult {
	s.CommitterName = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetCreatedAt(v string) *GetFileLastCommitResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetId(v string) *GetFileLastCommitResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetMessage(v string) *GetFileLastCommitResponseBodyResult {
	s.Message = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetParentIds(v []*string) *GetFileLastCommitResponseBodyResult {
	s.ParentIds = v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetShortId(v string) *GetFileLastCommitResponseBodyResult {
	s.ShortId = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetSignature(v *GetFileLastCommitResponseBodyResultSignature) *GetFileLastCommitResponseBodyResult {
	s.Signature = v
	return s
}

func (s *GetFileLastCommitResponseBodyResult) SetTitle(v string) *GetFileLastCommitResponseBodyResult {
	s.Title = &v
	return s
}

type GetFileLastCommitResponseBodyResultSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s GetFileLastCommitResponseBodyResultSignature) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponseBodyResultSignature) SetGpgKeyId(v string) *GetFileLastCommitResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetFileLastCommitResponseBodyResultSignature) SetVerificationStatus(v string) *GetFileLastCommitResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

type GetFileLastCommitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileLastCommitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileLastCommitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileLastCommitResponse) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponse) SetHeaders(v map[string]*string) *GetFileLastCommitResponse {
	s.Headers = v
	return s
}

func (s *GetFileLastCommitResponse) SetStatusCode(v int32) *GetFileLastCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileLastCommitResponse) SetBody(v *GetFileLastCommitResponseBody) *GetFileLastCommitResponse {
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
	ErrorCode    *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetGroupDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponseBody) SetErrorCode(v string) *GetGroupDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetErrorMessage(v string) *GetGroupDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetRequestId(v string) *GetGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupDetailResponseBody) SetResult(v *GetGroupDetailResponseBodyResult) *GetGroupDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetGroupDetailResponseBody) SetSuccess(v bool) *GetGroupDetailResponseBody {
	s.Success = &v
	return s
}

type GetGroupDetailResponseBodyResult struct {
	AvatarUrl         *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentId          *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s GetGroupDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetGroupDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupDetailResponseBodyResult) SetAvatarUrl(v string) *GetGroupDetailResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetDescription(v string) *GetGroupDetailResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetId(v int64) *GetGroupDetailResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetName(v string) *GetGroupDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetNameWithNamespace(v string) *GetGroupDetailResponseBodyResult {
	s.NameWithNamespace = &v
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

func (s *GetGroupDetailResponseBodyResult) SetPath(v string) *GetGroupDetailResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetPathWithNamespace(v string) *GetGroupDetailResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetType(v string) *GetGroupDetailResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetVisibilityLevel(v string) *GetGroupDetailResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *GetGroupDetailResponseBodyResult) SetWebUrl(v string) *GetGroupDetailResponseBodyResult {
	s.WebUrl = &v
	return s
}

type GetGroupDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetGroupDetailResponse) SetStatusCode(v int32) *GetGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupDetailResponse) SetBody(v *GetGroupDetailResponseBody) *GetGroupDetailResponse {
	s.Body = v
	return s
}

type GetMergeRequestApproveStatusRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetMergeRequestApproveStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestApproveStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMergeRequestApproveStatusRequest) SetAccessToken(v string) *GetMergeRequestApproveStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMergeRequestApproveStatusRequest) SetOrganizationId(v string) *GetMergeRequestApproveStatusRequest {
	s.OrganizationId = &v
	return s
}

type GetMergeRequestApproveStatusResponseBody struct {
	ErrorCode    *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetMergeRequestApproveStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMergeRequestApproveStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestApproveStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMergeRequestApproveStatusResponseBody) SetErrorCode(v string) *GetMergeRequestApproveStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMergeRequestApproveStatusResponseBody) SetErrorMessage(v string) *GetMergeRequestApproveStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMergeRequestApproveStatusResponseBody) SetRequestId(v string) *GetMergeRequestApproveStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMergeRequestApproveStatusResponseBody) SetResult(v *GetMergeRequestApproveStatusResponseBodyResult) *GetMergeRequestApproveStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetMergeRequestApproveStatusResponseBody) SetSuccess(v bool) *GetMergeRequestApproveStatusResponseBody {
	s.Success = &v
	return s
}

type GetMergeRequestApproveStatusResponseBodyResult struct {
	ApproveStatus *string `json:"ApproveStatus,omitempty" xml:"ApproveStatus,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetMergeRequestApproveStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestApproveStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMergeRequestApproveStatusResponseBodyResult) SetApproveStatus(v string) *GetMergeRequestApproveStatusResponseBodyResult {
	s.ApproveStatus = &v
	return s
}

func (s *GetMergeRequestApproveStatusResponseBodyResult) SetMessage(v string) *GetMergeRequestApproveStatusResponseBodyResult {
	s.Message = &v
	return s
}

type GetMergeRequestApproveStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMergeRequestApproveStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMergeRequestApproveStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestApproveStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMergeRequestApproveStatusResponse) SetHeaders(v map[string]*string) *GetMergeRequestApproveStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMergeRequestApproveStatusResponse) SetStatusCode(v int32) *GetMergeRequestApproveStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMergeRequestApproveStatusResponse) SetBody(v *GetMergeRequestApproveStatusResponseBody) *GetMergeRequestApproveStatusResponse {
	s.Body = v
	return s
}

type GetMergeRequestDetailRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetMergeRequestDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailRequest) SetAccessToken(v string) *GetMergeRequestDetailRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMergeRequestDetailRequest) SetOrganizationId(v string) *GetMergeRequestDetailRequest {
	s.OrganizationId = &v
	return s
}

type GetMergeRequestDetailResponseBody struct {
	ErrorCode    *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetMergeRequestDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMergeRequestDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBody) SetErrorCode(v string) *GetMergeRequestDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMergeRequestDetailResponseBody) SetErrorMessage(v string) *GetMergeRequestDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMergeRequestDetailResponseBody) SetRequestId(v string) *GetMergeRequestDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMergeRequestDetailResponseBody) SetResult(v *GetMergeRequestDetailResponseBodyResult) *GetMergeRequestDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetMergeRequestDetailResponseBody) SetSuccess(v bool) *GetMergeRequestDetailResponseBody {
	s.Success = &v
	return s
}

type GetMergeRequestDetailResponseBodyResult struct {
	AcceptedRevision   *string                                                    `json:"AcceptedRevision,omitempty" xml:"AcceptedRevision,omitempty"`
	AheadCommitCount   *int32                                                     `json:"AheadCommitCount,omitempty" xml:"AheadCommitCount,omitempty"`
	ApproveCheckResult *GetMergeRequestDetailResponseBodyResultApproveCheckResult `json:"ApproveCheckResult,omitempty" xml:"ApproveCheckResult,omitempty" type:"Struct"`
	AssigneeList       []*GetMergeRequestDetailResponseBodyResultAssigneeList     `json:"AssigneeList,omitempty" xml:"AssigneeList,omitempty" type:"Repeated"`
	Author             *GetMergeRequestDetailResponseBodyResultAuthor             `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	BehindCommitCount  *int32                                                     `json:"BehindCommitCount,omitempty" xml:"BehindCommitCount,omitempty"`
	CreatedAt          *string                                                    `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description        *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                 *int64                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	IsSupportMerge     *bool                                                      `json:"IsSupportMerge,omitempty" xml:"IsSupportMerge,omitempty"`
	MergeError         *string                                                    `json:"MergeError,omitempty" xml:"MergeError,omitempty"`
	MergeStatus        *string                                                    `json:"MergeStatus,omitempty" xml:"MergeStatus,omitempty"`
	MergeType          *string                                                    `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	MergedRevision     *string                                                    `json:"MergedRevision,omitempty" xml:"MergedRevision,omitempty"`
	NameWithNamespace  *string                                                    `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	ProjectId          *int64                                                     `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SourceBranch       *string                                                    `json:"SourceBranch,omitempty" xml:"SourceBranch,omitempty"`
	State              *string                                                    `json:"State,omitempty" xml:"State,omitempty"`
	TargetBranch       *string                                                    `json:"TargetBranch,omitempty" xml:"TargetBranch,omitempty"`
	Title              *string                                                    `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdatedAt          *string                                                    `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	WebUrl             *string                                                    `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s GetMergeRequestDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBodyResult) SetAcceptedRevision(v string) *GetMergeRequestDetailResponseBodyResult {
	s.AcceptedRevision = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetAheadCommitCount(v int32) *GetMergeRequestDetailResponseBodyResult {
	s.AheadCommitCount = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetApproveCheckResult(v *GetMergeRequestDetailResponseBodyResultApproveCheckResult) *GetMergeRequestDetailResponseBodyResult {
	s.ApproveCheckResult = v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetAssigneeList(v []*GetMergeRequestDetailResponseBodyResultAssigneeList) *GetMergeRequestDetailResponseBodyResult {
	s.AssigneeList = v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetAuthor(v *GetMergeRequestDetailResponseBodyResultAuthor) *GetMergeRequestDetailResponseBodyResult {
	s.Author = v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetBehindCommitCount(v int32) *GetMergeRequestDetailResponseBodyResult {
	s.BehindCommitCount = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetCreatedAt(v string) *GetMergeRequestDetailResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetDescription(v string) *GetMergeRequestDetailResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetId(v int64) *GetMergeRequestDetailResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetIsSupportMerge(v bool) *GetMergeRequestDetailResponseBodyResult {
	s.IsSupportMerge = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetMergeError(v string) *GetMergeRequestDetailResponseBodyResult {
	s.MergeError = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetMergeStatus(v string) *GetMergeRequestDetailResponseBodyResult {
	s.MergeStatus = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetMergeType(v string) *GetMergeRequestDetailResponseBodyResult {
	s.MergeType = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetMergedRevision(v string) *GetMergeRequestDetailResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetNameWithNamespace(v string) *GetMergeRequestDetailResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetProjectId(v int64) *GetMergeRequestDetailResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetSourceBranch(v string) *GetMergeRequestDetailResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetState(v string) *GetMergeRequestDetailResponseBodyResult {
	s.State = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetTargetBranch(v string) *GetMergeRequestDetailResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetTitle(v string) *GetMergeRequestDetailResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetUpdatedAt(v string) *GetMergeRequestDetailResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResult) SetWebUrl(v string) *GetMergeRequestDetailResponseBodyResult {
	s.WebUrl = &v
	return s
}

type GetMergeRequestDetailResponseBodyResultApproveCheckResult struct {
	SatisfiedCheckResults   []*GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults   `json:"SatisfiedCheckResults,omitempty" xml:"SatisfiedCheckResults,omitempty" type:"Repeated"`
	TotalCheckResult        *string                                                                             `json:"TotalCheckResult,omitempty" xml:"TotalCheckResult,omitempty"`
	UnsatisfiedCheckResults []*GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults `json:"UnsatisfiedCheckResults,omitempty" xml:"UnsatisfiedCheckResults,omitempty" type:"Repeated"`
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResult) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResult) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResult) SetSatisfiedCheckResults(v []*GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) *GetMergeRequestDetailResponseBodyResultApproveCheckResult {
	s.SatisfiedCheckResults = v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResult) SetTotalCheckResult(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResult {
	s.TotalCheckResult = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResult) SetUnsatisfiedCheckResults(v []*GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) *GetMergeRequestDetailResponseBodyResultApproveCheckResult {
	s.UnsatisfiedCheckResults = v
	return s
}

type GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults struct {
	CheckName        *string                                                                                     `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                     `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                     `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                   `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                   `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckName(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckName = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckStatus(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckType(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetExtraUsers(v []*GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetSatisfiedItems(v []*string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetId(v int64) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetName(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults struct {
	CheckName        *string                                                                                       `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                       `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                       `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                     `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                     `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckName(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckName = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckStatus(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckType(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetExtraUsers(v []*GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetSatisfiedItems(v []*string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetId(v int64) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetName(v string) *GetMergeRequestDetailResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type GetMergeRequestDetailResponseBodyResultAssigneeList struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMergeRequestDetailResponseBodyResultAssigneeList) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBodyResultAssigneeList) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBodyResultAssigneeList) SetAvatarUrl(v string) *GetMergeRequestDetailResponseBodyResultAssigneeList {
	s.AvatarUrl = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultAssigneeList) SetEmail(v string) *GetMergeRequestDetailResponseBodyResultAssigneeList {
	s.Email = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultAssigneeList) SetExternUserId(v string) *GetMergeRequestDetailResponseBodyResultAssigneeList {
	s.ExternUserId = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultAssigneeList) SetId(v string) *GetMergeRequestDetailResponseBodyResultAssigneeList {
	s.Id = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultAssigneeList) SetName(v string) *GetMergeRequestDetailResponseBodyResultAssigneeList {
	s.Name = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultAssigneeList) SetStatus(v string) *GetMergeRequestDetailResponseBodyResultAssigneeList {
	s.Status = &v
	return s
}

type GetMergeRequestDetailResponseBodyResultAuthor struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMergeRequestDetailResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponseBodyResultAuthor) SetAvatarUrl(v string) *GetMergeRequestDetailResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultAuthor) SetExternUserId(v string) *GetMergeRequestDetailResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultAuthor) SetId(v int64) *GetMergeRequestDetailResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *GetMergeRequestDetailResponseBodyResultAuthor) SetName(v string) *GetMergeRequestDetailResponseBodyResultAuthor {
	s.Name = &v
	return s
}

type GetMergeRequestDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMergeRequestDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMergeRequestDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMergeRequestDetailResponse) SetHeaders(v map[string]*string) *GetMergeRequestDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMergeRequestDetailResponse) SetStatusCode(v int32) *GetMergeRequestDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMergeRequestDetailResponse) SetBody(v *GetMergeRequestDetailResponseBody) *GetMergeRequestDetailResponse {
	s.Body = v
	return s
}

type GetMergeRequestSettingRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetMergeRequestSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestSettingRequest) GoString() string {
	return s.String()
}

func (s *GetMergeRequestSettingRequest) SetAccessToken(v string) *GetMergeRequestSettingRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMergeRequestSettingRequest) SetOrganizationId(v string) *GetMergeRequestSettingRequest {
	s.OrganizationId = &v
	return s
}

type GetMergeRequestSettingResponseBody struct {
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetMergeRequestSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMergeRequestSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetMergeRequestSettingResponseBody) SetErrorCode(v string) *GetMergeRequestSettingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMergeRequestSettingResponseBody) SetErrorMessage(v string) *GetMergeRequestSettingResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMergeRequestSettingResponseBody) SetRequestId(v string) *GetMergeRequestSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMergeRequestSettingResponseBody) SetResult(v *GetMergeRequestSettingResponseBodyResult) *GetMergeRequestSettingResponseBody {
	s.Result = v
	return s
}

func (s *GetMergeRequestSettingResponseBody) SetSuccess(v bool) *GetMergeRequestSettingResponseBody {
	s.Success = &v
	return s
}

type GetMergeRequestSettingResponseBodyResult struct {
	IsEnableSmartCodeReview *bool     `json:"IsEnableSmartCodeReview,omitempty" xml:"IsEnableSmartCodeReview,omitempty"`
	MergeTypes              []*string `json:"MergeTypes,omitempty" xml:"MergeTypes,omitempty" type:"Repeated"`
}

func (s GetMergeRequestSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMergeRequestSettingResponseBodyResult) SetIsEnableSmartCodeReview(v bool) *GetMergeRequestSettingResponseBodyResult {
	s.IsEnableSmartCodeReview = &v
	return s
}

func (s *GetMergeRequestSettingResponseBodyResult) SetMergeTypes(v []*string) *GetMergeRequestSettingResponseBodyResult {
	s.MergeTypes = v
	return s
}

type GetMergeRequestSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMergeRequestSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMergeRequestSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMergeRequestSettingResponse) GoString() string {
	return s.String()
}

func (s *GetMergeRequestSettingResponse) SetHeaders(v map[string]*string) *GetMergeRequestSettingResponse {
	s.Headers = v
	return s
}

func (s *GetMergeRequestSettingResponse) SetStatusCode(v int32) *GetMergeRequestSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMergeRequestSettingResponse) SetBody(v *GetMergeRequestSettingResponseBody) *GetMergeRequestSettingResponse {
	s.Body = v
	return s
}

type GetOrganizationRepositorySettingRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetOrganizationRepositorySettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationRepositorySettingRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationRepositorySettingRequest) SetAccessToken(v string) *GetOrganizationRepositorySettingRequest {
	s.AccessToken = &v
	return s
}

func (s *GetOrganizationRepositorySettingRequest) SetOrganizationId(v string) *GetOrganizationRepositorySettingRequest {
	s.OrganizationId = &v
	return s
}

type GetOrganizationRepositorySettingResponseBody struct {
	ErrorCode    *string                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetOrganizationRepositorySettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrganizationRepositorySettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationRepositorySettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationRepositorySettingResponseBody) SetErrorCode(v string) *GetOrganizationRepositorySettingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBody) SetErrorMessage(v string) *GetOrganizationRepositorySettingResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBody) SetRequestId(v string) *GetOrganizationRepositorySettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBody) SetResult(v *GetOrganizationRepositorySettingResponseBodyResult) *GetOrganizationRepositorySettingResponseBody {
	s.Result = v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBody) SetSuccess(v bool) *GetOrganizationRepositorySettingResponseBody {
	s.Success = &v
	return s
}

type GetOrganizationRepositorySettingResponseBodyResult struct {
	ForcePushForbidden             *bool                                                                           `json:"ForcePushForbidden,omitempty" xml:"ForcePushForbidden,omitempty"`
	GroupRequired                  *bool                                                                           `json:"GroupRequired,omitempty" xml:"GroupRequired,omitempty"`
	OpenCloneDownloadControl       *bool                                                                           `json:"OpenCloneDownloadControl,omitempty" xml:"OpenCloneDownloadControl,omitempty"`
	OrgCloneDownloadMethodList     []*GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList `json:"OrgCloneDownloadMethodList,omitempty" xml:"OrgCloneDownloadMethodList,omitempty" type:"Repeated"`
	OrgCloneDownloadRoleList       []*GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList   `json:"OrgCloneDownloadRoleList,omitempty" xml:"OrgCloneDownloadRoleList,omitempty" type:"Repeated"`
	RepoAdminAccessVisibilityLevel []*int64                                                                        `json:"RepoAdminAccessVisibilityLevel,omitempty" xml:"RepoAdminAccessVisibilityLevel,omitempty" type:"Repeated"`
	RepoAdminOperation             []*int64                                                                        `json:"RepoAdminOperation,omitempty" xml:"RepoAdminOperation,omitempty" type:"Repeated"`
	RepoCreatorIdentity            []*int64                                                                        `json:"RepoCreatorIdentity,omitempty" xml:"RepoCreatorIdentity,omitempty" type:"Repeated"`
	RepoVisibilityLevel            []*int64                                                                        `json:"RepoVisibilityLevel,omitempty" xml:"RepoVisibilityLevel,omitempty" type:"Repeated"`
}

func (s GetOrganizationRepositorySettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationRepositorySettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetForcePushForbidden(v bool) *GetOrganizationRepositorySettingResponseBodyResult {
	s.ForcePushForbidden = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetGroupRequired(v bool) *GetOrganizationRepositorySettingResponseBodyResult {
	s.GroupRequired = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetOpenCloneDownloadControl(v bool) *GetOrganizationRepositorySettingResponseBodyResult {
	s.OpenCloneDownloadControl = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetOrgCloneDownloadMethodList(v []*GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList) *GetOrganizationRepositorySettingResponseBodyResult {
	s.OrgCloneDownloadMethodList = v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetOrgCloneDownloadRoleList(v []*GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList) *GetOrganizationRepositorySettingResponseBodyResult {
	s.OrgCloneDownloadRoleList = v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetRepoAdminAccessVisibilityLevel(v []*int64) *GetOrganizationRepositorySettingResponseBodyResult {
	s.RepoAdminAccessVisibilityLevel = v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetRepoAdminOperation(v []*int64) *GetOrganizationRepositorySettingResponseBodyResult {
	s.RepoAdminOperation = v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetRepoCreatorIdentity(v []*int64) *GetOrganizationRepositorySettingResponseBodyResult {
	s.RepoCreatorIdentity = v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResult) SetRepoVisibilityLevel(v []*int64) *GetOrganizationRepositorySettingResponseBodyResult {
	s.RepoVisibilityLevel = v
	return s
}

type GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList struct {
	Allowed        *bool   `json:"Allowed,omitempty" xml:"Allowed,omitempty"`
	PermissionCode *string `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
}

func (s GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList) GoString() string {
	return s.String()
}

func (s *GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList) SetAllowed(v bool) *GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList {
	s.Allowed = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList) SetPermissionCode(v string) *GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadMethodList {
	s.PermissionCode = &v
	return s
}

type GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList struct {
	Allowed  *bool  `json:"Allowed,omitempty" xml:"Allowed,omitempty"`
	RoleCode *int64 `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
}

func (s GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList) GoString() string {
	return s.String()
}

func (s *GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList) SetAllowed(v bool) *GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList {
	s.Allowed = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList) SetRoleCode(v int64) *GetOrganizationRepositorySettingResponseBodyResultOrgCloneDownloadRoleList {
	s.RoleCode = &v
	return s
}

type GetOrganizationRepositorySettingResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizationRepositorySettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationRepositorySettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationRepositorySettingResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationRepositorySettingResponse) SetHeaders(v map[string]*string) *GetOrganizationRepositorySettingResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationRepositorySettingResponse) SetStatusCode(v int32) *GetOrganizationRepositorySettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationRepositorySettingResponse) SetBody(v *GetOrganizationRepositorySettingResponseBody) *GetOrganizationRepositorySettingResponse {
	s.Body = v
	return s
}

type GetOrganizationSecurityCenterStatusRequest struct {
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
}

func (s GetOrganizationSecurityCenterStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationSecurityCenterStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationSecurityCenterStatusRequest) SetAccessToken(v string) *GetOrganizationSecurityCenterStatusRequest {
	s.AccessToken = &v
	return s
}

type GetOrganizationSecurityCenterStatusResponseBody struct {
	ErrorCode    *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetOrganizationSecurityCenterStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrganizationSecurityCenterStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationSecurityCenterStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationSecurityCenterStatusResponseBody) SetErrorCode(v string) *GetOrganizationSecurityCenterStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrganizationSecurityCenterStatusResponseBody) SetErrorMessage(v string) *GetOrganizationSecurityCenterStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrganizationSecurityCenterStatusResponseBody) SetRequestId(v string) *GetOrganizationSecurityCenterStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrganizationSecurityCenterStatusResponseBody) SetResult(v *GetOrganizationSecurityCenterStatusResponseBodyResult) *GetOrganizationSecurityCenterStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetOrganizationSecurityCenterStatusResponseBody) SetSuccess(v bool) *GetOrganizationSecurityCenterStatusResponseBody {
	s.Success = &v
	return s
}

type GetOrganizationSecurityCenterStatusResponseBodyResult struct {
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s GetOrganizationSecurityCenterStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationSecurityCenterStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOrganizationSecurityCenterStatusResponseBodyResult) SetEnable(v bool) *GetOrganizationSecurityCenterStatusResponseBodyResult {
	s.Enable = &v
	return s
}

type GetOrganizationSecurityCenterStatusResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizationSecurityCenterStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationSecurityCenterStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationSecurityCenterStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationSecurityCenterStatusResponse) SetHeaders(v map[string]*string) *GetOrganizationSecurityCenterStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationSecurityCenterStatusResponse) SetStatusCode(v int32) *GetOrganizationSecurityCenterStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationSecurityCenterStatusResponse) SetBody(v *GetOrganizationSecurityCenterStatusResponseBody) *GetOrganizationSecurityCenterStatusResponse {
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
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetProjectMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBody) SetErrorCode(v string) *GetProjectMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetErrorMessage(v string) *GetProjectMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetRequestId(v string) *GetProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectMemberResponseBody) SetResult(v *GetProjectMemberResponseBodyResult) *GetProjectMemberResponseBody {
	s.Result = v
	return s
}

func (s *GetProjectMemberResponseBody) SetSuccess(v bool) *GetProjectMemberResponseBody {
	s.Success = &v
	return s
}

type GetProjectMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetProjectMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBodyResult) SetAccessLevel(v int32) *GetProjectMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetAvatarUrl(v string) *GetProjectMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetExternUserId(v string) *GetProjectMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetId(v int64) *GetProjectMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetProjectMemberResponseBodyResult) SetName(v string) *GetProjectMemberResponseBodyResult {
	s.Name = &v
	return s
}

type GetProjectMemberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetProjectMemberResponse) SetStatusCode(v int32) *GetProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectMemberResponse) SetBody(v *GetProjectMemberResponseBody) *GetProjectMemberResponse {
	s.Body = v
	return s
}

type GetRepositoryCommitRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetRepositoryCommitRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryCommitRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitRequest) SetAccessToken(v string) *GetRepositoryCommitRequest {
	s.AccessToken = &v
	return s
}

func (s *GetRepositoryCommitRequest) SetOrganizationId(v string) *GetRepositoryCommitRequest {
	s.OrganizationId = &v
	return s
}

type GetRepositoryCommitResponseBody struct {
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetRepositoryCommitResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRepositoryCommitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryCommitResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponseBody) SetErrorCode(v string) *GetRepositoryCommitResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryCommitResponseBody) SetErrorMessage(v string) *GetRepositoryCommitResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryCommitResponseBody) SetRequestId(v string) *GetRepositoryCommitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryCommitResponseBody) SetResult(v *GetRepositoryCommitResponseBodyResult) *GetRepositoryCommitResponseBody {
	s.Result = v
	return s
}

func (s *GetRepositoryCommitResponseBody) SetSuccess(v bool) *GetRepositoryCommitResponseBody {
	s.Success = &v
	return s
}

type GetRepositoryCommitResponseBodyResult struct {
	AuthorDate     *string                                         `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	AuthorEmail    *string                                         `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string                                         `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommittedDate  *string                                         `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string                                         `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string                                         `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string                                         `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string                                       `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string                                         `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Signature      *GetRepositoryCommitResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Title          *string                                         `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetRepositoryCommitResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryCommitResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponseBodyResult) SetAuthorDate(v string) *GetRepositoryCommitResponseBodyResult {
	s.AuthorDate = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetAuthorEmail(v string) *GetRepositoryCommitResponseBodyResult {
	s.AuthorEmail = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetAuthorName(v string) *GetRepositoryCommitResponseBodyResult {
	s.AuthorName = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCommittedDate(v string) *GetRepositoryCommitResponseBodyResult {
	s.CommittedDate = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCommitterEmail(v string) *GetRepositoryCommitResponseBodyResult {
	s.CommitterEmail = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCommitterName(v string) *GetRepositoryCommitResponseBodyResult {
	s.CommitterName = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetCreatedAt(v string) *GetRepositoryCommitResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetId(v string) *GetRepositoryCommitResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetMessage(v string) *GetRepositoryCommitResponseBodyResult {
	s.Message = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetParentIds(v []*string) *GetRepositoryCommitResponseBodyResult {
	s.ParentIds = v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetShortId(v string) *GetRepositoryCommitResponseBodyResult {
	s.ShortId = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetSignature(v *GetRepositoryCommitResponseBodyResultSignature) *GetRepositoryCommitResponseBodyResult {
	s.Signature = v
	return s
}

func (s *GetRepositoryCommitResponseBodyResult) SetTitle(v string) *GetRepositoryCommitResponseBodyResult {
	s.Title = &v
	return s
}

type GetRepositoryCommitResponseBodyResultSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s GetRepositoryCommitResponseBodyResultSignature) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryCommitResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponseBodyResultSignature) SetGpgKeyId(v string) *GetRepositoryCommitResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetRepositoryCommitResponseBodyResultSignature) SetVerificationStatus(v string) *GetRepositoryCommitResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

type GetRepositoryCommitResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRepositoryCommitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRepositoryCommitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryCommitResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitResponse) SetHeaders(v map[string]*string) *GetRepositoryCommitResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryCommitResponse) SetStatusCode(v int32) *GetRepositoryCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryCommitResponse) SetBody(v *GetRepositoryCommitResponseBody) *GetRepositoryCommitResponse {
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
	ErrorCode    *int32                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetRepositoryInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRepositoryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBody) SetErrorCode(v int32) *GetRepositoryInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryInfoResponseBody) SetErrorMessage(v string) *GetRepositoryInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryInfoResponseBody) SetRequestId(v string) *GetRepositoryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryInfoResponseBody) SetResult(v *GetRepositoryInfoResponseBodyResult) *GetRepositoryInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetRepositoryInfoResponseBody) SetSuccess(v bool) *GetRepositoryInfoResponseBody {
	s.Success = &v
	return s
}

type GetRepositoryInfoResponseBodyResult struct {
	AccessLevel          *int32                                          `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Archive              *bool                                           `json:"Archive,omitempty" xml:"Archive,omitempty"`
	AvatarUrl            *string                                         `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	CreatedAt            *string                                         `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CreatorId            *int64                                          `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DefaultBranch        *string                                         `json:"DefaultBranch,omitempty" xml:"DefaultBranch,omitempty"`
	DemoProjectStatus    *bool                                           `json:"DemoProjectStatus,omitempty" xml:"DemoProjectStatus,omitempty"`
	Description          *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpUrlToRepo        *string                                         `json:"HttpUrlToRepo,omitempty" xml:"HttpUrlToRepo,omitempty"`
	Id                   *int64                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	ImportFromSubversion *bool                                           `json:"ImportFromSubversion,omitempty" xml:"ImportFromSubversion,omitempty"`
	ImportStatus         *string                                         `json:"ImportStatus,omitempty" xml:"ImportStatus,omitempty"`
	ImportUrl            *string                                         `json:"ImportUrl,omitempty" xml:"ImportUrl,omitempty"`
	LastActivityAt       *string                                         `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	Name                 *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace    *string                                         `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	Namespace            *GetRepositoryInfoResponseBodyResultNamespace   `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Struct"`
	Path                 *string                                         `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace    *string                                         `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Permissions          *GetRepositoryInfoResponseBodyResultPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Struct"`
	Public               *bool                                           `json:"Public,omitempty" xml:"Public,omitempty"`
	SshUrlToRepo         *string                                         `json:"SshUrlToRepo,omitempty" xml:"SshUrlToRepo,omitempty"`
	TagList              []*string                                       `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	VisibilityLevel      *string                                         `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WebUrl               *string                                         `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s GetRepositoryInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBodyResult) SetAccessLevel(v int32) *GetRepositoryInfoResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetArchive(v bool) *GetRepositoryInfoResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetAvatarUrl(v string) *GetRepositoryInfoResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetCreatedAt(v string) *GetRepositoryInfoResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetCreatorId(v int64) *GetRepositoryInfoResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetDefaultBranch(v string) *GetRepositoryInfoResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetDemoProjectStatus(v bool) *GetRepositoryInfoResponseBodyResult {
	s.DemoProjectStatus = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetDescription(v string) *GetRepositoryInfoResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetHttpUrlToRepo(v string) *GetRepositoryInfoResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetId(v int64) *GetRepositoryInfoResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetImportFromSubversion(v bool) *GetRepositoryInfoResponseBodyResult {
	s.ImportFromSubversion = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetImportStatus(v string) *GetRepositoryInfoResponseBodyResult {
	s.ImportStatus = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetImportUrl(v string) *GetRepositoryInfoResponseBodyResult {
	s.ImportUrl = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetLastActivityAt(v string) *GetRepositoryInfoResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetName(v string) *GetRepositoryInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetNameWithNamespace(v string) *GetRepositoryInfoResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetNamespace(v *GetRepositoryInfoResponseBodyResultNamespace) *GetRepositoryInfoResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetPath(v string) *GetRepositoryInfoResponseBodyResult {
	s.Path = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetPathWithNamespace(v string) *GetRepositoryInfoResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetPermissions(v *GetRepositoryInfoResponseBodyResultPermissions) *GetRepositoryInfoResponseBodyResult {
	s.Permissions = v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetPublic(v bool) *GetRepositoryInfoResponseBodyResult {
	s.Public = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetSshUrlToRepo(v string) *GetRepositoryInfoResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetTagList(v []*string) *GetRepositoryInfoResponseBodyResult {
	s.TagList = v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetVisibilityLevel(v string) *GetRepositoryInfoResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResult) SetWebUrl(v string) *GetRepositoryInfoResponseBodyResult {
	s.WebUrl = &v
	return s
}

type GetRepositoryInfoResponseBodyResultNamespace struct {
	Avatar          *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	CreatedAt       *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Public          *bool   `json:"Public,omitempty" xml:"Public,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdatedAt       *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	VisibilityLevel *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
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

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetCreatedAt(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetDescription(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetId(v int64) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Id = &v
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

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetPath(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetPublic(v bool) *GetRepositoryInfoResponseBodyResultNamespace {
	s.Public = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetState(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.State = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetUpdatedAt(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultNamespace) SetVisibilityLevel(v string) *GetRepositoryInfoResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

type GetRepositoryInfoResponseBodyResultPermissions struct {
	GroupAccess   *GetRepositoryInfoResponseBodyResultPermissionsGroupAccess   `json:"GroupAccess,omitempty" xml:"GroupAccess,omitempty" type:"Struct"`
	ProjectAccess *GetRepositoryInfoResponseBodyResultPermissionsProjectAccess `json:"ProjectAccess,omitempty" xml:"ProjectAccess,omitempty" type:"Struct"`
}

func (s GetRepositoryInfoResponseBodyResultPermissions) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryInfoResponseBodyResultPermissions) GoString() string {
	return s.String()
}

func (s *GetRepositoryInfoResponseBodyResultPermissions) SetGroupAccess(v *GetRepositoryInfoResponseBodyResultPermissionsGroupAccess) *GetRepositoryInfoResponseBodyResultPermissions {
	s.GroupAccess = v
	return s
}

func (s *GetRepositoryInfoResponseBodyResultPermissions) SetProjectAccess(v *GetRepositoryInfoResponseBodyResultPermissionsProjectAccess) *GetRepositoryInfoResponseBodyResultPermissions {
	s.ProjectAccess = v
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

type GetRepositoryInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRepositoryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRepositoryInfoResponse) SetStatusCode(v int32) *GetRepositoryInfoResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetRepositoryTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRepositoryTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBody) SetErrorCode(v string) *GetRepositoryTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetErrorMessage(v string) *GetRepositoryTagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetRequestId(v string) *GetRepositoryTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryTagResponseBody) SetResult(v *GetRepositoryTagResponseBodyResult) *GetRepositoryTagResponseBody {
	s.Result = v
	return s
}

func (s *GetRepositoryTagResponseBody) SetSuccess(v bool) *GetRepositoryTagResponseBody {
	s.Success = &v
	return s
}

type GetRepositoryTagResponseBodyResult struct {
	Commit    *GetRepositoryTagResponseBodyResultCommit    `json:"Commit,omitempty" xml:"Commit,omitempty" type:"Struct"`
	Id        *string                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Name      *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Signature *GetRepositoryTagResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
}

func (s GetRepositoryTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResult) SetCommit(v *GetRepositoryTagResponseBodyResultCommit) *GetRepositoryTagResponseBodyResult {
	s.Commit = v
	return s
}

func (s *GetRepositoryTagResponseBodyResult) SetId(v string) *GetRepositoryTagResponseBodyResult {
	s.Id = &v
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

func (s *GetRepositoryTagResponseBodyResult) SetSignature(v *GetRepositoryTagResponseBodyResultSignature) *GetRepositoryTagResponseBodyResult {
	s.Signature = v
	return s
}

type GetRepositoryTagResponseBodyResultCommit struct {
	AuthorEmail    *string                                            `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string                                            `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	AuthoredDate   *string                                            `json:"AuthoredDate,omitempty" xml:"AuthoredDate,omitempty"`
	CommittedDate  *string                                            `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string                                            `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string                                            `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string                                            `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string                                          `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string                                            `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Signature      *GetRepositoryTagResponseBodyResultCommitSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Title          *string                                            `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetRepositoryTagResponseBodyResultCommit) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthorEmail(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthorName(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetAuthoredDate(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommittedDate(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommitterEmail(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCommitterName(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetCreatedAt(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetId(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetMessage(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetParentIds(v []*string) *GetRepositoryTagResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetShortId(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetSignature(v *GetRepositoryTagResponseBodyResultCommitSignature) *GetRepositoryTagResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *GetRepositoryTagResponseBodyResultCommit) SetTitle(v string) *GetRepositoryTagResponseBodyResultCommit {
	s.Title = &v
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

type GetRepositoryTagResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRepositoryTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRepositoryTagResponse) SetStatusCode(v int32) *GetRepositoryTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryTagResponse) SetBody(v *GetRepositoryTagResponseBody) *GetRepositoryTagResponse {
	s.Body = v
	return s
}

type GetRepositoryTagV2Request struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	TagName        *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetRepositoryTagV2Request) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagV2Request) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagV2Request) SetAccessToken(v string) *GetRepositoryTagV2Request {
	s.AccessToken = &v
	return s
}

func (s *GetRepositoryTagV2Request) SetOrganizationId(v string) *GetRepositoryTagV2Request {
	s.OrganizationId = &v
	return s
}

func (s *GetRepositoryTagV2Request) SetTagName(v string) *GetRepositoryTagV2Request {
	s.TagName = &v
	return s
}

type GetRepositoryTagV2ResponseBody struct {
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetRepositoryTagV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRepositoryTagV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagV2ResponseBody) SetErrorCode(v string) *GetRepositoryTagV2ResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBody) SetErrorMessage(v string) *GetRepositoryTagV2ResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBody) SetRequestId(v string) *GetRepositoryTagV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBody) SetResult(v *GetRepositoryTagV2ResponseBodyResult) *GetRepositoryTagV2ResponseBody {
	s.Result = v
	return s
}

func (s *GetRepositoryTagV2ResponseBody) SetSuccess(v bool) *GetRepositoryTagV2ResponseBody {
	s.Success = &v
	return s
}

type GetRepositoryTagV2ResponseBodyResult struct {
	Commit    *GetRepositoryTagV2ResponseBodyResultCommit    `json:"Commit,omitempty" xml:"Commit,omitempty" type:"Struct"`
	Id        *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Name      *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Signature *GetRepositoryTagV2ResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
}

func (s GetRepositoryTagV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagV2ResponseBodyResult) SetCommit(v *GetRepositoryTagV2ResponseBodyResultCommit) *GetRepositoryTagV2ResponseBodyResult {
	s.Commit = v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResult) SetId(v string) *GetRepositoryTagV2ResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResult) SetMessage(v string) *GetRepositoryTagV2ResponseBodyResult {
	s.Message = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResult) SetName(v string) *GetRepositoryTagV2ResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResult) SetSignature(v *GetRepositoryTagV2ResponseBodyResultSignature) *GetRepositoryTagV2ResponseBodyResult {
	s.Signature = v
	return s
}

type GetRepositoryTagV2ResponseBodyResultCommit struct {
	AuthorEmail    *string                                              `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string                                              `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	AuthoredDate   *string                                              `json:"AuthoredDate,omitempty" xml:"AuthoredDate,omitempty"`
	CommittedDate  *string                                              `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string                                              `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string                                              `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string                                              `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string                                            `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string                                              `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Signature      *GetRepositoryTagV2ResponseBodyResultCommitSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Title          *string                                              `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetRepositoryTagV2ResponseBodyResultCommit) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagV2ResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetAuthorEmail(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetAuthorName(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetAuthoredDate(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetCommittedDate(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetCommitterEmail(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetCommitterName(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetCreatedAt(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetId(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetMessage(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetParentIds(v []*string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetShortId(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetSignature(v *GetRepositoryTagV2ResponseBodyResultCommitSignature) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommit) SetTitle(v string) *GetRepositoryTagV2ResponseBodyResultCommit {
	s.Title = &v
	return s
}

type GetRepositoryTagV2ResponseBodyResultCommitSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s GetRepositoryTagV2ResponseBodyResultCommitSignature) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagV2ResponseBodyResultCommitSignature) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagV2ResponseBodyResultCommitSignature) SetGpgKeyId(v string) *GetRepositoryTagV2ResponseBodyResultCommitSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultCommitSignature) SetVerificationStatus(v string) *GetRepositoryTagV2ResponseBodyResultCommitSignature {
	s.VerificationStatus = &v
	return s
}

type GetRepositoryTagV2ResponseBodyResultSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s GetRepositoryTagV2ResponseBodyResultSignature) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagV2ResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagV2ResponseBodyResultSignature) SetGpgKeyId(v string) *GetRepositoryTagV2ResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *GetRepositoryTagV2ResponseBodyResultSignature) SetVerificationStatus(v string) *GetRepositoryTagV2ResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

type GetRepositoryTagV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRepositoryTagV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRepositoryTagV2Response) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryTagV2Response) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagV2Response) SetHeaders(v map[string]*string) *GetRepositoryTagV2Response {
	s.Headers = v
	return s
}

func (s *GetRepositoryTagV2Response) SetStatusCode(v int32) *GetRepositoryTagV2Response {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryTagV2Response) SetBody(v *GetRepositoryTagV2ResponseBody) *GetRepositoryTagV2Response {
	s.Body = v
	return s
}

type GetUserInfoRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
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

func (s *GetUserInfoRequest) SetOrganizationId(v string) *GetUserInfoRequest {
	s.OrganizationId = &v
	return s
}

type GetUserInfoResponseBody struct {
	ErrorCode    *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *GetUserInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBody) SetErrorCode(v string) *GetUserInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserInfoResponseBody) SetErrorMessage(v string) *GetUserInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserInfoResponseBody) SetRequestId(v string) *GetUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInfoResponseBody) SetResult(v *GetUserInfoResponseBodyResult) *GetUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetUserInfoResponseBody) SetSuccess(v bool) *GetUserInfoResponseBody {
	s.Success = &v
	return s
}

type GetUserInfoResponseBodyResult struct {
	AvatarUrl      *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBodyResult) SetAvatarUrl(v string) *GetUserInfoResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetEmail(v string) *GetUserInfoResponseBodyResult {
	s.Email = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetExternalUserId(v string) *GetUserInfoResponseBodyResult {
	s.ExternalUserId = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetId(v int64) *GetUserInfoResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetName(v string) *GetUserInfoResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetUserInfoResponseBodyResult) SetUsername(v string) *GetUserInfoResponseBodyResult {
	s.Username = &v
	return s
}

type GetUserInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserInfoResponse) SetStatusCode(v int32) *GetUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInfoResponse) SetBody(v *GetUserInfoResponseBody) *GetUserInfoResponse {
	s.Body = v
	return s
}

type IsSlsUserAuthrizedCodeupRequest struct {
	OrganizationId  *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	AliyunSubUserId *string `json:"aliyunSubUserId,omitempty" xml:"aliyunSubUserId,omitempty"`
	AliyunUserId    *string `json:"aliyunUserId,omitempty" xml:"aliyunUserId,omitempty"`
	CodeupProjectId *int64  `json:"codeupProjectId,omitempty" xml:"codeupProjectId,omitempty"`
	SlsLogStore     *string `json:"slsLogStore,omitempty" xml:"slsLogStore,omitempty"`
	SlsProject      *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
}

func (s IsSlsUserAuthrizedCodeupRequest) String() string {
	return tea.Prettify(s)
}

func (s IsSlsUserAuthrizedCodeupRequest) GoString() string {
	return s.String()
}

func (s *IsSlsUserAuthrizedCodeupRequest) SetOrganizationId(v string) *IsSlsUserAuthrizedCodeupRequest {
	s.OrganizationId = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupRequest) SetAliyunSubUserId(v string) *IsSlsUserAuthrizedCodeupRequest {
	s.AliyunSubUserId = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupRequest) SetAliyunUserId(v string) *IsSlsUserAuthrizedCodeupRequest {
	s.AliyunUserId = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupRequest) SetCodeupProjectId(v int64) *IsSlsUserAuthrizedCodeupRequest {
	s.CodeupProjectId = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupRequest) SetSlsLogStore(v string) *IsSlsUserAuthrizedCodeupRequest {
	s.SlsLogStore = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupRequest) SetSlsProject(v string) *IsSlsUserAuthrizedCodeupRequest {
	s.SlsProject = &v
	return s
}

type IsSlsUserAuthrizedCodeupResponseBody struct {
	ErrorCode    *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *IsSlsUserAuthrizedCodeupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *string                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IsSlsUserAuthrizedCodeupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsSlsUserAuthrizedCodeupResponseBody) GoString() string {
	return s.String()
}

func (s *IsSlsUserAuthrizedCodeupResponseBody) SetErrorCode(v string) *IsSlsUserAuthrizedCodeupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupResponseBody) SetErrorMessage(v string) *IsSlsUserAuthrizedCodeupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupResponseBody) SetRequestId(v string) *IsSlsUserAuthrizedCodeupResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupResponseBody) SetResult(v *IsSlsUserAuthrizedCodeupResponseBodyResult) *IsSlsUserAuthrizedCodeupResponseBody {
	s.Result = v
	return s
}

func (s *IsSlsUserAuthrizedCodeupResponseBody) SetSuccess(v string) *IsSlsUserAuthrizedCodeupResponseBody {
	s.Success = &v
	return s
}

type IsSlsUserAuthrizedCodeupResponseBodyResult struct {
	Authrized *bool `json:"authrized,omitempty" xml:"authrized,omitempty"`
}

func (s IsSlsUserAuthrizedCodeupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IsSlsUserAuthrizedCodeupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IsSlsUserAuthrizedCodeupResponseBodyResult) SetAuthrized(v bool) *IsSlsUserAuthrizedCodeupResponseBodyResult {
	s.Authrized = &v
	return s
}

type IsSlsUserAuthrizedCodeupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IsSlsUserAuthrizedCodeupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IsSlsUserAuthrizedCodeupResponse) String() string {
	return tea.Prettify(s)
}

func (s IsSlsUserAuthrizedCodeupResponse) GoString() string {
	return s.String()
}

func (s *IsSlsUserAuthrizedCodeupResponse) SetHeaders(v map[string]*string) *IsSlsUserAuthrizedCodeupResponse {
	s.Headers = v
	return s
}

func (s *IsSlsUserAuthrizedCodeupResponse) SetStatusCode(v int32) *IsSlsUserAuthrizedCodeupResponse {
	s.StatusCode = &v
	return s
}

func (s *IsSlsUserAuthrizedCodeupResponse) SetBody(v *IsSlsUserAuthrizedCodeupResponseBody) *IsSlsUserAuthrizedCodeupResponse {
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
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponseBody) SetErrorCode(v string) *ListGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetErrorMessage(v string) *ListGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetRequestId(v string) *ListGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetResult(v []*ListGroupMemberResponseBodyResult) *ListGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *ListGroupMemberResponseBody) SetSuccess(v bool) *ListGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetTotal(v int64) *ListGroupMemberResponseBody {
	s.Total = &v
	return s
}

type ListGroupMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponseBodyResult) SetAccessLevel(v int32) *ListGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetAvatarUrl(v string) *ListGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetEmail(v string) *ListGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetExternUserId(v string) *ListGroupMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetId(v int64) *ListGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListGroupMemberResponseBodyResult) SetState(v string) *ListGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

type ListGroupMemberResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListGroupMemberResponse) SetStatusCode(v int32) *ListGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupMemberResponse) SetBody(v *ListGroupMemberResponseBody) *ListGroupMemberResponse {
	s.Body = v
	return s
}

type ListGroupRepositoriesRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	IsMember       *bool   `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
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

func (s *ListGroupRepositoriesRequest) SetIsMember(v bool) *ListGroupRepositoriesRequest {
	s.IsMember = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetOrganizationId(v string) *ListGroupRepositoriesRequest {
	s.OrganizationId = &v
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

func (s *ListGroupRepositoriesRequest) SetSearch(v string) *ListGroupRepositoriesRequest {
	s.Search = &v
	return s
}

func (s *ListGroupRepositoriesRequest) SetSubUserId(v string) *ListGroupRepositoriesRequest {
	s.SubUserId = &v
	return s
}

type ListGroupRepositoriesResponseBody struct {
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListGroupRepositoriesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListGroupRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesResponseBody) SetErrorCode(v string) *ListGroupRepositoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetErrorMessage(v string) *ListGroupRepositoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetRequestId(v string) *ListGroupRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetResult(v []*ListGroupRepositoriesResponseBodyResult) *ListGroupRepositoriesResponseBody {
	s.Result = v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetSuccess(v bool) *ListGroupRepositoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupRepositoriesResponseBody) SetTotal(v int64) *ListGroupRepositoriesResponseBody {
	s.Total = &v
	return s
}

type ListGroupRepositoriesResponseBodyResult struct {
	Archive           *bool   `json:"Archive,omitempty" xml:"Archive,omitempty"`
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CreatorId         *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	HttpCloneUrl      *string `json:"HttpCloneUrl,omitempty" xml:"HttpCloneUrl,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ImportStatus      *string `json:"ImportStatus,omitempty" xml:"ImportStatus,omitempty"`
	LastActivityAt    *string `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	NamespaceId       *int64  `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	SshCloneUrl       *string `json:"SshCloneUrl,omitempty" xml:"SshCloneUrl,omitempty"`
	UpdatedAt         *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	VisibilityLevel   *int32  `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s ListGroupRepositoriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRepositoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupRepositoriesResponseBodyResult) SetArchive(v bool) *ListGroupRepositoriesResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetCreatedAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetCreatorId(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetHttpCloneUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.HttpCloneUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetId(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetImportStatus(v string) *ListGroupRepositoriesResponseBodyResult {
	s.ImportStatus = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetLastActivityAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetName(v string) *ListGroupRepositoriesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetNameWithNamespace(v string) *ListGroupRepositoriesResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetNamespaceId(v int64) *ListGroupRepositoriesResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetPath(v string) *ListGroupRepositoriesResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetPathWithNamespace(v string) *ListGroupRepositoriesResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetSshCloneUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.SshCloneUrl = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetUpdatedAt(v string) *ListGroupRepositoriesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetVisibilityLevel(v int32) *ListGroupRepositoriesResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListGroupRepositoriesResponseBodyResult) SetWebUrl(v string) *ListGroupRepositoriesResponseBodyResult {
	s.WebUrl = &v
	return s
}

type ListGroupRepositoriesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListGroupRepositoriesResponse) SetStatusCode(v int32) *ListGroupRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupRepositoriesResponse) SetBody(v *ListGroupRepositoriesResponseBody) *ListGroupRepositoriesResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	AccessToken     *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	IncludePersonal *bool   `json:"IncludePersonal,omitempty" xml:"IncludePersonal,omitempty"`
	OrganizationId  *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page            *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search          *string `json:"Search,omitempty" xml:"Search,omitempty"`
	SubUserId       *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
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

func (s *ListGroupsRequest) SetIncludePersonal(v bool) *ListGroupsRequest {
	s.IncludePersonal = &v
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

func (s *ListGroupsRequest) SetSearch(v string) *ListGroupsRequest {
	s.Search = &v
	return s
}

func (s *ListGroupsRequest) SetSubUserId(v string) *ListGroupsRequest {
	s.SubUserId = &v
	return s
}

type ListGroupsResponseBody struct {
	ErrorCode    *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListGroupsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetErrorCode(v string) *ListGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGroupsResponseBody) SetErrorMessage(v string) *ListGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetResult(v []*ListGroupsResponseBodyResult) *ListGroupsResponseBody {
	s.Result = v
	return s
}

func (s *ListGroupsResponseBody) SetSuccess(v bool) *ListGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotal(v int64) *ListGroupsResponseBody {
	s.Total = &v
	return s
}

type ListGroupsResponseBodyResult struct {
	AccessLevel       *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParentId          *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedAt         *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s ListGroupsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyResult) SetAccessLevel(v int32) *ListGroupsResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetCreatedAt(v string) *ListGroupsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetDescription(v string) *ListGroupsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetId(v int64) *ListGroupsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetName(v string) *ListGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetNameWithNamespace(v string) *ListGroupsResponseBodyResult {
	s.NameWithNamespace = &v
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

func (s *ListGroupsResponseBodyResult) SetPath(v string) *ListGroupsResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetPathWithNamespace(v string) *ListGroupsResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetType(v string) *ListGroupsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetUpdatedAt(v string) *ListGroupsResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetVisibilityLevel(v string) *ListGroupsResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListGroupsResponseBodyResult) SetWebUrl(v string) *ListGroupsResponseBodyResult {
	s.WebUrl = &v
	return s
}

type ListGroupsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListGroupsResponse) SetStatusCode(v int32) *ListGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListMergeRequestCommentsRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	FromCommit     *string `json:"FromCommit,omitempty" xml:"FromCommit,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	ToCommit       *string `json:"ToCommit,omitempty" xml:"ToCommit,omitempty"`
}

func (s ListMergeRequestCommentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestCommentsRequest) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsRequest) SetAccessToken(v string) *ListMergeRequestCommentsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetFromCommit(v string) *ListMergeRequestCommentsRequest {
	s.FromCommit = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetOrganizationId(v string) *ListMergeRequestCommentsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetToCommit(v string) *ListMergeRequestCommentsRequest {
	s.ToCommit = &v
	return s
}

type ListMergeRequestCommentsResponseBody struct {
	ErrorCode    *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListMergeRequestCommentsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMergeRequestCommentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBody) SetErrorCode(v string) *ListMergeRequestCommentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetErrorMessage(v string) *ListMergeRequestCommentsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetRequestId(v string) *ListMergeRequestCommentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetResult(v []*ListMergeRequestCommentsResponseBodyResult) *ListMergeRequestCommentsResponseBody {
	s.Result = v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetSuccess(v bool) *ListMergeRequestCommentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBody) SetTotal(v int64) *ListMergeRequestCommentsResponseBody {
	s.Total = &v
	return s
}

type ListMergeRequestCommentsResponseBodyResult struct {
	Author       *ListMergeRequestCommentsResponseBodyResultAuthor `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	Closed       *int32                                            `json:"Closed,omitempty" xml:"Closed,omitempty"`
	CreatedAt    *string                                           `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id           *int64                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDraft      *bool                                             `json:"IsDraft,omitempty" xml:"IsDraft,omitempty"`
	Line         *int64                                            `json:"Line,omitempty" xml:"Line,omitempty"`
	Note         *string                                           `json:"Note,omitempty" xml:"Note,omitempty"`
	OutDated     *bool                                             `json:"OutDated,omitempty" xml:"OutDated,omitempty"`
	ParentNoteId *int64                                            `json:"ParentNoteId,omitempty" xml:"ParentNoteId,omitempty"`
	Path         *string                                           `json:"Path,omitempty" xml:"Path,omitempty"`
	ProjectId    *int64                                            `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RangeContext *string                                           `json:"RangeContext,omitempty" xml:"RangeContext,omitempty"`
	Side         *string                                           `json:"Side,omitempty" xml:"Side,omitempty"`
	UpdatedAt    *string                                           `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetAuthor(v *ListMergeRequestCommentsResponseBodyResultAuthor) *ListMergeRequestCommentsResponseBodyResult {
	s.Author = v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetClosed(v int32) *ListMergeRequestCommentsResponseBodyResult {
	s.Closed = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetCreatedAt(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetId(v int64) *ListMergeRequestCommentsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetIsDraft(v bool) *ListMergeRequestCommentsResponseBodyResult {
	s.IsDraft = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetLine(v int64) *ListMergeRequestCommentsResponseBodyResult {
	s.Line = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetNote(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.Note = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetOutDated(v bool) *ListMergeRequestCommentsResponseBodyResult {
	s.OutDated = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetParentNoteId(v int64) *ListMergeRequestCommentsResponseBodyResult {
	s.ParentNoteId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetPath(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetProjectId(v int64) *ListMergeRequestCommentsResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetRangeContext(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.RangeContext = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetSide(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.Side = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResult) SetUpdatedAt(v string) *ListMergeRequestCommentsResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

type ListMergeRequestCommentsResponseBodyResultAuthor struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMergeRequestCommentsResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestCommentsResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetAvatarUrl(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetEmail(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.Email = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetExternUserId(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetId(v int64) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *ListMergeRequestCommentsResponseBodyResultAuthor) SetName(v string) *ListMergeRequestCommentsResponseBodyResultAuthor {
	s.Name = &v
	return s
}

type ListMergeRequestCommentsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMergeRequestCommentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMergeRequestCommentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestCommentsResponse) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsResponse) SetHeaders(v map[string]*string) *ListMergeRequestCommentsResponse {
	s.Headers = v
	return s
}

func (s *ListMergeRequestCommentsResponse) SetStatusCode(v int32) *ListMergeRequestCommentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMergeRequestCommentsResponse) SetBody(v *ListMergeRequestCommentsResponseBody) *ListMergeRequestCommentsResponse {
	s.Body = v
	return s
}

type ListMergeRequestsRequest struct {
	AccessToken            *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AfterDate              *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	AssigneeCodeupIdList   *string `json:"AssigneeCodeupIdList,omitempty" xml:"AssigneeCodeupIdList,omitempty"`
	AssigneeIdList         *string `json:"AssigneeIdList,omitempty" xml:"AssigneeIdList,omitempty"`
	AuthorCodeupIdList     *string `json:"AuthorCodeupIdList,omitempty" xml:"AuthorCodeupIdList,omitempty"`
	AuthorIdList           *string `json:"AuthorIdList,omitempty" xml:"AuthorIdList,omitempty"`
	BeforeDate             *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	GroupIdList            *string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty"`
	Order                  *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrganizationId         *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page                   *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize               *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectIdList          *string `json:"ProjectIdList,omitempty" xml:"ProjectIdList,omitempty"`
	Search                 *string `json:"Search,omitempty" xml:"Search,omitempty"`
	State                  *string `json:"State,omitempty" xml:"State,omitempty"`
	SubscriberCodeupIdList *string `json:"SubscriberCodeupIdList,omitempty" xml:"SubscriberCodeupIdList,omitempty"`
}

func (s ListMergeRequestsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsRequest) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsRequest) SetAccessToken(v string) *ListMergeRequestsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListMergeRequestsRequest) SetAfterDate(v string) *ListMergeRequestsRequest {
	s.AfterDate = &v
	return s
}

func (s *ListMergeRequestsRequest) SetAssigneeCodeupIdList(v string) *ListMergeRequestsRequest {
	s.AssigneeCodeupIdList = &v
	return s
}

func (s *ListMergeRequestsRequest) SetAssigneeIdList(v string) *ListMergeRequestsRequest {
	s.AssigneeIdList = &v
	return s
}

func (s *ListMergeRequestsRequest) SetAuthorCodeupIdList(v string) *ListMergeRequestsRequest {
	s.AuthorCodeupIdList = &v
	return s
}

func (s *ListMergeRequestsRequest) SetAuthorIdList(v string) *ListMergeRequestsRequest {
	s.AuthorIdList = &v
	return s
}

func (s *ListMergeRequestsRequest) SetBeforeDate(v string) *ListMergeRequestsRequest {
	s.BeforeDate = &v
	return s
}

func (s *ListMergeRequestsRequest) SetGroupIdList(v string) *ListMergeRequestsRequest {
	s.GroupIdList = &v
	return s
}

func (s *ListMergeRequestsRequest) SetOrder(v string) *ListMergeRequestsRequest {
	s.Order = &v
	return s
}

func (s *ListMergeRequestsRequest) SetOrganizationId(v string) *ListMergeRequestsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListMergeRequestsRequest) SetPage(v int64) *ListMergeRequestsRequest {
	s.Page = &v
	return s
}

func (s *ListMergeRequestsRequest) SetPageSize(v int64) *ListMergeRequestsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMergeRequestsRequest) SetProjectIdList(v string) *ListMergeRequestsRequest {
	s.ProjectIdList = &v
	return s
}

func (s *ListMergeRequestsRequest) SetSearch(v string) *ListMergeRequestsRequest {
	s.Search = &v
	return s
}

func (s *ListMergeRequestsRequest) SetState(v string) *ListMergeRequestsRequest {
	s.State = &v
	return s
}

func (s *ListMergeRequestsRequest) SetSubscriberCodeupIdList(v string) *ListMergeRequestsRequest {
	s.SubscriberCodeupIdList = &v
	return s
}

type ListMergeRequestsResponseBody struct {
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListMergeRequestsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMergeRequestsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBody) SetErrorCode(v string) *ListMergeRequestsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMergeRequestsResponseBody) SetErrorMessage(v string) *ListMergeRequestsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMergeRequestsResponseBody) SetRequestId(v string) *ListMergeRequestsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMergeRequestsResponseBody) SetResult(v []*ListMergeRequestsResponseBodyResult) *ListMergeRequestsResponseBody {
	s.Result = v
	return s
}

func (s *ListMergeRequestsResponseBody) SetSuccess(v bool) *ListMergeRequestsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMergeRequestsResponseBody) SetTotal(v int64) *ListMergeRequestsResponseBody {
	s.Total = &v
	return s
}

type ListMergeRequestsResponseBodyResult struct {
	AcceptedRevision   *string                                                `json:"AcceptedRevision,omitempty" xml:"AcceptedRevision,omitempty"`
	AheadCommitCount   *int32                                                 `json:"AheadCommitCount,omitempty" xml:"AheadCommitCount,omitempty"`
	ApproveCheckResult *ListMergeRequestsResponseBodyResultApproveCheckResult `json:"ApproveCheckResult,omitempty" xml:"ApproveCheckResult,omitempty" type:"Struct"`
	AssigneeList       []*ListMergeRequestsResponseBodyResultAssigneeList     `json:"AssigneeList,omitempty" xml:"AssigneeList,omitempty" type:"Repeated"`
	Author             *ListMergeRequestsResponseBodyResultAuthor             `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	BehindCommitCount  *int32                                                 `json:"BehindCommitCount,omitempty" xml:"BehindCommitCount,omitempty"`
	CreatedAt          *string                                                `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description        *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                 *int64                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsSupportMerge     *bool                                                  `json:"IsSupportMerge,omitempty" xml:"IsSupportMerge,omitempty"`
	MergeError         *string                                                `json:"MergeError,omitempty" xml:"MergeError,omitempty"`
	MergeStatus        *string                                                `json:"MergeStatus,omitempty" xml:"MergeStatus,omitempty"`
	MergeType          *string                                                `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	MergedRevision     *string                                                `json:"MergedRevision,omitempty" xml:"MergedRevision,omitempty"`
	NameWithNamespace  *string                                                `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	ProjectId          *int64                                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SourceBranch       *string                                                `json:"SourceBranch,omitempty" xml:"SourceBranch,omitempty"`
	State              *string                                                `json:"State,omitempty" xml:"State,omitempty"`
	TargetBranch       *string                                                `json:"TargetBranch,omitempty" xml:"TargetBranch,omitempty"`
	Title              *string                                                `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdatedAt          *string                                                `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	WebUrl             *string                                                `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s ListMergeRequestsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResult) SetAcceptedRevision(v string) *ListMergeRequestsResponseBodyResult {
	s.AcceptedRevision = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetAheadCommitCount(v int32) *ListMergeRequestsResponseBodyResult {
	s.AheadCommitCount = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetApproveCheckResult(v *ListMergeRequestsResponseBodyResultApproveCheckResult) *ListMergeRequestsResponseBodyResult {
	s.ApproveCheckResult = v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetAssigneeList(v []*ListMergeRequestsResponseBodyResultAssigneeList) *ListMergeRequestsResponseBodyResult {
	s.AssigneeList = v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetAuthor(v *ListMergeRequestsResponseBodyResultAuthor) *ListMergeRequestsResponseBodyResult {
	s.Author = v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetBehindCommitCount(v int32) *ListMergeRequestsResponseBodyResult {
	s.BehindCommitCount = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetCreatedAt(v string) *ListMergeRequestsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetDescription(v string) *ListMergeRequestsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetId(v int64) *ListMergeRequestsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetIsSupportMerge(v bool) *ListMergeRequestsResponseBodyResult {
	s.IsSupportMerge = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetMergeError(v string) *ListMergeRequestsResponseBodyResult {
	s.MergeError = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetMergeStatus(v string) *ListMergeRequestsResponseBodyResult {
	s.MergeStatus = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetMergeType(v string) *ListMergeRequestsResponseBodyResult {
	s.MergeType = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetMergedRevision(v string) *ListMergeRequestsResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetNameWithNamespace(v string) *ListMergeRequestsResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetProjectId(v int64) *ListMergeRequestsResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetSourceBranch(v string) *ListMergeRequestsResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetState(v string) *ListMergeRequestsResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetTargetBranch(v string) *ListMergeRequestsResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetTitle(v string) *ListMergeRequestsResponseBodyResult {
	s.Title = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetUpdatedAt(v string) *ListMergeRequestsResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResult) SetWebUrl(v string) *ListMergeRequestsResponseBodyResult {
	s.WebUrl = &v
	return s
}

type ListMergeRequestsResponseBodyResultApproveCheckResult struct {
	SatisfiedCheckResults   []*ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults   `json:"SatisfiedCheckResults,omitempty" xml:"SatisfiedCheckResults,omitempty" type:"Repeated"`
	TotalCheckResult        *string                                                                         `json:"TotalCheckResult,omitempty" xml:"TotalCheckResult,omitempty"`
	UnsatisfiedCheckResults []*ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults `json:"UnsatisfiedCheckResults,omitempty" xml:"UnsatisfiedCheckResults,omitempty" type:"Repeated"`
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResult) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResult) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResult) SetSatisfiedCheckResults(v []*ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) *ListMergeRequestsResponseBodyResultApproveCheckResult {
	s.SatisfiedCheckResults = v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResult) SetTotalCheckResult(v string) *ListMergeRequestsResponseBodyResultApproveCheckResult {
	s.TotalCheckResult = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResult) SetUnsatisfiedCheckResults(v []*ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) *ListMergeRequestsResponseBodyResultApproveCheckResult {
	s.UnsatisfiedCheckResults = v
	return s
}

type ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults struct {
	CheckName        *string                                                                                 `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                               `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                               `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckName(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckName = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckStatus(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckType(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetExtraUsers(v []*ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetSatisfiedItems(v []*string) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetId(v int64) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetName(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults struct {
	CheckName        *string                                                                                   `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                   `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                   `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                 `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                 `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckName(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckName = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckStatus(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckType(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetExtraUsers(v []*ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetSatisfiedItems(v []*string) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetId(v int64) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetName(v string) *ListMergeRequestsResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type ListMergeRequestsResponseBodyResultAssigneeList struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMergeRequestsResponseBodyResultAssigneeList) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultAssigneeList) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultAssigneeList) SetAvatarUrl(v string) *ListMergeRequestsResponseBodyResultAssigneeList {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAssigneeList) SetEmail(v string) *ListMergeRequestsResponseBodyResultAssigneeList {
	s.Email = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAssigneeList) SetExternUserId(v string) *ListMergeRequestsResponseBodyResultAssigneeList {
	s.ExternUserId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAssigneeList) SetId(v string) *ListMergeRequestsResponseBodyResultAssigneeList {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAssigneeList) SetName(v string) *ListMergeRequestsResponseBodyResultAssigneeList {
	s.Name = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAssigneeList) SetStatus(v string) *ListMergeRequestsResponseBodyResultAssigneeList {
	s.Status = &v
	return s
}

type ListMergeRequestsResponseBodyResultAuthor struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMergeRequestsResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetAvatarUrl(v string) *ListMergeRequestsResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetExternUserId(v string) *ListMergeRequestsResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetId(v int64) *ListMergeRequestsResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *ListMergeRequestsResponseBodyResultAuthor) SetName(v string) *ListMergeRequestsResponseBodyResultAuthor {
	s.Name = &v
	return s
}

type ListMergeRequestsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMergeRequestsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMergeRequestsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMergeRequestsResponse) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsResponse) SetHeaders(v map[string]*string) *ListMergeRequestsResponse {
	s.Headers = v
	return s
}

func (s *ListMergeRequestsResponse) SetStatusCode(v int32) *ListMergeRequestsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMergeRequestsResponse) SetBody(v *ListMergeRequestsResponseBody) *ListMergeRequestsResponse {
	s.Body = v
	return s
}

type ListOrganizationSecurityScoresRequest struct {
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
}

func (s ListOrganizationSecurityScoresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationSecurityScoresRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationSecurityScoresRequest) SetAccessToken(v string) *ListOrganizationSecurityScoresRequest {
	s.AccessToken = &v
	return s
}

type ListOrganizationSecurityScoresResponseBody struct {
	ErrorCode    *string                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListOrganizationSecurityScoresResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOrganizationSecurityScoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationSecurityScoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationSecurityScoresResponseBody) SetErrorCode(v string) *ListOrganizationSecurityScoresResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBody) SetErrorMessage(v string) *ListOrganizationSecurityScoresResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBody) SetRequestId(v string) *ListOrganizationSecurityScoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBody) SetResult(v []*ListOrganizationSecurityScoresResponseBodyResult) *ListOrganizationSecurityScoresResponseBody {
	s.Result = v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBody) SetSuccess(v bool) *ListOrganizationSecurityScoresResponseBody {
	s.Success = &v
	return s
}

type ListOrganizationSecurityScoresResponseBodyResult struct {
	Enable                    *bool                                                                      `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Id                        *int64                                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	OrganizationId            *string                                                                    `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	OrganizationSecurityScore *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore `json:"OrganizationSecurityScore,omitempty" xml:"OrganizationSecurityScore,omitempty" type:"Struct"`
}

func (s ListOrganizationSecurityScoresResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationSecurityScoresResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrganizationSecurityScoresResponseBodyResult) SetEnable(v bool) *ListOrganizationSecurityScoresResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBodyResult) SetId(v int64) *ListOrganizationSecurityScoresResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBodyResult) SetOrganizationId(v string) *ListOrganizationSecurityScoresResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBodyResult) SetOrganizationSecurityScore(v *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore) *ListOrganizationSecurityScoresResponseBodyResult {
	s.OrganizationSecurityScore = v
	return s
}

type ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore struct {
	AuthorityControlScore *int32  `json:"AuthorityControlScore,omitempty" xml:"AuthorityControlScore,omitempty"`
	CodeContentScore      *int32  `json:"CodeContentScore,omitempty" xml:"CodeContentScore,omitempty"`
	Level                 *string `json:"Level,omitempty" xml:"Level,omitempty"`
	MemberBehaviorScore   *int32  `json:"MemberBehaviorScore,omitempty" xml:"MemberBehaviorScore,omitempty"`
}

func (s ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore) GoString() string {
	return s.String()
}

func (s *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore) SetAuthorityControlScore(v int32) *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore {
	s.AuthorityControlScore = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore) SetCodeContentScore(v int32) *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore {
	s.CodeContentScore = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore) SetLevel(v string) *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore {
	s.Level = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore) SetMemberBehaviorScore(v int32) *ListOrganizationSecurityScoresResponseBodyResultOrganizationSecurityScore {
	s.MemberBehaviorScore = &v
	return s
}

type ListOrganizationSecurityScoresResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrganizationSecurityScoresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrganizationSecurityScoresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationSecurityScoresResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationSecurityScoresResponse) SetHeaders(v map[string]*string) *ListOrganizationSecurityScoresResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationSecurityScoresResponse) SetStatusCode(v int32) *ListOrganizationSecurityScoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationSecurityScoresResponse) SetBody(v *ListOrganizationSecurityScoresResponseBody) *ListOrganizationSecurityScoresResponse {
	s.Body = v
	return s
}

type ListOrganizationsRequest struct {
	AccessLevel    *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	MinAccessLevel *int32  `json:"MinAccessLevel,omitempty" xml:"MinAccessLevel,omitempty"`
}

func (s ListOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationsRequest) SetAccessLevel(v int32) *ListOrganizationsRequest {
	s.AccessLevel = &v
	return s
}

func (s *ListOrganizationsRequest) SetAccessToken(v string) *ListOrganizationsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListOrganizationsRequest) SetMinAccessLevel(v int32) *ListOrganizationsRequest {
	s.MinAccessLevel = &v
	return s
}

type ListOrganizationsResponseBody struct {
	ErrorCode    *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListOrganizationsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationsResponseBody) SetErrorCode(v string) *ListOrganizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetErrorMessage(v string) *ListOrganizationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetRequestId(v string) *ListOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationsResponseBody) SetResult(v []*ListOrganizationsResponseBodyResult) *ListOrganizationsResponseBody {
	s.Result = v
	return s
}

func (s *ListOrganizationsResponseBody) SetSuccess(v bool) *ListOrganizationsResponseBody {
	s.Success = &v
	return s
}

type ListOrganizationsResponseBodyResult struct {
	AccessLevel      *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	OrganizationRole *string `json:"OrganizationRole,omitempty" xml:"OrganizationRole,omitempty"`
}

func (s ListOrganizationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrganizationsResponseBodyResult) SetAccessLevel(v int32) *ListOrganizationsResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationId(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationName(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationName = &v
	return s
}

func (s *ListOrganizationsResponseBodyResult) SetOrganizationRole(v string) *ListOrganizationsResponseBodyResult {
	s.OrganizationRole = &v
	return s
}

type ListOrganizationsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListOrganizationsResponse) SetStatusCode(v int32) *ListOrganizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationsResponse) SetBody(v *ListOrganizationsResponseBody) *ListOrganizationsResponse {
	s.Body = v
	return s
}

type ListRepositoriesRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Archive        *bool   `json:"Archive,omitempty" xml:"Archive,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	Sort           *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
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

func (s *ListRepositoriesRequest) SetArchive(v bool) *ListRepositoriesRequest {
	s.Archive = &v
	return s
}

func (s *ListRepositoriesRequest) SetOrder(v string) *ListRepositoriesRequest {
	s.Order = &v
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

func (s *ListRepositoriesRequest) SetSearch(v string) *ListRepositoriesRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoriesRequest) SetSort(v string) *ListRepositoriesRequest {
	s.Sort = &v
	return s
}

type ListRepositoriesResponseBody struct {
	ErrorCode    *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoriesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBody) SetErrorCode(v int32) *ListRepositoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetErrorMessage(v string) *ListRepositoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetRequestId(v string) *ListRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetResult(v []*ListRepositoriesResponseBodyResult) *ListRepositoriesResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoriesResponseBody) SetSuccess(v bool) *ListRepositoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoriesResponseBody) SetTotal(v int64) *ListRepositoriesResponseBody {
	s.Total = &v
	return s
}

type ListRepositoriesResponseBodyResult struct {
	AccessLevel       *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	Archive           *bool   `json:"Archive,omitempty" xml:"Archive,omitempty"`
	AvatarUrl         *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DemoProjectStatus *bool   `json:"DemoProjectStatus,omitempty" xml:"DemoProjectStatus,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ImportStatus      *string `json:"ImportStatus,omitempty" xml:"ImportStatus,omitempty"`
	LastActivityAt    *string `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	NamespaceId       *int64  `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Star              *bool   `json:"Star,omitempty" xml:"Star,omitempty"`
	StarCount         *int64  `json:"StarCount,omitempty" xml:"StarCount,omitempty"`
	UpdatedAt         *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WebUrl            *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s ListRepositoriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoriesResponseBodyResult) SetAccessLevel(v int32) *ListRepositoriesResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetArchive(v bool) *ListRepositoriesResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetAvatarUrl(v string) *ListRepositoriesResponseBodyResult {
	s.AvatarUrl = &v
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

func (s *ListRepositoriesResponseBodyResult) SetDescription(v string) *ListRepositoriesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetId(v int64) *ListRepositoriesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetImportStatus(v string) *ListRepositoriesResponseBodyResult {
	s.ImportStatus = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetLastActivityAt(v string) *ListRepositoriesResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetName(v string) *ListRepositoriesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetNameWithNamespace(v string) *ListRepositoriesResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetNamespaceId(v int64) *ListRepositoriesResponseBodyResult {
	s.NamespaceId = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetPath(v string) *ListRepositoriesResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetPathWithNamespace(v string) *ListRepositoriesResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetStar(v bool) *ListRepositoriesResponseBodyResult {
	s.Star = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetStarCount(v int64) *ListRepositoriesResponseBodyResult {
	s.StarCount = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetUpdatedAt(v string) *ListRepositoriesResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetVisibilityLevel(v string) *ListRepositoriesResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *ListRepositoriesResponseBodyResult) SetWebUrl(v string) *ListRepositoriesResponseBodyResult {
	s.WebUrl = &v
	return s
}

type ListRepositoriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRepositoriesResponse) SetStatusCode(v int32) *ListRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoriesResponse) SetBody(v *ListRepositoriesResponseBody) *ListRepositoriesResponse {
	s.Body = v
	return s
}

type ListRepositoryBranchesRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
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

func (s *ListRepositoryBranchesRequest) SetSubUserId(v string) *ListRepositoryBranchesRequest {
	s.SubUserId = &v
	return s
}

type ListRepositoryBranchesResponseBody struct {
	ErrorCode    *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryBranchesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRepositoryBranchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBody) SetErrorCode(v string) *ListRepositoryBranchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetErrorMessage(v string) *ListRepositoryBranchesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetRequestId(v string) *ListRepositoryBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetResult(v []*ListRepositoryBranchesResponseBodyResult) *ListRepositoryBranchesResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetSuccess(v bool) *ListRepositoryBranchesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryBranchesResponseBody) SetTotal(v int64) *ListRepositoryBranchesResponseBody {
	s.Total = &v
	return s
}

type ListRepositoryBranchesResponseBodyResult struct {
	BranchName      *string                                             `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	CommitInfo      *ListRepositoryBranchesResponseBodyResultCommitInfo `json:"CommitInfo,omitempty" xml:"CommitInfo,omitempty" type:"Struct"`
	ProtectedBranch *bool                                               `json:"ProtectedBranch,omitempty" xml:"ProtectedBranch,omitempty"`
}

func (s ListRepositoryBranchesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBodyResult) SetBranchName(v string) *ListRepositoryBranchesResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) SetCommitInfo(v *ListRepositoryBranchesResponseBodyResultCommitInfo) *ListRepositoryBranchesResponseBodyResult {
	s.CommitInfo = v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResult) SetProtectedBranch(v bool) *ListRepositoryBranchesResponseBodyResult {
	s.ProtectedBranch = &v
	return s
}

type ListRepositoryBranchesResponseBodyResultCommitInfo struct {
	AuthorDate     *string   `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	AuthorEmail    *string   `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string   `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommittedDate  *string   `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string   `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string   `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string   `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListRepositoryBranchesResponseBodyResultCommitInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryBranchesResponseBodyResultCommitInfo) GoString() string {
	return s.String()
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetAuthorDate(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.AuthorDate = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetAuthorEmail(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetAuthorName(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetCommittedDate(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.CommittedDate = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetCommitterEmail(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetCommitterName(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetCreatedAt(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetId(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.Id = &v
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

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetShortId(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryBranchesResponseBodyResultCommitInfo) SetTitle(v string) *ListRepositoryBranchesResponseBodyResultCommitInfo {
	s.Title = &v
	return s
}

type ListRepositoryBranchesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRepositoryBranchesResponse) SetStatusCode(v int32) *ListRepositoryBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryBranchesResponse) SetBody(v *ListRepositoryBranchesResponseBody) *ListRepositoryBranchesResponse {
	s.Body = v
	return s
}

type ListRepositoryCodeRequest struct {
	OrganizationId *string                                  `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	FilePath       *ListRepositoryCodeRequestFilePath       `json:"FilePath,omitempty" xml:"FilePath,omitempty" type:"Struct"`
	IsCodeBlock    *bool                                    `json:"IsCodeBlock,omitempty" xml:"IsCodeBlock,omitempty"`
	KeyWord        *string                                  `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	Language       *string                                  `json:"Language,omitempty" xml:"Language,omitempty"`
	Order          *string                                  `json:"Order,omitempty" xml:"Order,omitempty"`
	Page           *int64                                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RepositoryPath *ListRepositoryCodeRequestRepositoryPath `json:"RepositoryPath,omitempty" xml:"RepositoryPath,omitempty" type:"Struct"`
	Scope          *string                                  `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Sort           *string                                  `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s ListRepositoryCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCodeRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryCodeRequest) SetOrganizationId(v string) *ListRepositoryCodeRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryCodeRequest) SetFilePath(v *ListRepositoryCodeRequestFilePath) *ListRepositoryCodeRequest {
	s.FilePath = v
	return s
}

func (s *ListRepositoryCodeRequest) SetIsCodeBlock(v bool) *ListRepositoryCodeRequest {
	s.IsCodeBlock = &v
	return s
}

func (s *ListRepositoryCodeRequest) SetKeyWord(v string) *ListRepositoryCodeRequest {
	s.KeyWord = &v
	return s
}

func (s *ListRepositoryCodeRequest) SetLanguage(v string) *ListRepositoryCodeRequest {
	s.Language = &v
	return s
}

func (s *ListRepositoryCodeRequest) SetOrder(v string) *ListRepositoryCodeRequest {
	s.Order = &v
	return s
}

func (s *ListRepositoryCodeRequest) SetPage(v int64) *ListRepositoryCodeRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryCodeRequest) SetPageSize(v int64) *ListRepositoryCodeRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryCodeRequest) SetRepositoryPath(v *ListRepositoryCodeRequestRepositoryPath) *ListRepositoryCodeRequest {
	s.RepositoryPath = v
	return s
}

func (s *ListRepositoryCodeRequest) SetScope(v string) *ListRepositoryCodeRequest {
	s.Scope = &v
	return s
}

func (s *ListRepositoryCodeRequest) SetSort(v string) *ListRepositoryCodeRequest {
	s.Sort = &v
	return s
}

type ListRepositoryCodeRequestFilePath struct {
	MatchType    *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRepositoryCodeRequestFilePath) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCodeRequestFilePath) GoString() string {
	return s.String()
}

func (s *ListRepositoryCodeRequestFilePath) SetMatchType(v string) *ListRepositoryCodeRequestFilePath {
	s.MatchType = &v
	return s
}

func (s *ListRepositoryCodeRequestFilePath) SetName(v string) *ListRepositoryCodeRequestFilePath {
	s.Name = &v
	return s
}

func (s *ListRepositoryCodeRequestFilePath) SetOperatorType(v string) *ListRepositoryCodeRequestFilePath {
	s.OperatorType = &v
	return s
}

func (s *ListRepositoryCodeRequestFilePath) SetValue(v string) *ListRepositoryCodeRequestFilePath {
	s.Value = &v
	return s
}

type ListRepositoryCodeRequestRepositoryPath struct {
	MatchType    *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRepositoryCodeRequestRepositoryPath) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCodeRequestRepositoryPath) GoString() string {
	return s.String()
}

func (s *ListRepositoryCodeRequestRepositoryPath) SetMatchType(v string) *ListRepositoryCodeRequestRepositoryPath {
	s.MatchType = &v
	return s
}

func (s *ListRepositoryCodeRequestRepositoryPath) SetName(v string) *ListRepositoryCodeRequestRepositoryPath {
	s.Name = &v
	return s
}

func (s *ListRepositoryCodeRequestRepositoryPath) SetOperatorType(v string) *ListRepositoryCodeRequestRepositoryPath {
	s.OperatorType = &v
	return s
}

func (s *ListRepositoryCodeRequestRepositoryPath) SetValue(v string) *ListRepositoryCodeRequestRepositoryPath {
	s.Value = &v
	return s
}

type ListRepositoryCodeResponseBody struct {
	ErrorCode    *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryCodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRepositoryCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryCodeResponseBody) SetErrorCode(v string) *ListRepositoryCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryCodeResponseBody) SetErrorMessage(v string) *ListRepositoryCodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryCodeResponseBody) SetRequestId(v string) *ListRepositoryCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryCodeResponseBody) SetResult(v []*ListRepositoryCodeResponseBodyResult) *ListRepositoryCodeResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryCodeResponseBody) SetSuccess(v bool) *ListRepositoryCodeResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryCodeResponseBody) SetTotal(v int64) *ListRepositoryCodeResponseBody {
	s.Total = &v
	return s
}

type ListRepositoryCodeResponseBodyResult struct {
	DocId            *string                                     `json:"DocId,omitempty" xml:"DocId,omitempty"`
	HighlightTextMap map[string]interface{}                      `json:"HighlightTextMap,omitempty" xml:"HighlightTextMap,omitempty"`
	Source           *ListRepositoryCodeResponseBodyResultSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
}

func (s ListRepositoryCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryCodeResponseBodyResult) SetDocId(v string) *ListRepositoryCodeResponseBodyResult {
	s.DocId = &v
	return s
}

func (s *ListRepositoryCodeResponseBodyResult) SetHighlightTextMap(v map[string]interface{}) *ListRepositoryCodeResponseBodyResult {
	s.HighlightTextMap = v
	return s
}

func (s *ListRepositoryCodeResponseBodyResult) SetSource(v *ListRepositoryCodeResponseBodyResultSource) *ListRepositoryCodeResponseBodyResult {
	s.Source = v
	return s
}

type ListRepositoryCodeResponseBodyResultSource struct {
	Branch         *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	CheckinDate    *string `json:"CheckinDate,omitempty" xml:"CheckinDate,omitempty"`
	FileName       *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FilePath       *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	RepoPath       *string `json:"RepoPath,omitempty" xml:"RepoPath,omitempty"`
}

func (s ListRepositoryCodeResponseBodyResultSource) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCodeResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *ListRepositoryCodeResponseBodyResultSource) SetBranch(v string) *ListRepositoryCodeResponseBodyResultSource {
	s.Branch = &v
	return s
}

func (s *ListRepositoryCodeResponseBodyResultSource) SetCheckinDate(v string) *ListRepositoryCodeResponseBodyResultSource {
	s.CheckinDate = &v
	return s
}

func (s *ListRepositoryCodeResponseBodyResultSource) SetFileName(v string) *ListRepositoryCodeResponseBodyResultSource {
	s.FileName = &v
	return s
}

func (s *ListRepositoryCodeResponseBodyResultSource) SetFilePath(v string) *ListRepositoryCodeResponseBodyResultSource {
	s.FilePath = &v
	return s
}

func (s *ListRepositoryCodeResponseBodyResultSource) SetLanguage(v string) *ListRepositoryCodeResponseBodyResultSource {
	s.Language = &v
	return s
}

func (s *ListRepositoryCodeResponseBodyResultSource) SetOrganizationId(v string) *ListRepositoryCodeResponseBodyResultSource {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryCodeResponseBodyResultSource) SetRepoPath(v string) *ListRepositoryCodeResponseBodyResultSource {
	s.RepoPath = &v
	return s
}

type ListRepositoryCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCodeResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryCodeResponse) SetHeaders(v map[string]*string) *ListRepositoryCodeResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryCodeResponse) SetStatusCode(v int32) *ListRepositoryCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryCodeResponse) SetBody(v *ListRepositoryCodeResponseBody) *ListRepositoryCodeResponse {
	s.Body = v
	return s
}

type ListRepositoryCommitDiffRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	ContextLine    *int32  `json:"ContextLine,omitempty" xml:"ContextLine,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRepositoryCommitDiffRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitDiffRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffRequest) SetAccessToken(v string) *ListRepositoryCommitDiffRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryCommitDiffRequest) SetContextLine(v int32) *ListRepositoryCommitDiffRequest {
	s.ContextLine = &v
	return s
}

func (s *ListRepositoryCommitDiffRequest) SetOrganizationId(v string) *ListRepositoryCommitDiffRequest {
	s.OrganizationId = &v
	return s
}

type ListRepositoryCommitDiffResponseBody struct {
	ErrorCode    *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryCommitDiffResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRepositoryCommitDiffResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitDiffResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponseBody) SetErrorCode(v string) *ListRepositoryCommitDiffResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetErrorMessage(v string) *ListRepositoryCommitDiffResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetRequestId(v string) *ListRepositoryCommitDiffResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetResult(v []*ListRepositoryCommitDiffResponseBodyResult) *ListRepositoryCommitDiffResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryCommitDiffResponseBody) SetSuccess(v bool) *ListRepositoryCommitDiffResponseBody {
	s.Success = &v
	return s
}

type ListRepositoryCommitDiffResponseBodyResult struct {
	AMode       *string `json:"AMode,omitempty" xml:"AMode,omitempty"`
	BMode       *string `json:"BMode,omitempty" xml:"BMode,omitempty"`
	DeletedFile *bool   `json:"DeletedFile,omitempty" xml:"DeletedFile,omitempty"`
	Diff        *string `json:"Diff,omitempty" xml:"Diff,omitempty"`
	IsBinary    *bool   `json:"IsBinary,omitempty" xml:"IsBinary,omitempty"`
	IsNewLfs    *bool   `json:"IsNewLfs,omitempty" xml:"IsNewLfs,omitempty"`
	IsOldLfs    *bool   `json:"IsOldLfs,omitempty" xml:"IsOldLfs,omitempty"`
	NewFile     *bool   `json:"NewFile,omitempty" xml:"NewFile,omitempty"`
	NewId       *string `json:"NewId,omitempty" xml:"NewId,omitempty"`
	NewPath     *string `json:"NewPath,omitempty" xml:"NewPath,omitempty"`
	OldId       *string `json:"OldId,omitempty" xml:"OldId,omitempty"`
	OldPath     *string `json:"OldPath,omitempty" xml:"OldPath,omitempty"`
	RenamedFile *bool   `json:"RenamedFile,omitempty" xml:"RenamedFile,omitempty"`
}

func (s ListRepositoryCommitDiffResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitDiffResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetAMode(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.AMode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetBMode(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.BMode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetDeletedFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.DeletedFile = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetDiff(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.Diff = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsBinary(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsBinary = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsNewLfs(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsNewLfs = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetIsOldLfs(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.IsOldLfs = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewFile = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewId(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetNewPath(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.NewPath = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetOldId(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.OldId = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetOldPath(v string) *ListRepositoryCommitDiffResponseBodyResult {
	s.OldPath = &v
	return s
}

func (s *ListRepositoryCommitDiffResponseBodyResult) SetRenamedFile(v bool) *ListRepositoryCommitDiffResponseBodyResult {
	s.RenamedFile = &v
	return s
}

type ListRepositoryCommitDiffResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryCommitDiffResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryCommitDiffResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitDiffResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffResponse) SetHeaders(v map[string]*string) *ListRepositoryCommitDiffResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryCommitDiffResponse) SetStatusCode(v int32) *ListRepositoryCommitDiffResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryCommitDiffResponse) SetBody(v *ListRepositoryCommitDiffResponseBody) *ListRepositoryCommitDiffResponse {
	s.Body = v
	return s
}

type ListRepositoryCommitsRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RefName        *string `json:"RefName,omitempty" xml:"RefName,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	ShowSignature  *bool   `json:"ShowSignature,omitempty" xml:"ShowSignature,omitempty"`
}

func (s ListRepositoryCommitsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitsRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsRequest) SetAccessToken(v string) *ListRepositoryCommitsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetOrganizationId(v string) *ListRepositoryCommitsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetPage(v int64) *ListRepositoryCommitsRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetPageSize(v int64) *ListRepositoryCommitsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetPath(v string) *ListRepositoryCommitsRequest {
	s.Path = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetRefName(v string) *ListRepositoryCommitsRequest {
	s.RefName = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetSearch(v string) *ListRepositoryCommitsRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetShowSignature(v bool) *ListRepositoryCommitsRequest {
	s.ShowSignature = &v
	return s
}

type ListRepositoryCommitsResponseBody struct {
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryCommitsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRepositoryCommitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponseBody) SetErrorCode(v string) *ListRepositoryCommitsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetErrorMessage(v string) *ListRepositoryCommitsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetRequestId(v string) *ListRepositoryCommitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetResult(v []*ListRepositoryCommitsResponseBodyResult) *ListRepositoryCommitsResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetSuccess(v bool) *ListRepositoryCommitsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryCommitsResponseBody) SetTotal(v int64) *ListRepositoryCommitsResponseBody {
	s.Total = &v
	return s
}

type ListRepositoryCommitsResponseBodyResult struct {
	AuthorDate     *string                                           `json:"AuthorDate,omitempty" xml:"AuthorDate,omitempty"`
	AuthorEmail    *string                                           `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string                                           `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	CommittedDate  *string                                           `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string                                           `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string                                           `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string                                           `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string                                         `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string                                           `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Signature      *ListRepositoryCommitsResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Title          *string                                           `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListRepositoryCommitsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponseBodyResult) SetAuthorDate(v string) *ListRepositoryCommitsResponseBodyResult {
	s.AuthorDate = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetAuthorEmail(v string) *ListRepositoryCommitsResponseBodyResult {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetAuthorName(v string) *ListRepositoryCommitsResponseBodyResult {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCommittedDate(v string) *ListRepositoryCommitsResponseBodyResult {
	s.CommittedDate = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCommitterEmail(v string) *ListRepositoryCommitsResponseBodyResult {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCommitterName(v string) *ListRepositoryCommitsResponseBodyResult {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetCreatedAt(v string) *ListRepositoryCommitsResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetId(v string) *ListRepositoryCommitsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetMessage(v string) *ListRepositoryCommitsResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetParentIds(v []*string) *ListRepositoryCommitsResponseBodyResult {
	s.ParentIds = v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetShortId(v string) *ListRepositoryCommitsResponseBodyResult {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetSignature(v *ListRepositoryCommitsResponseBodyResultSignature) *ListRepositoryCommitsResponseBodyResult {
	s.Signature = v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResult) SetTitle(v string) *ListRepositoryCommitsResponseBodyResult {
	s.Title = &v
	return s
}

type ListRepositoryCommitsResponseBodyResultSignature struct {
	GpgKeyId           *string `json:"GpgKeyId,omitempty" xml:"GpgKeyId,omitempty"`
	VerificationStatus *string `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s ListRepositoryCommitsResponseBodyResultSignature) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitsResponseBodyResultSignature) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponseBodyResultSignature) SetGpgKeyId(v string) *ListRepositoryCommitsResponseBodyResultSignature {
	s.GpgKeyId = &v
	return s
}

func (s *ListRepositoryCommitsResponseBodyResultSignature) SetVerificationStatus(v string) *ListRepositoryCommitsResponseBodyResultSignature {
	s.VerificationStatus = &v
	return s
}

type ListRepositoryCommitsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryCommitsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryCommitsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryCommitsResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsResponse) SetHeaders(v map[string]*string) *ListRepositoryCommitsResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryCommitsResponse) SetStatusCode(v int32) *ListRepositoryCommitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryCommitsResponse) SetBody(v *ListRepositoryCommitsResponseBody) *ListRepositoryCommitsResponse {
	s.Body = v
	return s
}

type ListRepositoryMemberRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
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

func (s *ListRepositoryMemberRequest) SetPage(v int64) *ListRepositoryMemberRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryMemberRequest) SetPageSize(v int64) *ListRepositoryMemberRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryMemberRequest) SetQuery(v string) *ListRepositoryMemberRequest {
	s.Query = &v
	return s
}

func (s *ListRepositoryMemberRequest) SetSubUserId(v string) *ListRepositoryMemberRequest {
	s.SubUserId = &v
	return s
}

type ListRepositoryMemberResponseBody struct {
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberResponseBody) SetErrorCode(v string) *ListRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetErrorMessage(v string) *ListRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetRequestId(v string) *ListRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetResult(v []*ListRepositoryMemberResponseBodyResult) *ListRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetSuccess(v bool) *ListRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryMemberResponseBody) SetTotal(v int64) *ListRepositoryMemberResponseBody {
	s.Total = &v
	return s
}

type ListRepositoryMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *ListRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *ListRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetEmail(v string) *ListRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetExternUserId(v string) *ListRepositoryMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetId(v int64) *ListRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetName(v string) *ListRepositoryMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetState(v string) *ListRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListRepositoryMemberResponseBodyResult) SetUsername(v string) *ListRepositoryMemberResponseBodyResult {
	s.Username = &v
	return s
}

type ListRepositoryMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRepositoryMemberResponse) SetStatusCode(v int32) *ListRepositoryMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryMemberResponse) SetBody(v *ListRepositoryMemberResponseBody) *ListRepositoryMemberResponse {
	s.Body = v
	return s
}

type ListRepositoryMemberWithInheritedRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRepositoryMemberWithInheritedRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedRequest) SetAccessToken(v string) *ListRepositoryMemberWithInheritedRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedRequest) SetOrganizationId(v string) *ListRepositoryMemberWithInheritedRequest {
	s.OrganizationId = &v
	return s
}

type ListRepositoryMemberWithInheritedResponseBody struct {
	ErrorCode    *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryMemberWithInheritedResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetErrorCode(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetErrorMessage(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetRequestId(v string) *ListRepositoryMemberWithInheritedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetResult(v []*ListRepositoryMemberWithInheritedResponseBodyResult) *ListRepositoryMemberWithInheritedResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBody) SetSuccess(v bool) *ListRepositoryMemberWithInheritedResponseBody {
	s.Success = &v
	return s
}

type ListRepositoryMemberWithInheritedResponseBodyResult struct {
	AccessLevel  *int32                                                        `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string                                                       `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string                                                       `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string                                                       `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Inherited    *ListRepositoryMemberWithInheritedResponseBodyResultInherited `json:"Inherited,omitempty" xml:"Inherited,omitempty" type:"Struct"`
	Name         *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	State        *string                                                       `json:"State,omitempty" xml:"State,omitempty"`
	Username     *string                                                       `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetAccessLevel(v int32) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetAvatarUrl(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetEmail(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Email = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetExternUserId(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetId(v int64) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetInherited(v *ListRepositoryMemberWithInheritedResponseBodyResultInherited) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Inherited = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetName(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetState(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResult) SetUsername(v string) *ListRepositoryMemberWithInheritedResponseBodyResult {
	s.Username = &v
	return s
}

type ListRepositoryMemberWithInheritedResponseBodyResultInherited struct {
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace *string `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	Path              *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace *string `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VisibilityLevel   *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponseBodyResultInherited) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponseBodyResultInherited) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetId(v int64) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Id = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetName(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Name = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetNameWithNamespace(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.NameWithNamespace = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetPath(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Path = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetPathWithNamespace(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.PathWithNamespace = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetType(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.Type = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponseBodyResultInherited) SetVisibilityLevel(v string) *ListRepositoryMemberWithInheritedResponseBodyResultInherited {
	s.VisibilityLevel = &v
	return s
}

type ListRepositoryMemberWithInheritedResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryMemberWithInheritedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryMemberWithInheritedResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponse) SetHeaders(v map[string]*string) *ListRepositoryMemberWithInheritedResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponse) SetStatusCode(v int32) *ListRepositoryMemberWithInheritedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponse) SetBody(v *ListRepositoryMemberWithInheritedResponseBody) *ListRepositoryMemberWithInheritedResponse {
	s.Body = v
	return s
}

type ListRepositoryProtectedBranchRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRepositoryProtectedBranchRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchRequest) SetAccessToken(v string) *ListRepositoryProtectedBranchRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryProtectedBranchRequest) SetOrganizationId(v string) *ListRepositoryProtectedBranchRequest {
	s.OrganizationId = &v
	return s
}

type ListRepositoryProtectedBranchResponseBody struct {
	ErrorCode    *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryProtectedBranchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRepositoryProtectedBranchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBody) SetErrorCode(v string) *ListRepositoryProtectedBranchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBody) SetErrorMessage(v string) *ListRepositoryProtectedBranchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBody) SetRequestId(v string) *ListRepositoryProtectedBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBody) SetResult(v []*ListRepositoryProtectedBranchResponseBodyResult) *ListRepositoryProtectedBranchResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBody) SetSuccess(v bool) *ListRepositoryProtectedBranchResponseBody {
	s.Success = &v
	return s
}

type ListRepositoryProtectedBranchResponseBodyResult struct {
	AllowMergeRoles     []*int32                                                            `json:"AllowMergeRoles,omitempty" xml:"AllowMergeRoles,omitempty" type:"Repeated"`
	AllowPushRoles      []*int32                                                            `json:"AllowPushRoles,omitempty" xml:"AllowPushRoles,omitempty" type:"Repeated"`
	Branch              *string                                                             `json:"Branch,omitempty" xml:"Branch,omitempty"`
	Id                  *int64                                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	MergeRequestSetting *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting `json:"MergeRequestSetting,omitempty" xml:"MergeRequestSetting,omitempty" type:"Struct"`
	TestSetting         *ListRepositoryProtectedBranchResponseBodyResultTestSetting         `json:"TestSetting,omitempty" xml:"TestSetting,omitempty" type:"Struct"`
}

func (s ListRepositoryProtectedBranchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBodyResult) SetAllowMergeRoles(v []*int32) *ListRepositoryProtectedBranchResponseBodyResult {
	s.AllowMergeRoles = v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResult) SetAllowPushRoles(v []*int32) *ListRepositoryProtectedBranchResponseBodyResult {
	s.AllowPushRoles = v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResult) SetBranch(v string) *ListRepositoryProtectedBranchResponseBodyResult {
	s.Branch = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResult) SetId(v int64) *ListRepositoryProtectedBranchResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResult) SetMergeRequestSetting(v *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) *ListRepositoryProtectedBranchResponseBodyResult {
	s.MergeRequestSetting = v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResult) SetTestSetting(v *ListRepositoryProtectedBranchResponseBodyResultTestSetting) *ListRepositoryProtectedBranchResponseBodyResult {
	s.TestSetting = v
	return s
}

type ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting struct {
	AllowMergeRequestRoles       []*int32                                                                              `json:"AllowMergeRequestRoles,omitempty" xml:"AllowMergeRequestRoles,omitempty" type:"Repeated"`
	AllowSelfApproval            *bool                                                                                 `json:"AllowSelfApproval,omitempty" xml:"AllowSelfApproval,omitempty"`
	DefaultAssignees             []*ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees `json:"DefaultAssignees,omitempty" xml:"DefaultAssignees,omitempty" type:"Repeated"`
	IsRequireDiscussionProcessed *bool                                                                                 `json:"IsRequireDiscussionProcessed,omitempty" xml:"IsRequireDiscussionProcessed,omitempty"`
	MergeRequestMode             *string                                                                               `json:"MergeRequestMode,omitempty" xml:"MergeRequestMode,omitempty"`
	MinimumApproval              *int32                                                                                `json:"MinimumApproval,omitempty" xml:"MinimumApproval,omitempty"`
	Required                     *bool                                                                                 `json:"Required,omitempty" xml:"Required,omitempty"`
	WhiteList                    *string                                                                               `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetAllowMergeRequestRoles(v []*int32) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.AllowMergeRequestRoles = v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetAllowSelfApproval(v bool) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.AllowSelfApproval = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetDefaultAssignees(v []*ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.DefaultAssignees = v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetIsRequireDiscussionProcessed(v bool) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.IsRequireDiscussionProcessed = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetMergeRequestMode(v string) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.MergeRequestMode = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetMinimumApproval(v int32) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.MinimumApproval = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetRequired(v bool) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.Required = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting) SetWhiteList(v string) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSetting {
	s.WhiteList = &v
	return s
}

type ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUid *string `json:"ExternUid,omitempty" xml:"ExternUid,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees) SetAvatarUrl(v string) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.AvatarUrl = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees) SetEmail(v string) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Email = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees) SetExternUid(v string) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.ExternUid = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees) SetId(v int64) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Id = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees) SetName(v string) *ListRepositoryProtectedBranchResponseBodyResultMergeRequestSettingDefaultAssignees {
	s.Name = &v
	return s
}

type ListRepositoryProtectedBranchResponseBodyResultTestSetting struct {
	CheckConfig               *ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig               `json:"CheckConfig,omitempty" xml:"CheckConfig,omitempty" type:"Struct"`
	CodingGuidelinesDetection *ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection `json:"CodingGuidelinesDetection,omitempty" xml:"CodingGuidelinesDetection,omitempty" type:"Struct"`
	Required                  *bool                                                                                `json:"Required,omitempty" xml:"Required,omitempty"`
	SensitiveInfoDetection    *ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection    `json:"SensitiveInfoDetection,omitempty" xml:"SensitiveInfoDetection,omitempty" type:"Struct"`
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSetting) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSetting) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSetting) SetCheckConfig(v *ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig) *ListRepositoryProtectedBranchResponseBodyResultTestSetting {
	s.CheckConfig = v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSetting) SetCodingGuidelinesDetection(v *ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) *ListRepositoryProtectedBranchResponseBodyResultTestSetting {
	s.CodingGuidelinesDetection = v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSetting) SetRequired(v bool) *ListRepositoryProtectedBranchResponseBodyResultTestSetting {
	s.Required = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSetting) SetSensitiveInfoDetection(v *ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) *ListRepositoryProtectedBranchResponseBodyResultTestSetting {
	s.SensitiveInfoDetection = v
	return s
}

type ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig struct {
	CheckItems []*ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems `json:"CheckItems,omitempty" xml:"CheckItems,omitempty" type:"Repeated"`
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig) SetCheckItems(v []*ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) *ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfig {
	s.CheckItems = v
	return s
}

type ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Required *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) SetName(v string) *ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems {
	s.Name = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems) SetRequired(v bool) *ListRepositoryProtectedBranchResponseBodyResultTestSettingCheckConfigCheckItems {
	s.Required = &v
	return s
}

type ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection struct {
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) SetEnabled(v bool) *ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection {
	s.Enabled = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection) SetMessage(v string) *ListRepositoryProtectedBranchResponseBodyResultTestSettingCodingGuidelinesDetection {
	s.Message = &v
	return s
}

type ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection struct {
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) SetEnabled(v bool) *ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection {
	s.Enabled = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection) SetMessage(v string) *ListRepositoryProtectedBranchResponseBodyResultTestSettingSensitiveInfoDetection {
	s.Message = &v
	return s
}

type ListRepositoryProtectedBranchResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryProtectedBranchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRepositoryProtectedBranchResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryProtectedBranchResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryProtectedBranchResponse) SetHeaders(v map[string]*string) *ListRepositoryProtectedBranchResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryProtectedBranchResponse) SetStatusCode(v int32) *ListRepositoryProtectedBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryProtectedBranchResponse) SetBody(v *ListRepositoryProtectedBranchResponseBody) *ListRepositoryProtectedBranchResponse {
	s.Body = v
	return s
}

type ListRepositoryTagsRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Page           *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search         *string `json:"Search,omitempty" xml:"Search,omitempty"`
	ShowSignature  *bool   `json:"ShowSignature,omitempty" xml:"ShowSignature,omitempty"`
	Sort           *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
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

func (s *ListRepositoryTagsRequest) SetSearch(v string) *ListRepositoryTagsRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetShowSignature(v bool) *ListRepositoryTagsRequest {
	s.ShowSignature = &v
	return s
}

func (s *ListRepositoryTagsRequest) SetSort(v string) *ListRepositoryTagsRequest {
	s.Sort = &v
	return s
}

type ListRepositoryTagsResponseBody struct {
	ErrorCode    *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryTagsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRepositoryTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBody) SetErrorCode(v string) *ListRepositoryTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetErrorMessage(v string) *ListRepositoryTagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetRequestId(v string) *ListRepositoryTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetResult(v []*ListRepositoryTagsResponseBodyResult) *ListRepositoryTagsResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetSuccess(v bool) *ListRepositoryTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryTagsResponseBody) SetTotal(v int64) *ListRepositoryTagsResponseBody {
	s.Total = &v
	return s
}

type ListRepositoryTagsResponseBodyResult struct {
	Commit    *ListRepositoryTagsResponseBodyResultCommit    `json:"Commit,omitempty" xml:"Commit,omitempty" type:"Struct"`
	Id        *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Name      *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Signature *ListRepositoryTagsResponseBodyResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
}

func (s ListRepositoryTagsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResult) SetCommit(v *ListRepositoryTagsResponseBodyResultCommit) *ListRepositoryTagsResponseBodyResult {
	s.Commit = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResult) SetId(v string) *ListRepositoryTagsResponseBodyResult {
	s.Id = &v
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

func (s *ListRepositoryTagsResponseBodyResult) SetSignature(v *ListRepositoryTagsResponseBodyResultSignature) *ListRepositoryTagsResponseBodyResult {
	s.Signature = v
	return s
}

type ListRepositoryTagsResponseBodyResultCommit struct {
	AuthorEmail    *string                                              `json:"AuthorEmail,omitempty" xml:"AuthorEmail,omitempty"`
	AuthorName     *string                                              `json:"AuthorName,omitempty" xml:"AuthorName,omitempty"`
	AuthoredDate   *string                                              `json:"AuthoredDate,omitempty" xml:"AuthoredDate,omitempty"`
	CommittedDate  *string                                              `json:"CommittedDate,omitempty" xml:"CommittedDate,omitempty"`
	CommitterEmail *string                                              `json:"CommitterEmail,omitempty" xml:"CommitterEmail,omitempty"`
	CommitterName  *string                                              `json:"CommitterName,omitempty" xml:"CommitterName,omitempty"`
	CreatedAt      *string                                              `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id             *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentIds      []*string                                            `json:"ParentIds,omitempty" xml:"ParentIds,omitempty" type:"Repeated"`
	ShortId        *string                                              `json:"ShortId,omitempty" xml:"ShortId,omitempty"`
	Signature      *ListRepositoryTagsResponseBodyResultCommitSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	Title          *string                                              `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListRepositoryTagsResponseBodyResultCommit) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTagsResponseBodyResultCommit) GoString() string {
	return s.String()
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthorEmail(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthorEmail = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthorName(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthorName = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetAuthoredDate(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.AuthoredDate = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommittedDate(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommittedDate = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommitterEmail(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommitterEmail = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCommitterName(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CommitterName = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetCreatedAt(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetId(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Id = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetMessage(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Message = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetParentIds(v []*string) *ListRepositoryTagsResponseBodyResultCommit {
	s.ParentIds = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetShortId(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.ShortId = &v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetSignature(v *ListRepositoryTagsResponseBodyResultCommitSignature) *ListRepositoryTagsResponseBodyResultCommit {
	s.Signature = v
	return s
}

func (s *ListRepositoryTagsResponseBodyResultCommit) SetTitle(v string) *ListRepositoryTagsResponseBodyResultCommit {
	s.Title = &v
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

type ListRepositoryTagsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRepositoryTagsResponse) SetStatusCode(v int32) *ListRepositoryTagsResponse {
	s.StatusCode = &v
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
	RefName        *string `json:"RefName,omitempty" xml:"RefName,omitempty"`
	SubUserId      *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListRepositoryTreeRequest) SetRefName(v string) *ListRepositoryTreeRequest {
	s.RefName = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetSubUserId(v string) *ListRepositoryTreeRequest {
	s.SubUserId = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetType(v string) *ListRepositoryTreeRequest {
	s.Type = &v
	return s
}

type ListRepositoryTreeResponseBody struct {
	ErrorCode    *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryTreeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRepositoryTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponseBody) SetErrorCode(v string) *ListRepositoryTreeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetErrorMessage(v string) *ListRepositoryTreeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetRequestId(v string) *ListRepositoryTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetResult(v []*ListRepositoryTreeResponseBodyResult) *ListRepositoryTreeResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryTreeResponseBody) SetSuccess(v bool) *ListRepositoryTreeResponseBody {
	s.Success = &v
	return s
}

type ListRepositoryTreeResponseBodyResult struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRepositoryTreeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryTreeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeResponseBodyResult) SetId(v string) *ListRepositoryTreeResponseBodyResult {
	s.Id = &v
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

func (s *ListRepositoryTreeResponseBodyResult) SetPath(v string) *ListRepositoryTreeResponseBodyResult {
	s.Path = &v
	return s
}

func (s *ListRepositoryTreeResponseBodyResult) SetType(v string) *ListRepositoryTreeResponseBodyResult {
	s.Type = &v
	return s
}

type ListRepositoryTreeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRepositoryTreeResponse) SetStatusCode(v int32) *ListRepositoryTreeResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*ListRepositoryWebhookResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Total        *int64                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRepositoryWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponseBody) SetErrorCode(v string) *ListRepositoryWebhookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetErrorMessage(v string) *ListRepositoryWebhookResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetRequestId(v string) *ListRepositoryWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetResult(v []*ListRepositoryWebhookResponseBodyResult) *ListRepositoryWebhookResponseBody {
	s.Result = v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetSuccess(v bool) *ListRepositoryWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *ListRepositoryWebhookResponseBody) SetTotal(v int64) *ListRepositoryWebhookResponseBody {
	s.Total = &v
	return s
}

type ListRepositoryWebhookResponseBodyResult struct {
	CreatedAt             *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableSslVerification *bool   `json:"EnableSslVerification,omitempty" xml:"EnableSslVerification,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LastTestResult        *string `json:"LastTestResult,omitempty" xml:"LastTestResult,omitempty"`
	MergeRequestsEvents   *bool   `json:"MergeRequestsEvents,omitempty" xml:"MergeRequestsEvents,omitempty"`
	NoteEvents            *bool   `json:"NoteEvents,omitempty" xml:"NoteEvents,omitempty"`
	ProjectId             *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PushEvents            *bool   `json:"PushEvents,omitempty" xml:"PushEvents,omitempty"`
	SecretToken           *string `json:"SecretToken,omitempty" xml:"SecretToken,omitempty"`
	TagPushEvents         *bool   `json:"TagPushEvents,omitempty" xml:"TagPushEvents,omitempty"`
	Url                   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListRepositoryWebhookResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryWebhookResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookResponseBodyResult) SetCreatedAt(v string) *ListRepositoryWebhookResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetDescription(v string) *ListRepositoryWebhookResponseBodyResult {
	s.Description = &v
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

func (s *ListRepositoryWebhookResponseBodyResult) SetLastTestResult(v string) *ListRepositoryWebhookResponseBodyResult {
	s.LastTestResult = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetMergeRequestsEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.MergeRequestsEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetNoteEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.NoteEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetProjectId(v int64) *ListRepositoryWebhookResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetPushEvents(v bool) *ListRepositoryWebhookResponseBodyResult {
	s.PushEvents = &v
	return s
}

func (s *ListRepositoryWebhookResponseBodyResult) SetSecretToken(v string) *ListRepositoryWebhookResponseBodyResult {
	s.SecretToken = &v
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

type ListRepositoryWebhookResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRepositoryWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRepositoryWebhookResponse) SetStatusCode(v int32) *ListRepositoryWebhookResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *MergeMergeRequestResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MergeMergeRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBody) SetErrorCode(v string) *MergeMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetErrorMessage(v string) *MergeMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetRequestId(v string) *MergeMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeMergeRequestResponseBody) SetResult(v *MergeMergeRequestResponseBodyResult) *MergeMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *MergeMergeRequestResponseBody) SetSuccess(v bool) *MergeMergeRequestResponseBody {
	s.Success = &v
	return s
}

type MergeMergeRequestResponseBodyResult struct {
	AcceptedRevision   *string                                                `json:"AcceptedRevision,omitempty" xml:"AcceptedRevision,omitempty"`
	AheadCommitCount   *int32                                                 `json:"AheadCommitCount,omitempty" xml:"AheadCommitCount,omitempty"`
	ApproveCheckResult *MergeMergeRequestResponseBodyResultApproveCheckResult `json:"ApproveCheckResult,omitempty" xml:"ApproveCheckResult,omitempty" type:"Struct"`
	AssigneeList       []*MergeMergeRequestResponseBodyResultAssigneeList     `json:"AssigneeList,omitempty" xml:"AssigneeList,omitempty" type:"Repeated"`
	Author             *MergeMergeRequestResponseBodyResultAuthor             `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	BehindCommitCount  *int32                                                 `json:"BehindCommitCount,omitempty" xml:"BehindCommitCount,omitempty"`
	CreatedAt          *string                                                `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description        *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                 *int64                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	MergeError         *string                                                `json:"MergeError,omitempty" xml:"MergeError,omitempty"`
	MergeStatus        *string                                                `json:"MergeStatus,omitempty" xml:"MergeStatus,omitempty"`
	MergeType          *string                                                `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	MergedRevision     *string                                                `json:"MergedRevision,omitempty" xml:"MergedRevision,omitempty"`
	NameWithNamespace  *string                                                `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	ProjectId          *int64                                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SourceBranch       *string                                                `json:"SourceBranch,omitempty" xml:"SourceBranch,omitempty"`
	State              *string                                                `json:"State,omitempty" xml:"State,omitempty"`
	TargetBranch       *string                                                `json:"TargetBranch,omitempty" xml:"TargetBranch,omitempty"`
	Title              *string                                                `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdatedAt          *string                                                `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	WebUrl             *string                                                `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s MergeMergeRequestResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResult) SetAcceptedRevision(v string) *MergeMergeRequestResponseBodyResult {
	s.AcceptedRevision = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetAheadCommitCount(v int32) *MergeMergeRequestResponseBodyResult {
	s.AheadCommitCount = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetApproveCheckResult(v *MergeMergeRequestResponseBodyResultApproveCheckResult) *MergeMergeRequestResponseBodyResult {
	s.ApproveCheckResult = v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetAssigneeList(v []*MergeMergeRequestResponseBodyResultAssigneeList) *MergeMergeRequestResponseBodyResult {
	s.AssigneeList = v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetAuthor(v *MergeMergeRequestResponseBodyResultAuthor) *MergeMergeRequestResponseBodyResult {
	s.Author = v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetBehindCommitCount(v int32) *MergeMergeRequestResponseBodyResult {
	s.BehindCommitCount = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetCreatedAt(v string) *MergeMergeRequestResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetDescription(v string) *MergeMergeRequestResponseBodyResult {
	s.Description = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetId(v int64) *MergeMergeRequestResponseBodyResult {
	s.Id = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergeError(v string) *MergeMergeRequestResponseBodyResult {
	s.MergeError = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergeStatus(v string) *MergeMergeRequestResponseBodyResult {
	s.MergeStatus = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergeType(v string) *MergeMergeRequestResponseBodyResult {
	s.MergeType = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetMergedRevision(v string) *MergeMergeRequestResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetNameWithNamespace(v string) *MergeMergeRequestResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetProjectId(v int64) *MergeMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetSourceBranch(v string) *MergeMergeRequestResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetState(v string) *MergeMergeRequestResponseBodyResult {
	s.State = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetTargetBranch(v string) *MergeMergeRequestResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetTitle(v string) *MergeMergeRequestResponseBodyResult {
	s.Title = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetUpdatedAt(v string) *MergeMergeRequestResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResult) SetWebUrl(v string) *MergeMergeRequestResponseBodyResult {
	s.WebUrl = &v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResult struct {
	SatisfiedCheckResults   []*MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults   `json:"SatisfiedCheckResults,omitempty" xml:"SatisfiedCheckResults,omitempty" type:"Repeated"`
	TotalCheckResult        *string                                                                         `json:"TotalCheckResult,omitempty" xml:"TotalCheckResult,omitempty"`
	UnsatisfiedCheckResults []*MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults `json:"UnsatisfiedCheckResults,omitempty" xml:"UnsatisfiedCheckResults,omitempty" type:"Repeated"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResult) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResult) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResult) SetSatisfiedCheckResults(v []*MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) *MergeMergeRequestResponseBodyResultApproveCheckResult {
	s.SatisfiedCheckResults = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResult) SetTotalCheckResult(v string) *MergeMergeRequestResponseBodyResultApproveCheckResult {
	s.TotalCheckResult = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResult) SetUnsatisfiedCheckResults(v []*MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) *MergeMergeRequestResponseBodyResultApproveCheckResult {
	s.UnsatisfiedCheckResults = v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults struct {
	CheckName        *string                                                                                 `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                               `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                               `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckName(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckName = &v
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

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetExtraUsers(v []*MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetSatisfiedItems(v []*string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetId(v int64) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetName(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults struct {
	CheckName        *string                                                                                   `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                   `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                   `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                 `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                 `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckName(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckName = &v
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

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetExtraUsers(v []*MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetSatisfiedItems(v []*string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetId(v int64) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetName(v string) *MergeMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type MergeMergeRequestResponseBodyResultAssigneeList struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultAssigneeList) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultAssigneeList) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultAssigneeList) SetAvatarUrl(v string) *MergeMergeRequestResponseBodyResultAssigneeList {
	s.AvatarUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAssigneeList) SetExternUserId(v string) *MergeMergeRequestResponseBodyResultAssigneeList {
	s.ExternUserId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAssigneeList) SetId(v string) *MergeMergeRequestResponseBodyResultAssigneeList {
	s.Id = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAssigneeList) SetName(v string) *MergeMergeRequestResponseBodyResultAssigneeList {
	s.Name = &v
	return s
}

type MergeMergeRequestResponseBodyResultAuthor struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MergeMergeRequestResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s MergeMergeRequestResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *MergeMergeRequestResponseBodyResultAuthor) SetAvatarUrl(v string) *MergeMergeRequestResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAuthor) SetExternUserId(v string) *MergeMergeRequestResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAuthor) SetId(v int64) *MergeMergeRequestResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *MergeMergeRequestResponseBodyResultAuthor) SetName(v string) *MergeMergeRequestResponseBodyResultAuthor {
	s.Name = &v
	return s
}

type MergeMergeRequestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MergeMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *MergeMergeRequestResponse) SetStatusCode(v int32) *MergeMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeMergeRequestResponse) SetBody(v *MergeMergeRequestResponseBody) *MergeMergeRequestResponse {
	s.Body = v
	return s
}

type QuerySlsRelationRequest struct {
	OrganizationId  *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	AliyunUserId    *string `json:"aliyunUserId,omitempty" xml:"aliyunUserId,omitempty"`
	CodeupProjectId *int64  `json:"codeupProjectId,omitempty" xml:"codeupProjectId,omitempty"`
	SlsLogStore     *string `json:"slsLogStore,omitempty" xml:"slsLogStore,omitempty"`
	SlsProject      *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
}

func (s QuerySlsRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsRelationRequest) GoString() string {
	return s.String()
}

func (s *QuerySlsRelationRequest) SetOrganizationId(v string) *QuerySlsRelationRequest {
	s.OrganizationId = &v
	return s
}

func (s *QuerySlsRelationRequest) SetAliyunUserId(v string) *QuerySlsRelationRequest {
	s.AliyunUserId = &v
	return s
}

func (s *QuerySlsRelationRequest) SetCodeupProjectId(v int64) *QuerySlsRelationRequest {
	s.CodeupProjectId = &v
	return s
}

func (s *QuerySlsRelationRequest) SetSlsLogStore(v string) *QuerySlsRelationRequest {
	s.SlsLogStore = &v
	return s
}

func (s *QuerySlsRelationRequest) SetSlsProject(v string) *QuerySlsRelationRequest {
	s.SlsProject = &v
	return s
}

type QuerySlsRelationResponseBody struct {
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       []*QuerySlsRelationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success      *string                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySlsRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySlsRelationResponseBody) SetErrorCode(v string) *QuerySlsRelationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QuerySlsRelationResponseBody) SetErrorMessage(v string) *QuerySlsRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySlsRelationResponseBody) SetRequestId(v string) *QuerySlsRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySlsRelationResponseBody) SetResult(v []*QuerySlsRelationResponseBodyResult) *QuerySlsRelationResponseBody {
	s.Result = v
	return s
}

func (s *QuerySlsRelationResponseBody) SetSuccess(v string) *QuerySlsRelationResponseBody {
	s.Success = &v
	return s
}

type QuerySlsRelationResponseBodyResult struct {
	AliyunUserId    *string `json:"aliyunUserId,omitempty" xml:"aliyunUserId,omitempty"`
	CodeupProjectId *int64  `json:"codeupProjectId,omitempty" xml:"codeupProjectId,omitempty"`
	DefaultViewer   *bool   `json:"defaultViewer,omitempty" xml:"defaultViewer,omitempty"`
	OrganizationId  *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	SlsLogStore     *string `json:"slsLogStore,omitempty" xml:"slsLogStore,omitempty"`
	SlsProject      *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
}

func (s QuerySlsRelationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsRelationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySlsRelationResponseBodyResult) SetAliyunUserId(v string) *QuerySlsRelationResponseBodyResult {
	s.AliyunUserId = &v
	return s
}

func (s *QuerySlsRelationResponseBodyResult) SetCodeupProjectId(v int64) *QuerySlsRelationResponseBodyResult {
	s.CodeupProjectId = &v
	return s
}

func (s *QuerySlsRelationResponseBodyResult) SetDefaultViewer(v bool) *QuerySlsRelationResponseBodyResult {
	s.DefaultViewer = &v
	return s
}

func (s *QuerySlsRelationResponseBodyResult) SetOrganizationId(v string) *QuerySlsRelationResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *QuerySlsRelationResponseBodyResult) SetSlsLogStore(v string) *QuerySlsRelationResponseBodyResult {
	s.SlsLogStore = &v
	return s
}

func (s *QuerySlsRelationResponseBodyResult) SetSlsProject(v string) *QuerySlsRelationResponseBodyResult {
	s.SlsProject = &v
	return s
}

type QuerySlsRelationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySlsRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySlsRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsRelationResponse) GoString() string {
	return s.String()
}

func (s *QuerySlsRelationResponse) SetHeaders(v map[string]*string) *QuerySlsRelationResponse {
	s.Headers = v
	return s
}

func (s *QuerySlsRelationResponse) SetStatusCode(v int32) *QuerySlsRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySlsRelationResponse) SetBody(v *QuerySlsRelationResponseBody) *QuerySlsRelationResponse {
	s.Body = v
	return s
}

type RelatedSlsLogStoreRequest struct {
	OrganizationId  *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	AliyunUserId    *string `json:"aliyunUserId,omitempty" xml:"aliyunUserId,omitempty"`
	CodeupProjectId *int64  `json:"codeupProjectId,omitempty" xml:"codeupProjectId,omitempty"`
	DefaultViewer   *bool   `json:"defaultViewer,omitempty" xml:"defaultViewer,omitempty"`
	SlsLogStore     *string `json:"slsLogStore,omitempty" xml:"slsLogStore,omitempty"`
	SlsProject      *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
}

func (s RelatedSlsLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s RelatedSlsLogStoreRequest) GoString() string {
	return s.String()
}

func (s *RelatedSlsLogStoreRequest) SetOrganizationId(v string) *RelatedSlsLogStoreRequest {
	s.OrganizationId = &v
	return s
}

func (s *RelatedSlsLogStoreRequest) SetAliyunUserId(v string) *RelatedSlsLogStoreRequest {
	s.AliyunUserId = &v
	return s
}

func (s *RelatedSlsLogStoreRequest) SetCodeupProjectId(v int64) *RelatedSlsLogStoreRequest {
	s.CodeupProjectId = &v
	return s
}

func (s *RelatedSlsLogStoreRequest) SetDefaultViewer(v bool) *RelatedSlsLogStoreRequest {
	s.DefaultViewer = &v
	return s
}

func (s *RelatedSlsLogStoreRequest) SetSlsLogStore(v string) *RelatedSlsLogStoreRequest {
	s.SlsLogStore = &v
	return s
}

func (s *RelatedSlsLogStoreRequest) SetSlsProject(v string) *RelatedSlsLogStoreRequest {
	s.SlsProject = &v
	return s
}

type RelatedSlsLogStoreResponseBody struct {
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *RelatedSlsLogStoreResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *string                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RelatedSlsLogStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RelatedSlsLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *RelatedSlsLogStoreResponseBody) SetErrorCode(v string) *RelatedSlsLogStoreResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RelatedSlsLogStoreResponseBody) SetErrorMessage(v string) *RelatedSlsLogStoreResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RelatedSlsLogStoreResponseBody) SetRequestId(v string) *RelatedSlsLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *RelatedSlsLogStoreResponseBody) SetResult(v *RelatedSlsLogStoreResponseBodyResult) *RelatedSlsLogStoreResponseBody {
	s.Result = v
	return s
}

func (s *RelatedSlsLogStoreResponseBody) SetSuccess(v string) *RelatedSlsLogStoreResponseBody {
	s.Success = &v
	return s
}

type RelatedSlsLogStoreResponseBodyResult struct {
	RelatedResult *bool `json:"RelatedResult,omitempty" xml:"RelatedResult,omitempty"`
}

func (s RelatedSlsLogStoreResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RelatedSlsLogStoreResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RelatedSlsLogStoreResponseBodyResult) SetRelatedResult(v bool) *RelatedSlsLogStoreResponseBodyResult {
	s.RelatedResult = &v
	return s
}

type RelatedSlsLogStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RelatedSlsLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RelatedSlsLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s RelatedSlsLogStoreResponse) GoString() string {
	return s.String()
}

func (s *RelatedSlsLogStoreResponse) SetHeaders(v map[string]*string) *RelatedSlsLogStoreResponse {
	s.Headers = v
	return s
}

func (s *RelatedSlsLogStoreResponse) SetStatusCode(v int32) *RelatedSlsLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *RelatedSlsLogStoreResponse) SetBody(v *RelatedSlsLogStoreResponseBody) *RelatedSlsLogStoreResponse {
	s.Body = v
	return s
}

type TriggerRepositoryMirrorSyncRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Account        *string `json:"Account,omitempty" xml:"Account,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s TriggerRepositoryMirrorSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncRequest) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncRequest) SetAccessToken(v string) *TriggerRepositoryMirrorSyncRequest {
	s.AccessToken = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetAccount(v string) *TriggerRepositoryMirrorSyncRequest {
	s.Account = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetOrganizationId(v string) *TriggerRepositoryMirrorSyncRequest {
	s.OrganizationId = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncRequest) SetToken(v string) *TriggerRepositoryMirrorSyncRequest {
	s.Token = &v
	return s
}

type TriggerRepositoryMirrorSyncResponseBody struct {
	ErrorCode    *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *TriggerRepositoryMirrorSyncResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TriggerRepositoryMirrorSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetErrorCode(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetErrorMessage(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetRequestId(v string) *TriggerRepositoryMirrorSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetResult(v *TriggerRepositoryMirrorSyncResponseBodyResult) *TriggerRepositoryMirrorSyncResponseBody {
	s.Result = v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponseBody) SetSuccess(v bool) *TriggerRepositoryMirrorSyncResponseBody {
	s.Success = &v
	return s
}

type TriggerRepositoryMirrorSyncResponseBodyResult struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TriggerRepositoryMirrorSyncResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponseBodyResult) SetResult(v bool) *TriggerRepositoryMirrorSyncResponseBodyResult {
	s.Result = &v
	return s
}

type TriggerRepositoryMirrorSyncResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TriggerRepositoryMirrorSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerRepositoryMirrorSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerRepositoryMirrorSyncResponse) GoString() string {
	return s.String()
}

func (s *TriggerRepositoryMirrorSyncResponse) SetHeaders(v map[string]*string) *TriggerRepositoryMirrorSyncResponse {
	s.Headers = v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponse) SetStatusCode(v int32) *TriggerRepositoryMirrorSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerRepositoryMirrorSyncResponse) SetBody(v *TriggerRepositoryMirrorSyncResponseBody) *TriggerRepositoryMirrorSyncResponse {
	s.Body = v
	return s
}

type UnRelatedSlsLogStoreRequest struct {
	OrganizationId  *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	AliyunUserId    *string `json:"aliyunUserId,omitempty" xml:"aliyunUserId,omitempty"`
	CodeupProjectId *int64  `json:"codeupProjectId,omitempty" xml:"codeupProjectId,omitempty"`
	SlsLogStore     *string `json:"slsLogStore,omitempty" xml:"slsLogStore,omitempty"`
	SlsProject      *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
}

func (s UnRelatedSlsLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UnRelatedSlsLogStoreRequest) GoString() string {
	return s.String()
}

func (s *UnRelatedSlsLogStoreRequest) SetOrganizationId(v string) *UnRelatedSlsLogStoreRequest {
	s.OrganizationId = &v
	return s
}

func (s *UnRelatedSlsLogStoreRequest) SetAliyunUserId(v string) *UnRelatedSlsLogStoreRequest {
	s.AliyunUserId = &v
	return s
}

func (s *UnRelatedSlsLogStoreRequest) SetCodeupProjectId(v int64) *UnRelatedSlsLogStoreRequest {
	s.CodeupProjectId = &v
	return s
}

func (s *UnRelatedSlsLogStoreRequest) SetSlsLogStore(v string) *UnRelatedSlsLogStoreRequest {
	s.SlsLogStore = &v
	return s
}

func (s *UnRelatedSlsLogStoreRequest) SetSlsProject(v string) *UnRelatedSlsLogStoreRequest {
	s.SlsProject = &v
	return s
}

type UnRelatedSlsLogStoreResponseBody struct {
	ErrorCode    *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *UnRelatedSlsLogStoreResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *string                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnRelatedSlsLogStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnRelatedSlsLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UnRelatedSlsLogStoreResponseBody) SetErrorCode(v string) *UnRelatedSlsLogStoreResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnRelatedSlsLogStoreResponseBody) SetErrorMessage(v string) *UnRelatedSlsLogStoreResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnRelatedSlsLogStoreResponseBody) SetRequestId(v string) *UnRelatedSlsLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnRelatedSlsLogStoreResponseBody) SetResult(v *UnRelatedSlsLogStoreResponseBodyResult) *UnRelatedSlsLogStoreResponseBody {
	s.Result = v
	return s
}

func (s *UnRelatedSlsLogStoreResponseBody) SetSuccess(v string) *UnRelatedSlsLogStoreResponseBody {
	s.Success = &v
	return s
}

type UnRelatedSlsLogStoreResponseBodyResult struct {
	UnRelatedResult *bool `json:"UnRelatedResult,omitempty" xml:"UnRelatedResult,omitempty"`
}

func (s UnRelatedSlsLogStoreResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UnRelatedSlsLogStoreResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UnRelatedSlsLogStoreResponseBodyResult) SetUnRelatedResult(v bool) *UnRelatedSlsLogStoreResponseBodyResult {
	s.UnRelatedResult = &v
	return s
}

type UnRelatedSlsLogStoreResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnRelatedSlsLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnRelatedSlsLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UnRelatedSlsLogStoreResponse) GoString() string {
	return s.String()
}

func (s *UnRelatedSlsLogStoreResponse) SetHeaders(v map[string]*string) *UnRelatedSlsLogStoreResponse {
	s.Headers = v
	return s
}

func (s *UnRelatedSlsLogStoreResponse) SetStatusCode(v int32) *UnRelatedSlsLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *UnRelatedSlsLogStoreResponse) SetBody(v *UnRelatedSlsLogStoreResponseBody) *UnRelatedSlsLogStoreResponse {
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
	ErrorCode    *string                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *UpdateFileResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBody) SetErrorCode(v string) *UpdateFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFileResponseBody) SetErrorMessage(v string) *UpdateFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFileResponseBody) SetRequestId(v string) *UpdateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileResponseBody) SetResult(v *UpdateFileResponseBodyResult) *UpdateFileResponseBody {
	s.Result = v
	return s
}

func (s *UpdateFileResponseBody) SetSuccess(v bool) *UpdateFileResponseBody {
	s.Success = &v
	return s
}

type UpdateFileResponseBodyResult struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	FilePath   *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s UpdateFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateFileResponseBodyResult) SetBranchName(v string) *UpdateFileResponseBodyResult {
	s.BranchName = &v
	return s
}

func (s *UpdateFileResponseBodyResult) SetFilePath(v string) *UpdateFileResponseBodyResult {
	s.FilePath = &v
	return s
}

type UpdateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateFileResponse) SetStatusCode(v int32) *UpdateFileResponse {
	s.StatusCode = &v
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
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *UpdateGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberResponseBody) SetErrorCode(v string) *UpdateGroupMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetErrorMessage(v string) *UpdateGroupMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetRequestId(v string) *UpdateGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetResult(v *UpdateGroupMemberResponseBodyResult) *UpdateGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *UpdateGroupMemberResponseBody) SetSuccess(v bool) *UpdateGroupMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateGroupMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdateGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberResponseBodyResult) SetAccessLevel(v int32) *UpdateGroupMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetAvatarUrl(v string) *UpdateGroupMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetEmail(v string) *UpdateGroupMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetExternUserId(v string) *UpdateGroupMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetId(v int64) *UpdateGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateGroupMemberResponseBodyResult) SetState(v string) *UpdateGroupMemberResponseBodyResult {
	s.State = &v
	return s
}

type UpdateGroupMemberResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateGroupMemberResponse) SetStatusCode(v int32) *UpdateGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupMemberResponse) SetBody(v *UpdateGroupMemberResponseBody) *UpdateGroupMemberResponse {
	s.Body = v
	return s
}

type UpdateMergeRequestRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateMergeRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestRequest) SetAccessToken(v string) *UpdateMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateMergeRequestRequest) SetOrganizationId(v string) *UpdateMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

type UpdateMergeRequestResponseBody struct {
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *UpdateMergeRequestResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMergeRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBody) SetErrorCode(v string) *UpdateMergeRequestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMergeRequestResponseBody) SetErrorMessage(v string) *UpdateMergeRequestResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMergeRequestResponseBody) SetRequestId(v string) *UpdateMergeRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMergeRequestResponseBody) SetResult(v *UpdateMergeRequestResponseBodyResult) *UpdateMergeRequestResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMergeRequestResponseBody) SetSuccess(v bool) *UpdateMergeRequestResponseBody {
	s.Success = &v
	return s
}

type UpdateMergeRequestResponseBodyResult struct {
	AcceptedRevision   *string                                                 `json:"AcceptedRevision,omitempty" xml:"AcceptedRevision,omitempty"`
	AheadCommitCount   *int32                                                  `json:"AheadCommitCount,omitempty" xml:"AheadCommitCount,omitempty"`
	ApproveCheckResult *UpdateMergeRequestResponseBodyResultApproveCheckResult `json:"ApproveCheckResult,omitempty" xml:"ApproveCheckResult,omitempty" type:"Struct"`
	AssigneeList       []*UpdateMergeRequestResponseBodyResultAssigneeList     `json:"AssigneeList,omitempty" xml:"AssigneeList,omitempty" type:"Repeated"`
	Author             *UpdateMergeRequestResponseBodyResultAuthor             `json:"Author,omitempty" xml:"Author,omitempty" type:"Struct"`
	BehindCommitCount  *int32                                                  `json:"BehindCommitCount,omitempty" xml:"BehindCommitCount,omitempty"`
	CreatedAt          *string                                                 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description        *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                 *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	MergeError         *string                                                 `json:"MergeError,omitempty" xml:"MergeError,omitempty"`
	MergeStatus        *string                                                 `json:"MergeStatus,omitempty" xml:"MergeStatus,omitempty"`
	MergeType          *string                                                 `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	MergedRevision     *string                                                 `json:"MergedRevision,omitempty" xml:"MergedRevision,omitempty"`
	NameWithNamespace  *string                                                 `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	ProjectId          *int64                                                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SourceBranch       *string                                                 `json:"SourceBranch,omitempty" xml:"SourceBranch,omitempty"`
	State              *string                                                 `json:"State,omitempty" xml:"State,omitempty"`
	TargetBranch       *string                                                 `json:"TargetBranch,omitempty" xml:"TargetBranch,omitempty"`
	Title              *string                                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdatedAt          *string                                                 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	WebUrl             *string                                                 `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s UpdateMergeRequestResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResult) SetAcceptedRevision(v string) *UpdateMergeRequestResponseBodyResult {
	s.AcceptedRevision = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetAheadCommitCount(v int32) *UpdateMergeRequestResponseBodyResult {
	s.AheadCommitCount = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetApproveCheckResult(v *UpdateMergeRequestResponseBodyResultApproveCheckResult) *UpdateMergeRequestResponseBodyResult {
	s.ApproveCheckResult = v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetAssigneeList(v []*UpdateMergeRequestResponseBodyResultAssigneeList) *UpdateMergeRequestResponseBodyResult {
	s.AssigneeList = v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetAuthor(v *UpdateMergeRequestResponseBodyResultAuthor) *UpdateMergeRequestResponseBodyResult {
	s.Author = v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetBehindCommitCount(v int32) *UpdateMergeRequestResponseBodyResult {
	s.BehindCommitCount = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetCreatedAt(v string) *UpdateMergeRequestResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetDescription(v string) *UpdateMergeRequestResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetId(v int64) *UpdateMergeRequestResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetMergeError(v string) *UpdateMergeRequestResponseBodyResult {
	s.MergeError = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetMergeStatus(v string) *UpdateMergeRequestResponseBodyResult {
	s.MergeStatus = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetMergeType(v string) *UpdateMergeRequestResponseBodyResult {
	s.MergeType = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetMergedRevision(v string) *UpdateMergeRequestResponseBodyResult {
	s.MergedRevision = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetNameWithNamespace(v string) *UpdateMergeRequestResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetProjectId(v int64) *UpdateMergeRequestResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetSourceBranch(v string) *UpdateMergeRequestResponseBodyResult {
	s.SourceBranch = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetState(v string) *UpdateMergeRequestResponseBodyResult {
	s.State = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetTargetBranch(v string) *UpdateMergeRequestResponseBodyResult {
	s.TargetBranch = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetTitle(v string) *UpdateMergeRequestResponseBodyResult {
	s.Title = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetUpdatedAt(v string) *UpdateMergeRequestResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResult) SetWebUrl(v string) *UpdateMergeRequestResponseBodyResult {
	s.WebUrl = &v
	return s
}

type UpdateMergeRequestResponseBodyResultApproveCheckResult struct {
	SatisfiedCheckResults   []*UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults   `json:"SatisfiedCheckResults,omitempty" xml:"SatisfiedCheckResults,omitempty" type:"Repeated"`
	TotalCheckResult        *string                                                                          `json:"TotalCheckResult,omitempty" xml:"TotalCheckResult,omitempty"`
	UnsatisfiedCheckResults []*UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults `json:"UnsatisfiedCheckResults,omitempty" xml:"UnsatisfiedCheckResults,omitempty" type:"Repeated"`
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResult) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResult) SetSatisfiedCheckResults(v []*UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) *UpdateMergeRequestResponseBodyResultApproveCheckResult {
	s.SatisfiedCheckResults = v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResult) SetTotalCheckResult(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResult {
	s.TotalCheckResult = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResult) SetUnsatisfiedCheckResults(v []*UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) *UpdateMergeRequestResponseBodyResultApproveCheckResult {
	s.UnsatisfiedCheckResults = v
	return s
}

type UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults struct {
	CheckName        *string                                                                                  `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                  `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                  `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckName(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckName = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckStatus(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetCheckType(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetExtraUsers(v []*UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetSatisfiedItems(v []*string) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetId(v int64) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers) SetName(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultSatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults struct {
	CheckName        *string                                                                                    `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	CheckStatus      *string                                                                                    `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	CheckType        *string                                                                                    `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExtraUsers       []*UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers `json:"ExtraUsers,omitempty" xml:"ExtraUsers,omitempty" type:"Repeated"`
	SatisfiedItems   []*string                                                                                  `json:"SatisfiedItems,omitempty" xml:"SatisfiedItems,omitempty" type:"Repeated"`
	UnsatisfiedItems []*string                                                                                  `json:"UnsatisfiedItems,omitempty" xml:"UnsatisfiedItems,omitempty" type:"Repeated"`
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckName(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckName = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckStatus(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckStatus = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetCheckType(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.CheckType = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetExtraUsers(v []*UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.ExtraUsers = v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetSatisfiedItems(v []*string) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.SatisfiedItems = v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults) SetUnsatisfiedItems(v []*string) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResults {
	s.UnsatisfiedItems = v
	return s
}

type UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetAvatarUrl(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetExternUserId(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.ExternUserId = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetId(v int64) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Id = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers) SetName(v string) *UpdateMergeRequestResponseBodyResultApproveCheckResultUnsatisfiedCheckResultsExtraUsers {
	s.Name = &v
	return s
}

type UpdateMergeRequestResponseBodyResultAssigneeList struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMergeRequestResponseBodyResultAssigneeList) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResultAssigneeList) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResultAssigneeList) SetAvatarUrl(v string) *UpdateMergeRequestResponseBodyResultAssigneeList {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultAssigneeList) SetExternUserId(v string) *UpdateMergeRequestResponseBodyResultAssigneeList {
	s.ExternUserId = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultAssigneeList) SetId(v string) *UpdateMergeRequestResponseBodyResultAssigneeList {
	s.Id = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultAssigneeList) SetName(v string) *UpdateMergeRequestResponseBodyResultAssigneeList {
	s.Name = &v
	return s
}

type UpdateMergeRequestResponseBodyResultAuthor struct {
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMergeRequestResponseBodyResultAuthor) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponseBodyResultAuthor) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponseBodyResultAuthor) SetAvatarUrl(v string) *UpdateMergeRequestResponseBodyResultAuthor {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultAuthor) SetExternUserId(v string) *UpdateMergeRequestResponseBodyResultAuthor {
	s.ExternUserId = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultAuthor) SetId(v int64) *UpdateMergeRequestResponseBodyResultAuthor {
	s.Id = &v
	return s
}

func (s *UpdateMergeRequestResponseBodyResultAuthor) SetName(v string) *UpdateMergeRequestResponseBodyResultAuthor {
	s.Name = &v
	return s
}

type UpdateMergeRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMergeRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestResponse) SetHeaders(v map[string]*string) *UpdateMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *UpdateMergeRequestResponse) SetStatusCode(v int32) *UpdateMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMergeRequestResponse) SetBody(v *UpdateMergeRequestResponseBody) *UpdateMergeRequestResponse {
	s.Body = v
	return s
}

type UpdateMergeRequestCommentRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateMergeRequestCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestCommentRequest) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestCommentRequest) SetAccessToken(v string) *UpdateMergeRequestCommentRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateMergeRequestCommentRequest) SetOrganizationId(v string) *UpdateMergeRequestCommentRequest {
	s.OrganizationId = &v
	return s
}

type UpdateMergeRequestCommentResponseBody struct {
	ErrorCode    *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *UpdateMergeRequestCommentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMergeRequestCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestCommentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestCommentResponseBody) SetErrorCode(v string) *UpdateMergeRequestCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMergeRequestCommentResponseBody) SetErrorMessage(v string) *UpdateMergeRequestCommentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMergeRequestCommentResponseBody) SetRequestId(v string) *UpdateMergeRequestCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMergeRequestCommentResponseBody) SetResult(v *UpdateMergeRequestCommentResponseBodyResult) *UpdateMergeRequestCommentResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMergeRequestCommentResponseBody) SetSuccess(v bool) *UpdateMergeRequestCommentResponseBody {
	s.Success = &v
	return s
}

type UpdateMergeRequestCommentResponseBodyResult struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateMergeRequestCommentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestCommentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestCommentResponseBodyResult) SetResult(v bool) *UpdateMergeRequestCommentResponseBodyResult {
	s.Result = &v
	return s
}

type UpdateMergeRequestCommentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMergeRequestCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMergeRequestCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestCommentResponse) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestCommentResponse) SetHeaders(v map[string]*string) *UpdateMergeRequestCommentResponse {
	s.Headers = v
	return s
}

func (s *UpdateMergeRequestCommentResponse) SetStatusCode(v int32) *UpdateMergeRequestCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMergeRequestCommentResponse) SetBody(v *UpdateMergeRequestCommentResponseBody) *UpdateMergeRequestCommentResponse {
	s.Body = v
	return s
}

type UpdateMergeRequestSettingRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateMergeRequestSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestSettingRequest) SetAccessToken(v string) *UpdateMergeRequestSettingRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateMergeRequestSettingRequest) SetOrganizationId(v string) *UpdateMergeRequestSettingRequest {
	s.OrganizationId = &v
	return s
}

type UpdateMergeRequestSettingResponseBody struct {
	ErrorCode    *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *UpdateMergeRequestSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMergeRequestSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestSettingResponseBody) SetErrorCode(v string) *UpdateMergeRequestSettingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMergeRequestSettingResponseBody) SetErrorMessage(v string) *UpdateMergeRequestSettingResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMergeRequestSettingResponseBody) SetRequestId(v string) *UpdateMergeRequestSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMergeRequestSettingResponseBody) SetResult(v *UpdateMergeRequestSettingResponseBodyResult) *UpdateMergeRequestSettingResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMergeRequestSettingResponseBody) SetSuccess(v bool) *UpdateMergeRequestSettingResponseBody {
	s.Success = &v
	return s
}

type UpdateMergeRequestSettingResponseBodyResult struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateMergeRequestSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestSettingResponseBodyResult) SetResult(v bool) *UpdateMergeRequestSettingResponseBodyResult {
	s.Result = &v
	return s
}

type UpdateMergeRequestSettingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMergeRequestSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMergeRequestSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMergeRequestSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestSettingResponse) SetHeaders(v map[string]*string) *UpdateMergeRequestSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateMergeRequestSettingResponse) SetStatusCode(v int32) *UpdateMergeRequestSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMergeRequestSettingResponse) SetBody(v *UpdateMergeRequestSettingResponseBody) *UpdateMergeRequestSettingResponse {
	s.Body = v
	return s
}

type UpdateRepositoryRequest struct {
	AccessToken    *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequest) SetAccessToken(v string) *UpdateRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateRepositoryRequest) SetOrganizationId(v string) *UpdateRepositoryRequest {
	s.OrganizationId = &v
	return s
}

type UpdateRepositoryResponseBody struct {
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *UpdateRepositoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBody) SetErrorCode(v string) *UpdateRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetErrorMessage(v string) *UpdateRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetRequestId(v string) *UpdateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetResult(v *UpdateRepositoryResponseBodyResult) *UpdateRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *UpdateRepositoryResponseBody) SetSuccess(v bool) *UpdateRepositoryResponseBody {
	s.Success = &v
	return s
}

type UpdateRepositoryResponseBodyResult struct {
	Archive           *bool                                        `json:"Archive,omitempty" xml:"Archive,omitempty"`
	AvatarUrl         *string                                      `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	CreatedAt         *string                                      `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CreatorId         *int64                                       `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DefaultBranch     *string                                      `json:"DefaultBranch,omitempty" xml:"DefaultBranch,omitempty"`
	Description       *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpUrlToRepo     *string                                      `json:"HttpUrlToRepo,omitempty" xml:"HttpUrlToRepo,omitempty"`
	Id                *int64                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	LastActivityAt    *string                                      `json:"LastActivityAt,omitempty" xml:"LastActivityAt,omitempty"`
	Name              *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	NameWithNamespace *string                                      `json:"NameWithNamespace,omitempty" xml:"NameWithNamespace,omitempty"`
	Namespace         *UpdateRepositoryResponseBodyResultNamespace `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Struct"`
	Path              *string                                      `json:"Path,omitempty" xml:"Path,omitempty"`
	PathWithNamespace *string                                      `json:"PathWithNamespace,omitempty" xml:"PathWithNamespace,omitempty"`
	SshUrlToRepo      *string                                      `json:"SshUrlToRepo,omitempty" xml:"SshUrlToRepo,omitempty"`
	VisibilityLevel   *string                                      `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
	WebUrl            *string                                      `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s UpdateRepositoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBodyResult) SetArchive(v bool) *UpdateRepositoryResponseBodyResult {
	s.Archive = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetAvatarUrl(v string) *UpdateRepositoryResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetCreatedAt(v string) *UpdateRepositoryResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetCreatorId(v int64) *UpdateRepositoryResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetDefaultBranch(v string) *UpdateRepositoryResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetDescription(v string) *UpdateRepositoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetHttpUrlToRepo(v string) *UpdateRepositoryResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetId(v int64) *UpdateRepositoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetLastActivityAt(v string) *UpdateRepositoryResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetName(v string) *UpdateRepositoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetNameWithNamespace(v string) *UpdateRepositoryResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetNamespace(v *UpdateRepositoryResponseBodyResultNamespace) *UpdateRepositoryResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetPath(v string) *UpdateRepositoryResponseBodyResult {
	s.Path = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetPathWithNamespace(v string) *UpdateRepositoryResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetSshUrlToRepo(v string) *UpdateRepositoryResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetVisibilityLevel(v string) *UpdateRepositoryResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetWebUrl(v string) *UpdateRepositoryResponseBodyResult {
	s.WebUrl = &v
	return s
}

type UpdateRepositoryResponseBodyResultNamespace struct {
	Avatar          *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	CreatedAt       *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Public          *bool   `json:"Public,omitempty" xml:"Public,omitempty"`
	UpdatedAt       *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	VisibilityLevel *string `json:"VisibilityLevel,omitempty" xml:"VisibilityLevel,omitempty"`
}

func (s UpdateRepositoryResponseBodyResultNamespace) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponseBodyResultNamespace) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetAvatar(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Avatar = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetCreatedAt(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetDescription(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetId(v int64) *UpdateRepositoryResponseBodyResultNamespace {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetName(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetOwnerId(v int64) *UpdateRepositoryResponseBodyResultNamespace {
	s.OwnerId = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetPath(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetPublic(v bool) *UpdateRepositoryResponseBodyResultNamespace {
	s.Public = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetUpdatedAt(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetVisibilityLevel(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

type UpdateRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponse) SetHeaders(v map[string]*string) *UpdateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepositoryResponse) SetStatusCode(v int32) *UpdateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepositoryResponse) SetBody(v *UpdateRepositoryResponseBody) *UpdateRepositoryResponse {
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
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *UpdateRepositoryMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRepositoryMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponseBody) SetErrorCode(v string) *UpdateRepositoryMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetErrorMessage(v string) *UpdateRepositoryMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetRequestId(v string) *UpdateRepositoryMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetResult(v *UpdateRepositoryMemberResponseBodyResult) *UpdateRepositoryMemberResponseBody {
	s.Result = v
	return s
}

func (s *UpdateRepositoryMemberResponseBody) SetSuccess(v bool) *UpdateRepositoryMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateRepositoryMemberResponseBodyResult struct {
	AccessLevel  *int32  `json:"AccessLevel,omitempty" xml:"AccessLevel,omitempty"`
	AvatarUrl    *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ExternUserId *string `json:"ExternUserId,omitempty" xml:"ExternUserId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdateRepositoryMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetAccessLevel(v int32) *UpdateRepositoryMemberResponseBodyResult {
	s.AccessLevel = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetAvatarUrl(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetEmail(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.Email = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetExternUserId(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.ExternUserId = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetId(v int64) *UpdateRepositoryMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryMemberResponseBodyResult) SetState(v string) *UpdateRepositoryMemberResponseBodyResult {
	s.State = &v
	return s
}

type UpdateRepositoryMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateRepositoryMemberResponse) SetStatusCode(v int32) *UpdateRepositoryMemberResponse {
	s.StatusCode = &v
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
	client.SignatureAlgorithm = tea.String("v2")
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

func (client *Client) AcceptMergeRequest(ProjectId *string, MergeRequestId *string, request *AcceptMergeRequestRequest) (_result *AcceptMergeRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AcceptMergeRequestResponse{}
	_body, _err := client.AcceptMergeRequestWithOptions(ProjectId, MergeRequestId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AcceptMergeRequestWithOptions(ProjectId *string, MergeRequestId *string, request *AcceptMergeRequestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AcceptMergeRequestResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("AcceptMergeRequest"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_request/" + tea.StringValue(openapiutil.GetEncodeParam(MergeRequestId)) + "/accept"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptMergeRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGroupMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/groups/" + tea.StringValue(openapiutil.GetEncodeParam(GroupId)) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRepositoryMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRepositoryMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("AddWebhook"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/hooks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		body["branchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.Ref)) {
		body["ref"] = request.Ref
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBranch"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/branches"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFile"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/files"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("CreateMergeRequest"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_requests"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMergeRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMergeRequestComment(ProjectId *string, MergeRequestId *string, request *CreateMergeRequestCommentRequest) (_result *CreateMergeRequestCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMergeRequestCommentResponse{}
	_body, _err := client.CreateMergeRequestCommentWithOptions(ProjectId, MergeRequestId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMergeRequestCommentWithOptions(ProjectId *string, MergeRequestId *string, request *CreateMergeRequestCommentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMergeRequestCommentResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("CreateMergeRequestComment"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_request/" + tea.StringValue(openapiutil.GetEncodeParam(MergeRequestId)) + "/comments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMergeRequestCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.CreateParentPath)) {
		query["CreateParentPath"] = request.CreateParentPath
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Sync)) {
		query["Sync"] = request.Sync
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepository"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRepositoryDeployKey(ProjectId *string, request *CreateRepositoryDeployKeyRequest) (_result *CreateRepositoryDeployKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRepositoryDeployKeyResponse{}
	_body, _err := client.CreateRepositoryDeployKeyWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRepositoryDeployKeyWithOptions(ProjectId *string, request *CreateRepositoryDeployKeyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRepositoryDeployKeyResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("CreateRepositoryDeployKey"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/keys"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepositoryDeployKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepositoryGroup"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepositoryGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRepositoryProtectedBranch(ProjectId *string, request *CreateRepositoryProtectedBranchRequest) (_result *CreateRepositoryProtectedBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRepositoryProtectedBranchResponse{}
	_body, _err := client.CreateRepositoryProtectedBranchWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRepositoryProtectedBranchWithOptions(ProjectId *string, request *CreateRepositoryProtectedBranchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRepositoryProtectedBranchResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("CreateRepositoryProtectedBranch"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/protect_branches"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepositoryProtectedBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSshKey(request *CreateSshKeyRequest) (_result *CreateSshKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSshKeyResponse{}
	_body, _err := client.CreateSshKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSshKeyWithOptions(request *CreateSshKeyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSshKeyResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("CreateSshKey"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/user/keys"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSshKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("CreateTag"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/tags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteBranch"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/branches/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.BranchName)) {
		query["BranchName"] = request.BranchName
	}

	if !tea.BoolValue(util.IsUnset(request.CommitMessage)) {
		query["CommitMessage"] = request.CommitMessage
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
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
	params := &openapi.Params{
		Action:      tea.String("DeleteFile"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/files"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteGroupMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/groups/" + tea.StringValue(openapiutil.GetEncodeParam(GroupId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(UserId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteRepository"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteRepositoryGroup"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/groups/" + tea.StringValue(openapiutil.GetEncodeParam(GroupId)) + "/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteRepositoryMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(UserId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepositoryMemberWithExternUid(ProjectId *string, request *DeleteRepositoryMemberWithExternUidRequest) (_result *DeleteRepositoryMemberWithExternUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryMemberWithExternUidResponse{}
	_body, _err := client.DeleteRepositoryMemberWithExternUidWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryMemberWithExternUidWithOptions(ProjectId *string, request *DeleteRepositoryMemberWithExternUidRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryMemberWithExternUidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.ExternUserId)) {
		query["ExternUserId"] = request.ExternUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepositoryMemberWithExternUid"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/members/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryMemberWithExternUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepositoryProtectedBranch(ProjectId *string, ProtectedBranchId *string, request *DeleteRepositoryProtectedBranchRequest) (_result *DeleteRepositoryProtectedBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryProtectedBranchResponse{}
	_body, _err := client.DeleteRepositoryProtectedBranchWithOptions(ProjectId, ProtectedBranchId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryProtectedBranchWithOptions(ProjectId *string, ProtectedBranchId *string, request *DeleteRepositoryProtectedBranchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryProtectedBranchResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("DeleteRepositoryProtectedBranch"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/protect_branches/" + tea.StringValue(openapiutil.GetEncodeParam(ProtectedBranchId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryProtectedBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteRepositoryTag"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/tags/" + tea.StringValue(openapiutil.GetEncodeParam(TagName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRepositoryTagV2(ProjectId *string, request *DeleteRepositoryTagV2Request) (_result *DeleteRepositoryTagV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRepositoryTagV2Response{}
	_body, _err := client.DeleteRepositoryTagV2WithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRepositoryTagV2WithOptions(ProjectId *string, request *DeleteRepositoryTagV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRepositoryTagV2Response, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepositoryTagV2"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/tag/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryTagV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteRepositoryWebhook"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/hooks/" + tea.StringValue(openapiutil.GetEncodeParam(WebhookId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableRepositoryDeployKey(ProjectId *string, KeyId *string, request *EnableRepositoryDeployKeyRequest) (_result *EnableRepositoryDeployKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableRepositoryDeployKeyResponse{}
	_body, _err := client.EnableRepositoryDeployKeyWithOptions(ProjectId, KeyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableRepositoryDeployKeyWithOptions(ProjectId *string, KeyId *string, request *EnableRepositoryDeployKeyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableRepositoryDeployKeyResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("EnableRepositoryDeployKey"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/keys/" + tea.StringValue(openapiutil.GetEncodeParam(KeyId)) + "/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableRepositoryDeployKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetBranchInfo"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/branches/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBranchInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCodeCompletion(ServiceName *string, request *GetCodeCompletionRequest) (_result *GetCodeCompletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeCompletionResponse{}
	_body, _err := client.GetCodeCompletionWithOptions(ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCodeCompletionWithOptions(ServiceName *string, request *GetCodeCompletionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCodeCompletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FetchKeys)) {
		query["FetchKeys"] = request.FetchKeys
	}

	if !tea.BoolValue(util.IsUnset(request.IsEncrypted)) {
		query["IsEncrypted"] = request.IsEncrypted
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCodeCompletion"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/service/invoke/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCodeCompletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetCodeupOrganization"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/organization/" + tea.StringValue(openapiutil.GetEncodeParam(OrganizationIdentity))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCodeupOrganizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Ref)) {
		query["Ref"] = request.Ref
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["To"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileBlobs"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/blobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileBlobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileLastCommit(ProjectId *string, request *GetFileLastCommitRequest) (_result *GetFileLastCommitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileLastCommitResponse{}
	_body, _err := client.GetFileLastCommitWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileLastCommitWithOptions(ProjectId *string, request *GetFileLastCommitRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileLastCommitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Sha)) {
		query["Sha"] = request.Sha
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileLastCommit"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/files/last_commit"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileLastCommitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetGroupDetail"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/groups/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMergeRequestApproveStatus(ProjectId *string, MergeRequestId *string, request *GetMergeRequestApproveStatusRequest) (_result *GetMergeRequestApproveStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMergeRequestApproveStatusResponse{}
	_body, _err := client.GetMergeRequestApproveStatusWithOptions(ProjectId, MergeRequestId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMergeRequestApproveStatusWithOptions(ProjectId *string, MergeRequestId *string, request *GetMergeRequestApproveStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMergeRequestApproveStatusResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetMergeRequestApproveStatus"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_request/" + tea.StringValue(openapiutil.GetEncodeParam(MergeRequestId)) + "/approve_status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMergeRequestApproveStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMergeRequestDetail(ProjectId *string, MergeRequestId *string, request *GetMergeRequestDetailRequest) (_result *GetMergeRequestDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMergeRequestDetailResponse{}
	_body, _err := client.GetMergeRequestDetailWithOptions(ProjectId, MergeRequestId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMergeRequestDetailWithOptions(ProjectId *string, MergeRequestId *string, request *GetMergeRequestDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMergeRequestDetailResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetMergeRequestDetail"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_request/" + tea.StringValue(openapiutil.GetEncodeParam(MergeRequestId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMergeRequestDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMergeRequestSetting(ProjectId *string, request *GetMergeRequestSettingRequest) (_result *GetMergeRequestSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMergeRequestSettingResponse{}
	_body, _err := client.GetMergeRequestSettingWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMergeRequestSettingWithOptions(ProjectId *string, request *GetMergeRequestSettingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMergeRequestSettingResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetMergeRequestSetting"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/settings/merge_requests"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMergeRequestSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizationRepositorySetting(request *GetOrganizationRepositorySettingRequest) (_result *GetOrganizationRepositorySettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrganizationRepositorySettingResponse{}
	_body, _err := client.GetOrganizationRepositorySettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationRepositorySettingWithOptions(request *GetOrganizationRepositorySettingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOrganizationRepositorySettingResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationRepositorySetting"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/organization/settings/repo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationRepositorySettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizationSecurityCenterStatus(request *GetOrganizationSecurityCenterStatusRequest) (_result *GetOrganizationSecurityCenterStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrganizationSecurityCenterStatusResponse{}
	_body, _err := client.GetOrganizationSecurityCenterStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationSecurityCenterStatusWithOptions(request *GetOrganizationSecurityCenterStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOrganizationSecurityCenterStatusResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationSecurityCenterStatus"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/organization/security/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationSecurityCenterStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetProjectMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(UserId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRepositoryCommit(ProjectId *string, Sha *string, request *GetRepositoryCommitRequest) (_result *GetRepositoryCommitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRepositoryCommitResponse{}
	_body, _err := client.GetRepositoryCommitWithOptions(ProjectId, Sha, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRepositoryCommitWithOptions(ProjectId *string, Sha *string, request *GetRepositoryCommitRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRepositoryCommitResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetRepositoryCommit"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/commits/" + tea.StringValue(openapiutil.GetEncodeParam(Sha))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepositoryCommitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetRepositoryInfo"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/info"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepositoryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetRepositoryTag"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/tags/" + tea.StringValue(openapiutil.GetEncodeParam(TagName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepositoryTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRepositoryTagV2(ProjectId *string, request *GetRepositoryTagV2Request) (_result *GetRepositoryTagV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRepositoryTagV2Response{}
	_body, _err := client.GetRepositoryTagV2WithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRepositoryTagV2WithOptions(ProjectId *string, request *GetRepositoryTagV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRepositoryTagV2Response, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepositoryTagV2"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/tag/info"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepositoryTagV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserInfo"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/user/current"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsSlsUserAuthrizedCodeup(request *IsSlsUserAuthrizedCodeupRequest) (_result *IsSlsUserAuthrizedCodeupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IsSlsUserAuthrizedCodeupResponse{}
	_body, _err := client.IsSlsUserAuthrizedCodeupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsSlsUserAuthrizedCodeupWithOptions(request *IsSlsUserAuthrizedCodeupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *IsSlsUserAuthrizedCodeupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunSubUserId)) {
		body["aliyunSubUserId"] = request.AliyunSubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunUserId)) {
		body["aliyunUserId"] = request.AliyunUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeupProjectId)) {
		body["codeupProjectId"] = request.CodeupProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SlsLogStore)) {
		body["slsLogStore"] = request.SlsLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProject)) {
		body["slsProject"] = request.SlsProject
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IsSlsUserAuthrizedCodeup"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/repository/is_codeup_authrized"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &IsSlsUserAuthrizedCodeupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ListGroupMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/groups/" + tea.StringValue(openapiutil.GetEncodeParam(GroupId)) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.IsMember)) {
		query["IsMember"] = request.IsMember
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

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupRepositories"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/groups/" + tea.StringValue(openapiutil.GetEncodeParam(Identity)) + "/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.IncludePersonal)) {
		query["IncludePersonal"] = request.IncludePersonal
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

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/groups/all"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMergeRequestComments(ProjectId *string, MergeRequestId *string, request *ListMergeRequestCommentsRequest) (_result *ListMergeRequestCommentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMergeRequestCommentsResponse{}
	_body, _err := client.ListMergeRequestCommentsWithOptions(ProjectId, MergeRequestId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMergeRequestCommentsWithOptions(ProjectId *string, MergeRequestId *string, request *ListMergeRequestCommentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMergeRequestCommentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.FromCommit)) {
		query["FromCommit"] = request.FromCommit
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.ToCommit)) {
		query["ToCommit"] = request.ToCommit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMergeRequestComments"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_request/" + tea.StringValue(openapiutil.GetEncodeParam(MergeRequestId)) + "/comments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMergeRequestCommentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMergeRequests(request *ListMergeRequestsRequest) (_result *ListMergeRequestsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMergeRequestsResponse{}
	_body, _err := client.ListMergeRequestsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMergeRequestsWithOptions(request *ListMergeRequestsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMergeRequestsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.AfterDate)) {
		query["AfterDate"] = request.AfterDate
	}

	if !tea.BoolValue(util.IsUnset(request.AssigneeCodeupIdList)) {
		query["AssigneeCodeupIdList"] = request.AssigneeCodeupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AssigneeIdList)) {
		query["AssigneeIdList"] = request.AssigneeIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorCodeupIdList)) {
		query["AuthorCodeupIdList"] = request.AuthorCodeupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorIdList)) {
		query["AuthorIdList"] = request.AuthorIdList
	}

	if !tea.BoolValue(util.IsUnset(request.BeforeDate)) {
		query["BeforeDate"] = request.BeforeDate
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIdList)) {
		query["GroupIdList"] = request.GroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
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

	if !tea.BoolValue(util.IsUnset(request.ProjectIdList)) {
		query["ProjectIdList"] = request.ProjectIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriberCodeupIdList)) {
		query["SubscriberCodeupIdList"] = request.SubscriberCodeupIdList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMergeRequests"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/merge_requests/advanced_search"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMergeRequestsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrganizationSecurityScores(request *ListOrganizationSecurityScoresRequest) (_result *ListOrganizationSecurityScoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOrganizationSecurityScoresResponse{}
	_body, _err := client.ListOrganizationSecurityScoresWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrganizationSecurityScoresWithOptions(request *ListOrganizationSecurityScoresRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOrganizationSecurityScoresResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationSecurityScores"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/organization/security/scores"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationSecurityScoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AccessLevel)) {
		query["AccessLevel"] = request.AccessLevel
	}

	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.MinAccessLevel)) {
		query["MinAccessLevel"] = request.MinAccessLevel
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizations"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/organization"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Archive)) {
		query["Archive"] = request.Archive
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
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

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositories"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/all"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryBranches"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/branches"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryBranchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryCode(request *ListRepositoryCodeRequest) (_result *ListRepositoryCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryCodeResponse{}
	_body, _err := client.ListRepositoryCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryCodeWithOptions(request *ListRepositoryCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FilePath))) {
		body["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.IsCodeBlock)) {
		body["IsCodeBlock"] = request.IsCodeBlock
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		body["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.RepositoryPath))) {
		body["RepositoryPath"] = request.RepositoryPath
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		body["Sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryCode"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/search/v3/code"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryCommitDiff(ProjectId *string, Sha *string, request *ListRepositoryCommitDiffRequest) (_result *ListRepositoryCommitDiffResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryCommitDiffResponse{}
	_body, _err := client.ListRepositoryCommitDiffWithOptions(ProjectId, Sha, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryCommitDiffWithOptions(ProjectId *string, Sha *string, request *ListRepositoryCommitDiffRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryCommitDiffResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContextLine)) {
		query["ContextLine"] = request.ContextLine
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryCommitDiff"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/commits/" + tea.StringValue(openapiutil.GetEncodeParam(Sha)) + "/diff"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryCommitDiffResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryCommits(ProjectId *string, request *ListRepositoryCommitsRequest) (_result *ListRepositoryCommitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryCommitsResponse{}
	_body, _err := client.ListRepositoryCommitsWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryCommitsWithOptions(ProjectId *string, request *ListRepositoryCommitsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryCommitsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.RefName)) {
		query["RefName"] = request.RefName
	}

	if !tea.BoolValue(util.IsUnset(request.Search)) {
		query["Search"] = request.Search
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSignature)) {
		query["ShowSignature"] = request.ShowSignature
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryCommits"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/commits"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryCommitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryMemberWithInherited(ProjectId *string, request *ListRepositoryMemberWithInheritedRequest) (_result *ListRepositoryMemberWithInheritedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryMemberWithInheritedResponse{}
	_body, _err := client.ListRepositoryMemberWithInheritedWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryMemberWithInheritedWithOptions(ProjectId *string, request *ListRepositoryMemberWithInheritedRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryMemberWithInheritedResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryMemberWithInherited"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/all_members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryMemberWithInheritedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRepositoryProtectedBranch(ProjectId *string, request *ListRepositoryProtectedBranchRequest) (_result *ListRepositoryProtectedBranchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRepositoryProtectedBranchResponse{}
	_body, _err := client.ListRepositoryProtectedBranchWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRepositoryProtectedBranchWithOptions(ProjectId *string, request *ListRepositoryProtectedBranchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRepositoryProtectedBranchResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryProtectedBranch"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/protect_branches"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryProtectedBranchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
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

	if !tea.BoolValue(util.IsUnset(request.ShowSignature)) {
		query["ShowSignature"] = request.ShowSignature
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryTags"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.RefName)) {
		query["RefName"] = request.RefName
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryTree"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/tree"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ListRepositoryWebhook"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/hooks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("MergeMergeRequest"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_request/" + tea.StringValue(openapiutil.GetEncodeParam(MergeRequestId)) + "/merge"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MergeMergeRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySlsRelation(request *QuerySlsRelationRequest) (_result *QuerySlsRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySlsRelationResponse{}
	_body, _err := client.QuerySlsRelationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySlsRelationWithOptions(request *QuerySlsRelationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QuerySlsRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUserId)) {
		body["aliyunUserId"] = request.AliyunUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeupProjectId)) {
		body["codeupProjectId"] = request.CodeupProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SlsLogStore)) {
		body["slsLogStore"] = request.SlsLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProject)) {
		body["slsProject"] = request.SlsProject
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySlsRelation"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/repository/query_sls_relation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySlsRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RelatedSlsLogStore(request *RelatedSlsLogStoreRequest) (_result *RelatedSlsLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RelatedSlsLogStoreResponse{}
	_body, _err := client.RelatedSlsLogStoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RelatedSlsLogStoreWithOptions(request *RelatedSlsLogStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RelatedSlsLogStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUserId)) {
		body["aliyunUserId"] = request.AliyunUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeupProjectId)) {
		body["codeupProjectId"] = request.CodeupProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultViewer)) {
		body["defaultViewer"] = request.DefaultViewer
	}

	if !tea.BoolValue(util.IsUnset(request.SlsLogStore)) {
		body["slsLogStore"] = request.SlsLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProject)) {
		body["slsProject"] = request.SlsProject
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RelatedSlsLogStore"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/repository/related_to_sls_log_store"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RelatedSlsLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerRepositoryMirrorSync(ProjectId *string, request *TriggerRepositoryMirrorSyncRequest) (_result *TriggerRepositoryMirrorSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TriggerRepositoryMirrorSyncResponse{}
	_body, _err := client.TriggerRepositoryMirrorSyncWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerRepositoryMirrorSyncWithOptions(ProjectId *string, request *TriggerRepositoryMirrorSyncRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TriggerRepositoryMirrorSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerRepositoryMirrorSync"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/mirror"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerRepositoryMirrorSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnRelatedSlsLogStore(request *UnRelatedSlsLogStoreRequest) (_result *UnRelatedSlsLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnRelatedSlsLogStoreResponse{}
	_body, _err := client.UnRelatedSlsLogStoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnRelatedSlsLogStoreWithOptions(request *UnRelatedSlsLogStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnRelatedSlsLogStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrganizationId)) {
		query["OrganizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUserId)) {
		body["aliyunUserId"] = request.AliyunUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeupProjectId)) {
		body["codeupProjectId"] = request.CodeupProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SlsLogStore)) {
		body["slsLogStore"] = request.SlsLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProject)) {
		body["slsProject"] = request.SlsProject
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnRelatedSlsLogStore"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/repository/unrelated_to_sls_log_store"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnRelatedSlsLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("UpdateFile"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/repository/files"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/groups/" + tea.StringValue(openapiutil.GetEncodeParam(GroupId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(UserId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMergeRequest(ProjectId *string, MergeRequestId *string, request *UpdateMergeRequestRequest) (_result *UpdateMergeRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMergeRequestResponse{}
	_body, _err := client.UpdateMergeRequestWithOptions(ProjectId, MergeRequestId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMergeRequestWithOptions(ProjectId *string, MergeRequestId *string, request *UpdateMergeRequestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMergeRequestResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("UpdateMergeRequest"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_request/" + tea.StringValue(openapiutil.GetEncodeParam(MergeRequestId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMergeRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMergeRequestComment(ProjectId *string, MergeRequestId *string, NoteId *string, request *UpdateMergeRequestCommentRequest) (_result *UpdateMergeRequestCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMergeRequestCommentResponse{}
	_body, _err := client.UpdateMergeRequestCommentWithOptions(ProjectId, MergeRequestId, NoteId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMergeRequestCommentWithOptions(ProjectId *string, MergeRequestId *string, NoteId *string, request *UpdateMergeRequestCommentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMergeRequestCommentResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("UpdateMergeRequestComment"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/merge_requests/" + tea.StringValue(openapiutil.GetEncodeParam(MergeRequestId)) + "/notes/" + tea.StringValue(openapiutil.GetEncodeParam(NoteId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMergeRequestCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMergeRequestSetting(ProjectId *string, request *UpdateMergeRequestSettingRequest) (_result *UpdateMergeRequestSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMergeRequestSettingResponse{}
	_body, _err := client.UpdateMergeRequestSettingWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMergeRequestSettingWithOptions(ProjectId *string, request *UpdateMergeRequestSettingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMergeRequestSettingResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("UpdateMergeRequestSetting"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v4/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/settings/merge_requests"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMergeRequestSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRepository(ProjectId *string, request *UpdateRepositoryRequest) (_result *UpdateRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.UpdateRepositoryWithOptions(ProjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRepositoryWithOptions(ProjectId *string, request *UpdateRepositoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRepositoryResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("UpdateRepository"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("UpdateRepositoryMember"),
		Version:     tea.String("2020-04-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v3/projects/" + tea.StringValue(openapiutil.GetEncodeParam(ProjectId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(UserId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRepositoryMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
