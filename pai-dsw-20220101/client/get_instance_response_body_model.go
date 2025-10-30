// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *GetInstanceResponseBody
	GetAcceleratorType() *string
	SetAccessibility(v string) *GetInstanceResponseBody
	GetAccessibility() *string
	SetAccumulatedRunningTimeInMs(v int64) *GetInstanceResponseBody
	GetAccumulatedRunningTimeInMs() *int64
	SetAffinity(v *GetInstanceResponseBodyAffinity) *GetInstanceResponseBody
	GetAffinity() *GetInstanceResponseBodyAffinity
	SetCloudDisks(v []*GetInstanceResponseBodyCloudDisks) *GetInstanceResponseBody
	GetCloudDisks() []*GetInstanceResponseBodyCloudDisks
	SetCode(v string) *GetInstanceResponseBody
	GetCode() *string
	SetCredentialConfig(v *CredentialConfig) *GetInstanceResponseBody
	GetCredentialConfig() *CredentialConfig
	SetDatasets(v []*GetInstanceResponseBodyDatasets) *GetInstanceResponseBody
	GetDatasets() []*GetInstanceResponseBodyDatasets
	SetDriver(v string) *GetInstanceResponseBody
	GetDriver() *string
	SetDynamicMount(v *DynamicMount) *GetInstanceResponseBody
	GetDynamicMount() *DynamicMount
	SetEcsSpec(v string) *GetInstanceResponseBody
	GetEcsSpec() *string
	SetEnvironmentVariables(v map[string]*string) *GetInstanceResponseBody
	GetEnvironmentVariables() map[string]*string
	SetGmtCreateTime(v string) *GetInstanceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetInstanceResponseBody
	GetGmtModifiedTime() *string
	SetHttpStatusCode(v int32) *GetInstanceResponseBody
	GetHttpStatusCode() *int32
	SetIdleInstanceCuller(v *GetInstanceResponseBodyIdleInstanceCuller) *GetInstanceResponseBody
	GetIdleInstanceCuller() *GetInstanceResponseBodyIdleInstanceCuller
	SetImageAuth(v string) *GetInstanceResponseBody
	GetImageAuth() *string
	SetImageId(v string) *GetInstanceResponseBody
	GetImageId() *string
	SetImageName(v string) *GetInstanceResponseBody
	GetImageName() *string
	SetImageUrl(v string) *GetInstanceResponseBody
	GetImageUrl() *string
	SetInstanceId(v string) *GetInstanceResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *GetInstanceResponseBody
	GetInstanceName() *string
	SetInstanceShutdownTimer(v *GetInstanceResponseBodyInstanceShutdownTimer) *GetInstanceResponseBody
	GetInstanceShutdownTimer() *GetInstanceResponseBodyInstanceShutdownTimer
	SetInstanceSnapshotList(v []*GetInstanceResponseBodyInstanceSnapshotList) *GetInstanceResponseBody
	GetInstanceSnapshotList() []*GetInstanceResponseBodyInstanceSnapshotList
	SetInstanceUrl(v string) *GetInstanceResponseBody
	GetInstanceUrl() *string
	SetJupyterlabUrl(v string) *GetInstanceResponseBody
	GetJupyterlabUrl() *string
	SetLabels(v []*GetInstanceResponseBodyLabels) *GetInstanceResponseBody
	GetLabels() []*GetInstanceResponseBodyLabels
	SetLatestSnapshot(v *GetInstanceResponseBodyLatestSnapshot) *GetInstanceResponseBody
	GetLatestSnapshot() *GetInstanceResponseBodyLatestSnapshot
	SetMessage(v string) *GetInstanceResponseBody
	GetMessage() *string
	SetNodeErrorRecovery(v *GetInstanceResponseBodyNodeErrorRecovery) *GetInstanceResponseBody
	GetNodeErrorRecovery() *GetInstanceResponseBodyNodeErrorRecovery
	SetPaymentType(v string) *GetInstanceResponseBody
	GetPaymentType() *string
	SetPriority(v int64) *GetInstanceResponseBody
	GetPriority() *int64
	SetProxyPath(v string) *GetInstanceResponseBody
	GetProxyPath() *string
	SetReasonCode(v string) *GetInstanceResponseBody
	GetReasonCode() *string
	SetReasonMessage(v string) *GetInstanceResponseBody
	GetReasonMessage() *string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetRequestedResource(v *GetInstanceResponseBodyRequestedResource) *GetInstanceResponseBody
	GetRequestedResource() *GetInstanceResponseBodyRequestedResource
	SetResourceId(v string) *GetInstanceResponseBody
	GetResourceId() *string
	SetResourceName(v string) *GetInstanceResponseBody
	GetResourceName() *string
	SetStatus(v string) *GetInstanceResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetInstanceResponseBody
	GetSuccess() *bool
	SetTags(v []*GetInstanceResponseBodyTags) *GetInstanceResponseBody
	GetTags() []*GetInstanceResponseBodyTags
	SetTerminalUrl(v string) *GetInstanceResponseBody
	GetTerminalUrl() *string
	SetUserCommandId(v string) *GetInstanceResponseBody
	GetUserCommandId() *string
	SetUserId(v string) *GetInstanceResponseBody
	GetUserId() *string
	SetUserName(v string) *GetInstanceResponseBody
	GetUserName() *string
	SetUserVpc(v *GetInstanceResponseBodyUserVpc) *GetInstanceResponseBody
	GetUserVpc() *GetInstanceResponseBodyUserVpc
	SetWebIDEUrl(v string) *GetInstanceResponseBody
	GetWebIDEUrl() *string
	SetWorkspaceId(v string) *GetInstanceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *GetInstanceResponseBody
	GetWorkspaceName() *string
	SetWorkspaceSource(v string) *GetInstanceResponseBody
	GetWorkspaceSource() *string
}

