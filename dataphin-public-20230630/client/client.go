// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataSourceInfo struct {
	Configs     map[string]*string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	CreateTime  *int64             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator     *string            `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorName *string            `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	Description *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Env         *string            `json:"Env,omitempty" xml:"Env,omitempty"`
	Id          *int64             `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifyTime  *int64             `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name        *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner       *string            `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerName   *string            `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	Scope       *string            `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TenantId    *int64             `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type        *string            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s DataSourceInfo) GoString() string {
	return s.String()
}

func (s *DataSourceInfo) SetConfigs(v map[string]*string) *DataSourceInfo {
	s.Configs = v
	return s
}

func (s *DataSourceInfo) SetCreateTime(v int64) *DataSourceInfo {
	s.CreateTime = &v
	return s
}

func (s *DataSourceInfo) SetCreator(v string) *DataSourceInfo {
	s.Creator = &v
	return s
}

func (s *DataSourceInfo) SetCreatorName(v string) *DataSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *DataSourceInfo) SetDescription(v string) *DataSourceInfo {
	s.Description = &v
	return s
}

func (s *DataSourceInfo) SetEnv(v string) *DataSourceInfo {
	s.Env = &v
	return s
}

func (s *DataSourceInfo) SetId(v int64) *DataSourceInfo {
	s.Id = &v
	return s
}

func (s *DataSourceInfo) SetModifyTime(v int64) *DataSourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *DataSourceInfo) SetName(v string) *DataSourceInfo {
	s.Name = &v
	return s
}

func (s *DataSourceInfo) SetOwner(v string) *DataSourceInfo {
	s.Owner = &v
	return s
}

func (s *DataSourceInfo) SetOwnerName(v string) *DataSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *DataSourceInfo) SetScope(v string) *DataSourceInfo {
	s.Scope = &v
	return s
}

func (s *DataSourceInfo) SetTenantId(v int64) *DataSourceInfo {
	s.TenantId = &v
	return s
}

func (s *DataSourceInfo) SetType(v string) *DataSourceInfo {
	s.Type = &v
	return s
}

type DatasourceCreate struct {
	CheckActivity *bool              `json:"CheckActivity,omitempty" xml:"CheckActivity,omitempty"`
	Configs       map[string]*string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	Description   *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Type          *string            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DatasourceCreate) String() string {
	return tea.Prettify(s)
}

func (s DatasourceCreate) GoString() string {
	return s.String()
}

func (s *DatasourceCreate) SetCheckActivity(v bool) *DatasourceCreate {
	s.CheckActivity = &v
	return s
}

func (s *DatasourceCreate) SetConfigs(v map[string]*string) *DatasourceCreate {
	s.Configs = v
	return s
}

func (s *DatasourceCreate) SetDescription(v string) *DatasourceCreate {
	s.Description = &v
	return s
}

func (s *DatasourceCreate) SetName(v string) *DatasourceCreate {
	s.Name = &v
	return s
}

func (s *DatasourceCreate) SetType(v string) *DatasourceCreate {
	s.Type = &v
	return s
}

type AddTenantMembersRequest struct {
	// This parameter is required.
	AddCommand *AddTenantMembersRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddTenantMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersRequest) GoString() string {
	return s.String()
}

func (s *AddTenantMembersRequest) SetAddCommand(v *AddTenantMembersRequestAddCommand) *AddTenantMembersRequest {
	s.AddCommand = v
	return s
}

func (s *AddTenantMembersRequest) SetOpTenantId(v int64) *AddTenantMembersRequest {
	s.OpTenantId = &v
	return s
}

type AddTenantMembersRequestAddCommand struct {
	// This parameter is required.
	UserList []*AddTenantMembersRequestAddCommandUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s AddTenantMembersRequestAddCommand) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddTenantMembersRequestAddCommand) SetUserList(v []*AddTenantMembersRequestAddCommandUserList) *AddTenantMembersRequestAddCommand {
	s.UserList = v
	return s
}

type AddTenantMembersRequestAddCommandUserList struct {
	// example:
	//
	// 1323241
	Id       *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	RoleList []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
}

func (s AddTenantMembersRequestAddCommandUserList) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersRequestAddCommandUserList) GoString() string {
	return s.String()
}

func (s *AddTenantMembersRequestAddCommandUserList) SetId(v string) *AddTenantMembersRequestAddCommandUserList {
	s.Id = &v
	return s
}

func (s *AddTenantMembersRequestAddCommandUserList) SetRoleList(v []*string) *AddTenantMembersRequestAddCommandUserList {
	s.RoleList = v
	return s
}

type AddTenantMembersShrinkRequest struct {
	// This parameter is required.
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddTenantMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTenantMembersShrinkRequest) SetAddCommandShrink(v string) *AddTenantMembersShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddTenantMembersShrinkRequest) SetOpTenantId(v int64) *AddTenantMembersShrinkRequest {
	s.OpTenantId = &v
	return s
}

type AddTenantMembersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTenantMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddTenantMembersResponseBody) SetCode(v string) *AddTenantMembersResponseBody {
	s.Code = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetData(v bool) *AddTenantMembersResponseBody {
	s.Data = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetHttpStatusCode(v int32) *AddTenantMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetMessage(v string) *AddTenantMembersResponseBody {
	s.Message = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetRequestId(v string) *AddTenantMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetSuccess(v bool) *AddTenantMembersResponseBody {
	s.Success = &v
	return s
}

type AddTenantMembersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTenantMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTenantMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersResponse) GoString() string {
	return s.String()
}

func (s *AddTenantMembersResponse) SetHeaders(v map[string]*string) *AddTenantMembersResponse {
	s.Headers = v
	return s
}

func (s *AddTenantMembersResponse) SetStatusCode(v int32) *AddTenantMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTenantMembersResponse) SetBody(v *AddTenantMembersResponseBody) *AddTenantMembersResponse {
	s.Body = v
	return s
}

type AddTenantMembersBySourceUserRequest struct {
	AddCommand *AddTenantMembersBySourceUserRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddTenantMembersBySourceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersBySourceUserRequest) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserRequest) SetAddCommand(v *AddTenantMembersBySourceUserRequestAddCommand) *AddTenantMembersBySourceUserRequest {
	s.AddCommand = v
	return s
}

func (s *AddTenantMembersBySourceUserRequest) SetOpTenantId(v int64) *AddTenantMembersBySourceUserRequest {
	s.OpTenantId = &v
	return s
}

type AddTenantMembersBySourceUserRequestAddCommand struct {
	SourceUserList []*AddTenantMembersBySourceUserRequestAddCommandSourceUserList `json:"SourceUserList,omitempty" xml:"SourceUserList,omitempty" type:"Repeated"`
}

func (s AddTenantMembersBySourceUserRequestAddCommand) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersBySourceUserRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserRequestAddCommand) SetSourceUserList(v []*AddTenantMembersBySourceUserRequestAddCommandSourceUserList) *AddTenantMembersBySourceUserRequestAddCommand {
	s.SourceUserList = v
	return s
}

type AddTenantMembersBySourceUserRequestAddCommandSourceUserList struct {
	// example:
	//
	// 123@xx.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 123@dingding
	DingNumber  *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123@xx.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 13888888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// example:
	//
	// 2323131
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s AddTenantMembersBySourceUserRequestAddCommandSourceUserList) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersBySourceUserRequestAddCommandSourceUserList) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetAccountName(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.AccountName = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetDingNumber(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.DingNumber = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetDisplayName(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.DisplayName = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetMail(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.Mail = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetMobilePhone(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.MobilePhone = &v
	return s
}

func (s *AddTenantMembersBySourceUserRequestAddCommandSourceUserList) SetSourceId(v string) *AddTenantMembersBySourceUserRequestAddCommandSourceUserList {
	s.SourceId = &v
	return s
}

type AddTenantMembersBySourceUserShrinkRequest struct {
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddTenantMembersBySourceUserShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersBySourceUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserShrinkRequest) SetAddCommandShrink(v string) *AddTenantMembersBySourceUserShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddTenantMembersBySourceUserShrinkRequest) SetOpTenantId(v int64) *AddTenantMembersBySourceUserShrinkRequest {
	s.OpTenantId = &v
	return s
}

type AddTenantMembersBySourceUserResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTenantMembersBySourceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersBySourceUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserResponseBody) SetCode(v string) *AddTenantMembersBySourceUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetData(v bool) *AddTenantMembersBySourceUserResponseBody {
	s.Data = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetHttpStatusCode(v int32) *AddTenantMembersBySourceUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetMessage(v string) *AddTenantMembersBySourceUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetRequestId(v string) *AddTenantMembersBySourceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetSuccess(v bool) *AddTenantMembersBySourceUserResponseBody {
	s.Success = &v
	return s
}

type AddTenantMembersBySourceUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTenantMembersBySourceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTenantMembersBySourceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTenantMembersBySourceUserResponse) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserResponse) SetHeaders(v map[string]*string) *AddTenantMembersBySourceUserResponse {
	s.Headers = v
	return s
}

func (s *AddTenantMembersBySourceUserResponse) SetStatusCode(v int32) *AddTenantMembersBySourceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponse) SetBody(v *AddTenantMembersBySourceUserResponseBody) *AddTenantMembersBySourceUserResponse {
	s.Body = v
	return s
}

type AddUserGroupMemberRequest struct {
	AddCommand *AddUserGroupMemberRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddUserGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberRequest) SetAddCommand(v *AddUserGroupMemberRequestAddCommand) *AddUserGroupMemberRequest {
	s.AddCommand = v
	return s
}

func (s *AddUserGroupMemberRequest) SetOpTenantId(v int64) *AddUserGroupMemberRequest {
	s.OpTenantId = &v
	return s
}

type AddUserGroupMemberRequestAddCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 132331
	UserGroupId *string   `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserIdList  []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s AddUserGroupMemberRequestAddCommand) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMemberRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberRequestAddCommand) SetUserGroupId(v string) *AddUserGroupMemberRequestAddCommand {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMemberRequestAddCommand) SetUserIdList(v []*string) *AddUserGroupMemberRequestAddCommand {
	s.UserIdList = v
	return s
}

type AddUserGroupMemberShrinkRequest struct {
	AddCommandShrink *string `json:"AddCommand,omitempty" xml:"AddCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddUserGroupMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserGroupMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberShrinkRequest) SetAddCommandShrink(v string) *AddUserGroupMemberShrinkRequest {
	s.AddCommandShrink = &v
	return s
}

func (s *AddUserGroupMemberShrinkRequest) SetOpTenantId(v int64) *AddUserGroupMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

type AddUserGroupMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *AddUserGroupMemberResponseBody) SetCode(v string) *AddUserGroupMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetData(v bool) *AddUserGroupMemberResponseBody {
	s.Data = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetHttpStatusCode(v int32) *AddUserGroupMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetMessage(v string) *AddUserGroupMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetRequestId(v string) *AddUserGroupMemberResponseBody {
	s.RequestId = &v
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

type CheckDataSourceConnectivityRequest struct {
	// This parameter is required.
	CheckCommand *CheckDataSourceConnectivityRequestCheckCommand `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckDataSourceConnectivityRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityRequest) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityRequest) SetCheckCommand(v *CheckDataSourceConnectivityRequestCheckCommand) *CheckDataSourceConnectivityRequest {
	s.CheckCommand = v
	return s
}

func (s *CheckDataSourceConnectivityRequest) SetOpTenantId(v int64) *CheckDataSourceConnectivityRequest {
	s.OpTenantId = &v
	return s
}

type CheckDataSourceConnectivityRequestCheckCommand struct {
	// This parameter is required.
	ConfigItemList []*CheckDataSourceConnectivityRequestCheckCommandConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckDataSourceConnectivityRequestCheckCommand) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityRequestCheckCommand) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityRequestCheckCommand) SetConfigItemList(v []*CheckDataSourceConnectivityRequestCheckCommandConfigItemList) *CheckDataSourceConnectivityRequestCheckCommand {
	s.ConfigItemList = v
	return s
}

func (s *CheckDataSourceConnectivityRequestCheckCommand) SetType(v string) *CheckDataSourceConnectivityRequestCheckCommand {
	s.Type = &v
	return s
}

type CheckDataSourceConnectivityRequestCheckCommandConfigItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CheckDataSourceConnectivityRequestCheckCommandConfigItemList) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityRequestCheckCommandConfigItemList) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityRequestCheckCommandConfigItemList) SetKey(v string) *CheckDataSourceConnectivityRequestCheckCommandConfigItemList {
	s.Key = &v
	return s
}

func (s *CheckDataSourceConnectivityRequestCheckCommandConfigItemList) SetValue(v string) *CheckDataSourceConnectivityRequestCheckCommandConfigItemList {
	s.Value = &v
	return s
}

type CheckDataSourceConnectivityShrinkRequest struct {
	// This parameter is required.
	CheckCommandShrink *string `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckDataSourceConnectivityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityShrinkRequest) SetCheckCommandShrink(v string) *CheckDataSourceConnectivityShrinkRequest {
	s.CheckCommandShrink = &v
	return s
}

func (s *CheckDataSourceConnectivityShrinkRequest) SetOpTenantId(v int64) *CheckDataSourceConnectivityShrinkRequest {
	s.OpTenantId = &v
	return s
}

type CheckDataSourceConnectivityResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDataSourceConnectivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityResponseBody) SetCode(v string) *CheckDataSourceConnectivityResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetData(v bool) *CheckDataSourceConnectivityResponseBody {
	s.Data = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetHttpStatusCode(v int32) *CheckDataSourceConnectivityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetMessage(v string) *CheckDataSourceConnectivityResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetRequestId(v string) *CheckDataSourceConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDataSourceConnectivityResponseBody) SetSuccess(v bool) *CheckDataSourceConnectivityResponseBody {
	s.Success = &v
	return s
}

type CheckDataSourceConnectivityResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDataSourceConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDataSourceConnectivityResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityResponse) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityResponse) SetHeaders(v map[string]*string) *CheckDataSourceConnectivityResponse {
	s.Headers = v
	return s
}

func (s *CheckDataSourceConnectivityResponse) SetStatusCode(v int32) *CheckDataSourceConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityResponse) SetBody(v *CheckDataSourceConnectivityResponseBody) *CheckDataSourceConnectivityResponse {
	s.Body = v
	return s
}

type CheckDataSourceConnectivityByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckDataSourceConnectivityByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityByIdRequest) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityByIdRequest) SetId(v int64) *CheckDataSourceConnectivityByIdRequest {
	s.Id = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdRequest) SetOpTenantId(v int64) *CheckDataSourceConnectivityByIdRequest {
	s.OpTenantId = &v
	return s
}

type CheckDataSourceConnectivityByIdResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDataSourceConnectivityByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityByIdResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetCode(v string) *CheckDataSourceConnectivityByIdResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetData(v bool) *CheckDataSourceConnectivityByIdResponseBody {
	s.Data = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetHttpStatusCode(v int32) *CheckDataSourceConnectivityByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetMessage(v string) *CheckDataSourceConnectivityByIdResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetRequestId(v string) *CheckDataSourceConnectivityByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponseBody) SetSuccess(v bool) *CheckDataSourceConnectivityByIdResponseBody {
	s.Success = &v
	return s
}

type CheckDataSourceConnectivityByIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDataSourceConnectivityByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDataSourceConnectivityByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDataSourceConnectivityByIdResponse) GoString() string {
	return s.String()
}

func (s *CheckDataSourceConnectivityByIdResponse) SetHeaders(v map[string]*string) *CheckDataSourceConnectivityByIdResponse {
	s.Headers = v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponse) SetStatusCode(v int32) *CheckDataSourceConnectivityByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDataSourceConnectivityByIdResponse) SetBody(v *CheckDataSourceConnectivityByIdResponseBody) *CheckDataSourceConnectivityByIdResponse {
	s.Body = v
	return s
}

type CheckResourcePermissionRequest struct {
	// This parameter is required.
	CheckCommand *CheckResourcePermissionRequestCheckCommand `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckResourcePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionRequest) SetCheckCommand(v *CheckResourcePermissionRequestCheckCommand) *CheckResourcePermissionRequest {
	s.CheckCommand = v
	return s
}

func (s *CheckResourcePermissionRequest) SetOpTenantId(v int64) *CheckResourcePermissionRequest {
	s.OpTenantId = &v
	return s
}

type CheckResourcePermissionRequestCheckCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// UPDATE
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// This parameter is required.
	ResourceList []*CheckResourcePermissionRequestCheckCommandResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PHYSICAL_TABLE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 323231
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CheckResourcePermissionRequestCheckCommand) String() string {
	return tea.Prettify(s)
}

func (s CheckResourcePermissionRequestCheckCommand) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionRequestCheckCommand) SetOperate(v string) *CheckResourcePermissionRequestCheckCommand {
	s.Operate = &v
	return s
}

func (s *CheckResourcePermissionRequestCheckCommand) SetResourceList(v []*CheckResourcePermissionRequestCheckCommandResourceList) *CheckResourcePermissionRequestCheckCommand {
	s.ResourceList = v
	return s
}

func (s *CheckResourcePermissionRequestCheckCommand) SetResourceType(v string) *CheckResourcePermissionRequestCheckCommand {
	s.ResourceType = &v
	return s
}

func (s *CheckResourcePermissionRequestCheckCommand) SetUserId(v string) *CheckResourcePermissionRequestCheckCommand {
	s.UserId = &v
	return s
}

type CheckResourcePermissionRequestCheckCommandResourceList struct {
	// This parameter is required.
	//
	// example:
	//
	// hadoop.300000806.data_distill.behavior_gameinfor_01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CheckResourcePermissionRequestCheckCommandResourceList) String() string {
	return tea.Prettify(s)
}

func (s CheckResourcePermissionRequestCheckCommandResourceList) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionRequestCheckCommandResourceList) SetResourceId(v string) *CheckResourcePermissionRequestCheckCommandResourceList {
	s.ResourceId = &v
	return s
}

type CheckResourcePermissionShrinkRequest struct {
	// This parameter is required.
	CheckCommandShrink *string `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckResourcePermissionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResourcePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionShrinkRequest) SetCheckCommandShrink(v string) *CheckResourcePermissionShrinkRequest {
	s.CheckCommandShrink = &v
	return s
}

func (s *CheckResourcePermissionShrinkRequest) SetOpTenantId(v int64) *CheckResourcePermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

type CheckResourcePermissionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePermissionList []*CheckResourcePermissionResponseBodyResourcePermissionList `json:"ResourcePermissionList,omitempty" xml:"ResourcePermissionList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckResourcePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckResourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionResponseBody) SetCode(v string) *CheckResourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetHttpStatusCode(v int32) *CheckResourcePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetMessage(v string) *CheckResourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetRequestId(v string) *CheckResourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetResourcePermissionList(v []*CheckResourcePermissionResponseBodyResourcePermissionList) *CheckResourcePermissionResponseBody {
	s.ResourcePermissionList = v
	return s
}

func (s *CheckResourcePermissionResponseBody) SetSuccess(v bool) *CheckResourcePermissionResponseBody {
	s.Success = &v
	return s
}

type CheckResourcePermissionResponseBodyResourcePermissionList struct {
	// example:
	//
	// true
	HasPermission *bool `json:"HasPermission,omitempty" xml:"HasPermission,omitempty"`
	// example:
	//
	// hadoop.300000806.data_distill.behavior_gameinfor_01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CheckResourcePermissionResponseBodyResourcePermissionList) String() string {
	return tea.Prettify(s)
}

func (s CheckResourcePermissionResponseBodyResourcePermissionList) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionResponseBodyResourcePermissionList) SetHasPermission(v bool) *CheckResourcePermissionResponseBodyResourcePermissionList {
	s.HasPermission = &v
	return s
}

func (s *CheckResourcePermissionResponseBodyResourcePermissionList) SetResourceId(v string) *CheckResourcePermissionResponseBodyResourcePermissionList {
	s.ResourceId = &v
	return s
}

type CheckResourcePermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResourcePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckResourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionResponse) SetHeaders(v map[string]*string) *CheckResourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *CheckResourcePermissionResponse) SetStatusCode(v int32) *CheckResourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResourcePermissionResponse) SetBody(v *CheckResourcePermissionResponseBody) *CheckResourcePermissionResponse {
	s.Body = v
	return s
}

type CreateAdHocFileRequest struct {
	// This parameter is required.
	CreateCommand *CreateAdHocFileRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateAdHocFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAdHocFileRequest) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileRequest) SetCreateCommand(v *CreateAdHocFileRequestCreateCommand) *CreateAdHocFileRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateAdHocFileRequest) SetOpTenantId(v int64) *CreateAdHocFileRequest {
	s.OpTenantId = &v
	return s
}

type CreateAdHocFileRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// select 1;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /xx1/xx2/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11212133
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateAdHocFileRequestCreateCommand) String() string {
	return tea.Prettify(s)
}

func (s CreateAdHocFileRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileRequestCreateCommand) SetContent(v string) *CreateAdHocFileRequestCreateCommand {
	s.Content = &v
	return s
}

func (s *CreateAdHocFileRequestCreateCommand) SetDirectory(v string) *CreateAdHocFileRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateAdHocFileRequestCreateCommand) SetName(v string) *CreateAdHocFileRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateAdHocFileRequestCreateCommand) SetProjectId(v int64) *CreateAdHocFileRequestCreateCommand {
	s.ProjectId = &v
	return s
}

type CreateAdHocFileShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateAdHocFileShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAdHocFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileShrinkRequest) SetCreateCommandShrink(v string) *CreateAdHocFileShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateAdHocFileShrinkRequest) SetOpTenantId(v int64) *CreateAdHocFileShrinkRequest {
	s.OpTenantId = &v
	return s
}

type CreateAdHocFileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1212313222
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAdHocFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAdHocFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileResponseBody) SetCode(v string) *CreateAdHocFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetFileId(v int64) *CreateAdHocFileResponseBody {
	s.FileId = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetHttpStatusCode(v int32) *CreateAdHocFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetMessage(v string) *CreateAdHocFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetRequestId(v string) *CreateAdHocFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdHocFileResponseBody) SetSuccess(v bool) *CreateAdHocFileResponseBody {
	s.Success = &v
	return s
}

type CreateAdHocFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAdHocFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAdHocFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAdHocFileResponse) GoString() string {
	return s.String()
}

func (s *CreateAdHocFileResponse) SetHeaders(v map[string]*string) *CreateAdHocFileResponse {
	s.Headers = v
	return s
}

func (s *CreateAdHocFileResponse) SetStatusCode(v int32) *CreateAdHocFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAdHocFileResponse) SetBody(v *CreateAdHocFileResponseBody) *CreateAdHocFileResponse {
	s.Body = v
	return s
}

type CreateDataSourceRequest struct {
	CreateCommand *CreateDataSourceRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) SetCreateCommand(v *CreateDataSourceRequestCreateCommand) *CreateDataSourceRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDataSourceRequest) SetOpTenantId(v int64) *CreateDataSourceRequest {
	s.OpTenantId = &v
	return s
}

type CreateDataSourceRequestCreateCommand struct {
	DevDataSourceCreate *CreateDataSourceRequestCreateCommandDevDataSourceCreate `json:"DevDataSourceCreate,omitempty" xml:"DevDataSourceCreate,omitempty" type:"Struct"`
	// 数据源创建结构体
	ProdDataSourceCreate *CreateDataSourceRequestCreateCommandProdDataSourceCreate `json:"ProdDataSourceCreate,omitempty" xml:"ProdDataSourceCreate,omitempty" type:"Struct"`
}

func (s CreateDataSourceRequestCreateCommand) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommand) SetDevDataSourceCreate(v *CreateDataSourceRequestCreateCommandDevDataSourceCreate) *CreateDataSourceRequestCreateCommand {
	s.DevDataSourceCreate = v
	return s
}

func (s *CreateDataSourceRequestCreateCommand) SetProdDataSourceCreate(v *CreateDataSourceRequestCreateCommandProdDataSourceCreate) *CreateDataSourceRequestCreateCommand {
	s.ProdDataSourceCreate = v
	return s
}

type CreateDataSourceRequestCreateCommandDevDataSourceCreate struct {
	// 数据源创建结构体
	DataSourceCreate *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate `json:"DataSourceCreate,omitempty" xml:"DataSourceCreate,omitempty" type:"Struct"`
	// example:
	//
	// 1011
	ProdDataSourceId *int64 `json:"ProdDataSourceId,omitempty" xml:"ProdDataSourceId,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreate) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreate) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreate) SetDataSourceCreate(v *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) *CreateDataSourceRequestCreateCommandDevDataSourceCreate {
	s.DataSourceCreate = v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreate) SetProdDataSourceId(v int64) *CreateDataSourceRequestCreateCommandDevDataSourceCreate {
	s.ProdDataSourceId = &v
	return s
}

type CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate struct {
	// example:
	//
	// true
	CheckActivity *bool `json:"CheckActivity,omitempty" xml:"CheckActivity,omitempty"`
	// This parameter is required.
	ConfigItemList []*CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// example:
	//
	// datasource for xxx in dev
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dp_test_dev
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetCheckActivity(v bool) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.CheckActivity = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetConfigItemList(v []*CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.ConfigItemList = v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetDescription(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetName(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate) SetType(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreate {
	s.Type = &v
	return s
}

type CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) SetKey(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList {
	s.Key = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList) SetValue(v string) *CreateDataSourceRequestCreateCommandDevDataSourceCreateDataSourceCreateConfigItemList {
	s.Value = &v
	return s
}

type CreateDataSourceRequestCreateCommandProdDataSourceCreate struct {
	// example:
	//
	// true
	CheckActivity *bool `json:"CheckActivity,omitempty" xml:"CheckActivity,omitempty"`
	// This parameter is required.
	ConfigItemList []*CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// example:
	//
	// datasource for xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dp_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandProdDataSourceCreate) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandProdDataSourceCreate) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetCheckActivity(v bool) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.CheckActivity = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetConfigItemList(v []*CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.ConfigItemList = v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetDescription(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetName(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreate) SetType(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreate {
	s.Type = &v
	return s
}

type CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) SetKey(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList {
	s.Key = &v
	return s
}

func (s *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList) SetValue(v string) *CreateDataSourceRequestCreateCommandProdDataSourceCreateConfigItemList {
	s.Value = &v
	return s
}

type CreateDataSourceShrinkRequest struct {
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataSourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceShrinkRequest) SetCreateCommandShrink(v string) *CreateDataSourceShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDataSourceShrinkRequest) SetOpTenantId(v int64) *CreateDataSourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

type CreateDataSourceResponseBody struct {
	// example:
	//
	// OK
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateResult *CreateDataSourceResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) SetCode(v string) *CreateDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetCreateResult(v *CreateDataSourceResponseBodyCreateResult) *CreateDataSourceResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateDataSourceResponseBody) SetHttpStatusCode(v int32) *CreateDataSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetMessage(v string) *CreateDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetSuccess(v bool) *CreateDataSourceResponseBody {
	s.Success = &v
	return s
}

type CreateDataSourceResponseBodyCreateResult struct {
	// example:
	//
	// 123
	DevDataSourceId *int64 `json:"DevDataSourceId,omitempty" xml:"DevDataSourceId,omitempty"`
	// example:
	//
	// 12345
	ProdDataSourceId *int64 `json:"ProdDataSourceId,omitempty" xml:"ProdDataSourceId,omitempty"`
}

func (s CreateDataSourceResponseBodyCreateResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBodyCreateResult) SetDevDataSourceId(v int64) *CreateDataSourceResponseBodyCreateResult {
	s.DevDataSourceId = &v
	return s
}

func (s *CreateDataSourceResponseBodyCreateResult) SetProdDataSourceId(v int64) *CreateDataSourceResponseBodyCreateResult {
	s.ProdDataSourceId = &v
	return s
}

type CreateDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponse) SetHeaders(v map[string]*string) *CreateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceResponse) SetStatusCode(v int32) *CreateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceResponse) SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse {
	s.Body = v
	return s
}

type CreateDirectoryRequest struct {
	// This parameter is required.
	CreateCommand *CreateDirectoryRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateDirectoryRequest) SetCreateCommand(v *CreateDirectoryRequestCreateCommand) *CreateDirectoryRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDirectoryRequest) SetOpTenantId(v int64) *CreateDirectoryRequest {
	s.OpTenantId = &v
	return s
}

type CreateDirectoryRequestCreateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// tempCode
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1212344
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateDirectoryRequestCreateCommand) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDirectoryRequestCreateCommand) SetCategory(v string) *CreateDirectoryRequestCreateCommand {
	s.Category = &v
	return s
}

func (s *CreateDirectoryRequestCreateCommand) SetDirectory(v string) *CreateDirectoryRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateDirectoryRequestCreateCommand) SetName(v string) *CreateDirectoryRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateDirectoryRequestCreateCommand) SetProjectId(v int64) *CreateDirectoryRequestCreateCommand {
	s.ProjectId = &v
	return s
}

type CreateDirectoryShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDirectoryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDirectoryShrinkRequest) SetCreateCommandShrink(v string) *CreateDirectoryShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateDirectoryShrinkRequest) SetOpTenantId(v int64) *CreateDirectoryShrinkRequest {
	s.OpTenantId = &v
	return s
}

type CreateDirectoryResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1311113211
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponseBody) SetCode(v string) *CreateDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetFileId(v int64) *CreateDirectoryResponseBody {
	s.FileId = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetHttpStatusCode(v int32) *CreateDirectoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetMessage(v string) *CreateDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetRequestId(v string) *CreateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetSuccess(v bool) *CreateDirectoryResponseBody {
	s.Success = &v
	return s
}

type CreateDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponse) SetHeaders(v map[string]*string) *CreateDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateDirectoryResponse) SetStatusCode(v int32) *CreateDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDirectoryResponse) SetBody(v *CreateDirectoryResponseBody) *CreateDirectoryResponse {
	s.Body = v
	return s
}

type CreateNodeSupplementRequest struct {
	// This parameter is required.
	CreateCommand *CreateNodeSupplementRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// example:
	//
	// DEV/PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateNodeSupplementRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequest) SetCreateCommand(v *CreateNodeSupplementRequestCreateCommand) *CreateNodeSupplementRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateNodeSupplementRequest) SetEnv(v string) *CreateNodeSupplementRequest {
	s.Env = &v
	return s
}

func (s *CreateNodeSupplementRequest) SetOpTenantId(v int64) *CreateNodeSupplementRequest {
	s.OpTenantId = &v
	return s
}

