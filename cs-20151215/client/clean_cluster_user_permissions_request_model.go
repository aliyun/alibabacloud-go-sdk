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
