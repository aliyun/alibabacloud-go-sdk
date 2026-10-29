// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *DescribeOperatorPermissionResponseBody
	GetCreatedTime() *string
	SetExpiredTime(v string) *DescribeOperatorPermissionResponseBody
	GetExpiredTime() *string
	SetPrivileges(v string) *DescribeOperatorPermissionResponseBody
	GetPrivileges() *string
	SetRequestId(v string) *DescribeOperatorPermissionResponseBody
	GetRequestId() *string
}

type DescribeOperatorPermissionResponseBody struct {
	// example:
	//
	// 2026-10-12T02:27:04Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2030-01-01T11:58:17Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// Control,Data
	Privileges *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	// example:
	//
	// 4C3535DC-743E-505A-B4F6-ABC***75C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOperatorPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorPermissionResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeOperatorPermissionResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeOperatorPermissionResponseBody) GetPrivileges() *string {
	return s.Privileges
}

func (s *DescribeOperatorPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOperatorPermissionResponseBody) SetCreatedTime(v string) *DescribeOperatorPermissionResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetExpiredTime(v string) *DescribeOperatorPermissionResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetPrivileges(v string) *DescribeOperatorPermissionResponseBody {
	s.Privileges = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) SetRequestId(v string) *DescribeOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperatorPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
