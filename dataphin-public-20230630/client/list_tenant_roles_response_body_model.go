// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTenantRolesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTenantRolesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTenantRolesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTenantRolesResponseBody
	GetRequestId() *string
	SetRoleList(v []*ListTenantRolesResponseBodyRoleList) *ListTenantRolesResponseBody
	GetRoleList() []*ListTenantRolesResponseBodyRoleList
	SetSuccess(v bool) *ListTenantRolesResponseBody
	GetSuccess() *bool
}

type ListTenantRolesResponseBody struct {
	// The request error code. OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code of the backend response.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message of the request.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role list.
	RoleList []*ListTenantRolesResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTenantRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTenantRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantRolesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTenantRolesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTenantRolesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTenantRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTenantRolesResponseBody) GetRoleList() []*ListTenantRolesResponseBodyRoleList {
	return s.RoleList
}

func (s *ListTenantRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTenantRolesResponseBody) SetCode(v string) *ListTenantRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantRolesResponseBody) SetHttpStatusCode(v int32) *ListTenantRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTenantRolesResponseBody) SetMessage(v string) *ListTenantRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantRolesResponseBody) SetRequestId(v string) *ListTenantRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantRolesResponseBody) SetRoleList(v []*ListTenantRolesResponseBodyRoleList) *ListTenantRolesResponseBody {
	s.RoleList = v
	return s
}

func (s *ListTenantRolesResponseBody) SetSuccess(v bool) *ListTenantRolesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTenantRolesResponseBody) Validate() error {
	if s.RoleList != nil {
		for _, item := range s.RoleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTenantRolesResponseBodyRoleList struct {
	// The JSON string of the permission set.
	//
	// example:
	//
	// {}
	AuthJson *string `json:"AuthJson,omitempty" xml:"AuthJson,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 30121201
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-07 16:45:30
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-01-07 16:45:30
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The user ID of the modifier.
	//
	// example:
	//
	// 30121201
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The role description.
	//
	// example:
	//
	// test
	RoleDesc *string `json:"RoleDesc,omitempty" xml:"RoleDesc,omitempty"`
	// The role key.
	//
	// example:
	//
	// test::1212
	RoleKey *string `json:"RoleKey,omitempty" xml:"RoleKey,omitempty"`
	// The role name.
	//
	// example:
	//
	// test
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The role type. Valid values:
	//
	// - BUILD_IN: built-in role
	//
	// - CUSTOM: custom role
	//
	// example:
	//
	// CUSTOM
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The role status. Valid values:
	//
	// - ON: enabled
	//
	// - OFF: disabled
	//
	// - DEL: deleted
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 30121001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The tenant type.
	//
	// example:
	//
	// CUSTOM
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
}

func (s ListTenantRolesResponseBodyRoleList) String() string {
	return dara.Prettify(s)
}

func (s ListTenantRolesResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *ListTenantRolesResponseBodyRoleList) GetAuthJson() *string {
	return s.AuthJson
}

func (s *ListTenantRolesResponseBodyRoleList) GetCreator() *string {
	return s.Creator
}

func (s *ListTenantRolesResponseBodyRoleList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListTenantRolesResponseBodyRoleList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListTenantRolesResponseBodyRoleList) GetModifier() *string {
	return s.Modifier
}

func (s *ListTenantRolesResponseBodyRoleList) GetRoleDesc() *string {
	return s.RoleDesc
}

func (s *ListTenantRolesResponseBodyRoleList) GetRoleKey() *string {
	return s.RoleKey
}

func (s *ListTenantRolesResponseBodyRoleList) GetRoleName() *string {
	return s.RoleName
}

func (s *ListTenantRolesResponseBodyRoleList) GetRoleType() *string {
	return s.RoleType
}

func (s *ListTenantRolesResponseBodyRoleList) GetStatus() *string {
	return s.Status
}

func (s *ListTenantRolesResponseBodyRoleList) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListTenantRolesResponseBodyRoleList) GetTenantType() *string {
	return s.TenantType
}

func (s *ListTenantRolesResponseBodyRoleList) SetAuthJson(v string) *ListTenantRolesResponseBodyRoleList {
	s.AuthJson = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetCreator(v string) *ListTenantRolesResponseBodyRoleList {
	s.Creator = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetGmtCreate(v string) *ListTenantRolesResponseBodyRoleList {
	s.GmtCreate = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetGmtModified(v string) *ListTenantRolesResponseBodyRoleList {
	s.GmtModified = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetModifier(v string) *ListTenantRolesResponseBodyRoleList {
	s.Modifier = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetRoleDesc(v string) *ListTenantRolesResponseBodyRoleList {
	s.RoleDesc = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetRoleKey(v string) *ListTenantRolesResponseBodyRoleList {
	s.RoleKey = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetRoleName(v string) *ListTenantRolesResponseBodyRoleList {
	s.RoleName = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetRoleType(v string) *ListTenantRolesResponseBodyRoleList {
	s.RoleType = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetStatus(v string) *ListTenantRolesResponseBodyRoleList {
	s.Status = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetTenantId(v int64) *ListTenantRolesResponseBodyRoleList {
	s.TenantId = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) SetTenantType(v string) *ListTenantRolesResponseBodyRoleList {
	s.TenantType = &v
	return s
}

func (s *ListTenantRolesResponseBodyRoleList) Validate() error {
	return dara.Validate(s)
}
