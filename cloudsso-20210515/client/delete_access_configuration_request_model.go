// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *DeleteAccessConfigurationRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *DeleteAccessConfigurationRequest
	GetDirectoryId() *string
	SetForceRemovePermissionPolicies(v bool) *DeleteAccessConfigurationRequest
	GetForceRemovePermissionPolicies() *bool
}

type DeleteAccessConfigurationRequest struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-001j9mcm3k7335bc****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Specifies whether to forcibly remove system policies and inline policies. Valid values:
	//
	// - true: When you delete the access configuration, the associated system policies and inline policies are forcibly removed.
	//
	// - false: When you delete the access configuration, the associated system policies and inline policies are not forcibly removed. This is the default value. If these policies exist in the access configuration, the deletion fails. Before you delete the access configuration, you must remove the system policies and inline policies. For more information, see [RemovePermissionPolicyFromAccessConfiguration](https://help.aliyun.com/document_detail/336904.html).
	//
	// example:
	//
	// false
	ForceRemovePermissionPolicies *bool `json:"ForceRemovePermissionPolicies,omitempty" xml:"ForceRemovePermissionPolicies,omitempty"`
}

func (s DeleteAccessConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessConfigurationRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *DeleteAccessConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteAccessConfigurationRequest) GetForceRemovePermissionPolicies() *bool {
	return s.ForceRemovePermissionPolicies
}

func (s *DeleteAccessConfigurationRequest) SetAccessConfigurationId(v string) *DeleteAccessConfigurationRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeleteAccessConfigurationRequest) SetDirectoryId(v string) *DeleteAccessConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteAccessConfigurationRequest) SetForceRemovePermissionPolicies(v bool) *DeleteAccessConfigurationRequest {
	s.ForceRemovePermissionPolicies = &v
	return s
}

func (s *DeleteAccessConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