type CreateNodeSupplementRequestCreateCommand struct {
	ContainAllDownStream *bool                                                           `json:"ContainAllDownStream,omitempty" xml:"ContainAllDownStream,omitempty"`
	DownStreamNodeIdList []*CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList `json:"DownStreamNodeIdList,omitempty" xml:"DownStreamNodeIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-06-01
	EndBizDate      *string                                                    `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	FilterList      []*CreateNodeSupplementRequestCreateCommandFilterList      `json:"FilterList,omitempty" xml:"FilterList,omitempty" type:"Repeated"`
	GlobalParamList []*CreateNodeSupplementRequestCreateCommandGlobalParamList `json:"GlobalParamList,omitempty" xml:"GlobalParamList,omitempty" type:"Repeated"`
	MaxDueTime      *string                                                    `json:"MaxDueTime,omitempty" xml:"MaxDueTime,omitempty"`
	MinDueTime      *string                                                    `json:"MinDueTime,omitempty" xml:"MinDueTime,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NodeIdList     []*CreateNodeSupplementRequestCreateCommandNodeIdList     `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	NodeParamsList []*CreateNodeSupplementRequestCreateCommandNodeParamsList `json:"NodeParamsList,omitempty" xml:"NodeParamsList,omitempty" type:"Repeated"`
	Parallelism    *int32                                                    `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 101121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-06-01
	StartBizDate *string `json:"StartBizDate,omitempty" xml:"StartBizDate,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommand) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommand) SetContainAllDownStream(v bool) *CreateNodeSupplementRequestCreateCommand {
	s.ContainAllDownStream = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetDownStreamNodeIdList(v []*CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) *CreateNodeSupplementRequestCreateCommand {
	s.DownStreamNodeIdList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetEndBizDate(v string) *CreateNodeSupplementRequestCreateCommand {
	s.EndBizDate = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetFilterList(v []*CreateNodeSupplementRequestCreateCommandFilterList) *CreateNodeSupplementRequestCreateCommand {
	s.FilterList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetGlobalParamList(v []*CreateNodeSupplementRequestCreateCommandGlobalParamList) *CreateNodeSupplementRequestCreateCommand {
	s.GlobalParamList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetMaxDueTime(v string) *CreateNodeSupplementRequestCreateCommand {
	s.MaxDueTime = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetMinDueTime(v string) *CreateNodeSupplementRequestCreateCommand {
	s.MinDueTime = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetName(v string) *CreateNodeSupplementRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetNodeIdList(v []*CreateNodeSupplementRequestCreateCommandNodeIdList) *CreateNodeSupplementRequestCreateCommand {
	s.NodeIdList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetNodeParamsList(v []*CreateNodeSupplementRequestCreateCommandNodeParamsList) *CreateNodeSupplementRequestCreateCommand {
	s.NodeParamsList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetParallelism(v int32) *CreateNodeSupplementRequestCreateCommand {
	s.Parallelism = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetProjectId(v int64) *CreateNodeSupplementRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommand) SetStartBizDate(v string) *CreateNodeSupplementRequestCreateCommand {
	s.StartBizDate = &v
	return s
}

type CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList struct {
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	Id          *string   `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) SetFieldIdList(v []*string) *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList {
	s.FieldIdList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList) SetId(v string) *CreateNodeSupplementRequestCreateCommandDownStreamNodeIdList {
	s.Id = &v
	return s
}

type CreateNodeSupplementRequestCreateCommandFilterList struct {
	Exclude   *bool     `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s CreateNodeSupplementRequestCreateCommandFilterList) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandFilterList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) SetExclude(v bool) *CreateNodeSupplementRequestCreateCommandFilterList {
	s.Exclude = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) SetKey(v string) *CreateNodeSupplementRequestCreateCommandFilterList {
	s.Key = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandFilterList) SetValueList(v []*string) *CreateNodeSupplementRequestCreateCommandFilterList {
	s.ValueList = v
	return s
}

type CreateNodeSupplementRequestCreateCommandGlobalParamList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommandGlobalParamList) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandGlobalParamList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandGlobalParamList) SetKey(v string) *CreateNodeSupplementRequestCreateCommandGlobalParamList {
	s.Key = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandGlobalParamList) SetValue(v string) *CreateNodeSupplementRequestCreateCommandGlobalParamList {
	s.Value = &v
	return s
}

type CreateNodeSupplementRequestCreateCommandNodeIdList struct {
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommandNodeIdList) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandNodeIdList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandNodeIdList) SetFieldIdList(v []*string) *CreateNodeSupplementRequestCreateCommandNodeIdList {
	s.FieldIdList = v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeIdList) SetId(v string) *CreateNodeSupplementRequestCreateCommandNodeIdList {
	s.Id = &v
	return s
}

type CreateNodeSupplementRequestCreateCommandNodeParamsList struct {
	NodeId    *string                                                            `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ParamList []*CreateNodeSupplementRequestCreateCommandNodeParamsListParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
}

func (s CreateNodeSupplementRequestCreateCommandNodeParamsList) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandNodeParamsList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsList) SetNodeId(v string) *CreateNodeSupplementRequestCreateCommandNodeParamsList {
	s.NodeId = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsList) SetParamList(v []*CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) *CreateNodeSupplementRequestCreateCommandNodeParamsList {
	s.ParamList = v
	return s
}

type CreateNodeSupplementRequestCreateCommandNodeParamsListParamList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) SetKey(v string) *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList {
	s.Key = &v
	return s
}

func (s *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList) SetValue(v string) *CreateNodeSupplementRequestCreateCommandNodeParamsListParamList {
	s.Value = &v
	return s
}

type CreateNodeSupplementShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// example:
	//
	// DEV/PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateNodeSupplementShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementShrinkRequest) SetCreateCommandShrink(v string) *CreateNodeSupplementShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateNodeSupplementShrinkRequest) SetEnv(v string) *CreateNodeSupplementShrinkRequest {
	s.Env = &v
	return s
}

func (s *CreateNodeSupplementShrinkRequest) SetOpTenantId(v int64) *CreateNodeSupplementShrinkRequest {
	s.OpTenantId = &v
	return s
}

type CreateNodeSupplementResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// f_2264518792396800000_20210223_2329354897145659392
	SubmitId *string `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	// example:
	//
	// true/false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNodeSupplementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementResponseBody) SetCode(v string) *CreateNodeSupplementResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetHttpStatusCode(v int32) *CreateNodeSupplementResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetMessage(v string) *CreateNodeSupplementResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetRequestId(v string) *CreateNodeSupplementResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetSubmitId(v string) *CreateNodeSupplementResponseBody {
	s.SubmitId = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetSuccess(v bool) *CreateNodeSupplementResponseBody {
	s.Success = &v
	return s
}

type CreateNodeSupplementResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeSupplementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeSupplementResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeSupplementResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementResponse) SetHeaders(v map[string]*string) *CreateNodeSupplementResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeSupplementResponse) SetStatusCode(v int32) *CreateNodeSupplementResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeSupplementResponse) SetBody(v *CreateNodeSupplementResponseBody) *CreateNodeSupplementResponse {
	s.Body = v
	return s
}

type CreateUserGroupRequest struct {
	CreateCommand *CreateUserGroupRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) SetCreateCommand(v *CreateUserGroupRequestCreateCommand) *CreateUserGroupRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateUserGroupRequest) SetOpTenantId(v int64) *CreateUserGroupRequest {
	s.OpTenantId = &v
	return s
}

type CreateUserGroupRequestCreateCommand struct {
	// example:
	//
	// true
	Active          *bool     `json:"Active,omitempty" xml:"Active,omitempty"`
	AdminUserIdList []*string `json:"AdminUserIdList,omitempty" xml:"AdminUserIdList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateUserGroupRequestCreateCommand) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequestCreateCommand) SetActive(v bool) *CreateUserGroupRequestCreateCommand {
	s.Active = &v
	return s
}

func (s *CreateUserGroupRequestCreateCommand) SetAdminUserIdList(v []*string) *CreateUserGroupRequestCreateCommand {
	s.AdminUserIdList = v
	return s
}

func (s *CreateUserGroupRequestCreateCommand) SetDescription(v string) *CreateUserGroupRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateUserGroupRequestCreateCommand) SetName(v string) *CreateUserGroupRequestCreateCommand {
	s.Name = &v
	return s
}

type CreateUserGroupShrinkRequest struct {
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateUserGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupShrinkRequest) SetCreateCommandShrink(v string) *CreateUserGroupShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateUserGroupShrinkRequest) SetOpTenantId(v int64) *CreateUserGroupShrinkRequest {
	s.OpTenantId = &v
	return s
}

type CreateUserGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 2313131
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s CreateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBody) SetCode(v string) *CreateUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetHttpStatusCode(v int32) *CreateUserGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetMessage(v string) *CreateUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetRequestId(v string) *CreateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetSuccess(v bool) *CreateUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetUserGroupId(v string) *CreateUserGroupResponseBody {
	s.UserGroupId = &v
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

type DeleteAdHocFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12121111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12132323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteAdHocFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAdHocFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteAdHocFileRequest) SetFileId(v int64) *DeleteAdHocFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteAdHocFileRequest) SetOpTenantId(v int64) *DeleteAdHocFileRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteAdHocFileRequest) SetProjectId(v int64) *DeleteAdHocFileRequest {
	s.ProjectId = &v
	return s
}

type DeleteAdHocFileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAdHocFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAdHocFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdHocFileResponseBody) SetCode(v string) *DeleteAdHocFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) SetHttpStatusCode(v int32) *DeleteAdHocFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) SetMessage(v string) *DeleteAdHocFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) SetRequestId(v string) *DeleteAdHocFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) SetSuccess(v bool) *DeleteAdHocFileResponseBody {
	s.Success = &v
	return s
}

type DeleteAdHocFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAdHocFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAdHocFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAdHocFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdHocFileResponse) SetHeaders(v map[string]*string) *DeleteAdHocFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdHocFileResponse) SetStatusCode(v int32) *DeleteAdHocFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdHocFileResponse) SetBody(v *DeleteAdHocFileResponseBody) *DeleteAdHocFileResponse {
	s.Body = v
	return s
}

type DeleteDataSourceRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteDataSourceRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) SetDeleteCommand(v *DeleteDataSourceRequestDeleteCommand) *DeleteDataSourceRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteDataSourceRequest) SetOpTenantId(v int64) *DeleteDataSourceRequest {
	s.OpTenantId = &v
	return s
}

type DeleteDataSourceRequestDeleteCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13121
	ProdDataSourceId *int64 `json:"ProdDataSourceId,omitempty" xml:"ProdDataSourceId,omitempty"`
}

func (s DeleteDataSourceRequestDeleteCommand) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequestDeleteCommand) SetMode(v string) *DeleteDataSourceRequestDeleteCommand {
	s.Mode = &v
	return s
}

func (s *DeleteDataSourceRequestDeleteCommand) SetProdDataSourceId(v int64) *DeleteDataSourceRequestDeleteCommand {
	s.ProdDataSourceId = &v
	return s
}

type DeleteDataSourceShrinkRequest struct {
	// This parameter is required.
	DeleteCommandShrink *string `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteDataSourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceShrinkRequest) SetDeleteCommandShrink(v string) *DeleteDataSourceShrinkRequest {
	s.DeleteCommandShrink = &v
	return s
}

func (s *DeleteDataSourceShrinkRequest) SetOpTenantId(v int64) *DeleteDataSourceShrinkRequest {
	s.OpTenantId = &v
	return s
}

type DeleteDataSourceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) SetCode(v string) *DeleteDataSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetData(v bool) *DeleteDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetHttpStatusCode(v int32) *DeleteDataSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetMessage(v string) *DeleteDataSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetSuccess(v bool) *DeleteDataSourceResponseBody {
	s.Success = &v
	return s
}

type DeleteDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetStatusCode(v int32) *DeleteDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
}

type DeleteDirectoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12121111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12132323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryRequest) SetFileId(v int64) *DeleteDirectoryRequest {
	s.FileId = &v
	return s
}

func (s *DeleteDirectoryRequest) SetOpTenantId(v int64) *DeleteDirectoryRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteDirectoryRequest) SetProjectId(v int64) *DeleteDirectoryRequest {
	s.ProjectId = &v
	return s
}

type DeleteDirectoryResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponseBody) SetCode(v string) *DeleteDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDirectoryResponseBody) SetHttpStatusCode(v int32) *DeleteDirectoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDirectoryResponseBody) SetMessage(v string) *DeleteDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDirectoryResponseBody) SetRequestId(v string) *DeleteDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDirectoryResponseBody) SetSuccess(v bool) *DeleteDirectoryResponseBody {
	s.Success = &v
	return s
}

type DeleteDirectoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponse) SetHeaders(v map[string]*string) *DeleteDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectoryResponse) SetStatusCode(v int32) *DeleteDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDirectoryResponse) SetBody(v *DeleteDirectoryResponseBody) *DeleteDirectoryResponse {
	s.Body = v
	return s
}

type DeleteUserGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 232131
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) SetOpTenantId(v int64) *DeleteUserGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetUserGroupId(v string) *DeleteUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type DeleteUserGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DeleteUserGroupResponseBody) SetCode(v string) *DeleteUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetData(v bool) *DeleteUserGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetHttpStatusCode(v int32) *DeleteUserGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetMessage(v string) *DeleteUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
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

type ExecuteManualNodeRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ExecuteCommand *ExecuteManualNodeRequestExecuteCommand `json:"ExecuteCommand,omitempty" xml:"ExecuteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ExecuteManualNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteManualNodeRequest) GoString() string {
	return s.String()
}

func (s *ExecuteManualNodeRequest) SetEnv(v string) *ExecuteManualNodeRequest {
	s.Env = &v
	return s
}

func (s *ExecuteManualNodeRequest) SetExecuteCommand(v *ExecuteManualNodeRequestExecuteCommand) *ExecuteManualNodeRequest {
	s.ExecuteCommand = v
	return s
}

func (s *ExecuteManualNodeRequest) SetOpTenantId(v int64) *ExecuteManualNodeRequest {
	s.OpTenantId = &v
	return s
}

type ExecuteManualNodeRequestExecuteCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-07
	EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	// example:
	//
	// xx测试
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// n_12132
	NodeId    *string                                            `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ParamList []*ExecuteManualNodeRequestExecuteCommandParamList `json:"ParamList,omitempty" xml:"ParamList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123324
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-01
	StartBizDate *string `json:"StartBizDate,omitempty" xml:"StartBizDate,omitempty"`
}

func (s ExecuteManualNodeRequestExecuteCommand) String() string {
	return tea.Prettify(s)
}

func (s ExecuteManualNodeRequestExecuteCommand) GoString() string {
	return s.String()
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetEndBizDate(v string) *ExecuteManualNodeRequestExecuteCommand {
	s.EndBizDate = &v
	return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetFlowName(v string) *ExecuteManualNodeRequestExecuteCommand {
	s.FlowName = &v
	return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetNodeId(v string) *ExecuteManualNodeRequestExecuteCommand {
	s.NodeId = &v
	return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetParamList(v []*ExecuteManualNodeRequestExecuteCommandParamList) *ExecuteManualNodeRequestExecuteCommand {
	s.ParamList = v
	return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetProjectId(v int64) *ExecuteManualNodeRequestExecuteCommand {
	s.ProjectId = &v
	return s
}

func (s *ExecuteManualNodeRequestExecuteCommand) SetStartBizDate(v string) *ExecuteManualNodeRequestExecuteCommand {
	s.StartBizDate = &v
	return s
}

type ExecuteManualNodeRequestExecuteCommandParamList struct {
	// example:
	//
	// param1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecuteManualNodeRequestExecuteCommandParamList) String() string {
	return tea.Prettify(s)
}

func (s ExecuteManualNodeRequestExecuteCommandParamList) GoString() string {
	return s.String()
}

func (s *ExecuteManualNodeRequestExecuteCommandParamList) SetKey(v string) *ExecuteManualNodeRequestExecuteCommandParamList {
	s.Key = &v
	return s
}

func (s *ExecuteManualNodeRequestExecuteCommandParamList) SetValue(v string) *ExecuteManualNodeRequestExecuteCommandParamList {
	s.Value = &v
	return s
}

type ExecuteManualNodeShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ExecuteCommandShrink *string `json:"ExecuteCommand,omitempty" xml:"ExecuteCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ExecuteManualNodeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteManualNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteManualNodeShrinkRequest) SetEnv(v string) *ExecuteManualNodeShrinkRequest {
	s.Env = &v
	return s
}

func (s *ExecuteManualNodeShrinkRequest) SetExecuteCommandShrink(v string) *ExecuteManualNodeShrinkRequest {
	s.ExecuteCommandShrink = &v
	return s
}

func (s *ExecuteManualNodeShrinkRequest) SetOpTenantId(v int64) *ExecuteManualNodeShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ExecuteManualNodeResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// f_1231_121324
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteManualNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteManualNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteManualNodeResponseBody) SetCode(v string) *ExecuteManualNodeResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteManualNodeResponseBody) SetFlowId(v string) *ExecuteManualNodeResponseBody {
	s.FlowId = &v
	return s
}

func (s *ExecuteManualNodeResponseBody) SetHttpStatusCode(v int32) *ExecuteManualNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteManualNodeResponseBody) SetMessage(v string) *ExecuteManualNodeResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteManualNodeResponseBody) SetRequestId(v string) *ExecuteManualNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteManualNodeResponseBody) SetSuccess(v bool) *ExecuteManualNodeResponseBody {
	s.Success = &v
	return s
}

type ExecuteManualNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteManualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteManualNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteManualNodeResponse) GoString() string {
	return s.String()
}

func (s *ExecuteManualNodeResponse) SetHeaders(v map[string]*string) *ExecuteManualNodeResponse {
	s.Headers = v
	return s
}

func (s *ExecuteManualNodeResponse) SetStatusCode(v int32) *ExecuteManualNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteManualNodeResponse) SetBody(v *ExecuteManualNodeResponseBody) *ExecuteManualNodeResponse {
	s.Body = v
	return s
}

type FixDataRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	FixDataCommand *FixDataRequestFixDataCommand `json:"FixDataCommand,omitempty" xml:"FixDataCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s FixDataRequest) String() string {
	return tea.Prettify(s)
}

func (s FixDataRequest) GoString() string {
	return s.String()
}

func (s *FixDataRequest) SetEnv(v string) *FixDataRequest {
	s.Env = &v
	return s
}

func (s *FixDataRequest) SetFixDataCommand(v *FixDataRequestFixDataCommand) *FixDataRequest {
	s.FixDataCommand = v
	return s
}

func (s *FixDataRequest) SetOpTenantId(v int64) *FixDataRequest {
	s.OpTenantId = &v
	return s
}

type FixDataRequestFixDataCommand struct {
	// example:
	//
	// false
	ContainRootInstance      *bool                                                   `json:"ContainRootInstance,omitempty" xml:"ContainRootInstance,omitempty"`
	DownStreamInstanceIdList []*FixDataRequestFixDataCommandDownStreamInstanceIdList `json:"DownStreamInstanceIdList,omitempty" xml:"DownStreamInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// ALL_INSTANCE
	DownstreamRange *string `json:"DownstreamRange,omitempty" xml:"DownstreamRange,omitempty"`
	// example:
	//
	// false
	ForceRerun *bool `json:"ForceRerun,omitempty" xml:"ForceRerun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 132344
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RootInstanceId *FixDataRequestFixDataCommandRootInstanceId `json:"RootInstanceId,omitempty" xml:"RootInstanceId,omitempty" type:"Struct"`
}

func (s FixDataRequestFixDataCommand) String() string {
	return tea.Prettify(s)
}

func (s FixDataRequestFixDataCommand) GoString() string {
	return s.String()
}

func (s *FixDataRequestFixDataCommand) SetContainRootInstance(v bool) *FixDataRequestFixDataCommand {
	s.ContainRootInstance = &v
	return s
}

func (s *FixDataRequestFixDataCommand) SetDownStreamInstanceIdList(v []*FixDataRequestFixDataCommandDownStreamInstanceIdList) *FixDataRequestFixDataCommand {
	s.DownStreamInstanceIdList = v
	return s
}

func (s *FixDataRequestFixDataCommand) SetDownstreamRange(v string) *FixDataRequestFixDataCommand {
	s.DownstreamRange = &v
	return s
}

func (s *FixDataRequestFixDataCommand) SetForceRerun(v bool) *FixDataRequestFixDataCommand {
	s.ForceRerun = &v
	return s
}

func (s *FixDataRequestFixDataCommand) SetProjectId(v int64) *FixDataRequestFixDataCommand {
	s.ProjectId = &v
	return s
}

func (s *FixDataRequestFixDataCommand) SetRootInstanceId(v *FixDataRequestFixDataCommandRootInstanceId) *FixDataRequestFixDataCommand {
	s.RootInstanceId = v
	return s
}

type FixDataRequestFixDataCommandDownStreamInstanceIdList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// t_2323421
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FixDataRequestFixDataCommandDownStreamInstanceIdList) String() string {
	return tea.Prettify(s)
}

func (s FixDataRequestFixDataCommandDownStreamInstanceIdList) GoString() string {
	return s.String()
}

func (s *FixDataRequestFixDataCommandDownStreamInstanceIdList) SetFieldInstanceIdList(v []*string) *FixDataRequestFixDataCommandDownStreamInstanceIdList {
	s.FieldInstanceIdList = v
	return s
}

func (s *FixDataRequestFixDataCommandDownStreamInstanceIdList) SetId(v string) *FixDataRequestFixDataCommandDownStreamInstanceIdList {
	s.Id = &v
	return s
}

type FixDataRequestFixDataCommandRootInstanceId struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// t_2323111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FixDataRequestFixDataCommandRootInstanceId) String() string {
	return tea.Prettify(s)
}

func (s FixDataRequestFixDataCommandRootInstanceId) GoString() string {
	return s.String()
}

func (s *FixDataRequestFixDataCommandRootInstanceId) SetFieldInstanceIdList(v []*string) *FixDataRequestFixDataCommandRootInstanceId {
	s.FieldInstanceIdList = v
	return s
}

func (s *FixDataRequestFixDataCommandRootInstanceId) SetId(v string) *FixDataRequestFixDataCommandRootInstanceId {
	s.Id = &v
	return s
}

type FixDataShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	FixDataCommandShrink *string `json:"FixDataCommand,omitempty" xml:"FixDataCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s FixDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FixDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *FixDataShrinkRequest) SetEnv(v string) *FixDataShrinkRequest {
	s.Env = &v
	return s
}

func (s *FixDataShrinkRequest) SetFixDataCommandShrink(v string) *FixDataShrinkRequest {
	s.FixDataCommandShrink = &v
	return s
}

func (s *FixDataShrinkRequest) SetOpTenantId(v int64) *FixDataShrinkRequest {
	s.OpTenantId = &v
	return s
}

type FixDataResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12324234
	SubmitId *string `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FixDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FixDataResponseBody) GoString() string {
	return s.String()
}

func (s *FixDataResponseBody) SetCode(v string) *FixDataResponseBody {
	s.Code = &v
	return s
}

func (s *FixDataResponseBody) SetHttpStatusCode(v int32) *FixDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FixDataResponseBody) SetMessage(v string) *FixDataResponseBody {
	s.Message = &v
	return s
}

func (s *FixDataResponseBody) SetRequestId(v string) *FixDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *FixDataResponseBody) SetSubmitId(v string) *FixDataResponseBody {
	s.SubmitId = &v
	return s
}

func (s *FixDataResponseBody) SetSuccess(v bool) *FixDataResponseBody {
	s.Success = &v
	return s
}

type FixDataResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FixDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FixDataResponse) String() string {
	return tea.Prettify(s)
}

func (s FixDataResponse) GoString() string {
	return s.String()
}

func (s *FixDataResponse) SetHeaders(v map[string]*string) *FixDataResponse {
	s.Headers = v
	return s
}

func (s *FixDataResponse) SetStatusCode(v int32) *FixDataResponse {
	s.StatusCode = &v
	return s
}

func (s *FixDataResponse) SetBody(v *FixDataResponseBody) *FixDataResponse {
	s.Body = v
	return s
}

type GetAdHocFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12121111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12132323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetAdHocFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdHocFileRequest) GoString() string {
	return s.String()
}

func (s *GetAdHocFileRequest) SetFileId(v int64) *GetAdHocFileRequest {
	s.FileId = &v
	return s
}

func (s *GetAdHocFileRequest) SetOpTenantId(v int64) *GetAdHocFileRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAdHocFileRequest) SetProjectId(v int64) *GetAdHocFileRequest {
	s.ProjectId = &v
	return s
}

type GetAdHocFileResponseBody struct {
	// example:
	//
	// OK
	Code     *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	FileInfo *GetAdHocFileResponseBodyFileInfo `json:"FileInfo,omitempty" xml:"FileInfo,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAdHocFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdHocFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdHocFileResponseBody) SetCode(v string) *GetAdHocFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetAdHocFileResponseBody) SetFileInfo(v *GetAdHocFileResponseBodyFileInfo) *GetAdHocFileResponseBody {
	s.FileInfo = v
	return s
}

func (s *GetAdHocFileResponseBody) SetHttpStatusCode(v int32) *GetAdHocFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAdHocFileResponseBody) SetMessage(v string) *GetAdHocFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetAdHocFileResponseBody) SetRequestId(v string) *GetAdHocFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdHocFileResponseBody) SetSuccess(v bool) *GetAdHocFileResponseBody {
	s.Success = &v
	return s
}

type GetAdHocFileResponseBodyFileInfo struct {
	// example:
	//
	// select 1;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 12121
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// /xx1/xx2/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// 12121111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 12121
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12132323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetAdHocFileResponseBodyFileInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAdHocFileResponseBodyFileInfo) GoString() string {
	return s.String()
}

func (s *GetAdHocFileResponseBodyFileInfo) SetContent(v string) *GetAdHocFileResponseBodyFileInfo {
	s.Content = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetCreator(v string) *GetAdHocFileResponseBodyFileInfo {
	s.Creator = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetDirectory(v string) *GetAdHocFileResponseBodyFileInfo {
	s.Directory = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetId(v int64) *GetAdHocFileResponseBodyFileInfo {
	s.Id = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetLastModifier(v string) *GetAdHocFileResponseBodyFileInfo {
	s.LastModifier = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetName(v string) *GetAdHocFileResponseBodyFileInfo {
	s.Name = &v
	return s
}

func (s *GetAdHocFileResponseBodyFileInfo) SetProjectId(v int64) *GetAdHocFileResponseBodyFileInfo {
	s.ProjectId = &v
	return s
}

type GetAdHocFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdHocFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdHocFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdHocFileResponse) GoString() string {
	return s.String()
}

func (s *GetAdHocFileResponse) SetHeaders(v map[string]*string) *GetAdHocFileResponse {
	s.Headers = v
	return s
}

func (s *GetAdHocFileResponse) SetStatusCode(v int32) *GetAdHocFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdHocFileResponse) SetBody(v *GetAdHocFileResponseBody) *GetAdHocFileResponse {
	s.Body = v
	return s
}

type GetDevObjectDependencyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DATA_PROCESS
	ObjectFrom *string `json:"ObjectFrom,omitempty" xml:"ObjectFrom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7026498387616064
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7026498387616064
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7021037162911616L
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDevObjectDependencyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevObjectDependencyRequest) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyRequest) SetObjectFrom(v string) *GetDevObjectDependencyRequest {
	s.ObjectFrom = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetObjectId(v string) *GetDevObjectDependencyRequest {
	s.ObjectId = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetObjectType(v string) *GetDevObjectDependencyRequest {
	s.ObjectType = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetOpTenantId(v int64) *GetDevObjectDependencyRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDevObjectDependencyRequest) SetProjectId(v int64) *GetDevObjectDependencyRequest {
	s.ProjectId = &v
	return s
}

type GetDevObjectDependencyResponseBody struct {
	// example:
	//
	// OK
	Code                    *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	DevObjectDependencyList []*GetDevObjectDependencyResponseBodyDevObjectDependencyList `json:"DevObjectDependencyList,omitempty" xml:"DevObjectDependencyList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDevObjectDependencyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevObjectDependencyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBody) SetCode(v string) *GetDevObjectDependencyResponseBody {
	s.Code = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetDevObjectDependencyList(v []*GetDevObjectDependencyResponseBodyDevObjectDependencyList) *GetDevObjectDependencyResponseBody {
	s.DevObjectDependencyList = v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetHttpStatusCode(v int32) *GetDevObjectDependencyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetMessage(v string) *GetDevObjectDependencyResponseBody {
	s.Message = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetRequestId(v string) *GetDevObjectDependencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetSuccess(v bool) *GetDevObjectDependencyResponseBody {
	s.Success = &v
	return s
}

type GetDevObjectDependencyResponseBodyDevObjectDependencyList struct {
	// example:
	//
	// true
	AutoParse *bool `json:"AutoParse,omitempty" xml:"AutoParse,omitempty"`
	// example:
	//
	// SCRIPT
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 13111
	BizUnitId   *string `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 0 0 0 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// true
	CustomCronExpression *bool                                                                      `json:"CustomCronExpression,omitempty" xml:"CustomCronExpression,omitempty"`
	DependFieldList      []*string                                                                  `json:"DependFieldList,omitempty" xml:"DependFieldList,omitempty" type:"Repeated"`
	DependencyPeriod     *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod `json:"DependencyPeriod,omitempty" xml:"DependencyPeriod,omitempty" type:"Struct"`
	// example:
	//
	// ALL
	DependencyStrategy *string `json:"DependencyStrategy,omitempty" xml:"DependencyStrategy,omitempty"`
	// example:
	//
	// true
	DimMidNode      *bool     `json:"DimMidNode,omitempty" xml:"DimMidNode,omitempty"`
	EffectFieldList []*string `json:"EffectFieldList,omitempty" xml:"EffectFieldList,omitempty" type:"Repeated"`
	ExternalBizInfo *string   `json:"ExternalBizInfo,omitempty" xml:"ExternalBizInfo,omitempty"`
	// example:
	//
	// false
	ManuallyAdd *bool `json:"ManuallyAdd,omitempty" xml:"ManuallyAdd,omitempty"`
	// example:
	//
	// n_13211
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// n_xx
	NodeOutputName *string `json:"NodeOutputName,omitempty" xml:"NodeOutputName,omitempty"`
	// example:
	//
	// t_xx
	NodeOutputTableName *string `json:"NodeOutputTableName,omitempty" xml:"NodeOutputTableName,omitempty"`
	// example:
	//
	// DATA_PROCESS
	NodeType               *string                                                                            `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OutputContextParamList []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList `json:"OutputContextParamList,omitempty" xml:"OutputContextParamList,omitempty" type:"Repeated"`
	OwnerList              []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList              `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PeriodDiff *int32 `json:"PeriodDiff,omitempty" xml:"PeriodDiff,omitempty"`
	// example:
	//
	// 123131
	ProjectId   *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// DAILY
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// example:
	//
	// true
	SelfDepend *bool `json:"SelfDepend,omitempty" xml:"SelfDepend,omitempty"`
	// example:
	//
	// SHELL
	SubBizType *string `json:"SubBizType,omitempty" xml:"SubBizType,omitempty"`
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyList) String() string {
	return tea.Prettify(s)
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyList) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetAutoParse(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.AutoParse = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetBizType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.BizType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetBizUnitId(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.BizUnitId = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetBizUnitName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.BizUnitName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetCronExpression(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.CronExpression = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetCustomCronExpression(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.CustomCronExpression = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetDependFieldList(v []*string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.DependFieldList = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetDependencyPeriod(v *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.DependencyPeriod = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetDependencyStrategy(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.DependencyStrategy = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetDimMidNode(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.DimMidNode = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetEffectFieldList(v []*string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.EffectFieldList = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetExternalBizInfo(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ExternalBizInfo = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetManuallyAdd(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ManuallyAdd = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeId(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeId = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeOutputName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeOutputName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeOutputTableName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeOutputTableName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetOutputContextParamList(v []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.OutputContextParamList = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetOwnerList(v []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.OwnerList = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetPeriodDiff(v int32) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.PeriodDiff = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetProjectId(v int64) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ProjectId = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetProjectName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ProjectName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetScheduleType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ScheduleType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetSelfDepend(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.SelfDepend = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetSubBizType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.SubBizType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetValid(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.Valid = &v
	return s
}

type GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod struct {
	// example:
	//
	// 1
	PeriodOffset *int32 `json:"PeriodOffset,omitempty" xml:"PeriodOffset,omitempty"`
	// example:
	//
	// CURRENT_PERIOD
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) String() string {
	return tea.Prettify(s)
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) SetPeriodOffset(v int32) *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod {
	s.PeriodOffset = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) SetPeriodType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod {
	s.PeriodType = &v
	return s
}

type GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList struct {
	// example:
	//
	// v1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// xxtest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// v1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) String() string {
	return tea.Prettify(s)
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) SetDefaultValue(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList {
	s.DefaultValue = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) SetDescription(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList {
	s.Description = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) SetKey(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList {
	s.Key = &v
	return s
}

type GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList struct {
	// example:
	//
	// 11123
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) String() string {
	return tea.Prettify(s)
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) SetId(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList {
	s.Id = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) SetName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList {
	s.Name = &v
	return s
}

type GetDevObjectDependencyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDevObjectDependencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDevObjectDependencyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevObjectDependencyResponse) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponse) SetHeaders(v map[string]*string) *GetDevObjectDependencyResponse {
	s.Headers = v
	return s
}

func (s *GetDevObjectDependencyResponse) SetStatusCode(v int32) *GetDevObjectDependencyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevObjectDependencyResponse) SetBody(v *GetDevObjectDependencyResponseBody) *GetDevObjectDependencyResponse {
	s.Body = v
	return s
}

type GetInstanceDownStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	InstanceGet *GetInstanceDownStreamRequestInstanceGet `json:"InstanceGet,omitempty" xml:"InstanceGet,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// SUCCESS
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
}

func (s GetInstanceDownStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDownStreamRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamRequest) SetDownStreamDepth(v int32) *GetInstanceDownStreamRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceDownStreamRequest) SetEnv(v string) *GetInstanceDownStreamRequest {
	s.Env = &v
	return s
}

func (s *GetInstanceDownStreamRequest) SetInstanceGet(v *GetInstanceDownStreamRequestInstanceGet) *GetInstanceDownStreamRequest {
	s.InstanceGet = v
	return s
}

func (s *GetInstanceDownStreamRequest) SetOpTenantId(v int64) *GetInstanceDownStreamRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetInstanceDownStreamRequest) SetRunStatus(v string) *GetInstanceDownStreamRequest {
	s.RunStatus = &v
	return s
}

type GetInstanceDownStreamRequestInstanceGet struct {
	// This parameter is required.
	//
	// example:
	//
	// t_5929472_20210411_9577721
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetInstanceDownStreamRequestInstanceGet) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDownStreamRequestInstanceGet) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamRequestInstanceGet) SetInstanceId(v string) *GetInstanceDownStreamRequestInstanceGet {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceDownStreamRequestInstanceGet) SetNodeType(v string) *GetInstanceDownStreamRequestInstanceGet {
	s.NodeType = &v
	return s
}

type GetInstanceDownStreamShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	InstanceGetShrink *string `json:"InstanceGet,omitempty" xml:"InstanceGet,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// SUCCESS
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
}

func (s GetInstanceDownStreamShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDownStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamShrinkRequest) SetDownStreamDepth(v int32) *GetInstanceDownStreamShrinkRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) SetEnv(v string) *GetInstanceDownStreamShrinkRequest {
	s.Env = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) SetInstanceGetShrink(v string) *GetInstanceDownStreamShrinkRequest {
	s.InstanceGetShrink = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) SetOpTenantId(v int64) *GetInstanceDownStreamShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetInstanceDownStreamShrinkRequest) SetRunStatus(v string) *GetInstanceDownStreamShrinkRequest {
	s.RunStatus = &v
	return s
}

type GetInstanceDownStreamResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode       *int32                                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceRelationList []*GetInstanceDownStreamResponseBodyInstanceRelationList `json:"InstanceRelationList,omitempty" xml:"InstanceRelationList,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceDownStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDownStreamResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponseBody) SetCode(v string) *GetInstanceDownStreamResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetHttpStatusCode(v int32) *GetInstanceDownStreamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetInstanceRelationList(v []*GetInstanceDownStreamResponseBodyInstanceRelationList) *GetInstanceDownStreamResponseBody {
	s.InstanceRelationList = v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetMessage(v string) *GetInstanceDownStreamResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetRequestId(v string) *GetInstanceDownStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceDownStreamResponseBody) SetSuccess(v bool) *GetInstanceDownStreamResponseBody {
	s.Success = &v
	return s
}

