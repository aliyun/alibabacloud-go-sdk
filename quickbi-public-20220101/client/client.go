// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDataLevelPermissionRuleUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"ruleId":"a5bb24da-***-a891683e14da","cubeId":"7c7223ae-***-3c744528014b","addModel":{"userGroups":["0d5fb19b-***-1248fc27ca51","3d2c23d4-***-f6390f325c2d"],"users":["4334***358","Huang***3fa822"]}}
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
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface. Valid values:\\n\\n	- true: The request was successful.\\n	- false: The request failed.\\n
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:\\n\\n	- true: The request was successful.\\n	- false: The request failed.\\n
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataLevelPermissionRuleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-***-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// Operation Type: You can set this parameter to one of the following values.
	//
	// 	- ADD: Add a whitelist
	//
	// 	- DELETE: deletes a whitelist.
	//
	// example:
	//
	// ADD
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The type of row-level permissions.
	//
	// 	- ROW_LEVEL: row-level permissions,
	//
	// 	- COLUMN_LEVEL: column-level permissions
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// 43342***435,1553a****41231
	TargetIds *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
	// Modify the type of the whitelist:
	//
	// 	- 1: user
	//
	// 	- 2: user group
	//
	// example:
	//
	// 1
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The scope of authorization. Valid values:
	//
	// 	- 1: view only
	//
	// 	- 3: View and export
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	AuthPoint *int32 `json:"AuthPoint,omitempty" xml:"AuthPoint,omitempty"`
	// The validity period of the share. The value is a timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1608202110838
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The ID of the person to be shared, which may be the user ID of the Quick BI or the user group ID.
	//
	// 	- If ShareToType is 0 (user), ShareTo is the user ID.
	//
	// 	- When ShareToType is set to 1 (user group), ShareTo is the user group ID.
	//
	// 	- When ShareToType=2 (organization), ShareTo is the ID of the organization.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	ShareToId *string `json:"ShareToId,omitempty" xml:"ShareToId,omitempty"`
	// The share type of the template. Valid values:
	//
	// 	- 0: user
	//
	// 	- 1: user group
	//
	// 	- 2: organization
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ShareToType *int32 `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
	// The ID of the shared work. The works here include BI portal, dashboards, spreadsheets, and self-service access.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 05739b8e-3de0-4204-9669-7f04f02522b9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddShareReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Add organization members.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// This parameter is required.
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	RoleIds  *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
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

func (s *AddUserRequest) SetRoleIds(v string) *AddUserRequest {
	s.RoleIds = &v
	return s
}

func (s *AddUserRequest) SetUserType(v int32) *AddUserRequest {
	s.UserType = &v
	return s
}

type AddUserResponseBody struct {
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddUserResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// xxxxxx@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// example:
	//
	// true
	AuthAdminUser *bool    `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	NickName      *string  `json:"NickName,omitempty" xml:"NickName,omitempty"`
	RoleIdList    []*int64 `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Repeated"`
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
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

func (s *AddUserResponseBodyResult) SetRoleIdList(v []*int64) *AddUserResponseBodyResult {
	s.RoleIdList = v
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The result of adding members to a user group is returned. Valid values:
	//
	// 	- true: The task is added.
	//
	// 	- false: The tag failed to be added.
	//
	// This parameter is required.
	//
	// example:
	//
	// 555c4cd****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e537a5****,3dadsu****
	UserIdList *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
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
	// example:
	//
	// B6141A5A-A9EF-5F16-BF34-EFB9C1CCE4F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the user groups. Separate the IDs with commas (,). Example: aGroupId,bGroupId,cGroupIds
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d5fb19b-****-****-99da-1248fc27ca51
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// The user ID of the Quick BI.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e5****37a5
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0822a7d9-****-****-****-f20163ab9b0d
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// example:
	//
	// 7AAB95D7-2E11-4FE2-94BC-858E4FC0C976
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddWorkspaceUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 2
	Failure *int32 `json:"Failure,omitempty" xml:"Failure,omitempty"`
	// example:
	//
	// {"2046274934845893" : "AE0150010001: This user already exists.", "1213444447906552" : "AE0150010001: This user already exists."}
	FailureDetail map[string]interface{} `json:"FailureDetail,omitempty" xml:"FailureDetail,omitempty"`
	// example:
	//
	// 1
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkspaceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type AllotDatasetAccelerationTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
}

func (s AllotDatasetAccelerationTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AllotDatasetAccelerationTaskRequest) GoString() string {
	return s.String()
}

func (s *AllotDatasetAccelerationTaskRequest) SetCubeId(v string) *AllotDatasetAccelerationTaskRequest {
	s.CubeId = &v
	return s
}

type AllotDatasetAccelerationTaskResponseBody struct {
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AllotDatasetAccelerationTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllotDatasetAccelerationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AllotDatasetAccelerationTaskResponseBody) SetRequestId(v string) *AllotDatasetAccelerationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllotDatasetAccelerationTaskResponseBody) SetResult(v bool) *AllotDatasetAccelerationTaskResponseBody {
	s.Result = &v
	return s
}

func (s *AllotDatasetAccelerationTaskResponseBody) SetSuccess(v bool) *AllotDatasetAccelerationTaskResponseBody {
	s.Success = &v
	return s
}

type AllotDatasetAccelerationTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllotDatasetAccelerationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllotDatasetAccelerationTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AllotDatasetAccelerationTaskResponse) GoString() string {
	return s.String()
}

func (s *AllotDatasetAccelerationTaskResponse) SetHeaders(v map[string]*string) *AllotDatasetAccelerationTaskResponse {
	s.Headers = v
	return s
}

func (s *AllotDatasetAccelerationTaskResponse) SetStatusCode(v int32) *AllotDatasetAccelerationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AllotDatasetAccelerationTaskResponse) SetBody(v *AllotDatasetAccelerationTaskResponseBody) *AllotDatasetAccelerationTaskResponse {
	s.Body = v
	return s
}

type AuthorizeMenuRequest struct {
	// Authorizes the permissions of the menu. Valid values:
	//
	// 	- 1: view
	//
	// 	- 3: View + Export (default)
	//
	// example:
	//
	// 3
	AuthPointsValue *int32 `json:"AuthPointsValue,omitempty" xml:"AuthPointsValue,omitempty"`
	// The ID of the BI portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// The menu ID of the BI portal leaf node.
	//
	// 	- The directory menu cannot be authorized.
	//
	// 	- You can upload multiple parameters at a time. Separate multiple IDs with commas (,). The maximum number of parameters that can be modified at a time is 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54kqgoa****,pg1n135****
	MenuIds *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	// The IDs of the user groups.
	//
	// 	- You can upload multiple parameters at a time. Separate multiple IDs with commas (,). The maximum number of parameters that can be modified at a time is 200.
	//
	// 	- UserGroupIds and UserIds cannot be empty at the same time
	//
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****,3d2c23d4-2b41-4af8-a1f5-f6390f32****
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// The IDs of the end users. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// 	- You can upload multiple parameters at a time. Separate multiple IDs with commas (,). The maximum number of parameters that can be modified at a time is 200.
	//
	// example:
	//
	// 204627493484****,121344444790****
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 188F0B12-00EF-41B3-944A-FB7EF06C9F43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of authorized menus.
	//
	// example:
	//
	// 2
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeMenuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type BatchAddFeishuUsersRequest struct {
	FeishuUsers *string `json:"FeishuUsers,omitempty" xml:"FeishuUsers,omitempty"`
	// example:
	//
	// False
	IsAdmin *bool `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	// example:
	//
	// true
	IsAuthAdmin *bool `json:"IsAuthAdmin,omitempty" xml:"IsAuthAdmin,omitempty"`
	// example:
	//
	// "0d5fb19b-5555-41f0-99da-1248fc27ca51,0f868dd6_68dd_4d13_8422_c5dca3bd4b61"
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s BatchAddFeishuUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFeishuUsersRequest) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersRequest) SetFeishuUsers(v string) *BatchAddFeishuUsersRequest {
	s.FeishuUsers = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) SetIsAdmin(v bool) *BatchAddFeishuUsersRequest {
	s.IsAdmin = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) SetIsAuthAdmin(v bool) *BatchAddFeishuUsersRequest {
	s.IsAuthAdmin = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) SetUserGroupIds(v string) *BatchAddFeishuUsersRequest {
	s.UserGroupIds = &v
	return s
}

func (s *BatchAddFeishuUsersRequest) SetUserType(v int32) *BatchAddFeishuUsersRequest {
	s.UserType = &v
	return s
}

type BatchAddFeishuUsersResponseBody struct {
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Result *BatchAddFeishuUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchAddFeishuUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFeishuUsersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponseBody) SetRequestId(v string) *BatchAddFeishuUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBody) SetResult(v *BatchAddFeishuUsersResponseBodyResult) *BatchAddFeishuUsersResponseBody {
	s.Result = v
	return s
}

func (s *BatchAddFeishuUsersResponseBody) SetSuccess(v bool) *BatchAddFeishuUsersResponseBody {
	s.Success = &v
	return s
}

type BatchAddFeishuUsersResponseBodyResult struct {
	// example:
	//
	// 10
	FailCount   *int32                                              `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	FailResults []*BatchAddFeishuUsersResponseBodyResultFailResults `json:"FailResults,omitempty" xml:"FailResults,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	OkCount *int32 `json:"OkCount,omitempty" xml:"OkCount,omitempty"`
}

func (s BatchAddFeishuUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFeishuUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponseBodyResult) SetFailCount(v int32) *BatchAddFeishuUsersResponseBodyResult {
	s.FailCount = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResult) SetFailResults(v []*BatchAddFeishuUsersResponseBodyResultFailResults) *BatchAddFeishuUsersResponseBodyResult {
	s.FailResults = v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResult) SetOkCount(v int32) *BatchAddFeishuUsersResponseBodyResult {
	s.OkCount = &v
	return s
}

type BatchAddFeishuUsersResponseBodyResultFailResults struct {
	FailInfos []*BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos `json:"FailInfos,omitempty" xml:"FailInfos,omitempty" type:"Repeated"`
}

func (s BatchAddFeishuUsersResponseBodyResultFailResults) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFeishuUsersResponseBodyResultFailResults) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResults) SetFailInfos(v []*BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) *BatchAddFeishuUsersResponseBodyResultFailResults {
	s.FailInfos = v
	return s
}

type BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos struct {
	// example:
	//
	// ACCOUNT_EXIST
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeDesc *string `json:"CodeDesc,omitempty" xml:"CodeDesc,omitempty"`
	// example:
	//
	// 20
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
}

func (s BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) SetCode(v string) *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos {
	s.Code = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) SetCodeDesc(v string) *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos {
	s.CodeDesc = &v
	return s
}

func (s *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos) SetInput(v string) *BatchAddFeishuUsersResponseBodyResultFailResultsFailInfos {
	s.Input = &v
	return s
}

type BatchAddFeishuUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddFeishuUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddFeishuUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddFeishuUsersResponse) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponse) SetHeaders(v map[string]*string) *BatchAddFeishuUsersResponse {
	s.Headers = v
	return s
}

func (s *BatchAddFeishuUsersResponse) SetStatusCode(v int32) *BatchAddFeishuUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddFeishuUsersResponse) SetBody(v *BatchAddFeishuUsersResponseBody) *BatchAddFeishuUsersResponse {
	s.Body = v
	return s
}

type CancelAuthorizationMenuRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 54kqgoa****,pg1n135****
	MenuIds *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****,3d2c23d4-2b41-4af8-a1f5-f6390f32****
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// example:
	//
	// 204627493484****,121344444790****
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
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
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAuthorizationMenuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the favorite user. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121344444790****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the work to cancel the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5d6ae4e7-cede-43cd-b4d3-d2fd442a9202
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
	// The ID of the request.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the work. The works here include BI portal, dashboards, spreadsheets, and self-service access.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the person to be shared, which may be the user ID of the Quick BI or the user group ID.
	//
	// 	- If ShareToType is 0 (user), ShareTo is the user ID.
	//
	// 	- When ShareToType is set to 1 (user group), ShareTo is the user group ID.
	//
	// 	- When ShareToType=2 (organization), ShareTo is the ID of the organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	ShareToIds *string `json:"ShareToIds,omitempty" xml:"ShareToIds,omitempty"`
	// The deletion method. Valid values:
	//
	// 	- 0: Delete by user
	//
	// 	- 1: Delete by user group
	//
	// 	- 2: Delete by organization
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ShareToType *int32 `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelReportShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The number of menus that are successfully modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54kqgoa****,pg1n135****
	MenuIds *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	ShowOnlyWithAccess *bool `json:"ShowOnlyWithAccess,omitempty" xml:"ShowOnlyWithAccess,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeVisibilityModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The user ID of the Quick BI to be checked.
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the work. Resources here include BI portal, dashboards, spreadsheets, and self-service access.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckReadableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// 0fc6a275c7f64f17b1****a306ce0f31
	CmptId *string `json:"CmptId,omitempty" xml:"CmptId,omitempty"`
	// example:
	//
	// 200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// [{"paramKey":"price","joinType":"and","conditionList":[{"operate":">","value":"0"}]}]
	GlobalParam *string `json:"GlobalParam,omitempty" xml:"GlobalParam,omitempty"`
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
	// example:
	//
	// 46e537466****92704c8
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WatermarkParam *string `json:"WatermarkParam,omitempty" xml:"WatermarkParam,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a206f5f3-****-e9b17c835b03
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ccd3428c-****-****-a608-26bae29dffee
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateTicket4CopilotRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccd3428c-dd23-460c-a608-26bae29dffee
	CopilotId *string `json:"CopilotId,omitempty" xml:"CopilotId,omitempty"`
	// example:
	//
	// 200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
	// example:
	//
	// 9c-asdawf-casxcasd-asdasd
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateTicket4CopilotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicket4CopilotRequest) GoString() string {
	return s.String()
}

func (s *CreateTicket4CopilotRequest) SetAccountName(v string) *CreateTicket4CopilotRequest {
	s.AccountName = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetAccountType(v int32) *CreateTicket4CopilotRequest {
	s.AccountType = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetCopilotId(v string) *CreateTicket4CopilotRequest {
	s.CopilotId = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetExpireTime(v int32) *CreateTicket4CopilotRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetTicketNum(v int32) *CreateTicket4CopilotRequest {
	s.TicketNum = &v
	return s
}

func (s *CreateTicket4CopilotRequest) SetUserId(v string) *CreateTicket4CopilotRequest {
	s.UserId = &v
	return s
}

type CreateTicket4CopilotResponseBody struct {
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicket4CopilotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicket4CopilotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicket4CopilotResponseBody) SetRequestId(v string) *CreateTicket4CopilotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicket4CopilotResponseBody) SetResult(v string) *CreateTicket4CopilotResponseBody {
	s.Result = &v
	return s
}

func (s *CreateTicket4CopilotResponseBody) SetSuccess(v bool) *CreateTicket4CopilotResponseBody {
	s.Success = &v
	return s
}

type CreateTicket4CopilotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTicket4CopilotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTicket4CopilotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicket4CopilotResponse) GoString() string {
	return s.String()
}

func (s *CreateTicket4CopilotResponse) SetHeaders(v map[string]*string) *CreateTicket4CopilotResponse {
	s.Headers = v
	return s
}

func (s *CreateTicket4CopilotResponse) SetStatusCode(v int32) *CreateTicket4CopilotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicket4CopilotResponse) SetBody(v *CreateTicket4CopilotResponseBody) *CreateTicket4CopilotResponse {
	s.Body = v
	return s
}

type CreateUserGroupRequest struct {
	// The ID of the parent user group. You can add new user groups to this group:
	//
	// 	- If you enter the ID of a parent user group, the new user group is added to the user group with the ID.
	//
	// 	- If you enter -1, the new user group is added to the root directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3d2c23d4-2b41-4af8-a1f5-f6390f32****
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	// The description of the user group.
	//
	// 	- Format verification: Maximum length 255
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// example:
	//
	// User group description
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	// The unique ID of the user group.
	//
	// 	- If you specify the UserGroupId parameter, the system automatically generates the UserGroupId parameter. If you specify the UserGroupId parameter, the user ID is used as the user group ID. You must ensure that the user ID is unique within the organization.
	//
	// 	- Format verification: Maximum length 64, cannot be -1,
	//
	// example:
	//
	// pop0001
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the RAM user group.
	//
	// 	- Format verification: Maximum length 255
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou Financial Report
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 36829379-0C38-5BC0-830A-92665BF77D4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the added user group is returned. An empty string \\"\\" is returned if the add fails.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DataSetBloodRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 234235234,234235235,234235235
	DataSetIds *string `json:"DataSetIds,omitempty" xml:"DataSetIds,omitempty"`
	// example:
	//
	// dasasgaj342351
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// PAGE
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
}

func (s DataSetBloodRequest) String() string {
	return tea.Prettify(s)
}

func (s DataSetBloodRequest) GoString() string {
	return s.String()
}

func (s *DataSetBloodRequest) SetDataSetIds(v string) *DataSetBloodRequest {
	s.DataSetIds = &v
	return s
}

func (s *DataSetBloodRequest) SetUserId(v string) *DataSetBloodRequest {
	s.UserId = &v
	return s
}

func (s *DataSetBloodRequest) SetWorksType(v string) *DataSetBloodRequest {
	s.WorksType = &v
	return s
}

type DataSetBloodResponseBody struct {
	// example:
	//
	// 46e537a5****,3dadsu****
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DataSetBloodResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DataSetBloodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DataSetBloodResponseBody) GoString() string {
	return s.String()
}

func (s *DataSetBloodResponseBody) SetRequestId(v string) *DataSetBloodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DataSetBloodResponseBody) SetResult(v []*DataSetBloodResponseBodyResult) *DataSetBloodResponseBody {
	s.Result = v
	return s
}

func (s *DataSetBloodResponseBody) SetSuccess(v bool) *DataSetBloodResponseBody {
	s.Success = &v
	return s
}

type DataSetBloodResponseBodyResult struct {
	// example:
	//
	// ccd3428c-****-****-a608-26bae29dffee
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// example:
	//
	// PAGE
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
}

func (s DataSetBloodResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DataSetBloodResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DataSetBloodResponseBodyResult) SetWorksId(v string) *DataSetBloodResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *DataSetBloodResponseBodyResult) SetWorksType(v string) *DataSetBloodResponseBodyResult {
	s.WorksType = &v
	return s
}

type DataSetBloodResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DataSetBloodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DataSetBloodResponse) String() string {
	return tea.Prettify(s)
}

func (s DataSetBloodResponse) GoString() string {
	return s.String()
}

func (s *DataSetBloodResponse) SetHeaders(v map[string]*string) *DataSetBloodResponse {
	s.Headers = v
	return s
}

func (s *DataSetBloodResponse) SetStatusCode(v int32) *DataSetBloodResponse {
	s.StatusCode = &v
	return s
}

func (s *DataSetBloodResponse) SetBody(v *DataSetBloodResponseBody) *DataSetBloodResponse {
	s.Body = v
	return s
}

type DataSourceBloodRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 44051300991327000048
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s DataSourceBloodRequest) String() string {
	return tea.Prettify(s)
}

func (s DataSourceBloodRequest) GoString() string {
	return s.String()
}

func (s *DataSourceBloodRequest) SetDataSourceId(v string) *DataSourceBloodRequest {
	s.DataSourceId = &v
	return s
}

type DataSourceBloodResponseBody struct {
	// example:
	//
	// 46e537a5****,3dadsu****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DataSourceBloodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DataSourceBloodResponseBody) GoString() string {
	return s.String()
}

func (s *DataSourceBloodResponseBody) SetRequestId(v string) *DataSourceBloodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DataSourceBloodResponseBody) SetResult(v []*string) *DataSourceBloodResponseBody {
	s.Result = v
	return s
}

func (s *DataSourceBloodResponseBody) SetSuccess(v bool) *DataSourceBloodResponseBody {
	s.Success = &v
	return s
}

type DataSourceBloodResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DataSourceBloodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DataSourceBloodResponse) String() string {
	return tea.Prettify(s)
}

func (s DataSourceBloodResponse) GoString() string {
	return s.String()
}

func (s *DataSourceBloodResponse) SetHeaders(v map[string]*string) *DataSourceBloodResponse {
	s.Headers = v
	return s
}

func (s *DataSourceBloodResponse) SetStatusCode(v int32) *DataSourceBloodResponse {
	s.StatusCode = &v
	return s
}

func (s *DataSourceBloodResponse) SetBody(v *DataSourceBloodResponseBody) *DataSourceBloodResponse {
	s.Body = v
	return s
}

type DelayTicketExpireTimeRequest struct {
	// The time to postpone.
	//
	// 	- Unit: minutes. Valid values: 0 to 240. Unit: minutes. Valid values: 4 hours.
	//
	// 	- Expired bills cannot be extended.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The value of the third-party embedded ticket, that is, the accessTicket value in the URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// 040e6f79d33444838e*****c7206c070
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the extension is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelayTicketExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// {"ruleId":"a5bb24da-***-a891683e14da","cubeId":"7c7223ae-***-3c744528014b","delModel":{"userGroups":["0d5fb19b-***-1248fc27ca51","3d2c23d4-***-f6390f325c2d"],"users":["4334***358","Huang***3fa822"]}}
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
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLevelPermissionRuleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a5bb24da-****-a891683e14da
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
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLevelRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Deletes a specified ticket from an embedded report.
	//
	// This parameter is required.
	//
	// example:
	//
	// 040e6f79d****7d283c7206c070
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// f5****afccd9e434a274
	TransferUserId *string `json:"TransferUserId,omitempty" xml:"TransferUserId,omitempty"`
	// Deletes a user from a specified organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserFromWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
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
	// The ID of the request.
	//
	// example:
	//
	// F2775AB6-DE99-5FA6-86A4-72EA0A8AFEE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e537****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user ID of the Quick BI.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2fe4fbd8****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of deleting a user group member. Valid values:
	//
	// 	- true: The task is deleted.
	//
	// 	- false: The deletion failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****,3d2c23d4-2b41-4af8-a1f5-f6390f32****
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 204627493484****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// ABBAD906-3925-5D18-B23D-714053AB0AA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// pop_001
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// The user group modifier. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
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
	// example:
	//
	// D7980306-1F08-5A88-9FE7-ECB8B9C4C0F5
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetUserGroupInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 2021-03-15 17:13:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	CreateUser     *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	IdentifiedPath *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	// example:
	//
	// 2021-03-15 20:36:40
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
	ParentUsergroupId *string `json:"ParentUsergroupId,omitempty" xml:"ParentUsergroupId,omitempty"`
	UsergroupDesc     *string `json:"UsergroupDesc,omitempty" xml:"UsergroupDesc,omitempty"`
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****
	UsergroupId   *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	UsergroupName *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListApiDatasourceRequest struct {
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListApiDatasourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApiDatasourceRequest) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceRequest) SetKeyWord(v string) *ListApiDatasourceRequest {
	s.KeyWord = &v
	return s
}

func (s *ListApiDatasourceRequest) SetPageNum(v int32) *ListApiDatasourceRequest {
	s.PageNum = &v
	return s
}

func (s *ListApiDatasourceRequest) SetPageSize(v int32) *ListApiDatasourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListApiDatasourceRequest) SetWorkspaceId(v string) *ListApiDatasourceRequest {
	s.WorkspaceId = &v
	return s
}

type ListApiDatasourceResponseBody struct {
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListApiDatasourceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListApiDatasourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApiDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceResponseBody) SetRequestId(v string) *ListApiDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiDatasourceResponseBody) SetResult(v *ListApiDatasourceResponseBodyResult) *ListApiDatasourceResponseBody {
	s.Result = v
	return s
}

func (s *ListApiDatasourceResponseBody) SetSuccess(v bool) *ListApiDatasourceResponseBody {
	s.Success = &v
	return s
}

type ListApiDatasourceResponseBodyResult struct {
	Data []*ListApiDatasourceResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListApiDatasourceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListApiDatasourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceResponseBodyResult) SetData(v []*ListApiDatasourceResponseBodyResultData) *ListApiDatasourceResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListApiDatasourceResponseBodyResult) SetPageNum(v int32) *ListApiDatasourceResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResult) SetPageSize(v int32) *ListApiDatasourceResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResult) SetTotalNum(v int32) *ListApiDatasourceResponseBodyResult {
	s.TotalNum = &v
	return s
}

type ListApiDatasourceResponseBodyResultData struct {
	// example:
	//
	// 0f2c3c6409be4dc0810f2a5785e816a8
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// {"key1":"value1"}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// example:
	//
	// 0.39746094
	DataSize *float32 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// 2022-05-25 16:19:43
	DateUpdateTime *string `json:"DateUpdateTime,omitempty" xml:"DateUpdateTime,omitempty"`
	// example:
	//
	// 2022-05-25 16:19:43
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-05-25 16:19:43
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// REST_API_SYNC_0f2c3c6409be4dc0810f2a5785e816a8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// [{"name":"token","value":"xxxxxxxxxxxx"},{"name":"pageSize","value":100}]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ShowName   *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// example:
	//
	// 2
	StatusType *int32 `json:"StatusType,omitempty" xml:"StatusType,omitempty"`
}

