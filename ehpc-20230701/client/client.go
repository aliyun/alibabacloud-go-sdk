// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddonNodeTemplate struct {
	AutoRenew          *bool                         `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod    *int32                        `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	DataDisks          []*AddonNodeTemplateDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	Duration           *int32                        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EnableHT           *bool                         `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	ImageId            *string                       `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType *string                       `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceType       *string                       `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OsName             *string                       `json:"OsName,omitempty" xml:"OsName,omitempty"`
	OsNameEN           *string                       `json:"OsNameEN,omitempty" xml:"OsNameEN,omitempty"`
	Period             *int32                        `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string                       `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SpotPriceLimit     *float32                      `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy       *string                       `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk         *AddonNodeTemplateSystemDisk  `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s AddonNodeTemplate) String() string {
	return tea.Prettify(s)
}

func (s AddonNodeTemplate) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplate) SetAutoRenew(v bool) *AddonNodeTemplate {
	s.AutoRenew = &v
	return s
}

func (s *AddonNodeTemplate) SetAutoRenewPeriod(v int32) *AddonNodeTemplate {
	s.AutoRenewPeriod = &v
	return s
}

func (s *AddonNodeTemplate) SetDataDisks(v []*AddonNodeTemplateDataDisks) *AddonNodeTemplate {
	s.DataDisks = v
	return s
}

func (s *AddonNodeTemplate) SetDuration(v int32) *AddonNodeTemplate {
	s.Duration = &v
	return s
}

func (s *AddonNodeTemplate) SetEnableHT(v bool) *AddonNodeTemplate {
	s.EnableHT = &v
	return s
}

func (s *AddonNodeTemplate) SetImageId(v string) *AddonNodeTemplate {
	s.ImageId = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceChargeType(v string) *AddonNodeTemplate {
	s.InstanceChargeType = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceType(v string) *AddonNodeTemplate {
	s.InstanceType = &v
	return s
}

func (s *AddonNodeTemplate) SetOsName(v string) *AddonNodeTemplate {
	s.OsName = &v
	return s
}

func (s *AddonNodeTemplate) SetOsNameEN(v string) *AddonNodeTemplate {
	s.OsNameEN = &v
	return s
}

func (s *AddonNodeTemplate) SetPeriod(v int32) *AddonNodeTemplate {
	s.Period = &v
	return s
}

func (s *AddonNodeTemplate) SetPeriodUnit(v string) *AddonNodeTemplate {
	s.PeriodUnit = &v
	return s
}

func (s *AddonNodeTemplate) SetSpotPriceLimit(v float32) *AddonNodeTemplate {
	s.SpotPriceLimit = &v
	return s
}

func (s *AddonNodeTemplate) SetSpotStrategy(v string) *AddonNodeTemplate {
	s.SpotStrategy = &v
	return s
}

func (s *AddonNodeTemplate) SetSystemDisk(v *AddonNodeTemplateSystemDisk) *AddonNodeTemplate {
	s.SystemDisk = v
	return s
}

type AddonNodeTemplateDataDisks struct {
	Category           *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Level              *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Size               *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddonNodeTemplateDataDisks) String() string {
	return tea.Prettify(s)
}

func (s AddonNodeTemplateDataDisks) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplateDataDisks) SetCategory(v string) *AddonNodeTemplateDataDisks {
	s.Category = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetDeleteWithInstance(v bool) *AddonNodeTemplateDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetLevel(v string) *AddonNodeTemplateDataDisks {
	s.Level = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetSize(v int32) *AddonNodeTemplateDataDisks {
	s.Size = &v
	return s
}

type AddonNodeTemplateSystemDisk struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddonNodeTemplateSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s AddonNodeTemplateSystemDisk) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplateSystemDisk) SetCategory(v string) *AddonNodeTemplateSystemDisk {
	s.Category = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) SetLevel(v string) *AddonNodeTemplateSystemDisk {
	s.Level = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) SetSize(v int32) *AddonNodeTemplateSystemDisk {
	s.Size = &v
	return s
}

type NodeTemplate struct {
	AutoRenew          *bool                    `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod    *int32                   `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	DataDisks          []*NodeTemplateDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	Duration           *int32                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EnableHT           *bool                    `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	ImageId            *string                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceChargeType *string                  `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceType       *string                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Period             *int32                   `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string                  `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SpotPriceLimit     *float32                 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy       *string                  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk         *NodeTemplateSystemDisk  `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s NodeTemplate) String() string {
	return tea.Prettify(s)
}

func (s NodeTemplate) GoString() string {
	return s.String()
}

func (s *NodeTemplate) SetAutoRenew(v bool) *NodeTemplate {
	s.AutoRenew = &v
	return s
}

func (s *NodeTemplate) SetAutoRenewPeriod(v int32) *NodeTemplate {
	s.AutoRenewPeriod = &v
	return s
}

func (s *NodeTemplate) SetDataDisks(v []*NodeTemplateDataDisks) *NodeTemplate {
	s.DataDisks = v
	return s
}

func (s *NodeTemplate) SetDuration(v int32) *NodeTemplate {
	s.Duration = &v
	return s
}

func (s *NodeTemplate) SetEnableHT(v bool) *NodeTemplate {
	s.EnableHT = &v
	return s
}

func (s *NodeTemplate) SetImageId(v string) *NodeTemplate {
	s.ImageId = &v
	return s
}

func (s *NodeTemplate) SetInstanceChargeType(v string) *NodeTemplate {
	s.InstanceChargeType = &v
	return s
}

func (s *NodeTemplate) SetInstanceType(v string) *NodeTemplate {
	s.InstanceType = &v
	return s
}

func (s *NodeTemplate) SetPeriod(v int32) *NodeTemplate {
	s.Period = &v
	return s
}

func (s *NodeTemplate) SetPeriodUnit(v string) *NodeTemplate {
	s.PeriodUnit = &v
	return s
}

func (s *NodeTemplate) SetSpotPriceLimit(v float32) *NodeTemplate {
	s.SpotPriceLimit = &v
	return s
}

func (s *NodeTemplate) SetSpotStrategy(v string) *NodeTemplate {
	s.SpotStrategy = &v
	return s
}

func (s *NodeTemplate) SetSystemDisk(v *NodeTemplateSystemDisk) *NodeTemplate {
	s.SystemDisk = v
	return s
}

