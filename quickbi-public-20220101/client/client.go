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

type AddDataLevelPermissionRuleUsersRequest struct {
	AddUserModel *string `json:"AddUserModel,omitempty" xml:"AddUserModel,omitempty"`
}

func (s AddDataLevelPermissionRuleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDataLevelPermissionRuleUsersRequest) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionRuleUsersRequest) SetAddUserModel(v string) *AddDataLevelPermissionRuleUsersRequest {
	s.AddUserModel = &v
	return s
}

type AddDataLevelPermissionRuleUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDataLevelPermissionRuleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDataLevelPermissionRuleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) SetRequestId(v string) *AddDataLevelPermissionRuleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) SetResult(v bool) *AddDataLevelPermissionRuleUsersResponseBody {
	s.Result = &v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponseBody) SetSuccess(v bool) *AddDataLevelPermissionRuleUsersResponseBody {
	s.Success = &v
	return s
}

type AddDataLevelPermissionRuleUsersResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddDataLevelPermissionRuleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDataLevelPermissionRuleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDataLevelPermissionRuleUsersResponse) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionRuleUsersResponse) SetHeaders(v map[string]*string) *AddDataLevelPermissionRuleUsersResponse {
	s.Headers = v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponse) SetStatusCode(v int32) *AddDataLevelPermissionRuleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataLevelPermissionRuleUsersResponse) SetBody(v *AddDataLevelPermissionRuleUsersResponseBody) *AddDataLevelPermissionRuleUsersResponse {
	s.Body = v
	return s
}

type AddDataLevelPermissionWhiteListRequest struct {
	CubeId      *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	TargetIds   *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
	TargetType  *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s AddDataLevelPermissionWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDataLevelPermissionWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionWhiteListRequest) SetCubeId(v string) *AddDataLevelPermissionWhiteListRequest {
	s.CubeId = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) SetOperateType(v string) *AddDataLevelPermissionWhiteListRequest {
	s.OperateType = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) SetRuleType(v string) *AddDataLevelPermissionWhiteListRequest {
	s.RuleType = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) SetTargetIds(v string) *AddDataLevelPermissionWhiteListRequest {
	s.TargetIds = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListRequest) SetTargetType(v string) *AddDataLevelPermissionWhiteListRequest {
	s.TargetType = &v
	return s
}

type AddDataLevelPermissionWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDataLevelPermissionWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDataLevelPermissionWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionWhiteListResponseBody) SetRequestId(v string) *AddDataLevelPermissionWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponseBody) SetResult(v bool) *AddDataLevelPermissionWhiteListResponseBody {
	s.Result = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponseBody) SetSuccess(v bool) *AddDataLevelPermissionWhiteListResponseBody {
	s.Success = &v
	return s
}

type AddDataLevelPermissionWhiteListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDataLevelPermissionWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDataLevelPermissionWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddDataLevelPermissionWhiteListResponse) SetHeaders(v map[string]*string) *AddDataLevelPermissionWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponse) SetStatusCode(v int32) *AddDataLevelPermissionWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataLevelPermissionWhiteListResponse) SetBody(v *AddDataLevelPermissionWhiteListResponseBody) *AddDataLevelPermissionWhiteListResponse {
	s.Body = v
	return s
}

type AddShareReportRequest struct {
	AuthPoint   *int32  `json:"AuthPoint,omitempty" xml:"AuthPoint,omitempty"`
	ExpireDate  *int64  `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	ShareToId   *string `json:"ShareToId,omitempty" xml:"ShareToId,omitempty"`
	ShareToType *int32  `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
	WorksId     *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s AddShareReportRequest) String() string {
	return tea.Prettify(s)
}

func (s AddShareReportRequest) GoString() string {
	return s.String()
}

func (s *AddShareReportRequest) SetAuthPoint(v int32) *AddShareReportRequest {
	s.AuthPoint = &v
	return s
}

func (s *AddShareReportRequest) SetExpireDate(v int64) *AddShareReportRequest {
	s.ExpireDate = &v
	return s
}

func (s *AddShareReportRequest) SetShareToId(v string) *AddShareReportRequest {
	s.ShareToId = &v
	return s
}

func (s *AddShareReportRequest) SetShareToType(v int32) *AddShareReportRequest {
	s.ShareToType = &v
	return s
}

func (s *AddShareReportRequest) SetWorksId(v string) *AddShareReportRequest {
	s.WorksId = &v
	return s
}

type AddShareReportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddShareReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddShareReportResponseBody) GoString() string {
	return s.String()
}

func (s *AddShareReportResponseBody) SetRequestId(v string) *AddShareReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddShareReportResponseBody) SetResult(v bool) *AddShareReportResponseBody {
	s.Result = &v
	return s
}

func (s *AddShareReportResponseBody) SetSuccess(v bool) *AddShareReportResponseBody {
	s.Success = &v
	return s
}

type AddShareReportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddShareReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddShareReportResponse) String() string {
	return tea.Prettify(s)
}

func (s AddShareReportResponse) GoString() string {
	return s.String()
}

func (s *AddShareReportResponse) SetHeaders(v map[string]*string) *AddShareReportResponse {
	s.Headers = v
	return s
}

func (s *AddShareReportResponse) SetStatusCode(v int32) *AddShareReportResponse {
	s.StatusCode = &v
	return s
}

func (s *AddShareReportResponse) SetBody(v *AddShareReportResponseBody) *AddShareReportResponse {
	s.Body = v
	return s
}

type AddUserRequest struct {
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AdminUser     *bool   `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	AuthAdminUser *bool   `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	NickName      *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	UserType      *int32  `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s AddUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserRequest) GoString() string {
	return s.String()
}

func (s *AddUserRequest) SetAccountName(v string) *AddUserRequest {
	s.AccountName = &v
	return s
}

func (s *AddUserRequest) SetAdminUser(v bool) *AddUserRequest {
	s.AdminUser = &v
	return s
}

func (s *AddUserRequest) SetAuthAdminUser(v bool) *AddUserRequest {
	s.AuthAdminUser = &v
	return s
}

func (s *AddUserRequest) SetNickName(v string) *AddUserRequest {
	s.NickName = &v
	return s
}

func (s *AddUserRequest) SetUserType(v int32) *AddUserRequest {
	s.UserType = &v
	return s
}

type AddUserResponseBody struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserResponseBody) SetRequestId(v string) *AddUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserResponseBody) SetResult(v *AddUserResponseBodyResult) *AddUserResponseBody {
	s.Result = v
	return s
}

func (s *AddUserResponseBody) SetSuccess(v bool) *AddUserResponseBody {
	s.Success = &v
	return s
}

type AddUserResponseBodyResult struct {
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AdminUser     *bool   `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	AuthAdminUser *bool   `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	NickName      *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType      *int32  `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s AddUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddUserResponseBodyResult) SetAccountName(v string) *AddUserResponseBodyResult {
	s.AccountName = &v
	return s
}

func (s *AddUserResponseBodyResult) SetAdminUser(v bool) *AddUserResponseBodyResult {
	s.AdminUser = &v
	return s
}

func (s *AddUserResponseBodyResult) SetAuthAdminUser(v bool) *AddUserResponseBodyResult {
	s.AuthAdminUser = &v
	return s
}

func (s *AddUserResponseBodyResult) SetNickName(v string) *AddUserResponseBodyResult {
	s.NickName = &v
	return s
}

func (s *AddUserResponseBodyResult) SetUserId(v string) *AddUserResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *AddUserResponseBodyResult) SetUserType(v int32) *AddUserResponseBodyResult {
	s.UserType = &v
	return s
}

type AddUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserResponse) GoString() string {
	return s.String()
}

func (s *AddUserResponse) SetHeaders(v map[string]*string) *AddUserResponse {
	s.Headers = v
	return s
}

func (s *AddUserResponse) SetStatusCode(v int32) *AddUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserResponse) SetBody(v *AddUserResponseBody) *AddUserResponse {
	s.Body = v
	return s
}

type AddUserGroupMemberRequest struct {
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserIdList  *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s AddUserGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberRequest) SetUserGroupId(v string) *AddUserGroupMemberRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMemberRequest) SetUserIdList(v string) *AddUserGroupMemberRequest {
	s.UserIdList = &v
	return s
}

type AddUserGroupMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberResponseBody) SetRequestId(v string) *AddUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetResult(v bool) *AddUserGroupMemberResponseBody {
	s.Result = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetSuccess(v bool) *AddUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

type AddUserGroupMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberResponse) SetHeaders(v map[string]*string) *AddUserGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *AddUserGroupMemberResponse) SetStatusCode(v int32) *AddUserGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserGroupMemberResponse) SetBody(v *AddUserGroupMemberResponseBody) *AddUserGroupMemberResponse {
	s.Body = v
	return s
}

type AddUserGroupMembersRequest struct {
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersRequest) SetUserGroupIds(v string) *AddUserGroupMembersRequest {
	s.UserGroupIds = &v
	return s
}

func (s *AddUserGroupMembersRequest) SetUserId(v string) *AddUserGroupMembersRequest {
	s.UserId = &v
	return s
}

type AddUserGroupMembersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersResponseBody) SetRequestId(v string) *AddUserGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetResult(v bool) *AddUserGroupMembersResponseBody {
	s.Result = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetSuccess(v bool) *AddUserGroupMembersResponseBody {
	s.Success = &v
	return s
}

type AddUserGroupMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersResponse) SetHeaders(v map[string]*string) *AddUserGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *AddUserGroupMembersResponse) SetStatusCode(v int32) *AddUserGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserGroupMembersResponse) SetBody(v *AddUserGroupMembersResponseBody) *AddUserGroupMembersResponse {
	s.Body = v
	return s
}

type AddUserTagMetaRequest struct {
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	TagName        *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s AddUserTagMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserTagMetaRequest) GoString() string {
	return s.String()
}

func (s *AddUserTagMetaRequest) SetTagDescription(v string) *AddUserTagMetaRequest {
	s.TagDescription = &v
	return s
}

func (s *AddUserTagMetaRequest) SetTagName(v string) *AddUserTagMetaRequest {
	s.TagName = &v
	return s
}

type AddUserTagMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserTagMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserTagMetaResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserTagMetaResponseBody) SetRequestId(v string) *AddUserTagMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserTagMetaResponseBody) SetResult(v string) *AddUserTagMetaResponseBody {
	s.Result = &v
	return s
}

func (s *AddUserTagMetaResponseBody) SetSuccess(v bool) *AddUserTagMetaResponseBody {
	s.Success = &v
	return s
}

type AddUserTagMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserTagMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserTagMetaResponse) GoString() string {
	return s.String()
}

func (s *AddUserTagMetaResponse) SetHeaders(v map[string]*string) *AddUserTagMetaResponse {
	s.Headers = v
	return s
}

func (s *AddUserTagMetaResponse) SetStatusCode(v int32) *AddUserTagMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserTagMetaResponse) SetBody(v *AddUserTagMetaResponseBody) *AddUserTagMetaResponse {
	s.Body = v
	return s
}

type AddUserToWorkspaceRequest struct {
	RoleId      *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddUserToWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *AddUserToWorkspaceRequest) SetRoleId(v int64) *AddUserToWorkspaceRequest {
	s.RoleId = &v
	return s
}

func (s *AddUserToWorkspaceRequest) SetUserId(v string) *AddUserToWorkspaceRequest {
	s.UserId = &v
	return s
}

func (s *AddUserToWorkspaceRequest) SetWorkspaceId(v string) *AddUserToWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

type AddUserToWorkspaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserToWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserToWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToWorkspaceResponseBody) SetRequestId(v string) *AddUserToWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToWorkspaceResponseBody) SetResult(v bool) *AddUserToWorkspaceResponseBody {
	s.Result = &v
	return s
}

func (s *AddUserToWorkspaceResponseBody) SetSuccess(v bool) *AddUserToWorkspaceResponseBody {
	s.Success = &v
	return s
}

type AddUserToWorkspaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUserToWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserToWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *AddUserToWorkspaceResponse) SetHeaders(v map[string]*string) *AddUserToWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *AddUserToWorkspaceResponse) SetStatusCode(v int32) *AddUserToWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToWorkspaceResponse) SetBody(v *AddUserToWorkspaceResponseBody) *AddUserToWorkspaceResponse {
	s.Body = v
	return s
}

type AddWorkspaceUsersRequest struct {
	RoleId      *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserIds     *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceUsersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUsersRequest) SetRoleId(v int64) *AddWorkspaceUsersRequest {
	s.RoleId = &v
	return s
}

func (s *AddWorkspaceUsersRequest) SetUserIds(v string) *AddWorkspaceUsersRequest {
	s.UserIds = &v
	return s
}

func (s *AddWorkspaceUsersRequest) SetWorkspaceId(v string) *AddWorkspaceUsersRequest {
	s.WorkspaceId = &v
	return s
}

type AddWorkspaceUsersResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddWorkspaceUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddWorkspaceUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUsersResponseBody) SetRequestId(v string) *AddWorkspaceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkspaceUsersResponseBody) SetResult(v *AddWorkspaceUsersResponseBodyResult) *AddWorkspaceUsersResponseBody {
	s.Result = v
	return s
}

func (s *AddWorkspaceUsersResponseBody) SetSuccess(v bool) *AddWorkspaceUsersResponseBody {
	s.Success = &v
	return s
}

type AddWorkspaceUsersResponseBodyResult struct {
	Failure       *int32                 `json:"Failure,omitempty" xml:"Failure,omitempty"`
	FailureDetail map[string]interface{} `json:"FailureDetail,omitempty" xml:"FailureDetail,omitempty"`
	Success       *int32                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total         *int32                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s AddWorkspaceUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUsersResponseBodyResult) SetFailure(v int32) *AddWorkspaceUsersResponseBodyResult {
	s.Failure = &v
	return s
}

func (s *AddWorkspaceUsersResponseBodyResult) SetFailureDetail(v map[string]interface{}) *AddWorkspaceUsersResponseBodyResult {
	s.FailureDetail = v
	return s
}

func (s *AddWorkspaceUsersResponseBodyResult) SetSuccess(v int32) *AddWorkspaceUsersResponseBodyResult {
	s.Success = &v
	return s
}

func (s *AddWorkspaceUsersResponseBodyResult) SetTotal(v int32) *AddWorkspaceUsersResponseBodyResult {
	s.Total = &v
	return s
}

type AddWorkspaceUsersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddWorkspaceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWorkspaceUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceUsersResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUsersResponse) SetHeaders(v map[string]*string) *AddWorkspaceUsersResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceUsersResponse) SetStatusCode(v int32) *AddWorkspaceUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceUsersResponse) SetBody(v *AddWorkspaceUsersResponseBody) *AddWorkspaceUsersResponse {
	s.Body = v
	return s
}

type AuthorizeMenuRequest struct {
	AuthPointsValue *int32  `json:"AuthPointsValue,omitempty" xml:"AuthPointsValue,omitempty"`
	DataPortalId    *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	MenuIds         *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	UserGroupIds    *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	UserIds         *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s AuthorizeMenuRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeMenuRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeMenuRequest) SetAuthPointsValue(v int32) *AuthorizeMenuRequest {
	s.AuthPointsValue = &v
	return s
}

func (s *AuthorizeMenuRequest) SetDataPortalId(v string) *AuthorizeMenuRequest {
	s.DataPortalId = &v
	return s
}

func (s *AuthorizeMenuRequest) SetMenuIds(v string) *AuthorizeMenuRequest {
	s.MenuIds = &v
	return s
}

func (s *AuthorizeMenuRequest) SetUserGroupIds(v string) *AuthorizeMenuRequest {
	s.UserGroupIds = &v
	return s
}

func (s *AuthorizeMenuRequest) SetUserIds(v string) *AuthorizeMenuRequest {
	s.UserIds = &v
	return s
}

type AuthorizeMenuResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeMenuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeMenuResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeMenuResponseBody) SetRequestId(v string) *AuthorizeMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeMenuResponseBody) SetResult(v int32) *AuthorizeMenuResponseBody {
	s.Result = &v
	return s
}

func (s *AuthorizeMenuResponseBody) SetSuccess(v bool) *AuthorizeMenuResponseBody {
	s.Success = &v
	return s
}

type AuthorizeMenuResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthorizeMenuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeMenuResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeMenuResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeMenuResponse) SetHeaders(v map[string]*string) *AuthorizeMenuResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeMenuResponse) SetStatusCode(v int32) *AuthorizeMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeMenuResponse) SetBody(v *AuthorizeMenuResponseBody) *AuthorizeMenuResponse {
	s.Body = v
	return s
}

type CancelAuthorizationMenuRequest struct {
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	MenuIds      *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	UserIds      *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s CancelAuthorizationMenuRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAuthorizationMenuRequest) GoString() string {
	return s.String()
}

func (s *CancelAuthorizationMenuRequest) SetDataPortalId(v string) *CancelAuthorizationMenuRequest {
	s.DataPortalId = &v
	return s
}

func (s *CancelAuthorizationMenuRequest) SetMenuIds(v string) *CancelAuthorizationMenuRequest {
	s.MenuIds = &v
	return s
}

func (s *CancelAuthorizationMenuRequest) SetUserGroupIds(v string) *CancelAuthorizationMenuRequest {
	s.UserGroupIds = &v
	return s
}

func (s *CancelAuthorizationMenuRequest) SetUserIds(v string) *CancelAuthorizationMenuRequest {
	s.UserIds = &v
	return s
}

type CancelAuthorizationMenuResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAuthorizationMenuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelAuthorizationMenuResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAuthorizationMenuResponseBody) SetRequestId(v string) *CancelAuthorizationMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAuthorizationMenuResponseBody) SetResult(v int32) *CancelAuthorizationMenuResponseBody {
	s.Result = &v
	return s
}

func (s *CancelAuthorizationMenuResponseBody) SetSuccess(v bool) *CancelAuthorizationMenuResponseBody {
	s.Success = &v
	return s
}

type CancelAuthorizationMenuResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelAuthorizationMenuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelAuthorizationMenuResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAuthorizationMenuResponse) GoString() string {
	return s.String()
}

func (s *CancelAuthorizationMenuResponse) SetHeaders(v map[string]*string) *CancelAuthorizationMenuResponse {
	s.Headers = v
	return s
}

func (s *CancelAuthorizationMenuResponse) SetStatusCode(v int32) *CancelAuthorizationMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAuthorizationMenuResponse) SetBody(v *CancelAuthorizationMenuResponseBody) *CancelAuthorizationMenuResponse {
	s.Body = v
	return s
}

type CancelCollectionRequest struct {
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s CancelCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCollectionRequest) GoString() string {
	return s.String()
}

func (s *CancelCollectionRequest) SetUserId(v string) *CancelCollectionRequest {
	s.UserId = &v
	return s
}

func (s *CancelCollectionRequest) SetWorksId(v string) *CancelCollectionRequest {
	s.WorksId = &v
	return s
}

type CancelCollectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCollectionResponseBody) SetRequestId(v string) *CancelCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCollectionResponseBody) SetResult(v bool) *CancelCollectionResponseBody {
	s.Result = &v
	return s
}

func (s *CancelCollectionResponseBody) SetSuccess(v bool) *CancelCollectionResponseBody {
	s.Success = &v
	return s
}

type CancelCollectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCollectionResponse) GoString() string {
	return s.String()
}

func (s *CancelCollectionResponse) SetHeaders(v map[string]*string) *CancelCollectionResponse {
	s.Headers = v
	return s
}

func (s *CancelCollectionResponse) SetStatusCode(v int32) *CancelCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCollectionResponse) SetBody(v *CancelCollectionResponseBody) *CancelCollectionResponse {
	s.Body = v
	return s
}

type CancelReportShareRequest struct {
	ReportId    *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ShareToIds  *string `json:"ShareToIds,omitempty" xml:"ShareToIds,omitempty"`
	ShareToType *int32  `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
}

func (s CancelReportShareRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelReportShareRequest) GoString() string {
	return s.String()
}

func (s *CancelReportShareRequest) SetReportId(v string) *CancelReportShareRequest {
	s.ReportId = &v
	return s
}

func (s *CancelReportShareRequest) SetShareToIds(v string) *CancelReportShareRequest {
	s.ShareToIds = &v
	return s
}

func (s *CancelReportShareRequest) SetShareToType(v int32) *CancelReportShareRequest {
	s.ShareToType = &v
	return s
}

type CancelReportShareResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelReportShareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelReportShareResponseBody) GoString() string {
	return s.String()
}

func (s *CancelReportShareResponseBody) SetRequestId(v string) *CancelReportShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelReportShareResponseBody) SetResult(v bool) *CancelReportShareResponseBody {
	s.Result = &v
	return s
}

func (s *CancelReportShareResponseBody) SetSuccess(v bool) *CancelReportShareResponseBody {
	s.Success = &v
	return s
}

type CancelReportShareResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelReportShareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelReportShareResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelReportShareResponse) GoString() string {
	return s.String()
}

func (s *CancelReportShareResponse) SetHeaders(v map[string]*string) *CancelReportShareResponse {
	s.Headers = v
	return s
}

func (s *CancelReportShareResponse) SetStatusCode(v int32) *CancelReportShareResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelReportShareResponse) SetBody(v *CancelReportShareResponseBody) *CancelReportShareResponse {
	s.Body = v
	return s
}

type ChangeVisibilityModelRequest struct {
	DataPortalId       *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	MenuIds            *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	ShowOnlyWithAccess *bool   `json:"ShowOnlyWithAccess,omitempty" xml:"ShowOnlyWithAccess,omitempty"`
}

func (s ChangeVisibilityModelRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeVisibilityModelRequest) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityModelRequest) SetDataPortalId(v string) *ChangeVisibilityModelRequest {
	s.DataPortalId = &v
	return s
}

func (s *ChangeVisibilityModelRequest) SetMenuIds(v string) *ChangeVisibilityModelRequest {
	s.MenuIds = &v
	return s
}

func (s *ChangeVisibilityModelRequest) SetShowOnlyWithAccess(v bool) *ChangeVisibilityModelRequest {
	s.ShowOnlyWithAccess = &v
	return s
}

type ChangeVisibilityModelResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeVisibilityModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeVisibilityModelResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityModelResponseBody) SetRequestId(v string) *ChangeVisibilityModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeVisibilityModelResponseBody) SetResult(v int32) *ChangeVisibilityModelResponseBody {
	s.Result = &v
	return s
}

func (s *ChangeVisibilityModelResponseBody) SetSuccess(v bool) *ChangeVisibilityModelResponseBody {
	s.Success = &v
	return s
}

type ChangeVisibilityModelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeVisibilityModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeVisibilityModelResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeVisibilityModelResponse) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityModelResponse) SetHeaders(v map[string]*string) *ChangeVisibilityModelResponse {
	s.Headers = v
	return s
}

func (s *ChangeVisibilityModelResponse) SetStatusCode(v int32) *ChangeVisibilityModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeVisibilityModelResponse) SetBody(v *ChangeVisibilityModelResponseBody) *ChangeVisibilityModelResponse {
	s.Body = v
	return s
}

type CheckReadableRequest struct {
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s CheckReadableRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckReadableRequest) GoString() string {
	return s.String()
}

func (s *CheckReadableRequest) SetUserId(v string) *CheckReadableRequest {
	s.UserId = &v
	return s
}

func (s *CheckReadableRequest) SetWorksId(v string) *CheckReadableRequest {
	s.WorksId = &v
	return s
}

type CheckReadableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckReadableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckReadableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckReadableResponseBody) SetRequestId(v string) *CheckReadableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckReadableResponseBody) SetResult(v bool) *CheckReadableResponseBody {
	s.Result = &v
	return s
}

func (s *CheckReadableResponseBody) SetSuccess(v bool) *CheckReadableResponseBody {
	s.Success = &v
	return s
}

type CheckReadableResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckReadableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckReadableResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckReadableResponse) GoString() string {
	return s.String()
}

func (s *CheckReadableResponse) SetHeaders(v map[string]*string) *CheckReadableResponse {
	s.Headers = v
	return s
}

func (s *CheckReadableResponse) SetStatusCode(v int32) *CheckReadableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckReadableResponse) SetBody(v *CheckReadableResponseBody) *CheckReadableResponse {
	s.Body = v
	return s
}

type CreateTicketRequest struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType    *int32  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	CmptId         *string `json:"CmptId,omitempty" xml:"CmptId,omitempty"`
	ExpireTime     *int32  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GlobalParam    *string `json:"GlobalParam,omitempty" xml:"GlobalParam,omitempty"`
	TicketNum      *int32  `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WatermarkParam *string `json:"WatermarkParam,omitempty" xml:"WatermarkParam,omitempty"`
	WorksId        *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetAccountName(v string) *CreateTicketRequest {
	s.AccountName = &v
	return s
}

func (s *CreateTicketRequest) SetAccountType(v int32) *CreateTicketRequest {
	s.AccountType = &v
	return s
}

func (s *CreateTicketRequest) SetCmptId(v string) *CreateTicketRequest {
	s.CmptId = &v
	return s
}

func (s *CreateTicketRequest) SetExpireTime(v int32) *CreateTicketRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateTicketRequest) SetGlobalParam(v string) *CreateTicketRequest {
	s.GlobalParam = &v
	return s
}