func (s ListApiDatasourceResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListApiDatasourceResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceResponseBodyResultData) SetApiId(v string) *ListApiDatasourceResponseBodyResultData {
	s.ApiId = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetBody(v string) *ListApiDatasourceResponseBodyResultData {
	s.Body = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetDataSize(v float32) *ListApiDatasourceResponseBodyResultData {
	s.DataSize = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetDateUpdateTime(v string) *ListApiDatasourceResponseBodyResultData {
	s.DateUpdateTime = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetGmtCreate(v string) *ListApiDatasourceResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetGmtModified(v string) *ListApiDatasourceResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetJobId(v string) *ListApiDatasourceResponseBodyResultData {
	s.JobId = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetParameters(v string) *ListApiDatasourceResponseBodyResultData {
	s.Parameters = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetShowName(v string) *ListApiDatasourceResponseBodyResultData {
	s.ShowName = &v
	return s
}

func (s *ListApiDatasourceResponseBodyResultData) SetStatusType(v int32) *ListApiDatasourceResponseBodyResultData {
	s.StatusType = &v
	return s
}

type ListApiDatasourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiDatasourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApiDatasourceResponse) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceResponse) SetHeaders(v map[string]*string) *ListApiDatasourceResponse {
	s.Headers = v
	return s
}

func (s *ListApiDatasourceResponse) SetStatusCode(v int32) *ListApiDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiDatasourceResponse) SetBody(v *ListApiDatasourceResponseBody) *ListApiDatasourceResponse {
	s.Body = v
	return s
}

type ListByUserGroupIdRequest struct {
	// The ID of the user group that you want to query. Separate multiple user groups with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 34fe-***-6dcb,84q9-****-4a274
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
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user group query result is returned.
	Result *ListByUserGroupIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	FailedUserGroupIds []*string `json:"FailedUserGroupIds,omitempty" xml:"FailedUserGroupIds,omitempty" type:"Repeated"`
	// The details of the user group that was queried.
	UserGroupModels []*ListByUserGroupIdResponseBodyResultUserGroupModels `json:"UserGroupModels,omitempty" xml:"UserGroupModels,omitempty" type:"Repeated"`
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
	// The time when the Secret was created.
	//
	// example:
	//
	// 2021-03-15 17:13:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The UserID of the creator in the Quick BI.
	//
	// example:
	//
	// 46e5*******ee22e2a292704c8
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The path of the user group.
	//
	// example:
	//
	// 2fe4fbd8-****-af083ea/34fd1-****-dcbc33f
	IdentifiedPath *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	// The time when the protection policy was last modified.
	//
	// example:
	//
	// 2021-03-15 20:36:40
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The UserID of the modifier in the Quick BI.
	//
	// example:
	//
	// 46e5*******ee22e2a292704c8
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The ID of the parent user group.
	//
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af083ea
	ParentUsergroupId *string `json:"ParentUsergroupId,omitempty" xml:"ParentUsergroupId,omitempty"`
	// The description of the user group.
	//
	// example:
	//
	// Description
	UsergroupDesc *string `json:"UsergroupDesc,omitempty" xml:"UsergroupDesc,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// 34fd141d-****-4093-8c33-8e066dcbc33f
	UsergroupId *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// Test user group
	UsergroupName *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListByUserGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the user. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
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
	// The ID of the request.
	//
	// example:
	//
	// 162A632E-0A88-51CF-98F8-94FDEE82DB7D
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCollectionsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The primary key ID of the favorite record.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 12373
	FavoriteId *int32 `json:"FavoriteId,omitempty" xml:"FavoriteId,omitempty"`
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorksId   *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	WorksName *string `json:"WorksName,omitempty" xml:"WorksName,omitempty"`
	// example:
	//
	// dashboardOfflineQuery
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
	// example:
	//
	// 9337d121-a78f-4c1b-a8bc-f81de117****
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// The type of the dataset row and column permission. Valid values:
	//
	// 	- ROW_LEVEL: row-level permissions
	//
	// 	- COLUMN_LEVEL: column-level permissions
	//
	// This parameter is required.
	//
	// example:
	//
	// ROW_LEVEL
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
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// { "isOpen": 1, "extraConfigModel": { // Additional configuration information "ruleType": "ROW_LEVEL", // The row-level permission type. "missHitPolicy": "NONE", // The hit rule policy: NONE has no permissions, and ALL has permissions. "cubeId": "7c7223 ae-31d1-4d2f-b11f-3c744528014b" // The ID of the dataset. }, "ruleType": "ROW_LEVEL", // Row-column permission type\\
	//
	// "ruleModels": [ { "ruleUsersModel": { // The target population. "userGroups": [ "0d5fb19b- ****-1248 fc27ca51", // The ID of the user group. "4aa3f089-****-85f0-0e8ac7c2dee9" ], "users": [ "HuangJia ***2e3fa822", // The ID of the user. "4334***84358" ] }, "ruleContentModel": { "ruleContentType": "ROW_FIELD", // The row-column permission type. "ruleContentJson": "{"conditionNode":{"caption": " Period ","isMeasure":false,"pathId":"7d3b073bc6","relationOperator":"not-null","name":"7d3b073bc6","value":{"value":[""}UM]," ENueType "} // The JSON string of the row-column permission rule. "ruleOriginConfigJson": "{"operator":"and","operands":[{"labelName": " Period ","isValid":true,"uniqueId":"5","fieldId":"7d3b073bc6","error":false,"fieldType":"string",": default "" value":{"conditionOp":"not-null","conditionValue":""},"valueType":"ENUM"}}],"isRelation":true}" }, // The fixed-format JSON string required by the frontend "isOpen": 1, // The status of the row-column permission configuration. 1. On. 0. Off. "hitTakeEffect": 1, // Specifies whether the rule takes effect after a column-level permission is hit. 1 takes effect and 0 takes effect. "ruleName": "Test row-level permission_Do not delete", // The name of the row-column permission rule. "ruleLevelType": "ROW_LEVEL", // The row-column permission type. "ruleId": "a5bb24 da-772f-45e8-a43c-a891683e14da", // The ID of the row-column permission rule. "cubeId": "7c7223 ae-31d1-4d2f-b11f-3c744528014b", // The ID of the dataset. "ruleTargetScope": "OTHERS" rule takes effect: ALL owner and OTHERS designated owner. } ], "cubeId": "7c7223 ae-31d1-4d2f-b11f-3c744528014b" // The ID of the dataset. }
	//
	// example:
	//
	// The JSON string of the row-column permission list. For more information, see the description.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCubeDataLevelPermissionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 3d5db23c-e4f2-49dd-a883-92285b48e14a
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ROW_LEVEL
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
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListDataLevelPermissionWhiteListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// example:
	//
	// ROW_LEVEL
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// PAGE
	TreeType *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListFavoriteReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Data []*ListFavoriteReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// example:
	//
	// true
	Favorite     *bool   `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	FavoriteDate *string `json:"FavoriteDate,omitempty" xml:"FavoriteDate,omitempty"`
	// example:
	//
	// 1640088615000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1640595729000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// true
	HasEditAuth *bool `json:"HasEditAuth,omitempty" xml:"HasEditAuth,omitempty"`
	// example:
	//
	// true
	HasViewAuth *bool   `json:"HasViewAuth,omitempty" xml:"HasViewAuth,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName   *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1365*****238860
	OwnerNum *string `json:"OwnerNum,omitempty" xml:"OwnerNum,omitempty"`
	// example:
	//
	// 1
	PublishStatus *int32 `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// example:
	//
	// 977c7698-****-****-****-44b7304d20fc
	TreeId *string `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
	// example:
	//
	// PAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 523793cb-****-****-****-aa71c65ffa39
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

func (s *ListFavoriteReportsResponseBodyResultData) SetFavoriteDate(v string) *ListFavoriteReportsResponseBodyResultData {
	s.FavoriteDate = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFavoriteReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListOrganizationRoleUsersRequest struct {
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111111111
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ListOrganizationRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersRequest) SetKeyword(v string) *ListOrganizationRoleUsersRequest {
	s.Keyword = &v
	return s
}

func (s *ListOrganizationRoleUsersRequest) SetPageNum(v int32) *ListOrganizationRoleUsersRequest {
	s.PageNum = &v
	return s
}

func (s *ListOrganizationRoleUsersRequest) SetPageSize(v int32) *ListOrganizationRoleUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationRoleUsersRequest) SetRoleId(v int64) *ListOrganizationRoleUsersRequest {
	s.RoleId = &v
	return s
}

type ListOrganizationRoleUsersResponseBody struct {
	// example:
	//
	// BCE45E6D-9304-4F94-86BB-5A772B1615FF
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListOrganizationRoleUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOrganizationRoleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersResponseBody) SetRequestId(v string) *ListOrganizationRoleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBody) SetResult(v *ListOrganizationRoleUsersResponseBodyResult) *ListOrganizationRoleUsersResponseBody {
	s.Result = v
	return s
}

func (s *ListOrganizationRoleUsersResponseBody) SetSuccess(v bool) *ListOrganizationRoleUsersResponseBody {
	s.Success = &v
	return s
}

type ListOrganizationRoleUsersResponseBodyResult struct {
	Data []*ListOrganizationRoleUsersResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListOrganizationRoleUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRoleUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetData(v []*ListOrganizationRoleUsersResponseBodyResultData) *ListOrganizationRoleUsersResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetPageNum(v int32) *ListOrganizationRoleUsersResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetPageSize(v int32) *ListOrganizationRoleUsersResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetTotalNum(v int32) *ListOrganizationRoleUsersResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResult) SetTotalPages(v int32) *ListOrganizationRoleUsersResponseBodyResult {
	s.TotalPages = &v
	return s
}

type ListOrganizationRoleUsersResponseBodyResultData struct {
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListOrganizationRoleUsersResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRoleUsersResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersResponseBodyResultData) SetNickName(v string) *ListOrganizationRoleUsersResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *ListOrganizationRoleUsersResponseBodyResultData) SetUserId(v string) *ListOrganizationRoleUsersResponseBodyResultData {
	s.UserId = &v
	return s
}

type ListOrganizationRoleUsersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersResponse) SetHeaders(v map[string]*string) *ListOrganizationRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationRoleUsersResponse) SetStatusCode(v int32) *ListOrganizationRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationRoleUsersResponse) SetBody(v *ListOrganizationRoleUsersResponseBody) *ListOrganizationRoleUsersResponse {
	s.Body = v
	return s
}

type ListOrganizationRolesResponseBody struct {
	// example:
	//
	// 7AAB95D7-2E11-4FE2-94BC-858E4FC0C976
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListOrganizationRolesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOrganizationRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationRolesResponseBody) SetRequestId(v string) *ListOrganizationRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationRolesResponseBody) SetResult(v []*ListOrganizationRolesResponseBodyResult) *ListOrganizationRolesResponseBody {
	s.Result = v
	return s
}

func (s *ListOrganizationRolesResponseBody) SetSuccess(v bool) *ListOrganizationRolesResponseBody {
	s.Success = &v
	return s
}

type ListOrganizationRolesResponseBodyResult struct {
	AuthConfigList []*ListOrganizationRolesResponseBodyResultAuthConfigList `json:"AuthConfigList,omitempty" xml:"AuthConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsSystemRole *bool `json:"IsSystemRole,omitempty" xml:"IsSystemRole,omitempty"`
	// example:
	//
	// 111111111
	RoleId   *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListOrganizationRolesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRolesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrganizationRolesResponseBodyResult) SetAuthConfigList(v []*ListOrganizationRolesResponseBodyResultAuthConfigList) *ListOrganizationRolesResponseBodyResult {
	s.AuthConfigList = v
	return s
}

func (s *ListOrganizationRolesResponseBodyResult) SetIsSystemRole(v bool) *ListOrganizationRolesResponseBodyResult {
	s.IsSystemRole = &v
	return s
}

func (s *ListOrganizationRolesResponseBodyResult) SetRoleId(v int64) *ListOrganizationRolesResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *ListOrganizationRolesResponseBodyResult) SetRoleName(v string) *ListOrganizationRolesResponseBodyResult {
	s.RoleName = &v
	return s
}

type ListOrganizationRolesResponseBodyResultAuthConfigList struct {
	// example:
	//
	// enterprise_safety
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
}

func (s ListOrganizationRolesResponseBodyResultAuthConfigList) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRolesResponseBodyResultAuthConfigList) GoString() string {
	return s.String()
}

func (s *ListOrganizationRolesResponseBodyResultAuthConfigList) SetAuthKey(v string) *ListOrganizationRolesResponseBodyResultAuthConfigList {
	s.AuthKey = &v
	return s
}

type ListOrganizationRolesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationRolesResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationRolesResponse) SetHeaders(v map[string]*string) *ListOrganizationRolesResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationRolesResponse) SetStatusCode(v int32) *ListOrganizationRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationRolesResponse) SetBody(v *ListOrganizationRolesResponseBody) *ListOrganizationRolesResponse {
	s.Body = v
	return s
}

type ListPortalMenuAuthorizationRequest struct {
	// The ID of the BI portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
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
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of authorization details of the portal menu.
	Result []*ListPortalMenuAuthorizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The menu ID of the BI portal leaf node.
	//
	// example:
	//
	// 54kqgoa****
	MenuId *string `json:"MenuId,omitempty" xml:"MenuId,omitempty"`
	// The details of the object to which the menu is authorized.
	Receivers []*ListPortalMenuAuthorizationResponseBodyResultReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// Whether only authorization is visible. Valid values:
	//
	// 	- true: Only the authorization is visible.
	//
	// 	- false: Both are visible.
	//
	// example:
	//
	// true
	ShowOnlyWithAccess *bool `json:"ShowOnlyWithAccess,omitempty" xml:"ShowOnlyWithAccess,omitempty"`
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
	// The ID of the authorization object.
	//
	// example:
	//
	// 121344444790****
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// The type of the authorization object. Valid values:
	//
	// 	- 0: user
	//
	// 	- 1: user group
	//
	// example:
	//
	// 0
	ReceiverType *int32 `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPortalMenuAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the BI portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// The user ID in the Quick BI. When passed in, the list displays only the menus that the user has permissions on.
	//
	// example:
	//
	// 1234567***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 75912036-5527-4B7E-9265-B481D6651AC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A JSON string that levels the details of the portal menu list. Valid values:
	//
	// 	- menuType: the type of the menu.
	//
	//     	- 0: dashboard
	//
	//     	- 1: outer chain
	//
	//     	- 2: workbook
	//
	//     	- 4: directory folder
	//
	//     	- 5: form filling
	//
	//     	- 6: self-service data retrieval
	//
	// 	- menuId: menu ID
	//
	// 	- uri: ID or URL of the resource associated with the menu
	//
	// 	- showOnlyWithAccess: Authorized Only Visible
	//
	// 	- menuName: menu display name
	//
	// 	- dependentPermisson: whether the report resource associated with the menu has permissions
	//
	// 	- children: submenu
	//
	// example:
	//
	// [{"children":[{"children":[{"children":[{"menuId":"54kqgoa\\*\\*\\*\\*","menuName":"Report menu","menuType":0,"showOnlyWithAccess":true,"dependentPermisson":false,"uri":"e5da4a3f-d7f9-4262-a39e-a840043c\\*\\*\\*\\*"},{\\*\\*\\*\\	- "menu1nId":"pName" 135 "Directory menu","menuType":4,"showOnlyWithAccess":false,"dependentPermisson":true}],"menuId":"23a7d5d8-e55a-4737-b6a1-3c585505\\*\\*\\*\\*","menuName":"pop level -3 menu","menuType":4,"showOnlyWithAccess":true,"dependentPermisson":true}],"menuId":"80764 f3c-affd-45a1-aaa1-bb039d8a\\*\\*\\*\\*","menuName":"pop menu","menuType":4,"showOnlyWithAccess":false,"dependentPermisson":true}],"menuId":"277 f968a-22 ff-4ce6-83f0-a82950f4\\*\\*\\*\\*","menuName":"pop menu","menuType":4,"showOnlyWithAccess":false,"dependentPermisson":true}]
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPortalMenusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 10
	OffsetDay *int32 `json:"OffsetDay,omitempty" xml:"OffsetDay,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// example:
	//
	// PAGE
	TreeType *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListRecentViewReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 当前API返回的分页参数如TotalNum等即将下线，如有使用请及时修改
	Attention *string                                        `json:"Attention,omitempty" xml:"Attention,omitempty"`
	Data      []*ListRecentViewReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListRecentViewReportsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRecentViewReportsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponseBodyResult) SetAttention(v string) *ListRecentViewReportsResponseBodyResult {
	s.Attention = &v
	return s
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
	// example:
	//
	// true
	Favorite *bool `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	// example:
	//
	// 1496651577000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1640595729000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// true
	HasEditAuth *bool `json:"HasEditAuth,omitempty" xml:"HasEditAuth,omitempty"`
	// example:
	//
	// true
	HasViewAuth *bool `json:"HasViewAuth,omitempty" xml:"HasViewAuth,omitempty"`
	// example:
	//
	// 1642067498000
	LatestViewTime *string `json:"LatestViewTime,omitempty" xml:"LatestViewTime,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName      *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1365*****238860
	OwnerNum *string `json:"OwnerNum,omitempty" xml:"OwnerNum,omitempty"`
	// example:
	//
	// 1
	PublishStatus *int32 `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// example:
	//
	// 977c7698-****-****-****-44b7304d20fc
	TreeId *string `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
	// example:
	//
	// PAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 7
	ViewCount *int64 `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
	// example:
	//
	// 523793cb-****-****-****-aa71c65ffa39
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentViewReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// PAGE
	TreeType *string `json:"TreeType,omitempty" xml:"TreeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListSharedReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Data []*ListSharedReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// example:
	//
	// true
	Favorite *bool `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	// example:
	//
	// 1640088615000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1644373980000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// true
	HasEditAuth *bool `json:"HasEditAuth,omitempty" xml:"HasEditAuth,omitempty"`
	// example:
	//
	// true
	HasViewAuth *bool   `json:"HasViewAuth,omitempty" xml:"HasViewAuth,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName   *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1365*****238860
	OwnerNum *string `json:"OwnerNum,omitempty" xml:"OwnerNum,omitempty"`
	// example:
	//
	// 1
	PublishStatus *int32 `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// example:
	//
	// 977c7698-****-****-****-44b7304d20fc
	TreeId *string `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
	// example:
	//
	// PAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// gfidm145-****-****-9426-8f93be23****
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSharedReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
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
	// example:
	//
	// E2440604-3059-561A-AD68-DEDBC870EB2B
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListUserGroupsByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The user group modifier. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 2021-03-15 17:13:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	CreateUser     *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	IdentifiedPath *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	// example:
	//
	// 2021-03-15 20:36:40
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
	ParentUsergroupId *string `json:"ParentUsergroupId,omitempty" xml:"ParentUsergroupId,omitempty"`
	UsergroupDesc     *string `json:"UsergroupDesc,omitempty" xml:"UsergroupDesc,omitempty"`
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****
	UsergroupId   *string `json:"UsergroupId,omitempty" xml:"UsergroupId,omitempty"`
	UsergroupName *string `json:"UsergroupName,omitempty" xml:"UsergroupName,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListWorkspaceRoleUsersRequest struct {
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 726bee5a-****-43e1-9a8e-b550f0120f35
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspaceRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersRequest) SetKeyword(v string) *ListWorkspaceRoleUsersRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) SetPageNum(v int32) *ListWorkspaceRoleUsersRequest {
	s.PageNum = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) SetPageSize(v int32) *ListWorkspaceRoleUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) SetRoleId(v int64) *ListWorkspaceRoleUsersRequest {
	s.RoleId = &v
	return s
}

func (s *ListWorkspaceRoleUsersRequest) SetWorkspaceId(v string) *ListWorkspaceRoleUsersRequest {
	s.WorkspaceId = &v
	return s
}

type ListWorkspaceRoleUsersResponseBody struct {
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListWorkspaceRoleUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkspaceRoleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersResponseBody) SetRequestId(v string) *ListWorkspaceRoleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBody) SetResult(v *ListWorkspaceRoleUsersResponseBodyResult) *ListWorkspaceRoleUsersResponseBody {
	s.Result = v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBody) SetSuccess(v bool) *ListWorkspaceRoleUsersResponseBody {
	s.Success = &v
	return s
}

type ListWorkspaceRoleUsersResponseBodyResult struct {
	Data []*ListWorkspaceRoleUsersResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListWorkspaceRoleUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRoleUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetData(v []*ListWorkspaceRoleUsersResponseBodyResultData) *ListWorkspaceRoleUsersResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetPageNum(v int32) *ListWorkspaceRoleUsersResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetPageSize(v int32) *ListWorkspaceRoleUsersResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetTotalNum(v int32) *ListWorkspaceRoleUsersResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResult) SetTotalPages(v int32) *ListWorkspaceRoleUsersResponseBodyResult {
	s.TotalPages = &v
	return s
}

type ListWorkspaceRoleUsersResponseBodyResultData struct {
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 7350a155-0e94-4c6c-8620-57bbec38****
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListWorkspaceRoleUsersResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRoleUsersResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) SetNickName(v string) *ListWorkspaceRoleUsersResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) SetUserId(v string) *ListWorkspaceRoleUsersResponseBodyResultData {
	s.UserId = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) SetWorkspaceId(v string) *ListWorkspaceRoleUsersResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponseBodyResultData) SetWorkspaceName(v string) *ListWorkspaceRoleUsersResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type ListWorkspaceRoleUsersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRoleUsersResponse) SetHeaders(v map[string]*string) *ListWorkspaceRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceRoleUsersResponse) SetStatusCode(v int32) *ListWorkspaceRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceRoleUsersResponse) SetBody(v *ListWorkspaceRoleUsersResponseBody) *ListWorkspaceRoleUsersResponse {
	s.Body = v
	return s
}

type ListWorkspaceRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspaceRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRolesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesRequest) SetWorkspaceId(v string) *ListWorkspaceRolesRequest {
	s.WorkspaceId = &v
	return s
}

type ListWorkspaceRolesResponseBody struct {
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListWorkspaceRolesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkspaceRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBody) SetRequestId(v string) *ListWorkspaceRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceRolesResponseBody) SetResult(v []*ListWorkspaceRolesResponseBodyResult) *ListWorkspaceRolesResponseBody {
	s.Result = v
	return s
}

func (s *ListWorkspaceRolesResponseBody) SetSuccess(v bool) *ListWorkspaceRolesResponseBody {
	s.Success = &v
	return s
}

type ListWorkspaceRolesResponseBodyResult struct {
	AuthConfigList []*ListWorkspaceRolesResponseBodyResultAuthConfigList `json:"AuthConfigList,omitempty" xml:"AuthConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsSystemRole *bool `json:"IsSystemRole,omitempty" xml:"IsSystemRole,omitempty"`
	// example:
	//
	// 25
	RoleId   *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListWorkspaceRolesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRolesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBodyResult) SetAuthConfigList(v []*ListWorkspaceRolesResponseBodyResultAuthConfigList) *ListWorkspaceRolesResponseBodyResult {
	s.AuthConfigList = v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResult) SetIsSystemRole(v bool) *ListWorkspaceRolesResponseBodyResult {
	s.IsSystemRole = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResult) SetRoleId(v int64) *ListWorkspaceRolesResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResult) SetRoleName(v string) *ListWorkspaceRolesResponseBodyResult {
	s.RoleName = &v
	return s
}

type ListWorkspaceRolesResponseBodyResultAuthConfigList struct {
	ActionAuthKeys []*string `json:"ActionAuthKeys,omitempty" xml:"ActionAuthKeys,omitempty" type:"Repeated"`
	// example:
	//
	// portal_create
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
}

func (s ListWorkspaceRolesResponseBodyResultAuthConfigList) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRolesResponseBodyResultAuthConfigList) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBodyResultAuthConfigList) SetActionAuthKeys(v []*string) *ListWorkspaceRolesResponseBodyResultAuthConfigList {
	s.ActionAuthKeys = v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResultAuthConfigList) SetAuthKey(v string) *ListWorkspaceRolesResponseBodyResultAuthConfigList {
	s.AuthKey = &v
	return s
}

type ListWorkspaceRolesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceRolesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponse) SetHeaders(v map[string]*string) *ListWorkspaceRolesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceRolesResponse) SetStatusCode(v int32) *ListWorkspaceRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceRolesResponse) SetBody(v *ListWorkspaceRolesResponseBody) *ListWorkspaceRolesResponse {
	s.Body = v
	return s
}

type ModifyApiDatasourceParametersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b66a66de51f24d149116c17718138194
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"name":"token","value":"xxxxxxxxxxxx"},{"name":"pageSize","value":100}]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 726bee5a-****-43e1-9a8e-b550f0120f35
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ModifyApiDatasourceParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiDatasourceParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiDatasourceParametersRequest) SetApiId(v string) *ModifyApiDatasourceParametersRequest {
	s.ApiId = &v
	return s
}

func (s *ModifyApiDatasourceParametersRequest) SetParameters(v string) *ModifyApiDatasourceParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyApiDatasourceParametersRequest) SetWorkspaceId(v string) *ModifyApiDatasourceParametersRequest {
	s.WorkspaceId = &v
	return s
}

