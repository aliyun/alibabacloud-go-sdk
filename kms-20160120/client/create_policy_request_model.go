// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlRules(v string) *CreatePolicyRequest
	GetAccessControlRules() *string
	SetDescription(v string) *CreatePolicyRequest
	GetDescription() *string
	SetKmsInstance(v string) *CreatePolicyRequest
	GetKmsInstance() *string
	SetName(v string) *CreatePolicyRequest
	GetName() *string
	SetPermissions(v string) *CreatePolicyRequest
	GetPermissions() *string
	SetResources(v string) *CreatePolicyRequest
	GetResources() *string
}

type CreatePolicyRequest struct {
	// The name of the access control rule.
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
	// The scope of the permission policy. You need to specify the KMS instance that you want to access.
	//
	// example:
	//
	// kst-hzz634e67d126u9p9****
	KmsInstance *string `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty"`
	// The name of the permission policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operations that can be performed. Valid values:
	//
	// 	- RbacPermission/Template/CryptoServiceKeyUser: allows you to perform cryptographic operations.
	//
	// 	- RbacPermission/Template/CryptoServiceSecretUser: allows you to perform secret-related operations.
	//
	// You can select both.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["RbacPermission/Template/CryptoServiceKeyUser", "RbacPermission/Template/CryptoServiceSecretUser"]
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The key and secret that are allowed to access.
	//
	// 	- Key: Enter a key in the `key/${KeyId}` format. To allow access to all keys of a KMS instance, enter key/\\*.
	//
	// 	- Secret: Enter a secret in the `secret/${SecretName}` format. To allow access to all secrets of a KMS instance, enter secret/\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["secret/acs/ram/user/ram-secret", "secret/acs/ram/user/acr-master", "key/key-hzz63d9c8d3dfv8cv****"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetAccessControlRules() *string {
	return s.AccessControlRules
}

func (s *CreatePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyRequest) GetKmsInstance() *string {
	return s.KmsInstance
}

func (s *CreatePolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreatePolicyRequest) GetPermissions() *string {
	return s.Permissions
}

func (s *CreatePolicyRequest) GetResources() *string {
	return s.Resources
}

func (s *CreatePolicyRequest) SetAccessControlRules(v string) *CreatePolicyRequest {
	s.AccessControlRules = &v
	return s
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetKmsInstance(v string) *CreatePolicyRequest {
	s.KmsInstance = &v
	return s
}

func (s *CreatePolicyRequest) SetName(v string) *CreatePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyRequest) SetPermissions(v string) *CreatePolicyRequest {
	s.Permissions = &v
	return s
}

func (s *CreatePolicyRequest) SetResources(v string) *CreatePolicyRequest {
	s.Resources = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	return dara.Validate(s)
}
