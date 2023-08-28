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

type AddWorkspaceDocMembersHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddWorkspaceDocMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddWorkspaceDocMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceDocMembersHeaders) SetAccountContext(v *AddWorkspaceDocMembersHeadersAccountContext) *AddWorkspaceDocMembersHeaders {
	s.AccountContext = v
	return s
}

type AddWorkspaceDocMembersHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddWorkspaceDocMembersHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersHeadersAccountContext) SetAccountId(v string) *AddWorkspaceDocMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

type AddWorkspaceDocMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddWorkspaceDocMembersShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceDocMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceDocMembersShrinkHeaders) SetAccountContextShrink(v string) *AddWorkspaceDocMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type AddWorkspaceDocMembersRequest struct {
	Members       []*AddWorkspaceDocMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	NodeId        *string                                     `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContext *AddWorkspaceDocMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkspaceId   *string                                     `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequest) SetMembers(v []*AddWorkspaceDocMembersRequestMembers) *AddWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *AddWorkspaceDocMembersRequest) SetNodeId(v string) *AddWorkspaceDocMembersRequest {
	s.NodeId = &v
	return s
}

func (s *AddWorkspaceDocMembersRequest) SetTenantContext(v *AddWorkspaceDocMembersRequestTenantContext) *AddWorkspaceDocMembersRequest {
	s.TenantContext = v
	return s
}

func (s *AddWorkspaceDocMembersRequest) SetWorkspaceId(v string) *AddWorkspaceDocMembersRequest {
	s.WorkspaceId = &v
	return s
}

type AddWorkspaceDocMembersRequestMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	RoleType   *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s AddWorkspaceDocMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequestMembers) SetMemberId(v string) *AddWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *AddWorkspaceDocMembersRequestMembers) SetMemberType(v string) *AddWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *AddWorkspaceDocMembersRequestMembers) SetRoleType(v string) *AddWorkspaceDocMembersRequestMembers {
	s.RoleType = &v
	return s
}

type AddWorkspaceDocMembersRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddWorkspaceDocMembersRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersRequestTenantContext) SetTenantId(v string) *AddWorkspaceDocMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

type AddWorkspaceDocMembersShrinkRequest struct {
	MembersShrink       *string `json:"Members,omitempty" xml:"Members,omitempty"`
	NodeId              *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceDocMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersShrinkRequest) SetMembersShrink(v string) *AddWorkspaceDocMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *AddWorkspaceDocMembersShrinkRequest) SetNodeId(v string) *AddWorkspaceDocMembersShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *AddWorkspaceDocMembersShrinkRequest) SetTenantContextShrink(v string) *AddWorkspaceDocMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddWorkspaceDocMembersShrinkRequest) SetWorkspaceId(v string) *AddWorkspaceDocMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type AddWorkspaceDocMembersResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddWorkspaceDocMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersResponseBody) SetRequestId(v string) *AddWorkspaceDocMembersResponseBody {
	s.RequestId = &v
	return s
}

type AddWorkspaceDocMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddWorkspaceDocMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWorkspaceDocMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *AddWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceDocMembersResponse) SetStatusCode(v int32) *AddWorkspaceDocMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceDocMembersResponse) SetBody(v *AddWorkspaceDocMembersResponseBody) *AddWorkspaceDocMembersResponse {
	s.Body = v
	return s
}

type AddWorkspaceMembersHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddWorkspaceMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddWorkspaceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceMembersHeaders) SetAccountContext(v *AddWorkspaceMembersHeadersAccountContext) *AddWorkspaceMembersHeaders {
	s.AccountContext = v
	return s
}

type AddWorkspaceMembersHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddWorkspaceMembersHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersHeadersAccountContext) SetAccountId(v string) *AddWorkspaceMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

type AddWorkspaceMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddWorkspaceMembersShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceMembersShrinkHeaders) SetAccountContextShrink(v string) *AddWorkspaceMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type AddWorkspaceMembersRequest struct {
	Members       []*AddWorkspaceMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	TenantContext *AddWorkspaceMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkspaceId   *string                                  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequest) SetMembers(v []*AddWorkspaceMembersRequestMembers) *AddWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *AddWorkspaceMembersRequest) SetTenantContext(v *AddWorkspaceMembersRequestTenantContext) *AddWorkspaceMembersRequest {
	s.TenantContext = v
	return s
}

func (s *AddWorkspaceMembersRequest) SetWorkspaceId(v string) *AddWorkspaceMembersRequest {
	s.WorkspaceId = &v
	return s
}

type AddWorkspaceMembersRequestMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	RoleType   *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s AddWorkspaceMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequestMembers) SetMemberId(v string) *AddWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *AddWorkspaceMembersRequestMembers) SetMemberType(v string) *AddWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *AddWorkspaceMembersRequestMembers) SetRoleType(v string) *AddWorkspaceMembersRequestMembers {
	s.RoleType = &v
	return s
}

type AddWorkspaceMembersRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddWorkspaceMembersRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersRequestTenantContext) SetTenantId(v string) *AddWorkspaceMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

type AddWorkspaceMembersShrinkRequest struct {
	MembersShrink       *string `json:"Members,omitempty" xml:"Members,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersShrinkRequest) SetMembersShrink(v string) *AddWorkspaceMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *AddWorkspaceMembersShrinkRequest) SetTenantContextShrink(v string) *AddWorkspaceMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddWorkspaceMembersShrinkRequest) SetWorkspaceId(v string) *AddWorkspaceMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type AddWorkspaceMembersResponseBody struct {
	NotInOrgList []*string `json:"NotInOrgList,omitempty" xml:"NotInOrgList,omitempty" type:"Repeated"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AddWorkspaceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersResponseBody) SetNotInOrgList(v []*string) *AddWorkspaceMembersResponseBody {
	s.NotInOrgList = v
	return s
}

func (s *AddWorkspaceMembersResponseBody) SetRequestId(v string) *AddWorkspaceMembersResponseBody {
	s.RequestId = &v
	return s
}

type AddWorkspaceMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddWorkspaceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWorkspaceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersResponse) SetHeaders(v map[string]*string) *AddWorkspaceMembersResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceMembersResponse) SetStatusCode(v int32) *AddWorkspaceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceMembersResponse) SetBody(v *AddWorkspaceMembersResponseBody) *AddWorkspaceMembersResponse {
	s.Body = v
	return s
}

type CreateSheetHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateSheetHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetHeaders) GoString() string {
	return s.String()
}

func (s *CreateSheetHeaders) SetCommonHeaders(v map[string]*string) *CreateSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSheetHeaders) SetAccountContext(v *CreateSheetHeadersAccountContext) *CreateSheetHeaders {
	s.AccountContext = v
	return s
}

type CreateSheetHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateSheetHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateSheetHeadersAccountContext) SetAccountId(v string) *CreateSheetHeadersAccountContext {
	s.AccountId = &v
	return s
}

type CreateSheetShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateSheetShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateSheetShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateSheetShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSheetShrinkHeaders) SetAccountContextShrink(v string) *CreateSheetShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type CreateSheetRequest struct {
	Name          *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContext *CreateSheetRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkbookId    *string                          `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s CreateSheetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetRequest) GoString() string {
	return s.String()
}

func (s *CreateSheetRequest) SetName(v string) *CreateSheetRequest {
	s.Name = &v
	return s
}

func (s *CreateSheetRequest) SetTenantContext(v *CreateSheetRequestTenantContext) *CreateSheetRequest {
	s.TenantContext = v
	return s
}

func (s *CreateSheetRequest) SetWorkbookId(v string) *CreateSheetRequest {
	s.WorkbookId = &v
	return s
}

type CreateSheetRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateSheetRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateSheetRequestTenantContext) SetTenantId(v string) *CreateSheetRequestTenantContext {
	s.TenantId = &v
	return s
}

type CreateSheetShrinkRequest struct {
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkbookId          *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s CreateSheetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSheetShrinkRequest) SetName(v string) *CreateSheetShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSheetShrinkRequest) SetTenantContextShrink(v string) *CreateSheetShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateSheetShrinkRequest) SetWorkbookId(v string) *CreateSheetShrinkRequest {
	s.WorkbookId = &v
	return s
}

type CreateSheetResponseBody struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s CreateSheetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSheetResponseBody) SetId(v string) *CreateSheetResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSheetResponseBody) SetName(v string) *CreateSheetResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSheetResponseBody) SetRequestId(v string) *CreateSheetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSheetResponseBody) SetVisibility(v string) *CreateSheetResponseBody {
	s.Visibility = &v
	return s
}

type CreateSheetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSheetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetResponse) GoString() string {
	return s.String()
}

func (s *CreateSheetResponse) SetHeaders(v map[string]*string) *CreateSheetResponse {
	s.Headers = v
	return s
}

func (s *CreateSheetResponse) SetStatusCode(v int32) *CreateSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSheetResponse) SetBody(v *CreateSheetResponseBody) *CreateSheetResponse {
	s.Body = v
	return s
}

type CreateTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTodoTaskHeaders) SetAccountContext(v *CreateTodoTaskHeadersAccountContext) *CreateTodoTaskHeaders {
	s.AccountContext = v
	return s
}

type CreateTodoTaskHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateTodoTaskHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskHeadersAccountContext) SetAccountId(v string) *CreateTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

type CreateTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateTodoTaskShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *CreateTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type CreateTodoTaskRequest struct {
	TenantContext      *CreateTodoTaskRequestTenantContext      `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	ContentFieldList   []*CreateTodoTaskRequestContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	CreatorId          *string                                  `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description        *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl          *CreateTodoTaskRequestDetailUrl          `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	DueTime            *int64                                   `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds        []*string                                `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	IsOnlyShowExecutor *bool                                    `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	NotifyConfigs      *CreateTodoTaskRequestNotifyConfigs      `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	OperatorId         *string                                  `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	ParticipantIds     []*string                                `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Priority           *int32                                   `json:"priority,omitempty" xml:"priority,omitempty"`
	SourceId           *string                                  `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Subject            *string                                  `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreateTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequest) SetTenantContext(v *CreateTodoTaskRequestTenantContext) *CreateTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *CreateTodoTaskRequest) SetContentFieldList(v []*CreateTodoTaskRequestContentFieldList) *CreateTodoTaskRequest {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTaskRequest) SetCreatorId(v string) *CreateTodoTaskRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetDescription(v string) *CreateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskRequest) SetDetailUrl(v *CreateTodoTaskRequestDetailUrl) *CreateTodoTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskRequest) SetDueTime(v int64) *CreateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskRequest) SetExecutorIds(v []*string) *CreateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetIsOnlyShowExecutor(v bool) *CreateTodoTaskRequest {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *CreateTodoTaskRequest) SetNotifyConfigs(v *CreateTodoTaskRequestNotifyConfigs) *CreateTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *CreateTodoTaskRequest) SetOperatorId(v string) *CreateTodoTaskRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetParticipantIds(v []*string) *CreateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetPriority(v int32) *CreateTodoTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateTodoTaskRequest) SetSourceId(v string) *CreateTodoTaskRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetSubject(v string) *CreateTodoTaskRequest {
	s.Subject = &v
	return s
}

type CreateTodoTaskRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateTodoTaskRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestTenantContext) SetTenantId(v string) *CreateTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

type CreateTodoTaskRequestContentFieldList struct {
	// fieldKey
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// fieldValue
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s CreateTodoTaskRequestContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestContentFieldList) SetFieldKey(v string) *CreateTodoTaskRequestContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTaskRequestContentFieldList) SetFieldValue(v string) *CreateTodoTaskRequestContentFieldList {
	s.FieldValue = &v
	return s
}

type CreateTodoTaskRequestDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s CreateTodoTaskRequestDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestDetailUrl) SetAppUrl(v string) *CreateTodoTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateTodoTaskRequestDetailUrl) SetPcUrl(v string) *CreateTodoTaskRequestDetailUrl {
	s.PcUrl = &v
	return s
}

type CreateTodoTaskRequestNotifyConfigs struct {
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
}

func (s CreateTodoTaskRequestNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

type CreateTodoTaskShrinkRequest struct {
	TenantContextShrink    *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	ContentFieldListShrink *string `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty"`
	CreatorId              *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description            *string `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrlShrink        *string `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	DueTime                *int64  `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIdsShrink      *string `json:"executorIds,omitempty" xml:"executorIds,omitempty"`
	IsOnlyShowExecutor     *bool   `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	NotifyConfigsShrink    *string `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty"`
	OperatorId             *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	ParticipantIdsShrink   *string `json:"participantIds,omitempty" xml:"participantIds,omitempty"`
	Priority               *int32  `json:"priority,omitempty" xml:"priority,omitempty"`
	SourceId               *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Subject                *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreateTodoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskShrinkRequest) SetTenantContextShrink(v string) *CreateTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetContentFieldListShrink(v string) *CreateTodoTaskShrinkRequest {
	s.ContentFieldListShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetCreatorId(v string) *CreateTodoTaskShrinkRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetDescription(v string) *CreateTodoTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetDetailUrlShrink(v string) *CreateTodoTaskShrinkRequest {
	s.DetailUrlShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetDueTime(v int64) *CreateTodoTaskShrinkRequest {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetExecutorIdsShrink(v string) *CreateTodoTaskShrinkRequest {
	s.ExecutorIdsShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetIsOnlyShowExecutor(v bool) *CreateTodoTaskShrinkRequest {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetNotifyConfigsShrink(v string) *CreateTodoTaskShrinkRequest {
	s.NotifyConfigsShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetOperatorId(v string) *CreateTodoTaskShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetParticipantIdsShrink(v string) *CreateTodoTaskShrinkRequest {
	s.ParticipantIdsShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetPriority(v int32) *CreateTodoTaskShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetSourceId(v string) *CreateTodoTaskShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetSubject(v string) *CreateTodoTaskShrinkRequest {
	s.Subject = &v
	return s
}

type CreateTodoTaskResponseBody struct {
	BizTag             *string                                       `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	ContentFieldList   []*CreateTodoTaskResponseBodyContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	CreatedTime        *int64                                        `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId          *string                                       `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description        *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl          *CreateTodoTaskResponseBodyDetailUrl          `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	Done               *bool                                         `json:"done,omitempty" xml:"done,omitempty"`
	DueTime            *int64                                        `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds        []*string                                     `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	FinishTime         *int64                                        `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	Id                 *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	IsOnlyShowExecutor *bool                                         `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	ModifiedTime       *int64                                        `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	ModifierId         *string                                       `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	NotifyConfigs      *CreateTodoTaskResponseBodyNotifyConfigs      `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	ParticipantIds     []*string                                     `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Priority           *int32                                        `json:"priority,omitempty" xml:"priority,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Source    *string `json:"source,omitempty" xml:"source,omitempty"`
	SourceId  *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Subject   *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreateTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBody) SetBizTag(v string) *CreateTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetContentFieldList(v []*CreateTodoTaskResponseBodyContentFieldList) *CreateTodoTaskResponseBody {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatedTime(v int64) *CreateTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatorId(v string) *CreateTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDescription(v string) *CreateTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDetailUrl(v *CreateTodoTaskResponseBodyDetailUrl) *CreateTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDone(v bool) *CreateTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDueTime(v int64) *CreateTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetExecutorIds(v []*string) *CreateTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetFinishTime(v int64) *CreateTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetId(v string) *CreateTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetIsOnlyShowExecutor(v bool) *CreateTodoTaskResponseBody {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifiedTime(v int64) *CreateTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifierId(v string) *CreateTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetNotifyConfigs(v *CreateTodoTaskResponseBodyNotifyConfigs) *CreateTodoTaskResponseBody {
	s.NotifyConfigs = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetParticipantIds(v []*string) *CreateTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetPriority(v int32) *CreateTodoTaskResponseBody {
	s.Priority = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetRequestId(v string) *CreateTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSource(v string) *CreateTodoTaskResponseBody {
	s.Source = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSourceId(v string) *CreateTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetStartTime(v int64) *CreateTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSubject(v string) *CreateTodoTaskResponseBody {
	s.Subject = &v
	return s
}

type CreateTodoTaskResponseBodyContentFieldList struct {
	// fieldKey
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// fieldValue
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s CreateTodoTaskResponseBodyContentFieldList) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyContentFieldList) SetFieldKey(v string) *CreateTodoTaskResponseBodyContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTaskResponseBodyContentFieldList) SetFieldValue(v string) *CreateTodoTaskResponseBodyContentFieldList {
	s.FieldValue = &v
	return s
}

type CreateTodoTaskResponseBodyDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s CreateTodoTaskResponseBodyDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

type CreateTodoTaskResponseBodyNotifyConfigs struct {
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
}

func (s CreateTodoTaskResponseBodyNotifyConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponseBodyNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskResponseBodyNotifyConfigs {
	s.DingNotify = &v
	return s
}

type CreateTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponse) SetHeaders(v map[string]*string) *CreateTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTodoTaskResponse) SetStatusCode(v int32) *CreateTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTodoTaskResponse) SetBody(v *CreateTodoTaskResponseBody) *CreateTodoTaskResponse {
	s.Body = v
	return s
}

type CreateWorkspaceHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateWorkspaceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateWorkspaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceHeaders) SetAccountContext(v *CreateWorkspaceHeadersAccountContext) *CreateWorkspaceHeaders {
	s.AccountContext = v
	return s
}

type CreateWorkspaceHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateWorkspaceHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceHeadersAccountContext) SetAccountId(v string) *CreateWorkspaceHeadersAccountContext {
	s.AccountId = &v
	return s
}

type CreateWorkspaceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateWorkspaceShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceShrinkHeaders) SetAccountContextShrink(v string) *CreateWorkspaceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type CreateWorkspaceRequest struct {
	Description   *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContext *CreateWorkspaceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetName(v string) *CreateWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceRequest) SetTenantContext(v *CreateWorkspaceRequestTenantContext) *CreateWorkspaceRequest {
	s.TenantContext = v
	return s
}

type CreateWorkspaceRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateWorkspaceRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestTenantContext) SetTenantId(v string) *CreateWorkspaceRequestTenantContext {
	s.TenantId = &v
	return s
}

type CreateWorkspaceShrinkRequest struct {
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s CreateWorkspaceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceShrinkRequest) SetDescription(v string) *CreateWorkspaceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceShrinkRequest) SetName(v string) *CreateWorkspaceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceShrinkRequest) SetTenantContextShrink(v string) *CreateWorkspaceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	// requestId
	RequestId   *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetDescription(v string) *CreateWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetName(v string) *CreateWorkspaceResponseBody {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetUrl(v string) *CreateWorkspaceResponseBody {
	s.Url = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspaceId(v string) *CreateWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResponse) SetStatusCode(v int32) *CreateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type CreateWorkspaceDocHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateWorkspaceDocHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateWorkspaceDocHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceDocHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceDocHeaders) SetAccountContext(v *CreateWorkspaceDocHeadersAccountContext) *CreateWorkspaceDocHeaders {
	s.AccountContext = v
	return s
}

type CreateWorkspaceDocHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateWorkspaceDocHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocHeadersAccountContext) SetAccountId(v string) *CreateWorkspaceDocHeadersAccountContext {
	s.AccountId = &v
	return s
}

type CreateWorkspaceDocShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateWorkspaceDocShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceDocShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceDocShrinkHeaders) SetAccountContextShrink(v string) *CreateWorkspaceDocShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type CreateWorkspaceDocRequest struct {
	DocType       *string                                 `json:"DocType,omitempty" xml:"DocType,omitempty"`
	Name          *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentNodeId  *string                                 `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	TemplateId    *string                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType  *string                                 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TenantContext *CreateWorkspaceDocRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkspaceId   *string                                 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceDocRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocRequest) SetDocType(v string) *CreateWorkspaceDocRequest {
	s.DocType = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetName(v string) *CreateWorkspaceDocRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetParentNodeId(v string) *CreateWorkspaceDocRequest {
	s.ParentNodeId = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetTemplateId(v string) *CreateWorkspaceDocRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetTemplateType(v string) *CreateWorkspaceDocRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateWorkspaceDocRequest) SetTenantContext(v *CreateWorkspaceDocRequestTenantContext) *CreateWorkspaceDocRequest {
	s.TenantContext = v
	return s
}

func (s *CreateWorkspaceDocRequest) SetWorkspaceId(v string) *CreateWorkspaceDocRequest {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceDocRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateWorkspaceDocRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocRequestTenantContext) SetTenantId(v string) *CreateWorkspaceDocRequestTenantContext {
	s.TenantId = &v
	return s
}

type CreateWorkspaceDocShrinkRequest struct {
	DocType             *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentNodeId        *string `json:"ParentNodeId,omitempty" xml:"ParentNodeId,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateType        *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceDocShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocShrinkRequest) SetDocType(v string) *CreateWorkspaceDocShrinkRequest {
	s.DocType = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetName(v string) *CreateWorkspaceDocShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetParentNodeId(v string) *CreateWorkspaceDocShrinkRequest {
	s.ParentNodeId = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetTemplateId(v string) *CreateWorkspaceDocShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetTemplateType(v string) *CreateWorkspaceDocShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetTenantContextShrink(v string) *CreateWorkspaceDocShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateWorkspaceDocShrinkRequest) SetWorkspaceId(v string) *CreateWorkspaceDocShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceDocResponseBody struct {
	DocKey *string `json:"docKey,omitempty" xml:"docKey,omitempty"`
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// requestId
	RequestId   *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocResponseBody) SetDocKey(v string) *CreateWorkspaceDocResponseBody {
	s.DocKey = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetNodeId(v string) *CreateWorkspaceDocResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetRequestId(v string) *CreateWorkspaceDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetUrl(v string) *CreateWorkspaceDocResponseBody {
	s.Url = &v
	return s
}

func (s *CreateWorkspaceDocResponseBody) SetWorkspaceId(v string) *CreateWorkspaceDocResponseBody {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceDocResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWorkspaceDocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkspaceDocResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceDocResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocResponse) SetHeaders(v map[string]*string) *CreateWorkspaceDocResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceDocResponse) SetStatusCode(v int32) *CreateWorkspaceDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceDocResponse) SetBody(v *CreateWorkspaceDocResponseBody) *CreateWorkspaceDocResponse {
	s.Body = v
	return s
}

type DeleteTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *DeleteTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTodoTaskHeaders) SetAccountContext(v *DeleteTodoTaskHeadersAccountContext) *DeleteTodoTaskHeaders {
	s.AccountContext = v
	return s
}

type DeleteTodoTaskHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteTodoTaskHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskHeadersAccountContext) SetAccountId(v string) *DeleteTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

type DeleteTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteTodoTaskShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *DeleteTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type DeleteTodoTaskRequest struct {
	TenantContext *DeleteTodoTaskRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	OperatorId    *string                             `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TaskId        *string                             `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskRequest) SetTenantContext(v *DeleteTodoTaskRequestTenantContext) *DeleteTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteTodoTaskRequest) SetOperatorId(v string) *DeleteTodoTaskRequest {
	s.OperatorId = &v
	return s
}

func (s *DeleteTodoTaskRequest) SetTaskId(v string) *DeleteTodoTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteTodoTaskRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteTodoTaskRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskRequestTenantContext) SetTenantId(v string) *DeleteTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

type DeleteTodoTaskShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	OperatorId          *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TaskId              *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteTodoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskShrinkRequest) SetTenantContextShrink(v string) *DeleteTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteTodoTaskShrinkRequest) SetOperatorId(v string) *DeleteTodoTaskShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *DeleteTodoTaskShrinkRequest) SetTaskId(v string) *DeleteTodoTaskShrinkRequest {
	s.TaskId = &v
	return s
}

type DeleteTodoTaskResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponseBody) SetRequestId(v string) *DeleteTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTodoTaskResponseBody) SetResult(v bool) *DeleteTodoTaskResponseBody {
	s.Result = &v
	return s
}

type DeleteTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponse) SetHeaders(v map[string]*string) *DeleteTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTodoTaskResponse) SetStatusCode(v int32) *DeleteTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTodoTaskResponse) SetBody(v *DeleteTodoTaskResponseBody) *DeleteTodoTaskResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceDocMembersHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteWorkspaceDocMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteWorkspaceDocMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceDocMembersHeaders) SetAccountContext(v *DeleteWorkspaceDocMembersHeadersAccountContext) *DeleteWorkspaceDocMembersHeaders {
	s.AccountContext = v
	return s
}

type DeleteWorkspaceDocMembersHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteWorkspaceDocMembersHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersHeadersAccountContext) SetAccountId(v string) *DeleteWorkspaceDocMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

type DeleteWorkspaceDocMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteWorkspaceDocMembersShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceDocMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkHeaders) SetAccountContextShrink(v string) *DeleteWorkspaceDocMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type DeleteWorkspaceDocMembersRequest struct {
	Members       []*DeleteWorkspaceDocMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	NodeId        *string                                        `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContext *DeleteWorkspaceDocMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkspaceId   *string                                        `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequest) SetMembers(v []*DeleteWorkspaceDocMembersRequestMembers) *DeleteWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) SetNodeId(v string) *DeleteWorkspaceDocMembersRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) SetTenantContext(v *DeleteWorkspaceDocMembersRequestTenantContext) *DeleteWorkspaceDocMembersRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteWorkspaceDocMembersRequest) SetWorkspaceId(v string) *DeleteWorkspaceDocMembersRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteWorkspaceDocMembersRequestMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s DeleteWorkspaceDocMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequestMembers) SetMemberId(v string) *DeleteWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersRequestMembers) SetMemberType(v string) *DeleteWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

type DeleteWorkspaceDocMembersRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteWorkspaceDocMembersRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersRequestTenantContext) SetTenantId(v string) *DeleteWorkspaceDocMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

type DeleteWorkspaceDocMembersShrinkRequest struct {
	MembersShrink       *string `json:"Members,omitempty" xml:"Members,omitempty"`
	NodeId              *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceDocMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) SetMembersShrink(v string) *DeleteWorkspaceDocMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) SetNodeId(v string) *DeleteWorkspaceDocMembersShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) SetTenantContextShrink(v string) *DeleteWorkspaceDocMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkRequest) SetWorkspaceId(v string) *DeleteWorkspaceDocMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteWorkspaceDocMembersResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteWorkspaceDocMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersResponseBody) SetRequestId(v string) *DeleteWorkspaceDocMembersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWorkspaceDocMembersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWorkspaceDocMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWorkspaceDocMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceDocMembersResponse) SetStatusCode(v int32) *DeleteWorkspaceDocMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceDocMembersResponse) SetBody(v *DeleteWorkspaceDocMembersResponseBody) *DeleteWorkspaceDocMembersResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceMembersHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteWorkspaceMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteWorkspaceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceMembersHeaders) SetAccountContext(v *DeleteWorkspaceMembersHeadersAccountContext) *DeleteWorkspaceMembersHeaders {
	s.AccountContext = v
	return s
}

type DeleteWorkspaceMembersHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteWorkspaceMembersHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersHeadersAccountContext) SetAccountId(v string) *DeleteWorkspaceMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

type DeleteWorkspaceMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteWorkspaceMembersShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceMembersShrinkHeaders) SetAccountContextShrink(v string) *DeleteWorkspaceMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type DeleteWorkspaceMembersRequest struct {
	Members       []*DeleteWorkspaceMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	TenantContext *DeleteWorkspaceMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkspaceId   *string                                     `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequest) SetMembers(v []*DeleteWorkspaceMembersRequestMembers) *DeleteWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *DeleteWorkspaceMembersRequest) SetTenantContext(v *DeleteWorkspaceMembersRequestTenantContext) *DeleteWorkspaceMembersRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteWorkspaceMembersRequest) SetWorkspaceId(v string) *DeleteWorkspaceMembersRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteWorkspaceMembersRequestMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s DeleteWorkspaceMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequestMembers) SetMemberId(v string) *DeleteWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *DeleteWorkspaceMembersRequestMembers) SetMemberType(v string) *DeleteWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

type DeleteWorkspaceMembersRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteWorkspaceMembersRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersRequestTenantContext) SetTenantId(v string) *DeleteWorkspaceMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

type DeleteWorkspaceMembersShrinkRequest struct {
	MembersShrink       *string `json:"Members,omitempty" xml:"Members,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersShrinkRequest) SetMembersShrink(v string) *DeleteWorkspaceMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *DeleteWorkspaceMembersShrinkRequest) SetTenantContextShrink(v string) *DeleteWorkspaceMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteWorkspaceMembersShrinkRequest) SetWorkspaceId(v string) *DeleteWorkspaceMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteWorkspaceMembersResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteWorkspaceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersResponseBody) SetRequestId(v string) *DeleteWorkspaceMembersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWorkspaceMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWorkspaceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWorkspaceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceMembersResponse) SetStatusCode(v int32) *DeleteWorkspaceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceMembersResponse) SetBody(v *DeleteWorkspaceMembersResponseBody) *DeleteWorkspaceMembersResponse {
	s.Body = v
	return s
}

type GetUserHeaders struct {
	CommonHeaders  map[string]*string            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetUserHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetAccountContext(v *GetUserHeadersAccountContext) *GetUserHeaders {
	s.AccountContext = v
	return s
}

type GetUserHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetUserHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetUserHeadersAccountContext) SetAccountId(v string) *GetUserHeadersAccountContext {
	s.AccountId = &v
	return s
}

type GetUserShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetUserShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetUserShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetUserShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserShrinkHeaders) SetAccountContextShrink(v string) *GetUserShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type GetUserRequest struct {
	TenantContext *GetUserRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	Language      *string                      `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetTenantContext(v *GetUserRequestTenantContext) *GetUserRequest {
	s.TenantContext = v
	return s
}

func (s *GetUserRequest) SetLanguage(v string) *GetUserRequest {
	s.Language = &v
	return s
}

type GetUserRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetUserRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetUserRequestTenantContext) SetTenantId(v string) *GetUserRequestTenantContext {
	s.TenantId = &v
	return s
}

type GetUserShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	Language            *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetUserShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserShrinkRequest) SetTenantContextShrink(v string) *GetUserShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetUserShrinkRequest) SetLanguage(v string) *GetUserShrinkRequest {
	s.Language = &v
	return s
}

type GetUserResponseBody struct {
	Active                   *bool                               `json:"active,omitempty" xml:"active,omitempty"`
	Admin                    *bool                               `json:"admin,omitempty" xml:"admin,omitempty"`
	Avatar                   *string                             `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Boss                     *bool                               `json:"boss,omitempty" xml:"boss,omitempty"`
	DeptIdList               []*int64                            `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	DeptOrderList            []*GetUserResponseBodyDeptOrderList `json:"deptOrderList,omitempty" xml:"deptOrderList,omitempty" type:"Repeated"`
	Email                    *string                             `json:"email,omitempty" xml:"email,omitempty"`
	ExclusiveAccount         *bool                               `json:"exclusiveAccount,omitempty" xml:"exclusiveAccount,omitempty"`
	ExclusiveAccountCorpId   *string                             `json:"exclusiveAccountCorpId,omitempty" xml:"exclusiveAccountCorpId,omitempty"`
	ExclusiveAccountCorpName *string                             `json:"exclusiveAccountCorpName,omitempty" xml:"exclusiveAccountCorpName,omitempty"`
	ExclusiveAccountType     *string                             `json:"exclusiveAccountType,omitempty" xml:"exclusiveAccountType,omitempty"`
	Extension                *string                             `json:"extension,omitempty" xml:"extension,omitempty"`
	HideMobile               *bool                               `json:"hideMobile,omitempty" xml:"hideMobile,omitempty"`
	HiredDate                *int64                              `json:"hiredDate,omitempty" xml:"hiredDate,omitempty"`
	JobNumber                *string                             `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	LeaderInDept             []*GetUserResponseBodyLeaderInDept  `json:"leaderInDept,omitempty" xml:"leaderInDept,omitempty" type:"Repeated"`
	LoginId                  *string                             `json:"loginId,omitempty" xml:"loginId,omitempty"`
	ManagerUserid            *string                             `json:"managerUserid,omitempty" xml:"managerUserid,omitempty"`
	Mobile                   *string                             `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name                     *string                             `json:"name,omitempty" xml:"name,omitempty"`
	Nickname                 *string                             `json:"nickname,omitempty" xml:"nickname,omitempty"`
	OrgEmail                 *string                             `json:"orgEmail,omitempty" xml:"orgEmail,omitempty"`
	RealAuthed               *bool                               `json:"realAuthed,omitempty" xml:"realAuthed,omitempty"`
	Remark                   *string                             `json:"remark,omitempty" xml:"remark,omitempty"`
	RequestId                *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RoleList                 []*GetUserResponseBodyRoleList      `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
	Senior                   *bool                               `json:"senior,omitempty" xml:"senior,omitempty"`
	StateCode                *string                             `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	Telephone                *string                             `json:"telephone,omitempty" xml:"telephone,omitempty"`
	Title                    *string                             `json:"title,omitempty" xml:"title,omitempty"`
	UnionEmpExt              *GetUserResponseBodyUnionEmpExt     `json:"unionEmpExt,omitempty" xml:"unionEmpExt,omitempty" type:"Struct"`
	Userid                   *string                             `json:"userid,omitempty" xml:"userid,omitempty"`
	WorkPlace                *string                             `json:"workPlace,omitempty" xml:"workPlace,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetActive(v bool) *GetUserResponseBody {
	s.Active = &v
	return s
}

func (s *GetUserResponseBody) SetAdmin(v bool) *GetUserResponseBody {
	s.Admin = &v
	return s
}

func (s *GetUserResponseBody) SetAvatar(v string) *GetUserResponseBody {
	s.Avatar = &v
	return s
}

func (s *GetUserResponseBody) SetBoss(v bool) *GetUserResponseBody {
	s.Boss = &v
	return s
}

func (s *GetUserResponseBody) SetDeptIdList(v []*int64) *GetUserResponseBody {
	s.DeptIdList = v
	return s
}

func (s *GetUserResponseBody) SetDeptOrderList(v []*GetUserResponseBodyDeptOrderList) *GetUserResponseBody {
	s.DeptOrderList = v
	return s
}

func (s *GetUserResponseBody) SetEmail(v string) *GetUserResponseBody {
	s.Email = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccount(v bool) *GetUserResponseBody {
	s.ExclusiveAccount = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccountCorpId(v string) *GetUserResponseBody {
	s.ExclusiveAccountCorpId = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccountCorpName(v string) *GetUserResponseBody {
	s.ExclusiveAccountCorpName = &v
	return s
}

func (s *GetUserResponseBody) SetExclusiveAccountType(v string) *GetUserResponseBody {
	s.ExclusiveAccountType = &v
	return s
}

func (s *GetUserResponseBody) SetExtension(v string) *GetUserResponseBody {
	s.Extension = &v
	return s
}

func (s *GetUserResponseBody) SetHideMobile(v bool) *GetUserResponseBody {
	s.HideMobile = &v
	return s
}

func (s *GetUserResponseBody) SetHiredDate(v int64) *GetUserResponseBody {
	s.HiredDate = &v
	return s
}

func (s *GetUserResponseBody) SetJobNumber(v string) *GetUserResponseBody {
	s.JobNumber = &v
	return s
}

func (s *GetUserResponseBody) SetLeaderInDept(v []*GetUserResponseBodyLeaderInDept) *GetUserResponseBody {
	s.LeaderInDept = v
	return s
}

func (s *GetUserResponseBody) SetLoginId(v string) *GetUserResponseBody {
	s.LoginId = &v
	return s
}

func (s *GetUserResponseBody) SetManagerUserid(v string) *GetUserResponseBody {
	s.ManagerUserid = &v
	return s
}

func (s *GetUserResponseBody) SetMobile(v string) *GetUserResponseBody {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBody) SetName(v string) *GetUserResponseBody {
	s.Name = &v
	return s
}

func (s *GetUserResponseBody) SetNickname(v string) *GetUserResponseBody {
	s.Nickname = &v
	return s
}

func (s *GetUserResponseBody) SetOrgEmail(v string) *GetUserResponseBody {
	s.OrgEmail = &v
	return s
}

func (s *GetUserResponseBody) SetRealAuthed(v bool) *GetUserResponseBody {
	s.RealAuthed = &v
	return s
}

func (s *GetUserResponseBody) SetRemark(v string) *GetUserResponseBody {
	s.Remark = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetRoleList(v []*GetUserResponseBodyRoleList) *GetUserResponseBody {
	s.RoleList = v
	return s
}

func (s *GetUserResponseBody) SetSenior(v bool) *GetUserResponseBody {
	s.Senior = &v
	return s
}

func (s *GetUserResponseBody) SetStateCode(v string) *GetUserResponseBody {
	s.StateCode = &v
	return s
}

func (s *GetUserResponseBody) SetTelephone(v string) *GetUserResponseBody {
	s.Telephone = &v
	return s
}

func (s *GetUserResponseBody) SetTitle(v string) *GetUserResponseBody {
	s.Title = &v
	return s
}

func (s *GetUserResponseBody) SetUnionEmpExt(v *GetUserResponseBodyUnionEmpExt) *GetUserResponseBody {
	s.UnionEmpExt = v
	return s
}

func (s *GetUserResponseBody) SetUserid(v string) *GetUserResponseBody {
	s.Userid = &v
	return s
}

func (s *GetUserResponseBody) SetWorkPlace(v string) *GetUserResponseBody {
	s.WorkPlace = &v
	return s
}

type GetUserResponseBodyDeptOrderList struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Order  *int64 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s GetUserResponseBodyDeptOrderList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyDeptOrderList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyDeptOrderList) SetDeptId(v int64) *GetUserResponseBodyDeptOrderList {
	s.DeptId = &v
	return s
}

func (s *GetUserResponseBodyDeptOrderList) SetOrder(v int64) *GetUserResponseBodyDeptOrderList {
	s.Order = &v
	return s
}

type GetUserResponseBodyLeaderInDept struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Leader *bool  `json:"leader,omitempty" xml:"leader,omitempty"`
}

func (s GetUserResponseBodyLeaderInDept) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyLeaderInDept) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyLeaderInDept) SetDeptId(v int64) *GetUserResponseBodyLeaderInDept {
	s.DeptId = &v
	return s
}

func (s *GetUserResponseBodyLeaderInDept) SetLeader(v bool) *GetUserResponseBodyLeaderInDept {
	s.Leader = &v
	return s
}

type GetUserResponseBodyRoleList struct {
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetUserResponseBodyRoleList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyRoleList) SetGroupName(v string) *GetUserResponseBodyRoleList {
	s.GroupName = &v
	return s
}

func (s *GetUserResponseBodyRoleList) SetId(v int64) *GetUserResponseBodyRoleList {
	s.Id = &v
	return s
}

func (s *GetUserResponseBodyRoleList) SetName(v string) *GetUserResponseBodyRoleList {
	s.Name = &v
	return s
}

type GetUserResponseBodyUnionEmpExt struct {
	CorpId          *string                                          `json:"corpId,omitempty" xml:"corpId,omitempty"`
	UnionEmpMapList []*GetUserResponseBodyUnionEmpExtUnionEmpMapList `json:"unionEmpMapList,omitempty" xml:"unionEmpMapList,omitempty" type:"Repeated"`
	Userid          *string                                          `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s GetUserResponseBodyUnionEmpExt) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUnionEmpExt) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUnionEmpExt) SetCorpId(v string) *GetUserResponseBodyUnionEmpExt {
	s.CorpId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) SetUnionEmpMapList(v []*GetUserResponseBodyUnionEmpExtUnionEmpMapList) *GetUserResponseBodyUnionEmpExt {
	s.UnionEmpMapList = v
	return s
}

