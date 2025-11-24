// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshCRAggregationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCPULimit(v string) *UpdateMeshCRAggregationRequest
	GetCPULimit() *string
	SetCPURequirement(v string) *UpdateMeshCRAggregationRequest
	GetCPURequirement() *string
	SetEnabled(v bool) *UpdateMeshCRAggregationRequest
	GetEnabled() *bool
	SetMemoryLimit(v string) *UpdateMeshCRAggregationRequest
	GetMemoryLimit() *string
	SetMemoryRequirement(v string) *UpdateMeshCRAggregationRequest
	GetMemoryRequirement() *string
	SetServiceMeshId(v string) *UpdateMeshCRAggregationRequest
	GetServiceMeshId() *string
	SetUsePublicApiServer(v bool) *UpdateMeshCRAggregationRequest
	GetUsePublicApiServer() *bool
}

type UpdateMeshCRAggregationRequest struct {
	// The maximum number of CPU cores that are available for the components installed in the ACK cluster on the data plane if you enable the Kubernetes API to access Istio resources in the ASM instance. You can specify the parameter value in the standard quantity representation used by Kubernetes.
	//
	// example:
	//
	// 1
	CPULimit *string `json:"CPULimit,omitempty" xml:"CPULimit,omitempty"`
	// The number of CPU cores that are requested by the components installed in the Container Service for Kubernetes (ACK) cluster on the data plane if you enable the Kubernetes API to access Istio resources in the ASM instance. You can specify the parameter value in the standard quantity representation used by Kubernetes.
	//
	// example:
	//
	// 1
	CPURequirement *string `json:"CPURequirement,omitempty" xml:"CPURequirement,omitempty"`
	// Specifies whether to enable the Kubernetes API on the data plane to access Istio resources in the ASM instance. Valid values:
	//
	// 	- `true`: enables the Kubernetes API to access Istio resources in the ASM instance.
	//
	// 	- `false`: disables the Kubernetes API to access Istio resources in the ASM instance.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum size of the memory that is available for the components installed in the ACK cluster on the data plane if you enable the Kubernetes API to access Istio resources in the ASM instance. You can specify the parameter value in the standard quantity representation used by Kubernetes. 1 Mi equals 1,024 KB.
	//
	// example:
	//
	// 500Mi
	MemoryLimit *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The size of the memory that is requested by the components installed in the ACK cluster on the data plane if you enable the Kubernetes API to access Istio resources in the ASM instance. You can specify the parameter value in the standard quantity representation used by Kubernetes. 1 Mi equals 1,024 KB.
	//
	// example:
	//
	// 500Mi
	MemoryRequirement *string `json:"MemoryRequirement,omitempty" xml:"MemoryRequirement,omitempty"`
	// The Service Mesh (ASM) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// Specifies whether the Kubernetes API on the data plane uses the public endpoint of the API server to access Istio resources in the ASM instance. Valid values:
	//
	// 	- `true`: The Kubernetes API on the data plane uses the public endpoint of the API server to access Istio resources in the ASM instance.
	//
	// 	- `false`: The Kubernetes API on the data plane uses the private endpoint of the API server to access Istio resources in the ASM instance.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	UsePublicApiServer *bool `json:"UsePublicApiServer,omitempty" xml:"UsePublicApiServer,omitempty"`
}

func (s UpdateMeshCRAggregationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshCRAggregationRequest) GoString() string {
	return s.String()
}

func (s *UpdateMeshCRAggregationRequest) GetCPULimit() *string {
	return s.CPULimit
}

func (s *UpdateMeshCRAggregationRequest) GetCPURequirement() *string {
	return s.CPURequirement
}

func (s *UpdateMeshCRAggregationRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMeshCRAggregationRequest) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *UpdateMeshCRAggregationRequest) GetMemoryRequirement() *string {
	return s.MemoryRequirement
}

func (s *UpdateMeshCRAggregationRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateMeshCRAggregationRequest) GetUsePublicApiServer() *bool {
	return s.UsePublicApiServer
}

func (s *UpdateMeshCRAggregationRequest) SetCPULimit(v string) *UpdateMeshCRAggregationRequest {
	s.CPULimit = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetCPURequirement(v string) *UpdateMeshCRAggregationRequest {
	s.CPURequirement = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetEnabled(v bool) *UpdateMeshCRAggregationRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetMemoryLimit(v string) *UpdateMeshCRAggregationRequest {
	s.MemoryLimit = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetMemoryRequirement(v string) *UpdateMeshCRAggregationRequest {
	s.MemoryRequirement = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetServiceMeshId(v string) *UpdateMeshCRAggregationRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) SetUsePublicApiServer(v bool) *UpdateMeshCRAggregationRequest {
	s.UsePublicApiServer = &v
	return s
}

func (s *UpdateMeshCRAggregationRequest) Validate() error {
	return dara.Validate(s)
}