type NodeTemplateDataDisks struct {
	Category           *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance *bool   `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Level              *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Size               *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s NodeTemplateDataDisks) String() string {
	return tea.Prettify(s)
}

func (s NodeTemplateDataDisks) GoString() string {
	return s.String()
}

func (s *NodeTemplateDataDisks) SetCategory(v string) *NodeTemplateDataDisks {
	s.Category = &v
	return s
}

func (s *NodeTemplateDataDisks) SetDeleteWithInstance(v bool) *NodeTemplateDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *NodeTemplateDataDisks) SetLevel(v string) *NodeTemplateDataDisks {
	s.Level = &v
	return s
}

func (s *NodeTemplateDataDisks) SetSize(v int32) *NodeTemplateDataDisks {
	s.Size = &v
	return s
}

type NodeTemplateSystemDisk struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s NodeTemplateSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s NodeTemplateSystemDisk) GoString() string {
	return s.String()
}

func (s *NodeTemplateSystemDisk) SetCategory(v string) *NodeTemplateSystemDisk {
	s.Category = &v
	return s
}

func (s *NodeTemplateSystemDisk) SetLevel(v string) *NodeTemplateSystemDisk {
	s.Level = &v
	return s
}

func (s *NodeTemplateSystemDisk) SetSize(v int32) *NodeTemplateSystemDisk {
	s.Size = &v
	return s
}

type QueueTemplate struct {
	AllocationStrategy *string         `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	ComputeNodes       []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	EnableScaleIn      *bool           `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	EnableScaleOut     *bool           `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	HostnamePrefix     *string         `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	HostnameSuffix     *string         `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	InitialCount       *int32          `json:"InitialCount,omitempty" xml:"InitialCount,omitempty"`
	InterConnect       *string         `json:"InterConnect,omitempty" xml:"InterConnect,omitempty"`
	KeepAliveNodes     []*string       `json:"KeepAliveNodes,omitempty" xml:"KeepAliveNodes,omitempty" type:"Repeated"`
	MaxCount           *int32          `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	MaxCountPerCycle   *int64          `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	MinCount           *int32          `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	Name               *string         `json:"Name,omitempty" xml:"Name,omitempty"`
	RamRole            *string         `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	VSwitchIds         []*string       `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s QueueTemplate) String() string {
	return tea.Prettify(s)
}

func (s QueueTemplate) GoString() string {
	return s.String()
}

func (s *QueueTemplate) SetAllocationStrategy(v string) *QueueTemplate {
	s.AllocationStrategy = &v
	return s
}

func (s *QueueTemplate) SetComputeNodes(v []*NodeTemplate) *QueueTemplate {
	s.ComputeNodes = v
	return s
}

func (s *QueueTemplate) SetEnableScaleIn(v bool) *QueueTemplate {
	s.EnableScaleIn = &v
	return s
}

func (s *QueueTemplate) SetEnableScaleOut(v bool) *QueueTemplate {
	s.EnableScaleOut = &v
	return s
}

func (s *QueueTemplate) SetHostnamePrefix(v string) *QueueTemplate {
	s.HostnamePrefix = &v
	return s
}

func (s *QueueTemplate) SetHostnameSuffix(v string) *QueueTemplate {
	s.HostnameSuffix = &v
	return s
}

func (s *QueueTemplate) SetInitialCount(v int32) *QueueTemplate {
	s.InitialCount = &v
	return s
}

func (s *QueueTemplate) SetInterConnect(v string) *QueueTemplate {
	s.InterConnect = &v
	return s
}

func (s *QueueTemplate) SetKeepAliveNodes(v []*string) *QueueTemplate {
	s.KeepAliveNodes = v
	return s
}

func (s *QueueTemplate) SetMaxCount(v int32) *QueueTemplate {
	s.MaxCount = &v
	return s
}

func (s *QueueTemplate) SetMaxCountPerCycle(v int64) *QueueTemplate {
	s.MaxCountPerCycle = &v
	return s
}

func (s *QueueTemplate) SetMinCount(v int32) *QueueTemplate {
	s.MinCount = &v
	return s
}

func (s *QueueTemplate) SetName(v string) *QueueTemplate {
	s.Name = &v
	return s
}

func (s *QueueTemplate) SetRamRole(v string) *QueueTemplate {
	s.RamRole = &v
	return s
}

func (s *QueueTemplate) SetVSwitchIds(v []*string) *QueueTemplate {
	s.VSwitchIds = v
	return s
}

type SharedStorageTemplate struct {
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountDirectory    *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
	MountOptions      *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NASDirectory      *string `json:"NASDirectory,omitempty" xml:"NASDirectory,omitempty"`
	ProtocolType      *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
}

func (s SharedStorageTemplate) String() string {
	return tea.Prettify(s)
}

func (s SharedStorageTemplate) GoString() string {
	return s.String()
}

func (s *SharedStorageTemplate) SetFileSystemId(v string) *SharedStorageTemplate {
	s.FileSystemId = &v
	return s
}

func (s *SharedStorageTemplate) SetMountDirectory(v string) *SharedStorageTemplate {
	s.MountDirectory = &v
	return s
}

func (s *SharedStorageTemplate) SetMountOptions(v string) *SharedStorageTemplate {
	s.MountOptions = &v
	return s
}

func (s *SharedStorageTemplate) SetMountTargetDomain(v string) *SharedStorageTemplate {
	s.MountTargetDomain = &v
	return s
}

func (s *SharedStorageTemplate) SetNASDirectory(v string) *SharedStorageTemplate {
	s.NASDirectory = &v
	return s
}

func (s *SharedStorageTemplate) SetProtocolType(v string) *SharedStorageTemplate {
	s.ProtocolType = &v
	return s
}

type AddImageRequest struct {
	ContainerImageSpec *AddImageRequestContainerImageSpec `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty" type:"Struct"`
	Description        *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Name               *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	VMImageSpec        *AddImageRequestVMImageSpec        `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty" type:"Struct"`
	Version            *string                            `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetContainerImageSpec(v *AddImageRequestContainerImageSpec) *AddImageRequest {
	s.ContainerImageSpec = v
	return s
}

func (s *AddImageRequest) SetDescription(v string) *AddImageRequest {
	s.Description = &v
	return s
}

func (s *AddImageRequest) SetName(v string) *AddImageRequest {
	s.Name = &v
	return s
}

func (s *AddImageRequest) SetVMImageSpec(v *AddImageRequestVMImageSpec) *AddImageRequest {
	s.VMImageSpec = v
	return s
}

func (s *AddImageRequest) SetVersion(v string) *AddImageRequest {
	s.Version = &v
	return s
}

type AddImageRequestContainerImageSpec struct {
	IsACREnterprise    *bool                                                `json:"IsACREnterprise,omitempty" xml:"IsACREnterprise,omitempty"`
	IsACRRegistry      *bool                                                `json:"IsACRRegistry,omitempty" xml:"IsACRRegistry,omitempty"`
	RegistryCredential *AddImageRequestContainerImageSpecRegistryCredential `json:"RegistryCredential,omitempty" xml:"RegistryCredential,omitempty" type:"Struct"`
	RegistryCriId      *string                                              `json:"RegistryCriId,omitempty" xml:"RegistryCriId,omitempty"`
	RegistryUrl        *string                                              `json:"RegistryUrl,omitempty" xml:"RegistryUrl,omitempty"`
}

func (s AddImageRequestContainerImageSpec) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequestContainerImageSpec) GoString() string {
	return s.String()
}

func (s *AddImageRequestContainerImageSpec) SetIsACREnterprise(v bool) *AddImageRequestContainerImageSpec {
	s.IsACREnterprise = &v
	return s
}

func (s *AddImageRequestContainerImageSpec) SetIsACRRegistry(v bool) *AddImageRequestContainerImageSpec {
	s.IsACRRegistry = &v
	return s
}

func (s *AddImageRequestContainerImageSpec) SetRegistryCredential(v *AddImageRequestContainerImageSpecRegistryCredential) *AddImageRequestContainerImageSpec {
	s.RegistryCredential = v
	return s
}

func (s *AddImageRequestContainerImageSpec) SetRegistryCriId(v string) *AddImageRequestContainerImageSpec {
	s.RegistryCriId = &v
	return s
}

func (s *AddImageRequestContainerImageSpec) SetRegistryUrl(v string) *AddImageRequestContainerImageSpec {
	s.RegistryUrl = &v
	return s
}

type AddImageRequestContainerImageSpecRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddImageRequestContainerImageSpecRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequestContainerImageSpecRegistryCredential) GoString() string {
	return s.String()
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) SetPassword(v string) *AddImageRequestContainerImageSpecRegistryCredential {
	s.Password = &v
	return s
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) SetServer(v string) *AddImageRequestContainerImageSpecRegistryCredential {
	s.Server = &v
	return s
}

func (s *AddImageRequestContainerImageSpecRegistryCredential) SetUserName(v string) *AddImageRequestContainerImageSpecRegistryCredential {
	s.UserName = &v
	return s
}

type AddImageRequestVMImageSpec struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s AddImageRequestVMImageSpec) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequestVMImageSpec) GoString() string {
	return s.String()
}

func (s *AddImageRequestVMImageSpec) SetImageId(v string) *AddImageRequestVMImageSpec {
	s.ImageId = &v
	return s
}

type AddImageShrinkRequest struct {
	ContainerImageSpecShrink *string `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VMImageSpecShrink        *string `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty"`
	Version                  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AddImageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddImageShrinkRequest) SetContainerImageSpecShrink(v string) *AddImageShrinkRequest {
	s.ContainerImageSpecShrink = &v
	return s
}

