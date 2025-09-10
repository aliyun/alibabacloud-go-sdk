// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateInstanceRequest
	GetAccessibility() *string
	SetAffinity(v *UpdateInstanceRequestAffinity) *UpdateInstanceRequest
	GetAffinity() *UpdateInstanceRequestAffinity
	SetAssignNodeSpec(v *UpdateInstanceRequestAssignNodeSpec) *UpdateInstanceRequest
	GetAssignNodeSpec() *UpdateInstanceRequestAssignNodeSpec
	SetCloudDisks(v []*UpdateInstanceRequestCloudDisks) *UpdateInstanceRequest
	GetCloudDisks() []*UpdateInstanceRequestCloudDisks
	SetCredentialConfig(v *CredentialConfig) *UpdateInstanceRequest
	GetCredentialConfig() *CredentialConfig
	SetDatasets(v []*UpdateInstanceRequestDatasets) *UpdateInstanceRequest
	GetDatasets() []*UpdateInstanceRequestDatasets
	SetDisassociateAssignNode(v bool) *UpdateInstanceRequest
	GetDisassociateAssignNode() *bool
	SetDisassociateCredential(v bool) *UpdateInstanceRequest
	GetDisassociateCredential() *bool
	SetDisassociateDatasets(v bool) *UpdateInstanceRequest
	GetDisassociateDatasets() *bool
	SetDisassociateDriver(v bool) *UpdateInstanceRequest
	GetDisassociateDriver() *bool
	SetDisassociateEnvironmentVariables(v bool) *UpdateInstanceRequest
	GetDisassociateEnvironmentVariables() *bool
	SetDisassociateForwardInfos(v bool) *UpdateInstanceRequest
	GetDisassociateForwardInfos() *bool
	SetDisassociateSpot(v bool) *UpdateInstanceRequest
	GetDisassociateSpot() *bool
	SetDisassociateUserCommand(v bool) *UpdateInstanceRequest
	GetDisassociateUserCommand() *bool
	SetDisassociateVpc(v bool) *UpdateInstanceRequest
	GetDisassociateVpc() *bool
	SetDriver(v string) *UpdateInstanceRequest
	GetDriver() *string
	SetDynamicMount(v *DynamicMount) *UpdateInstanceRequest
	GetDynamicMount() *DynamicMount
	SetEcsSpec(v string) *UpdateInstanceRequest
	GetEcsSpec() *string
	SetEnvironmentVariables(v map[string]interface{}) *UpdateInstanceRequest
	GetEnvironmentVariables() map[string]interface{}
	SetImageAuth(v string) *UpdateInstanceRequest
	GetImageAuth() *string
	SetImageId(v string) *UpdateInstanceRequest
	GetImageId() *string
	SetImageUrl(v string) *UpdateInstanceRequest
	GetImageUrl() *string
	SetInstanceName(v string) *UpdateInstanceRequest
	GetInstanceName() *string
	SetOversoldType(v string) *UpdateInstanceRequest
	GetOversoldType() *string
	SetPriority(v int64) *UpdateInstanceRequest
	GetPriority() *int64
	SetRequestedResource(v *UpdateInstanceRequestRequestedResource) *UpdateInstanceRequest
	GetRequestedResource() *UpdateInstanceRequestRequestedResource
	SetSpotSpec(v *UpdateInstanceRequestSpotSpec) *UpdateInstanceRequest
	GetSpotSpec() *UpdateInstanceRequestSpotSpec
	SetUserCommand(v *UpdateInstanceRequestUserCommand) *UpdateInstanceRequest
	GetUserCommand() *UpdateInstanceRequestUserCommand
	SetUserId(v string) *UpdateInstanceRequest
	GetUserId() *string
	SetUserVpc(v *UpdateInstanceRequestUserVpc) *UpdateInstanceRequest
	GetUserVpc() *UpdateInstanceRequestUserVpc
	SetWorkspaceSource(v string) *UpdateInstanceRequest
	GetWorkspaceSource() *string
}

