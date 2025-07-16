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
	// example:
	//
	// 0.444
	CurrentAmount *float32 `json:"CurrentAmount,omitempty" xml:"CurrentAmount,omitempty"`
	// example:
	//
	// 192.168.1.100
	ExternalIP *string `json:"ExternalIP,omitempty" xml:"ExternalIP,omitempty"`
	// example:
	//
	// 8080
	ExternalInstancePort *int32  `json:"ExternalInstancePort,omitempty" xml:"ExternalInstancePort,omitempty"`
	HostIP               *string `json:"HostIP,omitempty" xml:"HostIP,omitempty"`
	HostName             *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InnerIP              *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	InstanceName         *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstancePort         *int32  `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// false
	IsSpot *bool `json:"IsSpot,omitempty" xml:"IsSpot,omitempty"`
	// example:
	//
	// false
	Isolated  *bool                    `json:"Isolated,omitempty" xml:"Isolated,omitempty"`
	LastState []map[string]interface{} `json:"LastState,omitempty" xml:"LastState,omitempty" type:"Repeated"`
	Namespace *string                  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 2.2
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	ReadyProcesses *int32   `json:"ReadyProcesses,omitempty" xml:"ReadyProcesses,omitempty"`
	Reason         *string  `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// PublicResource
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RestartCount *int32  `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// example:
	//
	// Standard
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Deprecated
	StartAt   *string `json:"StartAt,omitempty" xml:"StartAt,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 192.168.0.39
	TenantHostIP *string `json:"TenantHostIP,omitempty" xml:"TenantHostIP,omitempty"`
	// example:
	//
	// 192.168.0.39
	TenantInstanceIP *string `json:"TenantInstanceIP,omitempty" xml:"TenantInstanceIP,omitempty"`
	TotalProcesses   *int32  `json:"TotalProcesses,omitempty" xml:"TotalProcesses,omitempty"`
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
