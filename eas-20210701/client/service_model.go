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
	SetCallerUid(v string) *Service
	GetCallerUid() *string
	SetCpu(v int32) *Service
	GetCpu() *int32
	SetCreateTime(v string) *Service
	GetCreateTime() *string
	SetCurrentVersion(v int32) *Service
	GetCurrentVersion() *int32
	SetExtraData(v string) *Service
	GetExtraData() *string
	SetGateway(v string) *Service
	GetGateway() *string
	SetGpu(v int32) *Service
	GetGpu() *int32
	SetImage(v string) *Service
	GetImage() *string
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
	Gateway                   *string          `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
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
	QuotaId                   *string          `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
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
	TrafficState              *string          `json:"TrafficState,omitempty" xml:"TrafficState,omitempty"`
	UpdateTime                *string          `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Weight                    *int32           `json:"Weight,omitempty" xml:"Weight,omitempty"`
	WorkspaceId               *string          `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *Service) GetCallerUid() *string {
	return s.CallerUid
}

func (s *Service) GetCpu() *int32 {
	return s.Cpu
}

func (s *Service) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Service) GetCurrentVersion() *int32 {
	return s.CurrentVersion
}

func (s *Service) GetExtraData() *string {
	return s.ExtraData
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
	return dara.Validate(s)
}

type ServiceLabels struct {
	LabelKey   *string `json:"LabelKey,omitempty" xml:"LabelKey,omitempty"`
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