type UpdateInstanceRequest struct {
	// The visibility of the instance.
	//
	// Valid values:
	//
	// 	- PUBLIC: Accessible to all members in the workspace.
	//
	// 	- PRIVATE: Accessible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The affinity configuration.
	Affinity       *UpdateInstanceRequestAffinity       `json:"Affinity,omitempty" xml:"Affinity,omitempty" type:"Struct"`
	AssignNodeSpec *UpdateInstanceRequestAssignNodeSpec `json:"AssignNodeSpec,omitempty" xml:"AssignNodeSpec,omitempty" type:"Struct"`
	// The cloud disks.
	//
	// example:
	//
	// []
	CloudDisks []*UpdateInstanceRequestCloudDisks `json:"CloudDisks,omitempty" xml:"CloudDisks,omitempty" type:"Repeated"`
	// The credential configuration.
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The datasets.
	Datasets               []*UpdateInstanceRequestDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	DisassociateAssignNode *bool                            `json:"DisassociateAssignNode,omitempty" xml:"DisassociateAssignNode,omitempty"`
	// Specifies whether to delete the credential injection information.
	//
	// example:
	//
	// false
	DisassociateCredential *bool `json:"DisassociateCredential,omitempty" xml:"DisassociateCredential,omitempty"`
	// Specifies whether to delete the associated datasets.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DisassociateDatasets *bool `json:"DisassociateDatasets,omitempty" xml:"DisassociateDatasets,omitempty"`
	// Specifies whether to delete the NVIDIA driver configuration.
	//
	// example:
	//
	// false
	DisassociateDriver               *bool `json:"DisassociateDriver,omitempty" xml:"DisassociateDriver,omitempty"`
	DisassociateEnvironmentVariables *bool `json:"DisassociateEnvironmentVariables,omitempty" xml:"DisassociateEnvironmentVariables,omitempty"`
	// Specifies whether to delete the associated forward information.
	//
	// example:
	//
	// false
	DisassociateForwardInfos *bool `json:"DisassociateForwardInfos,omitempty" xml:"DisassociateForwardInfos,omitempty"`
	DisassociateSpot         *bool `json:"DisassociateSpot,omitempty" xml:"DisassociateSpot,omitempty"`
	DisassociateUserCommand  *bool `json:"DisassociateUserCommand,omitempty" xml:"DisassociateUserCommand,omitempty"`
	// Specifies whether to delete the associated user VPC.
	//
	// example:
	//
	// false
	DisassociateVpc *bool `json:"DisassociateVpc,omitempty" xml:"DisassociateVpc,omitempty"`
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
	EcsSpec              *string                `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	EnvironmentVariables map[string]interface{} `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The Base64-encoded account and password for the user‘s private image. The password will be hidden.
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
	// The instance name. Format requirements:
	//
	// 	- The name can contain only letters, digits, and underscores (_).
	//
	// 	- The name can be up to 27 characters in length.
	//
	// example:
	//
	// training_data
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OversoldType *string `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
	// The priority based on which resources are allocated to instances. Valid values: 1 to 9.
	//
	// 	- 1: the lowest priority.
	//
	// 	- 9 is the highest priority.
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
	RequestedResource *UpdateInstanceRequestRequestedResource `json:"RequestedResource,omitempty" xml:"RequestedResource,omitempty" type:"Struct"`
	SpotSpec          *UpdateInstanceRequestSpotSpec          `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty" type:"Struct"`
	UserCommand       *UpdateInstanceRequestUserCommand       `json:"UserCommand,omitempty" xml:"UserCommand,omitempty" type:"Struct"`
	// the User ID of the instance.
	//
	// example:
	//
	// 16122**********
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The virtual private cloud (VPC) configurations.
	UserVpc *UpdateInstanceRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// Specifies the storage corresponding to the working directory. You can mount disks or datasets to /mnt/workspace at the same time. OSS datasets and dynamically mounted datasets are not supported.
	//
	// Valid values:
	//
	// 	- rootfsCloudDisk: Mount disk to the working directory.
	//
	// 	- Mount path of the dataset, such as /mnt/data: Datasets in URI format only support this method.
	//
	// 	- Dataset ID, such as d-vsqjvs\\*\\*\\*\\*rp5l206u: If a single dataset is mounted to multiple paths, the first path is selected. We recommend that you do not use this method, use the mount path instead.
	//
	// If you leave this parameter empty:
	//
	// 	- If the instance uses cloud disks, cloud disks are selected by default.
	//
	// 	- if no disks are available, the first NAS or CPFS dataset is selected as the working directory.
	//
	// 	- If no disk, NAS, or CPFS datasets is available, the host space is used.
	//
	// example:
	//
	// /mnt/data
	WorkspaceSource *string `json:"WorkspaceSource,omitempty" xml:"WorkspaceSource,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateInstanceRequest) GetAffinity() *UpdateInstanceRequestAffinity {
	return s.Affinity
}

func (s *UpdateInstanceRequest) GetAssignNodeSpec() *UpdateInstanceRequestAssignNodeSpec {
	return s.AssignNodeSpec
}

func (s *UpdateInstanceRequest) GetCloudDisks() []*UpdateInstanceRequestCloudDisks {
	return s.CloudDisks
}

func (s *UpdateInstanceRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *UpdateInstanceRequest) GetDatasets() []*UpdateInstanceRequestDatasets {
	return s.Datasets
}

func (s *UpdateInstanceRequest) GetDisassociateAssignNode() *bool {
	return s.DisassociateAssignNode
}

func (s *UpdateInstanceRequest) GetDisassociateCredential() *bool {
	return s.DisassociateCredential
}

func (s *UpdateInstanceRequest) GetDisassociateDatasets() *bool {
	return s.DisassociateDatasets
}

func (s *UpdateInstanceRequest) GetDisassociateDriver() *bool {
	return s.DisassociateDriver
}

func (s *UpdateInstanceRequest) GetDisassociateEnvironmentVariables() *bool {
	return s.DisassociateEnvironmentVariables
}

func (s *UpdateInstanceRequest) GetDisassociateForwardInfos() *bool {
	return s.DisassociateForwardInfos
}

func (s *UpdateInstanceRequest) GetDisassociateSpot() *bool {
	return s.DisassociateSpot
}

func (s *UpdateInstanceRequest) GetDisassociateUserCommand() *bool {
	return s.DisassociateUserCommand
}

func (s *UpdateInstanceRequest) GetDisassociateVpc() *bool {
	return s.DisassociateVpc
}

func (s *UpdateInstanceRequest) GetDriver() *string {
	return s.Driver
}

func (s *UpdateInstanceRequest) GetDynamicMount() *DynamicMount {
	return s.DynamicMount
}

func (s *UpdateInstanceRequest) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *UpdateInstanceRequest) GetEnvironmentVariables() map[string]interface{} {
	return s.EnvironmentVariables
}

func (s *UpdateInstanceRequest) GetImageAuth() *string {
	return s.ImageAuth
}

func (s *UpdateInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateInstanceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *UpdateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceRequest) GetOversoldType() *string {
	return s.OversoldType
}

func (s *UpdateInstanceRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *UpdateInstanceRequest) GetRequestedResource() *UpdateInstanceRequestRequestedResource {
	return s.RequestedResource
}

func (s *UpdateInstanceRequest) GetSpotSpec() *UpdateInstanceRequestSpotSpec {
	return s.SpotSpec
}

func (s *UpdateInstanceRequest) GetUserCommand() *UpdateInstanceRequestUserCommand {
	return s.UserCommand
}

func (s *UpdateInstanceRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateInstanceRequest) GetUserVpc() *UpdateInstanceRequestUserVpc {
	return s.UserVpc
}

func (s *UpdateInstanceRequest) GetWorkspaceSource() *string {
	return s.WorkspaceSource
}

func (s *UpdateInstanceRequest) SetAccessibility(v string) *UpdateInstanceRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateInstanceRequest) SetAffinity(v *UpdateInstanceRequestAffinity) *UpdateInstanceRequest {
	s.Affinity = v
	return s
}

func (s *UpdateInstanceRequest) SetAssignNodeSpec(v *UpdateInstanceRequestAssignNodeSpec) *UpdateInstanceRequest {
	s.AssignNodeSpec = v
	return s
}

func (s *UpdateInstanceRequest) SetCloudDisks(v []*UpdateInstanceRequestCloudDisks) *UpdateInstanceRequest {
	s.CloudDisks = v
	return s
}

func (s *UpdateInstanceRequest) SetCredentialConfig(v *CredentialConfig) *UpdateInstanceRequest {
	s.CredentialConfig = v
	return s
}

func (s *UpdateInstanceRequest) SetDatasets(v []*UpdateInstanceRequestDatasets) *UpdateInstanceRequest {
	s.Datasets = v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateAssignNode(v bool) *UpdateInstanceRequest {
	s.DisassociateAssignNode = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateCredential(v bool) *UpdateInstanceRequest {
	s.DisassociateCredential = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateDatasets(v bool) *UpdateInstanceRequest {
	s.DisassociateDatasets = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateDriver(v bool) *UpdateInstanceRequest {
	s.DisassociateDriver = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateEnvironmentVariables(v bool) *UpdateInstanceRequest {
	s.DisassociateEnvironmentVariables = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateForwardInfos(v bool) *UpdateInstanceRequest {
	s.DisassociateForwardInfos = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateSpot(v bool) *UpdateInstanceRequest {
	s.DisassociateSpot = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateUserCommand(v bool) *UpdateInstanceRequest {
	s.DisassociateUserCommand = &v
	return s
}

func (s *UpdateInstanceRequest) SetDisassociateVpc(v bool) *UpdateInstanceRequest {
	s.DisassociateVpc = &v
	return s
}

func (s *UpdateInstanceRequest) SetDriver(v string) *UpdateInstanceRequest {
	s.Driver = &v
	return s
}

func (s *UpdateInstanceRequest) SetDynamicMount(v *DynamicMount) *UpdateInstanceRequest {
	s.DynamicMount = v
	return s
}

func (s *UpdateInstanceRequest) SetEcsSpec(v string) *UpdateInstanceRequest {
	s.EcsSpec = &v
	return s
}

func (s *UpdateInstanceRequest) SetEnvironmentVariables(v map[string]interface{}) *UpdateInstanceRequest {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateInstanceRequest) SetImageAuth(v string) *UpdateInstanceRequest {
	s.ImageAuth = &v
	return s
}

func (s *UpdateInstanceRequest) SetImageId(v string) *UpdateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceRequest) SetImageUrl(v string) *UpdateInstanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetOversoldType(v string) *UpdateInstanceRequest {
	s.OversoldType = &v
	return s
}

func (s *UpdateInstanceRequest) SetPriority(v int64) *UpdateInstanceRequest {
	s.Priority = &v
	return s
}

func (s *UpdateInstanceRequest) SetRequestedResource(v *UpdateInstanceRequestRequestedResource) *UpdateInstanceRequest {
	s.RequestedResource = v
	return s
}

func (s *UpdateInstanceRequest) SetSpotSpec(v *UpdateInstanceRequestSpotSpec) *UpdateInstanceRequest {
	s.SpotSpec = v
	return s
}

func (s *UpdateInstanceRequest) SetUserCommand(v *UpdateInstanceRequestUserCommand) *UpdateInstanceRequest {
	s.UserCommand = v
	return s
}

func (s *UpdateInstanceRequest) SetUserId(v string) *UpdateInstanceRequest {
	s.UserId = &v
	return s
}

func (s *UpdateInstanceRequest) SetUserVpc(v *UpdateInstanceRequestUserVpc) *UpdateInstanceRequest {
	s.UserVpc = v
	return s
}

func (s *UpdateInstanceRequest) SetWorkspaceSource(v string) *UpdateInstanceRequest {
	s.WorkspaceSource = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestAffinity struct {
	// The CPU affinity configuration. Only subscription instances that use general-purpose computing resources support CPU affinity configuration.
	CPU *UpdateInstanceRequestAffinityCPU `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
}

