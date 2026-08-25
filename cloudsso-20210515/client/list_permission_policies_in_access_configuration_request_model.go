// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionPoliciesInAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *ListPermissionPoliciesInAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *ListPermissionPoliciesInAccessConfigurationRequest
	GetDirectoryId() *string
	SetPermissionPolicyType(v string) *ListPermissionPoliciesInAccessConfigurationRequest
	GetPermissionPolicyType() *string
}

type ListPermissionPoliciesInAccessConfigurationRequest struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The type of the policy. The type can be used to filter policies. Valid values:
	//
	// 	- System: system policy.
	//
	// 	- Inline: inline policy.
	//
	// If you do not specify this parameter, all types of policies are queried.
	//
	// example:
	//
	// System
	PermissionPolicyType *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
}

func (s ListPermissionPoliciesInAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionPoliciesInAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) GetPermissionPolicyType() *string {
	return s.PermissionPolicyType
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) SetAccessConfigurationId(v string) *ListPermissionPoliciesInAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) SetDirectoryId(v string) *ListPermissionPoliciesInAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) SetPermissionPolicyType(v string) *ListPermissionPoliciesInAccessConfigurationRequest {
	s.PermissionPolicyType = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
