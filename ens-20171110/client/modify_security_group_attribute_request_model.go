// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifySecurityGroupAttributeRequest
	GetDescription() *string
	SetSecurityGroupId(v string) *ModifySecurityGroupAttributeRequest
	GetSecurityGroupId() *string
	SetSecurityGroupName(v string) *ModifySecurityGroupAttributeRequest
	GetSecurityGroupName() *string
}

type ModifySecurityGroupAttributeRequest struct {
	// The description of the security group.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the security group. The name of a bucket must meet the following requirements:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name must start with a letter but cannot start with http:// or https://.
	//
	// 	- The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// example
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s ModifySecurityGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupAttributeRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifySecurityGroupAttributeRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *ModifySecurityGroupAttributeRequest) SetDescription(v string) *ModifySecurityGroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetSecurityGroupId(v string) *ModifySecurityGroupAttributeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetSecurityGroupName(v string) *ModifySecurityGroupAttributeRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
