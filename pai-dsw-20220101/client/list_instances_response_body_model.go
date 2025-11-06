// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListInstancesResponseBody
	GetHttpStatusCode() *int32
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListInstancesResponseBody
	GetTotalCount() *int64
}

type ListInstancesResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: an internal error. All errors, except for parameter validation errors, are classified as internal errors.
	//
	// 	- ValidationError: a parameter validation error.
	//
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// The instances returned on this page.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 35
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstances struct {
	// The accelerator type of the instance. Valid values:
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
	// 	- PRIVATE (default): The instances are accessible only to you and the administrator of the workspace.
	//
	// 	- PUBLIC: The instances are accessible only to all members in the workspace.
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
	Affinity *ListInstancesResponseBodyInstancesAffinity `json:"Affinity,omitempty" xml:"Affinity,omitempty" type:"Struct"`
	// The cloud disks of the instance.
	//
	// example:
	//
	// []
	CloudDisks []*ListInstancesResponseBodyInstancesCloudDisks `json:"CloudDisks,omitempty" xml:"CloudDisks,omitempty" type:"Repeated"`
	// The credential configuration.
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The datasets.
	Datasets []*ListInstancesResponseBodyInstancesDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The NVIDIA driver configuration.
	//
	// example:
	//
	// 535.54.03
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// The dynamic mount configurations.
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
	// The time when the instance was created.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the instance was modified.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The rule for stopping idle instances.
	//
	// example:
	//
	// {"InstanceId":"dsw-05cefd0be2e5a278","CpuPercentThreshold":20,"GpuPercentThreshold":10,"MaxIdleTimeInMinutes":120,"IdleTimeInMinutes":30}
	IdleInstanceCuller *ListInstancesResponseBodyInstancesIdleInstanceCuller `json:"IdleInstanceCuller,omitempty" xml:"IdleInstanceCuller,omitempty" type:"Struct"`
	// The Base64-encoded account and password for the user\\"s private image. The password will be hidden.
	//
	// example:
	//
	// aGFyYm9yYWlAeGltYWxheWE6KioqKioq
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
	// The scheduled stop task.
	InstanceShutdownTimer *ListInstancesResponseBodyInstancesInstanceShutdownTimer `json:"InstanceShutdownTimer,omitempty" xml:"InstanceShutdownTimer,omitempty" type:"Struct"`
	// The instance snapshots.
	//
	// example:
	//
	// []
	InstanceSnapshotList []*ListInstancesResponseBodyInstancesInstanceSnapshotList `json:"InstanceSnapshotList,omitempty" xml:"InstanceSnapshotList,omitempty" type:"Repeated"`
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
	// The custom labels.
	//
	// example:
	//
	// {\\"foo\\": \\"bar\\"}
	Labels []*ListInstancesResponseBodyInstancesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The user image that was latest saved.
	LatestSnapshot *ListInstancesResponseBodyInstancesLatestSnapshot `json:"LatestSnapshot,omitempty" xml:"LatestSnapshot,omitempty" type:"Struct"`
	OversoldInfo   *string                                           `json:"OversoldInfo,omitempty" xml:"OversoldInfo,omitempty"`
	OversoldType   *string                                           `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
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
	// The priority based on which resources are allocated to instances. Resources are preferentially allocated to instances with higher priorities.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
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
	// The resource configurations.
	//
	// example:
	//
	// {"CPU":"4","Memory":"8Gi","SharedMemory":"4Gi","GPU":"1","GPUType":"Tesla-V100-16G"}
	RequestedResource *ListInstancesResponseBodyInstancesRequestedResource `json:"RequestedResource,omitempty" xml:"RequestedResource,omitempty" type:"Struct"`
	// The resource ID. This parameter is valid only if you set PaymentType to Subscription.
	//
	// example:
	//
	// dsw-123456789
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The specifications.
	//
	// 	- In pay-as-you-go scenarios, the value is the specifications of the purchased ECS instance type.
	//
	// 	- In subscription scenarios, the value is the requested number of CPU cores and memory size.
	//
	// example:
	//
	// resource_group
	ResourceName  *string        `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ServiceConfig *ServiceConfig `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListInstancesResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The terminal URL.
	//
	// example:
	//
	// https://dsw-gateway-cn-shanghai.aliyun.com/dsw-39772/tty/
	TerminalUrl *string `json:"TerminalUrl,omitempty" xml:"TerminalUrl,omitempty"`
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
	UserVpc *ListInstancesResponseBodyInstancesUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
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

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListInstancesResponseBodyInstances) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListInstancesResponseBodyInstances) GetAccumulatedRunningTimeInMs() *int64 {
	return s.AccumulatedRunningTimeInMs
}

func (s *ListInstancesResponseBodyInstances) GetAffinity() *ListInstancesResponseBodyInstancesAffinity {
	return s.Affinity
}

func (s *ListInstancesResponseBodyInstances) GetCloudDisks() []*ListInstancesResponseBodyInstancesCloudDisks {
	return s.CloudDisks
}

func (s *ListInstancesResponseBodyInstances) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *ListInstancesResponseBodyInstances) GetDatasets() []*ListInstancesResponseBodyInstancesDatasets {
	return s.Datasets
}

func (s *ListInstancesResponseBodyInstances) GetDriver() *string {
	return s.Driver
}

func (s *ListInstancesResponseBodyInstances) GetDynamicMount() *DynamicMount {
	return s.DynamicMount
}

func (s *ListInstancesResponseBodyInstances) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *ListInstancesResponseBodyInstances) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *ListInstancesResponseBodyInstances) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListInstancesResponseBodyInstances) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListInstancesResponseBodyInstances) GetIdleInstanceCuller() *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	return s.IdleInstanceCuller
}

func (s *ListInstancesResponseBodyInstances) GetImageAuth() *string {
	return s.ImageAuth
}

func (s *ListInstancesResponseBodyInstances) GetImageId() *string {
	return s.ImageId
}

func (s *ListInstancesResponseBodyInstances) GetImageName() *string {
	return s.ImageName
}

func (s *ListInstancesResponseBodyInstances) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyInstances) GetInstanceShutdownTimer() *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	return s.InstanceShutdownTimer
}

func (s *ListInstancesResponseBodyInstances) GetInstanceSnapshotList() []*ListInstancesResponseBodyInstancesInstanceSnapshotList {
	return s.InstanceSnapshotList
}

func (s *ListInstancesResponseBodyInstances) GetInstanceUrl() *string {
	return s.InstanceUrl
}

func (s *ListInstancesResponseBodyInstances) GetJupyterlabUrl() *string {
	return s.JupyterlabUrl
}

func (s *ListInstancesResponseBodyInstances) GetLabels() []*ListInstancesResponseBodyInstancesLabels {
	return s.Labels
}

func (s *ListInstancesResponseBodyInstances) GetLatestSnapshot() *ListInstancesResponseBodyInstancesLatestSnapshot {
	return s.LatestSnapshot
}

func (s *ListInstancesResponseBodyInstances) GetOversoldInfo() *string {
	return s.OversoldInfo
}

func (s *ListInstancesResponseBodyInstances) GetOversoldType() *string {
	return s.OversoldType
}

func (s *ListInstancesResponseBodyInstances) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstancesResponseBodyInstances) GetPriority() *int64 {
	return s.Priority
}

func (s *ListInstancesResponseBodyInstances) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *ListInstancesResponseBodyInstances) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *ListInstancesResponseBodyInstances) GetRequestedResource() *ListInstancesResponseBodyInstancesRequestedResource {
	return s.RequestedResource
}

func (s *ListInstancesResponseBodyInstances) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListInstancesResponseBodyInstances) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListInstancesResponseBodyInstances) GetServiceConfig() *ServiceConfig {
	return s.ServiceConfig
}

func (s *ListInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstances) GetTags() []*ListInstancesResponseBodyInstancesTags {
	return s.Tags
}

func (s *ListInstancesResponseBodyInstances) GetTerminalUrl() *string {
	return s.TerminalUrl
}

func (s *ListInstancesResponseBodyInstances) GetUserId() *string {
	return s.UserId
}

func (s *ListInstancesResponseBodyInstances) GetUserName() *string {
	return s.UserName
}

func (s *ListInstancesResponseBodyInstances) GetUserVpc() *ListInstancesResponseBodyInstancesUserVpc {
	return s.UserVpc
}

func (s *ListInstancesResponseBodyInstances) GetWebIDEUrl() *string {
	return s.WebIDEUrl
}

func (s *ListInstancesResponseBodyInstances) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListInstancesResponseBodyInstances) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListInstancesResponseBodyInstances) GetWorkspaceSource() *string {
	return s.WorkspaceSource
}

func (s *ListInstancesResponseBodyInstances) SetAcceleratorType(v string) *ListInstancesResponseBodyInstances {
	s.AcceleratorType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetAccessibility(v string) *ListInstancesResponseBodyInstances {
	s.Accessibility = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetAccumulatedRunningTimeInMs(v int64) *ListInstancesResponseBodyInstances {
	s.AccumulatedRunningTimeInMs = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetAffinity(v *ListInstancesResponseBodyInstancesAffinity) *ListInstancesResponseBodyInstances {
	s.Affinity = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCloudDisks(v []*ListInstancesResponseBodyInstancesCloudDisks) *ListInstancesResponseBodyInstances {
	s.CloudDisks = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCredentialConfig(v *CredentialConfig) *ListInstancesResponseBodyInstances {
	s.CredentialConfig = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDatasets(v []*ListInstancesResponseBodyInstancesDatasets) *ListInstancesResponseBodyInstances {
	s.Datasets = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDriver(v string) *ListInstancesResponseBodyInstances {
	s.Driver = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDynamicMount(v *DynamicMount) *ListInstancesResponseBodyInstances {
	s.DynamicMount = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetEcsSpec(v string) *ListInstancesResponseBodyInstances {
	s.EcsSpec = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetEnvironmentVariables(v map[string]*string) *ListInstancesResponseBodyInstances {
	s.EnvironmentVariables = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetIdleInstanceCuller(v *ListInstancesResponseBodyInstancesIdleInstanceCuller) *ListInstancesResponseBodyInstances {
	s.IdleInstanceCuller = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageAuth(v string) *ListInstancesResponseBodyInstances {
	s.ImageAuth = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageId(v string) *ListInstancesResponseBodyInstances {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageName(v string) *ListInstancesResponseBodyInstances {
	s.ImageName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageUrl(v string) *ListInstancesResponseBodyInstances {
	s.ImageUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceShutdownTimer(v *ListInstancesResponseBodyInstancesInstanceShutdownTimer) *ListInstancesResponseBodyInstances {
	s.InstanceShutdownTimer = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceSnapshotList(v []*ListInstancesResponseBodyInstancesInstanceSnapshotList) *ListInstancesResponseBodyInstances {
	s.InstanceSnapshotList = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceUrl(v string) *ListInstancesResponseBodyInstances {
	s.InstanceUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetJupyterlabUrl(v string) *ListInstancesResponseBodyInstances {
	s.JupyterlabUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetLabels(v []*ListInstancesResponseBodyInstancesLabels) *ListInstancesResponseBodyInstances {
	s.Labels = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetLatestSnapshot(v *ListInstancesResponseBodyInstancesLatestSnapshot) *ListInstancesResponseBodyInstances {
	s.LatestSnapshot = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetOversoldInfo(v string) *ListInstancesResponseBodyInstances {
	s.OversoldInfo = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetOversoldType(v string) *ListInstancesResponseBodyInstances {
	s.OversoldType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPaymentType(v string) *ListInstancesResponseBodyInstances {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPriority(v int64) *ListInstancesResponseBodyInstances {
	s.Priority = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetReasonCode(v string) *ListInstancesResponseBodyInstances {
	s.ReasonCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetReasonMessage(v string) *ListInstancesResponseBodyInstances {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRequestedResource(v *ListInstancesResponseBodyInstancesRequestedResource) *ListInstancesResponseBodyInstances {
	s.RequestedResource = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceId(v string) *ListInstancesResponseBodyInstances {
	s.ResourceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceName(v string) *ListInstancesResponseBodyInstances {
	s.ResourceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetServiceConfig(v *ServiceConfig) *ListInstancesResponseBodyInstances {
	s.ServiceConfig = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetTags(v []*ListInstancesResponseBodyInstancesTags) *ListInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetTerminalUrl(v string) *ListInstancesResponseBodyInstances {
	s.TerminalUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserId(v string) *ListInstancesResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserName(v string) *ListInstancesResponseBodyInstances {
	s.UserName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserVpc(v *ListInstancesResponseBodyInstancesUserVpc) *ListInstancesResponseBodyInstances {
	s.UserVpc = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetWebIDEUrl(v string) *ListInstancesResponseBodyInstances {
	s.WebIDEUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetWorkspaceId(v string) *ListInstancesResponseBodyInstances {
	s.WorkspaceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetWorkspaceName(v string) *ListInstancesResponseBodyInstances {
	s.WorkspaceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetWorkspaceSource(v string) *ListInstancesResponseBodyInstances {
	s.WorkspaceSource = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
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
	if s.RequestedResource != nil {
		if err := s.RequestedResource.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceConfig != nil {
		if err := s.ServiceConfig.Validate(); err != nil {
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

type ListInstancesResponseBodyInstancesAffinity struct {
	// The CPU affinity configuration. Only subscription instances that use general-purpose computing resources support CPU affinity configuration.
	CPU *ListInstancesResponseBodyInstancesAffinityCPU `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
}

