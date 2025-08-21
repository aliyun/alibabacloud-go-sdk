// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *DeleteSecurityGroupRequest
	GetSecurityGroupId() *string
}

type DeleteSecurityGroupRequest struct {
	// The security group ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4ph***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DeleteSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