type GetInstanceDownStreamResponseBodyInstanceRelationList struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// {"a":"x"}
	ExtendInfo        *string                                                                   `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FieldInstanceList []*GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList `json:"FieldInstanceList,omitempty" xml:"FieldInstanceList,omitempty" type:"Repeated"`
	InstanceInfo      *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo        `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
	// example:
	//
	// OPTIONAL
	SelectStatus *string `json:"SelectStatus,omitempty" xml:"SelectStatus,omitempty"`
	// example:
	//
	// FIELD_DELETED
	SelectStatusCause *string `json:"SelectStatusCause,omitempty" xml:"SelectStatusCause,omitempty"`
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationList) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetDownStreamDepth(v int32) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetExtendInfo(v string) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.ExtendInfo = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetFieldInstanceList(v []*GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.FieldInstanceList = v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetInstanceInfo(v *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.InstanceInfo = v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetRunStatus(v string) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.RunStatus = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetSelectStatus(v string) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.SelectStatus = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationList) SetSelectStatusCause(v string) *GetInstanceDownStreamResponseBodyInstanceRelationList {
	s.SelectStatusCause = &v
	return s
}

type GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList struct {
	// example:
	//
	// t_23211
	FieldInstanceId *string `json:"FieldInstanceId,omitempty" xml:"FieldInstanceId,omitempty"`
	// example:
	//
	// SUCCESS
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
	// example:
	//
	// OPTIONAL
	SelectStatus *string `json:"SelectStatus,omitempty" xml:"SelectStatus,omitempty"`
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) SetFieldInstanceId(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList {
	s.FieldInstanceId = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) SetRunStatus(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList {
	s.RunStatus = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList) SetSelectStatus(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListFieldInstanceList {
	s.SelectStatus = &v
	return s
}

type GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo struct {
	// example:
	//
	// t_232411
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) SetId(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo {
	s.Id = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) SetName(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo {
	s.Name = &v
	return s
}

func (s *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo) SetType(v string) *GetInstanceDownStreamResponseBodyInstanceRelationListInstanceInfo {
	s.Type = &v
	return s
}

type GetInstanceDownStreamResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceDownStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceDownStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDownStreamResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponse) SetHeaders(v map[string]*string) *GetInstanceDownStreamResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceDownStreamResponse) SetStatusCode(v int32) *GetInstanceDownStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceDownStreamResponse) SetBody(v *GetInstanceDownStreamResponseBody) *GetInstanceDownStreamResponse {
	s.Body = v
	return s
}

type GetInstanceUpDownStreamRequest struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	InstanceId *GetInstanceUpDownStreamRequestInstanceId `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1
	UpStreamDepth *int32 `json:"UpStreamDepth,omitempty" xml:"UpStreamDepth,omitempty"`
}

func (s GetInstanceUpDownStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamRequest) SetDownStreamDepth(v int32) *GetInstanceUpDownStreamRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetEnv(v string) *GetInstanceUpDownStreamRequest {
	s.Env = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetInstanceId(v *GetInstanceUpDownStreamRequestInstanceId) *GetInstanceUpDownStreamRequest {
	s.InstanceId = v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetOpTenantId(v int64) *GetInstanceUpDownStreamRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetProjectId(v int64) *GetInstanceUpDownStreamRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceUpDownStreamRequest) SetUpStreamDepth(v int32) *GetInstanceUpDownStreamRequest {
	s.UpStreamDepth = &v
	return s
}

type GetInstanceUpDownStreamRequestInstanceId struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// t_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetInstanceUpDownStreamRequestInstanceId) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamRequestInstanceId) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamRequestInstanceId) SetFieldInstanceIdList(v []*string) *GetInstanceUpDownStreamRequestInstanceId {
	s.FieldInstanceIdList = v
	return s
}

func (s *GetInstanceUpDownStreamRequestInstanceId) SetId(v string) *GetInstanceUpDownStreamRequestInstanceId {
	s.Id = &v
	return s
}

type GetInstanceUpDownStreamShrinkRequest struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	InstanceIdShrink *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1001121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1
	UpStreamDepth *int32 `json:"UpStreamDepth,omitempty" xml:"UpStreamDepth,omitempty"`
}

func (s GetInstanceUpDownStreamShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetDownStreamDepth(v int32) *GetInstanceUpDownStreamShrinkRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetEnv(v string) *GetInstanceUpDownStreamShrinkRequest {
	s.Env = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetInstanceIdShrink(v string) *GetInstanceUpDownStreamShrinkRequest {
	s.InstanceIdShrink = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetOpTenantId(v int64) *GetInstanceUpDownStreamShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetProjectId(v int64) *GetInstanceUpDownStreamShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceUpDownStreamShrinkRequest) SetUpStreamDepth(v int32) *GetInstanceUpDownStreamShrinkRequest {
	s.UpStreamDepth = &v
	return s
}

type GetInstanceUpDownStreamResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode  *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceDagInfo *GetInstanceUpDownStreamResponseBodyInstanceDagInfo `json:"InstanceDagInfo,omitempty" xml:"InstanceDagInfo,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceUpDownStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBody) SetCode(v string) *GetInstanceUpDownStreamResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetHttpStatusCode(v int32) *GetInstanceUpDownStreamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetInstanceDagInfo(v *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) *GetInstanceUpDownStreamResponseBody {
	s.InstanceDagInfo = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetMessage(v string) *GetInstanceUpDownStreamResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetRequestId(v string) *GetInstanceUpDownStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBody) SetSuccess(v bool) *GetInstanceUpDownStreamResponseBody {
	s.Success = &v
	return s
}

type GetInstanceUpDownStreamResponseBodyInstanceDagInfo struct {
	DownInstanceList  []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList  `json:"DownInstanceList,omitempty" xml:"DownInstanceList,omitempty" type:"Repeated"`
	StartInstanceList []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList `json:"StartInstanceList,omitempty" xml:"StartInstanceList,omitempty" type:"Repeated"`
	UpInstanceList    []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList    `json:"UpInstanceList,omitempty" xml:"UpInstanceList,omitempty" type:"Repeated"`
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) SetDownInstanceList(v []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) *GetInstanceUpDownStreamResponseBodyInstanceDagInfo {
	s.DownInstanceList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) SetStartInstanceList(v []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) *GetInstanceUpDownStreamResponseBodyInstanceDagInfo {
	s.StartInstanceList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfo) SetUpInstanceList(v []*GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) *GetInstanceUpDownStreamResponseBodyInstanceDagInfo {
	s.UpInstanceList = v
	return s
}

type GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// t_1234567
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// n_1234567
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetFieldInstanceIdList(v []*string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.FieldInstanceIdList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.Id = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetName(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.Name = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetNodeId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.NodeId = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList) SetNodeType(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoDownInstanceList {
	s.NodeType = &v
	return s
}

type GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// t_1234567
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// n_1234567
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetFieldInstanceIdList(v []*string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.FieldInstanceIdList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.Id = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetName(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.Name = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetNodeId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.NodeId = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList) SetNodeType(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoStartInstanceList {
	s.NodeType = &v
	return s
}

type GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// t_1234567
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// n_1234567
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetFieldInstanceIdList(v []*string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.FieldInstanceIdList = v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.Id = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetName(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.Name = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetNodeId(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.NodeId = &v
	return s
}

func (s *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList) SetNodeType(v string) *GetInstanceUpDownStreamResponseBodyInstanceDagInfoUpInstanceList {
	s.NodeType = &v
	return s
}

type GetInstanceUpDownStreamResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceUpDownStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceUpDownStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUpDownStreamResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponse) SetHeaders(v map[string]*string) *GetInstanceUpDownStreamResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceUpDownStreamResponse) SetStatusCode(v int32) *GetInstanceUpDownStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceUpDownStreamResponse) SetBody(v *GetInstanceUpDownStreamResponseBody) *GetInstanceUpDownStreamResponse {
	s.Body = v
	return s
}

type GetMyRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetMyRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMyRolesRequest) GoString() string {
	return s.String()
}

func (s *GetMyRolesRequest) SetOpTenantId(v int64) *GetMyRolesRequest {
	s.OpTenantId = &v
	return s
}

type GetMyRolesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleList  []*GetMyRolesResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMyRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMyRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GetMyRolesResponseBody) SetCode(v string) *GetMyRolesResponseBody {
	s.Code = &v
	return s
}

func (s *GetMyRolesResponseBody) SetHttpStatusCode(v int32) *GetMyRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMyRolesResponseBody) SetMessage(v string) *GetMyRolesResponseBody {
	s.Message = &v
	return s
}

func (s *GetMyRolesResponseBody) SetRequestId(v string) *GetMyRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMyRolesResponseBody) SetRoleList(v []*GetMyRolesResponseBodyRoleList) *GetMyRolesResponseBody {
	s.RoleList = v
	return s
}

func (s *GetMyRolesResponseBody) SetSuccess(v bool) *GetMyRolesResponseBody {
	s.Success = &v
	return s
}

type GetMyRolesResponseBodyRoleList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 300047957
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// dataphinAdmin
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMyRolesResponseBodyRoleList) String() string {
	return tea.Prettify(s)
}

func (s GetMyRolesResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *GetMyRolesResponseBodyRoleList) SetDescription(v string) *GetMyRolesResponseBodyRoleList {
	s.Description = &v
	return s
}

func (s *GetMyRolesResponseBodyRoleList) SetId(v int64) *GetMyRolesResponseBodyRoleList {
	s.Id = &v
	return s
}

func (s *GetMyRolesResponseBodyRoleList) SetName(v string) *GetMyRolesResponseBodyRoleList {
	s.Name = &v
	return s
}

type GetMyRolesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMyRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMyRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMyRolesResponse) GoString() string {
	return s.String()
}

func (s *GetMyRolesResponse) SetHeaders(v map[string]*string) *GetMyRolesResponse {
	s.Headers = v
	return s
}

func (s *GetMyRolesResponse) SetStatusCode(v int32) *GetMyRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMyRolesResponse) SetBody(v *GetMyRolesResponseBody) *GetMyRolesResponse {
	s.Body = v
	return s
}

type GetMyTenantsRequest struct {
	FeatureCodeList []*string `json:"FeatureCodeList,omitempty" xml:"FeatureCodeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetMyTenantsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMyTenantsRequest) GoString() string {
	return s.String()
}

func (s *GetMyTenantsRequest) SetFeatureCodeList(v []*string) *GetMyTenantsRequest {
	s.FeatureCodeList = v
	return s
}

func (s *GetMyTenantsRequest) SetOpTenantId(v int64) *GetMyTenantsRequest {
	s.OpTenantId = &v
	return s
}

type GetMyTenantsShrinkRequest struct {
	FeatureCodeListShrink *string `json:"FeatureCodeList,omitempty" xml:"FeatureCodeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetMyTenantsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMyTenantsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMyTenantsShrinkRequest) SetFeatureCodeListShrink(v string) *GetMyTenantsShrinkRequest {
	s.FeatureCodeListShrink = &v
	return s
}

func (s *GetMyTenantsShrinkRequest) SetOpTenantId(v int64) *GetMyTenantsShrinkRequest {
	s.OpTenantId = &v
	return s
}

type GetMyTenantsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TenantList []*GetMyTenantsResponseBodyTenantList `json:"TenantList,omitempty" xml:"TenantList,omitempty" type:"Repeated"`
}

func (s GetMyTenantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMyTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMyTenantsResponseBody) SetCode(v string) *GetMyTenantsResponseBody {
	s.Code = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetHttpStatusCode(v int32) *GetMyTenantsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetMessage(v string) *GetMyTenantsResponseBody {
	s.Message = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetRequestId(v string) *GetMyTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetSuccess(v bool) *GetMyTenantsResponseBody {
	s.Success = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetTenantList(v []*GetMyTenantsResponseBodyTenantList) *GetMyTenantsResponseBody {
	s.TenantList = v
	return s
}

type GetMyTenantsResponseBodyTenantList struct {
	// example:
	//
	// 1717343597000
	DeleteTime *int64 `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// example:
	//
	// false
	Deleted     *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 132311
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	OpsTenant *bool `json:"OpsTenant,omitempty" xml:"OpsTenant,omitempty"`
	// example:
	//
	// 21323231
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// false
	ResourceLimited *bool     `json:"ResourceLimited,omitempty" xml:"ResourceLimited,omitempty"`
	TenantTypeList  []*string `json:"TenantTypeList,omitempty" xml:"TenantTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// icon
	TitleType *string `json:"TitleType,omitempty" xml:"TitleType,omitempty"`
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s GetMyTenantsResponseBodyTenantList) String() string {
	return tea.Prettify(s)
}

func (s GetMyTenantsResponseBodyTenantList) GoString() string {
	return s.String()
}

func (s *GetMyTenantsResponseBodyTenantList) SetDeleteTime(v int64) *GetMyTenantsResponseBodyTenantList {
	s.DeleteTime = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetDeleted(v bool) *GetMyTenantsResponseBodyTenantList {
	s.Deleted = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetDescription(v string) *GetMyTenantsResponseBodyTenantList {
	s.Description = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetId(v int64) *GetMyTenantsResponseBodyTenantList {
	s.Id = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetName(v string) *GetMyTenantsResponseBodyTenantList {
	s.Name = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetOpsTenant(v bool) *GetMyTenantsResponseBodyTenantList {
	s.OpsTenant = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetOwnerId(v string) *GetMyTenantsResponseBodyTenantList {
	s.OwnerId = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetResourceLimited(v bool) *GetMyTenantsResponseBodyTenantList {
	s.ResourceLimited = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetTenantTypeList(v []*string) *GetMyTenantsResponseBodyTenantList {
	s.TenantTypeList = v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetTitleType(v string) *GetMyTenantsResponseBodyTenantList {
	s.TitleType = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetVisible(v bool) *GetMyTenantsResponseBodyTenantList {
	s.Visible = &v
	return s
}

type GetMyTenantsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMyTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMyTenantsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMyTenantsResponse) GoString() string {
	return s.String()
}

func (s *GetMyTenantsResponse) SetHeaders(v map[string]*string) *GetMyTenantsResponse {
	s.Headers = v
	return s
}

func (s *GetMyTenantsResponse) SetStatusCode(v int32) *GetMyTenantsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMyTenantsResponse) SetBody(v *GetMyTenantsResponseBody) *GetMyTenantsResponse {
	s.Body = v
	return s
}

type GetNodeUpDownStreamRequest struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	NodeId *GetNodeUpDownStreamRequestNodeId `json:"NodeId,omitempty" xml:"NodeId,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 113123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1
	UpStreamDepth *int32 `json:"UpStreamDepth,omitempty" xml:"UpStreamDepth,omitempty"`
}

func (s GetNodeUpDownStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamRequest) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamRequest) SetDownStreamDepth(v int32) *GetNodeUpDownStreamRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetEnv(v string) *GetNodeUpDownStreamRequest {
	s.Env = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetNodeId(v *GetNodeUpDownStreamRequestNodeId) *GetNodeUpDownStreamRequest {
	s.NodeId = v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetOpTenantId(v int64) *GetNodeUpDownStreamRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetProjectId(v int64) *GetNodeUpDownStreamRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNodeUpDownStreamRequest) SetUpStreamDepth(v int32) *GetNodeUpDownStreamRequest {
	s.UpStreamDepth = &v
	return s
}

type GetNodeUpDownStreamRequestNodeId struct {
	// example:
	//
	// 12
	FieldIdList *string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11313
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetNodeUpDownStreamRequestNodeId) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamRequestNodeId) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamRequestNodeId) SetFieldIdList(v string) *GetNodeUpDownStreamRequestNodeId {
	s.FieldIdList = &v
	return s
}

func (s *GetNodeUpDownStreamRequestNodeId) SetId(v string) *GetNodeUpDownStreamRequestNodeId {
	s.Id = &v
	return s
}

type GetNodeUpDownStreamShrinkRequest struct {
	// example:
	//
	// 1
	DownStreamDepth *int32 `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	NodeIdShrink *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 113123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 1
	UpStreamDepth *int32 `json:"UpStreamDepth,omitempty" xml:"UpStreamDepth,omitempty"`
}

func (s GetNodeUpDownStreamShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamShrinkRequest) SetDownStreamDepth(v int32) *GetNodeUpDownStreamShrinkRequest {
	s.DownStreamDepth = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetEnv(v string) *GetNodeUpDownStreamShrinkRequest {
	s.Env = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetNodeIdShrink(v string) *GetNodeUpDownStreamShrinkRequest {
	s.NodeIdShrink = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetOpTenantId(v int64) *GetNodeUpDownStreamShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetProjectId(v int64) *GetNodeUpDownStreamShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNodeUpDownStreamShrinkRequest) SetUpStreamDepth(v int32) *GetNodeUpDownStreamShrinkRequest {
	s.UpStreamDepth = &v
	return s
}

type GetNodeUpDownStreamResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message     *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeDagInfo *GetNodeUpDownStreamResponseBodyNodeDagInfo `json:"NodeDagInfo,omitempty" xml:"NodeDagInfo,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeUpDownStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBody) SetCode(v string) *GetNodeUpDownStreamResponseBody {
	s.Code = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetHttpStatusCode(v int32) *GetNodeUpDownStreamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetMessage(v string) *GetNodeUpDownStreamResponseBody {
	s.Message = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetNodeDagInfo(v *GetNodeUpDownStreamResponseBodyNodeDagInfo) *GetNodeUpDownStreamResponseBody {
	s.NodeDagInfo = v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetRequestId(v string) *GetNodeUpDownStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBody) SetSuccess(v bool) *GetNodeUpDownStreamResponseBody {
	s.Success = &v
	return s
}

type GetNodeUpDownStreamResponseBodyNodeDagInfo struct {
	DownStreamNodeList []*GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList `json:"DownStreamNodeList,omitempty" xml:"DownStreamNodeList,omitempty" type:"Repeated"`
	StartNodeList      []*GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList      `json:"StartNodeList,omitempty" xml:"StartNodeList,omitempty" type:"Repeated"`
	UpStreamNodeList   []*GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList   `json:"UpStreamNodeList,omitempty" xml:"UpStreamNodeList,omitempty" type:"Repeated"`
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfo) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfo) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) SetDownStreamNodeList(v []*GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) *GetNodeUpDownStreamResponseBodyNodeDagInfo {
	s.DownStreamNodeList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) SetStartNodeList(v []*GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) *GetNodeUpDownStreamResponseBodyNodeDagInfo {
	s.StartNodeList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfo) SetUpStreamNodeList(v []*GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) *GetNodeUpDownStreamResponseBodyNodeDagInfo {
	s.UpStreamNodeList = v
	return s
}

type GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList struct {
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) SetFieldIdList(v []*string) *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	s.FieldIdList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) SetId(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	s.Id = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) SetName(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	s.Name = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList) SetType(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoDownStreamNodeList {
	s.Type = &v
	return s
}

type GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList struct {
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) SetFieldIdList(v []*string) *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	s.FieldIdList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) SetId(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	s.Id = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) SetName(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	s.Name = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList) SetType(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoStartNodeList {
	s.Type = &v
	return s
}

type GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList struct {
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) SetFieldIdList(v []*string) *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	s.FieldIdList = v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) SetId(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	s.Id = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) SetName(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	s.Name = &v
	return s
}

func (s *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList) SetType(v string) *GetNodeUpDownStreamResponseBodyNodeDagInfoUpStreamNodeList {
	s.Type = &v
	return s
}

type GetNodeUpDownStreamResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeUpDownStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeUpDownStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeUpDownStreamResponse) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponse) SetHeaders(v map[string]*string) *GetNodeUpDownStreamResponse {
	s.Headers = v
	return s
}

func (s *GetNodeUpDownStreamResponse) SetStatusCode(v int32) *GetNodeUpDownStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeUpDownStreamResponse) SetBody(v *GetNodeUpDownStreamResponseBody) *GetNodeUpDownStreamResponse {
	s.Body = v
	return s
}

type GetOperationSubmitStatusRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1324444131
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetOperationSubmitStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOperationSubmitStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOperationSubmitStatusRequest) SetEnv(v string) *GetOperationSubmitStatusRequest {
	s.Env = &v
	return s
}

func (s *GetOperationSubmitStatusRequest) SetJobId(v string) *GetOperationSubmitStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetOperationSubmitStatusRequest) SetOpTenantId(v int64) *GetOperationSubmitStatusRequest {
	s.OpTenantId = &v
	return s
}

type GetOperationSubmitStatusResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message            *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	OperationSubmitJob *GetOperationSubmitStatusResponseBodyOperationSubmitJob `json:"OperationSubmitJob,omitempty" xml:"OperationSubmitJob,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOperationSubmitStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOperationSubmitStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationSubmitStatusResponseBody) SetCode(v string) *GetOperationSubmitStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetHttpStatusCode(v int32) *GetOperationSubmitStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetMessage(v string) *GetOperationSubmitStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetOperationSubmitJob(v *GetOperationSubmitStatusResponseBodyOperationSubmitJob) *GetOperationSubmitStatusResponseBody {
	s.OperationSubmitJob = v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetRequestId(v string) *GetOperationSubmitStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBody) SetSuccess(v bool) *GetOperationSubmitStatusResponseBody {
	s.Success = &v
	return s
}

type GetOperationSubmitStatusResponseBodyOperationSubmitJob struct {
	// example:
	//
	// f_122_232342
	ExternalBizId *string `json:"ExternalBizId,omitempty" xml:"ExternalBizId,omitempty"`
	// example:
	//
	// 123456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// SUPPLY_DATA
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// SUCCESS
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// example:
	//
	// 132344
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 80
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s GetOperationSubmitStatusResponseBodyOperationSubmitJob) String() string {
	return tea.Prettify(s)
}

func (s GetOperationSubmitStatusResponseBodyOperationSubmitJob) GoString() string {
	return s.String()
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetExternalBizId(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.ExternalBizId = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetJobId(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.JobId = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetOperation(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.Operation = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetOperationStatus(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.OperationStatus = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetOperator(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.Operator = &v
	return s
}

func (s *GetOperationSubmitStatusResponseBodyOperationSubmitJob) SetProgress(v string) *GetOperationSubmitStatusResponseBodyOperationSubmitJob {
	s.Progress = &v
	return s
}

type GetOperationSubmitStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationSubmitStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationSubmitStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOperationSubmitStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOperationSubmitStatusResponse) SetHeaders(v map[string]*string) *GetOperationSubmitStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOperationSubmitStatusResponse) SetStatusCode(v int32) *GetOperationSubmitStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationSubmitStatusResponse) SetBody(v *GetOperationSubmitStatusResponseBody) *GetOperationSubmitStatusResponse {
	s.Body = v
	return s
}

type GetPhysicalInstanceRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_23231
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2323131
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetPhysicalInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceRequest) SetEnv(v string) *GetPhysicalInstanceRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalInstanceRequest) SetInstanceId(v string) *GetPhysicalInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPhysicalInstanceRequest) SetOpTenantId(v int64) *GetPhysicalInstanceRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalInstanceRequest) SetProjectId(v int64) *GetPhysicalInstanceRequest {
	s.ProjectId = &v
	return s
}

type GetPhysicalInstanceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Instance       *GetPhysicalInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBody) SetCode(v string) *GetPhysicalInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetHttpStatusCode(v int32) *GetPhysicalInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetInstance(v *GetPhysicalInstanceResponseBodyInstance) *GetPhysicalInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetMessage(v string) *GetPhysicalInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetRequestId(v string) *GetPhysicalInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalInstanceResponseBody) SetSuccess(v bool) *GetPhysicalInstanceResponseBody {
	s.Success = &v
	return s
}

type GetPhysicalInstanceResponseBodyInstance struct {
	// example:
	//
	// 2023-06-25
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 2023-06-27 00:30:00
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// example:
	//
	// 3600s
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2023-06-27 02:30:00
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// example:
	//
	// xx
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// t_23231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Index    *int32                                           `json:"Index,omitempty" xml:"Index,omitempty"`
	NodeInfo *GetPhysicalInstanceResponseBodyInstanceNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2023-06-27 01:30:00
	StartExecuteTime *int64    `json:"StartExecuteTime,omitempty" xml:"StartExecuteTime,omitempty"`
	StatusList       []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s GetPhysicalInstanceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetBizDate(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.BizDate = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetDueTime(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.DueTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetDuration(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.Duration = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetEndExecuteTime(v int64) *GetPhysicalInstanceResponseBodyInstance {
	s.EndExecuteTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetExtendInfo(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.ExtendInfo = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetId(v string) *GetPhysicalInstanceResponseBodyInstance {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetIndex(v int32) *GetPhysicalInstanceResponseBodyInstance {
	s.Index = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetNodeInfo(v *GetPhysicalInstanceResponseBodyInstanceNodeInfo) *GetPhysicalInstanceResponseBodyInstance {
	s.NodeInfo = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetStartExecuteTime(v int64) *GetPhysicalInstanceResponseBodyInstance {
	s.StartExecuteTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstance) SetStatusList(v []*string) *GetPhysicalInstanceResponseBodyInstance {
	s.StatusList = v
	return s
}

type GetPhysicalInstanceResponseBodyInstanceNodeInfo struct {
	// example:
	//
	// xx
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 2023-02-02 23:53:17
	CreateTime  *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator     *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	Description *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// true
	HasDev *bool `json:"HasDev,omitempty" xml:"HasDev,omitempty"`
	// example:
	//
	// true
	HasProd *bool `json:"HasProd,omitempty" xml:"HasProd,omitempty"`
	// example:
	//
	// n_3232312
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-02-02 23:53:17
	LastModifiedTime  *string                                                     `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier          *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier    `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name              *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList         []*GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	PriorityList      []*string                                                   `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	ResourceGroupList []*string                                                   `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// SHELL
	SubDetailType *string `json:"SubDetailType,omitempty" xml:"SubDetailType,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetBizUnitName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetCreateTime(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetCreator(v *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Creator = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetDescription(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Description = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetDryRun(v bool) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.DryRun = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetFrom(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.From = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetHasDev(v bool) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.HasDev = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetHasProd(v bool) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.HasProd = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetId(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetLastModifiedTime(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetModifier(v *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Modifier = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Name = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetOwnerList(v []*GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.OwnerList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetPriorityList(v []*string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.PriorityList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetResourceGroupList(v []*string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.ResourceGroupList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetSchedulePaused(v bool) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.SchedulePaused = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetSchedulePeriodList(v []*string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.SchedulePeriodList = v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetSubDetailType(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.SubDetailType = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfo) SetType(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfo {
	s.Type = &v
	return s
}

type GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator struct {
	// example:
	//
	// 2323111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) SetId(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator) SetName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoCreator {
	s.Name = &v
	return s
}

type GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier struct {
	// example:
	//
	// 2323111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) SetId(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier) SetName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoModifier {
	s.Name = &v
	return s
}

type GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList struct {
	// example:
	//
	// 2323111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) SetId(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList {
	s.Id = &v
	return s
}

func (s *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList) SetName(v string) *GetPhysicalInstanceResponseBodyInstanceNodeInfoOwnerList {
	s.Name = &v
	return s
}

type GetPhysicalInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceResponse) SetHeaders(v map[string]*string) *GetPhysicalInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalInstanceResponse) SetStatusCode(v int32) *GetPhysicalInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalInstanceResponse) SetBody(v *GetPhysicalInstanceResponseBody) *GetPhysicalInstanceResponse {
	s.Body = v
	return s
}

type GetPhysicalInstanceLogRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_5929472_20210411_9577721
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123131
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetPhysicalInstanceLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceLogRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceLogRequest) SetEnv(v string) *GetPhysicalInstanceLogRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalInstanceLogRequest) SetInstanceId(v string) *GetPhysicalInstanceLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPhysicalInstanceLogRequest) SetOpTenantId(v int64) *GetPhysicalInstanceLogRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalInstanceLogRequest) SetProjectId(v int64) *GetPhysicalInstanceLogRequest {
	s.ProjectId = &v
	return s
}

type GetPhysicalInstanceLogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success        *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskrunLogList []*GetPhysicalInstanceLogResponseBodyTaskrunLogList `json:"TaskrunLogList,omitempty" xml:"TaskrunLogList,omitempty" type:"Repeated"`
}

func (s GetPhysicalInstanceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceLogResponseBody) SetCode(v string) *GetPhysicalInstanceLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetHttpStatusCode(v int32) *GetPhysicalInstanceLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetMessage(v string) *GetPhysicalInstanceLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetRequestId(v string) *GetPhysicalInstanceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetSuccess(v bool) *GetPhysicalInstanceLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBody) SetTaskrunLogList(v []*GetPhysicalInstanceLogResponseBodyTaskrunLogList) *GetPhysicalInstanceLogResponseBody {
	s.TaskrunLogList = v
	return s
}

type GetPhysicalInstanceLogResponseBodyTaskrunLogList struct {
	// example:
	//
	// 60s
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-05-30 16:48:13
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xx
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// tr_23231
	TaskrunId *string `json:"TaskrunId,omitempty" xml:"TaskrunId,omitempty"`
}

func (s GetPhysicalInstanceLogResponseBodyTaskrunLogList) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceLogResponseBodyTaskrunLogList) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetDuration(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.Duration = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetEndTime(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.EndTime = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetLogContent(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.LogContent = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetStartTime(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.StartTime = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetStatus(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.Status = &v
	return s
}

func (s *GetPhysicalInstanceLogResponseBodyTaskrunLogList) SetTaskrunId(v string) *GetPhysicalInstanceLogResponseBodyTaskrunLogList {
	s.TaskrunId = &v
	return s
}

type GetPhysicalInstanceLogResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalInstanceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalInstanceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalInstanceLogResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalInstanceLogResponse) SetHeaders(v map[string]*string) *GetPhysicalInstanceLogResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalInstanceLogResponse) SetStatusCode(v int32) *GetPhysicalInstanceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalInstanceLogResponse) SetBody(v *GetPhysicalInstanceLogResponseBody) *GetPhysicalInstanceLogResponse {
	s.Body = v
	return s
}

type GetPhysicalNodeRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// n_232132
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetPhysicalNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeRequest) SetEnv(v string) *GetPhysicalNodeRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalNodeRequest) SetNodeId(v string) *GetPhysicalNodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetPhysicalNodeRequest) SetOpTenantId(v int64) *GetPhysicalNodeRequest {
	s.OpTenantId = &v
	return s
}

type GetPhysicalNodeResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeInfo *GetPhysicalNodeResponseBodyNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBody) SetCode(v string) *GetPhysicalNodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetHttpStatusCode(v int32) *GetPhysicalNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetMessage(v string) *GetPhysicalNodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetNodeInfo(v *GetPhysicalNodeResponseBodyNodeInfo) *GetPhysicalNodeResponseBody {
	s.NodeInfo = v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetRequestId(v string) *GetPhysicalNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalNodeResponseBody) SetSuccess(v bool) *GetPhysicalNodeResponseBody {
	s.Success = &v
	return s
}

type GetPhysicalNodeResponseBodyNodeInfo struct {
	// example:
	//
	// 1717343597000
	CreateTime *int64                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator    *GetPhysicalNodeResponseBodyNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// 0 0 10 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// 123456789
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// null
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// n_232132
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1717343597000
	LastModifiedTime *int64                                       `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier         *GetPhysicalNodeResponseBodyNodeInfoModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name             *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SHELL
	OperatorType   *string                                   `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	OutputNameList []*string                                 `json:"OutputNameList,omitempty" xml:"OutputNameList,omitempty" type:"Repeated"`
	Owner          *GetPhysicalNodeResponseBodyNodeInfoOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// example:
	//
	// MIDDLE
	Priority    *string                                         `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectInfo *GetPhysicalNodeResponseBodyNodeInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// example:
	//
	// DAILY
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {"expression":"any_success"}
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetCreateTime(v int64) *GetPhysicalNodeResponseBodyNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetCreator(v *GetPhysicalNodeResponseBodyNodeInfoCreator) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Creator = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetCronExpression(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.CronExpression = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetDataSourceId(v int64) *GetPhysicalNodeResponseBodyNodeInfo {
	s.DataSourceId = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetDataSourceSchema(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.DataSourceSchema = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetDescription(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Description = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetFrom(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.From = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetLastModifiedTime(v int64) *GetPhysicalNodeResponseBodyNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetModifier(v *GetPhysicalNodeResponseBodyNodeInfoModifier) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Modifier = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetOperatorType(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.OperatorType = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetOutputNameList(v []*string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.OutputNameList = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetOwner(v *GetPhysicalNodeResponseBodyNodeInfoOwner) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Owner = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetPriority(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Priority = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetProjectInfo(v *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) *GetPhysicalNodeResponseBodyNodeInfo {
	s.ProjectInfo = v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetScheduleType(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.ScheduleType = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetStatus(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.Status = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfo) SetTriggerConfig(v string) *GetPhysicalNodeResponseBodyNodeInfo {
	s.TriggerConfig = &v
	return s
}

type GetPhysicalNodeResponseBodyNodeInfoCreator struct {
	// example:
	//
	// 101312
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfoCreator) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfoCreator) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoCreator) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfoCreator {
	s.Name = &v
	return s
}

type GetPhysicalNodeResponseBodyNodeInfoModifier struct {
	// example:
	//
	// 101312
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfoModifier) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfoModifier) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoModifier) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfoModifier {
	s.Name = &v
	return s
}

type GetPhysicalNodeResponseBodyNodeInfoOwner struct {
	// example:
	//
	// 101312
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfoOwner) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfoOwner) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfoOwner) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfoOwner {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoOwner) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfoOwner {
	s.Name = &v
	return s
}

type GetPhysicalNodeResponseBodyNodeInfoProjectInfo struct {
	// example:
	//
	// 102132
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeResponseBodyNodeInfoProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeResponseBodyNodeInfoProjectInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) SetId(v string) *GetPhysicalNodeResponseBodyNodeInfoProjectInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeResponseBodyNodeInfoProjectInfo) SetName(v string) *GetPhysicalNodeResponseBodyNodeInfoProjectInfo {
	s.Name = &v
	return s
}

type GetPhysicalNodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponse) SetHeaders(v map[string]*string) *GetPhysicalNodeResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalNodeResponse) SetStatusCode(v int32) *GetPhysicalNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalNodeResponse) SetBody(v *GetPhysicalNodeResponseBody) *GetPhysicalNodeResponse {
	s.Body = v
	return s
}

type GetPhysicalNodeByOutputNameRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_test
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
}

func (s GetPhysicalNodeByOutputNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameRequest) SetEnv(v string) *GetPhysicalNodeByOutputNameRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameRequest) SetOpTenantId(v int64) *GetPhysicalNodeByOutputNameRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameRequest) SetOutputName(v string) *GetPhysicalNodeByOutputNameRequest {
	s.OutputName = &v
	return s
}

type GetPhysicalNodeByOutputNameResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeInfo *GetPhysicalNodeByOutputNameResponseBodyNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetCode(v string) *GetPhysicalNodeByOutputNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetHttpStatusCode(v int32) *GetPhysicalNodeByOutputNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetMessage(v string) *GetPhysicalNodeByOutputNameResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetNodeInfo(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) *GetPhysicalNodeByOutputNameResponseBody {
	s.NodeInfo = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetRequestId(v string) *GetPhysicalNodeByOutputNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBody) SetSuccess(v bool) *GetPhysicalNodeByOutputNameResponseBody {
	s.Success = &v
	return s
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfo struct {
	// example:
	//
	// 1717343597000
	CreateTime  *int64                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator     *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	Description *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// n_2321
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1717343597000
	LastModifiedTime *int64                                                   `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier         *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name             *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SHELL
	OperatorType *string                                               `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	Owner        *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// example:
	//
	// MIDDLE
	Priority    *string                                                     `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectInfo *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// example:
	//
	// DAILY
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {"expression":"any_success"}
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetCreateTime(v int64) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetCreator(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Creator = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetDescription(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Description = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetFrom(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.From = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetLastModifiedTime(v int64) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetModifier(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Modifier = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Name = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetOperatorType(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.OperatorType = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetOwner(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Owner = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetPriority(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Priority = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetProjectInfo(v *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.ProjectInfo = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetScheduleType(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.ScheduleType = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetStatus(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.Status = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfo) SetTriggerConfig(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfo {
	s.TriggerConfig = &v
	return s
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator struct {
	// example:
	//
	// 1311131
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoCreator {
	s.Name = &v
	return s
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier struct {
	// example:
	//
	// 1311131
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoModifier {
	s.Name = &v
	return s
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner struct {
	// example:
	//
	// 1311131
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoOwner {
	s.Name = &v
	return s
}

type GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo struct {
	// example:
	//
	// 1324211
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) SetId(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo {
	s.Id = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo) SetName(v string) *GetPhysicalNodeByOutputNameResponseBodyNodeInfoProjectInfo {
	s.Name = &v
	return s
}

type GetPhysicalNodeByOutputNameResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalNodeByOutputNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponse) SetHeaders(v map[string]*string) *GetPhysicalNodeByOutputNameResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponse) SetStatusCode(v int32) *GetPhysicalNodeByOutputNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponse) SetBody(v *GetPhysicalNodeByOutputNameResponseBody) *GetPhysicalNodeByOutputNameResponse {
	s.Body = v
	return s
}

type GetPhysicalNodeContentRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// n_232411
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetPhysicalNodeContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeContentRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeContentRequest) SetEnv(v string) *GetPhysicalNodeContentRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalNodeContentRequest) SetNodeId(v string) *GetPhysicalNodeContentRequest {
	s.NodeId = &v
	return s
}

func (s *GetPhysicalNodeContentRequest) SetOpTenantId(v int64) *GetPhysicalNodeContentRequest {
	s.OpTenantId = &v
	return s
}

type GetPhysicalNodeContentResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPhysicalNodeContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalNodeContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeContentResponseBody) SetCode(v string) *GetPhysicalNodeContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetData(v *GetPhysicalNodeContentResponseBodyData) *GetPhysicalNodeContentResponseBody {
	s.Data = v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetHttpStatusCode(v int32) *GetPhysicalNodeContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetMessage(v string) *GetPhysicalNodeContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetRequestId(v string) *GetPhysicalNodeContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBody) SetSuccess(v bool) *GetPhysicalNodeContentResponseBody {
	s.Success = &v
	return s
}

type GetPhysicalNodeContentResponseBodyData struct {
	// example:
	//
	// select 1;
	CodeContent *string `json:"CodeContent,omitempty" xml:"CodeContent,omitempty"`
	// example:
	//
	// n_232411
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetPhysicalNodeContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeContentResponseBodyData) SetCodeContent(v string) *GetPhysicalNodeContentResponseBodyData {
	s.CodeContent = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBodyData) SetNodeId(v string) *GetPhysicalNodeContentResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetPhysicalNodeContentResponseBodyData) SetNodeName(v string) *GetPhysicalNodeContentResponseBodyData {
	s.NodeName = &v
	return s
}

type GetPhysicalNodeContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalNodeContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalNodeContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeContentResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeContentResponse) SetHeaders(v map[string]*string) *GetPhysicalNodeContentResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalNodeContentResponse) SetStatusCode(v int32) *GetPhysicalNodeContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalNodeContentResponse) SetBody(v *GetPhysicalNodeContentResponseBody) *GetPhysicalNodeContentResponse {
	s.Body = v
	return s
}

type GetPhysicalNodeOperationLogRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// n_231131
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetPhysicalNodeOperationLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeOperationLogRequest) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeOperationLogRequest) SetEnv(v string) *GetPhysicalNodeOperationLogRequest {
	s.Env = &v
	return s
}

func (s *GetPhysicalNodeOperationLogRequest) SetNodeId(v string) *GetPhysicalNodeOperationLogRequest {
	s.NodeId = &v
	return s
}

func (s *GetPhysicalNodeOperationLogRequest) SetOpTenantId(v int64) *GetPhysicalNodeOperationLogRequest {
	s.OpTenantId = &v
	return s
}

type GetPhysicalNodeOperationLogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message          *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	OperationLogList []*GetPhysicalNodeOperationLogResponseBodyOperationLogList `json:"OperationLogList,omitempty" xml:"OperationLogList,omitempty" type:"Repeated"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPhysicalNodeOperationLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeOperationLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetCode(v string) *GetPhysicalNodeOperationLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetHttpStatusCode(v int32) *GetPhysicalNodeOperationLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetMessage(v string) *GetPhysicalNodeOperationLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetOperationLogList(v []*GetPhysicalNodeOperationLogResponseBodyOperationLogList) *GetPhysicalNodeOperationLogResponseBody {
	s.OperationLogList = v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetRequestId(v string) *GetPhysicalNodeOperationLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBody) SetSuccess(v bool) *GetPhysicalNodeOperationLogResponseBody {
	s.Success = &v
	return s
}

type GetPhysicalNodeOperationLogResponseBodyOperationLogList struct {
	// example:
	//
	// xx
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// example:
	//
	// PAUSE_TASK
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// 132222
	Operator     *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s GetPhysicalNodeOperationLogResponseBodyOperationLogList) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeOperationLogResponseBodyOperationLogList) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetContext(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.Context = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetOperationTime(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.OperationTime = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetOperationType(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.OperationType = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetOperator(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.Operator = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponseBodyOperationLogList) SetOperatorName(v string) *GetPhysicalNodeOperationLogResponseBodyOperationLogList {
	s.OperatorName = &v
	return s
}

type GetPhysicalNodeOperationLogResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalNodeOperationLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalNodeOperationLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhysicalNodeOperationLogResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeOperationLogResponse) SetHeaders(v map[string]*string) *GetPhysicalNodeOperationLogResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalNodeOperationLogResponse) SetStatusCode(v int32) *GetPhysicalNodeOperationLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalNodeOperationLogResponse) SetBody(v *GetPhysicalNodeOperationLogResponseBody) *GetPhysicalNodeOperationLogResponse {
	s.Body = v
	return s
}

type GetProjectProduceUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 131311111321
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectProduceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectProduceUserRequest) GoString() string {
	return s.String()
}

func (s *GetProjectProduceUserRequest) SetOpTenantId(v int64) *GetProjectProduceUserRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetProjectProduceUserRequest) SetProjectId(v int64) *GetProjectProduceUserRequest {
	s.ProjectId = &v
	return s
}

type GetProjectProduceUserResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	User    *GetProjectProduceUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetProjectProduceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectProduceUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectProduceUserResponseBody) SetCode(v string) *GetProjectProduceUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetHttpStatusCode(v int32) *GetProjectProduceUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetMessage(v string) *GetProjectProduceUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetRequestId(v string) *GetProjectProduceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetSuccess(v bool) *GetProjectProduceUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetUser(v *GetProjectProduceUserResponseBodyUser) *GetProjectProduceUserResponseBody {
	s.User = v
	return s
}

type GetProjectProduceUserResponseBodyUser struct {
	// example:
	//
	// 123131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProjectProduceUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetProjectProduceUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetProjectProduceUserResponseBodyUser) SetId(v string) *GetProjectProduceUserResponseBodyUser {
	s.Id = &v
	return s
}

type GetProjectProduceUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectProduceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectProduceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectProduceUserResponse) GoString() string {
	return s.String()
}

func (s *GetProjectProduceUserResponse) SetHeaders(v map[string]*string) *GetProjectProduceUserResponse {
	s.Headers = v
	return s
}

func (s *GetProjectProduceUserResponse) SetStatusCode(v int32) *GetProjectProduceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectProduceUserResponse) SetBody(v *GetProjectProduceUserResponseBody) *GetProjectProduceUserResponse {
	s.Body = v
	return s
}

type GetSupplementDagrunRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f_8241792_20201202_2099680
	SupplementId *string `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
}

func (s GetSupplementDagrunRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunRequest) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunRequest) SetEnv(v string) *GetSupplementDagrunRequest {
	s.Env = &v
	return s
}

func (s *GetSupplementDagrunRequest) SetOpTenantId(v int64) *GetSupplementDagrunRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetSupplementDagrunRequest) SetSupplementId(v string) *GetSupplementDagrunRequest {
	s.SupplementId = &v
	return s
}

type GetSupplementDagrunResponseBody struct {
	// example:
	//
	// OK
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	DagrunList []*GetSupplementDagrunResponseBodyDagrunList `json:"DagrunList,omitempty" xml:"DagrunList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSupplementDagrunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunResponseBody) SetCode(v string) *GetSupplementDagrunResponseBody {
	s.Code = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetDagrunList(v []*GetSupplementDagrunResponseBodyDagrunList) *GetSupplementDagrunResponseBody {
	s.DagrunList = v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetHttpStatusCode(v int32) *GetSupplementDagrunResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetMessage(v string) *GetSupplementDagrunResponseBody {
	s.Message = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetRequestId(v string) *GetSupplementDagrunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupplementDagrunResponseBody) SetSuccess(v bool) *GetSupplementDagrunResponseBody {
	s.Success = &v
	return s
}

type GetSupplementDagrunResponseBodyDagrunList struct {
	// example:
	//
	// 2024-04-01
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 60s
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1717081789000
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// Dagrun ID
	//
	// example:
	//
	// dr_2242792_14542
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1717081729000
	StartExecuteTime *int64 `json:"StartExecuteTime,omitempty" xml:"StartExecuteTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// f_8241792_20201202_2099680
	SupplementId *string `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
}

func (s GetSupplementDagrunResponseBodyDagrunList) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunResponseBodyDagrunList) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetBizDate(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.BizDate = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetDuration(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.Duration = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetEndExecuteTime(v int64) *GetSupplementDagrunResponseBodyDagrunList {
	s.EndExecuteTime = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetId(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetStartExecuteTime(v int64) *GetSupplementDagrunResponseBodyDagrunList {
	s.StartExecuteTime = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetStatus(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.Status = &v
	return s
}

func (s *GetSupplementDagrunResponseBodyDagrunList) SetSupplementId(v string) *GetSupplementDagrunResponseBodyDagrunList {
	s.SupplementId = &v
	return s
}

type GetSupplementDagrunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupplementDagrunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupplementDagrunResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunResponse) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunResponse) SetHeaders(v map[string]*string) *GetSupplementDagrunResponse {
	s.Headers = v
	return s
}

func (s *GetSupplementDagrunResponse) SetStatusCode(v int32) *GetSupplementDagrunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupplementDagrunResponse) SetBody(v *GetSupplementDagrunResponseBody) *GetSupplementDagrunResponse {
	s.Body = v
	return s
}

type GetSupplementDagrunInstanceRequest struct {
	// Dagrun ID
	//
	// This parameter is required.
	//
	// example:
	//
	// dr_2242792_14542
	DagrunId *string `json:"DagrunId,omitempty" xml:"DagrunId,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetSupplementDagrunInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceRequest) SetDagrunId(v string) *GetSupplementDagrunInstanceRequest {
	s.DagrunId = &v
	return s
}

func (s *GetSupplementDagrunInstanceRequest) SetEnv(v string) *GetSupplementDagrunInstanceRequest {
	s.Env = &v
	return s
}

func (s *GetSupplementDagrunInstanceRequest) SetOpTenantId(v int64) *GetSupplementDagrunInstanceRequest {
	s.OpTenantId = &v
	return s
}

type GetSupplementDagrunInstanceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceList   []*GetSupplementDagrunInstanceResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBody) SetCode(v string) *GetSupplementDagrunInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetHttpStatusCode(v int32) *GetSupplementDagrunInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetInstanceList(v []*GetSupplementDagrunInstanceResponseBodyInstanceList) *GetSupplementDagrunInstanceResponseBody {
	s.InstanceList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetMessage(v string) *GetSupplementDagrunInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetRequestId(v string) *GetSupplementDagrunInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetSuccess(v bool) *GetSupplementDagrunInstanceResponseBody {
	s.Success = &v
	return s
}

type GetSupplementDagrunInstanceResponseBodyInstanceList struct {
	// example:
	//
	// 2024-04-01
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 2024-04-02
	DueTime *int64 `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// example:
	//
	// 60
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-04-12 00:25:02
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// example:
	//
	// {"a":"b"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// t_239496_20210411_246982077481
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Index    *int32                                                       `json:"Index,omitempty" xml:"Index,omitempty"`
	NodeInfo *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2024-04-12 00:00:00
	StartExecuteTime *int64    `json:"StartExecuteTime,omitempty" xml:"StartExecuteTime,omitempty"`
	StatusList       []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// example:
	//
	// SUPPLEMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetBizDate(v int64) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.BizDate = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetDueTime(v int64) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.DueTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetDuration(v string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.Duration = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetEndExecuteTime(v int64) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.EndExecuteTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetExtendInfo(v string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.ExtendInfo = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetIndex(v int32) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.Index = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetNodeInfo(v *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.NodeInfo = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetStartExecuteTime(v int64) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.StartExecuteTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetStatusList(v []*string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.StatusList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetType(v string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.Type = &v
	return s
}

type GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo struct {
	// example:
	//
	// xx测试
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 2024-01-30 10:08:49
	CreateTime *string                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator    *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// xx测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// true
	HasDev *bool `json:"HasDev,omitempty" xml:"HasDev,omitempty"`
	// example:
	//
	// true
	HasProd *bool `json:"HasProd,omitempty" xml:"HasProd,omitempty"`
	// example:
	//
	// n_239496
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-01-30 10:08:49
	LastModifiedTime *string                                                              `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier         *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	// example:
	//
	// xx测试
	Name              *string                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList         []*GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	PriorityList      []*string                                                               `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	ResourceGroupList []*string                                                               `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// SHELL
	SubDetailType *string `json:"SubDetailType,omitempty" xml:"SubDetailType,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetBizUnitName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetCreateTime(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetCreator(v *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Creator = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetDescription(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Description = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetDryRun(v bool) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.DryRun = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetFrom(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.From = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetHasDev(v bool) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.HasDev = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetHasProd(v bool) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.HasProd = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetLastModifiedTime(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetModifier(v *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Modifier = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Name = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetOwnerList(v []*GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.OwnerList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetPriorityList(v []*string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.PriorityList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetResourceGroupList(v []*string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.ResourceGroupList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetSchedulePaused(v bool) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.SchedulePaused = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetSchedulePeriodList(v []*string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.SchedulePeriodList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetSubDetailType(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.SubDetailType = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetType(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Type = &v
	return s
}

type GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator struct {
	// example:
	//
	// 1001012
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) SetName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator {
	s.Name = &v
	return s
}

type GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier struct {
	// example:
	//
	// 1001012
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) SetName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier {
	s.Name = &v
	return s
}

type GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList struct {
	// example:
	//
	// 1001012
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) SetName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList {
	s.Name = &v
	return s
}

type GetSupplementDagrunInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupplementDagrunInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupplementDagrunInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponse) SetHeaders(v map[string]*string) *GetSupplementDagrunInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetSupplementDagrunInstanceResponse) SetStatusCode(v int32) *GetSupplementDagrunInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponse) SetBody(v *GetSupplementDagrunInstanceResponseBody) *GetSupplementDagrunInstanceResponse {
	s.Body = v
	return s
}

type GetUserBySourceIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 323131
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s GetUserBySourceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserBySourceIdRequest) GoString() string {
	return s.String()
}

func (s *GetUserBySourceIdRequest) SetOpTenantId(v int64) *GetUserBySourceIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUserBySourceIdRequest) SetSourceId(v string) *GetUserBySourceIdRequest {
	s.SourceId = &v
	return s
}

type GetUserBySourceIdResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	User    *GetUserBySourceIdResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserBySourceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserBySourceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserBySourceIdResponseBody) SetCode(v string) *GetUserBySourceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetHttpStatusCode(v int32) *GetUserBySourceIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetMessage(v string) *GetUserBySourceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetRequestId(v string) *GetUserBySourceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetSuccess(v bool) *GetUserBySourceIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetUser(v *GetUserBySourceIdResponseBodyUser) *GetUserBySourceIdResponseBody {
	s.User = v
	return s
}

type GetUserBySourceIdResponseBodyUser struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 23231231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserBySourceIdResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserBySourceIdResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserBySourceIdResponseBodyUser) SetDisplayName(v string) *GetUserBySourceIdResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserBySourceIdResponseBodyUser) SetId(v string) *GetUserBySourceIdResponseBodyUser {
	s.Id = &v
	return s
}

type GetUserBySourceIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserBySourceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserBySourceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserBySourceIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserBySourceIdResponse) SetHeaders(v map[string]*string) *GetUserBySourceIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserBySourceIdResponse) SetStatusCode(v int32) *GetUserBySourceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserBySourceIdResponse) SetBody(v *GetUserBySourceIdResponseBody) *GetUserBySourceIdResponse {
	s.Body = v
	return s
}

type GetUserGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1313213
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s GetUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupRequest) SetOpTenantId(v int64) *GetUserGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUserGroupRequest) SetUserGroupId(v string) *GetUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type GetUserGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success       *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	UserGroupInfo *GetUserGroupResponseBodyUserGroupInfo `json:"UserGroupInfo,omitempty" xml:"UserGroupInfo,omitempty" type:"Struct"`
}

