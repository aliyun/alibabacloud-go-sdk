// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ContainerInfo struct {
	CurrentReaon     *string `json:"CurrentReaon,omitempty" xml:"CurrentReaon,omitempty"`
	CurrentStatus    *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	CurrentTimestamp *string `json:"CurrentTimestamp,omitempty" xml:"CurrentTimestamp,omitempty"`
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	LastReason       *string `json:"LastReason,omitempty" xml:"LastReason,omitempty"`
	LastStatus       *string `json:"LastStatus,omitempty" xml:"LastStatus,omitempty"`
	LastTimestamp    *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Ready            *bool   `json:"Ready,omitempty" xml:"Ready,omitempty"`
	RestartCount     *int32  `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
}

func (s ContainerInfo) String() string {
	return tea.Prettify(s)
}

func (s ContainerInfo) GoString() string {
	return s.String()
}

func (s *ContainerInfo) SetCurrentReaon(v string) *ContainerInfo {
	s.CurrentReaon = &v
	return s
}

func (s *ContainerInfo) SetCurrentStatus(v string) *ContainerInfo {
	s.CurrentStatus = &v
	return s
}

func (s *ContainerInfo) SetCurrentTimestamp(v string) *ContainerInfo {
	s.CurrentTimestamp = &v
	return s
}

func (s *ContainerInfo) SetImage(v string) *ContainerInfo {
	s.Image = &v
	return s
}

func (s *ContainerInfo) SetLastReason(v string) *ContainerInfo {
	s.LastReason = &v
	return s
}

func (s *ContainerInfo) SetLastStatus(v string) *ContainerInfo {
	s.LastStatus = &v
	return s
}

func (s *ContainerInfo) SetLastTimestamp(v string) *ContainerInfo {
	s.LastTimestamp = &v
	return s
}

func (s *ContainerInfo) SetName(v string) *ContainerInfo {
	s.Name = &v
	return s
}

func (s *ContainerInfo) SetPort(v int32) *ContainerInfo {
	s.Port = &v
	return s
}

func (s *ContainerInfo) SetReady(v bool) *ContainerInfo {
	s.Ready = &v
	return s
}

func (s *ContainerInfo) SetRestartCount(v int32) *ContainerInfo {
	s.RestartCount = &v
	return s
}

type Group struct {
	AccessToken      *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	QueueService     *string `json:"QueueService,omitempty" xml:"QueueService,omitempty"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Group) String() string {
	return tea.Prettify(s)
}

func (s Group) GoString() string {
	return s.String()
}

func (s *Group) SetAccessToken(v string) *Group {
	s.AccessToken = &v
	return s
}

func (s *Group) SetClusterId(v string) *Group {
	s.ClusterId = &v
	return s
}

func (s *Group) SetCreateTime(v string) *Group {
	s.CreateTime = &v
	return s
}

func (s *Group) SetInternetEndpoint(v string) *Group {
	s.InternetEndpoint = &v
	return s
}

func (s *Group) SetIntranetEndpoint(v string) *Group {
	s.IntranetEndpoint = &v
	return s
}

func (s *Group) SetName(v string) *Group {
	s.Name = &v
	return s
}

func (s *Group) SetQueueService(v string) *Group {
	s.QueueService = &v
	return s
}

func (s *Group) SetUpdateTime(v string) *Group {
	s.UpdateTime = &v
	return s
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
	return tea.Prettify(s)
}

func (s Instance) GoString() string {
	return s.String()
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

type Resource struct {
	ClusterId             *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CpuCount              *int32                 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	CreateTime            *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtraData             map[string]interface{} `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	GpuCount              *int32                 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	InstanceCount         *int32                 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	Message               *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	PostPaidInstanceCount *int32                 `json:"PostPaidInstanceCount,omitempty" xml:"PostPaidInstanceCount,omitempty"`
	PrePaidInstanceCount  *int32                 `json:"PrePaidInstanceCount,omitempty" xml:"PrePaidInstanceCount,omitempty"`
	ResourceId            *string                `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName          *string                `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType          *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime            *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Vendor                *string                `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s Resource) String() string {
	return tea.Prettify(s)
}

func (s Resource) GoString() string {
	return s.String()
}

func (s *Resource) SetClusterId(v string) *Resource {
	s.ClusterId = &v
	return s
}

func (s *Resource) SetCpuCount(v int32) *Resource {
	s.CpuCount = &v
	return s
}

func (s *Resource) SetCreateTime(v string) *Resource {
	s.CreateTime = &v
	return s
}

func (s *Resource) SetExtraData(v map[string]interface{}) *Resource {
	s.ExtraData = v
	return s
}

func (s *Resource) SetGpuCount(v int32) *Resource {
	s.GpuCount = &v
	return s
}

func (s *Resource) SetInstanceCount(v int32) *Resource {
	s.InstanceCount = &v
	return s
}

func (s *Resource) SetMessage(v string) *Resource {
	s.Message = &v
	return s
}

func (s *Resource) SetPostPaidInstanceCount(v int32) *Resource {
	s.PostPaidInstanceCount = &v
	return s
}

func (s *Resource) SetPrePaidInstanceCount(v int32) *Resource {
	s.PrePaidInstanceCount = &v
	return s
}

func (s *Resource) SetResourceId(v string) *Resource {
	s.ResourceId = &v
	return s
}

func (s *Resource) SetResourceName(v string) *Resource {
	s.ResourceName = &v
	return s
}

func (s *Resource) SetResourceType(v string) *Resource {
	s.ResourceType = &v
	return s
}

func (s *Resource) SetStatus(v string) *Resource {
	s.Status = &v
	return s
}

func (s *Resource) SetUpdateTime(v string) *Resource {
	s.UpdateTime = &v
	return s
}

func (s *Resource) SetVendor(v string) *Resource {
	s.Vendor = &v
	return s
}

type ResourceInstance struct {
	Arch                   *string  `json:"Arch,omitempty" xml:"Arch,omitempty"`
	AutoRenewal            *bool    `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	ChargeType             *string  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime             *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpiredTime            *string  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceCpuCount       *int32   `json:"InstanceCpuCount,omitempty" xml:"InstanceCpuCount,omitempty"`
	InstanceGpuCount       *int32   `json:"InstanceGpuCount,omitempty" xml:"InstanceGpuCount,omitempty"`
	InstanceGpuMemory      *string  `json:"InstanceGpuMemory,omitempty" xml:"InstanceGpuMemory,omitempty"`
	InstanceId             *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIp             *string  `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	InstanceMemory         *string  `json:"InstanceMemory,omitempty" xml:"InstanceMemory,omitempty"`
	InstanceName           *string  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus         *string  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceSystemDiskSize *int32   `json:"InstanceSystemDiskSize,omitempty" xml:"InstanceSystemDiskSize,omitempty"`
	InstanceTenantIp       *string  `json:"InstanceTenantIp,omitempty" xml:"InstanceTenantIp,omitempty"`
	InstanceType           *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceUsedCpu        *float32 `json:"InstanceUsedCpu,omitempty" xml:"InstanceUsedCpu,omitempty"`
	InstanceUsedGpu        *float32 `json:"InstanceUsedGpu,omitempty" xml:"InstanceUsedGpu,omitempty"`
	InstanceUsedGpuMemory  *string  `json:"InstanceUsedGpuMemory,omitempty" xml:"InstanceUsedGpuMemory,omitempty"`
	InstanceUsedMemory     *string  `json:"InstanceUsedMemory,omitempty" xml:"InstanceUsedMemory,omitempty"`
	Region                 *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId             *string  `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Zone                   *string  `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s ResourceInstance) String() string {
	return tea.Prettify(s)
}

func (s ResourceInstance) GoString() string {
	return s.String()
}

func (s *ResourceInstance) SetArch(v string) *ResourceInstance {
	s.Arch = &v
	return s
}

func (s *ResourceInstance) SetAutoRenewal(v bool) *ResourceInstance {
	s.AutoRenewal = &v
	return s
}

func (s *ResourceInstance) SetChargeType(v string) *ResourceInstance {
	s.ChargeType = &v
	return s
}

func (s *ResourceInstance) SetCreateTime(v string) *ResourceInstance {
	s.CreateTime = &v
	return s
}

func (s *ResourceInstance) SetExpiredTime(v string) *ResourceInstance {
	s.ExpiredTime = &v
	return s
}

func (s *ResourceInstance) SetInstanceCpuCount(v int32) *ResourceInstance {
	s.InstanceCpuCount = &v
	return s
}

func (s *ResourceInstance) SetInstanceGpuCount(v int32) *ResourceInstance {
	s.InstanceGpuCount = &v
	return s
}

func (s *ResourceInstance) SetInstanceGpuMemory(v string) *ResourceInstance {
	s.InstanceGpuMemory = &v
	return s
}

func (s *ResourceInstance) SetInstanceId(v string) *ResourceInstance {
	s.InstanceId = &v
	return s
}

func (s *ResourceInstance) SetInstanceIp(v string) *ResourceInstance {
	s.InstanceIp = &v
	return s
}

func (s *ResourceInstance) SetInstanceMemory(v string) *ResourceInstance {
	s.InstanceMemory = &v
	return s
}

func (s *ResourceInstance) SetInstanceName(v string) *ResourceInstance {
	s.InstanceName = &v
	return s
}

func (s *ResourceInstance) SetInstanceStatus(v string) *ResourceInstance {
	s.InstanceStatus = &v
	return s
}

func (s *ResourceInstance) SetInstanceSystemDiskSize(v int32) *ResourceInstance {
	s.InstanceSystemDiskSize = &v
	return s
}

func (s *ResourceInstance) SetInstanceTenantIp(v string) *ResourceInstance {
	s.InstanceTenantIp = &v
	return s
}

func (s *ResourceInstance) SetInstanceType(v string) *ResourceInstance {
	s.InstanceType = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedCpu(v float32) *ResourceInstance {
	s.InstanceUsedCpu = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedGpu(v float32) *ResourceInstance {
	s.InstanceUsedGpu = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedGpuMemory(v string) *ResourceInstance {
	s.InstanceUsedGpuMemory = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedMemory(v string) *ResourceInstance {
	s.InstanceUsedMemory = &v
	return s
}

func (s *ResourceInstance) SetRegion(v string) *ResourceInstance {
	s.Region = &v
	return s
}

func (s *ResourceInstance) SetResourceId(v string) *ResourceInstance {
	s.ResourceId = &v
	return s
}

func (s *ResourceInstance) SetZone(v string) *ResourceInstance {
	s.Zone = &v
	return s
}

type ResourceInstanceWorker struct {
	CpuLimit     *int32  `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	CpuRequest   *int32  `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	GpuLimit     *int32  `json:"GpuLimit,omitempty" xml:"GpuLimit,omitempty"`
	GpuRequest   *int32  `json:"GpuRequest,omitempty" xml:"GpuRequest,omitempty"`
	MemoryLimit  *int32  `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	MemoryRquest *int32  `json:"MemoryRquest,omitempty" xml:"MemoryRquest,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Ready        *bool   `json:"Ready,omitempty" xml:"Ready,omitempty"`
	RestartCount *int32  `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResourceInstanceWorker) String() string {
	return tea.Prettify(s)
}

func (s ResourceInstanceWorker) GoString() string {
	return s.String()
}

func (s *ResourceInstanceWorker) SetCpuLimit(v int32) *ResourceInstanceWorker {
	s.CpuLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetCpuRequest(v int32) *ResourceInstanceWorker {
	s.CpuRequest = &v
	return s
}

func (s *ResourceInstanceWorker) SetGpuLimit(v int32) *ResourceInstanceWorker {
	s.GpuLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetGpuRequest(v int32) *ResourceInstanceWorker {
	s.GpuRequest = &v
	return s
}

func (s *ResourceInstanceWorker) SetMemoryLimit(v int32) *ResourceInstanceWorker {
	s.MemoryLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetMemoryRquest(v int32) *ResourceInstanceWorker {
	s.MemoryRquest = &v
	return s
}

func (s *ResourceInstanceWorker) SetName(v string) *ResourceInstanceWorker {
	s.Name = &v
	return s
}

func (s *ResourceInstanceWorker) SetReady(v bool) *ResourceInstanceWorker {
	s.Ready = &v
	return s
}

func (s *ResourceInstanceWorker) SetRestartCount(v int32) *ResourceInstanceWorker {
	s.RestartCount = &v
	return s
}

func (s *ResourceInstanceWorker) SetServiceName(v string) *ResourceInstanceWorker {
	s.ServiceName = &v
	return s
}

func (s *ResourceInstanceWorker) SetStartTime(v string) *ResourceInstanceWorker {
	s.StartTime = &v
	return s
}

func (s *ResourceInstanceWorker) SetStatus(v string) *ResourceInstanceWorker {
	s.Status = &v
	return s
}

type Service struct {
	AccessToken               *string          `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AppConfig                 *string          `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	AppSpecName               *string          `json:"AppSpecName,omitempty" xml:"AppSpecName,omitempty"`
	AppType                   *string          `json:"AppType,omitempty" xml:"AppType,omitempty"`
	AppVersion                *string          `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CallerUid                 *string          `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	Cpu                       *int32           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreateTime                *string          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentVersion            *int32           `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	ExtraData                 *string          `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	Gpu                       *int32           `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image                     *string          `json:"Image,omitempty" xml:"Image,omitempty"`
	InternetEndpoint          *string          `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetEndpoint          *string          `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Labels                    []*ServiceLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion             *int32           `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	Memory                    *int32           `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Message                   *string          `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace                 *string          `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ParentUid                 *string          `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	PendingInstance           *int32           `json:"PendingInstance,omitempty" xml:"PendingInstance,omitempty"`
	Reason                    *string          `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Region                    *string          `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestId                 *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource                  *string          `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceAlias             *string          `json:"ResourceAlias,omitempty" xml:"ResourceAlias,omitempty"`
	Role                      *string          `json:"Role,omitempty" xml:"Role,omitempty"`
	RoleAttrs                 *string          `json:"RoleAttrs,omitempty" xml:"RoleAttrs,omitempty"`
	RunningInstance           *int32           `json:"RunningInstance,omitempty" xml:"RunningInstance,omitempty"`
	SafetyLock                *string          `json:"SafetyLock,omitempty" xml:"SafetyLock,omitempty"`
	SecondaryInternetEndpoint *string          `json:"SecondaryInternetEndpoint,omitempty" xml:"SecondaryInternetEndpoint,omitempty"`
	SecondaryIntranetEndpoint *string          `json:"SecondaryIntranetEndpoint,omitempty" xml:"SecondaryIntranetEndpoint,omitempty"`
	ServiceConfig             *string          `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceGroup              *string          `json:"ServiceGroup,omitempty" xml:"ServiceGroup,omitempty"`
	ServiceId                 *string          `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName               *string          `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceUid                *string          `json:"ServiceUid,omitempty" xml:"ServiceUid,omitempty"`
	Source                    *string          `json:"Source,omitempty" xml:"Source,omitempty"`
	Status                    *string          `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalInstance             *int32           `json:"TotalInstance,omitempty" xml:"TotalInstance,omitempty"`
	UpdateTime                *string          `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Weight                    *int32           `json:"Weight,omitempty" xml:"Weight,omitempty"`
	WorkspaceId               *string          `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Service) String() string {
	return tea.Prettify(s)
}

func (s Service) GoString() string {
	return s.String()
}

func (s *Service) SetAccessToken(v string) *Service {
	s.AccessToken = &v
	return s
}

func (s *Service) SetAppConfig(v string) *Service {
	s.AppConfig = &v
	return s
}

func (s *Service) SetAppSpecName(v string) *Service {
	s.AppSpecName = &v
	return s
}

func (s *Service) SetAppType(v string) *Service {
	s.AppType = &v
	return s
}

func (s *Service) SetAppVersion(v string) *Service {
	s.AppVersion = &v
	return s
}

func (s *Service) SetCallerUid(v string) *Service {
	s.CallerUid = &v
	return s
}

func (s *Service) SetCpu(v int32) *Service {
	s.Cpu = &v
	return s
}

func (s *Service) SetCreateTime(v string) *Service {
	s.CreateTime = &v
	return s
}

func (s *Service) SetCurrentVersion(v int32) *Service {
	s.CurrentVersion = &v
	return s
}

func (s *Service) SetExtraData(v string) *Service {
	s.ExtraData = &v
	return s
}

func (s *Service) SetGpu(v int32) *Service {
	s.Gpu = &v
	return s
}

func (s *Service) SetImage(v string) *Service {
	s.Image = &v
	return s
}

func (s *Service) SetInternetEndpoint(v string) *Service {
	s.InternetEndpoint = &v
	return s
}

func (s *Service) SetIntranetEndpoint(v string) *Service {
	s.IntranetEndpoint = &v
	return s
}

func (s *Service) SetLabels(v []*ServiceLabels) *Service {
	s.Labels = v
	return s
}

func (s *Service) SetLatestVersion(v int32) *Service {
	s.LatestVersion = &v
	return s
}

func (s *Service) SetMemory(v int32) *Service {
	s.Memory = &v
	return s
}

func (s *Service) SetMessage(v string) *Service {
	s.Message = &v
	return s
}

func (s *Service) SetNamespace(v string) *Service {
	s.Namespace = &v
	return s
}

func (s *Service) SetParentUid(v string) *Service {
	s.ParentUid = &v
	return s
}

func (s *Service) SetPendingInstance(v int32) *Service {
	s.PendingInstance = &v
	return s
}

func (s *Service) SetReason(v string) *Service {
	s.Reason = &v
	return s
}

func (s *Service) SetRegion(v string) *Service {
	s.Region = &v
	return s
}

func (s *Service) SetRequestId(v string) *Service {
	s.RequestId = &v
	return s
}

func (s *Service) SetResource(v string) *Service {
	s.Resource = &v
	return s
}

func (s *Service) SetResourceAlias(v string) *Service {
	s.ResourceAlias = &v
	return s
}

func (s *Service) SetRole(v string) *Service {
	s.Role = &v
	return s
}

func (s *Service) SetRoleAttrs(v string) *Service {
	s.RoleAttrs = &v
	return s
}

func (s *Service) SetRunningInstance(v int32) *Service {
	s.RunningInstance = &v
	return s
}

func (s *Service) SetSafetyLock(v string) *Service {
	s.SafetyLock = &v
	return s
}

func (s *Service) SetSecondaryInternetEndpoint(v string) *Service {
	s.SecondaryInternetEndpoint = &v
	return s
}

func (s *Service) SetSecondaryIntranetEndpoint(v string) *Service {
	s.SecondaryIntranetEndpoint = &v
	return s
}

func (s *Service) SetServiceConfig(v string) *Service {
	s.ServiceConfig = &v
	return s
}

func (s *Service) SetServiceGroup(v string) *Service {
	s.ServiceGroup = &v
	return s
}

func (s *Service) SetServiceId(v string) *Service {
	s.ServiceId = &v
	return s
}

func (s *Service) SetServiceName(v string) *Service {
	s.ServiceName = &v
	return s
}

func (s *Service) SetServiceUid(v string) *Service {
	s.ServiceUid = &v
	return s
}

func (s *Service) SetSource(v string) *Service {
	s.Source = &v
	return s
}

func (s *Service) SetStatus(v string) *Service {
	s.Status = &v
	return s
}

func (s *Service) SetTotalInstance(v int32) *Service {
	s.TotalInstance = &v
	return s
}

func (s *Service) SetUpdateTime(v string) *Service {
	s.UpdateTime = &v
	return s
}

func (s *Service) SetWeight(v int32) *Service {
	s.Weight = &v
	return s
}

func (s *Service) SetWorkspaceId(v string) *Service {
	s.WorkspaceId = &v
	return s
}

type ServiceLabels struct {
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s ServiceLabels) String() string {
	return tea.Prettify(s)
}

func (s ServiceLabels) GoString() string {
	return s.String()
}

func (s *ServiceLabels) SetLabelKey(v string) *ServiceLabels {
	s.LabelKey = &v
	return s
}

func (s *ServiceLabels) SetLabelValue(v string) *ServiceLabels {
	s.LabelValue = &v
	return s
}

type CloneServiceRequest struct {
	// The label of the service to be cloned.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The request body. For more information, see [CreateService](https://help.aliyun.com/document_detail/412086.html).
	//
	// example:
	//
	// {   "name": "foo",   "model_path": "http://path/to/model.tar.gz",   "processor": "tensorflow_cpu",   "metadata": {     "instance": 2,     "memory": 7000,     "cpu": 4   } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneServiceRequest) GoString() string {
	return s.String()
}

func (s *CloneServiceRequest) SetLabels(v map[string]*string) *CloneServiceRequest {
	s.Labels = v
	return s
}

func (s *CloneServiceRequest) SetBody(v string) *CloneServiceRequest {
	s.Body = &v
	return s
}

type CloneServiceShrinkRequest struct {
	// The label of the service to be cloned.
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The request body. For more information, see [CreateService](https://help.aliyun.com/document_detail/412086.html).
	//
	// example:
	//
	// {   "name": "foo",   "model_path": "http://path/to/model.tar.gz",   "processor": "tensorflow_cpu",   "metadata": {     "instance": 2,     "memory": 7000,     "cpu": 4   } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneServiceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CloneServiceShrinkRequest) SetLabelsShrink(v string) *CloneServiceShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *CloneServiceShrinkRequest) SetBody(v string) *CloneServiceShrinkRequest {
	s.Body = &v
	return s
}

type CloneServiceResponseBody struct {
	// The public endpoint of the service.
	//
	// example:
	//
	// http://10123*****.cn-shanghai.aliyuncs.com/api/predict/echo
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The private endpoint of the service.
	//
	// example:
	//
	// http://10123*****.vpc.cn-shanghai.aliyuncs.com/api/predict/echo
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// eas-m-r9knx7n9guf2p*****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service status.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CloneServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CloneServiceResponseBody) SetInternetEndpoint(v string) *CloneServiceResponseBody {
	s.InternetEndpoint = &v
	return s
}

func (s *CloneServiceResponseBody) SetIntranetEndpoint(v string) *CloneServiceResponseBody {
	s.IntranetEndpoint = &v
	return s
}

func (s *CloneServiceResponseBody) SetRequestId(v string) *CloneServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneServiceResponseBody) SetServiceId(v string) *CloneServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CloneServiceResponseBody) SetServiceName(v string) *CloneServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CloneServiceResponseBody) SetStatus(v string) *CloneServiceResponseBody {
	s.Status = &v
	return s
}

type CloneServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneServiceResponse) GoString() string {
	return s.String()
}

func (s *CloneServiceResponse) SetHeaders(v map[string]*string) *CloneServiceResponse {
	s.Headers = v
	return s
}

func (s *CloneServiceResponse) SetStatusCode(v int32) *CloneServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneServiceResponse) SetBody(v *CloneServiceResponseBody) *CloneServiceResponse {
	s.Body = v
	return s
}

type CommitServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommitServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CommitServiceResponseBody) SetMessage(v string) *CommitServiceResponseBody {
	s.Message = &v
	return s
}

func (s *CommitServiceResponseBody) SetRequestId(v string) *CommitServiceResponseBody {
	s.RequestId = &v
	return s
}

type CommitServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommitServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommitServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitServiceResponse) GoString() string {
	return s.String()
}

func (s *CommitServiceResponse) SetHeaders(v map[string]*string) *CommitServiceResponse {
	s.Headers = v
	return s
}

func (s *CommitServiceResponse) SetStatusCode(v int32) *CommitServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitServiceResponse) SetBody(v *CommitServiceResponseBody) *CommitServiceResponse {
	s.Body = v
	return s
}

type CreateAclPolicyRequest struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway.
	//
	// This parameter is required.
	AclPolicyList []*CreateAclPolicyRequestAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateAclPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyRequest) SetAclPolicyList(v []*CreateAclPolicyRequestAclPolicyList) *CreateAclPolicyRequest {
	s.AclPolicyList = v
	return s
}

func (s *CreateAclPolicyRequest) SetVpcId(v string) *CreateAclPolicyRequest {
	s.VpcId = &v
	return s
}

type CreateAclPolicyRequestAclPolicyList struct {
	// The comment on the IP CIDR block in the VPC that can access the private gateway.
	//
	// example:
	//
	// default
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP CIDR block in the VPC that can access the private gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.23.XX.XX/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s CreateAclPolicyRequestAclPolicyList) String() string {
	return tea.Prettify(s)
}

func (s CreateAclPolicyRequestAclPolicyList) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyRequestAclPolicyList) SetComment(v string) *CreateAclPolicyRequestAclPolicyList {
	s.Comment = &v
	return s
}

func (s *CreateAclPolicyRequestAclPolicyList) SetEntry(v string) *CreateAclPolicyRequestAclPolicyList {
	s.Entry = &v
	return s
}

type CreateAclPolicyShrinkRequest struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway.
	//
	// This parameter is required.
	AclPolicyListShrink *string `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty"`
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateAclPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyShrinkRequest) SetAclPolicyListShrink(v string) *CreateAclPolicyShrinkRequest {
	s.AclPolicyListShrink = &v
	return s
}

func (s *CreateAclPolicyShrinkRequest) SetVpcId(v string) *CreateAclPolicyShrinkRequest {
	s.VpcId = &v
	return s
}

type CreateAclPolicyResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully add acl policy for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAclPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyResponseBody) SetGatewayId(v string) *CreateAclPolicyResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateAclPolicyResponseBody) SetMessage(v string) *CreateAclPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAclPolicyResponseBody) SetRequestId(v string) *CreateAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateAclPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAclPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAclPolicyResponse) SetHeaders(v map[string]*string) *CreateAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAclPolicyResponse) SetStatusCode(v int32) *CreateAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAclPolicyResponse) SetBody(v *CreateAclPolicyResponseBody) *CreateAclPolicyResponse {
	s.Body = v
	return s
}

type CreateAppServiceRequest struct {
	// The quota ID.
	//
	// example:
	//
	// abcdef
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The application service type.
	//
	// Valid values:
	//
	// 	- LLM
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// LLM
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The application version.
	//
	// example:
	//
	// v1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The additional configurations that are required for service deployment.
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The number of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service specifications. Valid values:
	//
	// 	- llama_7b_fp16
	//
	// 	- llama_7b_int8
	//
	// 	- llama_13b_fp16
	//
	// 	- llama_7b_int8
	//
	// 	- chatglm_6b_fp16
	//
	// 	- chatglm_6b_int8
	//
	// 	- chatglm2_6b_fp16
	//
	// 	- baichuan_7b_int8
	//
	// 	- baichuan_13b_fp16
	//
	// 	- baichuan_7b_fp16
	//
	// This parameter is required.
	//
	// example:
	//
	// llama_7b_fp16
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
}

func (s CreateAppServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateAppServiceRequest) SetQuotaId(v string) *CreateAppServiceRequest {
	s.QuotaId = &v
	return s
}

func (s *CreateAppServiceRequest) SetWorkspaceId(v string) *CreateAppServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAppServiceRequest) SetAppType(v string) *CreateAppServiceRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppServiceRequest) SetAppVersion(v string) *CreateAppServiceRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppServiceRequest) SetConfig(v map[string]interface{}) *CreateAppServiceRequest {
	s.Config = v
	return s
}

func (s *CreateAppServiceRequest) SetReplicas(v int32) *CreateAppServiceRequest {
	s.Replicas = &v
	return s
}

func (s *CreateAppServiceRequest) SetServiceName(v string) *CreateAppServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateAppServiceRequest) SetServiceSpec(v string) *CreateAppServiceRequest {
	s.ServiceSpec = &v
	return s
}

