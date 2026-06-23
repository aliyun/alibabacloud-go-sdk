// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasServiceLinkedRole(v bool) *CheckServiceLinkedRoleResponseBody
	GetHasServiceLinkedRole() *bool
	SetRequestId(v string) *CheckServiceLinkedRoleResponseBody
	GetRequestId() *string
}

type CheckServiceLinkedRoleResponseBody struct {
	// Indicates whether the service-linked role exists.
	//
	// example:
	//
	// true
	HasServiceLinkedRole *bool `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 66428721-xxxx-xxxx-xxxx-3BD1B67837E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) GetHasServiceLinkedRole() *bool {
	return s.HasServiceLinkedRole
}

func (s *CheckServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckServiceLinkedRoleResponseBody) SetHasServiceLinkedRole(v bool) *CheckServiceLinkedRoleResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