func (s GetUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBody) SetCode(v string) *GetUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserGroupResponseBody) SetHttpStatusCode(v int32) *GetUserGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserGroupResponseBody) SetMessage(v string) *GetUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserGroupResponseBody) SetRequestId(v string) *GetUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupResponseBody) SetSuccess(v bool) *GetUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserGroupResponseBody) SetUserGroupInfo(v *GetUserGroupResponseBodyUserGroupInfo) *GetUserGroupResponseBody {
	s.UserGroupInfo = v
	return s
}

type GetUserGroupResponseBodyUserGroupInfo struct {
	// example:
	//
	// true
	Active    *bool                                             `json:"Active,omitempty" xml:"Active,omitempty"`
	AdminList []*GetUserGroupResponseBodyUserGroupInfoAdminList `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1313213
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	MyRole *string `json:"MyRole,omitempty" xml:"MyRole,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroupInfo) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetActive(v bool) *GetUserGroupResponseBodyUserGroupInfo {
	s.Active = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetAdminList(v []*GetUserGroupResponseBodyUserGroupInfoAdminList) *GetUserGroupResponseBodyUserGroupInfo {
	s.AdminList = v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetDescription(v string) *GetUserGroupResponseBodyUserGroupInfo {
	s.Description = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetId(v string) *GetUserGroupResponseBodyUserGroupInfo {
	s.Id = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetMyRole(v string) *GetUserGroupResponseBodyUserGroupInfo {
	s.MyRole = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfo) SetName(v string) *GetUserGroupResponseBodyUserGroupInfo {
	s.Name = &v
	return s
}

type GetUserGroupResponseBodyUserGroupInfoAdminList struct {
	// example:
	//
	// xx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 131313
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroupInfoAdminList) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroupInfoAdminList) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) SetAccountName(v string) *GetUserGroupResponseBodyUserGroupInfoAdminList {
	s.AccountName = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) SetDisplayName(v string) *GetUserGroupResponseBodyUserGroupInfoAdminList {
	s.DisplayName = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupInfoAdminList) SetId(v string) *GetUserGroupResponseBodyUserGroupInfoAdminList {
	s.Id = &v
	return s
}

type GetUserGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponse) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponse) SetHeaders(v map[string]*string) *GetUserGroupResponse {
	s.Headers = v
	return s
}

func (s *GetUserGroupResponse) SetStatusCode(v int32) *GetUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGroupResponse) SetBody(v *GetUserGroupResponseBody) *GetUserGroupResponse {
	s.Body = v
	return s
}

type GetUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s GetUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUsersRequest) GoString() string {
	return s.String()
}

func (s *GetUsersRequest) SetOpTenantId(v int64) *GetUsersRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUsersRequest) SetUserIdList(v []*string) *GetUsersRequest {
	s.UserIdList = v
	return s
}

type GetUsersShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UserIdListShrink *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s GetUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUsersShrinkRequest) SetOpTenantId(v int64) *GetUsersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetUsersShrinkRequest) SetUserIdListShrink(v string) *GetUsersShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

type GetUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success  *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	UserList []*GetUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetUsersResponseBody) SetCode(v string) *GetUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GetUsersResponseBody) SetHttpStatusCode(v int32) *GetUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUsersResponseBody) SetMessage(v string) *GetUsersResponseBody {
	s.Message = &v
	return s
}

func (s *GetUsersResponseBody) SetRequestId(v string) *GetUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUsersResponseBody) SetSuccess(v bool) *GetUsersResponseBody {
	s.Success = &v
	return s
}

func (s *GetUsersResponseBody) SetUserList(v []*GetUsersResponseBodyUserList) *GetUsersResponseBody {
	s.UserList = v
	return s
}

type GetUsersResponseBodyUserList struct {
	// example:
	//
	// 123@xx.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 123@dingding
	DingNumber               *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	DisplayName              *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DisplayNameWithoutStatus *string `json:"DisplayNameWithoutStatus,omitempty" xml:"DisplayNameWithoutStatus,omitempty"`
	// example:
	//
	// true
	EnableWhiteIp *string `json:"EnableWhiteIp,omitempty" xml:"EnableWhiteIp,omitempty"`
	// example:
	//
	// xx
	FeiShuRobot *string `json:"FeiShuRobot,omitempty" xml:"FeiShuRobot,omitempty"`
	// example:
	//
	// 1717343597000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1717343597000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1233121
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123@xx.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 1388888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NickName    *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 231231
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// example:
	//
	// 123@xx.com
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// ALIYUN
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// xx
	WeChatRobot *string `json:"WeChatRobot,omitempty" xml:"WeChatRobot,omitempty"`
	// example:
	//
	// *
	WhiteIp *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
}

