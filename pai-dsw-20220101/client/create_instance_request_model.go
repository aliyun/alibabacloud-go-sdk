// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateInstanceRequest
	GetAccessibility() *string
	SetAffinity(v *CreateInstanceRequestAffinity) *CreateInstanceRequest
	GetAffinity() *CreateInstanceRequestAffinity
	SetAssignNodeSpec(v *CreateInstanceRequestAssignNodeSpec) *CreateInstanceRequest
	GetAssignNodeSpec() *CreateInstanceRequestAssignNodeSpec
	SetCloudDisks(v []*CreateInstanceRequestCloudDisks) *CreateInstanceRequest
	GetCloudDisks() []*CreateInstanceRequestCloudDisks
	SetCredentialConfig(v *CredentialConfig) *CreateInstanceRequest
	GetCredentialConfig() *CredentialConfig
	SetDatasets(v []*CreateInstanceRequestDatasets) *CreateInstanceRequest
	GetDatasets() []*CreateInstanceRequestDatasets
	SetDriver(v string) *CreateInstanceRequest
	GetDriver() *string
	SetDynamicMount(v *DynamicMount) *CreateInstanceRequest
	GetDynamicMount() *DynamicMount
	SetEcsSpec(v string) *CreateInstanceRequest
	GetEcsSpec() *string
	SetEnvironmentVariables(v map[string]*string) *CreateInstanceRequest
	GetEnvironmentVariables() map[string]*string
	SetImageAuth(v string) *CreateInstanceRequest
	GetImageAuth() *string
	SetImageId(v string) *CreateInstanceRequest
	GetImageId() *string
	SetImageUrl(v string) *CreateInstanceRequest
	GetImageUrl() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetLabels(v []*CreateInstanceRequestLabels) *CreateInstanceRequest
	GetLabels() []*CreateInstanceRequestLabels
	SetOversoldType(v string) *CreateInstanceRequest
	GetOversoldType() *string
	SetPriority(v int64) *CreateInstanceRequest
	GetPriority() *int64
	SetRequestedResource(v *CreateInstanceRequestRequestedResource) *CreateInstanceRequest
	GetRequestedResource() *CreateInstanceRequestRequestedResource
	SetResourceId(v string) *CreateInstanceRequest
	GetResourceId() *string
	SetSpotSpec(v *CreateInstanceRequestSpotSpec) *CreateInstanceRequest
	GetSpotSpec() *CreateInstanceRequestSpotSpec
	SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest
	GetTag() []*CreateInstanceRequestTag
	SetUserCommand(v *CreateInstanceRequestUserCommand) *CreateInstanceRequest
	GetUserCommand() *CreateInstanceRequestUserCommand
	SetUserId(v string) *CreateInstanceRequest
	GetUserId() *string
	SetUserVpc(v *CreateInstanceRequestUserVpc) *CreateInstanceRequest
	GetUserVpc() *CreateInstanceRequestUserVpc
	SetWorkspaceId(v string) *CreateInstanceRequest
	GetWorkspaceId() *string
	SetWorkspaceSource(v string) *CreateInstanceRequest
	GetWorkspaceSource() *string
}

