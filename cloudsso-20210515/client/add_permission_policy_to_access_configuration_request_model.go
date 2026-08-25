// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionPolicyToAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *AddPermissionPolicyToAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *AddPermissionPolicyToAccessConfigurationRequest
	GetDirectoryId() *string
	SetInlinePolicyDocument(v string) *AddPermissionPolicyToAccessConfigurationRequest
	GetInlinePolicyDocument() *string
	SetPermissionPolicyName(v string) *AddPermissionPolicyToAccessConfigurationRequest
	GetPermissionPolicyName() *string
	SetPermissionPolicyType(v string) *AddPermissionPolicyToAccessConfigurationRequest
	GetPermissionPolicyType() *string
}

type AddPermissionPolicyToAccessConfigurationRequest struct {
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
	// The configurations of the inline policy.
	//
	// The value can be up to 4,096 characters in length.
	//
	// If you set `PermissionPolicyType` to `Inline`, you must specify this parameter. For more information about the syntax and structure of RAM policies, see [Policy syntax and structure](https://help.aliyun.com/document_detail/93739.html).
	//
	// example:
	//
	// {"Statement": [{"Action": "*","Effect": "Allow","Resource": "*"}],"Version": "1"}
	InlinePolicyDocument *string `json:"InlinePolicyDocument,omitempty" xml:"InlinePolicyDocument,omitempty"`
	// The name of the policy.
	//
	// - If you set `PermissionPolicyType` to `System`, you must set PermissionPolicyName to the name of a system policy. You can obtain the name of the system policy from RAM.
	//
	// - If you set `PermissionPolicyType` to `Inline`, you must set PermissionPolicyName to the name of an inline policy. A custom value is supported. The value can be up to 32 characters in length.
	//
	// example:
	//
	// AliyunECSFullAccess
	PermissionPolicyName *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// - System: system policy. Resource Access Management (RAM) system policies are reused.
	//
	// - Inline: inline policy. Inline policies are created based on the RAM policy syntax and structure.
	//
	// example:
	//
	// System
	PermissionPolicyType *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
}

func (s AddPermissionPolicyToAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionPolicyToAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) GetInlinePolicyDocument() *string {
	return s.InlinePolicyDocument
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) GetPermissionPolicyName() *string {
	return s.PermissionPolicyName
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) GetPermissionPolicyType() *string {
	return s.PermissionPolicyType
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetAccessConfigurationId(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetDirectoryId(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetInlinePolicyDocument(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.InlinePolicyDocument = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetPermissionPolicyName(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.PermissionPolicyName = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) SetPermissionPolicyType(v string) *AddPermissionPolicyToAccessConfigurationRequest {
	s.PermissionPolicyType = &v
	return s
}

func (s *AddPermissionPolicyToAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