func (s GetUsersResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *GetUsersResponseBodyUserList) SetAccountName(v string) *GetUsersResponseBodyUserList {
	s.AccountName = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetDingNumber(v string) *GetUsersResponseBodyUserList {
	s.DingNumber = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetDisplayName(v string) *GetUsersResponseBodyUserList {
	s.DisplayName = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetDisplayNameWithoutStatus(v string) *GetUsersResponseBodyUserList {
	s.DisplayNameWithoutStatus = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetEnableWhiteIp(v string) *GetUsersResponseBodyUserList {
	s.EnableWhiteIp = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetFeiShuRobot(v string) *GetUsersResponseBodyUserList {
	s.FeiShuRobot = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetGmtCreate(v int64) *GetUsersResponseBodyUserList {
	s.GmtCreate = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetGmtModified(v int64) *GetUsersResponseBodyUserList {
	s.GmtModified = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetId(v string) *GetUsersResponseBodyUserList {
	s.Id = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetMail(v string) *GetUsersResponseBodyUserList {
	s.Mail = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetMobilePhone(v string) *GetUsersResponseBodyUserList {
	s.MobilePhone = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetName(v string) *GetUsersResponseBodyUserList {
	s.Name = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetNickName(v string) *GetUsersResponseBodyUserList {
	s.NickName = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetParentId(v string) *GetUsersResponseBodyUserList {
	s.ParentId = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetRealName(v string) *GetUsersResponseBodyUserList {
	s.RealName = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetSourceId(v string) *GetUsersResponseBodyUserList {
	s.SourceId = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetSourceType(v string) *GetUsersResponseBodyUserList {
	s.SourceType = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetWeChatRobot(v string) *GetUsersResponseBodyUserList {
	s.WeChatRobot = &v
	return s
}

func (s *GetUsersResponseBodyUserList) SetWhiteIp(v string) *GetUsersResponseBodyUserList {
	s.WhiteIp = &v
	return s
}

type GetUsersResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponse) GoString() string {
	return s.String()
}

func (s *GetUsersResponse) SetHeaders(v map[string]*string) *GetUsersResponse {
	s.Headers = v
	return s
}

func (s *GetUsersResponse) SetStatusCode(v int32) *GetUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUsersResponse) SetBody(v *GetUsersResponseBody) *GetUsersResponse {
	s.Body = v
	return s
}

type GrantResourcePermissionRequest struct {
	// This parameter is required.
	GrantCommand *GrantResourcePermissionRequestGrantCommand `json:"GrantCommand,omitempty" xml:"GrantCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GrantResourcePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantResourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionRequest) SetGrantCommand(v *GrantResourcePermissionRequestGrantCommand) *GrantResourcePermissionRequest {
	s.GrantCommand = v
	return s
}

func (s *GrantResourcePermissionRequest) SetOpTenantId(v int64) *GrantResourcePermissionRequest {
	s.OpTenantId = &v
	return s
}

type GrantResourcePermissionRequestGrantCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 1717343597000
	EffectiveEnd *string `json:"EffectiveEnd,omitempty" xml:"EffectiveEnd,omitempty"`
	// This parameter is required.
	OperateList []*string `json:"OperateList,omitempty" xml:"OperateList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// This parameter is required.
	ResourceList []*GrantResourcePermissionRequestGrantCommandResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PHYSICAL_TABLE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s GrantResourcePermissionRequestGrantCommand) String() string {
	return tea.Prettify(s)
}

func (s GrantResourcePermissionRequestGrantCommand) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionRequestGrantCommand) SetEffectiveEnd(v string) *GrantResourcePermissionRequestGrantCommand {
	s.EffectiveEnd = &v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetOperateList(v []*string) *GrantResourcePermissionRequestGrantCommand {
	s.OperateList = v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetReason(v string) *GrantResourcePermissionRequestGrantCommand {
	s.Reason = &v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetResourceList(v []*GrantResourcePermissionRequestGrantCommandResourceList) *GrantResourcePermissionRequestGrantCommand {
	s.ResourceList = v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetResourceType(v string) *GrantResourcePermissionRequestGrantCommand {
	s.ResourceType = &v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetUserIdList(v []*string) *GrantResourcePermissionRequestGrantCommand {
	s.UserIdList = v
	return s
}

type GrantResourcePermissionRequestGrantCommandResourceList struct {
	// example:
	//
	// hadoop.300000806.data_distill.behavior_gameinfor_01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GrantResourcePermissionRequestGrantCommandResourceList) String() string {
	return tea.Prettify(s)
}

func (s GrantResourcePermissionRequestGrantCommandResourceList) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionRequestGrantCommandResourceList) SetResourceId(v string) *GrantResourcePermissionRequestGrantCommandResourceList {
	s.ResourceId = &v
	return s
}

type GrantResourcePermissionShrinkRequest struct {
	// This parameter is required.
	GrantCommandShrink *string `json:"GrantCommand,omitempty" xml:"GrantCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GrantResourcePermissionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantResourcePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionShrinkRequest) SetGrantCommandShrink(v string) *GrantResourcePermissionShrinkRequest {
	s.GrantCommandShrink = &v
	return s
}

func (s *GrantResourcePermissionShrinkRequest) SetOpTenantId(v int64) *GrantResourcePermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

type GrantResourcePermissionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantResourcePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantResourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionResponseBody) SetCode(v string) *GrantResourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) SetHttpStatusCode(v int32) *GrantResourcePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) SetMessage(v string) *GrantResourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) SetRequestId(v string) *GrantResourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) SetSuccess(v bool) *GrantResourcePermissionResponseBody {
	s.Success = &v
	return s
}

type GrantResourcePermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantResourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantResourcePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantResourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionResponse) SetHeaders(v map[string]*string) *GrantResourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantResourcePermissionResponse) SetStatusCode(v int32) *GrantResourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantResourcePermissionResponse) SetBody(v *GrantResourcePermissionResponseBody) *GrantResourcePermissionResponse {
	s.Body = v
	return s
}

type ListAddableRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAddableRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAddableRolesRequest) GoString() string {
	return s.String()
}

func (s *ListAddableRolesRequest) SetOpTenantId(v int64) *ListAddableRolesRequest {
	s.OpTenantId = &v
	return s
}

type ListAddableRolesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleList  []*ListAddableRolesResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAddableRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAddableRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddableRolesResponseBody) SetCode(v string) *ListAddableRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAddableRolesResponseBody) SetHttpStatusCode(v int32) *ListAddableRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAddableRolesResponseBody) SetMessage(v string) *ListAddableRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAddableRolesResponseBody) SetRequestId(v string) *ListAddableRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddableRolesResponseBody) SetRoleList(v []*ListAddableRolesResponseBodyRoleList) *ListAddableRolesResponseBody {
	s.RoleList = v
	return s
}

func (s *ListAddableRolesResponseBody) SetSuccess(v bool) *ListAddableRolesResponseBody {
	s.Success = &v
	return s
}

type ListAddableRolesResponseBodyRoleList struct {
	// example:
	//
	// SECURITY_ADMIN
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAddableRolesResponseBodyRoleList) String() string {
	return tea.Prettify(s)
}

func (s ListAddableRolesResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *ListAddableRolesResponseBodyRoleList) SetCode(v string) *ListAddableRolesResponseBodyRoleList {
	s.Code = &v
	return s
}

func (s *ListAddableRolesResponseBodyRoleList) SetName(v string) *ListAddableRolesResponseBodyRoleList {
	s.Name = &v
	return s
}

type ListAddableRolesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddableRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddableRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAddableRolesResponse) GoString() string {
	return s.String()
}

func (s *ListAddableRolesResponse) SetHeaders(v map[string]*string) *ListAddableRolesResponse {
	s.Headers = v
	return s
}

func (s *ListAddableRolesResponse) SetStatusCode(v int32) *ListAddableRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddableRolesResponse) SetBody(v *ListAddableRolesResponseBody) *ListAddableRolesResponse {
	s.Body = v
	return s
}

type ListAddableUsersRequest struct {
	// This parameter is required.
	ListQuery *ListAddableUsersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAddableUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAddableUsersRequest) GoString() string {
	return s.String()
}

func (s *ListAddableUsersRequest) SetListQuery(v *ListAddableUsersRequestListQuery) *ListAddableUsersRequest {
	s.ListQuery = v
	return s
}

func (s *ListAddableUsersRequest) SetOpTenantId(v int64) *ListAddableUsersRequest {
	s.OpTenantId = &v
	return s
}

type ListAddableUsersRequestListQuery struct {
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
}

func (s ListAddableUsersRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListAddableUsersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListAddableUsersRequestListQuery) SetPage(v int32) *ListAddableUsersRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListAddableUsersRequestListQuery) SetPageSize(v int32) *ListAddableUsersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListAddableUsersRequestListQuery) SetSearchText(v string) *ListAddableUsersRequestListQuery {
	s.SearchText = &v
	return s
}

type ListAddableUsersShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListAddableUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAddableUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAddableUsersShrinkRequest) SetListQueryShrink(v string) *ListAddableUsersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListAddableUsersShrinkRequest) SetOpTenantId(v int64) *ListAddableUsersShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListAddableUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListAddableUsersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAddableUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAddableUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddableUsersResponseBody) SetCode(v string) *ListAddableUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListAddableUsersResponseBody) SetHttpStatusCode(v int32) *ListAddableUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAddableUsersResponseBody) SetMessage(v string) *ListAddableUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListAddableUsersResponseBody) SetPageResult(v *ListAddableUsersResponseBodyPageResult) *ListAddableUsersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListAddableUsersResponseBody) SetRequestId(v string) *ListAddableUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddableUsersResponseBody) SetSuccess(v bool) *ListAddableUsersResponseBody {
	s.Success = &v
	return s
}

type ListAddableUsersResponseBodyPageResult struct {
	// example:
	//
	// 66
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserList   []*ListAddableUsersResponseBodyPageResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListAddableUsersResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListAddableUsersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListAddableUsersResponseBodyPageResult) SetTotalCount(v int32) *ListAddableUsersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResult) SetUserList(v []*ListAddableUsersResponseBodyPageResultUserList) *ListAddableUsersResponseBodyPageResult {
	s.UserList = v
	return s
}

type ListAddableUsersResponseBodyPageResultUserList struct {
	// example:
	//
	// 123@xx.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 123@dingding
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// xx
	DisplayNameWithoutStatus *string `json:"DisplayNameWithoutStatus,omitempty" xml:"DisplayNameWithoutStatus,omitempty"`
	// example:
	//
	// true
	EnableWhiteIp *string `json:"EnableWhiteIp,omitempty" xml:"EnableWhiteIp,omitempty"`
	// example:
	//
	// xx
	FeiShuRobot *string `json:"FeiShuRobot,omitempty" xml:"FeiShuRobot,omitempty"`
	// example:
	//
	// 1717343597000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1717343597000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 123231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123@xx.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 13888888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// xx
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 231313
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// xx
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	// example:
	//
	// 123@xx.com
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// aliyun
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// xx
	WeChatRobot *string `json:"WeChatRobot,omitempty" xml:"WeChatRobot,omitempty"`
	// example:
	//
	// *
	WhiteIp *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
}

func (s ListAddableUsersResponseBodyPageResultUserList) String() string {
	return tea.Prettify(s)
}

func (s ListAddableUsersResponseBodyPageResultUserList) GoString() string {
	return s.String()
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetAccountName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.AccountName = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetDingNumber(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.DingNumber = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetDisplayName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.DisplayName = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetDisplayNameWithoutStatus(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.DisplayNameWithoutStatus = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetEnableWhiteIp(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.EnableWhiteIp = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetFeiShuRobot(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.FeiShuRobot = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetGmtCreate(v int64) *ListAddableUsersResponseBodyPageResultUserList {
	s.GmtCreate = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetGmtModified(v int64) *ListAddableUsersResponseBodyPageResultUserList {
	s.GmtModified = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetId(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.Id = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetMail(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.Mail = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetMobilePhone(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.MobilePhone = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.Name = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetNickName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.NickName = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetParentId(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.ParentId = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetRealName(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.RealName = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetSourceId(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.SourceId = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetSourceType(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.SourceType = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetWeChatRobot(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.WeChatRobot = &v
	return s
}

func (s *ListAddableUsersResponseBodyPageResultUserList) SetWhiteIp(v string) *ListAddableUsersResponseBodyPageResultUserList {
	s.WhiteIp = &v
	return s
}

type ListAddableUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddableUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddableUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAddableUsersResponse) GoString() string {
	return s.String()
}

func (s *ListAddableUsersResponse) SetHeaders(v map[string]*string) *ListAddableUsersResponse {
	s.Headers = v
	return s
}

func (s *ListAddableUsersResponse) SetStatusCode(v int32) *ListAddableUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddableUsersResponse) SetBody(v *ListAddableUsersResponseBody) *ListAddableUsersResponse {
	s.Body = v
	return s
}

type ListDataSourceWithConfigRequest struct {
	// This parameter is required.
	ListQuery *ListDataSourceWithConfigRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDataSourceWithConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigRequest) SetListQuery(v *ListDataSourceWithConfigRequestListQuery) *ListDataSourceWithConfigRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataSourceWithConfigRequest) SetOpTenantId(v int64) *ListDataSourceWithConfigRequest {
	s.OpTenantId = &v
	return s
}

type ListDataSourceWithConfigRequestListQuery struct {
	// example:
	//
	// vcns-test
	Name      *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScopeList []*string `json:"ScopeList,omitempty" xml:"ScopeList,omitempty" type:"Repeated"`
	Tag       *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TypeList  []*string `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s ListDataSourceWithConfigRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigRequestListQuery) SetName(v string) *ListDataSourceWithConfigRequestListQuery {
	s.Name = &v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetOwnerList(v []*string) *ListDataSourceWithConfigRequestListQuery {
	s.OwnerList = v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetPage(v int32) *ListDataSourceWithConfigRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetPageSize(v int32) *ListDataSourceWithConfigRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetScopeList(v []*string) *ListDataSourceWithConfigRequestListQuery {
	s.ScopeList = v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetTag(v string) *ListDataSourceWithConfigRequestListQuery {
	s.Tag = &v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetTypeList(v []*string) *ListDataSourceWithConfigRequestListQuery {
	s.TypeList = v
	return s
}

type ListDataSourceWithConfigShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDataSourceWithConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigShrinkRequest) SetListQueryShrink(v string) *ListDataSourceWithConfigShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataSourceWithConfigShrinkRequest) SetOpTenantId(v int64) *ListDataSourceWithConfigShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListDataSourceWithConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListDataSourceWithConfigResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataSourceWithConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBody) SetCode(v string) *ListDataSourceWithConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetHttpStatusCode(v int32) *ListDataSourceWithConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetMessage(v string) *ListDataSourceWithConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetPageResult(v *ListDataSourceWithConfigResponseBodyPageResult) *ListDataSourceWithConfigResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetRequestId(v string) *ListDataSourceWithConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBody) SetSuccess(v bool) *ListDataSourceWithConfigResponseBody {
	s.Success = &v
	return s
}

type ListDataSourceWithConfigResponseBodyPageResult struct {
	DataSourceList []*ListDataSourceWithConfigResponseBodyPageResultDataSourceList `json:"DataSourceList,omitempty" xml:"DataSourceList,omitempty" type:"Repeated"`
	// example:
	//
	// 39
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResult) SetDataSourceList(v []*ListDataSourceWithConfigResponseBodyPageResultDataSourceList) *ListDataSourceWithConfigResponseBodyPageResult {
	s.DataSourceList = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResult) SetTotalCount(v int64) *ListDataSourceWithConfigResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceList struct {
	// 开发环境数据源信息
	DevDataSourceInfo *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo `json:"DevDataSourceInfo,omitempty" xml:"DevDataSourceInfo,omitempty" type:"Struct"`
	// 生产环境数据源
	ProdDataSourceInfo *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo `json:"ProdDataSourceInfo,omitempty" xml:"ProdDataSourceInfo,omitempty" type:"Struct"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceList) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceList) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceList) SetDevDataSourceInfo(v *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) *ListDataSourceWithConfigResponseBodyPageResultDataSourceList {
	s.DevDataSourceInfo = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceList) SetProdDataSourceInfo(v *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) *ListDataSourceWithConfigResponseBodyPageResultDataSourceList {
	s.ProdDataSourceInfo = v
	return s
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo struct {
	ConfigItemList []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// example:
	//
	// 1710209552704
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 212211111
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 12313123131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1710209552704
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 212211111
	Owner     *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// ALL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetConfigItemList(v []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.ConfigItemList = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetCreateTime(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetCreator(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Creator = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetCreatorName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetDescription(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Description = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetEnv(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Env = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetId(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Id = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetModifyTime(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetOwner(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Owner = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetOwnerName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetScope(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Scope = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo) SetType(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfo {
	s.Type = &v
	return s
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList struct {
	// example:
	//
	// param1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) SetKey(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList {
	s.Key = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList) SetValue(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListDevDataSourceInfoConfigItemList {
	s.Value = &v
	return s
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo struct {
	ConfigItemList []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// example:
	//
	// 1708303959000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 212211111
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 300000028799
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1708303959000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 212211111
	Owner     *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// ALL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetConfigItemList(v []*ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.ConfigItemList = v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetCreateTime(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetCreator(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Creator = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetCreatorName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.CreatorName = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetDescription(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Description = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetEnv(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Env = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetId(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Id = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetModifyTime(v int64) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetOwner(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Owner = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetOwnerName(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.OwnerName = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetScope(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Scope = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo) SetType(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfo {
	s.Type = &v
	return s
}

type ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList struct {
	// example:
	//
	// param1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) SetKey(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList {
	s.Key = &v
	return s
}

func (s *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList) SetValue(v string) *ListDataSourceWithConfigResponseBodyPageResultDataSourceListProdDataSourceInfoConfigItemList {
	s.Value = &v
	return s
}

type ListDataSourceWithConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceWithConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceWithConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceWithConfigResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigResponse) SetHeaders(v map[string]*string) *ListDataSourceWithConfigResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceWithConfigResponse) SetStatusCode(v int32) *ListDataSourceWithConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceWithConfigResponse) SetBody(v *ListDataSourceWithConfigResponseBody) *ListDataSourceWithConfigResponse {
	s.Body = v
	return s
}

type ListFilesRequest struct {
	// This parameter is required.
	ListQuery *ListFilesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFilesRequest) SetListQuery(v *ListFilesRequestListQuery) *ListFilesRequest {
	s.ListQuery = v
	return s
}

func (s *ListFilesRequest) SetOpTenantId(v int64) *ListFilesRequest {
	s.OpTenantId = &v
	return s
}

type ListFilesRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// tempCode
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /xx/x
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11112311
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
}

func (s ListFilesRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListFilesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListFilesRequestListQuery) SetCategory(v string) *ListFilesRequestListQuery {
	s.Category = &v
	return s
}

func (s *ListFilesRequestListQuery) SetDirectory(v string) *ListFilesRequestListQuery {
	s.Directory = &v
	return s
}

func (s *ListFilesRequestListQuery) SetEnv(v string) *ListFilesRequestListQuery {
	s.Env = &v
	return s
}

func (s *ListFilesRequestListQuery) SetProjectId(v int64) *ListFilesRequestListQuery {
	s.ProjectId = &v
	return s
}

func (s *ListFilesRequestListQuery) SetRecursive(v bool) *ListFilesRequestListQuery {
	s.Recursive = &v
	return s
}

type ListFilesShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListFilesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFilesShrinkRequest) SetListQueryShrink(v string) *ListFilesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListFilesShrinkRequest) SetOpTenantId(v int64) *ListFilesShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListFilesResponseBody struct {
	// example:
	//
	// OK
	Code     *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	FileList []*ListFilesResponseBodyFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBody) SetCode(v string) *ListFilesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFilesResponseBody) SetFileList(v []*ListFilesResponseBodyFileList) *ListFilesResponseBody {
	s.FileList = v
	return s
}

func (s *ListFilesResponseBody) SetHttpStatusCode(v int32) *ListFilesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFilesResponseBody) SetMessage(v string) *ListFilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFilesResponseBody) SetRequestId(v string) *ListFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFilesResponseBody) SetSuccess(v bool) *ListFilesResponseBody {
	s.Success = &v
	return s
}

type ListFilesResponseBodyFileList struct {
	// example:
	//
	// tempCode
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// select 1;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1212111
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// /xx/x
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// directory
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// 1717483193830
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1717483193830
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 111231112
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1212111
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 312112121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListFilesResponseBodyFileList) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponseBodyFileList) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBodyFileList) SetCategory(v string) *ListFilesResponseBodyFileList {
	s.Category = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetContent(v string) *ListFilesResponseBodyFileList {
	s.Content = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetCreator(v string) *ListFilesResponseBodyFileList {
	s.Creator = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetDirectory(v string) *ListFilesResponseBodyFileList {
	s.Directory = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetFileType(v string) *ListFilesResponseBodyFileList {
	s.FileType = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetGmtCreate(v int64) *ListFilesResponseBodyFileList {
	s.GmtCreate = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetGmtModified(v int64) *ListFilesResponseBodyFileList {
	s.GmtModified = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetId(v int64) *ListFilesResponseBodyFileList {
	s.Id = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetLastModifier(v string) *ListFilesResponseBodyFileList {
	s.LastModifier = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetName(v string) *ListFilesResponseBodyFileList {
	s.Name = &v
	return s
}

func (s *ListFilesResponseBodyFileList) SetProjectId(v int64) *ListFilesResponseBodyFileList {
	s.ProjectId = &v
	return s
}

type ListFilesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponse) GoString() string {
	return s.String()
}

func (s *ListFilesResponse) SetHeaders(v map[string]*string) *ListFilesResponse {
	s.Headers = v
	return s
}

func (s *ListFilesResponse) SetStatusCode(v int32) *ListFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFilesResponse) SetBody(v *ListFilesResponseBody) *ListFilesResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// example:
	//
	// PROD
	Env       *string                        `json:"Env,omitempty" xml:"Env,omitempty"`
	ListQuery *ListInstancesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetEnv(v string) *ListInstancesRequest {
	s.Env = &v
	return s
}

func (s *ListInstancesRequest) SetListQuery(v *ListInstancesRequestListQuery) *ListInstancesRequest {
	s.ListQuery = v
	return s
}

func (s *ListInstancesRequest) SetOpTenantId(v int64) *ListInstancesRequest {
	s.OpTenantId = &v
	return s
}

type ListInstancesRequestListQuery struct {
	// example:
	//
	// SCRIPT
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizUnitId *int64  `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// 2024-05-31
	MaxBizDate *string `json:"MaxBizDate,omitempty" xml:"MaxBizDate,omitempty"`
	// example:
	//
	// 2024-05-31
	MaxRunDate *string `json:"MaxRunDate,omitempty" xml:"MaxRunDate,omitempty"`
	// example:
	//
	// 2024-05-30
	MinBizDate *string `json:"MinBizDate,omitempty" xml:"MinBizDate,omitempty"`
	// example:
	//
	// 2024-05-30
	MinRunDate *string `json:"MinRunDate,omitempty" xml:"MinRunDate,omitempty"`
	// example:
	//
	// n_23131
	NodeId    *string   `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerList []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PriorityList []*string `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 131311111321
	ProjectId          *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RunStatusList      []*string `json:"RunStatusList,omitempty" xml:"RunStatusList,omitempty" type:"Repeated"`
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// example:
	//
	// xx
	SearchText     *string   `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	SubBizTypeList []*string `json:"SubBizTypeList,omitempty" xml:"SubBizTypeList,omitempty" type:"Repeated"`
}

func (s ListInstancesRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestListQuery) SetBizType(v string) *ListInstancesRequestListQuery {
	s.BizType = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetBizUnitId(v int64) *ListInstancesRequestListQuery {
	s.BizUnitId = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetMaxBizDate(v string) *ListInstancesRequestListQuery {
	s.MaxBizDate = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetMaxRunDate(v string) *ListInstancesRequestListQuery {
	s.MaxRunDate = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetMinBizDate(v string) *ListInstancesRequestListQuery {
	s.MinBizDate = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetMinRunDate(v string) *ListInstancesRequestListQuery {
	s.MinRunDate = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetNodeId(v string) *ListInstancesRequestListQuery {
	s.NodeId = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetOwnerList(v []*string) *ListInstancesRequestListQuery {
	s.OwnerList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetPage(v int32) *ListInstancesRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetPageSize(v int32) *ListInstancesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetPriorityList(v []*string) *ListInstancesRequestListQuery {
	s.PriorityList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetProjectId(v int64) *ListInstancesRequestListQuery {
	s.ProjectId = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetRunStatusList(v []*string) *ListInstancesRequestListQuery {
	s.RunStatusList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetSchedulePaused(v bool) *ListInstancesRequestListQuery {
	s.SchedulePaused = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetSchedulePeriodList(v []*string) *ListInstancesRequestListQuery {
	s.SchedulePeriodList = v
	return s
}

func (s *ListInstancesRequestListQuery) SetScheduleType(v string) *ListInstancesRequestListQuery {
	s.ScheduleType = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetSearchText(v string) *ListInstancesRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListInstancesRequestListQuery) SetSubBizTypeList(v []*string) *ListInstancesRequestListQuery {
	s.SubBizTypeList = v
	return s
}

type ListInstancesShrinkRequest struct {
	// example:
	//
	// PROD
	Env             *string `json:"Env,omitempty" xml:"Env,omitempty"`
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) SetEnv(v string) *ListInstancesShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetListQueryShrink(v string) *ListInstancesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetOpTenantId(v int64) *ListInstancesShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListInstancesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListInstancesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageResult(v *ListInstancesResponseBodyPageResult) *ListInstancesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

type ListInstancesResponseBodyPageResult struct {
	Data []*ListInstancesResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 107
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResult) SetData(v []*ListInstancesResponseBodyPageResultData) *ListInstancesResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBodyPageResult) SetTotalCount(v int32) *ListInstancesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyPageResultData struct {
	// example:
	//
	// 2024-05-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// example:
	//
	// 60s
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// example:
	//
	// xx
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// t_23231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Index    *int32                                           `json:"Index,omitempty" xml:"Index,omitempty"`
	NodeInfo *ListInstancesResponseBodyPageResultDataNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2024-05-30 16:46:13
	StartExecuteTime *int64    `json:"StartExecuteTime,omitempty" xml:"StartExecuteTime,omitempty"`
	StatusList       []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyPageResultData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultData) SetBizDate(v string) *ListInstancesResponseBodyPageResultData {
	s.BizDate = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetDueTime(v string) *ListInstancesResponseBodyPageResultData {
	s.DueTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetDuration(v string) *ListInstancesResponseBodyPageResultData {
	s.Duration = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetEndExecuteTime(v int64) *ListInstancesResponseBodyPageResultData {
	s.EndExecuteTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetExtendInfo(v string) *ListInstancesResponseBodyPageResultData {
	s.ExtendInfo = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetId(v string) *ListInstancesResponseBodyPageResultData {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetIndex(v int32) *ListInstancesResponseBodyPageResultData {
	s.Index = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetNodeInfo(v *ListInstancesResponseBodyPageResultDataNodeInfo) *ListInstancesResponseBodyPageResultData {
	s.NodeInfo = v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetStartExecuteTime(v int64) *ListInstancesResponseBodyPageResultData {
	s.StartExecuteTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetStatusList(v []*string) *ListInstancesResponseBodyPageResultData {
	s.StatusList = v
	return s
}

type ListInstancesResponseBodyPageResultDataNodeInfo struct {
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	CreateTime  *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator     *ListInstancesResponseBodyPageResultDataNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	Description *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// DATA_PROCES
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// false
	HasDev *bool `json:"HasDev,omitempty" xml:"HasDev,omitempty"`
	// example:
	//
	// true
	HasProd *bool `json:"HasProd,omitempty" xml:"HasProd,omitempty"`
	// example:
	//
	// n_132331
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	LastModifiedTime  *string                                                     `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier          *ListInstancesResponseBodyPageResultDataNodeInfoModifier    `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name              *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList         []*ListInstancesResponseBodyPageResultDataNodeInfoOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	PriorityList      []*string                                                   `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	ResourceGroupList []*string                                                   `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// SHELL
	SubDetailType *string `json:"SubDetailType,omitempty" xml:"SubDetailType,omitempty"`
	// example:
	//
	// DATA_PROCES
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyPageResultDataNodeInfo) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultDataNodeInfo) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetBizUnitName(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.BizUnitName = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetCreateTime(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetCreator(v *ListInstancesResponseBodyPageResultDataNodeInfoCreator) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Creator = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetDescription(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetDryRun(v bool) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.DryRun = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetFrom(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.From = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetHasDev(v bool) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.HasDev = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetHasProd(v bool) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.HasProd = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetId(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetLastModifiedTime(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetModifier(v *ListInstancesResponseBodyPageResultDataNodeInfoModifier) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Modifier = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetName(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetOwnerList(v []*ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.OwnerList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetPriorityList(v []*string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.PriorityList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetResourceGroupList(v []*string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.ResourceGroupList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetSchedulePaused(v bool) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.SchedulePaused = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetSchedulePeriodList(v []*string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.SchedulePeriodList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetSubDetailType(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.SubDetailType = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetType(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Type = &v
	return s
}

type ListInstancesResponseBodyPageResultDataNodeInfoCreator struct {
	// example:
	//
	// 21313112
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoCreator) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoCreator) SetId(v string) *ListInstancesResponseBodyPageResultDataNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoCreator) SetName(v string) *ListInstancesResponseBodyPageResultDataNodeInfoCreator {
	s.Name = &v
	return s
}

type ListInstancesResponseBodyPageResultDataNodeInfoModifier struct {
	// example:
	//
	// 21313112
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoModifier) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoModifier) SetId(v string) *ListInstancesResponseBodyPageResultDataNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoModifier) SetName(v string) *ListInstancesResponseBodyPageResultDataNodeInfoModifier {
	s.Name = &v
	return s
}

type ListInstancesResponseBodyPageResultDataNodeInfoOwnerList struct {
	// example:
	//
	// 21313112
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) SetId(v string) *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) SetName(v string) *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList {
	s.Name = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListNodeDownStreamRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ListQuery *ListNodeDownStreamRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListNodeDownStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDownStreamRequest) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamRequest) SetEnv(v string) *ListNodeDownStreamRequest {
	s.Env = &v
	return s
}

func (s *ListNodeDownStreamRequest) SetListQuery(v *ListNodeDownStreamRequestListQuery) *ListNodeDownStreamRequest {
	s.ListQuery = v
	return s
}

func (s *ListNodeDownStreamRequest) SetOpTenantId(v int64) *ListNodeDownStreamRequest {
	s.OpTenantId = &v
	return s
}

type ListNodeDownStreamRequestListQuery struct {
	// example:
	//
	// 1
	DownStreamDepth *int32                                          `json:"DownStreamDepth,omitempty" xml:"DownStreamDepth,omitempty"`
	FilterList      []*ListNodeDownStreamRequestListQueryFilterList `json:"FilterList,omitempty" xml:"FilterList,omitempty" type:"Repeated"`
	// This parameter is required.
	NodeIdList []*ListNodeDownStreamRequestListQueryNodeIdList `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 123011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListNodeDownStreamRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDownStreamRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamRequestListQuery) SetDownStreamDepth(v int32) *ListNodeDownStreamRequestListQuery {
	s.DownStreamDepth = &v
	return s
}

func (s *ListNodeDownStreamRequestListQuery) SetFilterList(v []*ListNodeDownStreamRequestListQueryFilterList) *ListNodeDownStreamRequestListQuery {
	s.FilterList = v
	return s
}

func (s *ListNodeDownStreamRequestListQuery) SetNodeIdList(v []*ListNodeDownStreamRequestListQueryNodeIdList) *ListNodeDownStreamRequestListQuery {
	s.NodeIdList = v
	return s
}

func (s *ListNodeDownStreamRequestListQuery) SetProjectId(v int64) *ListNodeDownStreamRequestListQuery {
	s.ProjectId = &v
	return s
}

type ListNodeDownStreamRequestListQueryFilterList struct {
	// example:
	//
	// false
	Exclude *bool `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// example:
	//
	// PROJECT
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s ListNodeDownStreamRequestListQueryFilterList) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDownStreamRequestListQueryFilterList) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamRequestListQueryFilterList) SetExclude(v bool) *ListNodeDownStreamRequestListQueryFilterList {
	s.Exclude = &v
	return s
}

func (s *ListNodeDownStreamRequestListQueryFilterList) SetKey(v string) *ListNodeDownStreamRequestListQueryFilterList {
	s.Key = &v
	return s
}

func (s *ListNodeDownStreamRequestListQueryFilterList) SetValueList(v []*string) *ListNodeDownStreamRequestListQueryFilterList {
	s.ValueList = v
	return s
}

type ListNodeDownStreamRequestListQueryNodeIdList struct {
	// example:
	//
	// 112
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_23431
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListNodeDownStreamRequestListQueryNodeIdList) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDownStreamRequestListQueryNodeIdList) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamRequestListQueryNodeIdList) SetFieldIdList(v []*string) *ListNodeDownStreamRequestListQueryNodeIdList {
	s.FieldIdList = v
	return s
}

func (s *ListNodeDownStreamRequestListQueryNodeIdList) SetId(v string) *ListNodeDownStreamRequestListQueryNodeIdList {
	s.Id = &v
	return s
}

type ListNodeDownStreamShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListNodeDownStreamShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDownStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamShrinkRequest) SetEnv(v string) *ListNodeDownStreamShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListNodeDownStreamShrinkRequest) SetListQueryShrink(v string) *ListNodeDownStreamShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListNodeDownStreamShrinkRequest) SetOpTenantId(v int64) *ListNodeDownStreamShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListNodeDownStreamResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message      *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeInfoList []*ListNodeDownStreamResponseBodyNodeInfoList `json:"NodeInfoList,omitempty" xml:"NodeInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeDownStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDownStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamResponseBody) SetCode(v string) *ListNodeDownStreamResponseBody {
	s.Code = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetHttpStatusCode(v int32) *ListNodeDownStreamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetMessage(v string) *ListNodeDownStreamResponseBody {
	s.Message = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetNodeInfoList(v []*ListNodeDownStreamResponseBodyNodeInfoList) *ListNodeDownStreamResponseBody {
	s.NodeInfoList = v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetRequestId(v string) *ListNodeDownStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeDownStreamResponseBody) SetSuccess(v bool) *ListNodeDownStreamResponseBody {
	s.Success = &v
	return s
}

type ListNodeDownStreamResponseBodyNodeInfoList struct {
	// example:
	//
	// 1
	Depth       *int32    `json:"Depth,omitempty" xml:"Depth,omitempty"`
	FieldIdList []*string `json:"FieldIdList,omitempty" xml:"FieldIdList,omitempty" type:"Repeated"`
	// example:
	//
	// n_2423351
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDownStreamResponseBodyNodeInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDownStreamResponseBodyNodeInfoList) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetDepth(v int32) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.Depth = &v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetFieldIdList(v []*string) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.FieldIdList = v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetId(v string) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.Id = &v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetName(v string) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.Name = &v
	return s
}

func (s *ListNodeDownStreamResponseBodyNodeInfoList) SetType(v string) *ListNodeDownStreamResponseBodyNodeInfoList {
	s.Type = &v
	return s
}

type ListNodeDownStreamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeDownStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeDownStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDownStreamResponse) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamResponse) SetHeaders(v map[string]*string) *ListNodeDownStreamResponse {
	s.Headers = v
	return s
}

func (s *ListNodeDownStreamResponse) SetStatusCode(v int32) *ListNodeDownStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeDownStreamResponse) SetBody(v *ListNodeDownStreamResponseBody) *ListNodeDownStreamResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ListQuery *ListNodesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetEnv(v string) *ListNodesRequest {
	s.Env = &v
	return s
}

func (s *ListNodesRequest) SetListQuery(v *ListNodesRequestListQuery) *ListNodesRequest {
	s.ListQuery = v
	return s
}

func (s *ListNodesRequest) SetOpTenantId(v int64) *ListNodesRequest {
	s.OpTenantId = &v
	return s
}

type ListNodesRequestListQuery struct {
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SCRIPT
	NodeBizType *string `json:"NodeBizType,omitempty" xml:"NodeBizType,omitempty"`
	// This parameter is required.
	NodeSubBizTypeList []*string `json:"NodeSubBizTypeList,omitempty" xml:"NodeSubBizTypeList,omitempty" type:"Repeated"`
	OwnerList          []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PriorityList []*string `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 12111
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// true
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// NORMAL
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// example:
	//
	// xx
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
}

func (s ListNodesRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListNodesRequestListQuery) SetDryRun(v bool) *ListNodesRequestListQuery {
	s.DryRun = &v
	return s
}

func (s *ListNodesRequestListQuery) SetNodeBizType(v string) *ListNodesRequestListQuery {
	s.NodeBizType = &v
	return s
}

func (s *ListNodesRequestListQuery) SetNodeSubBizTypeList(v []*string) *ListNodesRequestListQuery {
	s.NodeSubBizTypeList = v
	return s
}

func (s *ListNodesRequestListQuery) SetOwnerList(v []*string) *ListNodesRequestListQuery {
	s.OwnerList = v
	return s
}

func (s *ListNodesRequestListQuery) SetPage(v int32) *ListNodesRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListNodesRequestListQuery) SetPageSize(v int32) *ListNodesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequestListQuery) SetPriorityList(v []*string) *ListNodesRequestListQuery {
	s.PriorityList = v
	return s
}

func (s *ListNodesRequestListQuery) SetProjectId(v int64) *ListNodesRequestListQuery {
	s.ProjectId = &v
	return s
}

func (s *ListNodesRequestListQuery) SetSchedulePaused(v bool) *ListNodesRequestListQuery {
	s.SchedulePaused = &v
	return s
}

func (s *ListNodesRequestListQuery) SetSchedulePeriodList(v []*string) *ListNodesRequestListQuery {
	s.SchedulePeriodList = v
	return s
}

func (s *ListNodesRequestListQuery) SetScheduleType(v string) *ListNodesRequestListQuery {
	s.ScheduleType = &v
	return s
}

func (s *ListNodesRequestListQuery) SetSearchText(v string) *ListNodesRequestListQuery {
	s.SearchText = &v
	return s
}

type ListNodesShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodesShrinkRequest) SetEnv(v string) *ListNodesShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListNodesShrinkRequest) SetListQueryShrink(v string) *ListNodesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetOpTenantId(v int64) *ListNodesShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListNodesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListNodesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetCode(v string) *ListNodesResponseBody {
	s.Code = &v
	return s
}

func (s *ListNodesResponseBody) SetHttpStatusCode(v int32) *ListNodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodesResponseBody) SetMessage(v string) *ListNodesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNodesResponseBody) SetPageResult(v *ListNodesResponseBodyPageResult) *ListNodesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetSuccess(v bool) *ListNodesResponseBody {
	s.Success = &v
	return s
}

type ListNodesResponseBodyPageResult struct {
	NodeList []*ListNodesResponseBodyPageResultNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResult) SetNodeList(v []*ListNodesResponseBodyPageResultNodeList) *ListNodesResponseBodyPageResult {
	s.NodeList = v
	return s
}

func (s *ListNodesResponseBodyPageResult) SetTotalCount(v int32) *ListNodesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

type ListNodesResponseBodyPageResultNodeList struct {
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	CreateTime *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator    *ListNodesResponseBodyPageResultNodeListCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// xx test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// {"xx":"zz"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// true
	HasDev *bool `json:"HasDev,omitempty" xml:"HasDev,omitempty"`
	// example:
	//
	// true
	HasProd *bool `json:"HasProd,omitempty" xml:"HasProd,omitempty"`
	// example:
	//
	// n_31111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	LastModifiedTime *string                                             `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier         *ListNodesResponseBodyPageResultNodeListModifier    `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name             *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList        []*ListNodesResponseBodyPageResultNodeListOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	PriorityList     []*string                                           `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	ProjectInfo      *ListNodesResponseBodyPageResultNodeListProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// example:
	//
	// true
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// SHELL
	SubDetailType *string `json:"SubDetailType,omitempty" xml:"SubDetailType,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeList) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeList) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeList) SetBizUnitName(v string) *ListNodesResponseBodyPageResultNodeList {
	s.BizUnitName = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetCreateTime(v string) *ListNodesResponseBodyPageResultNodeList {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetCreator(v *ListNodesResponseBodyPageResultNodeListCreator) *ListNodesResponseBodyPageResultNodeList {
	s.Creator = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetDescription(v string) *ListNodesResponseBodyPageResultNodeList {
	s.Description = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetDryRun(v bool) *ListNodesResponseBodyPageResultNodeList {
	s.DryRun = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetExtendInfo(v string) *ListNodesResponseBodyPageResultNodeList {
	s.ExtendInfo = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetFrom(v string) *ListNodesResponseBodyPageResultNodeList {
	s.From = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetHasDev(v bool) *ListNodesResponseBodyPageResultNodeList {
	s.HasDev = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetHasProd(v bool) *ListNodesResponseBodyPageResultNodeList {
	s.HasProd = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetId(v string) *ListNodesResponseBodyPageResultNodeList {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetLastModifiedTime(v string) *ListNodesResponseBodyPageResultNodeList {
	s.LastModifiedTime = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetModifier(v *ListNodesResponseBodyPageResultNodeListModifier) *ListNodesResponseBodyPageResultNodeList {
	s.Modifier = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetName(v string) *ListNodesResponseBodyPageResultNodeList {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetOwnerList(v []*ListNodesResponseBodyPageResultNodeListOwnerList) *ListNodesResponseBodyPageResultNodeList {
	s.OwnerList = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetPriorityList(v []*string) *ListNodesResponseBodyPageResultNodeList {
	s.PriorityList = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetProjectInfo(v *ListNodesResponseBodyPageResultNodeListProjectInfo) *ListNodesResponseBodyPageResultNodeList {
	s.ProjectInfo = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetSchedulePaused(v bool) *ListNodesResponseBodyPageResultNodeList {
	s.SchedulePaused = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetSchedulePeriodList(v []*string) *ListNodesResponseBodyPageResultNodeList {
	s.SchedulePeriodList = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetSubDetailType(v string) *ListNodesResponseBodyPageResultNodeList {
	s.SubDetailType = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetType(v string) *ListNodesResponseBodyPageResultNodeList {
	s.Type = &v
	return s
}

type ListNodesResponseBodyPageResultNodeListCreator struct {
	// example:
	//
	// 23222
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeListCreator) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeListCreator) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeListCreator) SetId(v string) *ListNodesResponseBodyPageResultNodeListCreator {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListCreator) SetName(v string) *ListNodesResponseBodyPageResultNodeListCreator {
	s.Name = &v
	return s
}

type ListNodesResponseBodyPageResultNodeListModifier struct {
	// example:
	//
	// 311131
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeListModifier) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeListModifier) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeListModifier) SetId(v string) *ListNodesResponseBodyPageResultNodeListModifier {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListModifier) SetName(v string) *ListNodesResponseBodyPageResultNodeListModifier {
	s.Name = &v
	return s
}

type ListNodesResponseBodyPageResultNodeListOwnerList struct {
	// example:
	//
	// 23222
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeListOwnerList) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeListOwnerList) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeListOwnerList) SetId(v string) *ListNodesResponseBodyPageResultNodeListOwnerList {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListOwnerList) SetName(v string) *ListNodesResponseBodyPageResultNodeListOwnerList {
	s.Name = &v
	return s
}

type ListNodesResponseBodyPageResultNodeListProjectInfo struct {
	// example:
	//
	// 1121321
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeListProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeListProjectInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeListProjectInfo) SetId(v string) *ListNodesResponseBodyPageResultNodeListProjectInfo {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListProjectInfo) SetName(v string) *ListNodesResponseBodyPageResultNodeListProjectInfo {
	s.Name = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListResourcePermissionOperationLogRequest struct {
	// This parameter is required.
	ListQuery *ListResourcePermissionOperationLogRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListResourcePermissionOperationLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogRequest) SetListQuery(v *ListResourcePermissionOperationLogRequestListQuery) *ListResourcePermissionOperationLogRequest {
	s.ListQuery = v
	return s
}

func (s *ListResourcePermissionOperationLogRequest) SetOpTenantId(v int64) *ListResourcePermissionOperationLogRequest {
	s.OpTenantId = &v
	return s
}

type ListResourcePermissionOperationLogRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx测试
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	TabType *string `json:"TabType,omitempty" xml:"TabType,omitempty"`
}

func (s ListResourcePermissionOperationLogRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogRequestListQuery) SetPage(v int32) *ListResourcePermissionOperationLogRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListResourcePermissionOperationLogRequestListQuery) SetPageSize(v int32) *ListResourcePermissionOperationLogRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListResourcePermissionOperationLogRequestListQuery) SetSearchText(v string) *ListResourcePermissionOperationLogRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListResourcePermissionOperationLogRequestListQuery) SetTabType(v string) *ListResourcePermissionOperationLogRequestListQuery {
	s.TabType = &v
	return s
}

type ListResourcePermissionOperationLogShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListResourcePermissionOperationLogShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogShrinkRequest) SetListQueryShrink(v string) *ListResourcePermissionOperationLogShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListResourcePermissionOperationLogShrinkRequest) SetOpTenantId(v int64) *ListResourcePermissionOperationLogShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListResourcePermissionOperationLogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListResourcePermissionOperationLogResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBody) SetCode(v string) *ListResourcePermissionOperationLogResponseBody {
	s.Code = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetHttpStatusCode(v int32) *ListResourcePermissionOperationLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetMessage(v string) *ListResourcePermissionOperationLogResponseBody {
	s.Message = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetPageResult(v *ListResourcePermissionOperationLogResponseBodyPageResult) *ListResourcePermissionOperationLogResponseBody {
	s.PageResult = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetRequestId(v string) *ListResourcePermissionOperationLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBody) SetSuccess(v bool) *ListResourcePermissionOperationLogResponseBody {
	s.Success = &v
	return s
}

type ListResourcePermissionOperationLogResponseBodyPageResult struct {
	Data []*ListResourcePermissionOperationLogResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 121
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResult) SetData(v []*ListResourcePermissionOperationLogResponseBodyPageResultData) *ListResourcePermissionOperationLogResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResult) SetTotalCount(v int64) *ListResourcePermissionOperationLogResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

type ListResourcePermissionOperationLogResponseBodyPageResultData struct {
	Account *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Struct"`
	// example:
	//
	// selectTable
	AuthScope *string `json:"AuthScope,omitempty" xml:"AuthScope,omitempty"`
	// example:
	//
	// 123133
	OperateId *int64 `json:"OperateId,omitempty" xml:"OperateId,omitempty"`
	// example:
	//
	// 1710012121000
	OperateTime *int64 `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// example:
	//
	// APPLY
	OperateType *string                                                             `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Period      *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// example:
	//
	// xx测试
	Reason        *string                                                                    `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ResourceInfo  *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo  `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Struct"`
	TargetAccount *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty" type:"Struct"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultData) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetAccount(v *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.Account = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetAuthScope(v string) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.AuthScope = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetOperateId(v int64) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.OperateId = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetOperateTime(v int64) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.OperateTime = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetOperateType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.OperateType = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetPeriod(v *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.Period = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetReason(v string) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.Reason = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetResourceInfo(v *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.ResourceInfo = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultData) SetTargetAccount(v *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) *ListResourcePermissionOperationLogResponseBodyPageResultData {
	s.TargetAccount = v
	return s
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataAccount struct {
	// example:
	//
	// 1212131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) SetId(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount) SetType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataAccount {
	s.Type = &v
	return s
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod struct {
	// example:
	//
	// 1712000000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) SetEndTime(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod {
	s.EndTime = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod) SetType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataPeriod {
	s.Type = &v
	return s
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo struct {
	BizUnitInfo *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo `json:"BizUnitInfo,omitempty" xml:"BizUnitInfo,omitempty" type:"Struct"`
	// example:
	//
	// tb1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// a.tb1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// tb1
	Name        *string                                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectInfo *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// example:
	//
	// PHYSICAL_TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetBizUnitInfo(v *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.BizUnitInfo = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetDisplayName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetEnv(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetId(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetProjectInfo(v *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.ProjectInfo = v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo) SetType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfo {
	s.Type = &v
	return s
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo struct {
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Id
	//
	// example:
	//
	// 121323
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) SetDisplayName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) SetEnv(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) SetId(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Name = &v
	return s
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo struct {
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 1123131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) SetDisplayName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) SetEnv(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) SetId(v int64) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Name = &v
	return s
}

type ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount struct {
	// example:
	//
	// 1212131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) SetId(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) SetName(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount) SetType(v string) *ListResourcePermissionOperationLogResponseBodyPageResultDataTargetAccount {
	s.Type = &v
	return s
}

type ListResourcePermissionOperationLogResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePermissionOperationLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePermissionOperationLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionOperationLogResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogResponse) SetHeaders(v map[string]*string) *ListResourcePermissionOperationLogResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePermissionOperationLogResponse) SetStatusCode(v int32) *ListResourcePermissionOperationLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePermissionOperationLogResponse) SetBody(v *ListResourcePermissionOperationLogResponseBody) *ListResourcePermissionOperationLogResponse {
	s.Body = v
	return s
}

type ListResourcePermissionsRequest struct {
	// This parameter is required.
	ListQuery *ListResourcePermissionsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListResourcePermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsRequest) SetListQuery(v *ListResourcePermissionsRequestListQuery) *ListResourcePermissionsRequest {
	s.ListQuery = v
	return s
}

func (s *ListResourcePermissionsRequest) SetOpTenantId(v int64) *ListResourcePermissionsRequest {
	s.OpTenantId = &v
	return s
}

type ListResourcePermissionsRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx测试
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	TabType *string `json:"TabType,omitempty" xml:"TabType,omitempty"`
}

func (s ListResourcePermissionsRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsRequestListQuery) SetPage(v int32) *ListResourcePermissionsRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListResourcePermissionsRequestListQuery) SetPageSize(v int32) *ListResourcePermissionsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListResourcePermissionsRequestListQuery) SetSearchText(v string) *ListResourcePermissionsRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListResourcePermissionsRequestListQuery) SetTabType(v string) *ListResourcePermissionsRequestListQuery {
	s.TabType = &v
	return s
}

type ListResourcePermissionsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListResourcePermissionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsShrinkRequest) SetListQueryShrink(v string) *ListResourcePermissionsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListResourcePermissionsShrinkRequest) SetOpTenantId(v int64) *ListResourcePermissionsShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListResourcePermissionsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListResourcePermissionsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourcePermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBody) SetCode(v string) *ListResourcePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetHttpStatusCode(v int32) *ListResourcePermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetMessage(v string) *ListResourcePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetPageResult(v *ListResourcePermissionsResponseBodyPageResult) *ListResourcePermissionsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetRequestId(v string) *ListResourcePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcePermissionsResponseBody) SetSuccess(v bool) *ListResourcePermissionsResponseBody {
	s.Success = &v
	return s
}

type ListResourcePermissionsResponseBodyPageResult struct {
	Data []*ListResourcePermissionsResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 121
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResult) SetData(v []*ListResourcePermissionsResponseBodyPageResultData) *ListResourcePermissionsResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResult) SetTotalCount(v int64) *ListResourcePermissionsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

type ListResourcePermissionsResponseBodyPageResultData struct {
	// example:
	//
	// selectTable
	AuthScope            *string                                                                  `json:"AuthScope,omitempty" xml:"AuthScope,omitempty"`
	Period               *ListResourcePermissionsResponseBodyPageResultDataPeriod                 `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	PermissionPeriodList []*ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList `json:"PermissionPeriodList,omitempty" xml:"PermissionPeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// 12123111
	RecordId      *string                                                         `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceInfo  *ListResourcePermissionsResponseBodyPageResultDataResourceInfo  `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty" type:"Struct"`
	TargetAccount *ListResourcePermissionsResponseBodyPageResultDataTargetAccount `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty" type:"Struct"`
}

func (s ListResourcePermissionsResponseBodyPageResultData) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetAuthScope(v string) *ListResourcePermissionsResponseBodyPageResultData {
	s.AuthScope = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetPeriod(v *ListResourcePermissionsResponseBodyPageResultDataPeriod) *ListResourcePermissionsResponseBodyPageResultData {
	s.Period = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetPermissionPeriodList(v []*ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) *ListResourcePermissionsResponseBodyPageResultData {
	s.PermissionPeriodList = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetRecordId(v string) *ListResourcePermissionsResponseBodyPageResultData {
	s.RecordId = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetResourceInfo(v *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) *ListResourcePermissionsResponseBodyPageResultData {
	s.ResourceInfo = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultData) SetTargetAccount(v *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) *ListResourcePermissionsResponseBodyPageResultData {
	s.TargetAccount = v
	return s
}

type ListResourcePermissionsResponseBodyPageResultDataPeriod struct {
	// example:
	//
	// 1712000000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataPeriod) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataPeriod) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPeriod) SetEndTime(v string) *ListResourcePermissionsResponseBodyPageResultDataPeriod {
	s.EndTime = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPeriod) SetType(v string) *ListResourcePermissionsResponseBodyPageResultDataPeriod {
	s.Type = &v
	return s
}

type ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList struct {
	Period *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// example:
	//
	// SELECT
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) SetPeriod(v *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList {
	s.Period = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList) SetPermissionType(v string) *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodList {
	s.PermissionType = &v
	return s
}

type ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod struct {
	// example:
	//
	// 1712000000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) SetEndTime(v string) *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod {
	s.EndTime = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod) SetType(v string) *ListResourcePermissionsResponseBodyPageResultDataPermissionPeriodListPeriod {
	s.Type = &v
	return s
}

type ListResourcePermissionsResponseBodyPageResultDataResourceInfo struct {
	BizUnitInfo *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo `json:"BizUnitInfo,omitempty" xml:"BizUnitInfo,omitempty" type:"Struct"`
	// example:
	//
	// tb1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// a.tb1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// tb1
	Name        *string                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectInfo *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// example:
	//
	// PHYSICAL_TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetBizUnitInfo(v *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.BizUnitInfo = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetDisplayName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetEnv(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetId(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetProjectInfo(v *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.ProjectInfo = v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfo) SetType(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfo {
	s.Type = &v
	return s
}

type ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo struct {
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Id
	//
	// example:
	//
	// 121323
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) SetDisplayName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) SetEnv(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) SetId(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo) SetName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoBizUnitInfo {
	s.Name = &v
	return s
}

type ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo struct {
	// example:
	//
	// xx
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 1123131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) SetDisplayName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	s.DisplayName = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) SetEnv(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Env = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) SetId(v int64) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo) SetName(v string) *ListResourcePermissionsResponseBodyPageResultDataResourceInfoProjectInfo {
	s.Name = &v
	return s
}

type ListResourcePermissionsResponseBodyPageResultDataTargetAccount struct {
	// example:
	//
	// 1212131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcePermissionsResponseBodyPageResultDataTargetAccount) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponseBodyPageResultDataTargetAccount) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) SetId(v string) *ListResourcePermissionsResponseBodyPageResultDataTargetAccount {
	s.Id = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) SetName(v string) *ListResourcePermissionsResponseBodyPageResultDataTargetAccount {
	s.Name = &v
	return s
}

func (s *ListResourcePermissionsResponseBodyPageResultDataTargetAccount) SetType(v string) *ListResourcePermissionsResponseBodyPageResultDataTargetAccount {
	s.Type = &v
	return s
}

type ListResourcePermissionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcePermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionsResponse) SetHeaders(v map[string]*string) *ListResourcePermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePermissionsResponse) SetStatusCode(v int32) *ListResourcePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePermissionsResponse) SetBody(v *ListResourcePermissionsResponseBody) *ListResourcePermissionsResponse {
	s.Body = v
	return s
}

type ListTenantMembersRequest struct {
	// This parameter is required.
	ListQuery *ListTenantMembersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListTenantMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTenantMembersRequest) GoString() string {
	return s.String()
}

func (s *ListTenantMembersRequest) SetListQuery(v *ListTenantMembersRequestListQuery) *ListTenantMembersRequest {
	s.ListQuery = v
	return s
}

func (s *ListTenantMembersRequest) SetOpTenantId(v int64) *ListTenantMembersRequest {
	s.OpTenantId = &v
	return s
}

type ListTenantMembersRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RoleList        []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	SearchText      *string   `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	UserGroupIdList []*string `json:"UserGroupIdList,omitempty" xml:"UserGroupIdList,omitempty" type:"Repeated"`
}

func (s ListTenantMembersRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListTenantMembersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListTenantMembersRequestListQuery) SetPage(v int32) *ListTenantMembersRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListTenantMembersRequestListQuery) SetPageSize(v int32) *ListTenantMembersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListTenantMembersRequestListQuery) SetRoleList(v []*string) *ListTenantMembersRequestListQuery {
	s.RoleList = v
	return s
}

func (s *ListTenantMembersRequestListQuery) SetSearchText(v string) *ListTenantMembersRequestListQuery {
	s.SearchText = &v
	return s
}

func (s *ListTenantMembersRequestListQuery) SetUserGroupIdList(v []*string) *ListTenantMembersRequestListQuery {
	s.UserGroupIdList = v
	return s
}

type ListTenantMembersShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListTenantMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTenantMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTenantMembersShrinkRequest) SetListQueryShrink(v string) *ListTenantMembersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListTenantMembersShrinkRequest) SetOpTenantId(v int64) *ListTenantMembersShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListTenantMembersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListTenantMembersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTenantMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTenantMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponseBody) SetCode(v string) *ListTenantMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantMembersResponseBody) SetHttpStatusCode(v int32) *ListTenantMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTenantMembersResponseBody) SetMessage(v string) *ListTenantMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantMembersResponseBody) SetPageResult(v *ListTenantMembersResponseBodyPageResult) *ListTenantMembersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListTenantMembersResponseBody) SetRequestId(v string) *ListTenantMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantMembersResponseBody) SetSuccess(v bool) *ListTenantMembersResponseBody {
	s.Success = &v
	return s
}

type ListTenantMembersResponseBodyPageResult struct {
	// example:
	//
	// 110
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserList   []*ListTenantMembersResponseBodyPageResultUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListTenantMembersResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListTenantMembersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponseBodyPageResult) SetTotalCount(v int32) *ListTenantMembersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResult) SetUserList(v []*ListTenantMembersResponseBodyPageResultUserList) *ListTenantMembersResponseBodyPageResult {
	s.UserList = v
	return s
}

type ListTenantMembersResponseBodyPageResultUserList struct {
	// example:
	//
	// zhangsan
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// dd123123
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// zhangsan
	DisplayNameWithoutStatus *string `json:"DisplayNameWithoutStatus,omitempty" xml:"DisplayNameWithoutStatus,omitempty"`
	// example:
	//
	// true
	EnableWhiteIp *string `json:"EnableWhiteIp,omitempty" xml:"EnableWhiteIp,omitempty"`
	// example:
	//
	// 1730000000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1730000000000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 132321
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 123@aliyun.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 13888888888
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// susan
	NickName *string   `json:"NickName,omitempty" xml:"NickName,omitempty"`
	RealName *string   `json:"RealName,omitempty" xml:"RealName,omitempty"`
	RoleList []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// example:
	//
	// 213213232422222
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// aliyun
	SourceType    *string                                                         `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	UserGroupList []*ListTenantMembersResponseBodyPageResultUserListUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
	// example:
	//
	// 0.0.0.0/0
	WhiteIp *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
}

func (s ListTenantMembersResponseBodyPageResultUserList) String() string {
	return tea.Prettify(s)
}

func (s ListTenantMembersResponseBodyPageResultUserList) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetAccountName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.AccountName = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetDingNumber(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.DingNumber = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetDisplayName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.DisplayName = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetDisplayNameWithoutStatus(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.DisplayNameWithoutStatus = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetEnableWhiteIp(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.EnableWhiteIp = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetGmtCreate(v int64) *ListTenantMembersResponseBodyPageResultUserList {
	s.GmtCreate = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetGmtModified(v int64) *ListTenantMembersResponseBodyPageResultUserList {
	s.GmtModified = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetId(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.Id = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetMail(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.Mail = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetMobilePhone(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.MobilePhone = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.Name = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetNickName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.NickName = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetRealName(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.RealName = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetRoleList(v []*string) *ListTenantMembersResponseBodyPageResultUserList {
	s.RoleList = v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetSourceId(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.SourceId = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetSourceType(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.SourceType = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetUserGroupList(v []*ListTenantMembersResponseBodyPageResultUserListUserGroupList) *ListTenantMembersResponseBodyPageResultUserList {
	s.UserGroupList = v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserList) SetWhiteIp(v string) *ListTenantMembersResponseBodyPageResultUserList {
	s.WhiteIp = &v
	return s
}

type ListTenantMembersResponseBodyPageResultUserListUserGroupList struct {
	// example:
	//
	// true
	Active      *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 121313
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTenantMembersResponseBodyPageResultUserListUserGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListTenantMembersResponseBodyPageResultUserListUserGroupList) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) SetActive(v bool) *ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	s.Active = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) SetDescription(v string) *ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	s.Description = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) SetId(v string) *ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	s.Id = &v
	return s
}

func (s *ListTenantMembersResponseBodyPageResultUserListUserGroupList) SetName(v string) *ListTenantMembersResponseBodyPageResultUserListUserGroupList {
	s.Name = &v
	return s
}

type ListTenantMembersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTenantMembersResponse) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponse) SetHeaders(v map[string]*string) *ListTenantMembersResponse {
	s.Headers = v
	return s
}

func (s *ListTenantMembersResponse) SetStatusCode(v int32) *ListTenantMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantMembersResponse) SetBody(v *ListTenantMembersResponseBody) *ListTenantMembersResponse {
	s.Body = v
	return s
}

type ListUserGroupMembersRequest struct {
	// This parameter is required.
	ListQuery *ListUserGroupMembersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListUserGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersRequest) SetListQuery(v *ListUserGroupMembersRequestListQuery) *ListUserGroupMembersRequest {
	s.ListQuery = v
	return s
}

func (s *ListUserGroupMembersRequest) SetOpTenantId(v int64) *ListUserGroupMembersRequest {
	s.OpTenantId = &v
	return s
}

type ListUserGroupMembersRequestListQuery struct {
	// example:
	//
	// a
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 232231
	UserGroupId *string   `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserIdList  []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s ListUserGroupMembersRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersRequestListQuery) SetKeyword(v string) *ListUserGroupMembersRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) SetPageNo(v int32) *ListUserGroupMembersRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) SetPageSize(v int32) *ListUserGroupMembersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) SetUserGroupId(v string) *ListUserGroupMembersRequestListQuery {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) SetUserIdList(v []*string) *ListUserGroupMembersRequestListQuery {
	s.UserIdList = v
	return s
}

type ListUserGroupMembersShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListUserGroupMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersShrinkRequest) SetListQueryShrink(v string) *ListUserGroupMembersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListUserGroupMembersShrinkRequest) SetOpTenantId(v int64) *ListUserGroupMembersShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListUserGroupMembersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListUserGroupMembersResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBody) SetCode(v string) *ListUserGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetHttpStatusCode(v int32) *ListUserGroupMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetMessage(v string) *ListUserGroupMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetPageResult(v *ListUserGroupMembersResponseBodyPageResult) *ListUserGroupMembersResponseBody {
	s.PageResult = v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetRequestId(v string) *ListUserGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupMembersResponseBody) SetSuccess(v bool) *ListUserGroupMembersResponseBody {
	s.Success = &v
	return s
}

type ListUserGroupMembersResponseBodyPageResult struct {
	MemberList []*ListUserGroupMembersResponseBodyPageResultMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// example:
	//
	// 217
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserGroupMembersResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBodyPageResult) SetMemberList(v []*ListUserGroupMembersResponseBodyPageResultMemberList) *ListUserGroupMembersResponseBodyPageResult {
	s.MemberList = v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResult) SetTotalCount(v int32) *ListUserGroupMembersResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

type ListUserGroupMembersResponseBodyPageResultMemberList struct {
	Creator *ListUserGroupMembersResponseBodyPageResultMemberListCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// zhangsan
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2324211
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 231111
	UserGroupId *string                                                       `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserInfo    *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	// example:
	//
	// SECURITY_ADMIN
	UserRole *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
}

func (s ListUserGroupMembersResponseBodyPageResultMemberList) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersResponseBodyPageResultMemberList) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetCreator(v *ListUserGroupMembersResponseBodyPageResultMemberListCreator) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.Creator = v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetGmtCreate(v int64) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.GmtCreate = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetId(v string) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.Id = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetUserGroupId(v string) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetUserInfo(v *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.UserInfo = v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberList) SetUserRole(v string) *ListUserGroupMembersResponseBodyPageResultMemberList {
	s.UserRole = &v
	return s
}

type ListUserGroupMembersResponseBodyPageResultMemberListCreator struct {
	// example:
	//
	// 12121111
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 12121111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserGroupMembersResponseBodyPageResultMemberListCreator) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersResponseBodyPageResultMemberListCreator) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) SetAccountName(v string) *ListUserGroupMembersResponseBodyPageResultMemberListCreator {
	s.AccountName = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) SetDisplayName(v string) *ListUserGroupMembersResponseBodyPageResultMemberListCreator {
	s.DisplayName = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListCreator) SetId(v string) *ListUserGroupMembersResponseBodyPageResultMemberListCreator {
	s.Id = &v
	return s
}

type ListUserGroupMembersResponseBodyPageResultMemberListUserInfo struct {
	// example:
	//
	// atest
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 13232
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) SetAccountName(v string) *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo {
	s.AccountName = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) SetDisplayName(v string) *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo {
	s.DisplayName = &v
	return s
}

func (s *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo) SetId(v string) *ListUserGroupMembersResponseBodyPageResultMemberListUserInfo {
	s.Id = &v
	return s
}

type ListUserGroupMembersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersResponse) SetHeaders(v map[string]*string) *ListUserGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupMembersResponse) SetStatusCode(v int32) *ListUserGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupMembersResponse) SetBody(v *ListUserGroupMembersResponseBody) *ListUserGroupMembersResponse {
	s.Body = v
	return s
}

type ListUserGroupsRequest struct {
	// This parameter is required.
	ListQuery *ListUserGroupsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) SetListQuery(v *ListUserGroupsRequestListQuery) *ListUserGroupsRequest {
	s.ListQuery = v
	return s
}

func (s *ListUserGroupsRequest) SetOpTenantId(v int64) *ListUserGroupsRequest {
	s.OpTenantId = &v
	return s
}

type ListUserGroupsRequestListQuery struct {
	// example:
	//
	// true
	Active      *bool     `json:"Active,omitempty" xml:"Active,omitempty"`
	AdminIdList []*string `json:"AdminIdList,omitempty" xml:"AdminIdList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	FilterMine *bool   `json:"FilterMine,omitempty" xml:"FilterMine,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserGroupsRequestListQuery) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequestListQuery) SetActive(v bool) *ListUserGroupsRequestListQuery {
	s.Active = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetAdminIdList(v []*string) *ListUserGroupsRequestListQuery {
	s.AdminIdList = v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetFilterMine(v bool) *ListUserGroupsRequestListQuery {
	s.FilterMine = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetKeyword(v string) *ListUserGroupsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetPageNo(v int32) *ListUserGroupsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListUserGroupsRequestListQuery) SetPageSize(v int32) *ListUserGroupsRequestListQuery {
	s.PageSize = &v
	return s
}

type ListUserGroupsShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListUserGroupsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsShrinkRequest) SetListQueryShrink(v string) *ListUserGroupsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListUserGroupsShrinkRequest) SetOpTenantId(v int64) *ListUserGroupsShrinkRequest {
	s.OpTenantId = &v
	return s
}

type ListUserGroupsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListUserGroupsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) SetCode(v string) *ListUserGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetHttpStatusCode(v int32) *ListUserGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetMessage(v string) *ListUserGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetPageResult(v *ListUserGroupsResponseBodyPageResult) *ListUserGroupsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetSuccess(v bool) *ListUserGroupsResponseBody {
	s.Success = &v
	return s
}

type ListUserGroupsResponseBodyPageResult struct {
	// example:
	//
	// 49
	TotalCount    *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserGroupList []*ListUserGroupsResponseBodyPageResultUserGroupList `json:"UserGroupList,omitempty" xml:"UserGroupList,omitempty" type:"Repeated"`
}

func (s ListUserGroupsResponseBodyPageResult) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyPageResult) SetTotalCount(v int32) *ListUserGroupsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResult) SetUserGroupList(v []*ListUserGroupsResponseBodyPageResultUserGroupList) *ListUserGroupsResponseBodyPageResult {
	s.UserGroupList = v
	return s
}