type GetInstanceResponseBody struct {
	// The accelerator type of the instance.
	//
	// Valid values:
	//
	// 	- CPU
	//
	// 	- GPU
	//
	// example:
	//
	// CPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The accessibility. Valid values:
	//
	// 	- PRIVATE: Accessible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: Accessible to all members in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The accumulated running duration. Unit: milliseconds.
	//
	// example:
	//
	// 3600000
	AccumulatedRunningTimeInMs *int64 `json:"AccumulatedRunningTimeInMs,omitempty" xml:"AccumulatedRunningTimeInMs,omitempty"`
	// The affinity configuration.
	Affinity *GetInstanceResponseBodyAffinity `json:"Affinity,omitempty" xml:"Affinity,omitempty" type:"Struct"`
	// The cloud disks of the instance.
	//
	// example:
	//
	// []
	CloudDisks []*GetInstanceResponseBodyCloudDisks `json:"CloudDisks,omitempty" xml:"CloudDisks,omitempty" type:"Repeated"`
	// The status code. Valid values:
	//
	// 	- InternalError: All errors, except for parameter validation errors, are internal errors.
	//
	// 	- ValidationError: A parameter validation error.
	//
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The credential injection configuration.
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The datasets.
	Datasets []*GetInstanceResponseBodyDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The NVIDIA driver configuration.
	//
	// example:
	//
	// 535.54.03
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// The dynamic mount configuration.
	DynamicMount *DynamicMount `json:"DynamicMount,omitempty" xml:"DynamicMount,omitempty"`
	// The ECS instance type of the instance.
	//
	// example:
	//
	// ecs.c6.large
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// The environment variables.
	//
	// example:
	//
	// {userName: "Chris"}
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The creation time of the instance.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modified time of the instance.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The HTTP status code. Valid values:
	//
	// 	- 400
	//
	// 	- 404
	//
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The automatic shutdown settings.
	//
	// example:
	//
	// {"InstanceId":"dsw-05cefd0be2e5a278","CpuPercentThreshold":20,"GpuPercentThreshold":10,"MaxIdleTimeInMinutes":120,"IdleTimeInMinutes":30}
	IdleInstanceCuller *GetInstanceResponseBodyIdleInstanceCuller `json:"IdleInstanceCuller,omitempty" xml:"IdleInstanceCuller,omitempty" type:"Struct"`
	// The Base64-encoded account and password for the user‘s private image. The password will be hidden.
	//
	// example:
	//
	// YWxpeXVuNjUzMzM5MjIwMzoqKioqKio=
	ImageAuth *string `json:"ImageAuth,omitempty" xml:"ImageAuth,omitempty"`
	// The image ID.
	//
	// example:
	//
	// image-05cefd0be2exxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// py36_cpu_tf1.12_ubuntu
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image address.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/pai_product/tensorflow:py36_cpu_tf1.12_ubuntu
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// training_data
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The scheduled stop tasks.
	InstanceShutdownTimer *GetInstanceResponseBodyInstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	// The instance snapshots.
	//
	// example:
	//
	// []
	InstanceSnapshotList []*GetInstanceResponseBodyInstanceSnapshotList `json:"InstanceSnapshotList,omitempty" xml:"InstanceSnapshotList,omitempty" type:"Repeated"`
	// The instance URL.
	//
	// example:
	//
	// https://dsw-cn-shanghai.data.aliyun.com/notebook.htm?instance=39772#/
	InstanceUrl *string `json:"InstanceUrl,omitempty" xml:"InstanceUrl,omitempty"`
	// The JupyterLab URL.
	//
	// example:
	//
	// https://dsw-gateway-cn-shanghai.aliyun.com/dsw-39772/lab/
	JupyterlabUrl *string `json:"JupyterlabUrl,omitempty" xml:"JupyterlabUrl,omitempty"`
	// The custom tags.
	//
	// example:
	//
	// {\\"foo\\": \\"bar\\"}
	Labels []*GetInstanceResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The latest user image saved.
	LatestSnapshot *GetInstanceResponseBodyLatestSnapshot `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	// The error message. Valid values:
	//
	// 	- If the request is successful, null is returned.
	//
	// 	- If the request fails, the cause for the failure is returned.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The error recovery configuration of the node.
	NodeErrorRecovery *GetInstanceResponseBodyNodeErrorRecovery `json:"NodeErrorRecovery,omitempty" xml:"NodeErrorRecovery,omitempty" type:"Struct"`
	// The billing method. Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Subscription
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The priority based on which resources are allocated to instances.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The proxy path.
	//
	// example:
	//
	// dsw-170197/proxy/
	ProxyPath *string `json:"ProxyPath,omitempty" xml:"ProxyPath,omitempty"`
	// The error code of the instance.
	//
	// example:
	//
	// Internal Error
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The cause of the instance error.
	//
	// example:
	//
	// ImagePullBackOff
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource configurations in subscription scenarios.
	//
	// example:
	//
	// {"CPU":"4","Memory":"8Gi","SharedMemory":"4Gi","GPU":"1","GPUType":"Tesla-V100-16G"}
	RequestedResource *GetInstanceResponseBodyRequestedResource `json:"RequestedResource,omitempty" xml:"RequestedResource,omitempty" type:"Struct"`
	// The resource ID. This parameter is available if the billing method is subscription.
	//
	// example:
	//
	// dsw-123456789
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The specification type.
	//
	// 	- For subscription, this is the requested CPU and memory size.
	//
	// 	- For pay-as-you-go, this is the selected ECS instance type.
	//
	// example:
	//
	// ecs.g7.xlarge
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The instance status.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- SaveFailed
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- ResourceAllocating
	//
	// 	- Stopping
	//
	// 	- Updating
	//
	// 	- Saving
	//
	// 	- Queuing
	//
	// 	- Recovering
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Saved
	//
	// 	- Deleting
	//
	// 	- EnvPreparing
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The tags.
	Tags []*GetInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The terminal URL.
	//
	// example:
	//
	// https://dsw-gateway-cn-shanghai.aliyun.com/dsw-39772/tty/
	TerminalUrl   *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
	UserCommandId *string `json:"UserCommandId,omitempty" xml:"UserCommandId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1612285282502324
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// 测试用户
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The virtual private cloud (VPC) configurations.
	UserVpc *GetInstanceResponseBodyUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The Web IDE URL.
	//
	// example:
	//
	// https://dsw-gateway-cn-shanghai.aliyun.com/dsw-39772/ide/
	WebIDEUrl *string `json:"WebIDEUrl,omitempty" xml:"WebIDEUrl,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 40823
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// training_data
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// The storage for the workspace. If you leave this parameter empty, the workspace uses File Storage NAS (NAS) storage, cloud disks, or local disks in sequence.
	//
	// example:
	//
	// d-123456789
	WorkspaceSource *string `json:"WorkspaceSource,omitempty" xml:"WorkspaceSource,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *GetInstanceResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetInstanceResponseBody) GetAccumulatedRunningTimeInMs() *int64 {
	return s.AccumulatedRunningTimeInMs
}

