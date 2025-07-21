// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlRules(v string) *UpdatePolicyRequest
	GetAccessControlRules() *string
	SetDescription(v string) *UpdatePolicyRequest
	GetDescription() *string
	SetName(v string) *UpdatePolicyRequest
	GetName() *string
	SetPermissions(v string) *UpdatePolicyRequest
	GetPermissions() *string
	SetResources(v string) *UpdatePolicyRequest
	GetResources() *string
}

type UpdatePolicyRequest struct {
	// The access control rule.
	//
	// > For more information about how to query created access control rules, see [ListNetworkRules](https://help.aliyun.com/document_detail/2539433.html).
	//
	// example:
	//
	// {"NetworkRules":["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]}
	AccessControlRules *string `json:"AccessControlRules,omitempty" xml:"AccessControlRules,omitempty"`
	// The description.
	//
	// example:
	//
	// policy  description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission policy that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operations that are supported by the updated policy. Valid values:
	//
	// 	- RbacPermission/Template/CryptoServiceKeyUser: allows you to perform cryptographic operations.
	//
	// 	- RbacPermission/Template/CryptoServiceSecretUser: allows you to perform secret-related operations.
	//
	// You can select both.
	//
	// example:
	//
	// ["RbacPermission/Template/CryptoServiceKeyUser", "RbacPermission/Template/CryptoServiceSecretUser"]
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The key and secret that are allowed to access after the update.
	//
	// 	- Key: Enter a key in the `key/${KeyId}` format. To allow access to all keys of a KMS instance, enter key/\\*.
	//
	// 	- Secret: Enter a secret in the `secret/${SecretName}` format. To allow access to all secrets of a KMS instance, enter secret/\\*.
	//
	// example:
	//
	// ["secret/acs/ram/user/ram-secret", "secret/acs/ram/user/acr-master", "key/key-hzz63d9c8d3dfv8cv****"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s UpdatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequest) GetAccessControlRules() *string {
	return s.AccessControlRules
}

func (s *UpdatePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePolicyRequest) GetPermissions() *string {
	return s.Permissions
}

func (s *UpdatePolicyRequest) GetResources() *string {
	return s.Resources
}

func (s *UpdatePolicyRequest) SetAccessControlRules(v string) *UpdatePolicyRequest {
	s.AccessControlRules = &v
	return s
}

func (s *UpdatePolicyRequest) SetDescription(v string) *UpdatePolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolicyRequest) SetName(v string) *UpdatePolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdatePolicyRequest) SetPermissions(v string) *UpdatePolicyRequest {
	s.Permissions = &v
	return s
}

func (s *UpdatePolicyRequest) SetResources(v string) *UpdatePolicyRequest {
	s.Resources = &v
	return s
}

func (s *UpdatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