type ModifyApiDatasourceParametersResponseBody struct {
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyApiDatasourceParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiDatasourceParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiDatasourceParametersResponseBody) SetRequestId(v string) *ModifyApiDatasourceParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiDatasourceParametersResponseBody) SetResult(v bool) *ModifyApiDatasourceParametersResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyApiDatasourceParametersResponseBody) SetSuccess(v bool) *ModifyApiDatasourceParametersResponseBody {
	s.Success = &v
	return s
}

type ModifyApiDatasourceParametersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiDatasourceParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiDatasourceParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiDatasourceParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiDatasourceParametersResponse) SetHeaders(v map[string]*string) *ModifyApiDatasourceParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiDatasourceParametersResponse) SetStatusCode(v int32) *ModifyApiDatasourceParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiDatasourceParametersResponse) SetBody(v *ModifyApiDatasourceParametersResponseBody) *ModifyApiDatasourceParametersResponse {
	s.Body = v
	return s
}

type ModifyCopilotEmbedConfigRequest struct {
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccd3428c-dd2xxxxxxxxxxxxdffee
	CopilotId *string `json:"CopilotId,omitempty" xml:"CopilotId,omitempty"`
	// example:
	//
	// Map<String,Object> data=new HashMap<>();
	//
	//         data.put("allTheme",true);
	//
	//         //data.put("allCube",true);
	//
	//         //data.put("themes",Lists.newArrayList("1111","22222"));
	//
	//         //data.put("llmCubes",Lists.newArrayList("33333","44444"));
	//
	//         request.setDataRange(JSON.toJSONString(data));
	DataRange  *string `json:"DataRange,omitempty" xml:"DataRange,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s ModifyCopilotEmbedConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCopilotEmbedConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyCopilotEmbedConfigRequest) SetAgentName(v string) *ModifyCopilotEmbedConfigRequest {
	s.AgentName = &v
	return s
}

func (s *ModifyCopilotEmbedConfigRequest) SetCopilotId(v string) *ModifyCopilotEmbedConfigRequest {
	s.CopilotId = &v
	return s
}

func (s *ModifyCopilotEmbedConfigRequest) SetDataRange(v string) *ModifyCopilotEmbedConfigRequest {
	s.DataRange = &v
	return s
}

func (s *ModifyCopilotEmbedConfigRequest) SetModuleName(v string) *ModifyCopilotEmbedConfigRequest {
	s.ModuleName = &v
	return s
}

type ModifyCopilotEmbedConfigResponseBody struct {
	// example:
	//
	// 4BAA4694-CC27-555F-B15A-688AA9289FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCopilotEmbedConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCopilotEmbedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCopilotEmbedConfigResponseBody) SetRequestId(v string) *ModifyCopilotEmbedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCopilotEmbedConfigResponseBody) SetResult(v bool) *ModifyCopilotEmbedConfigResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyCopilotEmbedConfigResponseBody) SetSuccess(v bool) *ModifyCopilotEmbedConfigResponseBody {
	s.Success = &v
	return s
}

type ModifyCopilotEmbedConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCopilotEmbedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCopilotEmbedConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCopilotEmbedConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyCopilotEmbedConfigResponse) SetHeaders(v map[string]*string) *ModifyCopilotEmbedConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyCopilotEmbedConfigResponse) SetStatusCode(v int32) *ModifyCopilotEmbedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCopilotEmbedConfigResponse) SetBody(v *ModifyCopilotEmbedConfigResponseBody) *ModifyCopilotEmbedConfigResponse {
	s.Body = v
	return s
}

type QueryApprovalInfoRequest struct {
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12352fasdavsa
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryApprovalInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoRequest) SetPage(v int32) *QueryApprovalInfoRequest {
	s.Page = &v
	return s
}

func (s *QueryApprovalInfoRequest) SetPageSize(v int32) *QueryApprovalInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryApprovalInfoRequest) SetStatus(v int32) *QueryApprovalInfoRequest {
	s.Status = &v
	return s
}

func (s *QueryApprovalInfoRequest) SetUserId(v string) *QueryApprovalInfoRequest {
	s.UserId = &v
	return s
}

type QueryApprovalInfoResponseBody struct {
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryApprovalInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryApprovalInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoResponseBody) SetRequestId(v string) *QueryApprovalInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryApprovalInfoResponseBody) SetResult(v *QueryApprovalInfoResponseBodyResult) *QueryApprovalInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryApprovalInfoResponseBody) SetSuccess(v bool) *QueryApprovalInfoResponseBody {
	s.Success = &v
	return s
}

type QueryApprovalInfoResponseBodyResult struct {
	Data []*QueryApprovalInfoResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryApprovalInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoResponseBodyResult) SetData(v []*QueryApprovalInfoResponseBodyResultData) *QueryApprovalInfoResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetPage(v int32) *QueryApprovalInfoResponseBodyResult {
	s.Page = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetPageSize(v int32) *QueryApprovalInfoResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetStart(v int32) *QueryApprovalInfoResponseBodyResult {
	s.Start = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetTotal(v int32) *QueryApprovalInfoResponseBodyResult {
	s.Total = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResult) SetTotalPages(v int32) *QueryApprovalInfoResponseBodyResult {
	s.TotalPages = &v
	return s
}

type QueryApprovalInfoResponseBodyResultData struct {
	// example:
	//
	// 1359508
	ApplicantId   *string `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicantName *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	// example:
	//
	// 64813ef6da58e80eef8ed2f9
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ApplyReason   *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// example:
	//
	// sdasascasxasd
	ApproverId *string `json:"ApproverId,omitempty" xml:"ApproverId,omitempty"`
	// example:
	//
	// data_fusion_002
	ApproverName *string `json:"ApproverName,omitempty" xml:"ApproverName,omitempty"`
	// example:
	//
	// true
	DeleteFlag *bool `json:"DeleteFlag,omitempty" xml:"DeleteFlag,omitempty"`
	// example:
	//
	// 1708568097135
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// 0
	FlagStatus *int32 `json:"FlagStatus,omitempty" xml:"FlagStatus,omitempty"`
	// example:
	//
	// 1687315758
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1640595729000
	GmtModified  *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HandleReason *string `json:"HandleReason,omitempty" xml:"HandleReason,omitempty"`
	// example:
	//
	// acl-ct4t2e4u2x4ej1bzur
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// DASHBOARD
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryApprovalInfoResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalInfoResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApplicantId(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApplicantId = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApplicantName(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApplicantName = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApplicationId(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApplicationId = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApplyReason(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApplyReason = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApproverId(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApproverId = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetApproverName(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ApproverName = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetDeleteFlag(v bool) *QueryApprovalInfoResponseBodyResultData {
	s.DeleteFlag = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetExpireDate(v int64) *QueryApprovalInfoResponseBodyResultData {
	s.ExpireDate = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetFlagStatus(v int32) *QueryApprovalInfoResponseBodyResultData {
	s.FlagStatus = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetGmtCreate(v int64) *QueryApprovalInfoResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetGmtModified(v int64) *QueryApprovalInfoResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetHandleReason(v string) *QueryApprovalInfoResponseBodyResultData {
	s.HandleReason = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetResourceId(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ResourceId = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetResourceName(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ResourceName = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetResourceType(v string) *QueryApprovalInfoResponseBodyResultData {
	s.ResourceType = &v
	return s
}

func (s *QueryApprovalInfoResponseBodyResultData) SetWorkspaceName(v string) *QueryApprovalInfoResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type QueryApprovalInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApprovalInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApprovalInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryApprovalInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoResponse) SetHeaders(v map[string]*string) *QueryApprovalInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryApprovalInfoResponse) SetStatusCode(v int32) *QueryApprovalInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApprovalInfoResponse) SetBody(v *QueryApprovalInfoResponseBody) *QueryApprovalInfoResponse {
	s.Body = v
	return s
}

type QueryAuditLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1715856218001
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// function
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0***
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// example:
	//
	// MODIFY
	OperatorTypes *string `json:"OperatorTypes,omitempty" xml:"OperatorTypes,omitempty"`
	// example:
	//
	// cube
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1715856218001
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryAuditLogRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAuditLogRequest) GoString() string {
	return s.String()
}

func (s *QueryAuditLogRequest) SetEndDate(v string) *QueryAuditLogRequest {
	s.EndDate = &v
	return s
}

func (s *QueryAuditLogRequest) SetLogType(v string) *QueryAuditLogRequest {
	s.LogType = &v
	return s
}

func (s *QueryAuditLogRequest) SetOperatorId(v string) *QueryAuditLogRequest {
	s.OperatorId = &v
	return s
}

func (s *QueryAuditLogRequest) SetOperatorTypes(v string) *QueryAuditLogRequest {
	s.OperatorTypes = &v
	return s
}

func (s *QueryAuditLogRequest) SetResourceType(v string) *QueryAuditLogRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryAuditLogRequest) SetStartDate(v string) *QueryAuditLogRequest {
	s.StartDate = &v
	return s
}

func (s *QueryAuditLogRequest) SetWorkspaceId(v string) *QueryAuditLogRequest {
	s.WorkspaceId = &v
	return s
}

type QueryAuditLogResponseBody struct {
	// example:
	//
	// 78C1AA2D-9201-599E-A0BA-6FC462E57A95
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryAuditLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAuditLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuditLogResponseBody) SetRequestId(v string) *QueryAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuditLogResponseBody) SetResult(v []*QueryAuditLogResponseBodyResult) *QueryAuditLogResponseBody {
	s.Result = v
	return s
}

func (s *QueryAuditLogResponseBody) SetSuccess(v bool) *QueryAuditLogResponseBody {
	s.Success = &v
	return s
}

type QueryAuditLogResponseBodyResult struct {
	// example:
	//
	// 2024-04-16 13:17:39
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// wukaibis
	OperatorAccountName *string `json:"OperatorAccountName,omitempty" xml:"OperatorAccountName,omitempty"`
	// example:
	//
	// buc_344078
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// example:
	//
	// CREATE
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	TargetName   *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// example:
	//
	// USER
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// example:
	//
	// 87c6b145-090c-43e1-9426-8f93be23****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryAuditLogResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAuditLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAuditLogResponseBodyResult) SetGmtCreate(v string) *QueryAuditLogResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetOperatorAccountName(v string) *QueryAuditLogResponseBodyResult {
	s.OperatorAccountName = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetOperatorName(v string) *QueryAuditLogResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetOperatorType(v string) *QueryAuditLogResponseBodyResult {
	s.OperatorType = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetTargetName(v string) *QueryAuditLogResponseBodyResult {
	s.TargetName = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetTargetType(v string) *QueryAuditLogResponseBodyResult {
	s.TargetType = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetWorkspaceId(v string) *QueryAuditLogResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

type QueryAuditLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAuditLogResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAuditLogResponse) GoString() string {
	return s.String()
}

func (s *QueryAuditLogResponse) SetHeaders(v map[string]*string) *QueryAuditLogResponse {
	s.Headers = v
	return s
}

func (s *QueryAuditLogResponse) SetStatusCode(v int32) *QueryAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuditLogResponse) SetBody(v *QueryAuditLogResponseBody) *QueryAuditLogResponse {
	s.Body = v
	return s
}

type QueryComponentPerformanceRequest struct {
	// example:
	//
	// 1
	CostTimeAvgMin *int32 `json:"CostTimeAvgMin,omitempty" xml:"CostTimeAvgMin,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sevenDays
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// example:
	//
	// report
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryComponentPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentPerformanceRequest) GoString() string {
	return s.String()
}

func (s *QueryComponentPerformanceRequest) SetCostTimeAvgMin(v int32) *QueryComponentPerformanceRequest {
	s.CostTimeAvgMin = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetPageNum(v int32) *QueryComponentPerformanceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetPageSize(v int32) *QueryComponentPerformanceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetQueryType(v string) *QueryComponentPerformanceRequest {
	s.QueryType = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetReportId(v string) *QueryComponentPerformanceRequest {
	s.ReportId = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetResourceType(v string) *QueryComponentPerformanceRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetWorkspaceId(v string) *QueryComponentPerformanceRequest {
	s.WorkspaceId = &v
	return s
}

type QueryComponentPerformanceResponseBody struct {
	// example:
	//
	// BCE45E6D-9304-4F94-86BB-5A772B1615FF
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryComponentPerformanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryComponentPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryComponentPerformanceResponseBody) SetRequestId(v string) *QueryComponentPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryComponentPerformanceResponseBody) SetResult(v []*QueryComponentPerformanceResponseBodyResult) *QueryComponentPerformanceResponseBody {
	s.Result = v
	return s
}

func (s *QueryComponentPerformanceResponseBody) SetSuccess(v bool) *QueryComponentPerformanceResponseBody {
	s.Success = &v
	return s
}

type QueryComponentPerformanceResponseBodyResult struct {
	// example:
	//
	// 0.3
	CacheCostTimeAvg *float64 `json:"CacheCostTimeAvg,omitempty" xml:"CacheCostTimeAvg,omitempty"`
	// example:
	//
	// 3
	CacheQueryCount *int32 `json:"CacheQueryCount,omitempty" xml:"CacheQueryCount,omitempty"`
	// example:
	//
	// 0696083a-ca72-4d89-8e7a-c017910e0***
	ComponentId   *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// example:
	//
	// 0.3
	CostTimeAvg *float64 `json:"CostTimeAvg,omitempty" xml:"CostTimeAvg,omitempty"`
	// example:
	//
	// 5
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// example:
	//
	// 0.3
	QueryCountAvg *float64 `json:"QueryCountAvg,omitempty" xml:"QueryCountAvg,omitempty"`
	// example:
	//
	// 5
	QueryOverFivePercentNum *float64 `json:"QueryOverFivePercentNum,omitempty" xml:"QueryOverFivePercentNum,omitempty"`
	// example:
	//
	// 0.3
	QueryOverFiveSecPercent *string `json:"QueryOverFiveSecPercent,omitempty" xml:"QueryOverFiveSecPercent,omitempty"`
	// example:
	//
	// 0.3
	QueryOverTenSecPercent *string `json:"QueryOverTenSecPercent,omitempty" xml:"QueryOverTenSecPercent,omitempty"`
	// example:
	//
	// 0.3
	QueryOverTenSecPercentNum *float64 `json:"QueryOverTenSecPercentNum,omitempty" xml:"QueryOverTenSecPercentNum,omitempty"`
	// example:
	//
	// 1
	QueryTimeoutCount *int32 `json:"QueryTimeoutCount,omitempty" xml:"QueryTimeoutCount,omitempty"`
	// example:
	//
	// 0.3
	QueryTimeoutCountPercent *float64 `json:"QueryTimeoutCountPercent,omitempty" xml:"QueryTimeoutCountPercent,omitempty"`
	// example:
	//
	// 0.3
	QuickIndexCostTimeAvg *float64 `json:"QuickIndexCostTimeAvg,omitempty" xml:"QuickIndexCostTimeAvg,omitempty"`
	// example:
	//
	// 3
	QuickIndexQueryCount *int32 `json:"QuickIndexQueryCount,omitempty" xml:"QuickIndexQueryCount,omitempty"`
	// example:
	//
	// 0.3
	RepeatQueryPercent *string `json:"RepeatQueryPercent,omitempty" xml:"RepeatQueryPercent,omitempty"`
	// example:
	//
	// 2
	RepeatQueryPercentNum *float64 `json:"RepeatQueryPercentNum,omitempty" xml:"RepeatQueryPercentNum,omitempty"`
	// example:
	//
	// 5
	RepeatSqlQueryCount *int32 `json:"RepeatSqlQueryCount,omitempty" xml:"RepeatSqlQueryCount,omitempty"`
	// example:
	//
	// 0.3
	RepeatSqlQueryPercent *string `json:"RepeatSqlQueryPercent,omitempty" xml:"RepeatSqlQueryPercent,omitempty"`
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// example:
	//
	// ClusterRiskReport
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// example:
	//
	// report
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryComponentPerformanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentPerformanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryComponentPerformanceResponseBodyResult) SetCacheCostTimeAvg(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.CacheCostTimeAvg = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetCacheQueryCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.CacheQueryCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetComponentId(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ComponentId = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetComponentName(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ComponentName = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetCostTimeAvg(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.CostTimeAvg = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.QueryCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryCountAvg(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QueryCountAvg = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryOverFivePercentNum(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QueryOverFivePercentNum = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryOverFiveSecPercent(v string) *QueryComponentPerformanceResponseBodyResult {
	s.QueryOverFiveSecPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryOverTenSecPercent(v string) *QueryComponentPerformanceResponseBodyResult {
	s.QueryOverTenSecPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryOverTenSecPercentNum(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QueryOverTenSecPercentNum = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryTimeoutCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.QueryTimeoutCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQueryTimeoutCountPercent(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QueryTimeoutCountPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQuickIndexCostTimeAvg(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.QuickIndexCostTimeAvg = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetQuickIndexQueryCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.QuickIndexQueryCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetRepeatQueryPercent(v string) *QueryComponentPerformanceResponseBodyResult {
	s.RepeatQueryPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetRepeatQueryPercentNum(v float64) *QueryComponentPerformanceResponseBodyResult {
	s.RepeatQueryPercentNum = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetRepeatSqlQueryCount(v int32) *QueryComponentPerformanceResponseBodyResult {
	s.RepeatSqlQueryCount = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetRepeatSqlQueryPercent(v string) *QueryComponentPerformanceResponseBodyResult {
	s.RepeatSqlQueryPercent = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetReportId(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetReportName(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ReportName = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetReportType(v string) *QueryComponentPerformanceResponseBodyResult {
	s.ReportType = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetWorkspaceId(v string) *QueryComponentPerformanceResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryComponentPerformanceResponseBodyResult) SetWorkspaceName(v string) *QueryComponentPerformanceResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

type QueryComponentPerformanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryComponentPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryComponentPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentPerformanceResponse) GoString() string {
	return s.String()
}

func (s *QueryComponentPerformanceResponse) SetHeaders(v map[string]*string) *QueryComponentPerformanceResponse {
	s.Headers = v
	return s
}

func (s *QueryComponentPerformanceResponse) SetStatusCode(v int32) *QueryComponentPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryComponentPerformanceResponse) SetBody(v *QueryComponentPerformanceResponseBody) *QueryComponentPerformanceResponse {
	s.Body = v
	return s
}

type QueryCopilotEmbedConfigRequest struct {
	// example:
	//
	// 06-ELive
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s QueryCopilotEmbedConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCopilotEmbedConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigRequest) SetKeyword(v string) *QueryCopilotEmbedConfigRequest {
	s.Keyword = &v
	return s
}

type QueryCopilotEmbedConfigResponseBody struct {
	// example:
	//
	// 1FC71085-D5FD-08E0-813A-4D4BD1031BC5
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryCopilotEmbedConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCopilotEmbedConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCopilotEmbedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigResponseBody) SetRequestId(v string) *QueryCopilotEmbedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBody) SetResult(v []*QueryCopilotEmbedConfigResponseBodyResult) *QueryCopilotEmbedConfigResponseBody {
	s.Result = v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBody) SetSuccess(v bool) *QueryCopilotEmbedConfigResponseBody {
	s.Success = &v
	return s
}

type QueryCopilotEmbedConfigResponseBodyResult struct {
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// 9c079710-ddbe-48b3-b495-7c83c8d57cc4
	CopilotId *string `json:"CopilotId,omitempty" xml:"CopilotId,omitempty"`
	// example:
	//
	// qweqw12312423521
	CreateUser     *string                                             `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	CreateUserName *string                                             `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DataRange      *QueryCopilotEmbedConfigResponseBodyResultDataRange `json:"DataRange,omitempty" xml:"DataRange,omitempty" type:"Struct"`
	// example:
	//
	// asda1231231dfs
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// 0327
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s QueryCopilotEmbedConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCopilotEmbedConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetAgentName(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.AgentName = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetCopilotId(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.CopilotId = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetCreateUser(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.CreateUser = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetCreateUserName(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.CreateUserName = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetDataRange(v *QueryCopilotEmbedConfigResponseBodyResultDataRange) *QueryCopilotEmbedConfigResponseBodyResult {
	s.DataRange = v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetModifyUser(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.ModifyUser = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetModuleName(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.ModuleName = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResult) SetShowName(v string) *QueryCopilotEmbedConfigResponseBodyResult {
	s.ShowName = &v
	return s
}

type QueryCopilotEmbedConfigResponseBodyResultDataRange struct {
	// example:
	//
	// true/false
	AllCube *bool `json:"AllCube,omitempty" xml:"AllCube,omitempty"`
	// example:
	//
	// true/false
	AllTheme *bool     `json:"AllTheme,omitempty" xml:"AllTheme,omitempty"`
	LlmCubes []*string `json:"LlmCubes,omitempty" xml:"LlmCubes,omitempty" type:"Repeated"`
	Themes   []*string `json:"Themes,omitempty" xml:"Themes,omitempty" type:"Repeated"`
}

func (s QueryCopilotEmbedConfigResponseBodyResultDataRange) String() string {
	return tea.Prettify(s)
}

func (s QueryCopilotEmbedConfigResponseBodyResultDataRange) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) SetAllCube(v bool) *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	s.AllCube = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) SetAllTheme(v bool) *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	s.AllTheme = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) SetLlmCubes(v []*string) *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	s.LlmCubes = v
	return s
}

func (s *QueryCopilotEmbedConfigResponseBodyResultDataRange) SetThemes(v []*string) *QueryCopilotEmbedConfigResponseBodyResultDataRange {
	s.Themes = v
	return s
}

type QueryCopilotEmbedConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCopilotEmbedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCopilotEmbedConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCopilotEmbedConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigResponse) SetHeaders(v map[string]*string) *QueryCopilotEmbedConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryCopilotEmbedConfigResponse) SetStatusCode(v int32) *QueryCopilotEmbedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponse) SetBody(v *QueryCopilotEmbedConfigResponseBody) *QueryCopilotEmbedConfigResponse {
	s.Body = v
	return s
}

type QueryCubeOptimizationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryCubeOptimizationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCubeOptimizationRequest) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationRequest) SetWorkspaceId(v string) *QueryCubeOptimizationRequest {
	s.WorkspaceId = &v
	return s
}

type QueryCubeOptimizationResponseBody struct {
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryCubeOptimizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCubeOptimizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCubeOptimizationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationResponseBody) SetRequestId(v string) *QueryCubeOptimizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCubeOptimizationResponseBody) SetResult(v []*QueryCubeOptimizationResponseBodyResult) *QueryCubeOptimizationResponseBody {
	s.Result = v
	return s
}

func (s *QueryCubeOptimizationResponseBody) SetSuccess(v bool) *QueryCubeOptimizationResponseBody {
	s.Success = &v
	return s
}

type QueryCubeOptimizationResponseBodyResult struct {
	// example:
	//
	// OPENQUICKENGINE
	AdviceType                   *string                                                              `json:"AdviceType,omitempty" xml:"AdviceType,omitempty"`
	CubePerformanceDiagnoseModel *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel `json:"CubePerformanceDiagnoseModel,omitempty" xml:"CubePerformanceDiagnoseModel,omitempty" type:"Struct"`
}

func (s QueryCubeOptimizationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCubeOptimizationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationResponseBodyResult) SetAdviceType(v string) *QueryCubeOptimizationResponseBodyResult {
	s.AdviceType = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResult) SetCubePerformanceDiagnoseModel(v *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) *QueryCubeOptimizationResponseBodyResult {
	s.CubePerformanceDiagnoseModel = v
	return s
}

type QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel struct {
	// example:
	//
	// 1
	CacheCostTimeAvg *float64 `json:"CacheCostTimeAvg,omitempty" xml:"CacheCostTimeAvg,omitempty"`
	// example:
	//
	// 2
	CacheQueryCount *int32 `json:"CacheQueryCount,omitempty" xml:"CacheQueryCount,omitempty"`
	// example:
	//
	// 1.0
	CostTimeAvg *float64 `json:"CostTimeAvg,omitempty" xml:"CostTimeAvg,omitempty"`
	// example:
	//
	// 3e45b61a-9ba8-4c7c-8248-8dbe69945636
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// example:
	//
	// test
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// example:
	//
	// 50
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// example:
	//
	// 2
	QueryCountAvg *float64 `json:"QueryCountAvg,omitempty" xml:"QueryCountAvg,omitempty"`
	// example:
	//
	// 0.1
	QueryOverFivePercentNum *float64 `json:"QueryOverFivePercentNum,omitempty" xml:"QueryOverFivePercentNum,omitempty"`
	// example:
	//
	// 0.5
	QueryOverFiveSecPercent *string `json:"QueryOverFiveSecPercent,omitempty" xml:"QueryOverFiveSecPercent,omitempty"`
	// example:
	//
	// 0.1
	QueryOverTenSecPercent *string `json:"QueryOverTenSecPercent,omitempty" xml:"QueryOverTenSecPercent,omitempty"`
	// example:
	//
	// 0.3
	QueryOverTenSecPercentNum *float64 `json:"QueryOverTenSecPercentNum,omitempty" xml:"QueryOverTenSecPercentNum,omitempty"`
	// example:
	//
	// 1
	QueryTimeoutCount *int32 `json:"QueryTimeoutCount,omitempty" xml:"QueryTimeoutCount,omitempty"`
	// example:
	//
	// 0.3
	QueryTimeoutCountPercent *float64 `json:"QueryTimeoutCountPercent,omitempty" xml:"QueryTimeoutCountPercent,omitempty"`
	// example:
	//
	// 1
	QuickIndexCostTimeAvg *float64 `json:"QuickIndexCostTimeAvg,omitempty" xml:"QuickIndexCostTimeAvg,omitempty"`
	// example:
	//
	// 2
	QuickIndexQueryCount *int32 `json:"QuickIndexQueryCount,omitempty" xml:"QuickIndexQueryCount,omitempty"`
	// example:
	//
	// 0.1
	RepeatQueryPercent *string `json:"RepeatQueryPercent,omitempty" xml:"RepeatQueryPercent,omitempty"`
	// example:
	//
	// 2
	RepeatQueryPercentNum *float64 `json:"RepeatQueryPercentNum,omitempty" xml:"RepeatQueryPercentNum,omitempty"`
	// example:
	//
	// 2
	RepeatSqlQueryCount *int32 `json:"RepeatSqlQueryCount,omitempty" xml:"RepeatSqlQueryCount,omitempty"`
	// example:
	//
	// 0.3
	RepeatSqlQueryPercent *string `json:"RepeatSqlQueryPercent,omitempty" xml:"RepeatSqlQueryPercent,omitempty"`
	// example:
	//
	// 6ea74bff-c818-4188-b462-dbb45a24dbac
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// eco0sh0prods
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) String() string {
	return tea.Prettify(s)
}

func (s QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCacheCostTimeAvg(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CacheCostTimeAvg = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCacheQueryCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CacheQueryCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCostTimeAvg(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CostTimeAvg = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCubeId(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CubeId = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetCubeName(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.CubeName = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryCountAvg(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryCountAvg = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryOverFivePercentNum(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryOverFivePercentNum = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryOverFiveSecPercent(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryOverFiveSecPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryOverTenSecPercent(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryOverTenSecPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryOverTenSecPercentNum(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryOverTenSecPercentNum = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryTimeoutCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryTimeoutCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQueryTimeoutCountPercent(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QueryTimeoutCountPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQuickIndexCostTimeAvg(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QuickIndexCostTimeAvg = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetQuickIndexQueryCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.QuickIndexQueryCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetRepeatQueryPercent(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.RepeatQueryPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetRepeatQueryPercentNum(v float64) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.RepeatQueryPercentNum = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetRepeatSqlQueryCount(v int32) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.RepeatSqlQueryCount = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetRepeatSqlQueryPercent(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.RepeatSqlQueryPercent = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetWorkspaceId(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel) SetWorkspaceName(v string) *QueryCubeOptimizationResponseBodyResultCubePerformanceDiagnoseModel {
	s.WorkspaceName = &v
	return s
}

type QueryCubeOptimizationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCubeOptimizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCubeOptimizationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCubeOptimizationResponse) GoString() string {
	return s.String()
}

func (s *QueryCubeOptimizationResponse) SetHeaders(v map[string]*string) *QueryCubeOptimizationResponse {
	s.Headers = v
	return s
}

func (s *QueryCubeOptimizationResponse) SetStatusCode(v int32) *QueryCubeOptimizationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCubeOptimizationResponse) SetBody(v *QueryCubeOptimizationResponseBody) *QueryCubeOptimizationResponse {
	s.Body = v
	return s
}

type QueryCubePerformanceRequest struct {
	// example:
	//
	// 1
	CostTimeAvgMin *int32 `json:"CostTimeAvgMin,omitempty" xml:"CostTimeAvgMin,omitempty"`
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sevenDays
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryCubePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCubePerformanceRequest) GoString() string {
	return s.String()
}

func (s *QueryCubePerformanceRequest) SetCostTimeAvgMin(v int32) *QueryCubePerformanceRequest {
	s.CostTimeAvgMin = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetCubeId(v string) *QueryCubePerformanceRequest {
	s.CubeId = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetPageNum(v int32) *QueryCubePerformanceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetPageSize(v int32) *QueryCubePerformanceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetQueryType(v string) *QueryCubePerformanceRequest {
	s.QueryType = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetWorkspaceId(v string) *QueryCubePerformanceRequest {
	s.WorkspaceId = &v
	return s
}

type QueryCubePerformanceResponseBody struct {
	// example:
	//
	// 685072a0-1fd5-40ef-ae6b-cf94e79e718f
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryCubePerformanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCubePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCubePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCubePerformanceResponseBody) SetRequestId(v string) *QueryCubePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCubePerformanceResponseBody) SetResult(v []*QueryCubePerformanceResponseBodyResult) *QueryCubePerformanceResponseBody {
	s.Result = v
	return s
}

func (s *QueryCubePerformanceResponseBody) SetSuccess(v bool) *QueryCubePerformanceResponseBody {
	s.Success = &v
	return s
}

type QueryCubePerformanceResponseBodyResult struct {
	// example:
	//
	// 1
	CacheCostTimeAvg *float64 `json:"CacheCostTimeAvg,omitempty" xml:"CacheCostTimeAvg,omitempty"`
	// example:
	//
	// 1
	CacheQueryCount *int32 `json:"CacheQueryCount,omitempty" xml:"CacheQueryCount,omitempty"`
	// example:
	//
	// 1
	CostTimeAvg *float64 `json:"CostTimeAvg,omitempty" xml:"CostTimeAvg,omitempty"`
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// example:
	//
	// test
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// example:
	//
	// 50
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// example:
	//
	// 1
	QueryCountAvg *float64 `json:"QueryCountAvg,omitempty" xml:"QueryCountAvg,omitempty"`
	// example:
	//
	// 1.0
	QueryOverFivePercentNum *float64 `json:"QueryOverFivePercentNum,omitempty" xml:"QueryOverFivePercentNum,omitempty"`
	// example:
	//
	// 1.0
	QueryOverFiveSecPercent *string `json:"QueryOverFiveSecPercent,omitempty" xml:"QueryOverFiveSecPercent,omitempty"`
	// example:
	//
	// 1.0
	QueryOverTenSecPercent *string `json:"QueryOverTenSecPercent,omitempty" xml:"QueryOverTenSecPercent,omitempty"`
	// example:
	//
	// 1.0
	QueryOverTenSecPercentNum *float64 `json:"QueryOverTenSecPercentNum,omitempty" xml:"QueryOverTenSecPercentNum,omitempty"`
	// example:
	//
	// 1
	QueryTimeoutCount *int32 `json:"QueryTimeoutCount,omitempty" xml:"QueryTimeoutCount,omitempty"`
	// example:
	//
	// 1
	QueryTimeoutCountPercent *float64 `json:"QueryTimeoutCountPercent,omitempty" xml:"QueryTimeoutCountPercent,omitempty"`
	// example:
	//
	// 1
	QuickIndexCostTimeAvg *float64 `json:"QuickIndexCostTimeAvg,omitempty" xml:"QuickIndexCostTimeAvg,omitempty"`
	// example:
	//
	// 1
	QuickIndexQueryCount *int32 `json:"QuickIndexQueryCount,omitempty" xml:"QuickIndexQueryCount,omitempty"`
	// example:
	//
	// 0.3
	RepeatQueryPercent *string `json:"RepeatQueryPercent,omitempty" xml:"RepeatQueryPercent,omitempty"`
	// example:
	//
	// 1
	RepeatQueryPercentNum *float64 `json:"RepeatQueryPercentNum,omitempty" xml:"RepeatQueryPercentNum,omitempty"`
	// example:
	//
	// 1
	RepeatSqlQueryCount *int32 `json:"RepeatSqlQueryCount,omitempty" xml:"RepeatSqlQueryCount,omitempty"`
	// example:
	//
	// 1
	RepeatSqlQueryPercent *string `json:"RepeatSqlQueryPercent,omitempty" xml:"RepeatSqlQueryPercent,omitempty"`
	// example:
	//
	// 87c6b145-090c-43e1-9426-8f93be23****
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryCubePerformanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCubePerformanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCubePerformanceResponseBodyResult) SetCacheCostTimeAvg(v float64) *QueryCubePerformanceResponseBodyResult {
	s.CacheCostTimeAvg = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetCacheQueryCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.CacheQueryCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetCostTimeAvg(v float64) *QueryCubePerformanceResponseBodyResult {
	s.CostTimeAvg = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetCubeId(v string) *QueryCubePerformanceResponseBodyResult {
	s.CubeId = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetCubeName(v string) *QueryCubePerformanceResponseBodyResult {
	s.CubeName = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.QueryCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryCountAvg(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QueryCountAvg = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryOverFivePercentNum(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QueryOverFivePercentNum = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryOverFiveSecPercent(v string) *QueryCubePerformanceResponseBodyResult {
	s.QueryOverFiveSecPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryOverTenSecPercent(v string) *QueryCubePerformanceResponseBodyResult {
	s.QueryOverTenSecPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryOverTenSecPercentNum(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QueryOverTenSecPercentNum = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryTimeoutCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.QueryTimeoutCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQueryTimeoutCountPercent(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QueryTimeoutCountPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQuickIndexCostTimeAvg(v float64) *QueryCubePerformanceResponseBodyResult {
	s.QuickIndexCostTimeAvg = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetQuickIndexQueryCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.QuickIndexQueryCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetRepeatQueryPercent(v string) *QueryCubePerformanceResponseBodyResult {
	s.RepeatQueryPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetRepeatQueryPercentNum(v float64) *QueryCubePerformanceResponseBodyResult {
	s.RepeatQueryPercentNum = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetRepeatSqlQueryCount(v int32) *QueryCubePerformanceResponseBodyResult {
	s.RepeatSqlQueryCount = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetRepeatSqlQueryPercent(v string) *QueryCubePerformanceResponseBodyResult {
	s.RepeatSqlQueryPercent = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetWorkspaceId(v string) *QueryCubePerformanceResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCubePerformanceResponseBodyResult) SetWorkspaceName(v string) *QueryCubePerformanceResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

type QueryCubePerformanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCubePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCubePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCubePerformanceResponse) GoString() string {
	return s.String()
}

func (s *QueryCubePerformanceResponse) SetHeaders(v map[string]*string) *QueryCubePerformanceResponse {
	s.Headers = v
	return s
}

func (s *QueryCubePerformanceResponse) SetStatusCode(v int32) *QueryCubePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCubePerformanceResponse) SetBody(v *QueryCubePerformanceResponseBody) *QueryCubePerformanceResponse {
	s.Body = v
	return s
}

type QueryDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f4cc43bc3***
	ApiId      *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// example:
	//
	// ["area", "city", "price", "date"]
	ReturnFields *string `json:"ReturnFields,omitempty" xml:"ReturnFields,omitempty"`
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDataRequest) SetApiId(v string) *QueryDataRequest {
	s.ApiId = &v
	return s
}

func (s *QueryDataRequest) SetConditions(v string) *QueryDataRequest {
	s.Conditions = &v
	return s
}

func (s *QueryDataRequest) SetReturnFields(v string) *QueryDataRequest {
	s.ReturnFields = &v
	return s
}

func (s *QueryDataRequest) SetUserId(v string) *QueryDataRequest {
	s.UserId = &v
	return s
}

type QueryDataResponseBody struct {
	// example:
	//
	// a4d1a221d-41za1-****
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDataResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataResponseBody) SetRequestId(v string) *QueryDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataResponseBody) SetResult(v *QueryDataResponseBodyResult) *QueryDataResponseBody {
	s.Result = v
	return s
}

func (s *QueryDataResponseBody) SetSuccess(v bool) *QueryDataResponseBody {
	s.Success = &v
	return s
}

type QueryDataResponseBodyResult struct {
	Headers []*QueryDataResponseBodyResultHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	Sql     *string                               `json:"Sql,omitempty" xml:"Sql,omitempty"`
	Values  []map[string]interface{}              `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDataResponseBodyResult) SetHeaders(v []*QueryDataResponseBodyResultHeaders) *QueryDataResponseBodyResult {
	s.Headers = v
	return s
}

func (s *QueryDataResponseBodyResult) SetSql(v string) *QueryDataResponseBodyResult {
	s.Sql = &v
	return s
}

func (s *QueryDataResponseBodyResult) SetValues(v []map[string]interface{}) *QueryDataResponseBodyResult {
	s.Values = v
	return s
}

type QueryDataResponseBodyResultHeaders struct {
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	Column     *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// REGION
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// area
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// Dimension
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataResponseBodyResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDataResponseBodyResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryDataResponseBodyResultHeaders) SetAggregator(v string) *QueryDataResponseBodyResultHeaders {
	s.Aggregator = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetColumn(v string) *QueryDataResponseBodyResultHeaders {
	s.Column = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetDataType(v string) *QueryDataResponseBodyResultHeaders {
	s.DataType = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetGranularity(v string) *QueryDataResponseBodyResultHeaders {
	s.Granularity = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetLabel(v string) *QueryDataResponseBodyResultHeaders {
	s.Label = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetType(v string) *QueryDataResponseBodyResultHeaders {
	s.Type = &v
	return s
}

type QueryDataResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDataResponse) SetHeaders(v map[string]*string) *QueryDataResponse {
	s.Headers = v
	return s
}

func (s *QueryDataResponse) SetStatusCode(v int32) *QueryDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataResponse) SetBody(v *QueryDataResponseBody) *QueryDataResponse {
	s.Body = v
	return s
}

type QueryDataRangeRequest struct {
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llmCube
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDataRangeRequest) GoString() string {
	return s.String()
}

func (s *QueryDataRangeRequest) SetKeyword(v string) *QueryDataRangeRequest {
	s.Keyword = &v
	return s
}

func (s *QueryDataRangeRequest) SetType(v string) *QueryDataRangeRequest {
	s.Type = &v
	return s
}

type QueryDataRangeResponseBody struct {
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDataRangeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDataRangeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBody) SetRequestId(v string) *QueryDataRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataRangeResponseBody) SetResult(v *QueryDataRangeResponseBodyResult) *QueryDataRangeResponseBody {
	s.Result = v
	return s
}

func (s *QueryDataRangeResponseBody) SetSuccess(v bool) *QueryDataRangeResponseBody {
	s.Success = &v
	return s
}

type QueryDataRangeResponseBodyResult struct {
	ApiCopilotLlmCubeModels []*QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels `json:"ApiCopilotLlmCubeModels,omitempty" xml:"ApiCopilotLlmCubeModels,omitempty" type:"Repeated"`
	ApiCopilotThemeModels   []*QueryDataRangeResponseBodyResultApiCopilotThemeModels   `json:"ApiCopilotThemeModels,omitempty" xml:"ApiCopilotThemeModels,omitempty" type:"Repeated"`
}

func (s QueryDataRangeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDataRangeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBodyResult) SetApiCopilotLlmCubeModels(v []*QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) *QueryDataRangeResponseBodyResult {
	s.ApiCopilotLlmCubeModels = v
	return s
}

func (s *QueryDataRangeResponseBodyResult) SetApiCopilotThemeModels(v []*QueryDataRangeResponseBodyResultApiCopilotThemeModels) *QueryDataRangeResponseBodyResult {
	s.ApiCopilotThemeModels = v
	return s
}

type QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels struct {
	Alias      *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// sdasdafas23342342342
	LlmCubeId *string `json:"LlmCubeId,omitempty" xml:"LlmCubeId,omitempty"`
}

func (s QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) String() string {
	return tea.Prettify(s)
}

func (s QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) SetAlias(v string) *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels {
	s.Alias = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) SetCreateUser(v string) *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels {
	s.CreateUser = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels) SetLlmCubeId(v string) *QueryDataRangeResponseBodyResultApiCopilotLlmCubeModels {
	s.LlmCubeId = &v
	return s
}

type QueryDataRangeResponseBodyResultApiCopilotThemeModels struct {
	ApiCopilotLlmCubeModels []*QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels `json:"ApiCopilotLlmCubeModels,omitempty" xml:"ApiCopilotLlmCubeModels,omitempty" type:"Repeated"`
	CreateUser              *string                                                                         `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 36631232342312312
	ThemeId   *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	ThemeName *string `json:"ThemeName,omitempty" xml:"ThemeName,omitempty"`
}

func (s QueryDataRangeResponseBodyResultApiCopilotThemeModels) String() string {
	return tea.Prettify(s)
}

func (s QueryDataRangeResponseBodyResultApiCopilotThemeModels) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) SetApiCopilotLlmCubeModels(v []*QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) *QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	s.ApiCopilotLlmCubeModels = v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) SetCreateUser(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	s.CreateUser = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) SetThemeId(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	s.ThemeId = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModels) SetThemeName(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModels {
	s.ThemeName = &v
	return s
}

type QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels struct {
	Alias      *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 1231242231asdasda
	LlmCubeId *string `json:"LlmCubeId,omitempty" xml:"LlmCubeId,omitempty"`
}

func (s QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) String() string {
	return tea.Prettify(s)
}

func (s QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) SetAlias(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels {
	s.Alias = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) SetCreateUser(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels {
	s.CreateUser = &v
	return s
}

func (s *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels) SetLlmCubeId(v string) *QueryDataRangeResponseBodyResultApiCopilotThemeModelsApiCopilotLlmCubeModels {
	s.LlmCubeId = &v
	return s
}

type QueryDataRangeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDataRangeResponse) GoString() string {
	return s.String()
}

func (s *QueryDataRangeResponse) SetHeaders(v map[string]*string) *QueryDataRangeResponse {
	s.Headers = v
	return s
}

func (s *QueryDataRangeResponse) SetStatusCode(v int32) *QueryDataRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataRangeResponse) SetBody(v *QueryDataRangeResponseBody) *QueryDataRangeResponse {
	s.Body = v
	return s
}

type QueryDataServiceRequest struct {
	// Call an API that is created in DataService Studio.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4cc43bc3***
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// # Prerequisites
	//
	// You can use the Quick BI data service to create an API for the data service. For more information, see [Data service](https://help.aliyun.com/document_detail/144980.html).
	//
	// # Limits
	//
	// 	- The Data Service feature is available only to Professional customers.
	//
	// 	- The timeout period for API calls is 60s. The QPS of a single API is 10 times per second.
	//
	// 	- If row-level permissions are enabled for datasets that are referenced by a Data Service API, the API may be blocked by row-level permission policies.
	//
	// example:
	//
	// For more information about API IDs in DataService, see [Data Service](https://help.aliyun.com/document_detail/144980.html).
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The query conditions of the data service. The query conditions are specified in the form of keys and values. A string of the map type. Key is the name of the request parameters parameter, and Value is the value of the request parameters parameter. Key and Value must appear in pairs.
	//
	// **Note:**
	//
	// 	- If a value contains multiple values, the value is a List in the JSON format. Example: `area=["East China","North China","South China"]`
	//
	// 	- For dates, different input parameter formats are provided based on different types:
	//
	//     	- Year: 2019
	//
	//     	- Season: 2019Q1
	//
	//     	- Month: 201901 (carry 0)
	//
	//     	- Week: 2019-52
	//
	//     	- Day: 20190101
	//
	//     	- Hours: 14:00:00 (minutes and seconds are 00)
	//
	//     	- Minutes: 14:12:00 (seconds are 00)
	//
	//     	- Seconds: 14:34:34
	//
	// example:
	//
	// ["area", "city", "price", "date"]
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
	// The list of parameter names of the returned parameters. The value is a string of the List type.
	//
	// example:
	//
	// 78C1AA2D-9201-599E-A0BA-6FC462E57A95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	Result *QueryDataServiceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// { "area": ["East China", "North China"], "shopping_date": "2019Q1", }
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The SQL of the request query.
	Headers []*QueryDataServiceResponseBodyResultHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// The query result of the API operation is returned.
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// Physical Field Name
	Values []map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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
	// The field name, which corresponds to the physical table field name.
	//
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// The granularity of the dimension field. This field is returned only when the requested field is a date dimension or a geographical dimension. Valid values:
	//
	// 	- Date granularity: yearRegion (year), monthRegion (month), weekRegion (week), dayRegion (day), hourRegion (hour), minRegion (minute), secRegion (second)
	//
	// 	- Geographic information granularity: COUNTRY (international level), PROVINCE (provincial level), CITY (municipal level), XIAN (district /county), and REGION (regional level)
	//
	// example:
	//
	// The alias of the field. The key of the map data row in the result parameter values.
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The column header.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The field type, which is used to distinguish whether the field type is a dimension or a measure.
	//
	// example:
	//
	// yearRegion
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The data type of the field. generally have number, string, date, datetime, time, and geographic.
	//
	// example:
	//
	// area
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// SELECT COMPANY_T_1_.\\"area\\" AS D_AREA_2_, COMPANY_T_1_.\\"city\\" AS D_CITY_3_, SUM(COMPANY_T_1_.\\"profit_amt\\") AS D_PROFIT_4_ FROM \\"quickbi_test\\".\\"company_sales_record_copy\\" AS COMPANY_T_1_ WHERE COMPANY_T_1_.\\"area\\" LIKE \\"% China East %\\" GROUP BY COMPANY_T_1_.\\"area\\", COMPANY_T_1_.\\"city\\" HAVING SUM(COMPANY_T_1_.\\"order_amt\\") > 1 LIMIT 0,10
	//
	// example:
	//
	// StandardDimension
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryDataServiceListRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// dasdfdsa-csddf-dsadsa
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDataServiceListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListRequest) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListRequest) SetName(v string) *QueryDataServiceListRequest {
	s.Name = &v
	return s
}

func (s *QueryDataServiceListRequest) SetPageNo(v int32) *QueryDataServiceListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryDataServiceListRequest) SetPageSize(v int32) *QueryDataServiceListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDataServiceListRequest) SetUserId(v string) *QueryDataServiceListRequest {
	s.UserId = &v
	return s
}

type QueryDataServiceListResponseBody struct {
	// example:
	//
	// 78C1AA2D-9201-599E-A0BA-6FC462E57A95
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDataServiceListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataServiceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBody) SetRequestId(v string) *QueryDataServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataServiceListResponseBody) SetResult(v *QueryDataServiceListResponseBodyResult) *QueryDataServiceListResponseBody {
	s.Result = v
	return s
}

func (s *QueryDataServiceListResponseBody) SetSuccess(v bool) *QueryDataServiceListResponseBody {
	s.Success = &v
	return s
}

type QueryDataServiceListResponseBodyResult struct {
	Data []*QueryDataServiceListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s QueryDataServiceListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResult) SetData(v []*QueryDataServiceListResponseBodyResultData) *QueryDataServiceListResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) SetPageNum(v int32) *QueryDataServiceListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) SetPageSize(v int32) *QueryDataServiceListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) SetTotalNum(v int32) *QueryDataServiceListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResult) SetTotalPages(v int32) *QueryDataServiceListResponseBodyResult {
	s.TotalPages = &v
	return s
}

type QueryDataServiceListResponseBodyResultData struct {
	Content *QueryDataServiceListResponseBodyResultDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// 7cb94cd48701
	CreatorId   *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// d14e7448-0eb3-40d3-9375-4afef8de29fd
	CubeId   *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	Desc     *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// 2023-05-18 14:00:02.0
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2023-03-21 18:02:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 7cb94cd48701
	ModifierId   *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	ModifierName *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 862801339
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// dtsuq3i31f5j8v848b
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// example:
	//
	// 7350a155-0e94-4c6c-8620-57bbec38****
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryDataServiceListResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultData) SetContent(v *QueryDataServiceListResponseBodyResultDataContent) *QueryDataServiceListResponseBodyResultData {
	s.Content = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetCreatorId(v string) *QueryDataServiceListResponseBodyResultData {
	s.CreatorId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetCreatorName(v string) *QueryDataServiceListResponseBodyResultData {
	s.CreatorName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetCubeId(v string) *QueryDataServiceListResponseBodyResultData {
	s.CubeId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetCubeName(v string) *QueryDataServiceListResponseBodyResultData {
	s.CubeName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetDesc(v string) *QueryDataServiceListResponseBodyResultData {
	s.Desc = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetGmtCreate(v string) *QueryDataServiceListResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetGmtModified(v string) *QueryDataServiceListResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetModifierId(v string) *QueryDataServiceListResponseBodyResultData {
	s.ModifierId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetModifierName(v string) *QueryDataServiceListResponseBodyResultData {
	s.ModifierName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetName(v string) *QueryDataServiceListResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetOwnerId(v string) *QueryDataServiceListResponseBodyResultData {
	s.OwnerId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetOwnerName(v string) *QueryDataServiceListResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetSid(v string) *QueryDataServiceListResponseBodyResultData {
	s.Sid = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetWorkspaceId(v string) *QueryDataServiceListResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultData) SetWorkspaceName(v string) *QueryDataServiceListResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

type QueryDataServiceListResponseBodyResultDataContent struct {
	// example:
	//
	// 56f9f34a-bdba-496a-91a3-a18b1ff73a80
	CubeId   *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	CubeName *string `json:"CubeName,omitempty" xml:"CubeName,omitempty"`
	// example:
	//
	// true
	Detail       *bool                                                            `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Filter       *QueryDataServiceListResponseBodyResultDataContentFilter         `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	ReturnFields []*QueryDataServiceListResponseBodyResultDataContentReturnFields `json:"ReturnFields,omitempty" xml:"ReturnFields,omitempty" type:"Repeated"`
}

func (s QueryDataServiceListResponseBodyResultDataContent) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultDataContent) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetCubeId(v string) *QueryDataServiceListResponseBodyResultDataContent {
	s.CubeId = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetCubeName(v string) *QueryDataServiceListResponseBodyResultDataContent {
	s.CubeName = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetDetail(v bool) *QueryDataServiceListResponseBodyResultDataContent {
	s.Detail = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetFilter(v *QueryDataServiceListResponseBodyResultDataContentFilter) *QueryDataServiceListResponseBodyResultDataContent {
	s.Filter = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContent) SetReturnFields(v []*QueryDataServiceListResponseBodyResultDataContentReturnFields) *QueryDataServiceListResponseBodyResultDataContent {
	s.ReturnFields = v
	return s
}

type QueryDataServiceListResponseBodyResultDataContentFilter struct {
	Filters []map[string]interface{} `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// example:
	//
	// and
	LogicalOperator *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataServiceListResponseBodyResultDataContentFilter) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultDataContentFilter) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) SetFilters(v []map[string]interface{}) *QueryDataServiceListResponseBodyResultDataContentFilter {
	s.Filters = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) SetLogicalOperator(v string) *QueryDataServiceListResponseBodyResultDataContentFilter {
	s.LogicalOperator = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentFilter) SetType(v string) *QueryDataServiceListResponseBodyResultDataContentFilter {
	s.Type = &v
	return s
}

type QueryDataServiceListResponseBodyResultDataContentReturnFields struct {
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// example:
	//
	// s_number
	Alias *string                                                             `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Desc  *string                                                             `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Field *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField `json:"Field,omitempty" xml:"Field,omitempty" type:"Struct"`
	// example:
	//
	// no
	Orderby *string `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
}

func (s QueryDataServiceListResponseBodyResultDataContentReturnFields) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultDataContentReturnFields) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetAggregator(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Aggregator = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetAlias(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Alias = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetDesc(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Desc = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetField(v *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Field = v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFields) SetOrderby(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFields {
	s.Orderby = &v
	return s
}

type QueryDataServiceListResponseBodyResultDataContentReturnFieldsField struct {
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// example:
	//
	// shid_star
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// example:
	//
	// datetime
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// 1c1f88cb7d
	Fid *string `json:"Fid,omitempty" xml:"Fid,omitempty"`
	// example:
	//
	// yearRegion
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// sss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// dimension
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetCaption(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Caption = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetColumn(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Column = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetDataType(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.DataType = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetFid(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Fid = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetGranularity(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Granularity = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetName(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Name = &v
	return s
}

func (s *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField) SetType(v string) *QueryDataServiceListResponseBodyResultDataContentReturnFieldsField {
	s.Type = &v
	return s
}

type QueryDataServiceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataServiceListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDataServiceListResponse) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponse) SetHeaders(v map[string]*string) *QueryDataServiceListResponse {
	s.Headers = v
	return s
}

func (s *QueryDataServiceListResponse) SetStatusCode(v int32) *QueryDataServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataServiceListResponse) SetBody(v *QueryDataServiceListResponseBody) *QueryDataServiceListResponse {
	s.Body = v
	return s
}

type QueryDatasetDetailInfoRequest struct {
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5820f58c-c734-4d8a-baf1-7979af4f****
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
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the dataset data in JSON format: `{ "cube": { "dimensions": [ { "caption": "customer name", "dataType": "string", "dimensionType": "standard_dimension", "factColumn": "customer_name", "uid": "N5820f5_customer_name" }, { "caption": "datastring", "" standard_dimension", "factColumn": "order_id", "uid": "N5820f5_order_id" }, ], "measures": [ { "caption": "order amount ", "dataType": "number", "factColumn": "order_amt", "measureType": "standard_measure ": " Nderamid " }, " { "customsql": false, "dsId": "261b252d-c3c3-498a-a0a7-5d1ec6cd****", "tableName": "company_sales_record_copy" } }, "datasetId": "5820f58c-c734-4d8a-baf1-7979af4f****", "datasetName": "company_sales_record_copy12", "datasource": { "dsId": "261b252d-c3c3-498a-a0a7-5d1ec6cd****", "dsName": "Self-use", "dsType": "mysql" }, "directory" { "id": "schemaad8aad00-9c55-4984-a767-b4e0ec60****", "name": "My dataset", "pathId": "schemaad8aad00-9c55-4984-a767-b4e0ec60****", "pathName": "My dataset" }, "ownerId": "13651626232****", "ownerName": "Zhang San", "rowLevel": false, "workspaceId": "95296e95-ca89-4c7d-8af9-dedf0ad0****", "workspaceName": "Test Workspace" }`
	//
	// example:
	//
	// A JSON dataset is returned. For more information, see the description on the left.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Queries information about a specified dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// a201c85c-******
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
	// Whether the operation is successfully returned. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// a4d1a221d-41za1-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the request.
	Result *QueryDatasetInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The unique ID of the dataset.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The unique ID of the dataset.
	CubeTableList []*QueryDatasetInfoResponseBodyResultCubeTableList `json:"CubeTableList,omitempty" xml:"CubeTableList,omitempty" type:"Repeated"`
	// The unique ID of the workspace to which the dataset belongs.
	//
	// example:
	//
	// false
	CustimzeSql *bool `json:"CustimzeSql,omitempty" xml:"CustimzeSql,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- mysql
	//
	// 	- odps
	//
	// 	- oracle
	//
	// 	- ... Data source types supported by Quick BI such as
	//
	// example:
	//
	// a201c85c-******
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The user ID of the dataset owner in the Quick BI.
	//
	// example:
	//
	// opds_40
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// If it is a custom SQL table, this is the specific SQL.
	DimensionList []*QueryDatasetInfoResponseBodyResultDimensionList `json:"DimensionList,omitempty" xml:"DimensionList,omitempty" type:"Repeated"`
	// The unique ID of the metric.
	Directory *QueryDatasetInfoResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The name of the data source.
	//
	// example:
	//
	// a201c85c-******
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// The time when the dataset was last modified.
	//
	// example:
	//
	// odps
	DsName *string `json:"DsName,omitempty" xml:"DsName,omitempty"`
	// The point in time when the training dataset was created.
	//
	// example:
	//
	// odps
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// Indicates whether to customize SQL statements. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 1629450382000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The information about the dataset.
	//
	// example:
	//
	// 1629450382000
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// The unique ID of the table to which the table belongs, which corresponds to the UniqueId of the CubeTypeList.
	MeasureList             []*QueryDatasetInfoResponseBodyResultMeasureList `json:"MeasureList,omitempty" xml:"MeasureList,omitempty" type:"Repeated"`
	OpenOfflineAcceleration *bool                                            `json:"OpenOfflineAcceleration,omitempty" xml:"OpenOfflineAcceleration,omitempty"`
	// Test Space
	//
	// example:
	//
	// b8494aab26124*****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unique ID of the data source.
	//
	// example:
	//
	// The name of the dataset owner.
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The name of the training dataset.
	//
	// example:
	//
	// false
	RowLevel *bool `json:"RowLevel,omitempty" xml:"RowLevel,omitempty"`
	// Whether row-level permissions are enabled. Valid values:
	//
	// 	- true: The VIP Netty channel is enabled.
	//
	// 	- false: The VIP Netty channel is disabled.
	//
	// example:
	//
	// 420abef4-a79b-4289-b12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// Big Baby
	//
	// example:
	//
	// The name of the workspace in which the dataset resides.
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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

func (s *QueryDatasetInfoResponseBodyResult) SetOpenOfflineAcceleration(v bool) *QueryDatasetInfoResponseBodyResult {
	s.OpenOfflineAcceleration = &v
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
	// Indicates whether the data source table is valid. Valid values:
	//
	// 	- true: data source table
	//
	// 	- false: custom table
	//
	// example:
	//
	// odps_40
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The display name of the table.
	//
	// example:
	//
	// false
	Customsql *bool `json:"Customsql,omitempty" xml:"Customsql,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// dfefd7f4-fc6e-42c9-b4******
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// maxcompute
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// The unique ID of the table.
	//
	// example:
	//
	// true
	FactTable *bool `json:"FactTable,omitempty" xml:"FactTable,omitempty"`
	// Indicates whether the table is a custom SQL table. Valid values:
	//
	// 	- true: custom SQL table
	//
	// 	- false: non-custom SQL table
	//
	// example:
	//
	// select 	- from ****
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The list of tables used by the dataset.
	//
	// example:
	//
	// viewdasb8494aab2612473cb74992159fe****
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- mysql
	//
	// 	- odps
	//
	// 	- oracle
	//
	// 	- ... and other data source types supported by Quick BI
	//
	// example:
	//
	// 7a62530b36
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
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
	// The unique ID of the field that is referenced by the group measure. Non-NULL if and only if the metric is a grouping metric.
	//
	// example:
	//
	// city
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// A list of all dimensions in the dataset.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The actual physical field.
	//
	// example:
	//
	// group_dimension
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	// Data type; value:
	//
	// 	- string: character
	//
	// 	- number: a number
	//
	// 	- datetime: time
	//
	// example:
	//
	// example_expression
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// Expression for a calculated dimension; valid only for calculated dimensions.
	//
	// example:
	//
	// city
	FactColumn *string `json:"FactColumn,omitempty" xml:"FactColumn,omitempty"`
	// example:
	//
	// hhhh
	FieldDescription *string `json:"FieldDescription,omitempty" xml:"FieldDescription,omitempty"`
	// The type of the dimension. Valid values:
	//
	// 	- standard_dimension: General Dimension
	//
	// 	- calculate_dimension: calculating dimensions
	//
	// 	- group_dimension: grouping dimensions
	//
	// example:
	//
	// example_granularity
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The granularity.
	//
	// example:
	//
	// 308f7****
	RefUid *string `json:"RefUid,omitempty" xml:"RefUid,omitempty"`
	// The ARN.
	//
	// example:
	//
	// 7a62530***
	TableUniqueId *string `json:"TableUniqueId,omitempty" xml:"TableUniqueId,omitempty"`
	// The display name of the dimension.
	//
	// example:
	//
	// a69774***
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
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

func (s *QueryDatasetInfoResponseBodyResultDimensionList) SetFieldDescription(v string) *QueryDatasetInfoResponseBodyResultDimensionList {
	s.FieldDescription = &v
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
	// Test directory
	//
	// example:
	//
	// a3eecab7-618d-4f9f-*****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Test directory
	//
	// example:
	//
	// The name of the directory.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the directory to which the dataset belongs.
	//
	// example:
	//
	// 88b680****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The path of the directory ID, for example, aa/bb/cc/dd.
	//
	// example:
	//
	// The path name of the directory ID, for example, one-level directory /two-level directory.
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
	// The actual physical field.
	//
	// example:
	//
	// profit_amt
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// A list of all measures for the dataset.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Data type; value:
	//
	// 	- string: character
	//
	// 	- number: a number
	//
	// 	- datetime: time
	//
	// example:
	//
	// example_expression
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The type of the measure. Valid values:
	//
	// 	- standard_measure: General Metrics
	//
	// 	- calculate_measure: Calculating Measures
	//
	// example:
	//
	// profit_amt
	FactColumn *string `json:"FactColumn,omitempty" xml:"FactColumn,omitempty"`
	// example:
	//
	// hhhh
	FieldDescription *string `json:"FieldDescription,omitempty" xml:"FieldDescription,omitempty"`
	// An expression that calculates a measure; valid only for calculated measures.
	//
	// example:
	//
	// calculate_measure
	MeasureType *string `json:"MeasureType,omitempty" xml:"MeasureType,omitempty"`
	// The display name of the metric.
	//
	// example:
	//
	// 7a62530b36
	TableUniqueId *string `json:"TableUniqueId,omitempty" xml:"TableUniqueId,omitempty"`
	// The unique ID of the table to which the table belongs, which corresponds to the UniqueId of the CubeTypeList.
	//
	// example:
	//
	// 88b680****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
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

func (s *QueryDatasetInfoResponseBodyResultMeasureList) SetFieldDescription(v string) *QueryDatasetInfoResponseBodyResultMeasureList {
	s.FieldDescription = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// schemaad8aad00-9c55-4984-a767-b4e0ec60****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Information about the directory where the dataset is located
	//
	// example:
	//
	// Queries the datasets of a specified workspace. The datasets are sorted in descending order by creation time.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Specifies the directory ID.
	//
	// 	- If this field is not empty, all datasets in the directory are obtained.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// true
	WithChildren *bool `json:"WithChildren,omitempty" xml:"WithChildren,omitempty"`
	// The name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// The keyword used to search for the dataset name.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Test dataset
	Result *QueryDatasetListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Whether to recursively wrap the dataset in the subdirectory. Valid values:
	//
	// 	- true: returns datasets in all recursive subdirectories in the directoryId directory.
	//
	// 	- false: Only datasets in the directory specified by directoryId are returned, excluding subdirectories.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// Returns the pagination results of the dataset list. The detailed information of the dataset list is stored in the response parameter Data.
	Data []*QueryDatasetListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The number of rows per page in a paged query.
	//
	// 	- Default value: 10.
	//
	// 	- Maximum value: 1,000.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// Current page number for dataset list:
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// The details of the dataset list.
	//
	// example:
	//
	// 2020-11-02 10:36:05
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Test Space
	DataSource *QueryDatasetListResponseBodyResultDataDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The name of the workspace.
	//
	// example:
	//
	// 5820f58c-c734-4d8a-baf1-7979af4f****
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// Tom
	//
	// example:
	//
	// company_sales_record_copy12
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The number of rows per page set when the interface is requested.
	//
	// example:
	//
	// The total number of rows in the table.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the data source to which the dataset belongs.
	Directory *QueryDatasetListResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The nickname of the dataset owner.
	//
	// example:
	//
	// 2020-11-02 10:36:05
	ModifyTime              *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OpenOfflineAcceleration *bool   `json:"OpenOfflineAcceleration,omitempty" xml:"OpenOfflineAcceleration,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 136516262323****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Whether to enable row-level permissions. Valid values:
	//
	// 	- true: The VIP Netty channel is enabled.
	//
	// 	- false: The incremental log backup feature is disabled.
	//
	// example:
	//
	// The ID of the workspace.
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// true
	RowLevel *bool `json:"RowLevel,omitempty" xml:"RowLevel,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad06adf
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The description of the dataset.
	//
	// example:
	//
	// Test dataset
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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

func (s *QueryDatasetListResponseBodyResultData) SetOpenOfflineAcceleration(v bool) *QueryDatasetListResponseBodyResultData {
	s.OpenOfflineAcceleration = &v
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
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// example:
	//
	// 261b252d-c3c3-498a-a0a7-5d1ec6cd****
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// The time when the scaling group was modified.
	//
	// example:
	//
	// The name of the training dataset.
	DsName *string `json:"DsName,omitempty" xml:"DsName,omitempty"`
	// The user ID of the dataset owner in the Quick BI.
	//
	// example:
	//
	// mysql
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
	// The ID of the directory path.
	//
	// example:
	//
	// schemaad8aad00-9c55-4984-a767-b4e0ec60****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// Information about the directory where the dataset is located
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// schemaad8aad00-9c55-4984-a767-b4e0ec60****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// Test a data source
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
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
	// example:
	//
	// FAECEFA8-09BB-58AB-BC58-C8ACEFE4D232
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryDatasetSwitchInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// example:
	//
	// 1
	IsOpenColumnLevelPermission *int32 `json:"IsOpenColumnLevelPermission,omitempty" xml:"IsOpenColumnLevelPermission,omitempty"`
	// example:
	//
	// 1
	IsOpenRowLevelPermission *int32 `json:"IsOpenRowLevelPermission,omitempty" xml:"IsOpenRowLevelPermission,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetSwitchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryEmbeddedInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Detail *QueryEmbeddedInfoResponseBodyResultDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// example:
	//
	// 3
	EmbeddedCount *int32 `json:"EmbeddedCount,omitempty" xml:"EmbeddedCount,omitempty"`
	// example:
	//
	// 100
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
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
	// example:
	//
	// 1
	DashboardOfflineQuery *int32 `json:"DashboardOfflineQuery,omitempty" xml:"DashboardOfflineQuery,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 1
	Report *int32 `json:"Report,omitempty" xml:"Report,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEmbeddedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The work ID of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the work is enabled for embedding. Valid values:
	//
	// 	- true: embedded
	//
	// 	- false: not embedded
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEmbeddedStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryOrganizationRoleConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111111111
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s QueryOrganizationRoleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationRoleConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigRequest) SetRoleId(v int64) *QueryOrganizationRoleConfigRequest {
	s.RoleId = &v
	return s
}

type QueryOrganizationRoleConfigResponseBody struct {
	// example:
	//
	// BCE45E6D-9304-4F94-86BB-5A772B1615FF
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryOrganizationRoleConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOrganizationRoleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationRoleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigResponseBody) SetRequestId(v string) *QueryOrganizationRoleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBody) SetResult(v *QueryOrganizationRoleConfigResponseBodyResult) *QueryOrganizationRoleConfigResponseBody {
	s.Result = v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBody) SetSuccess(v bool) *QueryOrganizationRoleConfigResponseBody {
	s.Success = &v
	return s
}

type QueryOrganizationRoleConfigResponseBodyResult struct {
	AuthConfigList []*QueryOrganizationRoleConfigResponseBodyResultAuthConfigList `json:"AuthConfigList,omitempty" xml:"AuthConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsSystemRole *bool `json:"IsSystemRole,omitempty" xml:"IsSystemRole,omitempty"`
	// example:
	//
	// 111111111
	RoleId   *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s QueryOrganizationRoleConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationRoleConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) SetAuthConfigList(v []*QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) *QueryOrganizationRoleConfigResponseBodyResult {
	s.AuthConfigList = v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) SetIsSystemRole(v bool) *QueryOrganizationRoleConfigResponseBodyResult {
	s.IsSystemRole = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) SetRoleId(v int64) *QueryOrganizationRoleConfigResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) SetRoleName(v string) *QueryOrganizationRoleConfigResponseBodyResult {
	s.RoleName = &v
	return s
}

type QueryOrganizationRoleConfigResponseBodyResultAuthConfigList struct {
	// example:
	//
	// quick_monitor
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
}

func (s QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) SetAuthKey(v string) *QueryOrganizationRoleConfigResponseBodyResultAuthConfigList {
	s.AuthKey = &v
	return s
}

type QueryOrganizationRoleConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrganizationRoleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrganizationRoleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrganizationRoleConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigResponse) SetHeaders(v map[string]*string) *QueryOrganizationRoleConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryOrganizationRoleConfigResponse) SetStatusCode(v int32) *QueryOrganizationRoleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponse) SetBody(v *QueryOrganizationRoleConfigResponseBody) *QueryOrganizationRoleConfigResponse {
	s.Body = v
	return s
}

type QueryOrganizationWorkspaceListRequest struct {
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryOrganizationWorkspaceListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Data []*QueryOrganizationWorkspaceListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// example:
	//
	// true
	AllowPublishOperation *bool `json:"AllowPublishOperation,omitempty" xml:"AllowPublishOperation,omitempty"`
	// example:
	//
	// true
	AllowShareOperation *bool `json:"AllowShareOperation,omitempty" xml:"AllowShareOperation,omitempty"`
	// example:
	//
	// 2020-11-10 17:51:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// pop****@aliyun.com
	CreateUserAccountName *string `json:"CreateUserAccountName,omitempty" xml:"CreateUserAccountName,omitempty"`
	// example:
	//
	// 2020-11-10 17:51:07
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// pop****@aliyun.com
	ModifyUserAccountName *string `json:"ModifyUserAccountName,omitempty" xml:"ModifyUserAccountName,omitempty"`
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// pop****@aliyun.com
	OwnerAccountName     *string `json:"OwnerAccountName,omitempty" xml:"OwnerAccountName,omitempty"`
	WorkspaceDescription *string `json:"WorkspaceDescription,omitempty" xml:"WorkspaceDescription,omitempty"`
	// example:
	//
	// 7350a155-0e94-4c6c-8620-57bbec38****
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrganizationWorkspaceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Quick BI the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
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
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of works that the user has permission to view.
	Result []*QueryReadableResourcesListByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The timestamp of the creation time in milliseconds.
	//
	// example:
	//
	// 1611023338000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Remarks on the work.
	//
	// example:
	//
	// Description
	Description *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory   *QueryReadableResourcesListByUserIdResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The name of the Alibaba Cloud account to which the modifier belongs.
	//
	// example:
	//
	// Li Si
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// example:
	//
	// 1611023338000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The Quick BI UserID of the work owner.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud account name of the owner.
	//
	// example:
	//
	// Tom
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Security policies for collaborative authorization of works. Valid values:
	//
	// 	- 0: private
	//
	// 	- 12: Authorize specified members
	//
	// 	- 1 or 11: Authorize all workspace members
	//
	// >
	//
	// 	- If you use legacy permissions, the return value is 1.
	//
	// 	- If you use the new permissions, the return value is 11.
	//
	// example:
	//
	// 0
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The status of the report. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 1: published
	//
	// 	- 2: modified but not published
	//
	// 	- 3: unpublished
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Third-party embedding status. Valid values:
	//
	// 	- 0: The embed service is not enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 1
	ThirdPartAuthFlag *int32 `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The name of the work.
	WorkName *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// The type of the work. Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- FULLPAGE: full-screen dashboards
	//
	// 	- REPORT: workbook
	//
	// example:
	//
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 03366b16-69ce-43c8-b782-56c2f6ec****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// Test Workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	// example:
	//
	// e4276ea5-b232-4fb1-8f0f-efcee4a2****
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// e4276ea5-b232-4fb1-8f0f-efcee4a2****
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReadableResourcesListByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryReportPerformanceRequest struct {
	// example:
	//
	// 1
	CostTimeAvgMin *int32 `json:"CostTimeAvgMin,omitempty" xml:"CostTimeAvgMin,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sevenDays
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// example:
	//
	// report
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryReportPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReportPerformanceRequest) GoString() string {
	return s.String()
}

func (s *QueryReportPerformanceRequest) SetCostTimeAvgMin(v int32) *QueryReportPerformanceRequest {
	s.CostTimeAvgMin = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetPageNum(v int32) *QueryReportPerformanceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetPageSize(v int32) *QueryReportPerformanceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetQueryType(v string) *QueryReportPerformanceRequest {
	s.QueryType = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetReportId(v string) *QueryReportPerformanceRequest {
	s.ReportId = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetResourceType(v string) *QueryReportPerformanceRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetWorkspaceId(v string) *QueryReportPerformanceRequest {
	s.WorkspaceId = &v
	return s
}

type QueryReportPerformanceResponseBody struct {
	// example:
	//
	// 1FC71085-D5FD-08E0-813A-4D4BD1031BC5
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryReportPerformanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryReportPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReportPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReportPerformanceResponseBody) SetRequestId(v string) *QueryReportPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReportPerformanceResponseBody) SetResult(v []*QueryReportPerformanceResponseBodyResult) *QueryReportPerformanceResponseBody {
	s.Result = v
	return s
}

func (s *QueryReportPerformanceResponseBody) SetSuccess(v bool) *QueryReportPerformanceResponseBody {
	s.Success = &v
	return s
}

type QueryReportPerformanceResponseBodyResult struct {
	// example:
	//
	// 2.2
	CacheCostTimeAvg *float64 `json:"CacheCostTimeAvg,omitempty" xml:"CacheCostTimeAvg,omitempty"`
	// example:
	//
	// 1
	CacheQueryCount *int32 `json:"CacheQueryCount,omitempty" xml:"CacheQueryCount,omitempty"`
	// example:
	//
	// 1
	ComponentQueryCount *int32 `json:"ComponentQueryCount,omitempty" xml:"ComponentQueryCount,omitempty"`
	// example:
	//
	// 2.0
	ComponentQueryCountAvg *float64 `json:"ComponentQueryCountAvg,omitempty" xml:"ComponentQueryCountAvg,omitempty"`
	// example:
	//
	// 0.2
	CostTimeAvg *float64 `json:"CostTimeAvg,omitempty" xml:"CostTimeAvg,omitempty"`
	// example:
	//
	// 50
	QueryCount *int32 `json:"QueryCount,omitempty" xml:"QueryCount,omitempty"`
	// example:
	//
	// 3.3
	QueryCountAvg *float64 `json:"QueryCountAvg,omitempty" xml:"QueryCountAvg,omitempty"`
	// example:
	//
	// 0.5
	QueryOverFivePercentNum *float64 `json:"QueryOverFivePercentNum,omitempty" xml:"QueryOverFivePercentNum,omitempty"`
	// example:
	//
	// 0.5
	QueryOverFiveSecPercent *string `json:"QueryOverFiveSecPercent,omitempty" xml:"QueryOverFiveSecPercent,omitempty"`
	// example:
	//
	// 0.5
	QueryOverTenSecPercent *string `json:"QueryOverTenSecPercent,omitempty" xml:"QueryOverTenSecPercent,omitempty"`
	// example:
	//
	// 0.5
	QueryOverTenSecPercentNum *float64 `json:"QueryOverTenSecPercentNum,omitempty" xml:"QueryOverTenSecPercentNum,omitempty"`
	// example:
	//
	// 8
	QueryTimeoutCount *int32 `json:"QueryTimeoutCount,omitempty" xml:"QueryTimeoutCount,omitempty"`
	// example:
	//
	// 0.5
	QueryTimeoutCountPercent *float64 `json:"QueryTimeoutCountPercent,omitempty" xml:"QueryTimeoutCountPercent,omitempty"`
	// example:
	//
	// 10
	QuickIndexCostTimeAvg *float64 `json:"QuickIndexCostTimeAvg,omitempty" xml:"QuickIndexCostTimeAvg,omitempty"`
	// example:
	//
	// 5
	QuickIndexQueryCount *int32 `json:"QuickIndexQueryCount,omitempty" xml:"QuickIndexQueryCount,omitempty"`
	// example:
	//
	// 0.8
	RepeatQueryPercent *string `json:"RepeatQueryPercent,omitempty" xml:"RepeatQueryPercent,omitempty"`
	// example:
	//
	// 3
	RepeatQueryPercentNum *float64 `json:"RepeatQueryPercentNum,omitempty" xml:"RepeatQueryPercentNum,omitempty"`
	// example:
	//
	// 1
	RepeatSqlQueryCount *int32 `json:"RepeatSqlQueryCount,omitempty" xml:"RepeatSqlQueryCount,omitempty"`
	// example:
	//
	// 0.7
	RepeatSqlQueryPercent *string `json:"RepeatSqlQueryPercent,omitempty" xml:"RepeatSqlQueryPercent,omitempty"`
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// example:
	//
	// ClusterAddonUpgradeReport
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// example:
	//
	// report
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// ab46ed33-6278-4ef7-8013-8c1335f266ee
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s QueryReportPerformanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryReportPerformanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryReportPerformanceResponseBodyResult) SetCacheCostTimeAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.CacheCostTimeAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetCacheQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.CacheQueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetComponentQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.ComponentQueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetComponentQueryCountAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.ComponentQueryCountAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetCostTimeAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.CostTimeAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.QueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryCountAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QueryCountAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryOverFivePercentNum(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QueryOverFivePercentNum = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryOverFiveSecPercent(v string) *QueryReportPerformanceResponseBodyResult {
	s.QueryOverFiveSecPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryOverTenSecPercent(v string) *QueryReportPerformanceResponseBodyResult {
	s.QueryOverTenSecPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryOverTenSecPercentNum(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QueryOverTenSecPercentNum = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryTimeoutCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.QueryTimeoutCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQueryTimeoutCountPercent(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QueryTimeoutCountPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQuickIndexCostTimeAvg(v float64) *QueryReportPerformanceResponseBodyResult {
	s.QuickIndexCostTimeAvg = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetQuickIndexQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.QuickIndexQueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetRepeatQueryPercent(v string) *QueryReportPerformanceResponseBodyResult {
	s.RepeatQueryPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetRepeatQueryPercentNum(v float64) *QueryReportPerformanceResponseBodyResult {
	s.RepeatQueryPercentNum = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetRepeatSqlQueryCount(v int32) *QueryReportPerformanceResponseBodyResult {
	s.RepeatSqlQueryCount = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetRepeatSqlQueryPercent(v string) *QueryReportPerformanceResponseBodyResult {
	s.RepeatSqlQueryPercent = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetReportId(v string) *QueryReportPerformanceResponseBodyResult {
	s.ReportId = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetReportName(v string) *QueryReportPerformanceResponseBodyResult {
	s.ReportName = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetReportType(v string) *QueryReportPerformanceResponseBodyResult {
	s.ReportType = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetWorkspaceId(v string) *QueryReportPerformanceResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryReportPerformanceResponseBodyResult) SetWorkspaceName(v string) *QueryReportPerformanceResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

type QueryReportPerformanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReportPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReportPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReportPerformanceResponse) GoString() string {
	return s.String()
}

func (s *QueryReportPerformanceResponse) SetHeaders(v map[string]*string) *QueryReportPerformanceResponse {
	s.Headers = v
	return s
}

func (s *QueryReportPerformanceResponse) SetStatusCode(v int32) *QueryReportPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReportPerformanceResponse) SetBody(v *QueryReportPerformanceResponseBody) *QueryReportPerformanceResponse {
	s.Body = v
	return s
}

type QueryShareListRequest struct {
	// The type of work being shared. Valid values:
	//
	// 	- product: BI portal
	//
	// 	- dashboard: dashboard
	//
	// 	- worksheet: workbook
	//
	// 	- dashboardOfflineQuery: self-service data retrieval
	//
	// 	- Analysis: Ad hoc analysis
	//
	// 	- DATAFORM
	//
	// 	- SCREEN: Data dashboard
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
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
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryShareListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 3
	AuthPoint *int32 `json:"AuthPoint,omitempty" xml:"AuthPoint,omitempty"`
	// example:
	//
	// 1640102400000
	ExpireDate *int64 `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// example:
	//
	// 0ab9659e-29cf-47d7-a364-3a91553b****
	ShareId *string `json:"ShareId,omitempty" xml:"ShareId,omitempty"`
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	ShareToId   *string `json:"ShareToId,omitempty" xml:"ShareToId,omitempty"`
	ShareToName *string `json:"ShareToName,omitempty" xml:"ShareToName,omitempty"`
	// example:
	//
	// 0
	ShareToType *int32 `json:"ShareToType,omitempty" xml:"ShareToType,omitempty"`
	// example:
	//
	// product
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryShareListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the user. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e53****5ba4b679ee22e2a2927****
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
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns a list of works authorized to the user.
	Result []*QuerySharesToUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The timestamp of the creation time in milliseconds.
	//
	// example:
	//
	// 1530078690000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Remarks on the work.
	//
	// example:
	//
	// Description of the test report
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Information about the directory where the work is located.
	Directory *QuerySharesToUserListResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The name of the Alibaba Cloud account to which the modifier belongs.
	//
	// example:
	//
	// 13855265****@163.com
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The timestamp of the modification time in milliseconds.
	//
	// example:
	//
	// 1530078690000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The UserID of the work owner in Quickbi.
	//
	// example:
	//
	// 74f5527216d14e9892245320ebf2****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud account name of the work owner.
	//
	// example:
	//
	// w****@aliyun.com
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Security policies for collaborative authorization of works. Valid values:
	//
	// 	- 0: private
	//
	// 	- 12: Authorize specified members
	//
	// 	- 1 or 11: Authorize all workspace members
	//
	// >
	//
	// 	- If you use legacy permissions, the return value is 1.
	//
	// 	- If you use the new permissions, the return value is 11.
	//
	// example:
	//
	// 0
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The publishing status of the report. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 1: published
	//
	// 	- 2: modified but not published
	//
	// 	- 3: unpublished
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Third-party embedding status. Valid values:
	//
	// 	- 0: No embedding is enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 0
	ThirdPartAuthFlag *int32 `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The name of the report.
	//
	// example:
	//
	// Test report
	WorkName *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// The type of the work. Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- FULLPAGE: full-screen dashboards
	//
	// 	- REPORT: workbook
	//
	// 	- dashboardOfflineQuery: self-service data retrieval
	//
	// example:
	//
	// DATAFORM
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The ID of the operations report.
	//
	// example:
	//
	// 97f7f4c1-543a-4069-8e8d-a56cfcd6****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// The ID of the workspace to which the report belongs.
	//
	// example:
	//
	// c5f86ad2-ef53-4c51-8720-162ecfdb****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the report belongs.
	//
	// example:
	//
	// Return to Professional Edition
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	// The ID of the directory where the resource is located.
	//
	// example:
	//
	// f7f6e22b-83be-47fd-b49d-9ca686a9****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// Chart Report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path ID of the directory where the resource is located.
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The path name of the directory where the resource is located.
	//
	// example:
	//
	// Level -1 Directory /Level -2 Directory
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySharesToUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Obtains the details of a specified ticket for a report that is not embedded in the report.
	//
	// This parameter is required.
	//
	// example:
	//
	// a27a9aec-****-****-bd40-1a21ea41d7c5
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryTicketInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// a27a9aec-****-****-bd40-1a21ea41d7c5
	AccessTicket *string `json:"AccessTicket,omitempty" xml:"AccessTicket,omitempty"`
	// example:
	//
	// sfdgsds-****-****-a608-mghdgd
	CmptId      *string `json:"CmptId,omitempty" xml:"CmptId,omitempty"`
	GlobalParam *string `json:"GlobalParam,omitempty" xml:"GlobalParam,omitempty"`
	// example:
	//
	// 2022-01-30 03:03:49
	InvalidTime *string `json:"InvalidTime,omitempty" xml:"InvalidTime,omitempty"`
	// example:
	//
	// 9999
	MaxTicketNum *int32 `json:"MaxTicketNum,omitempty" xml:"MaxTicketNum,omitempty"`
	// example:
	//
	// 2fe4fbd8-****-****-b3e1-e92c7af083ea
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// 2022-01-09 22:23:49
	RegisterTime *string `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// example:
	//
	// 47
	UsedTicketNum *int32 `json:"UsedTicketNum,omitempty" xml:"UsedTicketNum,omitempty"`
	// example:
	//
	// 974e50**********9033f46
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WatermarkParam *string `json:"WatermarkParam,omitempty" xml:"WatermarkParam,omitempty"`
	// example:
	//
	// ccd3428c-****-****-a608-26bae29dffee
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTicketInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the parent user group.
	//
	// 	- If you enter the ID of the parent user group, you can obtain the information of the child user group under this ID.
	//
	// 	- If you enter -1, you can obtain the sub-user group information under the root directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3d2c23d4-2b41-4af8-a1f5-f6390f32****
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
	// The ID of the request.
	//
	// example:
	//
	// 72B19D61-B37A-5C7A-9389-0856CD7935B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the sub-user group.
	Result []*QueryUserGroupListByParentIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The time when the sub-user group was created.
	//
	// example:
	//
	// 2020-10-30 10:03:09
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the sub-user group. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// example:
	//
	// 136516262323****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Directory level of the sub-user group.
	IdentifiedPath *string `json:"IdentifiedPath,omitempty" xml:"IdentifiedPath,omitempty"`
	// The time when the sub-user group was last modified.
	//
	// example:
	//
	// 2020-11-16 15:49:08
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The user who modified the subgroup. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// example:
	//
	// 136516262323****
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The ID of the parent user group.
	//
	// example:
	//
	// 3d2c23d4-2b41-4af8-a1f5-f6390f32****
	ParentUserGroupId *string `json:"ParentUserGroupId,omitempty" xml:"ParentUserGroupId,omitempty"`
	// The description of the sub-user group.
	//
	// example:
	//
	// User Group for Testing
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	// The ID of the sub-user group.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the sub-user group.
	//
	// example:
	//
	// popapi test group
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserGroupListByParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
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
	// example:
	//
	// 48C930FF-DFCF-5986-902B-E24C202E2443
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryUserGroupMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 3d2c23d4-2b41-4af8-a1f5-f6390f32****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsUserGroup *bool   `json:"IsUserGroup,omitempty" xml:"IsUserGroup,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2fe4fbd8-588f-489a-b3e1-e92c7af0****
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Enter the name or ID of the Alibaba Cloud account that you want to query.
	//
	// 	- When you enter an account name:
	//
	//     	- If the organization user is a master account, such as main_account, the query account format is master account. That is, the main account main_account to be entered.
	//
	//     	- If the organization user is a RAM user, such as a <zhangsan@test.onaliyun.com>, the query account format is the head of the RAM user, that is, the RAM user to be entered is zhangsan.
	//
	// 	- ID：
	//
	//     	- Enter the UID of the account to query the account information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1386587****@163.com
	Account           *string `json:"Account,omitempty" xml:"Account,omitempty"`
	ParentAccountName *string `json:"ParentAccountName,omitempty" xml:"ParentAccountName,omitempty"`
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

func (s *QueryUserInfoByAccountRequest) SetParentAccountName(v string) *QueryUserInfoByAccountRequest {
	s.ParentAccountName = &v
	return s
}

type QueryUserInfoByAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned organization user information.
	Result *QueryUserInfoByAccountResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 135****5848
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account that corresponds to the member. (If you use a RAM user, the domain name information that follows @ is removed. For example, if you use a <test@test.com>, test is returned.)
	//
	// example:
	//
	// 1386587****@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Whether you are an administrator of the organization. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Whether you are a permission administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// 1386587****@163.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The nickname of the account.
	//
	// example:
	//
	// Test user
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The phone number of the alert contact.
	//
	// example:
	//
	// 1386587****
	Phone      *string  `json:"Phone,omitempty" xml:"Phone,omitempty"`
	RoleIdList []*int64 `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Repeated"`
	// The UserID in the Quick BI.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The role type of the organization member. Valid values:
	//
	// 	- 1 : developer
	//
	// 	- 2 : visitors
	//
	// 	- 3 : Analyst
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
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

func (s *QueryUserInfoByAccountResponseBodyResult) SetRoleIdList(v []*int64) *QueryUserInfoByAccountResponseBodyResult {
	s.RoleIdList = v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserInfoByAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the user. The UserID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
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
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned organization user information.
	Result *QueryUserInfoByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 135****5848
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account that corresponds to the member.
	//
	// example:
	//
	// 1386587****@163.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Whether you are an administrator of the organization. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Whether you are a permission administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// 1386587****@163.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The nickname of the account.
	//
	// example:
	//
	// Test user
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The phone number of the alert contact.
	//
	// example:
	//
	// 1386587****
	Phone      *string  `json:"Phone,omitempty" xml:"Phone,omitempty"`
	RoleIdList []*int64 `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Repeated"`
	// The UserID in the Quick BI.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The role type of the organization member. Valid values:
	//
	// 	- 1 : developer
	//
	// 	- 2 : visitors
	//
	// 	- 3 : Analyst
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
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

func (s *QueryUserInfoByUserIdResponseBodyResult) SetRoleIdList(v []*int64) *QueryUserInfoByUserIdResponseBodyResult {
	s.RoleIdList = v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserInfoByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Data []*QueryUserListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// example:
	//
	// 1355********
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// example:
	//
	// true
	AuthAdminUser *bool    `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	JoinedDate    *int64   `json:"JoinedDate,omitempty" xml:"JoinedDate,omitempty"`
	LastLoginTime *int64   `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	NickName      *string  `json:"NickName,omitempty" xml:"NickName,omitempty"`
	RoleIdList    []*int64 `json:"RoleIdList,omitempty" xml:"RoleIdList,omitempty" type:"Repeated"`
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
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

func (s *QueryUserListResponseBodyResultData) SetJoinedDate(v int64) *QueryUserListResponseBodyResultData {
	s.JoinedDate = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetLastLoginTime(v int64) *QueryUserListResponseBodyResultData {
	s.LastLoginTime = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetNickName(v string) *QueryUserListResponseBodyResultData {
	s.NickName = &v
	return s
}

func (s *QueryUserListResponseBodyResultData) SetRoleIdList(v []*int64) *QueryUserListResponseBodyResultData {
	s.RoleIdList = v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// f5698bedeb384b1986afccd9e434****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryUserRoleInfoInWorkspaceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// role_workspace_admin
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// example:
	//
	// 25
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserRoleInfoInWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryUserTagMetaListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Queries the metadata list of member tags in an organization.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// pop_001
	TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserTagMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*QueryUserTagValueListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// pop_001
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserTagValueListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// abcd****
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryWorksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 0
	Auth3rdFlag *int32                                 `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	Description *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory   *QueryWorksResponseBodyResultDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// example:
	//
	// 1496651577000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1496651577000
	GmtModify  *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// example:
	//
	// 9187a612aa474e2a8ac1414d5529****
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName         *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	PublicFlag        *bool   `json:"PublicFlag,omitempty" xml:"PublicFlag,omitempty"`
	PublicInvalidTime *int64  `json:"PublicInvalidTime,omitempty" xml:"PublicInvalidTime,omitempty"`
	// example:
	//
	// 0
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// example:
	//
	// 0
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkName *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// example:
	//
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// example:
	//
	// 87c6b145-090c-43e1-9426-8f93be23****
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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

func (s *QueryWorksResponseBodyResult) SetPublicFlag(v bool) *QueryWorksResponseBodyResult {
	s.PublicFlag = &v
	return s
}

func (s *QueryWorksResponseBodyResult) SetPublicInvalidTime(v int64) *QueryWorksResponseBodyResult {
	s.PublicInvalidTime = &v
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
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Obtains the kinship of a data work, including the datasets referenced by each component and query field information. Currently, only supported data works include dashboards, workbooks, and self-service data retrieval.
	//
	// This parameter is required.
	//
	// example:
	//
	// abcd****
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
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the request.
	Result []*QueryWorksBloodRelationshipResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The response.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// List of work blood information.
	//
	// example:
	//
	// 0696083a-ca72-4d89-8e7a-c017910e0***
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The ID of the component that you want to modify.
	//
	// example:
	//
	// The name of the component.
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// Line
	//
	// example:
	//
	// 3
	ComponentType *int32 `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The type of the image component.
	//
	// example:
	//
	// LINE
	ComponentTypeName *string `json:"ComponentTypeName,omitempty" xml:"ComponentTypeName,omitempty"`
	// Column (Measure)
	//
	// example:
	//
	// dc78a4ed-880d-452e-b017-90cfc10c83e5_company_sales_record
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The name of the component type.
	QueryParams []*QueryWorksBloodRelationshipResponseBodyResultQueryParams `json:"QueryParams,omitempty" xml:"QueryParams,omitempty" type:"Repeated"`
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
	// Indices whether the metric. Valid values:
	//
	// true false
	//
	// example:
	//
	// area_column
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The ID of the owning location.
	//
	// example:
	//
	// The name of the owning location.
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The globally unique PathId.
	//
	// example:
	//
	// order_number
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The display name of the field.
	//
	// example:
	//
	// number
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The type of the field. Valid values:
	//
	// 	- string: string type
	//
	// 	- date: a date type that contains only the year, month, and day parts
	//
	// 	- datetime: a common date type
	//
	// 	- time: a date type that contains only hours, minutes, and seconds.
	//
	// 	- number: numeric
	//
	// 	- boolean: Boolean type
	//
	// 	- geographical: geographical location
	//
	// 	- url: string type
	//
	// 	- imageUrl: the type of the image link.
	//
	// 	- multivalue: a multi-value column
	//
	// example:
	//
	// true
	IsMeasure *bool `json:"IsMeasure,omitempty" xml:"IsMeasure,omitempty"`
	// The unique ID of the field.
	//
	// example:
	//
	// schema7d1944eb-443e-48c6-8123-bf45a99e7e74.dc78a4ed-880d-452e-b017-90cfc10c83e5_company_sales_record.[Ndc78a4_order_level].[Ndc78a4_order_level].[Ndc78a4_order_level]
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// A list of query parameter reference columns.
	//
	// example:
	//
	// Ndc78a4_order_number
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorksBloodRelationshipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page set when the interface is requested.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Returns a list of all works in the organization that meet the requested criteria.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	ThirdPartAuthFlag *int32 `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// PAGE
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
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
	// The details of the list of works.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the report. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 1: published
	//
	// 	- 2: modified but not published
	//
	// 	- 3: unpublished
	Result *QueryWorksByOrganizationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The total number of rows in the table.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The Alibaba Cloud account name of the work owner.
	Data []*QueryWorksByOrganizationResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The timestamp of the modification of the work in milliseconds.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the work. Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- FULLPAGE: full-screen dashboards
	//
	// 	- REPORT: workbook
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// Third-party embedding status. Valid values:
	//
	// 	- 0: The embed service is not enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// 1
	Auth3rdFlag *int32 `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	// The hierarchical structure of the directory ID to which the directory belongs. Separate the hierarchical structure with a /.
	//
	// example:
	//
	// The hierarchical structure of the directory to which the directory belongs. Separate the hierarchical structure with a (/).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the directory.
	Directory *QueryWorksByOrganizationResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// Test directory
	//
	// example:
	//
	// 1496651577000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Test Workspace
	//
	// example:
	//
	// 1572334870000
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// Description
	//
	// example:
	//
	// The name of the work.
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// Security policies for collaborative authorization of works. Valid values:
	//
	// 	- 0: private
	//
	// 	- 12: Authorize specified members
	//
	// 	- 1 or 11: Authorize all workspace members
	//
	// >
	//
	// 	- If you use legacy permissions, the return value is 1.
	//
	// 	- If you use the new permissions, the return value is 11.
	//
	// example:
	//
	// Remarks on the work.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Alibaba Cloud account name of the person who modified the work.
	//
	// example:
	//
	// Tom
	OwnerName         *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	PublicFlag        *bool   `json:"PublicFlag,omitempty" xml:"PublicFlag,omitempty"`
	PublicInvalidTime *int64  `json:"PublicInvalidTime,omitempty" xml:"PublicInvalidTime,omitempty"`
	// The directory to which the work belongs.
	//
	// example:
	//
	// 1
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Li Si
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Test directory
	//
	// example:
	//
	// The name of the directory.
	WorkName *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The user ID of the work owner in the Quick BI.
	//
	// example:
	//
	// 897ce25e-****-****-af84-d13c5610****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// Test report
	//
	// example:
	//
	// The timestamp of the creation of the work in milliseconds.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// The name of the Alibaba Cloud account that modified the work.
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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

func (s *QueryWorksByOrganizationResponseBodyResultData) SetPublicFlag(v bool) *QueryWorksByOrganizationResponseBodyResultData {
	s.PublicFlag = &v
	return s
}

func (s *QueryWorksByOrganizationResponseBodyResultData) SetPublicInvalidTime(v int64) *QueryWorksByOrganizationResponseBodyResultData {
	s.PublicInvalidTime = &v
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
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorksByOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The page number of the returned page.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the work. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 1: published
	//
	// 	- 2: modified but not published
	//
	// 	- 3: unpublished
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Third-party embedding status. Valid values:
	//
	// 	- 0: The embed service is not enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 0
	ThirdPartAuthFlag *int32 `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The type of the work. Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- FULLPAGE: full-screen dashboards
	//
	// 	- REPORT: workbook
	//
	// example:
	//
	// PAGE
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns a list of all works in the organization workspace that meet the requested criteria.
	Result *QueryWorksByWorkspaceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The details of the list of works.
	Data []*QueryWorksByWorkspaceResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page set when the interface is requested.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of rows in the table.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// Third-party embedding status. Valid values:
	//
	// 	- 0: The embed service is not enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 1
	Auth3rdFlag *int32 `json:"Auth3rdFlag,omitempty" xml:"Auth3rdFlag,omitempty"`
	// Remarks on the work.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory to which the work belongs.
	Directory *QueryWorksByWorkspaceResponseBodyResultDataDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The timestamp of the creation of the work in milliseconds.
	//
	// example:
	//
	// 1496651577000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp of the modification of the work in milliseconds.
	//
	// example:
	//
	// 1572334870000
	GmtModify *string `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// Nickname of the work modifier.
	//
	// example:
	//
	// Tom
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// The user ID of the work owner in the Quick BI.
	//
	// example:
	//
	// The name of the workspace to which the work belongs.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The nickname of the work owner.
	//
	// example:
	//
	// Li Si
	OwnerName         *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	PublicFlag        *bool   `json:"PublicFlag,omitempty" xml:"PublicFlag,omitempty"`
	PublicInvalidTime *int64  `json:"PublicInvalidTime,omitempty" xml:"PublicInvalidTime,omitempty"`
	// Security policies for collaborative authorization of works. Valid values:
	//
	// 	- 0: private
	//
	// 	- 12: Authorize specified members
	//
	// 	- 1 or 11: Authorize all workspace members
	//
	// >
	//
	// 	- If you use legacy permissions, the return value is 1.
	//
	// 	- If you use the new permissions, the return value is 11.
	//
	// example:
	//
	// 0
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Status of dashboards, full-screen dashboards, spreadsheets. The default value of other work types is 1. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 1: published
	//
	// 	- 2: modified but not published
	//
	// 	- 3: unpublished
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the work.
	//
	// example:
	//
	// Test report
	WorkName *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
	// The type of the work. Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- FULLPAGE: full-screen dashboards
	//
	// 	- REPORT: workbook
	//
	// 	- dashboardOfflineQuery: self-service data retrieval
	//
	// 	- Analysis: Ad hoc analysis
	//
	// 	- DATAFORM: form filling
	//
	// example:
	//
	// PAGE
	WorkType *string `json:"WorkType,omitempty" xml:"WorkType,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 897ce25e-f993-4abd-af84-d13c5610****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// 87c6b145-090c-43e1-9426-8f93be23****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// Test Workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetPublicFlag(v bool) *QueryWorksByWorkspaceResponseBodyResultData {
	s.PublicFlag = &v
	return s
}

func (s *QueryWorksByWorkspaceResponseBodyResultData) SetPublicInvalidTime(v int64) *QueryWorksByWorkspaceResponseBodyResultData {
	s.PublicInvalidTime = &v
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
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The hierarchical structure of the directory ID to which the directory belongs. Separate the hierarchical structure with a /.
	//
	// example:
	//
	// 83d37ba6-d909-48a2-a517-f4d05c3a****
	PathId *string `json:"PathId,omitempty" xml:"PathId,omitempty"`
	// The hierarchical structure of the directory to which the directory belongs. Separate the hierarchical structure with a (/).
	//
	// example:
	//
	// Test directory
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorksByWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryWorkspaceRoleConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s QueryWorkspaceRoleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceRoleConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigRequest) SetRoleId(v int64) *QueryWorkspaceRoleConfigRequest {
	s.RoleId = &v
	return s
}

type QueryWorkspaceRoleConfigResponseBody struct {
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryWorkspaceRoleConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWorkspaceRoleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceRoleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigResponseBody) SetRequestId(v string) *QueryWorkspaceRoleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBody) SetResult(v *QueryWorkspaceRoleConfigResponseBodyResult) *QueryWorkspaceRoleConfigResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBody) SetSuccess(v bool) *QueryWorkspaceRoleConfigResponseBody {
	s.Success = &v
	return s
}

type QueryWorkspaceRoleConfigResponseBodyResult struct {
	AuthConfigList []*QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList `json:"AuthConfigList,omitempty" xml:"AuthConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsSystemRole *bool `json:"IsSystemRole,omitempty" xml:"IsSystemRole,omitempty"`
	// example:
	//
	// 25
	RoleId   *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s QueryWorkspaceRoleConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceRoleConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) SetAuthConfigList(v []*QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) *QueryWorkspaceRoleConfigResponseBodyResult {
	s.AuthConfigList = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) SetIsSystemRole(v bool) *QueryWorkspaceRoleConfigResponseBodyResult {
	s.IsSystemRole = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) SetRoleId(v int64) *QueryWorkspaceRoleConfigResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) SetRoleName(v string) *QueryWorkspaceRoleConfigResponseBodyResult {
	s.RoleName = &v
	return s
}

type QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList struct {
	ActionAuthKeys []*string `json:"ActionAuthKeys,omitempty" xml:"ActionAuthKeys,omitempty" type:"Repeated"`
	// example:
	//
	// portal_create
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
}

func (s QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) SetActionAuthKeys(v []*string) *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList {
	s.ActionAuthKeys = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) SetAuthKey(v string) *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList {
	s.AuthKey = &v
	return s
}

type QueryWorkspaceRoleConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorkspaceRoleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWorkspaceRoleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWorkspaceRoleConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigResponse) SetHeaders(v map[string]*string) *QueryWorkspaceRoleConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponse) SetStatusCode(v int32) *QueryWorkspaceRoleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponse) SetBody(v *QueryWorkspaceRoleConfigResponseBody) *QueryWorkspaceRoleConfigResponse {
	s.Body = v
	return s
}

type QueryWorkspaceUserListRequest struct {
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *QueryWorkspaceUserListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Data []*QueryWorkspaceUserListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// example:
	//
	// 16020915****8429
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// pop****@aliyunid.test
	AccountName *string                                           `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	NickName    *string                                           `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Role        *QueryWorkspaceUserListResponseBodyResultDataRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// role_workspace_admin
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// example:
	//
	// 25
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorkspaceUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// c5ea0db8-****-****-9081-04bc0df4c6a3
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	HandleReason *string `json:"HandleReason,omitempty" xml:"HandleReason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResultCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The user ID of the collection. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121344444790****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// d23e84a1-82a0-4292-bfdb-521306c3****
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
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveFavoritesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-******-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NONE
	MissHitPolicy *string `json:"MissHitPolicy,omitempty" xml:"MissHitPolicy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ROW_LEVEL
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
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
	// example:
	//
	// B70E1FBD-E533-52F2-A7A1-E02B92F78DDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDataLevelPermissionExtraConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDataLevelPermissionRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// { "ruleType": "ROW_LEVEL", // The row-level permission type. "usersModel": { "userGroups": [ "0d5fb19b- ***-1248 fc27ca51", // The ID of the user group. "3d2c23d4-***-f6390f325c2d" ], "users": [ "4334 ***358", // Quick BI the UserID of the user. "Huang***3fa822" ] }, "cubeId": "7c7223ae-31d1-4d2f-b11f-3c744528014b" }
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ruleType":"ROW_LEVEL","usersModel":{"userGroups":["26edcb76-****-bdbab78267cb","187e6dd5-1611-4cf7-a034-1a93bd5fecf9"],"users":["4334***358","Huang***3fa822"]},"cubeId":"7c7223ae-****44528014b"}
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
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDataLevelPermissionWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the training dataset that you want to remove from the specified custom linguistic model.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsOpen *int32 `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ROW_LEVEL
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
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataLevelPermissionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// true
	ThirdPartAuthFlag *bool `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 897ce25e-f993-4abd-af84-d13c5610****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEmbeddedStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 040e6f79d33444838***83c7206c070
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTicketNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Indicates whether the organization administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AdminUser *bool `json:"AdminUser,omitempty" xml:"AdminUser,omitempty"`
	// Indicate whether the RAM user is a permission administrator. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AuthAdminUser *bool `json:"AuthAdminUser,omitempty" xml:"AuthAdminUser,omitempty"`
	IsDeleted     *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The nickname of the account.
	//
	// 	- Format check: The value can be up to 50 characters in length.
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// example:
	//
	// Xiao Zhang
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	RoleIds  *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The ID of the user to be updated. The user ID is the UserID of the Quick BI, not the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The role type of the organization member. Valid values:
	//
	// 	- 1 : developer
	//
	// 	- 2 : visitors
	//
	// 	- 3 : Analyst
	//
	// example:
	//
	// 1
	UserType *int32 `json:"UserType,omitempty" xml:"UserType,omitempty"`
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

func (s *UpdateUserRequest) SetIsDeleted(v bool) *UpdateUserRequest {
	s.IsDeleted = &v
	return s
}

func (s *UpdateUserRequest) SetNickName(v string) *UpdateUserRequest {
	s.NickName = &v
	return s
}

func (s *UpdateUserRequest) SetRoleIds(v string) *UpdateUserRequest {
	s.RoleIds = &v
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
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The description of the user group.
	//
	// 	- Format verification: Maximum length 255
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// example:
	//
	// Description
	UserGroupDescription *string `json:"UserGroupDescription,omitempty" xml:"UserGroupDescription,omitempty"`
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The name of the user group.
	//
	// 	- Format verification: Maximum length 255
	//
	// 	- Special format verification: Chinese and English digits_ \\ / | () ] [
	//
	// example:
	//
	// pop0001
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 4AEF8C5C-D5D2-55D3-BB2F-9D3AA1B6F4FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the interface is successfully executed. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// e82f6c6c0333431bad0225b2f85e****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// This parameter is required.
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// pop_001
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// This parameter is required.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// 46e5374665ba4b679ee22e2a29270
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserTagValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f5698bedeb384b1986afccd9e434****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceUserRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 136516262323****,124498444445****
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
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
	// example:
	//
	// 7AAB95D7-2E11-4FE2-94BC-858E4FC0C976
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateWorkspaceUsersRoleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// 0
	Failure       *int32                 `json:"Failure,omitempty" xml:"Failure,omitempty"`
	FailureDetail map[string]interface{} `json:"FailureDetail,omitempty" xml:"FailureDetail,omitempty"`
	// example:
	//
	// 2
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceUsersRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the user. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
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
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawAllUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// Summary:
//
// Add selected groups of people incrementally for a single row and column permission rule.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.\\n
//
// @param request - AddDataLevelPermissionRuleUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataLevelPermissionRuleUsersResponse
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

// Summary:
//
// Add selected groups of people incrementally for a single row and column permission rule.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.\\n
//
// @param request - AddDataLevelPermissionRuleUsersRequest
//
// @return AddDataLevelPermissionRuleUsersResponse
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

// Summary:
//
// 43342***435,1553a****41231
//
// Description:
//
// ROW_LEVEL
//
// @param request - AddDataLevelPermissionWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataLevelPermissionWhiteListResponse
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

// Summary:
//
// 43342***435,1553a****41231
//
// Description:
//
// ROW_LEVEL
//
// @param request - AddDataLevelPermissionWhiteListRequest
//
// @return AddDataLevelPermissionWhiteListResponse
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

// Summary:
//
// Add a sharing configuration for data works.
//
// @param request - AddShareReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShareReportResponse
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

// Summary:
//
// Add a sharing configuration for data works.
//
// @param request - AddShareReportRequest
//
// @return AddShareReportResponse
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

// Summary:
//
// auditing
//
// @param request - AddUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserResponse
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleIds)) {
		body["RoleIds"] = request.RoleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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

// Summary:
//
// auditing
//
// @param request - AddUserRequest
//
// @return AddUserResponse
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

// Summary:
//
// The ID of the request.
//
// @param request - AddUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMemberResponse
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

// Summary:
//
// The ID of the request.
//
// @param request - AddUserGroupMemberRequest
//
// @return AddUserGroupMemberResponse
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

// Summary:
//
// Add users to a specified user group at a time.
//
// @param request - AddUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMembersResponse
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

// Summary:
//
// Add users to a specified user group at a time.
//
// @param request - AddUserGroupMembersRequest
//
// @return AddUserGroupMembersResponse
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

// Summary:
//
// Add the metadata of an organization member tag.
//
// @param request - AddUserTagMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserTagMetaResponse
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

// Summary:
//
// Add the metadata of an organization member tag.
//
// @param request - AddUserTagMetaRequest
//
// @return AddUserTagMetaResponse
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

// Summary:
//
// 添加成员到指定工作空间。
//
// @param request - AddUserToWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToWorkspaceResponse
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

// Summary:
//
// 添加成员到指定工作空间。
//
// @param request - AddUserToWorkspaceRequest
//
// @return AddUserToWorkspaceResponse
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

// Summary:
//
// 批量添加成员到工作空间。
//
// @param request - AddWorkspaceUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWorkspaceUsersResponse
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

// Summary:
//
// 批量添加成员到工作空间。
//
// @param request - AddWorkspaceUsersRequest
//
// @return AddWorkspaceUsersResponse
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

// Summary:
//
// 触发数据集抽取加速。
//
// @param request - AllotDatasetAccelerationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllotDatasetAccelerationTaskResponse
func (client *Client) AllotDatasetAccelerationTaskWithOptions(request *AllotDatasetAccelerationTaskRequest, runtime *util.RuntimeOptions) (_result *AllotDatasetAccelerationTaskResponse, _err error) {
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
		Action:      tea.String("AllotDatasetAccelerationTask"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllotDatasetAccelerationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发数据集抽取加速。
//
// @param request - AllotDatasetAccelerationTaskRequest
//
// @return AllotDatasetAccelerationTaskResponse
func (client *Client) AllotDatasetAccelerationTask(request *AllotDatasetAccelerationTaskRequest) (_result *AllotDatasetAccelerationTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllotDatasetAccelerationTaskResponse{}
	_body, _err := client.AllotDatasetAccelerationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch authorization of BI portal menu will be skipped automatically.
//
// @param request - AuthorizeMenuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeMenuResponse
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

// Summary:
//
// Batch authorization of BI portal menu will be skipped automatically.
//
// @param request - AuthorizeMenuRequest
//
// @return AuthorizeMenuResponse
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

// Summary:
//
// 批量添加飞书用户。
//
// @param request - BatchAddFeishuUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddFeishuUsersResponse
func (client *Client) BatchAddFeishuUsersWithOptions(request *BatchAddFeishuUsersRequest, runtime *util.RuntimeOptions) (_result *BatchAddFeishuUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeishuUsers)) {
		query["FeishuUsers"] = request.FeishuUsers
	}

	if !tea.BoolValue(util.IsUnset(request.IsAdmin)) {
		query["IsAdmin"] = request.IsAdmin
	}

	if !tea.BoolValue(util.IsUnset(request.IsAuthAdmin)) {
		query["IsAuthAdmin"] = request.IsAuthAdmin
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAddFeishuUsers"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddFeishuUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量添加飞书用户。
//
// @param request - BatchAddFeishuUsersRequest
//
// @return BatchAddFeishuUsersResponse
func (client *Client) BatchAddFeishuUsers(request *BatchAddFeishuUsersRequest) (_result *BatchAddFeishuUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAddFeishuUsersResponse{}
	_body, _err := client.BatchAddFeishuUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据门户菜单ID，取消指定用户、用户组的授权记录。
//
// @param request - CancelAuthorizationMenuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAuthorizationMenuResponse
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

// Summary:
//
// 根据门户菜单ID，取消指定用户、用户组的授权记录。
//
// @param request - CancelAuthorizationMenuRequest
//
// @return CancelAuthorizationMenuResponse
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

// Summary:
//
// Cancel the data works from the user\\"s collection.
//
// @param request - CancelCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCollectionResponse
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

// Summary:
//
// Cancel the data works from the user\\"s collection.
//
// @param request - CancelCollectionRequest
//
// @return CancelCollectionResponse
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

// Summary:
//
// Delete a share authorization for a data work.
//
// @param request - CancelReportShareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelReportShareResponse
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

// Summary:
//
// Delete a share authorization for a data work.
//
// @param request - CancelReportShareRequest
//
// @return CancelReportShareResponse
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

// Summary:
//
// The ID of the request.
//
// @param request - ChangeVisibilityModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeVisibilityModelResponse
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

// Summary:
//
// The ID of the request.
//
// @param request - ChangeVisibilityModelRequest
//
// @return ChangeVisibilityModelResponse
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

// Summary:
//
// Queries whether a user has permissions to view data works, such as dashboards and workbooks.
//
// @param request - CheckReadableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckReadableResponse
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

// Summary:
//
// Queries whether a user has permissions to view data works, such as dashboards and workbooks.
//
// @param request - CheckReadableRequest
//
// @return CheckReadableResponse
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

// Summary:
//
// 生成三方嵌入的ticket。
//
// @param request - CreateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
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

// Summary:
//
// 生成三方嵌入的ticket。
//
// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
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

// Summary:
//
// 生成智能小Q嵌入ticket。
//
// @param request - CreateTicket4CopilotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicket4CopilotResponse
func (client *Client) CreateTicket4CopilotWithOptions(request *CreateTicket4CopilotRequest, runtime *util.RuntimeOptions) (_result *CreateTicket4CopilotResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.CopilotId)) {
		query["CopilotId"] = request.CopilotId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.TicketNum)) {
		query["TicketNum"] = request.TicketNum
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTicket4Copilot"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicket4CopilotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成智能小Q嵌入ticket。
//
// @param request - CreateTicket4CopilotRequest
//
// @return CreateTicket4CopilotResponse
func (client *Client) CreateTicket4Copilot(request *CreateTicket4CopilotRequest) (_result *CreateTicket4CopilotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTicket4CopilotResponse{}
	_body, _err := client.CreateTicket4CopilotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a user group. You can specify a parent user group.
//
// @param request - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
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

// Summary:
//
// Create a user group. You can specify a parent user group.
//
// @param request - CreateUserGroupRequest
//
// @return CreateUserGroupResponse
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

// Summary:
//
// 查询引用指定数据集下的作品信息。
//
// @param request - DataSetBloodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataSetBloodResponse
func (client *Client) DataSetBloodWithOptions(request *DataSetBloodRequest, runtime *util.RuntimeOptions) (_result *DataSetBloodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSetIds)) {
		query["DataSetIds"] = request.DataSetIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WorksType)) {
		query["WorksType"] = request.WorksType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DataSetBlood"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DataSetBloodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询引用指定数据集下的作品信息。
//
// @param request - DataSetBloodRequest
//
// @return DataSetBloodResponse
func (client *Client) DataSetBlood(request *DataSetBloodRequest) (_result *DataSetBloodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DataSetBloodResponse{}
	_body, _err := client.DataSetBloodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询引用指定数据源下的数据集信息。
//
// @param request - DataSourceBloodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataSourceBloodResponse
func (client *Client) DataSourceBloodWithOptions(request *DataSourceBloodRequest, runtime *util.RuntimeOptions) (_result *DataSourceBloodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["DataSourceId"] = request.DataSourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DataSourceBlood"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DataSourceBloodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询引用指定数据源下的数据集信息。
//
// @param request - DataSourceBloodRequest
//
// @return DataSourceBloodResponse
func (client *Client) DataSourceBlood(request *DataSourceBloodRequest) (_result *DataSourceBloodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DataSourceBloodResponse{}
	_body, _err := client.DataSourceBloodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the expiration time of the ticket embedded in the report.
//
// @param request - DelayTicketExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelayTicketExpireTimeResponse
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

// Summary:
//
// Update the expiration time of the ticket embedded in the report.
//
// @param request - DelayTicketExpireTimeRequest
//
// @return DelayTicketExpireTimeResponse
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

// Summary:
//
// { "ruleId": "a5bb24da- ***-a891683e14da", // The ID of the row-column permission rule. "cubeId": "7c7223ae- ***-3c744528014b", // The ID of the dataset. "delModel": { "userGroups": [ "0d5fb19b- ***-1248 fc27ca51", // Delete the user group ID of the user group. "3d2c23d4-***-f6390f325c2d" ], "users": [ "4334 ***358", // Delete the UserID of the user group. "Huang***3fa822" ] } }
//
// Description:
//
// {"ruleId":"a5bb24da-***-a891683e14da","cubeId":"7c7223ae-***-3c744528014b","delModel":{"userGroups":["0d5fb19b-***-1248fc27ca51","3d2c23d4-***-f6390f325c2d"],"users":["4334***358","Huang***3fa822"]}}
//
// @param request - DeleteDataLevelPermissionRuleUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLevelPermissionRuleUsersResponse
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

// Summary:
//
// { "ruleId": "a5bb24da- ***-a891683e14da", // The ID of the row-column permission rule. "cubeId": "7c7223ae- ***-3c744528014b", // The ID of the dataset. "delModel": { "userGroups": [ "0d5fb19b- ***-1248 fc27ca51", // Delete the user group ID of the user group. "3d2c23d4-***-f6390f325c2d" ], "users": [ "4334 ***358", // Delete the UserID of the user group. "Huang***3fa822" ] } }
//
// Description:
//
// {"ruleId":"a5bb24da-***-a891683e14da","cubeId":"7c7223ae-***-3c744528014b","delModel":{"userGroups":["0d5fb19b-***-1248fc27ca51","3d2c23d4-***-f6390f325c2d"],"users":["4334***358","Huang***3fa822"]}}
//
// @param request - DeleteDataLevelPermissionRuleUsersRequest
//
// @return DeleteDataLevelPermissionRuleUsersResponse
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

// Summary:
//
// The ID of the request.
//
// Description:
//
// The ID of the training dataset that you want to remove from the specified custom linguistic model.
//
// @param request - DeleteDataLevelRuleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLevelRuleConfigResponse
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

// Summary:
//
// The ID of the request.
//
// Description:
//
// The ID of the training dataset that you want to remove from the specified custom linguistic model.
//
// @param request - DeleteDataLevelRuleConfigRequest
//
// @return DeleteDataLevelRuleConfigResponse
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

// Summary:
//
// auditing
//
// @param request - DeleteTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTicketResponse
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

// Summary:
//
// auditing
//
// @param request - DeleteTicketRequest
//
// @return DeleteTicketResponse
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

// Summary:
//
// auditing
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
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

// Summary:
//
// auditing
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
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

// Summary:
//
// 删除指定工作空间的成员。
//
// @param request - DeleteUserFromWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserFromWorkspaceResponse
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

// Summary:
//
// 删除指定工作空间的成员。
//
// @param request - DeleteUserFromWorkspaceRequest
//
// @return DeleteUserFromWorkspaceResponse
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

// Summary:
//
// Deletes a user group in an organization.
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
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

// Summary:
//
// Deletes a user group in an organization.
//
// @param request - DeleteUserGroupRequest
//
// @return DeleteUserGroupResponse
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

// Summary:
//
// Deletes a specified member from a specified user group.
//
// @param request - DeleteUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupMemberResponse
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

// Summary:
//
// Deletes a specified member from a specified user group.
//
// @param request - DeleteUserGroupMemberRequest
//
// @return DeleteUserGroupMemberResponse
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

// Summary:
//
// Indicates whether the request is successful. Valid values:
//
// 	- true: The request was successful.
//
// 	- false: The request failed.
//
// @param request - DeleteUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupMembersResponse
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

// Summary:
//
// Indicates whether the request is successful. Valid values:
//
// 	- true: The request was successful.
//
// 	- false: The request failed.
//
// @param request - DeleteUserGroupMembersRequest
//
// @return DeleteUserGroupMembersResponse
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

// Summary:
//
// Deletes the tag metadata of an organization member.
//
// @param request - DeleteUserTagMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserTagMetaResponse
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

// Summary:
//
// Deletes the tag metadata of an organization member.
//
// @param request - DeleteUserTagMetaRequest
//
// @return DeleteUserTagMetaResponse
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

// Summary:
//
// Test description
//
// @param request - GetUserGroupInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupInfoResponse
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

// Summary:
//
// Test description
//
// @param request - GetUserGroupInfoRequest
//
// @return GetUserGroupInfoResponse
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

// Summary:
//
// 查询API数据源列表。
//
// @param request - ListApiDatasourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiDatasourceResponse
func (client *Client) ListApiDatasourceWithOptions(request *ListApiDatasourceRequest, runtime *util.RuntimeOptions) (_result *ListApiDatasourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
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
		Action:      tea.String("ListApiDatasource"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApiDatasourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询API数据源列表。
//
// @param request - ListApiDatasourceRequest
//
// @return ListApiDatasourceResponse
func (client *Client) ListApiDatasource(request *ListApiDatasourceRequest) (_result *ListApiDatasourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApiDatasourceResponse{}
	_body, _err := client.ListApiDatasourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user group information at a time by user group ID.
//
// @param request - ListByUserGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListByUserGroupIdResponse
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

// Summary:
//
// Queries user group information at a time by user group ID.
//
// @param request - ListByUserGroupIdRequest
//
// @return ListByUserGroupIdResponse
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

// Summary:
//
// The ID of the work.
//
// @param request - ListCollectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCollectionsResponse
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

// Summary:
//
// The ID of the work.
//
// @param request - ListCollectionsRequest
//
// @return ListCollectionsResponse
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

// Summary:
//
// You can this operation to obtain a list of row and column permission configurations for a specified dataset.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.
//
// @param request - ListCubeDataLevelPermissionConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCubeDataLevelPermissionConfigResponse
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

// Summary:
//
// You can this operation to obtain a list of row and column permission configurations for a specified dataset.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.
//
// @param request - ListCubeDataLevelPermissionConfigRequest
//
// @return ListCubeDataLevelPermissionConfigResponse
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

// Summary:
//
// 根据行列权限种类，获取数据集行列权限的白名单列表。
//
// @param request - ListDataLevelPermissionWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLevelPermissionWhiteListResponse
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

// Summary:
//
// 根据行列权限种类，获取数据集行列权限的白名单列表。
//
// @param request - ListDataLevelPermissionWhiteListRequest
//
// @return ListDataLevelPermissionWhiteListResponse
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

// Summary:
//
// 获取指定用户在首页看板中展示的收藏作品列表。
//
// @param request - ListFavoriteReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFavoriteReportsResponse
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

// Summary:
//
// 获取指定用户在首页看板中展示的收藏作品列表。
//
// @param request - ListFavoriteReportsRequest
//
// @return ListFavoriteReportsResponse
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

// Summary:
//
// 获取指定组织角色下的用户列表。
//
// @param request - ListOrganizationRoleUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationRoleUsersResponse
func (client *Client) ListOrganizationRoleUsersWithOptions(request *ListOrganizationRoleUsersRequest, runtime *util.RuntimeOptions) (_result *ListOrganizationRoleUsersResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["RoleId"] = request.RoleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationRoleUsers"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationRoleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定组织角色下的用户列表。
//
// @param request - ListOrganizationRoleUsersRequest
//
// @return ListOrganizationRoleUsersResponse
func (client *Client) ListOrganizationRoleUsers(request *ListOrganizationRoleUsersRequest) (_result *ListOrganizationRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrganizationRoleUsersResponse{}
	_body, _err := client.ListOrganizationRoleUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取组织级别自定义角色列表。
//
// @param request - ListOrganizationRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationRolesResponse
func (client *Client) ListOrganizationRolesWithOptions(runtime *util.RuntimeOptions) (_result *ListOrganizationRolesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationRoles"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取组织级别自定义角色列表。
//
// @return ListOrganizationRolesResponse
func (client *Client) ListOrganizationRoles() (_result *ListOrganizationRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrganizationRolesResponse{}
	_body, _err := client.ListOrganizationRolesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list of authorization details for a BI portal menu.
//
// @param request - ListPortalMenuAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPortalMenuAuthorizationResponse
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

// Summary:
//
// Obtains the list of authorization details for a BI portal menu.
//
// @param request - ListPortalMenuAuthorizationRequest
//
// @return ListPortalMenuAuthorizationResponse
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

// Summary:
//
// Gets a hierarchical list of menus under a specific BI portal.
//
// @param request - ListPortalMenusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPortalMenusResponse
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

// Summary:
//
// Gets a hierarchical list of menus under a specific BI portal.
//
// @param request - ListPortalMenusRequest
//
// @return ListPortalMenusResponse
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

// Summary:
//
// 获取首页看板常看和足迹列表。
//
// @param request - ListRecentViewReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentViewReportsResponse
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

// Summary:
//
// 获取首页看板常看和足迹列表。
//
// @param request - ListRecentViewReportsRequest
//
// @return ListRecentViewReportsResponse
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

// Summary:
//
// 获取指定用户在首页看板中展示的被授权的作品列表。
//
// @param request - ListSharedReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSharedReportsResponse
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

// Summary:
//
// 获取指定用户在首页看板中展示的被授权的作品列表。
//
// @param request - ListSharedReportsRequest
//
// @return ListSharedReportsResponse
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

// Summary:
//
// Description
//
// @param request - ListUserGroupsByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsByUserIdResponse
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

// Summary:
//
// Description
//
// @param request - ListUserGroupsByUserIdRequest
//
// @return ListUserGroupsByUserIdResponse
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

// Summary:
//
// 获取指定空间角色下的用户列表。
//
// @param request - ListWorkspaceRoleUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceRoleUsersResponse
func (client *Client) ListWorkspaceRoleUsersWithOptions(request *ListWorkspaceRoleUsersRequest, runtime *util.RuntimeOptions) (_result *ListWorkspaceRoleUsersResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["RoleId"] = request.RoleId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaceRoleUsers"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspaceRoleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定空间角色下的用户列表。
//
// @param request - ListWorkspaceRoleUsersRequest
//
// @return ListWorkspaceRoleUsersResponse
func (client *Client) ListWorkspaceRoleUsers(request *ListWorkspaceRoleUsersRequest) (_result *ListWorkspaceRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkspaceRoleUsersResponse{}
	_body, _err := client.ListWorkspaceRoleUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取空间角色列表。
//
// @param request - ListWorkspaceRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceRolesResponse
func (client *Client) ListWorkspaceRolesWithOptions(request *ListWorkspaceRolesRequest, runtime *util.RuntimeOptions) (_result *ListWorkspaceRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaceRoles"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspaceRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取空间角色列表。
//
// @param request - ListWorkspaceRolesRequest
//
// @return ListWorkspaceRolesResponse
func (client *Client) ListWorkspaceRoles(request *ListWorkspaceRolesRequest) (_result *ListWorkspaceRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkspaceRolesResponse{}
	_body, _err := client.ListWorkspaceRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改指定API数据源参数值。
//
// @param request - ModifyApiDatasourceParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiDatasourceParametersResponse
func (client *Client) ModifyApiDatasourceParametersWithOptions(request *ModifyApiDatasourceParametersRequest, runtime *util.RuntimeOptions) (_result *ModifyApiDatasourceParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApiDatasourceParameters"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApiDatasourceParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改指定API数据源参数值。
//
// @param request - ModifyApiDatasourceParametersRequest
//
// @return ModifyApiDatasourceParametersResponse
func (client *Client) ModifyApiDatasourceParameters(request *ModifyApiDatasourceParametersRequest) (_result *ModifyApiDatasourceParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApiDatasourceParametersResponse{}
	_body, _err := client.ModifyApiDatasourceParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改智能问数嵌入配置。
//
// @param request - ModifyCopilotEmbedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCopilotEmbedConfigResponse
func (client *Client) ModifyCopilotEmbedConfigWithOptions(request *ModifyCopilotEmbedConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyCopilotEmbedConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentName)) {
		query["AgentName"] = request.AgentName
	}

	if !tea.BoolValue(util.IsUnset(request.CopilotId)) {
		query["CopilotId"] = request.CopilotId
	}

	if !tea.BoolValue(util.IsUnset(request.DataRange)) {
		query["DataRange"] = request.DataRange
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCopilotEmbedConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCopilotEmbedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改智能问数嵌入配置。
//
// @param request - ModifyCopilotEmbedConfigRequest
//
// @return ModifyCopilotEmbedConfigResponse
func (client *Client) ModifyCopilotEmbedConfig(request *ModifyCopilotEmbedConfigRequest) (_result *ModifyCopilotEmbedConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCopilotEmbedConfigResponse{}
	_body, _err := client.ModifyCopilotEmbedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据审批人获取相应的审批流信息。
//
// @param request - QueryApprovalInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApprovalInfoResponse
func (client *Client) QueryApprovalInfoWithOptions(request *QueryApprovalInfoRequest, runtime *util.RuntimeOptions) (_result *QueryApprovalInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryApprovalInfo"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryApprovalInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据审批人获取相应的审批流信息。
//
// @param request - QueryApprovalInfoRequest
//
// @return QueryApprovalInfoResponse
func (client *Client) QueryApprovalInfo(request *QueryApprovalInfoRequest) (_result *QueryApprovalInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryApprovalInfoResponse{}
	_body, _err := client.QueryApprovalInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询审计日志信息。
//
// @param request - QueryAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAuditLogResponse
func (client *Client) QueryAuditLogWithOptions(request *QueryAuditLogRequest, runtime *util.RuntimeOptions) (_result *QueryAuditLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorTypes)) {
		query["OperatorTypes"] = request.OperatorTypes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAuditLog"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审计日志信息。
//
// @param request - QueryAuditLogRequest
//
// @return QueryAuditLogResponse
func (client *Client) QueryAuditLog(request *QueryAuditLogRequest) (_result *QueryAuditLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAuditLogResponse{}
	_body, _err := client.QueryAuditLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询组件性能列表。
//
// @param request - QueryComponentPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryComponentPerformanceResponse
func (client *Client) QueryComponentPerformanceWithOptions(request *QueryComponentPerformanceRequest, runtime *util.RuntimeOptions) (_result *QueryComponentPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CostTimeAvgMin)) {
		query["CostTimeAvgMin"] = request.CostTimeAvgMin
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryComponentPerformance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryComponentPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组件性能列表。
//
// @param request - QueryComponentPerformanceRequest
//
// @return QueryComponentPerformanceResponse
func (client *Client) QueryComponentPerformance(request *QueryComponentPerformanceRequest) (_result *QueryComponentPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryComponentPerformanceResponse{}
	_body, _err := client.QueryComponentPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取开通小Q嵌入的配置列表。
//
// @param request - QueryCopilotEmbedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCopilotEmbedConfigResponse
func (client *Client) QueryCopilotEmbedConfigWithOptions(request *QueryCopilotEmbedConfigRequest, runtime *util.RuntimeOptions) (_result *QueryCopilotEmbedConfigResponse, _err error) {
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
		Action:      tea.String("QueryCopilotEmbedConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCopilotEmbedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取开通小Q嵌入的配置列表。
//
// @param request - QueryCopilotEmbedConfigRequest
//
// @return QueryCopilotEmbedConfigResponse
func (client *Client) QueryCopilotEmbedConfig(request *QueryCopilotEmbedConfigRequest) (_result *QueryCopilotEmbedConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCopilotEmbedConfigResponse{}
	_body, _err := client.QueryCopilotEmbedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据集优化建议。
//
// @param request - QueryCubeOptimizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCubeOptimizationResponse
func (client *Client) QueryCubeOptimizationWithOptions(request *QueryCubeOptimizationRequest, runtime *util.RuntimeOptions) (_result *QueryCubeOptimizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCubeOptimization"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCubeOptimizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据集优化建议。
//
// @param request - QueryCubeOptimizationRequest
//
// @return QueryCubeOptimizationResponse
func (client *Client) QueryCubeOptimization(request *QueryCubeOptimizationRequest) (_result *QueryCubeOptimizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCubeOptimizationResponse{}
	_body, _err := client.QueryCubeOptimizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据集性能。
//
// @param request - QueryCubePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCubePerformanceResponse
func (client *Client) QueryCubePerformanceWithOptions(request *QueryCubePerformanceRequest, runtime *util.RuntimeOptions) (_result *QueryCubePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CostTimeAvgMin)) {
		query["CostTimeAvgMin"] = request.CostTimeAvgMin
	}

	if !tea.BoolValue(util.IsUnset(request.CubeId)) {
		query["CubeId"] = request.CubeId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCubePerformance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCubePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据集性能。
//
// @param request - QueryCubePerformanceRequest
//
// @return QueryCubePerformanceResponse
func (client *Client) QueryCubePerformance(request *QueryCubePerformanceRequest) (_result *QueryCubePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCubePerformanceResponse{}
	_body, _err := client.QueryCubePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用开放数据服务API。
//
// @param request - QueryDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataResponse
func (client *Client) QueryDataWithOptions(request *QueryDataRequest, runtime *util.RuntimeOptions) (_result *QueryDataResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryData"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用开放数据服务API。
//
// @param request - QueryDataRequest
//
// @return QueryDataResponse
func (client *Client) QueryData(request *QueryDataRequest) (_result *QueryDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDataResponse{}
	_body, _err := client.QueryDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据范围目录列表。
//
// @param request - QueryDataRangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataRangeResponse
func (client *Client) QueryDataRangeWithOptions(request *QueryDataRangeRequest, runtime *util.RuntimeOptions) (_result *QueryDataRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDataRange"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDataRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据范围目录列表。
//
// @param request - QueryDataRangeRequest
//
// @return QueryDataRangeResponse
func (client *Client) QueryDataRange(request *QueryDataRangeRequest) (_result *QueryDataRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDataRangeResponse{}
	_body, _err := client.QueryDataRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The operator for the aggregate operation. Metric fields are available, such as SUM, AVG, and MAX.
//
// Description:
//
// f4cc43bc3***
//
// @param request - QueryDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataServiceResponse
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

// Summary:
//
// The operator for the aggregate operation. Metric fields are available, such as SUM, AVG, and MAX.
//
// Description:
//
// f4cc43bc3***
//
// @param request - QueryDataServiceRequest
//
// @return QueryDataServiceResponse
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

// Summary:
//
// 查询数据服务API列表。
//
// @param request - QueryDataServiceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataServiceListResponse
func (client *Client) QueryDataServiceListWithOptions(request *QueryDataServiceListRequest, runtime *util.RuntimeOptions) (_result *QueryDataServiceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
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
		Action:      tea.String("QueryDataServiceList"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDataServiceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据服务API列表。
//
// @param request - QueryDataServiceListRequest
//
// @return QueryDataServiceListResponse
func (client *Client) QueryDataServiceList(request *QueryDataServiceListRequest) (_result *QueryDataServiceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDataServiceListResponse{}
	_body, _err := client.QueryDataServiceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified dataset, including the data source, directory, and dataset model.
//
// Description:
//
// The data source, directory, and dataset model (including dimensions, measures, physical fields, custom SQL text, and association relationships).
//
// @param request - QueryDatasetDetailInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetDetailInfoResponse
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

// Summary:
//
// Queries the details of a specified dataset, including the data source, directory, and dataset model.
//
// Description:
//
// The data source, directory, and dataset model (including dimensions, measures, physical fields, custom SQL text, and association relationships).
//
// @param request - QueryDatasetDetailInfoRequest
//
// @return QueryDatasetDetailInfoResponse
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

// Summary:
//
// Indicates whether the table is a custom SQL table. Valid values:
//
// 	- true: custom SQL table
//
// 	- false: non-custom SQL table
//
// @param request - QueryDatasetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetInfoResponse
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

// Summary:
//
// Indicates whether the table is a custom SQL table. Valid values:
//
// 	- true: custom SQL table
//
// 	- false: non-custom SQL table
//
// @param request - QueryDatasetInfoRequest
//
// @return QueryDatasetInfoResponse
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

// Summary:
//
// The name of the training dataset.
//
// @param request - QueryDatasetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetListResponse
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

// Summary:
//
// The name of the training dataset.
//
// @param request - QueryDatasetListRequest
//
// @return QueryDatasetListResponse
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

// Summary:
//
// 获取指定数据集的行级权限开关状态。
//
// @param request - QueryDatasetSwitchInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetSwitchInfoResponse
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

// Summary:
//
// 获取指定数据集的行级权限开关状态。
//
// @param request - QueryDatasetSwitchInfoRequest
//
// @return QueryDatasetSwitchInfoResponse
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

// Summary:
//
// The number of reports that are currently embedded.
//
// @param request - QueryEmbeddedInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEmbeddedInfoResponse
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

// Summary:
//
// The number of reports that are currently embedded.
//
// @return QueryEmbeddedInfoResponse
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

// Summary:
//
// Queries whether embedding is enabled for a report.
//
// @param request - QueryEmbeddedStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEmbeddedStatusResponse
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

// Summary:
//
// Queries whether embedding is enabled for a report.
//
// @param request - QueryEmbeddedStatusRequest
//
// @return QueryEmbeddedStatusResponse
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

// Summary:
//
// 获取指定组织角色的配置信息
//
// @param request - QueryOrganizationRoleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrganizationRoleConfigResponse
func (client *Client) QueryOrganizationRoleConfigWithOptions(request *QueryOrganizationRoleConfigRequest, runtime *util.RuntimeOptions) (_result *QueryOrganizationRoleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["RoleId"] = request.RoleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrganizationRoleConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrganizationRoleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定组织角色的配置信息
//
// @param request - QueryOrganizationRoleConfigRequest
//
// @return QueryOrganizationRoleConfigResponse
func (client *Client) QueryOrganizationRoleConfig(request *QueryOrganizationRoleConfigRequest) (_result *QueryOrganizationRoleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrganizationRoleConfigResponse{}
	_body, _err := client.QueryOrganizationRoleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前组织下的工作空间列表信息。
//
// @param request - QueryOrganizationWorkspaceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrganizationWorkspaceListResponse
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

// Summary:
//
// 获取当前组织下的工作空间列表信息。
//
// @param request - QueryOrganizationWorkspaceListRequest
//
// @return QueryOrganizationWorkspaceListResponse
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

// Summary:
//
// The Alibaba Cloud account name of the owner.
//
// @param request - QueryReadableResourcesListByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReadableResourcesListByUserIdResponse
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

// Summary:
//
// The Alibaba Cloud account name of the owner.
//
// @param request - QueryReadableResourcesListByUserIdRequest
//
// @return QueryReadableResourcesListByUserIdResponse
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

// Summary:
//
// 查询报表性能列表
//
// @param request - QueryReportPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReportPerformanceResponse
func (client *Client) QueryReportPerformanceWithOptions(request *QueryReportPerformanceRequest, runtime *util.RuntimeOptions) (_result *QueryReportPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CostTimeAvgMin)) {
		query["CostTimeAvgMin"] = request.CostTimeAvgMin
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryReportPerformance"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryReportPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报表性能列表
//
// @param request - QueryReportPerformanceRequest
//
// @return QueryReportPerformanceResponse
func (client *Client) QueryReportPerformance(request *QueryReportPerformanceRequest) (_result *QueryReportPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryReportPerformanceResponse{}
	_body, _err := client.QueryReportPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Xiao Zhang
//
// @param request - QueryShareListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryShareListResponse
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

// Summary:
//
// Xiao Zhang
//
// @param request - QueryShareListRequest
//
// @return QueryShareListResponse
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

// Summary:
//
// You can call this operation to query the list of works authorized to a user.
//
// @param request - QuerySharesToUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySharesToUserListResponse
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

// Summary:
//
// You can call this operation to query the list of works authorized to a user.
//
// @param request - QuerySharesToUserListRequest
//
// @return QuerySharesToUserListResponse
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

// Summary:
//
// auditing
//
// @param request - QueryTicketInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTicketInfoResponse
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

// Summary:
//
// auditing
//
// @param request - QueryTicketInfoRequest
//
// @return QueryTicketInfoResponse
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

// Summary:
//
// You can this operation to obtain information about child user groups under a specified parent user group.
//
// @param request - QueryUserGroupListByParentIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserGroupListByParentIdResponse
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

// Summary:
//
// You can this operation to obtain information about child user groups under a specified parent user group.
//
// @param request - QueryUserGroupListByParentIdRequest
//
// @return QueryUserGroupListByParentIdResponse
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

// Summary:
//
// 获取用户组下的成员列表信息。
//
// @param request - QueryUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserGroupMemberResponse
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

// Summary:
//
// 获取用户组下的成员列表信息。
//
// @param request - QueryUserGroupMemberRequest
//
// @return QueryUserGroupMemberResponse
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

// Summary:
//
// Queries user information based on the Alibaba Cloud ID or Alibaba Cloud account name.
//
// @param request - QueryUserInfoByAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserInfoByAccountResponse
func (client *Client) QueryUserInfoByAccountWithOptions(request *QueryUserInfoByAccountRequest, runtime *util.RuntimeOptions) (_result *QueryUserInfoByAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.ParentAccountName)) {
		query["ParentAccountName"] = request.ParentAccountName
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

// Summary:
//
// Queries user information based on the Alibaba Cloud ID or Alibaba Cloud account name.
//
// @param request - QueryUserInfoByAccountRequest
//
// @return QueryUserInfoByAccountResponse
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

// Summary:
//
// Queries user information based on the user ID.
//
// @param request - QueryUserInfoByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserInfoByUserIdResponse
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

// Summary:
//
// Queries user information based on the user ID.
//
// @param request - QueryUserInfoByUserIdRequest
//
// @return QueryUserInfoByUserIdResponse
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

// Summary:
//
// 查询组织成员列表信息。
//
// @param request - QueryUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserListResponse
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

// Summary:
//
// 查询组织成员列表信息。
//
// @param request - QueryUserListRequest
//
// @return QueryUserListResponse
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

// Summary:
//
// 获取指定工作空间成员的预置空间角色信息。
//
// @param request - QueryUserRoleInfoInWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserRoleInfoInWorkspaceResponse
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

// Summary:
//
// 获取指定工作空间成员的预置空间角色信息。
//
// @param request - QueryUserRoleInfoInWorkspaceRequest
//
// @return QueryUserRoleInfoInWorkspaceResponse
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

// Summary:
//
// auditing
//
// @param request - QueryUserTagMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserTagMetaListResponse
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

// Summary:
//
// auditing
//
// @return QueryUserTagMetaListResponse
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

// Summary:
//
// Queries the list of tag values for a specific user.
//
// @param request - QueryUserTagValueListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserTagValueListResponse
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

// Summary:
//
// Queries the list of tag values for a specific user.
//
// @param request - QueryUserTagValueListRequest
//
// @return QueryUserTagValueListResponse
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

// Summary:
//
// 获取报表详情
//
// @param request - QueryWorksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksResponse
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

// Summary:
//
// 获取报表详情
//
// @param request - QueryWorksRequest
//
// @return QueryWorksResponse
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

// Summary:
//
// abcd****
//
// @param request - QueryWorksBloodRelationshipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksBloodRelationshipResponse
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

// Summary:
//
// abcd****
//
// @param request - QueryWorksBloodRelationshipRequest
//
// @return QueryWorksBloodRelationshipResponse
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

// Summary:
//
// The total number of rows in the table.
//
// @param request - QueryWorksByOrganizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksByOrganizationResponse
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

// Summary:
//
// The total number of rows in the table.
//
// @param request - QueryWorksByOrganizationRequest
//
// @return QueryWorksByOrganizationResponse
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

// Summary:
//
// The name of the directory.
//
// @param request - QueryWorksByWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksByWorkspaceResponse
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

// Summary:
//
// The name of the directory.
//
// @param request - QueryWorksByWorkspaceRequest
//
// @return QueryWorksByWorkspaceResponse
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

// Summary:
//
// 获取指定空间角色的配置信息
//
// @param request - QueryWorkspaceRoleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorkspaceRoleConfigResponse
func (client *Client) QueryWorkspaceRoleConfigWithOptions(request *QueryWorkspaceRoleConfigRequest, runtime *util.RuntimeOptions) (_result *QueryWorkspaceRoleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["RoleId"] = request.RoleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWorkspaceRoleConfig"),
		Version:     tea.String("2022-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWorkspaceRoleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定空间角色的配置信息
//
// @param request - QueryWorkspaceRoleConfigRequest
//
// @return QueryWorkspaceRoleConfigResponse
func (client *Client) QueryWorkspaceRoleConfig(request *QueryWorkspaceRoleConfigRequest) (_result *QueryWorkspaceRoleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWorkspaceRoleConfigResponse{}
	_body, _err := client.QueryWorkspaceRoleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定工作空间下的成员列表信息。
//
// @param request - QueryWorkspaceUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorkspaceUserListResponse
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

// Summary:
//
// 查询指定工作空间下的成员列表信息。
//
// @param request - QueryWorkspaceUserListRequest
//
// @return QueryWorkspaceUserListResponse
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

// Summary:
//
// 第三方资源审批回调接口
//
// @param request - ResultCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResultCallbackResponse
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

// Summary:
//
// 第三方资源审批回调接口
//
// @param request - ResultCallbackRequest
//
// @return ResultCallbackResponse
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

// Summary:
//
// Add the user\\"s collection data works.
//
// @param request - SaveFavoritesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveFavoritesResponse
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

// Summary:
//
// Add the user\\"s collection data works.
//
// @param request - SaveFavoritesRequest
//
// @return SaveFavoritesResponse
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

// Summary:
//
// 设置行列权限的额外配置
//
// @param request - SetDataLevelPermissionExtraConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionExtraConfigResponse
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

// Summary:
//
// 设置行列权限的额外配置
//
// @param request - SetDataLevelPermissionExtraConfigRequest
//
// @return SetDataLevelPermissionExtraConfigResponse
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

// Summary:
//
// 设置单条数据集行列权限配置信息（新增和更新）
//
// @param request - SetDataLevelPermissionRuleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionRuleConfigResponse
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

// Summary:
//
// 设置单条数据集行列权限配置信息（新增和更新）
//
// @param request - SetDataLevelPermissionRuleConfigRequest
//
// @return SetDataLevelPermissionRuleConfigResponse
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

// Summary:
//
// Sets the whitelist for the specified row-level permissions.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.
//
// @param request - SetDataLevelPermissionWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionWhiteListResponse
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

// Summary:
//
// Sets the whitelist for the specified row-level permissions.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.
//
// @param request - SetDataLevelPermissionWhiteListRequest
//
// @return SetDataLevelPermissionWhiteListResponse
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

// Summary:
//
// Indicates whether the request is successful. Valid values:
//
// 	- true: The request was successful.
//
// 	- false: The request failed.
//
// Description:
//
// The execution result of the interface. Valid values:
//
// 	- true: The request was successful.
//
// 	- false: The request failed.
//
// @param request - UpdateDataLevelPermissionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLevelPermissionStatusResponse
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

// Summary:
//
// Indicates whether the request is successful. Valid values:
//
// 	- true: The request was successful.
//
// 	- false: The request failed.
//
// Description:
//
// The execution result of the interface. Valid values:
//
// 	- true: The request was successful.
//
// 	- false: The request failed.
//
// @param request - UpdateDataLevelPermissionStatusRequest
//
// @return UpdateDataLevelPermissionStatusResponse
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

// Summary:
//
// The ID of the request.
//
// @param request - UpdateEmbeddedStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEmbeddedStatusResponse
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

// Summary:
//
// The ID of the request.
//
// @param request - UpdateEmbeddedStatusRequest
//
// @return UpdateEmbeddedStatusResponse
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

// Summary:
//
// 更新三方嵌入ticket的票据数量
//
// @param request - UpdateTicketNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketNumResponse
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

// Summary:
//
// 更新三方嵌入ticket的票据数量
//
// @param request - UpdateTicketNumRequest
//
// @return UpdateTicketNumResponse
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

// Summary:
//
// Updates the information of a specified member in an organization.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
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

	if !tea.BoolValue(util.IsUnset(request.IsDeleted)) {
		query["IsDeleted"] = request.IsDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		query["NickName"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.RoleIds)) {
		query["RoleIds"] = request.RoleIds
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

// Summary:
//
// Updates the information of a specified member in an organization.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
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

// Summary:
//
// Updates information about a specified user group in an organization.
//
// @param request - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
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

// Summary:
//
// Updates information about a specified user group in an organization.
//
// @param request - UpdateUserGroupRequest
//
// @return UpdateUserGroupResponse
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

// Summary:
//
// 用于更新组织成员标签元信息
//
// @param request - UpdateUserTagMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserTagMetaResponse
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

// Summary:
//
// 用于更新组织成员标签元信息
//
// @param request - UpdateUserTagMetaRequest
//
// @return UpdateUserTagMetaResponse
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

// Summary:
//
// 更新组织成员标签值
//
// @param request - UpdateUserTagValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserTagValueResponse
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

// Summary:
//
// 更新组织成员标签值
//
// @param request - UpdateUserTagValueRequest
//
// @return UpdateUserTagValueResponse
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

// Summary:
//
// 修改工作空间下指定成员的角色，已有的角色会被覆盖
//
// @param request - UpdateWorkspaceUserRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceUserRoleResponse
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

// Summary:
//
// 修改工作空间下指定成员的角色，已有的角色会被覆盖
//
// @param request - UpdateWorkspaceUserRoleRequest
//
// @return UpdateWorkspaceUserRoleResponse
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

// Summary:
//
// 批量更新工作空间成员的角色信息，已有角色会被覆盖
//
// @param request - UpdateWorkspaceUsersRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceUsersRoleResponse
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

// Summary:
//
// 批量更新工作空间成员的角色信息，已有角色会被覆盖
//
// @param request - UpdateWorkspaceUsersRoleRequest
//
// @return UpdateWorkspaceUsersRoleResponse
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

// Summary:
//
// Make the user exit all user groups. This process is irreversible. Exercise caution when performing this operation.
//
// @param request - WithdrawAllUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawAllUserGroupsResponse
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

// Summary:
//
// Make the user exit all user groups. This process is irreversible. Exercise caution when performing this operation.
//
// @param request - WithdrawAllUserGroupsRequest
//
// @return WithdrawAllUserGroupsResponse
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
