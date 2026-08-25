// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePermissionPolicyFromAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *RemovePermissionPolicyFromAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *RemovePermissionPolicyFromAccessConfigurationRequest
	GetDirectoryId() *string
	SetPermissionPolicyName(v string) *RemovePermissionPolicyFromAccessConfigurationRequest
	GetPermissionPolicyName() *string
	SetPermissionPolicyType(v string) *RemovePermissionPolicyFromAccessConfigurationRequest
	GetPermissionPolicyType() *string
}

type RemovePermissionPolicyFromAccessConfigurationRequest struct {
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
	// The name of the policy.
	//
	// example:
	//
	// AliyunECSFullAccess
	PermissionPolicyName *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// - System: system policy.
	//
	// - Inline: inline policy.
	//
	// example:
	//
	// System
	PermissionPolicyType *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
}

func (s RemovePermissionPolicyFromAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePermissionPolicyFromAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) GetPermissionPolicyName() *string {
	return s.PermissionPolicyName
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) GetPermissionPolicyType() *string {
	return s.PermissionPolicyType
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) SetAccessConfigurationId(v string) *RemovePermissionPolicyFromAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) SetDirectoryId(v string) *RemovePermissionPolicyFromAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) SetPermissionPolicyName(v string) *RemovePermissionPolicyFromAccessConfigurationRequest {
	s.PermissionPolicyName = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) SetPermissionPolicyType(v string) *RemovePermissionPolicyFromAccessConfigurationRequest {
	s.PermissionPolicyType = &v
	return s
}

func (s *RemovePermissionPolicyFromAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
