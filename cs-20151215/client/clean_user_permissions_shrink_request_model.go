// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanUserPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIdsShrink(v string) *CleanUserPermissionsShrinkRequest
	GetClusterIdsShrink() *string
	SetForce(v bool) *CleanUserPermissionsShrinkRequest
	GetForce() *bool
}

type CleanUserPermissionsShrinkRequest struct {
	// The list of cluster IDs. If you specify this list, only the KubeConfig credentials and RBAC permissions of the specified user in the listed clusters are cleared.
	ClusterIdsShrink *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// Specifies whether to force delete the specified KubeConfig. Valid values:
	//
	// - false (default): Checks whether there are cluster access records within the last seven days before deleting the KubeConfig. If such records exist or cannot be found, the deletion is not allowed.
	//
	// - true: Force deletes the KubeConfig without checking cluster access records.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s CleanUserPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CleanUserPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CleanUserPermissionsShrinkRequest) GetClusterIdsShrink() *string {
	return s.ClusterIdsShrink
}

func (s *CleanUserPermissionsShrinkRequest) GetForce() *bool {
	return s.Force
}

func (s *CleanUserPermissionsShrinkRequest) SetClusterIdsShrink(v string) *CleanUserPermissionsShrinkRequest {
	s.ClusterIdsShrink = &v
	return s
}

func (s *CleanUserPermissionsShrinkRequest) SetForce(v bool) *CleanUserPermissionsShrinkRequest {
	s.Force = &v
	return s
}

func (s *CleanUserPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
