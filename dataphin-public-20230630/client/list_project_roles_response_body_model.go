// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProjectRolesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListProjectRolesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListProjectRolesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProjectRolesResponseBody
	GetRequestId() *string
	SetRoleList(v []*ListProjectRolesResponseBodyRoleList) *ListProjectRolesResponseBody
	GetRoleList() []*ListProjectRolesResponseBodyRoleList
	SetSuccess(v bool) *ListProjectRolesResponseBody
	GetSuccess() *bool
}

type ListProjectRolesResponseBody struct {
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
	RoleList  []*ListProjectRolesResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProjectRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProjectRolesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListProjectRolesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProjectRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectRolesResponseBody) GetRoleList() []*ListProjectRolesResponseBodyRoleList {
	return s.RoleList
}

func (s *ListProjectRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectRolesResponseBody) SetCode(v string) *ListProjectRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectRolesResponseBody) SetHttpStatusCode(v int32) *ListProjectRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListProjectRolesResponseBody) SetMessage(v string) *ListProjectRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectRolesResponseBody) SetRequestId(v string) *ListProjectRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectRolesResponseBody) SetRoleList(v []*ListProjectRolesResponseBodyRoleList) *ListProjectRolesResponseBody {
	s.RoleList = v
	return s
}

func (s *ListProjectRolesResponseBody) SetSuccess(v bool) *ListProjectRolesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectRolesResponseBody) Validate() error {
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

type ListProjectRolesResponseBodyRoleList struct {
	// example:
	//
	// {}
	AuthJson *string `json:"AuthJson,omitempty" xml:"AuthJson,omitempty"`
	// example:
	//
	// 30112011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2026-01-30 17:38:32
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2026-01-30 17:38:32
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 30112011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// BASIC
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	// example:
	//
	// test
	RoleDesc *string `json:"RoleDesc,omitempty" xml:"RoleDesc,omitempty"`
	// example:
	//
	// abc::01121
	RoleKey *string `json:"RoleKey,omitempty" xml:"RoleKey,omitempty"`
	// example:
	//
	// test
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// CUSTOM
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 30110110
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListProjectRolesResponseBodyRoleList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRolesResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBodyRoleList) GetAuthJson() *string {
	return s.AuthJson
}

func (s *ListProjectRolesResponseBodyRoleList) GetCreator() *string {
	return s.Creator
}

func (s *ListProjectRolesResponseBodyRoleList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListProjectRolesResponseBodyRoleList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListProjectRolesResponseBodyRoleList) GetModifier() *string {
	return s.Modifier
}

func (s *ListProjectRolesResponseBodyRoleList) GetProjectType() *string {
	return s.ProjectType
}

func (s *ListProjectRolesResponseBodyRoleList) GetRoleDesc() *string {
	return s.RoleDesc
}

func (s *ListProjectRolesResponseBodyRoleList) GetRoleKey() *string {
	return s.RoleKey
}

func (s *ListProjectRolesResponseBodyRoleList) GetRoleName() *string {
	return s.RoleName
}

func (s *ListProjectRolesResponseBodyRoleList) GetRoleType() *string {
	return s.RoleType
}

func (s *ListProjectRolesResponseBodyRoleList) GetStatus() *string {
	return s.Status
}

func (s *ListProjectRolesResponseBodyRoleList) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListProjectRolesResponseBodyRoleList) SetAuthJson(v string) *ListProjectRolesResponseBodyRoleList {
	s.AuthJson = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetCreator(v string) *ListProjectRolesResponseBodyRoleList {
	s.Creator = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetGmtCreate(v string) *ListProjectRolesResponseBodyRoleList {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetGmtModified(v string) *ListProjectRolesResponseBodyRoleList {
	s.GmtModified = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetModifier(v string) *ListProjectRolesResponseBodyRoleList {
	s.Modifier = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetProjectType(v string) *ListProjectRolesResponseBodyRoleList {
	s.ProjectType = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetRoleDesc(v string) *ListProjectRolesResponseBodyRoleList {
	s.RoleDesc = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetRoleKey(v string) *ListProjectRolesResponseBodyRoleList {
	s.RoleKey = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetRoleName(v string) *ListProjectRolesResponseBodyRoleList {
	s.RoleName = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetRoleType(v string) *ListProjectRolesResponseBodyRoleList {
	s.RoleType = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetStatus(v string) *ListProjectRolesResponseBodyRoleList {
	s.Status = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) SetTenantId(v int64) *ListProjectRolesResponseBodyRoleList {
	s.TenantId = &v
	return s
}

func (s *ListProjectRolesResponseBodyRoleList) Validate() error {
	return dara.Validate(s)
}