func (s UpdateInstanceRequestAffinity) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestAffinity) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestAffinity) GetCPU() *UpdateInstanceRequestAffinityCPU {
	return s.CPU
}

func (s *UpdateInstanceRequestAffinity) SetCPU(v *UpdateInstanceRequestAffinityCPU) *UpdateInstanceRequestAffinity {
	s.CPU = v
	return s
}

func (s *UpdateInstanceRequestAffinity) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestAffinityCPU struct {
	// Specifies whether CPU affinity is enabled.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s UpdateInstanceRequestAffinityCPU) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestAffinityCPU) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestAffinityCPU) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateInstanceRequestAffinityCPU) SetEnable(v bool) *UpdateInstanceRequestAffinityCPU {
	s.Enable = &v
	return s
}

func (s *UpdateInstanceRequestAffinityCPU) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestAssignNodeSpec struct {
	// example:
	//
	// node-b
	AntiAffinityNodeNames *string `json:"AntiAffinityNodeNames,omitempty" xml:"AntiAffinityNodeNames,omitempty"`
	// example:
	//
	// node-a
	NodeNames *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
}

func (s UpdateInstanceRequestAssignNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestAssignNodeSpec) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestAssignNodeSpec) GetAntiAffinityNodeNames() *string {
	return s.AntiAffinityNodeNames
}