type CreateAppServiceResponseBody struct {
	// The public endpoint of the service.
	//
	// example:
	//
	// http://pai-eas.cn-shanghai.****
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The internal endpoint of the service.
	//
	// example:
	//
	// http://pai-eas-vpc.cn-shanghai.****
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// The region ID of the service.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// eas-m-aaxxxddf
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service state.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAppServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppServiceResponseBody) SetInternetEndpoint(v string) *CreateAppServiceResponseBody {
	s.InternetEndpoint = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetIntranetEndpoint(v string) *CreateAppServiceResponseBody {
	s.IntranetEndpoint = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetRegion(v string) *CreateAppServiceResponseBody {
	s.Region = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetRequestId(v string) *CreateAppServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetServiceId(v string) *CreateAppServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetServiceName(v string) *CreateAppServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateAppServiceResponseBody) SetStatus(v string) *CreateAppServiceResponseBody {
	s.Status = &v
	return s
}

type CreateAppServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateAppServiceResponse) SetHeaders(v map[string]*string) *CreateAppServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateAppServiceResponse) SetStatusCode(v int32) *CreateAppServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppServiceResponse) SetBody(v *CreateAppServiceResponseBody) *CreateAppServiceResponse {
	s.Body = v
	return s
}

type CreateBenchmarkTaskRequest struct {
	// The request body. The body includes the parameters that are set to create a stress testing task.
	//
	// example:
	//
	// {
	//
	//     "base": {
	//
	//         "duration": 600
	//
	//     },
	//
	//     "service": {
	//
	//         "serviceName": "test_service",
	//
	//         "requestToken": "test_token"
	//
	//     },
	//
	//     "data": {
	//
	//         "path": "https://larec-benchmark-cd.oss-cn-chengdu.aliyuncs.com/youbei/sv_dbmtl/data/youbei.warmup.tf.bin",
	//
	//         "dataType": "binary"
	//
	//     },
	//
	//     "optional": {
	//
	//        "maxRt": 100
	//
	//     }
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBenchmarkTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskRequest) SetBody(v string) *CreateBenchmarkTaskRequest {
	s.Body = &v
	return s
}

type CreateBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark  task [foo] is Creating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the region where the stress testing task is performed.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the stress testing task.
	//
	// example:
	//
	// benchmark-larec-test-1076
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskResponseBody) SetMessage(v string) *CreateBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetRegion(v string) *CreateBenchmarkTaskResponseBody {
	s.Region = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetRequestId(v string) *CreateBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetTaskName(v string) *CreateBenchmarkTaskResponseBody {
	s.TaskName = &v
	return s
}

type CreateBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskResponse) SetHeaders(v map[string]*string) *CreateBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateBenchmarkTaskResponse) SetStatusCode(v int32) *CreateBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBenchmarkTaskResponse) SetBody(v *CreateBenchmarkTaskResponseBody) *CreateBenchmarkTaskResponse {
	s.Body = v
	return s
}

type CreateGatewayRequest struct {
	// The resource group ID. To obtain a resource group ID, see the ResourceId field in the response of the [ListResources](https://help.aliyun.com/document_detail/412133.html) operation.
	//
	// example:
	//
	// eas-r-4gt8twzwllfo******
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	AutoRenewal  *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	ChargeType   *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Specifies whether to enable Internet access. Default value: false.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableInternet *bool `json:"EnableInternet,omitempty" xml:"EnableInternet,omitempty"`
	// Specifies whether to enable private access. Default value: true.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableIntranet *bool `json:"EnableIntranet,omitempty" xml:"EnableIntranet,omitempty"`
	// The instance type used by the private gateway. Valid values:
	//
	// 	- 2c4g
	//
	// 	- 4c8g
	//
	// 	- 8c16g
	//
	// 	- 16c32g
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.c6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The alias of the private gateway.
	//
	// example:
	//
	// mygateway1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of nodes in the private gateway.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s CreateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) SetResourceName(v string) *CreateGatewayRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateGatewayRequest) SetAutoRenewal(v bool) *CreateGatewayRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateGatewayRequest) SetChargeType(v string) *CreateGatewayRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateGatewayRequest) SetEnableInternet(v bool) *CreateGatewayRequest {
	s.EnableInternet = &v
	return s
}

func (s *CreateGatewayRequest) SetEnableIntranet(v bool) *CreateGatewayRequest {
	s.EnableIntranet = &v
	return s
}

func (s *CreateGatewayRequest) SetInstanceType(v string) *CreateGatewayRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateGatewayRequest) SetName(v string) *CreateGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateGatewayRequest) SetReplicas(v int32) *CreateGatewayRequest {
	s.Replicas = &v
	return s
}

type CreateGatewayResponseBody struct {
	// The region ID of the private gateway.
	//
	// example:
	//
	// cn-hangzhou
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully create gateway.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponseBody) SetClusterId(v string) *CreateGatewayResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetGatewayId(v string) *CreateGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayResponseBody) SetMessage(v string) *CreateGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayResponseBody) SetRequestId(v string) *CreateGatewayResponseBody {
	s.RequestId = &v
	return s
}

type CreateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayResponse) SetHeaders(v map[string]*string) *CreateGatewayResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayResponse) SetStatusCode(v int32) *CreateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayResponse) SetBody(v *CreateGatewayResponseBody) *CreateGatewayResponse {
	s.Body = v
	return s
}

type CreateGatewayIntranetLinkedVpcRequest struct {
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vbqn2at0kljjxxxx****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcRequest) SetVSwitchId(v string) *CreateGatewayIntranetLinkedVpcRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcRequest) SetVpcId(v string) *CreateGatewayIntranetLinkedVpcRequest {
	s.VpcId = &v
	return s
}

type CreateGatewayIntranetLinkedVpcResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully add intranet linked vpc for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) SetGatewayId(v string) *CreateGatewayIntranetLinkedVpcResponseBody {
	s.GatewayId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) SetMessage(v string) *CreateGatewayIntranetLinkedVpcResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponseBody) SetRequestId(v string) *CreateGatewayIntranetLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

type CreateGatewayIntranetLinkedVpcResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayIntranetLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcResponse) SetHeaders(v map[string]*string) *CreateGatewayIntranetLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponse) SetStatusCode(v int32) *CreateGatewayIntranetLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcResponse) SetBody(v *CreateGatewayIntranetLinkedVpcResponseBody) *CreateGatewayIntranetLinkedVpcResponse {
	s.Body = v
	return s
}

type CreateResourceRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: the subscription billing method.
	//
	// 	- PostPaid: the pay-as-you-go billing method.
	//
	// >  This parameter is required when the ResourceType parameter is set to Dedicated.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of ECS instances.
	//
	// >  This parameter is required when the ResourceType parameter is set to Dedicated.
	//
	// example:
	//
	// 5
	EcsInstanceCount *int32 `json:"EcsInstanceCount,omitempty" xml:"EcsInstanceCount,omitempty"`
	// The type of the Elastic Compute Service (ECS) instance.
	//
	// >  This parameter is required when the ResourceType parameter is set to Dedicated.
	//
	// example:
	//
	// ecs.c6.8xlarge
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- Dedicated: the dedicated resource group.
	//
	// 	- SelfManaged: the self-managed resource group.
	//
	// >  If you use a self-managed resource group, you must configure a whitelist.
	//
	// example:
	//
	// Dedicated
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The configurations of the self-managed resource group.
	SelfManagedResourceOptions *CreateResourceRequestSelfManagedResourceOptions `json:"SelfManagedResourceOptions,omitempty" xml:"SelfManagedResourceOptions,omitempty" type:"Struct"`
	// The size of the system disk. Unit: GiB. Valid values: 200 to 2000. Default value: 200.
	//
	// example:
	//
	// 200
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The ID of the zone in which the instance resides.
	//
	// example:
	//
	// cn-shanghai-f
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetAutoRenewal(v bool) *CreateResourceRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateResourceRequest) SetChargeType(v string) *CreateResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateResourceRequest) SetEcsInstanceCount(v int32) *CreateResourceRequest {
	s.EcsInstanceCount = &v
	return s
}

func (s *CreateResourceRequest) SetEcsInstanceType(v string) *CreateResourceRequest {
	s.EcsInstanceType = &v
	return s
}

func (s *CreateResourceRequest) SetResourceType(v string) *CreateResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceRequest) SetSelfManagedResourceOptions(v *CreateResourceRequestSelfManagedResourceOptions) *CreateResourceRequest {
	s.SelfManagedResourceOptions = v
	return s
}

func (s *CreateResourceRequest) SetSystemDiskSize(v int32) *CreateResourceRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateResourceRequest) SetZone(v string) *CreateResourceRequest {
	s.Zone = &v
	return s
}

type CreateResourceRequestSelfManagedResourceOptions struct {
	// The ID of the self-managed cluster.
	//
	// example:
	//
	// cf0386f250f2545689ca7fdd1cd******
	ExternalClusterId *string `json:"ExternalClusterId,omitempty" xml:"ExternalClusterId,omitempty"`
	// The tag key-value pairs for nodes.
	NodeMatchLabels map[string]*string `json:"NodeMatchLabels,omitempty" xml:"NodeMatchLabels,omitempty"`
	// Tolerations for nodes.
	NodeTolerations []*CreateResourceRequestSelfManagedResourceOptionsNodeTolerations `json:"NodeTolerations,omitempty" xml:"NodeTolerations,omitempty" type:"Repeated"`
	// The name of the RAM user to which the permissions on Elastic Algorithm Service of Platform for AI (PAI-EAS) are granted.
	//
	// example:
	//
	// clusterrole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateResourceRequestSelfManagedResourceOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequestSelfManagedResourceOptions) GoString() string {
	return s.String()
}

func (s *CreateResourceRequestSelfManagedResourceOptions) SetExternalClusterId(v string) *CreateResourceRequestSelfManagedResourceOptions {
	s.ExternalClusterId = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptions) SetNodeMatchLabels(v map[string]*string) *CreateResourceRequestSelfManagedResourceOptions {
	s.NodeMatchLabels = v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptions) SetNodeTolerations(v []*CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) *CreateResourceRequestSelfManagedResourceOptions {
	s.NodeTolerations = v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptions) SetRoleName(v string) *CreateResourceRequestSelfManagedResourceOptions {
	s.RoleName = &v
	return s
}

type CreateResourceRequestSelfManagedResourceOptionsNodeTolerations struct {
	// The result.
	//
	// Valid values:
	//
	// 	- PreferNoSchedule
	//
	// 	- NoSchedule
	//
	// 	- NoExecute
	//
	// example:
	//
	// NoSchedule
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The key name.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The relationship between key names and key values.
	//
	// Valid values:
	//
	// 	- Equal
	//
	// 	- Exists
	//
	// example:
	//
	// Equal
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The key value.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) GoString() string {
	return s.String()
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetEffect(v string) *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Effect = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetKey(v string) *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Key = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetOperator(v string) *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Operator = &v
	return s
}

func (s *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetValue(v string) *CreateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Value = &v
	return s
}

type CreateResourceResponseBody struct {
	// The ID of the cluster to which the resource group belongs.
	//
	// example:
	//
	// cn-shanghai
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The user ID (UID) of the resource group owner.
	//
	// example:
	//
	// 14401087478****
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// eas-r-h7lcw24dyqztwxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// MyResource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetClusterId(v string) *CreateResourceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateResourceResponseBody) SetInstanceIds(v []*string) *CreateResourceResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateResourceResponseBody) SetOwnerUid(v string) *CreateResourceResponseBody {
	s.OwnerUid = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceId(v string) *CreateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceName(v string) *CreateResourceResponseBody {
	s.ResourceName = &v
	return s
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type CreateResourceInstancesRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of instances that you want to create. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	EcsInstanceCount *int32 `json:"EcsInstanceCount,omitempty" xml:"EcsInstanceCount,omitempty"`
	// The type of the Elastic Compute Service (ECS) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.s6-c1m2.xlarge
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values: 200 to 2000. Default value: 200.
	//
	// example:
	//
	// 200
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The user-defined information. This parameter is not in use.
	//
	// example:
	//
	// x112223333
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The zone to which the instance belongs.
	//
	// example:
	//
	// cn-shanghai-f
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s CreateResourceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesRequest) SetAutoRenewal(v bool) *CreateResourceInstancesRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetChargeType(v string) *CreateResourceInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetEcsInstanceCount(v int32) *CreateResourceInstancesRequest {
	s.EcsInstanceCount = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetEcsInstanceType(v string) *CreateResourceInstancesRequest {
	s.EcsInstanceType = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetSystemDiskSize(v int32) *CreateResourceInstancesRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetUserData(v string) *CreateResourceInstancesRequest {
	s.UserData = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetZone(v string) *CreateResourceInstancesRequest {
	s.Zone = &v
	return s
}

type CreateResourceInstancesResponseBody struct {
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// Create 5 new ecs instance(s) in resource [eas-r-asdasdasd] successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesResponseBody) SetInstanceIds(v []*string) *CreateResourceInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateResourceInstancesResponseBody) SetMessage(v string) *CreateResourceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourceInstancesResponseBody) SetRequestId(v string) *CreateResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesResponse) SetHeaders(v map[string]*string) *CreateResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceInstancesResponse) SetStatusCode(v int32) *CreateResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceInstancesResponse) SetBody(v *CreateResourceInstancesResponseBody) *CreateResourceInstancesResponse {
	s.Body = v
	return s
}

type CreateResourceLogRequest struct {
	// The Logstore of Log Service. For more information about how to query a Logstore, see [ListLogStores](https://help.aliyun.com/document_detail/426970.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// access_log
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project that is associated with the resource group. For more information about how to query the project, see [ListProject](https://help.aliyun.com/document_detail/74955.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// eas-r-asdasdasd-sls
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateResourceLogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceLogRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceLogRequest) SetLogStore(v string) *CreateResourceLogRequest {
	s.LogStore = &v
	return s
}

func (s *CreateResourceLogRequest) SetProjectName(v string) *CreateResourceLogRequest {
	s.ProjectName = &v
	return s
}

type CreateResourceLogResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Resource [eas-r-asdasdasd] is deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceLogResponseBody) SetMessage(v string) *CreateResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourceLogResponseBody) SetRequestId(v string) *CreateResourceLogResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceLogResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceLogResponse) SetHeaders(v map[string]*string) *CreateResourceLogResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceLogResponse) SetStatusCode(v int32) *CreateResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceLogResponse) SetBody(v *CreateResourceLogResponseBody) *CreateResourceLogResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	// Specifies whether to enter development mode.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Develop *string `json:"Develop,omitempty" xml:"Develop,omitempty"`
	// The custom label.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The request body. For more information about the key request parameters, see **Table 1. Request body parameters*	- and **Table 2. Metadata parameters**. For more information about all related parameters, see [Parameters of model services](https://help.aliyun.com/document_detail/450525.html).
	//
	// example:
	//
	// Service deployment by using an image:
	//
	// {
	//
	//   "name": "foo",
	//
	//   "metadata": {
	//
	//     "instance": 2,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4
	//
	//     },
	//
	//   "containers": [
	//
	//     {
	//
	//       "image": "****",
	//
	//       "script": "***	- --listen=0.0.0.0 --server_port=8000 --headless",
	//
	//       "port": 8000
	//
	//     }
	//
	//   ],
	//
	//   "storage": [
	//
	//     {
	//
	//       "oss": {
	//
	//         "path": "oss://examplebuket/data111/",
	//
	//         "readOnly": false
	//
	//       },
	//
	//       "properties": {
	//
	//         "resource_type": "model"
	//
	//       },
	//
	//       "mount_path": "/data"
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// AI-Web application deployment by using an image:
	//
	// {
	//
	//   "name": "foo",
	//
	//   "metadata": {
	//
	//     "instance": 1,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4,
	//
	//     "enable_webservice": true
	//
	//   },
	//
	//   "containers": [
	//
	//     {
	//
	//       "image": "****",
	//
	//       "script": "***	- --listen=0.0.0.0 --server_port=8000 --headless",
	//
	//       "port": 8000
	//
	//     }
	//
	//   ],
	//
	//   "storage": [
	//
	//     {
	//
	//       "oss": {
	//
	//         "path": "oss://examplebucket/data111/",
	//
	//         "readOnly": false
	//
	//       },
	//
	//       "properties": {
	//
	//       "resource_type": "model"
	//
	//       },
	//
	//       "mount_path": "/data"
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// Service deployment by using models and processors:
	//
	// {
	//
	//   "metadata": {
	//
	//     "instance": 1,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4
	//
	//   },
	//
	//   "name": "foo",
	//
	//   "model_config": {},
	//
	//   "processor_type": "python",
	//
	//   "processor_path": "oss://****",
	//
	//   "processor_entry": "a.py",
	//
	//   "model_path": "oss://****"
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetDevelop(v string) *CreateServiceRequest {
	s.Develop = &v
	return s
}

func (s *CreateServiceRequest) SetLabels(v map[string]*string) *CreateServiceRequest {
	s.Labels = v
	return s
}

func (s *CreateServiceRequest) SetWorkspaceId(v string) *CreateServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateServiceRequest) SetBody(v string) *CreateServiceRequest {
	s.Body = &v
	return s
}

type CreateServiceShrinkRequest struct {
	// Specifies whether to enter development mode.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Develop *string `json:"Develop,omitempty" xml:"Develop,omitempty"`
	// The custom label.
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The request body. For more information about the key request parameters, see **Table 1. Request body parameters*	- and **Table 2. Metadata parameters**. For more information about all related parameters, see [Parameters of model services](https://help.aliyun.com/document_detail/450525.html).
	//
	// example:
	//
	// Service deployment by using an image:
	//
	// {
	//
	//   "name": "foo",
	//
	//   "metadata": {
	//
	//     "instance": 2,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4
	//
	//     },
	//
	//   "containers": [
	//
	//     {
	//
	//       "image": "****",
	//
	//       "script": "***	- --listen=0.0.0.0 --server_port=8000 --headless",
	//
	//       "port": 8000
	//
	//     }
	//
	//   ],
	//
	//   "storage": [
	//
	//     {
	//
	//       "oss": {
	//
	//         "path": "oss://examplebuket/data111/",
	//
	//         "readOnly": false
	//
	//       },
	//
	//       "properties": {
	//
	//         "resource_type": "model"
	//
	//       },
	//
	//       "mount_path": "/data"
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// AI-Web application deployment by using an image:
	//
	// {
	//
	//   "name": "foo",
	//
	//   "metadata": {
	//
	//     "instance": 1,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4,
	//
	//     "enable_webservice": true
	//
	//   },
	//
	//   "containers": [
	//
	//     {
	//
	//       "image": "****",
	//
	//       "script": "***	- --listen=0.0.0.0 --server_port=8000 --headless",
	//
	//       "port": 8000
	//
	//     }
	//
	//   ],
	//
	//   "storage": [
	//
	//     {
	//
	//       "oss": {
	//
	//         "path": "oss://examplebucket/data111/",
	//
	//         "readOnly": false
	//
	//       },
	//
	//       "properties": {
	//
	//       "resource_type": "model"
	//
	//       },
	//
	//       "mount_path": "/data"
	//
	//     }
	//
	//   ]
	//
	// }
	//
	// Service deployment by using models and processors:
	//
	// {
	//
	//   "metadata": {
	//
	//     "instance": 1,
	//
	//     "memory": 7000,
	//
	//     "cpu": 4
	//
	//   },
	//
	//   "name": "foo",
	//
	//   "model_config": {},
	//
	//   "processor_type": "python",
	//
	//   "processor_path": "oss://****",
	//
	//   "processor_entry": "a.py",
	//
	//   "model_path": "oss://****"
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceShrinkRequest) SetDevelop(v string) *CreateServiceShrinkRequest {
	s.Develop = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetLabelsShrink(v string) *CreateServiceShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetWorkspaceId(v string) *CreateServiceShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateServiceShrinkRequest) SetBody(v string) *CreateServiceShrinkRequest {
	s.Body = &v
	return s
}

type CreateServiceResponseBody struct {
	// The public endpoint of the created service.
	//
	// example:
	//
	// http://pai-eas.vpc.cn-shanghai.****
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The internal endpoint of the created service.
	//
	// example:
	//
	// http://pai-eas.cn-shanghai.****
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// The region ID of the created service.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created service.
	//
	// example:
	//
	// eas-m-aaxxxddf
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the created service.
	//
	// example:
	//
	// yourname
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service state.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetInternetEndpoint(v string) *CreateServiceResponseBody {
	s.InternetEndpoint = &v
	return s
}

func (s *CreateServiceResponseBody) SetIntranetEndpoint(v string) *CreateServiceResponseBody {
	s.IntranetEndpoint = &v
	return s
}

func (s *CreateServiceResponseBody) SetRegion(v string) *CreateServiceResponseBody {
	s.Region = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceName(v string) *CreateServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceResponseBody) SetStatus(v string) *CreateServiceResponseBody {
	s.Status = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateServiceAutoScalerRequest struct {
	// The Autoscaler operation.
	Behavior *CreateServiceAutoScalerRequestBehavior `json:"behavior,omitempty" xml:"behavior,omitempty" type:"Struct"`
	// The maximum number of instances in the service. The value of max must be greater than the value of min.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	Max *int32 `json:"max,omitempty" xml:"max,omitempty"`
	// The minimum number of instances in the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// The service for which the metric is specified. If you do not set this parameter, the current service is specified by default.
	//
	// This parameter is required.
	ScaleStrategies []*CreateServiceAutoScalerRequestScaleStrategies `json:"scaleStrategies,omitempty" xml:"scaleStrategies,omitempty" type:"Repeated"`
}

func (s CreateServiceAutoScalerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequest) SetBehavior(v *CreateServiceAutoScalerRequestBehavior) *CreateServiceAutoScalerRequest {
	s.Behavior = v
	return s
}

func (s *CreateServiceAutoScalerRequest) SetMax(v int32) *CreateServiceAutoScalerRequest {
	s.Max = &v
	return s
}

func (s *CreateServiceAutoScalerRequest) SetMin(v int32) *CreateServiceAutoScalerRequest {
	s.Min = &v
	return s
}

func (s *CreateServiceAutoScalerRequest) SetScaleStrategies(v []*CreateServiceAutoScalerRequestScaleStrategies) *CreateServiceAutoScalerRequest {
	s.ScaleStrategies = v
	return s
}

type CreateServiceAutoScalerRequestBehavior struct {
	// The operation that reduces the number of instances to 0.
	OnZero *CreateServiceAutoScalerRequestBehaviorOnZero `json:"onZero,omitempty" xml:"onZero,omitempty" type:"Struct"`
	// The scale-in operation.
	ScaleDown *CreateServiceAutoScalerRequestBehaviorScaleDown `json:"scaleDown,omitempty" xml:"scaleDown,omitempty" type:"Struct"`
	// The scale-out operation.
	ScaleUp *CreateServiceAutoScalerRequestBehaviorScaleUp `json:"scaleUp,omitempty" xml:"scaleUp,omitempty" type:"Struct"`
}

func (s CreateServiceAutoScalerRequestBehavior) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerRequestBehavior) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestBehavior) SetOnZero(v *CreateServiceAutoScalerRequestBehaviorOnZero) *CreateServiceAutoScalerRequestBehavior {
	s.OnZero = v
	return s
}

func (s *CreateServiceAutoScalerRequestBehavior) SetScaleDown(v *CreateServiceAutoScalerRequestBehaviorScaleDown) *CreateServiceAutoScalerRequestBehavior {
	s.ScaleDown = v
	return s
}

func (s *CreateServiceAutoScalerRequestBehavior) SetScaleUp(v *CreateServiceAutoScalerRequestBehaviorScaleUp) *CreateServiceAutoScalerRequestBehavior {
	s.ScaleUp = v
	return s
}

type CreateServiceAutoScalerRequestBehaviorOnZero struct {
	// The time window that is required before the number of instances is reduced to 0. The number of instances can be reduced to 0 only if no request is available or no traffic exists in the specified time window. Default value: 600.
	//
	// example:
	//
	// 600
	ScaleDownGracePeriodSeconds *int32 `json:"scaleDownGracePeriodSeconds,omitempty" xml:"scaleDownGracePeriodSeconds,omitempty"`
	// The number of instances that you want to create at a time if the number of instances is 0. Default value: 1.
	//
	// example:
	//
	// 1
	ScaleUpActivationReplicas *int32 `json:"scaleUpActivationReplicas,omitempty" xml:"scaleUpActivationReplicas,omitempty"`
}

func (s CreateServiceAutoScalerRequestBehaviorOnZero) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerRequestBehaviorOnZero) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestBehaviorOnZero) SetScaleDownGracePeriodSeconds(v int32) *CreateServiceAutoScalerRequestBehaviorOnZero {
	s.ScaleDownGracePeriodSeconds = &v
	return s
}

func (s *CreateServiceAutoScalerRequestBehaviorOnZero) SetScaleUpActivationReplicas(v int32) *CreateServiceAutoScalerRequestBehaviorOnZero {
	s.ScaleUpActivationReplicas = &v
	return s
}

type CreateServiceAutoScalerRequestBehaviorScaleDown struct {
	// The time window that is required before the scale-in operation is performed. The scale-in operation can be performed only if the specified metric drops below the specified threshold in the specified time window. Default value: 300.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty" xml:"stabilizationWindowSeconds,omitempty"`
}

func (s CreateServiceAutoScalerRequestBehaviorScaleDown) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerRequestBehaviorScaleDown) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestBehaviorScaleDown) SetStabilizationWindowSeconds(v int32) *CreateServiceAutoScalerRequestBehaviorScaleDown {
	s.StabilizationWindowSeconds = &v
	return s
}

type CreateServiceAutoScalerRequestBehaviorScaleUp struct {
	// The time window that is required before the scale-out operation is performed. The scale-out operation can be performed only if the specified metric exceeds the specified threshold in the specified time window. Default value: 0.
	//
	// example:
	//
	// 0
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty" xml:"stabilizationWindowSeconds,omitempty"`
}

func (s CreateServiceAutoScalerRequestBehaviorScaleUp) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerRequestBehaviorScaleUp) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestBehaviorScaleUp) SetStabilizationWindowSeconds(v int32) *CreateServiceAutoScalerRequestBehaviorScaleUp {
	s.StabilizationWindowSeconds = &v
	return s
}

type CreateServiceAutoScalerRequestScaleStrategies struct {
	// The name of the metric for triggering auto scaling. Valid values:
	//
	// 	- qps: the queries per second (qps) for an individual instance.
	//
	// 	- cpu: the cpu utilization.
	//
	// 	- gpu[util]: gpu utilization.
	//
	// This parameter is required.
	//
	// example:
	//
	// qps
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The service for which the metric is specified. If you do not set this parameter, the current service is specified by default.
	//
	// example:
	//
	// demo_svc
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// The threshold of the metric that triggers auto scaling.
	//
	// 	- If you set metricName to qps, scale-out is triggered when the average qps for a single instance is greater than this threshold.
	//
	// 	- If you set metricName to cpu, scale-out is triggered when the average cpu utilization for a single instance is greater than this threshold.
	//
	// 	- If you set metricName to gpu, scale-out is triggered when the average cpu utilization for a single instance is greater than this threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s CreateServiceAutoScalerRequestScaleStrategies) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerRequestScaleStrategies) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) SetMetricName(v string) *CreateServiceAutoScalerRequestScaleStrategies {
	s.MetricName = &v
	return s
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) SetService(v string) *CreateServiceAutoScalerRequestScaleStrategies {
	s.Service = &v
	return s
}

func (s *CreateServiceAutoScalerRequestScaleStrategies) SetThreshold(v float32) *CreateServiceAutoScalerRequestScaleStrategies {
	s.Threshold = &v
	return s
}

type CreateServiceAutoScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to auto scale service [foo]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceAutoScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerResponseBody) SetMessage(v string) *CreateServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceAutoScalerResponseBody) SetRequestId(v string) *CreateServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceAutoScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerResponse) SetHeaders(v map[string]*string) *CreateServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceAutoScalerResponse) SetStatusCode(v int32) *CreateServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceAutoScalerResponse) SetBody(v *CreateServiceAutoScalerResponseBody) *CreateServiceAutoScalerResponse {
	s.Body = v
	return s
}

