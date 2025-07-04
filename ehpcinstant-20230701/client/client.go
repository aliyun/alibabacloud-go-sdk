// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddImageRequest struct {
	ContainerImageSpec *AddImageRequestContainerImageSpec `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty" type:"Struct"`
	Description        *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageType          *string                            `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// V1.0
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// app-image
	Name        *string                     `json:"Name,omitempty" xml:"Name,omitempty"`
	VMImageSpec *AddImageRequestVMImageSpec `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty" type:"Struct"`
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

func (s *AddImageRequest) SetImageType(v string) *AddImageRequest {
	s.ImageType = &v
	return s
}

func (s *AddImageRequest) SetImageVersion(v string) *AddImageRequest {
	s.ImageVersion = &v
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

type AddImageRequestContainerImageSpec struct {
	// example:
	//
	// True
	IsACREnterprise *bool `json:"IsACREnterprise,omitempty" xml:"IsACREnterprise,omitempty"`
	// example:
	//
	// True
	IsACRRegistry      *bool                                                `json:"IsACRRegistry,omitempty" xml:"IsACRRegistry,omitempty"`
	RegistryCredential *AddImageRequestContainerImageSpecRegistryCredential `json:"RegistryCredential,omitempty" xml:"RegistryCredential,omitempty" type:"Struct"`
	// example:
	//
	// cri-xyz795ygf8k9****
	RegistryCriId *string `json:"RegistryCriId,omitempty" xml:"RegistryCriId,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc_open/nginx:latest
	RegistryUrl *string `json:"RegistryUrl,omitempty" xml:"RegistryUrl,omitempty"`
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
	// example:
	//
	// userpassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// example:
	//
	// username
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
	// example:
	//
	// m-bp1akkkr1rkxtb******
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
	ImageType                *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// V1.0
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// app-image
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VMImageSpecShrink *string `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty"`
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

func (s *AddImageShrinkRequest) SetImageType(v string) *AddImageShrinkRequest {
	s.ImageType = &v
	return s
}

func (s *AddImageShrinkRequest) SetImageVersion(v string) *AddImageShrinkRequest {
	s.ImageVersion = &v
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

type AddImageResponseBody struct {
	// example:
	//
	// m-bp1akkkr1rkxtb******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testjob
	JobName        *string                         `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobScheduler   *string                         `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	SecurityPolicy *CreateJobRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// This parameter is required.
	Tasks []*CreateJobRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
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

func (s *CreateJobRequest) SetJobScheduler(v string) *CreateJobRequest {
	s.JobScheduler = &v
	return s
}

func (s *CreateJobRequest) SetSecurityPolicy(v *CreateJobRequestSecurityPolicy) *CreateJobRequest {
	s.SecurityPolicy = v
	return s
}

func (s *CreateJobRequest) SetTasks(v []*CreateJobRequestTasks) *CreateJobRequest {
	s.Tasks = v
	return s
}

type CreateJobRequestDeploymentPolicy struct {
	// example:
	//
	// Dedicated
	AllocationSpec *string                                  `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	Level          *string                                  `json:"Level,omitempty" xml:"Level,omitempty"`
	Network        *CreateJobRequestDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	Pool           *string                                  `json:"Pool,omitempty" xml:"Pool,omitempty"`
	Priority       *int32                                   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Tag            []*CreateJobRequestDeploymentPolicyTag   `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateJobRequestDeploymentPolicy) SetLevel(v string) *CreateJobRequestDeploymentPolicy {
	s.Level = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetNetwork(v *CreateJobRequestDeploymentPolicyNetwork) *CreateJobRequestDeploymentPolicy {
	s.Network = v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetPool(v string) *CreateJobRequestDeploymentPolicy {
	s.Pool = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetPriority(v int32) *CreateJobRequestDeploymentPolicy {
	s.Priority = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetTag(v []*CreateJobRequestDeploymentPolicyTag) *CreateJobRequestDeploymentPolicy {
	s.Tag = v
	return s
}

type CreateJobRequestDeploymentPolicyNetwork struct {
	EnableExternalIpAddress *bool     `json:"EnableExternalIpAddress,omitempty" xml:"EnableExternalIpAddress,omitempty"`
	Vswitch                 []*string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
}

func (s CreateJobRequestDeploymentPolicyNetwork) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDeploymentPolicyNetwork) SetEnableExternalIpAddress(v bool) *CreateJobRequestDeploymentPolicyNetwork {
	s.EnableExternalIpAddress = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicyNetwork) SetVswitch(v []*string) *CreateJobRequestDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

type CreateJobRequestDeploymentPolicyTag struct {
	// This parameter is required.
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobRequestDeploymentPolicyTag) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestDeploymentPolicyTag) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDeploymentPolicyTag) SetKey(v string) *CreateJobRequestDeploymentPolicyTag {
	s.Key = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicyTag) SetValue(v string) *CreateJobRequestDeploymentPolicyTag {
	s.Value = &v
	return s
}

type CreateJobRequestSecurityPolicy struct {
	SecurityGroup *CreateJobRequestSecurityPolicySecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Struct"`
}

func (s CreateJobRequestSecurityPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobRequestSecurityPolicy) SetSecurityGroup(v *CreateJobRequestSecurityPolicySecurityGroup) *CreateJobRequestSecurityPolicy {
	s.SecurityGroup = v
	return s
}

type CreateJobRequestSecurityPolicySecurityGroup struct {
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s CreateJobRequestSecurityPolicySecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestSecurityPolicySecurityGroup) GoString() string {
	return s.String()
}

func (s *CreateJobRequestSecurityPolicySecurityGroup) SetSecurityGroupIds(v []*string) *CreateJobRequestSecurityPolicySecurityGroup {
	s.SecurityGroupIds = v
	return s
}

type CreateJobRequestTasks struct {
	ExecutorPolicy *CreateJobRequestTasksExecutorPolicy `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	// example:
	//
	// task0
	TaskName *string                        `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskSpec *CreateJobRequestTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	// example:
	//
	// true
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
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
	// example:
	//
	// 1
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
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
	// example:
	//
	// 9
	IndexEnd *int32 `json:"IndexEnd,omitempty" xml:"IndexEnd,omitempty"`
	// example:
	//
	// 0
	IndexStart *int32 `json:"IndexStart,omitempty" xml:"IndexStart,omitempty"`
	// example:
	//
	// 1
	IndexStep *int32 `json:"IndexStep,omitempty" xml:"IndexStep,omitempty"`
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
	Resource *CreateJobRequestTasksTaskSpecResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// This parameter is required.
	TaskExecutor []*CreateJobRequestTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
	VolumeMount  []*CreateJobRequestTasksTaskSpecVolumeMount  `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
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

func (s *CreateJobRequestTasksTaskSpec) SetVolumeMount(v []*CreateJobRequestTasksTaskSpecVolumeMount) *CreateJobRequestTasksTaskSpec {
	s.VolumeMount = v
	return s
}

type CreateJobRequestTasksTaskSpecResource struct {
	// example:
	//
	// 2
	Cores         *float32                                      `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Disks         []*CreateJobRequestTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	InstanceTypes []*string                                     `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
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

func (s *CreateJobRequestTasksTaskSpecResource) SetInstanceTypes(v []*string) *CreateJobRequestTasksTaskSpecResource {
	s.InstanceTypes = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetMemory(v float32) *CreateJobRequestTasksTaskSpecResource {
	s.Memory = &v
	return s
}

type CreateJobRequestTasksTaskSpecResourceDisks struct {
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// System
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
	Container *CreateJobRequestTasksTaskSpecTaskExecutorContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	VM        *CreateJobRequestTasksTaskSpecTaskExecutorVM        `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutor) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutor) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutor) SetContainer(v *CreateJobRequestTasksTaskSpecTaskExecutorContainer) *CreateJobRequestTasksTaskSpecTaskExecutor {
	s.Container = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutor) SetVM(v *CreateJobRequestTasksTaskSpecTaskExecutorVM) *CreateJobRequestTasksTaskSpecTaskExecutor {
	s.VM = v
	return s
}

type CreateJobRequestTasksTaskSpecTaskExecutorContainer struct {
	AppId           *string                                                              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Arg             []*string                                                            `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command         []*string                                                            `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	EnvironmentVars []*CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc/hpl:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// /usr/local/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorContainer) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorContainer) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetAppId(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.AppId = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetArg(v []*string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.Arg = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetCommand(v []*string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.Command = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetEnvironmentVars(v []*CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.EnvironmentVars = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetImage(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.Image = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetWorkingDir(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.WorkingDir = &v
	return s
}

type CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars struct {
	// example:
	//
	// PATH
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /usr/local/bin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetName(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Name = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetValue(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Value = &v
	return s
}

type CreateJobRequestTasksTaskSpecTaskExecutorVM struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// m-xxxx
	Image    *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorVM) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorVM) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetAppId(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.AppId = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetImage(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.Image = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetPassword(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.Password = &v
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

type CreateJobRequestTasksTaskSpecVolumeMount struct {
	// example:
	//
	// {"server":"xxxxx-xxxxx.cn-heyuan.nas.aliyuncs.com","vers":"3","path":"/data","options":"nolock,tcp,noresvport"}
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// example:
	//
	// /mnt
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly  *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// example:
	//
	// alicloud/nas
	VolumeDriver *string `json:"VolumeDriver,omitempty" xml:"VolumeDriver,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) SetMountOptions(v string) *CreateJobRequestTasksTaskSpecVolumeMount {
	s.MountOptions = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) SetMountPath(v string) *CreateJobRequestTasksTaskSpecVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) SetReadOnly(v bool) *CreateJobRequestTasksTaskSpecVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) SetVolumeDriver(v string) *CreateJobRequestTasksTaskSpecVolumeMount {
	s.VolumeDriver = &v
	return s
}

type CreateJobShrinkRequest struct {
	DeploymentPolicyShrink *string `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty"`
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testjob
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobScheduler         *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	// This parameter is required.
	TasksShrink *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
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

func (s *CreateJobShrinkRequest) SetJobScheduler(v string) *CreateJobShrinkRequest {
	s.JobScheduler = &v
	return s
}

func (s *CreateJobShrinkRequest) SetSecurityPolicyShrink(v string) *CreateJobShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) SetTasksShrink(v string) *CreateJobShrinkRequest {
	s.TasksShrink = &v
	return s
}

type CreateJobResponseBody struct {
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*CreateJobResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
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

func (s *CreateJobResponseBody) SetTasks(v []*CreateJobResponseBodyTasks) *CreateJobResponseBody {
	s.Tasks = v
	return s
}

type CreateJobResponseBodyTasks struct {
	ExecutorIds []*string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	TaskName    *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateJobResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBodyTasks) SetExecutorIds(v []*string) *CreateJobResponseBodyTasks {
	s.ExecutorIds = v
	return s
}

func (s *CreateJobResponseBodyTasks) SetTaskName(v string) *CreateJobResponseBodyTasks {
	s.TaskName = &v
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

type CreatePoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority       *int32                           `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceLimits *CreatePoolRequestResourceLimits `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty" type:"Struct"`
}

func (s CreatePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePoolRequest) GoString() string {
	return s.String()
}

func (s *CreatePoolRequest) SetPoolName(v string) *CreatePoolRequest {
	s.PoolName = &v
	return s
}

func (s *CreatePoolRequest) SetPriority(v int32) *CreatePoolRequest {
	s.Priority = &v
	return s
}

func (s *CreatePoolRequest) SetResourceLimits(v *CreatePoolRequestResourceLimits) *CreatePoolRequest {
	s.ResourceLimits = v
	return s
}

type CreatePoolRequestResourceLimits struct {
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
}

func (s CreatePoolRequestResourceLimits) String() string {
	return tea.Prettify(s)
}

func (s CreatePoolRequestResourceLimits) GoString() string {
	return s.String()
}

func (s *CreatePoolRequestResourceLimits) SetMaxExectorNum(v int32) *CreatePoolRequestResourceLimits {
	s.MaxExectorNum = &v
	return s
}

type CreatePoolShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority             *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceLimitsShrink *string `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
}

func (s CreatePoolShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePoolShrinkRequest) SetPoolName(v string) *CreatePoolShrinkRequest {
	s.PoolName = &v
	return s
}

func (s *CreatePoolShrinkRequest) SetPriority(v int32) *CreatePoolShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreatePoolShrinkRequest) SetResourceLimitsShrink(v string) *CreatePoolShrinkRequest {
	s.ResourceLimitsShrink = &v
	return s
}

type CreatePoolResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePoolResponseBody) SetRequestId(v string) *CreatePoolResponseBody {
	s.RequestId = &v
	return s
}

type CreatePoolResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePoolResponse) GoString() string {
	return s.String()
}

func (s *CreatePoolResponse) SetHeaders(v map[string]*string) *CreatePoolResponse {
	s.Headers = v
	return s
}

func (s *CreatePoolResponse) SetStatusCode(v int32) *CreatePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePoolResponse) SetBody(v *CreatePoolResponseBody) *CreatePoolResponse {
	s.Body = v
	return s
}

type DeleteJobsRequest struct {
	ExecutorIds  []*string                   `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	JobScheduler *string                     `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	JobSpec      []*DeleteJobsRequestJobSpec `json:"JobSpec,omitempty" xml:"JobSpec,omitempty" type:"Repeated"`
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

func (s *DeleteJobsRequest) SetJobScheduler(v string) *DeleteJobsRequest {
	s.JobScheduler = &v
	return s
}

func (s *DeleteJobsRequest) SetJobSpec(v []*DeleteJobsRequestJobSpec) *DeleteJobsRequest {
	s.JobSpec = v
	return s
}

type DeleteJobsRequestJobSpec struct {
	// example:
	//
	// job-xxxx
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
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
	JobScheduler      *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
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

func (s *DeleteJobsShrinkRequest) SetJobScheduler(v string) *DeleteJobsShrinkRequest {
	s.JobScheduler = &v
	return s
}

func (s *DeleteJobsShrinkRequest) SetJobSpecShrink(v string) *DeleteJobsShrinkRequest {
	s.JobSpecShrink = &v
	return s
}

type DeleteJobsResponseBody struct {
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
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

type DeletePoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
}

func (s DeletePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePoolRequest) GoString() string {
	return s.String()
}

func (s *DeletePoolRequest) SetPoolName(v string) *DeletePoolRequest {
	s.PoolName = &v
	return s
}

type DeletePoolResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePoolResponseBody) SetRequestId(v string) *DeletePoolResponseBody {
	s.RequestId = &v
	return s
}

type DeletePoolResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePoolResponse) GoString() string {
	return s.String()
}

func (s *DeletePoolResponse) SetHeaders(v map[string]*string) *DeletePoolResponse {
	s.Headers = v
	return s
}

func (s *DeletePoolResponse) SetStatusCode(v int32) *DeletePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePoolResponse) SetBody(v *DeletePoolResponseBody) *DeletePoolResponse {
	s.Body = v
	return s
}

type DescribeJobMetricDataRequest struct {
	ArrayIndex []*int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty" type:"Repeated"`
	// example:
	//
	// job-xxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// cpu_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// Task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeJobMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricDataRequest) SetArrayIndex(v []*int32) *DescribeJobMetricDataRequest {
	s.ArrayIndex = v
	return s
}

func (s *DescribeJobMetricDataRequest) SetJobId(v string) *DescribeJobMetricDataRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobMetricDataRequest) SetMetricName(v string) *DescribeJobMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeJobMetricDataRequest) SetTaskName(v string) *DescribeJobMetricDataRequest {
	s.TaskName = &v
	return s
}

type DescribeJobMetricDataShrinkRequest struct {
	ArrayIndexShrink *string `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// example:
	//
	// job-xxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// cpu_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// Task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeJobMetricDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricDataShrinkRequest) SetArrayIndexShrink(v string) *DescribeJobMetricDataShrinkRequest {
	s.ArrayIndexShrink = &v
	return s
}

func (s *DescribeJobMetricDataShrinkRequest) SetJobId(v string) *DescribeJobMetricDataShrinkRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobMetricDataShrinkRequest) SetMetricName(v string) *DescribeJobMetricDataShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeJobMetricDataShrinkRequest) SetTaskName(v string) *DescribeJobMetricDataShrinkRequest {
	s.TaskName = &v
	return s
}

type DescribeJobMetricDataResponseBody struct {
	// example:
	//
	// [{"timestamp":1709540685000,"Minimum":28.45,"Maximum":28.45,"Average":28.45}]
	DataPoints *string `json:"DataPoints,omitempty" xml:"DataPoints,omitempty"`
	// example:
	//
	// 15
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricDataResponseBody) SetDataPoints(v string) *DescribeJobMetricDataResponseBody {
	s.DataPoints = &v
	return s
}

func (s *DescribeJobMetricDataResponseBody) SetPeriod(v int32) *DescribeJobMetricDataResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeJobMetricDataResponseBody) SetRequestId(v string) *DescribeJobMetricDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeJobMetricDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricDataResponse) SetHeaders(v map[string]*string) *DescribeJobMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobMetricDataResponse) SetStatusCode(v int32) *DescribeJobMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobMetricDataResponse) SetBody(v *DescribeJobMetricDataResponseBody) *DescribeJobMetricDataResponse {
	s.Body = v
	return s
}

type DescribeJobMetricLastRequest struct {
	ArrayIndex []*int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty" type:"Repeated"`
	// example:
	//
	// job-xxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// Task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeJobMetricLastRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricLastRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastRequest) SetArrayIndex(v []*int32) *DescribeJobMetricLastRequest {
	s.ArrayIndex = v
	return s
}

func (s *DescribeJobMetricLastRequest) SetJobId(v string) *DescribeJobMetricLastRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobMetricLastRequest) SetTaskName(v string) *DescribeJobMetricLastRequest {
	s.TaskName = &v
	return s
}

type DescribeJobMetricLastShrinkRequest struct {
	ArrayIndexShrink *string `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// example:
	//
	// job-xxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// Task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeJobMetricLastShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricLastShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastShrinkRequest) SetArrayIndexShrink(v string) *DescribeJobMetricLastShrinkRequest {
	s.ArrayIndexShrink = &v
	return s
}

func (s *DescribeJobMetricLastShrinkRequest) SetJobId(v string) *DescribeJobMetricLastShrinkRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobMetricLastShrinkRequest) SetTaskName(v string) *DescribeJobMetricLastShrinkRequest {
	s.TaskName = &v
	return s
}

type DescribeJobMetricLastResponseBody struct {
	Metrics []*DescribeJobMetricLastResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobMetricLastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricLastResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastResponseBody) SetMetrics(v []*DescribeJobMetricLastResponseBodyMetrics) *DescribeJobMetricLastResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribeJobMetricLastResponseBody) SetRequestId(v string) *DescribeJobMetricLastResponseBody {
	s.RequestId = &v
	return s
}

type DescribeJobMetricLastResponseBodyMetrics struct {
	// example:
	//
	// 1
	ArrayIndex *int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// example:
	//
	// {"memory_utilization": 37.42,"cpu_utilization": 1.008, "disk_utilization": 3.562}
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
}

func (s DescribeJobMetricLastResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricLastResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastResponseBodyMetrics) SetArrayIndex(v int32) *DescribeJobMetricLastResponseBodyMetrics {
	s.ArrayIndex = &v
	return s
}

func (s *DescribeJobMetricLastResponseBodyMetrics) SetMetric(v string) *DescribeJobMetricLastResponseBodyMetrics {
	s.Metric = &v
	return s
}

type DescribeJobMetricLastResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobMetricLastResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobMetricLastResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMetricLastResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastResponse) SetHeaders(v map[string]*string) *DescribeJobMetricLastResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobMetricLastResponse) SetStatusCode(v int32) *DescribeJobMetricLastResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobMetricLastResponse) SetBody(v *DescribeJobMetricLastResponseBody) *DescribeJobMetricLastResponse {
	s.Body = v
	return s
}

type GetAppVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// V-Ray
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// Public
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetAppVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppVersionsRequest) GoString() string {
	return s.String()
}

func (s *GetAppVersionsRequest) SetAppName(v string) *GetAppVersionsRequest {
	s.AppName = &v
	return s
}

func (s *GetAppVersionsRequest) SetImageCategory(v string) *GetAppVersionsRequest {
	s.ImageCategory = &v
	return s
}

func (s *GetAppVersionsRequest) SetImageType(v string) *GetAppVersionsRequest {
	s.ImageType = &v
	return s
}

func (s *GetAppVersionsRequest) SetPageNumber(v int64) *GetAppVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAppVersionsRequest) SetPageSize(v int64) *GetAppVersionsRequest {
	s.PageSize = &v
	return s
}

type GetAppVersionsResponseBody struct {
	AppVersions []*GetAppVersionsResponseBodyAppVersions `json:"AppVersions,omitempty" xml:"AppVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAppVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppVersionsResponseBody) SetAppVersions(v []*GetAppVersionsResponseBodyAppVersions) *GetAppVersionsResponseBody {
	s.AppVersions = v
	return s
}

func (s *GetAppVersionsResponseBody) SetPageNumber(v int64) *GetAppVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetAppVersionsResponseBody) SetPageSize(v int64) *GetAppVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAppVersionsResponseBody) SetRequestId(v string) *GetAppVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppVersionsResponseBody) SetSuccess(v bool) *GetAppVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppVersionsResponseBody) SetTotalCount(v int32) *GetAppVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type GetAppVersionsResponseBodyAppVersions struct {
	// This parameter is required.
	//
	// example:
	//
	// m-f8z0dfa96luomqly****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// V-Ray
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// v1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAppVersionsResponseBodyAppVersions) String() string {
	return tea.Prettify(s)
}

