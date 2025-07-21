// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlRules(v string) *DescribePolicyResponseBody
	GetAccessControlRules() *string
	SetArn(v string) *DescribePolicyResponseBody
	GetArn() *string
	SetDescription(v string) *DescribePolicyResponseBody
	GetDescription() *string
	SetKmsInstance(v string) *DescribePolicyResponseBody
	GetKmsInstance() *string
	SetName(v string) *DescribePolicyResponseBody
	GetName() *string
	SetPermissions(v []*string) *DescribePolicyResponseBody
	GetPermissions() []*string
	SetRequestId(v string) *DescribePolicyResponseBody
	GetRequestId() *string
	SetResources(v []*string) *DescribePolicyResponseBody
	GetResources() []*string
}

type DescribePolicyResponseBody struct {
	// The network access rule that is associated with the permission policy.
	//
	// example:
	//
	// {"NetworkRules":["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]}
	AccessControlRules *string `json:"AccessControlRules,omitempty" xml:"AccessControlRules,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the permission policy.
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
	// A list of operations that can be performed.
	//
	// example:
	//
	// ["RbacPermission/Template/CryptoServiceKeyUser", "RbacPermission/Template/CryptoServiceSecretUser"]
	Permissions []*string `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// f455324b-e229-4066-9f58-9c1cf3fe83a9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of keys and secrets that are allowed to access.
	//
	// example:
	//
	// ["secret/acs/ram/user/ram-secret", "secret/acs/ram/user/acr-master", "key/key-hzz63d9c8d3dfv8cv****"]
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyResponseBody) GetAccessControlRules() *string {
	return s.AccessControlRules
}

func (s *DescribePolicyResponseBody) GetArn() *string {
	return s.Arn
}

func (s *DescribePolicyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribePolicyResponseBody) GetKmsInstance() *string {
	return s.KmsInstance
}

func (s *DescribePolicyResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribePolicyResponseBody) GetPermissions() []*string {
	return s.Permissions
}

func (s *DescribePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolicyResponseBody) GetResources() []*string {
	return s.Resources
}

func (s *DescribePolicyResponseBody) SetAccessControlRules(v string) *DescribePolicyResponseBody {
	s.AccessControlRules = &v
	return s
}

func (s *DescribePolicyResponseBody) SetArn(v string) *DescribePolicyResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribePolicyResponseBody) SetDescription(v string) *DescribePolicyResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePolicyResponseBody) SetKmsInstance(v string) *DescribePolicyResponseBody {
	s.KmsInstance = &v
	return s
}

func (s *DescribePolicyResponseBody) SetName(v string) *DescribePolicyResponseBody {
	s.Name = &v
	return s
}

func (s *DescribePolicyResponseBody) SetPermissions(v []*string) *DescribePolicyResponseBody {
	s.Permissions = v
	return s
}

func (s *DescribePolicyResponseBody) SetRequestId(v string) *DescribePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyResponseBody) SetResources(v []*string) *DescribePolicyResponseBody {
	s.Resources = v
	return s
}

func (s *DescribePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