type CreateServiceCronScalerRequest struct {
	// The points in time that are excluded when you schedule a CronHPA job. The points in time must be specified by using a cron expression.
	ExcludeDates []*string `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	// The description of the CronHPA job.
	//
	// This parameter is required.
	ScaleJobs []*CreateServiceCronScalerRequestScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
}

func (s CreateServiceCronScalerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceCronScalerRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerRequest) SetExcludeDates(v []*string) *CreateServiceCronScalerRequest {
	s.ExcludeDates = v
	return s
}

func (s *CreateServiceCronScalerRequest) SetScaleJobs(v []*CreateServiceCronScalerRequestScaleJobs) *CreateServiceCronScalerRequest {
	s.ScaleJobs = v
	return s
}

type CreateServiceCronScalerRequestScaleJobs struct {
	// The name of the CronHPA job.
	//
	// example:
	//
	// scale-job-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cron expression that is used to configure the execution time of the CronHPA job. For more information about how to configure cron expressions, see **Description of special characters*	- in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 18 	- 	- 	- *
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The number of instances that you want to configure for the CronHPA job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TargetSize *int32 `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s CreateServiceCronScalerRequestScaleJobs) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceCronScalerRequestScaleJobs) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetName(v string) *CreateServiceCronScalerRequestScaleJobs {
	s.Name = &v
	return s
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetSchedule(v string) *CreateServiceCronScalerRequestScaleJobs {
	s.Schedule = &v
	return s
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetTargetSize(v int32) *CreateServiceCronScalerRequestScaleJobs {
	s.TargetSize = &v
	return s
}

type CreateServiceCronScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Create cron scaler for service [foo] successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceCronScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerResponseBody) SetMessage(v string) *CreateServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceCronScalerResponseBody) SetRequestId(v string) *CreateServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceCronScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerResponse) SetHeaders(v map[string]*string) *CreateServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceCronScalerResponse) SetStatusCode(v int32) *CreateServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceCronScalerResponse) SetBody(v *CreateServiceCronScalerResponseBody) *CreateServiceCronScalerResponse {
	s.Body = v
	return s
}

type CreateServiceMirrorRequest struct {
	// The percentage of the traffic that is mirrored to the destination service. Valid values: 0 to 100.
	//
	// example:
	//
	// 30
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The instances.
	Target []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s CreateServiceMirrorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMirrorRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorRequest) SetRatio(v int32) *CreateServiceMirrorRequest {
	s.Ratio = &v
	return s
}

func (s *CreateServiceMirrorRequest) SetTarget(v []*string) *CreateServiceMirrorRequest {
	s.Target = v
	return s
}

type CreateServiceMirrorResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Traffic mirroring is ON for service [foo] in region [cn-shanghia], ratio [70%]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceMirrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorResponseBody) SetMessage(v string) *CreateServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceMirrorResponseBody) SetRequestId(v string) *CreateServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceMirrorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorResponse) SetHeaders(v map[string]*string) *CreateServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceMirrorResponse) SetStatusCode(v int32) *CreateServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceMirrorResponse) SetBody(v *CreateServiceMirrorResponseBody) *CreateServiceMirrorResponse {
	s.Body = v
	return s
}

type DeleteAclPolicyRequest struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway.
	AclPolicyList []*DeleteAclPolicyRequestAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteAclPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyRequest) SetAclPolicyList(v []*DeleteAclPolicyRequestAclPolicyList) *DeleteAclPolicyRequest {
	s.AclPolicyList = v
	return s
}

func (s *DeleteAclPolicyRequest) SetVpcId(v string) *DeleteAclPolicyRequest {
	s.VpcId = &v
	return s
}

type DeleteAclPolicyRequestAclPolicyList struct {
	// The comment on the IP CIDR block in the VPC that can access the private gateway.
	//
	// example:
	//
	// default
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP CIDR block in the VPC that can access the private gateway.
	//
	// example:
	//
	// 10.23.XX.XX/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s DeleteAclPolicyRequestAclPolicyList) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclPolicyRequestAclPolicyList) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyRequestAclPolicyList) SetComment(v string) *DeleteAclPolicyRequestAclPolicyList {
	s.Comment = &v
	return s
}

func (s *DeleteAclPolicyRequestAclPolicyList) SetEntry(v string) *DeleteAclPolicyRequestAclPolicyList {
	s.Entry = &v
	return s
}

type DeleteAclPolicyShrinkRequest struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway.
	AclPolicyListShrink *string `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty"`
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteAclPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyShrinkRequest) SetAclPolicyListShrink(v string) *DeleteAclPolicyShrinkRequest {
	s.AclPolicyListShrink = &v
	return s
}

func (s *DeleteAclPolicyShrinkRequest) SetVpcId(v string) *DeleteAclPolicyShrinkRequest {
	s.VpcId = &v
	return s
}

type DeleteAclPolicyResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully delete acl policy for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAclPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyResponseBody) SetGatewayId(v string) *DeleteAclPolicyResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DeleteAclPolicyResponseBody) SetMessage(v string) *DeleteAclPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAclPolicyResponseBody) SetRequestId(v string) *DeleteAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAclPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAclPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclPolicyResponse) SetHeaders(v map[string]*string) *DeleteAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclPolicyResponse) SetStatusCode(v int32) *DeleteAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAclPolicyResponse) SetBody(v *DeleteAclPolicyResponseBody) *DeleteAclPolicyResponse {
	s.Body = v
	return s
}

type DeleteBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-test-service-234c] is Deleting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBenchmarkTaskResponseBody) SetMessage(v string) *DeleteBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBenchmarkTaskResponseBody) SetRequestId(v string) *DeleteBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteBenchmarkTaskResponse) SetHeaders(v map[string]*string) *DeleteBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteBenchmarkTaskResponse) SetStatusCode(v int32) *DeleteBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBenchmarkTaskResponse) SetBody(v *DeleteBenchmarkTaskResponseBody) *DeleteBenchmarkTaskResponse {
	s.Body = v
	return s
}

type DeleteGatewayResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Gateway is deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) SetGatewayId(v string) *DeleteGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetMessage(v string) *DeleteGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponse) SetHeaders(v map[string]*string) *DeleteGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayResponse) SetStatusCode(v int32) *DeleteGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayResponse) SetBody(v *DeleteGatewayResponseBody) *DeleteGatewayResponse {
	s.Body = v
	return s
}

type DeleteGatewayIntranetLinkedVpcRequest struct {
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-8vbqn2at0kljjxxxx****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcRequest) SetVSwitchId(v string) *DeleteGatewayIntranetLinkedVpcRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcRequest) SetVpcId(v string) *DeleteGatewayIntranetLinkedVpcRequest {
	s.VpcId = &v
	return s
}

type DeleteGatewayIntranetLinkedVpcResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully delete intranet linked vpc for gateway
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) SetGatewayId(v string) *DeleteGatewayIntranetLinkedVpcResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) SetMessage(v string) *DeleteGatewayIntranetLinkedVpcResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponseBody) SetRequestId(v string) *DeleteGatewayIntranetLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGatewayIntranetLinkedVpcResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayIntranetLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) SetHeaders(v map[string]*string) *DeleteGatewayIntranetLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) SetStatusCode(v int32) *DeleteGatewayIntranetLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcResponse) SetBody(v *DeleteGatewayIntranetLinkedVpcResponseBody) *DeleteGatewayIntranetLinkedVpcResponse {
	s.Body = v
	return s
}

type DeleteResourceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Resource [eas-r-asdasdasd] is deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82-9624-EC2B1779848E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) SetMessage(v string) *DeleteResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetStatusCode(v int32) *DeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

type DeleteResourceDLinkResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Network interfaces are deleting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceDLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceDLinkResponseBody) SetMessage(v string) *DeleteResourceDLinkResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceDLinkResponseBody) SetRequestId(v string) *DeleteResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceDLinkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceDLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceDLinkResponse) SetHeaders(v map[string]*string) *DeleteResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceDLinkResponse) SetStatusCode(v int32) *DeleteResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceDLinkResponse) SetBody(v *DeleteResourceDLinkResponseBody) *DeleteResourceDLinkResponse {
	s.Body = v
	return s
}

type DeleteResourceInstancesRequest struct {
	// Specifies whether to delete all the instances that fail to be created. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllFailed *bool `json:"AllFailed,omitempty" xml:"AllFailed,omitempty"`
	// The instances. Separate multiple instances with commas (,), such as `instanceId1,instanceId2`. For more information about how to query the instances, see [ListResourceInstances](https://help.aliyun.com/document_detail/412129.html).
	//
	// example:
	//
	// eas-i-xxxxxxx,eas-i-xxxxxxx
	InstanceList *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
}

func (s DeleteResourceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesRequest) SetAllFailed(v bool) *DeleteResourceInstancesRequest {
	s.AllFailed = &v
	return s
}

func (s *DeleteResourceInstancesRequest) SetInstanceList(v string) *DeleteResourceInstancesRequest {
	s.InstanceList = &v
	return s
}

type DeleteResourceInstancesResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Instances eas-i-011227132046,eas-i-011227132046 are deleting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesResponseBody) SetMessage(v string) *DeleteResourceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceInstancesResponseBody) SetRequestId(v string) *DeleteResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesResponse) SetHeaders(v map[string]*string) *DeleteResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceInstancesResponse) SetStatusCode(v int32) *DeleteResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceInstancesResponse) SetBody(v *DeleteResourceInstancesResponseBody) *DeleteResourceInstancesResponse {
	s.Body = v
	return s
}

type DeleteResourceLogResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Log service for resource [eas-r-asdasdasd] is deleting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceLogResponseBody) SetMessage(v string) *DeleteResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceLogResponseBody) SetRequestId(v string) *DeleteResourceLogResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceLogResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceLogResponse) SetHeaders(v map[string]*string) *DeleteResourceLogResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceLogResponse) SetStatusCode(v int32) *DeleteResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceLogResponse) SetBody(v *DeleteResourceLogResponseBody) *DeleteResourceLogResponse {
	s.Body = v
	return s
}

type DeleteServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Service [foo] in region [cn-shanghai] is terminating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetMessage(v string) *DeleteServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DeleteServiceAutoScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Autoscaler for service [foo] deleted successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceAutoScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceAutoScalerResponseBody) SetMessage(v string) *DeleteServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceAutoScalerResponseBody) SetRequestId(v string) *DeleteServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceAutoScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceAutoScalerResponse) SetHeaders(v map[string]*string) *DeleteServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceAutoScalerResponse) SetStatusCode(v int32) *DeleteServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceAutoScalerResponse) SetBody(v *DeleteServiceAutoScalerResponseBody) *DeleteServiceAutoScalerResponse {
	s.Body = v
	return s
}

type DeleteServiceCronScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Cronscaler for service [foo] deleted successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceCronScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceCronScalerResponseBody) SetMessage(v string) *DeleteServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceCronScalerResponseBody) SetRequestId(v string) *DeleteServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceCronScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceCronScalerResponse) SetHeaders(v map[string]*string) *DeleteServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceCronScalerResponse) SetStatusCode(v int32) *DeleteServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceCronScalerResponse) SetBody(v *DeleteServiceCronScalerResponseBody) *DeleteServiceCronScalerResponse {
	s.Body = v
	return s
}

type DeleteServiceInstancesRequest struct {
	// The name of the container whose process needs to be restarted. This parameter takes effect only if the SoftRestart parameter is set to true.
	//
	// example:
	//
	// worker0
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The instances that you want to restart. Separate multiple instance names with commas (,). For more information about how to query the instance name, see [ListServiceInstances](https://help.aliyun.com/document_detail/412108.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// foo-rdsbxxxx,foo-rdsaxxxx
	InstanceList *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
	// Specifies whether to restart only the container process without recreating the instance. Default value: false. Valid values: true and false.
	//
	// example:
	//
	// true
	SoftRestart *bool `json:"SoftRestart,omitempty" xml:"SoftRestart,omitempty"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) SetContainer(v string) *DeleteServiceInstancesRequest {
	s.Container = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetInstanceList(v string) *DeleteServiceInstancesRequest {
	s.InstanceList = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetSoftRestart(v bool) *DeleteServiceInstancesRequest {
	s.SoftRestart = &v
	return s
}

type DeleteServiceInstancesResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Instance(s) [foo-rdsb,foo-rdsa]  for service [foo] in region [cn-shanghai] was deleted successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) SetMessage(v string) *DeleteServiceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponse) SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceInstancesResponse) SetStatusCode(v int32) *DeleteServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

type DeleteServiceLabelRequest struct {
	// The service tags that you want to delete.
	//
	// This parameter is required.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s DeleteServiceLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceLabelRequest) SetKeys(v []*string) *DeleteServiceLabelRequest {
	s.Keys = v
	return s
}

type DeleteServiceLabelShrinkRequest struct {
	// The service tags that you want to delete.
	//
	// This parameter is required.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s DeleteServiceLabelShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceLabelShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceLabelShrinkRequest) SetKeysShrink(v string) *DeleteServiceLabelShrinkRequest {
	s.KeysShrink = &v
	return s
}

type DeleteServiceLabelResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to delete service [service_from_zxxx] labels.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceLabelResponseBody) SetMessage(v string) *DeleteServiceLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceLabelResponseBody) SetRequestId(v string) *DeleteServiceLabelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceLabelResponse) SetHeaders(v map[string]*string) *DeleteServiceLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceLabelResponse) SetStatusCode(v int32) *DeleteServiceLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceLabelResponse) SetBody(v *DeleteServiceLabelResponseBody) *DeleteServiceLabelResponse {
	s.Body = v
	return s
}

type DeleteServiceMirrorResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Traffic mirroring is OFF for service [foo] in region [cn-shanghia], ratio [70%]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceMirrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceMirrorResponseBody) SetMessage(v string) *DeleteServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceMirrorResponseBody) SetRequestId(v string) *DeleteServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceMirrorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceMirrorResponse) SetHeaders(v map[string]*string) *DeleteServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceMirrorResponse) SetStatusCode(v int32) *DeleteServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceMirrorResponse) SetBody(v *DeleteServiceMirrorResponseBody) *DeleteServiceMirrorResponse {
	s.Body = v
	return s
}

type DescribeBenchmarkTaskResponseBody struct {
	// The number of instances that you can test.
	//
	// example:
	//
	// 4
	AvailableAgent *int64 `json:"AvailableAgent,omitempty" xml:"AvailableAgent,omitempty"`
	// The ID of the operation caller.
	//
	// example:
	//
	// 1640133467****
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The number of instances that you want to test.
	//
	// example:
	//
	// 4
	DesiredAgent *int64 `json:"DesiredAgent,omitempty" xml:"DesiredAgent,omitempty"`
	// The endpoint of the service gateway.
	//
	// example:
	//
	// 192342311234.pai-eas.cn-chengdu.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Benchmar task is Running
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the Alibaba Cloud account that is used to call the operation.
	//
	// example:
	//
	// 1029728669****
	ParentUid *string `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	// The event or reason that causes the current state of the stress testing task.
	//
	// example:
	//
	// RUNNING
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the service that you want to test.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the stress testing task.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DeleteFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Error
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Updating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CreateFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the stress testing task.
	//
	// example:
	//
	// eas-b-gv4y86u****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the stress testing task.
	//
	// example:
	//
	// benchmark-larec-test-ae70
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The token for authentication when a stress testing task is created.
	//
	// example:
	//
	// 6062787a-9301****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskResponseBody) SetAvailableAgent(v int64) *DescribeBenchmarkTaskResponseBody {
	s.AvailableAgent = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetCallerUid(v string) *DescribeBenchmarkTaskResponseBody {
	s.CallerUid = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetDesiredAgent(v int64) *DescribeBenchmarkTaskResponseBody {
	s.DesiredAgent = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetEndpoint(v string) *DescribeBenchmarkTaskResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetMessage(v string) *DescribeBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetParentUid(v string) *DescribeBenchmarkTaskResponseBody {
	s.ParentUid = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetReason(v string) *DescribeBenchmarkTaskResponseBody {
	s.Reason = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetRequestId(v string) *DescribeBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetServiceName(v string) *DescribeBenchmarkTaskResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetStatus(v string) *DescribeBenchmarkTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetTaskId(v string) *DescribeBenchmarkTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetTaskName(v string) *DescribeBenchmarkTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetToken(v string) *DescribeBenchmarkTaskResponseBody {
	s.Token = &v
	return s
}

type DescribeBenchmarkTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskResponse) SetHeaders(v map[string]*string) *DescribeBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeBenchmarkTaskResponse) SetStatusCode(v int32) *DescribeBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBenchmarkTaskResponse) SetBody(v *DescribeBenchmarkTaskResponseBody) *DescribeBenchmarkTaskResponse {
	s.Body = v
	return s
}

type DescribeBenchmarkTaskReportRequest struct {
	// The report type of the stress testing task. Valid values: RAW and Report.
	//
	// example:
	//
	// report
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s DescribeBenchmarkTaskReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskReportRequest) SetReportType(v string) *DescribeBenchmarkTaskReportRequest {
	s.ReportType = &v
	return s
}

type DescribeBenchmarkTaskReportResponseBody struct {
	// If the value of ReportType is set to RAW, the details about the stress testing report are returned.
	//
	// example:
	//
	// {
	//
	//     "TimestampList": ["int64"],
	//
	//     "QPSList": ["float32"],
	//
	//     "RTList": [
	//
	//       {
	//
	//         "AVG": "float32",
	//
	//         "TP100": "float32",
	//
	//         "TP99": "float32",
	//
	//         "TP90": "float32",
	//
	//         "TP50": "float32",
	//
	//         "TP10": "float32"
	//
	//       }
	//
	//     ],
	//
	//     "TrafficList": [
	//
	//       {
	//
	//         "Send": "float64",
	//
	//         "Receive": "float64"
	//
	//       }
	//
	//     ],
	//
	//     "StatusCode": {
	//
	//       "200": "uint64",
	//
	//       "450": "uint64",
	//
	//       "500": "uint64"
	//
	//     },
	//
	//     "Count": "uint64",
	//
	//     "Total": "float64",
	//
	//     "MinRT": "float32",
	//
	//     "AvgRT": "float32",
	//
	//     "MaxRT": "float32",
	//
	//     "QPS": "float32",
	//
	//     "TotalSend": "float64",
	//
	//     "TotalReceive": "float64",
	//
	//     "RTDistribution": [
	//
	//       {
	//
	//         "Latency": "float32",
	//
	//         "Percentage": "int"
	//
	//       }
	//
	//     ],
	//
	//     "RTHistogram": [
	//
	//       {
	//
	//         "Count": "int",
	//
	//         "Mark": "float32",
	//
	//         "Frequency": "float32"
	//
	//       }
	//
	//     ]
	//
	//   }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// If the value of ReportType is set to Report, the URL of the stress testing report is returned.
	//
	// example:
	//
	// http://eas-benchmark.oss-cn-chengdu.aliyuncs.com/summary/benchmark-larec-test-015d-10007.html
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBenchmarkTaskReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskReportResponseBody) SetData(v interface{}) *DescribeBenchmarkTaskReportResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBenchmarkTaskReportResponseBody) SetReportUrl(v string) *DescribeBenchmarkTaskReportResponseBody {
	s.ReportUrl = &v
	return s
}

func (s *DescribeBenchmarkTaskReportResponseBody) SetRequestId(v string) *DescribeBenchmarkTaskReportResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBenchmarkTaskReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBenchmarkTaskReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBenchmarkTaskReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskReportResponse) SetHeaders(v map[string]*string) *DescribeBenchmarkTaskReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeBenchmarkTaskReportResponse) SetStatusCode(v int32) *DescribeBenchmarkTaskReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBenchmarkTaskReportResponse) SetBody(v *DescribeBenchmarkTaskReportResponseBody) *DescribeBenchmarkTaskReportResponse {
	s.Body = v
	return s
}

type DescribeGatewayResponseBody struct {
	// The time when the private gateway was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-19T14:19:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the self-managed cluster.
	//
	// example:
	//
	// c935eadf284c14c2da57a2a13ad6******
	ExternalClusterId *string `json:"ExternalClusterId,omitempty" xml:"ExternalClusterId,omitempty"`
	// The ID of the private gateway.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The alias of the private gateway.
	//
	// example:
	//
	// mygateway1
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The instance type used by the private gateway.
	//
	// Valid values:
	//
	// 	- 8c16g
	//
	// 	- 4c8g
	//
	// 	- 2c4g
	//
	// 	- 16c32g
	//
	// example:
	//
	// ecs.c6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******-1801786532******.cn-hangzhou.pai-eas.aliyuncs.com
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// Indicates whether Internet access is enabled.
	//
	// example:
	//
	// true
	InternetEnabled *bool `json:"InternetEnabled,omitempty" xml:"InternetEnabled,omitempty"`
	// Indicates whether Internet access is enabled.
	//
	// Valid values:
	//
	// 	- Creating: Internet access is being enabled.
	//
	// 	- Failed: Internet access failed to be enabled or deleted.
	//
	// 	- Running: Internet access is running.
	//
	// 	- Deleted: Internet access is deleted.
	//
	// 	- Deleting: Internet access is being deleted.
	//
	// example:
	//
	// Running
	InternetStatus *string `json:"InternetStatus,omitempty" xml:"InternetStatus,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******-1801786532******-vpc.cn-hangzhou.pai-eas.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// Indicates whether it is the default private gateway.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The number of nodes in the private gateway.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the private gateway.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Deleted
	//
	// 	- Deleting
	//
	// 	- Waiting
	//
	// example:
	//
	// PrivateGatewayRunning
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the private gateway was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-02-24T11:52:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayResponseBody) SetCreateTime(v string) *DescribeGatewayResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetExternalClusterId(v string) *DescribeGatewayResponseBody {
	s.ExternalClusterId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayId(v string) *DescribeGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetGatewayName(v string) *DescribeGatewayResponseBody {
	s.GatewayName = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInstanceType(v string) *DescribeGatewayResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInternetDomain(v string) *DescribeGatewayResponseBody {
	s.InternetDomain = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInternetEnabled(v bool) *DescribeGatewayResponseBody {
	s.InternetEnabled = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetInternetStatus(v string) *DescribeGatewayResponseBody {
	s.InternetStatus = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetIntranetDomain(v string) *DescribeGatewayResponseBody {
	s.IntranetDomain = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetIsDefault(v bool) *DescribeGatewayResponseBody {
	s.IsDefault = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetReplicas(v int32) *DescribeGatewayResponseBody {
	s.Replicas = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetRequestId(v string) *DescribeGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetStatus(v string) *DescribeGatewayResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGatewayResponseBody) SetUpdateTime(v string) *DescribeGatewayResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeGatewayResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGatewayResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayResponse) SetHeaders(v map[string]*string) *DescribeGatewayResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayResponse) SetStatusCode(v int32) *DescribeGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayResponse) SetBody(v *DescribeGatewayResponseBody) *DescribeGatewayResponse {
	s.Body = v
	return s
}

type DescribeGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Group             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupResponse) SetHeaders(v map[string]*string) *DescribeGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupResponse) SetStatusCode(v int32) *DescribeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupResponse) SetBody(v *Group) *DescribeGroupResponse {
	s.Body = v
	return s
}

type DescribeResourceResponseBody struct {
	// The ID of the cluster to which the resource group belongs.
	//
	// example:
	//
	// cn-beijing
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The total number of CPU cores.
	//
	// example:
	//
	// 16
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The time when the resource group was created.
	//
	// example:
	//
	// 2020-05-19T14:19:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The additional information, such as the connection status of a virtual private cloud (VPC) and the log status of Log Service.
	//
	// example:
	//
	// {"vswitch_id":"vsw-bp17uo6xebcusy****","gpu_share":true,"aux_vswitch_id_list":["vsw-bp13b3pvjap3vxn****","vsw-bp1nls8o5hk8mt8*****"],"security_group_id":"sg-bp1j1z7297hcink*****","vpc_id":"vpc-bp1kjr3rfyhx01*****","destination_cidr":"172.16.0.12/28","role_arn":"acs:ram::1157703270*****:role/AliyunServiceRoleForPaiEas","sls_project":"","sls_logstore":"","sls_status":"ResourceReady","sls_message":"","update_time":""}
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// The total number of GPUs.
	//
	// example:
	//
	// 1
	GpuCount *int32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The total number of instances in the resource group.
	//
	// example:
	//
	// 4
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Resource is ready
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the resource group owner.
	//
	// example:
	//
	// 14401087478****
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The total number of pay-as-you-go instances in the resource group.
	//
	// example:
	//
	// 3
	PostPaidInstanceCount *int32 `json:"PostPaidInstanceCount,omitempty" xml:"PostPaidInstanceCount,omitempty"`
	// The total number of subscription instances in the resource group.
	//
	// example:
	//
	// 1
	PrePaidInstanceCount *int32 `json:"PrePaidInstanceCount,omitempty" xml:"PrePaidInstanceCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 902976F2-6FAF-5404-8A4D-6CC223***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the Elastic Algorithm Service (EAS) resource.
	//
	// example:
	//
	// eas-r-glkfpsxuw57x1h*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the EAS resource.
	//
	// example:
	//
	// my-resouce****
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- Dedicated: the dedicated resource group.
	//
	// 	- SelfManaged: the self-managed resource group.
	//
	// example:
	//
	// Dedicated
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The state of the resource group.
	//
	// example:
	//
	// ResourceReady
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the resource group was last updated.
	//
	// example:
	//
	// 2021-02-24T11:52:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceResponseBody) SetClusterId(v string) *DescribeResourceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetCpuCount(v int32) *DescribeResourceResponseBody {
	s.CpuCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetCreateTime(v string) *DescribeResourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourceResponseBody) SetExtraData(v string) *DescribeResourceResponseBody {
	s.ExtraData = &v
	return s
}

func (s *DescribeResourceResponseBody) SetGpuCount(v int32) *DescribeResourceResponseBody {
	s.GpuCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetInstanceCount(v int32) *DescribeResourceResponseBody {
	s.InstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetMessage(v string) *DescribeResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceResponseBody) SetOwnerUid(v string) *DescribeResourceResponseBody {
	s.OwnerUid = &v
	return s
}

func (s *DescribeResourceResponseBody) SetPostPaidInstanceCount(v int32) *DescribeResourceResponseBody {
	s.PostPaidInstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetPrePaidInstanceCount(v int32) *DescribeResourceResponseBody {
	s.PrePaidInstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetRequestId(v string) *DescribeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetResourceId(v string) *DescribeResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetResourceName(v string) *DescribeResourceResponseBody {
	s.ResourceName = &v
	return s
}

func (s *DescribeResourceResponseBody) SetResourceType(v string) *DescribeResourceResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceResponseBody) SetStatus(v string) *DescribeResourceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeResourceResponseBody) SetUpdateTime(v string) *DescribeResourceResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceResponse) SetHeaders(v map[string]*string) *DescribeResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceResponse) SetStatusCode(v int32) *DescribeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceResponse) SetBody(v *DescribeResourceResponseBody) *DescribeResourceResponse {
	s.Body = v
	return s
}

