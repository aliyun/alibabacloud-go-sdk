// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceMyAppPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataServiceMyAppPermissionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDataServiceMyAppPermissionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataServiceMyAppPermissionsResponseBody
	GetMessage() *string
	SetPageResult(v *ListDataServiceMyAppPermissionsResponseBodyPageResult) *ListDataServiceMyAppPermissionsResponseBody
	GetPageResult() *ListDataServiceMyAppPermissionsResponseBodyPageResult
	SetRequestId(v string) *ListDataServiceMyAppPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceMyAppPermissionsResponseBody
	GetSuccess() *bool
}

type ListDataServiceMyAppPermissionsResponseBody struct {
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
	// internal error
	Message    *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListDataServiceMyAppPermissionsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceMyAppPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyAppPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyAppPermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataServiceMyAppPermissionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceMyAppPermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataServiceMyAppPermissionsResponseBody) GetPageResult() *ListDataServiceMyAppPermissionsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDataServiceMyAppPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceMyAppPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceMyAppPermissionsResponseBody) SetCode(v string) *ListDataServiceMyAppPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBody) SetHttpStatusCode(v int32) *ListDataServiceMyAppPermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBody) SetMessage(v string) *ListDataServiceMyAppPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBody) SetPageResult(v *ListDataServiceMyAppPermissionsResponseBodyPageResult) *ListDataServiceMyAppPermissionsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBody) SetRequestId(v string) *ListDataServiceMyAppPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBody) SetSuccess(v bool) *ListDataServiceMyAppPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceMyAppPermissionsResponseBodyPageResult struct {
	PermissionList []*ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList `json:"PermissionList,omitempty" xml:"PermissionList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceMyAppPermissionsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyAppPermissionsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResult) GetPermissionList() []*ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	return s.PermissionList
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResult) SetPermissionList(v []*ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) *ListDataServiceMyAppPermissionsResponseBodyPageResult {
	s.PermissionList = v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResult) SetTotalCount(v int32) *ListDataServiceMyAppPermissionsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResult) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList struct {
	// AppId
	//
	// example:
	//
	// 1021
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 1121
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// NormalUser
	CurrentUserRole *string `json:"CurrentUserRole,omitempty" xml:"CurrentUserRole,omitempty"`
	// example:
	//
	// 1121
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// test
	OwnerUserName *string `json:"OwnerUserName,omitempty" xml:"OwnerUserName,omitempty"`
	// example:
	//
	// 1121
	PrivilegeBelongTo *string `json:"PrivilegeBelongTo,omitempty" xml:"PrivilegeBelongTo,omitempty"`
	// example:
	//
	// 0
	PrivilegeFrom *int32 `json:"PrivilegeFrom,omitempty" xml:"PrivilegeFrom,omitempty"`
	// example:
	//
	// 112101
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName        *string                                                                                  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RemarkForDebugList []*ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList `json:"RemarkForDebugList,omitempty" xml:"RemarkForDebugList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetAppId() *int32 {
	return s.AppId
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetAppName() *string {
	return s.AppName
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetCreator() *string {
	return s.Creator
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetCurrentUserRole() *string {
	return s.CurrentUserRole
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetOwner() *string {
	return s.Owner
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetOwnerUserName() *string {
	return s.OwnerUserName
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetPrivilegeBelongTo() *string {
	return s.PrivilegeBelongTo
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetPrivilegeFrom() *int32 {
	return s.PrivilegeFrom
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetRemarkForDebugList() []*ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList {
	return s.RemarkForDebugList
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) GetRole() *int32 {
	return s.Role
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetAppId(v int32) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.AppId = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetAppName(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.AppName = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetCreateUserName(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.CreateUserName = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetCreator(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.Creator = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetCurrentUserRole(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.CurrentUserRole = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetOwner(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.Owner = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetOwnerUserName(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.OwnerUserName = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetPrivilegeBelongTo(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.PrivilegeBelongTo = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetPrivilegeFrom(v int32) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.PrivilegeFrom = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetProjectId(v int32) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetProjectName(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.ProjectName = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetRemarkForDebugList(v []*ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.RemarkForDebugList = v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) SetRole(v int32) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList {
	s.Role = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionList) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList struct {
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList) GetKey() *string {
	return s.Key
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList) GetValue() *string {
	return s.Value
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList) SetKey(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList {
	s.Key = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList) SetValue(v string) *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList {
	s.Value = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsResponseBodyPageResultPermissionListRemarkForDebugList) Validate() error {
	return dara.Validate(s)
}