func (s ListInstancesResponseBodyInstancesAffinity) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesAffinity) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesAffinity) GetCPU() *ListInstancesResponseBodyInstancesAffinityCPU {
	return s.CPU
}

func (s *ListInstancesResponseBodyInstancesAffinity) SetCPU(v *ListInstancesResponseBodyInstancesAffinityCPU) *ListInstancesResponseBodyInstancesAffinity {
	s.CPU = v
	return s
}

func (s *ListInstancesResponseBodyInstancesAffinity) Validate() error {
	if s.CPU != nil {
		if err := s.CPU.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyInstancesAffinityCPU struct {
	// Indicates whether the CPU affinity feature was enabled.
	//
	// true false
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ListInstancesResponseBodyInstancesAffinityCPU) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesAffinityCPU) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesAffinityCPU) GetEnable() *bool {
	return s.Enable
}

func (s *ListInstancesResponseBodyInstancesAffinityCPU) SetEnable(v bool) *ListInstancesResponseBodyInstancesAffinityCPU {
	s.Enable = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesAffinityCPU) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesCloudDisks struct {
	// The cloud disk capacity.
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
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The cloud disk type. The value rootfs indicates that the cloud disk is used as the root file system (rootfs).
	//
	// example:
	//
	// rootfs
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s ListInstancesResponseBodyInstancesCloudDisks) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesCloudDisks) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) GetCapacity() *string {
	return s.Capacity
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) GetMountPath() *string {
	return s.MountPath
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) GetPath() *string {
	return s.Path
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) GetSubType() *string {
	return s.SubType
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) SetCapacity(v string) *ListInstancesResponseBodyInstancesCloudDisks {
	s.Capacity = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) SetMountPath(v string) *ListInstancesResponseBodyInstancesCloudDisks {
	s.MountPath = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) SetPath(v string) *ListInstancesResponseBodyInstancesCloudDisks {
	s.Path = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) SetSubType(v string) *ListInstancesResponseBodyInstancesCloudDisks {
	s.SubType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesCloudDisks) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesDatasets struct {
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
	// Indicates whether dynamic mounting was enabled. Default value: false.
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
	// The type of the mount option.
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

func (s ListInstancesResponseBodyInstancesDatasets) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDatasets) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDatasets) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ListInstancesResponseBodyInstancesDatasets) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *ListInstancesResponseBodyInstancesDatasets) GetDynamic() *bool {
	return s.Dynamic
}

