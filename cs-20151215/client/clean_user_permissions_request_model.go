// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanUserPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v []*string) *CleanUserPermissionsRequest
	GetClusterIds() []*string
	SetForce(v bool) *CleanUserPermissionsRequest
	GetForce() *bool
}

type CleanUserPermissionsRequest struct {
	// The cluster IDs. If you specify a list of cluster IDs, only the kubeconfig files and RBAC permissions of the clusters that belong to the current user in the list are revoked.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
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

func (s CleanUserPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CleanUserPermissionsRequest) GoString() string {
	return s.String()
}

func (s *CleanUserPermissionsRequest) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *CleanUserPermissionsRequest) GetForce() *bool {
	return s.Force
}

func (s *CleanUserPermissionsRequest) SetClusterIds(v []*string) *CleanUserPermissionsRequest {
	s.ClusterIds = v
	return s
}

func (s *CleanUserPermissionsRequest) SetForce(v bool) *CleanUserPermissionsRequest {
	s.Force = &v
	return s
}

func (s *CleanUserPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
