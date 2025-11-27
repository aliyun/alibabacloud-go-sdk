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
	// Indicates whether an SLR is created.
	//
	// example:
	//
	// true
	HasServiceLinkedRole *string `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AB44DC0A-7E77-442A-97A9-C6418694CB22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the service-linked role is required. Default value: true.
	//
	// example:
	//
	// true
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
