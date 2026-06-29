// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddableRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAddableRolesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAddableRolesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAddableRolesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAddableRolesResponseBody
	GetRequestId() *string
	SetRoleList(v []*ListAddableRolesResponseBodyRoleList) *ListAddableRolesResponseBody
	GetRoleList() []*ListAddableRolesResponseBodyRoleList
	SetSuccess(v bool) *ListAddableRolesResponseBody
	GetSuccess() *bool
}

type ListAddableRolesResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
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
	// The list of roles.
	RoleList []*ListAddableRolesResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAddableRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddableRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddableRolesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAddableRolesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAddableRolesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAddableRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddableRolesResponseBody) GetRoleList() []*ListAddableRolesResponseBodyRoleList {
	return s.RoleList
}

func (s *ListAddableRolesResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *ListAddableRolesResponseBody) Validate() error {
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

type ListAddableRolesResponseBodyRoleList struct {
	// The role code.
	//
	// example:
	//
	// SECURITY_ADMIN
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The role name.
	//
	// example:
	//
	// 安全管理员
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAddableRolesResponseBodyRoleList) String() string {
	return dara.Prettify(s)
}

func (s ListAddableRolesResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *ListAddableRolesResponseBodyRoleList) GetCode() *string {
	return s.Code
}

func (s *ListAddableRolesResponseBodyRoleList) GetName() *string {
	return s.Name
}

func (s *ListAddableRolesResponseBodyRoleList) SetCode(v string) *ListAddableRolesResponseBodyRoleList {
	s.Code = &v
	return s
}

func (s *ListAddableRolesResponseBodyRoleList) SetName(v string) *ListAddableRolesResponseBodyRoleList {
	s.Name = &v
	return s
}

func (s *ListAddableRolesResponseBodyRoleList) Validate() error {
	return dara.Validate(s)
}