type DescribeResourceDLinkResponseBody struct {
	// The IDs of the secondary vSwitches that are directly connected.
	AuxVSwitchList []*string `json:"AuxVSwitchList,omitempty" xml:"AuxVSwitchList,omitempty" type:"Repeated"`
	// The CIDR blocks of the clients that you want to connect to. After this parameter is specified, the CIDR blocks are added to the back-to-origin route of the server. Either this parameter or the VSwitchIdList parameter can be used to determine CIDR blocks.
	//
	// example:
	//
	// 72.16.0.0/16
	DestinationCIDRs *string `json:"DestinationCIDRs,omitempty" xml:"DestinationCIDRs,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the security group that is directly connected.
	//
	// example:
	//
	// sg-bp1j1z7297hcink9d****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the primary vSwitch that is directly connected.
	//
	// example:
	//
	// vsw-8vbqn2at0kljjxxxx****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC that is directly connected.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeResourceDLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceDLinkResponseBody) SetAuxVSwitchList(v []*string) *DescribeResourceDLinkResponseBody {
	s.AuxVSwitchList = v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetDestinationCIDRs(v string) *DescribeResourceDLinkResponseBody {
	s.DestinationCIDRs = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetRequestId(v string) *DescribeResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetSecurityGroupId(v string) *DescribeResourceDLinkResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetVSwitchId(v string) *DescribeResourceDLinkResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetVpcId(v string) *DescribeResourceDLinkResponseBody {
	s.VpcId = &v
	return s
}

type DescribeResourceDLinkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceDLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceDLinkResponse) SetHeaders(v map[string]*string) *DescribeResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceDLinkResponse) SetStatusCode(v int32) *DescribeResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceDLinkResponse) SetBody(v *DescribeResourceDLinkResponseBody) *DescribeResourceDLinkResponse {
	s.Body = v
	return s
}

type DescribeResourceLogResponseBody struct {
	// The Logstore of Log Service.
	//
	// example:
	//
	// access_log
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Network interfaces are updating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The Log Service project that is associated with the resource group.
	//
	// example:
	//
	// eas-r-asdasdasd-sls
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the resource group.
	//
	// example:
	//
	// ResourceReady
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogResponseBody) SetLogStore(v string) *DescribeResourceLogResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetMessage(v string) *DescribeResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetProjectName(v string) *DescribeResourceLogResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetRequestId(v string) *DescribeResourceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetStatus(v string) *DescribeResourceLogResponseBody {
	s.Status = &v
	return s
}

type DescribeResourceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogResponse) SetHeaders(v map[string]*string) *DescribeResourceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLogResponse) SetStatusCode(v int32) *DescribeResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLogResponse) SetBody(v *DescribeResourceLogResponseBody) *DescribeResourceLogResponse {
	s.Body = v
	return s
}

type DescribeServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Service           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponse) SetHeaders(v map[string]*string) *DescribeServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceResponse) SetStatusCode(v int32) *DescribeServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceResponse) SetBody(v *Service) *DescribeServiceResponse {
	s.Body = v
	return s
}

type DescribeServiceAutoScalerResponseBody struct {
	// The additional information about the Autoscaler policy, such as the interval of triggering Autoscaler.
	//
	// example:
	//
	// {
	//
	//   "behavior": {
	//
	//     "scaleDown": {
	//
	//       "stabilizationWindowSeconds": 150
	//
	//     }
	//
	//   }
	//
	// }
	Behavior map[string]interface{} `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// The metrics.
	CurrentMetrics []*DescribeServiceAutoScalerResponseBodyCurrentMetrics `json:"CurrentMetrics,omitempty" xml:"CurrentMetrics,omitempty" type:"Repeated"`
	// The maximum number of instances in the service.
	//
	// example:
	//
	// 8
	MaxReplica *int32 `json:"MaxReplica,omitempty" xml:"MaxReplica,omitempty"`
	// The minimum number of instances in the service.
	//
	// example:
	//
	// 3
	MinReplica *int32 `json:"MinReplica,omitempty" xml:"MinReplica,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The auto scaling policies.
	ScaleStrategies []*DescribeServiceAutoScalerResponseBodyScaleStrategies `json:"ScaleStrategies,omitempty" xml:"ScaleStrategies,omitempty" type:"Repeated"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeServiceAutoScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponseBody) SetBehavior(v map[string]interface{}) *DescribeServiceAutoScalerResponseBody {
	s.Behavior = v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetCurrentMetrics(v []*DescribeServiceAutoScalerResponseBodyCurrentMetrics) *DescribeServiceAutoScalerResponseBody {
	s.CurrentMetrics = v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetMaxReplica(v int32) *DescribeServiceAutoScalerResponseBody {
	s.MaxReplica = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetMinReplica(v int32) *DescribeServiceAutoScalerResponseBody {
	s.MinReplica = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetRequestId(v string) *DescribeServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetScaleStrategies(v []*DescribeServiceAutoScalerResponseBodyScaleStrategies) *DescribeServiceAutoScalerResponseBody {
	s.ScaleStrategies = v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetServiceName(v string) *DescribeServiceAutoScalerResponseBody {
	s.ServiceName = &v
	return s
}

type DescribeServiceAutoScalerResponseBodyCurrentMetrics struct {
	// The metric name. Valid values:
	//
	// 	- QPS
	//
	// 	- CPU
	//
	// example:
	//
	// qps
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The service for which the metric is specified.
	//
	// example:
	//
	// demo_svc
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 10
	Value *float32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeServiceAutoScalerResponseBodyCurrentMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceAutoScalerResponseBodyCurrentMetrics) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) SetMetricName(v string) *DescribeServiceAutoScalerResponseBodyCurrentMetrics {
	s.MetricName = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) SetService(v string) *DescribeServiceAutoScalerResponseBodyCurrentMetrics {
	s.Service = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) SetValue(v float32) *DescribeServiceAutoScalerResponseBodyCurrentMetrics {
	s.Value = &v
	return s
}

type DescribeServiceAutoScalerResponseBodyScaleStrategies struct {
	// The metric name. Valid values:
	//
	// 	- QPS: the queries per second (QPS) for an individual instance.
	//
	// 	- CPU: the CPU utilization.
	//
	// example:
	//
	// QPS
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The service for which the metric is specified. If you do not set this parameter, the current service is specified by default.
	//
	// example:
	//
	// demo_svc
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// The threshold of the metric that triggers auto scaling.
	//
	// 	- If you set metricName to QPS, scale-out is triggered when the average QPS for a single instance is greater than this threshold.
	//
	// 	- If you set metricName to CPU, scale-out is triggered when the average CPU utilization for a single instance is greater than this threshold.
	//
	// example:
	//
	// 10
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s DescribeServiceAutoScalerResponseBodyScaleStrategies) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceAutoScalerResponseBodyScaleStrategies) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) SetMetricName(v string) *DescribeServiceAutoScalerResponseBodyScaleStrategies {
	s.MetricName = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) SetService(v string) *DescribeServiceAutoScalerResponseBodyScaleStrategies {
	s.Service = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) SetThreshold(v float32) *DescribeServiceAutoScalerResponseBodyScaleStrategies {
	s.Threshold = &v
	return s
}

type DescribeServiceAutoScalerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceAutoScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponse) SetHeaders(v map[string]*string) *DescribeServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceAutoScalerResponse) SetStatusCode(v int32) *DescribeServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceAutoScalerResponse) SetBody(v *DescribeServiceAutoScalerResponseBody) *DescribeServiceAutoScalerResponse {
	s.Body = v
	return s
}

type DescribeServiceCronScalerResponseBody struct {
	// The points in time that are excluded when you schedule a CronHPA job. The points in time must be specified by using a cron expression.
	ExcludeDates []*string `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The CronHPA jobs.
	ScaleJobs []*DescribeServiceCronScalerResponseBodyScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeServiceCronScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponseBody) SetExcludeDates(v []*string) *DescribeServiceCronScalerResponseBody {
	s.ExcludeDates = v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetRequestId(v string) *DescribeServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetScaleJobs(v []*DescribeServiceCronScalerResponseBodyScaleJobs) *DescribeServiceCronScalerResponseBody {
	s.ScaleJobs = v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetServiceName(v string) *DescribeServiceCronScalerResponseBody {
	s.ServiceName = &v
	return s
}

type DescribeServiceCronScalerResponseBodyScaleJobs struct {
	// The time when the most recent CronHPA job was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-24T02:11:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the most recent CronHPA job ran. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-02-24T06:31:00Z
	LastProbeTime *string `json:"LastProbeTime,omitempty" xml:"LastProbeTime,omitempty"`
	// The returned message.
	//
	// example:
	//
	// "cron hpa job scale-jobs-0 executed successfully. current replicas:3, desired replicas:2."
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the CronHPA job.
	//
	// example:
	//
	// scale-job-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cron expression that is used to configure the execution time of the CronHPA job.
	//
	// example:
	//
	// 0 18 	- 	- 	- *
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The status of the most recent CronHPA job.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The number of instances that you expect to configure for the CronHPA job.
	//
	// example:
	//
	// 1
	TargetSize *int32 `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s DescribeServiceCronScalerResponseBodyScaleJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceCronScalerResponseBodyScaleJobs) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetCreateTime(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.CreateTime = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetLastProbeTime(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.LastProbeTime = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetMessage(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Message = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetName(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Name = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetSchedule(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Schedule = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetState(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.State = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetTargetSize(v int32) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.TargetSize = &v
	return s
}

type DescribeServiceCronScalerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceCronScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponse) SetHeaders(v map[string]*string) *DescribeServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceCronScalerResponse) SetStatusCode(v int32) *DescribeServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceCronScalerResponse) SetBody(v *DescribeServiceCronScalerResponseBody) *DescribeServiceCronScalerResponse {
	s.Body = v
	return s
}

type DescribeServiceDiagnosisResponseBody struct {
	// The diagnostics list.
	DiagnosisList []*DescribeServiceDiagnosisResponseBodyDiagnosisList `json:"DiagnosisList,omitempty" xml:"DiagnosisList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceDiagnosisResponseBody) SetDiagnosisList(v []*DescribeServiceDiagnosisResponseBodyDiagnosisList) *DescribeServiceDiagnosisResponseBody {
	s.DiagnosisList = v
	return s
}

func (s *DescribeServiceDiagnosisResponseBody) SetRequestId(v string) *DescribeServiceDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceDiagnosisResponseBodyDiagnosisList struct {
	// The suggestions about how to handle the errors.
	Advices []*string `json:"Advices,omitempty" xml:"Advices,omitempty" type:"Repeated"`
	// The causes of the errors.
	Causes []*string `json:"Causes,omitempty" xml:"Causes,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// Container worker0 failed to pull image.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s DescribeServiceDiagnosisResponseBodyDiagnosisList) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceDiagnosisResponseBodyDiagnosisList) GoString() string {
	return s.String()
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) SetAdvices(v []*string) *DescribeServiceDiagnosisResponseBodyDiagnosisList {
	s.Advices = v
	return s
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) SetCauses(v []*string) *DescribeServiceDiagnosisResponseBodyDiagnosisList {
	s.Causes = v
	return s
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) SetError(v string) *DescribeServiceDiagnosisResponseBodyDiagnosisList {
	s.Error = &v
	return s
}

type DescribeServiceDiagnosisResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceDiagnosisResponse) SetHeaders(v map[string]*string) *DescribeServiceDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceDiagnosisResponse) SetStatusCode(v int32) *DescribeServiceDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceDiagnosisResponse) SetBody(v *DescribeServiceDiagnosisResponseBody) *DescribeServiceDiagnosisResponse {
	s.Body = v
	return s
}

type DescribeServiceEventRequest struct {
	// The end of the time range to query. By default, the current point in time is the end of the time range to query.
	//
	// example:
	//
	// 2006-01-02 15:04:05
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event type. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// example:
	//
	// Normal
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The instance name. For more information about how to obtain the instance name, see [ListServiceInstances](https://help.aliyun.com/document_detail/412108.html).
	//
	// example:
	//
	// echo-test-784xxxx85d-hhnd8
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. The time must be in UTC. The default value is seven days ago.
	//
	// example:
	//
	// 2006-01-02 15:04:05
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeServiceEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceEventRequest) SetEndTime(v string) *DescribeServiceEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeServiceEventRequest) SetEventType(v string) *DescribeServiceEventRequest {
	s.EventType = &v
	return s
}

func (s *DescribeServiceEventRequest) SetInstanceName(v string) *DescribeServiceEventRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeServiceEventRequest) SetPageNum(v string) *DescribeServiceEventRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceEventRequest) SetPageSize(v string) *DescribeServiceEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeServiceEventRequest) SetStartTime(v string) *DescribeServiceEventRequest {
	s.StartTime = &v
	return s
}

type DescribeServiceEventResponseBody struct {
	// The events.
	Events []*DescribeServiceEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D491C94-6239-5318-B4B4-799D859***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 29
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 12
	TotalPageNum *int64 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s DescribeServiceEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceEventResponseBody) SetEvents(v []*DescribeServiceEventResponseBodyEvents) *DescribeServiceEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeServiceEventResponseBody) SetPageNum(v int64) *DescribeServiceEventResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceEventResponseBody) SetRequestId(v string) *DescribeServiceEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceEventResponseBody) SetTotalCount(v int64) *DescribeServiceEventResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeServiceEventResponseBody) SetTotalPageNum(v int64) *DescribeServiceEventResponseBody {
	s.TotalPageNum = &v
	return s
}

type DescribeServiceEventResponseBodyEvents struct {
	// The returned message. The message is formatted and returned in the JSON format.
	//
	// example:
	//
	// {\\"versionId\\":1,\\"message\\":\\"Stage scale complete\\",\\"availableInstance\\":1,\\"unavailableInstance\\":0}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The cause of the event. The information about the change in the service status is returned.
	//
	// example:
	//
	// Updating
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The time when the event occurred. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-09 06:30:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The event type. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeServiceEventResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeServiceEventResponseBodyEvents) SetMessage(v string) *DescribeServiceEventResponseBodyEvents {
	s.Message = &v
	return s
}

func (s *DescribeServiceEventResponseBodyEvents) SetReason(v string) *DescribeServiceEventResponseBodyEvents {
	s.Reason = &v
	return s
}

func (s *DescribeServiceEventResponseBodyEvents) SetTime(v string) *DescribeServiceEventResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *DescribeServiceEventResponseBodyEvents) SetType(v string) *DescribeServiceEventResponseBodyEvents {
	s.Type = &v
	return s
}

type DescribeServiceEventResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceEventResponse) SetHeaders(v map[string]*string) *DescribeServiceEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceEventResponse) SetStatusCode(v int32) *DescribeServiceEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceEventResponse) SetBody(v *DescribeServiceEventResponseBody) *DescribeServiceEventResponse {
	s.Body = v
	return s
}

type DescribeServiceInstanceDiagnosisResponseBody struct {
	// The diagnostics information.
	Diagnosis *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis `json:"Diagnosis,omitempty" xml:"Diagnosis,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceInstanceDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceInstanceDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceInstanceDiagnosisResponseBody) SetDiagnosis(v *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) *DescribeServiceInstanceDiagnosisResponseBody {
	s.Diagnosis = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponseBody) SetRequestId(v string) *DescribeServiceInstanceDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type DescribeServiceInstanceDiagnosisResponseBodyDiagnosis struct {
	// The solutions to the errors.
	Advices []*string `json:"Advices,omitempty" xml:"Advices,omitempty" type:"Repeated"`
	// The causes of the errors.
	Causes []*string `json:"Causes,omitempty" xml:"Causes,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// Container worker0 failed to pull image.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) GoString() string {
	return s.String()
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) SetAdvices(v []*string) *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis {
	s.Advices = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) SetCauses(v []*string) *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis {
	s.Causes = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) SetError(v string) *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis {
	s.Error = &v
	return s
}

type DescribeServiceInstanceDiagnosisResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceInstanceDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceInstanceDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceInstanceDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceInstanceDiagnosisResponse) SetHeaders(v map[string]*string) *DescribeServiceInstanceDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponse) SetStatusCode(v int32) *DescribeServiceInstanceDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponse) SetBody(v *DescribeServiceInstanceDiagnosisResponseBody) *DescribeServiceInstanceDiagnosisResponse {
	s.Body = v
	return s
}

type DescribeServiceLogRequest struct {
	// The name of the container that runs the service.
	//
	// example:
	//
	// worker0
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The end of the time range to query. The time must be in UTC.
	//
	// example:
	//
	// 2006-01-02 15:04:05
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the instance that runs the service. For more information about how to query the instance name, see [ListServiceInstances](https://help.aliyun.com/document_detail/412108.html).
	//
	// example:
	//
	// echo-da290ac8-7fckm
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP address of the instance whose logs you want to query. For more information about how to query the IP address of an instance, see [ListServiceInstances](https://help.aliyun.com/document_detail/412108.html).
	//
	// example:
	//
	// 10.0.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The keyword that you use to query the logs of the service.
	//
	// example:
	//
	// key
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 500.
	//
	// example:
	//
	// 500
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to query the logs that are generated before the instance last restarts. This parameter is available only if the instance restarts.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Previous *bool `json:"Previous,omitempty" xml:"Previous,omitempty"`
	// The beginning of the time range to query. The time must be in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2006-01-02 15:04:05
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeServiceLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogRequest) SetContainerName(v string) *DescribeServiceLogRequest {
	s.ContainerName = &v
	return s
}

func (s *DescribeServiceLogRequest) SetEndTime(v string) *DescribeServiceLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeServiceLogRequest) SetInstanceName(v string) *DescribeServiceLogRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeServiceLogRequest) SetIp(v string) *DescribeServiceLogRequest {
	s.Ip = &v
	return s
}

func (s *DescribeServiceLogRequest) SetKeyword(v string) *DescribeServiceLogRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeServiceLogRequest) SetPageNum(v int64) *DescribeServiceLogRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceLogRequest) SetPageSize(v int64) *DescribeServiceLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeServiceLogRequest) SetPrevious(v bool) *DescribeServiceLogRequest {
	s.Previous = &v
	return s
}

func (s *DescribeServiceLogRequest) SetStartTime(v string) *DescribeServiceLogRequest {
	s.StartTime = &v
	return s
}

type DescribeServiceLogResponseBody struct {
	// The returned logs.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 500
	TotalPageNum *int64 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s DescribeServiceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogResponseBody) SetLogs(v []*string) *DescribeServiceLogResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeServiceLogResponseBody) SetPageNum(v int64) *DescribeServiceLogResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetRequestId(v string) *DescribeServiceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetTotalCount(v int64) *DescribeServiceLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetTotalPageNum(v int64) *DescribeServiceLogResponseBody {
	s.TotalPageNum = &v
	return s
}

type DescribeServiceLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogResponse) SetHeaders(v map[string]*string) *DescribeServiceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceLogResponse) SetStatusCode(v int32) *DescribeServiceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceLogResponse) SetBody(v *DescribeServiceLogResponseBody) *DescribeServiceLogResponse {
	s.Body = v
	return s
}

type DescribeServiceMirrorResponseBody struct {
	// The percentage of traffic that you want to mirror. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The destination services to which you want to mirror traffic.
	//
	// example:
	//
	// foo2,foo3
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s DescribeServiceMirrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMirrorResponseBody) SetRatio(v string) *DescribeServiceMirrorResponseBody {
	s.Ratio = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetRequestId(v string) *DescribeServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetServiceName(v string) *DescribeServiceMirrorResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetTarget(v string) *DescribeServiceMirrorResponseBody {
	s.Target = &v
	return s
}

type DescribeServiceMirrorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMirrorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMirrorResponse) SetHeaders(v map[string]*string) *DescribeServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMirrorResponse) SetStatusCode(v int32) *DescribeServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMirrorResponse) SetBody(v *DescribeServiceMirrorResponseBody) *DescribeServiceMirrorResponse {
	s.Body = v
	return s
}

type DescribeSpotDiscountHistoryRequest struct {
	// The type of the Elastic Algorithm Service (EAS) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.c6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the preemptible instance has a protection period. During the 1-hour protection period of the preemptible instance, the preemptible instance will not be released.
	//
	// example:
	//
	// false
	IsProtect *bool `json:"IsProtect,omitempty" xml:"IsProtect,omitempty"`
}

func (s DescribeSpotDiscountHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpotDiscountHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpotDiscountHistoryRequest) SetInstanceType(v string) *DescribeSpotDiscountHistoryRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeSpotDiscountHistoryRequest) SetIsProtect(v bool) *DescribeSpotDiscountHistoryRequest {
	s.IsProtect = &v
	return s
}

type DescribeSpotDiscountHistoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The discount for the preemptible instance.
	SpotDiscounts []*DescribeSpotDiscountHistoryResponseBodySpotDiscounts `json:"SpotDiscounts,omitempty" xml:"SpotDiscounts,omitempty" type:"Repeated"`
}

func (s DescribeSpotDiscountHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpotDiscountHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpotDiscountHistoryResponseBody) SetRequestId(v string) *DescribeSpotDiscountHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBody) SetSpotDiscounts(v []*DescribeSpotDiscountHistoryResponseBodySpotDiscounts) *DescribeSpotDiscountHistoryResponseBody {
	s.SpotDiscounts = v
	return s
}

type DescribeSpotDiscountHistoryResponseBodySpotDiscounts struct {
	// The type of the ECS instance.
	//
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The discount for the preemptible instance. For example, 0.1 represents a 90% discount.
	//
	// example:
	//
	// 0.1
	SpotDiscount *string `json:"SpotDiscount,omitempty" xml:"SpotDiscount,omitempty"`
	// The time when the discount is available. The time must be in UTC.
	//
	// example:
	//
	// 2024-04-10T10:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSpotDiscountHistoryResponseBodySpotDiscounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpotDiscountHistoryResponseBodySpotDiscounts) GoString() string {
	return s.String()
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) SetInstanceType(v string) *DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	s.InstanceType = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) SetSpotDiscount(v string) *DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	s.SpotDiscount = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) SetTimestamp(v string) *DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	s.Timestamp = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) SetZoneId(v string) *DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	s.ZoneId = &v
	return s
}

type DescribeSpotDiscountHistoryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSpotDiscountHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSpotDiscountHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpotDiscountHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpotDiscountHistoryResponse) SetHeaders(v map[string]*string) *DescribeSpotDiscountHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpotDiscountHistoryResponse) SetStatusCode(v int32) *DescribeSpotDiscountHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponse) SetBody(v *DescribeSpotDiscountHistoryResponseBody) *DescribeSpotDiscountHistoryResponse {
	s.Body = v
	return s
}

type DevelopServiceRequest struct {
	// Specifies whether to exit development mode. Valid values:
	//
	// 	- true: exits development mode.
	//
	// 	- false (default): enters development mode.
	//
	// example:
	//
	// true
	Exit *string `json:"Exit,omitempty" xml:"Exit,omitempty"`
}

func (s DevelopServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DevelopServiceRequest) GoString() string {
	return s.String()
}

func (s *DevelopServiceRequest) SetExit(v string) *DevelopServiceRequest {
	s.Exit = &v
	return s
}

type DevelopServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DevelopServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DevelopServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DevelopServiceResponseBody) SetMessage(v string) *DevelopServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DevelopServiceResponseBody) SetRequestId(v string) *DevelopServiceResponseBody {
	s.RequestId = &v
	return s
}

type DevelopServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DevelopServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DevelopServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DevelopServiceResponse) GoString() string {
	return s.String()
}

func (s *DevelopServiceResponse) SetHeaders(v map[string]*string) *DevelopServiceResponse {
	s.Headers = v
	return s
}

func (s *DevelopServiceResponse) SetStatusCode(v int32) *DevelopServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DevelopServiceResponse) SetBody(v *DevelopServiceResponseBody) *DevelopServiceResponse {
	s.Body = v
	return s
}

type ListAclPolicyRequest struct {
	// The ID of the virtual private cloud (VPC). For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAclPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListAclPolicyRequest) SetVpcId(v string) *ListAclPolicyRequest {
	s.VpcId = &v
	return s
}

type ListAclPolicyResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The access control policies of the private gateway over the Internet.
	InternetAclPolicyList []*ListAclPolicyResponseBodyInternetAclPolicyList `json:"InternetAclPolicyList,omitempty" xml:"InternetAclPolicyList,omitempty" type:"Repeated"`
	// The access control policies of the private gateway over the internal network.
	IntranetVpcAclPolicyList []*ListAclPolicyResponseBodyIntranetVpcAclPolicyList `json:"IntranetVpcAclPolicyList,omitempty" xml:"IntranetVpcAclPolicyList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAclPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBody) SetGatewayId(v string) *ListAclPolicyResponseBody {
	s.GatewayId = &v
	return s
}

func (s *ListAclPolicyResponseBody) SetInternetAclPolicyList(v []*ListAclPolicyResponseBodyInternetAclPolicyList) *ListAclPolicyResponseBody {
	s.InternetAclPolicyList = v
	return s
}

func (s *ListAclPolicyResponseBody) SetIntranetVpcAclPolicyList(v []*ListAclPolicyResponseBodyIntranetVpcAclPolicyList) *ListAclPolicyResponseBody {
	s.IntranetVpcAclPolicyList = v
	return s
}

func (s *ListAclPolicyResponseBody) SetRequestId(v string) *ListAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ListAclPolicyResponseBodyInternetAclPolicyList struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway over the Internet.
	AclPolicyList []*ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
}

func (s ListAclPolicyResponseBodyInternetAclPolicyList) String() string {
	return tea.Prettify(s)
}

func (s ListAclPolicyResponseBodyInternetAclPolicyList) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyList) SetAclPolicyList(v []*ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) *ListAclPolicyResponseBodyInternetAclPolicyList {
	s.AclPolicyList = v
	return s
}

type ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList struct {
	// The comment on the IP CIDR block in the VPC that can access the private gateway over the Internet.
	//
	// example:
	//
	// default
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP CIDR block in the VPC that can access the private gateway over the Internet.
	//
	// example:
	//
	// 10.23.XX.XX/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) String() string {
	return tea.Prettify(s)
}

func (s ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) SetComment(v string) *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList {
	s.Comment = &v
	return s
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) SetEntry(v string) *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList {
	s.Entry = &v
	return s
}

type ListAclPolicyResponseBodyIntranetVpcAclPolicyList struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway over the internal network.
	AclPolicyList []*ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
	// The VPC ID. For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAclPolicyResponseBodyIntranetVpcAclPolicyList) String() string {
	return tea.Prettify(s)
}

func (s ListAclPolicyResponseBodyIntranetVpcAclPolicyList) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyList) SetAclPolicyList(v []*ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) *ListAclPolicyResponseBodyIntranetVpcAclPolicyList {
	s.AclPolicyList = v
	return s
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyList) SetVpcId(v string) *ListAclPolicyResponseBodyIntranetVpcAclPolicyList {
	s.VpcId = &v
	return s
}

type ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList struct {
	// The comment on the IP CIDR block in the VPC that can access the private gateway over the internal network.
	//
	// example:
	//
	// Test Entry
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP CIDR block in the VPC that can access the private gateway over the internal network.
	//
	// example:
	//
	// 192.168.XX.XX/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) String() string {
	return tea.Prettify(s)
}

func (s ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) SetComment(v string) *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList {
	s.Comment = &v
	return s
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) SetEntry(v string) *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList {
	s.Entry = &v
	return s
}

type ListAclPolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAclPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponse) SetHeaders(v map[string]*string) *ListAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListAclPolicyResponse) SetStatusCode(v int32) *ListAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAclPolicyResponse) SetBody(v *ListAclPolicyResponseBody) *ListAclPolicyResponse {
	s.Body = v
	return s
}

type ListBenchmarkTaskRequest struct {
	// The keyword used to query required stress testing tasks. If this parameter is specified, the system returns stress testing tasks based on the names of the stress testing tasks in the matched Elastic Algorithm Service (EAS).
	//
	// example:
	//
	// test_bench
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the EAS service that corresponds to the stress testing task. For more information about how to query the service name, see [ListServices](https://help.aliyun.com/document_detail/412109.html).
	//
	// example:
	//
	// test_bench_srv
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListBenchmarkTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskRequest) SetFilter(v string) *ListBenchmarkTaskRequest {
	s.Filter = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetPageNumber(v string) *ListBenchmarkTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetPageSize(v string) *ListBenchmarkTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetServiceName(v string) *ListBenchmarkTaskRequest {
	s.ServiceName = &v
	return s
}

type ListBenchmarkTaskResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stress testing tasks.
	Tasks []*ListBenchmarkTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponseBody) SetPageNumber(v int32) *ListBenchmarkTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetPageSize(v int32) *ListBenchmarkTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetRequestId(v string) *ListBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetTasks(v []*ListBenchmarkTaskResponseBodyTasks) *ListBenchmarkTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetTotalCount(v int32) *ListBenchmarkTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListBenchmarkTaskResponseBodyTasks struct {
	// The number of instances that are available for stress testing.
	//
	// example:
	//
	// 2
	AvailableAgent *int64 `json:"AvailableAgent,omitempty" xml:"AvailableAgent,omitempty"`
	// The time when the stress testing task was created.
	//
	// example:
	//
	// 2020-12-04T02:43:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-larec-test-1076] is Running
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID of the stress testing task.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the service on which you want to perform a stress testing.
	//
	// example:
	//
	// test_quota
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the stress testing task.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DeleteFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Error
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Updating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CreateFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the stress testing task.
	//
	// example:
	//
	// eas-b-gv4y86uvgt****i
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the stress testing task.
	//
	// example:
	//
	// benchmark-larec-test-1076
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The time when the stress testing task was updated.
	//
	// example:
	//
	// 2020-06-24T03:11:30Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListBenchmarkTaskResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListBenchmarkTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetAvailableAgent(v int64) *ListBenchmarkTaskResponseBodyTasks {
	s.AvailableAgent = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetCreateTime(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetMessage(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Message = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetRegion(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Region = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetServiceName(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.ServiceName = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetStatus(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetTaskId(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetTaskName(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetUpdateTime(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.UpdateTime = &v
	return s
}

type ListBenchmarkTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponse) SetHeaders(v map[string]*string) *ListBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *ListBenchmarkTaskResponse) SetStatusCode(v int32) *ListBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBenchmarkTaskResponse) SetBody(v *ListBenchmarkTaskResponseBody) *ListBenchmarkTaskResponse {
	s.Body = v
	return s
}

type ListGatewayRequest struct {
	// The private gateway ID. To obtain the private gateway ID, see the private_gateway_id parameter in the response parameters of the ListResources operation.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The private gateway alias.
	//
	// example:
	//
	// mygateway1
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ListGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRequest) SetGatewayId(v string) *ListGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayRequest) SetGatewayName(v string) *ListGatewayRequest {
	s.GatewayName = &v
	return s
}

func (s *ListGatewayRequest) SetPageNumber(v int32) *ListGatewayRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRequest) SetPageSize(v int32) *ListGatewayRequest {
	s.PageSize = &v
	return s
}

func (s *ListGatewayRequest) SetResourceName(v string) *ListGatewayRequest {
	s.ResourceName = &v
	return s
}

type ListGatewayResponseBody struct {
	// The private gateways.
	Gateways []*ListGatewayResponseBodyGateways `json:"Gateways,omitempty" xml:"Gateways,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of private gateways returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBody) SetGateways(v []*ListGatewayResponseBodyGateways) *ListGatewayResponseBody {
	s.Gateways = v
	return s
}