func (s *GetUserResponseBodyUnionEmpExt) SetUserid(v string) *GetUserResponseBodyUnionEmpExt {
	s.Userid = &v
	return s
}

type GetUserResponseBodyUnionEmpExtUnionEmpMapList struct {
	CropId *string `json:"cropId,omitempty" xml:"cropId,omitempty"`
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s GetUserResponseBodyUnionEmpExtUnionEmpMapList) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUnionEmpExtUnionEmpMapList) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) SetCropId(v string) *GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	s.CropId = &v
	return s
}

func (s *GetUserResponseBodyUnionEmpExtUnionEmpMapList) SetUserid(v string) *GetUserResponseBodyUnionEmpExtUnionEmpMapList {
	s.Userid = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type InsertRowsBeforeHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InsertRowsBeforeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InsertRowsBeforeHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeHeaders) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeHeaders) SetCommonHeaders(v map[string]*string) *InsertRowsBeforeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertRowsBeforeHeaders) SetAccountContext(v *InsertRowsBeforeHeadersAccountContext) *InsertRowsBeforeHeaders {
	s.AccountContext = v
	return s
}

type InsertRowsBeforeHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InsertRowsBeforeHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeHeadersAccountContext) SetAccountId(v string) *InsertRowsBeforeHeadersAccountContext {
	s.AccountId = &v
	return s
}

type InsertRowsBeforeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InsertRowsBeforeShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeShrinkHeaders) SetCommonHeaders(v map[string]*string) *InsertRowsBeforeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertRowsBeforeShrinkHeaders) SetAccountContextShrink(v string) *InsertRowsBeforeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type InsertRowsBeforeRequest struct {
	Row           *int64                                `json:"Row,omitempty" xml:"Row,omitempty"`
	RowCount      *int64                                `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	SheetId       *string                               `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContext *InsertRowsBeforeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkbookId    *string                               `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s InsertRowsBeforeRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeRequest) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeRequest) SetRow(v int64) *InsertRowsBeforeRequest {
	s.Row = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetRowCount(v int64) *InsertRowsBeforeRequest {
	s.RowCount = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetSheetId(v string) *InsertRowsBeforeRequest {
	s.SheetId = &v
	return s
}

func (s *InsertRowsBeforeRequest) SetTenantContext(v *InsertRowsBeforeRequestTenantContext) *InsertRowsBeforeRequest {
	s.TenantContext = v
	return s
}

func (s *InsertRowsBeforeRequest) SetWorkbookId(v string) *InsertRowsBeforeRequest {
	s.WorkbookId = &v
	return s
}

type InsertRowsBeforeRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s InsertRowsBeforeRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeRequestTenantContext) SetTenantId(v string) *InsertRowsBeforeRequestTenantContext {
	s.TenantId = &v
	return s
}

type InsertRowsBeforeShrinkRequest struct {
	Row                 *int64  `json:"Row,omitempty" xml:"Row,omitempty"`
	RowCount            *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	SheetId             *string `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkbookId          *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s InsertRowsBeforeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeShrinkRequest) SetRow(v int64) *InsertRowsBeforeShrinkRequest {
	s.Row = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetRowCount(v int64) *InsertRowsBeforeShrinkRequest {
	s.RowCount = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetSheetId(v string) *InsertRowsBeforeShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetTenantContextShrink(v string) *InsertRowsBeforeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetWorkbookId(v string) *InsertRowsBeforeShrinkRequest {
	s.WorkbookId = &v
	return s
}

type InsertRowsBeforeResponseBody struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InsertRowsBeforeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeResponseBody) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeResponseBody) SetId(v string) *InsertRowsBeforeResponseBody {
	s.Id = &v
	return s
}

func (s *InsertRowsBeforeResponseBody) SetRequestId(v string) *InsertRowsBeforeResponseBody {
	s.RequestId = &v
	return s
}

type InsertRowsBeforeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InsertRowsBeforeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertRowsBeforeResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertRowsBeforeResponse) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeResponse) SetHeaders(v map[string]*string) *InsertRowsBeforeResponse {
	s.Headers = v
	return s
}

func (s *InsertRowsBeforeResponse) SetStatusCode(v int32) *InsertRowsBeforeResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertRowsBeforeResponse) SetBody(v *InsertRowsBeforeResponseBody) *InsertRowsBeforeResponse {
	s.Body = v
	return s
}

type QueryOrgTodoTasksHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryOrgTodoTasksHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryOrgTodoTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgTodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgTodoTasksHeaders) SetAccountContext(v *QueryOrgTodoTasksHeadersAccountContext) *QueryOrgTodoTasksHeaders {
	s.AccountContext = v
	return s
}

type QueryOrgTodoTasksHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryOrgTodoTasksHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksHeadersAccountContext) SetAccountId(v string) *QueryOrgTodoTasksHeadersAccountContext {
	s.AccountId = &v
	return s
}

type QueryOrgTodoTasksShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryOrgTodoTasksShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgTodoTasksShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgTodoTasksShrinkHeaders) SetAccountContextShrink(v string) *QueryOrgTodoTasksShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type QueryOrgTodoTasksRequest struct {
	TenantContext *QueryOrgTodoTasksRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	IsDone        *bool                                  `json:"isDone,omitempty" xml:"isDone,omitempty"`
	NextToken     *string                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryOrgTodoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksRequest) SetTenantContext(v *QueryOrgTodoTasksRequestTenantContext) *QueryOrgTodoTasksRequest {
	s.TenantContext = v
	return s
}

func (s *QueryOrgTodoTasksRequest) SetIsDone(v bool) *QueryOrgTodoTasksRequest {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoTasksRequest) SetNextToken(v string) *QueryOrgTodoTasksRequest {
	s.NextToken = &v
	return s
}

type QueryOrgTodoTasksRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryOrgTodoTasksRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksRequestTenantContext) SetTenantId(v string) *QueryOrgTodoTasksRequestTenantContext {
	s.TenantId = &v
	return s
}

type QueryOrgTodoTasksShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	IsDone              *bool   `json:"isDone,omitempty" xml:"isDone,omitempty"`
	NextToken           *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryOrgTodoTasksShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksShrinkRequest) SetTenantContextShrink(v string) *QueryOrgTodoTasksShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryOrgTodoTasksShrinkRequest) SetIsDone(v bool) *QueryOrgTodoTasksShrinkRequest {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoTasksShrinkRequest) SetNextToken(v string) *QueryOrgTodoTasksShrinkRequest {
	s.NextToken = &v
	return s
}

type QueryOrgTodoTasksResponseBody struct {
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TodoCards []*QueryOrgTodoTasksResponseBodyTodoCards `json:"todoCards,omitempty" xml:"todoCards,omitempty" type:"Repeated"`
}

func (s QueryOrgTodoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBody) SetNextToken(v string) *QueryOrgTodoTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBody) SetRequestId(v string) *QueryOrgTodoTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBody) SetTodoCards(v []*QueryOrgTodoTasksResponseBodyTodoCards) *QueryOrgTodoTasksResponseBody {
	s.TodoCards = v
	return s
}

type QueryOrgTodoTasksResponseBodyTodoCards struct {
	BizTag       *string                                          `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	CreatedTime  *int64                                           `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CreatorId    *string                                          `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DetailUrl    *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	DueTime      *int64                                           `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	IsDone       *bool                                            `json:"isDone,omitempty" xml:"isDone,omitempty"`
	ModifiedTime *int64                                           `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Priority     *int32                                           `json:"priority,omitempty" xml:"priority,omitempty"`
	SourceId     *string                                          `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	Subject      *string                                          `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId       *string                                          `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryOrgTodoTasksResponseBodyTodoCards) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBodyTodoCards) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetBizTag(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.BizTag = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetCreatedTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.CreatedTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetCreatorId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.CreatorId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetDetailUrl(v *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.DetailUrl = v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetDueTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.DueTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetIsDone(v bool) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetModifiedTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.ModifiedTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetPriority(v int32) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.Priority = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetSourceId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.SourceId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetSubject(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.Subject = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetTaskId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.TaskId = &v
	return s
}

type QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) SetAppUrl(v string) *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) SetPcUrl(v string) *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl {
	s.PcUrl = &v
	return s
}

type QueryOrgTodoTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOrgTodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgTodoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgTodoTasksResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponse) SetHeaders(v map[string]*string) *QueryOrgTodoTasksResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgTodoTasksResponse) SetStatusCode(v int32) *QueryOrgTodoTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrgTodoTasksResponse) SetBody(v *QueryOrgTodoTasksResponseBody) *QueryOrgTodoTasksResponse {
	s.Body = v
	return s
}

type UpdateTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskHeaders) SetAccountContext(v *UpdateTodoTaskHeadersAccountContext) *UpdateTodoTaskHeaders {
	s.AccountContext = v
	return s
}

type UpdateTodoTaskHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateTodoTaskHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskHeadersAccountContext) SetAccountId(v string) *UpdateTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

type UpdateTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateTodoTaskShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *UpdateTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type UpdateTodoTaskRequest struct {
	TenantContext  *UpdateTodoTaskRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	Description    *string                             `json:"description,omitempty" xml:"description,omitempty"`
	Done           *bool                               `json:"done,omitempty" xml:"done,omitempty"`
	DueTime        *int64                              `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds    []*string                           `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	ParticipantIds []*string                           `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	Subject        *string                             `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId         *string                             `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequest) SetTenantContext(v *UpdateTodoTaskRequestTenantContext) *UpdateTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateTodoTaskRequest) SetDescription(v string) *UpdateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDone(v bool) *UpdateTodoTaskRequest {
	s.Done = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDueTime(v int64) *UpdateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetExecutorIds(v []*string) *UpdateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *UpdateTodoTaskRequest) SetParticipantIds(v []*string) *UpdateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *UpdateTodoTaskRequest) SetSubject(v string) *UpdateTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetTaskId(v string) *UpdateTodoTaskRequest {
	s.TaskId = &v
	return s
}

type UpdateTodoTaskRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateTodoTaskRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequestTenantContext) SetTenantId(v string) *UpdateTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

type UpdateTodoTaskShrinkRequest struct {
	TenantContextShrink  *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	Done                 *bool   `json:"done,omitempty" xml:"done,omitempty"`
	DueTime              *int64  `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIdsShrink    *string `json:"executorIds,omitempty" xml:"executorIds,omitempty"`
	ParticipantIdsShrink *string `json:"participantIds,omitempty" xml:"participantIds,omitempty"`
	Subject              *string `json:"subject,omitempty" xml:"subject,omitempty"`
	TaskId               *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTodoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskShrinkRequest) SetTenantContextShrink(v string) *UpdateTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetDescription(v string) *UpdateTodoTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetDone(v bool) *UpdateTodoTaskShrinkRequest {
	s.Done = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetDueTime(v int64) *UpdateTodoTaskShrinkRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetExecutorIdsShrink(v string) *UpdateTodoTaskShrinkRequest {
	s.ExecutorIdsShrink = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetParticipantIdsShrink(v string) *UpdateTodoTaskShrinkRequest {
	s.ParticipantIdsShrink = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetSubject(v string) *UpdateTodoTaskShrinkRequest {
	s.Subject = &v
	return s
}

func (s *UpdateTodoTaskShrinkRequest) SetTaskId(v string) *UpdateTodoTaskShrinkRequest {
	s.TaskId = &v
	return s
}

type UpdateTodoTaskResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskResponseBody) SetRequestId(v string) *UpdateTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTodoTaskResponseBody) SetResult(v bool) *UpdateTodoTaskResponseBody {
	s.Result = &v
	return s
}

type UpdateTodoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskResponse) SetHeaders(v map[string]*string) *UpdateTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTaskResponse) SetStatusCode(v int32) *UpdateTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTodoTaskResponse) SetBody(v *UpdateTodoTaskResponseBody) *UpdateTodoTaskResponse {
	s.Body = v
	return s
}

type UpdateTodoTaskExecutorStatusHeaders struct {
	CommonHeaders  map[string]*string                                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateTodoTaskExecutorStatusHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateTodoTaskExecutorStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusHeaders) SetAccountContext(v *UpdateTodoTaskExecutorStatusHeadersAccountContext) *UpdateTodoTaskExecutorStatusHeaders {
	s.AccountContext = v
	return s
}

type UpdateTodoTaskExecutorStatusHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusHeadersAccountContext) SetAccountId(v string) *UpdateTodoTaskExecutorStatusHeadersAccountContext {
	s.AccountId = &v
	return s
}

type UpdateTodoTaskExecutorStatusShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkHeaders) SetAccountContextShrink(v string) *UpdateTodoTaskExecutorStatusShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type UpdateTodoTaskExecutorStatusRequest struct {
	TenantContext      *UpdateTodoTaskExecutorStatusRequestTenantContext        `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	ExecutorStatusList []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList `json:"executorStatusList,omitempty" xml:"executorStatusList,omitempty" type:"Repeated"`
	OperatorId         *string                                                  `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TaskId             *string                                                  `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetTenantContext(v *UpdateTodoTaskExecutorStatusRequestTenantContext) *UpdateTodoTaskExecutorStatusRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetExecutorStatusList(v []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList) *UpdateTodoTaskExecutorStatusRequest {
	s.ExecutorStatusList = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetOperatorId(v string) *UpdateTodoTaskExecutorStatusRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetTaskId(v string) *UpdateTodoTaskExecutorStatusRequest {
	s.TaskId = &v
	return s
}

type UpdateTodoTaskExecutorStatusRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequestTenantContext) SetTenantId(v string) *UpdateTodoTaskExecutorStatusRequestTenantContext {
	s.TenantId = &v
	return s
}

type UpdateTodoTaskExecutorStatusRequestExecutorStatusList struct {
	Id     *string `json:"id,omitempty" xml:"id,omitempty"`
	IsDone *bool   `json:"isDone,omitempty" xml:"isDone,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequestExecutorStatusList) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequestExecutorStatusList) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) SetId(v string) *UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	s.Id = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) SetIsDone(v bool) *UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	s.IsDone = &v
	return s
}

