// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstance interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentAmount(v float32) *Instance
	GetCurrentAmount() *float32
	SetDetached(v bool) *Instance
	GetDetached() *bool
	SetExternalIP(v string) *Instance
	GetExternalIP() *string
	SetExternalInstancePort(v int32) *Instance
	GetExternalInstancePort() *int32
	SetHostIP(v string) *Instance
	GetHostIP() *string
	SetHostName(v string) *Instance
	GetHostName() *string
	SetInnerIP(v string) *Instance
	GetInnerIP() *string
	SetInstanceName(v string) *Instance
	GetInstanceName() *string
	SetInstancePort(v int32) *Instance
	GetInstancePort() *int32
	SetInstanceType(v string) *Instance
	GetInstanceType() *string
	SetIsLatest(v bool) *Instance
	GetIsLatest() *bool
	SetIsReplica(v bool) *Instance
	GetIsReplica() *bool
	SetIsSpot(v bool) *Instance
	GetIsSpot() *bool
	SetIsolated(v bool) *Instance
	GetIsolated() *bool
	SetLastState(v []map[string]interface{}) *Instance
	GetLastState() []map[string]interface{}
	SetNamespace(v string) *Instance
	GetNamespace() *string
	SetOriginalAmount(v float32) *Instance
	GetOriginalAmount() *float32
	SetReadyProcesses(v int32) *Instance
	GetReadyProcesses() *int32
	SetReason(v string) *Instance
	GetReason() *string
	SetReplicaName(v string) *Instance
	GetReplicaName() *string
	SetResourceType(v string) *Instance
	GetResourceType() *string
	SetRestartCount(v int32) *Instance
	GetRestartCount() *int32
	SetRole(v string) *Instance
	GetRole() *string
	SetStartAt(v string) *Instance
	GetStartAt() *string
	SetStartTime(v string) *Instance
	GetStartTime() *string
	SetStatus(v string) *Instance
	GetStatus() *string
	SetTenantHostIP(v string) *Instance
	GetTenantHostIP() *string
	SetTenantInstanceIP(v string) *Instance
	GetTenantInstanceIP() *string
	SetTotalProcesses(v int32) *Instance
	GetTotalProcesses() *int32
	SetZone(v string) *Instance
	GetZone() *string
}