func (s *ListGatewayResponseBody) SetPageNumber(v int32) *ListGatewayResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayResponseBody) SetPageSize(v int32) *ListGatewayResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGatewayResponseBody) SetRequestId(v string) *ListGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayResponseBody) SetTotalCount(v int64) *ListGatewayResponseBody {
	s.TotalCount = &v
	return s
}

type ListGatewayResponseBodyGateways struct {
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the private gateway was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-19T14:19:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The private gateway alias.
	//
	// example:
	//
	// mygateway1
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The type of instances used for the private gateway.
	//
	// example:
	//
	// 2c4g
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******-1801786532******.cn-wulanchabu.pai-eas.aliyuncs.com
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// Indicates whether Internet access is enabled.
	//
	// example:
	//
	// true
	InternetEnabled *bool `json:"InternetEnabled,omitempty" xml:"InternetEnabled,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******-1801786532******-vpc.cn-wulanchabu.pai-eas.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// Indicates whether it is the default private gateway.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The number of nodes in the private gateway.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The state of the private gateway.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Deleted
	//
	// 	- Deleting
	//
	// 	- Waiting
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the private gateway was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-02-24T11:52:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGatewayResponseBodyGateways) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponseBodyGateways) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyGateways) SetChargeType(v string) *ListGatewayResponseBodyGateways {
	s.ChargeType = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetCreateTime(v string) *ListGatewayResponseBodyGateways {
	s.CreateTime = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetGatewayId(v string) *ListGatewayResponseBodyGateways {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetGatewayName(v string) *ListGatewayResponseBodyGateways {
	s.GatewayName = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetInstanceType(v string) *ListGatewayResponseBodyGateways {
	s.InstanceType = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetInternetDomain(v string) *ListGatewayResponseBodyGateways {
	s.InternetDomain = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetInternetEnabled(v bool) *ListGatewayResponseBodyGateways {
	s.InternetEnabled = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetIntranetDomain(v string) *ListGatewayResponseBodyGateways {
	s.IntranetDomain = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetIsDefault(v bool) *ListGatewayResponseBodyGateways {
	s.IsDefault = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetReplicas(v int32) *ListGatewayResponseBodyGateways {
	s.Replicas = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetStatus(v string) *ListGatewayResponseBodyGateways {
	s.Status = &v
	return s
}

func (s *ListGatewayResponseBodyGateways) SetUpdateTime(v string) *ListGatewayResponseBodyGateways {
	s.UpdateTime = &v
	return s
}

type ListGatewayResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayResponse) SetHeaders(v map[string]*string) *ListGatewayResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayResponse) SetStatusCode(v int32) *ListGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayResponse) SetBody(v *ListGatewayResponseBody) *ListGatewayResponse {
	s.Body = v
	return s
}

type ListGatewayIntranetLinkedVpcResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The internal endpoints.
	IntranetLinkedVpcList []*ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList `json:"IntranetLinkedVpcList,omitempty" xml:"IntranetLinkedVpcList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) SetGatewayId(v string) *ListGatewayIntranetLinkedVpcResponseBody {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) SetIntranetLinkedVpcList(v []*ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) *ListGatewayIntranetLinkedVpcResponseBody {
	s.IntranetLinkedVpcList = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) SetRequestId(v string) *ListGatewayIntranetLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

type ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.10.11
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2ze4pgstgszvgq******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The state of the private gateway.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The private gateway is being created.
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The private gateway is running.
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vb2qjoiio6m9pg******
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-2zetuli9ws0qgjd******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetIp(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.Ip = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetSecurityGroupId(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.SecurityGroupId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetStatus(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.Status = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetVSwitchId(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.VSwitchId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetVpcId(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.VpcId = &v
	return s
}

type ListGatewayIntranetLinkedVpcResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayIntranetLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcResponse) SetHeaders(v map[string]*string) *ListGatewayIntranetLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponse) SetStatusCode(v int32) *ListGatewayIntranetLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponse) SetBody(v *ListGatewayIntranetLinkedVpcResponseBody) *ListGatewayIntranetLinkedVpcResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	// The name of the filter that is used to filter out unwanted service groups. Fuzzy match is supported.
	//
	// example:
	//
	// foo
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetFilter(v string) *ListGroupsRequest {
	s.Filter = &v
	return s
}

func (s *ListGroupsRequest) SetPageNumber(v string) *ListGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsRequest) SetPageSize(v string) *ListGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsRequest) SetWorkspaceId(v string) *ListGroupsRequest {
	s.WorkspaceId = &v
	return s
}

type ListGroupsResponseBody struct {
	// The service groups.
	Groups []*Group `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetGroups(v []*Group) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetPageNumber(v int64) *ListGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsResponseBody) SetPageSize(v int64) *ListGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCount(v int64) *ListGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetHeaders(v map[string]*string) *ListGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsResponse) SetStatusCode(v int32) *ListGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListResourceInstanceWorkerRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListResourceInstanceWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstanceWorkerRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerRequest) SetPageNumber(v int32) *ListResourceInstanceWorkerRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetPageSize(v int32) *ListResourceInstanceWorkerRequest {
	s.PageSize = &v
	return s
}

type ListResourceInstanceWorkerResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workers.
	Pods []*ResourceInstanceWorker `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceInstanceWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstanceWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerResponseBody) SetPageNumber(v int32) *ListResourceInstanceWorkerResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetPageSize(v int32) *ListResourceInstanceWorkerResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetPods(v []*ResourceInstanceWorker) *ListResourceInstanceWorkerResponseBody {
	s.Pods = v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetRequestId(v string) *ListResourceInstanceWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetTotalCount(v int32) *ListResourceInstanceWorkerResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceInstanceWorkerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceInstanceWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceInstanceWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstanceWorkerResponse) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerResponse) SetHeaders(v map[string]*string) *ListResourceInstanceWorkerResponse {
	s.Headers = v
	return s
}

func (s *ListResourceInstanceWorkerResponse) SetStatusCode(v int32) *ListResourceInstanceWorkerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceInstanceWorkerResponse) SetBody(v *ListResourceInstanceWorkerResponseBody) *ListResourceInstanceWorkerResponse {
	s.Body = v
	return s
}

type ListResourceInstancesRequest struct {
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The keyword used to query instances. Instances can be queried by instance ID or instance IP address.
	//
	// example:
	//
	// 10.224.xx.xx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.224.xx.xx
	InstanceIP *string `json:"InstanceIP,omitempty" xml:"InstanceIP,omitempty"`
	// The instance ID. For more information about how to query the instance ID, see [ListResourceInstances](https://help.aliyun.com/document_detail/412129.html).
	//
	// example:
	//
	// i-bp1jd6x3uotsv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// e-xxxx***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance state.
	//
	// Valid values:
	//
	// 	- Ready-SchedulingDisabled
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instance is available but unschedulable
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Ready
	//
	//     <!-- -->
	//
	//     : The instance
	//
	//     <!-- -->
	//
	//     is running
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- NotReady
	//
	//     <!-- -->
	//
	//     : The instance is unready.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopped
	//
	//     <!-- -->
	//
	//     : The instance has stopped.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NotReady-SchedulingDisabled
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instance is unavailable and unschedulable
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Attaching
	//
	//     <!-- -->
	//
	//     : The instance
	//
	//     <!-- -->
	//
	//     is starting
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     : The instance is being deleted.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CreateFailed: The instance failed to be created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Ready
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The sorting order.
	//
	// Valid values:
	//
	// 	- asc: The instances are sorted in ascending order.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- desc
	//
	//     <!-- -->
	//
	//     : The instances are sorted in descending order.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field that you use to sort the query results.
	//
	// Valid values:
	//
	// 	- CreateTime
	//
	//     <!-- -->
	//
	//     : The instances are sorted based on the time when the instances were created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MemoryUsed
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instances are sorted based on the memory usage of the instances
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- GpuUsed
	//
	//     <!-- -->
	//
	//     : The instances are sorted based on the
	//
	//     <!-- -->
	//
	//     GPU usage of the instances.
	//
	//     <!-- -->
	//
	// 	- ExpireTime: The instances are sorted based on the time when the instances expired.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CpuUsed
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instances are sorted based on the CPU utilization of the instances.
	//
	//     <!-- -->
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s ListResourceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesRequest) SetChargeType(v string) *ListResourceInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *ListResourceInstancesRequest) SetFilter(v string) *ListResourceInstancesRequest {
	s.Filter = &v
	return s
}

func (s *ListResourceInstancesRequest) SetInstanceIP(v string) *ListResourceInstancesRequest {
	s.InstanceIP = &v
	return s
}

func (s *ListResourceInstancesRequest) SetInstanceId(v string) *ListResourceInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListResourceInstancesRequest) SetInstanceName(v string) *ListResourceInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListResourceInstancesRequest) SetInstanceStatus(v string) *ListResourceInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListResourceInstancesRequest) SetOrder(v string) *ListResourceInstancesRequest {
	s.Order = &v
	return s
}

func (s *ListResourceInstancesRequest) SetPageNumber(v int32) *ListResourceInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstancesRequest) SetPageSize(v int32) *ListResourceInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstancesRequest) SetSort(v string) *ListResourceInstancesRequest {
	s.Sort = &v
	return s
}

type ListResourceInstancesResponseBody struct {
	// The instances.
	Instances []*ResourceInstance `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponseBody) SetInstances(v []*ResourceInstance) *ListResourceInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListResourceInstancesResponseBody) SetPageNumber(v int32) *ListResourceInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetPageSize(v int32) *ListResourceInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetRequestId(v string) *ListResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetTotalCount(v int32) *ListResourceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponse) SetHeaders(v map[string]*string) *ListResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceInstancesResponse) SetStatusCode(v int32) *ListResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceInstancesResponse) SetBody(v *ListResourceInstancesResponseBody) *ListResourceInstancesResponse {
	s.Body = v
	return s
}

type ListResourceServicesRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListResourceServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceServicesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceServicesRequest) SetPageNumber(v int32) *ListResourceServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceServicesRequest) SetPageSize(v int32) *ListResourceServicesRequest {
	s.PageSize = &v
	return s
}

type ListResourceServicesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The services.
	Services []*Service `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceServicesResponseBody) SetPageNumber(v int32) *ListResourceServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetPageSize(v int32) *ListResourceServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetRequestId(v string) *ListResourceServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetServices(v []*Service) *ListResourceServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListResourceServicesResponseBody) SetTotalCount(v int32) *ListResourceServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceServicesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceServicesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceServicesResponse) SetHeaders(v map[string]*string) *ListResourceServicesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceServicesResponse) SetStatusCode(v int32) *ListResourceServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceServicesResponse) SetBody(v *ListResourceServicesResponseBody) *ListResourceServicesResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. You can call the [CreateResource](https://help.aliyun.com/document_detail/412111.html) operation to query the ID of the resource group.
	//
	// example:
	//
	// eas-r-h7lcw24dyqztwxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource group. You can call the [CreateResource](https://help.aliyun.com/document_detail/412111.html) operation to query the name of the resource group.
	//
	// example:
	//
	// MyResource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- Dedicated: the dedicated resource group.
	//
	// 	- SelfManaged: the self-managed resource group.
	//
	// example:
	//
	// Dedicated
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetResourceId(v string) *ListResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesRequest) SetResourceName(v string) *ListResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesRequest) SetResourceType(v string) *ListResourcesRequest {
	s.ResourceType = &v
	return s
}

type ListResourcesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource groups.
	Resources []*Resource `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetPageNumber(v int32) *ListResourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBody) SetPageSize(v int32) *ListResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetResources(v []*Resource) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int32) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListServiceContainersResponseBody struct {
	// The containers of the service.
	Containers []*ContainerInfo `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListServiceContainersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceContainersResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceContainersResponseBody) SetContainers(v []*ContainerInfo) *ListServiceContainersResponseBody {
	s.Containers = v
	return s
}

func (s *ListServiceContainersResponseBody) SetRequestId(v string) *ListServiceContainersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceContainersResponseBody) SetServiceName(v string) *ListServiceContainersResponseBody {
	s.ServiceName = &v
	return s
}

type ListServiceContainersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceContainersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceContainersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceContainersResponse) GoString() string {
	return s.String()
}

func (s *ListServiceContainersResponse) SetHeaders(v map[string]*string) *ListServiceContainersResponse {
	s.Headers = v
	return s
}

func (s *ListServiceContainersResponse) SetStatusCode(v int32) *ListServiceContainersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceContainersResponse) SetBody(v *ListServiceContainersResponseBody) *ListServiceContainersResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	// The keyword used to query instances. Instances can be queried based on instance name, instance IP address, IP address of the server where the instance resides, and instance type.
	//
	// example:
	//
	// 10.118.xx.xx
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The IP address of the server where the instance resides.
	//
	// example:
	//
	// 10.224.xx.xx
	HostIP *string `json:"HostIP,omitempty" xml:"HostIP,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.224.xx.xx
	InstanceIP *string `json:"InstanceIP,omitempty" xml:"InstanceIP,omitempty"`
	// The instance name.
	//
	// example:
	//
	// foo-bdc5xxxx-8l7rk
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance state.
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the instance is a preemptible instance.
	//
	// example:
	//
	// false
	IsSpot *bool `json:"IsSpot,omitempty" xml:"IsSpot,omitempty"`
	// The sorting order.
	//
	// Valid values:
	//
	// 	- asc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     The instances are sorted in ascending order.
	//
	// 	- desc
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     The instances are sorted in descending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource group to which the instance belongs.
	//
	// Valid values:
	//
	// 	- PublicResource
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DedicatedResource
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// PublicResource
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service role.
	//
	// Valid values:
	//
	// 	- DataSet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     dataset service
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- SDProxy
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Stable-Diffusion proxy service
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Standard
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     standard service
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Queue
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     queue service
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Queue
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The field that you use to sort the query results.
	//
	// 	- Set the value to StartTime.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     The value specifies that the query results are sorted based on the time when the instances were created
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// StartTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetFilter(v string) *ListServiceInstancesRequest {
	s.Filter = &v
	return s
}

func (s *ListServiceInstancesRequest) SetHostIP(v string) *ListServiceInstancesRequest {
	s.HostIP = &v
	return s
}

func (s *ListServiceInstancesRequest) SetInstanceIP(v string) *ListServiceInstancesRequest {
	s.InstanceIP = &v
	return s
}

func (s *ListServiceInstancesRequest) SetInstanceName(v string) *ListServiceInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListServiceInstancesRequest) SetInstanceStatus(v string) *ListServiceInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListServiceInstancesRequest) SetInstanceType(v string) *ListServiceInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListServiceInstancesRequest) SetIsSpot(v bool) *ListServiceInstancesRequest {
	s.IsSpot = &v
	return s
}

func (s *ListServiceInstancesRequest) SetOrder(v string) *ListServiceInstancesRequest {
	s.Order = &v
	return s
}

func (s *ListServiceInstancesRequest) SetPageNumber(v int32) *ListServiceInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceInstancesRequest) SetPageSize(v int32) *ListServiceInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceInstancesRequest) SetResourceType(v string) *ListServiceInstancesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRole(v string) *ListServiceInstancesRequest {
	s.Role = &v
	return s
}

func (s *ListServiceInstancesRequest) SetSort(v string) *ListServiceInstancesRequest {
	s.Sort = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	// The instances.
	Instances []*Instance `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetInstances(v []*Instance) *ListServiceInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetPageNumber(v int32) *ListServiceInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetPageSize(v int32) *ListServiceInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int32) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

type ListServiceVersionsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListServiceVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsRequest) SetPageNumber(v int32) *ListServiceVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceVersionsRequest) SetPageSize(v int32) *ListServiceVersionsRequest {
	s.PageSize = &v
	return s
}

type ListServiceVersionsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E089D584-B6F4-50C4-9902-DA2295B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 166
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The historical versions of the service.
	Versions []*ListServiceVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListServiceVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponseBody) SetPageNumber(v int32) *ListServiceVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetPageSize(v int32) *ListServiceVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetRequestId(v string) *ListServiceVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetTotalCount(v int64) *ListServiceVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceVersionsResponseBody) SetVersions(v []*ListServiceVersionsResponseBodyVersions) *ListServiceVersionsResponseBody {
	s.Versions = v
	return s
}

type ListServiceVersionsResponseBodyVersions struct {
	// The time when the service version was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-29T22:02:14Z
	BuildTime *string `json:"BuildTime,omitempty" xml:"BuildTime,omitempty"`
	// Indicates whether the image is available. Valid values:
	//
	// 	- true: The image is available.
	//
	// 	- false: The image is unavailable.
	//
	// 	- unknown: The availability of the image is unknown.
	//
	// example:
	//
	// true
	ImageAvailable *string `json:"ImageAvailable,omitempty" xml:"ImageAvailable,omitempty"`
	// The image ID.
	//
	// example:
	//
	// 4
	ImageId *int32 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Service is Running
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The service deployment configurations. This parameter is returned only if the service is deployed by using a custom image.
	//
	// example:
	//
	// {
	//
	//     "metadata": {
	//
	//         "cpu": 1,
	//
	//         "instance": 1,
	//
	//         "memory": 1024
	//
	//     },
	//
	//     "name": "echo"
	//
	// }
	ServiceConfig *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	// Indicates whether Elastic Algorithm service (EAS) is activated. Valid values:
	//
	// 	- true: EAS is activated.
	//
	// 	- false: EAS is not activated.
	//
	// 	- unknown: The activation of EAS is unknown.
	//
	// example:
	//
	// true
	ServiceRunnable *string `json:"ServiceRunnable,omitempty" xml:"ServiceRunnable,omitempty"`
}

func (s ListServiceVersionsResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponseBodyVersions) SetBuildTime(v string) *ListServiceVersionsResponseBodyVersions {
	s.BuildTime = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetImageAvailable(v string) *ListServiceVersionsResponseBodyVersions {
	s.ImageAvailable = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetImageId(v int32) *ListServiceVersionsResponseBodyVersions {
	s.ImageId = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetMessage(v string) *ListServiceVersionsResponseBodyVersions {
	s.Message = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetServiceConfig(v string) *ListServiceVersionsResponseBodyVersions {
	s.ServiceConfig = &v
	return s
}

func (s *ListServiceVersionsResponseBodyVersions) SetServiceRunnable(v string) *ListServiceVersionsResponseBodyVersions {
	s.ServiceRunnable = &v
	return s
}

type ListServiceVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceVersionsResponse) SetHeaders(v map[string]*string) *ListServiceVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceVersionsResponse) SetStatusCode(v int32) *ListServiceVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceVersionsResponse) SetBody(v *ListServiceVersionsResponseBody) *ListServiceVersionsResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	// The field that is used for fuzzy matches. The system performs fuzzy matches only by service name.
	//
	// example:
	//
	// foo
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The name of the service group. For more information about how to query the name of a service group, see [ListServices](https://help.aliyun.com/document_detail/412109.html).
	//
	// example:
	//
	// foo
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag that is used to filter services.
	Label map[string]*string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- desc (default): The query results are sorted in descending order.
	//
	// 	- asc: The query results are sorted in ascending order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the primary service that corresponds to the Band member service.
	//
	// example:
	//
	// eas-m-ijafy3c8cxxxx
	ParentServiceUid *string `json:"ParentServiceUid,omitempty" xml:"ParentServiceUid,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// quota12345
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The name or ID of the resource group to which the service belongs.
	//
	// example:
	//
	// eas-r-hd0qwy8cxxxx
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The server role.
	//
	// Valid values:
	//
	// 	- DataLoader
	//
	// 	- FrontEnd
	//
	// 	- DataSet
	//
	// 	- SDProxy
	//
	// 	- LLMSscheduler
	//
	// 	- ScalableJob
	//
	// 	- LLMGateway
	//
	// 	- Job
	//
	// 	- Queue
	//
	// example:
	//
	// LLMGateway
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The service name.
	//
	// example:
	//
	// echo_test
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service state.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopped
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Failed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Complete
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Cloning
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Updating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Waiting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- HotUpdate
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Committing
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DeleteFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Developing
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Scaling
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleted
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Pending
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The service type. Valid values:
	//
	// 	- Async
	//
	// 	- Standard
	//
	// 	- Offline Task
	//
	// 	- Proxima
	//
	// Valid values:
	//
	// 	- Async
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Standard
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- OfflineTask
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Proxima
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Standard
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The user ID (UID) of the service.
	//
	// example:
	//
	// eas-m-c9iw3yitxxxx
	ServiceUid *string `json:"ServiceUid,omitempty" xml:"ServiceUid,omitempty"`
	// The sort field. By default, the query results are sorted by the timestamp type in descending order.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetFilter(v string) *ListServicesRequest {
	s.Filter = &v
	return s
}

func (s *ListServicesRequest) SetGateway(v string) *ListServicesRequest {
	s.Gateway = &v
	return s
}

func (s *ListServicesRequest) SetGroupName(v string) *ListServicesRequest {
	s.GroupName = &v
	return s
}

func (s *ListServicesRequest) SetLabel(v map[string]*string) *ListServicesRequest {
	s.Label = v
	return s
}

func (s *ListServicesRequest) SetOrder(v string) *ListServicesRequest {
	s.Order = &v
	return s
}

func (s *ListServicesRequest) SetPageNumber(v int32) *ListServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServicesRequest) SetPageSize(v int32) *ListServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesRequest) SetParentServiceUid(v string) *ListServicesRequest {
	s.ParentServiceUid = &v
	return s
}

func (s *ListServicesRequest) SetQuotaId(v string) *ListServicesRequest {
	s.QuotaId = &v
	return s
}

func (s *ListServicesRequest) SetResourceName(v string) *ListServicesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListServicesRequest) SetRole(v string) *ListServicesRequest {
	s.Role = &v
	return s
}

func (s *ListServicesRequest) SetServiceName(v string) *ListServicesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListServicesRequest) SetServiceStatus(v string) *ListServicesRequest {
	s.ServiceStatus = &v
	return s
}

func (s *ListServicesRequest) SetServiceType(v string) *ListServicesRequest {
	s.ServiceType = &v
	return s
}

func (s *ListServicesRequest) SetServiceUid(v string) *ListServicesRequest {
	s.ServiceUid = &v
	return s
}

func (s *ListServicesRequest) SetSort(v string) *ListServicesRequest {
	s.Sort = &v
	return s
}

func (s *ListServicesRequest) SetWorkspaceId(v string) *ListServicesRequest {
	s.WorkspaceId = &v
	return s
}

type ListServicesShrinkRequest struct {
	// The field that is used for fuzzy matches. The system performs fuzzy matches only by service name.
	//
	// example:
	//
	// foo
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The name of the service group. For more information about how to query the name of a service group, see [ListServices](https://help.aliyun.com/document_detail/412109.html).
	//
	// example:
	//
	// foo
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag that is used to filter services.
	LabelShrink *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- desc (default): The query results are sorted in descending order.
	//
	// 	- asc: The query results are sorted in ascending order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the primary service that corresponds to the Band member service.
	//
	// example:
	//
	// eas-m-ijafy3c8cxxxx
	ParentServiceUid *string `json:"ParentServiceUid,omitempty" xml:"ParentServiceUid,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// quota12345
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The name or ID of the resource group to which the service belongs.
	//
	// example:
	//
	// eas-r-hd0qwy8cxxxx
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The server role.
	//
	// Valid values:
	//
	// 	- DataLoader
	//
	// 	- FrontEnd
	//
	// 	- DataSet
	//
	// 	- SDProxy
	//
	// 	- LLMSscheduler
	//
	// 	- ScalableJob
	//
	// 	- LLMGateway
	//
	// 	- Job
	//
	// 	- Queue
	//
	// example:
	//
	// LLMGateway
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The service name.
	//
	// example:
	//
	// echo_test
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service state.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopped
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Failed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Complete
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Cloning
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Updating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Waiting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- HotUpdate
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Committing
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DeleteFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Developing
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Scaling
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleted
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Pending
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The service type. Valid values:
	//
	// 	- Async
	//
	// 	- Standard
	//
	// 	- Offline Task
	//
	// 	- Proxima
	//
	// Valid values:
	//
	// 	- Async
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Standard
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- OfflineTask
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Proxima
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Standard
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The user ID (UID) of the service.
	//
	// example:
	//
	// eas-m-c9iw3yitxxxx
	ServiceUid *string `json:"ServiceUid,omitempty" xml:"ServiceUid,omitempty"`
	// The sort field. By default, the query results are sorted by the timestamp type in descending order.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListServicesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListServicesShrinkRequest) SetFilter(v string) *ListServicesShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListServicesShrinkRequest) SetGateway(v string) *ListServicesShrinkRequest {
	s.Gateway = &v
	return s
}

func (s *ListServicesShrinkRequest) SetGroupName(v string) *ListServicesShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *ListServicesShrinkRequest) SetLabelShrink(v string) *ListServicesShrinkRequest {
	s.LabelShrink = &v
	return s
}

