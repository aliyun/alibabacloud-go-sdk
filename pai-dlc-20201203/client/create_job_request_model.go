// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateJobRequest
	GetAccessibility() *string
	SetCodeSource(v *CreateJobRequestCodeSource) *CreateJobRequest
	GetCodeSource() *CreateJobRequestCodeSource
	SetCredentialConfig(v *CredentialConfig) *CreateJobRequest
	GetCredentialConfig() *CredentialConfig
	SetCustomEnvs(v []*CreateJobRequestCustomEnvs) *CreateJobRequest
	GetCustomEnvs() []*CreateJobRequestCustomEnvs
	SetDataSources(v []*CreateJobRequestDataSources) *CreateJobRequest
	GetDataSources() []*CreateJobRequestDataSources
	SetDebuggerConfigContent(v string) *CreateJobRequest
	GetDebuggerConfigContent() *string
	SetDisplayName(v string) *CreateJobRequest
	GetDisplayName() *string
	SetElasticSpec(v *JobElasticSpec) *CreateJobRequest
	GetElasticSpec() *JobElasticSpec
	SetEnvs(v map[string]*string) *CreateJobRequest
	GetEnvs() map[string]*string
	SetJobMaxRunningTimeMinutes(v int64) *CreateJobRequest
	GetJobMaxRunningTimeMinutes() *int64
	SetJobSpecs(v []*JobSpec) *CreateJobRequest
	GetJobSpecs() []*JobSpec
	SetJobType(v string) *CreateJobRequest
	GetJobType() *string
	SetOptions(v string) *CreateJobRequest
	GetOptions() *string
	SetPriority(v int32) *CreateJobRequest
	GetPriority() *int32
	SetResourceId(v string) *CreateJobRequest
	GetResourceId() *string
	SetSettings(v *JobSettings) *CreateJobRequest
	GetSettings() *JobSettings
	SetSuccessPolicy(v string) *CreateJobRequest
	GetSuccessPolicy() *string
	SetTemplateId(v string) *CreateJobRequest
	GetTemplateId() *string
	SetTemplateVersion(v int32) *CreateJobRequest
	GetTemplateVersion() *int32
	SetThirdpartyLibDir(v string) *CreateJobRequest
	GetThirdpartyLibDir() *string
	SetThirdpartyLibs(v []*string) *CreateJobRequest
	GetThirdpartyLibs() []*string
	SetUserCommand(v string) *CreateJobRequest
	GetUserCommand() *string
	SetUserVpc(v *CreateJobRequestUserVpc) *CreateJobRequest
	GetUserVpc() *CreateJobRequestUserVpc
	SetWorkspaceId(v string) *CreateJobRequest
	GetWorkspaceId() *string
}

