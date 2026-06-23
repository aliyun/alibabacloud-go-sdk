// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanClusterUserPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *CleanClusterUserPermissionsRequest
	GetForce() *bool
}

type CleanClusterUserPermissionsRequest struct {
	// Specifies whether to force delete the specified KubeConfig. Valid values:
	//
	// - false (default): Before deleting the KubeConfig, the system checks whether cluster access records exist within the last seven days. If access records exist or cannot be retrieved, the deletion is not allowed.
	//
	// - true: Force deletes the KubeConfig without checking cluster access records.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s CleanClusterUserPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CleanClusterUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *CleanClusterUserPermissionsRequest) GetForce() *bool {
	return s.Force
}

func (s *CleanClusterUserPermissionsRequest) SetForce(v bool) *CleanClusterUserPermissionsRequest {
	s.Force = &v
	return s
}

func (s *CleanClusterUserPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