func (s *CreateTicketRequest) SetTicketNum(v int32) *CreateTicketRequest {
	s.TicketNum = &v
	return s
}

func (s *CreateTicketRequest) SetUserId(v string) *CreateTicketRequest {
	s.UserId = &v
	return s
}

func (s *CreateTicketRequest) SetWatermarkParam(v string) *CreateTicketRequest {
	s.WatermarkParam = &v
	return s
}

func (s *CreateTicketRequest) SetWorksId(v string) *CreateTicketRequest {
	s.WorksId = &v
	return s
}

type CreateTicketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetResult(v string) *CreateTicketResponseBody {
	s.Result = &v
	return s
}

func (s *CreateTicketResponseBody) SetSuccess(v bool) *CreateTicketResponseBody {
	s.Success = &v
	return s
}

type CreateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetStatusCode(v int32) *CreateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type CreateUserGroupRequest struct {
	ParentUserGroupId    *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	UserGroupId          *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName        *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) SetParentUserGroupId(v string) *CreateUserGroupRequest {
	s.ParentUserGroupId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupDescription(v string) *CreateUserGroupRequest {
	s.UserGroupDescription = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupId(v string) *CreateUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupName(v string) *CreateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

type CreateUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBody) SetRequestId(v string) *CreateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetResult(v string) *CreateUserGroupResponseBody {
	s.Result = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetSuccess(v bool) *CreateUserGroupResponseBody {
	s.Success = &v
	return s
}

type CreateUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponse) SetHeaders(v map[string]*string) *CreateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateUserGroupResponse) SetStatusCode(v int32) *CreateUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserGroupResponse) SetBody(v *CreateUserGroupResponseBody) *CreateUserGroupResponse {
	s.Body = v
	return s
}

type DelayTicketExpireTimeRequest struct {
	ExpireTime *int32  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Ticket     *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s DelayTicketExpireTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s DelayTicketExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *DelayTicketExpireTimeRequest) SetExpireTime(v int32) *DelayTicketExpireTimeRequest {
	s.ExpireTime = &v
	return s
}

func (s *DelayTicketExpireTimeRequest) SetTicket(v string) *DelayTicketExpireTimeRequest {
	s.Ticket = &v
	return s
}

type DelayTicketExpireTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelayTicketExpireTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelayTicketExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DelayTicketExpireTimeResponseBody) SetRequestId(v string) *DelayTicketExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelayTicketExpireTimeResponseBody) SetResult(v bool) *DelayTicketExpireTimeResponseBody {
	s.Result = &v
	return s
}

func (s *DelayTicketExpireTimeResponseBody) SetSuccess(v bool) *DelayTicketExpireTimeResponseBody {
	s.Success = &v
	return s
}

type DelayTicketExpireTimeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DelayTicketExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DelayTicketExpireTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s DelayTicketExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *DelayTicketExpireTimeResponse) SetHeaders(v map[string]*string) *DelayTicketExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *DelayTicketExpireTimeResponse) SetStatusCode(v int32) *DelayTicketExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DelayTicketExpireTimeResponse) SetBody(v *DelayTicketExpireTimeResponseBody) *DelayTicketExpireTimeResponse {
	s.Body = v
	return s
}

type DeleteDataLevelPermissionRuleUsersRequest struct {
	DeleteUserModel *string `json:"DeleteUserModel,omitempty" xml:"DeleteUserModel,omitempty"`
}

func (s DeleteDataLevelPermissionRuleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLevelPermissionRuleUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelPermissionRuleUsersRequest) SetDeleteUserModel(v string) *DeleteDataLevelPermissionRuleUsersRequest {
	s.DeleteUserModel = &v
	return s
}

type DeleteDataLevelPermissionRuleUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataLevelPermissionRuleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLevelPermissionRuleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) SetRequestId(v string) *DeleteDataLevelPermissionRuleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) SetResult(v bool) *DeleteDataLevelPermissionRuleUsersResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponseBody) SetSuccess(v bool) *DeleteDataLevelPermissionRuleUsersResponseBody {
	s.Success = &v
	return s
}

type DeleteDataLevelPermissionRuleUsersResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDataLevelPermissionRuleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataLevelPermissionRuleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLevelPermissionRuleUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) SetHeaders(v map[string]*string) *DeleteDataLevelPermissionRuleUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) SetStatusCode(v int32) *DeleteDataLevelPermissionRuleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLevelPermissionRuleUsersResponse) SetBody(v *DeleteDataLevelPermissionRuleUsersResponseBody) *DeleteDataLevelPermissionRuleUsersResponse {
	s.Body = v
	return s
}

type DeleteDataLevelRuleConfigRequest struct {
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteDataLevelRuleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLevelRuleConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelRuleConfigRequest) SetCubeId(v string) *DeleteDataLevelRuleConfigRequest {
	s.CubeId = &v
	return s
}

func (s *DeleteDataLevelRuleConfigRequest) SetRuleId(v string) *DeleteDataLevelRuleConfigRequest {
	s.RuleId = &v
	return s
}

type DeleteDataLevelRuleConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataLevelRuleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLevelRuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelRuleConfigResponseBody) SetRequestId(v string) *DeleteDataLevelRuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLevelRuleConfigResponseBody) SetResult(v bool) *DeleteDataLevelRuleConfigResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteDataLevelRuleConfigResponseBody) SetSuccess(v bool) *DeleteDataLevelRuleConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteDataLevelRuleConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDataLevelRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataLevelRuleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataLevelRuleConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLevelRuleConfigResponse) SetHeaders(v map[string]*string) *DeleteDataLevelRuleConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLevelRuleConfigResponse) SetStatusCode(v int32) *DeleteDataLevelRuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLevelRuleConfigResponse) SetBody(v *DeleteDataLevelRuleConfigResponseBody) *DeleteDataLevelRuleConfigResponse {
	s.Body = v
	return s
}

type DeleteTicketRequest struct {
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s DeleteTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTicketRequest) GoString() string {
	return s.String()
}

func (s *DeleteTicketRequest) SetTicket(v string) *DeleteTicketRequest {
	s.Ticket = &v
	return s
}

type DeleteTicketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTicketResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTicketResponseBody) SetRequestId(v string) *DeleteTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTicketResponseBody) SetResult(v bool) *DeleteTicketResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteTicketResponseBody) SetSuccess(v bool) *DeleteTicketResponseBody {
	s.Success = &v
	return s
}

type DeleteTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTicketResponse) GoString() string {
	return s.String()
}

func (s *DeleteTicketResponse) SetHeaders(v map[string]*string) *DeleteTicketResponse {
	s.Headers = v
	return s
}

func (s *DeleteTicketResponse) SetStatusCode(v int32) *DeleteTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTicketResponse) SetBody(v *DeleteTicketResponseBody) *DeleteTicketResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	TransferUserId *string `json:"TransferUserId,omitempty" xml:"TransferUserId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetTransferUserId(v string) *DeleteUserRequest {
	s.TransferUserId = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetResult(v bool) *DeleteUserResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeleteUserFromWorkspaceRequest struct {
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteUserFromWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserFromWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserFromWorkspaceRequest) SetUserId(v string) *DeleteUserFromWorkspaceRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserFromWorkspaceRequest) SetWorkspaceId(v string) *DeleteUserFromWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteUserFromWorkspaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserFromWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserFromWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserFromWorkspaceResponseBody) SetRequestId(v string) *DeleteUserFromWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserFromWorkspaceResponseBody) SetResult(v bool) *DeleteUserFromWorkspaceResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserFromWorkspaceResponseBody) SetSuccess(v bool) *DeleteUserFromWorkspaceResponseBody {
	s.Success = &v
	return s
}

type DeleteUserFromWorkspaceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserFromWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserFromWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserFromWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserFromWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteUserFromWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserFromWorkspaceResponse) SetStatusCode(v int32) *DeleteUserFromWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserFromWorkspaceResponse) SetBody(v *DeleteUserFromWorkspaceResponseBody) *DeleteUserFromWorkspaceResponse {
	s.Body = v
	return s
}

type DeleteUserGroupRequest struct {
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) SetUserGroupId(v string) *DeleteUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type DeleteUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetResult(v bool) *DeleteUserGroupResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetSuccess(v bool) *DeleteUserGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponse) SetHeaders(v map[string]*string) *DeleteUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupResponse) SetStatusCode(v int32) *DeleteUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupResponse) SetBody(v *DeleteUserGroupResponseBody) *DeleteUserGroupResponse {
	s.Body = v
	return s
}

type DeleteUserGroupMemberRequest struct {
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMemberRequest) SetUserGroupId(v string) *DeleteUserGroupMemberRequest {
	s.UserGroupId = &v
	return s
}

func (s *DeleteUserGroupMemberRequest) SetUserId(v string) *DeleteUserGroupMemberRequest {
	s.UserId = &v
	return s
}

type DeleteUserGroupMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMemberResponseBody) SetRequestId(v string) *DeleteUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupMemberResponseBody) SetResult(v bool) *DeleteUserGroupMemberResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserGroupMemberResponseBody) SetSuccess(v bool) *DeleteUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

type DeleteUserGroupMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMemberResponse) SetHeaders(v map[string]*string) *DeleteUserGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupMemberResponse) SetStatusCode(v int32) *DeleteUserGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupMemberResponse) SetBody(v *DeleteUserGroupMemberResponseBody) *DeleteUserGroupMemberResponse {
	s.Body = v
	return s
}

type DeleteUserGroupMembersRequest struct {
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMembersRequest) SetUserGroupIds(v string) *DeleteUserGroupMembersRequest {
	s.UserGroupIds = &v
	return s
}

func (s *DeleteUserGroupMembersRequest) SetUserId(v string) *DeleteUserGroupMembersRequest {
	s.UserId = &v
	return s
}

type DeleteUserGroupMembersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMembersResponseBody) SetRequestId(v string) *DeleteUserGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupMembersResponseBody) SetResult(v bool) *DeleteUserGroupMembersResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserGroupMembersResponseBody) SetSuccess(v bool) *DeleteUserGroupMembersResponseBody {
	s.Success = &v
	return s
}

type DeleteUserGroupMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMembersResponse) SetHeaders(v map[string]*string) *DeleteUserGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupMembersResponse) SetStatusCode(v int32) *DeleteUserGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupMembersResponse) SetBody(v *DeleteUserGroupMembersResponseBody) *DeleteUserGroupMembersResponse {
	s.Body = v
	return s
}

type DeleteUserTagMetaRequest struct {
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DeleteUserTagMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserTagMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserTagMetaRequest) SetTagId(v string) *DeleteUserTagMetaRequest {
	s.TagId = &v
	return s
}

type DeleteUserTagMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserTagMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserTagMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserTagMetaResponseBody) SetRequestId(v string) *DeleteUserTagMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserTagMetaResponseBody) SetResult(v bool) *DeleteUserTagMetaResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserTagMetaResponseBody) SetSuccess(v bool) *DeleteUserTagMetaResponseBody {
	s.Success = &v
	return s
}

type DeleteUserTagMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserTagMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserTagMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserTagMetaResponse) SetHeaders(v map[string]*string) *DeleteUserTagMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserTagMetaResponse) SetStatusCode(v int32) *DeleteUserTagMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserTagMetaResponse) SetBody(v *DeleteUserTagMetaResponseBody) *DeleteUserTagMetaResponse {
	s.Body = v
	return s
}

type GetUserGroupInfoRequest struct {
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s GetUserGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupInfoRequest) SetKeyword(v string) *GetUserGroupInfoRequest {
	s.Keyword = &v
	return s
}

type GetUserGroupInfoResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetUserGroupInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupInfoResponseBody) SetRequestId(v string) *GetUserGroupInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupInfoResponseBody) SetResult(v []*GetUserGroupInfoResponseBodyResult) *GetUserGroupInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetUserGroupInfoResponseBody) SetSuccess(v bool) *GetUserGroupInfoResponseBody {
	s.Success = &v
	return s
}

type GetUserGroupInfoResponseBodyResult struct {
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser        *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	IdentifiedPath    *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModifyUser        *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	ParentUsergroupId *string `json:"ParentUsergroupId,omitempty" xml:"ParentUsergroupId,omitempty"`
	UsergroupDesc     *string `json:"UsergroupDesc,omitempty" xml:"UsergroupDesc,omitempty"`
	UsergroupId       *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	UsergroupName     *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
}

func (s GetUserGroupInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserGroupInfoResponseBodyResult) SetCreateTime(v string) *GetUserGroupInfoResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetCreateUser(v string) *GetUserGroupInfoResponseBodyResult {
	s.CreateUser = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetIdentifiedPath(v string) *GetUserGroupInfoResponseBodyResult {
	s.IdentifiedPath = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetModifiedTime(v string) *GetUserGroupInfoResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetModifyUser(v string) *GetUserGroupInfoResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetParentUsergroupId(v string) *GetUserGroupInfoResponseBodyResult {
	s.ParentUsergroupId = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetUsergroupDesc(v string) *GetUserGroupInfoResponseBodyResult {
	s.UsergroupDesc = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetUsergroupId(v string) *GetUserGroupInfoResponseBodyResult {
	s.UsergroupId = &v
	return s
}

func (s *GetUserGroupInfoResponseBodyResult) SetUsergroupName(v string) *GetUserGroupInfoResponseBodyResult {
	s.UsergroupName = &v
	return s
}

type GetUserGroupInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserGroupInfoResponse) SetHeaders(v map[string]*string) *GetUserGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserGroupInfoResponse) SetStatusCode(v int32) *GetUserGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGroupInfoResponse) SetBody(v *GetUserGroupInfoResponseBody) *GetUserGroupInfoResponse {
	s.Body = v
	return s
}

type ListByUserGroupIdRequest struct {
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
}

func (s ListByUserGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListByUserGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdRequest) SetUserGroupIds(v string) *ListByUserGroupIdRequest {
	s.UserGroupIds = &v
	return s
}

type ListByUserGroupIdResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListByUserGroupIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListByUserGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListByUserGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdResponseBody) SetRequestId(v string) *ListByUserGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListByUserGroupIdResponseBody) SetResult(v *ListByUserGroupIdResponseBodyResult) *ListByUserGroupIdResponseBody {
	s.Result = v
	return s
}

func (s *ListByUserGroupIdResponseBody) SetSuccess(v bool) *ListByUserGroupIdResponseBody {
	s.Success = &v
	return s
}

type ListByUserGroupIdResponseBodyResult struct {
	FailedUserGroupIds []*string                                             `json:"FailedUserGroupIds,omitempty" xml:"FailedUserGroupIds,omitempty" type:"Repeated"`
	UserGroupModels    []*ListByUserGroupIdResponseBodyResultUserGroupModels `json:"UserGroupModels,omitempty" xml:"UserGroupModels,omitempty" type:"Repeated"`
}

func (s ListByUserGroupIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListByUserGroupIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdResponseBodyResult) SetFailedUserGroupIds(v []*string) *ListByUserGroupIdResponseBodyResult {
	s.FailedUserGroupIds = v
	return s
}

func (s *ListByUserGroupIdResponseBodyResult) SetUserGroupModels(v []*ListByUserGroupIdResponseBodyResultUserGroupModels) *ListByUserGroupIdResponseBodyResult {
	s.UserGroupModels = v
	return s
}

type ListByUserGroupIdResponseBodyResultUserGroupModels struct {
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser        *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	IdentifiedPath    *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModifyUser        *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	ParentUsergroupId *string `json:"ParentUsergroupId,omitempty" xml:"ParentUsergroupId,omitempty"`
	UsergroupDesc     *string `json:"UsergroupDesc,omitempty" xml:"UsergroupDesc,omitempty"`
	UsergroupId       *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	UsergroupName     *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
}

func (s ListByUserGroupIdResponseBodyResultUserGroupModels) String() string {
	return tea.Prettify(s)
}

func (s ListByUserGroupIdResponseBodyResultUserGroupModels) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetCreateTime(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.CreateTime = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetCreateUser(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.CreateUser = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetIdentifiedPath(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.IdentifiedPath = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetModifiedTime(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.ModifiedTime = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetModifyUser(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.ModifyUser = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetParentUsergroupId(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.ParentUsergroupId = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetUsergroupDesc(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.UsergroupDesc = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetUsergroupId(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.UsergroupId = &v
	return s
}

func (s *ListByUserGroupIdResponseBodyResultUserGroupModels) SetUsergroupName(v string) *ListByUserGroupIdResponseBodyResultUserGroupModels {
	s.UsergroupName = &v
	return s
}

type ListByUserGroupIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListByUserGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListByUserGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListByUserGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListByUserGroupIdResponse) SetHeaders(v map[string]*string) *ListByUserGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListByUserGroupIdResponse) SetStatusCode(v int32) *ListByUserGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListByUserGroupIdResponse) SetBody(v *ListByUserGroupIdResponseBody) *ListByUserGroupIdResponse {
	s.Body = v
	return s
}

type ListCollectionsRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCollectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListCollectionsRequest) SetUserId(v string) *ListCollectionsRequest {
	s.UserId = &v
	return s
}

type ListCollectionsResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCollectionsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCollectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponseBody) SetRequestId(v string) *ListCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollectionsResponseBody) SetResult(v []*ListCollectionsResponseBodyResult) *ListCollectionsResponseBody {
	s.Result = v
	return s
}

func (s *ListCollectionsResponseBody) SetSuccess(v bool) *ListCollectionsResponseBody {
	s.Success = &v
	return s
}

type ListCollectionsResponseBodyResult struct {
	FavoriteId    *int32  `json:"FavoriteId,omitempty" xml:"FavoriteId,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	WorksId       *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	WorksName     *string `json:"WorksName,omitempty" xml:"WorksName,omitempty"`
	WorksType     *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListCollectionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponseBodyResult) SetFavoriteId(v int32) *ListCollectionsResponseBodyResult {
	s.FavoriteId = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetOwnerId(v string) *ListCollectionsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorksId(v string) *ListCollectionsResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorksName(v string) *ListCollectionsResponseBodyResult {
	s.WorksName = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorksType(v string) *ListCollectionsResponseBodyResult {
	s.WorksType = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorkspaceId(v string) *ListCollectionsResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *ListCollectionsResponseBodyResult) SetWorkspaceName(v string) *ListCollectionsResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

type ListCollectionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCollectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListCollectionsResponse) SetHeaders(v map[string]*string) *ListCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListCollectionsResponse) SetStatusCode(v int32) *ListCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCollectionsResponse) SetBody(v *ListCollectionsResponseBody) *ListCollectionsResponse {
	s.Body = v
	return s
}

type ListCubeDataLevelPermissionConfigRequest struct {
	CubeId   *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListCubeDataLevelPermissionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCubeDataLevelPermissionConfigRequest) GoString() string {
	return s.String()
}

func (s *ListCubeDataLevelPermissionConfigRequest) SetCubeId(v string) *ListCubeDataLevelPermissionConfigRequest {
	s.CubeId = &v
	return s
}

func (s *ListCubeDataLevelPermissionConfigRequest) SetRuleType(v string) *ListCubeDataLevelPermissionConfigRequest {
	s.RuleType = &v
	return s
}

type ListCubeDataLevelPermissionConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCubeDataLevelPermissionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCubeDataLevelPermissionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListCubeDataLevelPermissionConfigResponseBody) SetRequestId(v string) *ListCubeDataLevelPermissionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCubeDataLevelPermissionConfigResponseBody) SetResult(v string) *ListCubeDataLevelPermissionConfigResponseBody {
	s.Result = &v
	return s
}

func (s *ListCubeDataLevelPermissionConfigResponseBody) SetSuccess(v bool) *ListCubeDataLevelPermissionConfigResponseBody {
	s.Success = &v
	return s
}

type ListCubeDataLevelPermissionConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCubeDataLevelPermissionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCubeDataLevelPermissionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCubeDataLevelPermissionConfigResponse) GoString() string {
	return s.String()
}

func (s *ListCubeDataLevelPermissionConfigResponse) SetHeaders(v map[string]*string) *ListCubeDataLevelPermissionConfigResponse {
	s.Headers = v
	return s
}

func (s *ListCubeDataLevelPermissionConfigResponse) SetStatusCode(v int32) *ListCubeDataLevelPermissionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCubeDataLevelPermissionConfigResponse) SetBody(v *ListCubeDataLevelPermissionConfigResponseBody) *ListCubeDataLevelPermissionConfigResponse {
	s.Body = v
	return s
}

type ListDataLevelPermissionWhiteListRequest struct {
	CubeId   *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListDataLevelPermissionWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListRequest) SetCubeId(v string) *ListDataLevelPermissionWhiteListRequest {
	s.CubeId = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListRequest) SetRuleType(v string) *ListDataLevelPermissionWhiteListRequest {
	s.RuleType = &v
	return s
}

type ListDataLevelPermissionWhiteListResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListDataLevelPermissionWhiteListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLevelPermissionWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListResponseBody) SetRequestId(v string) *ListDataLevelPermissionWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBody) SetResult(v *ListDataLevelPermissionWhiteListResponseBodyResult) *ListDataLevelPermissionWhiteListResponseBody {
	s.Result = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBody) SetSuccess(v bool) *ListDataLevelPermissionWhiteListResponseBody {
	s.Success = &v
	return s
}

type ListDataLevelPermissionWhiteListResponseBodyResult struct {
	CubeId     *string                                                       `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	RuleType   *string                                                       `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	UsersModel *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel `json:"UsersModel,omitempty" xml:"UsersModel,omitempty" type:"Struct"`
}

func (s ListDataLevelPermissionWhiteListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) SetCubeId(v string) *ListDataLevelPermissionWhiteListResponseBodyResult {
	s.CubeId = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) SetRuleType(v string) *ListDataLevelPermissionWhiteListResponseBodyResult {
	s.RuleType = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResult) SetUsersModel(v *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) *ListDataLevelPermissionWhiteListResponseBodyResult {
	s.UsersModel = v
	return s
}

