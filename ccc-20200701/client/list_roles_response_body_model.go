// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRolesResponseBody
	GetCode() *string
	SetData(v []*ListRolesResponseBodyData) *ListRolesResponseBody
	GetData() []*ListRolesResponseBodyData
	SetHttpStatusCode(v int32) *ListRolesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRolesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRolesResponseBody
	GetRequestId() *string
}

type ListRolesResponseBody struct {
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListRolesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 711D948F-C616-4E23-8573-0F260513CE09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRolesResponseBody) GetData() []*ListRolesResponseBodyData {
	return s.Data
}

func (s *ListRolesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRolesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRolesResponseBody) SetCode(v string) *ListRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRolesResponseBody) SetData(v []*ListRolesResponseBodyData) *ListRolesResponseBody {
	s.Data = v
	return s
}

func (s *ListRolesResponseBody) SetHttpStatusCode(v int32) *ListRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRolesResponseBody) SetMessage(v string) *ListRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRolesResponseBodyData struct {
	// example:
	//
	// Admin
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ListRolesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListRolesResponseBodyData) GetRoleId() *string {
	return s.RoleId
}

func (s *ListRolesResponseBodyData) SetName(v string) *ListRolesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyData) SetRoleId(v string) *ListRolesResponseBodyData {
	s.RoleId = &v
	return s
}

func (s *ListRolesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
