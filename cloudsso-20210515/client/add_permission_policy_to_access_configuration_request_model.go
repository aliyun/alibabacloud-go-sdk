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
	// The access configuration ID.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The content of the inline policy.
	//
	// The maximum length is 4,096 characters.
	//
	// If PermissionPolicyType is set to Inline, this parameter is required. For more information about the syntax and structure of access policies, see [Access policy syntax and structure](https://help.aliyun.com/document_detail/93739.html).
	//
	// example:
	//
	// {"Statement": [{"Action": "*","Effect": "Allow","Resource": "*"}],"Version": "1"}
	InlinePolicyDocument *string `json:"InlinePolicyDocument,omitempty" xml:"InlinePolicyDocument,omitempty"`
	// The name of the access policy.
	//
	// - If PermissionPolicyType is set to System, this parameter specifies the name of the system policy. Because RAM system policies are reused, you must obtain the system policy name from RAM.
	//
	// - If PermissionPolicyType is set to Inline, this parameter specifies the name of the inline policy. You can customize the name. The name can be up to 32 characters in length.
	//
	// example:
	//
	// AliyunECSFullAccess
	PermissionPolicyName *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	// The type of the access policy. Valid values:
	//
	// - System: system policy. Reuses the system policies of Resource Access Management (RAM).
	//
	// - Inline: inline policy. A custom policy written based on the syntax and structure of RAM policies.
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