type ListUserGroupsResponseBodyPageResultUserGroupList struct {
	// example:
	//
	// true
	Active      *bool                                                         `json:"Active,omitempty" xml:"Active,omitempty"`
	AdminList   []*ListUserGroupsResponseBodyPageResultUserGroupListAdminList `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	Description *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 31313232
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// SECURITY_ADMIN
	MyRole *string `json:"MyRole,omitempty" xml:"MyRole,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUserGroupsResponseBodyPageResultUserGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBodyPageResultUserGroupList) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetActive(v bool) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.Active = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetAdminList(v []*ListUserGroupsResponseBodyPageResultUserGroupListAdminList) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.AdminList = v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetDescription(v string) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.Description = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetId(v string) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.Id = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetMyRole(v string) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.MyRole = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupList) SetName(v string) *ListUserGroupsResponseBodyPageResultUserGroupList {
	s.Name = &v
	return s
}

type ListUserGroupsResponseBodyPageResultUserGroupListAdminList struct {
	// example:
	//
	// zhangsan
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 32313131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserGroupsResponseBodyPageResultUserGroupListAdminList) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBodyPageResultUserGroupListAdminList) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) SetAccountName(v string) *ListUserGroupsResponseBodyPageResultUserGroupListAdminList {
	s.AccountName = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) SetDisplayName(v string) *ListUserGroupsResponseBodyPageResultUserGroupListAdminList {
	s.DisplayName = &v
	return s
}

func (s *ListUserGroupsResponseBodyPageResultUserGroupListAdminList) SetId(v string) *ListUserGroupsResponseBodyPageResultUserGroupListAdminList {
	s.Id = &v
	return s
}

type ListUserGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponse) SetHeaders(v map[string]*string) *ListUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsResponse) SetStatusCode(v int32) *ListUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsResponse) SetBody(v *ListUserGroupsResponseBody) *ListUserGroupsResponse {
	s.Body = v
	return s
}

type OperateInstanceRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	OperateCommand *OperateInstanceRequestOperateCommand `json:"OperateCommand,omitempty" xml:"OperateCommand,omitempty" type:"Struct"`
}

func (s OperateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateInstanceRequest) GoString() string {
	return s.String()
}

func (s *OperateInstanceRequest) SetEnv(v string) *OperateInstanceRequest {
	s.Env = &v
	return s
}

func (s *OperateInstanceRequest) SetOpTenantId(v int64) *OperateInstanceRequest {
	s.OpTenantId = &v
	return s
}

func (s *OperateInstanceRequest) SetOperateCommand(v *OperateInstanceRequestOperateCommand) *OperateInstanceRequest {
	s.OperateCommand = v
	return s
}

type OperateInstanceRequestOperateCommand struct {
	// This parameter is required.
	InstanceIdList []*OperateInstanceRequestOperateCommandInstanceIdList `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// RERUN
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 132311
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s OperateInstanceRequestOperateCommand) String() string {
	return tea.Prettify(s)
}

func (s OperateInstanceRequestOperateCommand) GoString() string {
	return s.String()
}

func (s *OperateInstanceRequestOperateCommand) SetInstanceIdList(v []*OperateInstanceRequestOperateCommandInstanceIdList) *OperateInstanceRequestOperateCommand {
	s.InstanceIdList = v
	return s
}

func (s *OperateInstanceRequestOperateCommand) SetOperation(v string) *OperateInstanceRequestOperateCommand {
	s.Operation = &v
	return s
}

func (s *OperateInstanceRequestOperateCommand) SetProjectId(v int64) *OperateInstanceRequestOperateCommand {
	s.ProjectId = &v
	return s
}

type OperateInstanceRequestOperateCommandInstanceIdList struct {
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// t_32111312
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s OperateInstanceRequestOperateCommandInstanceIdList) String() string {
	return tea.Prettify(s)
}

func (s OperateInstanceRequestOperateCommandInstanceIdList) GoString() string {
	return s.String()
}

func (s *OperateInstanceRequestOperateCommandInstanceIdList) SetFieldInstanceIdList(v []*string) *OperateInstanceRequestOperateCommandInstanceIdList {
	s.FieldInstanceIdList = v
	return s
}

func (s *OperateInstanceRequestOperateCommandInstanceIdList) SetId(v string) *OperateInstanceRequestOperateCommandInstanceIdList {
	s.Id = &v
	return s
}

type OperateInstanceShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	OperateCommandShrink *string `json:"OperateCommand,omitempty" xml:"OperateCommand,omitempty"`
}

func (s OperateInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateInstanceShrinkRequest) SetEnv(v string) *OperateInstanceShrinkRequest {
	s.Env = &v
	return s
}

func (s *OperateInstanceShrinkRequest) SetOpTenantId(v int64) *OperateInstanceShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *OperateInstanceShrinkRequest) SetOperateCommandShrink(v string) *OperateInstanceShrinkRequest {
	s.OperateCommandShrink = &v
	return s
}

type OperateInstanceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode     *int32                                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceStatusList []*OperateInstanceResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *OperateInstanceResponseBody) SetCode(v string) *OperateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *OperateInstanceResponseBody) SetHttpStatusCode(v int32) *OperateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OperateInstanceResponseBody) SetInstanceStatusList(v []*OperateInstanceResponseBodyInstanceStatusList) *OperateInstanceResponseBody {
	s.InstanceStatusList = v
	return s
}

func (s *OperateInstanceResponseBody) SetMessage(v string) *OperateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *OperateInstanceResponseBody) SetRequestId(v string) *OperateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateInstanceResponseBody) SetSuccess(v bool) *OperateInstanceResponseBody {
	s.Success = &v
	return s
}

type OperateInstanceResponseBodyInstanceStatusList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// t_132435
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 121311
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s OperateInstanceResponseBodyInstanceStatusList) String() string {
	return tea.Prettify(s)
}

func (s OperateInstanceResponseBodyInstanceStatusList) GoString() string {
	return s.String()
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetDisplayName(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.DisplayName = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetErrorMessage(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.ErrorMessage = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetId(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.Id = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetName(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.Name = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetOwnerId(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.OwnerId = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetOwnerName(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.OwnerName = &v
	return s
}

func (s *OperateInstanceResponseBodyInstanceStatusList) SetStatus(v string) *OperateInstanceResponseBodyInstanceStatusList {
	s.Status = &v
	return s
}

type OperateInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateInstanceResponse) GoString() string {
	return s.String()
}

func (s *OperateInstanceResponse) SetHeaders(v map[string]*string) *OperateInstanceResponse {
	s.Headers = v
	return s
}

func (s *OperateInstanceResponse) SetStatusCode(v int32) *OperateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateInstanceResponse) SetBody(v *OperateInstanceResponseBody) *OperateInstanceResponse {
	s.Body = v
	return s
}

type PausePhysicalNodeRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PauseCommand *PausePhysicalNodeRequestPauseCommand `json:"PauseCommand,omitempty" xml:"PauseCommand,omitempty" type:"Struct"`
}

func (s PausePhysicalNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s PausePhysicalNodeRequest) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeRequest) SetEnv(v string) *PausePhysicalNodeRequest {
	s.Env = &v
	return s
}

func (s *PausePhysicalNodeRequest) SetOpTenantId(v int64) *PausePhysicalNodeRequest {
	s.OpTenantId = &v
	return s
}

func (s *PausePhysicalNodeRequest) SetPauseCommand(v *PausePhysicalNodeRequestPauseCommand) *PausePhysicalNodeRequest {
	s.PauseCommand = v
	return s
}

type PausePhysicalNodeRequestPauseCommand struct {
	// This parameter is required.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 13222210
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PausePhysicalNodeRequestPauseCommand) String() string {
	return tea.Prettify(s)
}

func (s PausePhysicalNodeRequestPauseCommand) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeRequestPauseCommand) SetNodeIdList(v []*string) *PausePhysicalNodeRequestPauseCommand {
	s.NodeIdList = v
	return s
}

func (s *PausePhysicalNodeRequestPauseCommand) SetProjectId(v int64) *PausePhysicalNodeRequestPauseCommand {
	s.ProjectId = &v
	return s
}

type PausePhysicalNodeShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PauseCommandShrink *string `json:"PauseCommand,omitempty" xml:"PauseCommand,omitempty"`
}

func (s PausePhysicalNodeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PausePhysicalNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeShrinkRequest) SetEnv(v string) *PausePhysicalNodeShrinkRequest {
	s.Env = &v
	return s
}

func (s *PausePhysicalNodeShrinkRequest) SetOpTenantId(v int64) *PausePhysicalNodeShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *PausePhysicalNodeShrinkRequest) SetPauseCommandShrink(v string) *PausePhysicalNodeShrinkRequest {
	s.PauseCommandShrink = &v
	return s
}

type PausePhysicalNodeResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message               *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeOperateResultList []*PausePhysicalNodeResponseBodyNodeOperateResultList `json:"NodeOperateResultList,omitempty" xml:"NodeOperateResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PausePhysicalNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PausePhysicalNodeResponseBody) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeResponseBody) SetCode(v string) *PausePhysicalNodeResponseBody {
	s.Code = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetHttpStatusCode(v int32) *PausePhysicalNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetMessage(v string) *PausePhysicalNodeResponseBody {
	s.Message = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetNodeOperateResultList(v []*PausePhysicalNodeResponseBodyNodeOperateResultList) *PausePhysicalNodeResponseBody {
	s.NodeOperateResultList = v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetRequestId(v string) *PausePhysicalNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PausePhysicalNodeResponseBody) SetSuccess(v bool) *PausePhysicalNodeResponseBody {
	s.Success = &v
	return s
}

type PausePhysicalNodeResponseBodyNodeOperateResultList struct {
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// n_123456
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PausePhysicalNodeResponseBodyNodeOperateResultList) String() string {
	return tea.Prettify(s)
}

func (s PausePhysicalNodeResponseBodyNodeOperateResultList) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) SetErrorMessage(v string) *PausePhysicalNodeResponseBodyNodeOperateResultList {
	s.ErrorMessage = &v
	return s
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) SetNodeId(v string) *PausePhysicalNodeResponseBodyNodeOperateResultList {
	s.NodeId = &v
	return s
}

func (s *PausePhysicalNodeResponseBodyNodeOperateResultList) SetStatus(v string) *PausePhysicalNodeResponseBodyNodeOperateResultList {
	s.Status = &v
	return s
}

type PausePhysicalNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PausePhysicalNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PausePhysicalNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s PausePhysicalNodeResponse) GoString() string {
	return s.String()
}

