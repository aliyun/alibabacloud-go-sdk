// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceLinkedRole(v string) *DescribeServiceLinkedRoleStatusRequest
	GetServiceLinkedRole() *string
}

type DescribeServiceLinkedRoleStatusRequest struct {
	// example:
	//
	// AliyunServiceRoleForOssMfd
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
}

func (s DescribeServiceLinkedRoleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *DescribeServiceLinkedRoleStatusRequest) SetServiceLinkedRole(v string) *DescribeServiceLinkedRoleStatusRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusRequest) Validate() error {
	return dara.Validate(s)
}
