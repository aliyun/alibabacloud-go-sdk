// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListK8sResourceRequest
	GetClusterId() *string
	SetK8sClusterId(v string) *ListK8sResourceRequest
	GetK8sClusterId() *string
	SetK8sNamespace(v string) *ListK8sResourceRequest
	GetK8sNamespace() *string
	SetResourceType(v string) *ListK8sResourceRequest
	GetResourceType() *string
	SetVpcId(v string) *ListK8sResourceRequest
	GetVpcId() *string
}

type ListK8sResourceRequest struct {
	// The Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the Kubernetes cluster used for synchronizing namespaces. The cluster is added to the service mesh specified by ServiceMeshId. This parameter is required when you query a k8s namespace or a k8s service.
	//
	// example:
	//
	// cadb451ed7af04b0297f4be396a2196fc
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The name of the namespace in the k8s cluster. This parameter is required when you query a k8s service.
	//
	// example:
	//
	// default
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The VPC ID. This parameter is required if ResourceType is set to \\"cluster\\".
	//
	// example:
	//
	// vpc-aa1a18236n90rqhuhhnhh
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListK8sResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListK8sResourceRequest) GoString() string {
	return s.String()
}

func (s *ListK8sResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sResourceRequest) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *ListK8sResourceRequest) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *ListK8sResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListK8sResourceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListK8sResourceRequest) SetClusterId(v string) *ListK8sResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *ListK8sResourceRequest) SetK8sClusterId(v string) *ListK8sResourceRequest {
	s.K8sClusterId = &v
	return s
}

func (s *ListK8sResourceRequest) SetK8sNamespace(v string) *ListK8sResourceRequest {
	s.K8sNamespace = &v
	return s
}

func (s *ListK8sResourceRequest) SetResourceType(v string) *ListK8sResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *ListK8sResourceRequest) SetVpcId(v string) *ListK8sResourceRequest {
	s.VpcId = &v
	return s
}

func (s *ListK8sResourceRequest) Validate() error {
	return dara.Validate(s)
}