func (s *ListServicesShrinkRequest) SetOrder(v string) *ListServicesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListServicesShrinkRequest) SetPageNumber(v int32) *ListServicesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServicesShrinkRequest) SetPageSize(v int32) *ListServicesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesShrinkRequest) SetParentServiceUid(v string) *ListServicesShrinkRequest {
	s.ParentServiceUid = &v
	return s
}

func (s *ListServicesShrinkRequest) SetQuotaId(v string) *ListServicesShrinkRequest {
	s.QuotaId = &v
	return s
}

func (s *ListServicesShrinkRequest) SetResourceName(v string) *ListServicesShrinkRequest {
	s.ResourceName = &v
	return s
}

func (s *ListServicesShrinkRequest) SetRole(v string) *ListServicesShrinkRequest {
	s.Role = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceName(v string) *ListServicesShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceStatus(v string) *ListServicesShrinkRequest {
	s.ServiceStatus = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceType(v string) *ListServicesShrinkRequest {
	s.ServiceType = &v
	return s
}

func (s *ListServicesShrinkRequest) SetServiceUid(v string) *ListServicesShrinkRequest {
	s.ServiceUid = &v
	return s
}

func (s *ListServicesShrinkRequest) SetSort(v string) *ListServicesShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListServicesShrinkRequest) SetWorkspaceId(v string) *ListServicesShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type ListServicesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The services.
	Services []*Service `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetPageNumber(v int32) *ListServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServicesResponseBody) SetPageSize(v int32) *ListServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v []*Service) *ListServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v int32) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListTenantAddonsResponseBody struct {
	// The information about the plug-in.
	Addons []*ListTenantAddonsResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTenantAddonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTenantAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantAddonsResponseBody) SetAddons(v []*ListTenantAddonsResponseBodyAddons) *ListTenantAddonsResponseBody {
	s.Addons = v
	return s
}

func (s *ListTenantAddonsResponseBody) SetRequestId(v string) *ListTenantAddonsResponseBody {
	s.RequestId = &v
	return s
}

type ListTenantAddonsResponseBodyAddons struct {
	// The attributes of the plug-in.
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// prometheus_discovery
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListTenantAddonsResponseBodyAddons) String() string {
	return tea.Prettify(s)
}

func (s ListTenantAddonsResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListTenantAddonsResponseBodyAddons) SetAttributes(v map[string]*string) *ListTenantAddonsResponseBodyAddons {
	s.Attributes = v
	return s
}

func (s *ListTenantAddonsResponseBodyAddons) SetName(v string) *ListTenantAddonsResponseBodyAddons {
	s.Name = &v
	return s
}

type ListTenantAddonsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantAddonsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTenantAddonsResponse) GoString() string {
	return s.String()
}

func (s *ListTenantAddonsResponse) SetHeaders(v map[string]*string) *ListTenantAddonsResponse {
	s.Headers = v
	return s
}

func (s *ListTenantAddonsResponse) SetStatusCode(v int32) *ListTenantAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantAddonsResponse) SetBody(v *ListTenantAddonsResponseBody) *ListTenantAddonsResponse {
	s.Body = v
	return s
}

type ReinstallTenantAddonResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Addon prometheus_discovery is successfully reinstalled
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReinstallTenantAddonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReinstallTenantAddonResponseBody) GoString() string {
	return s.String()
}

func (s *ReinstallTenantAddonResponseBody) SetMessage(v string) *ReinstallTenantAddonResponseBody {
	s.Message = &v
	return s
}

func (s *ReinstallTenantAddonResponseBody) SetRequestId(v string) *ReinstallTenantAddonResponseBody {
	s.RequestId = &v
	return s
}

type ReinstallTenantAddonResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReinstallTenantAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReinstallTenantAddonResponse) String() string {
	return tea.Prettify(s)
}

func (s ReinstallTenantAddonResponse) GoString() string {
	return s.String()
}

func (s *ReinstallTenantAddonResponse) SetHeaders(v map[string]*string) *ReinstallTenantAddonResponse {
	s.Headers = v
	return s
}

func (s *ReinstallTenantAddonResponse) SetStatusCode(v int32) *ReinstallTenantAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *ReinstallTenantAddonResponse) SetBody(v *ReinstallTenantAddonResponseBody) *ReinstallTenantAddonResponse {
	s.Body = v
	return s
}

type ReleaseServiceRequest struct {
	// The traffic state. Valid values:
	//
	// 	- standalone: independent traffic.
	//
	// 	- grouping: grouped traffic.
	//
	// example:
	//
	// grouping
	TrafficState *string `json:"TrafficState,omitempty" xml:"TrafficState,omitempty"`
	// The weight of the canary release. Valid values: 0 to 100.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ReleaseServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseServiceRequest) SetTrafficState(v string) *ReleaseServiceRequest {
	s.TrafficState = &v
	return s
}

func (s *ReleaseServiceRequest) SetWeight(v int32) *ReleaseServiceRequest {
	s.Weight = &v
	return s
}

type ReleaseServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Releasing service [foo] in region [cn-shanghai] with weight [40], service status: [Running]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseServiceResponseBody) SetMessage(v string) *ReleaseServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseServiceResponseBody) SetRequestId(v string) *ReleaseServiceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseServiceResponse) SetHeaders(v map[string]*string) *ReleaseServiceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseServiceResponse) SetStatusCode(v int32) *ReleaseServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseServiceResponse) SetBody(v *ReleaseServiceResponseBody) *ReleaseServiceResponse {
	s.Body = v
	return s
}

type RestartServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Service is restarting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartServiceResponseBody) SetMessage(v string) *RestartServiceResponseBody {
	s.Message = &v
	return s
}

func (s *RestartServiceResponseBody) SetRequestId(v string) *RestartServiceResponseBody {
	s.RequestId = &v
	return s
}

type RestartServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceResponse) GoString() string {
	return s.String()
}

func (s *RestartServiceResponse) SetHeaders(v map[string]*string) *RestartServiceResponse {
	s.Headers = v
	return s
}

func (s *RestartServiceResponse) SetStatusCode(v int32) *RestartServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartServiceResponse) SetBody(v *RestartServiceResponseBody) *RestartServiceResponse {
	s.Body = v
	return s
}

type StartBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-larec-test-1076] is Starting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartBenchmarkTaskResponseBody) SetMessage(v string) *StartBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartBenchmarkTaskResponseBody) SetRequestId(v string) *StartBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type StartBenchmarkTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *StartBenchmarkTaskResponse) SetHeaders(v map[string]*string) *StartBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *StartBenchmarkTaskResponse) SetStatusCode(v int32) *StartBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBenchmarkTaskResponse) SetBody(v *StartBenchmarkTaskResponseBody) *StartBenchmarkTaskResponse {
	s.Body = v
	return s
}

type StartServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Service [foo] in region [cn-shanghai] is starting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceResponseBody) SetMessage(v string) *StartServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartServiceResponseBody) SetRequestId(v string) *StartServiceResponseBody {
	s.RequestId = &v
	return s
}

type StartServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartServiceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceResponse) SetHeaders(v map[string]*string) *StartServiceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceResponse) SetStatusCode(v int32) *StartServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceResponse) SetBody(v *StartServiceResponseBody) *StartServiceResponse {
	s.Body = v
	return s
}

type StopBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-larec-test-1076] is Stopping
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopBenchmarkTaskResponseBody) SetMessage(v string) *StopBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopBenchmarkTaskResponseBody) SetRequestId(v string) *StopBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopBenchmarkTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *StopBenchmarkTaskResponse) SetHeaders(v map[string]*string) *StopBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *StopBenchmarkTaskResponse) SetStatusCode(v int32) *StopBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopBenchmarkTaskResponse) SetBody(v *StopBenchmarkTaskResponseBody) *StopBenchmarkTaskResponse {
	s.Body = v
	return s
}

type StopServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to auto scale service [foo]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServiceResponseBody) SetMessage(v string) *StopServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StopServiceResponseBody) SetRequestId(v string) *StopServiceResponseBody {
	s.RequestId = &v
	return s
}

type StopServiceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopServiceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceResponse) SetHeaders(v map[string]*string) *StopServiceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceResponse) SetStatusCode(v int32) *StopServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceResponse) SetBody(v *StopServiceResponseBody) *StopServiceResponse {
	s.Body = v
	return s
}

type UpdateAppServiceRequest struct {
	// The quota ID.
	//
	// example:
	//
	// abcdef
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123456
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The application type.
	//
	// Valid values:
	//
	// 	- LLM: the large language model (LLM) application
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// LLM
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The application version.
	//
	// example:
	//
	// v1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The additional configurations that are required for service deployment.
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The number of instances. This value must be greater than 0.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The service specifications. Valid values:
	//
	// 	- llama_7b_fp16
	//
	// 	- llama_7b_int8
	//
	// 	- llama_13b_fp16
	//
	// 	- llama_7b_int8
	//
	// 	- chatglm_6b_fp16
	//
	// 	- chatglm_6b_int8
	//
	// 	- chatglm2_6b_fp16
	//
	// 	- baichuan_7b_int8
	//
	// 	- baichuan_13b_fp16
	//
	// 	- baichuan_7b_fp16
	//
	// example:
	//
	// llama_7b_fp16
	ServiceSpec *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
}

func (s UpdateAppServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppServiceRequest) SetQuotaId(v string) *UpdateAppServiceRequest {
	s.QuotaId = &v
	return s
}

func (s *UpdateAppServiceRequest) SetWorkspaceId(v string) *UpdateAppServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateAppServiceRequest) SetAppType(v string) *UpdateAppServiceRequest {
	s.AppType = &v
	return s
}

func (s *UpdateAppServiceRequest) SetAppVersion(v string) *UpdateAppServiceRequest {
	s.AppVersion = &v
	return s
}

func (s *UpdateAppServiceRequest) SetConfig(v map[string]interface{}) *UpdateAppServiceRequest {
	s.Config = v
	return s
}

func (s *UpdateAppServiceRequest) SetReplicas(v int32) *UpdateAppServiceRequest {
	s.Replicas = &v
	return s
}

func (s *UpdateAppServiceRequest) SetServiceSpec(v string) *UpdateAppServiceRequest {
	s.ServiceSpec = &v
	return s
}

type UpdateAppServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succee
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppServiceResponseBody) SetMessage(v string) *UpdateAppServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppServiceResponseBody) SetRequestId(v string) *UpdateAppServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppServiceResponse) SetHeaders(v map[string]*string) *UpdateAppServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppServiceResponse) SetStatusCode(v int32) *UpdateAppServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppServiceResponse) SetBody(v *UpdateAppServiceResponseBody) *UpdateAppServiceResponse {
	s.Body = v
	return s
}

type UpdateBenchmarkTaskRequest struct {
	// The request body. The body includes the parameters that are set to create a stress testing task. For more information, see **Table 1. Fields in the base parameter**.
	//
	// example:
	//
	// {
	//
	//     "base":  {
	//
	//          "qps": 200
	//
	//     }
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBenchmarkTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskRequest) SetBody(v string) *UpdateBenchmarkTaskRequest {
	s.Body = &v
	return s
}

type UpdateBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-larec-test-1076] is Updating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskResponseBody) SetMessage(v string) *UpdateBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBenchmarkTaskResponseBody) SetRequestId(v string) *UpdateBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskResponse) SetHeaders(v map[string]*string) *UpdateBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateBenchmarkTaskResponse) SetStatusCode(v int32) *UpdateBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBenchmarkTaskResponse) SetBody(v *UpdateBenchmarkTaskResponseBody) *UpdateBenchmarkTaskResponse {
	s.Body = v
	return s
}

type UpdateGatewayRequest struct {
	// Specifies whether to enable Internet access. Default value: false.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	EnableInternet *bool `json:"EnableInternet,omitempty" xml:"EnableInternet,omitempty"`
	// Specifies whether to enable internal network access. Default value: true.
	//
	// example:
	//
	// true
	EnableIntranet *bool `json:"EnableIntranet,omitempty" xml:"EnableIntranet,omitempty"`
	// The instance type used for the private gateway.
	//
	// example:
	//
	// ecs.c6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Indicates whether it is the default private gateway.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The private gateway alias.
	//
	// example:
	//
	// mygateway1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of nodes in the private gateway.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s UpdateGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRequest) SetEnableInternet(v bool) *UpdateGatewayRequest {
	s.EnableInternet = &v
	return s
}

func (s *UpdateGatewayRequest) SetEnableIntranet(v bool) *UpdateGatewayRequest {
	s.EnableIntranet = &v
	return s
}

func (s *UpdateGatewayRequest) SetInstanceType(v string) *UpdateGatewayRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateGatewayRequest) SetIsDefault(v bool) *UpdateGatewayRequest {
	s.IsDefault = &v
	return s
}

func (s *UpdateGatewayRequest) SetName(v string) *UpdateGatewayRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRequest) SetReplicas(v int32) *UpdateGatewayRequest {
	s.Replicas = &v
	return s
}

type UpdateGatewayResponseBody struct {
	// The ID of the gateway.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Gateway is updated
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponseBody) SetGatewayId(v string) *UpdateGatewayResponseBody {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetMessage(v string) *UpdateGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayResponseBody) SetRequestId(v string) *UpdateGatewayResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGatewayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayResponse) SetHeaders(v map[string]*string) *UpdateGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayResponse) SetStatusCode(v int32) *UpdateGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayResponse) SetBody(v *UpdateGatewayResponseBody) *UpdateGatewayResponse {
	s.Body = v
	return s
}

type UpdateResourceRequest struct {
	// The new name of the resource group after the update. The name can be up to 27 characters in length.
	//
	// example:
	//
	// iot
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The configuration items of the self-managed resource group.
	SelfManagedResourceOptions *UpdateResourceRequestSelfManagedResourceOptions `json:"SelfManagedResourceOptions,omitempty" xml:"SelfManagedResourceOptions,omitempty" type:"Struct"`
}

func (s UpdateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) SetResourceName(v string) *UpdateResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *UpdateResourceRequest) SetSelfManagedResourceOptions(v *UpdateResourceRequestSelfManagedResourceOptions) *UpdateResourceRequest {
	s.SelfManagedResourceOptions = v
	return s
}

type UpdateResourceRequestSelfManagedResourceOptions struct {
	// Tag tag key-value pairs for nodes.
	NodeMatchLabels map[string]*string `json:"NodeMatchLabels,omitempty" xml:"NodeMatchLabels,omitempty"`
	// Tolerations for nodes.
	NodeTolerations []*UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations `json:"NodeTolerations,omitempty" xml:"NodeTolerations,omitempty" type:"Repeated"`
}

func (s UpdateResourceRequestSelfManagedResourceOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequestSelfManagedResourceOptions) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequestSelfManagedResourceOptions) SetNodeMatchLabels(v map[string]*string) *UpdateResourceRequestSelfManagedResourceOptions {
	s.NodeMatchLabels = v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptions) SetNodeTolerations(v []*UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) *UpdateResourceRequestSelfManagedResourceOptions {
	s.NodeTolerations = v
	return s
}

type UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations struct {
	// The effect.
	//
	// Valid values:
	//
	// - PreferNoSchedule
	//
	// - NoSchedule
	//
	// - NoExecute
	//
	// example:
	//
	// NoSchedule
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The key name.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Relationship between key names and key values.
	//
	// Valid values:
	//
	// - Equal
	//
	// - Exists
	//
	// example:
	//
	// Equal
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The key value.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetEffect(v string) *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Effect = &v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetKey(v string) *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Key = &v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetOperator(v string) *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Operator = &v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetValue(v string) *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Value = &v
	return s
}

type UpdateResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// eas-r-asdasdasd
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// iot
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetResourceId(v string) *UpdateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetResourceName(v string) *UpdateResourceResponseBody {
	s.ResourceName = &v
	return s
}

type UpdateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetStatusCode(v int32) *UpdateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *UpdateResourceResponseBody) *UpdateResourceResponse {
	s.Body = v
	return s
}

type UpdateResourceDLinkRequest struct {
	// The CIDR blocks of the clients that you want to connect to. After this parameter is specified, the CIDR blocks are added to the back-to-origin route of the server. Either this parameter or the VSwitchIdList parameter can be used to determine CIDR blocks.
	//
	// example:
	//
	// 72.16.0.0/16
	DestinationCIDRs *string `json:"DestinationCIDRs,omitempty" xml:"DestinationCIDRs,omitempty"`
	// The ID of the security group to which the Elastic Compute Service (ECS) instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp149cedsfx2rfspd2d
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the peer primary vSwitch. After this parameter is specified, an elastic network interface (ENI) is created in the VSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitches of the clients that you want to connect to. After this parameter is specified, the CIDR blocks of these vSwitches are added to the back-to-origin route of the server.
	VSwitchIdList []*string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty" type:"Repeated"`
}

func (s UpdateResourceDLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceDLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkRequest) SetDestinationCIDRs(v string) *UpdateResourceDLinkRequest {
	s.DestinationCIDRs = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetSecurityGroupId(v string) *UpdateResourceDLinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetVSwitchId(v string) *UpdateResourceDLinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetVSwitchIdList(v []*string) *UpdateResourceDLinkRequest {
	s.VSwitchIdList = v
	return s
}

type UpdateResourceDLinkResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Network interfaces are updating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceDLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkResponseBody) SetMessage(v string) *UpdateResourceDLinkResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateResourceDLinkResponseBody) SetRequestId(v string) *UpdateResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceDLinkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceDLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkResponse) SetHeaders(v map[string]*string) *UpdateResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceDLinkResponse) SetStatusCode(v int32) *UpdateResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceDLinkResponse) SetBody(v *UpdateResourceDLinkResponseBody) *UpdateResourceDLinkResponse {
	s.Body = v
	return s
}

type UpdateResourceInstanceRequest struct {
	// The operation that updates the scheduling state of the instance in a dedicated resource group. Valid values:
	//
	// 	- Uncordon: allows scheduling the service to this instance.
	//
	// 	- Cordon: prohibits scheduling the service to this instance.
	//
	// 	- Drain: evicts the service that has been scheduled to this instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cordon
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
}

func (s UpdateResourceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceRequest) SetAction(v string) *UpdateResourceInstanceRequest {
	s.Action = &v
	return s
}

type UpdateResourceInstanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// eas-i-asdasdasd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// eas-r-asdasdasd
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s UpdateResourceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceResponseBody) SetInstanceId(v string) *UpdateResourceInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceInstanceResponseBody) SetRequestId(v string) *UpdateResourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceInstanceResponseBody) SetResourceId(v string) *UpdateResourceInstanceResponseBody {
	s.ResourceId = &v
	return s
}

type UpdateResourceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceResponse) SetHeaders(v map[string]*string) *UpdateResourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceInstanceResponse) SetStatusCode(v int32) *UpdateResourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceInstanceResponse) SetBody(v *UpdateResourceInstanceResponseBody) *UpdateResourceInstanceResponse {
	s.Body = v
	return s
}

type UpdateServiceRequest struct {
	// The type of the service update. Valid values: merge and replace. By default, merge is used if you do not specify this parameter.
	//
	// 	- merge: If the JSON string configured for the existing service is `{"a":"b"}` and the JSON string specified in the body parameter is `{"c":"d"}`, the JSON string is `{"a":"b","c":"d"}` after the service update.
	//
	// 	- replace: If the JSON string configured for the existing service is `{"a":"b"}` and the JSON string specified in the body parameter is `{"c":"d"}`, the JSON string is `{"c":"d"}` after the service update.
	//
	// example:
	//
	// merge
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// The request body. The body includes the request parameters that you want to update. For more information about the request parameters, see [CreateService](https://help.aliyun.com/document_detail/412086.html).
	//
	// example:
	//
	// {   "name": "foo",   "model_path": "http://path/to/model.tar.gz",   "processor": "tensorflow_cpu",   "metadata": {     "instance": 2,     "memory": 7000,     "cpu": 4   } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetUpdateType(v string) *UpdateServiceRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateServiceRequest) SetBody(v string) *UpdateServiceRequest {
	s.Body = &v
	return s
}

type UpdateServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to update service [foo] in region [cn-shanghai]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetMessage(v string) *UpdateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetStatusCode(v int32) *UpdateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type UpdateServiceAutoScalerRequest struct {
	// The Autoscaler operation.
	Behavior *UpdateServiceAutoScalerRequestBehavior `json:"behavior,omitempty" xml:"behavior,omitempty" type:"Struct"`
	// The maximum number of instances. The value must be greater than that of the min parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	Max *int32 `json:"max,omitempty" xml:"max,omitempty"`
	// The minimum number of instances. The value must be greater than 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// The auto scaling policies.
	//
	// This parameter is required.
	ScaleStrategies []*UpdateServiceAutoScalerRequestScaleStrategies `json:"scaleStrategies,omitempty" xml:"scaleStrategies,omitempty" type:"Repeated"`
}

func (s UpdateServiceAutoScalerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequest) SetBehavior(v *UpdateServiceAutoScalerRequestBehavior) *UpdateServiceAutoScalerRequest {
	s.Behavior = v
	return s
}

func (s *UpdateServiceAutoScalerRequest) SetMax(v int32) *UpdateServiceAutoScalerRequest {
	s.Max = &v
	return s
}

func (s *UpdateServiceAutoScalerRequest) SetMin(v int32) *UpdateServiceAutoScalerRequest {
	s.Min = &v
	return s
}

func (s *UpdateServiceAutoScalerRequest) SetScaleStrategies(v []*UpdateServiceAutoScalerRequestScaleStrategies) *UpdateServiceAutoScalerRequest {
	s.ScaleStrategies = v
	return s
}

type UpdateServiceAutoScalerRequestBehavior struct {
	// The operation that reduces the number of instances to 0.
	OnZero *UpdateServiceAutoScalerRequestBehaviorOnZero `json:"onZero,omitempty" xml:"onZero,omitempty" type:"Struct"`
	// The scale-in operation.
	ScaleDown *UpdateServiceAutoScalerRequestBehaviorScaleDown `json:"scaleDown,omitempty" xml:"scaleDown,omitempty" type:"Struct"`
	// The scale-out operation.
	ScaleUp *UpdateServiceAutoScalerRequestBehaviorScaleUp `json:"scaleUp,omitempty" xml:"scaleUp,omitempty" type:"Struct"`
}

func (s UpdateServiceAutoScalerRequestBehavior) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestBehavior) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestBehavior) SetOnZero(v *UpdateServiceAutoScalerRequestBehaviorOnZero) *UpdateServiceAutoScalerRequestBehavior {
	s.OnZero = v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehavior) SetScaleDown(v *UpdateServiceAutoScalerRequestBehaviorScaleDown) *UpdateServiceAutoScalerRequestBehavior {
	s.ScaleDown = v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehavior) SetScaleUp(v *UpdateServiceAutoScalerRequestBehaviorScaleUp) *UpdateServiceAutoScalerRequestBehavior {
	s.ScaleUp = v
	return s
}

type UpdateServiceAutoScalerRequestBehaviorOnZero struct {
	// The time window that is required before the number of instances is reduced to 0. Default value: 600. The number of instances can be reduced to 0 only if no request is available or no traffic exists in the specified time window.
	//
	// example:
	//
	// 600
	ScaleDownGracePeriodSeconds *int32 `json:"scaleDownGracePeriodSeconds,omitempty" xml:"scaleDownGracePeriodSeconds,omitempty"`
	// The number of instances that you want to create at a time if the number of instances is scaled out from 0. Default value: 1.
	//
	// example:
	//
	// 1
	ScaleUpActivationReplicas *int32 `json:"scaleUpActivationReplicas,omitempty" xml:"scaleUpActivationReplicas,omitempty"`
}

func (s UpdateServiceAutoScalerRequestBehaviorOnZero) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestBehaviorOnZero) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestBehaviorOnZero) SetScaleDownGracePeriodSeconds(v int32) *UpdateServiceAutoScalerRequestBehaviorOnZero {
	s.ScaleDownGracePeriodSeconds = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestBehaviorOnZero) SetScaleUpActivationReplicas(v int32) *UpdateServiceAutoScalerRequestBehaviorOnZero {
	s.ScaleUpActivationReplicas = &v
	return s
}

type UpdateServiceAutoScalerRequestBehaviorScaleDown struct {
	// The time window that is required before the scale-in operation is performed. Default value: 300. The scale-in operation can be performed only if the specified metric drops below the threshold in the specified time window.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty" xml:"stabilizationWindowSeconds,omitempty"`
}

func (s UpdateServiceAutoScalerRequestBehaviorScaleDown) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestBehaviorScaleDown) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestBehaviorScaleDown) SetStabilizationWindowSeconds(v int32) *UpdateServiceAutoScalerRequestBehaviorScaleDown {
	s.StabilizationWindowSeconds = &v
	return s
}

type UpdateServiceAutoScalerRequestBehaviorScaleUp struct {
	// The time window that is required before the scale-out operation is performed. Default value: 0. The scale-out operation can be performed only if the specified metric exceeds the specified threshold in the specified time window.
	//
	// example:
	//
	// 0
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty" xml:"stabilizationWindowSeconds,omitempty"`
}

func (s UpdateServiceAutoScalerRequestBehaviorScaleUp) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestBehaviorScaleUp) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestBehaviorScaleUp) SetStabilizationWindowSeconds(v int32) *UpdateServiceAutoScalerRequestBehaviorScaleUp {
	s.StabilizationWindowSeconds = &v
	return s
}

type UpdateServiceAutoScalerRequestScaleStrategies struct {
	// The name of the metric for triggering auto scaling. Valid values:
	//
	// 	- qps: the queries per second (QPS) for an individual instance.
	//
	// 	- cpu: the CPU utilization.
	//
	// This parameter is required.
	//
	// example:
	//
	// qps
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The service for which the metric is specified. If you do not set this parameter, the current service is specified by default.
	//
	// example:
	//
	// demo_svc
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// The threshold of the metric that triggers auto scaling.
	//
	// 	- If you set metricName to QPS, scale-out is triggered when the average QPS for a single instance is greater than this threshold.
	//
	// 	- If you set metricName to CPU, scale-out is triggered when the average CPU utilization for a single instance is greater than this threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s UpdateServiceAutoScalerRequestScaleStrategies) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestScaleStrategies) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) SetMetricName(v string) *UpdateServiceAutoScalerRequestScaleStrategies {
	s.MetricName = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) SetService(v string) *UpdateServiceAutoScalerRequestScaleStrategies {
	s.Service = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestScaleStrategies) SetThreshold(v float32) *UpdateServiceAutoScalerRequestScaleStrategies {
	s.Threshold = &v
	return s
}

type UpdateServiceAutoScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Update auto scale for service [foo] successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceAutoScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerResponseBody) SetMessage(v string) *UpdateServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceAutoScalerResponseBody) SetRequestId(v string) *UpdateServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceAutoScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerResponse) SetHeaders(v map[string]*string) *UpdateServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceAutoScalerResponse) SetStatusCode(v int32) *UpdateServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceAutoScalerResponse) SetBody(v *UpdateServiceAutoScalerResponseBody) *UpdateServiceAutoScalerResponse {
	s.Body = v
	return s
}