func (s *UpdateInstanceRequestAssignNodeSpec) GetNodeNames() *string {
	return s.NodeNames
}

func (s *UpdateInstanceRequestAssignNodeSpec) SetAntiAffinityNodeNames(v string) *UpdateInstanceRequestAssignNodeSpec {
	s.AntiAffinityNodeNames = &v
	return s
}

func (s *UpdateInstanceRequestAssignNodeSpec) SetNodeNames(v string) *UpdateInstanceRequestAssignNodeSpec {
	s.NodeNames = &v
	return s
}

func (s *UpdateInstanceRequestAssignNodeSpec) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestCloudDisks struct {
	// If **Resource Type*	- is **Public Resource*	- or if **Resource Quota*	- is subscription-based general-purpose computing resources (CPU cores ≥ 2 and memory ≥ 4 GB, or configured with GPU):
	//
	// Each instance has a free system disk quota of 100 GiB for persistent storage. **If the DSW instance is stopped and not launched for more than 15 days, the disk is cleared**. The disk can be expanded. For specific pricing, refer to the console.
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
	// Disk type:
	//
	// 	- rootfs: Mounts the disk as a system disk. The system environment is stored on the disk.
	//
	// example:
	//
	// rootfs
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s UpdateInstanceRequestCloudDisks) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestCloudDisks) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestCloudDisks) GetCapacity() *string {
	return s.Capacity
}

