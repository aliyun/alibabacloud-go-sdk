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
	// The cluster IDs. If you specify a list of cluster IDs, only the kubeconfig files and RBAC permissions of the clusters that belong to the current user in the list are revoked.
	ClusterIdsShrink *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// Specifies whether to forcefully delete the specified kubeconfig files. Valid values:
	//
	// 	- false (default): checks the cluster access records within the previous seven days before deleting the kubeconfig files. The kubeconfig files are not deleted if cluster access records are found or fail to be retrieved.
	//
	// 	- true: forcefully deletes the kubeconfig files without checking the cluster access records.
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