type UpdateServiceCronScalerRequest struct {
	// The points in time that are excluded when you schedule a CronHPA job. The points in time must be specified by using a cron expression.
	ExcludeDates []*string `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	// The description of the CronHPA job.
	//
	// This parameter is required.
	ScaleJobs []*UpdateServiceCronScalerRequestScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
}

func (s UpdateServiceCronScalerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceCronScalerRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerRequest) SetExcludeDates(v []*string) *UpdateServiceCronScalerRequest {
	s.ExcludeDates = v
	return s
}

func (s *UpdateServiceCronScalerRequest) SetScaleJobs(v []*UpdateServiceCronScalerRequestScaleJobs) *UpdateServiceCronScalerRequest {
	s.ScaleJobs = v
	return s
}

type UpdateServiceCronScalerRequestScaleJobs struct {
	// The name of the CronHPA job.
	//
	// example:
	//
	// scale-job-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cron expression that is used to configure the execution time of the CronHPA job. For more information about how to configure cron expressions, see **Description of special characters*	- in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 18 	- 	- 	- *
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The number of instances that you want to configure for the CronHPA job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	TargetSize *int32 `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s UpdateServiceCronScalerRequestScaleJobs) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceCronScalerRequestScaleJobs) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetName(v string) *UpdateServiceCronScalerRequestScaleJobs {
	s.Name = &v
	return s
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetSchedule(v string) *UpdateServiceCronScalerRequestScaleJobs {
	s.Schedule = &v
	return s
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetTargetSize(v int32) *UpdateServiceCronScalerRequestScaleJobs {
	s.TargetSize = &v
	return s
}

type UpdateServiceCronScalerResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Update cron scaler for service [foo] successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceCronScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerResponseBody) SetMessage(v string) *UpdateServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceCronScalerResponseBody) SetRequestId(v string) *UpdateServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceCronScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerResponse) SetHeaders(v map[string]*string) *UpdateServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceCronScalerResponse) SetStatusCode(v int32) *UpdateServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceCronScalerResponse) SetBody(v *UpdateServiceCronScalerResponseBody) *UpdateServiceCronScalerResponse {
	s.Body = v
	return s
}

type UpdateServiceInstanceRequest struct {
	// Specifies whether to isolate the service instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Isolate *bool `json:"Isolate,omitempty" xml:"Isolate,omitempty"`
}

func (s UpdateServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceRequest) SetIsolate(v bool) *UpdateServiceInstanceRequest {
	s.Isolate = &v
	return s
}

type UpdateServiceInstanceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceResponseBody) SetMessage(v string) *UpdateServiceInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceInstanceResponseBody) SetRequestId(v string) *UpdateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceResponse) SetStatusCode(v int32) *UpdateServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceResponse) SetBody(v *UpdateServiceInstanceResponseBody) *UpdateServiceInstanceResponse {
	s.Body = v
	return s
}

type UpdateServiceLabelRequest struct {
	// The custom service tags.
	//
	// This parameter is required.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateServiceLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceLabelRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceLabelRequest) SetLabels(v map[string]*string) *UpdateServiceLabelRequest {
	s.Labels = v
	return s
}

type UpdateServiceLabelResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to update service [service_from_XXXX] labels.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceLabelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceLabelResponseBody) SetMessage(v string) *UpdateServiceLabelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceLabelResponseBody) SetRequestId(v string) *UpdateServiceLabelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceLabelResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceLabelResponse) SetHeaders(v map[string]*string) *UpdateServiceLabelResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceLabelResponse) SetStatusCode(v int32) *UpdateServiceLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceLabelResponse) SetBody(v *UpdateServiceLabelResponseBody) *UpdateServiceLabelResponse {
	s.Body = v
	return s
}

type UpdateServiceMirrorRequest struct {
	// The percentage of traffic that you want to mirror. Valid values: 0 to 100.
	//
	// example:
	//
	// 30
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The service instances.
	Target []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s UpdateServiceMirrorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMirrorRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorRequest) SetRatio(v int32) *UpdateServiceMirrorRequest {
	s.Ratio = &v
	return s
}

func (s *UpdateServiceMirrorRequest) SetTarget(v []*string) *UpdateServiceMirrorRequest {
	s.Target = v
	return s
}

type UpdateServiceMirrorResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Traffic mirroring is updating for service [foo] in region [cn-shanghia], ratio [70%]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceMirrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorResponseBody) SetMessage(v string) *UpdateServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceMirrorResponseBody) SetRequestId(v string) *UpdateServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceMirrorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorResponse) SetHeaders(v map[string]*string) *UpdateServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceMirrorResponse) SetStatusCode(v int32) *UpdateServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceMirrorResponse) SetBody(v *UpdateServiceMirrorResponseBody) *UpdateServiceMirrorResponse {
	s.Body = v
	return s
}

type UpdateServiceSafetyLockRequest struct {
	// The lock scope. Valid values:
	//
	// 	- all: locks all operations.
	//
	// 	- dangerous: locks dangerous operations such as delete and stop operations.
	//
	// 	- none: locks no operations.
	//
	// This parameter is required.
	//
	// example:
	//
	// dangerous
	Lock *string `json:"Lock,omitempty" xml:"Lock,omitempty"`
}

func (s UpdateServiceSafetyLockRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceSafetyLockRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceSafetyLockRequest) SetLock(v string) *UpdateServiceSafetyLockRequest {
	s.Lock = &v
	return s
}

type UpdateServiceSafetyLockResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// service safety lock updated to dangerous
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E089D584-B6F4-50C4-9902-DA2295B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceSafetyLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceSafetyLockResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceSafetyLockResponseBody) SetMessage(v string) *UpdateServiceSafetyLockResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceSafetyLockResponseBody) SetRequestId(v string) *UpdateServiceSafetyLockResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceSafetyLockResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceSafetyLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceSafetyLockResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceSafetyLockResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceSafetyLockResponse) SetHeaders(v map[string]*string) *UpdateServiceSafetyLockResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceSafetyLockResponse) SetStatusCode(v int32) *UpdateServiceSafetyLockResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceSafetyLockResponse) SetBody(v *UpdateServiceSafetyLockResponseBody) *UpdateServiceSafetyLockResponse {
	s.Body = v
	return s
}

type UpdateServiceVersionRequest struct {
	// The destination version of the service. The value must be of the INT type. The value must be greater than 0 and smaller than the current version of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateServiceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionRequest) SetVersion(v int32) *UpdateServiceVersionRequest {
	s.Version = &v
	return s
}

type UpdateServiceVersionResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Service [foo] in region [cn-shanghai] is starting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionResponseBody) SetMessage(v string) *UpdateServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceVersionResponseBody) SetRequestId(v string) *UpdateServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionResponse) SetHeaders(v map[string]*string) *UpdateServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceVersionResponse) SetStatusCode(v int32) *UpdateServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceVersionResponse) SetBody(v *UpdateServiceVersionResponseBody) *UpdateServiceVersionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":            tea.String("pai-eas.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("pai-eas.cn-zhangjiakou.aliyuncs.com"),
		"cn-hangzhou":           tea.String("pai-eas.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("pai-eas.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("pai-eas.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":           tea.String("pai-eas.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1":        tea.String("pai-eas.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("pai-eas.ap-southeast-5.aliyuncs.com"),
		"us-east-1":             tea.String("pai-eas.us-east-1.aliyuncs.com"),
		"us-west-1":             tea.String("pai-eas.us-west-1.aliyuncs.com"),
		"eu-central-1":          tea.String("pai-eas.eu-central-1.aliyuncs.com"),
		"ap-south-1":            tea.String("pai-eas.ap-south-1.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("pai-eas.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("pai-eas.cn-north-2-gov-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("pai-eas.cn-chengdu.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones a service.
//
// @param tmpReq - CloneServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneServiceResponse
func (client *Client) CloneServiceWithOptions(ClusterId *string, ServiceName *string, tmpReq *CloneServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneServiceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CloneServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Labels)) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, tea.String("Labels"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelsShrink)) {
		query["Labels"] = request.LabelsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CloneService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones a service.
//
// @param request - CloneServiceRequest
//
// @return CloneServiceResponse
func (client *Client) CloneService(ClusterId *string, ServiceName *string, request *CloneServiceRequest) (_result *CloneServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneServiceResponse{}
	_body, _err := client.CloneServiceWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Commits the Worker0 container in the custom container service and deploys the container as a new image.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitServiceResponse
func (client *Client) CommitServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CommitServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CommitService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/commit"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CommitServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Commits the Worker0 container in the custom container service and deploys the container as a new image.
//
// @return CommitServiceResponse
func (client *Client) CommitService(ClusterId *string, ServiceName *string) (_result *CommitServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CommitServiceResponse{}
	_body, _err := client.CommitServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) for a private gateway. The IP CIDR blocks added to the ACL can access the private gateway.
//
// @param tmpReq - CreateAclPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAclPolicyResponse
func (client *Client) CreateAclPolicyWithOptions(ClusterId *string, GatewayId *string, tmpReq *CreateAclPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAclPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAclPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AclPolicyList)) {
		request.AclPolicyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AclPolicyList, tea.String("AclPolicyList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclPolicyListShrink)) {
		query["AclPolicyList"] = request.AclPolicyListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAclPolicy"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId)) + "/acl_policy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAclPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) for a private gateway. The IP CIDR blocks added to the ACL can access the private gateway.
//
// @param request - CreateAclPolicyRequest
//
// @return CreateAclPolicyResponse
func (client *Client) CreateAclPolicy(ClusterId *string, GatewayId *string, request *CreateAclPolicyRequest) (_result *CreateAclPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAclPolicyResponse{}
	_body, _err := client.CreateAclPolicyWithOptions(ClusterId, GatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application service.
//
// @param request - CreateAppServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppServiceResponse
func (client *Client) CreateAppServiceWithOptions(request *CreateAppServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QuotaId)) {
		query["QuotaId"] = request.QuotaId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		body["Replicas"] = request.Replicas
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceSpec)) {
		body["ServiceSpec"] = request.ServiceSpec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/app_services"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application service.
//
// @param request - CreateAppServiceRequest
//
// @return CreateAppServiceResponse
func (client *Client) CreateAppService(request *CreateAppServiceRequest) (_result *CreateAppServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppServiceResponse{}
	_body, _err := client.CreateAppServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a stress testing task.
//
// @param request - CreateBenchmarkTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBenchmarkTaskResponse
func (client *Client) CreateBenchmarkTaskWithOptions(request *CreateBenchmarkTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBenchmarkTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a stress testing task.
//
// @param request - CreateBenchmarkTaskRequest
//
// @return CreateBenchmarkTaskResponse
func (client *Client) CreateBenchmarkTask(request *CreateBenchmarkTaskRequest) (_result *CreateBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBenchmarkTaskResponse{}
	_body, _err := client.CreateBenchmarkTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a private gateway. You can create a private gateway only in a self-managed resource group.
//
// @param request - CreateGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayResponse
func (client *Client) CreateGatewayWithOptions(request *CreateGatewayRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewal)) {
		body["AutoRenewal"] = request.AutoRenewal
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.EnableInternet)) {
		body["EnableInternet"] = request.EnableInternet
	}

	if !tea.BoolValue(util.IsUnset(request.EnableIntranet)) {
		body["EnableIntranet"] = request.EnableIntranet
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		body["Replicas"] = request.Replicas
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGateway"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a private gateway. You can create a private gateway only in a self-managed resource group.
//
// @param request - CreateGatewayRequest
//
// @return CreateGatewayResponse
func (client *Client) CreateGateway(request *CreateGatewayRequest) (_result *CreateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGatewayResponse{}
	_body, _err := client.CreateGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an internal endpoint of a private gateway.
//
// @param request - CreateGatewayIntranetLinkedVpcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayIntranetLinkedVpcResponse
func (client *Client) CreateGatewayIntranetLinkedVpcWithOptions(ClusterId *string, GatewayId *string, request *CreateGatewayIntranetLinkedVpcRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGatewayIntranetLinkedVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGatewayIntranetLinkedVpc"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId)) + "/intranet_endpoint_linked_vpc"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an internal endpoint of a private gateway.
//
// @param request - CreateGatewayIntranetLinkedVpcRequest
//
// @return CreateGatewayIntranetLinkedVpcResponse
func (client *Client) CreateGatewayIntranetLinkedVpc(ClusterId *string, GatewayId *string, request *CreateGatewayIntranetLinkedVpcRequest) (_result *CreateGatewayIntranetLinkedVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.CreateGatewayIntranetLinkedVpcWithOptions(ClusterId, GatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource group.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/144261.html) of Elastic Algorithm Service (EAS).
//
// @param request - CreateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(request *CreateResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewal)) {
		body["AutoRenewal"] = request.AutoRenewal
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceCount)) {
		body["EcsInstanceCount"] = request.EcsInstanceCount
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceType)) {
		body["EcsInstanceType"] = request.EcsInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SelfManagedResourceOptions)) {
		body["SelfManagedResourceOptions"] = request.SelfManagedResourceOptions
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDiskSize)) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.Zone)) {
		body["Zone"] = request.Zone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource group.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/144261.html) of Elastic Algorithm Service (EAS).
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates instances in a dedicated resource group.
//
// @param request - CreateResourceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceInstancesResponse
func (client *Client) CreateResourceInstancesWithOptions(ClusterId *string, ResourceId *string, request *CreateResourceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewal)) {
		body["AutoRenewal"] = request.AutoRenewal
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceCount)) {
		body["EcsInstanceCount"] = request.EcsInstanceCount
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceType)) {
		body["EcsInstanceType"] = request.EcsInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDiskSize)) {
		body["SystemDiskSize"] = request.SystemDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.Zone)) {
		body["Zone"] = request.Zone
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates instances in a dedicated resource group.
//
// @param request - CreateResourceInstancesRequest
//
// @return CreateResourceInstancesResponse
func (client *Client) CreateResourceInstances(ClusterId *string, ResourceId *string, request *CreateResourceInstancesRequest) (_result *CreateResourceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceInstancesResponse{}
	_body, _err := client.CreateResourceInstancesWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the LogShipper feature of Log Service for a resource group.
//
// @param request - CreateResourceLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceLogResponse
func (client *Client) CreateResourceLogWithOptions(ClusterId *string, ResourceId *string, request *CreateResourceLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogStore)) {
		body["LogStore"] = request.LogStore
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceLog"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/log"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the LogShipper feature of Log Service for a resource group.
//
// @param request - CreateResourceLogRequest
//
// @return CreateResourceLogResponse
func (client *Client) CreateResourceLog(ClusterId *string, ResourceId *string, request *CreateResourceLogRequest) (_result *CreateResourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceLogResponse{}
	_body, _err := client.CreateResourceLogWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a model service in Elastic Algorithm Service (EAS).
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/144261.html) of Elastic Algorithm Service (EAS).
//
// @param tmpReq - CreateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithOptions(tmpReq *CreateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Labels)) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, tea.String("Labels"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Develop)) {
		query["Develop"] = request.Develop
	}

	if !tea.BoolValue(util.IsUnset(request.LabelsShrink)) {
		query["Labels"] = request.LabelsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a model service in Elastic Algorithm Service (EAS).
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/144261.html) of Elastic Algorithm Service (EAS).
//
// @param request - CreateServiceRequest
//
// @return CreateServiceResponse
func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the Autoscaler feature and creates an Autoscaler controller for a service.
//
// @param request - CreateServiceAutoScalerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceAutoScalerResponse
func (client *Client) CreateServiceAutoScalerWithOptions(ClusterId *string, ServiceName *string, request *CreateServiceAutoScalerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceAutoScalerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Behavior)) {
		body["behavior"] = request.Behavior
	}

	if !tea.BoolValue(util.IsUnset(request.Max)) {
		body["max"] = request.Max
	}

	if !tea.BoolValue(util.IsUnset(request.Min)) {
		body["min"] = request.Min
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleStrategies)) {
		body["scaleStrategies"] = request.ScaleStrategies
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceAutoScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/autoscaler"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceAutoScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Autoscaler feature and creates an Autoscaler controller for a service.
//
// @param request - CreateServiceAutoScalerRequest
//
// @return CreateServiceAutoScalerResponse
func (client *Client) CreateServiceAutoScaler(ClusterId *string, ServiceName *string, request *CreateServiceAutoScalerRequest) (_result *CreateServiceAutoScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceAutoScalerResponse{}
	_body, _err := client.CreateServiceAutoScalerWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the Cron Horizontal Pod Autoscaler (CronHPA) feature for a service.
//
// @param request - CreateServiceCronScalerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceCronScalerResponse
func (client *Client) CreateServiceCronScalerWithOptions(ClusterId *string, ServiceName *string, request *CreateServiceCronScalerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceCronScalerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeDates)) {
		body["ExcludeDates"] = request.ExcludeDates
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleJobs)) {
		body["ScaleJobs"] = request.ScaleJobs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceCronScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/cronscaler"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceCronScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Cron Horizontal Pod Autoscaler (CronHPA) feature for a service.
//
// @param request - CreateServiceCronScalerRequest
//
// @return CreateServiceCronScalerResponse
func (client *Client) CreateServiceCronScaler(ClusterId *string, ServiceName *string, request *CreateServiceCronScalerRequest) (_result *CreateServiceCronScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceCronScalerResponse{}
	_body, _err := client.CreateServiceCronScalerWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the traffic mirroring feature for a service. After the feature is enabled, requests received by the service can be mirrored to another service.
//
// @param request - CreateServiceMirrorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceMirrorResponse
func (client *Client) CreateServiceMirrorWithOptions(ClusterId *string, ServiceName *string, request *CreateServiceMirrorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceMirrorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ratio)) {
		body["Ratio"] = request.Ratio
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		body["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceMirror"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/mirror"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceMirrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the traffic mirroring feature for a service. After the feature is enabled, requests received by the service can be mirrored to another service.
//
// @param request - CreateServiceMirrorRequest
//
// @return CreateServiceMirrorResponse
func (client *Client) CreateServiceMirror(ClusterId *string, ServiceName *string, request *CreateServiceMirrorRequest) (_result *CreateServiceMirrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceMirrorResponse{}
	_body, _err := client.CreateServiceMirrorWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access control list (ACL) for a private gateway. The IP CIDR block that is deleted from the ACL cannot access the private gateway.
//
// @param tmpReq - DeleteAclPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAclPolicyResponse
func (client *Client) DeleteAclPolicyWithOptions(ClusterId *string, GatewayId *string, tmpReq *DeleteAclPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAclPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteAclPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AclPolicyList)) {
		request.AclPolicyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AclPolicyList, tea.String("AclPolicyList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclPolicyListShrink)) {
		query["AclPolicyList"] = request.AclPolicyListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAclPolicy"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId)) + "/acl_policy"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAclPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control list (ACL) for a private gateway. The IP CIDR block that is deleted from the ACL cannot access the private gateway.
//
// @param request - DeleteAclPolicyRequest
//
// @return DeleteAclPolicyResponse
func (client *Client) DeleteAclPolicy(ClusterId *string, GatewayId *string, request *DeleteAclPolicyRequest) (_result *DeleteAclPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAclPolicyResponse{}
	_body, _err := client.DeleteAclPolicyWithOptions(ClusterId, GatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a stress testing task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBenchmarkTaskResponse
func (client *Client) DeleteBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBenchmarkTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(TaskName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stress testing task.
//
// @return DeleteBenchmarkTaskResponse
func (client *Client) DeleteBenchmarkTask(ClusterId *string, TaskName *string) (_result *DeleteBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBenchmarkTaskResponse{}
	_body, _err := client.DeleteBenchmarkTaskWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a private gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithOptions(ClusterId *string, GatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGateway"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a private gateway.
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGateway(ClusterId *string, GatewayId *string) (_result *DeleteGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(ClusterId, GatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关内网访问端点
//
// @param request - DeleteGatewayIntranetLinkedVpcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayIntranetLinkedVpcResponse
func (client *Client) DeleteGatewayIntranetLinkedVpcWithOptions(ClusterId *string, GatewayId *string, request *DeleteGatewayIntranetLinkedVpcRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGatewayIntranetLinkedVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGatewayIntranetLinkedVpc"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId)) + "/intranet_endpoint_linked_vpc"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关内网访问端点
//
// @param request - DeleteGatewayIntranetLinkedVpcRequest
//
// @return DeleteGatewayIntranetLinkedVpcResponse
func (client *Client) DeleteGatewayIntranetLinkedVpc(ClusterId *string, GatewayId *string, request *DeleteGatewayIntranetLinkedVpcRequest) (_result *DeleteGatewayIntranetLinkedVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.DeleteGatewayIntranetLinkedVpcWithOptions(ClusterId, GatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource group that contains no resources or instances.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResource"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource group that contains no resources or instances.
//
// @return DeleteResourceResponse
func (client *Client) DeleteResource(ClusterId *string, ResourceId *string) (_result *DeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the virtual private cloud (VPC) direct connection feature for a dedicated resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceDLinkResponse
func (client *Client) DeleteResourceDLinkWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceDLinkResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceDLink"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/dlink"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceDLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the virtual private cloud (VPC) direct connection feature for a dedicated resource group.
//
// @return DeleteResourceDLinkResponse
func (client *Client) DeleteResourceDLink(ClusterId *string, ResourceId *string) (_result *DeleteResourceDLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceDLinkResponse{}
	_body, _err := client.DeleteResourceDLinkWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes instances in a dedicated resource group. You can delete only pay-as-you-go instances as a regular user.
//
// @param request - DeleteResourceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceInstancesResponse
func (client *Client) DeleteResourceInstancesWithOptions(ClusterId *string, ResourceId *string, request *DeleteResourceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllFailed)) {
		query["AllFailed"] = request.AllFailed
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceList)) {
		query["InstanceList"] = request.InstanceList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/instances"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes instances in a dedicated resource group. You can delete only pay-as-you-go instances as a regular user.
//
// @param request - DeleteResourceInstancesRequest
//
// @return DeleteResourceInstancesResponse
func (client *Client) DeleteResourceInstances(ClusterId *string, ResourceId *string, request *DeleteResourceInstancesRequest) (_result *DeleteResourceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceInstancesResponse{}
	_body, _err := client.DeleteResourceInstancesWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the LogShipper feature of Log Service for a dedicated resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceLogResponse
func (client *Client) DeleteResourceLogWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceLogResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceLog"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/log"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the LogShipper feature of Log Service for a dedicated resource group.
//
// @return DeleteResourceLogResponse
func (client *Client) DeleteResourceLog(ClusterId *string, ResourceId *string) (_result *DeleteResourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceLogResponse{}
	_body, _err := client.DeleteResourceLogWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service.
//
// @return DeleteServiceResponse
func (client *Client) DeleteService(ClusterId *string, ServiceName *string) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the existing Autoscaler controller and disables the Autoscaler feature for a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceAutoScalerResponse
func (client *Client) DeleteServiceAutoScalerWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceAutoScalerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceAutoScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/autoscaler"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceAutoScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the existing Autoscaler controller and disables the Autoscaler feature for a service.
//
// @return DeleteServiceAutoScalerResponse
func (client *Client) DeleteServiceAutoScaler(ClusterId *string, ServiceName *string) (_result *DeleteServiceAutoScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceAutoScalerResponse{}
	_body, _err := client.DeleteServiceAutoScalerWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the Cronscaler feature for a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceCronScalerResponse
func (client *Client) DeleteServiceCronScalerWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceCronScalerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceCronScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/cronscaler"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceCronScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the Cronscaler feature for a service.
//
// @return DeleteServiceCronScalerResponse
func (client *Client) DeleteServiceCronScaler(ClusterId *string, ServiceName *string) (_result *DeleteServiceCronScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceCronScalerResponse{}
	_body, _err := client.DeleteServiceCronScalerWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts the instances of a service.
//
// @param request - DeleteServiceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstancesWithOptions(ClusterId *string, ServiceName *string, request *DeleteServiceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Container)) {
		query["Container"] = request.Container
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceList)) {
		query["InstanceList"] = request.InstanceList
	}

	if !tea.BoolValue(util.IsUnset(request.SoftRestart)) {
		query["SoftRestart"] = request.SoftRestart
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/instances"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts the instances of a service.
//
// @param request - DeleteServiceInstancesRequest
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstances(ClusterId *string, ServiceName *string, request *DeleteServiceInstancesRequest) (_result *DeleteServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.DeleteServiceInstancesWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes existing service tags.
//
// @param tmpReq - DeleteServiceLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceLabelResponse
func (client *Client) DeleteServiceLabelWithOptions(ClusterId *string, ServiceName *string, tmpReq *DeleteServiceLabelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceLabelResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteServiceLabelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Keys)) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, tea.String("Keys"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeysShrink)) {
		query["Keys"] = request.KeysShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceLabel"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/label"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes existing service tags.
//
// @param request - DeleteServiceLabelRequest
//
// @return DeleteServiceLabelResponse
func (client *Client) DeleteServiceLabel(ClusterId *string, ServiceName *string, request *DeleteServiceLabelRequest) (_result *DeleteServiceLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceLabelResponse{}
	_body, _err := client.DeleteServiceLabelWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the traffic mirroring feature for a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceMirrorResponse
func (client *Client) DeleteServiceMirrorWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceMirrorResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceMirror"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/mirror"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceMirrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the traffic mirroring feature for a service.
//
// @return DeleteServiceMirrorResponse
func (client *Client) DeleteServiceMirror(ClusterId *string, ServiceName *string) (_result *DeleteServiceMirrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceMirrorResponse{}
	_body, _err := client.DeleteServiceMirrorWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about the configurations of a stress testing task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBenchmarkTaskResponse
func (client *Client) DescribeBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeBenchmarkTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(TaskName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about the configurations of a stress testing task.
//
// @return DescribeBenchmarkTaskResponse
func (client *Client) DescribeBenchmarkTask(ClusterId *string, TaskName *string) (_result *DescribeBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeBenchmarkTaskResponse{}
	_body, _err := client.DescribeBenchmarkTaskWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the report of a stress testing task.
//
// @param request - DescribeBenchmarkTaskReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBenchmarkTaskReportResponse
func (client *Client) DescribeBenchmarkTaskReportWithOptions(ClusterId *string, TaskName *string, request *DescribeBenchmarkTaskReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeBenchmarkTaskReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		query["ReportType"] = request.ReportType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBenchmarkTaskReport"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(TaskName)) + "/report"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBenchmarkTaskReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the report of a stress testing task.
//
// @param request - DescribeBenchmarkTaskReportRequest
//
// @return DescribeBenchmarkTaskReportResponse
func (client *Client) DescribeBenchmarkTaskReport(ClusterId *string, TaskName *string, request *DescribeBenchmarkTaskReportRequest) (_result *DescribeBenchmarkTaskReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeBenchmarkTaskReportResponse{}
	_body, _err := client.DescribeBenchmarkTaskReportWithOptions(ClusterId, TaskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a private gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGatewayResponse
func (client *Client) DescribeGatewayWithOptions(ClusterId *string, GatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeGatewayResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGateway"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a private gateway.
//
// @return DescribeGatewayResponse
func (client *Client) DescribeGateway(ClusterId *string, GatewayId *string) (_result *DescribeGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeGatewayResponse{}
	_body, _err := client.DescribeGatewayWithOptions(ClusterId, GatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a service group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupResponse
func (client *Client) DescribeGroupWithOptions(ClusterId *string, GroupName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGroup"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/groups/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GroupName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a service group.
//
// @return DescribeGroupResponse
func (client *Client) DescribeGroup(ClusterId *string, GroupName *string) (_result *DescribeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeGroupResponse{}
	_body, _err := client.DescribeGroupWithOptions(ClusterId, GroupName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceResponse
func (client *Client) DescribeResourceWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeResourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResource"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a resource group.
//
// @return DescribeResourceResponse
func (client *Client) DescribeResource(ClusterId *string, ResourceId *string) (_result *DescribeResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourceResponse{}
	_body, _err := client.DescribeResourceWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries detailed configurations about a virtual private cloud (VPC) direct connection of a dedicated resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceDLinkResponse
func (client *Client) DescribeResourceDLinkWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeResourceDLinkResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceDLink"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/dlink"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceDLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed configurations about a virtual private cloud (VPC) direct connection of a dedicated resource group.
//
// @return DescribeResourceDLinkResponse
func (client *Client) DescribeResourceDLink(ClusterId *string, ResourceId *string) (_result *DescribeResourceDLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourceDLinkResponse{}
	_body, _err := client.DescribeResourceDLinkWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about the LogShipper configurations of Log Service for a dedicated resource group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceLogResponse
func (client *Client) DescribeResourceLogWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeResourceLogResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceLog"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/log"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the LogShipper configurations of Log Service for a dedicated resource group.
//
// @return DescribeResourceLogResponse
func (client *Client) DescribeResourceLog(ClusterId *string, ResourceId *string) (_result *DescribeResourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourceLogResponse{}
	_body, _err := client.DescribeResourceLogWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceResponse
func (client *Client) DescribeServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a service.
//
// @return DescribeServiceResponse
func (client *Client) DescribeService(ClusterId *string, ServiceName *string) (_result *DescribeServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceResponse{}
	_body, _err := client.DescribeServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the Autoscaler configurations of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceAutoScalerResponse
func (client *Client) DescribeServiceAutoScalerWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceAutoScalerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceAutoScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/autoscaler"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceAutoScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the Autoscaler configurations of a service.
//
// @return DescribeServiceAutoScalerResponse
func (client *Client) DescribeServiceAutoScaler(ClusterId *string, ServiceName *string) (_result *DescribeServiceAutoScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceAutoScalerResponse{}
	_body, _err := client.DescribeServiceAutoScalerWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Cron Horizontal Pod Autoscaler (CronHPA) configurations of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceCronScalerResponse
func (client *Client) DescribeServiceCronScalerWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceCronScalerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceCronScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/cronscaler"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceCronScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Cron Horizontal Pod Autoscaler (CronHPA) configurations of a service.
//
// @return DescribeServiceCronScalerResponse
func (client *Client) DescribeServiceCronScaler(ClusterId *string, ServiceName *string) (_result *DescribeServiceCronScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceCronScalerResponse{}
	_body, _err := client.DescribeServiceCronScalerWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnostics details of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceDiagnosisResponse
func (client *Client) DescribeServiceDiagnosisWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceDiagnosisResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceDiagnosis"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/diagnosis"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostics details of a service.
//
// @return DescribeServiceDiagnosisResponse
func (client *Client) DescribeServiceDiagnosis(ClusterId *string, ServiceName *string) (_result *DescribeServiceDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceDiagnosisResponse{}
	_body, _err := client.DescribeServiceDiagnosisWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about recent service deployment events.
//
// @param request - DescribeServiceEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceEventResponse
func (client *Client) DescribeServiceEventWithOptions(ClusterId *string, ServiceName *string, request *DescribeServiceEventRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceEvent"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about recent service deployment events.
//
// @param request - DescribeServiceEventRequest
//
// @return DescribeServiceEventResponse
func (client *Client) DescribeServiceEvent(ClusterId *string, ServiceName *string, request *DescribeServiceEventRequest) (_result *DescribeServiceEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceEventResponse{}
	_body, _err := client.DescribeServiceEventWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnostics details of an instance that runs Elastic Algorithm Service (EAS).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceInstanceDiagnosisResponse
func (client *Client) DescribeServiceInstanceDiagnosisWithOptions(ClusterId *string, ServiceName *string, InstanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceInstanceDiagnosisResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceInstanceDiagnosis"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/diagnosis"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceInstanceDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnostics details of an instance that runs Elastic Algorithm Service (EAS).
//
// @return DescribeServiceInstanceDiagnosisResponse
func (client *Client) DescribeServiceInstanceDiagnosis(ClusterId *string, ServiceName *string, InstanceName *string) (_result *DescribeServiceInstanceDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceInstanceDiagnosisResponse{}
	_body, _err := client.DescribeServiceInstanceDiagnosisWithOptions(ClusterId, ServiceName, InstanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the logs of a service.
//
// @param request - DescribeServiceLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceLogResponse
func (client *Client) DescribeServiceLogWithOptions(ClusterId *string, ServiceName *string, request *DescribeServiceLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerName)) {
		query["ContainerName"] = request.ContainerName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Previous)) {
		query["Previous"] = request.Previous
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceLog"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the logs of a service.
//
// @param request - DescribeServiceLogRequest
//
// @return DescribeServiceLogResponse
func (client *Client) DescribeServiceLog(ClusterId *string, ServiceName *string, request *DescribeServiceLogRequest) (_result *DescribeServiceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceLogResponse{}
	_body, _err := client.DescribeServiceLogWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about the traffic mirroring settings of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceMirrorResponse
func (client *Client) DescribeServiceMirrorWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceMirrorResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMirror"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/mirror"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMirrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about the traffic mirroring settings of a service.
//
// @return DescribeServiceMirrorResponse
func (client *Client) DescribeServiceMirror(ClusterId *string, ServiceName *string) (_result *DescribeServiceMirrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceMirrorResponse{}
	_body, _err := client.DescribeServiceMirrorWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical prices of preemptible instances. For more information about preemptible instances, see Create and use preemptible instances.
//
// @param request - DescribeSpotDiscountHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSpotDiscountHistoryResponse
func (client *Client) DescribeSpotDiscountHistoryWithOptions(request *DescribeSpotDiscountHistoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSpotDiscountHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.IsProtect)) {
		query["IsProtect"] = request.IsProtect
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSpotDiscountHistory"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/public/spot_discount"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSpotDiscountHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical prices of preemptible instances. For more information about preemptible instances, see Create and use preemptible instances.
//
// @param request - DescribeSpotDiscountHistoryRequest
//
// @return DescribeSpotDiscountHistoryResponse
func (client *Client) DescribeSpotDiscountHistory(request *DescribeSpotDiscountHistoryRequest) (_result *DescribeSpotDiscountHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSpotDiscountHistoryResponse{}
	_body, _err := client.DescribeSpotDiscountHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches a container service to development mode or exits development mode.
//
// @param request - DevelopServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DevelopServiceResponse
func (client *Client) DevelopServiceWithOptions(ClusterId *string, ServiceName *string, request *DevelopServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DevelopServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Exit)) {
		query["Exit"] = request.Exit
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DevelopService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/develop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DevelopServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches a container service to development mode or exits development mode.
//
// @param request - DevelopServiceRequest
//
// @return DevelopServiceResponse
func (client *Client) DevelopService(ClusterId *string, ServiceName *string, request *DevelopServiceRequest) (_result *DevelopServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DevelopServiceResponse{}
	_body, _err := client.DevelopServiceWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries access control lists (ACLs) created for a private gateway.
//
// @param request - ListAclPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAclPolicyResponse
func (client *Client) ListAclPolicyWithOptions(ClusterId *string, GatewayId *string, request *ListAclPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAclPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAclPolicy"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId)) + "/acl_policy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAclPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access control lists (ACLs) created for a private gateway.
//
// @param request - ListAclPolicyRequest
//
// @return ListAclPolicyResponse
func (client *Client) ListAclPolicy(ClusterId *string, GatewayId *string, request *ListAclPolicyRequest) (_result *ListAclPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAclPolicyResponse{}
	_body, _err := client.ListAclPolicyWithOptions(ClusterId, GatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of stress testing tasks that are created by the current user.
//
// @param request - ListBenchmarkTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBenchmarkTaskResponse
func (client *Client) ListBenchmarkTaskWithOptions(request *ListBenchmarkTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListBenchmarkTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of stress testing tasks that are created by the current user.
//
// @param request - ListBenchmarkTaskRequest
//
// @return ListBenchmarkTaskResponse
func (client *Client) ListBenchmarkTask(request *ListBenchmarkTaskRequest) (_result *ListBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBenchmarkTaskResponse{}
	_body, _err := client.ListBenchmarkTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of private gateways.
//
// @param request - ListGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayResponse
func (client *Client) ListGatewayWithOptions(request *ListGatewayRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GatewayId)) {
		query["GatewayId"] = request.GatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayName)) {
		query["GatewayName"] = request.GatewayName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGateway"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of private gateways.
//
// @param request - ListGatewayRequest
//
// @return ListGatewayResponse
func (client *Client) ListGateway(request *ListGatewayRequest) (_result *ListGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewayResponse{}
	_body, _err := client.ListGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of the internal endpoints of a private gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayIntranetLinkedVpcResponse
func (client *Client) ListGatewayIntranetLinkedVpcWithOptions(ClusterId *string, GatewayId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGatewayIntranetLinkedVpcResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListGatewayIntranetLinkedVpc"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId)) + "/intranet_endpoint_linked_vpc"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of the internal endpoints of a private gateway.
//
// @return ListGatewayIntranetLinkedVpcResponse
func (client *Client) ListGatewayIntranetLinkedVpc(ClusterId *string, GatewayId *string) (_result *ListGatewayIntranetLinkedVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewayIntranetLinkedVpcResponse{}
	_body, _err := client.ListGatewayIntranetLinkedVpcWithOptions(ClusterId, GatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries created service groups.
//
// @param request - ListGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries created service groups.
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of workers in a resource group.
//
// @param request - ListResourceInstanceWorkerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceInstanceWorkerResponse
func (client *Client) ListResourceInstanceWorkerWithOptions(ClusterId *string, ResourceId *string, InstanceName *string, request *ListResourceInstanceWorkerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceInstanceWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceInstanceWorker"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/instance/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/workers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceInstanceWorkerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workers in a resource group.
//
// @param request - ListResourceInstanceWorkerRequest
//
// @return ListResourceInstanceWorkerResponse
func (client *Client) ListResourceInstanceWorker(ClusterId *string, ResourceId *string, InstanceName *string, request *ListResourceInstanceWorkerRequest) (_result *ListResourceInstanceWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceInstanceWorkerResponse{}
	_body, _err := client.ListResourceInstanceWorkerWithOptions(ClusterId, ResourceId, InstanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances in a dedicated resource group.
//
// @param request - ListResourceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceInstancesResponse
func (client *Client) ListResourceInstancesWithOptions(ClusterId *string, ResourceId *string, request *ListResourceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIP)) {
		query["InstanceIP"] = request.InstanceIP
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances in a dedicated resource group.
//
// @param request - ListResourceInstancesRequest
//
// @return ListResourceInstancesResponse
func (client *Client) ListResourceInstances(ClusterId *string, ResourceId *string, request *ListResourceInstancesRequest) (_result *ListResourceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceInstancesResponse{}
	_body, _err := client.ListResourceInstancesWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListResourceServices is deprecated
//
// Summary:
//
// Queries a list of services that are deployed in the dedicated resource group.
//
// @param request - ListResourceServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceServicesResponse
// Deprecated
func (client *Client) ListResourceServicesWithOptions(ClusterId *string, ResourceId *string, request *ListResourceServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceServices"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/services"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListResourceServices is deprecated
//
// Summary:
//
// Queries a list of services that are deployed in the dedicated resource group.
//
// @param request - ListResourceServicesRequest
//
// @return ListResourceServicesResponse
// Deprecated
func (client *Client) ListResourceServices(ClusterId *string, ResourceId *string, request *ListResourceServicesRequest) (_result *ListResourceServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceServicesResponse{}
	_body, _err := client.ListResourceServicesWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of dedicated resource groups for the current user.
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of dedicated resource groups for the current user.
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the containers of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceContainersResponse
func (client *Client) ListServiceContainersWithOptions(ClusterId *string, ServiceName *string, InstanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceContainersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceContainers"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName)) + "/containers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceContainersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the containers of a service.
//
// @return ListServiceContainersResponse
func (client *Client) ListServiceContainers(ClusterId *string, ServiceName *string, InstanceName *string) (_result *ListServiceContainersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceContainersResponse{}
	_body, _err := client.ListServiceContainersWithOptions(ClusterId, ServiceName, InstanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instances of a service.
//
// @param request - ListServiceInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstancesWithOptions(ClusterId *string, ServiceName *string, request *ListServiceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.HostIP)) {
		query["HostIP"] = request.HostIP
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIP)) {
		query["InstanceIP"] = request.InstanceIP
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.IsSpot)) {
		query["IsSpot"] = request.IsSpot
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instances of a service.
//
// @param request - ListServiceInstancesRequest
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstances(ClusterId *string, ServiceName *string, request *ListServiceInstancesRequest) (_result *ListServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.ListServiceInstancesWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the historical versions of a service.
//
// @param request - ListServiceVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceVersionsResponse
func (client *Client) ListServiceVersionsWithOptions(ClusterId *string, ServiceName *string, request *ListServiceVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceVersions"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the historical versions of a service.
//
// @param request - ListServiceVersionsRequest
//
// @return ListServiceVersionsResponse
func (client *Client) ListServiceVersions(ClusterId *string, ServiceName *string, request *ListServiceVersionsRequest) (_result *ListServiceVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceVersionsResponse{}
	_body, _err := client.ListServiceVersionsWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists services.
//
// @param tmpReq - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(tmpReq *ListServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListServicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Label)) {
		request.LabelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Label, tea.String("Label"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.Gateway)) {
		query["Gateway"] = request.Gateway
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.LabelShrink)) {
		query["Label"] = request.LabelShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentServiceUid)) {
		query["ParentServiceUid"] = request.ParentServiceUid
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaId)) {
		query["QuotaId"] = request.QuotaId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceStatus)) {
		query["ServiceStatus"] = request.ServiceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceUid)) {
		query["ServiceUid"] = request.ServiceUid
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists services.
//
// @param request - ListServicesRequest
//
// @return ListServicesResponse
func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tenant plug-ins.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantAddonsResponse
func (client *Client) ListTenantAddonsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTenantAddonsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListTenantAddons"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/tenantaddons"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTenantAddonsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tenant plug-ins.
//
// @return ListTenantAddonsResponse
func (client *Client) ListTenantAddons() (_result *ListTenantAddonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTenantAddonsResponse{}
	_body, _err := client.ListTenantAddonsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a tenant plug-in.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReinstallTenantAddonResponse
func (client *Client) ReinstallTenantAddonWithOptions(ClusterId *string, TenantAddonName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReinstallTenantAddonResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ReinstallTenantAddon"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/tenantaddons/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(TenantAddonName)) + "/reinstall"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReinstallTenantAddonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a tenant plug-in.
//
// @return ReinstallTenantAddonResponse
func (client *Client) ReinstallTenantAddon(ClusterId *string, TenantAddonName *string) (_result *ReinstallTenantAddonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReinstallTenantAddonResponse{}
	_body, _err := client.ReinstallTenantAddonWithOptions(ClusterId, TenantAddonName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs canary release or blue-green release of a service.
//
// @param request - ReleaseServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseServiceResponse
func (client *Client) ReleaseServiceWithOptions(ClusterId *string, ServiceName *string, request *ReleaseServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReleaseServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TrafficState)) {
		body["TrafficState"] = request.TrafficState
	}

	if !tea.BoolValue(util.IsUnset(request.Weight)) {
		body["Weight"] = request.Weight
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/release"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs canary release or blue-green release of a service.
//
// @param request - ReleaseServiceRequest
//
// @return ReleaseServiceResponse
func (client *Client) ReleaseService(ClusterId *string, ServiceName *string, request *ReleaseServiceRequest) (_result *ReleaseServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseServiceResponse{}
	_body, _err := client.ReleaseServiceWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartServiceResponse
func (client *Client) RestartServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RestartService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/restart"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a service.
//
// @return RestartServiceResponse
func (client *Client) RestartService(ClusterId *string, ServiceName *string) (_result *RestartServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartServiceResponse{}
	_body, _err := client.RestartServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a stress testing task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBenchmarkTaskResponse
func (client *Client) StartBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartBenchmarkTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(TaskName)) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a stress testing task.
//
// @return StartBenchmarkTaskResponse
func (client *Client) StartBenchmarkTask(ClusterId *string, TaskName *string) (_result *StartBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartBenchmarkTaskResponse{}
	_body, _err := client.StartBenchmarkTaskWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartServiceResponse
func (client *Client) StartServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a service.
//
// @return StartServiceResponse
func (client *Client) StartService(ClusterId *string, ServiceName *string) (_result *StartServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartServiceResponse{}
	_body, _err := client.StartServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a stress testing task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopBenchmarkTaskResponse
func (client *Client) StopBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopBenchmarkTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(TaskName)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a stress testing task.
//
// @return StopBenchmarkTaskResponse
func (client *Client) StopBenchmarkTask(ClusterId *string, TaskName *string) (_result *StopBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopBenchmarkTaskResponse{}
	_body, _err := client.StopBenchmarkTaskWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a running service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopServiceResponse
func (client *Client) StopServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a running service.
//
// @return StopServiceResponse
func (client *Client) StopService(ClusterId *string, ServiceName *string) (_result *StopServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopServiceResponse{}
	_body, _err := client.StopServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an application service.
//
// @param request - UpdateAppServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppServiceResponse
func (client *Client) UpdateAppServiceWithOptions(ClusterId *string, ServiceName *string, request *UpdateAppServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAppServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QuotaId)) {
		query["QuotaId"] = request.QuotaId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		body["Replicas"] = request.Replicas
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceSpec)) {
		body["ServiceSpec"] = request.ServiceSpec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/app_services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an application service.
//
// @param request - UpdateAppServiceRequest
//
// @return UpdateAppServiceResponse
func (client *Client) UpdateAppService(ClusterId *string, ServiceName *string, request *UpdateAppServiceRequest) (_result *UpdateAppServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAppServiceResponse{}
	_body, _err := client.UpdateAppServiceWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a stress testing task.
//
// @param request - UpdateBenchmarkTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBenchmarkTaskResponse
func (client *Client) UpdateBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, request *UpdateBenchmarkTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateBenchmarkTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(TaskName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a stress testing task.
//
// @param request - UpdateBenchmarkTaskRequest
//
// @return UpdateBenchmarkTaskResponse
func (client *Client) UpdateBenchmarkTask(ClusterId *string, TaskName *string, request *UpdateBenchmarkTaskRequest) (_result *UpdateBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBenchmarkTaskResponse{}
	_body, _err := client.UpdateBenchmarkTaskWithOptions(ClusterId, TaskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update a private gateway.
//
// @param request - UpdateGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayResponse
func (client *Client) UpdateGatewayWithOptions(GatewayId *string, ClusterId *string, request *UpdateGatewayRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableInternet)) {
		body["EnableInternet"] = request.EnableInternet
	}

	if !tea.BoolValue(util.IsUnset(request.EnableIntranet)) {
		body["EnableIntranet"] = request.EnableIntranet
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.IsDefault)) {
		body["IsDefault"] = request.IsDefault
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Replicas)) {
		body["Replicas"] = request.Replicas
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGateway"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/gateways/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(GatewayId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a private gateway.
//
// @param request - UpdateGatewayRequest
//
// @return UpdateGatewayResponse
func (client *Client) UpdateGateway(GatewayId *string, ClusterId *string, request *UpdateGatewayRequest) (_result *UpdateGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayResponse{}
	_body, _err := client.UpdateGatewayWithOptions(GatewayId, ClusterId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a dedicated resource group. Only the name of a dedicated resource group can be updated.
//
// @param request - UpdateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(ClusterId *string, ResourceId *string, request *UpdateResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		body["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.SelfManagedResourceOptions)) {
		body["SelfManagedResourceOptions"] = request.SelfManagedResourceOptions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResource"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a dedicated resource group. Only the name of a dedicated resource group can be updated.
//
// @param request - UpdateResourceRequest
//
// @return UpdateResourceResponse
func (client *Client) UpdateResource(ClusterId *string, ResourceId *string, request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of a virtual private cloud (VPC) direct connection for a dedicated resource group.
//
// @param request - UpdateResourceDLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceDLinkResponse
func (client *Client) UpdateResourceDLinkWithOptions(ClusterId *string, ResourceId *string, request *UpdateResourceDLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceDLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCIDRs)) {
		body["DestinationCIDRs"] = request.DestinationCIDRs
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIdList)) {
		body["VSwitchIdList"] = request.VSwitchIdList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceDLink"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/dlink"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceDLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a virtual private cloud (VPC) direct connection for a dedicated resource group.
//
// @param request - UpdateResourceDLinkRequest
//
// @return UpdateResourceDLinkResponse
func (client *Client) UpdateResourceDLink(ClusterId *string, ResourceId *string, request *UpdateResourceDLinkRequest) (_result *UpdateResourceDLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceDLinkResponse{}
	_body, _err := client.UpdateResourceDLinkWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the service scheduling status of an instance in a dedicated resource group.
//
// @param request - UpdateResourceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceInstanceResponse
func (client *Client) UpdateResourceInstanceWithOptions(ClusterId *string, ResourceId *string, InstanceId *string, request *UpdateResourceInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["Action"] = request.Action
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceInstance"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ResourceId)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the service scheduling status of an instance in a dedicated resource group.
//
// @param request - UpdateResourceInstanceRequest
//
// @return UpdateResourceInstanceResponse
func (client *Client) UpdateResourceInstance(ClusterId *string, ResourceId *string, InstanceId *string, request *UpdateResourceInstanceRequest) (_result *UpdateResourceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceInstanceResponse{}
	_body, _err := client.UpdateResourceInstanceWithOptions(ClusterId, ResourceId, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a model or processor of a service. If only the metadata.instance field is updated, manual scaling can be performed.
//
// @param request - UpdateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateType)) {
		query["UpdateType"] = request.UpdateType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a model or processor of a service. If only the metadata.instance field is updated, manual scaling can be performed.
//
// @param request - UpdateServiceRequest
//
// @return UpdateServiceResponse
func (client *Client) UpdateService(ClusterId *string, ServiceName *string, request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the Autoscaler configurations of a service.
//
// @param request - UpdateServiceAutoScalerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceAutoScalerResponse
func (client *Client) UpdateServiceAutoScalerWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceAutoScalerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceAutoScalerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Behavior)) {
		body["behavior"] = request.Behavior
	}

	if !tea.BoolValue(util.IsUnset(request.Max)) {
		body["max"] = request.Max
	}

	if !tea.BoolValue(util.IsUnset(request.Min)) {
		body["min"] = request.Min
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleStrategies)) {
		body["scaleStrategies"] = request.ScaleStrategies
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceAutoScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/autoscaler"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceAutoScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Autoscaler configurations of a service.
//
// @param request - UpdateServiceAutoScalerRequest
//
// @return UpdateServiceAutoScalerResponse
func (client *Client) UpdateServiceAutoScaler(ClusterId *string, ServiceName *string, request *UpdateServiceAutoScalerRequest) (_result *UpdateServiceAutoScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceAutoScalerResponse{}
	_body, _err := client.UpdateServiceAutoScalerWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the Cron Horizontal Pod Autoscaler (CronHPA) settings of a service.
//
// @param request - UpdateServiceCronScalerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceCronScalerResponse
func (client *Client) UpdateServiceCronScalerWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceCronScalerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceCronScalerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeDates)) {
		body["ExcludeDates"] = request.ExcludeDates
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleJobs)) {
		body["ScaleJobs"] = request.ScaleJobs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceCronScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/cronscaler"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceCronScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Cron Horizontal Pod Autoscaler (CronHPA) settings of a service.
//
// @param request - UpdateServiceCronScalerRequest
//
// @return UpdateServiceCronScalerResponse
func (client *Client) UpdateServiceCronScaler(ClusterId *string, ServiceName *string, request *UpdateServiceCronScalerRequest) (_result *UpdateServiceCronScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceCronScalerResponse{}
	_body, _err := client.UpdateServiceCronScalerWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates attributes of service instances. Only isolation can be performed for service instances.
//
// @param request - UpdateServiceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceResponse
func (client *Client) UpdateServiceInstanceWithOptions(ClusterId *string, ServiceName *string, InstanceName *string, request *UpdateServiceInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Isolate)) {
		body["Isolate"] = request.Isolate
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceInstance"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(InstanceName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates attributes of service instances. Only isolation can be performed for service instances.
//
// @param request - UpdateServiceInstanceRequest
//
// @return UpdateServiceInstanceResponse
func (client *Client) UpdateServiceInstance(ClusterId *string, ServiceName *string, InstanceName *string, request *UpdateServiceInstanceRequest) (_result *UpdateServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceInstanceResponse{}
	_body, _err := client.UpdateServiceInstanceWithOptions(ClusterId, ServiceName, InstanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds service tags or updates existing service tags.
//
// @param request - UpdateServiceLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceLabelResponse
func (client *Client) UpdateServiceLabelWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceLabelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceLabel"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/label"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds service tags or updates existing service tags.
//
// @param request - UpdateServiceLabelRequest
//
// @return UpdateServiceLabelResponse
func (client *Client) UpdateServiceLabel(ClusterId *string, ServiceName *string, request *UpdateServiceLabelRequest) (_result *UpdateServiceLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceLabelResponse{}
	_body, _err := client.UpdateServiceLabelWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the traffic mirroring configurations of a service.
//
// @param request - UpdateServiceMirrorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceMirrorResponse
func (client *Client) UpdateServiceMirrorWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceMirrorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceMirrorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ratio)) {
		body["Ratio"] = request.Ratio
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		body["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceMirror"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/mirror"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceMirrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the traffic mirroring configurations of a service.
//
// @param request - UpdateServiceMirrorRequest
//
// @return UpdateServiceMirrorResponse
func (client *Client) UpdateServiceMirror(ClusterId *string, ServiceName *string, request *UpdateServiceMirrorRequest) (_result *UpdateServiceMirrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceMirrorResponse{}
	_body, _err := client.UpdateServiceMirrorWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the safety lock of a service to minimize misoperations on the service.
//
// @param request - UpdateServiceSafetyLockRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceSafetyLockResponse
func (client *Client) UpdateServiceSafetyLockWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceSafetyLockRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceSafetyLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lock)) {
		body["Lock"] = request.Lock
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceSafetyLock"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/lock"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceSafetyLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the safety lock of a service to minimize misoperations on the service.
//
// @param request - UpdateServiceSafetyLockRequest
//
// @return UpdateServiceSafetyLockResponse
func (client *Client) UpdateServiceSafetyLock(ClusterId *string, ServiceName *string, request *UpdateServiceSafetyLockRequest) (_result *UpdateServiceSafetyLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceSafetyLockResponse{}
	_body, _err := client.UpdateServiceSafetyLockWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the version of a service or rolls back the service to a specific version.
//
// @param request - UpdateServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceVersionResponse
func (client *Client) UpdateServiceVersionWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceVersion"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(openapiutil.GetEncodeParam(ClusterId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceName)) + "/version"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version of a service or rolls back the service to a specific version.
//
// @param request - UpdateServiceVersionRequest
//
// @return UpdateServiceVersionResponse
func (client *Client) UpdateServiceVersion(ClusterId *string, ServiceName *string, request *UpdateServiceVersionRequest) (_result *UpdateServiceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceVersionResponse{}
	_body, _err := client.UpdateServiceVersionWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
