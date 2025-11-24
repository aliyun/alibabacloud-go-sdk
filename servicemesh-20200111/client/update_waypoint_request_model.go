// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaypointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateWaypointRequest
	GetClusterId() *string
	SetHPAEnabled(v bool) *UpdateWaypointRequest
	GetHPAEnabled() *bool
	SetHPAMaxReplicas(v int32) *UpdateWaypointRequest
	GetHPAMaxReplicas() *int32
	SetHPAMinReplicas(v int32) *UpdateWaypointRequest
	GetHPAMinReplicas() *int32
	SetHPATargetCPU(v int32) *UpdateWaypointRequest
	GetHPATargetCPU() *int32
	SetHPATargetMemory(v int32) *UpdateWaypointRequest
	GetHPATargetMemory() *int32
	SetLimitCPU(v string) *UpdateWaypointRequest
	GetLimitCPU() *string
	SetLimitMemory(v string) *UpdateWaypointRequest
	GetLimitMemory() *string
	SetName(v string) *UpdateWaypointRequest
	GetName() *string
	SetNamespace(v string) *UpdateWaypointRequest
	GetNamespace() *string
	SetPreferECI(v bool) *UpdateWaypointRequest
	GetPreferECI() *bool
	SetReplicas(v int32) *UpdateWaypointRequest
	GetReplicas() *int32
	SetRequestCPU(v string) *UpdateWaypointRequest
	GetRequestCPU() *string
	SetRequestMemory(v string) *UpdateWaypointRequest
	GetRequestMemory() *string
	SetServiceMeshId(v string) *UpdateWaypointRequest
	GetServiceMeshId() *string
}

type UpdateWaypointRequest struct {
	// The ID of the cluster on the data plane.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable Horizontal Pod Autoscaling (HPA).
	//
	// example:
	//
	// false
	HPAEnabled *bool `json:"HPAEnabled,omitempty" xml:"HPAEnabled,omitempty"`
	// The maximum number of waypoint proxy pods when HPA is enabled.
	//
	// example:
	//
	// 4
	HPAMaxReplicas *int32 `json:"HPAMaxReplicas,omitempty" xml:"HPAMaxReplicas,omitempty"`
	// The minimum number of waypoint proxy pods when HPA is enabled.
	//
	// example:
	//
	// 2
	HPAMinReplicas *int32 `json:"HPAMinReplicas,omitempty" xml:"HPAMinReplicas,omitempty"`
	// The expected CPU utilization when HPA is enabled.
	//
	// example:
	//
	// 91
	HPATargetCPU *int32 `json:"HPATargetCPU,omitempty" xml:"HPATargetCPU,omitempty"`
	// The expected memory usage when HPA is enabled.
	//
	// example:
	//
	// 89
	HPATargetMemory *int32 `json:"HPATargetMemory,omitempty" xml:"HPATargetMemory,omitempty"`
	// The maximum number of CPU cores that are available to the waypoint proxy pods.
	//
	// example:
	//
	// 2000m
	LimitCPU *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	// The maximum size of the memory that is available to the waypoint proxy pods.
	//
	// example:
	//
	// 1024Mi
	LimitMemory *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	// Waypoint名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// bookinfo-reviews
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Specifies whether to deploy waypoint proxy pods based on Elastic Container Instance (ECI).
	//
	// example:
	//
	// false
	PreferECI *bool `json:"PreferECI,omitempty" xml:"PreferECI,omitempty"`
	// The number of waypoint proxy pods.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The number of CPU cores requested by the waypoint proxy pods.
	//
	// example:
	//
	// 100m
	RequestCPU *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The size of the memory requested by the waypoint proxy pods.
	//
	// example:
	//
	// 128Mi
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// The Service Mesh (ASM) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpdateWaypointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaypointRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaypointRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateWaypointRequest) GetHPAEnabled() *bool {
	return s.HPAEnabled
}

func (s *UpdateWaypointRequest) GetHPAMaxReplicas() *int32 {
	return s.HPAMaxReplicas
}

func (s *UpdateWaypointRequest) GetHPAMinReplicas() *int32 {
	return s.HPAMinReplicas
}

func (s *UpdateWaypointRequest) GetHPATargetCPU() *int32 {
	return s.HPATargetCPU
}

func (s *UpdateWaypointRequest) GetHPATargetMemory() *int32 {
	return s.HPATargetMemory
}

func (s *UpdateWaypointRequest) GetLimitCPU() *string {
	return s.LimitCPU
}

func (s *UpdateWaypointRequest) GetLimitMemory() *string {
	return s.LimitMemory
}

func (s *UpdateWaypointRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWaypointRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateWaypointRequest) GetPreferECI() *bool {
	return s.PreferECI
}

func (s *UpdateWaypointRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *UpdateWaypointRequest) GetRequestCPU() *string {
	return s.RequestCPU
}

func (s *UpdateWaypointRequest) GetRequestMemory() *string {
	return s.RequestMemory
}

func (s *UpdateWaypointRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpdateWaypointRequest) SetClusterId(v string) *UpdateWaypointRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateWaypointRequest) SetHPAEnabled(v bool) *UpdateWaypointRequest {
	s.HPAEnabled = &v
	return s
}

func (s *UpdateWaypointRequest) SetHPAMaxReplicas(v int32) *UpdateWaypointRequest {
	s.HPAMaxReplicas = &v
	return s
}

func (s *UpdateWaypointRequest) SetHPAMinReplicas(v int32) *UpdateWaypointRequest {
	s.HPAMinReplicas = &v
	return s
}

func (s *UpdateWaypointRequest) SetHPATargetCPU(v int32) *UpdateWaypointRequest {
	s.HPATargetCPU = &v
	return s
}

func (s *UpdateWaypointRequest) SetHPATargetMemory(v int32) *UpdateWaypointRequest {
	s.HPATargetMemory = &v
	return s
}

func (s *UpdateWaypointRequest) SetLimitCPU(v string) *UpdateWaypointRequest {
	s.LimitCPU = &v
	return s
}

func (s *UpdateWaypointRequest) SetLimitMemory(v string) *UpdateWaypointRequest {
	s.LimitMemory = &v
	return s
}

func (s *UpdateWaypointRequest) SetName(v string) *UpdateWaypointRequest {
	s.Name = &v
	return s
}

func (s *UpdateWaypointRequest) SetNamespace(v string) *UpdateWaypointRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateWaypointRequest) SetPreferECI(v bool) *UpdateWaypointRequest {
	s.PreferECI = &v
	return s
}

func (s *UpdateWaypointRequest) SetReplicas(v int32) *UpdateWaypointRequest {
	s.Replicas = &v
	return s
}

func (s *UpdateWaypointRequest) SetRequestCPU(v string) *UpdateWaypointRequest {
	s.RequestCPU = &v
	return s
}

func (s *UpdateWaypointRequest) SetRequestMemory(v string) *UpdateWaypointRequest {
	s.RequestMemory = &v
	return s
}

func (s *UpdateWaypointRequest) SetServiceMeshId(v string) *UpdateWaypointRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpdateWaypointRequest) Validate() error {
	return dara.Validate(s)
}
