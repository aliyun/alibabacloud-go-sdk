// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityRoleGrant(v bool) *DescribeServiceLinkedRoleResponseBody
	GetEntityRoleGrant() *bool
	SetRequestId(v string) *DescribeServiceLinkedRoleResponseBody
	GetRequestId() *string
}

type DescribeServiceLinkedRoleResponseBody struct {
	// Whether the service-linked role permission is granted:
	//
	// - true: Granted.
	//
	// - false: Not granted.
	//
	// example:
	//
	// true
	EntityRoleGrant *bool `json:"EntityRoleGrant,omitempty" xml:"EntityRoleGrant,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 02F8BBF3-2D61-5982-8911-EEB387BE3AF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleResponseBody) GetEntityRoleGrant() *bool {
	return s.EntityRoleGrant
}

func (s *DescribeServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceLinkedRoleResponseBody) SetEntityRoleGrant(v bool) *DescribeServiceLinkedRoleResponseBody {
	s.EntityRoleGrant = &v
	return s
}

func (s *DescribeServiceLinkedRoleResponseBody) SetRequestId(v string) *DescribeServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceLinkedRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