func (s *GetInstanceResponseBody) GetAffinity() *GetInstanceResponseBodyAffinity {
	return s.Affinity
}

func (s *GetInstanceResponseBody) GetCloudDisks() []*GetInstanceResponseBodyCloudDisks {
	return s.CloudDisks
}

func (s *GetInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceResponseBody) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *GetInstanceResponseBody) GetDatasets() []*GetInstanceResponseBodyDatasets {
	return s.Datasets
}

func (s *GetInstanceResponseBody) GetDriver() *string {
	return s.Driver
}

func (s *GetInstanceResponseBody) GetDynamicMount() *DynamicMount {
	return s.DynamicMount
}

func (s *GetInstanceResponseBody) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *GetInstanceResponseBody) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *GetInstanceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceResponseBody) GetIdleInstanceCuller() *GetInstanceResponseBodyIdleInstanceCuller {
	return s.IdleInstanceCuller
}

func (s *GetInstanceResponseBody) GetImageAuth() *string {
	return s.ImageAuth
}

func (s *GetInstanceResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *GetInstanceResponseBody) GetImageName() *string {
	return s.ImageName
}

func (s *GetInstanceResponseBody) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceResponseBody) GetInstanceShutdownTimer() *GetInstanceResponseBodyInstanceShutdownTimer {
	return s.InstanceShutdownTimer
}

func (s *GetInstanceResponseBody) GetInstanceSnapshotList() []*GetInstanceResponseBodyInstanceSnapshotList {
	return s.InstanceSnapshotList
}

func (s *GetInstanceResponseBody) GetInstanceUrl() *string {
	return s.InstanceUrl
}

func (s *GetInstanceResponseBody) GetJupyterlabUrl() *string {
	return s.JupyterlabUrl
}

func (s *GetInstanceResponseBody) GetLabels() []*GetInstanceResponseBodyLabels {
	return s.Labels
}

func (s *GetInstanceResponseBody) GetLatestSnapshot() *GetInstanceResponseBodyLatestSnapshot {
	return s.LatestSnapshot
}

func (s *GetInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceResponseBody) GetNodeErrorRecovery() *GetInstanceResponseBodyNodeErrorRecovery {
	return s.NodeErrorRecovery
}

func (s *GetInstanceResponseBody) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetInstanceResponseBody) GetPriority() *int64 {
	return s.Priority
}

func (s *GetInstanceResponseBody) GetProxyPath() *string {
	return s.ProxyPath
}

func (s *GetInstanceResponseBody) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetInstanceResponseBody) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetRequestedResource() *GetInstanceResponseBodyRequestedResource {
	return s.RequestedResource
}

func (s *GetInstanceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetInstanceResponseBody) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceResponseBody) GetTags() []*GetInstanceResponseBodyTags {
	return s.Tags
}