type CreateJobRequest struct {
	// The job visibility. Valid values:
	//
	// 	- PUBLIC: The job is visible to all members in the workspace.
	//
	// 	- PRIVATE: The job is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The code source of the job. Before the node of the job runs, DLC automatically downloads the configured code from the code source and mounts the code to the local path of the container.
	CodeSource *CreateJobRequestCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// The access credential configuration.
	CredentialConfig *CredentialConfig             `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	CustomEnvs       []*CreateJobRequestCustomEnvs `json:"CustomEnvs,omitempty" xml:"CustomEnvs,omitempty" type:"Repeated"`
	// The data sources for job running.
	DataSources []*CreateJobRequestDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// This parameter is not supported.
	//
	// example:
	//
	// “”
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	// The job name. The name must be in the following format:
	//
	// 	- The name must be 1 to 256 characters in length.
	//
	// 	- The name can contain digits, letters, underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is not supported.
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// The environment variables.
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The maximum running duration of the job. Unit: minutes.
	//
	// example:
	//
	// 1024
	JobMaxRunningTimeMinutes *int64 `json:"JobMaxRunningTimeMinutes,omitempty" xml:"JobMaxRunningTimeMinutes,omitempty"`
	// **JobSpecs*	- describes the configurations for job running, such as the image address, startup command, node resource declaration, and number of replicas.
	//
	// A DLC job consists of different types of nodes. If nodes of the same type have exactly the same configuration, the configuration is called JobSpec. **JobSpecs*	- specifies the configurations of all types of nodes. The value is of the array type.
	//
	// This parameter is required.
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// The job type. The value is case-sensitive. The following job types are supported:
	//
	// 	- TFJob
	//
	// 	- PyTorchJob
	//
	// 	- MPIJob
	//
	// 	- XGBoostJob
	//
	// 	- OneFlowJob
	//
	// 	- ElasticBatchJob
	//
	// 	- SlurmJob
	//
	// 	- RayJob
	//
	// Valid values and corresponding frameworks:
	//
	// 	- OneFlowJob: OneFlow.
	//
	// 	- PyTorchJob: PyTorch.
	//
	// 	- SlurmJob: Slurm.
	//
	// 	- XGBoostJob: XGBoost.
	//
	// 	- ElasticBatchJob: ElasticBatch.
	//
	// 	- MPIJob: MPIJob.
	//
	// 	- TFJob: Tensorflow.
	//
	// 	- RayJob: Ray.
	//
	// This parameter is required.
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The additional configuration of the job. You can use this parameter to adjust the behavior of the attached data source. For example, if the attached data source of the job is of the OSS type, you can use this parameter to add the following configurations to override the default parameters of JindoFS: `fs.oss.download.thread.concurrency=4,fs.oss.download.queue.size=16`.
	//
	// example:
	//
	// key1=value1,key2=value2
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The priority of the job. Default value: 1. Valid values: 1 to 9.
	//
	// 	- 1 is the lowest priority.
	//
	// 	- 9: the highest priority.
	//
	// example:
	//
	// 8
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group. This parameter is optional.
	//
	// 	- If you leave this parameter empty, the job is submitted to a public resource group.
	//
	// 	- If a resource quota is bound to the current workspace, you can specify the resource quota ID. For more information about how to query the resource quota ID, see [Manage resource quotas](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// rs-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The additional parameter configurations of the job.
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// The policy that is used to check whether a distributed multi-node job is successful. Only TensorFlow distributed multi-node jobs are supported.
	//
	// 	- ChiefWorker: If you use this policy, the job is considered successful when the pod on the chief node completes operations.
	//
	// 	- AllWorkers (default): If you use this policy, the job is considered successful when all worker nodes complete operations.
	//
	// example:
	//
	// AllWorkers
	SuccessPolicy *string `json:"SuccessPolicy,omitempty" xml:"SuccessPolicy,omitempty"`
	// 任务模板的 ID。指定后将基于模板创建作业，作业参数需符合模板约束规则。
	//
	// example:
	//
	// tplxxxxxxxxxxxxxxxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 指定使用的模板版本号，不传则使用模板默认版本。
	//
	// example:
	//
	// 1
	TemplateVersion *int32 `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The folder in which the third-party Python library file requirements.txt is stored. Before the startup command specified by the UserCommand parameter is run on each node, DLC fetches the requirements.txt file from the folder and runs `pip install -r` to install the required package and library.
	//
	// example:
	//
	// /root/code/
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// The third-party Python libraries to be installed.
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// The startup command for all nodes of the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// python /root/code/mnist.py
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// The VPC settings.
	UserVpc *CreateJobRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// ws-20210126170216-xxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateJobRequest) GetCodeSource() *CreateJobRequestCodeSource {
	return s.CodeSource
}

func (s *CreateJobRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateJobRequest) GetCustomEnvs() []*CreateJobRequestCustomEnvs {
	return s.CustomEnvs
}

func (s *CreateJobRequest) GetDataSources() []*CreateJobRequestDataSources {
	return s.DataSources
}

func (s *CreateJobRequest) GetDebuggerConfigContent() *string {
	return s.DebuggerConfigContent
}

func (s *CreateJobRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateJobRequest) GetElasticSpec() *JobElasticSpec {
	return s.ElasticSpec
}

func (s *CreateJobRequest) GetEnvs() map[string]*string {
	return s.Envs
}

func (s *CreateJobRequest) GetJobMaxRunningTimeMinutes() *int64 {
	return s.JobMaxRunningTimeMinutes
}

func (s *CreateJobRequest) GetJobSpecs() []*JobSpec {
	return s.JobSpecs
}

func (s *CreateJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateJobRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateJobRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateJobRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateJobRequest) GetSettings() *JobSettings {
	return s.Settings
}

func (s *CreateJobRequest) GetSuccessPolicy() *string {
	return s.SuccessPolicy
}

func (s *CreateJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateJobRequest) GetTemplateVersion() *int32 {
	return s.TemplateVersion
}

func (s *CreateJobRequest) GetThirdpartyLibDir() *string {
	return s.ThirdpartyLibDir
}

func (s *CreateJobRequest) GetThirdpartyLibs() []*string {
	return s.ThirdpartyLibs
}

func (s *CreateJobRequest) GetUserCommand() *string {
	return s.UserCommand
}

func (s *CreateJobRequest) GetUserVpc() *CreateJobRequestUserVpc {
	return s.UserVpc
}

func (s *CreateJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateJobRequest) SetAccessibility(v string) *CreateJobRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateJobRequest) SetCodeSource(v *CreateJobRequestCodeSource) *CreateJobRequest {
	s.CodeSource = v
	return s
}