func (s *AddImageShrinkRequest) SetDescription(v string) *AddImageShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddImageShrinkRequest) SetName(v string) *AddImageShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddImageShrinkRequest) SetVMImageSpecShrink(v string) *AddImageShrinkRequest {
	s.VMImageSpecShrink = &v
	return s
}

func (s *AddImageShrinkRequest) SetVersion(v string) *AddImageShrinkRequest {
	s.Version = &v
	return s
}

type AddImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetImageId(v string) *AddImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetSuccess(v bool) *AddImageResponseBody {
	s.Success = &v
	return s
}

type AddImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponse) GoString() string {
	return s.String()
}

func (s *AddImageResponse) SetHeaders(v map[string]*string) *AddImageResponse {
	s.Headers = v
	return s
}

func (s *AddImageResponse) SetStatusCode(v int32) *AddImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	DeploymentPolicy *CreateJobRequestDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	JobDescription   *string                           `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	JobName          *string                           `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Tasks            []*CreateJobRequestTasks          `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetDeploymentPolicy(v *CreateJobRequestDeploymentPolicy) *CreateJobRequest {
	s.DeploymentPolicy = v
	return s
}

func (s *CreateJobRequest) SetJobDescription(v string) *CreateJobRequest {
	s.JobDescription = &v
	return s
}

func (s *CreateJobRequest) SetJobName(v string) *CreateJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobRequest) SetTasks(v []*CreateJobRequestTasks) *CreateJobRequest {
	s.Tasks = v
	return s
}

type CreateJobRequestDeploymentPolicy struct {
	AllocationSpec *string                                  `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	Network        *CreateJobRequestDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
}

func (s CreateJobRequestDeploymentPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestDeploymentPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDeploymentPolicy) SetAllocationSpec(v string) *CreateJobRequestDeploymentPolicy {
	s.AllocationSpec = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetNetwork(v *CreateJobRequestDeploymentPolicyNetwork) *CreateJobRequestDeploymentPolicy {
	s.Network = v
	return s
}

type CreateJobRequestDeploymentPolicyNetwork struct {
	Vswitch []*string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
}

func (s CreateJobRequestDeploymentPolicyNetwork) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDeploymentPolicyNetwork) SetVswitch(v []*string) *CreateJobRequestDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

type CreateJobRequestTasks struct {
	ExecutorPolicy  *CreateJobRequestTasksExecutorPolicy `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	TaskName        *string                              `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskSpec        *CreateJobRequestTasksTaskSpec       `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	TaskSustainable *bool                                `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s CreateJobRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasks) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasks) SetExecutorPolicy(v *CreateJobRequestTasksExecutorPolicy) *CreateJobRequestTasks {
	s.ExecutorPolicy = v
	return s
}

func (s *CreateJobRequestTasks) SetTaskName(v string) *CreateJobRequestTasks {
	s.TaskName = &v
	return s
}

func (s *CreateJobRequestTasks) SetTaskSpec(v *CreateJobRequestTasksTaskSpec) *CreateJobRequestTasks {
	s.TaskSpec = v
	return s
}

func (s *CreateJobRequestTasks) SetTaskSustainable(v bool) *CreateJobRequestTasks {
	s.TaskSustainable = &v
	return s
}

type CreateJobRequestTasksExecutorPolicy struct {
	ArraySpec *CreateJobRequestTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
	MaxCount  *int32                                        `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s CreateJobRequestTasksExecutorPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksExecutorPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksExecutorPolicy) SetArraySpec(v *CreateJobRequestTasksExecutorPolicyArraySpec) *CreateJobRequestTasksExecutorPolicy {
	s.ArraySpec = v
	return s
}

func (s *CreateJobRequestTasksExecutorPolicy) SetMaxCount(v int32) *CreateJobRequestTasksExecutorPolicy {
	s.MaxCount = &v
	return s
}

type CreateJobRequestTasksExecutorPolicyArraySpec struct {
	IndexEnd   *int32 `json:"IndexEnd,omitempty" xml:"IndexEnd,omitempty"`
	IndexStart *int32 `json:"IndexStart,omitempty" xml:"IndexStart,omitempty"`
	IndexStep  *int32 `json:"IndexStep,omitempty" xml:"IndexStep,omitempty"`
}

func (s CreateJobRequestTasksExecutorPolicyArraySpec) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksExecutorPolicyArraySpec) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) SetIndexEnd(v int32) *CreateJobRequestTasksExecutorPolicyArraySpec {
	s.IndexEnd = &v
	return s
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) SetIndexStart(v int32) *CreateJobRequestTasksExecutorPolicyArraySpec {
	s.IndexStart = &v
	return s
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) SetIndexStep(v int32) *CreateJobRequestTasksExecutorPolicyArraySpec {
	s.IndexStep = &v
	return s
}

type CreateJobRequestTasksTaskSpec struct {
	Resource     *CreateJobRequestTasksTaskSpecResource       `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	TaskExecutor []*CreateJobRequestTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
}

