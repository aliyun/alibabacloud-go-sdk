// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceJobId(v string) *CreateWorkspaceRoleResponseBody
	GetInstanceJobId() *string
	SetRequestId(v string) *CreateWorkspaceRoleResponseBody
	GetRequestId() *string
}

type CreateWorkspaceRoleResponseBody struct {
	// The job ID for the request.
	//
	// example:
	//
	// CreateWorkspaceCustomRole-role-***abc*******
	InstanceJobId *string `json:"InstanceJobId,omitempty" xml:"InstanceJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A519F77D-28A0-52F5-AB82-5********8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWorkspaceRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRoleResponseBody) GetInstanceJobId() *string {
	return s.InstanceJobId
}

func (s *CreateWorkspaceRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceRoleResponseBody) SetInstanceJobId(v string) *CreateWorkspaceRoleResponseBody {
	s.InstanceJobId = &v
	return s
}

func (s *CreateWorkspaceRoleResponseBody) SetRequestId(v string) *CreateWorkspaceRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
