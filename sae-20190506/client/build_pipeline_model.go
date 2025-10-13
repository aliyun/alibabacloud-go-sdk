// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildPipeline interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *BuildPipeline
	GetApplicationId() *string
	SetApplicationName(v string) *BuildPipeline
	GetApplicationName() *string
	SetBuildConfig(v *BuildPipelineBuildConfig) *BuildPipeline
	GetBuildConfig() *BuildPipelineBuildConfig
	SetCodeConfig(v *BuildPipelineCodeConfig) *BuildPipeline
	GetCodeConfig() *BuildPipelineCodeConfig
	SetDeployConfig(v *BuildPipelineDeployConfig) *BuildPipeline
	GetDeployConfig() *BuildPipelineDeployConfig
	SetEnabled(v bool) *BuildPipeline
	GetEnabled() *bool
	SetImageConfig(v *BuildPipelineImageConfig) *BuildPipeline
	GetImageConfig() *BuildPipelineImageConfig
	SetPackageConfig(v *BuildPipelinePackageConfig) *BuildPipeline
	GetPackageConfig() *BuildPipelinePackageConfig
	SetTriggerConfig(v *BuildPipelineTriggerConfig) *BuildPipeline
	GetTriggerConfig() *BuildPipelineTriggerConfig
}

type BuildPipeline struct {
	// This parameter is required.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// This parameter is required.
	BuildConfig *BuildPipelineBuildConfig `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
	// This parameter is required.
	CodeConfig    *BuildPipelineCodeConfig    `json:"CodeConfig,omitempty" xml:"CodeConfig,omitempty" type:"Struct"`
	DeployConfig  *BuildPipelineDeployConfig  `json:"DeployConfig,omitempty" xml:"DeployConfig,omitempty" type:"Struct"`
	Enabled       *bool                       `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ImageConfig   *BuildPipelineImageConfig   `json:"ImageConfig,omitempty" xml:"ImageConfig,omitempty" type:"Struct"`
	PackageConfig *BuildPipelinePackageConfig `json:"PackageConfig,omitempty" xml:"PackageConfig,omitempty" type:"Struct"`
	// This parameter is required.
	TriggerConfig *BuildPipelineTriggerConfig `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty" type:"Struct"`
}

func (s BuildPipeline) String() string {
	return dara.Prettify(s)
}

func (s BuildPipeline) GoString() string {
	return s.String()
}

func (s *BuildPipeline) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *BuildPipeline) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *BuildPipeline) GetBuildConfig() *BuildPipelineBuildConfig {
	return s.BuildConfig
}

func (s *BuildPipeline) GetCodeConfig() *BuildPipelineCodeConfig {
	return s.CodeConfig
}

func (s *BuildPipeline) GetDeployConfig() *BuildPipelineDeployConfig {
	return s.DeployConfig
}

func (s *BuildPipeline) GetEnabled() *bool {
	return s.Enabled
}

func (s *BuildPipeline) GetImageConfig() *BuildPipelineImageConfig {
	return s.ImageConfig
}

func (s *BuildPipeline) GetPackageConfig() *BuildPipelinePackageConfig {
	return s.PackageConfig
}

func (s *BuildPipeline) GetTriggerConfig() *BuildPipelineTriggerConfig {
	return s.TriggerConfig
}

func (s *BuildPipeline) SetApplicationId(v string) *BuildPipeline {
	s.ApplicationId = &v
	return s
}

func (s *BuildPipeline) SetApplicationName(v string) *BuildPipeline {
	s.ApplicationName = &v
	return s
}

func (s *BuildPipeline) SetBuildConfig(v *BuildPipelineBuildConfig) *BuildPipeline {
	s.BuildConfig = v
	return s
}

func (s *BuildPipeline) SetCodeConfig(v *BuildPipelineCodeConfig) *BuildPipeline {
	s.CodeConfig = v
	return s
}