type ListDataLevelPermissionWhiteListResponseBodyResultUsersModel struct {
	UserGroups []*string `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
	Users      []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) String() string {
	return tea.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) SetUserGroups(v []*string) *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel {
	s.UserGroups = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel) SetUsers(v []*string) *ListDataLevelPermissionWhiteListResponseBodyResultUsersModel {
	s.Users = v
	return s
}

type ListDataLevelPermissionWhiteListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataLevelPermissionWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataLevelPermissionWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ListDataLevelPermissionWhiteListResponse) SetHeaders(v map[string]*string) *ListDataLevelPermissionWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponse) SetStatusCode(v int32) *ListDataLevelPermissionWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLevelPermissionWhiteListResponse) SetBody(v *ListDataLevelPermissionWhiteListResponseBody) *ListDataLevelPermissionWhiteListResponse {
	s.Body = v
	return s
}

type ListFavoriteReportsRequest struct {
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TreeType *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListFavoriteReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFavoriteReportsRequest) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsRequest) SetKeyword(v string) *ListFavoriteReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListFavoriteReportsRequest) SetPageSize(v int32) *ListFavoriteReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFavoriteReportsRequest) SetTreeType(v string) *ListFavoriteReportsRequest {
	s.TreeType = &v
	return s
}

func (s *ListFavoriteReportsRequest) SetUserId(v string) *ListFavoriteReportsRequest {
	s.UserId = &v
	return s
}

type ListFavoriteReportsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListFavoriteReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFavoriteReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFavoriteReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsResponseBody) SetRequestId(v string) *ListFavoriteReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFavoriteReportsResponseBody) SetResult(v *ListFavoriteReportsResponseBodyResult) *ListFavoriteReportsResponseBody {
	s.Result = v
	return s
}

func (s *ListFavoriteReportsResponseBody) SetSuccess(v bool) *ListFavoriteReportsResponseBody {
	s.Success = &v
	return s
}

type ListFavoriteReportsResponseBodyResult struct {
	Data       []*ListFavoriteReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                       `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                       `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                       `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListFavoriteReportsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFavoriteReportsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsResponseBodyResult) SetData(v []*ListFavoriteReportsResponseBodyResultData) *ListFavoriteReportsResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) SetPageNum(v int32) *ListFavoriteReportsResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) SetPageSize(v int32) *ListFavoriteReportsResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) SetTotalNum(v int32) *ListFavoriteReportsResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) SetTotalPages(v int32) *ListFavoriteReportsResponseBodyResult {
	s.TotalPages = &v
	return s
}

type ListFavoriteReportsResponseBodyResultData struct {
	Favorite      *bool   `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	GmtCreate     *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HasEditAuth   *bool   `json:"HasEditAuth,omitempty" xml:"HasEditAuth,omitempty"`
	HasViewAuth   *bool   `json:"HasViewAuth,omitempty" xml:"HasViewAuth,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName     *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerNum      *string `json:"OwnerNum,omitempty" xml:"OwnerNum,omitempty"`
	PublishStatus *int32  `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	TreeId        *string `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListFavoriteReportsResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListFavoriteReportsResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsResponseBodyResultData) SetFavorite(v bool) *ListFavoriteReportsResponseBodyResultData {
	s.Favorite = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetGmtCreate(v string) *ListFavoriteReportsResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetGmtModified(v string) *ListFavoriteReportsResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetHasEditAuth(v bool) *ListFavoriteReportsResponseBodyResultData {
	s.HasEditAuth = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetHasViewAuth(v bool) *ListFavoriteReportsResponseBodyResultData {
	s.HasViewAuth = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetName(v string) *ListFavoriteReportsResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetOwnerName(v string) *ListFavoriteReportsResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetOwnerNum(v string) *ListFavoriteReportsResponseBodyResultData {
	s.OwnerNum = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetPublishStatus(v int32) *ListFavoriteReportsResponseBodyResultData {
	s.PublishStatus = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetTreeId(v string) *ListFavoriteReportsResponseBodyResultData {
	s.TreeId = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetType(v string) *ListFavoriteReportsResponseBodyResultData {
	s.Type = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetWorkspaceId(v string) *ListFavoriteReportsResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetWorkspaceName(v string) *ListFavoriteReportsResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type ListFavoriteReportsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFavoriteReportsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFavoriteReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFavoriteReportsResponse) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsResponse) SetHeaders(v map[string]*string) *ListFavoriteReportsResponse {
	s.Headers = v
	return s
}

func (s *ListFavoriteReportsResponse) SetStatusCode(v int32) *ListFavoriteReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFavoriteReportsResponse) SetBody(v *ListFavoriteReportsResponseBody) *ListFavoriteReportsResponse {
	s.Body = v
	return s
}

type ListPortalMenuAuthorizationRequest struct {
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
}

func (s ListPortalMenuAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPortalMenuAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationRequest) SetDataPortalId(v string) *ListPortalMenuAuthorizationRequest {
	s.DataPortalId = &v
	return s
}

type ListPortalMenuAuthorizationResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListPortalMenuAuthorizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPortalMenuAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPortalMenuAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationResponseBody) SetRequestId(v string) *ListPortalMenuAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBody) SetResult(v []*ListPortalMenuAuthorizationResponseBodyResult) *ListPortalMenuAuthorizationResponseBody {
	s.Result = v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBody) SetSuccess(v bool) *ListPortalMenuAuthorizationResponseBody {
	s.Success = &v
	return s
}

type ListPortalMenuAuthorizationResponseBodyResult struct {
	MenuId             *string                                                   `json:"MenuId,omitempty" xml:"MenuId,omitempty"`
	Receivers          []*ListPortalMenuAuthorizationResponseBodyResultReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	ShowOnlyWithAccess *bool                                                     `json:"ShowOnlyWithAccess,omitempty" xml:"ShowOnlyWithAccess,omitempty"`
}

func (s ListPortalMenuAuthorizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListPortalMenuAuthorizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) SetMenuId(v string) *ListPortalMenuAuthorizationResponseBodyResult {
	s.MenuId = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) SetReceivers(v []*ListPortalMenuAuthorizationResponseBodyResultReceivers) *ListPortalMenuAuthorizationResponseBodyResult {
	s.Receivers = v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResult) SetShowOnlyWithAccess(v bool) *ListPortalMenuAuthorizationResponseBodyResult {
	s.ShowOnlyWithAccess = &v
	return s
}

type ListPortalMenuAuthorizationResponseBodyResultReceivers struct {
	ReceiverId   *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ReceiverType *int32  `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
}

func (s ListPortalMenuAuthorizationResponseBodyResultReceivers) String() string {
	return tea.Prettify(s)
}

func (s ListPortalMenuAuthorizationResponseBodyResultReceivers) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) SetReceiverId(v string) *ListPortalMenuAuthorizationResponseBodyResultReceivers {
	s.ReceiverId = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponseBodyResultReceivers) SetReceiverType(v int32) *ListPortalMenuAuthorizationResponseBodyResultReceivers {
	s.ReceiverType = &v
	return s
}

type ListPortalMenuAuthorizationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPortalMenuAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPortalMenuAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPortalMenuAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ListPortalMenuAuthorizationResponse) SetHeaders(v map[string]*string) *ListPortalMenuAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ListPortalMenuAuthorizationResponse) SetStatusCode(v int32) *ListPortalMenuAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPortalMenuAuthorizationResponse) SetBody(v *ListPortalMenuAuthorizationResponseBody) *ListPortalMenuAuthorizationResponse {
	s.Body = v
	return s
}

type ListPortalMenusRequest struct {
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPortalMenusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPortalMenusRequest) GoString() string {
	return s.String()
}

func (s *ListPortalMenusRequest) SetDataPortalId(v string) *ListPortalMenusRequest {
	s.DataPortalId = &v
	return s
}

func (s *ListPortalMenusRequest) SetUserId(v string) *ListPortalMenusRequest {
	s.UserId = &v
	return s
}

type ListPortalMenusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPortalMenusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPortalMenusResponseBody) GoString() string {
	return s.String()
}

func (s *ListPortalMenusResponseBody) SetRequestId(v string) *ListPortalMenusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPortalMenusResponseBody) SetResult(v string) *ListPortalMenusResponseBody {
	s.Result = &v
	return s
}

func (s *ListPortalMenusResponseBody) SetSuccess(v bool) *ListPortalMenusResponseBody {
	s.Success = &v
	return s
}

type ListPortalMenusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPortalMenusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPortalMenusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPortalMenusResponse) GoString() string {
	return s.String()
}

func (s *ListPortalMenusResponse) SetHeaders(v map[string]*string) *ListPortalMenusResponse {
	s.Headers = v
	return s
}

func (s *ListPortalMenusResponse) SetStatusCode(v int32) *ListPortalMenusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPortalMenusResponse) SetBody(v *ListPortalMenusResponseBody) *ListPortalMenusResponse {
	s.Body = v
	return s
}

type ListRecentViewReportsRequest struct {
	Keyword   *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OffsetDay *int32  `json:"OffsetDay,omitempty" xml:"OffsetDay,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	TreeType  *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRecentViewReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecentViewReportsRequest) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsRequest) SetKeyword(v string) *ListRecentViewReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetOffsetDay(v int32) *ListRecentViewReportsRequest {
	s.OffsetDay = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetPageSize(v int32) *ListRecentViewReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetQueryMode(v string) *ListRecentViewReportsRequest {
	s.QueryMode = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetTreeType(v string) *ListRecentViewReportsRequest {
	s.TreeType = &v
	return s
}

func (s *ListRecentViewReportsRequest) SetUserId(v string) *ListRecentViewReportsRequest {
	s.UserId = &v
	return s
}

type ListRecentViewReportsResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListRecentViewReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRecentViewReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentViewReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponseBody) SetRequestId(v string) *ListRecentViewReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentViewReportsResponseBody) SetResult(v *ListRecentViewReportsResponseBodyResult) *ListRecentViewReportsResponseBody {
	s.Result = v
	return s
}

func (s *ListRecentViewReportsResponseBody) SetSuccess(v bool) *ListRecentViewReportsResponseBody {
	s.Success = &v
	return s
}

type ListRecentViewReportsResponseBodyResult struct {
	Data       []*ListRecentViewReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                         `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                         `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                         `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListRecentViewReportsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRecentViewReportsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponseBodyResult) SetData(v []*ListRecentViewReportsResponseBodyResultData) *ListRecentViewReportsResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetPageNum(v int32) *ListRecentViewReportsResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetPageSize(v int32) *ListRecentViewReportsResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetTotalNum(v int32) *ListRecentViewReportsResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetTotalPages(v int32) *ListRecentViewReportsResponseBodyResult {
	s.TotalPages = &v
	return s
}

type ListRecentViewReportsResponseBodyResultData struct {
	Favorite       *bool   `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HasEditAuth    *bool   `json:"HasEditAuth,omitempty" xml:"HasEditAuth,omitempty"`
	HasViewAuth    *bool   `json:"HasViewAuth,omitempty" xml:"HasViewAuth,omitempty"`
	LatestViewTime *string `json:"LatestViewTime,omitempty" xml:"LatestViewTime,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName      *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerNum       *string `json:"OwnerNum,omitempty" xml:"OwnerNum,omitempty"`
	PublishStatus  *int32  `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	TreeId         *string `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ViewCount      *int64  `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
	WorkspaceId    *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName  *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListRecentViewReportsResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListRecentViewReportsResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponseBodyResultData) SetFavorite(v bool) *ListRecentViewReportsResponseBodyResultData {
	s.Favorite = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetGmtCreate(v string) *ListRecentViewReportsResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetGmtModified(v string) *ListRecentViewReportsResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetHasEditAuth(v bool) *ListRecentViewReportsResponseBodyResultData {
	s.HasEditAuth = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetHasViewAuth(v bool) *ListRecentViewReportsResponseBodyResultData {
	s.HasViewAuth = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetLatestViewTime(v string) *ListRecentViewReportsResponseBodyResultData {
	s.LatestViewTime = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetName(v string) *ListRecentViewReportsResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetOwnerName(v string) *ListRecentViewReportsResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetOwnerNum(v string) *ListRecentViewReportsResponseBodyResultData {
	s.OwnerNum = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetPublishStatus(v int32) *ListRecentViewReportsResponseBodyResultData {
	s.PublishStatus = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetTreeId(v string) *ListRecentViewReportsResponseBodyResultData {
	s.TreeId = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetType(v string) *ListRecentViewReportsResponseBodyResultData {
	s.Type = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetViewCount(v int64) *ListRecentViewReportsResponseBodyResultData {
	s.ViewCount = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetWorkspaceId(v string) *ListRecentViewReportsResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetWorkspaceName(v string) *ListRecentViewReportsResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type ListRecentViewReportsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecentViewReportsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecentViewReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecentViewReportsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponse) SetHeaders(v map[string]*string) *ListRecentViewReportsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentViewReportsResponse) SetStatusCode(v int32) *ListRecentViewReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentViewReportsResponse) SetBody(v *ListRecentViewReportsResponseBody) *ListRecentViewReportsResponse {
	s.Body = v
	return s
}

type ListSharedReportsRequest struct {
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TreeType *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSharedReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSharedReportsRequest) GoString() string {
	return s.String()
}

func (s *ListSharedReportsRequest) SetKeyword(v string) *ListSharedReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListSharedReportsRequest) SetPageSize(v int32) *ListSharedReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSharedReportsRequest) SetTreeType(v string) *ListSharedReportsRequest {
	s.TreeType = &v
	return s
}

func (s *ListSharedReportsRequest) SetUserId(v string) *ListSharedReportsRequest {
	s.UserId = &v
	return s
}

type ListSharedReportsResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListSharedReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSharedReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSharedReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSharedReportsResponseBody) SetRequestId(v string) *ListSharedReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSharedReportsResponseBody) SetResult(v *ListSharedReportsResponseBodyResult) *ListSharedReportsResponseBody {
	s.Result = v
	return s
}

func (s *ListSharedReportsResponseBody) SetSuccess(v bool) *ListSharedReportsResponseBody {
	s.Success = &v
	return s
}

type ListSharedReportsResponseBodyResult struct {
	Data       []*ListSharedReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                     `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                     `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListSharedReportsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSharedReportsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSharedReportsResponseBodyResult) SetData(v []*ListSharedReportsResponseBodyResultData) *ListSharedReportsResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListSharedReportsResponseBodyResult) SetPageNum(v int32) *ListSharedReportsResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListSharedReportsResponseBodyResult) SetPageSize(v int32) *ListSharedReportsResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListSharedReportsResponseBodyResult) SetTotalNum(v int32) *ListSharedReportsResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListSharedReportsResponseBodyResult) SetTotalPages(v int32) *ListSharedReportsResponseBodyResult {
	s.TotalPages = &v
	return s
}

type ListSharedReportsResponseBodyResultData struct {
	Favorite      *bool   `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	GmtCreate     *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HasEditAuth   *bool   `json:"HasEditAuth,omitempty" xml:"HasEditAuth,omitempty"`
	HasViewAuth   *bool   `json:"HasViewAuth,omitempty" xml:"HasViewAuth,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName     *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerNum      *string `json:"OwnerNum,omitempty" xml:"OwnerNum,omitempty"`
	PublishStatus *int32  `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	TreeId        *string `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListSharedReportsResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListSharedReportsResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListSharedReportsResponseBodyResultData) SetFavorite(v bool) *ListSharedReportsResponseBodyResultData {
	s.Favorite = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetGmtCreate(v string) *ListSharedReportsResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetGmtModified(v string) *ListSharedReportsResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetHasEditAuth(v bool) *ListSharedReportsResponseBodyResultData {
	s.HasEditAuth = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetHasViewAuth(v bool) *ListSharedReportsResponseBodyResultData {
	s.HasViewAuth = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetName(v string) *ListSharedReportsResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetOwnerName(v string) *ListSharedReportsResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetOwnerNum(v string) *ListSharedReportsResponseBodyResultData {
	s.OwnerNum = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetPublishStatus(v int32) *ListSharedReportsResponseBodyResultData {
	s.PublishStatus = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetTreeId(v string) *ListSharedReportsResponseBodyResultData {
	s.TreeId = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetType(v string) *ListSharedReportsResponseBodyResultData {
	s.Type = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetWorkspaceId(v string) *ListSharedReportsResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetWorkspaceName(v string) *ListSharedReportsResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type ListSharedReportsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSharedReportsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSharedReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSharedReportsResponse) GoString() string {
	return s.String()
}

func (s *ListSharedReportsResponse) SetHeaders(v map[string]*string) *ListSharedReportsResponse {
	s.Headers = v
	return s
}

func (s *ListSharedReportsResponse) SetStatusCode(v int32) *ListSharedReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSharedReportsResponse) SetBody(v *ListSharedReportsResponseBody) *ListSharedReportsResponse {
	s.Body = v
	return s
}

type ListUserGroupsByUserIdRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserGroupsByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsByUserIdRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsByUserIdRequest) SetUserId(v string) *ListUserGroupsByUserIdRequest {
	s.UserId = &v
	return s
}

type ListUserGroupsByUserIdResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListUserGroupsByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserGroupsByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsByUserIdResponseBody) SetRequestId(v string) *ListUserGroupsByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBody) SetResult(v []*ListUserGroupsByUserIdResponseBodyResult) *ListUserGroupsByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *ListUserGroupsByUserIdResponseBody) SetSuccess(v bool) *ListUserGroupsByUserIdResponseBody {
	s.Success = &v
	return s
}

type ListUserGroupsByUserIdResponseBodyResult struct {
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser        *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	IdentifiedPath    *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	ModifiedTime      *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModifyUser        *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	ParentUsergroupId *string `json:"ParentUsergroupId,omitempty" xml:"ParentUsergroupId,omitempty"`
	UsergroupDesc     *string `json:"UsergroupDesc,omitempty" xml:"UsergroupDesc,omitempty"`
	UsergroupId       *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	UsergroupName     *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
}

func (s ListUserGroupsByUserIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetCreateTime(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetCreateUser(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.CreateUser = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetIdentifiedPath(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.IdentifiedPath = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetModifiedTime(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetModifyUser(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetParentUsergroupId(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.ParentUsergroupId = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetUsergroupDesc(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.UsergroupDesc = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetUsergroupId(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.UsergroupId = &v
	return s
}

func (s *ListUserGroupsByUserIdResponseBodyResult) SetUsergroupName(v string) *ListUserGroupsByUserIdResponseBodyResult {
	s.UsergroupName = &v
	return s
}

type ListUserGroupsByUserIdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserGroupsByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserGroupsByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsByUserIdResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsByUserIdResponse) SetHeaders(v map[string]*string) *ListUserGroupsByUserIdResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsByUserIdResponse) SetStatusCode(v int32) *ListUserGroupsByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsByUserIdResponse) SetBody(v *ListUserGroupsByUserIdResponseBody) *ListUserGroupsByUserIdResponse {
	s.Body = v
	return s
}

type QueryDataServiceRequest struct {
	ApiId        *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	Conditions   *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	ReturnFields *string `json:"ReturnFields,omitempty" xml:"ReturnFields,omitempty"`
}

func (s QueryDataServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceRequest) GoString() string {
	return s.String()
}

func (s *QueryDataServiceRequest) SetApiId(v string) *QueryDataServiceRequest {
	s.ApiId = &v
	return s
}

func (s *QueryDataServiceRequest) SetConditions(v string) *QueryDataServiceRequest {
	s.Conditions = &v
	return s
}

func (s *QueryDataServiceRequest) SetReturnFields(v string) *QueryDataServiceRequest {
	s.ReturnFields = &v
	return s
}

type QueryDataServiceResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDataServiceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataServiceResponseBody) SetRequestId(v string) *QueryDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataServiceResponseBody) SetResult(v *QueryDataServiceResponseBodyResult) *QueryDataServiceResponseBody {
	s.Result = v
	return s
}

func (s *QueryDataServiceResponseBody) SetSuccess(v bool) *QueryDataServiceResponseBody {
	s.Success = &v
	return s
}

type QueryDataServiceResponseBodyResult struct {
	Headers []*QueryDataServiceResponseBodyResultHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	Sql     *string                                      `json:"Sql,omitempty" xml:"Sql,omitempty"`
	Values  []map[string]interface{}                     `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryDataServiceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDataServiceResponseBodyResult) SetHeaders(v []*QueryDataServiceResponseBodyResultHeaders) *QueryDataServiceResponseBodyResult {
	s.Headers = v
	return s
}

func (s *QueryDataServiceResponseBodyResult) SetSql(v string) *QueryDataServiceResponseBodyResult {
	s.Sql = &v
	return s
}

func (s *QueryDataServiceResponseBodyResult) SetValues(v []map[string]interface{}) *QueryDataServiceResponseBodyResult {
	s.Values = v
	return s
}

type QueryDataServiceResponseBodyResultHeaders struct {
	Aggregator  *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	Column      *string `json:"Column,omitempty" xml:"Column,omitempty"`
	DataType    *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataServiceResponseBodyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceResponseBodyResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetAggregator(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Aggregator = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetColumn(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Column = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetDataType(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.DataType = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetGranularity(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Granularity = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetLabel(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Label = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetType(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Type = &v
	return s
}

type QueryDataServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDataServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryDataServiceResponse) SetHeaders(v map[string]*string) *QueryDataServiceResponse {
	s.Headers = v
	return s
}

func (s *QueryDataServiceResponse) SetStatusCode(v int32) *QueryDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataServiceResponse) SetBody(v *QueryDataServiceResponseBody) *QueryDataServiceResponse {
	s.Body = v
	return s
}

type QueryDatasetDetailInfoRequest struct {
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
}

func (s QueryDatasetDetailInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetDetailInfoRequest) SetDatasetId(v string) *QueryDatasetDetailInfoRequest {
	s.DatasetId = &v
	return s
}

type QueryDatasetDetailInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDatasetDetailInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetDetailInfoResponseBody) SetRequestId(v string) *QueryDatasetDetailInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetDetailInfoResponseBody) SetResult(v string) *QueryDatasetDetailInfoResponseBody {
	s.Result = &v
	return s
}

func (s *QueryDatasetDetailInfoResponseBody) SetSuccess(v bool) *QueryDatasetDetailInfoResponseBody {
	s.Success = &v
	return s
}

type QueryDatasetDetailInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDatasetDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDatasetDetailInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetDetailInfoResponse) SetHeaders(v map[string]*string) *QueryDatasetDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetDetailInfoResponse) SetStatusCode(v int32) *QueryDatasetDetailInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetDetailInfoResponse) SetBody(v *QueryDatasetDetailInfoResponseBody) *QueryDatasetDetailInfoResponse {
	s.Body = v
	return s
}

type QueryDatasetInfoRequest struct {
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
}

func (s QueryDatasetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoRequest) SetDatasetId(v string) *QueryDatasetInfoRequest {
	s.DatasetId = &v
	return s
}

type QueryDatasetInfoResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDatasetInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDatasetInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBody) SetRequestId(v string) *QueryDatasetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetInfoResponseBody) SetResult(v *QueryDatasetInfoResponseBodyResult) *QueryDatasetInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryDatasetInfoResponseBody) SetSuccess(v bool) *QueryDatasetInfoResponseBody {
	s.Success = &v
	return s
}

type QueryDatasetInfoResponseBodyResult struct {
	CubeTableList []*QueryDatasetInfoResponseBodyResultCubeTableList `json:"CubeTableList,omitempty" xml:"CubeTableList,omitempty" type:"Repeated"`
	CustimzeSql   *bool                                              `json:"CustimzeSql,omitempty" xml:"CustimzeSql,omitempty"`
	DatasetId     *string                                            `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetName   *string                                            `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	DimensionList []*QueryDatasetInfoResponseBodyResultDimensionList `json:"DimensionList,omitempty" xml:"DimensionList,omitempty" type:"Repeated"`
	Directory     *QueryDatasetInfoResponseBodyResultDirectory       `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	DsId          *string                                            `json:"DsId,omitempty" xml:"DsId,omitempty"`
	DsName        *string                                            `json:"DsName,omitempty" xml:"DsName,omitempty"`
	DsType        *string                                            `json:"DsType,omitempty" xml:"DsType,omitempty"`
	GmtCreate     *string                                            `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify     *string                                            `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	MeasureList   []*QueryDatasetInfoResponseBodyResultMeasureList   `json:"MeasureList,omitempty" xml:"MeasureList,omitempty" type:"Repeated"`
	OwnerId       *string                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName     *string                                            `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	RowLevel      *bool                                              `json:"RowLevel,omitempty" xml:"RowLevel,omitempty"`
	WorkspaceId   *string                                            `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string                                            `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResult) SetCubeTableList(v []*QueryDatasetInfoResponseBodyResultCubeTableList) *QueryDatasetInfoResponseBodyResult {
	s.CubeTableList = v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetCustimzeSql(v bool) *QueryDatasetInfoResponseBodyResult {
	s.CustimzeSql = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDatasetId(v string) *QueryDatasetInfoResponseBodyResult {
	s.DatasetId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDatasetName(v string) *QueryDatasetInfoResponseBodyResult {
	s.DatasetName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDimensionList(v []*QueryDatasetInfoResponseBodyResultDimensionList) *QueryDatasetInfoResponseBodyResult {
	s.DimensionList = v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDirectory(v *QueryDatasetInfoResponseBodyResultDirectory) *QueryDatasetInfoResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDsId(v string) *QueryDatasetInfoResponseBodyResult {
	s.DsId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDsName(v string) *QueryDatasetInfoResponseBodyResult {
	s.DsName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetDsType(v string) *QueryDatasetInfoResponseBodyResult {
	s.DsType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetGmtCreate(v string) *QueryDatasetInfoResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetGmtModify(v string) *QueryDatasetInfoResponseBodyResult {
	s.GmtModify = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetMeasureList(v []*QueryDatasetInfoResponseBodyResultMeasureList) *QueryDatasetInfoResponseBodyResult {
	s.MeasureList = v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetOwnerId(v string) *QueryDatasetInfoResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetOwnerName(v string) *QueryDatasetInfoResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetRowLevel(v bool) *QueryDatasetInfoResponseBodyResult {
	s.RowLevel = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetWorkspaceId(v string) *QueryDatasetInfoResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResult) SetWorkspaceName(v string) *QueryDatasetInfoResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

type QueryDatasetInfoResponseBodyResultCubeTableList struct {
	Caption      *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	Customsql    *bool   `json:"Customsql,omitempty" xml:"Customsql,omitempty"`
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	DsType       *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	FactTable    *bool   `json:"FactTable,omitempty" xml:"FactTable,omitempty"`
	Sql          *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	UniqueId     *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResultCubeTableList) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResultCubeTableList) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetCaption(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.Caption = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetCustomsql(v bool) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.Customsql = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetDatasourceId(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.DatasourceId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetDsType(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.DsType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetFactTable(v bool) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.FactTable = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetSql(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.Sql = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetTableName(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.TableName = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultCubeTableList) SetUniqueId(v string) *QueryDatasetInfoResponseBodyResultCubeTableList {
	s.UniqueId = &v
	return s
}

type QueryDatasetInfoResponseBodyResultDimensionList struct {
	Caption       *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	DataType      *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	Expression    *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	FactColumn    *string `json:"FactColumn,omitempty" xml:"FactColumn,omitempty"`
	Granularity   *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	RefUid        *string `json:"RefUid,omitempty" xml:"RefUid,omitempty"`
	TableUniqueId *string `json:"TableUniqueId,omitempty" xml:"TableUniqueId,omitempty"`
	Uid           *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResultDimensionList) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResultDimensionList) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetCaption(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.Caption = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetDataType(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.DataType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetDimensionType(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.DimensionType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetExpression(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.Expression = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetFactColumn(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.FactColumn = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetGranularity(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.Granularity = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetRefUid(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.RefUid = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetTableUniqueId(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.TableUniqueId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetUid(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.Uid = &v
	return s
}

type QueryDatasetInfoResponseBodyResultDirectory struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PathId   *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResultDirectory) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) SetId(v string) *QueryDatasetInfoResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) SetName(v string) *QueryDatasetInfoResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) SetPathId(v string) *QueryDatasetInfoResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultDirectory) SetPathName(v string) *QueryDatasetInfoResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

type QueryDatasetInfoResponseBodyResultMeasureList struct {
	Caption       *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	DataType      *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Expression    *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	FactColumn    *string `json:"FactColumn,omitempty" xml:"FactColumn,omitempty"`
	MeasureType   *string `json:"MeasureType,omitempty" xml:"MeasureType,omitempty"`
	TableUniqueId *string `json:"TableUniqueId,omitempty" xml:"TableUniqueId,omitempty"`
	Uid           *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryDatasetInfoResponseBodyResultMeasureList) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetInfoResponseBodyResultMeasureList) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetCaption(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.Caption = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetDataType(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.DataType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetExpression(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.Expression = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetFactColumn(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.FactColumn = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetMeasureType(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.MeasureType = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetTableUniqueId(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.TableUniqueId = &v
	return s
}

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetUid(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.Uid = &v
	return s
}

type QueryDatasetInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDatasetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDatasetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetInfoResponse) SetHeaders(v map[string]*string) *QueryDatasetInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetInfoResponse) SetStatusCode(v int32) *QueryDatasetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetInfoResponse) SetBody(v *QueryDatasetInfoResponseBody) *QueryDatasetInfoResponse {
	s.Body = v
	return s
}

type QueryDatasetListRequest struct {
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	WithChildren *bool   `json:"WithChildren,omitempty" xml:"WithChildren,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryDatasetListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetListRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetListRequest) SetDirectoryId(v string) *QueryDatasetListRequest {
	s.DirectoryId = &v
	return s
}

func (s *QueryDatasetListRequest) SetKeyword(v string) *QueryDatasetListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryDatasetListRequest) SetPageNum(v int32) *QueryDatasetListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryDatasetListRequest) SetPageSize(v int32) *QueryDatasetListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDatasetListRequest) SetWithChildren(v bool) *QueryDatasetListRequest {
	s.WithChildren = &v
	return s
}

func (s *QueryDatasetListRequest) SetWorkspaceId(v string) *QueryDatasetListRequest {
	s.WorkspaceId = &v
	return s
}

type QueryDatasetListResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDatasetListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDatasetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBody) SetRequestId(v string) *QueryDatasetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetListResponseBody) SetResult(v *QueryDatasetListResponseBodyResult) *QueryDatasetListResponseBody {
	s.Result = v
	return s
}

