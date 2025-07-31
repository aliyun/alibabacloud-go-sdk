// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *DescribeAuditLogFilterResponseBody
	GetFilter() *string
	SetRequestId(v string) *DescribeAuditLogFilterResponseBody
	GetRequestId() *string
	SetRoleType(v string) *DescribeAuditLogFilterResponseBody
	GetRoleType() *string
}

type DescribeAuditLogFilterResponseBody struct {
	// The type of the audit log entries. Valid values:
	//
	// 	- **admin**: O\\&M and management operations
	//
	// 	- **slow**: slow query logs
	//
	// 	- **query**: query operations
	//
	// 	- **insert**: insert operations
	//
	// 	- **update**: update operations
	//
	// 	- **delete**: delete operations
	//
	// 	- **command**: protocol commands such as the aggregate method
	//
	// example:
	//
	// admin,slow,insert,query,update,delete,command
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7BAFB0B3-2A54-5B65-B13E-3937CF08FEE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role of the node.
	//
	// example:
	//
	// logic
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAuditLogFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogFilterResponseBody) GetFilter() *string {
	return s.Filter
}

func (s *DescribeAuditLogFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditLogFilterResponseBody) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeAuditLogFilterResponseBody) SetFilter(v string) *DescribeAuditLogFilterResponseBody {
	s.Filter = &v
	return s
}

func (s *DescribeAuditLogFilterResponseBody) SetRequestId(v string) *DescribeAuditLogFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditLogFilterResponseBody) SetRoleType(v string) *DescribeAuditLogFilterResponseBody {
	s.RoleType = &v
	return s
}

func (s *DescribeAuditLogFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
