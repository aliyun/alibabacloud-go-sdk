// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleTagStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRoleTagStatusResponseBody
	GetRequestId() *string
	SetRoleTagStatus(v string) *DescribeRoleTagStatusResponseBody
	GetRoleTagStatus() *string
	SetShardRoleTagStatus(v map[string]interface{}) *DescribeRoleTagStatusResponseBody
	GetShardRoleTagStatus() map[string]interface{}
}

type DescribeRoleTagStatusResponseBody struct {
	// example:
	//
	// 69D55DEC-B219-569F-A686-AC2F67A1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	RoleTagStatus *string `json:"RoleTagStatus,omitempty" xml:"RoleTagStatus,omitempty"`
	// example:
	//
	// {
	//
	//     "d-2ze204b4db58****": "false",
	//
	//     "d-2zeb9716646e***": "false"
	//
	// }
	ShardRoleTagStatus map[string]interface{} `json:"ShardRoleTagStatus,omitempty" xml:"ShardRoleTagStatus,omitempty"`
}

func (s DescribeRoleTagStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleTagStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoleTagStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoleTagStatusResponseBody) GetRoleTagStatus() *string {
	return s.RoleTagStatus
}

func (s *DescribeRoleTagStatusResponseBody) GetShardRoleTagStatus() map[string]interface{} {
	return s.ShardRoleTagStatus
}

func (s *DescribeRoleTagStatusResponseBody) SetRequestId(v string) *DescribeRoleTagStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoleTagStatusResponseBody) SetRoleTagStatus(v string) *DescribeRoleTagStatusResponseBody {
	s.RoleTagStatus = &v
	return s
}

func (s *DescribeRoleTagStatusResponseBody) SetShardRoleTagStatus(v map[string]interface{}) *DescribeRoleTagStatusResponseBody {
	s.ShardRoleTagStatus = v
	return s
}

func (s *DescribeRoleTagStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