func (s *UpdateInstanceRequestCloudDisks) GetSubType() *string {
	return s.SubType
}

func (s *UpdateInstanceRequestCloudDisks) SetCapacity(v string) *UpdateInstanceRequestCloudDisks {
	s.Capacity = &v
	return s
}

func (s *UpdateInstanceRequestCloudDisks) SetSubType(v string) *UpdateInstanceRequestCloudDisks {
	s.SubType = &v
	return s
}

func (s *UpdateInstanceRequestCloudDisks) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestDatasets struct {
	// The dataset ID. If the dataset is read-only, you cannot change the dataset pemission from read-only to read and write by using MountAccess.
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
	// Specifies whether dynamic mounting is enabled. Default value: false.
	//
	// 	- Currently, only instances using general-purpose computing resources are supported.
	//
	// 	- Currently, only OSS datasets are supported. The mounted datasets are read-only.
	//
	// 	- The MountPath of the dynamically mounted dataset must be a subpath of the root path. Example: /mnt/dynamic/data1/
	//
	// 	- A dynamically mounted dataset must be after non-dynamic datasets.
	//
	// example:
	//
	// false
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
	// The mount type. You cannot specify Options at the same time. This is deprecated, you can use Options instead.
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

func (s UpdateInstanceRequestDatasets) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestDatasets) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestDatasets) GetDatasetId() *string {
	return s.DatasetId
}