func (s *BuildPipeline) SetDeployConfig(v *BuildPipelineDeployConfig) *BuildPipeline {
	s.DeployConfig = v
	return s
}

func (s *BuildPipeline) SetEnabled(v bool) *BuildPipeline {
	s.Enabled = &v
	return s
}

func (s *BuildPipeline) SetImageConfig(v *BuildPipelineImageConfig) *BuildPipeline {
	s.ImageConfig = v
	return s
}

func (s *BuildPipeline) SetPackageConfig(v *BuildPipelinePackageConfig) *BuildPipeline {
	s.PackageConfig = v
	return s
}

func (s *BuildPipeline) SetTriggerConfig(v *BuildPipelineTriggerConfig) *BuildPipeline {
	s.TriggerConfig = v
	return s
}

func (s *BuildPipeline) Validate() error {
	if s.BuildConfig != nil {
		if err := s.BuildConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CodeConfig != nil {
		if err := s.CodeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DeployConfig != nil {
		if err := s.DeployConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ImageConfig != nil {
		if err := s.ImageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PackageConfig != nil {
		if err := s.PackageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerConfig != nil {
		if err := s.TriggerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BuildPipelineBuildConfig struct {
	// example:
	//
	// mvn clean package
	BeforeBuildCommand *string `json:"BeforeBuildCommand,omitempty" xml:"BeforeBuildCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// buildpacks/dockerfile
	BuildType *string `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	// example:
	//
	// code/Dockerfile
	DockerfilePath *string `json:"DockerfilePath,omitempty" xml:"DockerfilePath,omitempty"`
	// example:
	//
	// java -jar target/app.jar
	RunCommand     *string       `json:"RunCommand,omitempty" xml:"RunCommand,omitempty"`
	RuntimeType    *string       `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
	RuntimeVersion *string       `json:"RuntimeVersion,omitempty" xml:"RuntimeVersion,omitempty"`
	TomcatConfig   *TomcatConfig `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// example:
	//
	// code
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s BuildPipelineBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineBuildConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineBuildConfig) GetBeforeBuildCommand() *string {
	return s.BeforeBuildCommand
}

func (s *BuildPipelineBuildConfig) GetBuildType() *string {
	return s.BuildType
}

func (s *BuildPipelineBuildConfig) GetDockerfilePath() *string {
	return s.DockerfilePath
}

func (s *BuildPipelineBuildConfig) GetRunCommand() *string {
	return s.RunCommand
}

func (s *BuildPipelineBuildConfig) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *BuildPipelineBuildConfig) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *BuildPipelineBuildConfig) GetTomcatConfig() *TomcatConfig {
	return s.TomcatConfig
}

func (s *BuildPipelineBuildConfig) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *BuildPipelineBuildConfig) SetBeforeBuildCommand(v string) *BuildPipelineBuildConfig {
	s.BeforeBuildCommand = &v
	return s
}

func (s *BuildPipelineBuildConfig) SetBuildType(v string) *BuildPipelineBuildConfig {
	s.BuildType = &v
	return s
}

func (s *BuildPipelineBuildConfig) SetDockerfilePath(v string) *BuildPipelineBuildConfig {
	s.DockerfilePath = &v
	return s
}

func (s *BuildPipelineBuildConfig) SetRunCommand(v string) *BuildPipelineBuildConfig {
	s.RunCommand = &v
	return s
}

func (s *BuildPipelineBuildConfig) SetRuntimeType(v string) *BuildPipelineBuildConfig {
	s.RuntimeType = &v
	return s
}

func (s *BuildPipelineBuildConfig) SetRuntimeVersion(v string) *BuildPipelineBuildConfig {
	s.RuntimeVersion = &v
	return s
}

func (s *BuildPipelineBuildConfig) SetTomcatConfig(v *TomcatConfig) *BuildPipelineBuildConfig {
	s.TomcatConfig = v
	return s
}

func (s *BuildPipelineBuildConfig) SetWorkingDir(v string) *BuildPipelineBuildConfig {
	s.WorkingDir = &v
	return s
}

func (s *BuildPipelineBuildConfig) Validate() error {
	if s.TomcatConfig != nil {
		if err := s.TomcatConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BuildPipelineCodeConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// master
	BranchName     *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	CommitId       *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	CommitUrl      *string `json:"CommitUrl,omitempty" xml:"CommitUrl,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// github/gitee/gitlab等
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sae-samples/java-maven-demo
	RepoFullName *string `json:"RepoFullName,omitempty" xml:"RepoFullName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3001
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s BuildPipelineCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineCodeConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineCodeConfig) GetAccountId() *string {
	return s.AccountId
}

func (s *BuildPipelineCodeConfig) GetBranchName() *string {
	return s.BranchName
}

func (s *BuildPipelineCodeConfig) GetCommitId() *string {
	return s.CommitId
}

func (s *BuildPipelineCodeConfig) GetCommitUrl() *string {
	return s.CommitUrl
}

func (s *BuildPipelineCodeConfig) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *BuildPipelineCodeConfig) GetProvider() *string {
	return s.Provider
}

func (s *BuildPipelineCodeConfig) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *BuildPipelineCodeConfig) GetRepoId() *string {
	return s.RepoId
}

func (s *BuildPipelineCodeConfig) SetAccountId(v string) *BuildPipelineCodeConfig {
	s.AccountId = &v
	return s
}

func (s *BuildPipelineCodeConfig) SetBranchName(v string) *BuildPipelineCodeConfig {
	s.BranchName = &v
	return s
}

func (s *BuildPipelineCodeConfig) SetCommitId(v string) *BuildPipelineCodeConfig {
	s.CommitId = &v
	return s
}

func (s *BuildPipelineCodeConfig) SetCommitUrl(v string) *BuildPipelineCodeConfig {
	s.CommitUrl = &v
	return s
}

func (s *BuildPipelineCodeConfig) SetOrganizationId(v string) *BuildPipelineCodeConfig {
	s.OrganizationId = &v
	return s
}

func (s *BuildPipelineCodeConfig) SetProvider(v string) *BuildPipelineCodeConfig {
	s.Provider = &v
	return s
}

func (s *BuildPipelineCodeConfig) SetRepoFullName(v string) *BuildPipelineCodeConfig {
	s.RepoFullName = &v
	return s
}

func (s *BuildPipelineCodeConfig) SetRepoId(v string) *BuildPipelineCodeConfig {
	s.RepoId = &v
	return s
}

func (s *BuildPipelineCodeConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineDeployConfig struct {
	AlwaysAllocateCPU      *bool   `json:"AlwaysAllocateCPU,omitempty" xml:"AlwaysAllocateCPU,omitempty"`
	MaximumInstanceCount   *int32  `json:"MaximumInstanceCount,omitempty" xml:"MaximumInstanceCount,omitempty"`
	MinimumInstanceCount   *int32  `json:"MinimumInstanceCount,omitempty" xml:"MinimumInstanceCount,omitempty"`
	UpdateApplicationInput *string `json:"UpdateApplicationInput,omitempty" xml:"UpdateApplicationInput,omitempty"`
	UpdateTraffic          *bool   `json:"UpdateTraffic,omitempty" xml:"UpdateTraffic,omitempty"`
}

func (s BuildPipelineDeployConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineDeployConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineDeployConfig) GetAlwaysAllocateCPU() *bool {
	return s.AlwaysAllocateCPU
}

func (s *BuildPipelineDeployConfig) GetMaximumInstanceCount() *int32 {
	return s.MaximumInstanceCount
}

func (s *BuildPipelineDeployConfig) GetMinimumInstanceCount() *int32 {
	return s.MinimumInstanceCount
}

func (s *BuildPipelineDeployConfig) GetUpdateApplicationInput() *string {
	return s.UpdateApplicationInput
}

func (s *BuildPipelineDeployConfig) GetUpdateTraffic() *bool {
	return s.UpdateTraffic
}

func (s *BuildPipelineDeployConfig) SetAlwaysAllocateCPU(v bool) *BuildPipelineDeployConfig {
	s.AlwaysAllocateCPU = &v
	return s
}

func (s *BuildPipelineDeployConfig) SetMaximumInstanceCount(v int32) *BuildPipelineDeployConfig {
	s.MaximumInstanceCount = &v
	return s
}

func (s *BuildPipelineDeployConfig) SetMinimumInstanceCount(v int32) *BuildPipelineDeployConfig {
	s.MinimumInstanceCount = &v
	return s
}

func (s *BuildPipelineDeployConfig) SetUpdateApplicationInput(v string) *BuildPipelineDeployConfig {
	s.UpdateApplicationInput = &v
	return s
}

func (s *BuildPipelineDeployConfig) SetUpdateTraffic(v bool) *BuildPipelineDeployConfig {
	s.UpdateTraffic = &v
	return s
}

func (s *BuildPipelineDeployConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineImageConfig struct {
	// example:
	//
	// ACR/ACREE
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Repository   *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s BuildPipelineImageConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineImageConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineImageConfig) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BuildPipelineImageConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *BuildPipelineImageConfig) GetRepository() *string {
	return s.Repository
}

func (s *BuildPipelineImageConfig) SetInstanceType(v string) *BuildPipelineImageConfig {
	s.InstanceType = &v
	return s
}

func (s *BuildPipelineImageConfig) SetNamespace(v string) *BuildPipelineImageConfig {
	s.Namespace = &v
	return s
}

func (s *BuildPipelineImageConfig) SetRepository(v string) *BuildPipelineImageConfig {
	s.Repository = &v
	return s
}

func (s *BuildPipelineImageConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelinePackageConfig struct {
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// war/jar/zip
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// example:
	//
	// http://myoss.oss-cn-****.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// example:
	//
	// 1.0.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
}

func (s BuildPipelinePackageConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelinePackageConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelinePackageConfig) GetPackageName() *string {
	return s.PackageName
}

func (s *BuildPipelinePackageConfig) GetPackageType() *string {
	return s.PackageType
}

func (s *BuildPipelinePackageConfig) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *BuildPipelinePackageConfig) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *BuildPipelinePackageConfig) SetPackageName(v string) *BuildPipelinePackageConfig {
	s.PackageName = &v
	return s
}

func (s *BuildPipelinePackageConfig) SetPackageType(v string) *BuildPipelinePackageConfig {
	s.PackageType = &v
	return s
}

func (s *BuildPipelinePackageConfig) SetPackageUrl(v string) *BuildPipelinePackageConfig {
	s.PackageUrl = &v
	return s
}

func (s *BuildPipelinePackageConfig) SetPackageVersion(v string) *BuildPipelinePackageConfig {
	s.PackageVersion = &v
	return s
}

func (s *BuildPipelinePackageConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineTriggerConfig struct {
	// example:
	//
	// master
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// example:
	//
	// v1
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PUSH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BuildPipelineTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineTriggerConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineTriggerConfig) GetBranchName() *string {
	return s.BranchName
}

func (s *BuildPipelineTriggerConfig) GetTagName() *string {
	return s.TagName
}

func (s *BuildPipelineTriggerConfig) GetType() *string {
	return s.Type
}

func (s *BuildPipelineTriggerConfig) SetBranchName(v string) *BuildPipelineTriggerConfig {
	s.BranchName = &v
	return s
}

func (s *BuildPipelineTriggerConfig) SetTagName(v string) *BuildPipelineTriggerConfig {
	s.TagName = &v
	return s
}

func (s *BuildPipelineTriggerConfig) SetType(v string) *BuildPipelineTriggerConfig {
	s.Type = &v
	return s
}

func (s *BuildPipelineTriggerConfig) Validate() error {
	return dara.Validate(s)
}
