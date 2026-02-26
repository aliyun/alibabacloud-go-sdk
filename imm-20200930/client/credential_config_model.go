// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialConfig interface {
	dara.Model
	String() string
	GoString() string
	SetChain(v []*CredentialConfigChain) *CredentialConfig
	GetChain() []*CredentialConfigChain
	SetPolicy(v string) *CredentialConfig
	GetPolicy() *string
	SetServiceRole(v string) *CredentialConfig
	GetServiceRole() *string
}

type CredentialConfig struct {
	// The authorization chains. All roles in the array must have the `sts:AssumeRole` permission. You need to only grant other permissions, such as read and write permissions on OSS, to the last role in the array. You can grant permissions in the RAM console.
	Chain []*CredentialConfigChain `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Repeated"`
	// The policy that is attached to the role specified by the ServiceRole parameter. For example, the policy allows access to OSS. This parameter is optional.
	//
	// example:
	//
	// {"Statement": [{"Action": "oss:*","Effect": "Allow","Resource": "*"}],"Version": "1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The service role in the account that is used to call an IMM API operation. The role must have the `sts:AssumeRole` permission. You can configure permissions for the role in the Resource Access Management (RAM) console.
	//
	// example:
	//
	// AliyunSTSAssumeForIMMServiceRole
	ServiceRole *string `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
}

func (s CredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfig) GoString() string {
	return s.String()
}

func (s *CredentialConfig) GetChain() []*CredentialConfigChain {
	return s.Chain
}

func (s *CredentialConfig) GetPolicy() *string {
	return s.Policy
}

func (s *CredentialConfig) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *CredentialConfig) SetChain(v []*CredentialConfigChain) *CredentialConfig {
	s.Chain = v
	return s
}

func (s *CredentialConfig) SetPolicy(v string) *CredentialConfig {
	s.Policy = &v
	return s
}

func (s *CredentialConfig) SetServiceRole(v string) *CredentialConfig {
	s.ServiceRole = &v
	return s
}

func (s *CredentialConfig) Validate() error {
	if s.Chain != nil {
		for _, item := range s.Chain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CredentialConfigChain struct {
	// The ID of the account that you use to grant permissions.
	//
	// example:
	//
	// 10232100246xxxxx
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// The RAM role that can be assumed.
	//
	// example:
	//
	// AliyunOSSRole
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The role type. Valid values:
	//
	// 	- user: Alibaba Cloud account.
	//
	// 	- service: Alibaba Cloud service.
	//
	// example:
	//
	// user
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CredentialConfigChain) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfigChain) GoString() string {
	return s.String()
}

func (s *CredentialConfigChain) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *CredentialConfigChain) GetRole() *string {
	return s.Role
}

func (s *CredentialConfigChain) GetRoleType() *string {
	return s.RoleType
}

func (s *CredentialConfigChain) SetAssumeRoleFor(v string) *CredentialConfigChain {
	s.AssumeRoleFor = &v
	return s
}

func (s *CredentialConfigChain) SetRole(v string) *CredentialConfigChain {
	s.Role = &v
	return s
}

func (s *CredentialConfigChain) SetRoleType(v string) *CredentialConfigChain {
	s.RoleType = &v
	return s
}

func (s *CredentialConfigChain) Validate() error {
	return dara.Validate(s)
}
