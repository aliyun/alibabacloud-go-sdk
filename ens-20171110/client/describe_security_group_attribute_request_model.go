// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeRequest
	GetSecurityGroupId() *string
}

type DescribeSecurityGroupAttributeRequest struct {
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4ph***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeSecurityGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupAttributeRequest) SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