type Instance struct {
	// The current hourly price of the spot instance.
	//
	// example:
	//
	// 0.444
	CurrentAmount *float32 `json:"CurrentAmount,omitempty" xml:"CurrentAmount,omitempty"`
	Detached      *bool    `json:"Detached,omitempty" xml:"Detached,omitempty"`
	// The IP address of the instance in the user-created VPC.
	//
	// example:
	//
	// 192.168.1.100
	ExternalIP *string `json:"ExternalIP,omitempty" xml:"ExternalIP,omitempty"`
	// The port number of the instance in the user-created VPC.
	//
	// example:
	//
	// 8080
	ExternalInstancePort *int32 `json:"ExternalInstancePort,omitempty" xml:"ExternalInstancePort,omitempty"`
	// The IP address of the host where the instance resides.
	//
	// example:
	//
	// 11.0.XX.XX
	HostIP *string `json:"HostIP,omitempty" xml:"HostIP,omitempty"`
	// The name of the host where the instance resides.
	//
	// example:
	//
	// smart-scene-cls-854dbdc99d-****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The internal IP address of the instance.
	//
	// example:
	//
	// 172.17.0.17
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// The instance name.
	//
	// example:
	//
	// foo-5fc8946767-v****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network port of the instance.
	//
	// example:
	//
	// 8080
	InstancePort *int32 `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	// The instance specification.
	//
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IsLatest     *bool   `json:"IsLatest,omitempty" xml:"IsLatest,omitempty"`
	IsReplica    *bool   `json:"IsReplica,omitempty" xml:"IsReplica,omitempty"`
	// Indicates whether the instance is a spot instance.
	//
	// example:
	//
	// false
	IsSpot *bool `json:"IsSpot,omitempty" xml:"IsSpot,omitempty"`
	// Indicates whether the instance is isolated.
	//
	// example:
	//
	// false
	Isolated *bool `json:"Isolated,omitempty" xml:"Isolated,omitempty"`
	// The last state of the instance.
	LastState []map[string]interface{} `json:"LastState,omitempty" xml:"LastState,omitempty" type:"Repeated"`
	// The namespace of the instance.
	//
	// example:
	//
	// foo
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The original hourly price of the spot instance before a discount is used.
	//
	// example:
	//
	// 2.2
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The number of processes that have started for the instance.
	//
	// example:
	//
	// 1
	ReadyProcesses *int32 `json:"ReadyProcesses,omitempty" xml:"ReadyProcesses,omitempty"`
	// The reason for which the instance is in the current state.
	//
	// example:
	//
	// RUNNING
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ReplicaName *string `json:"ReplicaName,omitempty" xml:"ReplicaName,omitempty"`
	// The type of the resource group to which the instance belongs. Valid values: PublicResource and PrivateResource.
	//
	// example:
	//
	// PublicResource
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of times for which the instance is restarted.
	//
	// example:
	//
	// 1
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// The service role of the instance. Valid values: Queue, DataLoader, and Standard.
	//
	// example:
	//
	// Standard
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Deprecated
	//
	// The time when the instance was started. This parameter is deprecated. StartTime is used instead.
	//
	// example:
	//
	// 2021-05-27T09:46:05Z
	StartAt *string `json:"StartAt,omitempty" xml:"StartAt,omitempty"`
	// The time when the instance was started.
	//
	// example:
	//
	// 2021-05-27T09:46:05Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The current state of the instance.
	//
	// Valid values:
	//
	// 	- Terminating
	//
	// 	- Succeeded
	//
	// 	- Unknown
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Pending
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IP address of the host in the VPC.
	//
	// example:
	//
	// 192.168.xx.xx
	TenantHostIP *string `json:"TenantHostIP,omitempty" xml:"TenantHostIP,omitempty"`
	// The IP address of the instance in the VPC.
	//
	// example:
	//
	// 192.168.xx.xx
	TenantInstanceIP *string `json:"TenantInstanceIP,omitempty" xml:"TenantInstanceIP,omitempty"`
	// The total number of processes that the instance contains.
	//
	// example:
	//
	// 1
	TotalProcesses *int32 `json:"TotalProcesses,omitempty" xml:"TotalProcesses,omitempty"`
	// The zone to which the instance belongs.
	//
	// example:
	//
	// cn-shanghai-a
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s Instance) String() string {
	return dara.Prettify(s)
}

func (s Instance) GoString() string {
	return s.String()
}

func (s *Instance) GetCurrentAmount() *float32 {
	return s.CurrentAmount
}

func (s *Instance) GetDetached() *bool {
	return s.Detached
}

func (s *Instance) GetExternalIP() *string {
	return s.ExternalIP
}

func (s *Instance) GetExternalInstancePort() *int32 {
	return s.ExternalInstancePort
}

func (s *Instance) GetHostIP() *string {
	return s.HostIP
}

func (s *Instance) GetHostName() *string {
	return s.HostName
}

func (s *Instance) GetInnerIP() *string {
	return s.InnerIP
}

func (s *Instance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *Instance) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *Instance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *Instance) GetIsLatest() *bool {
	return s.IsLatest
}

func (s *Instance) GetIsReplica() *bool {
	return s.IsReplica
}

func (s *Instance) GetIsSpot() *bool {
	return s.IsSpot
}

func (s *Instance) GetIsolated() *bool {
	return s.Isolated
}

func (s *Instance) GetLastState() []map[string]interface{} {
	return s.LastState
}

func (s *Instance) GetNamespace() *string {
	return s.Namespace
}

func (s *Instance) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *Instance) GetReadyProcesses() *int32 {
	return s.ReadyProcesses
}

func (s *Instance) GetReason() *string {
	return s.Reason
}

func (s *Instance) GetReplicaName() *string {
	return s.ReplicaName
}

func (s *Instance) GetResourceType() *string {
	return s.ResourceType
}

func (s *Instance) GetRestartCount() *int32 {
	return s.RestartCount
}

func (s *Instance) GetRole() *string {
	return s.Role
}

func (s *Instance) GetStartAt() *string {
	return s.StartAt
}

func (s *Instance) GetStartTime() *string {
	return s.StartTime
}

func (s *Instance) GetStatus() *string {
	return s.Status
}

func (s *Instance) GetTenantHostIP() *string {
	return s.TenantHostIP
}

func (s *Instance) GetTenantInstanceIP() *string {
	return s.TenantInstanceIP
}

func (s *Instance) GetTotalProcesses() *int32 {
	return s.TotalProcesses
}

func (s *Instance) GetZone() *string {
	return s.Zone
}

func (s *Instance) SetCurrentAmount(v float32) *Instance {
	s.CurrentAmount = &v
	return s
}

func (s *Instance) SetDetached(v bool) *Instance {
	s.Detached = &v
	return s
}

func (s *Instance) SetExternalIP(v string) *Instance {
	s.ExternalIP = &v
	return s
}

func (s *Instance) SetExternalInstancePort(v int32) *Instance {
	s.ExternalInstancePort = &v
	return s
}

func (s *Instance) SetHostIP(v string) *Instance {
	s.HostIP = &v
	return s
}

func (s *Instance) SetHostName(v string) *Instance {
	s.HostName = &v
	return s
}

func (s *Instance) SetInnerIP(v string) *Instance {
	s.InnerIP = &v
	return s
}

func (s *Instance) SetInstanceName(v string) *Instance {
	s.InstanceName = &v
	return s
}

func (s *Instance) SetInstancePort(v int32) *Instance {
	s.InstancePort = &v
	return s
}

func (s *Instance) SetInstanceType(v string) *Instance {
	s.InstanceType = &v
	return s
}

func (s *Instance) SetIsLatest(v bool) *Instance {
	s.IsLatest = &v
	return s
}

func (s *Instance) SetIsReplica(v bool) *Instance {
	s.IsReplica = &v
	return s
}

func (s *Instance) SetIsSpot(v bool) *Instance {
	s.IsSpot = &v
	return s
}

func (s *Instance) SetIsolated(v bool) *Instance {
	s.Isolated = &v
	return s
}

func (s *Instance) SetLastState(v []map[string]interface{}) *Instance {
	s.LastState = v
	return s
}

func (s *Instance) SetNamespace(v string) *Instance {
	s.Namespace = &v
	return s
}

func (s *Instance) SetOriginalAmount(v float32) *Instance {
	s.OriginalAmount = &v
	return s
}

func (s *Instance) SetReadyProcesses(v int32) *Instance {
	s.ReadyProcesses = &v
	return s
}

func (s *Instance) SetReason(v string) *Instance {
	s.Reason = &v
	return s
}

func (s *Instance) SetReplicaName(v string) *Instance {
	s.ReplicaName = &v
	return s
}

func (s *Instance) SetResourceType(v string) *Instance {
	s.ResourceType = &v
	return s
}

func (s *Instance) SetRestartCount(v int32) *Instance {
	s.RestartCount = &v
	return s
}

func (s *Instance) SetRole(v string) *Instance {
	s.Role = &v
	return s
}

func (s *Instance) SetStartAt(v string) *Instance {
	s.StartAt = &v
	return s
}

func (s *Instance) SetStartTime(v string) *Instance {
	s.StartTime = &v
	return s
}

func (s *Instance) SetStatus(v string) *Instance {
	s.Status = &v
	return s
}

func (s *Instance) SetTenantHostIP(v string) *Instance {
	s.TenantHostIP = &v
	return s
}

func (s *Instance) SetTenantInstanceIP(v string) *Instance {
	s.TenantInstanceIP = &v
	return s
}

func (s *Instance) SetTotalProcesses(v int32) *Instance {
	s.TotalProcesses = &v
	return s
}

func (s *Instance) SetZone(v string) *Instance {
	s.Zone = &v
	return s
}

func (s *Instance) Validate() error {
	return dara.Validate(s)
}
