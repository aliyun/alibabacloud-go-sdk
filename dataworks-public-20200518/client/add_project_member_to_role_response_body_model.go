// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMemberToRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddProjectMemberToRoleResponseBody
	GetRequestId() *string
}

type AddProjectMemberToRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddProjectMemberToRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberToRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectMemberToRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddProjectMemberToRoleResponseBody) SetRequestId(v string) *AddProjectMemberToRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProjectMemberToRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
