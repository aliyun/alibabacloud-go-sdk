// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceJobId(v string) *UpdateWorkspaceRoleResponseBody
	GetInstanceJobId() *string
	SetRequestId(v string) *UpdateWorkspaceRoleResponseBody
	GetRequestId() *string
}

type UpdateWorkspaceRoleResponseBody struct {
	// The task ID.
	//
	// example:
	//
	// UpdateWorkspaceCustomRole-role-***abc*******
	InstanceJobId *string `json:"InstanceJobId,omitempty" xml:"InstanceJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A519F77D-28A0-52F5-AB82-5********8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkspaceRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRoleResponseBody) GetInstanceJobId() *string {
	return s.InstanceJobId
}

func (s *UpdateWorkspaceRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceRoleResponseBody) SetInstanceJobId(v string) *UpdateWorkspaceRoleResponseBody {
	s.InstanceJobId = &v
	return s
}

func (s *UpdateWorkspaceRoleResponseBody) SetRequestId(v string) *UpdateWorkspaceRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
