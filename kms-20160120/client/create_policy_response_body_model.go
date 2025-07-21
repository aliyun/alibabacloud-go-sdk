// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlRules(v string) *CreatePolicyResponseBody
	GetAccessControlRules() *string
	SetArn(v string) *CreatePolicyResponseBody
	GetArn() *string
	SetDescription(v string) *CreatePolicyResponseBody
	GetDescription() *string
	SetKmsInstance(v string) *CreatePolicyResponseBody
	GetKmsInstance() *string
	SetName(v string) *CreatePolicyResponseBody
	GetName() *string
	SetPermissions(v string) *CreatePolicyResponseBody
	GetPermissions() *string
	SetRequestId(v string) *CreatePolicyResponseBody
	GetRequestId() *string
	SetResources(v string) *CreatePolicyResponseBody
	GetResources() *string
}

type CreatePolicyResponseBody struct {
	// The name of the access control rule.
	//
	// example:
	//
	// {"NetworkRules":["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]}
	AccessControlRules *string `json:"AccessControlRules,omitempty" xml:"AccessControlRules,omitempty"`
	// The ARN of the permission policy.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:policy/policy_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The description.
	//
	// example:
	//
	// policy  description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The scope of the permission policy.
	//
	// example:
	//
	// kst-hzz634e67d126u9p9****
	KmsInstance *string `json:"KmsInstance,omitempty" xml:"KmsInstance,omitempty"`
	// The name of the permission policy.
	//
	// example:
	//
	// policy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operations that can be performed.
	//
	// example:
	//
	// ["RbacPermission/Template/CryptoServiceKeyUser", "RbacPermission/Template/CryptoServiceSecretUser"]
	Permissions *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f34-be0f-c4543fda2d33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The key and secret that are allowed to access.
	//
	// 	- `key/*` indicates that all keys of the KMS instance can be accessed.
	//
	// 	- `secret/*` indicates all secrets of the KMS instance can be accessed.
	//
	// example:
	//
	// ["secret/acs/ram/user/ram-secret", "secret/acs/ram/user/acr-master", "key/key-hzz63d9c8d3dfv8cv****"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) GetAccessControlRules() *string {
	return s.AccessControlRules
}

func (s *CreatePolicyResponseBody) GetArn() *string {
	return s.Arn
}

func (s *CreatePolicyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyResponseBody) GetKmsInstance() *string {
	return s.KmsInstance
}

func (s *CreatePolicyResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePolicyResponseBody) GetPermissions() *string {
	return s.Permissions
}

func (s *CreatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyResponseBody) GetResources() *string {
	return s.Resources
}

func (s *CreatePolicyResponseBody) SetAccessControlRules(v string) *CreatePolicyResponseBody {
	s.AccessControlRules = &v
	return s
}

func (s *CreatePolicyResponseBody) SetArn(v string) *CreatePolicyResponseBody {
	s.Arn = &v
	return s
}

func (s *CreatePolicyResponseBody) SetDescription(v string) *CreatePolicyResponseBody {
	s.Description = &v
	return s
}

func (s *CreatePolicyResponseBody) SetKmsInstance(v string) *CreatePolicyResponseBody {
	s.KmsInstance = &v
	return s
}

func (s *CreatePolicyResponseBody) SetName(v string) *CreatePolicyResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePolicyResponseBody) SetPermissions(v string) *CreatePolicyResponseBody {
	s.Permissions = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) SetResources(v string) *CreatePolicyResponseBody {
	s.Resources = &v
	return s
}

func (s *CreatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