func (s *CreateJobRequest) SetCredentialConfig(v *CredentialConfig) *CreateJobRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateJobRequest) SetCustomEnvs(v []*CreateJobRequestCustomEnvs) *CreateJobRequest {
	s.CustomEnvs = v
	return s
}

func (s *CreateJobRequest) SetDataSources(v []*CreateJobRequestDataSources) *CreateJobRequest {
	s.DataSources = v
	return s
}

func (s *CreateJobRequest) SetDebuggerConfigContent(v string) *CreateJobRequest {
	s.DebuggerConfigContent = &v
	return s
}

func (s *CreateJobRequest) SetDisplayName(v string) *CreateJobRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateJobRequest) SetElasticSpec(v *JobElasticSpec) *CreateJobRequest {
	s.ElasticSpec = v
	return s
}

func (s *CreateJobRequest) SetEnvs(v map[string]*string) *CreateJobRequest {
	s.Envs = v
	return s
}

func (s *CreateJobRequest) SetJobMaxRunningTimeMinutes(v int64) *CreateJobRequest {
	s.JobMaxRunningTimeMinutes = &v
	return s
}

func (s *CreateJobRequest) SetJobSpecs(v []*JobSpec) *CreateJobRequest {
	s.JobSpecs = v
	return s
}

func (s *CreateJobRequest) SetJobType(v string) *CreateJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateJobRequest) SetOptions(v string) *CreateJobRequest {
	s.Options = &v
	return s
}

func (s *CreateJobRequest) SetPriority(v int32) *CreateJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobRequest) SetResourceId(v string) *CreateJobRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateJobRequest) SetSettings(v *JobSettings) *CreateJobRequest {
	s.Settings = v
	return s
}

func (s *CreateJobRequest) SetSuccessPolicy(v string) *CreateJobRequest {
	s.SuccessPolicy = &v
	return s
}

func (s *CreateJobRequest) SetTemplateId(v string) *CreateJobRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateJobRequest) SetTemplateVersion(v int32) *CreateJobRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateJobRequest) SetThirdpartyLibDir(v string) *CreateJobRequest {
	s.ThirdpartyLibDir = &v
	return s
}

func (s *CreateJobRequest) SetThirdpartyLibs(v []*string) *CreateJobRequest {
	s.ThirdpartyLibs = v
	return s
}

func (s *CreateJobRequest) SetUserCommand(v string) *CreateJobRequest {
	s.UserCommand = &v
	return s
}

func (s *CreateJobRequest) SetUserVpc(v *CreateJobRequestUserVpc) *CreateJobRequest {
	s.UserVpc = v
	return s
}

