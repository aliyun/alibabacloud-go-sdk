// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubAccountUserId(v string) *DescribeUserPermissionsRequest
	GetSubAccountUserId() *string
}

type DescribeUserPermissionsRequest struct {
	// The ID of a RAM user or RAM role.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27852573609480****
	SubAccountUserId *string `json:"SubAccountUserId,omitempty" xml:"SubAccountUserId,omitempty"`
}

func (s DescribeUserPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsRequest) GetSubAccountUserId() *string {
	return s.SubAccountUserId
}

func (s *DescribeUserPermissionsRequest) SetSubAccountUserId(v string) *DescribeUserPermissionsRequest {
	s.SubAccountUserId = &v
	return s
}

func (s *DescribeUserPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
