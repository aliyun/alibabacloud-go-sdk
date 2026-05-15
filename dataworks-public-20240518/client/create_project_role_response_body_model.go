// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateProjectRoleResponseBody
	GetCode() *string
	SetRequestId(v string) *CreateProjectRoleResponseBody
	GetRequestId() *string
}

type CreateProjectRoleResponseBody struct {
	// example:
	//
	// base_role_dte
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// B313938A-4475-599B-98EB-A0875019FD5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateProjectRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectRoleResponseBody) SetCode(v string) *CreateProjectRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProjectRoleResponseBody) SetRequestId(v string) *CreateProjectRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
