// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClusterFromServiceMeshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RemoveClusterFromServiceMeshRequest
	GetClusterId() *string
	SetReserveNamespace(v bool) *RemoveClusterFromServiceMeshRequest
	GetReserveNamespace() *bool
	SetServiceMeshId(v string) *RemoveClusterFromServiceMeshRequest
	GetServiceMeshId() *string
}

type RemoveClusterFromServiceMeshRequest struct {
	// The ID of the cluster that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to reserve the istio-system namespace when you delete the cluster. Valid values: true and false.
	//
	// example:
	//
	// false
	ReserveNamespace *bool `json:"ReserveNamespace,omitempty" xml:"ReserveNamespace,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s RemoveClusterFromServiceMeshRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveClusterFromServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *RemoveClusterFromServiceMeshRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RemoveClusterFromServiceMeshRequest) GetReserveNamespace() *bool {
	return s.ReserveNamespace
}

func (s *RemoveClusterFromServiceMeshRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *RemoveClusterFromServiceMeshRequest) SetClusterId(v string) *RemoveClusterFromServiceMeshRequest {
	s.ClusterId = &v
	return s
}

func (s *RemoveClusterFromServiceMeshRequest) SetReserveNamespace(v bool) *RemoveClusterFromServiceMeshRequest {
	s.ReserveNamespace = &v
	return s
}

func (s *RemoveClusterFromServiceMeshRequest) SetServiceMeshId(v string) *RemoveClusterFromServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *RemoveClusterFromServiceMeshRequest) Validate() error {
	return dara.Validate(s)
}
