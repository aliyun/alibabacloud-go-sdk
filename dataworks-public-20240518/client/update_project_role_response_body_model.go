// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateProjectRoleResponseBody
	GetRequestId() *string
}

type UpdateProjectRoleResponseBody struct {
	// example:
	//
	// 037DFCE4-ABA5-51D7-9F2D-CCF709252DAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectRoleResponseBody) SetRequestId(v string) *UpdateProjectRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
