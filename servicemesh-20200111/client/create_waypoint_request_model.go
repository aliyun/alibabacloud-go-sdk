// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaypointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateWaypointRequest
	GetClusterId() *string
	SetHPAEnabled(v bool) *CreateWaypointRequest
	GetHPAEnabled() *bool
	SetHPAMaxReplicas(v int32) *CreateWaypointRequest
	GetHPAMaxReplicas() *int32
	SetHPAMinReplicas(v int32) *CreateWaypointRequest
	GetHPAMinReplicas() *int32
	SetHPATargetCPU(v int32) *CreateWaypointRequest
	GetHPATargetCPU() *int32
	SetHPATargetMemory(v int32) *CreateWaypointRequest
	GetHPATargetMemory() *int32
	SetLimitCPU(v string) *CreateWaypointRequest
	GetLimitCPU() *string
	SetLimitMemory(v string) *CreateWaypointRequest
	GetLimitMemory() *string
	SetNamespace(v string) *CreateWaypointRequest
	GetNamespace() *string
	SetPreferECI(v bool) *CreateWaypointRequest
	GetPreferECI() *bool
	SetReplicas(v int32) *CreateWaypointRequest
	GetReplicas() *int32
	SetRequestCPU(v string) *CreateWaypointRequest
	GetRequestCPU() *string
	SetRequestMemory(v string) *CreateWaypointRequest
	GetRequestMemory() *string
	SetServiceAccount(v string) *CreateWaypointRequest
	GetServiceAccount() *string
	SetServiceMeshId(v string) *CreateWaypointRequest
	GetServiceMeshId() *string
}

type CreateWaypointRequest struct {
	// The ID of the cluster on the data plane.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
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
	// 93
	HPATargetCPU *int32 `json:"HPATargetCPU,omitempty" xml:"HPATargetCPU,omitempty"`
	// The expected memory usage when HPA is enabled.
	//
	// example:
	//
	// 91
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
	// The service account on which the waypoint proxy takes effect. If you do not specify this parameter, the waypoint proxy takes effect for the entire namespace.
	//
	// example:
	//
	// bookinfo-productpage
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	// The Service Mesh (ASM) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce134b0727aa2492db69f6c3880e****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s CreateWaypointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWaypointRequest) GoString() string {
	return s.String()
}

func (s *CreateWaypointRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateWaypointRequest) GetHPAEnabled() *bool {
	return s.HPAEnabled
}

func (s *CreateWaypointRequest) GetHPAMaxReplicas() *int32 {
	return s.HPAMaxReplicas
}

func (s *CreateWaypointRequest) GetHPAMinReplicas() *int32 {
	return s.HPAMinReplicas
}

func (s *CreateWaypointRequest) GetHPATargetCPU() *int32 {
	return s.HPATargetCPU
}

func (s *CreateWaypointRequest) GetHPATargetMemory() *int32 {
	return s.HPATargetMemory
}

func (s *CreateWaypointRequest) GetLimitCPU() *string {
	return s.LimitCPU
}

func (s *CreateWaypointRequest) GetLimitMemory() *string {
	return s.LimitMemory
}

func (s *CreateWaypointRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateWaypointRequest) GetPreferECI() *bool {
	return s.PreferECI
}

func (s *CreateWaypointRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateWaypointRequest) GetRequestCPU() *string {
	return s.RequestCPU
}

func (s *CreateWaypointRequest) GetRequestMemory() *string {
	return s.RequestMemory
}

func (s *CreateWaypointRequest) GetServiceAccount() *string {
	return s.ServiceAccount
}

func (s *CreateWaypointRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateWaypointRequest) SetClusterId(v string) *CreateWaypointRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateWaypointRequest) SetHPAEnabled(v bool) *CreateWaypointRequest {
	s.HPAEnabled = &v
	return s
}

func (s *CreateWaypointRequest) SetHPAMaxReplicas(v int32) *CreateWaypointRequest {
	s.HPAMaxReplicas = &v
	return s
}

func (s *CreateWaypointRequest) SetHPAMinReplicas(v int32) *CreateWaypointRequest {
	s.HPAMinReplicas = &v
	return s
}

func (s *CreateWaypointRequest) SetHPATargetCPU(v int32) *CreateWaypointRequest {
	s.HPATargetCPU = &v
	return s
}

func (s *CreateWaypointRequest) SetHPATargetMemory(v int32) *CreateWaypointRequest {
	s.HPATargetMemory = &v
	return s
}

func (s *CreateWaypointRequest) SetLimitCPU(v string) *CreateWaypointRequest {
	s.LimitCPU = &v
	return s
}

func (s *CreateWaypointRequest) SetLimitMemory(v string) *CreateWaypointRequest {
	s.LimitMemory = &v
	return s
}

func (s *CreateWaypointRequest) SetNamespace(v string) *CreateWaypointRequest {
	s.Namespace = &v
	return s
}

func (s *CreateWaypointRequest) SetPreferECI(v bool) *CreateWaypointRequest {
	s.PreferECI = &v
	return s
}

func (s *CreateWaypointRequest) SetReplicas(v int32) *CreateWaypointRequest {
	s.Replicas = &v
	return s
}

func (s *CreateWaypointRequest) SetRequestCPU(v string) *CreateWaypointRequest {
	s.RequestCPU = &v
	return s
}

func (s *CreateWaypointRequest) SetRequestMemory(v string) *CreateWaypointRequest {
	s.RequestMemory = &v
	return s
}

func (s *CreateWaypointRequest) SetServiceAccount(v string) *CreateWaypointRequest {
	s.ServiceAccount = &v
	return s
}

func (s *CreateWaypointRequest) SetServiceMeshId(v string) *CreateWaypointRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateWaypointRequest) Validate() error {
	return dara.Validate(s)
}
