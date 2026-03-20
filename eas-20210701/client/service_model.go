// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iService interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *Service
	GetAccessToken() *string
	SetAppConfig(v string) *Service
	GetAppConfig() *string
	SetAppSpecName(v string) *Service
	GetAppSpecName() *string
	SetAppType(v string) *Service
	GetAppType() *string
	SetAppVersion(v string) *Service
	GetAppVersion() *string
	SetAutoscalerEnabled(v bool) *Service
	GetAutoscalerEnabled() *bool
	SetCallerUid(v string) *Service
	GetCallerUid() *string
	SetCpu(v int32) *Service
	GetCpu() *int32
	SetCreateTime(v string) *Service
	GetCreateTime() *string
	SetCronscalerEnabled(v bool) *Service
	GetCronscalerEnabled() *bool
	SetCurrentVersion(v int32) *Service
	GetCurrentVersion() *int32
	SetExtraData(v string) *Service
	GetExtraData() *string
	SetGPUCorePercentage(v int32) *Service
	GetGPUCorePercentage() *int32
	SetGPUMemory(v int32) *Service
	GetGPUMemory() *int32
	SetGateway(v string) *Service
	GetGateway() *string
	SetGpu(v int32) *Service
	GetGpu() *int32
	SetImage(v string) *Service
	GetImage() *string
	SetInstanceCountInResource(v *ServiceInstanceCountInResource) *Service
	GetInstanceCountInResource() *ServiceInstanceCountInResource
	SetInternetEndpoint(v string) *Service
	GetInternetEndpoint() *string
	SetIntranetEndpoint(v string) *Service
	GetIntranetEndpoint() *string
	SetLabels(v []*ServiceLabels) *Service
	GetLabels() []*ServiceLabels
	SetLatestVersion(v int32) *Service
	GetLatestVersion() *int32
	SetMemory(v int32) *Service
	GetMemory() *int32
	SetMessage(v string) *Service
	GetMessage() *string
	SetNamespace(v string) *Service
	GetNamespace() *string
	SetParentUid(v string) *Service
	GetParentUid() *string
	SetPendingInstance(v int32) *Service
	GetPendingInstance() *int32
	SetQuotaId(v string) *Service
	GetQuotaId() *string
	SetReason(v string) *Service
	GetReason() *string
	SetRegion(v string) *Service
	GetRegion() *string
	SetRequestId(v string) *Service
	GetRequestId() *string
	SetResource(v string) *Service
	GetResource() *string
	SetResourceAlias(v string) *Service
	GetResourceAlias() *string
	SetResourceBurstable(v bool) *Service
	GetResourceBurstable() *bool
	SetRole(v string) *Service
	GetRole() *string
	SetRoleAttrs(v string) *Service
	GetRoleAttrs() *string
	SetRunningInstance(v int32) *Service
	GetRunningInstance() *int32
	SetSafetyLock(v string) *Service
	GetSafetyLock() *string
	SetSecondaryInternetEndpoint(v string) *Service
	GetSecondaryInternetEndpoint() *string
	SetSecondaryIntranetEndpoint(v string) *Service
	GetSecondaryIntranetEndpoint() *string
	SetServiceConfig(v string) *Service
	GetServiceConfig() *string
	SetServiceGroup(v string) *Service
	GetServiceGroup() *string
	SetServiceId(v string) *Service
	GetServiceId() *string
	SetServiceName(v string) *Service
	GetServiceName() *string
	SetServiceUid(v string) *Service
	GetServiceUid() *string
	SetSource(v string) *Service
	GetSource() *string
	SetStatus(v string) *Service
	GetStatus() *string
	SetTotalInstance(v int32) *Service
	GetTotalInstance() *int32
	SetTrafficState(v string) *Service
	GetTrafficState() *string
	SetUpdateTime(v string) *Service
	GetUpdateTime() *string
	SetWeight(v int32) *Service
	GetWeight() *int32
	SetWorkspaceId(v string) *Service
	GetWorkspaceId() *string
}