type CreateInstanceRequest struct {
	// The instance accessibility.
	//
	// Valid values:
	//
	// 	- PUBLIC: The instances are accessible to all members in the workspace.
	//
	// 	- PRIVATE: The instances are accessible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The affinity configuration.
	Affinity       *CreateInstanceRequestAffinity       `json:"Affinity,omitempty" xml:"Affinity,omitempty" type:"Struct"`
	AssignNodeSpec *CreateInstanceRequestAssignNodeSpec `json:"AssignNodeSpec,omitempty" xml:"AssignNodeSpec,omitempty" type:"Struct"`
	// The cloud disks.
	//
	// example:
	//
	// []
	CloudDisks []*CreateInstanceRequestCloudDisks `json:"CloudDisks,omitempty" xml:"CloudDisks,omitempty" type:"Repeated"`
	// The credential configuration.
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The datasets.
	Datasets []*CreateInstanceRequestDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The NVIDIA driver configuration.
	//
	// example:
	//
	// 535.54.03
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// The dynamic mount configuration.
	DynamicMount *DynamicMount `json:"DynamicMount,omitempty" xml:"DynamicMount,omitempty"`
	// The ECS instance type of the instance. You can call [ListEcsSpecs](https://help.aliyun.com/document_detail/470423.html) to obtain the ECS instance type.
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
	// The Base64-encoded account and password for the user\\"s private image. The password will be hidden.
	//
	// example:
	//
	// ****
	ImageAuth *string `json:"ImageAuth,omitempty" xml:"ImageAuth,omitempty"`
	// The image ID. You can call [ListImages](https://help.aliyun.com/document_detail/449118.html) to obtain the image ID.
	//
	// example:
	//
	// image-05cefd0be2exxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image address. You can call [ListImages](https://help.aliyun.com/document_detail/449118.html) to obtain the image address.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/pai_product/tensorflow:py36_cpu_tf1.12_ubuntu
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The instance name. The name must meet the following requirements:
	//
	// 	- The name can contain only letters, digits, and underscores (_).
	//
	// 	- The name can be up to 27 characters in length.
	//
	// example:
	//
	// training_data
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The custom labels.
	//
	// example:
	//
	// {\\"foo\\": \\"bar\\"}
	Labels       []*CreateInstanceRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OversoldType *string                        `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
	// The priority based on which resources are allocated to instances. Valid values: 1 to 9.
	//
	// 	- 1: the lowest priority.
	//
	// 	- 9: the highest priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource configurations.
	//
	// example:
	//
	// {"CPU":"4","Memory":"8Gi","SharedMemory":"4Gi","GPU":"1","GPUType":"Tesla-V100-16G"}
	RequestedResource *CreateInstanceRequestRequestedResource `json:"RequestedResource,omitempty" xml:"RequestedResource,omitempty" type:"Struct"`
	// The ID of the resource group. This parameter is configured during prepayment. For information about how to create a dedicated resource group, see [Create a dedicated resource group and purchase general computing resources](https://help.aliyun.com/document_detail/202827.html).
	//
	// example:
	//
	// dsw-123456789
	ResourceId *string                        `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	SpotSpec   *CreateInstanceRequestSpotSpec `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty" type:"Struct"`
	// The tags.
	Tag         []*CreateInstanceRequestTag       `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UserCommand *CreateInstanceRequestUserCommand `json:"UserCommand,omitempty" xml:"UserCommand,omitempty" type:"Struct"`
	// The ID of the instance owner. Valid values: Alibaba Cloud account and RAM user.
	//
	// example:
	//
	// 161228528250****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The virtual private cloud (VPC) configurations.
	UserVpc *CreateInstanceRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 40823
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The storage corresponding to the working directory. You can mount disks or datasets to /mnt/workspace at the same time. OSS datasets and dynamically mounted datasets are not supported.
	//
	// Valid values:
	//
	// 	- rootfsCloudDisk: Mount the disk to the working directory.
	//
	// 	- Mount path of the dataset, such as /mnt/data: Datasets in URI format only support this method.
	//
	// 	- Dataset ID, such as d-vsqjvs\\*\\*\\*\\*rp5l206u: If a single dataset is mounted to multiple paths, the first path is selected. We recommend that you do not use this method, use the mount path instead.
	//
	// If you leave this parameter empty:
	//
	// 	- If the instance uses cloud disks, cloud disks are selected by default.
	//
	// 	- if no cloud disks are available, the first NAS or CPFS dataset is selected as the working directory.
	//
	// 	- If no cloud disks, and NAS or CPFS datasets are available, the host space is used.
	//
	// example:
	//
	// rootfsCloudDisk
	WorkspaceSource *string `json:"WorkspaceSource,omitempty" xml:"WorkspaceSource,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateInstanceRequest) GetAffinity() *CreateInstanceRequestAffinity {
	return s.Affinity
}

func (s *CreateInstanceRequest) GetAssignNodeSpec() *CreateInstanceRequestAssignNodeSpec {
	return s.AssignNodeSpec
}

func (s *CreateInstanceRequest) GetCloudDisks() []*CreateInstanceRequestCloudDisks {
	return s.CloudDisks
}

func (s *CreateInstanceRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateInstanceRequest) GetDatasets() []*CreateInstanceRequestDatasets {
	return s.Datasets
}

func (s *CreateInstanceRequest) GetDriver() *string {
	return s.Driver
}

func (s *CreateInstanceRequest) GetDynamicMount() *DynamicMount {
	return s.DynamicMount
}

func (s *CreateInstanceRequest) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *CreateInstanceRequest) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateInstanceRequest) GetImageAuth() *string {
	return s.ImageAuth
}