func (s CreateJobRequestTasksTaskSpec) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpec) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpec) SetResource(v *CreateJobRequestTasksTaskSpecResource) *CreateJobRequestTasksTaskSpec {
	s.Resource = v
	return s
}

func (s *CreateJobRequestTasksTaskSpec) SetTaskExecutor(v []*CreateJobRequestTasksTaskSpecTaskExecutor) *CreateJobRequestTasksTaskSpec {
	s.TaskExecutor = v
	return s
}

type CreateJobRequestTasksTaskSpecResource struct {
	Cores  *float32                                      `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Disks  []*CreateJobRequestTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	Memory *float32                                      `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecResource) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecResource) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecResource) SetCores(v float32) *CreateJobRequestTasksTaskSpecResource {
	s.Cores = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetDisks(v []*CreateJobRequestTasksTaskSpecResourceDisks) *CreateJobRequestTasksTaskSpecResource {
	s.Disks = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetMemory(v float32) *CreateJobRequestTasksTaskSpecResource {
	s.Memory = &v
	return s
}

type CreateJobRequestTasksTaskSpecResourceDisks struct {
	Size *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecResourceDisks) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecResourceDisks) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecResourceDisks) SetSize(v int32) *CreateJobRequestTasksTaskSpecResourceDisks {
	s.Size = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResourceDisks) SetType(v string) *CreateJobRequestTasksTaskSpecResourceDisks {
	s.Type = &v
	return s
}

type CreateJobRequestTasksTaskSpecTaskExecutor struct {
	VM *CreateJobRequestTasksTaskSpecTaskExecutorVM `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutor) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutor) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutor) SetVM(v *CreateJobRequestTasksTaskSpecTaskExecutorVM) *CreateJobRequestTasksTaskSpecTaskExecutor {
	s.VM = v
	return s
}

type CreateJobRequestTasksTaskSpecTaskExecutorVM struct {
	Image        *string `json:"Image,omitempty" xml:"Image,omitempty"`
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	Script       *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorVM) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorVM) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetImage(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.Image = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetPrologScript(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.PrologScript = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetScript(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.Script = &v
	return s
}

type CreateJobShrinkRequest struct {
	DeploymentPolicyShrink *string `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty"`
	JobDescription         *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	JobName                *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	TasksShrink            *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
}

func (s CreateJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobShrinkRequest) SetDeploymentPolicyShrink(v string) *CreateJobShrinkRequest {
	s.DeploymentPolicyShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobDescription(v string) *CreateJobShrinkRequest {
	s.JobDescription = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobName(v string) *CreateJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTasksShrink(v string) *CreateJobShrinkRequest {
	s.TasksShrink = &v
	return s
}

type CreateJobResponseBody struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type DeleteJobsRequest struct {
	ExecutorIds []*string                   `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	JobSpec     []*DeleteJobsRequestJobSpec `json:"JobSpec,omitempty" xml:"JobSpec,omitempty" type:"Repeated"`
}

func (s DeleteJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) SetExecutorIds(v []*string) *DeleteJobsRequest {
	s.ExecutorIds = v
	return s
}

func (s *DeleteJobsRequest) SetJobSpec(v []*DeleteJobsRequestJobSpec) *DeleteJobsRequest {
	s.JobSpec = v
	return s
}

type DeleteJobsRequestJobSpec struct {
	JobId    *string                             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	TaskSpec []*DeleteJobsRequestJobSpecTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Repeated"`
}

func (s DeleteJobsRequestJobSpec) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequestJobSpec) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequestJobSpec) SetJobId(v string) *DeleteJobsRequestJobSpec {
	s.JobId = &v
	return s
}

func (s *DeleteJobsRequestJobSpec) SetTaskSpec(v []*DeleteJobsRequestJobSpecTaskSpec) *DeleteJobsRequestJobSpec {
	s.TaskSpec = v
	return s
}

type DeleteJobsRequestJobSpecTaskSpec struct {
	ArrayIndex []*int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty" type:"Repeated"`
	TaskName   *string  `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DeleteJobsRequestJobSpecTaskSpec) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequestJobSpecTaskSpec) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequestJobSpecTaskSpec) SetArrayIndex(v []*int32) *DeleteJobsRequestJobSpecTaskSpec {
	s.ArrayIndex = v
	return s
}

func (s *DeleteJobsRequestJobSpecTaskSpec) SetTaskName(v string) *DeleteJobsRequestJobSpecTaskSpec {
	s.TaskName = &v
	return s
}

type DeleteJobsShrinkRequest struct {
	ExecutorIdsShrink *string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty"`
	JobSpecShrink     *string `json:"JobSpec,omitempty" xml:"JobSpec,omitempty"`
}

func (s DeleteJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsShrinkRequest) SetExecutorIdsShrink(v string) *DeleteJobsShrinkRequest {
	s.ExecutorIdsShrink = &v
	return s
}

func (s *DeleteJobsShrinkRequest) SetJobSpecShrink(v string) *DeleteJobsShrinkRequest {
	s.JobSpecShrink = &v
	return s
}

type DeleteJobsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponse) SetHeaders(v map[string]*string) *DeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobsResponse) SetStatusCode(v int32) *DeleteJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobsResponse) SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetImageId(v string) *GetImageRequest {
	s.ImageId = &v
	return s
}

type GetImageResponseBody struct {
	Image      *GetImageResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	RequestId  *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) SetImage(v *GetImageResponseBodyImage) *GetImageResponseBody {
	s.Image = v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSuccess(v bool) *GetImageResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageResponseBody) SetTotalCount(v int32) *GetImageResponseBody {
	s.TotalCount = &v
	return s
}

