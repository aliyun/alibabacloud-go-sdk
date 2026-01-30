// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServiceLinkedRoleStatusResponseBody
	GetRequestId() *string
	SetRoleStatus(v *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus) *DescribeServiceLinkedRoleStatusResponseBody
	GetRoleStatus() *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus
}

type DescribeServiceLinkedRoleStatusResponseBody struct {
	// example:
	//
	// E99D386F-5546-5BCD-BADD-F4EF4B******
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleStatus *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus `json:"RoleStatus,omitempty" xml:"RoleStatus,omitempty" type:"Struct"`
}

func (s DescribeServiceLinkedRoleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) GetRoleStatus() *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus {
	return s.RoleStatus
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) SetRequestId(v string) *DescribeServiceLinkedRoleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) SetRoleStatus(v *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus) *DescribeServiceLinkedRoleStatusResponseBody {
	s.RoleStatus = v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) Validate() error {
	if s.RoleStatus != nil {
		if err := s.RoleStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceLinkedRoleStatusResponseBodyRoleStatus struct {
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeServiceLinkedRoleStatusResponseBodyRoleStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusResponseBodyRoleStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus) GetStatus() *bool {
	return s.Status
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus) SetStatus(v bool) *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus {
	s.Status = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyRoleStatus) Validate() error {
	return dara.Validate(s)
}