func (s *CreateInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateInstanceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetLabels() []*CreateInstanceRequestLabels {
	return s.Labels
}

func (s *CreateInstanceRequest) GetOversoldType() *string {
	return s.OversoldType
}

func (s *CreateInstanceRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *CreateInstanceRequest) GetRequestedResource() *CreateInstanceRequestRequestedResource {
	return s.RequestedResource
}

func (s *CreateInstanceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateInstanceRequest) GetSpotSpec() *CreateInstanceRequestSpotSpec {
	return s.SpotSpec
}

func (s *CreateInstanceRequest) GetTag() []*CreateInstanceRequestTag {
	return s.Tag
}

func (s *CreateInstanceRequest) GetUserCommand() *CreateInstanceRequestUserCommand {
	return s.UserCommand
}

func (s *CreateInstanceRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateInstanceRequest) GetUserVpc() *CreateInstanceRequestUserVpc {
	return s.UserVpc
}

func (s *CreateInstanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateInstanceRequest) GetWorkspaceSource() *string {
	return s.WorkspaceSource
}

func (s *CreateInstanceRequest) SetAccessibility(v string) *CreateInstanceRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateInstanceRequest) SetAffinity(v *CreateInstanceRequestAffinity) *CreateInstanceRequest {
	s.Affinity = v
	return s
}

func (s *CreateInstanceRequest) SetAssignNodeSpec(v *CreateInstanceRequestAssignNodeSpec) *CreateInstanceRequest {
	s.AssignNodeSpec = v
	return s
}

func (s *CreateInstanceRequest) SetCloudDisks(v []*CreateInstanceRequestCloudDisks) *CreateInstanceRequest {
	s.CloudDisks = v
	return s
}

func (s *CreateInstanceRequest) SetCredentialConfig(v *CredentialConfig) *CreateInstanceRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateInstanceRequest) SetDatasets(v []*CreateInstanceRequestDatasets) *CreateInstanceRequest {
	s.Datasets = v
	return s
}

func (s *CreateInstanceRequest) SetDriver(v string) *CreateInstanceRequest {
	s.Driver = &v
	return s
}

func (s *CreateInstanceRequest) SetDynamicMount(v *DynamicMount) *CreateInstanceRequest {
	s.DynamicMount = v
	return s
}

func (s *CreateInstanceRequest) SetEcsSpec(v string) *CreateInstanceRequest {
	s.EcsSpec = &v
	return s
}

func (s *CreateInstanceRequest) SetEnvironmentVariables(v map[string]*string) *CreateInstanceRequest {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateInstanceRequest) SetImageAuth(v string) *CreateInstanceRequest {
	s.ImageAuth = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetImageUrl(v string) *CreateInstanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetLabels(v []*CreateInstanceRequestLabels) *CreateInstanceRequest {
	s.Labels = v
	return s
}

func (s *CreateInstanceRequest) SetOversoldType(v string) *CreateInstanceRequest {
	s.OversoldType = &v
	return s
}

func (s *CreateInstanceRequest) SetPriority(v int64) *CreateInstanceRequest {
	s.Priority = &v
	return s
}

func (s *CreateInstanceRequest) SetRequestedResource(v *CreateInstanceRequestRequestedResource) *CreateInstanceRequest {
	s.RequestedResource = v
	return s
}

func (s *CreateInstanceRequest) SetResourceId(v string) *CreateInstanceRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotSpec(v *CreateInstanceRequestSpotSpec) *CreateInstanceRequest {
	s.SpotSpec = v
	return s
}

func (s *CreateInstanceRequest) SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateInstanceRequest) SetUserCommand(v *CreateInstanceRequestUserCommand) *CreateInstanceRequest {
	s.UserCommand = v
	return s
}

func (s *CreateInstanceRequest) SetUserId(v string) *CreateInstanceRequest {
	s.UserId = &v
	return s
}

func (s *CreateInstanceRequest) SetUserVpc(v *CreateInstanceRequestUserVpc) *CreateInstanceRequest {
	s.UserVpc = v
	return s
}

