// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMNamespaceFromGuestClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetK8sClusterId(v string) *UpdateASMNamespaceFromGuestClusterRequest
	GetK8sClusterId() *string
	SetServiceMeshId(v string) *UpdateASMNamespaceFromGuestClusterRequest
	GetServiceMeshId() *string
}

type UpdateASMNamespaceFromGuestClusterRequest struct {
	// The ID of the Kubernetes cluster whose namespace information you want to synchronize to ASM. The Kubernetes cluster is added to the ASM instance that is specified by the ServiceMeshId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// c6f6d46583def494ba1f2e2937c8*****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbf9ca9e6d5dc4c87a3ecd0ebf1e*****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateASMNamespaceFromGuestClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMNamespaceFromGuestClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateASMNamespaceFromGuestClusterRequest) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *UpdateASMNamespaceFromGuestClusterRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateASMNamespaceFromGuestClusterRequest) SetK8sClusterId(v string) *UpdateASMNamespaceFromGuestClusterRequest {
	s.K8sClusterId = &v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterRequest) SetServiceMeshId(v string) *UpdateASMNamespaceFromGuestClusterRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateASMNamespaceFromGuestClusterRequest) Validate() error {
	return dara.Validate(s)
}