type UpdateTodoTaskExecutorStatusShrinkRequest struct {
	TenantContextShrink      *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	ExecutorStatusListShrink *string `json:"executorStatusList,omitempty" xml:"executorStatusList,omitempty"`
	OperatorId               *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	TaskId                   *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) SetTenantContextShrink(v string) *UpdateTodoTaskExecutorStatusShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) SetExecutorStatusListShrink(v string) *UpdateTodoTaskExecutorStatusShrinkRequest {
	s.ExecutorStatusListShrink = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) SetOperatorId(v string) *UpdateTodoTaskExecutorStatusShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkRequest) SetTaskId(v string) *UpdateTodoTaskExecutorStatusShrinkRequest {
	s.TaskId = &v
	return s
}

type UpdateTodoTaskExecutorStatusResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) SetRequestId(v string) *UpdateTodoTaskExecutorStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) SetResult(v bool) *UpdateTodoTaskExecutorStatusResponseBody {
	s.Result = &v
	return s
}

type UpdateTodoTaskExecutorStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTodoTaskExecutorStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTodoTaskExecutorStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetStatusCode(v int32) *UpdateTodoTaskExecutorStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponse) SetBody(v *UpdateTodoTaskExecutorStatusResponseBody) *UpdateTodoTaskExecutorStatusResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceDocMembersHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateWorkspaceDocMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceDocMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceDocMembersHeaders) SetAccountContext(v *UpdateWorkspaceDocMembersHeadersAccountContext) *UpdateWorkspaceDocMembersHeaders {
	s.AccountContext = v
	return s
}

type UpdateWorkspaceDocMembersHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateWorkspaceDocMembersHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersHeadersAccountContext) SetAccountId(v string) *UpdateWorkspaceDocMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

type UpdateWorkspaceDocMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateWorkspaceDocMembersShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceDocMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkHeaders) SetAccountContextShrink(v string) *UpdateWorkspaceDocMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type UpdateWorkspaceDocMembersRequest struct {
	Members       []*UpdateWorkspaceDocMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	NodeId        *string                                        `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContext *UpdateWorkspaceDocMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkspaceId   *string                                        `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceDocMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequest) SetMembers(v []*UpdateWorkspaceDocMembersRequestMembers) *UpdateWorkspaceDocMembersRequest {
	s.Members = v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) SetNodeId(v string) *UpdateWorkspaceDocMembersRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) SetTenantContext(v *UpdateWorkspaceDocMembersRequestTenantContext) *UpdateWorkspaceDocMembersRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateWorkspaceDocMembersRequest) SetWorkspaceId(v string) *UpdateWorkspaceDocMembersRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateWorkspaceDocMembersRequestMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	RoleType   *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s UpdateWorkspaceDocMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetMemberId(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetMemberType(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *UpdateWorkspaceDocMembersRequestMembers) SetRoleType(v string) *UpdateWorkspaceDocMembersRequestMembers {
	s.RoleType = &v
	return s
}

type UpdateWorkspaceDocMembersRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateWorkspaceDocMembersRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersRequestTenantContext) SetTenantId(v string) *UpdateWorkspaceDocMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

type UpdateWorkspaceDocMembersShrinkRequest struct {
	MembersShrink       *string `json:"Members,omitempty" xml:"Members,omitempty"`
	NodeId              *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceDocMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) SetMembersShrink(v string) *UpdateWorkspaceDocMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) SetNodeId(v string) *UpdateWorkspaceDocMembersShrinkRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) SetTenantContextShrink(v string) *UpdateWorkspaceDocMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkRequest) SetWorkspaceId(v string) *UpdateWorkspaceDocMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateWorkspaceDocMembersResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateWorkspaceDocMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersResponseBody) SetRequestId(v string) *UpdateWorkspaceDocMembersResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWorkspaceDocMembersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWorkspaceDocMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkspaceDocMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceDocMembersResponse) SetStatusCode(v int32) *UpdateWorkspaceDocMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceDocMembersResponse) SetBody(v *UpdateWorkspaceDocMembersResponseBody) *UpdateWorkspaceDocMembersResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceMembersHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateWorkspaceMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceMembersHeaders) SetAccountContext(v *UpdateWorkspaceMembersHeadersAccountContext) *UpdateWorkspaceMembersHeaders {
	s.AccountContext = v
	return s
}

type UpdateWorkspaceMembersHeadersAccountContext struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateWorkspaceMembersHeadersAccountContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersHeadersAccountContext) SetAccountId(v string) *UpdateWorkspaceMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

type UpdateWorkspaceMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateWorkspaceMembersShrinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceMembersShrinkHeaders) SetAccountContextShrink(v string) *UpdateWorkspaceMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

type UpdateWorkspaceMembersRequest struct {
	Members       []*UpdateWorkspaceMembersRequestMembers     `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	TenantContext *UpdateWorkspaceMembersRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WorkspaceId   *string                                     `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequest) SetMembers(v []*UpdateWorkspaceMembersRequestMembers) *UpdateWorkspaceMembersRequest {
	s.Members = v
	return s
}

func (s *UpdateWorkspaceMembersRequest) SetTenantContext(v *UpdateWorkspaceMembersRequestTenantContext) *UpdateWorkspaceMembersRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateWorkspaceMembersRequest) SetWorkspaceId(v string) *UpdateWorkspaceMembersRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateWorkspaceMembersRequestMembers struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	RoleType   *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s UpdateWorkspaceMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequestMembers) SetMemberId(v string) *UpdateWorkspaceMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *UpdateWorkspaceMembersRequestMembers) SetMemberType(v string) *UpdateWorkspaceMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *UpdateWorkspaceMembersRequestMembers) SetRoleType(v string) *UpdateWorkspaceMembersRequestMembers {
	s.RoleType = &v
	return s
}

type UpdateWorkspaceMembersRequestTenantContext struct {
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateWorkspaceMembersRequestTenantContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersRequestTenantContext) SetTenantId(v string) *UpdateWorkspaceMembersRequestTenantContext {
	s.TenantId = &v
	return s
}

