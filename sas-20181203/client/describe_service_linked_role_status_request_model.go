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
	// The service-linked role. Default value: **AliyunServiceRoleForSas**. Valid values:
	//
	// 	- **AliyunServiceRoleForSas**: the service-linked role of Security Center. Security Center assumes this role to access the resources of other cloud services within your account.
	//
	// 	- **AliyunServiceRoleForSasCspm**: the service-linked role of Security Center-CSPM. Security Center-CSPM assumes this role to access the resources of other cloud services within your account.
	//
	// example:
	//
	// AliyunServiceRoleForSas
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