func (s *GetInstanceResponseBody) GetTerminalUrl() *string {
	return s.TerminalUrl
}

func (s *GetInstanceResponseBody) GetUserCommandId() *string {
	return s.UserCommandId
}

func (s *GetInstanceResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetInstanceResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *GetInstanceResponseBody) GetUserVpc() *GetInstanceResponseBodyUserVpc {
	return s.UserVpc
}

func (s *GetInstanceResponseBody) GetWebIDEUrl() *string {
	return s.WebIDEUrl
}

func (s *GetInstanceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetInstanceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetInstanceResponseBody) GetWorkspaceSource() *string {
	return s.WorkspaceSource
}

func (s *GetInstanceResponseBody) SetAcceleratorType(v string) *GetInstanceResponseBody {
	s.AcceleratorType = &v
	return s
}

func (s *GetInstanceResponseBody) SetAccessibility(v string) *GetInstanceResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetInstanceResponseBody) SetAccumulatedRunningTimeInMs(v int64) *GetInstanceResponseBody {
	s.AccumulatedRunningTimeInMs = &v
	return s
}

func (s *GetInstanceResponseBody) SetAffinity(v *GetInstanceResponseBodyAffinity) *GetInstanceResponseBody {
	s.Affinity = v
	return s
}

func (s *GetInstanceResponseBody) SetCloudDisks(v []*GetInstanceResponseBodyCloudDisks) *GetInstanceResponseBody {
	s.CloudDisks = v
	return s
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetCredentialConfig(v *CredentialConfig) *GetInstanceResponseBody {
	s.CredentialConfig = v
	return s
}

func (s *GetInstanceResponseBody) SetDatasets(v []*GetInstanceResponseBodyDatasets) *GetInstanceResponseBody {
	s.Datasets = v
	return s
}

func (s *GetInstanceResponseBody) SetDriver(v string) *GetInstanceResponseBody {
	s.Driver = &v
	return s
}

func (s *GetInstanceResponseBody) SetDynamicMount(v *DynamicMount) *GetInstanceResponseBody {
	s.DynamicMount = v
	return s
}

func (s *GetInstanceResponseBody) SetEcsSpec(v string) *GetInstanceResponseBody {
	s.EcsSpec = &v
	return s
}

func (s *GetInstanceResponseBody) SetEnvironmentVariables(v map[string]*string) *GetInstanceResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *GetInstanceResponseBody) SetGmtCreateTime(v string) *GetInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetGmtModifiedTime(v string) *GetInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetIdleInstanceCuller(v *GetInstanceResponseBodyIdleInstanceCuller) *GetInstanceResponseBody {
	s.IdleInstanceCuller = v
	return s
}