type UpdateWorkspaceMembersShrinkRequest struct {
	MembersShrink       *string `json:"Members,omitempty" xml:"Members,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersShrinkRequest) SetMembersShrink(v string) *UpdateWorkspaceMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *UpdateWorkspaceMembersShrinkRequest) SetTenantContextShrink(v string) *UpdateWorkspaceMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateWorkspaceMembersShrinkRequest) SetWorkspaceId(v string) *UpdateWorkspaceMembersShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateWorkspaceMembersResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateWorkspaceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersResponseBody) SetRequestId(v string) *UpdateWorkspaceMembersResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWorkspaceMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWorkspaceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkspaceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceMembersResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceMembersResponse) SetStatusCode(v int32) *UpdateWorkspaceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceMembersResponse) SetBody(v *UpdateWorkspaceMembersResponseBody) *UpdateWorkspaceMembersResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aliding"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddWorkspaceDocMembersWithOptions(tmpReq *AddWorkspaceDocMembersRequest, tmpHeader *AddWorkspaceDocMembersHeaders, runtime *util.RuntimeOptions) (_result *AddWorkspaceDocMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddWorkspaceDocMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddWorkspaceDocMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Members)) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, tea.String("Members"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MembersShrink)) {
		body["Members"] = request.MembersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddWorkspaceDocMembers"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/addWorkspaceDocMembers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWorkspaceDocMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWorkspaceDocMembers(request *AddWorkspaceDocMembersRequest) (_result *AddWorkspaceDocMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddWorkspaceDocMembersHeaders{}
	_result = &AddWorkspaceDocMembersResponse{}
	_body, _err := client.AddWorkspaceDocMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWorkspaceMembersWithOptions(tmpReq *AddWorkspaceMembersRequest, tmpHeader *AddWorkspaceMembersHeaders, runtime *util.RuntimeOptions) (_result *AddWorkspaceMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddWorkspaceMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &AddWorkspaceMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Members)) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, tea.String("Members"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MembersShrink)) {
		body["Members"] = request.MembersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddWorkspaceMembers"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/addWorkspaceMembers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWorkspaceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWorkspaceMembers(request *AddWorkspaceMembersRequest) (_result *AddWorkspaceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddWorkspaceMembersHeaders{}
	_result = &AddWorkspaceMembersResponse{}
	_body, _err := client.AddWorkspaceMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSheetWithOptions(tmpReq *CreateSheetRequest, tmpHeader *CreateSheetHeaders, runtime *util.RuntimeOptions) (_result *CreateSheetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSheetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateSheetShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkbookId)) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSheet"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/createSheet"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSheetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSheet(request *CreateSheetRequest) (_result *CreateSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSheetHeaders{}
	_result = &CreateSheetResponse{}
	_body, _err := client.CreateSheetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTodoTaskWithOptions(tmpReq *CreateTodoTaskRequest, tmpHeader *CreateTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateTodoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ContentFieldList)) {
		request.ContentFieldListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContentFieldList, tea.String("contentFieldList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DetailUrl)) {
		request.DetailUrlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetailUrl, tea.String("detailUrl"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ExecutorIds)) {
		request.ExecutorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorIds, tea.String("executorIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NotifyConfigs)) {
		request.NotifyConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyConfigs, tea.String("notifyConfigs"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ParticipantIds)) {
		request.ParticipantIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParticipantIds, tea.String("participantIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["operatorId"] = request.OperatorId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ContentFieldListShrink)) {
		body["contentFieldList"] = request.ContentFieldListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["creatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DetailUrlShrink)) {
		body["detailUrl"] = request.DetailUrlShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIdsShrink)) {
		body["executorIds"] = request.ExecutorIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IsOnlyShowExecutor)) {
		body["isOnlyShowExecutor"] = request.IsOnlyShowExecutor
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyConfigsShrink)) {
		body["notifyConfigs"] = request.NotifyConfigsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParticipantIdsShrink)) {
		body["participantIds"] = request.ParticipantIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["sourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTodoTask"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/task/createTodoTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTodoTask(request *CreateTodoTaskRequest) (_result *CreateTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTodoTaskHeaders{}
	_result = &CreateTodoTaskResponse{}
	_body, _err := client.CreateTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceWithOptions(tmpReq *CreateWorkspaceRequest, tmpHeader *CreateWorkspaceHeaders, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateWorkspaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateWorkspaceShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkspace"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/createWorkspace"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateWorkspaceHeaders{}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceDocWithOptions(tmpReq *CreateWorkspaceDocRequest, tmpHeader *CreateWorkspaceDocHeaders, runtime *util.RuntimeOptions) (_result *CreateWorkspaceDocResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateWorkspaceDocShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &CreateWorkspaceDocShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentNodeId)) {
		body["ParentNodeId"] = request.ParentNodeId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkspaceDoc"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/createWorkspaceDoc"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkspaceDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspaceDoc(request *CreateWorkspaceDocRequest) (_result *CreateWorkspaceDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateWorkspaceDocHeaders{}
	_result = &CreateWorkspaceDocResponse{}
	_body, _err := client.CreateWorkspaceDocWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTodoTaskWithOptions(tmpReq *DeleteTodoTaskRequest, tmpHeader *DeleteTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *DeleteTodoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTodoTask"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/task/deleteTodoTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTodoTask(request *DeleteTodoTaskRequest) (_result *DeleteTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTodoTaskHeaders{}
	_result = &DeleteTodoTaskResponse{}
	_body, _err := client.DeleteTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceDocMembersWithOptions(tmpReq *DeleteWorkspaceDocMembersRequest, tmpHeader *DeleteWorkspaceDocMembersHeaders, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceDocMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteWorkspaceDocMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteWorkspaceDocMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Members)) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, tea.String("Members"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MembersShrink)) {
		body["Members"] = request.MembersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkspaceDocMembers"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/deleteWorkspaceDocMembers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkspaceDocMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspaceDocMembers(request *DeleteWorkspaceDocMembersRequest) (_result *DeleteWorkspaceDocMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteWorkspaceDocMembersHeaders{}
	_result = &DeleteWorkspaceDocMembersResponse{}
	_body, _err := client.DeleteWorkspaceDocMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceMembersWithOptions(tmpReq *DeleteWorkspaceMembersRequest, tmpHeader *DeleteWorkspaceMembersHeaders, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteWorkspaceMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &DeleteWorkspaceMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Members)) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, tea.String("Members"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MembersShrink)) {
		body["Members"] = request.MembersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkspaceMembers"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/deleteWorkspaceMembers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkspaceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspaceMembers(request *DeleteWorkspaceMembersRequest) (_result *DeleteWorkspaceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteWorkspaceMembersHeaders{}
	_result = &DeleteWorkspaceMembersResponse{}
	_body, _err := client.DeleteWorkspaceMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(tmpReq *GetUserRequest, tmpHeader *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &GetUserShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/im/getUser"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertRowsBeforeWithOptions(tmpReq *InsertRowsBeforeRequest, tmpHeader *InsertRowsBeforeHeaders, runtime *util.RuntimeOptions) (_result *InsertRowsBeforeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InsertRowsBeforeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &InsertRowsBeforeShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Row)) {
		body["Row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.RowCount)) {
		body["RowCount"] = request.RowCount
	}

	if !tea.BoolValue(util.IsUnset(request.SheetId)) {
		body["SheetId"] = request.SheetId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkbookId)) {
		body["WorkbookId"] = request.WorkbookId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertRowsBefore"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/insertRowsBefore"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertRowsBeforeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertRowsBefore(request *InsertRowsBeforeRequest) (_result *InsertRowsBeforeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertRowsBeforeHeaders{}
	_result = &InsertRowsBeforeResponse{}
	_body, _err := client.InsertRowsBeforeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgTodoTasksWithOptions(tmpReq *QueryOrgTodoTasksRequest, tmpHeader *QueryOrgTodoTasksHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgTodoTasksResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryOrgTodoTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &QueryOrgTodoTasksShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["isDone"] = request.IsDone
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrgTodoTasks"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/task/queryOrgTodoTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrgTodoTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgTodoTasks(request *QueryOrgTodoTasksRequest) (_result *QueryOrgTodoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgTodoTasksHeaders{}
	_result = &QueryOrgTodoTasksResponse{}
	_body, _err := client.QueryOrgTodoTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTodoTaskWithOptions(tmpReq *UpdateTodoTaskRequest, tmpHeader *UpdateTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTodoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateTodoTaskShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ExecutorIds)) {
		request.ExecutorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorIds, tea.String("executorIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ParticipantIds)) {
		request.ParticipantIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParticipantIds, tea.String("participantIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Done)) {
		body["done"] = request.Done
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorIdsShrink)) {
		body["executorIds"] = request.ExecutorIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParticipantIdsShrink)) {
		body["participantIds"] = request.ParticipantIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		body["subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTodoTask"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/task/updateTodoTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTodoTask(request *UpdateTodoTaskRequest) (_result *UpdateTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTaskHeaders{}
	_result = &UpdateTodoTaskResponse{}
	_body, _err := client.UpdateTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTodoTaskExecutorStatusWithOptions(tmpReq *UpdateTodoTaskExecutorStatusRequest, tmpHeader *UpdateTodoTaskExecutorStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTodoTaskExecutorStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateTodoTaskExecutorStatusShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ExecutorStatusList)) {
		request.ExecutorStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorStatusList, tea.String("executorStatusList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorStatusListShrink)) {
		body["executorStatusList"] = request.ExecutorStatusListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTodoTaskExecutorStatus"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/task/updateTodoTaskExecutorStatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTodoTaskExecutorStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTodoTaskExecutorStatus(request *UpdateTodoTaskExecutorStatusRequest) (_result *UpdateTodoTaskExecutorStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTodoTaskExecutorStatusHeaders{}
	_result = &UpdateTodoTaskExecutorStatusResponse{}
	_body, _err := client.UpdateTodoTaskExecutorStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceDocMembersWithOptions(tmpReq *UpdateWorkspaceDocMembersRequest, tmpHeader *UpdateWorkspaceDocMembersHeaders, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceDocMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateWorkspaceDocMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateWorkspaceDocMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Members)) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, tea.String("Members"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MembersShrink)) {
		body["Members"] = request.MembersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspaceDocMembers"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/updateWorkspaceDocMembers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkspaceDocMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspaceDocMembers(request *UpdateWorkspaceDocMembersRequest) (_result *UpdateWorkspaceDocMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateWorkspaceDocMembersHeaders{}
	_result = &UpdateWorkspaceDocMembersResponse{}
	_body, _err := client.UpdateWorkspaceDocMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceMembersWithOptions(tmpReq *UpdateWorkspaceMembersRequest, tmpHeader *UpdateWorkspaceMembersHeaders, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateWorkspaceMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	headers := &UpdateWorkspaceMembersShrinkHeaders{}
	openapiutil.Convert(tmpHeader, headers)
	if !tea.BoolValue(util.IsUnset(tmpHeader.AccountContext)) {
		headers.AccountContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpHeader.AccountContext, tea.String("AccountContext"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Members)) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, tea.String("Members"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantContext)) {
		request.TenantContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantContext, tea.String("TenantContext"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MembersShrink)) {
		body["Members"] = request.MembersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantContextShrink)) {
		body["TenantContext"] = request.TenantContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccountContextShrink)) {
		realHeaders["AccountContext"] = util.ToJSONString(headers.AccountContextShrink)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspaceMembers"),
		Version:     tea.String("2023-04-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dingtalk/v1/documents/updateWorkspaceMembers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkspaceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspaceMembers(request *UpdateWorkspaceMembersRequest) (_result *UpdateWorkspaceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateWorkspaceMembersHeaders{}
	_result = &UpdateWorkspaceMembersResponse{}
	_body, _err := client.UpdateWorkspaceMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