func (s *CreateInstanceRequest) SetWorkspaceId(v string) *CreateInstanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateInstanceRequest) SetWorkspaceSource(v string) *CreateInstanceRequest {
	s.WorkspaceSource = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestAffinity struct {
	// The CPU affinity configuration. Only subscription instances that use general-purpose computing resources support CPU affinity configuration.
	CPU *CreateInstanceRequestAffinityCPU `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
}

func (s CreateInstanceRequestAffinity) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestAffinity) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestAffinity) GetCPU() *CreateInstanceRequestAffinityCPU {
	return s.CPU
}

func (s *CreateInstanceRequestAffinity) SetCPU(v *CreateInstanceRequestAffinityCPU) *CreateInstanceRequestAffinity {
	s.CPU = v
	return s
}

func (s *CreateInstanceRequestAffinity) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestAffinityCPU struct {
	// Specifies whether to enable the CPU affinity feature.
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s CreateInstanceRequestAffinityCPU) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestAffinityCPU) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestAffinityCPU) GetEnable() *bool {
	return s.Enable
}

func (s *CreateInstanceRequestAffinityCPU) SetEnable(v bool) *CreateInstanceRequestAffinityCPU {
	s.Enable = &v
	return s
}

func (s *CreateInstanceRequestAffinityCPU) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestAssignNodeSpec struct {
	// example:
	//
	// node-b
	AntiAffinityNodeNames *string `json:"AntiAffinityNodeNames,omitempty" xml:"AntiAffinityNodeNames,omitempty"`
	// example:
	//
	// node-a
	NodeNames *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
}

func (s CreateInstanceRequestAssignNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestAssignNodeSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestAssignNodeSpec) GetAntiAffinityNodeNames() *string {
	return s.AntiAffinityNodeNames
}

func (s *CreateInstanceRequestAssignNodeSpec) GetNodeNames() *string {
	return s.NodeNames
}

func (s *CreateInstanceRequestAssignNodeSpec) SetAntiAffinityNodeNames(v string) *CreateInstanceRequestAssignNodeSpec {
	s.AntiAffinityNodeNames = &v
	return s
}

func (s *CreateInstanceRequestAssignNodeSpec) SetNodeNames(v string) *CreateInstanceRequestAssignNodeSpec {
	s.NodeNames = &v
	return s
}

func (s *CreateInstanceRequestAssignNodeSpec) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestCloudDisks struct {
	// If **Resource Type*	- is **Public Resource*	- or if **Resource Quota*	- is subscription-based general-purpose computing resources (CPU cores ≥ 2 and memory ≥ 4 GB, or configured with GPU):
	//
	// Each instance has a free system disk of 100 GiB for persistent storage. **If the DSW instance is stopped and not launched for more than 15 days, the disk is cleared**. The disk can be expanded. For specific pricing, refer to the console.
	//
	// **
	//
	// **Warning**
	//
	// 	- After the expansion, you cannot reduce the storage space. Proceed with caution.
	//
	// 	- After the expansion, the disk is not cleared if the instance is stopped for more than 15 days. However, it will continue to incur fees.
	//
	// 	- If you delete the instance, the system disk is also released and the data stored in the disk is deleted. Make sure that you have backed up your data before you delete the instance.
	//
	// If you need persistent storage, you can **mount a dataset*	- or add the OSS, NAS, or CPFS path to the **storage path**.
	//
	// example:
	//
	// 100Gi
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The mount path of the cloud disk.
	//
	// example:
	//
	// /mnt/systemDisk
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The subpath of the cloud disk that is mounted to the instance.
	//
	// example:
	//
	// workspace
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The disk or snapshot usage.
	Status *CreateInstanceRequestCloudDisksStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	// The cloud disk type.
	//
	// 	- rootfs: Mounts the disk as a system disk. The system environment is stored on the disk.
	//
	// example:
	//
	// rootfs
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s CreateInstanceRequestCloudDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestCloudDisks) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCloudDisks) GetCapacity() *string {
	return s.Capacity
}

func (s *CreateInstanceRequestCloudDisks) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateInstanceRequestCloudDisks) GetPath() *string {
	return s.Path
}

func (s *CreateInstanceRequestCloudDisks) GetStatus() *CreateInstanceRequestCloudDisksStatus {
	return s.Status
}

func (s *CreateInstanceRequestCloudDisks) GetSubType() *string {
	return s.SubType
}

func (s *CreateInstanceRequestCloudDisks) SetCapacity(v string) *CreateInstanceRequestCloudDisks {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceRequestCloudDisks) SetMountPath(v string) *CreateInstanceRequestCloudDisks {
	s.MountPath = &v
	return s
}

func (s *CreateInstanceRequestCloudDisks) SetPath(v string) *CreateInstanceRequestCloudDisks {
	s.Path = &v
	return s
}

func (s *CreateInstanceRequestCloudDisks) SetStatus(v *CreateInstanceRequestCloudDisksStatus) *CreateInstanceRequestCloudDisks {
	s.Status = v
	return s
}

func (s *CreateInstanceRequestCloudDisks) SetSubType(v string) *CreateInstanceRequestCloudDisks {
	s.SubType = &v
	return s
}

func (s *CreateInstanceRequestCloudDisks) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestCloudDisksStatus struct {
	// The available capacity. Unit: bytes.
	//
	// example:
	//
	// 31841058816
	Available *int64 `json:"Available,omitempty" xml:"Available,omitempty"`
	// The capacity. Unit: bytes.
	//
	// example:
	//
	// 32212254720
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The used capacity. Unit: bytes.
	//
	// example:
	//
	// 371195904
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s CreateInstanceRequestCloudDisksStatus) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestCloudDisksStatus) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCloudDisksStatus) GetAvailable() *int64 {
	return s.Available
}

func (s *CreateInstanceRequestCloudDisksStatus) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CreateInstanceRequestCloudDisksStatus) GetUsage() *int64 {
	return s.Usage
}

func (s *CreateInstanceRequestCloudDisksStatus) SetAvailable(v int64) *CreateInstanceRequestCloudDisksStatus {
	s.Available = &v
	return s
}

func (s *CreateInstanceRequestCloudDisksStatus) SetCapacity(v int64) *CreateInstanceRequestCloudDisksStatus {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceRequestCloudDisksStatus) SetUsage(v int64) *CreateInstanceRequestCloudDisksStatus {
	s.Usage = &v
	return s
}

func (s *CreateInstanceRequestCloudDisksStatus) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestDatasets struct {
	// The dataset ID. If the dataset is read-only, you cannot change the dataset permission from read-only to read and write by using MountAccess.
	//
	// You can call [ListDatasets](https://help.aliyun.com/document_detail/457222.html) to obtain the dataset ID. If you configure the dataset ID, you cannot configure the dataset URI.
	//
	// example:
	//
	// d-vsqjvsjp4orp5l206u
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset version. You must also configure DatasetId. If you leave this parameter empty, the value v1 is used by default.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// Specifies whether to enable dynamic mounting. Default value: false.
	//
	// 	- Currently, only instances using general-purpose computing resources are supported.
	//
	// 	- Currently, only OSS datasets are supported. The mounted datasets are read-only.
	//
	// 	- The mount path of the dynamically mounted dataset must be a subpath of the root path. Example: /mnt/dynamic/data1/
	//
	// 	- A dynamically mounted dataset must be after non-dynamic datasets.
	//
	// example:
	//
	// true
	Dynamic *bool `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	// The read and write permissions of the dataset. If the dataset is read-only, it cannot be changed to read and write.
	//
	// example:
	//
	// RW
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The mount path of the dataset.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// Deprecated
	//
	// The mount type. You cannot specify Options at the same time. This is deprecated, and you can use Options instead.
	//
	// example:
	//
	// ReadOnly
	OptionType *string `json:"OptionType,omitempty" xml:"OptionType,omitempty"`
	// The custom dataset mount options. Only OSS is supported. You cannot specify OptionType at the same time. For more information, see [DSW mount configurations](https://www.alibabacloud.com/help/en/pai/user-guide/read-and-write-dataset-data).
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
	// The URI of the storage service directory, which can be directly mounted. This parameter is mutually exclusive with DatasetId.
	//
	// URI formats of different types of storage:
	//
	// 	- OSS: oss://bucket-name.oss-cn-shanghai-internal.aliyuncs.com/data/path/
	//
	// 	- NAS: nas://29\\*\\*d-b12\\*\\*\\*\\*446.cn-hangzhou.nas.aliyuncs.com/data/path/
	//
	// 	- Extreme NAS: nas://29\\*\\*\\*\\*123-y\\*\\*r.cn-hangzhou.extreme.nas.aliyuncs.com/data/path/
	//
	// 	- CPFS: cpfs://cpfs-213\\*\\*\\*\\*87.cn-wulanchabu/ptc-292\\*\\*\\*\\*\\*cbb/exp-290\\*\\*\\*\\*\\*\\*\\*\\*03e/data/path/
	//
	// 	- Lingjun CPFS: bmcpfs://cpfs-290\\*\\*\\*\\*\\*\\*foflh-vpc-x\\*\\*\\*\\*8r.cn-wulanchabu.cpfs.aliyuncs.com/data/path/
	//
	// example:
	//
	// oss://bucket-name.oss-cn-shanghai-internal.aliyuncs.com/data/path/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateInstanceRequestDatasets) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestDatasets) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestDatasets) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateInstanceRequestDatasets) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *CreateInstanceRequestDatasets) GetDynamic() *bool {
	return s.Dynamic
}

func (s *CreateInstanceRequestDatasets) GetMountAccess() *string {
	return s.MountAccess
}

func (s *CreateInstanceRequestDatasets) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateInstanceRequestDatasets) GetOptionType() *string {
	return s.OptionType
}