func (s *CreateJobRequest) SetWorkspaceId(v string) *CreateJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateJobRequest) Validate() error {
	if s.CodeSource != nil {
		if err := s.CodeSource.Validate(); err != nil {
			return err
		}
	}
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CustomEnvs != nil {
		for _, item := range s.CustomEnvs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ElasticSpec != nil {
		if err := s.ElasticSpec.Validate(); err != nil {
			return err
		}
	}
	if s.JobSpecs != nil {
		for _, item := range s.JobSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			return err
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobRequestCodeSource struct {
	// The branch of the referenced code repository. By default, the branch configured in the code source is used. This parameter is optional.
	//
	// example:
	//
	// master
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The ID of the code source.
	//
	// example:
	//
	// code-20210111103721-xxxxxxx
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The commit ID of the code to be downloaded. By default, the commit ID configured in the code source is used. This parameter is optional.
	//
	// example:
	//
	// 44da109b5******
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// The path to which the job is mounted. By default, the mount path configured in the data source is used. This parameter is optional.
	//
	// example:
	//
	// /root/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s CreateJobRequestCodeSource) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestCodeSource) GoString() string {
	return s.String()
}

func (s *CreateJobRequestCodeSource) GetBranch() *string {
	return s.Branch
}

func (s *CreateJobRequestCodeSource) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *CreateJobRequestCodeSource) GetCommit() *string {
	return s.Commit
}

func (s *CreateJobRequestCodeSource) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateJobRequestCodeSource) SetBranch(v string) *CreateJobRequestCodeSource {
	s.Branch = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetCodeSourceId(v string) *CreateJobRequestCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetCommit(v string) *CreateJobRequestCodeSource {
	s.Commit = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetMountPath(v string) *CreateJobRequestCodeSource {
	s.MountPath = &v
	return s
}

func (s *CreateJobRequestCodeSource) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestCustomEnvs struct {
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value   *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Visible *string `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s CreateJobRequestCustomEnvs) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestCustomEnvs) GoString() string {
	return s.String()
}

func (s *CreateJobRequestCustomEnvs) GetKey() *string {
	return s.Key
}

func (s *CreateJobRequestCustomEnvs) GetValue() *string {
	return s.Value
}

func (s *CreateJobRequestCustomEnvs) GetVisible() *string {
	return s.Visible
}

func (s *CreateJobRequestCustomEnvs) SetKey(v string) *CreateJobRequestCustomEnvs {
	s.Key = &v
	return s
}

func (s *CreateJobRequestCustomEnvs) SetValue(v string) *CreateJobRequestCustomEnvs {
	s.Value = &v
	return s
}

func (s *CreateJobRequestCustomEnvs) SetVisible(v string) *CreateJobRequestCustomEnvs {
	s.Visible = &v
	return s
}

func (s *CreateJobRequestCustomEnvs) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestDataSources struct {
	// The data source ID.
	//
	// example:
	//
	// d-cn9dl*******
	DataSourceId      *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataSourceVersion *string `json:"DataSourceVersion,omitempty" xml:"DataSourceVersion,omitempty"`
	EnableCache       *bool   `json:"EnableCache,omitempty" xml:"EnableCache,omitempty"`
	MountAccess       *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The path to which the job is mounted. By default, the mount path in the data source configuration is used. This parameter is optional.
	//
	// example:
	//
	// /root/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount attribute of the custom dataset. Set the value to OSS.
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
	// The data source path.
	//
	// example:
	//
	// oss://bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateJobRequestDataSources) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestDataSources) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDataSources) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateJobRequestDataSources) GetDataSourceVersion() *string {
	return s.DataSourceVersion
}

func (s *CreateJobRequestDataSources) GetEnableCache() *bool {
	return s.EnableCache
}

func (s *CreateJobRequestDataSources) GetMountAccess() *string {
	return s.MountAccess
}

func (s *CreateJobRequestDataSources) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateJobRequestDataSources) GetOptions() *string {
	return s.Options
}

func (s *CreateJobRequestDataSources) GetUri() *string {
	return s.Uri
}

func (s *CreateJobRequestDataSources) SetDataSourceId(v string) *CreateJobRequestDataSources {
	s.DataSourceId = &v
	return s
}

func (s *CreateJobRequestDataSources) SetDataSourceVersion(v string) *CreateJobRequestDataSources {
	s.DataSourceVersion = &v
	return s
}

func (s *CreateJobRequestDataSources) SetEnableCache(v bool) *CreateJobRequestDataSources {
	s.EnableCache = &v
	return s
}

func (s *CreateJobRequestDataSources) SetMountAccess(v string) *CreateJobRequestDataSources {
	s.MountAccess = &v
	return s
}

func (s *CreateJobRequestDataSources) SetMountPath(v string) *CreateJobRequestDataSources {
	s.MountPath = &v
	return s
}

func (s *CreateJobRequestDataSources) SetOptions(v string) *CreateJobRequestDataSources {
	s.Options = &v
	return s
}

func (s *CreateJobRequestDataSources) SetUri(v string) *CreateJobRequestDataSources {
	s.Uri = &v
	return s
}

func (s *CreateJobRequestDataSources) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestUserVpc struct {
	// The default route. Default value: false. Valid values:
	//
	// 	- eth0: The default network interface is used to access the Internet through the public gateway.
	//
	// 	- eth1: The user\\"s elastic network interface (ENI) is used to access the Internet through the private gateway. For more information about the configuration method, see [Enable Internet access for a DSW instance by using a private Internet NAT gateway](https://help.aliyun.com/document_detail/2525343.html).
	//
	// example:
	//
	// eth0
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR block.
	//
	// 	- If you leave the SwitchId and ExtendedCIDRs parameters empty, the system automatically obtains all CIDR blocks in a VPC.
	//
	// 	- If you configure the SwitchId and ExtendedCIDRs parameters, we recommend that you specify all CIDR blocks in a VPC.
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID. This parameter is optional.
	//
	// 	- If you leave this parameter empty, the system automatically selects a vSwitch based on the inventory status.
	//
	// 	- You can also specify a vSwitch ID.
	//
	// example:
	//
	// vs-abcdef****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-abcdef****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateJobRequestUserVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateJobRequestUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *CreateJobRequestUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *CreateJobRequestUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateJobRequestUserVpc) GetSwitchId() *string {
	return s.SwitchId
}

func (s *CreateJobRequestUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateJobRequestUserVpc) SetDefaultRoute(v string) *CreateJobRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetExtendedCIDRs(v []*string) *CreateJobRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *CreateJobRequestUserVpc) SetSecurityGroupId(v string) *CreateJobRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetSwitchId(v string) *CreateJobRequestUserVpc {
	s.SwitchId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetVpcId(v string) *CreateJobRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *CreateJobRequestUserVpc) Validate() error {
	return dara.Validate(s)
}