func (s *GetInstanceResponseBody) SetImageAuth(v string) *GetInstanceResponseBody {
	s.ImageAuth = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageId(v string) *GetInstanceResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageName(v string) *GetInstanceResponseBody {
	s.ImageName = &v
	return s
}

func (s *GetInstanceResponseBody) SetImageUrl(v string) *GetInstanceResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceId(v string) *GetInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceName(v string) *GetInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceShutdownTimer(v *GetInstanceResponseBodyInstanceShutdownTimer) *GetInstanceResponseBody {
	s.InstanceShutdownTimer = v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceSnapshotList(v []*GetInstanceResponseBodyInstanceSnapshotList) *GetInstanceResponseBody {
	s.InstanceSnapshotList = v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceUrl(v string) *GetInstanceResponseBody {
	s.InstanceUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetJupyterlabUrl(v string) *GetInstanceResponseBody {
	s.JupyterlabUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetLabels(v []*GetInstanceResponseBodyLabels) *GetInstanceResponseBody {
	s.Labels = v
	return s
}

func (s *GetInstanceResponseBody) SetLatestSnapshot(v *GetInstanceResponseBodyLatestSnapshot) *GetInstanceResponseBody {
	s.LatestSnapshot = v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetNodeErrorRecovery(v *GetInstanceResponseBodyNodeErrorRecovery) *GetInstanceResponseBody {
	s.NodeErrorRecovery = v
	return s
}

func (s *GetInstanceResponseBody) SetPaymentType(v string) *GetInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceResponseBody) SetPriority(v int64) *GetInstanceResponseBody {
	s.Priority = &v
	return s
}

func (s *GetInstanceResponseBody) SetProxyPath(v string) *GetInstanceResponseBody {
	s.ProxyPath = &v
	return s
}

func (s *GetInstanceResponseBody) SetReasonCode(v string) *GetInstanceResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetReasonMessage(v string) *GetInstanceResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestedResource(v *GetInstanceResponseBodyRequestedResource) *GetInstanceResponseBody {
	s.RequestedResource = v
	return s
}

func (s *GetInstanceResponseBody) SetResourceId(v string) *GetInstanceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResourceName(v string) *GetInstanceResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetStatus(v string) *GetInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceResponseBody) SetTags(v []*GetInstanceResponseBodyTags) *GetInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBody) SetTerminalUrl(v string) *GetInstanceResponseBody {
	s.TerminalUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserCommandId(v string) *GetInstanceResponseBody {
	s.UserCommandId = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserId(v string) *GetInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserName(v string) *GetInstanceResponseBody {
	s.UserName = &v
	return s
}

func (s *GetInstanceResponseBody) SetUserVpc(v *GetInstanceResponseBodyUserVpc) *GetInstanceResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetInstanceResponseBody) SetWebIDEUrl(v string) *GetInstanceResponseBody {
	s.WebIDEUrl = &v
	return s
}

func (s *GetInstanceResponseBody) SetWorkspaceId(v string) *GetInstanceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetWorkspaceName(v string) *GetInstanceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetWorkspaceSource(v string) *GetInstanceResponseBody {
	s.WorkspaceSource = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.Affinity != nil {
		if err := s.Affinity.Validate(); err != nil {
			return err
		}
	}
	if s.CloudDisks != nil {
		for _, item := range s.CloudDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Datasets != nil {
		for _, item := range s.Datasets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DynamicMount != nil {
		if err := s.DynamicMount.Validate(); err != nil {
			return err
		}
	}
	if s.IdleInstanceCuller != nil {
		if err := s.IdleInstanceCuller.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceShutdownTimer != nil {
		if err := s.InstanceShutdownTimer.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceSnapshotList != nil {
		for _, item := range s.InstanceSnapshotList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.LatestSnapshot != nil {
		if err := s.LatestSnapshot.Validate(); err != nil {
			return err
		}
	}
	if s.NodeErrorRecovery != nil {
		if err := s.NodeErrorRecovery.Validate(); err != nil {
			return err
		}
	}
	if s.RequestedResource != nil {
		if err := s.RequestedResource.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyAffinity struct {
	// The CPU affinity configuration. Only subscription instances that use general-purpose computing resources support CPU affinity configuration.
	CPU *GetInstanceResponseBodyAffinityCPU `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBodyAffinity) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyAffinity) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyAffinity) GetCPU() *GetInstanceResponseBodyAffinityCPU {
	return s.CPU
}

func (s *GetInstanceResponseBodyAffinity) SetCPU(v *GetInstanceResponseBodyAffinityCPU) *GetInstanceResponseBodyAffinity {
	s.CPU = v
	return s
}

func (s *GetInstanceResponseBodyAffinity) Validate() error {
	if s.CPU != nil {
		if err := s.CPU.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyAffinityCPU struct {
	// Indicates whether CPU affinity is enabled.
	//
	// true false
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s GetInstanceResponseBodyAffinityCPU) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyAffinityCPU) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyAffinityCPU) GetEnable() *bool {
	return s.Enable
}

func (s *GetInstanceResponseBodyAffinityCPU) SetEnable(v bool) *GetInstanceResponseBodyAffinityCPU {
	s.Enable = &v
	return s
}

func (s *GetInstanceResponseBodyAffinityCPU) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyCloudDisks struct {
	// Disk Capacity
	//
	// example:
	//
	// 30Gi
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The mount path of the cloud disk in the container.
	//
	// example:
	//
	// /mmt/workspace
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The directory on the cloud disk that is mounted to the container.
	//
	// example:
	//
	// /workspace
	Path   *string                                  `json:"Path,omitempty" xml:"Path,omitempty"`
	Status *GetInstanceResponseBodyCloudDisksStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	// The usage mode of the cloud disk. The value rootfs indicates that the cloud disk is used as the root file system.
	//
	// example:
	//
	// rootfs
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s GetInstanceResponseBodyCloudDisks) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyCloudDisks) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyCloudDisks) GetCapacity() *string {
	return s.Capacity
}

func (s *GetInstanceResponseBodyCloudDisks) GetMountPath() *string {
	return s.MountPath
}

func (s *GetInstanceResponseBodyCloudDisks) GetPath() *string {
	return s.Path
}

func (s *GetInstanceResponseBodyCloudDisks) GetStatus() *GetInstanceResponseBodyCloudDisksStatus {
	return s.Status
}

func (s *GetInstanceResponseBodyCloudDisks) GetSubType() *string {
	return s.SubType
}

func (s *GetInstanceResponseBodyCloudDisks) SetCapacity(v string) *GetInstanceResponseBodyCloudDisks {
	s.Capacity = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisks) SetMountPath(v string) *GetInstanceResponseBodyCloudDisks {
	s.MountPath = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisks) SetPath(v string) *GetInstanceResponseBodyCloudDisks {
	s.Path = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisks) SetStatus(v *GetInstanceResponseBodyCloudDisksStatus) *GetInstanceResponseBodyCloudDisks {
	s.Status = v
	return s
}

func (s *GetInstanceResponseBodyCloudDisks) SetSubType(v string) *GetInstanceResponseBodyCloudDisks {
	s.SubType = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisks) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyCloudDisksStatus struct {
	Available *int64 `json:"Available,omitempty" xml:"Available,omitempty"`
	Capacity  *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Usage     *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetInstanceResponseBodyCloudDisksStatus) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyCloudDisksStatus) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyCloudDisksStatus) GetAvailable() *int64 {
	return s.Available
}

func (s *GetInstanceResponseBodyCloudDisksStatus) GetCapacity() *int64 {
	return s.Capacity
}

func (s *GetInstanceResponseBodyCloudDisksStatus) GetUsage() *int64 {
	return s.Usage
}

func (s *GetInstanceResponseBodyCloudDisksStatus) SetAvailable(v int64) *GetInstanceResponseBodyCloudDisksStatus {
	s.Available = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisksStatus) SetCapacity(v int64) *GetInstanceResponseBodyCloudDisksStatus {
	s.Capacity = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisksStatus) SetUsage(v int64) *GetInstanceResponseBodyCloudDisksStatus {
	s.Usage = &v
	return s
}

func (s *GetInstanceResponseBodyCloudDisksStatus) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDatasets struct {
	// The dataset ID.
	//
	// example:
	//
	// d-vsqjvsjp4orp5l206u
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// Indicates whether dynamic mounting is enabled. Default value: false.
	//
	// example:
	//
	// false
	Dynamic *bool `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	// The read and write permissions. Valid values: RW and RO.
	//
	// example:
	//
	// RW
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The mount path in the container.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount type of the dataset (deprecated).
	//
	// example:
	//
	// FastReadWrite
	OptionType *string `json:"OptionType,omitempty" xml:"OptionType,omitempty"`
	// The mount type of the dataset.
	//
	// example:
	//
	// {
	//
	//   "fs.oss.download.thread.concurrency": "10",
	//
	//   "fs.oss.upload.thread.concurrency": "10",
	//
	//   "fs.jindo.args": "-oattr_timeout=3 -oentry_timeout=0 -onegative_timeout=0 -oauto_cache -ono_symlink"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The dataset URI.
	//
	// example:
	//
	// oss://bucket-name.oss-cn-shanghai-internal.aliyuncs.com/data/path/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetInstanceResponseBodyDatasets) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDatasets) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetInstanceResponseBodyDatasets) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *GetInstanceResponseBodyDatasets) GetDynamic() *bool {
	return s.Dynamic
}

func (s *GetInstanceResponseBodyDatasets) GetMountAccess() *string {
	return s.MountAccess
}

func (s *GetInstanceResponseBodyDatasets) GetMountPath() *string {
	return s.MountPath
}

func (s *GetInstanceResponseBodyDatasets) GetOptionType() *string {
	return s.OptionType
}

func (s *GetInstanceResponseBodyDatasets) GetOptions() *string {
	return s.Options
}

func (s *GetInstanceResponseBodyDatasets) GetUri() *string {
	return s.Uri
}

func (s *GetInstanceResponseBodyDatasets) SetDatasetId(v string) *GetInstanceResponseBodyDatasets {
	s.DatasetId = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) SetDatasetVersion(v string) *GetInstanceResponseBodyDatasets {
	s.DatasetVersion = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) SetDynamic(v bool) *GetInstanceResponseBodyDatasets {
	s.Dynamic = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) SetMountAccess(v string) *GetInstanceResponseBodyDatasets {
	s.MountAccess = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) SetMountPath(v string) *GetInstanceResponseBodyDatasets {
	s.MountPath = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) SetOptionType(v string) *GetInstanceResponseBodyDatasets {
	s.OptionType = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) SetOptions(v string) *GetInstanceResponseBodyDatasets {
	s.Options = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) SetUri(v string) *GetInstanceResponseBodyDatasets {
	s.Uri = &v
	return s
}

func (s *GetInstanceResponseBodyDatasets) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyIdleInstanceCuller struct {
	// The CPU utilization threshold. Unit: percentage. Valid values: 1 to 100. If the CPU utilization of the instance is lower than this threshold, the instance is considered idle.
	//
	// example:
	//
	// 20
	CpuPercentThreshold *int32 `json:"CpuPercentThreshold,omitempty" xml:"CpuPercentThreshold,omitempty"`
	// The GPU utilization threshold. Unit: percentage. Valid values: 1 to 100. This parameter takes effect only if the instance is of the GPU instance type. If both CPU and GPU utilization is lower than the thresholds, the instance is considered idle.
	//
	// example:
	//
	// 10
	GpuPercentThreshold *int32 `json:"GpuPercentThreshold,omitempty" xml:"GpuPercentThreshold,omitempty"`
	// The current time duration for which the instance is idle. Unit: minutes.
	//
	// example:
	//
	// 30
	IdleTimeInMinutes *int32 `json:"IdleTimeInMinutes,omitempty" xml:"IdleTimeInMinutes,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum time duration for which the instance is idle. Unit: minutes. If the time duration for which the instance is idle exceeds this value, the system automatically stops the instance.
	//
	// example:
	//
	// 60
	MaxIdleTimeInMinutes *int32 `json:"MaxIdleTimeInMinutes,omitempty" xml:"MaxIdleTimeInMinutes,omitempty"`
}

func (s GetInstanceResponseBodyIdleInstanceCuller) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyIdleInstanceCuller) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) GetCpuPercentThreshold() *int32 {
	return s.CpuPercentThreshold
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) GetGpuPercentThreshold() *int32 {
	return s.GpuPercentThreshold
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) GetIdleTimeInMinutes() *int32 {
	return s.IdleTimeInMinutes
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) GetMaxIdleTimeInMinutes() *int32 {
	return s.MaxIdleTimeInMinutes
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetCpuPercentThreshold(v int32) *GetInstanceResponseBodyIdleInstanceCuller {
	s.CpuPercentThreshold = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetGpuPercentThreshold(v int32) *GetInstanceResponseBodyIdleInstanceCuller {
	s.GpuPercentThreshold = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetIdleTimeInMinutes(v int32) *GetInstanceResponseBodyIdleInstanceCuller {
	s.IdleTimeInMinutes = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetInstanceId(v string) *GetInstanceResponseBodyIdleInstanceCuller {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) SetMaxIdleTimeInMinutes(v int32) *GetInstanceResponseBodyIdleInstanceCuller {
	s.MaxIdleTimeInMinutes = &v
	return s
}

func (s *GetInstanceResponseBodyIdleInstanceCuller) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceShutdownTimer struct {
	// The scheduled stop time.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The modified time.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The remaining time before the instance is stopped. Unit: milliseconds.
	//
	// example:
	//
	// 3600000
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s GetInstanceResponseBodyInstanceShutdownTimer) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) GetDueTime() *string {
	return s.DueTime
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) GetRemainingTimeInMs() *int64 {
	return s.RemainingTimeInMs
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetDueTime(v string) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.DueTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetGmtCreateTime(v string) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetGmtModifiedTime(v string) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetInstanceId(v string) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) SetRemainingTimeInMs(v int64) *GetInstanceResponseBodyInstanceShutdownTimer {
	s.RemainingTimeInMs = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceShutdownTimer) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceSnapshotList struct {
	// The time when the snapshot was created.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the snapshot was modified.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The image ID.
	//
	// example:
	//
	// image-05cefd0be2exxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// py36_cpu_tf1.12_ubuntu
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image URL.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/pai_product/tensorflow:py36_cpu_tf1.12_ubuntu
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The error code of the instance snapshot.
	//
	// example:
	//
	// Internal Error
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The error message of the instance snapshot.
	//
	// example:
	//
	// ImagePullBackOff
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The image repository URL.
	//
	// example:
	//
	// https://cr.console.aliyun.com/repository/cn-hangzhou/zouxu/kf/images
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	// The instance snapshot status.
	//
	// example:
	//
	// Pushing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyInstanceSnapshotList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceSnapshotList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetImageId() *string {
	return s.ImageId
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetImageName() *string {
	return s.ImageName
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetRepositoryUrl() *string {
	return s.RepositoryUrl
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetGmtCreateTime(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetGmtModifiedTime(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetImageId(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ImageId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetImageName(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ImageName = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetImageUrl(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetReasonCode(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetReasonMessage(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetRepositoryUrl(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.RepositoryUrl = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) SetStatus(v string) *GetInstanceResponseBodyInstanceSnapshotList {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceSnapshotList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyLabels struct {
	// The tag key.
	//
	// example:
	//
	// stsTokenOwner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyLabels) GetKey() *string {
	return s.Key
}

func (s *GetInstanceResponseBodyLabels) GetValue() *string {
	return s.Value
}

func (s *GetInstanceResponseBodyLabels) SetKey(v string) *GetInstanceResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyLabels) SetValue(v string) *GetInstanceResponseBodyLabels {
	s.Value = &v
	return s
}

func (s *GetInstanceResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyLatestSnapshot struct {
	// The time when the snapshot was created.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the snapshot was modified.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The image ID.
	//
	// example:
	//
	// image-05cefd0be2exxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// py36_cpu_tf1.12_ubuntu
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image URL.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/pai_product/tensorflow:py36_cpu_tf1.12_ubuntu
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The error code of the instance snapshot.
	//
	// example:
	//
	// Internal Error
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The error message of the instance snapshot.
	//
	// example:
	//
	// ImagePullBackOff
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The image repository URL.
	//
	// example:
	//
	// https://cr.console.aliyun.com/repository/cn-hangzhou/zouxu/kf/images
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	// The instance snapshot status.
	//
	// Valid values:
	//
	// 	- Committing
	//
	// 	- Pushing
	//
	// 	- Failed
	//
	// 	- Saved
	//
	// example:
	//
	// Pushing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyLatestSnapshot) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyLatestSnapshot) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetImageId() *string {
	return s.ImageId
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetImageName() *string {
	return s.ImageName
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetRepositoryUrl() *string {
	return s.RepositoryUrl
}

func (s *GetInstanceResponseBodyLatestSnapshot) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetGmtCreateTime(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetGmtModifiedTime(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetImageId(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ImageId = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetImageName(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ImageName = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetImageUrl(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ImageUrl = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetReasonCode(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ReasonCode = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetReasonMessage(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.ReasonMessage = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetRepositoryUrl(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) SetStatus(v string) *GetInstanceResponseBodyLatestSnapshot {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyLatestSnapshot) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyNodeErrorRecovery struct {
	// The number of seconds to wait before automatic switchover.
	//
	// example:
	//
	// 30
	AutoSwitchCountdownSeconds *int64 `json:"autoSwitchCountdownSeconds,omitempty" xml:"autoSwitchCountdownSeconds,omitempty"`
	// Indicates whether to enable automatic switchover when a node error occurs.
	//
	// example:
	//
	// true
	EnableAutoSwitchOnNodeError *bool `json:"enableAutoSwitchOnNodeError,omitempty" xml:"enableAutoSwitchOnNodeError,omitempty"`
	// Indicates whether the node has an error.
	//
	// example:
	//
	// false
	HasNodeError *bool `json:"hasNodeError,omitempty" xml:"hasNodeError,omitempty"`
}

func (s GetInstanceResponseBodyNodeErrorRecovery) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyNodeErrorRecovery) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyNodeErrorRecovery) GetAutoSwitchCountdownSeconds() *int64 {
	return s.AutoSwitchCountdownSeconds
}

func (s *GetInstanceResponseBodyNodeErrorRecovery) GetEnableAutoSwitchOnNodeError() *bool {
	return s.EnableAutoSwitchOnNodeError
}

func (s *GetInstanceResponseBodyNodeErrorRecovery) GetHasNodeError() *bool {
	return s.HasNodeError
}

func (s *GetInstanceResponseBodyNodeErrorRecovery) SetAutoSwitchCountdownSeconds(v int64) *GetInstanceResponseBodyNodeErrorRecovery {
	s.AutoSwitchCountdownSeconds = &v
	return s
}

func (s *GetInstanceResponseBodyNodeErrorRecovery) SetEnableAutoSwitchOnNodeError(v bool) *GetInstanceResponseBodyNodeErrorRecovery {
	s.EnableAutoSwitchOnNodeError = &v
	return s
}

func (s *GetInstanceResponseBodyNodeErrorRecovery) SetHasNodeError(v bool) *GetInstanceResponseBodyNodeErrorRecovery {
	s.HasNodeError = &v
	return s
}

func (s *GetInstanceResponseBodyNodeErrorRecovery) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyRequestedResource struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 32
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 4
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// The GPU type. Valid values:
	//
	// 	- V100
	//
	// 	- A100
	//
	// 	- T4
	//
	// 	- A10
	//
	// 	- P100
	//
	// example:
	//
	// v100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 32
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The shared memory size. Unit: GB.
	//
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s GetInstanceResponseBodyRequestedResource) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyRequestedResource) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyRequestedResource) GetCPU() *string {
	return s.CPU
}

func (s *GetInstanceResponseBodyRequestedResource) GetGPU() *string {
	return s.GPU
}

func (s *GetInstanceResponseBodyRequestedResource) GetGPUType() *string {
	return s.GPUType
}

func (s *GetInstanceResponseBodyRequestedResource) GetMemory() *string {
	return s.Memory
}

func (s *GetInstanceResponseBodyRequestedResource) GetSharedMemory() *string {
	return s.SharedMemory
}

func (s *GetInstanceResponseBodyRequestedResource) SetCPU(v string) *GetInstanceResponseBodyRequestedResource {
	s.CPU = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) SetGPU(v string) *GetInstanceResponseBodyRequestedResource {
	s.GPU = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) SetGPUType(v string) *GetInstanceResponseBodyRequestedResource {
	s.GPUType = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) SetMemory(v string) *GetInstanceResponseBodyRequestedResource {
	s.Memory = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) SetSharedMemory(v string) *GetInstanceResponseBodyRequestedResource {
	s.SharedMemory = &v
	return s
}

func (s *GetInstanceResponseBodyRequestedResource) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetInstanceResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetInstanceResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetInstanceResponseBodyTags) SetTagKey(v string) *GetInstanceResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetTagValue(v string) *GetInstanceResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetInstanceResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyUserVpc struct {
	BandwidthLimit *BandwidthLimit `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// Default Route
	//
	// example:
	//
	// eth0 | eth1
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR block.
	//
	// 	- If you leave VSwitchId empty, this parameter is not required and the system automatically obtains all CIDR blocks in the VPC.
	//
	// 	- If VSwitchId is not empty, this parameter is required. Specify all CIDR blocks in the VPC.
	//
	// example:
	//
	// ["192.168.0.1/24", "192.168.1.1/24"]
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The forward information.
	ForwardInfos []*ForwardInfoResponse `json:"ForwardInfos,omitempty" xml:"ForwardInfos,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// example:
	//
	// sg-xxxxxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-xxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetInstanceResponseBodyUserVpc) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyUserVpc) GetBandwidthLimit() *BandwidthLimit {
	return s.BandwidthLimit
}

func (s *GetInstanceResponseBodyUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *GetInstanceResponseBodyUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *GetInstanceResponseBodyUserVpc) GetForwardInfos() []*ForwardInfoResponse {
	return s.ForwardInfos
}

func (s *GetInstanceResponseBodyUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetInstanceResponseBodyUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetInstanceResponseBodyUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceResponseBodyUserVpc) SetBandwidthLimit(v *BandwidthLimit) *GetInstanceResponseBodyUserVpc {
	s.BandwidthLimit = v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetDefaultRoute(v string) *GetInstanceResponseBodyUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetExtendedCIDRs(v []*string) *GetInstanceResponseBodyUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetForwardInfos(v []*ForwardInfoResponse) *GetInstanceResponseBodyUserVpc {
	s.ForwardInfos = v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetSecurityGroupId(v string) *GetInstanceResponseBodyUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetVSwitchId(v string) *GetInstanceResponseBodyUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) SetVpcId(v string) *GetInstanceResponseBodyUserVpc {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyUserVpc) Validate() error {
	if s.BandwidthLimit != nil {
		if err := s.BandwidthLimit.Validate(); err != nil {
			return err
		}
	}
	if s.ForwardInfos != nil {
		for _, item := range s.ForwardInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