func (s *ListInstancesResponseBodyInstancesDatasets) GetMountAccess() *string {
	return s.MountAccess
}

func (s *ListInstancesResponseBodyInstancesDatasets) GetMountPath() *string {
	return s.MountPath
}

func (s *ListInstancesResponseBodyInstancesDatasets) GetOptionType() *string {
	return s.OptionType
}

func (s *ListInstancesResponseBodyInstancesDatasets) GetOptions() *string {
	return s.Options
}

func (s *ListInstancesResponseBodyInstancesDatasets) GetUri() *string {
	return s.Uri
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetDatasetId(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.DatasetId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetDatasetVersion(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.DatasetVersion = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetDynamic(v bool) *ListInstancesResponseBodyInstancesDatasets {
	s.Dynamic = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetMountAccess(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.MountAccess = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetMountPath(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.MountPath = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetOptionType(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.OptionType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetOptions(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.Options = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) SetUri(v string) *ListInstancesResponseBodyInstancesDatasets {
	s.Uri = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDatasets) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesIdleInstanceCuller struct {
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
	// The time duration for which the instance is idle. Unit: minutes.
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

func (s ListInstancesResponseBodyInstancesIdleInstanceCuller) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesIdleInstanceCuller) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) GetCpuPercentThreshold() *int32 {
	return s.CpuPercentThreshold
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) GetGpuPercentThreshold() *int32 {
	return s.GpuPercentThreshold
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) GetIdleTimeInMinutes() *int32 {
	return s.IdleTimeInMinutes
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) GetMaxIdleTimeInMinutes() *int32 {
	return s.MaxIdleTimeInMinutes
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetCpuPercentThreshold(v int32) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.CpuPercentThreshold = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetGpuPercentThreshold(v int32) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.GpuPercentThreshold = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetIdleTimeInMinutes(v int32) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.IdleTimeInMinutes = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetInstanceId(v string) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) SetMaxIdleTimeInMinutes(v int32) *ListInstancesResponseBodyInstancesIdleInstanceCuller {
	s.MaxIdleTimeInMinutes = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesIdleInstanceCuller) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesInstanceShutdownTimer struct {
	// The scheduled stop time.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the instance was modified.
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
	// The remaining time before the instance is stopped.
	//
	// example:
	//
	// 3600000
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s ListInstancesResponseBodyInstancesInstanceShutdownTimer) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstanceShutdownTimer) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) GetDueTime() *string {
	return s.DueTime
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) GetRemainingTimeInMs() *int64 {
	return s.RemainingTimeInMs
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetDueTime(v string) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.DueTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetInstanceId(v string) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) SetRemainingTimeInMs(v int64) *ListInstancesResponseBodyInstancesInstanceShutdownTimer {
	s.RemainingTimeInMs = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceShutdownTimer) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesInstanceSnapshotList struct {
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
	// The URL of the image repository.
	//
	// example:
	//
	// https://cr.console.aliyun.com/repository/cn-hangzhou/zouxu/kf/images
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	// The status of the instance snapshot.
	//
	// example:
	//
	// Pushing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesInstanceSnapshotList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstanceSnapshotList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetImageId() *string {
	return s.ImageId
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetImageName() *string {
	return s.ImageName
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetRepositoryUrl() *string {
	return s.RepositoryUrl
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetImageId(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetImageName(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ImageName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetImageUrl(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ImageUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetReasonCode(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ReasonCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetReasonMessage(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetRepositoryUrl(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.RepositoryUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) SetStatus(v string) *ListInstancesResponseBodyInstancesInstanceSnapshotList {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceSnapshotList) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesLabels struct {
	// The custom label key.
	//
	// example:
	//
	// stsTokenOwner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The custom label value.
	//
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesLabels) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesLabels) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesLabels) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyInstancesLabels) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyInstancesLabels) SetKey(v string) *ListInstancesResponseBodyInstancesLabels {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLabels) SetValue(v string) *ListInstancesResponseBodyInstancesLabels {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLabels) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesLatestSnapshot struct {
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
	// The URL of the image repository.
	//
	// example:
	//
	// https://cr.console.aliyun.com/repository/cn-hangzhou/zouxu/kf/images
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	// The status of the instance snapshot.
	//
	// example:
	//
	// Pushing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesLatestSnapshot) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesLatestSnapshot) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetImageId() *string {
	return s.ImageId
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetImageName() *string {
	return s.ImageName
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetRepositoryUrl() *string {
	return s.RepositoryUrl
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetImageId(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetImageName(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ImageName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetImageUrl(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ImageUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetReasonCode(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ReasonCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetReasonMessage(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.ReasonMessage = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetRepositoryUrl(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.RepositoryUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) SetStatus(v string) *ListInstancesResponseBodyInstancesLatestSnapshot {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesLatestSnapshot) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesRequestedResource struct {
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
	// The GPU memory type.
	//
	// example:
	//
	// v100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 32
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The size of the shared memory.
	//
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s ListInstancesResponseBodyInstancesRequestedResource) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesRequestedResource) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) GetCPU() *string {
	return s.CPU
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) GetGPU() *string {
	return s.GPU
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) GetGPUType() *string {
	return s.GPUType
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) GetMemory() *string {
	return s.Memory
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) GetSharedMemory() *string {
	return s.SharedMemory
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetCPU(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.CPU = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetGPU(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.GPU = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetGPUType(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.GPUType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetMemory(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.Memory = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) SetSharedMemory(v string) *ListInstancesResponseBodyInstancesRequestedResource {
	s.SharedMemory = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesRequestedResource) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesTags struct {
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

func (s ListInstancesResponseBodyInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListInstancesResponseBodyInstancesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListInstancesResponseBodyInstancesTags) SetTagKey(v string) *ListInstancesResponseBodyInstancesTags {
	s.TagKey = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesTags) SetTagValue(v string) *ListInstancesResponseBodyInstancesTags {
	s.TagValue = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesTags) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesUserVpc struct {
	BandwidthLimit *BandwidthLimit `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The default route.
	//
	// example:
	//
	// eth0 | eth1
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR blocks.
	//
	// example:
	//
	// ["192.168.0.1/24", "192.168.1.1/24"]
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The forward information.
	ForwardInfos []*ForwardInfoResponse `json:"ForwardInfos,omitempty" xml:"ForwardInfos,omitempty" type:"Repeated"`
	Ip           *string                `json:"Ip,omitempty" xml:"Ip,omitempty"`
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

func (s ListInstancesResponseBodyInstancesUserVpc) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesUserVpc) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesUserVpc) GetBandwidthLimit() *BandwidthLimit {
	return s.BandwidthLimit
}

func (s *ListInstancesResponseBodyInstancesUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *ListInstancesResponseBodyInstancesUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *ListInstancesResponseBodyInstancesUserVpc) GetForwardInfos() []*ForwardInfoResponse {
	return s.ForwardInfos
}

func (s *ListInstancesResponseBodyInstancesUserVpc) GetIp() *string {
	return s.Ip
}

func (s *ListInstancesResponseBodyInstancesUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListInstancesResponseBodyInstancesUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListInstancesResponseBodyInstancesUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetBandwidthLimit(v *BandwidthLimit) *ListInstancesResponseBodyInstancesUserVpc {
	s.BandwidthLimit = v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetDefaultRoute(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetExtendedCIDRs(v []*string) *ListInstancesResponseBodyInstancesUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetForwardInfos(v []*ForwardInfoResponse) *ListInstancesResponseBodyInstancesUserVpc {
	s.ForwardInfos = v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetIp(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.Ip = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetSecurityGroupId(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetVSwitchId(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) SetVpcId(v string) *ListInstancesResponseBodyInstancesUserVpc {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesUserVpc) Validate() error {
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