func (s *QueryDatasetListResponseBody) SetSuccess(v bool) *QueryDatasetListResponseBody {
	s.Success = &v
	return s
}

type QueryDatasetListResponseBodyResult struct {
	Data       []*QueryDatasetListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                    `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                    `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryDatasetListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBodyResult) SetData(v []*QueryDatasetListResponseBodyResultData) *QueryDatasetListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryDatasetListResponseBodyResult) SetPageNum(v int32) *QueryDatasetListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryDatasetListResponseBodyResult) SetPageSize(v int32) *QueryDatasetListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryDatasetListResponseBodyResult) SetTotalNum(v int32) *QueryDatasetListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryDatasetListResponseBodyResult) SetTotalPages(v int32) *QueryDatasetListResponseBodyResult {
	s.TotalPages = &v
	return s
}

type QueryDatasetListResponseBodyResultData struct {
	CreateTime    *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataSource    *QueryDatasetListResponseBodyResultDataDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	DatasetId     *string                                           `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetName   *string                                           `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Description   *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory     *QueryDatasetListResponseBodyResultDataDirectory  `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	ModifyTime    *string                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OwnerId       *string                                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName     *string                                           `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	RowLevel      *bool                                             `json:"RowLevel,omitempty" xml:"RowLevel,omitempty"`
	WorkspaceId   *string                                           `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string                                           `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryDatasetListResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBodyResultData) SetCreateTime(v string) *QueryDatasetListResponseBodyResultData {
	s.CreateTime = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDataSource(v *QueryDatasetListResponseBodyResultDataDataSource) *QueryDatasetListResponseBodyResultData {
	s.DataSource = v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDatasetId(v string) *QueryDatasetListResponseBodyResultData {
	s.DatasetId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDatasetName(v string) *QueryDatasetListResponseBodyResultData {
	s.DatasetName = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDescription(v string) *QueryDatasetListResponseBodyResultData {
	s.Description = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetDirectory(v *QueryDatasetListResponseBodyResultDataDirectory) *QueryDatasetListResponseBodyResultData {
	s.Directory = v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetModifyTime(v string) *QueryDatasetListResponseBodyResultData {
	s.ModifyTime = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetOwnerId(v string) *QueryDatasetListResponseBodyResultData {
	s.OwnerId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetOwnerName(v string) *QueryDatasetListResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetRowLevel(v bool) *QueryDatasetListResponseBodyResultData {
	s.RowLevel = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetWorkspaceId(v string) *QueryDatasetListResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultData) SetWorkspaceName(v string) *QueryDatasetListResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type QueryDatasetListResponseBodyResultDataDataSource struct {
	DsId   *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	DsName *string `json:"DsName,omitempty" xml:"DsName,omitempty"`
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
}

func (s QueryDatasetListResponseBodyResultDataDataSource) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetListResponseBodyResultDataDataSource) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) SetDsId(v string) *QueryDatasetListResponseBodyResultDataDataSource {
	s.DsId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) SetDsName(v string) *QueryDatasetListResponseBodyResultDataDataSource {
	s.DsName = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDataSource) SetDsType(v string) *QueryDatasetListResponseBodyResultDataDataSource {
	s.DsType = &v
	return s
}

type QueryDatasetListResponseBodyResultDataDirectory struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PathId   *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryDatasetListResponseBodyResultDataDirectory) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetListResponseBodyResultDataDirectory) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) SetId(v string) *QueryDatasetListResponseBodyResultDataDirectory {
	s.Id = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) SetName(v string) *QueryDatasetListResponseBodyResultDataDirectory {
	s.Name = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) SetPathId(v string) *QueryDatasetListResponseBodyResultDataDirectory {
	s.PathId = &v
	return s
}

func (s *QueryDatasetListResponseBodyResultDataDirectory) SetPathName(v string) *QueryDatasetListResponseBodyResultDataDirectory {
	s.PathName = &v
	return s
}

type QueryDatasetListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDatasetListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDatasetListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetListResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetListResponse) SetHeaders(v map[string]*string) *QueryDatasetListResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetListResponse) SetStatusCode(v int32) *QueryDatasetListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetListResponse) SetBody(v *QueryDatasetListResponseBody) *QueryDatasetListResponse {
	s.Body = v
	return s
}

type QueryDatasetSwitchInfoRequest struct {
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
}

func (s QueryDatasetSwitchInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetSwitchInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDatasetSwitchInfoRequest) SetCubeId(v string) *QueryDatasetSwitchInfoRequest {
	s.CubeId = &v
	return s
}

type QueryDatasetSwitchInfoResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDatasetSwitchInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDatasetSwitchInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetSwitchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetSwitchInfoResponseBody) SetRequestId(v string) *QueryDatasetSwitchInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBody) SetResult(v *QueryDatasetSwitchInfoResponseBodyResult) *QueryDatasetSwitchInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBody) SetSuccess(v bool) *QueryDatasetSwitchInfoResponseBody {
	s.Success = &v
	return s
}

type QueryDatasetSwitchInfoResponseBodyResult struct {
	CubeId                      *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	IsOpenColumnLevelPermission *int32  `json:"IsOpenColumnLevelPermission,omitempty" xml:"IsOpenColumnLevelPermission,omitempty"`
	IsOpenRowLevelPermission    *int32  `json:"IsOpenRowLevelPermission,omitempty" xml:"IsOpenRowLevelPermission,omitempty"`
}

func (s QueryDatasetSwitchInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetSwitchInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) SetCubeId(v string) *QueryDatasetSwitchInfoResponseBodyResult {
	s.CubeId = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) SetIsOpenColumnLevelPermission(v int32) *QueryDatasetSwitchInfoResponseBodyResult {
	s.IsOpenColumnLevelPermission = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponseBodyResult) SetIsOpenRowLevelPermission(v int32) *QueryDatasetSwitchInfoResponseBodyResult {
	s.IsOpenRowLevelPermission = &v
	return s
}

type QueryDatasetSwitchInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDatasetSwitchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDatasetSwitchInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDatasetSwitchInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetSwitchInfoResponse) SetHeaders(v map[string]*string) *QueryDatasetSwitchInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetSwitchInfoResponse) SetStatusCode(v int32) *QueryDatasetSwitchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetSwitchInfoResponse) SetBody(v *QueryDatasetSwitchInfoResponseBody) *QueryDatasetSwitchInfoResponse {
	s.Body = v
	return s
}

type QueryEmbeddedInfoResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryEmbeddedInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEmbeddedInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEmbeddedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedInfoResponseBody) SetRequestId(v string) *QueryEmbeddedInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBody) SetResult(v *QueryEmbeddedInfoResponseBodyResult) *QueryEmbeddedInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryEmbeddedInfoResponseBody) SetSuccess(v bool) *QueryEmbeddedInfoResponseBody {
	s.Success = &v
	return s
}