func (s *UpdateInstanceRequestDatasets) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *UpdateInstanceRequestDatasets) GetDynamic() *bool {
	return s.Dynamic
}

func (s *UpdateInstanceRequestDatasets) GetMountAccess() *string {
	return s.MountAccess
}

func (s *UpdateInstanceRequestDatasets) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateInstanceRequestDatasets) GetOptionType() *string {
	return s.OptionType
}

func (s *UpdateInstanceRequestDatasets) GetOptions() *string {
	return s.Options
}

func (s *UpdateInstanceRequestDatasets) GetUri() *string {
	return s.Uri
}

func (s *UpdateInstanceRequestDatasets) SetDatasetId(v string) *UpdateInstanceRequestDatasets {
	s.DatasetId = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) SetDatasetVersion(v string) *UpdateInstanceRequestDatasets {
	s.DatasetVersion = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) SetDynamic(v bool) *UpdateInstanceRequestDatasets {
	s.Dynamic = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) SetMountAccess(v string) *UpdateInstanceRequestDatasets {
	s.MountAccess = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) SetMountPath(v string) *UpdateInstanceRequestDatasets {
	s.MountPath = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) SetOptionType(v string) *UpdateInstanceRequestDatasets {
	s.OptionType = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) SetOptions(v string) *UpdateInstanceRequestDatasets {
	s.Options = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) SetUri(v string) *UpdateInstanceRequestDatasets {
	s.Uri = &v
	return s
}

func (s *UpdateInstanceRequestDatasets) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestRequestedResource struct {
	// The number of vCPU cores.
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
	// The GPU type.
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

func (s UpdateInstanceRequestRequestedResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestRequestedResource) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestRequestedResource) GetCPU() *string {
	return s.CPU
}

func (s *UpdateInstanceRequestRequestedResource) GetGPU() *string {
	return s.GPU
}

func (s *UpdateInstanceRequestRequestedResource) GetGPUType() *string {
	return s.GPUType
}

func (s *UpdateInstanceRequestRequestedResource) GetMemory() *string {
	return s.Memory
}

func (s *UpdateInstanceRequestRequestedResource) GetSharedMemory() *string {
	return s.SharedMemory
}

func (s *UpdateInstanceRequestRequestedResource) SetCPU(v string) *UpdateInstanceRequestRequestedResource {
	s.CPU = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) SetGPU(v string) *UpdateInstanceRequestRequestedResource {
	s.GPU = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) SetGPUType(v string) *UpdateInstanceRequestRequestedResource {
	s.GPUType = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) SetMemory(v string) *UpdateInstanceRequestRequestedResource {
	s.Memory = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) SetSharedMemory(v string) *UpdateInstanceRequestRequestedResource {
	s.SharedMemory = &v
	return s
}

func (s *UpdateInstanceRequestRequestedResource) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestSpotSpec struct {
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

func (s UpdateInstanceRequestSpotSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestSpotSpec) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestSpotSpec) GetSpotDiscountLimit() *string {
	return s.SpotDiscountLimit
}

func (s *UpdateInstanceRequestSpotSpec) GetSpotDuration() *string {
	return s.SpotDuration
}

func (s *UpdateInstanceRequestSpotSpec) GetSpotPriceLimit() *string {
	return s.SpotPriceLimit
}

