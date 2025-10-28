// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InsertRoleResponseBody
	GetCode() *int32
	SetMessage(v string) *InsertRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertRoleResponseBody
	GetRequestId() *string
	SetRoleId(v int32) *InsertRoleResponseBody
	GetRoleId() *int32
}

type InsertRoleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57609587-DFA2-41EC-****-*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the role.
	//
	// example:
	//
	// 33
	RoleId *int32 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s InsertRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InsertRoleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertRoleResponseBody) GetRoleId() *int32 {
	return s.RoleId
}

func (s *InsertRoleResponseBody) SetCode(v int32) *InsertRoleResponseBody {
	s.Code = &v
	return s
}

func (s *InsertRoleResponseBody) SetMessage(v string) *InsertRoleResponseBody {
	s.Message = &v
	return s
}

func (s *InsertRoleResponseBody) SetRequestId(v string) *InsertRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertRoleResponseBody) SetRoleId(v int32) *InsertRoleResponseBody {
	s.RoleId = &v
	return s
}

func (s *InsertRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