type QueryEmbeddedInfoResponseBodyResult struct {
	Detail        *QueryEmbeddedInfoResponseBodyResultDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	EmbeddedCount *int32                                     `json:"EmbeddedCount,omitempty" xml:"EmbeddedCount,omitempty"`
	MaxCount      *int32                                     `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s QueryEmbeddedInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryEmbeddedInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedInfoResponseBodyResult) SetDetail(v *QueryEmbeddedInfoResponseBodyResultDetail) *QueryEmbeddedInfoResponseBodyResult {
	s.Detail = v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResult) SetEmbeddedCount(v int32) *QueryEmbeddedInfoResponseBodyResult {
	s.EmbeddedCount = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResult) SetMaxCount(v int32) *QueryEmbeddedInfoResponseBodyResult {
	s.MaxCount = &v
	return s
}

type QueryEmbeddedInfoResponseBodyResultDetail struct {
	DashboardOfflineQuery *int32 `json:"DashboardOfflineQuery,omitempty" xml:"DashboardOfflineQuery,omitempty"`
	Page                  *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	Report                *int32 `json:"Report,omitempty" xml:"Report,omitempty"`
}

func (s QueryEmbeddedInfoResponseBodyResultDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryEmbeddedInfoResponseBodyResultDetail) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) SetDashboardOfflineQuery(v int32) *QueryEmbeddedInfoResponseBodyResultDetail {
	s.DashboardOfflineQuery = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) SetPage(v int32) *QueryEmbeddedInfoResponseBodyResultDetail {
	s.Page = &v
	return s
}

func (s *QueryEmbeddedInfoResponseBodyResultDetail) SetReport(v int32) *QueryEmbeddedInfoResponseBodyResultDetail {
	s.Report = &v
	return s
}

type QueryEmbeddedInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEmbeddedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEmbeddedInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEmbeddedInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedInfoResponse) SetHeaders(v map[string]*string) *QueryEmbeddedInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryEmbeddedInfoResponse) SetStatusCode(v int32) *QueryEmbeddedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmbeddedInfoResponse) SetBody(v *QueryEmbeddedInfoResponseBody) *QueryEmbeddedInfoResponse {
	s.Body = v
	return s
}

type QueryEmbeddedStatusRequest struct {
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s QueryEmbeddedStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEmbeddedStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedStatusRequest) SetWorksId(v string) *QueryEmbeddedStatusRequest {
	s.WorksId = &v
	return s
}

type QueryEmbeddedStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEmbeddedStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEmbeddedStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedStatusResponseBody) SetRequestId(v string) *QueryEmbeddedStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEmbeddedStatusResponseBody) SetResult(v bool) *QueryEmbeddedStatusResponseBody {
	s.Result = &v
	return s
}

func (s *QueryEmbeddedStatusResponseBody) SetSuccess(v bool) *QueryEmbeddedStatusResponseBody {
	s.Success = &v
	return s
}

type QueryEmbeddedStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEmbeddedStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEmbeddedStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEmbeddedStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedStatusResponse) SetHeaders(v map[string]*string) *QueryEmbeddedStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryEmbeddedStatusResponse) SetStatusCode(v int32) *QueryEmbeddedStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmbeddedStatusResponse) SetBody(v *QueryEmbeddedStatusResponseBody) *QueryEmbeddedStatusResponse {
	s.Body = v
	return s
}

type QueryOrganizationWorkspaceListRequest struct {
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryOrganizationWorkspaceListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationWorkspaceListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListRequest) SetKeyword(v string) *QueryOrganizationWorkspaceListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryOrganizationWorkspaceListRequest) SetPageNum(v int32) *QueryOrganizationWorkspaceListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryOrganizationWorkspaceListRequest) SetPageSize(v int32) *QueryOrganizationWorkspaceListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrganizationWorkspaceListRequest) SetUserId(v string) *QueryOrganizationWorkspaceListRequest {
	s.UserId = &v
	return s
}

type QueryOrganizationWorkspaceListResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryOrganizationWorkspaceListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOrganizationWorkspaceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationWorkspaceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListResponseBody) SetRequestId(v string) *QueryOrganizationWorkspaceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBody) SetResult(v *QueryOrganizationWorkspaceListResponseBodyResult) *QueryOrganizationWorkspaceListResponseBody {
	s.Result = v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBody) SetSuccess(v bool) *QueryOrganizationWorkspaceListResponseBody {
	s.Success = &v
	return s
}

type QueryOrganizationWorkspaceListResponseBodyResult struct {
	Data       []*QueryOrganizationWorkspaceListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                                  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                                  `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                                  `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryOrganizationWorkspaceListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationWorkspaceListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetData(v []*QueryOrganizationWorkspaceListResponseBodyResultData) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetPageNum(v int32) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetPageSize(v int32) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetTotalNum(v int32) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResult) SetTotalPages(v int32) *QueryOrganizationWorkspaceListResponseBodyResult {
	s.TotalPages = &v
	return s
}

type QueryOrganizationWorkspaceListResponseBodyResultData struct {
	AllowPublishOperation *bool   `json:"AllowPublishOperation,omitempty" xml:"AllowPublishOperation,omitempty"`
	AllowShareOperation   *bool   `json:"AllowShareOperation,omitempty" xml:"AllowShareOperation,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser            *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	CreateUserAccountName *string `json:"CreateUserAccountName,omitempty" xml:"CreateUserAccountName,omitempty"`
	ModifiedTime          *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModifyUser            *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	ModifyUserAccountName *string `json:"ModifyUserAccountName,omitempty" xml:"ModifyUserAccountName,omitempty"`
	OrganizationId        *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Owner                 *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerAccountName      *string `json:"OwnerAccountName,omitempty" xml:"OwnerAccountName,omitempty"`
	WorkspaceDescription  *string `json:"WorkspaceDescription,omitempty" xml:"WorkspaceDescription,omitempty"`
	WorkspaceId           *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName         *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryOrganizationWorkspaceListResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationWorkspaceListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetAllowPublishOperation(v bool) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.AllowPublishOperation = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetAllowShareOperation(v bool) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.AllowShareOperation = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetCreateTime(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.CreateTime = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetCreateUser(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.CreateUser = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetCreateUserAccountName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.CreateUserAccountName = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetModifiedTime(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.ModifiedTime = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetModifyUser(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.ModifyUser = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetModifyUserAccountName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.ModifyUserAccountName = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetOrganizationId(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.OrganizationId = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetOwner(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.Owner = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetOwnerAccountName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.OwnerAccountName = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetWorkspaceDescription(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.WorkspaceDescription = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetWorkspaceId(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponseBodyResultData) SetWorkspaceName(v string) *QueryOrganizationWorkspaceListResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type QueryOrganizationWorkspaceListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOrganizationWorkspaceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrganizationWorkspaceListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationWorkspaceListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListResponse) SetHeaders(v map[string]*string) *QueryOrganizationWorkspaceListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrganizationWorkspaceListResponse) SetStatusCode(v int32) *QueryOrganizationWorkspaceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrganizationWorkspaceListResponse) SetBody(v *QueryOrganizationWorkspaceListResponseBody) *QueryOrganizationWorkspaceListResponse {
	s.Body = v
	return s
}

type QueryReadableResourcesListByUserIdRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryReadableResourcesListByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdRequest) SetUserId(v string) *QueryReadableResourcesListByUserIdRequest {
	s.UserId = &v
	return s
}

type QueryReadableResourcesListByUserIdResponseBody struct {
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryReadableResourcesListByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryReadableResourcesListByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdResponseBody) SetRequestId(v string) *QueryReadableResourcesListByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBody) SetResult(v []*QueryReadableResourcesListByUserIdResponseBodyResult) *QueryReadableResourcesListByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBody) SetSuccess(v bool) *QueryReadableResourcesListByUserIdResponseBody {
	s.Success = &v
	return s
}

type QueryReadableResourcesListByUserIdResponseBodyResult struct {
	CreateTime        *string                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory         *QueryReadableResourcesListByUserIdResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	ModifyName        *string                                                        `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	ModifyTime        *string                                                        `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OwnerId           *string                                                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName         *string                                                        `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	SecurityLevel     *string                                                        `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	Status            *int32                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	ThirdPartAuthFlag *int32                                                         `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	WorkName          *string                                                        `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	WorkType          *string                                                        `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	WorksId           *string                                                        `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	WorkspaceId       *string                                                        `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName     *string                                                        `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryReadableResourcesListByUserIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetCreateTime(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetDescription(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetDirectory(v *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetModifyName(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.ModifyName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetModifyTime(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetOwnerId(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetOwnerName(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetSecurityLevel(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.SecurityLevel = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetStatus(v int32) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetThirdPartAuthFlag(v int32) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorkName(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorkName = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorkType(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorkType = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorksId(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorkspaceId(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResult) SetWorkspaceName(v string) *QueryReadableResourcesListByUserIdResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

type QueryReadableResourcesListByUserIdResponseBodyResultDirectory struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PathId   *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryReadableResourcesListByUserIdResponseBodyResultDirectory) String() string {
	return tea.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) SetId(v string) *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) SetName(v string) *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) SetPathId(v string) *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponseBodyResultDirectory) SetPathName(v string) *QueryReadableResourcesListByUserIdResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

type QueryReadableResourcesListByUserIdResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryReadableResourcesListByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryReadableResourcesListByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdResponse) SetHeaders(v map[string]*string) *QueryReadableResourcesListByUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponse) SetStatusCode(v int32) *QueryReadableResourcesListByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponse) SetBody(v *QueryReadableResourcesListByUserIdResponseBody) *QueryReadableResourcesListByUserIdResponse {
	s.Body = v
	return s
}

type QueryShareListRequest struct {
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s QueryShareListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryShareListRequest) GoString() string {
	return s.String()
}

func (s *QueryShareListRequest) SetReportId(v string) *QueryShareListRequest {
	s.ReportId = &v
	return s
}

type QueryShareListResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryShareListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryShareListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryShareListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryShareListResponseBody) SetRequestId(v string) *QueryShareListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryShareListResponseBody) SetResult(v []*QueryShareListResponseBodyResult) *QueryShareListResponseBody {
	s.Result = v
	return s
}

func (s *QueryShareListResponseBody) SetSuccess(v bool) *QueryShareListResponseBody {
	s.Success = &v
	return s
}

type QueryShareListResponseBodyResult struct {
	AuthPoint   *int32  `json:"AuthPoint,omitempty" xml:"AuthPoint,omitempty"`
	ExpireDate  *int64  `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	ReportId    *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ShareId     *string `json:"ShareId,omitempty" xml:"ShareId,omitempty"`
	ShareToId   *string `json:"ShareToId,omitempty" xml:"ShareToId,omitempty"`
	ShareToName *string `json:"ShareToName,omitempty" xml:"ShareToName,omitempty"`
	ShareToType *int32  `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
	ShareType   *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s QueryShareListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryShareListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryShareListResponseBodyResult) SetAuthPoint(v int32) *QueryShareListResponseBodyResult {
	s.AuthPoint = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetExpireDate(v int64) *QueryShareListResponseBodyResult {
	s.ExpireDate = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetReportId(v string) *QueryShareListResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareId(v string) *QueryShareListResponseBodyResult {
	s.ShareId = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareToId(v string) *QueryShareListResponseBodyResult {
	s.ShareToId = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareToName(v string) *QueryShareListResponseBodyResult {
	s.ShareToName = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareToType(v int32) *QueryShareListResponseBodyResult {
	s.ShareToType = &v
	return s
}

func (s *QueryShareListResponseBodyResult) SetShareType(v string) *QueryShareListResponseBodyResult {
	s.ShareType = &v
	return s
}

type QueryShareListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryShareListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryShareListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryShareListResponse) GoString() string {
	return s.String()
}

func (s *QueryShareListResponse) SetHeaders(v map[string]*string) *QueryShareListResponse {
	s.Headers = v
	return s
}

func (s *QueryShareListResponse) SetStatusCode(v int32) *QueryShareListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryShareListResponse) SetBody(v *QueryShareListResponseBody) *QueryShareListResponse {
	s.Body = v
	return s
}

type QuerySharesToUserListRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySharesToUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySharesToUserListRequest) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListRequest) SetUserId(v string) *QuerySharesToUserListRequest {
	s.UserId = &v
	return s
}

type QuerySharesToUserListResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QuerySharesToUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySharesToUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySharesToUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListResponseBody) SetRequestId(v string) *QuerySharesToUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySharesToUserListResponseBody) SetResult(v []*QuerySharesToUserListResponseBodyResult) *QuerySharesToUserListResponseBody {
	s.Result = v
	return s
}

func (s *QuerySharesToUserListResponseBody) SetSuccess(v bool) *QuerySharesToUserListResponseBody {
	s.Success = &v
	return s
}

type QuerySharesToUserListResponseBodyResult struct {
	CreateTime        *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory         *QuerySharesToUserListResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	ModifyName        *string                                           `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	ModifyTime        *string                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OwnerId           *string                                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName         *string                                           `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	SecurityLevel     *string                                           `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	Status            *int32                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	ThirdPartAuthFlag *int32                                            `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	WorkName          *string                                           `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	WorkType          *string                                           `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	WorksId           *string                                           `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	WorkspaceId       *string                                           `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName     *string                                           `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QuerySharesToUserListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySharesToUserListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListResponseBodyResult) SetCreateTime(v string) *QuerySharesToUserListResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetDescription(v string) *QuerySharesToUserListResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetDirectory(v *QuerySharesToUserListResponseBodyResultDirectory) *QuerySharesToUserListResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetModifyName(v string) *QuerySharesToUserListResponseBodyResult {
	s.ModifyName = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetModifyTime(v string) *QuerySharesToUserListResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetOwnerId(v string) *QuerySharesToUserListResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetOwnerName(v string) *QuerySharesToUserListResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetSecurityLevel(v string) *QuerySharesToUserListResponseBodyResult {
	s.SecurityLevel = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetStatus(v int32) *QuerySharesToUserListResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetThirdPartAuthFlag(v int32) *QuerySharesToUserListResponseBodyResult {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorkName(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorkName = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorkType(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorkType = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorksId(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorkspaceId(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResult) SetWorkspaceName(v string) *QuerySharesToUserListResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

type QuerySharesToUserListResponseBodyResultDirectory struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PathId   *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QuerySharesToUserListResponseBodyResultDirectory) String() string {
	return tea.Prettify(s)
}

func (s QuerySharesToUserListResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) SetId(v string) *QuerySharesToUserListResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) SetName(v string) *QuerySharesToUserListResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) SetPathId(v string) *QuerySharesToUserListResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QuerySharesToUserListResponseBodyResultDirectory) SetPathName(v string) *QuerySharesToUserListResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

type QuerySharesToUserListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySharesToUserListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySharesToUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySharesToUserListResponse) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListResponse) SetHeaders(v map[string]*string) *QuerySharesToUserListResponse {
	s.Headers = v
	return s
}

func (s *QuerySharesToUserListResponse) SetStatusCode(v int32) *QuerySharesToUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySharesToUserListResponse) SetBody(v *QuerySharesToUserListResponseBody) *QuerySharesToUserListResponse {
	s.Body = v
	return s
}

type QueryTicketInfoRequest struct {
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s QueryTicketInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketInfoRequest) SetTicket(v string) *QueryTicketInfoRequest {
	s.Ticket = &v
	return s
}

type QueryTicketInfoResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryTicketInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketInfoResponseBody) SetRequestId(v string) *QueryTicketInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketInfoResponseBody) SetResult(v *QueryTicketInfoResponseBodyResult) *QueryTicketInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryTicketInfoResponseBody) SetSuccess(v bool) *QueryTicketInfoResponseBody {
	s.Success = &v
	return s
}

type QueryTicketInfoResponseBodyResult struct {
	AccessTicket   *string `json:"AccessTicket,omitempty" xml:"AccessTicket,omitempty"`
	CmptId         *string `json:"CmptId,omitempty" xml:"CmptId,omitempty"`
	GlobalParam    *string `json:"GlobalParam,omitempty" xml:"GlobalParam,omitempty"`
	InvalidTime    *string `json:"InvalidTime,omitempty" xml:"InvalidTime,omitempty"`
	MaxTicketNum   *int32  `json:"MaxTicketNum,omitempty" xml:"MaxTicketNum,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	RegisterTime   *string `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	UsedTicketNum  *int32  `json:"UsedTicketNum,omitempty" xml:"UsedTicketNum,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WatermarkParam *string `json:"WatermarkParam,omitempty" xml:"WatermarkParam,omitempty"`
	WorksId        *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s QueryTicketInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryTicketInfoResponseBodyResult) SetAccessTicket(v string) *QueryTicketInfoResponseBodyResult {
	s.AccessTicket = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetCmptId(v string) *QueryTicketInfoResponseBodyResult {
	s.CmptId = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetGlobalParam(v string) *QueryTicketInfoResponseBodyResult {
	s.GlobalParam = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetInvalidTime(v string) *QueryTicketInfoResponseBodyResult {
	s.InvalidTime = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetMaxTicketNum(v int32) *QueryTicketInfoResponseBodyResult {
	s.MaxTicketNum = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetOrganizationId(v string) *QueryTicketInfoResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetRegisterTime(v string) *QueryTicketInfoResponseBodyResult {
	s.RegisterTime = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetUsedTicketNum(v int32) *QueryTicketInfoResponseBodyResult {
	s.UsedTicketNum = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetUserId(v string) *QueryTicketInfoResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetWatermarkParam(v string) *QueryTicketInfoResponseBodyResult {
	s.WatermarkParam = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetWorksId(v string) *QueryTicketInfoResponseBodyResult {
	s.WorksId = &v
	return s
}

type QueryTicketInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTicketInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTicketInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketInfoResponse) SetHeaders(v map[string]*string) *QueryTicketInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketInfoResponse) SetStatusCode(v int32) *QueryTicketInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTicketInfoResponse) SetBody(v *QueryTicketInfoResponseBody) *QueryTicketInfoResponse {
	s.Body = v
	return s
}

type QueryUserGroupListByParentIdRequest struct {
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
}

func (s QueryUserGroupListByParentIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserGroupListByParentIdRequest) GoString() string {
	return s.String()
}

func (s *QueryUserGroupListByParentIdRequest) SetParentUserGroupId(v string) *QueryUserGroupListByParentIdRequest {
	s.ParentUserGroupId = &v
	return s
}

type QueryUserGroupListByParentIdResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryUserGroupListByParentIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserGroupListByParentIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserGroupListByParentIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserGroupListByParentIdResponseBody) SetRequestId(v string) *QueryUserGroupListByParentIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBody) SetResult(v []*QueryUserGroupListByParentIdResponseBodyResult) *QueryUserGroupListByParentIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBody) SetSuccess(v bool) *QueryUserGroupListByParentIdResponseBody {
	s.Success = &v
	return s
}

type QueryUserGroupListByParentIdResponseBodyResult struct {
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser           *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	IdentifiedPath       *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	ModifiedTime         *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ModifyUser           *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	ParentUserGroupId    *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	UserGroupId          *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName        *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s QueryUserGroupListByParentIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserGroupListByParentIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetCreateTime(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetCreateUser(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.CreateUser = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetIdentifiedPath(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.IdentifiedPath = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetModifiedTime(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetModifyUser(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetParentUserGroupId(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.ParentUserGroupId = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetUserGroupDescription(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.UserGroupDescription = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetUserGroupId(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.UserGroupId = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponseBodyResult) SetUserGroupName(v string) *QueryUserGroupListByParentIdResponseBodyResult {
	s.UserGroupName = &v
	return s
}

type QueryUserGroupListByParentIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserGroupListByParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserGroupListByParentIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserGroupListByParentIdResponse) GoString() string {
	return s.String()
}

func (s *QueryUserGroupListByParentIdResponse) SetHeaders(v map[string]*string) *QueryUserGroupListByParentIdResponse {
	s.Headers = v
	return s
}

func (s *QueryUserGroupListByParentIdResponse) SetStatusCode(v int32) *QueryUserGroupListByParentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponse) SetBody(v *QueryUserGroupListByParentIdResponseBody) *QueryUserGroupListByParentIdResponse {
	s.Body = v
	return s
}

type QueryUserGroupMemberRequest struct {
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s QueryUserGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *QueryUserGroupMemberRequest) SetKeyword(v string) *QueryUserGroupMemberRequest {
	s.Keyword = &v
	return s
}

func (s *QueryUserGroupMemberRequest) SetUserGroupId(v string) *QueryUserGroupMemberRequest {
	s.UserGroupId = &v
	return s
}

type QueryUserGroupMemberResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryUserGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserGroupMemberResponseBody) SetRequestId(v string) *QueryUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserGroupMemberResponseBody) SetResult(v []*QueryUserGroupMemberResponseBodyResult) *QueryUserGroupMemberResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserGroupMemberResponseBody) SetSuccess(v bool) *QueryUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

type QueryUserGroupMemberResponseBodyResult struct {
	Id                  *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsUserGroup         *bool   `json:"IsUserGroup,omitempty" xml:"IsUserGroup,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentUserGroupId   *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	ParentUserGroupName *string `json:"ParentUserGroupName,omitempty" xml:"ParentUserGroupName,omitempty"`
}

func (s QueryUserGroupMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserGroupMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserGroupMemberResponseBodyResult) SetId(v string) *QueryUserGroupMemberResponseBodyResult {
	s.Id = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) SetIsUserGroup(v bool) *QueryUserGroupMemberResponseBodyResult {
	s.IsUserGroup = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) SetName(v string) *QueryUserGroupMemberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) SetParentUserGroupId(v string) *QueryUserGroupMemberResponseBodyResult {
	s.ParentUserGroupId = &v
	return s
}

func (s *QueryUserGroupMemberResponseBodyResult) SetParentUserGroupName(v string) *QueryUserGroupMemberResponseBodyResult {
	s.ParentUserGroupName = &v
	return s
}

type QueryUserGroupMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *QueryUserGroupMemberResponse) SetHeaders(v map[string]*string) *QueryUserGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *QueryUserGroupMemberResponse) SetStatusCode(v int32) *QueryUserGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserGroupMemberResponse) SetBody(v *QueryUserGroupMemberResponseBody) *QueryUserGroupMemberResponse {
	s.Body = v
	return s
}

type QueryUserInfoByAccountRequest struct {
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
}

func (s QueryUserInfoByAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoByAccountRequest) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByAccountRequest) SetAccount(v string) *QueryUserInfoByAccountRequest {
	s.Account = &v
	return s
}

type QueryUserInfoByAccountResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryUserInfoByAccountResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserInfoByAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoByAccountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByAccountResponseBody) SetRequestId(v string) *QueryUserInfoByAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBody) SetResult(v *QueryUserInfoByAccountResponseBodyResult) *QueryUserInfoByAccountResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserInfoByAccountResponseBody) SetSuccess(v bool) *QueryUserInfoByAccountResponseBody {
	s.Success = &v
	return s
}

type QueryUserInfoByAccountResponseBodyResult struct {
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AdminUser     *bool   `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	AuthAdminUser *bool   `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	NickName      *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Phone         *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType      *int32  `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s QueryUserInfoByAccountResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoByAccountResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetAccountId(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.AccountId = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetAccountName(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.AccountName = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetAdminUser(v bool) *QueryUserInfoByAccountResponseBodyResult {
	s.AdminUser = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetAuthAdminUser(v bool) *QueryUserInfoByAccountResponseBodyResult {
	s.AuthAdminUser = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetEmail(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.Email = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetNickName(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.NickName = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetPhone(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.Phone = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetUserId(v string) *QueryUserInfoByAccountResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryUserInfoByAccountResponseBodyResult) SetUserType(v int32) *QueryUserInfoByAccountResponseBodyResult {
	s.UserType = &v
	return s
}

type QueryUserInfoByAccountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserInfoByAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserInfoByAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoByAccountResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByAccountResponse) SetHeaders(v map[string]*string) *QueryUserInfoByAccountResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoByAccountResponse) SetStatusCode(v int32) *QueryUserInfoByAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserInfoByAccountResponse) SetBody(v *QueryUserInfoByAccountResponseBody) *QueryUserInfoByAccountResponse {
	s.Body = v
	return s
}

type QueryUserInfoByUserIdRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserInfoByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByUserIdRequest) SetUserId(v string) *QueryUserInfoByUserIdRequest {
	s.UserId = &v
	return s
}

type QueryUserInfoByUserIdResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryUserInfoByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserInfoByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByUserIdResponseBody) SetRequestId(v string) *QueryUserInfoByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBody) SetResult(v *QueryUserInfoByUserIdResponseBodyResult) *QueryUserInfoByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserInfoByUserIdResponseBody) SetSuccess(v bool) *QueryUserInfoByUserIdResponseBody {
	s.Success = &v
	return s
}

type QueryUserInfoByUserIdResponseBodyResult struct {
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AdminUser     *bool   `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	AuthAdminUser *bool   `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	Email         *string `json:"Email,omitempty" xml:"Email,omitempty"`
	NickName      *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Phone         *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType      *int32  `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s QueryUserInfoByUserIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetAccountId(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.AccountId = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetAccountName(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.AccountName = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetAdminUser(v bool) *QueryUserInfoByUserIdResponseBodyResult {
	s.AdminUser = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetAuthAdminUser(v bool) *QueryUserInfoByUserIdResponseBodyResult {
	s.AuthAdminUser = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetEmail(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.Email = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetNickName(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.NickName = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetPhone(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.Phone = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetUserId(v string) *QueryUserInfoByUserIdResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryUserInfoByUserIdResponseBodyResult) SetUserType(v int32) *QueryUserInfoByUserIdResponseBodyResult {
	s.UserType = &v
	return s
}

type QueryUserInfoByUserIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserInfoByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserInfoByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoByUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByUserIdResponse) SetHeaders(v map[string]*string) *QueryUserInfoByUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoByUserIdResponse) SetStatusCode(v int32) *QueryUserInfoByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserInfoByUserIdResponse) SetBody(v *QueryUserInfoByUserIdResponseBody) *QueryUserInfoByUserIdResponse {
	s.Body = v
	return s
}

type QueryUserListRequest struct {
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserListRequest) GoString() string {
	return s.String()
}

func (s *QueryUserListRequest) SetKeyword(v string) *QueryUserListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryUserListRequest) SetPageNum(v int32) *QueryUserListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryUserListRequest) SetPageSize(v int32) *QueryUserListRequest {
	s.PageSize = &v
	return s
}

type QueryUserListResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserListResponseBody) SetRequestId(v string) *QueryUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserListResponseBody) SetResult(v *QueryUserListResponseBodyResult) *QueryUserListResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserListResponseBody) SetSuccess(v bool) *QueryUserListResponseBody {
	s.Success = &v
	return s
}

type QueryUserListResponseBodyResult struct {
	Data       []*QueryUserListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryUserListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserListResponseBodyResult) SetData(v []*QueryUserListResponseBodyResultData) *QueryUserListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryUserListResponseBodyResult) SetPageNum(v int32) *QueryUserListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryUserListResponseBodyResult) SetPageSize(v int32) *QueryUserListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryUserListResponseBodyResult) SetTotalNum(v int32) *QueryUserListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryUserListResponseBodyResult) SetTotalPages(v int32) *QueryUserListResponseBodyResult {
	s.TotalPages = &v
	return s
}

type QueryUserListResponseBodyResultData struct {
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AdminUser     *bool   `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	AuthAdminUser *bool   `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	NickName      *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType      *int32  `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s QueryUserListResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryUserListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryUserListResponseBodyResultData) SetAccountId(v string) *QueryUserListResponseBodyResultData {
	s.AccountId = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetAccountName(v string) *QueryUserListResponseBodyResultData {
	s.AccountName = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetAdminUser(v bool) *QueryUserListResponseBodyResultData {
	s.AdminUser = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetAuthAdminUser(v bool) *QueryUserListResponseBodyResultData {
	s.AuthAdminUser = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetNickName(v string) *QueryUserListResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetUserId(v string) *QueryUserListResponseBodyResultData {
	s.UserId = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetUserType(v int32) *QueryUserListResponseBodyResultData {
	s.UserType = &v
	return s
}

type QueryUserListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserListResponse) GoString() string {
	return s.String()
}

func (s *QueryUserListResponse) SetHeaders(v map[string]*string) *QueryUserListResponse {
	s.Headers = v
	return s
}

func (s *QueryUserListResponse) SetStatusCode(v int32) *QueryUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserListResponse) SetBody(v *QueryUserListResponseBody) *QueryUserListResponse {
	s.Body = v
	return s
}

type QueryUserRoleInfoInWorkspaceRequest struct {
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryUserRoleInfoInWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleInfoInWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *QueryUserRoleInfoInWorkspaceRequest) SetUserId(v string) *QueryUserRoleInfoInWorkspaceRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceRequest) SetWorkspaceId(v string) *QueryUserRoleInfoInWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

type QueryUserRoleInfoInWorkspaceResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryUserRoleInfoInWorkspaceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserRoleInfoInWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleInfoInWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) SetRequestId(v string) *QueryUserRoleInfoInWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) SetResult(v *QueryUserRoleInfoInWorkspaceResponseBodyResult) *QueryUserRoleInfoInWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) SetSuccess(v bool) *QueryUserRoleInfoInWorkspaceResponseBody {
	s.Success = &v
	return s
}

type QueryUserRoleInfoInWorkspaceResponseBodyResult struct {
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	RoleId   *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s QueryUserRoleInfoInWorkspaceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleInfoInWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) SetRoleCode(v string) *QueryUserRoleInfoInWorkspaceResponseBodyResult {
	s.RoleCode = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) SetRoleId(v int64) *QueryUserRoleInfoInWorkspaceResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) SetRoleName(v string) *QueryUserRoleInfoInWorkspaceResponseBodyResult {
	s.RoleName = &v
	return s
}

type QueryUserRoleInfoInWorkspaceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserRoleInfoInWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserRoleInfoInWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRoleInfoInWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *QueryUserRoleInfoInWorkspaceResponse) SetHeaders(v map[string]*string) *QueryUserRoleInfoInWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponse) SetStatusCode(v int32) *QueryUserRoleInfoInWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponse) SetBody(v *QueryUserRoleInfoInWorkspaceResponseBody) *QueryUserRoleInfoInWorkspaceResponse {
	s.Body = v
	return s
}

type QueryUserTagMetaListResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryUserTagMetaListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserTagMetaListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserTagMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserTagMetaListResponseBody) SetRequestId(v string) *QueryUserTagMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserTagMetaListResponseBody) SetResult(v []*QueryUserTagMetaListResponseBodyResult) *QueryUserTagMetaListResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserTagMetaListResponseBody) SetSuccess(v bool) *QueryUserTagMetaListResponseBody {
	s.Success = &v
	return s
}

type QueryUserTagMetaListResponseBodyResult struct {
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	TagId          *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName        *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryUserTagMetaListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserTagMetaListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserTagMetaListResponseBodyResult) SetTagDescription(v string) *QueryUserTagMetaListResponseBodyResult {
	s.TagDescription = &v
	return s
}

func (s *QueryUserTagMetaListResponseBodyResult) SetTagId(v string) *QueryUserTagMetaListResponseBodyResult {
	s.TagId = &v
	return s
}

func (s *QueryUserTagMetaListResponseBodyResult) SetTagName(v string) *QueryUserTagMetaListResponseBodyResult {
	s.TagName = &v
	return s
}

type QueryUserTagMetaListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserTagMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserTagMetaListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserTagMetaListResponse) GoString() string {
	return s.String()
}

func (s *QueryUserTagMetaListResponse) SetHeaders(v map[string]*string) *QueryUserTagMetaListResponse {
	s.Headers = v
	return s
}

func (s *QueryUserTagMetaListResponse) SetStatusCode(v int32) *QueryUserTagMetaListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserTagMetaListResponse) SetBody(v *QueryUserTagMetaListResponseBody) *QueryUserTagMetaListResponse {
	s.Body = v
	return s
}

type QueryUserTagValueListRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserTagValueListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserTagValueListRequest) GoString() string {
	return s.String()
}

func (s *QueryUserTagValueListRequest) SetUserId(v string) *QueryUserTagValueListRequest {
	s.UserId = &v
	return s
}

type QueryUserTagValueListResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryUserTagValueListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserTagValueListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserTagValueListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserTagValueListResponseBody) SetRequestId(v string) *QueryUserTagValueListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserTagValueListResponseBody) SetResult(v []*QueryUserTagValueListResponseBodyResult) *QueryUserTagValueListResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserTagValueListResponseBody) SetSuccess(v bool) *QueryUserTagValueListResponseBody {
	s.Success = &v
	return s
}

type QueryUserTagValueListResponseBodyResult struct {
	TagId    *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName  *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s QueryUserTagValueListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserTagValueListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserTagValueListResponseBodyResult) SetTagId(v string) *QueryUserTagValueListResponseBodyResult {
	s.TagId = &v
	return s
}

