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
	SetRegionId(v string) *CheckServiceLinkedRoleResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CheckServiceLinkedRoleResponseBody
	GetRequestId() *string
}

type CheckServiceLinkedRoleResponseBody struct {
	// Indicates whether an SLR is created.
	//
	// example:
	//
	// true
	HasServiceLinkedRole *string `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CheckServiceLinkedRoleResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckServiceLinkedRoleResponseBody) SetHasServiceLinkedRole(v string) *CheckServiceLinkedRoleResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRegionId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