func (s GetAppVersionsResponseBodyAppVersions) GoString() string {
	return s.String()
}

func (s *GetAppVersionsResponseBodyAppVersions) SetImageId(v string) *GetAppVersionsResponseBodyAppVersions {
	s.ImageId = &v
	return s
}

func (s *GetAppVersionsResponseBodyAppVersions) SetName(v string) *GetAppVersionsResponseBodyAppVersions {
	s.Name = &v
	return s
}

func (s *GetAppVersionsResponseBodyAppVersions) SetVersion(v string) *GetAppVersionsResponseBodyAppVersions {
	s.Version = &v
	return s
}

type GetAppVersionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppVersionsResponse) GoString() string {
	return s.String()
}

func (s *GetAppVersionsResponse) SetHeaders(v map[string]*string) *GetAppVersionsResponse {
	s.Headers = v
	return s
}

func (s *GetAppVersionsResponse) SetStatusCode(v int32) *GetAppVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppVersionsResponse) SetBody(v *GetAppVersionsResponseBody) *GetAppVersionsResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	AdditionalRegionIds []*string `json:"AdditionalRegionIds,omitempty" xml:"AdditionalRegionIds,omitempty" type:"Repeated"`
	ImageCategory       *string   `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// example:
	//
	// m-2ze74g5mvy4pjg*****
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetAdditionalRegionIds(v []*string) *GetImageRequest {
	s.AdditionalRegionIds = v
	return s
}

func (s *GetImageRequest) SetImageCategory(v string) *GetImageRequest {
	s.ImageCategory = &v
	return s
}

func (s *GetImageRequest) SetImageId(v string) *GetImageRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageRequest) SetImageType(v string) *GetImageRequest {
	s.ImageType = &v
	return s
}

type GetImageShrinkRequest struct {
	AdditionalRegionIdsShrink *string `json:"AdditionalRegionIds,omitempty" xml:"AdditionalRegionIds,omitempty"`
	ImageCategory             *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// example:
	//
	// m-2ze74g5mvy4pjg*****
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s GetImageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetImageShrinkRequest) SetAdditionalRegionIdsShrink(v string) *GetImageShrinkRequest {
	s.AdditionalRegionIdsShrink = &v
	return s
}

func (s *GetImageShrinkRequest) SetImageCategory(v string) *GetImageShrinkRequest {
	s.ImageCategory = &v
	return s
}

func (s *GetImageShrinkRequest) SetImageId(v string) *GetImageShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageShrinkRequest) SetImageType(v string) *GetImageShrinkRequest {
	s.ImageType = &v
	return s
}

type GetImageResponseBody struct {
	Image *GetImageResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AdditionalRegionsInfo []*GetImageResponseBodyImageAdditionalRegionsInfo `json:"AdditionalRegionsInfo,omitempty" xml:"AdditionalRegionsInfo,omitempty" type:"Repeated"`
	AppId                 *string                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ContainerImageSpec    *GetImageResponseBodyImageContainerImageSpec      `json:"ContainerImageSpec,omitempty" xml:"ContainerImageSpec,omitempty" type:"Struct"`
	// example:
	//
	// 2022-12-23T09:51:39Z
	CreateTime   *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentInfo *GetImageResponseBodyImageDocumentInfo `json:"DocumentInfo,omitempty" xml:"DocumentInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// app-image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 40 GiB
	Size        *string                               `json:"Size,omitempty" xml:"Size,omitempty"`
	Status      *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	VMImageSpec *GetImageResponseBodyImageVMImageSpec `json:"VMImageSpec,omitempty" xml:"VMImageSpec,omitempty" type:"Struct"`
	// example:
	//
	// v1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetImageResponseBodyImage) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImage) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImage) SetAdditionalRegionsInfo(v []*GetImageResponseBodyImageAdditionalRegionsInfo) *GetImageResponseBodyImage {
	s.AdditionalRegionsInfo = v
	return s
}