func (s *QueryUserTagValueListResponseBodyResult) SetTagName(v string) *QueryUserTagValueListResponseBodyResult {
	s.TagName = &v
	return s
}

func (s *QueryUserTagValueListResponseBodyResult) SetTagValue(v string) *QueryUserTagValueListResponseBodyResult {
	s.TagValue = &v
	return s
}

type QueryUserTagValueListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserTagValueListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserTagValueListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserTagValueListResponse) GoString() string {
	return s.String()
}

func (s *QueryUserTagValueListResponse) SetHeaders(v map[string]*string) *QueryUserTagValueListResponse {
	s.Headers = v
	return s
}

func (s *QueryUserTagValueListResponse) SetStatusCode(v int32) *QueryUserTagValueListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserTagValueListResponse) SetBody(v *QueryUserTagValueListResponseBody) *QueryUserTagValueListResponse {
	s.Body = v
	return s
}

type QueryWorksRequest struct {
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s QueryWorksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksRequest) GoString() string {
	return s.String()
}

func (s *QueryWorksRequest) SetWorksId(v string) *QueryWorksRequest {
	s.WorksId = &v
	return s
}

type QueryWorksResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryWorksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWorksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorksResponseBody) SetRequestId(v string) *QueryWorksResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorksResponseBody) SetResult(v *QueryWorksResponseBodyResult) *QueryWorksResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorksResponseBody) SetSuccess(v bool) *QueryWorksResponseBody {
	s.Success = &v
	return s
}

type QueryWorksResponseBodyResult struct {
	Auth3rdFlag   *int32                                 `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	Description   *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory     *QueryWorksResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	GmtCreate     *string                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify     *string                                `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	ModifyName    *string                                `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	OwnerId       *string                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName     *string                                `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	SecurityLevel *string                                `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	Status        *int32                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkName      *string                                `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	WorkType      *string                                `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	WorksId       *string                                `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	WorkspaceId   *string                                `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string                                `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryWorksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorksResponseBodyResult) SetAuth3rdFlag(v int32) *QueryWorksResponseBodyResult {
	s.Auth3rdFlag = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetDescription(v string) *QueryWorksResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetDirectory(v *QueryWorksResponseBodyResultDirectory) *QueryWorksResponseBodyResult {
	s.Directory = v
	return s
}

func (s *QueryWorksResponseBodyResult) SetGmtCreate(v string) *QueryWorksResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetGmtModify(v string) *QueryWorksResponseBodyResult {
	s.GmtModify = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetModifyName(v string) *QueryWorksResponseBodyResult {
	s.ModifyName = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetOwnerId(v string) *QueryWorksResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetOwnerName(v string) *QueryWorksResponseBodyResult {
	s.OwnerName = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetSecurityLevel(v string) *QueryWorksResponseBodyResult {
	s.SecurityLevel = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetStatus(v int32) *QueryWorksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorkName(v string) *QueryWorksResponseBodyResult {
	s.WorkName = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorkType(v string) *QueryWorksResponseBodyResult {
	s.WorkType = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorksId(v string) *QueryWorksResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorkspaceId(v string) *QueryWorksResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetWorkspaceName(v string) *QueryWorksResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

type QueryWorksResponseBodyResultDirectory struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PathId   *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryWorksResponseBodyResultDirectory) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksResponseBodyResultDirectory) GoString() string {
	return s.String()
}

func (s *QueryWorksResponseBodyResultDirectory) SetId(v string) *QueryWorksResponseBodyResultDirectory {
	s.Id = &v
	return s
}

func (s *QueryWorksResponseBodyResultDirectory) SetName(v string) *QueryWorksResponseBodyResultDirectory {
	s.Name = &v
	return s
}

func (s *QueryWorksResponseBodyResultDirectory) SetPathId(v string) *QueryWorksResponseBodyResultDirectory {
	s.PathId = &v
	return s
}

func (s *QueryWorksResponseBodyResultDirectory) SetPathName(v string) *QueryWorksResponseBodyResultDirectory {
	s.PathName = &v
	return s
}

type QueryWorksResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryWorksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWorksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksResponse) GoString() string {
	return s.String()
}

func (s *QueryWorksResponse) SetHeaders(v map[string]*string) *QueryWorksResponse {
	s.Headers = v
	return s
}

func (s *QueryWorksResponse) SetStatusCode(v int32) *QueryWorksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorksResponse) SetBody(v *QueryWorksResponseBody) *QueryWorksResponse {
	s.Body = v
	return s
}

type QueryWorksBloodRelationshipRequest struct {
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s QueryWorksBloodRelationshipRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksBloodRelationshipRequest) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipRequest) SetWorksId(v string) *QueryWorksBloodRelationshipRequest {
	s.WorksId = &v
	return s
}

type QueryWorksBloodRelationshipResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryWorksBloodRelationshipResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWorksBloodRelationshipResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksBloodRelationshipResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipResponseBody) SetRequestId(v string) *QueryWorksBloodRelationshipResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBody) SetResult(v []*QueryWorksBloodRelationshipResponseBodyResult) *QueryWorksBloodRelationshipResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBody) SetSuccess(v bool) *QueryWorksBloodRelationshipResponseBody {
	s.Success = &v
	return s
}

type QueryWorksBloodRelationshipResponseBodyResult struct {
	ComponentId       *string                                                     `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	ComponentName     *string                                                     `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	ComponentType     *int32                                                      `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	ComponentTypeName *string                                                     `json:"ComponentTypeName,omitempty" xml:"ComponentTypeName,omitempty"`
	DatasetId         *string                                                     `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	QueryParams       []*QueryWorksBloodRelationshipResponseBodyResultQueryParams `json:"QueryParams,omitempty" xml:"QueryParams,omitempty" type:"Repeated"`
}

func (s QueryWorksBloodRelationshipResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksBloodRelationshipResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentId(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentName(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentName = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentType(v int32) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentType = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetComponentTypeName(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.ComponentTypeName = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetDatasetId(v string) *QueryWorksBloodRelationshipResponseBodyResult {
	s.DatasetId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResult) SetQueryParams(v []*QueryWorksBloodRelationshipResponseBodyResultQueryParams) *QueryWorksBloodRelationshipResponseBodyResult {
	s.QueryParams = v
	return s
}

type QueryWorksBloodRelationshipResponseBodyResultQueryParams struct {
	AreaId    *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	AreaName  *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	Caption   *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	IsMeasure *bool   `json:"IsMeasure,omitempty" xml:"IsMeasure,omitempty"`
	PathId    *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryWorksBloodRelationshipResponseBodyResultQueryParams) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksBloodRelationshipResponseBodyResultQueryParams) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetAreaId(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.AreaId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetAreaName(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.AreaName = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetCaption(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.Caption = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetDataType(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.DataType = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetIsMeasure(v bool) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.IsMeasure = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetPathId(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.PathId = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponseBodyResultQueryParams) SetUid(v string) *QueryWorksBloodRelationshipResponseBodyResultQueryParams {
	s.Uid = &v
	return s
}

type QueryWorksBloodRelationshipResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryWorksBloodRelationshipResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWorksBloodRelationshipResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksBloodRelationshipResponse) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipResponse) SetHeaders(v map[string]*string) *QueryWorksBloodRelationshipResponse {
	s.Headers = v
	return s
}

func (s *QueryWorksBloodRelationshipResponse) SetStatusCode(v int32) *QueryWorksBloodRelationshipResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponse) SetBody(v *QueryWorksBloodRelationshipResponseBody) *QueryWorksBloodRelationshipResponse {
	s.Body = v
	return s
}

type QueryWorksByOrganizationRequest struct {
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ThirdPartAuthFlag *int32  `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	WorksType         *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
}

func (s QueryWorksByOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByOrganizationRequest) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationRequest) SetPageNum(v int32) *QueryWorksByOrganizationRequest {
	s.PageNum = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) SetPageSize(v int32) *QueryWorksByOrganizationRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) SetStatus(v int32) *QueryWorksByOrganizationRequest {
	s.Status = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) SetThirdPartAuthFlag(v int32) *QueryWorksByOrganizationRequest {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QueryWorksByOrganizationRequest) SetWorksType(v string) *QueryWorksByOrganizationRequest {
	s.WorksType = &v
	return s
}

type QueryWorksByOrganizationResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryWorksByOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWorksByOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponseBody) SetRequestId(v string) *QueryWorksByOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBody) SetResult(v *QueryWorksByOrganizationResponseBodyResult) *QueryWorksByOrganizationResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorksByOrganizationResponseBody) SetSuccess(v bool) *QueryWorksByOrganizationResponseBody {
	s.Success = &v
	return s
}

type QueryWorksByOrganizationResponseBodyResult struct {
	Data       []*QueryWorksByOrganizationResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                            `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryWorksByOrganizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByOrganizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetData(v []*QueryWorksByOrganizationResponseBodyResultData) *QueryWorksByOrganizationResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetPageNum(v int32) *QueryWorksByOrganizationResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetPageSize(v int32) *QueryWorksByOrganizationResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetTotalNum(v int32) *QueryWorksByOrganizationResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResult) SetTotalPages(v int32) *QueryWorksByOrganizationResponseBodyResult {
	s.TotalPages = &v
	return s
}

type QueryWorksByOrganizationResponseBodyResultData struct {
	Auth3rdFlag   *int32                                                   `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	Description   *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory     *QueryWorksByOrganizationResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	GmtCreate     *string                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify     *string                                                  `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	ModifyName    *string                                                  `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	OwnerId       *string                                                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName     *string                                                  `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	SecurityLevel *string                                                  `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	Status        *int32                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkName      *string                                                  `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	WorkType      *string                                                  `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	WorksId       *string                                                  `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	WorkspaceId   *string                                                  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string                                                  `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryWorksByOrganizationResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByOrganizationResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetAuth3rdFlag(v int32) *QueryWorksByOrganizationResponseBodyResultData {
	s.Auth3rdFlag = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetDescription(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.Description = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetDirectory(v *QueryWorksByOrganizationResponseBodyResultDataDirectory) *QueryWorksByOrganizationResponseBodyResultData {
	s.Directory = v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetGmtCreate(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetGmtModify(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.GmtModify = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetModifyName(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.ModifyName = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetOwnerId(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.OwnerId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetOwnerName(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetSecurityLevel(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.SecurityLevel = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetStatus(v int32) *QueryWorksByOrganizationResponseBodyResultData {
	s.Status = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorkName(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorkName = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorkType(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorkType = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorksId(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorksId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorkspaceId(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetWorkspaceName(v string) *QueryWorksByOrganizationResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type QueryWorksByOrganizationResponseBodyResultDataDirectory struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PathId   *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryWorksByOrganizationResponseBodyResultDataDirectory) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByOrganizationResponseBodyResultDataDirectory) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) SetId(v string) *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	s.Id = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) SetName(v string) *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	s.Name = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) SetPathId(v string) *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	s.PathId = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultDataDirectory) SetPathName(v string) *QueryWorksByOrganizationResponseBodyResultDataDirectory {
	s.PathName = &v
	return s
}

type QueryWorksByOrganizationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryWorksByOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWorksByOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByOrganizationResponse) GoString() string {
	return s.String()
}

func (s *QueryWorksByOrganizationResponse) SetHeaders(v map[string]*string) *QueryWorksByOrganizationResponse {
	s.Headers = v
	return s
}

func (s *QueryWorksByOrganizationResponse) SetStatusCode(v int32) *QueryWorksByOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorksByOrganizationResponse) SetBody(v *QueryWorksByOrganizationResponseBody) *QueryWorksByOrganizationResponse {
	s.Body = v
	return s
}

type QueryWorksByWorkspaceRequest struct {
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ThirdPartAuthFlag *int32  `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	WorksType         *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
	WorkspaceId       *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryWorksByWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceRequest) SetPageNum(v int32) *QueryWorksByWorkspaceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetPageSize(v int32) *QueryWorksByWorkspaceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetStatus(v int32) *QueryWorksByWorkspaceRequest {
	s.Status = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetThirdPartAuthFlag(v int32) *QueryWorksByWorkspaceRequest {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetWorksType(v string) *QueryWorksByWorkspaceRequest {
	s.WorksType = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetWorkspaceId(v string) *QueryWorksByWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

type QueryWorksByWorkspaceResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryWorksByWorkspaceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWorksByWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponseBody) SetRequestId(v string) *QueryWorksByWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBody) SetResult(v *QueryWorksByWorkspaceResponseBodyResult) *QueryWorksByWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorksByWorkspaceResponseBody) SetSuccess(v bool) *QueryWorksByWorkspaceResponseBody {
	s.Success = &v
	return s
}

type QueryWorksByWorkspaceResponseBodyResult struct {
	Data       []*QueryWorksByWorkspaceResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                         `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                         `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                         `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryWorksByWorkspaceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetData(v []*QueryWorksByWorkspaceResponseBodyResultData) *QueryWorksByWorkspaceResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetPageNum(v int32) *QueryWorksByWorkspaceResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetPageSize(v int32) *QueryWorksByWorkspaceResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetTotalNum(v int32) *QueryWorksByWorkspaceResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResult) SetTotalPages(v int32) *QueryWorksByWorkspaceResponseBodyResult {
	s.TotalPages = &v
	return s
}

type QueryWorksByWorkspaceResponseBodyResultData struct {
	Auth3rdFlag   *int32                                                `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	Description   *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory     *QueryWorksByWorkspaceResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	GmtCreate     *string                                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModify     *string                                               `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	ModifyName    *string                                               `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	OwnerId       *string                                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName     *string                                               `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	SecurityLevel *string                                               `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	Status        *int32                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkName      *string                                               `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	WorkType      *string                                               `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	WorksId       *string                                               `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	WorkspaceId   *string                                               `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string                                               `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryWorksByWorkspaceResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByWorkspaceResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetAuth3rdFlag(v int32) *QueryWorksByWorkspaceResponseBodyResultData {
	s.Auth3rdFlag = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetDescription(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.Description = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetDirectory(v *QueryWorksByWorkspaceResponseBodyResultDataDirectory) *QueryWorksByWorkspaceResponseBodyResultData {
	s.Directory = v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetGmtCreate(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetGmtModify(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.GmtModify = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetModifyName(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.ModifyName = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetOwnerId(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.OwnerId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetOwnerName(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetSecurityLevel(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.SecurityLevel = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetStatus(v int32) *QueryWorksByWorkspaceResponseBodyResultData {
	s.Status = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorkName(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorkName = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorkType(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorkType = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorksId(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorksId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorkspaceId(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetWorkspaceName(v string) *QueryWorksByWorkspaceResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type QueryWorksByWorkspaceResponseBodyResultDataDirectory struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PathId   *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	PathName *string `json:"PathName,omitempty" xml:"PathName,omitempty"`
}

func (s QueryWorksByWorkspaceResponseBodyResultDataDirectory) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByWorkspaceResponseBodyResultDataDirectory) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) SetId(v string) *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	s.Id = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) SetName(v string) *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	s.Name = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) SetPathId(v string) *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	s.PathId = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultDataDirectory) SetPathName(v string) *QueryWorksByWorkspaceResponseBodyResultDataDirectory {
	s.PathName = &v
	return s
}

type QueryWorksByWorkspaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryWorksByWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWorksByWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWorksByWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceResponse) SetHeaders(v map[string]*string) *QueryWorksByWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *QueryWorksByWorkspaceResponse) SetStatusCode(v int32) *QueryWorksByWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorksByWorkspaceResponse) SetBody(v *QueryWorksByWorkspaceResponseBody) *QueryWorksByWorkspaceResponse {
	s.Body = v
	return s
}

type QueryWorkspaceUserListRequest struct {
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryWorkspaceUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceUserListRequest) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListRequest) SetKeyword(v string) *QueryWorkspaceUserListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryWorkspaceUserListRequest) SetPageNum(v int32) *QueryWorkspaceUserListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryWorkspaceUserListRequest) SetPageSize(v int32) *QueryWorkspaceUserListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorkspaceUserListRequest) SetWorkspaceId(v string) *QueryWorkspaceUserListRequest {
	s.WorkspaceId = &v
	return s
}

type QueryWorkspaceUserListResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryWorkspaceUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWorkspaceUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceUserListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponseBody) SetRequestId(v string) *QueryWorkspaceUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBody) SetResult(v *QueryWorkspaceUserListResponseBodyResult) *QueryWorkspaceUserListResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorkspaceUserListResponseBody) SetSuccess(v bool) *QueryWorkspaceUserListResponseBody {
	s.Success = &v
	return s
}

type QueryWorkspaceUserListResponseBodyResult struct {
	Data       []*QueryWorkspaceUserListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNum    *int32                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum   *int32                                          `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPages *int32                                          `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryWorkspaceUserListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceUserListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetData(v []*QueryWorkspaceUserListResponseBodyResultData) *QueryWorkspaceUserListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetPageNum(v int32) *QueryWorkspaceUserListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetPageSize(v int32) *QueryWorkspaceUserListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetTotalNum(v int32) *QueryWorkspaceUserListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResult) SetTotalPages(v int32) *QueryWorkspaceUserListResponseBodyResult {
	s.TotalPages = &v
	return s
}

type QueryWorkspaceUserListResponseBodyResultData struct {
	AccountId   *string                                           `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName *string                                           `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	NickName    *string                                           `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Role        *QueryWorkspaceUserListResponseBodyResultDataRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
	UserId      *string                                           `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryWorkspaceUserListResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceUserListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetAccountId(v string) *QueryWorkspaceUserListResponseBodyResultData {
	s.AccountId = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetAccountName(v string) *QueryWorkspaceUserListResponseBodyResultData {
	s.AccountName = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetNickName(v string) *QueryWorkspaceUserListResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetRole(v *QueryWorkspaceUserListResponseBodyResultDataRole) *QueryWorkspaceUserListResponseBodyResultData {
	s.Role = v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultData) SetUserId(v string) *QueryWorkspaceUserListResponseBodyResultData {
	s.UserId = &v
	return s
}

type QueryWorkspaceUserListResponseBodyResultDataRole struct {
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	RoleId   *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s QueryWorkspaceUserListResponseBodyResultDataRole) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceUserListResponseBodyResultDataRole) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) SetRoleCode(v string) *QueryWorkspaceUserListResponseBodyResultDataRole {
	s.RoleCode = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) SetRoleId(v int64) *QueryWorkspaceUserListResponseBodyResultDataRole {
	s.RoleId = &v
	return s
}

func (s *QueryWorkspaceUserListResponseBodyResultDataRole) SetRoleName(v string) *QueryWorkspaceUserListResponseBodyResultDataRole {
	s.RoleName = &v
	return s
}

type QueryWorkspaceUserListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryWorkspaceUserListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWorkspaceUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceUserListResponse) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListResponse) SetHeaders(v map[string]*string) *QueryWorkspaceUserListResponse {
	s.Headers = v
	return s
}

func (s *QueryWorkspaceUserListResponse) SetStatusCode(v int32) *QueryWorkspaceUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorkspaceUserListResponse) SetBody(v *QueryWorkspaceUserListResponseBody) *QueryWorkspaceUserListResponse {
	s.Body = v
	return s
}

type ResultCallbackRequest struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	HandleReason  *string `json:"HandleReason,omitempty" xml:"HandleReason,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResultCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s ResultCallbackRequest) GoString() string {
	return s.String()
}

func (s *ResultCallbackRequest) SetApplicationId(v string) *ResultCallbackRequest {
	s.ApplicationId = &v
	return s
}

func (s *ResultCallbackRequest) SetHandleReason(v string) *ResultCallbackRequest {
	s.HandleReason = &v
	return s
}

func (s *ResultCallbackRequest) SetStatus(v int32) *ResultCallbackRequest {
	s.Status = &v
	return s
}

type ResultCallbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResultCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResultCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ResultCallbackResponseBody) SetRequestId(v string) *ResultCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResultCallbackResponseBody) SetResult(v bool) *ResultCallbackResponseBody {
	s.Result = &v
	return s
}

func (s *ResultCallbackResponseBody) SetSuccess(v bool) *ResultCallbackResponseBody {
	s.Success = &v
	return s
}

type ResultCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResultCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResultCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s ResultCallbackResponse) GoString() string {
	return s.String()
}

func (s *ResultCallbackResponse) SetHeaders(v map[string]*string) *ResultCallbackResponse {
	s.Headers = v
	return s
}

func (s *ResultCallbackResponse) SetStatusCode(v int32) *ResultCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ResultCallbackResponse) SetBody(v *ResultCallbackResponseBody) *ResultCallbackResponse {
	s.Body = v
	return s
}

type SaveFavoritesRequest struct {
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s SaveFavoritesRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveFavoritesRequest) GoString() string {
	return s.String()
}

func (s *SaveFavoritesRequest) SetUserId(v string) *SaveFavoritesRequest {
	s.UserId = &v
	return s
}

func (s *SaveFavoritesRequest) SetWorksId(v string) *SaveFavoritesRequest {
	s.WorksId = &v
	return s
}

type SaveFavoritesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveFavoritesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveFavoritesResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFavoritesResponseBody) SetRequestId(v string) *SaveFavoritesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveFavoritesResponseBody) SetResult(v bool) *SaveFavoritesResponseBody {
	s.Result = &v
	return s
}

func (s *SaveFavoritesResponseBody) SetSuccess(v bool) *SaveFavoritesResponseBody {
	s.Success = &v
	return s
}

type SaveFavoritesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveFavoritesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveFavoritesResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveFavoritesResponse) GoString() string {
	return s.String()
}

func (s *SaveFavoritesResponse) SetHeaders(v map[string]*string) *SaveFavoritesResponse {
	s.Headers = v
	return s
}

func (s *SaveFavoritesResponse) SetStatusCode(v int32) *SaveFavoritesResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveFavoritesResponse) SetBody(v *SaveFavoritesResponseBody) *SaveFavoritesResponse {
	s.Body = v
	return s
}

type SetDataLevelPermissionExtraConfigRequest struct {
	CubeId        *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	MissHitPolicy *string `json:"MissHitPolicy,omitempty" xml:"MissHitPolicy,omitempty"`
	RuleType      *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s SetDataLevelPermissionExtraConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionExtraConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionExtraConfigRequest) SetCubeId(v string) *SetDataLevelPermissionExtraConfigRequest {
	s.CubeId = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigRequest) SetMissHitPolicy(v string) *SetDataLevelPermissionExtraConfigRequest {
	s.MissHitPolicy = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigRequest) SetRuleType(v string) *SetDataLevelPermissionExtraConfigRequest {
	s.RuleType = &v
	return s
}

type SetDataLevelPermissionExtraConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDataLevelPermissionExtraConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionExtraConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) SetRequestId(v string) *SetDataLevelPermissionExtraConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) SetResult(v bool) *SetDataLevelPermissionExtraConfigResponseBody {
	s.Result = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponseBody) SetSuccess(v bool) *SetDataLevelPermissionExtraConfigResponseBody {
	s.Success = &v
	return s
}

type SetDataLevelPermissionExtraConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDataLevelPermissionExtraConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDataLevelPermissionExtraConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionExtraConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionExtraConfigResponse) SetHeaders(v map[string]*string) *SetDataLevelPermissionExtraConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponse) SetStatusCode(v int32) *SetDataLevelPermissionExtraConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataLevelPermissionExtraConfigResponse) SetBody(v *SetDataLevelPermissionExtraConfigResponseBody) *SetDataLevelPermissionExtraConfigResponse {
	s.Body = v
	return s
}

