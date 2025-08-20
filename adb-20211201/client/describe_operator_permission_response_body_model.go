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
	SetDBClusterId(v string) *DescribeOperatorPermissionResponseBody
	GetDBClusterId() *string
	SetExpiredTime(v string) *DescribeOperatorPermissionResponseBody
	GetExpiredTime() *string
	SetPrivileges(v string) *DescribeOperatorPermissionResponseBody
	GetPrivileges() *string
	SetRequestId(v string) *DescribeOperatorPermissionResponseBody
	GetRequestId() *string
}

type DescribeOperatorPermissionResponseBody struct {
	// The time when the permissions take effect.
	//
	// example:
	//
	// 2024-02-25T03:35:02Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-uf6wjk5xxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The time when the permissions expire.
	//
	// example:
	//
	// 2024-01-10T02:19:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The queried permissions.
	//
	// example:
	//
	// Control,Data
	Privileges *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
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

func (s *DescribeOperatorPermissionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DescribeOperatorPermissionResponseBody) SetDBClusterId(v string) *DescribeOperatorPermissionResponseBody {
	s.DBClusterId = &v
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