func (s *GetImageResponseBodyImage) SetAppId(v string) *GetImageResponseBodyImage {
	s.AppId = &v
	return s
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

func (s *GetImageResponseBodyImage) SetDocumentInfo(v *GetImageResponseBodyImageDocumentInfo) *GetImageResponseBodyImage {
	s.DocumentInfo = v
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

func (s *GetImageResponseBodyImage) SetStatus(v string) *GetImageResponseBodyImage {
	s.Status = &v
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

type GetImageResponseBodyImageAdditionalRegionsInfo struct {
	ImageId  *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetImageResponseBodyImageAdditionalRegionsInfo) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImageAdditionalRegionsInfo) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) SetImageId(v string) *GetImageResponseBodyImageAdditionalRegionsInfo {
	s.ImageId = &v
	return s
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) SetRegionId(v string) *GetImageResponseBodyImageAdditionalRegionsInfo {
	s.RegionId = &v
	return s
}

func (s *GetImageResponseBodyImageAdditionalRegionsInfo) SetStatus(v string) *GetImageResponseBodyImageAdditionalRegionsInfo {
	s.Status = &v
	return s
}

type GetImageResponseBodyImageContainerImageSpec struct {
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// True
	IsACREnterprise *bool `json:"IsACREnterprise,omitempty" xml:"IsACREnterprise,omitempty"`
	// example:
	//
	// True
	IsACRRegistry      *bool                                                          `json:"IsACRRegistry,omitempty" xml:"IsACRRegistry,omitempty"`
	OsTag              *string                                                        `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Platform           *string                                                        `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RegistryCredential *GetImageResponseBodyImageContainerImageSpecRegistryCredential `json:"RegistryCredential,omitempty" xml:"RegistryCredential,omitempty" type:"Struct"`
	// example:
	//
	// cri-xyz795ygf8k9****
	RegistryCriId *string `json:"RegistryCriId,omitempty" xml:"RegistryCriId,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc_open/nginx:latest
	RegistryUrl *string `json:"RegistryUrl,omitempty" xml:"RegistryUrl,omitempty"`
}

func (s GetImageResponseBodyImageContainerImageSpec) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImageContainerImageSpec) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetArchitecture(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.Architecture = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetIsACREnterprise(v bool) *GetImageResponseBodyImageContainerImageSpec {
	s.IsACREnterprise = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetIsACRRegistry(v bool) *GetImageResponseBodyImageContainerImageSpec {
	s.IsACRRegistry = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetOsTag(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.OsTag = &v
	return s
}

func (s *GetImageResponseBodyImageContainerImageSpec) SetPlatform(v string) *GetImageResponseBodyImageContainerImageSpec {
	s.Platform = &v
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
	// example:
	//
	// userpassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// example:
	//
	// username
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

type GetImageResponseBodyImageDocumentInfo struct {
	Document     *string `json:"Document,omitempty" xml:"Document,omitempty"`
	DocumentId   *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	EncodingMode *string `json:"EncodingMode,omitempty" xml:"EncodingMode,omitempty"`
}

func (s GetImageResponseBodyImageDocumentInfo) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImageDocumentInfo) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageDocumentInfo) SetDocument(v string) *GetImageResponseBodyImageDocumentInfo {
	s.Document = &v
	return s
}

func (s *GetImageResponseBodyImageDocumentInfo) SetDocumentId(v string) *GetImageResponseBodyImageDocumentInfo {
	s.DocumentId = &v
	return s
}

func (s *GetImageResponseBodyImageDocumentInfo) SetEncodingMode(v string) *GetImageResponseBodyImageDocumentInfo {
	s.EncodingMode = &v
	return s
}

type GetImageResponseBodyImageVMImageSpec struct {
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// m-uf60twafjtaart******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// CentOS  7.6 64 bit
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	// example:
	//
	// CentOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
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
	// example:
	//
	// job-xxxx
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
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
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
	AppExtraInfo *string `json:"AppExtraInfo,omitempty" xml:"AppExtraInfo,omitempty"`
	// example:
	//
	// 2024-03-05 20:00:46
	CreateTime       *string                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeploymentPolicy *GetJobResponseBodyJobInfoDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	// example:
	//
	// 2024-03-05 20:01:48
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// testJob
	JobName      *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobScheduler *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	// example:
	//
	// 2024-03-05 20:00:48
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Succeed
	Status *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tasks  []*GetJobResponseBodyJobInfoTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfo) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfo) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfo) SetAppExtraInfo(v string) *GetJobResponseBodyJobInfo {
	s.AppExtraInfo = &v
	return s
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

func (s *GetJobResponseBodyJobInfo) SetJobScheduler(v string) *GetJobResponseBodyJobInfo {
	s.JobScheduler = &v
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
	// example:
	//
	// Dedicated
	AllocationSpec *string                                           `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	Level          *string                                           `json:"Level,omitempty" xml:"Level,omitempty"`
	Network        *GetJobResponseBodyJobInfoDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	Tags           []*GetJobResponseBodyJobInfoDeploymentPolicyTags  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetLevel(v string) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.Level = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetNetwork(v *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.Network = v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetTags(v []*GetJobResponseBodyJobInfoDeploymentPolicyTags) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.Tags = v
	return s
}

type GetJobResponseBodyJobInfoDeploymentPolicyNetwork struct {
	EnableENIMapping        *bool     `json:"EnableENIMapping,omitempty" xml:"EnableENIMapping,omitempty"`
	EnableExternalIpAddress *bool     `json:"EnableExternalIpAddress,omitempty" xml:"EnableExternalIpAddress,omitempty"`
	Vswitch                 []*string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyNetwork) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) SetEnableENIMapping(v bool) *GetJobResponseBodyJobInfoDeploymentPolicyNetwork {
	s.EnableENIMapping = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) SetEnableExternalIpAddress(v bool) *GetJobResponseBodyJobInfoDeploymentPolicyNetwork {
	s.EnableExternalIpAddress = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) SetVswitch(v []*string) *GetJobResponseBodyJobInfoDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

type GetJobResponseBodyJobInfoDeploymentPolicyTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyTags) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyTags) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyTags) SetTagKey(v string) *GetJobResponseBodyJobInfoDeploymentPolicyTags {
	s.TagKey = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyTags) SetTagValue(v string) *GetJobResponseBodyJobInfoDeploymentPolicyTags {
	s.TagValue = &v
	return s
}