type SetDataLevelPermissionRuleConfigRequest struct {
	RuleModel *string `json:"RuleModel,omitempty" xml:"RuleModel,omitempty"`
}

func (s SetDataLevelPermissionRuleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionRuleConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionRuleConfigRequest) SetRuleModel(v string) *SetDataLevelPermissionRuleConfigRequest {
	s.RuleModel = &v
	return s
}

type SetDataLevelPermissionRuleConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDataLevelPermissionRuleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionRuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) SetRequestId(v string) *SetDataLevelPermissionRuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) SetResult(v string) *SetDataLevelPermissionRuleConfigResponseBody {
	s.Result = &v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponseBody) SetSuccess(v bool) *SetDataLevelPermissionRuleConfigResponseBody {
	s.Success = &v
	return s
}

type SetDataLevelPermissionRuleConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDataLevelPermissionRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDataLevelPermissionRuleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionRuleConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionRuleConfigResponse) SetHeaders(v map[string]*string) *SetDataLevelPermissionRuleConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponse) SetStatusCode(v int32) *SetDataLevelPermissionRuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataLevelPermissionRuleConfigResponse) SetBody(v *SetDataLevelPermissionRuleConfigResponseBody) *SetDataLevelPermissionRuleConfigResponse {
	s.Body = v
	return s
}

type SetDataLevelPermissionWhiteListRequest struct {
	WhiteListModel *string `json:"WhiteListModel,omitempty" xml:"WhiteListModel,omitempty"`
}

func (s SetDataLevelPermissionWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionWhiteListRequest) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionWhiteListRequest) SetWhiteListModel(v string) *SetDataLevelPermissionWhiteListRequest {
	s.WhiteListModel = &v
	return s
}

type SetDataLevelPermissionWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDataLevelPermissionWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionWhiteListResponseBody) SetRequestId(v string) *SetDataLevelPermissionWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponseBody) SetResult(v bool) *SetDataLevelPermissionWhiteListResponseBody {
	s.Result = &v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponseBody) SetSuccess(v bool) *SetDataLevelPermissionWhiteListResponseBody {
	s.Success = &v
	return s
}

type SetDataLevelPermissionWhiteListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDataLevelPermissionWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDataLevelPermissionWhiteListResponse) GoString() string {
	return s.String()
}

func (s *SetDataLevelPermissionWhiteListResponse) SetHeaders(v map[string]*string) *SetDataLevelPermissionWhiteListResponse {
	s.Headers = v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponse) SetStatusCode(v int32) *SetDataLevelPermissionWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDataLevelPermissionWhiteListResponse) SetBody(v *SetDataLevelPermissionWhiteListResponseBody) *SetDataLevelPermissionWhiteListResponse {
	s.Body = v
	return s
}

type UpdateDataLevelPermissionStatusRequest struct {
	CubeId   *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	IsOpen   *int32  `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s UpdateDataLevelPermissionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataLevelPermissionStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLevelPermissionStatusRequest) SetCubeId(v string) *UpdateDataLevelPermissionStatusRequest {
	s.CubeId = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusRequest) SetIsOpen(v int32) *UpdateDataLevelPermissionStatusRequest {
	s.IsOpen = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusRequest) SetRuleType(v string) *UpdateDataLevelPermissionStatusRequest {
	s.RuleType = &v
	return s
}

type UpdateDataLevelPermissionStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataLevelPermissionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataLevelPermissionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataLevelPermissionStatusResponseBody) SetRequestId(v string) *UpdateDataLevelPermissionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponseBody) SetResult(v bool) *UpdateDataLevelPermissionStatusResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponseBody) SetSuccess(v bool) *UpdateDataLevelPermissionStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateDataLevelPermissionStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDataLevelPermissionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDataLevelPermissionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataLevelPermissionStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataLevelPermissionStatusResponse) SetHeaders(v map[string]*string) *UpdateDataLevelPermissionStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponse) SetStatusCode(v int32) *UpdateDataLevelPermissionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataLevelPermissionStatusResponse) SetBody(v *UpdateDataLevelPermissionStatusResponseBody) *UpdateDataLevelPermissionStatusResponse {
	s.Body = v
	return s
}

type UpdateEmbeddedStatusRequest struct {
	ThirdPartAuthFlag *bool   `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	WorksId           *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s UpdateEmbeddedStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmbeddedStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmbeddedStatusRequest) SetThirdPartAuthFlag(v bool) *UpdateEmbeddedStatusRequest {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *UpdateEmbeddedStatusRequest) SetWorksId(v string) *UpdateEmbeddedStatusRequest {
	s.WorksId = &v
	return s
}

type UpdateEmbeddedStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEmbeddedStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmbeddedStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEmbeddedStatusResponseBody) SetRequestId(v string) *UpdateEmbeddedStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEmbeddedStatusResponseBody) SetResult(v int32) *UpdateEmbeddedStatusResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateEmbeddedStatusResponseBody) SetSuccess(v bool) *UpdateEmbeddedStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateEmbeddedStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEmbeddedStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEmbeddedStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEmbeddedStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateEmbeddedStatusResponse) SetHeaders(v map[string]*string) *UpdateEmbeddedStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateEmbeddedStatusResponse) SetStatusCode(v int32) *UpdateEmbeddedStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEmbeddedStatusResponse) SetBody(v *UpdateEmbeddedStatusResponseBody) *UpdateEmbeddedStatusResponse {
	s.Body = v
	return s
}

type UpdateTicketNumRequest struct {
	Ticket    *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
	TicketNum *int32  `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
}

func (s UpdateTicketNumRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketNumRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketNumRequest) SetTicket(v string) *UpdateTicketNumRequest {
	s.Ticket = &v
	return s
}

func (s *UpdateTicketNumRequest) SetTicketNum(v int32) *UpdateTicketNumRequest {
	s.TicketNum = &v
	return s
}

type UpdateTicketNumResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTicketNumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketNumResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketNumResponseBody) SetRequestId(v string) *UpdateTicketNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketNumResponseBody) SetResult(v bool) *UpdateTicketNumResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateTicketNumResponseBody) SetSuccess(v bool) *UpdateTicketNumResponseBody {
	s.Success = &v
	return s
}

type UpdateTicketNumResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTicketNumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTicketNumResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketNumResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketNumResponse) SetHeaders(v map[string]*string) *UpdateTicketNumResponse {
	s.Headers = v
	return s
}

func (s *UpdateTicketNumResponse) SetStatusCode(v int32) *UpdateTicketNumResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTicketNumResponse) SetBody(v *UpdateTicketNumResponseBody) *UpdateTicketNumResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	AdminUser     *bool   `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	AuthAdminUser *bool   `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	NickName      *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserType      *int32  `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetAdminUser(v bool) *UpdateUserRequest {
	s.AdminUser = &v
	return s
}

func (s *UpdateUserRequest) SetAuthAdminUser(v bool) *UpdateUserRequest {
	s.AuthAdminUser = &v
	return s
}

func (s *UpdateUserRequest) SetNickName(v string) *UpdateUserRequest {
	s.NickName = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) SetUserType(v int32) *UpdateUserRequest {
	s.UserType = &v
	return s
}

type UpdateUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetResult(v bool) *UpdateUserResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type UpdateUserGroupRequest struct {
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	UserGroupId          *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName        *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s UpdateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) SetUserGroupDescription(v string) *UpdateUserGroupRequest {
	s.UserGroupDescription = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupId(v string) *UpdateUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupName(v string) *UpdateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

type UpdateUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetResult(v bool) *UpdateUserGroupResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetSuccess(v bool) *UpdateUserGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponse) SetHeaders(v map[string]*string) *UpdateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserGroupResponse) SetStatusCode(v int32) *UpdateUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserGroupResponse) SetBody(v *UpdateUserGroupResponseBody) *UpdateUserGroupResponse {
	s.Body = v
	return s
}

type UpdateUserTagMetaRequest struct {
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	TagId          *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName        *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s UpdateUserTagMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTagMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserTagMetaRequest) SetTagDescription(v string) *UpdateUserTagMetaRequest {
	s.TagDescription = &v
	return s
}

func (s *UpdateUserTagMetaRequest) SetTagId(v string) *UpdateUserTagMetaRequest {
	s.TagId = &v
	return s
}

func (s *UpdateUserTagMetaRequest) SetTagName(v string) *UpdateUserTagMetaRequest {
	s.TagName = &v
	return s
}

type UpdateUserTagMetaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserTagMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTagMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserTagMetaResponseBody) SetRequestId(v string) *UpdateUserTagMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserTagMetaResponseBody) SetResult(v bool) *UpdateUserTagMetaResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateUserTagMetaResponseBody) SetSuccess(v bool) *UpdateUserTagMetaResponseBody {
	s.Success = &v
	return s
}

type UpdateUserTagMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserTagMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTagMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserTagMetaResponse) SetHeaders(v map[string]*string) *UpdateUserTagMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserTagMetaResponse) SetStatusCode(v int32) *UpdateUserTagMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserTagMetaResponse) SetBody(v *UpdateUserTagMetaResponseBody) *UpdateUserTagMetaResponse {
	s.Body = v
	return s
}

type UpdateUserTagValueRequest struct {
	TagId    *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserTagValueRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTagValueRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserTagValueRequest) SetTagId(v string) *UpdateUserTagValueRequest {
	s.TagId = &v
	return s
}

func (s *UpdateUserTagValueRequest) SetTagValue(v string) *UpdateUserTagValueRequest {
	s.TagValue = &v
	return s
}

func (s *UpdateUserTagValueRequest) SetUserId(v string) *UpdateUserTagValueRequest {
	s.UserId = &v
	return s
}

type UpdateUserTagValueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserTagValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTagValueResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserTagValueResponseBody) SetRequestId(v string) *UpdateUserTagValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserTagValueResponseBody) SetResult(v bool) *UpdateUserTagValueResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateUserTagValueResponseBody) SetSuccess(v bool) *UpdateUserTagValueResponseBody {
	s.Success = &v
	return s
}

type UpdateUserTagValueResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserTagValueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserTagValueResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserTagValueResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserTagValueResponse) SetHeaders(v map[string]*string) *UpdateUserTagValueResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserTagValueResponse) SetStatusCode(v int32) *UpdateUserTagValueResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserTagValueResponse) SetBody(v *UpdateUserTagValueResponseBody) *UpdateUserTagValueResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceUserRoleRequest struct {
	RoleId      *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceUserRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceUserRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserRoleRequest) SetRoleId(v int64) *UpdateWorkspaceUserRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateWorkspaceUserRoleRequest) SetUserId(v string) *UpdateWorkspaceUserRoleRequest {
	s.UserId = &v
	return s
}

func (s *UpdateWorkspaceUserRoleRequest) SetWorkspaceId(v string) *UpdateWorkspaceUserRoleRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateWorkspaceUserRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkspaceUserRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceUserRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserRoleResponseBody) SetRequestId(v string) *UpdateWorkspaceUserRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceUserRoleResponseBody) SetResult(v bool) *UpdateWorkspaceUserRoleResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateWorkspaceUserRoleResponseBody) SetSuccess(v bool) *UpdateWorkspaceUserRoleResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkspaceUserRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWorkspaceUserRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkspaceUserRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceUserRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserRoleResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceUserRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceUserRoleResponse) SetStatusCode(v int32) *UpdateWorkspaceUserRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceUserRoleResponse) SetBody(v *UpdateWorkspaceUserRoleResponseBody) *UpdateWorkspaceUserRoleResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceUsersRoleRequest struct {
	RoleId      *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserIds     *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceUsersRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceUsersRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUsersRoleRequest) SetRoleId(v int64) *UpdateWorkspaceUsersRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleRequest) SetUserIds(v string) *UpdateWorkspaceUsersRoleRequest {
	s.UserIds = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleRequest) SetWorkspaceId(v string) *UpdateWorkspaceUsersRoleRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateWorkspaceUsersRoleResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateWorkspaceUsersRoleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkspaceUsersRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceUsersRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUsersRoleResponseBody) SetRequestId(v string) *UpdateWorkspaceUsersRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBody) SetResult(v *UpdateWorkspaceUsersRoleResponseBodyResult) *UpdateWorkspaceUsersRoleResponseBody {
	s.Result = v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBody) SetSuccess(v bool) *UpdateWorkspaceUsersRoleResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkspaceUsersRoleResponseBodyResult struct {
	Failure       *int32                 `json:"Failure,omitempty" xml:"Failure,omitempty"`
	FailureDetail map[string]interface{} `json:"FailureDetail,omitempty" xml:"FailureDetail,omitempty"`
	Success       *int32                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total         *int32                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s UpdateWorkspaceUsersRoleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceUsersRoleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) SetFailure(v int32) *UpdateWorkspaceUsersRoleResponseBodyResult {
	s.Failure = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) SetFailureDetail(v map[string]interface{}) *UpdateWorkspaceUsersRoleResponseBodyResult {
	s.FailureDetail = v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) SetSuccess(v int32) *UpdateWorkspaceUsersRoleResponseBodyResult {
	s.Success = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponseBodyResult) SetTotal(v int32) *UpdateWorkspaceUsersRoleResponseBodyResult {
	s.Total = &v
	return s
}

type UpdateWorkspaceUsersRoleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWorkspaceUsersRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkspaceUsersRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceUsersRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUsersRoleResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceUsersRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponse) SetStatusCode(v int32) *UpdateWorkspaceUsersRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleResponse) SetBody(v *UpdateWorkspaceUsersRoleResponseBody) *UpdateWorkspaceUsersRoleResponse {
	s.Body = v
	return s
}

type WithdrawAllUserGroupsRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s WithdrawAllUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s WithdrawAllUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *WithdrawAllUserGroupsRequest) SetUserId(v string) *WithdrawAllUserGroupsRequest {
	s.UserId = &v
	return s
}

type WithdrawAllUserGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WithdrawAllUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WithdrawAllUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawAllUserGroupsResponseBody) SetRequestId(v string) *WithdrawAllUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawAllUserGroupsResponseBody) SetResult(v bool) *WithdrawAllUserGroupsResponseBody {
	s.Result = &v
	return s
}

func (s *WithdrawAllUserGroupsResponseBody) SetSuccess(v bool) *WithdrawAllUserGroupsResponseBody {
	s.Success = &v
	return s
}

type WithdrawAllUserGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WithdrawAllUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WithdrawAllUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s WithdrawAllUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *WithdrawAllUserGroupsResponse) SetHeaders(v map[string]*string) *WithdrawAllUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *WithdrawAllUserGroupsResponse) SetStatusCode(v int32) *WithdrawAllUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawAllUserGroupsResponse) SetBody(v *WithdrawAllUserGroupsResponseBody) *WithdrawAllUserGroupsResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("quickbi-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddDataLevelPermissionRuleUsersWithOptions(request *AddDataLevelPermissionRuleUsersRequest, runtime *util.RuntimeOptions) (_result *AddDataLevelPermissionRuleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddUserModel)) {
		query["AddUserModel"] = request.AddUserModel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDataLevelPermissionRuleUsers"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDataLevelPermissionRuleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDataLevelPermissionRuleUsers(request *AddDataLevelPermissionRuleUsersRequest) (_result *AddDataLevelPermissionRuleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDataLevelPermissionRuleUsersResponse{}
	_body, _err := client.AddDataLevelPermissionRuleUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDataLevelPermissionWhiteListWithOptions(request *AddDataLevelPermissionWhiteListRequest, runtime *util.RuntimeOptions) (_result *AddDataLevelPermissionWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CubeId)) {
		query["CubeId"] = request.CubeId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetIds)) {
		query["TargetIds"] = request.TargetIds
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDataLevelPermissionWhiteList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDataLevelPermissionWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDataLevelPermissionWhiteList(request *AddDataLevelPermissionWhiteListRequest) (_result *AddDataLevelPermissionWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDataLevelPermissionWhiteListResponse{}
	_body, _err := client.AddDataLevelPermissionWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddShareReportWithOptions(request *AddShareReportRequest, runtime *util.RuntimeOptions) (_result *AddShareReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthPoint)) {
		query["AuthPoint"] = request.AuthPoint
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireDate)) {
		query["ExpireDate"] = request.ExpireDate
	}

	if !tea.BoolValue(util.IsUnset(request.ShareToId)) {
		query["ShareToId"] = request.ShareToId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareToType)) {
		query["ShareToType"] = request.ShareToType
	}

	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddShareReport"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddShareReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddShareReport(request *AddShareReportRequest) (_result *AddShareReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddShareReportResponse{}
	_body, _err := client.AddShareReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserWithOptions(request *AddUserRequest, runtime *util.RuntimeOptions) (_result *AddUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AdminUser)) {
		query["AdminUser"] = request.AdminUser
	}

	if !tea.BoolValue(util.IsUnset(request.AuthAdminUser)) {
		query["AuthAdminUser"] = request.AuthAdminUser
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		query["NickName"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUser"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUser(request *AddUserRequest) (_result *AddUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserResponse{}
	_body, _err := client.AddUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserGroupMemberWithOptions(request *AddUserGroupMemberRequest, runtime *util.RuntimeOptions) (_result *AddUserGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserGroupMember"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserGroupMember(request *AddUserGroupMemberRequest) (_result *AddUserGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserGroupMemberResponse{}
	_body, _err := client.AddUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserGroupMembersWithOptions(request *AddUserGroupMembersRequest, runtime *util.RuntimeOptions) (_result *AddUserGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserGroupMembers"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserGroupMembers(request *AddUserGroupMembersRequest) (_result *AddUserGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserGroupMembersResponse{}
	_body, _err := client.AddUserGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserTagMetaWithOptions(request *AddUserTagMetaRequest, runtime *util.RuntimeOptions) (_result *AddUserTagMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagDescription)) {
		query["TagDescription"] = request.TagDescription
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserTagMeta"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserTagMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserTagMeta(request *AddUserTagMetaRequest) (_result *AddUserTagMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserTagMetaResponse{}
	_body, _err := client.AddUserTagMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserToWorkspaceWithOptions(request *AddUserToWorkspaceRequest, runtime *util.RuntimeOptions) (_result *AddUserToWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["RoleId"] = request.RoleId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToWorkspace"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserToWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserToWorkspace(request *AddUserToWorkspaceRequest) (_result *AddUserToWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToWorkspaceResponse{}
	_body, _err := client.AddUserToWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWorkspaceUsersWithOptions(request *AddWorkspaceUsersRequest, runtime *util.RuntimeOptions) (_result *AddWorkspaceUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["RoleId"] = request.RoleId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddWorkspaceUsers"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWorkspaceUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWorkspaceUsers(request *AddWorkspaceUsersRequest) (_result *AddWorkspaceUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddWorkspaceUsersResponse{}
	_body, _err := client.AddWorkspaceUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeMenuWithOptions(request *AuthorizeMenuRequest, runtime *util.RuntimeOptions) (_result *AuthorizeMenuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthPointsValue)) {
		query["AuthPointsValue"] = request.AuthPointsValue
	}

	if !tea.BoolValue(util.IsUnset(request.DataPortalId)) {
		query["DataPortalId"] = request.DataPortalId
	}

	if !tea.BoolValue(util.IsUnset(request.MenuIds)) {
		query["MenuIds"] = request.MenuIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeMenu"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeMenuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeMenu(request *AuthorizeMenuRequest) (_result *AuthorizeMenuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeMenuResponse{}
	_body, _err := client.AuthorizeMenuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelAuthorizationMenuWithOptions(request *CancelAuthorizationMenuRequest, runtime *util.RuntimeOptions) (_result *CancelAuthorizationMenuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataPortalId)) {
		query["DataPortalId"] = request.DataPortalId
	}

	if !tea.BoolValue(util.IsUnset(request.MenuIds)) {
		query["MenuIds"] = request.MenuIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelAuthorizationMenu"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelAuthorizationMenuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelAuthorizationMenu(request *CancelAuthorizationMenuRequest) (_result *CancelAuthorizationMenuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelAuthorizationMenuResponse{}
	_body, _err := client.CancelAuthorizationMenuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelCollectionWithOptions(request *CancelCollectionRequest, runtime *util.RuntimeOptions) (_result *CancelCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelCollection"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelCollection(request *CancelCollectionRequest) (_result *CancelCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelCollectionResponse{}
	_body, _err := client.CancelCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelReportShareWithOptions(request *CancelReportShareRequest, runtime *util.RuntimeOptions) (_result *CancelReportShareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareToIds)) {
		query["ShareToIds"] = request.ShareToIds
	}

	if !tea.BoolValue(util.IsUnset(request.ShareToType)) {
		query["ShareToType"] = request.ShareToType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelReportShare"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelReportShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelReportShare(request *CancelReportShareRequest) (_result *CancelReportShareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelReportShareResponse{}
	_body, _err := client.CancelReportShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeVisibilityModelWithOptions(request *ChangeVisibilityModelRequest, runtime *util.RuntimeOptions) (_result *ChangeVisibilityModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataPortalId)) {
		query["DataPortalId"] = request.DataPortalId
	}

	if !tea.BoolValue(util.IsUnset(request.MenuIds)) {
		query["MenuIds"] = request.MenuIds
	}

	if !tea.BoolValue(util.IsUnset(request.ShowOnlyWithAccess)) {
		query["ShowOnlyWithAccess"] = request.ShowOnlyWithAccess
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeVisibilityModel"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeVisibilityModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeVisibilityModel(request *ChangeVisibilityModelRequest) (_result *ChangeVisibilityModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeVisibilityModelResponse{}
	_body, _err := client.ChangeVisibilityModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckReadableWithOptions(request *CheckReadableRequest, runtime *util.RuntimeOptions) (_result *CheckReadableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckReadable"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckReadableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckReadable(request *CheckReadableRequest) (_result *CheckReadableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckReadableResponse{}
	_body, _err := client.CheckReadableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.CmptId)) {
		query["CmptId"] = request.CmptId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalParam)) {
		query["GlobalParam"] = request.GlobalParam
	}

	if !tea.BoolValue(util.IsUnset(request.TicketNum)) {
		query["TicketNum"] = request.TicketNum
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WatermarkParam)) {
		query["WatermarkParam"] = request.WatermarkParam
	}

	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTicket"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, runtime *util.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParentUserGroupId)) {
		query["ParentUserGroupId"] = request.ParentUserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupDescription)) {
		query["UserGroupDescription"] = request.UserGroupDescription
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupName)) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserGroup"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DelayTicketExpireTimeWithOptions(request *DelayTicketExpireTimeRequest, runtime *util.RuntimeOptions) (_result *DelayTicketExpireTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		query["Ticket"] = request.Ticket
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DelayTicketExpireTime"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DelayTicketExpireTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DelayTicketExpireTime(request *DelayTicketExpireTimeRequest) (_result *DelayTicketExpireTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DelayTicketExpireTimeResponse{}
	_body, _err := client.DelayTicketExpireTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataLevelPermissionRuleUsersWithOptions(request *DeleteDataLevelPermissionRuleUsersRequest, runtime *util.RuntimeOptions) (_result *DeleteDataLevelPermissionRuleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteUserModel)) {
		query["DeleteUserModel"] = request.DeleteUserModel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataLevelPermissionRuleUsers"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataLevelPermissionRuleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataLevelPermissionRuleUsers(request *DeleteDataLevelPermissionRuleUsersRequest) (_result *DeleteDataLevelPermissionRuleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataLevelPermissionRuleUsersResponse{}
	_body, _err := client.DeleteDataLevelPermissionRuleUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataLevelRuleConfigWithOptions(request *DeleteDataLevelRuleConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteDataLevelRuleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CubeId)) {
		query["CubeId"] = request.CubeId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataLevelRuleConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataLevelRuleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataLevelRuleConfig(request *DeleteDataLevelRuleConfigRequest) (_result *DeleteDataLevelRuleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataLevelRuleConfigResponse{}
	_body, _err := client.DeleteDataLevelRuleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTicketWithOptions(request *DeleteTicketRequest, runtime *util.RuntimeOptions) (_result *DeleteTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		query["Ticket"] = request.Ticket
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTicket"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTicket(request *DeleteTicketRequest) (_result *DeleteTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTicketResponse{}
	_body, _err := client.DeleteTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TransferUserId)) {
		query["TransferUserId"] = request.TransferUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserFromWorkspaceWithOptions(request *DeleteUserFromWorkspaceRequest, runtime *util.RuntimeOptions) (_result *DeleteUserFromWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserFromWorkspace"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserFromWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserFromWorkspace(request *DeleteUserFromWorkspaceRequest) (_result *DeleteUserFromWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserFromWorkspaceResponse{}
	_body, _err := client.DeleteUserFromWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserGroup"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserGroupMemberWithOptions(request *DeleteUserGroupMemberRequest, runtime *util.RuntimeOptions) (_result *DeleteUserGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserGroupMember"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserGroupMember(request *DeleteUserGroupMemberRequest) (_result *DeleteUserGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserGroupMemberResponse{}
	_body, _err := client.DeleteUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserGroupMembersWithOptions(request *DeleteUserGroupMembersRequest, runtime *util.RuntimeOptions) (_result *DeleteUserGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserGroupMembers"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserGroupMembers(request *DeleteUserGroupMembersRequest) (_result *DeleteUserGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserGroupMembersResponse{}
	_body, _err := client.DeleteUserGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserTagMetaWithOptions(request *DeleteUserTagMetaRequest, runtime *util.RuntimeOptions) (_result *DeleteUserTagMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserTagMeta"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserTagMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserTagMeta(request *DeleteUserTagMetaRequest) (_result *DeleteUserTagMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserTagMetaResponse{}
	_body, _err := client.DeleteUserTagMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserGroupInfoWithOptions(request *GetUserGroupInfoRequest, runtime *util.RuntimeOptions) (_result *GetUserGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGroupInfo"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGroupInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserGroupInfo(request *GetUserGroupInfoRequest) (_result *GetUserGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserGroupInfoResponse{}
	_body, _err := client.GetUserGroupInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListByUserGroupIdWithOptions(request *ListByUserGroupIdRequest, runtime *util.RuntimeOptions) (_result *ListByUserGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListByUserGroupId"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListByUserGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListByUserGroupId(request *ListByUserGroupIdRequest) (_result *ListByUserGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListByUserGroupIdResponse{}
	_body, _err := client.ListByUserGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCollectionsWithOptions(request *ListCollectionsRequest, runtime *util.RuntimeOptions) (_result *ListCollectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCollections"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCollections(request *ListCollectionsRequest) (_result *ListCollectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCollectionsResponse{}
	_body, _err := client.ListCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCubeDataLevelPermissionConfigWithOptions(request *ListCubeDataLevelPermissionConfigRequest, runtime *util.RuntimeOptions) (_result *ListCubeDataLevelPermissionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CubeId)) {
		query["CubeId"] = request.CubeId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCubeDataLevelPermissionConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCubeDataLevelPermissionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCubeDataLevelPermissionConfig(request *ListCubeDataLevelPermissionConfigRequest) (_result *ListCubeDataLevelPermissionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCubeDataLevelPermissionConfigResponse{}
	_body, _err := client.ListCubeDataLevelPermissionConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataLevelPermissionWhiteListWithOptions(request *ListDataLevelPermissionWhiteListRequest, runtime *util.RuntimeOptions) (_result *ListDataLevelPermissionWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CubeId)) {
		query["CubeId"] = request.CubeId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataLevelPermissionWhiteList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataLevelPermissionWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataLevelPermissionWhiteList(request *ListDataLevelPermissionWhiteListRequest) (_result *ListDataLevelPermissionWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataLevelPermissionWhiteListResponse{}
	_body, _err := client.ListDataLevelPermissionWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFavoriteReportsWithOptions(request *ListFavoriteReportsRequest, runtime *util.RuntimeOptions) (_result *ListFavoriteReportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TreeType)) {
		query["TreeType"] = request.TreeType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFavoriteReports"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFavoriteReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFavoriteReports(request *ListFavoriteReportsRequest) (_result *ListFavoriteReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFavoriteReportsResponse{}
	_body, _err := client.ListFavoriteReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPortalMenuAuthorizationWithOptions(request *ListPortalMenuAuthorizationRequest, runtime *util.RuntimeOptions) (_result *ListPortalMenuAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataPortalId)) {
		query["DataPortalId"] = request.DataPortalId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPortalMenuAuthorization"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPortalMenuAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPortalMenuAuthorization(request *ListPortalMenuAuthorizationRequest) (_result *ListPortalMenuAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPortalMenuAuthorizationResponse{}
	_body, _err := client.ListPortalMenuAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPortalMenusWithOptions(request *ListPortalMenusRequest, runtime *util.RuntimeOptions) (_result *ListPortalMenusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataPortalId)) {
		query["DataPortalId"] = request.DataPortalId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPortalMenus"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPortalMenusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPortalMenus(request *ListPortalMenusRequest) (_result *ListPortalMenusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPortalMenusResponse{}
	_body, _err := client.ListPortalMenusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecentViewReportsWithOptions(request *ListRecentViewReportsRequest, runtime *util.RuntimeOptions) (_result *ListRecentViewReportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.OffsetDay)) {
		query["OffsetDay"] = request.OffsetDay
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryMode)) {
		query["QueryMode"] = request.QueryMode
	}

	if !tea.BoolValue(util.IsUnset(request.TreeType)) {
		query["TreeType"] = request.TreeType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecentViewReports"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecentViewReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecentViewReports(request *ListRecentViewReportsRequest) (_result *ListRecentViewReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecentViewReportsResponse{}
	_body, _err := client.ListRecentViewReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSharedReportsWithOptions(request *ListSharedReportsRequest, runtime *util.RuntimeOptions) (_result *ListSharedReportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TreeType)) {
		query["TreeType"] = request.TreeType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSharedReports"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSharedReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSharedReports(request *ListSharedReportsRequest) (_result *ListSharedReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSharedReportsResponse{}
	_body, _err := client.ListSharedReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserGroupsByUserIdWithOptions(request *ListUserGroupsByUserIdRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsByUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserGroupsByUserId"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserGroupsByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserGroupsByUserId(request *ListUserGroupsByUserIdRequest) (_result *ListUserGroupsByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupsByUserIdResponse{}
	_body, _err := client.ListUserGroupsByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDataServiceWithOptions(request *QueryDataServiceRequest, runtime *util.RuntimeOptions) (_result *QueryDataServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.Conditions)) {
		query["Conditions"] = request.Conditions
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnFields)) {
		query["ReturnFields"] = request.ReturnFields
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDataService"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDataService(request *QueryDataServiceRequest) (_result *QueryDataServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDataServiceResponse{}
	_body, _err := client.QueryDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDatasetDetailInfoWithOptions(request *QueryDatasetDetailInfoRequest, runtime *util.RuntimeOptions) (_result *QueryDatasetDetailInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		query["DatasetId"] = request.DatasetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDatasetDetailInfo"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDatasetDetailInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDatasetDetailInfo(request *QueryDatasetDetailInfoRequest) (_result *QueryDatasetDetailInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDatasetDetailInfoResponse{}
	_body, _err := client.QueryDatasetDetailInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDatasetInfoWithOptions(request *QueryDatasetInfoRequest, runtime *util.RuntimeOptions) (_result *QueryDatasetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		query["DatasetId"] = request.DatasetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDatasetInfo"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDatasetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDatasetInfo(request *QueryDatasetInfoRequest) (_result *QueryDatasetInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDatasetInfoResponse{}
	_body, _err := client.QueryDatasetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDatasetListWithOptions(request *QueryDatasetListRequest, runtime *util.RuntimeOptions) (_result *QueryDatasetListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.WithChildren)) {
		query["WithChildren"] = request.WithChildren
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDatasetList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDatasetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDatasetList(request *QueryDatasetListRequest) (_result *QueryDatasetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDatasetListResponse{}
	_body, _err := client.QueryDatasetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDatasetSwitchInfoWithOptions(request *QueryDatasetSwitchInfoRequest, runtime *util.RuntimeOptions) (_result *QueryDatasetSwitchInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CubeId)) {
		query["CubeId"] = request.CubeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDatasetSwitchInfo"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDatasetSwitchInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDatasetSwitchInfo(request *QueryDatasetSwitchInfoRequest) (_result *QueryDatasetSwitchInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDatasetSwitchInfoResponse{}
	_body, _err := client.QueryDatasetSwitchInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEmbeddedInfoWithOptions(runtime *util.RuntimeOptions) (_result *QueryEmbeddedInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryEmbeddedInfo"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEmbeddedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEmbeddedInfo() (_result *QueryEmbeddedInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEmbeddedInfoResponse{}
	_body, _err := client.QueryEmbeddedInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEmbeddedStatusWithOptions(request *QueryEmbeddedStatusRequest, runtime *util.RuntimeOptions) (_result *QueryEmbeddedStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEmbeddedStatus"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEmbeddedStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEmbeddedStatus(request *QueryEmbeddedStatusRequest) (_result *QueryEmbeddedStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEmbeddedStatusResponse{}
	_body, _err := client.QueryEmbeddedStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrganizationWorkspaceListWithOptions(request *QueryOrganizationWorkspaceListRequest, runtime *util.RuntimeOptions) (_result *QueryOrganizationWorkspaceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrganizationWorkspaceList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrganizationWorkspaceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrganizationWorkspaceList(request *QueryOrganizationWorkspaceListRequest) (_result *QueryOrganizationWorkspaceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrganizationWorkspaceListResponse{}
	_body, _err := client.QueryOrganizationWorkspaceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryReadableResourcesListByUserIdWithOptions(request *QueryReadableResourcesListByUserIdRequest, runtime *util.RuntimeOptions) (_result *QueryReadableResourcesListByUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryReadableResourcesListByUserId"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryReadableResourcesListByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryReadableResourcesListByUserId(request *QueryReadableResourcesListByUserIdRequest) (_result *QueryReadableResourcesListByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryReadableResourcesListByUserIdResponse{}
	_body, _err := client.QueryReadableResourcesListByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryShareListWithOptions(request *QueryShareListRequest, runtime *util.RuntimeOptions) (_result *QueryShareListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryShareList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryShareListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryShareList(request *QueryShareListRequest) (_result *QueryShareListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryShareListResponse{}
	_body, _err := client.QueryShareListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySharesToUserListWithOptions(request *QuerySharesToUserListRequest, runtime *util.RuntimeOptions) (_result *QuerySharesToUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySharesToUserList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySharesToUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySharesToUserList(request *QuerySharesToUserListRequest) (_result *QuerySharesToUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySharesToUserListResponse{}
	_body, _err := client.QuerySharesToUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTicketInfoWithOptions(request *QueryTicketInfoRequest, runtime *util.RuntimeOptions) (_result *QueryTicketInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		query["Ticket"] = request.Ticket
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTicketInfo"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTicketInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTicketInfo(request *QueryTicketInfoRequest) (_result *QueryTicketInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTicketInfoResponse{}
	_body, _err := client.QueryTicketInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserGroupListByParentIdWithOptions(request *QueryUserGroupListByParentIdRequest, runtime *util.RuntimeOptions) (_result *QueryUserGroupListByParentIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParentUserGroupId)) {
		query["ParentUserGroupId"] = request.ParentUserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserGroupListByParentId"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserGroupListByParentIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserGroupListByParentId(request *QueryUserGroupListByParentIdRequest) (_result *QueryUserGroupListByParentIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserGroupListByParentIdResponse{}
	_body, _err := client.QueryUserGroupListByParentIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserGroupMemberWithOptions(request *QueryUserGroupMemberRequest, runtime *util.RuntimeOptions) (_result *QueryUserGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserGroupMember"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserGroupMember(request *QueryUserGroupMemberRequest) (_result *QueryUserGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserGroupMemberResponse{}
	_body, _err := client.QueryUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserInfoByAccountWithOptions(request *QueryUserInfoByAccountRequest, runtime *util.RuntimeOptions) (_result *QueryUserInfoByAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserInfoByAccount"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserInfoByAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserInfoByAccount(request *QueryUserInfoByAccountRequest) (_result *QueryUserInfoByAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserInfoByAccountResponse{}
	_body, _err := client.QueryUserInfoByAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserInfoByUserIdWithOptions(request *QueryUserInfoByUserIdRequest, runtime *util.RuntimeOptions) (_result *QueryUserInfoByUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserInfoByUserId"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserInfoByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserInfoByUserId(request *QueryUserInfoByUserIdRequest) (_result *QueryUserInfoByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserInfoByUserIdResponse{}
	_body, _err := client.QueryUserInfoByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserListWithOptions(request *QueryUserListRequest, runtime *util.RuntimeOptions) (_result *QueryUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserList(request *QueryUserListRequest) (_result *QueryUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserListResponse{}
	_body, _err := client.QueryUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserRoleInfoInWorkspaceWithOptions(request *QueryUserRoleInfoInWorkspaceRequest, runtime *util.RuntimeOptions) (_result *QueryUserRoleInfoInWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserRoleInfoInWorkspace"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserRoleInfoInWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserRoleInfoInWorkspace(request *QueryUserRoleInfoInWorkspaceRequest) (_result *QueryUserRoleInfoInWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserRoleInfoInWorkspaceResponse{}
	_body, _err := client.QueryUserRoleInfoInWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserTagMetaListWithOptions(runtime *util.RuntimeOptions) (_result *QueryUserTagMetaListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryUserTagMetaList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserTagMetaListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserTagMetaList() (_result *QueryUserTagMetaListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserTagMetaListResponse{}
	_body, _err := client.QueryUserTagMetaListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserTagValueListWithOptions(request *QueryUserTagValueListRequest, runtime *util.RuntimeOptions) (_result *QueryUserTagValueListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserTagValueList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserTagValueListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserTagValueList(request *QueryUserTagValueListRequest) (_result *QueryUserTagValueListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserTagValueListResponse{}
	_body, _err := client.QueryUserTagValueListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWorksWithOptions(request *QueryWorksRequest, runtime *util.RuntimeOptions) (_result *QueryWorksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWorks"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWorksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWorks(request *QueryWorksRequest) (_result *QueryWorksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWorksResponse{}
	_body, _err := client.QueryWorksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWorksBloodRelationshipWithOptions(request *QueryWorksBloodRelationshipRequest, runtime *util.RuntimeOptions) (_result *QueryWorksBloodRelationshipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWorksBloodRelationship"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWorksBloodRelationshipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWorksBloodRelationship(request *QueryWorksBloodRelationshipRequest) (_result *QueryWorksBloodRelationshipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWorksBloodRelationshipResponse{}
	_body, _err := client.QueryWorksBloodRelationshipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWorksByOrganizationWithOptions(request *QueryWorksByOrganizationRequest, runtime *util.RuntimeOptions) (_result *QueryWorksByOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartAuthFlag)) {
		query["ThirdPartAuthFlag"] = request.ThirdPartAuthFlag
	}

	if !tea.BoolValue(util.IsUnset(request.WorksType)) {
		query["WorksType"] = request.WorksType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWorksByOrganization"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWorksByOrganizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWorksByOrganization(request *QueryWorksByOrganizationRequest) (_result *QueryWorksByOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWorksByOrganizationResponse{}
	_body, _err := client.QueryWorksByOrganizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWorksByWorkspaceWithOptions(request *QueryWorksByWorkspaceRequest, runtime *util.RuntimeOptions) (_result *QueryWorksByWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdPartAuthFlag)) {
		query["ThirdPartAuthFlag"] = request.ThirdPartAuthFlag
	}

	if !tea.BoolValue(util.IsUnset(request.WorksType)) {
		query["WorksType"] = request.WorksType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWorksByWorkspace"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWorksByWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWorksByWorkspace(request *QueryWorksByWorkspaceRequest) (_result *QueryWorksByWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWorksByWorkspaceResponse{}
	_body, _err := client.QueryWorksByWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWorkspaceUserListWithOptions(request *QueryWorkspaceUserListRequest, runtime *util.RuntimeOptions) (_result *QueryWorkspaceUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWorkspaceUserList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWorkspaceUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWorkspaceUserList(request *QueryWorkspaceUserListRequest) (_result *QueryWorkspaceUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWorkspaceUserListResponse{}
	_body, _err := client.QueryWorkspaceUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResultCallbackWithOptions(request *ResultCallbackRequest, runtime *util.RuntimeOptions) (_result *ResultCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.HandleReason)) {
		query["HandleReason"] = request.HandleReason
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResultCallback"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResultCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResultCallback(request *ResultCallbackRequest) (_result *ResultCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResultCallbackResponse{}
	_body, _err := client.ResultCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveFavoritesWithOptions(request *SaveFavoritesRequest, runtime *util.RuntimeOptions) (_result *SaveFavoritesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveFavorites"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveFavoritesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveFavorites(request *SaveFavoritesRequest) (_result *SaveFavoritesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveFavoritesResponse{}
	_body, _err := client.SaveFavoritesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDataLevelPermissionExtraConfigWithOptions(request *SetDataLevelPermissionExtraConfigRequest, runtime *util.RuntimeOptions) (_result *SetDataLevelPermissionExtraConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CubeId)) {
		query["CubeId"] = request.CubeId
	}

	if !tea.BoolValue(util.IsUnset(request.MissHitPolicy)) {
		query["MissHitPolicy"] = request.MissHitPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDataLevelPermissionExtraConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDataLevelPermissionExtraConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDataLevelPermissionExtraConfig(request *SetDataLevelPermissionExtraConfigRequest) (_result *SetDataLevelPermissionExtraConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDataLevelPermissionExtraConfigResponse{}
	_body, _err := client.SetDataLevelPermissionExtraConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDataLevelPermissionRuleConfigWithOptions(request *SetDataLevelPermissionRuleConfigRequest, runtime *util.RuntimeOptions) (_result *SetDataLevelPermissionRuleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleModel)) {
		query["RuleModel"] = request.RuleModel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDataLevelPermissionRuleConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDataLevelPermissionRuleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDataLevelPermissionRuleConfig(request *SetDataLevelPermissionRuleConfigRequest) (_result *SetDataLevelPermissionRuleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDataLevelPermissionRuleConfigResponse{}
	_body, _err := client.SetDataLevelPermissionRuleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDataLevelPermissionWhiteListWithOptions(request *SetDataLevelPermissionWhiteListRequest, runtime *util.RuntimeOptions) (_result *SetDataLevelPermissionWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WhiteListModel)) {
		query["WhiteListModel"] = request.WhiteListModel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDataLevelPermissionWhiteList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDataLevelPermissionWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDataLevelPermissionWhiteList(request *SetDataLevelPermissionWhiteListRequest) (_result *SetDataLevelPermissionWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDataLevelPermissionWhiteListResponse{}
	_body, _err := client.SetDataLevelPermissionWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDataLevelPermissionStatusWithOptions(request *UpdateDataLevelPermissionStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateDataLevelPermissionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CubeId)) {
		query["CubeId"] = request.CubeId
	}

	if !tea.BoolValue(util.IsUnset(request.IsOpen)) {
		query["IsOpen"] = request.IsOpen
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataLevelPermissionStatus"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataLevelPermissionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDataLevelPermissionStatus(request *UpdateDataLevelPermissionStatusRequest) (_result *UpdateDataLevelPermissionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDataLevelPermissionStatusResponse{}
	_body, _err := client.UpdateDataLevelPermissionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEmbeddedStatusWithOptions(request *UpdateEmbeddedStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateEmbeddedStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ThirdPartAuthFlag)) {
		query["ThirdPartAuthFlag"] = request.ThirdPartAuthFlag
	}

	if !tea.BoolValue(util.IsUnset(request.WorksId)) {
		query["WorksId"] = request.WorksId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEmbeddedStatus"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEmbeddedStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEmbeddedStatus(request *UpdateEmbeddedStatusRequest) (_result *UpdateEmbeddedStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEmbeddedStatusResponse{}
	_body, _err := client.UpdateEmbeddedStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTicketNumWithOptions(request *UpdateTicketNumRequest, runtime *util.RuntimeOptions) (_result *UpdateTicketNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ticket)) {
		query["Ticket"] = request.Ticket
	}

	if !tea.BoolValue(util.IsUnset(request.TicketNum)) {
		query["TicketNum"] = request.TicketNum
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTicketNum"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTicketNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTicketNum(request *UpdateTicketNumRequest) (_result *UpdateTicketNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTicketNumResponse{}
	_body, _err := client.UpdateTicketNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdminUser)) {
		query["AdminUser"] = request.AdminUser
	}

	if !tea.BoolValue(util.IsUnset(request.AuthAdminUser)) {
		query["AuthAdminUser"] = request.AuthAdminUser
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		query["NickName"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserGroupWithOptions(request *UpdateUserGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupDescription)) {
		query["UserGroupDescription"] = request.UserGroupDescription
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupName)) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserGroup"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (_result *UpdateUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.UpdateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserTagMetaWithOptions(request *UpdateUserTagMetaRequest, runtime *util.RuntimeOptions) (_result *UpdateUserTagMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagDescription)) {
		query["TagDescription"] = request.TagDescription
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserTagMeta"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserTagMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserTagMeta(request *UpdateUserTagMetaRequest) (_result *UpdateUserTagMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserTagMetaResponse{}
	_body, _err := client.UpdateUserTagMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserTagValueWithOptions(request *UpdateUserTagValueRequest, runtime *util.RuntimeOptions) (_result *UpdateUserTagValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	if !tea.BoolValue(util.IsUnset(request.TagValue)) {
		query["TagValue"] = request.TagValue
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserTagValue"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserTagValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserTagValue(request *UpdateUserTagValueRequest) (_result *UpdateUserTagValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserTagValueResponse{}
	_body, _err := client.UpdateUserTagValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceUserRoleWithOptions(request *UpdateWorkspaceUserRoleRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceUserRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["RoleId"] = request.RoleId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspaceUserRole"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkspaceUserRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspaceUserRole(request *UpdateWorkspaceUserRoleRequest) (_result *UpdateWorkspaceUserRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkspaceUserRoleResponse{}
	_body, _err := client.UpdateWorkspaceUserRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceUsersRoleWithOptions(request *UpdateWorkspaceUsersRoleRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceUsersRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["RoleId"] = request.RoleId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspaceUsersRole"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkspaceUsersRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspaceUsersRole(request *UpdateWorkspaceUsersRoleRequest) (_result *UpdateWorkspaceUsersRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkspaceUsersRoleResponse{}
	_body, _err := client.UpdateWorkspaceUsersRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WithdrawAllUserGroupsWithOptions(request *WithdrawAllUserGroupsRequest, runtime *util.RuntimeOptions) (_result *WithdrawAllUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("WithdrawAllUserGroups"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &WithdrawAllUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WithdrawAllUserGroups(request *WithdrawAllUserGroupsRequest) (_result *WithdrawAllUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WithdrawAllUserGroupsResponse{}
	_body, _err := client.WithdrawAllUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