type GetImageResponseBodyImage struct {
	ContainerImageSpec *GetImageResponseBodyImageContainerImageSpec `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty" type:"Struct"`
	CreateTime         *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description        *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageType          *string                                      `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Name               *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Size               *string                                      `json:"Size,omitempty" xml:"Size,omitempty"`
	VMImageSpec        *GetImageResponseBodyImageVMImageSpec        `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty" type:"Struct"`
	Version            *string                                      `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetImageResponseBodyImage) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImage) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImage) SetContainerImageSpec(v *GetImageResponseBodyImageContainerImageSpec) *GetImageResponseBodyImage {
	s.ContainerImageSpec = v
	return s
}

func (s *GetImageResponseBodyImage) SetCreateTime(v string) *GetImageResponseBodyImage {
	s.CreateTime = &v
	return s
}

func (s *GetImageResponseBodyImage) SetDescription(v string) *GetImageResponseBodyImage {
	s.Description = &v
	return s
}

func (s *GetImageResponseBodyImage) SetImageType(v string) *GetImageResponseBodyImage {
	s.ImageType = &v
	return s
}

func (s *GetImageResponseBodyImage) SetName(v string) *GetImageResponseBodyImage {
	s.Name = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSize(v string) *GetImageResponseBodyImage {
	s.Size = &v
	return s
}

func (s *GetImageResponseBodyImage) SetVMImageSpec(v *GetImageResponseBodyImageVMImageSpec) *GetImageResponseBodyImage {
	s.VMImageSpec = v
	return s
}

func (s *GetImageResponseBodyImage) SetVersion(v string) *GetImageResponseBodyImage {
	s.Version = &v
	return s
}

type GetImageResponseBodyImageContainerImageSpec struct {
	IsACREnterprise    *bool                                                          `json:"IsACREnterprise,omitempty" xml:"IsACREnterprise,omitempty"`
	IsACRRegistry      *bool                                                          `json:"IsACRRegistry,omitempty" xml:"IsACRRegistry,omitempty"`
	RegistryCredential *GetImageResponseBodyImageContainerImageSpecRegistryCredential `json:"RegistryCredential,omitempty" xml:"RegistryCredential,omitempty" type:"Struct"`
	RegistryCriId      *string                                                        `json:"RegistryCriId,omitempty" xml:"RegistryCriId,omitempty"`
	RegistryUrl        *string                                                        `json:"RegistryUrl,omitempty" xml:"RegistryUrl,omitempty"`
}

func (s GetImageResponseBodyImageContainerImageSpec) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImageContainerImageSpec) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetIsACREnterprise(v bool) *GetImageResponseBodyImageContainerImageSpec {
	s.IsACREnterprise = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetIsACRRegistry(v bool) *GetImageResponseBodyImageContainerImageSpec {
	s.IsACRRegistry = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetRegistryCredential(v *GetImageResponseBodyImageContainerImageSpecRegistryCredential) *GetImageResponseBodyImageContainerImageSpec {
	s.RegistryCredential = v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetRegistryCriId(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.RegistryCriId = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetRegistryUrl(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.RegistryUrl = &v
	return s
}

type GetImageResponseBodyImageContainerImageSpecRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetImageResponseBodyImageContainerImageSpecRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImageContainerImageSpecRegistryCredential) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) SetPassword(v string) *GetImageResponseBodyImageContainerImageSpecRegistryCredential {
	s.Password = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) SetServer(v string) *GetImageResponseBodyImageContainerImageSpecRegistryCredential {
	s.Server = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpecRegistryCredential) SetUserName(v string) *GetImageResponseBodyImageContainerImageSpecRegistryCredential {
	s.UserName = &v
	return s
}

type GetImageResponseBodyImageVMImageSpec struct {
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s GetImageResponseBodyImageVMImageSpec) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImageVMImageSpec) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageVMImageSpec) SetArchitecture(v string) *GetImageResponseBodyImageVMImageSpec {
	s.Architecture = &v
	return s
}

func (s *GetImageResponseBodyImageVMImageSpec) SetImageId(v string) *GetImageResponseBodyImageVMImageSpec {
	s.ImageId = &v
	return s
}

func (s *GetImageResponseBodyImageVMImageSpec) SetOsTag(v string) *GetImageResponseBodyImageVMImageSpec {
	s.OsTag = &v
	return s
}

func (s *GetImageResponseBodyImageVMImageSpec) SetPlatform(v string) *GetImageResponseBodyImageVMImageSpec {
	s.Platform = &v
	return s
}

type GetImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s *GetImageResponse) SetHeaders(v map[string]*string) *GetImageResponse {
	s.Headers = v
	return s
}

func (s *GetImageResponse) SetStatusCode(v int32) *GetImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetJobId(v string) *GetJobRequest {
	s.JobId = &v
	return s
}

type GetJobResponseBody struct {
	JobInfo *GetJobResponseBodyJobInfo `json:"JobInfo,omitempty" xml:"JobInfo,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetJobInfo(v *GetJobResponseBodyJobInfo) *GetJobResponseBody {
	s.JobInfo = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

type GetJobResponseBodyJobInfo struct {
	CreateTime       *string                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeploymentPolicy *GetJobResponseBodyJobInfoDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	EndTime          *string                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobDescription   *string                                    `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	JobId            *string                                    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName          *string                                    `json:"JobName,omitempty" xml:"JobName,omitempty"`
	StartTime        *string                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status           *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tasks            []*GetJobResponseBodyJobInfoTasks          `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfo) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfo) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfo) SetCreateTime(v string) *GetJobResponseBodyJobInfo {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetDeploymentPolicy(v *GetJobResponseBodyJobInfoDeploymentPolicy) *GetJobResponseBodyJobInfo {
	s.DeploymentPolicy = v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetEndTime(v string) *GetJobResponseBodyJobInfo {
	s.EndTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobDescription(v string) *GetJobResponseBodyJobInfo {
	s.JobDescription = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobId(v string) *GetJobResponseBodyJobInfo {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobName(v string) *GetJobResponseBodyJobInfo {
	s.JobName = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetStartTime(v string) *GetJobResponseBodyJobInfo {
	s.StartTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetStatus(v string) *GetJobResponseBodyJobInfo {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetTasks(v []*GetJobResponseBodyJobInfoTasks) *GetJobResponseBodyJobInfo {
	s.Tasks = v
	return s
}

type GetJobResponseBodyJobInfoDeploymentPolicy struct {
	AllocationSpec *string                                           `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	Network        *GetJobResponseBodyJobInfoDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
}

func (s GetJobResponseBodyJobInfoDeploymentPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDeploymentPolicy) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetAllocationSpec(v string) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.AllocationSpec = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetNetwork(v *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.Network = v
	return s
}

type GetJobResponseBodyJobInfoDeploymentPolicyNetwork struct {
	Vswitch []*string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyNetwork) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) SetVswitch(v []*string) *GetJobResponseBodyJobInfoDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

type GetJobResponseBodyJobInfoTasks struct {
	ExecutorPolicy  *GetJobResponseBodyJobInfoTasksExecutorPolicy   `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	ExecutorStatus  []*GetJobResponseBodyJobInfoTasksExecutorStatus `json:"ExecutorStatus,omitempty" xml:"ExecutorStatus,omitempty" type:"Repeated"`
	TaskName        *string                                         `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskSpec        *GetJobResponseBodyJobInfoTasksTaskSpec         `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	TaskSustainable *bool                                           `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasks) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasks) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasks) SetExecutorPolicy(v *GetJobResponseBodyJobInfoTasksExecutorPolicy) *GetJobResponseBodyJobInfoTasks {
	s.ExecutorPolicy = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) SetExecutorStatus(v []*GetJobResponseBodyJobInfoTasksExecutorStatus) *GetJobResponseBodyJobInfoTasks {
	s.ExecutorStatus = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) SetTaskName(v string) *GetJobResponseBodyJobInfoTasks {
	s.TaskName = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) SetTaskSpec(v *GetJobResponseBodyJobInfoTasksTaskSpec) *GetJobResponseBodyJobInfoTasks {
	s.TaskSpec = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) SetTaskSustainable(v bool) *GetJobResponseBodyJobInfoTasks {
	s.TaskSustainable = &v
	return s
}

type GetJobResponseBodyJobInfoTasksExecutorPolicy struct {
	ArraySpec *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
	MaxCount  *int32                                                 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksExecutorPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksExecutorPolicy) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicy) SetArraySpec(v *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) *GetJobResponseBodyJobInfoTasksExecutorPolicy {
	s.ArraySpec = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicy) SetMaxCount(v int32) *GetJobResponseBodyJobInfoTasksExecutorPolicy {
	s.MaxCount = &v
	return s
}

type GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec struct {
	IndexEnd   *int32 `json:"IndexEnd,omitempty" xml:"IndexEnd,omitempty"`
	IndexStart *int32 `json:"IndexStart,omitempty" xml:"IndexStart,omitempty"`
	IndexStep  *int32 `json:"IndexStep,omitempty" xml:"IndexStep,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) SetIndexEnd(v int32) *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec {
	s.IndexEnd = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) SetIndexStart(v int32) *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec {
	s.IndexStart = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) SetIndexStep(v int32) *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec {
	s.IndexStep = &v
	return s
}

type GetJobResponseBodyJobInfoTasksExecutorStatus struct {
	ArrayId      *int32  `json:"ArrayId,omitempty" xml:"ArrayId,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksExecutorStatus) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksExecutorStatus) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetArrayId(v int32) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.ArrayId = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetCreateTime(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetEndTime(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.EndTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetStartTime(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.StartTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetStatus(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetStatusReason(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.StatusReason = &v
	return s
}

type GetJobResponseBodyJobInfoTasksTaskSpec struct {
	Resource     *GetJobResponseBodyJobInfoTasksTaskSpecResource       `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	TaskExecutor []*GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpec) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpec) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) SetResource(v *GetJobResponseBodyJobInfoTasksTaskSpecResource) *GetJobResponseBodyJobInfoTasksTaskSpec {
	s.Resource = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) SetTaskExecutor(v []*GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) *GetJobResponseBodyJobInfoTasksTaskSpec {
	s.TaskExecutor = v
	return s
}

type GetJobResponseBodyJobInfoTasksTaskSpecResource struct {
	Cores  *float32                                               `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Disks  []*GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	Memory *int32                                                 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecResource) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecResource) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetCores(v float32) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.Cores = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetDisks(v []*GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.Disks = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetMemory(v int32) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.Memory = &v
	return s
}

type GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks struct {
	Size *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) SetSize(v int32) *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks {
	s.Size = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) SetType(v string) *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks {
	s.Type = &v
	return s
}

type GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor struct {
	VM *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) SetVM(v *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor {
	s.VM = v
	return s
}

type GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM struct {
	Image        *string `json:"Image,omitempty" xml:"Image,omitempty"`
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	Script       *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) SetImage(v string) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM {
	s.Image = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) SetPrologScript(v string) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM {
	s.PrologScript = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) SetScript(v string) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM {
	s.Script = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type ListExecutorsRequest struct {
	Filter     *ListExecutorsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	PageNumber *string                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExecutorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorsRequest) SetFilter(v *ListExecutorsRequestFilter) *ListExecutorsRequest {
	s.Filter = v
	return s
}

func (s *ListExecutorsRequest) SetPageNumber(v string) *ListExecutorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsRequest) SetPageSize(v string) *ListExecutorsRequest {
	s.PageSize = &v
	return s
}

type ListExecutorsRequestFilter struct {
	ExecutorIds       []*string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	IpAddresses       []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	JobName           *string   `json:"JobName,omitempty" xml:"JobName,omitempty"`
	TimeCreatedAfter  *int32    `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	TimeCreatedBefore *int32    `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
}

func (s ListExecutorsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListExecutorsRequestFilter) SetExecutorIds(v []*string) *ListExecutorsRequestFilter {
	s.ExecutorIds = v
	return s
}

func (s *ListExecutorsRequestFilter) SetIpAddresses(v []*string) *ListExecutorsRequestFilter {
	s.IpAddresses = v
	return s
}

func (s *ListExecutorsRequestFilter) SetJobName(v string) *ListExecutorsRequestFilter {
	s.JobName = &v
	return s
}

func (s *ListExecutorsRequestFilter) SetTimeCreatedAfter(v int32) *ListExecutorsRequestFilter {
	s.TimeCreatedAfter = &v
	return s
}

func (s *ListExecutorsRequestFilter) SetTimeCreatedBefore(v int32) *ListExecutorsRequestFilter {
	s.TimeCreatedBefore = &v
	return s
}

type ListExecutorsShrinkRequest struct {
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExecutorsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorsShrinkRequest) SetFilterShrink(v string) *ListExecutorsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListExecutorsShrinkRequest) SetPageNumber(v string) *ListExecutorsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsShrinkRequest) SetPageSize(v string) *ListExecutorsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListExecutorsResponseBody struct {
	Executors  []*ListExecutorsResponseBodyExecutors `json:"Executors,omitempty" xml:"Executors,omitempty" type:"Repeated"`
	JobId      *string                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber *string                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskName   *string                               `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TotalCount *string                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExecutorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBody) SetExecutors(v []*ListExecutorsResponseBodyExecutors) *ListExecutorsResponseBody {
	s.Executors = v
	return s
}

func (s *ListExecutorsResponseBody) SetJobId(v string) *ListExecutorsResponseBody {
	s.JobId = &v
	return s
}

func (s *ListExecutorsResponseBody) SetPageNumber(v string) *ListExecutorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsResponseBody) SetPageSize(v string) *ListExecutorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExecutorsResponseBody) SetRequestId(v string) *ListExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorsResponseBody) SetTaskName(v string) *ListExecutorsResponseBody {
	s.TaskName = &v
	return s
}

func (s *ListExecutorsResponseBody) SetTotalCount(v string) *ListExecutorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListExecutorsResponseBodyExecutors struct {
	ArrayIndex   *int32    `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	CreateTime   *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime      *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutorId   *string   `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	HostName     []*string `json:"HostName,omitempty" xml:"HostName,omitempty" type:"Repeated"`
	IpAddress    []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
	JobId        *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName      *string   `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusReason *string   `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	TaskName     *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListExecutorsResponseBodyExecutors) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutors) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutors) SetArrayIndex(v int32) *ListExecutorsResponseBodyExecutors {
	s.ArrayIndex = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetCreateTime(v string) *ListExecutorsResponseBodyExecutors {
	s.CreateTime = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetEndTime(v string) *ListExecutorsResponseBodyExecutors {
	s.EndTime = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetExecutorId(v string) *ListExecutorsResponseBodyExecutors {
	s.ExecutorId = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetHostName(v []*string) *ListExecutorsResponseBodyExecutors {
	s.HostName = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetIpAddress(v []*string) *ListExecutorsResponseBodyExecutors {
	s.IpAddress = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetJobId(v string) *ListExecutorsResponseBodyExecutors {
	s.JobId = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetJobName(v string) *ListExecutorsResponseBodyExecutors {
	s.JobName = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetStatus(v string) *ListExecutorsResponseBodyExecutors {
	s.Status = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetStatusReason(v string) *ListExecutorsResponseBodyExecutors {
	s.StatusReason = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetTaskName(v string) *ListExecutorsResponseBodyExecutors {
	s.TaskName = &v
	return s
}

type ListExecutorsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponse) SetHeaders(v map[string]*string) *ListExecutorsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutorsResponse) SetStatusCode(v int32) *ListExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutorsResponse) SetBody(v *ListExecutorsResponseBody) *ListExecutorsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	ImageIds   []*string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty" type:"Repeated"`
	ImageNames []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	PageNumber *int64    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetImageIds(v []*string) *ListImagesRequest {
	s.ImageIds = v
	return s
}

func (s *ListImagesRequest) SetImageNames(v []*string) *ListImagesRequest {
	s.ImageNames = v
	return s
}

func (s *ListImagesRequest) SetPageNumber(v int64) *ListImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int64) *ListImagesRequest {
	s.PageSize = &v
	return s
}

type ListImagesShrinkRequest struct {
	ImageIdsShrink   *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	ImageNamesShrink *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	PageNumber       *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImagesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListImagesShrinkRequest) SetImageIdsShrink(v string) *ListImagesShrinkRequest {
	s.ImageIdsShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetImageNamesShrink(v string) *ListImagesShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetPageNumber(v int64) *ListImagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesShrinkRequest) SetPageSize(v int64) *ListImagesShrinkRequest {
	s.PageSize = &v
	return s
}

type ListImagesResponseBody struct {
	Images     []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	PageNumber *int64                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetPageNumber(v int64) *ListImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListImagesResponseBody) SetPageSize(v int64) *ListImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetSuccess(v bool) *ListImagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int32) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListImagesResponseBodyImages struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType   *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetCreateTime(v string) *ListImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageType(v string) *ListImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetName(v string) *ListImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetVersion(v string) *ListImagesResponseBodyImages {
	s.Version = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListJobExecutorsRequest struct {
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Task ID
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListJobExecutorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutorsRequest) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsRequest) SetJobId(v string) *ListJobExecutorsRequest {
	s.JobId = &v
	return s
}

func (s *ListJobExecutorsRequest) SetPageNumber(v string) *ListJobExecutorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobExecutorsRequest) SetPageSize(v string) *ListJobExecutorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutorsRequest) SetTaskName(v string) *ListJobExecutorsRequest {
	s.TaskName = &v
	return s
}

type ListJobExecutorsResponseBody struct {
	Executors  []*ListJobExecutorsResponseBodyExecutors `json:"Executors,omitempty" xml:"Executors,omitempty" type:"Repeated"`
	JobId      *string                                  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber *string                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobExecutorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBody) SetExecutors(v []*ListJobExecutorsResponseBodyExecutors) *ListJobExecutorsResponseBody {
	s.Executors = v
	return s
}

func (s *ListJobExecutorsResponseBody) SetJobId(v string) *ListJobExecutorsResponseBody {
	s.JobId = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetPageNumber(v string) *ListJobExecutorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetPageSize(v string) *ListJobExecutorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetRequestId(v string) *ListJobExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetTaskName(v string) *ListJobExecutorsResponseBody {
	s.TaskName = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetTotalCount(v string) *ListJobExecutorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobExecutorsResponseBodyExecutors struct {
	ArrayIndex   *int32    `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	CreateTime   *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime      *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HostName     []*string `json:"HostName,omitempty" xml:"HostName,omitempty" type:"Repeated"`
	IpAddress    []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusReason *string   `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s ListJobExecutorsResponseBodyExecutors) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutorsResponseBodyExecutors) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBodyExecutors) SetArrayIndex(v int32) *ListJobExecutorsResponseBodyExecutors {
	s.ArrayIndex = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetCreateTime(v string) *ListJobExecutorsResponseBodyExecutors {
	s.CreateTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetEndTime(v string) *ListJobExecutorsResponseBodyExecutors {
	s.EndTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetHostName(v []*string) *ListJobExecutorsResponseBodyExecutors {
	s.HostName = v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetIpAddress(v []*string) *ListJobExecutorsResponseBodyExecutors {
	s.IpAddress = v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetStatus(v string) *ListJobExecutorsResponseBodyExecutors {
	s.Status = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetStatusReason(v string) *ListJobExecutorsResponseBodyExecutors {
	s.StatusReason = &v
	return s
}

type ListJobExecutorsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobExecutorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutorsResponse) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponse) SetHeaders(v map[string]*string) *ListJobExecutorsResponse {
	s.Headers = v
	return s
}

func (s *ListJobExecutorsResponse) SetStatusCode(v int32) *ListJobExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobExecutorsResponse) SetBody(v *ListJobExecutorsResponseBody) *ListJobExecutorsResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	Filter     *ListJobsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	PageNumber *string                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *ListJobsRequestSortBy `json:"SortBy,omitempty" xml:"SortBy,omitempty" type:"Struct"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetFilter(v *ListJobsRequestFilter) *ListJobsRequest {
	s.Filter = v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v string) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v string) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetSortBy(v *ListJobsRequestSortBy) *ListJobsRequest {
	s.SortBy = v
	return s
}

type ListJobsRequestFilter struct {
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName           *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeCreatedAfter  *int32  `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	TimeCreatedBefore *int32  `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
}

func (s ListJobsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListJobsRequestFilter) SetJobId(v string) *ListJobsRequestFilter {
	s.JobId = &v
	return s
}

func (s *ListJobsRequestFilter) SetJobName(v string) *ListJobsRequestFilter {
	s.JobName = &v
	return s
}

func (s *ListJobsRequestFilter) SetStatus(v string) *ListJobsRequestFilter {
	s.Status = &v
	return s
}

func (s *ListJobsRequestFilter) SetTimeCreatedAfter(v int32) *ListJobsRequestFilter {
	s.TimeCreatedAfter = &v
	return s
}

func (s *ListJobsRequestFilter) SetTimeCreatedBefore(v int32) *ListJobsRequestFilter {
	s.TimeCreatedBefore = &v
	return s
}

type ListJobsRequestSortBy struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListJobsRequestSortBy) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequestSortBy) GoString() string {
	return s.String()
}

func (s *ListJobsRequestSortBy) SetLabel(v string) *ListJobsRequestSortBy {
	s.Label = &v
	return s
}

func (s *ListJobsRequestSortBy) SetOrder(v string) *ListJobsRequestSortBy {
	s.Order = &v
	return s
}

type ListJobsShrinkRequest struct {
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortByShrink *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) SetFilterShrink(v string) *ListJobsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageNumber(v string) *ListJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageSize(v string) *ListJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsShrinkRequest) SetSortByShrink(v string) *ListJobsShrinkRequest {
	s.SortByShrink = &v
	return s
}

type ListJobsResponseBody struct {
	JobList    []*ListJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	PageNumber *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobList(v []*ListJobsResponseBodyJobList) *ListJobsResponseBody {
	s.JobList = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponseBodyJobList struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutorCount   *int32  `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	JobDescription  *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	JobId           *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName         *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	OwnerUid        *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskCount       *int32  `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	TaskSustainable *bool   `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s ListJobsResponseBodyJobList) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobList) SetCreateTime(v string) *ListJobsResponseBodyJobList {
	s.CreateTime = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetEndTime(v string) *ListJobsResponseBodyJobList {
	s.EndTime = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetExecutorCount(v int32) *ListJobsResponseBodyJobList {
	s.ExecutorCount = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetJobDescription(v string) *ListJobsResponseBodyJobList {
	s.JobDescription = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetJobId(v string) *ListJobsResponseBodyJobList {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetJobName(v string) *ListJobsResponseBodyJobList {
	s.JobName = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetOwnerUid(v string) *ListJobsResponseBodyJobList {
	s.OwnerUid = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetStartTime(v string) *ListJobsResponseBodyJobList {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetStatus(v string) *ListJobsResponseBodyJobList {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetTaskCount(v int32) *ListJobsResponseBodyJobList {
	s.TaskCount = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetTaskSustainable(v bool) *ListJobsResponseBodyJobList {
	s.TaskSustainable = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type RemoveImageRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s RemoveImageRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageRequest) SetImageId(v string) *RemoveImageRequest {
	s.ImageId = &v
	return s
}

type RemoveImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageResponseBody) SetRequestId(v string) *RemoveImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveImageResponseBody) SetSuccess(v bool) *RemoveImageResponseBody {
	s.Success = &v
	return s
}

type RemoveImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageResponse) SetHeaders(v map[string]*string) *RemoveImageResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageResponse) SetStatusCode(v int32) *RemoveImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageResponse) SetBody(v *RemoveImageResponseBody) *RemoveImageResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ehpc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddImageWithOptions(tmpReq *AddImageRequest, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ContainerImageSpec)) {
		request.ContainerImageSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContainerImageSpec, tea.String("ContainerImageSpec"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VMImageSpec)) {
		request.VMImageSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VMImageSpec, tea.String("VMImageSpec"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerImageSpecShrink)) {
		query["ContainerImageSpec"] = request.ContainerImageSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.VMImageSpecShrink)) {
		query["VMImageSpec"] = request.VMImageSpecShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImage"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobWithOptions(tmpReq *CreateJobRequest, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeploymentPolicy)) {
		request.DeploymentPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentPolicy, tea.String("DeploymentPolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tasks)) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, tea.String("Tasks"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeploymentPolicyShrink)) {
		query["DeploymentPolicy"] = request.DeploymentPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.JobDescription)) {
		query["JobDescription"] = request.JobDescription
	}

	if !tea.BoolValue(util.IsUnset(request.JobName)) {
		query["JobName"] = request.JobName
	}

	if !tea.BoolValue(util.IsUnset(request.TasksShrink)) {
		query["Tasks"] = request.TasksShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobsWithOptions(tmpReq *DeleteJobsRequest, runtime *util.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExecutorIds)) {
		request.ExecutorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutorIds, tea.String("ExecutorIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.JobSpec)) {
		request.JobSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSpec, tea.String("JobSpec"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutorIdsShrink)) {
		query["ExecutorIds"] = request.ExecutorIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.JobSpecShrink)) {
		query["JobSpec"] = request.JobSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobs"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobs(request *DeleteJobsRequest) (_result *DeleteJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobsResponse{}
	_body, _err := client.DeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageWithOptions(request *GetImageRequest, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImage"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImage(request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobWithOptions(request *GetJobRequest, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExecutorsWithOptions(tmpReq *ListExecutorsRequest, runtime *util.RuntimeOptions) (_result *ListExecutorsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListExecutorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filter)) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, tea.String("Filter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterShrink)) {
		query["Filter"] = request.FilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExecutors"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExecutorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExecutors(request *ListExecutorsRequest) (_result *ListExecutorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutorsResponse{}
	_body, _err := client.ListExecutorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(tmpReq *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListImagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ImageIds)) {
		request.ImageIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageIds, tea.String("ImageIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ImageNames)) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, tea.String("ImageNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageIdsShrink)) {
		query["ImageIds"] = request.ImageIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageNamesShrink)) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobExecutorsWithOptions(request *ListJobExecutorsRequest, runtime *util.RuntimeOptions) (_result *ListJobExecutorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobExecutors"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobExecutorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobExecutors(request *ListJobExecutorsRequest) (_result *ListJobExecutorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobExecutorsResponse{}
	_body, _err := client.ListJobExecutorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(tmpReq *ListJobsRequest, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filter)) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, tea.String("Filter"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SortBy)) {
		request.SortByShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SortBy, tea.String("SortBy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterShrink)) {
		query["Filter"] = request.FilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortByShrink)) {
		query["SortBy"] = request.SortByShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveImageWithOptions(request *RemoveImageRequest, runtime *util.RuntimeOptions) (_result *RemoveImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImage"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveImage(request *RemoveImageRequest) (_result *RemoveImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveImageResponse{}
	_body, _err := client.RemoveImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