func (s *CreateInstanceRequestDatasets) GetOptions() *string {
	return s.Options
}

func (s *CreateInstanceRequestDatasets) GetUri() *string {
	return s.Uri
}

func (s *CreateInstanceRequestDatasets) SetDatasetId(v string) *CreateInstanceRequestDatasets {
	s.DatasetId = &v
	return s
}

func (s *CreateInstanceRequestDatasets) SetDatasetVersion(v string) *CreateInstanceRequestDatasets {
	s.DatasetVersion = &v
	return s
}

func (s *CreateInstanceRequestDatasets) SetDynamic(v bool) *CreateInstanceRequestDatasets {
	s.Dynamic = &v
	return s
}

func (s *CreateInstanceRequestDatasets) SetMountAccess(v string) *CreateInstanceRequestDatasets {
	s.MountAccess = &v
	return s
}

func (s *CreateInstanceRequestDatasets) SetMountPath(v string) *CreateInstanceRequestDatasets {
	s.MountPath = &v
	return s
}

func (s *CreateInstanceRequestDatasets) SetOptionType(v string) *CreateInstanceRequestDatasets {
	s.OptionType = &v
	return s
}

func (s *CreateInstanceRequestDatasets) SetOptions(v string) *CreateInstanceRequestDatasets {
	s.Options = &v
	return s
}