type Service struct {
	// The token that is used to access the service.
	//
	// example:
	//
	// MzJiMDI5MDliODc0MTlkYmI0ZDhlYmExYjczYTIyZTE3Zm********
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The application service configuration.
	//
	// example:
	//
	// {"ModelStorage":"oss"}
	AppConfig *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The name of the application service specification.
	//
	// example:
	//
	// llama_7b_fp16
	AppSpecName *string `json:"AppSpecName,omitempty" xml:"AppSpecName,omitempty"`
	// The application service type.
	//
	// example:
	//
	// LLM
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The application service version.
	//
	// example:
	//
	// v1
	AppVersion        *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AutoscalerEnabled *bool   `json:"AutoscalerEnabled,omitempty" xml:"AutoscalerEnabled,omitempty"`
	// The user ID (UID) of the Alibaba Cloud account that is used to create the service.
	//
	// example:
	//
	// 20123*******
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The number of CPU cores that you applied for each instance.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the service was created. The time is displayed in the UTC RFC3339 format.
	//
	// example:
	//
	// 2021-01-29T11:13:20Z
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CronscalerEnabled *bool   `json:"CronscalerEnabled,omitempty" xml:"CronscalerEnabled,omitempty"`
	// The version of the model that is running.
	//
	// example:
	//
	// 1
	CurrentVersion *int32 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The additional information about the service.
	//
	// example:
	//
	// {\\"blue_green_services\\":[\\"test\\",\\"testxxxx\\"]}
	ExtraData         *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	GPUCorePercentage *int32  `json:"GPUCorePercentage,omitempty" xml:"GPUCorePercentage,omitempty"`
	GPUMemory         *int32  `json:"GPUMemory,omitempty" xml:"GPUMemory,omitempty"`
	// The ID of the dedicated gateway for the service. This parameter is available only for services that are associated with dedicated gateways.
	//
	// example:
	//
	// gw-xxxxxx
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The number of GPUs that you applied for each instance.
	//
	// example:
	//
	// 0
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The data image of the service.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/eas/echo_cn-shanghai:v0.0.1-20210129111320
	Image                   *string                         `json:"Image,omitempty" xml:"Image,omitempty"`
	InstanceCountInResource *ServiceInstanceCountInResource `json:"InstanceCountInResource,omitempty" xml:"InstanceCountInResource,omitempty" type:"Struct"`
	// The public endpoint of the service. This parameter is returned only in the DescribeService API operation.
	//
	// example:
	//
	// http://10123*****.cn-shanghai.aliyuncs.com/api/predict/echo
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The internal endpoint of the service. This parameter is returned only in the DescribeService API operation.
	//
	// example:
	//
	// http://10123*****.vpc.cn-shanghai.aliyuncs.com/api/predict/echo
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// The labels.
	Labels []*ServiceLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest version of the service.
	//
	// example:
	//
	// 1
	LatestVersion *int32 `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The memory size that you applied for each instance. Unit: MB.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The service summary.
	//
	// example:
	//
	// Service start successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The namespace in which the service resides.
	//
	// example:
	//
	// echo
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The UID of the Alibaba Cloud account that is used to create the service.
	//
	// example:
	//
	// 11234*******
	ParentUid *string `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	// The number of instances for the pending service.
	//
	// example:
	//
	// 0
	PendingInstance *int32 `json:"PendingInstance,omitempty" xml:"PendingInstance,omitempty"`
	// The quota ID for the service. This parameter is available only for services deployed by using Lingjun resource quotas.
	//
	// example:
	//
	// quotaxxxxx
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The reason for which the service is in the current state.
	//
	// example:
	//
	// RUNNING
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region in which the service resides.
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
	// The resource group to which the service belongs.
	//
	// example:
	//
	// eas-r-xxxxxxx
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The alias of the resource group to which the service belongs.
	//
	// example:
	//
	// my_resource
	ResourceAlias     *string `json:"ResourceAlias,omitempty" xml:"ResourceAlias,omitempty"`
	ResourceBurstable *bool   `json:"ResourceBurstable,omitempty" xml:"ResourceBurstable,omitempty"`
	// The service role.
	//
	// example:
	//
	// Queue
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The additional attributes of the service role. This parameter is returned only in the DescribeService API operation.
	//
	// example:
	//
	// "{\\"ApproxMaxLength\\":null,\\"Length\\":null,\\"MaxPayloadBytes\\":null}"
	RoleAttrs *string `json:"RoleAttrs,omitempty" xml:"RoleAttrs,omitempty"`
	// The number of instances for the running service.
	//
	// example:
	//
	// 1
	RunningInstance *int32 `json:"RunningInstance,omitempty" xml:"RunningInstance,omitempty"`
	// The security lock of the service.
	//
	// Valid values:
	//
	// 	- all: forbids all operations.
	//
	// 	- dangerous: forbids the operation of deleting or stopping the service.
	//
	// 	- none: forbids no operations.
	//
	// example:
	//
	// dangerous
	SafetyLock *string `json:"SafetyLock,omitempty" xml:"SafetyLock,omitempty"`
	// The public endpoint that is used in the asynchronization request of the service. This parameter is returned only in the DescribeService API operation.
	//
	// example:
	//
	// http://10123*****.cn-shanghai.aliyuncs.com/api/predict/async_path.echo
	SecondaryInternetEndpoint *string `json:"SecondaryInternetEndpoint,omitempty" xml:"SecondaryInternetEndpoint,omitempty"`
	// The internal endpoint that is used in the asynchronization request of the service. This parameter is returned only in the DescribeService API operation.
	//
	// example:
	//
	// http://10123*****.vpc.cn-shanghai.aliyuncs.com/api/predict/async_path.echo
	SecondaryIntranetEndpoint *string `json:"SecondaryIntranetEndpoint,omitempty" xml:"SecondaryIntranetEndpoint,omitempty"`
	// The service configurations.
	//
	// example:
	//
	// {        "metadata": {             "cpu":1,             "instance":1,             "memory":1024           },         "name":"echo",         "processor_entry":"libecho.so",         "processor_path":"http://oss-cn-hangzhou-zmf.aliyuncs.com/059247/echo_processor_release.tar.gz",         "processor_type":"cpp"     }
	ServiceConfig *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	// The group to which the service belongs.
	//
	// example:
	//
	// my_group
	ServiceGroup *string `json:"ServiceGroup,omitempty" xml:"ServiceGroup,omitempty"`
	// The unique ID of the service.
	//
	// example:
	//
	// eas-m-xxasdat
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// echo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service ID. ServiceUid has the same meaning as ServiceId, and the values of the two parameters are the same.
	//
	// example:
	//
	// eas-m-xxasdat
	ServiceUid *string `json:"ServiceUid,omitempty" xml:"ServiceUid,omitempty"`
	// The source from which the service deployment request is initiated.
	//
	// example:
	//
	// dsw
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The service status.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- Deploying
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Updating
	//
	// 	- Stopping
	//
	// 	- Waiting
	//
	// 	- HotUpdate
	//
	// 	- Starting
	//
	// 	- DeleteFailed
	//
	// 	- Running
	//
	// 	- Scaling
	//
	// 	- Pending
	//
	// 	- Deleting
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of instances for the service.
	//
	// example:
	//
	// 1
	TotalInstance *int32 `json:"TotalInstance,omitempty" xml:"TotalInstance,omitempty"`
	// The traffic state.
	//
	// Valid values:
	//
	// 	- standalone: independent traffic.
	//
	// 	- grouping: grouped traffic.
	//
	// example:
	//
	// standalone
	TrafficState *string `json:"TrafficState,omitempty" xml:"TrafficState,omitempty"`
	// The time when the service was updated. The time is displayed in the UTC RFC3339 format.
	//
	// example:
	//
	// 2021-01-29T11:13:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The weight of the service in canary release.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The ID of the workspace to which the service belongs.
	//
	// example:
	//
	// 123445
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Service) String() string {
	return dara.Prettify(s)
}

