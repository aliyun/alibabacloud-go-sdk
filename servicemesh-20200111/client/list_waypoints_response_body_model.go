// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaypointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContinue(v string) *ListWaypointsResponseBody
	GetContinue() *string
	SetRequestId(v string) *ListWaypointsResponseBody
	GetRequestId() *string
	SetWaypoints(v []*ListWaypointsResponseBodyWaypoints) *ListWaypointsResponseBody
	GetWaypoints() []*ListWaypointsResponseBodyWaypoints
}

type ListWaypointsResponseBody struct {
	// Base64 encoded string. If it is empty, it means that all waypoints have been obtained; if it is not empty, this value should be included in the next list, and you can continue to obtain it from the offset that ends this time.
	//
	// example:
	//
	// eyJ2IjoibWV0YS5rOHMuaW8vdjEiLCJydiI6MjY4Njc5Miwic3RhcnQiOiJkZWZhdWx0L2Jvb2tpbmZvLXByb2R1Y3RwYWdlXHUwMDAwIn0
	Continue *string `json:"Continue,omitempty" xml:"Continue,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of waypoint proxies.
	Waypoints []*ListWaypointsResponseBodyWaypoints `json:"Waypoints,omitempty" xml:"Waypoints,omitempty" type:"Repeated"`
}

func (s ListWaypointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWaypointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaypointsResponseBody) GetContinue() *string {
	return s.Continue
}

func (s *ListWaypointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWaypointsResponseBody) GetWaypoints() []*ListWaypointsResponseBodyWaypoints {
	return s.Waypoints
}

func (s *ListWaypointsResponseBody) SetContinue(v string) *ListWaypointsResponseBody {
	s.Continue = &v
	return s
}

func (s *ListWaypointsResponseBody) SetRequestId(v string) *ListWaypointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaypointsResponseBody) SetWaypoints(v []*ListWaypointsResponseBodyWaypoints) *ListWaypointsResponseBody {
	s.Waypoints = v
	return s
}

func (s *ListWaypointsResponseBody) Validate() error {
	if s.Waypoints != nil {
		for _, item := range s.Waypoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWaypointsResponseBodyWaypoints struct {
	// Indicates whether Horizontal Pod Autoscaling (HPA) is enabled.
	//
	// example:
	//
	// false
	HPAEnabled *string `json:"HPAEnabled,omitempty" xml:"HPAEnabled,omitempty"`
	// The maximum number of waypoint proxy pods when HPA is enabled.
	//
	// example:
	//
	// 4
	HPAMaxReplicas *string `json:"HPAMaxReplicas,omitempty" xml:"HPAMaxReplicas,omitempty"`
	// The minimum number of waypoint proxy pods when HPA is enabled.
	//
	// example:
	//
	// 2
	HPAMinReplicas *string `json:"HPAMinReplicas,omitempty" xml:"HPAMinReplicas,omitempty"`
	// The expected CPU utilization when HPA is enabled.
	//
	// example:
	//
	// 93
	HPATargetCPU *string `json:"HPATargetCPU,omitempty" xml:"HPATargetCPU,omitempty"`
	// The expected memory usage when HPA is enabled.
	//
	// example:
	//
	// 91
	HPATargetMemory *string `json:"HPATargetMemory,omitempty" xml:"HPATargetMemory,omitempty"`
	// The maximum amount of CPU resources that are available to the waypoint proxy pods.
	//
	// example:
	//
	// 2000m
	LimitCPU *string `json:"LimitCPU,omitempty" xml:"LimitCPU,omitempty"`
	// The maximum amount of memory resources that are available to the waypoint proxy pods.
	//
	// example:
	//
	// 1024Mi
	LimitMemory *string `json:"LimitMemory,omitempty" xml:"LimitMemory,omitempty"`
	// The name of the gateway resource for the waypoint proxy. If the waypoint proxy applies to a service account, the name is the service account name. If the waypoint proxy applies to the entire namespace, the name is `namespace`.
	//
	// example:
	//
	// namespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Indicates whether waypoint proxy pods are preferentially scheduled to elastic container instances.
	//
	// example:
	//
	// false
	PreferECI *string `json:"PreferECI,omitempty" xml:"PreferECI,omitempty"`
	// The number of waypoint proxy pods.
	//
	// example:
	//
	// 1
	Replicas *string `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The amount of CPU resources requested by the waypoint proxy pods.
	//
	// example:
	//
	// 100m
	RequestCPU *string `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// The amount of memory resources requested by the waypoint proxy pods.
	//
	// example:
	//
	// 128Mi
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// The service account to which the waypoint proxy applies. If no value is returned for this parameter, the waypoint proxy applies to the entire namespace.
	//
	// example:
	//
	// bookinfo-productpage
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
}

func (s ListWaypointsResponseBodyWaypoints) String() string {
	return dara.Prettify(s)
}

func (s ListWaypointsResponseBodyWaypoints) GoString() string {
	return s.String()
}

func (s *ListWaypointsResponseBodyWaypoints) GetHPAEnabled() *string {
	return s.HPAEnabled
}

func (s *ListWaypointsResponseBodyWaypoints) GetHPAMaxReplicas() *string {
	return s.HPAMaxReplicas
}

func (s *ListWaypointsResponseBodyWaypoints) GetHPAMinReplicas() *string {
	return s.HPAMinReplicas
}

func (s *ListWaypointsResponseBodyWaypoints) GetHPATargetCPU() *string {
	return s.HPATargetCPU
}

func (s *ListWaypointsResponseBodyWaypoints) GetHPATargetMemory() *string {
	return s.HPATargetMemory
}

func (s *ListWaypointsResponseBodyWaypoints) GetLimitCPU() *string {
	return s.LimitCPU
}

func (s *ListWaypointsResponseBodyWaypoints) GetLimitMemory() *string {
	return s.LimitMemory
}

func (s *ListWaypointsResponseBodyWaypoints) GetName() *string {
	return s.Name
}

func (s *ListWaypointsResponseBodyWaypoints) GetNamespace() *string {
	return s.Namespace
}

func (s *ListWaypointsResponseBodyWaypoints) GetPreferECI() *string {
	return s.PreferECI
}

func (s *ListWaypointsResponseBodyWaypoints) GetReplicas() *string {
	return s.Replicas
}

func (s *ListWaypointsResponseBodyWaypoints) GetRequestCPU() *string {
	return s.RequestCPU
}

func (s *ListWaypointsResponseBodyWaypoints) GetRequestMemory() *string {
	return s.RequestMemory
}

func (s *ListWaypointsResponseBodyWaypoints) GetServiceAccount() *string {
	return s.ServiceAccount
}

func (s *ListWaypointsResponseBodyWaypoints) SetHPAEnabled(v string) *ListWaypointsResponseBodyWaypoints {
	s.HPAEnabled = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetHPAMaxReplicas(v string) *ListWaypointsResponseBodyWaypoints {
	s.HPAMaxReplicas = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetHPAMinReplicas(v string) *ListWaypointsResponseBodyWaypoints {
	s.HPAMinReplicas = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetHPATargetCPU(v string) *ListWaypointsResponseBodyWaypoints {
	s.HPATargetCPU = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetHPATargetMemory(v string) *ListWaypointsResponseBodyWaypoints {
	s.HPATargetMemory = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetLimitCPU(v string) *ListWaypointsResponseBodyWaypoints {
	s.LimitCPU = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetLimitMemory(v string) *ListWaypointsResponseBodyWaypoints {
	s.LimitMemory = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetName(v string) *ListWaypointsResponseBodyWaypoints {
	s.Name = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetNamespace(v string) *ListWaypointsResponseBodyWaypoints {
	s.Namespace = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetPreferECI(v string) *ListWaypointsResponseBodyWaypoints {
	s.PreferECI = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetReplicas(v string) *ListWaypointsResponseBodyWaypoints {
	s.Replicas = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetRequestCPU(v string) *ListWaypointsResponseBodyWaypoints {
	s.RequestCPU = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetRequestMemory(v string) *ListWaypointsResponseBodyWaypoints {
	s.RequestMemory = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) SetServiceAccount(v string) *ListWaypointsResponseBodyWaypoints {
	s.ServiceAccount = &v
	return s
}

func (s *ListWaypointsResponseBodyWaypoints) Validate() error {
	return dara.Validate(s)
}
