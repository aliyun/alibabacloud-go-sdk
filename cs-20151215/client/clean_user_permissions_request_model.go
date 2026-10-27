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
	// The list of cluster IDs. If you specify this list, only the KubeConfig credentials and RBAC permissions of the specified user in the listed clusters are cleared.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
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