func (s Service) GoString() string {
	return s.String()
}

func (s *Service) GetAccessToken() *string {
	return s.AccessToken
}

func (s *Service) GetAppConfig() *string {
	return s.AppConfig
}

func (s *Service) GetAppSpecName() *string {
	return s.AppSpecName
}

func (s *Service) GetAppType() *string {
	return s.AppType
}

func (s *Service) GetAppVersion() *string {
	return s.AppVersion
}

func (s *Service) GetAutoscalerEnabled() *bool {
	return s.AutoscalerEnabled
}

func (s *Service) GetCallerUid() *string {
	return s.CallerUid
}

func (s *Service) GetCpu() *int32 {
	return s.Cpu
}

func (s *Service) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Service) GetCronscalerEnabled() *bool {
	return s.CronscalerEnabled
}

func (s *Service) GetCurrentVersion() *int32 {
	return s.CurrentVersion
}

func (s *Service) GetExtraData() *string {
	return s.ExtraData
}

func (s *Service) GetGPUCorePercentage() *int32 {
	return s.GPUCorePercentage
}

func (s *Service) GetGPUMemory() *int32 {
	return s.GPUMemory
}

func (s *Service) GetGateway() *string {
	return s.Gateway
}

func (s *Service) GetGpu() *int32 {
	return s.Gpu
}

func (s *Service) GetImage() *string {
	return s.Image
}

func (s *Service) GetInstanceCountInResource() *ServiceInstanceCountInResource {
	return s.InstanceCountInResource
}

func (s *Service) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *Service) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *Service) GetLabels() []*ServiceLabels {
	return s.Labels
}

func (s *Service) GetLatestVersion() *int32 {
	return s.LatestVersion
}

func (s *Service) GetMemory() *int32 {
	return s.Memory
}

func (s *Service) GetMessage() *string {
	return s.Message
}

func (s *Service) GetNamespace() *string {
	return s.Namespace
}

func (s *Service) GetParentUid() *string {
	return s.ParentUid
}

func (s *Service) GetPendingInstance() *int32 {
	return s.PendingInstance
}

func (s *Service) GetQuotaId() *string {
	return s.QuotaId
}

func (s *Service) GetReason() *string {
	return s.Reason
}

func (s *Service) GetRegion() *string {
	return s.Region
}

func (s *Service) GetRequestId() *string {
	return s.RequestId
}

func (s *Service) GetResource() *string {
	return s.Resource
}