type GetJobResponseBodyJobInfoTasks struct {
	ExecutorPolicy *GetJobResponseBodyJobInfoTasksExecutorPolicy   `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	ExecutorStatus []*GetJobResponseBodyJobInfoTasksExecutorStatus `json:"ExecutorStatus,omitempty" xml:"ExecutorStatus,omitempty" type:"Repeated"`
	// example:
	//
	// task0
	TaskName *string                                 `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskSpec *GetJobResponseBodyJobInfoTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	// example:
	//
	// true
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
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
	// example:
	//
	// 10
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
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
	// example:
	//
	// 9
	IndexEnd *int32 `json:"IndexEnd,omitempty" xml:"IndexEnd,omitempty"`
	// example:
	//
	// 0
	IndexStart *int32 `json:"IndexStart,omitempty" xml:"IndexStart,omitempty"`
	// example:
	//
	// 1
	IndexStep *int32 `json:"IndexStep,omitempty" xml:"IndexStep,omitempty"`
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
	// example:
	//
	// 0
	ArrayId *int32 `json:"ArrayId,omitempty" xml:"ArrayId,omitempty"`
	// example:
	//
	// 2024-02-04 13:54:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-02-04 13:54:10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2024-02-04 13:54:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Creating executor
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
	// example:
	//
	// 1
	Cores         *float32                                               `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Disks         []*GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	InstanceTypes []*string                                              `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
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

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetInstanceTypes(v []*string) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.InstanceTypes = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetMemory(v int32) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.Memory = &v
	return s
}

type GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks struct {
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// System
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
	// example:
	//
	// m-xxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// ZWNobyAiMTIzNCIgPiBgZGF0ZSArJXNg
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	// example:
	//
	// ZWNobyAiMTIzNCIgPiBgZGF0ZSArJXNg
	Script   *string `json:"Script,omitempty" xml:"Script,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) SetUserName(v string) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM {
	s.UserName = &v
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

type GetPoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
}

func (s GetPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPoolRequest) GoString() string {
	return s.String()
}

func (s *GetPoolRequest) SetPoolName(v string) *GetPoolRequest {
	s.PoolName = &v
	return s
}

type GetPoolResponseBody struct {
	PoolInfo *GetPoolResponseBodyPoolInfo `json:"PoolInfo,omitempty" xml:"PoolInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPoolResponseBody) GoString() string {
	return s.String()
}

func (s *GetPoolResponseBody) SetPoolInfo(v *GetPoolResponseBodyPoolInfo) *GetPoolResponseBody {
	s.PoolInfo = v
	return s
}

func (s *GetPoolResponseBody) SetRequestId(v string) *GetPoolResponseBody {
	s.RequestId = &v
	return s
}

type GetPoolResponseBodyPoolInfo struct {
	// example:
	//
	// 2024-12-01 20:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	ExectorUsage *int32 `json:"ExectorUsage,omitempty" xml:"ExectorUsage,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Fails to **	- pool: ***.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-12-01 20:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPoolResponseBodyPoolInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPoolResponseBodyPoolInfo) GoString() string {
	return s.String()
}