func (s *CreateInstanceRequestDatasets) SetUri(v string) *CreateInstanceRequestDatasets {
	s.Uri = &v
	return s
}

func (s *CreateInstanceRequestDatasets) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestLabels struct {
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

func (s CreateInstanceRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestLabels) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestLabels) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestLabels) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestLabels) SetKey(v string) *CreateInstanceRequestLabels {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestLabels) SetValue(v string) *CreateInstanceRequestLabels {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestLabels) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestRequestedResource struct {
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
	// The GPU memory type. Valid values:
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
	// The size of the shared memory. Unit: GB.
	//
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s CreateInstanceRequestRequestedResource) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestRequestedResource) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestRequestedResource) GetCPU() *string {
	return s.CPU
}

func (s *CreateInstanceRequestRequestedResource) GetGPU() *string {
	return s.GPU
}

func (s *CreateInstanceRequestRequestedResource) GetGPUType() *string {
	return s.GPUType
}

func (s *CreateInstanceRequestRequestedResource) GetMemory() *string {
	return s.Memory
}

func (s *CreateInstanceRequestRequestedResource) GetSharedMemory() *string {
	return s.SharedMemory
}

func (s *CreateInstanceRequestRequestedResource) SetCPU(v string) *CreateInstanceRequestRequestedResource {
	s.CPU = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) SetGPU(v string) *CreateInstanceRequestRequestedResource {
	s.GPU = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) SetGPUType(v string) *CreateInstanceRequestRequestedResource {
	s.GPUType = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) SetMemory(v string) *CreateInstanceRequestRequestedResource {
	s.Memory = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) SetSharedMemory(v string) *CreateInstanceRequestRequestedResource {
	s.SharedMemory = &v
	return s
}

func (s *CreateInstanceRequestRequestedResource) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestSpotSpec struct {
	// example:
	//
	// 0.1
	SpotDiscountLimit *string `json:"SpotDiscountLimit,omitempty" xml:"SpotDiscountLimit,omitempty"`
	// example:
	//
	// 0
	SpotDuration *string `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// example:
	//
	// 0.12
	SpotPriceLimit *string `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s CreateInstanceRequestSpotSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestSpotSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestSpotSpec) GetSpotDiscountLimit() *string {
	return s.SpotDiscountLimit
}