func (s *Service) GetResourceAlias() *string {
	return s.ResourceAlias
}

func (s *Service) GetResourceBurstable() *bool {
	return s.ResourceBurstable
}

func (s *Service) GetRole() *string {
	return s.Role
}

func (s *Service) GetRoleAttrs() *string {
	return s.RoleAttrs
}

func (s *Service) GetRunningInstance() *int32 {
	return s.RunningInstance
}

func (s *Service) GetSafetyLock() *string {
	return s.SafetyLock
}

func (s *Service) GetSecondaryInternetEndpoint() *string {
	return s.SecondaryInternetEndpoint
}

func (s *Service) GetSecondaryIntranetEndpoint() *string {
	return s.SecondaryIntranetEndpoint
}

func (s *Service) GetServiceConfig() *string {
	return s.ServiceConfig
}

func (s *Service) GetServiceGroup() *string {
	return s.ServiceGroup
}

func (s *Service) GetServiceId() *string {
	return s.ServiceId
}

func (s *Service) GetServiceName() *string {
	return s.ServiceName
}

func (s *Service) GetServiceUid() *string {
	return s.ServiceUid
}

func (s *Service) GetSource() *string {
	return s.Source
}

func (s *Service) GetStatus() *string {
	return s.Status
}

func (s *Service) GetTotalInstance() *int32 {
	return s.TotalInstance
}

func (s *Service) GetTrafficState() *string {
	return s.TrafficState
}

func (s *Service) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Service) GetWeight() *int32 {
	return s.Weight
}

func (s *Service) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *Service) SetAutoscalerEnabled(v bool) *Service {
	s.AutoscalerEnabled = &v
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

func (s *Service) SetCronscalerEnabled(v bool) *Service {
	s.CronscalerEnabled = &v
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

func (s *Service) SetGPUCorePercentage(v int32) *Service {
	s.GPUCorePercentage = &v
	return s
}

func (s *Service) SetGPUMemory(v int32) *Service {
	s.GPUMemory = &v
	return s
}

func (s *Service) SetGateway(v string) *Service {
	s.Gateway = &v
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

func (s *Service) SetInstanceCountInResource(v *ServiceInstanceCountInResource) *Service {
	s.InstanceCountInResource = v
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

func (s *Service) SetQuotaId(v string) *Service {
	s.QuotaId = &v
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

func (s *Service) SetResourceBurstable(v bool) *Service {
	s.ResourceBurstable = &v
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

func (s *Service) SetTrafficState(v string) *Service {
	s.TrafficState = &v
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

func (s *Service) Validate() error {
	if s.InstanceCountInResource != nil {
		if err := s.InstanceCountInResource.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ServiceInstanceCountInResource struct {
	Dedicated *int32 `json:"Dedicated,omitempty" xml:"Dedicated,omitempty"`
	Public    *int32 `json:"Public,omitempty" xml:"Public,omitempty"`
	Quota     *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s ServiceInstanceCountInResource) String() string {
	return dara.Prettify(s)
}

func (s ServiceInstanceCountInResource) GoString() string {
	return s.String()
}

func (s *ServiceInstanceCountInResource) GetDedicated() *int32 {
	return s.Dedicated
}

func (s *ServiceInstanceCountInResource) GetPublic() *int32 {
	return s.Public
}

func (s *ServiceInstanceCountInResource) GetQuota() *int32 {
	return s.Quota
}

func (s *ServiceInstanceCountInResource) SetDedicated(v int32) *ServiceInstanceCountInResource {
	s.Dedicated = &v
	return s
}

func (s *ServiceInstanceCountInResource) SetPublic(v int32) *ServiceInstanceCountInResource {
	s.Public = &v
	return s
}

func (s *ServiceInstanceCountInResource) SetQuota(v int32) *ServiceInstanceCountInResource {
	s.Quota = &v
	return s
}

func (s *ServiceInstanceCountInResource) Validate() error {
	return dara.Validate(s)
}

type ServiceLabels struct {
	// The label key.
	//
	// example:
	//
	// key1
	LabelKey *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
	// The label value.
	//
	// example:
	//
	// value1
	LabelValue *string `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
}

func (s ServiceLabels) String() string {
	return dara.Prettify(s)
}

func (s ServiceLabels) GoString() string {
	return s.String()
}

func (s *ServiceLabels) GetLabelKey() *string {
	return s.LabelKey
}

func (s *ServiceLabels) GetLabelValue() *string {
	return s.LabelValue
}

func (s *ServiceLabels) SetLabelKey(v string) *ServiceLabels {
	s.LabelKey = &v
	return s
}

func (s *ServiceLabels) SetLabelValue(v string) *ServiceLabels {
	s.LabelValue = &v
	return s
}

func (s *ServiceLabels) Validate() error {
	return dara.Validate(s)
}