func (s *PausePhysicalNodeResponse) SetHeaders(v map[string]*string) *PausePhysicalNodeResponse {
	s.Headers = v
	return s
}

func (s *PausePhysicalNodeResponse) SetStatusCode(v int32) *PausePhysicalNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *PausePhysicalNodeResponse) SetBody(v *PausePhysicalNodeResponseBody) *PausePhysicalNodeResponse {
	s.Body = v
	return s
}

type RemoveTenantMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommand *RemoveTenantMemberRequestRemoveCommand `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty" type:"Struct"`
}

func (s RemoveTenantMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberRequest) SetOpTenantId(v int64) *RemoveTenantMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveTenantMemberRequest) SetRemoveCommand(v *RemoveTenantMemberRequestRemoveCommand) *RemoveTenantMemberRequest {
	s.RemoveCommand = v
	return s
}

type RemoveTenantMemberRequestRemoveCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@xx.com
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s RemoveTenantMemberRequestRemoveCommand) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantMemberRequestRemoveCommand) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberRequestRemoveCommand) SetSourceId(v string) *RemoveTenantMemberRequestRemoveCommand {
	s.SourceId = &v
	return s
}

type RemoveTenantMemberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommandShrink *string `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty"`
}

func (s RemoveTenantMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberShrinkRequest) SetOpTenantId(v int64) *RemoveTenantMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveTenantMemberShrinkRequest) SetRemoveCommandShrink(v string) *RemoveTenantMemberShrinkRequest {
	s.RemoveCommandShrink = &v
	return s
}

type RemoveTenantMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveTenantMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberResponseBody) SetCode(v string) *RemoveTenantMemberResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) SetHttpStatusCode(v int32) *RemoveTenantMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) SetMessage(v string) *RemoveTenantMemberResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) SetRequestId(v string) *RemoveTenantMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) SetSuccess(v bool) *RemoveTenantMemberResponseBody {
	s.Success = &v
	return s
}

type RemoveTenantMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTenantMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTenantMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberResponse) SetHeaders(v map[string]*string) *RemoveTenantMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveTenantMemberResponse) SetStatusCode(v int32) *RemoveTenantMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTenantMemberResponse) SetBody(v *RemoveTenantMemberResponseBody) *RemoveTenantMemberResponse {
	s.Body = v
	return s
}

type RemoveUserGroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommand *RemoveUserGroupMemberRequestRemoveCommand `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty" type:"Struct"`
}

func (s RemoveUserGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberRequest) SetOpTenantId(v int64) *RemoveUserGroupMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveUserGroupMemberRequest) SetRemoveCommand(v *RemoveUserGroupMemberRequestRemoveCommand) *RemoveUserGroupMemberRequest {
	s.RemoveCommand = v
	return s
}

type RemoveUserGroupMemberRequestRemoveCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 2311
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s RemoveUserGroupMemberRequestRemoveCommand) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserGroupMemberRequestRemoveCommand) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberRequestRemoveCommand) SetUserGroupId(v string) *RemoveUserGroupMemberRequestRemoveCommand {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUserGroupMemberRequestRemoveCommand) SetUserIdList(v []*string) *RemoveUserGroupMemberRequestRemoveCommand {
	s.UserIdList = v
	return s
}

type RemoveUserGroupMemberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommandShrink *string `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty"`
}

func (s RemoveUserGroupMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserGroupMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberShrinkRequest) SetOpTenantId(v int64) *RemoveUserGroupMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveUserGroupMemberShrinkRequest) SetRemoveCommandShrink(v string) *RemoveUserGroupMemberShrinkRequest {
	s.RemoveCommandShrink = &v
	return s
}

type RemoveUserGroupMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveUserGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberResponseBody) SetCode(v string) *RemoveUserGroupMemberResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetData(v bool) *RemoveUserGroupMemberResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetHttpStatusCode(v int32) *RemoveUserGroupMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetMessage(v string) *RemoveUserGroupMemberResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetRequestId(v string) *RemoveUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserGroupMemberResponseBody) SetSuccess(v bool) *RemoveUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

type RemoveUserGroupMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMemberResponse) SetHeaders(v map[string]*string) *RemoveUserGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserGroupMemberResponse) SetStatusCode(v int32) *RemoveUserGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserGroupMemberResponse) SetBody(v *RemoveUserGroupMemberResponseBody) *RemoveUserGroupMemberResponse {
	s.Body = v
	return s
}

type ResumePhysicalNodeRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	ResumeCommand *ResumePhysicalNodeRequestResumeCommand `json:"ResumeCommand,omitempty" xml:"ResumeCommand,omitempty" type:"Struct"`
}

func (s ResumePhysicalNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumePhysicalNodeRequest) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeRequest) SetEnv(v string) *ResumePhysicalNodeRequest {
	s.Env = &v
	return s
}

func (s *ResumePhysicalNodeRequest) SetOpTenantId(v int64) *ResumePhysicalNodeRequest {
	s.OpTenantId = &v
	return s
}

func (s *ResumePhysicalNodeRequest) SetResumeCommand(v *ResumePhysicalNodeRequestResumeCommand) *ResumePhysicalNodeRequest {
	s.ResumeCommand = v
	return s
}

type ResumePhysicalNodeRequestResumeCommand struct {
	// This parameter is required.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 102011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ResumePhysicalNodeRequestResumeCommand) String() string {
	return tea.Prettify(s)
}

func (s ResumePhysicalNodeRequestResumeCommand) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeRequestResumeCommand) SetNodeIdList(v []*string) *ResumePhysicalNodeRequestResumeCommand {
	s.NodeIdList = v
	return s
}

func (s *ResumePhysicalNodeRequestResumeCommand) SetProjectId(v int64) *ResumePhysicalNodeRequestResumeCommand {
	s.ProjectId = &v
	return s
}

type ResumePhysicalNodeShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	ResumeCommandShrink *string `json:"ResumeCommand,omitempty" xml:"ResumeCommand,omitempty"`
}

func (s ResumePhysicalNodeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumePhysicalNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeShrinkRequest) SetEnv(v string) *ResumePhysicalNodeShrinkRequest {
	s.Env = &v
	return s
}

func (s *ResumePhysicalNodeShrinkRequest) SetOpTenantId(v int64) *ResumePhysicalNodeShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ResumePhysicalNodeShrinkRequest) SetResumeCommandShrink(v string) *ResumePhysicalNodeShrinkRequest {
	s.ResumeCommandShrink = &v
	return s
}

type ResumePhysicalNodeResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message               *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NodeOperateResultList []*ResumePhysicalNodeResponseBodyNodeOperateResultList `json:"NodeOperateResultList,omitempty" xml:"NodeOperateResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumePhysicalNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumePhysicalNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeResponseBody) SetCode(v string) *ResumePhysicalNodeResponseBody {
	s.Code = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetHttpStatusCode(v int32) *ResumePhysicalNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetMessage(v string) *ResumePhysicalNodeResponseBody {
	s.Message = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetNodeOperateResultList(v []*ResumePhysicalNodeResponseBodyNodeOperateResultList) *ResumePhysicalNodeResponseBody {
	s.NodeOperateResultList = v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetRequestId(v string) *ResumePhysicalNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumePhysicalNodeResponseBody) SetSuccess(v bool) *ResumePhysicalNodeResponseBody {
	s.Success = &v
	return s
}

type ResumePhysicalNodeResponseBodyNodeOperateResultList struct {
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// n_123456
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResumePhysicalNodeResponseBodyNodeOperateResultList) String() string {
	return tea.Prettify(s)
}

func (s ResumePhysicalNodeResponseBodyNodeOperateResultList) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) SetErrorMessage(v string) *ResumePhysicalNodeResponseBodyNodeOperateResultList {
	s.ErrorMessage = &v
	return s
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) SetNodeId(v string) *ResumePhysicalNodeResponseBodyNodeOperateResultList {
	s.NodeId = &v
	return s
}

func (s *ResumePhysicalNodeResponseBodyNodeOperateResultList) SetStatus(v string) *ResumePhysicalNodeResponseBodyNodeOperateResultList {
	s.Status = &v
	return s
}

type ResumePhysicalNodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumePhysicalNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumePhysicalNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumePhysicalNodeResponse) GoString() string {
	return s.String()
}

func (s *ResumePhysicalNodeResponse) SetHeaders(v map[string]*string) *ResumePhysicalNodeResponse {
	s.Headers = v
	return s
}

func (s *ResumePhysicalNodeResponse) SetStatusCode(v int32) *ResumePhysicalNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumePhysicalNodeResponse) SetBody(v *ResumePhysicalNodeResponseBody) *ResumePhysicalNodeResponse {
	s.Body = v
	return s
}

type RevokeResourcePermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RevokeCommand *RevokeResourcePermissionRequestRevokeCommand `json:"RevokeCommand,omitempty" xml:"RevokeCommand,omitempty" type:"Struct"`
}

func (s RevokeResourcePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeResourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionRequest) SetOpTenantId(v int64) *RevokeResourcePermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *RevokeResourcePermissionRequest) SetRevokeCommand(v *RevokeResourcePermissionRequestRevokeCommand) *RevokeResourcePermissionRequest {
	s.RevokeCommand = v
	return s
}

type RevokeResourcePermissionRequestRevokeCommand struct {
	OperateList []*string `json:"OperateList,omitempty" xml:"OperateList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// This parameter is required.
	ResourceList []*RevokeResourcePermissionRequestRevokeCommandResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PHYSICAL_TABLE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13131
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeResourcePermissionRequestRevokeCommand) String() string {
	return tea.Prettify(s)
}

func (s RevokeResourcePermissionRequestRevokeCommand) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetOperateList(v []*string) *RevokeResourcePermissionRequestRevokeCommand {
	s.OperateList = v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetReason(v string) *RevokeResourcePermissionRequestRevokeCommand {
	s.Reason = &v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetResourceList(v []*RevokeResourcePermissionRequestRevokeCommandResourceList) *RevokeResourcePermissionRequestRevokeCommand {
	s.ResourceList = v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetResourceType(v string) *RevokeResourcePermissionRequestRevokeCommand {
	s.ResourceType = &v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetUserId(v string) *RevokeResourcePermissionRequestRevokeCommand {
	s.UserId = &v
	return s
}

type RevokeResourcePermissionRequestRevokeCommandResourceList struct {
	// example:
	//
	// odps.300002102.beginner_test.amin_table
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s RevokeResourcePermissionRequestRevokeCommandResourceList) String() string {
	return tea.Prettify(s)
}

func (s RevokeResourcePermissionRequestRevokeCommandResourceList) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionRequestRevokeCommandResourceList) SetResourceId(v string) *RevokeResourcePermissionRequestRevokeCommandResourceList {
	s.ResourceId = &v
	return s
}

type RevokeResourcePermissionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RevokeCommandShrink *string `json:"RevokeCommand,omitempty" xml:"RevokeCommand,omitempty"`
}

func (s RevokeResourcePermissionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeResourcePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionShrinkRequest) SetOpTenantId(v int64) *RevokeResourcePermissionShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *RevokeResourcePermissionShrinkRequest) SetRevokeCommandShrink(v string) *RevokeResourcePermissionShrinkRequest {
	s.RevokeCommandShrink = &v
	return s
}

type RevokeResourcePermissionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeResourcePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeResourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionResponseBody) SetCode(v string) *RevokeResourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) SetHttpStatusCode(v int32) *RevokeResourcePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) SetMessage(v string) *RevokeResourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) SetRequestId(v string) *RevokeResourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) SetSuccess(v bool) *RevokeResourcePermissionResponseBody {
	s.Success = &v
	return s
}

type RevokeResourcePermissionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResourcePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeResourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionResponse) SetHeaders(v map[string]*string) *RevokeResourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeResourcePermissionResponse) SetStatusCode(v int32) *RevokeResourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResourcePermissionResponse) SetBody(v *RevokeResourcePermissionResponseBody) *RevokeResourcePermissionResponse {
	s.Body = v
	return s
}

type UpdateAdHocFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateAdHocFileRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateAdHocFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdHocFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileRequest) SetOpTenantId(v int64) *UpdateAdHocFileRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateAdHocFileRequest) SetUpdateCommand(v *UpdateAdHocFileRequestUpdateCommand) *UpdateAdHocFileRequest {
	s.UpdateCommand = v
	return s
}

type UpdateAdHocFileRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// select 1;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2311113
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1212313
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateAdHocFileRequestUpdateCommand) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdHocFileRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileRequestUpdateCommand) SetContent(v string) *UpdateAdHocFileRequestUpdateCommand {
	s.Content = &v
	return s
}

func (s *UpdateAdHocFileRequestUpdateCommand) SetFileId(v int64) *UpdateAdHocFileRequestUpdateCommand {
	s.FileId = &v
	return s
}

func (s *UpdateAdHocFileRequestUpdateCommand) SetProjectId(v int64) *UpdateAdHocFileRequestUpdateCommand {
	s.ProjectId = &v
	return s
}

type UpdateAdHocFileShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateAdHocFileShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdHocFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileShrinkRequest) SetOpTenantId(v int64) *UpdateAdHocFileShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateAdHocFileShrinkRequest) SetUpdateCommandShrink(v string) *UpdateAdHocFileShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

type UpdateAdHocFileResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAdHocFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdHocFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileResponseBody) SetCode(v string) *UpdateAdHocFileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) SetHttpStatusCode(v int32) *UpdateAdHocFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) SetMessage(v string) *UpdateAdHocFileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) SetRequestId(v string) *UpdateAdHocFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdHocFileResponseBody) SetSuccess(v bool) *UpdateAdHocFileResponseBody {
	s.Success = &v
	return s
}

type UpdateAdHocFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAdHocFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAdHocFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAdHocFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdHocFileResponse) SetHeaders(v map[string]*string) *UpdateAdHocFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdHocFileResponse) SetStatusCode(v int32) *UpdateAdHocFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdHocFileResponse) SetBody(v *UpdateAdHocFileResponseBody) *UpdateAdHocFileResponse {
	s.Body = v
	return s
}

type UpdateDataSourceBasicInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateDataSourceBasicInfoRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataSourceBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoRequest) SetOpTenantId(v int64) *UpdateDataSourceBasicInfoRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataSourceBasicInfoRequest) SetUpdateCommand(v *UpdateDataSourceBasicInfoRequestUpdateCommand) *UpdateDataSourceBasicInfoRequest {
	s.UpdateCommand = v
	return s
}

type UpdateDataSourceBasicInfoRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23231
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDataSourceBasicInfoRequestUpdateCommand) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceBasicInfoRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) SetDescription(v string) *UpdateDataSourceBasicInfoRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) SetId(v int64) *UpdateDataSourceBasicInfoRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateDataSourceBasicInfoRequestUpdateCommand) SetName(v string) *UpdateDataSourceBasicInfoRequestUpdateCommand {
	s.Name = &v
	return s
}

type UpdateDataSourceBasicInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataSourceBasicInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoShrinkRequest) SetOpTenantId(v int64) *UpdateDataSourceBasicInfoShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataSourceBasicInfoShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataSourceBasicInfoShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

type UpdateDataSourceBasicInfoResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataSourceBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetCode(v string) *UpdateDataSourceBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetData(v bool) *UpdateDataSourceBasicInfoResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetHttpStatusCode(v int32) *UpdateDataSourceBasicInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetMessage(v string) *UpdateDataSourceBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetRequestId(v string) *UpdateDataSourceBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponseBody) SetSuccess(v bool) *UpdateDataSourceBasicInfoResponseBody {
	s.Success = &v
	return s
}

type UpdateDataSourceBasicInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourceBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourceBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceBasicInfoResponse) SetHeaders(v map[string]*string) *UpdateDataSourceBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourceBasicInfoResponse) SetStatusCode(v int32) *UpdateDataSourceBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourceBasicInfoResponse) SetBody(v *UpdateDataSourceBasicInfoResponseBody) *UpdateDataSourceBasicInfoResponse {
	s.Body = v
	return s
}

type UpdateDataSourceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateDataSourceConfigRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataSourceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigRequest) SetOpTenantId(v int64) *UpdateDataSourceConfigRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataSourceConfigRequest) SetUpdateCommand(v *UpdateDataSourceConfigRequestUpdateCommand) *UpdateDataSourceConfigRequest {
	s.UpdateCommand = v
	return s
}

type UpdateDataSourceConfigRequestUpdateCommand struct {
	// This parameter is required.
	ConfigItemList []*UpdateDataSourceConfigRequestUpdateCommandConfigItemList `json:"ConfigItemList,omitempty" xml:"ConfigItemList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 13231313
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateDataSourceConfigRequestUpdateCommand) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceConfigRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigRequestUpdateCommand) SetConfigItemList(v []*UpdateDataSourceConfigRequestUpdateCommandConfigItemList) *UpdateDataSourceConfigRequestUpdateCommand {
	s.ConfigItemList = v
	return s
}

func (s *UpdateDataSourceConfigRequestUpdateCommand) SetId(v int64) *UpdateDataSourceConfigRequestUpdateCommand {
	s.Id = &v
	return s
}

type UpdateDataSourceConfigRequestUpdateCommandConfigItemList struct {
	// This parameter is required.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDataSourceConfigRequestUpdateCommandConfigItemList) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceConfigRequestUpdateCommandConfigItemList) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigRequestUpdateCommandConfigItemList) SetKey(v string) *UpdateDataSourceConfigRequestUpdateCommandConfigItemList {
	s.Key = &v
	return s
}

func (s *UpdateDataSourceConfigRequestUpdateCommandConfigItemList) SetValue(v string) *UpdateDataSourceConfigRequestUpdateCommandConfigItemList {
	s.Value = &v
	return s
}

type UpdateDataSourceConfigShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateDataSourceConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigShrinkRequest) SetOpTenantId(v int64) *UpdateDataSourceConfigShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataSourceConfigShrinkRequest) SetUpdateCommandShrink(v string) *UpdateDataSourceConfigShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

type UpdateDataSourceConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataSourceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigResponseBody) SetCode(v string) *UpdateDataSourceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetData(v bool) *UpdateDataSourceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetHttpStatusCode(v int32) *UpdateDataSourceConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetMessage(v string) *UpdateDataSourceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetRequestId(v string) *UpdateDataSourceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceConfigResponseBody) SetSuccess(v bool) *UpdateDataSourceConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateDataSourceConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigResponse) SetHeaders(v map[string]*string) *UpdateDataSourceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourceConfigResponse) SetStatusCode(v int32) *UpdateDataSourceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourceConfigResponse) SetBody(v *UpdateDataSourceConfigResponseBody) *UpdateDataSourceConfigResponse {
	s.Body = v
	return s
}

type UpdateFileDirectoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /xx测试/目录new
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12121111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12132323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateFileDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileDirectoryRequest) SetDirectory(v string) *UpdateFileDirectoryRequest {
	s.Directory = &v
	return s
}

func (s *UpdateFileDirectoryRequest) SetFileId(v int64) *UpdateFileDirectoryRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileDirectoryRequest) SetOpTenantId(v int64) *UpdateFileDirectoryRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateFileDirectoryRequest) SetProjectId(v int64) *UpdateFileDirectoryRequest {
	s.ProjectId = &v
	return s
}

type UpdateFileDirectoryResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFileDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileDirectoryResponseBody) SetCode(v string) *UpdateFileDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) SetHttpStatusCode(v int32) *UpdateFileDirectoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) SetMessage(v string) *UpdateFileDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) SetRequestId(v string) *UpdateFileDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileDirectoryResponseBody) SetSuccess(v bool) *UpdateFileDirectoryResponseBody {
	s.Success = &v
	return s
}

type UpdateFileDirectoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileDirectoryResponse) SetHeaders(v map[string]*string) *UpdateFileDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileDirectoryResponse) SetStatusCode(v int32) *UpdateFileDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileDirectoryResponse) SetBody(v *UpdateFileDirectoryResponseBody) *UpdateFileDirectoryResponse {
	s.Body = v
	return s
}

type UpdateFileNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12121111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxNew
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12132323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateFileNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileNameRequest) SetFileId(v int64) *UpdateFileNameRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileNameRequest) SetName(v string) *UpdateFileNameRequest {
	s.Name = &v
	return s
}

func (s *UpdateFileNameRequest) SetOpTenantId(v int64) *UpdateFileNameRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateFileNameRequest) SetProjectId(v int64) *UpdateFileNameRequest {
	s.ProjectId = &v
	return s
}

type UpdateFileNameResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFileNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileNameResponseBody) SetCode(v string) *UpdateFileNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFileNameResponseBody) SetHttpStatusCode(v int32) *UpdateFileNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFileNameResponseBody) SetMessage(v string) *UpdateFileNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFileNameResponseBody) SetRequestId(v string) *UpdateFileNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileNameResponseBody) SetSuccess(v bool) *UpdateFileNameResponseBody {
	s.Success = &v
	return s
}

type UpdateFileNameResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileNameResponse) SetHeaders(v map[string]*string) *UpdateFileNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileNameResponse) SetStatusCode(v int32) *UpdateFileNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileNameResponse) SetBody(v *UpdateFileNameResponseBody) *UpdateFileNameResponse {
	s.Body = v
	return s
}

type UpdateTenantMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateTenantMemberRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateTenantMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberRequest) SetOpTenantId(v int64) *UpdateTenantMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateTenantMemberRequest) SetUpdateCommand(v *UpdateTenantMemberRequestUpdateCommand) *UpdateTenantMemberRequest {
	s.UpdateCommand = v
	return s
}

type UpdateTenantMemberRequestUpdateCommand struct {
	// This parameter is required.
	MemberList []*UpdateTenantMemberRequestUpdateCommandMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s UpdateTenantMemberRequestUpdateCommand) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantMemberRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberRequestUpdateCommand) SetMemberList(v []*UpdateTenantMemberRequestUpdateCommandMemberList) *UpdateTenantMemberRequestUpdateCommand {
	s.MemberList = v
	return s
}

type UpdateTenantMemberRequestUpdateCommandMemberList struct {
	// example:
	//
	// 123@dingding
	DingNumber *string `json:"DingNumber,omitempty" xml:"DingNumber,omitempty"`
	// example:
	//
	// 123@xx.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// example:
	//
	// 13888888888
	MobilePhone *string   `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	RoleList    []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// example:
	//
	// 2331
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateTenantMemberRequestUpdateCommandMemberList) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantMemberRequestUpdateCommandMemberList) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetDingNumber(v string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.DingNumber = &v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetMail(v string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.Mail = &v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetMobilePhone(v string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.MobilePhone = &v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetRoleList(v []*string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.RoleList = v
	return s
}

func (s *UpdateTenantMemberRequestUpdateCommandMemberList) SetUserId(v string) *UpdateTenantMemberRequestUpdateCommandMemberList {
	s.UserId = &v
	return s
}

type UpdateTenantMemberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateTenantMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberShrinkRequest) SetOpTenantId(v int64) *UpdateTenantMemberShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateTenantMemberShrinkRequest) SetUpdateCommandShrink(v string) *UpdateTenantMemberShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

type UpdateTenantMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTenantMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberResponseBody) SetCode(v string) *UpdateTenantMemberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) SetHttpStatusCode(v int32) *UpdateTenantMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) SetMessage(v string) *UpdateTenantMemberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) SetRequestId(v string) *UpdateTenantMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) SetSuccess(v bool) *UpdateTenantMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateTenantMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTenantMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTenantMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberResponse) SetHeaders(v map[string]*string) *UpdateTenantMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateTenantMemberResponse) SetStatusCode(v int32) *UpdateTenantMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTenantMemberResponse) SetBody(v *UpdateTenantMemberResponseBody) *UpdateTenantMemberResponse {
	s.Body = v
	return s
}

type UpdateUserGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId    *int64                               `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	UpdateCommand *UpdateUserGroupRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) SetOpTenantId(v int64) *UpdateUserGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUpdateCommand(v *UpdateUserGroupRequestUpdateCommand) *UpdateUserGroupRequest {
	s.UpdateCommand = v
	return s
}

type UpdateUserGroupRequestUpdateCommand struct {
	AdminUserIdList []*string `json:"AdminUserIdList,omitempty" xml:"AdminUserIdList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13423
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateUserGroupRequestUpdateCommand) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequestUpdateCommand) SetAdminUserIdList(v []*string) *UpdateUserGroupRequestUpdateCommand {
	s.AdminUserIdList = v
	return s
}

func (s *UpdateUserGroupRequestUpdateCommand) SetDescription(v string) *UpdateUserGroupRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateUserGroupRequestUpdateCommand) SetId(v string) *UpdateUserGroupRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateUserGroupRequestUpdateCommand) SetName(v string) *UpdateUserGroupRequestUpdateCommand {
	s.Name = &v
	return s
}

type UpdateUserGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId          *int64  `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s UpdateUserGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupShrinkRequest) SetOpTenantId(v int64) *UpdateUserGroupShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateUserGroupShrinkRequest) SetUpdateCommandShrink(v string) *UpdateUserGroupShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

type UpdateUserGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *UpdateUserGroupResponseBody) SetCode(v string) *UpdateUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetData(v bool) *UpdateUserGroupResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetHttpStatusCode(v int32) *UpdateUserGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetMessage(v string) *UpdateUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
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

type UpdateUserGroupSwitchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 31323
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s UpdateUserGroupSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupSwitchRequest) SetActive(v bool) *UpdateUserGroupSwitchRequest {
	s.Active = &v
	return s
}

func (s *UpdateUserGroupSwitchRequest) SetOpTenantId(v int64) *UpdateUserGroupSwitchRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateUserGroupSwitchRequest) SetUserGroupId(v string) *UpdateUserGroupSwitchRequest {
	s.UserGroupId = &v
	return s
}

type UpdateUserGroupSwitchResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserGroupSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupSwitchResponseBody) SetCode(v string) *UpdateUserGroupSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetData(v bool) *UpdateUserGroupSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetHttpStatusCode(v int32) *UpdateUserGroupSwitchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetMessage(v string) *UpdateUserGroupSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetRequestId(v string) *UpdateUserGroupSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupSwitchResponseBody) SetSuccess(v bool) *UpdateUserGroupSwitchResponseBody {
	s.Success = &v
	return s
}

type UpdateUserGroupSwitchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserGroupSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserGroupSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupSwitchResponse) SetHeaders(v map[string]*string) *UpdateUserGroupSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserGroupSwitchResponse) SetStatusCode(v int32) *UpdateUserGroupSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserGroupSwitchResponse) SetBody(v *UpdateUserGroupSwitchResponseBody) *UpdateUserGroupSwitchResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dataphin-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 新增租户成员
//
// @param tmpReq - AddTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersResponse
func (client *Client) AddTenantMembersWithOptions(tmpReq *AddTenantMembersRequest, runtime *util.RuntimeOptions) (_result *AddTenantMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddTenantMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AddCommand)) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, tea.String("AddCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddCommandShrink)) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTenantMembers"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTenantMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增租户成员
//
// @param request - AddTenantMembersRequest
//
// @return AddTenantMembersResponse
func (client *Client) AddTenantMembers(request *AddTenantMembersRequest) (_result *AddTenantMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTenantMembersResponse{}
	_body, _err := client.AddTenantMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过原始用户添加租户成员.
//
// @param tmpReq - AddTenantMembersBySourceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersBySourceUserResponse
func (client *Client) AddTenantMembersBySourceUserWithOptions(tmpReq *AddTenantMembersBySourceUserRequest, runtime *util.RuntimeOptions) (_result *AddTenantMembersBySourceUserResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddTenantMembersBySourceUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AddCommand)) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, tea.String("AddCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddCommandShrink)) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTenantMembersBySourceUser"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTenantMembersBySourceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过原始用户添加租户成员.
//
// @param request - AddTenantMembersBySourceUserRequest
//
// @return AddTenantMembersBySourceUserResponse
func (client *Client) AddTenantMembersBySourceUser(request *AddTenantMembersBySourceUserRequest) (_result *AddTenantMembersBySourceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTenantMembersBySourceUserResponse{}
	_body, _err := client.AddTenantMembersBySourceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加用户组成员.
//
// @param tmpReq - AddUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMemberWithOptions(tmpReq *AddUserGroupMemberRequest, runtime *util.RuntimeOptions) (_result *AddUserGroupMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddUserGroupMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AddCommand)) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, tea.String("AddCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddCommandShrink)) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserGroupMember"),
		Version:     tea.String("2023-06-30"),
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
// 添加用户组成员.
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
// 检查数据源连通性
//
// @param tmpReq - CheckDataSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityResponse
func (client *Client) CheckDataSourceConnectivityWithOptions(tmpReq *CheckDataSourceConnectivityRequest, runtime *util.RuntimeOptions) (_result *CheckDataSourceConnectivityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CheckDataSourceConnectivityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CheckCommand)) {
		request.CheckCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckCommand, tea.String("CheckCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckCommandShrink)) {
		body["CheckCommand"] = request.CheckCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDataSourceConnectivity"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDataSourceConnectivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查数据源连通性
//
// @param request - CheckDataSourceConnectivityRequest
//
// @return CheckDataSourceConnectivityResponse
func (client *Client) CheckDataSourceConnectivity(request *CheckDataSourceConnectivityRequest) (_result *CheckDataSourceConnectivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDataSourceConnectivityResponse{}
	_body, _err := client.CheckDataSourceConnectivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查已创建的数据源是否正常连通
//
// @param request - CheckDataSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityByIdResponse
func (client *Client) CheckDataSourceConnectivityByIdWithOptions(request *CheckDataSourceConnectivityByIdRequest, runtime *util.RuntimeOptions) (_result *CheckDataSourceConnectivityByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDataSourceConnectivityById"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDataSourceConnectivityByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查已创建的数据源是否正常连通
//
// @param request - CheckDataSourceConnectivityByIdRequest
//
// @return CheckDataSourceConnectivityByIdResponse
func (client *Client) CheckDataSourceConnectivityById(request *CheckDataSourceConnectivityByIdRequest) (_result *CheckDataSourceConnectivityByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDataSourceConnectivityByIdResponse{}
	_body, _err := client.CheckDataSourceConnectivityByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验用户是否有指定资源权限点.
//
// @param tmpReq - CheckResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourcePermissionResponse
func (client *Client) CheckResourcePermissionWithOptions(tmpReq *CheckResourcePermissionRequest, runtime *util.RuntimeOptions) (_result *CheckResourcePermissionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CheckResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CheckCommand)) {
		request.CheckCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckCommand, tea.String("CheckCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckCommandShrink)) {
		body["CheckCommand"] = request.CheckCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckResourcePermission"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckResourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验用户是否有指定资源权限点.
//
// @param request - CheckResourcePermissionRequest
//
// @return CheckResourcePermissionResponse
func (client *Client) CheckResourcePermission(request *CheckResourcePermissionRequest) (_result *CheckResourcePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckResourcePermissionResponse{}
	_body, _err := client.CheckResourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建即席查询文件
//
// @param tmpReq - CreateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdHocFileResponse
func (client *Client) CreateAdHocFileWithOptions(tmpReq *CreateAdHocFileRequest, runtime *util.RuntimeOptions) (_result *CreateAdHocFileResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAdHocFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CreateCommand)) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, tea.String("CreateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateCommandShrink)) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAdHocFile"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAdHocFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建即席查询文件
//
// @param request - CreateAdHocFileRequest
//
// @return CreateAdHocFileResponse
func (client *Client) CreateAdHocFile(request *CreateAdHocFileRequest) (_result *CreateAdHocFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAdHocFileResponse{}
	_body, _err := client.CreateAdHocFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建数据源
//
// @param tmpReq - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(tmpReq *CreateDataSourceRequest, runtime *util.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDataSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CreateCommand)) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, tea.String("CreateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateCommandShrink)) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataSource"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建数据源
//
// @param request - CreateDataSourceRequest
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建菜单树文件目录
//
// @param tmpReq - CreateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectoryWithOptions(tmpReq *CreateDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDirectoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CreateCommand)) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, tea.String("CreateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateCommandShrink)) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDirectory"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建菜单树文件目录
//
// @param request - CreateDirectoryRequest
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectory(request *CreateDirectoryRequest) (_result *CreateDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CreateDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用补数据接口 1.会生成补数据实例运行：影响相关产产出表数据 2.会进行任务运行：造成计算的费用以及存储的费用
//
// @param tmpReq - CreateNodeSupplementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeSupplementResponse
func (client *Client) CreateNodeSupplementWithOptions(tmpReq *CreateNodeSupplementRequest, runtime *util.RuntimeOptions) (_result *CreateNodeSupplementResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateNodeSupplementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CreateCommand)) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, tea.String("CreateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateCommandShrink)) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNodeSupplement"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodeSupplementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用补数据接口 1.会生成补数据实例运行：影响相关产产出表数据 2.会进行任务运行：造成计算的费用以及存储的费用
//
// @param request - CreateNodeSupplementRequest
//
// @return CreateNodeSupplementResponse
func (client *Client) CreateNodeSupplement(request *CreateNodeSupplementRequest) (_result *CreateNodeSupplementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeSupplementResponse{}
	_body, _err := client.CreateNodeSupplementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建用户组.
//
// @param tmpReq - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithOptions(tmpReq *CreateUserGroupRequest, runtime *util.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateUserGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CreateCommand)) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, tea.String("CreateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateCommandShrink)) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserGroup"),
		Version:     tea.String("2023-06-30"),
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
// 新建用户组.
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
// 删除菜单树即席查询文件
//
// @param request - DeleteAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdHocFileResponse
func (client *Client) DeleteAdHocFileWithOptions(request *DeleteAdHocFileRequest, runtime *util.RuntimeOptions) (_result *DeleteAdHocFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAdHocFile"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAdHocFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除菜单树即席查询文件
//
// @param request - DeleteAdHocFileRequest
//
// @return DeleteAdHocFileResponse
func (client *Client) DeleteAdHocFile(request *DeleteAdHocFileRequest) (_result *DeleteAdHocFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAdHocFileResponse{}
	_body, _err := client.DeleteAdHocFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据源
//
// @param tmpReq - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(tmpReq *DeleteDataSourceRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteDataSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeleteCommand)) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, tea.String("DeleteCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteCommandShrink)) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSource"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据源
//
// @param request - DeleteDataSourceRequest
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除菜单树文件目录
//
// @param request - DeleteDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectoryWithOptions(request *DeleteDirectoryRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDirectory"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除菜单树文件目录
//
// @param request - DeleteDirectoryRequest
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectory(request *DeleteDirectoryRequest) (_result *DeleteDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.DeleteDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户组.
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
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserGroup"),
		Version:     tea.String("2023-06-30"),
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
// 删除用户组.
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
// 运行手动调度节点。
//
// @param tmpReq - ExecuteManualNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteManualNodeResponse
func (client *Client) ExecuteManualNodeWithOptions(tmpReq *ExecuteManualNodeRequest, runtime *util.RuntimeOptions) (_result *ExecuteManualNodeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteManualNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExecuteCommand)) {
		request.ExecuteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecuteCommand, tea.String("ExecuteCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecuteCommandShrink)) {
		body["ExecuteCommand"] = request.ExecuteCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteManualNode"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteManualNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行手动调度节点。
//
// @param request - ExecuteManualNodeRequest
//
// @return ExecuteManualNodeResponse
func (client *Client) ExecuteManualNode(request *ExecuteManualNodeRequest) (_result *ExecuteManualNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteManualNodeResponse{}
	_body, _err := client.ExecuteManualNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重跑下游(修复链路数据), 支持强制重跑下游。影响范围: 1. 会产生计算成本；2. 会影响数据产出
//
// @param tmpReq - FixDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FixDataResponse
func (client *Client) FixDataWithOptions(tmpReq *FixDataRequest, runtime *util.RuntimeOptions) (_result *FixDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FixDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FixDataCommand)) {
		request.FixDataCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FixDataCommand, tea.String("FixDataCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FixDataCommandShrink)) {
		body["FixDataCommand"] = request.FixDataCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FixData"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FixDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重跑下游(修复链路数据), 支持强制重跑下游。影响范围: 1. 会产生计算成本；2. 会影响数据产出
//
// @param request - FixDataRequest
//
// @return FixDataResponse
func (client *Client) FixData(request *FixDataRequest) (_result *FixDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FixDataResponse{}
	_body, _err := client.FixDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询即席查询文件。
//
// @param request - GetAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocFileResponse
func (client *Client) GetAdHocFileWithOptions(request *GetAdHocFileRequest, runtime *util.RuntimeOptions) (_result *GetAdHocFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAdHocFile"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdHocFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询即席查询文件。
//
// @param request - GetAdHocFileRequest
//
// @return GetAdHocFileResponse
func (client *Client) GetAdHocFile(request *GetAdHocFileRequest) (_result *GetAdHocFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAdHocFileResponse{}
	_body, _err := client.GetAdHocFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询开发态对象上游依赖。
//
// @param request - GetDevObjectDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDevObjectDependencyResponse
func (client *Client) GetDevObjectDependencyWithOptions(request *GetDevObjectDependencyRequest, runtime *util.RuntimeOptions) (_result *GetDevObjectDependencyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectFrom)) {
		query["ObjectFrom"] = request.ObjectFrom
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		query["ObjectId"] = request.ObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDevObjectDependency"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDevObjectDependencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询开发态对象上游依赖。
//
// @param request - GetDevObjectDependencyRequest
//
// @return GetDevObjectDependencyResponse
func (client *Client) GetDevObjectDependency(request *GetDevObjectDependencyRequest) (_result *GetDevObjectDependencyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDevObjectDependencyResponse{}
	_body, _err := client.GetDevObjectDependencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据起始的实例查询该实例的下游
//
// @param tmpReq - GetInstanceDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceDownStreamResponse
func (client *Client) GetInstanceDownStreamWithOptions(tmpReq *GetInstanceDownStreamRequest, runtime *util.RuntimeOptions) (_result *GetInstanceDownStreamResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetInstanceDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceGet)) {
		request.InstanceGetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceGet, tea.String("InstanceGet"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DownStreamDepth)) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.RunStatus)) {
		query["RunStatus"] = request.RunStatus
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceGetShrink)) {
		body["InstanceGet"] = request.InstanceGetShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceDownStream"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceDownStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起始的实例查询该实例的下游
//
// @param request - GetInstanceDownStreamRequest
//
// @return GetInstanceDownStreamResponse
func (client *Client) GetInstanceDownStream(request *GetInstanceDownStreamRequest) (_result *GetInstanceDownStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceDownStreamResponse{}
	_body, _err := client.GetInstanceDownStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例的上下游，支持逻辑表和代码任务。
//
// @param tmpReq - GetInstanceUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceUpDownStreamResponse
func (client *Client) GetInstanceUpDownStreamWithOptions(tmpReq *GetInstanceUpDownStreamRequest, runtime *util.RuntimeOptions) (_result *GetInstanceUpDownStreamResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetInstanceUpDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceId)) {
		request.InstanceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceId, tea.String("InstanceId"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DownStreamDepth)) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.UpStreamDepth)) {
		query["UpStreamDepth"] = request.UpStreamDepth
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdShrink)) {
		body["InstanceId"] = request.InstanceIdShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceUpDownStream"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceUpDownStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例的上下游，支持逻辑表和代码任务。
//
// @param request - GetInstanceUpDownStreamRequest
//
// @return GetInstanceUpDownStreamResponse
func (client *Client) GetInstanceUpDownStream(request *GetInstanceUpDownStreamRequest) (_result *GetInstanceUpDownStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceUpDownStreamResponse{}
	_body, _err := client.GetInstanceUpDownStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - GetMyRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyRolesResponse
func (client *Client) GetMyRolesWithOptions(request *GetMyRolesRequest, runtime *util.RuntimeOptions) (_result *GetMyRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMyRoles"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMyRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - GetMyRolesRequest
//
// @return GetMyRolesResponse
func (client *Client) GetMyRoles(request *GetMyRolesRequest) (_result *GetMyRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMyRolesResponse{}
	_body, _err := client.GetMyRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前用户归属租户.
//
// @param tmpReq - GetMyTenantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyTenantsResponse
func (client *Client) GetMyTenantsWithOptions(tmpReq *GetMyTenantsRequest, runtime *util.RuntimeOptions) (_result *GetMyTenantsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetMyTenantsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FeatureCodeList)) {
		request.FeatureCodeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FeatureCodeList, tea.String("FeatureCodeList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FeatureCodeListShrink)) {
		body["FeatureCodeList"] = request.FeatureCodeListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMyTenants"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMyTenantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前用户归属租户.
//
// @param request - GetMyTenantsRequest
//
// @return GetMyTenantsResponse
func (client *Client) GetMyTenants(request *GetMyTenantsRequest) (_result *GetMyTenantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMyTenantsResponse{}
	_body, _err := client.GetMyTenantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用查询节点上下游接口
//
// @param tmpReq - GetNodeUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeUpDownStreamResponse
func (client *Client) GetNodeUpDownStreamWithOptions(tmpReq *GetNodeUpDownStreamRequest, runtime *util.RuntimeOptions) (_result *GetNodeUpDownStreamResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetNodeUpDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeId)) {
		request.NodeIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeId, tea.String("NodeId"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DownStreamDepth)) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.UpStreamDepth)) {
		query["UpStreamDepth"] = request.UpStreamDepth
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeIdShrink)) {
		body["NodeId"] = request.NodeIdShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeUpDownStream"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeUpDownStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用查询节点上下游接口
//
// @param request - GetNodeUpDownStreamRequest
//
// @return GetNodeUpDownStreamResponse
func (client *Client) GetNodeUpDownStream(request *GetNodeUpDownStreamRequest) (_result *GetNodeUpDownStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeUpDownStreamResponse{}
	_body, _err := client.GetNodeUpDownStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询补数据提交的状态
//
// @param request - GetOperationSubmitStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationSubmitStatusResponse
func (client *Client) GetOperationSubmitStatusWithOptions(request *GetOperationSubmitStatusRequest, runtime *util.RuntimeOptions) (_result *GetOperationSubmitStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOperationSubmitStatus"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOperationSubmitStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询补数据提交的状态
//
// @param request - GetOperationSubmitStatusRequest
//
// @return GetOperationSubmitStatusResponse
func (client *Client) GetOperationSubmitStatus(request *GetOperationSubmitStatusRequest) (_result *GetOperationSubmitStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOperationSubmitStatusResponse{}
	_body, _err := client.GetOperationSubmitStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询脚本的实例信息, 包括实例状态、运行时间等信息.
//
// @param request - GetPhysicalInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceResponse
func (client *Client) GetPhysicalInstanceWithOptions(request *GetPhysicalInstanceRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhysicalInstance"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhysicalInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询脚本的实例信息, 包括实例状态、运行时间等信息.
//
// @param request - GetPhysicalInstanceRequest
//
// @return GetPhysicalInstanceResponse
func (client *Client) GetPhysicalInstance(request *GetPhysicalInstanceRequest) (_result *GetPhysicalInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalInstanceResponse{}
	_body, _err := client.GetPhysicalInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例执行的日志，如果实例重跑了多次，则会有多条日志
//
// @param request - GetPhysicalInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceLogResponse
func (client *Client) GetPhysicalInstanceLogWithOptions(request *GetPhysicalInstanceLogRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalInstanceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhysicalInstanceLog"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhysicalInstanceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例执行的日志，如果实例重跑了多次，则会有多条日志
//
// @param request - GetPhysicalInstanceLogRequest
//
// @return GetPhysicalInstanceLogResponse
func (client *Client) GetPhysicalInstanceLog(request *GetPhysicalInstanceLogRequest) (_result *GetPhysicalInstanceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalInstanceLogResponse{}
	_body, _err := client.GetPhysicalInstanceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询物理调度节点。
//
// @param request - GetPhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeResponse
func (client *Client) GetPhysicalNodeWithOptions(request *GetPhysicalNodeRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhysicalNode"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhysicalNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询物理调度节点。
//
// @param request - GetPhysicalNodeRequest
//
// @return GetPhysicalNodeResponse
func (client *Client) GetPhysicalNode(request *GetPhysicalNodeRequest) (_result *GetPhysicalNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalNodeResponse{}
	_body, _err := client.GetPhysicalNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据输出名查询对应的物理节点。
//
// @param request - GetPhysicalNodeByOutputNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeByOutputNameResponse
func (client *Client) GetPhysicalNodeByOutputNameWithOptions(request *GetPhysicalNodeByOutputNameRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalNodeByOutputNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.OutputName)) {
		query["OutputName"] = request.OutputName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhysicalNodeByOutputName"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhysicalNodeByOutputNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据输出名查询对应的物理节点。
//
// @param request - GetPhysicalNodeByOutputNameRequest
//
// @return GetPhysicalNodeByOutputNameResponse
func (client *Client) GetPhysicalNodeByOutputName(request *GetPhysicalNodeByOutputNameRequest) (_result *GetPhysicalNodeByOutputNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalNodeByOutputNameResponse{}
	_body, _err := client.GetPhysicalNodeByOutputNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询调度节点代码内容。
//
// @param request - GetPhysicalNodeContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeContentResponse
func (client *Client) GetPhysicalNodeContentWithOptions(request *GetPhysicalNodeContentRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalNodeContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhysicalNodeContent"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhysicalNodeContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调度节点代码内容。
//
// @param request - GetPhysicalNodeContentRequest
//
// @return GetPhysicalNodeContentResponse
func (client *Client) GetPhysicalNodeContent(request *GetPhysicalNodeContentRequest) (_result *GetPhysicalNodeContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalNodeContentResponse{}
	_body, _err := client.GetPhysicalNodeContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询节点的操作日志。
//
// @param request - GetPhysicalNodeOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeOperationLogResponse
func (client *Client) GetPhysicalNodeOperationLogWithOptions(request *GetPhysicalNodeOperationLogRequest, runtime *util.RuntimeOptions) (_result *GetPhysicalNodeOperationLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhysicalNodeOperationLog"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhysicalNodeOperationLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询节点的操作日志。
//
// @param request - GetPhysicalNodeOperationLogRequest
//
// @return GetPhysicalNodeOperationLogResponse
func (client *Client) GetPhysicalNodeOperationLog(request *GetPhysicalNodeOperationLogRequest) (_result *GetPhysicalNodeOperationLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhysicalNodeOperationLogResponse{}
	_body, _err := client.GetPhysicalNodeOperationLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取项目生产账号
//
// @param request - GetProjectProduceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectProduceUserResponse
func (client *Client) GetProjectProduceUserWithOptions(request *GetProjectProduceUserRequest, runtime *util.RuntimeOptions) (_result *GetProjectProduceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectProduceUser"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectProduceUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目生产账号
//
// @param request - GetProjectProduceUserRequest
//
// @return GetProjectProduceUserResponse
func (client *Client) GetProjectProduceUser(request *GetProjectProduceUserRequest) (_result *GetProjectProduceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectProduceUserResponse{}
	_body, _err := client.GetProjectProduceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取补数据工作流所有业务日期的Dagrun信息。
//
// @param request - GetSupplementDagrunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunResponse
func (client *Client) GetSupplementDagrunWithOptions(request *GetSupplementDagrunRequest, runtime *util.RuntimeOptions) (_result *GetSupplementDagrunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.SupplementId)) {
		query["SupplementId"] = request.SupplementId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSupplementDagrun"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSupplementDagrunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取补数据工作流所有业务日期的Dagrun信息。
//
// @param request - GetSupplementDagrunRequest
//
// @return GetSupplementDagrunResponse
func (client *Client) GetSupplementDagrun(request *GetSupplementDagrunRequest) (_result *GetSupplementDagrunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSupplementDagrunResponse{}
	_body, _err := client.GetSupplementDagrunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出补数据工作流下具体一个业务日期的所有节点的实例。
//
// @param request - GetSupplementDagrunInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunInstanceResponse
func (client *Client) GetSupplementDagrunInstanceWithOptions(request *GetSupplementDagrunInstanceRequest, runtime *util.RuntimeOptions) (_result *GetSupplementDagrunInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DagrunId)) {
		query["DagrunId"] = request.DagrunId
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSupplementDagrunInstance"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSupplementDagrunInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出补数据工作流下具体一个业务日期的所有节点的实例。
//
// @param request - GetSupplementDagrunInstanceRequest
//
// @return GetSupplementDagrunInstanceResponse
func (client *Client) GetSupplementDagrunInstance(request *GetSupplementDagrunInstanceRequest) (_result *GetSupplementDagrunInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSupplementDagrunInstanceResponse{}
	_body, _err := client.GetSupplementDagrunInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过用户原始Id（如阿里云Id）获取用户详情
//
// @param request - GetUserBySourceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBySourceIdResponse
func (client *Client) GetUserBySourceIdWithOptions(request *GetUserBySourceIdRequest, runtime *util.RuntimeOptions) (_result *GetUserBySourceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserBySourceId"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserBySourceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过用户原始Id（如阿里云Id）获取用户详情
//
// @param request - GetUserBySourceIdRequest
//
// @return GetUserBySourceIdResponse
func (client *Client) GetUserBySourceId(request *GetUserBySourceIdRequest) (_result *GetUserBySourceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserBySourceIdResponse{}
	_body, _err := client.GetUserBySourceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户组详情.
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, runtime *util.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGroup"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户组详情.
//
// @param request - GetUserGroupRequest
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroup(request *GetUserGroupRequest) (_result *GetUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserGroupResponse{}
	_body, _err := client.GetUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户详情
//
// @param tmpReq - GetUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUsersResponse
func (client *Client) GetUsersWithOptions(tmpReq *GetUsersRequest, runtime *util.RuntimeOptions) (_result *GetUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserIdList)) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, tea.String("UserIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIdListShrink)) {
		body["UserIdList"] = request.UserIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUsers"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户详情
//
// @param request - GetUsersRequest
//
// @return GetUsersResponse
func (client *Client) GetUsers(request *GetUsersRequest) (_result *GetUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUsersResponse{}
	_body, _err := client.GetUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过资源点对用户授权
//
// @param tmpReq - GrantResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantResourcePermissionResponse
func (client *Client) GrantResourcePermissionWithOptions(tmpReq *GrantResourcePermissionRequest, runtime *util.RuntimeOptions) (_result *GrantResourcePermissionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GrantResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GrantCommand)) {
		request.GrantCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GrantCommand, tea.String("GrantCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GrantCommandShrink)) {
		body["GrantCommand"] = request.GrantCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantResourcePermission"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantResourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过资源点对用户授权
//
// @param request - GrantResourcePermissionRequest
//
// @return GrantResourcePermissionResponse
func (client *Client) GrantResourcePermission(request *GrantResourcePermissionRequest) (_result *GrantResourcePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantResourcePermissionResponse{}
	_body, _err := client.GrantResourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - ListAddableRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableRolesResponse
func (client *Client) ListAddableRolesWithOptions(request *ListAddableRolesRequest, runtime *util.RuntimeOptions) (_result *ListAddableRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAddableRoles"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAddableRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - ListAddableRolesRequest
//
// @return ListAddableRolesResponse
func (client *Client) ListAddableRoles(request *ListAddableRolesRequest) (_result *ListAddableRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAddableRolesResponse{}
	_body, _err := client.ListAddableRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取可加入租户成员列表的用户
//
// @param tmpReq - ListAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableUsersResponse
func (client *Client) ListAddableUsersWithOptions(tmpReq *ListAddableUsersRequest, runtime *util.RuntimeOptions) (_result *ListAddableUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAddableUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAddableUsers"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAddableUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取可加入租户成员列表的用户
//
// @param request - ListAddableUsersRequest
//
// @return ListAddableUsersResponse
func (client *Client) ListAddableUsers(request *ListAddableUsersRequest) (_result *ListAddableUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAddableUsersResponse{}
	_body, _err := client.ListAddableUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索数据源，所属结果包含数据源配置项
//
// @param tmpReq - ListDataSourceWithConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceWithConfigResponse
func (client *Client) ListDataSourceWithConfigWithOptions(tmpReq *ListDataSourceWithConfigRequest, runtime *util.RuntimeOptions) (_result *ListDataSourceWithConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDataSourceWithConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceWithConfig"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceWithConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索数据源，所属结果包含数据源配置项
//
// @param request - ListDataSourceWithConfigRequest
//
// @return ListDataSourceWithConfigResponse
func (client *Client) ListDataSourceWithConfig(request *ListDataSourceWithConfigRequest) (_result *ListDataSourceWithConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourceWithConfigResponse{}
	_body, _err := client.ListDataSourceWithConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 遍历菜单树目录文件。
//
// @param tmpReq - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithOptions(tmpReq *ListFilesRequest, runtime *util.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFiles"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 遍历菜单树目录文件。
//
// @param request - ListFilesRequest
//
// @return ListFilesResponse
func (client *Client) ListFiles(request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询实例。
//
// @param tmpReq - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询实例。
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询节点下游，创建补数据工作流时可以作为数据参考
//
// @param tmpReq - ListNodeDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDownStreamResponse
func (client *Client) ListNodeDownStreamWithOptions(tmpReq *ListNodeDownStreamRequest, runtime *util.RuntimeOptions) (_result *ListNodeDownStreamResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListNodeDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeDownStream"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeDownStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询节点下游，创建补数据工作流时可以作为数据参考
//
// @param request - ListNodeDownStreamRequest
//
// @return ListNodeDownStreamResponse
func (client *Client) ListNodeDownStream(request *ListNodeDownStreamRequest) (_result *ListNodeDownStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeDownStreamResponse{}
	_body, _err := client.ListNodeDownStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询调度节点列表。
//
// @param tmpReq - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(tmpReq *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调度节点列表。
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取权限操作列表
//
// @param tmpReq - ListResourcePermissionOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionOperationLogResponse
func (client *Client) ListResourcePermissionOperationLogWithOptions(tmpReq *ListResourcePermissionOperationLogRequest, runtime *util.RuntimeOptions) (_result *ListResourcePermissionOperationLogResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListResourcePermissionOperationLogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourcePermissionOperationLog"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcePermissionOperationLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取权限操作列表
//
// @param request - ListResourcePermissionOperationLogRequest
//
// @return ListResourcePermissionOperationLogResponse
func (client *Client) ListResourcePermissionOperationLog(request *ListResourcePermissionOperationLogRequest) (_result *ListResourcePermissionOperationLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcePermissionOperationLogResponse{}
	_body, _err := client.ListResourcePermissionOperationLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取权限记录列表
//
// @param tmpReq - ListResourcePermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionsResponse
func (client *Client) ListResourcePermissionsWithOptions(tmpReq *ListResourcePermissionsRequest, runtime *util.RuntimeOptions) (_result *ListResourcePermissionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListResourcePermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourcePermissions"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取权限记录列表
//
// @param request - ListResourcePermissionsRequest
//
// @return ListResourcePermissionsResponse
func (client *Client) ListResourcePermissions(request *ListResourcePermissionsRequest) (_result *ListResourcePermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcePermissionsResponse{}
	_body, _err := client.ListResourcePermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询租户成员列表
//
// @param tmpReq - ListTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantMembersResponse
func (client *Client) ListTenantMembersWithOptions(tmpReq *ListTenantMembersRequest, runtime *util.RuntimeOptions) (_result *ListTenantMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTenantMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTenantMembers"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTenantMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询租户成员列表
//
// @param request - ListTenantMembersRequest
//
// @return ListTenantMembersResponse
func (client *Client) ListTenantMembers(request *ListTenantMembersRequest) (_result *ListTenantMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTenantMembersResponse{}
	_body, _err := client.ListTenantMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户组成员列表分页查询.
//
// @param tmpReq - ListUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupMembersResponse
func (client *Client) ListUserGroupMembersWithOptions(tmpReq *ListUserGroupMembersRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListUserGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserGroupMembers"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户组成员列表分页查询.
//
// @param request - ListUserGroupMembersRequest
//
// @return ListUserGroupMembersResponse
func (client *Client) ListUserGroupMembers(request *ListUserGroupMembersRequest) (_result *ListUserGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupMembersResponse{}
	_body, _err := client.ListUserGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户组列表分页查询.
//
// @param tmpReq - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithOptions(tmpReq *ListUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListUserGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListQuery)) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, tea.String("ListQuery"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListQueryShrink)) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserGroups"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户组列表分页查询.
//
// @param request - ListUserGroupsRequest
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运维实例。
//
// @param tmpReq - OperateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateInstanceResponse
func (client *Client) OperateInstanceWithOptions(tmpReq *OperateInstanceRequest, runtime *util.RuntimeOptions) (_result *OperateInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &OperateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OperateCommand)) {
		request.OperateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperateCommand, tea.String("OperateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperateCommandShrink)) {
		body["OperateCommand"] = request.OperateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateInstance"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运维实例。
//
// @param request - OperateInstanceRequest
//
// @return OperateInstanceResponse
func (client *Client) OperateInstance(request *OperateInstanceRequest) (_result *OperateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateInstanceResponse{}
	_body, _err := client.OperateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂停物理节点调度。
//
// @param tmpReq - PausePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PausePhysicalNodeResponse
func (client *Client) PausePhysicalNodeWithOptions(tmpReq *PausePhysicalNodeRequest, runtime *util.RuntimeOptions) (_result *PausePhysicalNodeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PausePhysicalNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PauseCommand)) {
		request.PauseCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PauseCommand, tea.String("PauseCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PauseCommandShrink)) {
		body["PauseCommand"] = request.PauseCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PausePhysicalNode"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PausePhysicalNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停物理节点调度。
//
// @param request - PausePhysicalNodeRequest
//
// @return PausePhysicalNodeResponse
func (client *Client) PausePhysicalNode(request *PausePhysicalNodeRequest) (_result *PausePhysicalNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PausePhysicalNodeResponse{}
	_body, _err := client.PausePhysicalNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除租户成员
//
// @param tmpReq - RemoveTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTenantMemberResponse
func (client *Client) RemoveTenantMemberWithOptions(tmpReq *RemoveTenantMemberRequest, runtime *util.RuntimeOptions) (_result *RemoveTenantMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveTenantMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RemoveCommand)) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, tea.String("RemoveCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RemoveCommandShrink)) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTenantMember"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTenantMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除租户成员
//
// @param request - RemoveTenantMemberRequest
//
// @return RemoveTenantMemberResponse
func (client *Client) RemoveTenantMember(request *RemoveTenantMemberRequest) (_result *RemoveTenantMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTenantMemberResponse{}
	_body, _err := client.RemoveTenantMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除用户组成员.
//
// @param tmpReq - RemoveUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserGroupMemberResponse
func (client *Client) RemoveUserGroupMemberWithOptions(tmpReq *RemoveUserGroupMemberRequest, runtime *util.RuntimeOptions) (_result *RemoveUserGroupMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RemoveUserGroupMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RemoveCommand)) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, tea.String("RemoveCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RemoveCommandShrink)) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserGroupMember"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除用户组成员.
//
// @param request - RemoveUserGroupMemberRequest
//
// @return RemoveUserGroupMemberResponse
func (client *Client) RemoveUserGroupMember(request *RemoveUserGroupMemberRequest) (_result *RemoveUserGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserGroupMemberResponse{}
	_body, _err := client.RemoveUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复物理节点调度。
//
// @param tmpReq - ResumePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumePhysicalNodeResponse
func (client *Client) ResumePhysicalNodeWithOptions(tmpReq *ResumePhysicalNodeRequest, runtime *util.RuntimeOptions) (_result *ResumePhysicalNodeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ResumePhysicalNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResumeCommand)) {
		request.ResumeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResumeCommand, tea.String("ResumeCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Env)) {
		query["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResumeCommandShrink)) {
		body["ResumeCommand"] = request.ResumeCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumePhysicalNode"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumePhysicalNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复物理节点调度。
//
// @param request - ResumePhysicalNodeRequest
//
// @return ResumePhysicalNodeResponse
func (client *Client) ResumePhysicalNode(request *ResumePhysicalNodeRequest) (_result *ResumePhysicalNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumePhysicalNodeResponse{}
	_body, _err := client.ResumePhysicalNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 回收用户资源授权
//
// @param tmpReq - RevokeResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourcePermissionResponse
func (client *Client) RevokeResourcePermissionWithOptions(tmpReq *RevokeResourcePermissionRequest, runtime *util.RuntimeOptions) (_result *RevokeResourcePermissionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RevokeResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RevokeCommand)) {
		request.RevokeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RevokeCommand, tea.String("RevokeCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RevokeCommandShrink)) {
		body["RevokeCommand"] = request.RevokeCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeResourcePermission"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeResourcePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回收用户资源授权
//
// @param request - RevokeResourcePermissionRequest
//
// @return RevokeResourcePermissionResponse
func (client *Client) RevokeResourcePermission(request *RevokeResourcePermissionRequest) (_result *RevokeResourcePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeResourcePermissionResponse{}
	_body, _err := client.RevokeResourcePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑即席查询文件。
//
// @param tmpReq - UpdateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdHocFileResponse
func (client *Client) UpdateAdHocFileWithOptions(tmpReq *UpdateAdHocFileRequest, runtime *util.RuntimeOptions) (_result *UpdateAdHocFileResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAdHocFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateCommand)) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, tea.String("UpdateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateCommandShrink)) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAdHocFile"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAdHocFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑即席查询文件。
//
// @param request - UpdateAdHocFileRequest
//
// @return UpdateAdHocFileResponse
func (client *Client) UpdateAdHocFile(request *UpdateAdHocFileRequest) (_result *UpdateAdHocFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAdHocFileResponse{}
	_body, _err := client.UpdateAdHocFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑数据源基本信息
//
// @param tmpReq - UpdateDataSourceBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceBasicInfoResponse
func (client *Client) UpdateDataSourceBasicInfoWithOptions(tmpReq *UpdateDataSourceBasicInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateDataSourceBasicInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDataSourceBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateCommand)) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, tea.String("UpdateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateCommandShrink)) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataSourceBasicInfo"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataSourceBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑数据源基本信息
//
// @param request - UpdateDataSourceBasicInfoRequest
//
// @return UpdateDataSourceBasicInfoResponse
func (client *Client) UpdateDataSourceBasicInfo(request *UpdateDataSourceBasicInfoRequest) (_result *UpdateDataSourceBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDataSourceBasicInfoResponse{}
	_body, _err := client.UpdateDataSourceBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑数据源连接配置项
//
// @param tmpReq - UpdateDataSourceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceConfigResponse
func (client *Client) UpdateDataSourceConfigWithOptions(tmpReq *UpdateDataSourceConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateDataSourceConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDataSourceConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateCommand)) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, tea.String("UpdateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateCommandShrink)) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataSourceConfig"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataSourceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑数据源连接配置项
//
// @param request - UpdateDataSourceConfigRequest
//
// @return UpdateDataSourceConfigResponse
func (client *Client) UpdateDataSourceConfig(request *UpdateDataSourceConfigRequest) (_result *UpdateDataSourceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDataSourceConfigResponse{}
	_body, _err := client.UpdateDataSourceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改菜单树文件所在目录
//
// @param request - UpdateFileDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileDirectoryResponse
func (client *Client) UpdateFileDirectoryWithOptions(request *UpdateFileDirectoryRequest, runtime *util.RuntimeOptions) (_result *UpdateFileDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Directory)) {
		query["Directory"] = request.Directory
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFileDirectory"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFileDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改菜单树文件所在目录
//
// @param request - UpdateFileDirectoryRequest
//
// @return UpdateFileDirectoryResponse
func (client *Client) UpdateFileDirectory(request *UpdateFileDirectoryRequest) (_result *UpdateFileDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFileDirectoryResponse{}
	_body, _err := client.UpdateFileDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改菜单树文件名称
//
// @param request - UpdateFileNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileNameResponse
func (client *Client) UpdateFileNameWithOptions(request *UpdateFileNameRequest, runtime *util.RuntimeOptions) (_result *UpdateFileNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["FileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFileName"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFileNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改菜单树文件名称
//
// @param request - UpdateFileNameRequest
//
// @return UpdateFileNameResponse
func (client *Client) UpdateFileName(request *UpdateFileNameRequest) (_result *UpdateFileNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFileNameResponse{}
	_body, _err := client.UpdateFileNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑租户成员
//
// @param tmpReq - UpdateTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantMemberResponse
func (client *Client) UpdateTenantMemberWithOptions(tmpReq *UpdateTenantMemberRequest, runtime *util.RuntimeOptions) (_result *UpdateTenantMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTenantMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateCommand)) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, tea.String("UpdateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateCommandShrink)) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTenantMember"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTenantMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑租户成员
//
// @param request - UpdateTenantMemberRequest
//
// @return UpdateTenantMemberResponse
func (client *Client) UpdateTenantMember(request *UpdateTenantMemberRequest) (_result *UpdateTenantMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTenantMemberResponse{}
	_body, _err := client.UpdateTenantMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑用户组.
//
// @param tmpReq - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithOptions(tmpReq *UpdateUserGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateUserGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UpdateCommand)) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, tea.String("UpdateCommand"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateCommandShrink)) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserGroup"),
		Version:     tea.String("2023-06-30"),
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
// 编辑用户组.
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
// 编辑用户组启用开关.
//
// @param request - UpdateUserGroupSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupSwitchResponse
func (client *Client) UpdateUserGroupSwitchWithOptions(request *UpdateUserGroupSwitchRequest, runtime *util.RuntimeOptions) (_result *UpdateUserGroupSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		query["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.OpTenantId)) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserGroupSwitch"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserGroupSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑用户组启用开关.
//
// @param request - UpdateUserGroupSwitchRequest
//
// @return UpdateUserGroupSwitchResponse
func (client *Client) UpdateUserGroupSwitch(request *UpdateUserGroupSwitchRequest) (_result *UpdateUserGroupSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserGroupSwitchResponse{}
	_body, _err := client.UpdateUserGroupSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
