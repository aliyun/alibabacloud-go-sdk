// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasServiceLinkedRole(v string) *CheckServiceLinkedRoleResponseBody
	GetHasServiceLinkedRole() *string
	SetRequestId(v string) *CheckServiceLinkedRoleResponseBody
	GetRequestId() *string
	SetRequireServiceLinkedRole(v string) *CheckServiceLinkedRoleResponseBody
	GetRequireServiceLinkedRole() *string
}

type CheckServiceLinkedRoleResponseBody struct {
	HasServiceLinkedRole     *string `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequireServiceLinkedRole *string `json:"RequireServiceLinkedRole,omitempty" xml:"RequireServiceLinkedRole,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) GetHasServiceLinkedRole() *string {
	return s.HasServiceLinkedRole
}

func (s *CheckServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckServiceLinkedRoleResponseBody) GetRequireServiceLinkedRole() *string {
	return s.RequireServiceLinkedRole
}

func (s *CheckServiceLinkedRoleResponseBody) SetHasServiceLinkedRole(v string) *CheckServiceLinkedRoleResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequireServiceLinkedRole(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequireServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
