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
	SetDescription(v string) *CreateJobRequest
	GetDescription() *string
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
	SetSchedulingStrategy(v string) *CreateJobRequest
	GetSchedulingStrategy() *string
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
	// The visibility of the job. Valid values:
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The code source used by this job. Before the job nodes start, DLC automatically downloads the code configured in the code source and mounts it to a local directory in the container.
	CodeSource *CreateJobRequestCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// The access credential configuration.
	CredentialConfig *CredentialConfig             `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	CustomEnvs       []*CreateJobRequestCustomEnvs `json:"CustomEnvs,omitempty" xml:"CustomEnvs,omitempty" type:"Repeated"`
	// The list of data sources used by the job.
	DataSources []*CreateJobRequestDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// This parameter is not currently supported. Ignore this parameter.
	//
	// example:
	//
	// “”
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the job. The naming format is as follows:
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is not currently supported. Ignore this parameter.
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// The environment variable configuration.
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The maximum running duration of the job, in minutes.
	//
	// example:
	//
	// 1024
	JobMaxRunningTimeMinutes *int64 `json:"JobMaxRunningTimeMinutes,omitempty" xml:"JobMaxRunningTimeMinutes,omitempty"`
	// **JobSpecs*	- describes various configurations for job runtime, such as image address, startup command, node resource declarations, and number of replicas.
	//
	// This parameter is required.
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// The job type. This parameter is case-sensitive. Currently supported job types:
	//
	// This parameter is required.
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The additional configuration for this node. You can use this parameter to adjust certain behaviors of mounted data sources. For example, if the node has an OSS-type data source mounted, you can set this parameter to `fs.oss.download.thread.concurrency=4,fs.oss.download.queue.size=16` to overwrite the default JindoFS parameter settings.
	//
	// example:
	//
	// key1=value1,key2=value2
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The priority of the job. This is an optional parameter. The default value is 1. Valid values: 1 to 9. Specifically:
	//
	// example:
	//
	// 8
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource group ID. This is an optional parameter.
	//
	// example:
	//
	// rs-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Auto
	SchedulingStrategy *string `json:"SchedulingStrategy,omitempty" xml:"SchedulingStrategy,omitempty"`
	// The additional parameter settings for the job.
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// The success policy for distributed multi-node jobs. Currently only TensorFlow multi-node jobs support this parameter.
	//
	// example:
	//
	// AllWorkers
	SuccessPolicy *string `json:"SuccessPolicy,omitempty" xml:"SuccessPolicy,omitempty"`
	// The job template ID.
	//
	// example:
	//
	// tplabc1234567
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The job template version.
	//
	// example:
	//
	// 1
	TemplateVersion *int32 `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The folder name where the third-party Python library (requirements.txt) file is located. Before running the specified UserCommand on each node, PAI-DLC retrieves the requirements.txt file from the specified folder and runs `pip install -r` to install the libraries.
	//
	// example:
	//
	// /root/code/
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// The list of third-party Python libraries to install.
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// The startup command for all nodes of the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// python /root/code/mnist.py
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// The user VPC configuration.
	UserVpc *CreateJobRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The workspace ID. <props="china">For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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

func (s *CreateJobRequest) GetDescription() *string {
	return s.Description
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

func (s *CreateJobRequest) GetSchedulingStrategy() *string {
	return s.SchedulingStrategy
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

func (s *CreateJobRequest) SetDescription(v string) *CreateJobRequest {
	s.Description = &v
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

func (s *CreateJobRequest) SetSchedulingStrategy(v string) *CreateJobRequest {
	s.SchedulingStrategy = &v
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
	// The branch of the code repository referenced when this job runs. This is an optional parameter. By default, the branch configured in the code source is used.
	//
	// example:
	//
	// master
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The code source ID. <props="china">For information about how to obtain the code source ID, see [ListCodeSources](https://help.aliyun.com/document_detail/459922.html).
	//
	// example:
	//
	// code-20210111103721-xxxxxxx
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The commit ID of the code to download for this job. This is an optional parameter. By default, the CommitID configured in the code source is used.
	//
	// example:
	//
	// 44da109b5******
	Commit            *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	IsSharedMountPath *bool   `json:"IsSharedMountPath,omitempty" xml:"IsSharedMountPath,omitempty"`
	// The mount path for this job. This is an optional parameter. By default, the mount path configured in the code source is used.
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

func (s *CreateJobRequestCodeSource) GetIsSharedMountPath() *bool {
	return s.IsSharedMountPath
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

func (s *CreateJobRequestCodeSource) SetIsSharedMountPath(v bool) *CreateJobRequestCodeSource {
	s.IsSharedMountPath = &v
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
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The ID of the data source. <props="china">For information about how to view the data source ID, see [ListDatasets](https://help.aliyun.com/document_detail/457222.html).
	//
	// example:
	//
	// d-cn9dl*******
	DataSourceId      *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataSourceVersion *string `json:"DataSourceVersion,omitempty" xml:"DataSourceVersion,omitempty"`
	EnableCache       *bool   `json:"EnableCache,omitempty" xml:"EnableCache,omitempty"`
	MountAccess       *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The mount path for this job. This is an optional parameter. By default, the mount path configured in the data source is used.
	//
	// example:
	//
	// /root/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// Custom dataset mount properties. Currently only OSS is supported.
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
	Options   *string `json:"Options,omitempty" xml:"Options,omitempty"`
	RoleChain *string `json:"RoleChain,omitempty" xml:"RoleChain,omitempty"`
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

func (s *CreateJobRequestDataSources) GetAccessPointId() *string {
	return s.AccessPointId
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

func (s *CreateJobRequestDataSources) GetRoleChain() *string {
	return s.RoleChain
}

func (s *CreateJobRequestDataSources) GetUri() *string {
	return s.Uri
}

func (s *CreateJobRequestDataSources) SetAccessPointId(v string) *CreateJobRequestDataSources {
	s.AccessPointId = &v
	return s
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

func (s *CreateJobRequestDataSources) SetRoleChain(v string) *CreateJobRequestDataSources {
	s.RoleChain = &v
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
	// The default route. Valid values:
	//
	// example:
	//
	// eth0
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR blocks.
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The ID of the user security group.
	//
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the user vSwitch. This is an optional parameter.
	//
	// example:
	//
	// vs-abcdef****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The ID of the user VPC.
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
