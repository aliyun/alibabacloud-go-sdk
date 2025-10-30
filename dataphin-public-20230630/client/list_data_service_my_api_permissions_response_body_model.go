// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceMyApiPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataServiceMyApiPermissionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDataServiceMyApiPermissionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataServiceMyApiPermissionsResponseBody
	GetMessage() *string
	SetPageResult(v *ListDataServiceMyApiPermissionsResponseBodyPageResult) *ListDataServiceMyApiPermissionsResponseBody
	GetPageResult() *ListDataServiceMyApiPermissionsResponseBodyPageResult
	SetRequestId(v string) *ListDataServiceMyApiPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceMyApiPermissionsResponseBody
	GetSuccess() *bool
}

type ListDataServiceMyApiPermissionsResponseBody struct {
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
	PageResult *ListDataServiceMyApiPermissionsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceMyApiPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyApiPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyApiPermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataServiceMyApiPermissionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceMyApiPermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataServiceMyApiPermissionsResponseBody) GetPageResult() *ListDataServiceMyApiPermissionsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDataServiceMyApiPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceMyApiPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceMyApiPermissionsResponseBody) SetCode(v string) *ListDataServiceMyApiPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBody) SetHttpStatusCode(v int32) *ListDataServiceMyApiPermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBody) SetMessage(v string) *ListDataServiceMyApiPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBody) SetPageResult(v *ListDataServiceMyApiPermissionsResponseBodyPageResult) *ListDataServiceMyApiPermissionsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBody) SetRequestId(v string) *ListDataServiceMyApiPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBody) SetSuccess(v bool) *ListDataServiceMyApiPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceMyApiPermissionsResponseBodyPageResult struct {
	PermissionList []*ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList `json:"PermissionList,omitempty" xml:"PermissionList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceMyApiPermissionsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyApiPermissionsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResult) GetPermissionList() []*ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	return s.PermissionList
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResult) SetPermissionList(v []*ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) *ListDataServiceMyApiPermissionsResponseBodyPageResult {
	s.PermissionList = v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResult) SetTotalCount(v int32) *ListDataServiceMyApiPermissionsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResult) Validate() error {
	if s.PermissionList != nil {
		for _, item := range s.PermissionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList struct {
	// example:
	//
	// 1322
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// teset
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
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
	// 102122
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 0
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetCreator() *string {
	return s.Creator
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetOwner() *string {
	return s.Owner
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetOwnerUserName() *string {
	return s.OwnerUserName
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetPrivilegeBelongTo() *string {
	return s.PrivilegeBelongTo
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetPrivilegeFrom() *int32 {
	return s.PrivilegeFrom
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) GetRole() *int32 {
	return s.Role
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetApiId(v int64) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetApiName(v string) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.ApiName = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetCreateUserName(v string) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.CreateUserName = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetCreator(v string) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.Creator = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetOwner(v string) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.Owner = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetOwnerUserName(v string) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.OwnerUserName = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetPrivilegeBelongTo(v string) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.PrivilegeBelongTo = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetPrivilegeFrom(v int32) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.PrivilegeFrom = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetProjectId(v int32) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetProjectName(v string) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.ProjectName = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) SetRole(v int32) *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList {
	s.Role = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsResponseBodyPageResultPermissionList) Validate() error {
	return dara.Validate(s)
}
