// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleExistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExistsServiceLinkedRole(v bool) *DescribeServiceLinkedRoleExistsResponseBody
	GetExistsServiceLinkedRole() *bool
	SetRequestId(v string) *DescribeServiceLinkedRoleExistsResponseBody
	GetRequestId() *string
}

type DescribeServiceLinkedRoleExistsResponseBody struct {
	// Indicates whether the service-linked role has been authorized for the current account. Valid values:
	//
	// - **true**: Authorized.
	//
	// - **false**: Not authorized.
	//
	// example:
	//
	// true
	ExistsServiceLinkedRole *bool `json:"ExistsServiceLinkedRole,omitempty" xml:"ExistsServiceLinkedRole,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90B82DB7-FB28-4CC2-ADBF-1F8659F3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceLinkedRoleExistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleExistsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleExistsResponseBody) GetExistsServiceLinkedRole() *bool {
	return s.ExistsServiceLinkedRole
}

func (s *DescribeServiceLinkedRoleExistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceLinkedRoleExistsResponseBody) SetExistsServiceLinkedRole(v bool) *DescribeServiceLinkedRoleExistsResponseBody {
	s.ExistsServiceLinkedRole = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsResponseBody) SetRequestId(v string) *DescribeServiceLinkedRoleExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsResponseBody) Validate() error {
	return dara.Validate(s)
}