func (s *GetPoolResponseBodyPoolInfo) SetCreateTime(v string) *GetPoolResponseBodyPoolInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetExectorUsage(v int32) *GetPoolResponseBodyPoolInfo {
	s.ExectorUsage = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetIsDefault(v bool) *GetPoolResponseBodyPoolInfo {
	s.IsDefault = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetMaxExectorNum(v int32) *GetPoolResponseBodyPoolInfo {
	s.MaxExectorNum = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetPoolName(v string) *GetPoolResponseBodyPoolInfo {
	s.PoolName = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetPriority(v int32) *GetPoolResponseBodyPoolInfo {
	s.Priority = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetReason(v string) *GetPoolResponseBodyPoolInfo {
	s.Reason = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetStatus(v string) *GetPoolResponseBodyPoolInfo {
	s.Status = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetUpdateTime(v string) *GetPoolResponseBodyPoolInfo {
	s.UpdateTime = &v
	return s
}

type GetPoolResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPoolResponse) GoString() string {
	return s.String()
}

func (s *GetPoolResponse) SetHeaders(v map[string]*string) *GetPoolResponse {
	s.Headers = v
	return s
}

func (s *GetPoolResponse) SetStatusCode(v int32) *GetPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPoolResponse) SetBody(v *GetPoolResponseBody) *GetPoolResponse {
	s.Body = v
	return s
}

type ListExecutorsRequest struct {
	Filter *ListExecutorsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListExecutorsRequest) SetPageNumber(v int32) *ListExecutorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsRequest) SetPageSize(v int32) *ListExecutorsRequest {
	s.PageSize = &v
	return s
}

type ListExecutorsRequestFilter struct {
	ExecutorIds []*string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// testJob
	JobName *string   `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Status  []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	// example:
	//
	// 1703820113
	TimeCreatedBefore *int32 `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
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

func (s *ListExecutorsRequestFilter) SetStatus(v []*string) *ListExecutorsRequestFilter {
	s.Status = v
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
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListExecutorsShrinkRequest) SetPageNumber(v int32) *ListExecutorsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsShrinkRequest) SetPageSize(v int32) *ListExecutorsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListExecutorsResponseBody struct {
	Executors []*ListExecutorsResponseBodyExecutors `json:"Executors,omitempty" xml:"Executors,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 40
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListExecutorsResponseBody) SetPageNumber(v int32) *ListExecutorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsResponseBody) SetPageSize(v int32) *ListExecutorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExecutorsResponseBody) SetRequestId(v string) *ListExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorsResponseBody) SetTotalCount(v string) *ListExecutorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListExecutorsResponseBodyExecutors struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 0
	ArrayIndex *int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// example:
	//
	// 2024-02-20 10:04:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-02-20 10:04:18
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// job-xxxx-task0-1
	ExecutorId        *string   `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	ExpirationTime    *string   `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	ExternalIpAddress []*string `json:"ExternalIpAddress,omitempty" xml:"ExternalIpAddress,omitempty" type:"Repeated"`
	HostName          []*string `json:"HostName,omitempty" xml:"HostName,omitempty" type:"Repeated"`
	Image             *string   `json:"Image,omitempty" xml:"Image,omitempty"`
	IpAddress         []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
	// example:
	//
	// job-hy1nggvyukuvkr******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// testJob
	JobName      *string                                     `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Resource     *ListExecutorsResponseBodyExecutorsResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	ResourceType *string                                     `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime    *string                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Succeeded to release executor resource
	StatusReason *string                                   `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	Tags         []*ListExecutorsResponseBodyExecutorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// task0
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskSustainable *bool   `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListExecutorsResponseBodyExecutors) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutors) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutors) SetAppName(v string) *ListExecutorsResponseBodyExecutors {
	s.AppName = &v
	return s
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

func (s *ListExecutorsResponseBodyExecutors) SetExpirationTime(v string) *ListExecutorsResponseBodyExecutors {
	s.ExpirationTime = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetExternalIpAddress(v []*string) *ListExecutorsResponseBodyExecutors {
	s.ExternalIpAddress = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetHostName(v []*string) *ListExecutorsResponseBodyExecutors {
	s.HostName = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetImage(v string) *ListExecutorsResponseBodyExecutors {
	s.Image = &v
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

func (s *ListExecutorsResponseBodyExecutors) SetResource(v *ListExecutorsResponseBodyExecutorsResource) *ListExecutorsResponseBodyExecutors {
	s.Resource = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetResourceType(v string) *ListExecutorsResponseBodyExecutors {
	s.ResourceType = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetStartTime(v string) *ListExecutorsResponseBodyExecutors {
	s.StartTime = &v
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

func (s *ListExecutorsResponseBodyExecutors) SetTags(v []*ListExecutorsResponseBodyExecutorsTags) *ListExecutorsResponseBodyExecutors {
	s.Tags = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetTaskName(v string) *ListExecutorsResponseBodyExecutors {
	s.TaskName = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetTaskSustainable(v bool) *ListExecutorsResponseBodyExecutors {
	s.TaskSustainable = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetVswitchId(v string) *ListExecutorsResponseBodyExecutors {
	s.VswitchId = &v
	return s
}

type ListExecutorsResponseBodyExecutorsResource struct {
	Cores  *float32                                           `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Disks  []*ListExecutorsResponseBodyExecutorsResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	Memory *float32                                           `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListExecutorsResponseBodyExecutorsResource) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutorsResource) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutorsResource) SetCores(v float32) *ListExecutorsResponseBodyExecutorsResource {
	s.Cores = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResource) SetDisks(v []*ListExecutorsResponseBodyExecutorsResourceDisks) *ListExecutorsResponseBodyExecutorsResource {
	s.Disks = v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResource) SetMemory(v float32) *ListExecutorsResponseBodyExecutorsResource {
	s.Memory = &v
	return s
}

type ListExecutorsResponseBodyExecutorsResourceDisks struct {
	Size *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListExecutorsResponseBodyExecutorsResourceDisks) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutorsResourceDisks) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutorsResourceDisks) SetSize(v int32) *ListExecutorsResponseBodyExecutorsResourceDisks {
	s.Size = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResourceDisks) SetType(v string) *ListExecutorsResponseBodyExecutorsResourceDisks {
	s.Type = &v
	return s
}

type ListExecutorsResponseBodyExecutorsTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListExecutorsResponseBodyExecutorsTags) String() string {
	return tea.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutorsTags) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutorsTags) SetTagKey(v string) *ListExecutorsResponseBodyExecutorsTags {
	s.TagKey = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsTags) SetTagValue(v string) *ListExecutorsResponseBodyExecutorsTags {
	s.TagValue = &v
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
	ImageCategory *string   `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	ImageIds      []*string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty" type:"Repeated"`
	ImageNames    []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageType     *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Mode          *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetImageCategory(v string) *ListImagesRequest {
	s.ImageCategory = &v
	return s
}

func (s *ListImagesRequest) SetImageIds(v []*string) *ListImagesRequest {
	s.ImageIds = v
	return s
}

func (s *ListImagesRequest) SetImageNames(v []*string) *ListImagesRequest {
	s.ImageNames = v
	return s
}

func (s *ListImagesRequest) SetImageType(v string) *ListImagesRequest {
	s.ImageType = &v
	return s
}

func (s *ListImagesRequest) SetMode(v string) *ListImagesRequest {
	s.Mode = &v
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
	ImageCategory    *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	ImageIdsShrink   *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	ImageNamesShrink *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageType        *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Mode             *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImagesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListImagesShrinkRequest) SetImageCategory(v string) *ListImagesShrinkRequest {
	s.ImageCategory = &v
	return s
}

func (s *ListImagesShrinkRequest) SetImageIdsShrink(v string) *ListImagesShrinkRequest {
	s.ImageIdsShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetImageNamesShrink(v string) *ListImagesShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetImageType(v string) *ListImagesShrinkRequest {
	s.ImageType = &v
	return s
}

func (s *ListImagesShrinkRequest) SetMode(v string) *ListImagesShrinkRequest {
	s.Mode = &v
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
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2022-12-09T07:06:34Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentId  *int32  `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// m-bp181x855551ww5yq****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// app-image
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OsTag      *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// v1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Weight  *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetAppId(v string) *ListImagesResponseBodyImages {
	s.AppId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCreateTime(v string) *ListImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDocumentId(v int32) *ListImagesResponseBodyImages {
	s.DocumentId = &v
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

func (s *ListImagesResponseBodyImages) SetOsTag(v string) *ListImagesResponseBodyImages {
	s.OsTag = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetUpdateTime(v string) *ListImagesResponseBodyImages {
	s.UpdateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetVersion(v string) *ListImagesResponseBodyImages {
	s.Version = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetWeight(v int32) *ListImagesResponseBodyImages {
	s.Weight = &v
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
	// example:
	//
	// job-xxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// task0
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

func (s *ListJobExecutorsRequest) SetPageNumber(v int32) *ListJobExecutorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobExecutorsRequest) SetPageSize(v int32) *ListJobExecutorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutorsRequest) SetTaskName(v string) *ListJobExecutorsRequest {
	s.TaskName = &v
	return s
}

type ListJobExecutorsResponseBody struct {
	ExecutorStatus *ListJobExecutorsResponseBodyExecutorStatus `json:"ExecutorStatus,omitempty" xml:"ExecutorStatus,omitempty" type:"Struct"`
	Executors      []*ListJobExecutorsResponseBodyExecutors    `json:"Executors,omitempty" xml:"Executors,omitempty" type:"Repeated"`
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 50
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobExecutorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBody) SetExecutorStatus(v *ListJobExecutorsResponseBodyExecutorStatus) *ListJobExecutorsResponseBody {
	s.ExecutorStatus = v
	return s
}

func (s *ListJobExecutorsResponseBody) SetExecutors(v []*ListJobExecutorsResponseBodyExecutors) *ListJobExecutorsResponseBody {
	s.Executors = v
	return s
}

func (s *ListJobExecutorsResponseBody) SetJobId(v string) *ListJobExecutorsResponseBody {
	s.JobId = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetPageNumber(v int32) *ListJobExecutorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobExecutorsResponseBody) SetPageSize(v int32) *ListJobExecutorsResponseBody {
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

type ListJobExecutorsResponseBodyExecutorStatus struct {
	Deleted   *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Exception *int32 `json:"Exception,omitempty" xml:"Exception,omitempty"`
	Failed    *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	Initing   *int32 `json:"Initing,omitempty" xml:"Initing,omitempty"`
	Pending   *int32 `json:"Pending,omitempty" xml:"Pending,omitempty"`
	Running   *int32 `json:"Running,omitempty" xml:"Running,omitempty"`
	Succeeded *int32 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
}

func (s ListJobExecutorsResponseBodyExecutorStatus) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutorsResponseBodyExecutorStatus) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetDeleted(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Deleted = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetException(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Exception = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetFailed(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Failed = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetIniting(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Initing = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetPending(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Pending = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetRunning(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Running = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorStatus) SetSucceeded(v int32) *ListJobExecutorsResponseBodyExecutorStatus {
	s.Succeeded = &v
	return s
}

type ListJobExecutorsResponseBodyExecutors struct {
	// example:
	//
	// 0
	ArrayIndex *int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// example:
	//
	// 2024-02-20 10:04:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-02-20 10:04:18
	EndTime           *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutorId        *string   `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	ExpirationTime    *string   `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	ExternalIpAddress []*string `json:"ExternalIpAddress,omitempty" xml:"ExternalIpAddress,omitempty" type:"Repeated"`
	HostName          []*string `json:"HostName,omitempty" xml:"HostName,omitempty" type:"Repeated"`
	IpAddress         []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
	StartTime         *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Creating executor
	StatusReason *string                                      `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	Tags         []*ListJobExecutorsResponseBodyExecutorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *ListJobExecutorsResponseBodyExecutors) SetExecutorId(v string) *ListJobExecutorsResponseBodyExecutors {
	s.ExecutorId = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetExpirationTime(v string) *ListJobExecutorsResponseBodyExecutors {
	s.ExpirationTime = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutors) SetExternalIpAddress(v []*string) *ListJobExecutorsResponseBodyExecutors {
	s.ExternalIpAddress = v
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

func (s *ListJobExecutorsResponseBodyExecutors) SetStartTime(v string) *ListJobExecutorsResponseBodyExecutors {
	s.StartTime = &v
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

func (s *ListJobExecutorsResponseBodyExecutors) SetTags(v []*ListJobExecutorsResponseBodyExecutorsTags) *ListJobExecutorsResponseBodyExecutors {
	s.Tags = v
	return s
}

type ListJobExecutorsResponseBodyExecutorsTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListJobExecutorsResponseBodyExecutorsTags) String() string {
	return tea.Prettify(s)
}

func (s ListJobExecutorsResponseBodyExecutorsTags) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponseBodyExecutorsTags) SetTagKey(v string) *ListJobExecutorsResponseBodyExecutorsTags {
	s.TagKey = &v
	return s
}

func (s *ListJobExecutorsResponseBodyExecutorsTags) SetTagValue(v string) *ListJobExecutorsResponseBodyExecutorsTags {
	s.TagValue = &v
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
	Filter *ListJobsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy   *ListJobsRequestSortBy `json:"SortBy,omitempty" xml:"SortBy,omitempty" type:"Struct"`
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

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetSortBy(v *ListJobsRequestSortBy) *ListJobsRequest {
	s.SortBy = v
	return s
}

type ListJobsRequestFilter struct {
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// testJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	// example:
	//
	// 1703820113
	TimeCreatedBefore *int32 `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
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
	// example:
	//
	// time_start
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// ASC
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
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListJobsShrinkRequest) SetPageNumber(v int32) *ListJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageSize(v int32) *ListJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsShrinkRequest) SetSortByShrink(v string) *ListJobsShrinkRequest {
	s.SortByShrink = &v
	return s
}

type ListJobsResponseBody struct {
	JobList []*ListJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AppExtraInfo *string `json:"AppExtraInfo,omitempty" xml:"AppExtraInfo,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 2024-01-25 12:29:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-01-25 12:35:23
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	ExecutorCount *int32 `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// example:
	//
	// job-xxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// testJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 129**********
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// example:
	//
	// 2024-01-25 12:29:23
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListJobsResponseBodyJobListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TaskCount *int32 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	// example:
	//
	// true
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s ListJobsResponseBodyJobList) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobList) SetAppExtraInfo(v string) *ListJobsResponseBodyJobList {
	s.AppExtraInfo = &v
	return s
}

func (s *ListJobsResponseBodyJobList) SetAppName(v string) *ListJobsResponseBodyJobList {
	s.AppName = &v
	return s
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

func (s *ListJobsResponseBodyJobList) SetTags(v []*ListJobsResponseBodyJobListTags) *ListJobsResponseBodyJobList {
	s.Tags = v
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

type ListJobsResponseBodyJobListTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListJobsResponseBodyJobListTags) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobListTags) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobListTags) SetTagKey(v string) *ListJobsResponseBodyJobListTags {
	s.TagKey = &v
	return s
}

func (s *ListJobsResponseBodyJobListTags) SetTagValue(v string) *ListJobsResponseBodyJobListTags {
	s.TagValue = &v
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

type ListPoolsRequest struct {
	Filter *ListPoolsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPoolsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoolsRequest) GoString() string {
	return s.String()
}

func (s *ListPoolsRequest) SetFilter(v *ListPoolsRequestFilter) *ListPoolsRequest {
	s.Filter = v
	return s
}

func (s *ListPoolsRequest) SetPageNumber(v int32) *ListPoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoolsRequest) SetPageSize(v int32) *ListPoolsRequest {
	s.PageSize = &v
	return s
}

type ListPoolsRequestFilter struct {
	PoolName []*string `json:"PoolName,omitempty" xml:"PoolName,omitempty" type:"Repeated"`
	Status   []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	// example:
	//
	// 1703820113
	TimeCreatedBefore *int32 `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
}

func (s ListPoolsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListPoolsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListPoolsRequestFilter) SetPoolName(v []*string) *ListPoolsRequestFilter {
	s.PoolName = v
	return s
}

func (s *ListPoolsRequestFilter) SetStatus(v []*string) *ListPoolsRequestFilter {
	s.Status = v
	return s
}

func (s *ListPoolsRequestFilter) SetTimeCreatedAfter(v int32) *ListPoolsRequestFilter {
	s.TimeCreatedAfter = &v
	return s
}

func (s *ListPoolsRequestFilter) SetTimeCreatedBefore(v int32) *ListPoolsRequestFilter {
	s.TimeCreatedBefore = &v
	return s
}

type ListPoolsShrinkRequest struct {
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPoolsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoolsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPoolsShrinkRequest) SetFilterShrink(v string) *ListPoolsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListPoolsShrinkRequest) SetPageNumber(v int32) *ListPoolsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoolsShrinkRequest) SetPageSize(v int32) *ListPoolsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListPoolsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PoolList []*ListPoolsResponseBodyPoolList `json:"PoolList,omitempty" xml:"PoolList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPoolsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoolsResponseBody) SetPageNumber(v int32) *ListPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPoolsResponseBody) SetPageSize(v int32) *ListPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPoolsResponseBody) SetPoolList(v []*ListPoolsResponseBodyPoolList) *ListPoolsResponseBody {
	s.PoolList = v
	return s
}

func (s *ListPoolsResponseBody) SetRequestId(v string) *ListPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoolsResponseBody) SetTotalCount(v int32) *ListPoolsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPoolsResponseBodyPoolList struct {
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPoolsResponseBodyPoolList) String() string {
	return tea.Prettify(s)
}

func (s ListPoolsResponseBodyPoolList) GoString() string {
	return s.String()
}

func (s *ListPoolsResponseBodyPoolList) SetIsDefault(v bool) *ListPoolsResponseBodyPoolList {
	s.IsDefault = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetMaxExectorNum(v int32) *ListPoolsResponseBodyPoolList {
	s.MaxExectorNum = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetPoolName(v string) *ListPoolsResponseBodyPoolList {
	s.PoolName = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetPriority(v int32) *ListPoolsResponseBodyPoolList {
	s.Priority = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetStatus(v string) *ListPoolsResponseBodyPoolList {
	s.Status = &v
	return s
}

type ListPoolsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoolsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListPoolsResponse) SetHeaders(v map[string]*string) *ListPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListPoolsResponse) SetStatusCode(v int32) *ListPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoolsResponse) SetBody(v *ListPoolsResponseBody) *ListPoolsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// example:
	//
	// 20
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// example:
	//
	// 1d2db86scXXXXXXXXXX
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Job
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResult(v int32) *ListTagResourcesRequest {
	s.MaxResult = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// example:
	//
	// 1d2db86scXXXXXXXXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// example:
	//
	// job-xxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Job
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type RemoveImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// m-bp14wakr1rkxtb******
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
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

func (s *RemoveImageRequest) SetImageType(v string) *RemoveImageRequest {
	s.ImageType = &v
	return s
}

type RemoveImageResponseBody struct {
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

type TagResourcesRequest struct {
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Job
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnTagResourcesRequest struct {
	// example:
	//
	// False
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Job
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesResponseBody struct {
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type UpdatePoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority       *int32                           `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceLimits *UpdatePoolRequestResourceLimits `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty" type:"Struct"`
}

func (s UpdatePoolRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePoolRequest) GoString() string {
	return s.String()
}

func (s *UpdatePoolRequest) SetPoolName(v string) *UpdatePoolRequest {
	s.PoolName = &v
	return s
}

func (s *UpdatePoolRequest) SetPriority(v int32) *UpdatePoolRequest {
	s.Priority = &v
	return s
}

func (s *UpdatePoolRequest) SetResourceLimits(v *UpdatePoolRequestResourceLimits) *UpdatePoolRequest {
	s.ResourceLimits = v
	return s
}

type UpdatePoolRequestResourceLimits struct {
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
}

func (s UpdatePoolRequestResourceLimits) String() string {
	return tea.Prettify(s)
}

func (s UpdatePoolRequestResourceLimits) GoString() string {
	return s.String()
}

func (s *UpdatePoolRequestResourceLimits) SetMaxExectorNum(v int32) *UpdatePoolRequestResourceLimits {
	s.MaxExectorNum = &v
	return s
}

type UpdatePoolShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// example:
	//
	// 1
	Priority             *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceLimitsShrink *string `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
}

func (s UpdatePoolShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePoolShrinkRequest) SetPoolName(v string) *UpdatePoolShrinkRequest {
	s.PoolName = &v
	return s
}

func (s *UpdatePoolShrinkRequest) SetPriority(v int32) *UpdatePoolShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdatePoolShrinkRequest) SetResourceLimitsShrink(v string) *UpdatePoolShrinkRequest {
	s.ResourceLimitsShrink = &v
	return s
}

type UpdatePoolResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePoolResponseBody) SetRequestId(v string) *UpdatePoolResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePoolResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePoolResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePoolResponse) GoString() string {
	return s.String()
}

func (s *UpdatePoolResponse) SetHeaders(v map[string]*string) *UpdatePoolResponse {
	s.Headers = v
	return s
}

func (s *UpdatePoolResponse) SetStatusCode(v int32) *UpdatePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePoolResponse) SetBody(v *UpdatePoolResponseBody) *UpdatePoolResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ehpcinstant"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加托管侧用户自定义镜像
//
// @param tmpReq - AddImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageResponse
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

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageVersion)) {
		query["ImageVersion"] = request.ImageVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.VMImageSpecShrink)) {
		query["VMImageSpec"] = request.VMImageSpecShrink
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

// Summary:
//
// 添加托管侧用户自定义镜像
//
// @param request - AddImageRequest
//
// @return AddImageResponse
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

// Summary:
//
// 提交任务
//
// @param tmpReq - CreateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
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

	if !tea.BoolValue(util.IsUnset(tmpReq.SecurityPolicy)) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, tea.String("SecurityPolicy"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.JobScheduler)) {
		query["JobScheduler"] = request.JobScheduler
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyShrink)) {
		query["SecurityPolicy"] = request.SecurityPolicyShrink
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

// Summary:
//
// 提交任务
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
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

// Summary:
//
// 创建资源池
//
// @param tmpReq - CreatePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePoolResponse
func (client *Client) CreatePoolWithOptions(tmpReq *CreatePoolRequest, runtime *util.RuntimeOptions) (_result *CreatePoolResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceLimits)) {
		request.ResourceLimitsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceLimits, tea.String("ResourceLimits"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceLimitsShrink)) {
		query["ResourceLimits"] = request.ResourceLimitsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePool"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源池
//
// @param request - CreatePoolRequest
//
// @return CreatePoolResponse
func (client *Client) CreatePool(request *CreatePoolRequest) (_result *CreatePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePoolResponse{}
	_body, _err := client.CreatePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除作业
//
// @param tmpReq - DeleteJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobsResponse
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

	if !tea.BoolValue(util.IsUnset(request.JobScheduler)) {
		query["JobScheduler"] = request.JobScheduler
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

// Summary:
//
// 删除作业
//
// @param request - DeleteJobsRequest
//
// @return DeleteJobsResponse
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

// Summary:
//
// 删除资源池
//
// @param request - DeletePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePoolResponse
func (client *Client) DeletePoolWithOptions(request *DeletePoolRequest, runtime *util.RuntimeOptions) (_result *DeletePoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePool"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源池
//
// @param request - DeletePoolRequest
//
// @return DeletePoolResponse
func (client *Client) DeletePool(request *DeletePoolRequest) (_result *DeletePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePoolResponse{}
	_body, _err := client.DeletePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询作业性能数据
//
// @param tmpReq - DescribeJobMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobMetricDataResponse
func (client *Client) DescribeJobMetricDataWithOptions(tmpReq *DescribeJobMetricDataRequest, runtime *util.RuntimeOptions) (_result *DescribeJobMetricDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeJobMetricDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ArrayIndex)) {
		request.ArrayIndexShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArrayIndex, tea.String("ArrayIndex"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArrayIndexShrink)) {
		query["ArrayIndex"] = request.ArrayIndexShrink
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeJobMetricData"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeJobMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作业性能数据
//
// @param request - DescribeJobMetricDataRequest
//
// @return DescribeJobMetricDataResponse
func (client *Client) DescribeJobMetricData(request *DescribeJobMetricDataRequest) (_result *DescribeJobMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobMetricDataResponse{}
	_body, _err := client.DescribeJobMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询作业即时监控项
//
// @param tmpReq - DescribeJobMetricLastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobMetricLastResponse
func (client *Client) DescribeJobMetricLastWithOptions(tmpReq *DescribeJobMetricLastRequest, runtime *util.RuntimeOptions) (_result *DescribeJobMetricLastResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeJobMetricLastShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ArrayIndex)) {
		request.ArrayIndexShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArrayIndex, tea.String("ArrayIndex"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArrayIndexShrink)) {
		query["ArrayIndex"] = request.ArrayIndexShrink
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeJobMetricLast"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeJobMetricLastResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作业即时监控项
//
// @param request - DescribeJobMetricLastRequest
//
// @return DescribeJobMetricLastResponse
func (client *Client) DescribeJobMetricLast(request *DescribeJobMetricLastRequest) (_result *DescribeJobMetricLastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobMetricLastResponse{}
	_body, _err := client.DescribeJobMetricLastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看应用版本列表
//
// @param request - GetAppVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppVersionsResponse
func (client *Client) GetAppVersionsWithOptions(request *GetAppVersionsRequest, runtime *util.RuntimeOptions) (_result *GetAppVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCategory)) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
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
		Action:      tea.String("GetAppVersions"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看应用版本列表
//
// @param request - GetAppVersionsRequest
//
// @return GetAppVersionsResponse
func (client *Client) GetAppVersions(request *GetAppVersionsRequest) (_result *GetAppVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppVersionsResponse{}
	_body, _err := client.GetAppVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询托管侧镜像详情。
//
// @param tmpReq - GetImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
func (client *Client) GetImageWithOptions(tmpReq *GetImageRequest, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AdditionalRegionIds)) {
		request.AdditionalRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdditionalRegionIds, tea.String("AdditionalRegionIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalRegionIdsShrink)) {
		query["AdditionalRegionIds"] = request.AdditionalRegionIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCategory)) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
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

// Summary:
//
// 查询托管侧镜像详情。
//
// @param request - GetImageRequest
//
// @return GetImageResponse
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

// Summary:
//
// 查询作业详情
//
// @param request - GetJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
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

// Summary:
//
// 查询作业详情
//
// @param request - GetJobRequest
//
// @return GetJobResponse
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

// Summary:
//
// 查询队列详细信息
//
// @param request - GetPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPoolResponse
func (client *Client) GetPoolWithOptions(request *GetPoolRequest, runtime *util.RuntimeOptions) (_result *GetPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPool"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询队列详细信息
//
// @param request - GetPoolRequest
//
// @return GetPoolResponse
func (client *Client) GetPool(request *GetPoolRequest) (_result *GetPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPoolResponse{}
	_body, _err := client.GetPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询全局Executor信息
//
// @param tmpReq - ListExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutorsResponse
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

// Summary:
//
// 查询全局Executor信息
//
// @param request - ListExecutorsRequest
//
// @return ListExecutorsResponse
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

// Summary:
//
// 查看托管侧镜像列表
//
// @param tmpReq - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
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
	if !tea.BoolValue(util.IsUnset(request.ImageCategory)) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !tea.BoolValue(util.IsUnset(request.ImageIdsShrink)) {
		query["ImageIds"] = request.ImageIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageNamesShrink)) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
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

// Summary:
//
// 查看托管侧镜像列表
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
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

// Summary:
//
// 查询作业Executor信息
//
// @param request - ListJobExecutorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobExecutorsResponse
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

// Summary:
//
// 查询作业Executor信息
//
// @param request - ListJobExecutorsRequest
//
// @return ListJobExecutorsResponse
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

// Summary:
//
// 查询作业列表
//
// @param tmpReq - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
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

// Summary:
//
// 查询作业列表
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
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

// Summary:
//
// 查询资源池列表
//
// @param tmpReq - ListPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoolsResponse
func (client *Client) ListPoolsWithOptions(tmpReq *ListPoolsRequest, runtime *util.RuntimeOptions) (_result *ListPoolsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPoolsShrinkRequest{}
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
		Action:      tea.String("ListPools"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源池列表
//
// @param request - ListPoolsRequest
//
// @return ListPoolsResponse
func (client *Client) ListPools(request *ListPoolsRequest) (_result *ListPoolsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoolsResponse{}
	_body, _err := client.ListPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一个或多个资源已经绑定的标签列表
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一个或多个资源已经绑定的标签列表
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除托管侧镜像信息。
//
// @param request - RemoveImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageResponse
func (client *Client) RemoveImageWithOptions(request *RemoveImageRequest, runtime *util.RuntimeOptions) (_result *RemoveImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
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

// Summary:
//
// 移除托管侧镜像信息。
//
// @param request - RemoveImageRequest
//
// @return RemoveImageResponse
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

// Summary:
//
// 为指定的资源列表统一创建并绑定标签
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为指定的资源列表统一创建并绑定标签
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为指定的ECS资源列表统一解绑标签
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为指定的ECS资源列表统一解绑标签
//
// @param request - UnTagResourcesRequest
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源池
//
// @param tmpReq - UpdatePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePoolResponse
func (client *Client) UpdatePoolWithOptions(tmpReq *UpdatePoolRequest, runtime *util.RuntimeOptions) (_result *UpdatePoolResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdatePoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceLimits)) {
		request.ResourceLimitsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceLimits, tea.String("ResourceLimits"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceLimitsShrink)) {
		query["ResourceLimits"] = request.ResourceLimitsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePool"),
		Version:     tea.String("2023-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源池
//
// @param request - UpdatePoolRequest
//
// @return UpdatePoolResponse
func (client *Client) UpdatePool(request *UpdatePoolRequest) (_result *UpdatePoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePoolResponse{}
	_body, _err := client.UpdatePoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
