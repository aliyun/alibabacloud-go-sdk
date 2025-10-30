// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleAuthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthStatus(v bool) *DescribeRoleAuthStatusResponseBody
	GetAuthStatus() *bool
	SetRequestId(v string) *DescribeRoleAuthStatusResponseBody
	GetRequestId() *string
}

type DescribeRoleAuthStatusResponseBody struct {
	// example:
	//
	// true
	AuthStatus *bool `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// example:
	//
	// 79ECBB08-079C-57C5-A676-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRoleAuthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleAuthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoleAuthStatusResponseBody) GetAuthStatus() *bool {
	return s.AuthStatus
}

func (s *DescribeRoleAuthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoleAuthStatusResponseBody) SetAuthStatus(v bool) *DescribeRoleAuthStatusResponseBody {
	s.AuthStatus = &v
	return s
}

func (s *DescribeRoleAuthStatusResponseBody) SetRequestId(v string) *DescribeRoleAuthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoleAuthStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