func (s *CreateInstanceRequestSpotSpec) GetSpotDuration() *string {
	return s.SpotDuration
}

func (s *CreateInstanceRequestSpotSpec) GetSpotPriceLimit() *string {
	return s.SpotPriceLimit
}

func (s *CreateInstanceRequestSpotSpec) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateInstanceRequestSpotSpec) SetSpotDiscountLimit(v string) *CreateInstanceRequestSpotSpec {
	s.SpotDiscountLimit = &v
	return s
}

func (s *CreateInstanceRequestSpotSpec) SetSpotDuration(v string) *CreateInstanceRequestSpotSpec {
	s.SpotDuration = &v
	return s
}

func (s *CreateInstanceRequestSpotSpec) SetSpotPriceLimit(v string) *CreateInstanceRequestSpotSpec {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateInstanceRequestSpotSpec) SetSpotStrategy(v string) *CreateInstanceRequestSpotSpec {
	s.SpotStrategy = &v
	return s
}

func (s *CreateInstanceRequestSpotSpec) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTag) SetKey(v string) *CreateInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTag) SetValue(v string) *CreateInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestUserCommand struct {
	OnStart *CreateInstanceRequestUserCommandOnStart `json:"OnStart,omitempty" xml:"OnStart,omitempty" type:"Struct"`
}

func (s CreateInstanceRequestUserCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestUserCommand) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestUserCommand) GetOnStart() *CreateInstanceRequestUserCommandOnStart {
	return s.OnStart
}

func (s *CreateInstanceRequestUserCommand) SetOnStart(v *CreateInstanceRequestUserCommandOnStart) *CreateInstanceRequestUserCommand {
	s.OnStart = v
	return s
}

func (s *CreateInstanceRequestUserCommand) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestUserCommandOnStart struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateInstanceRequestUserCommandOnStart) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestUserCommandOnStart) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestUserCommandOnStart) GetContent() *string {
	return s.Content
}

func (s *CreateInstanceRequestUserCommandOnStart) SetContent(v string) *CreateInstanceRequestUserCommandOnStart {
	s.Content = &v
	return s
}

func (s *CreateInstanceRequestUserCommandOnStart) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestUserVpc struct {
	BandwidthLimit *BandwidthLimit `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The default route. Valid values:
	//
	// 	- eth0: The default network interface is used to access the Internet through the public gateway.
	//
	// 	- eth1: The user\\"s elastic network interface (ENI) is used to access the Internet through the private gateway. For more information about the configuration method, see [Enable Internet access for a DSW instance by using a private Internet NAT gateway](https://help.aliyun.com/document_detail/2525343.html).
	//
	// example:
	//
	// eth0
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR blocks.
	//
	// 	- If you leave the SwitchId and ExtendedCIDRs parameters empty, the system automatically obtains all CIDR blocks in a VPC.
	//
	// 	- If you configure the SwitchId and ExtendedCIDRs parameters, we recommend that you specify all CIDR blocks in a VPC.
	//
	// example:
	//
	// ["192.168.0.1/24", "192.168.1.1/24"]
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The forward information.
	ForwardInfos []*ForwardInfo `json:"ForwardInfos,omitempty" xml:"ForwardInfos,omitempty" type:"Repeated"`
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

func (s CreateInstanceRequestUserVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestUserVpc) GetBandwidthLimit() *BandwidthLimit {
	return s.BandwidthLimit
}

func (s *CreateInstanceRequestUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *CreateInstanceRequestUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *CreateInstanceRequestUserVpc) GetForwardInfos() []*ForwardInfo {
	return s.ForwardInfos
}

func (s *CreateInstanceRequestUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateInstanceRequestUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequestUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequestUserVpc) SetBandwidthLimit(v *BandwidthLimit) *CreateInstanceRequestUserVpc {
	s.BandwidthLimit = v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetDefaultRoute(v string) *CreateInstanceRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetExtendedCIDRs(v []*string) *CreateInstanceRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetForwardInfos(v []*ForwardInfo) *CreateInstanceRequestUserVpc {
	s.ForwardInfos = v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetSecurityGroupId(v string) *CreateInstanceRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetVSwitchId(v string) *CreateInstanceRequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequestUserVpc) SetVpcId(v string) *CreateInstanceRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequestUserVpc) Validate() error {
	return dara.Validate(s)
}