func (s *UpdateInstanceRequestSpotSpec) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *UpdateInstanceRequestSpotSpec) SetSpotDiscountLimit(v string) *UpdateInstanceRequestSpotSpec {
	s.SpotDiscountLimit = &v
	return s
}

func (s *UpdateInstanceRequestSpotSpec) SetSpotDuration(v string) *UpdateInstanceRequestSpotSpec {
	s.SpotDuration = &v
	return s
}

func (s *UpdateInstanceRequestSpotSpec) SetSpotPriceLimit(v string) *UpdateInstanceRequestSpotSpec {
	s.SpotPriceLimit = &v
	return s
}

func (s *UpdateInstanceRequestSpotSpec) SetSpotStrategy(v string) *UpdateInstanceRequestSpotSpec {
	s.SpotStrategy = &v
	return s
}

func (s *UpdateInstanceRequestSpotSpec) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestUserCommand struct {
	OnStart *UpdateInstanceRequestUserCommandOnStart `json:"OnStart,omitempty" xml:"OnStart,omitempty" type:"Struct"`
}

func (s UpdateInstanceRequestUserCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestUserCommand) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestUserCommand) GetOnStart() *UpdateInstanceRequestUserCommandOnStart {
	return s.OnStart
}

func (s *UpdateInstanceRequestUserCommand) SetOnStart(v *UpdateInstanceRequestUserCommandOnStart) *UpdateInstanceRequestUserCommand {
	s.OnStart = v
	return s
}

func (s *UpdateInstanceRequestUserCommand) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestUserCommandOnStart struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s UpdateInstanceRequestUserCommandOnStart) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestUserCommandOnStart) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestUserCommandOnStart) GetContent() *string {
	return s.Content
}

func (s *UpdateInstanceRequestUserCommandOnStart) SetContent(v string) *UpdateInstanceRequestUserCommandOnStart {
	s.Content = &v
	return s
}

func (s *UpdateInstanceRequestUserCommandOnStart) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestUserVpc struct {
	BandwidthLimit *BandwidthLimit `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The default route. Valid values:
	//
	// 	- eth0: The default network interface is used to access the Internet through the public gateway.
	//
	// 	- eth1: The user\\"s Elastic Network Interface is used to access the Internet through the private gateway.
	//
	// example:
	//
	// eth0
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR blocks.
	//
	// 	- If you leave VSwitchId empty, this parameter is not required and the system automatically obtains all CIDR blocks in the VPC.
	//
	// 	- If VSwitchId is not empty, this parameter is required. Specify all CIDR blocks in the VPC.
	//
	// example:
	//
	// ["192.168.0.1/24", "192.168.1.1/24"]
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The forward configuration of the instance.
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

func (s UpdateInstanceRequestUserVpc) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestUserVpc) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestUserVpc) GetBandwidthLimit() *BandwidthLimit {
	return s.BandwidthLimit
}

func (s *UpdateInstanceRequestUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *UpdateInstanceRequestUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *UpdateInstanceRequestUserVpc) GetForwardInfos() []*ForwardInfo {
	return s.ForwardInfos
}

func (s *UpdateInstanceRequestUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateInstanceRequestUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateInstanceRequestUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateInstanceRequestUserVpc) SetBandwidthLimit(v *BandwidthLimit) *UpdateInstanceRequestUserVpc {
	s.BandwidthLimit = v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetDefaultRoute(v string) *UpdateInstanceRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetExtendedCIDRs(v []*string) *UpdateInstanceRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetForwardInfos(v []*ForwardInfo) *UpdateInstanceRequestUserVpc {
	s.ForwardInfos = v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetSecurityGroupId(v string) *UpdateInstanceRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetVSwitchId(v string) *UpdateInstanceRequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *UpdateInstanceRequestUserVpc) SetVpcId(v string) *UpdateInstanceRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *UpdateInstanceRequestUserVpc) Validate() error {
	return dara.Validate(s)
}
