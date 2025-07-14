// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildPipelineRun interface {
	dara.Model
	String() string
	GoString() string
	SetBuildConfig(v *BuildPipelineRunBuildConfig) *BuildPipelineRun
	GetBuildConfig() *BuildPipelineRunBuildConfig
	SetBuildDuration(v int64) *BuildPipelineRun
	GetBuildDuration() *int64
	SetCodeConfig(v *BuildPipelineRunCodeConfig) *BuildPipelineRun
	GetCodeConfig() *BuildPipelineRunCodeConfig
	SetCreateTime(v int64) *BuildPipelineRun
	GetCreateTime() *int64
	SetDeployConfig(v *BuildPipelineRunDeployConfig) *BuildPipelineRun
	GetDeployConfig() *BuildPipelineRunDeployConfig
	SetDeployDuration(v int64) *BuildPipelineRun
	GetDeployDuration() *int64
	SetEndTime(v int64) *BuildPipelineRun
	GetEndTime() *int64
	SetImageConfig(v *BuildPipelineRunImageConfig) *BuildPipelineRun
	GetImageConfig() *BuildPipelineRunImageConfig
	SetPackageConfig(v *BuildPipelineRunPackageConfig) *BuildPipelineRun
	GetPackageConfig() *BuildPipelineRunPackageConfig
	SetPipelineId(v string) *BuildPipelineRun
	GetPipelineId() *string
	SetPipelineRunId(v string) *BuildPipelineRun
	GetPipelineRunId() *string
	SetStartTime(v int64) *BuildPipelineRun
	GetStartTime() *int64
	SetStatus(v string) *BuildPipelineRun
	GetStatus() *string
	SetSteps(v []*BuildPipelineRunSteps) *BuildPipelineRun
	GetSteps() []*BuildPipelineRunSteps
	SetTriggerConfig(v *BuildPipelineRunTriggerConfig) *BuildPipelineRun
	GetTriggerConfig() *BuildPipelineRunTriggerConfig
	SetVersionId(v string) *BuildPipelineRun
	GetVersionId() *string
	SetWaitDuration(v int64) *BuildPipelineRun
	GetWaitDuration() *int64
}

type BuildPipelineRun struct {
	BuildConfig    *BuildPipelineRunBuildConfig   `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
	BuildDuration  *int64                         `json:"BuildDuration,omitempty" xml:"BuildDuration,omitempty"`
	CodeConfig     *BuildPipelineRunCodeConfig    `json:"CodeConfig,omitempty" xml:"CodeConfig,omitempty" type:"Struct"`
	CreateTime     *int64                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployConfig   *BuildPipelineRunDeployConfig  `json:"DeployConfig,omitempty" xml:"DeployConfig,omitempty" type:"Struct"`
	DeployDuration *int64                         `json:"DeployDuration,omitempty" xml:"DeployDuration,omitempty"`
	EndTime        *int64                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ImageConfig    *BuildPipelineRunImageConfig   `json:"ImageConfig,omitempty" xml:"ImageConfig,omitempty" type:"Struct"`
	PackageConfig  *BuildPipelineRunPackageConfig `json:"PackageConfig,omitempty" xml:"PackageConfig,omitempty" type:"Struct"`
	PipelineId     *string                        `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	PipelineRunId  *string                        `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	StartTime      *int64                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status         *string                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Steps          []*BuildPipelineRunSteps       `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	TriggerConfig  *BuildPipelineRunTriggerConfig `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty" type:"Struct"`
	VersionId      *string                        `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	WaitDuration   *int64                         `json:"WaitDuration,omitempty" xml:"WaitDuration,omitempty"`
}

func (s BuildPipelineRun) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRun) GoString() string {
	return s.String()
}

func (s *BuildPipelineRun) GetBuildConfig() *BuildPipelineRunBuildConfig {
	return s.BuildConfig
}

func (s *BuildPipelineRun) GetBuildDuration() *int64 {
	return s.BuildDuration
}

func (s *BuildPipelineRun) GetCodeConfig() *BuildPipelineRunCodeConfig {
	return s.CodeConfig
}

func (s *BuildPipelineRun) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *BuildPipelineRun) GetDeployConfig() *BuildPipelineRunDeployConfig {
	return s.DeployConfig
}

func (s *BuildPipelineRun) GetDeployDuration() *int64 {
	return s.DeployDuration
}

func (s *BuildPipelineRun) GetEndTime() *int64 {
	return s.EndTime
}

func (s *BuildPipelineRun) GetImageConfig() *BuildPipelineRunImageConfig {
	return s.ImageConfig
}

func (s *BuildPipelineRun) GetPackageConfig() *BuildPipelineRunPackageConfig {
	return s.PackageConfig
}

func (s *BuildPipelineRun) GetPipelineId() *string {
	return s.PipelineId
}

func (s *BuildPipelineRun) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *BuildPipelineRun) GetStartTime() *int64 {
	return s.StartTime
}

func (s *BuildPipelineRun) GetStatus() *string {
	return s.Status
}

func (s *BuildPipelineRun) GetSteps() []*BuildPipelineRunSteps {
	return s.Steps
}

func (s *BuildPipelineRun) GetTriggerConfig() *BuildPipelineRunTriggerConfig {
	return s.TriggerConfig
}

func (s *BuildPipelineRun) GetVersionId() *string {
	return s.VersionId
}

func (s *BuildPipelineRun) GetWaitDuration() *int64 {
	return s.WaitDuration
}

func (s *BuildPipelineRun) SetBuildConfig(v *BuildPipelineRunBuildConfig) *BuildPipelineRun {
	s.BuildConfig = v
	return s
}

func (s *BuildPipelineRun) SetBuildDuration(v int64) *BuildPipelineRun {
	s.BuildDuration = &v
	return s
}

func (s *BuildPipelineRun) SetCodeConfig(v *BuildPipelineRunCodeConfig) *BuildPipelineRun {
	s.CodeConfig = v
	return s
}

func (s *BuildPipelineRun) SetCreateTime(v int64) *BuildPipelineRun {
	s.CreateTime = &v
	return s
}

func (s *BuildPipelineRun) SetDeployConfig(v *BuildPipelineRunDeployConfig) *BuildPipelineRun {
	s.DeployConfig = v
	return s
}

func (s *BuildPipelineRun) SetDeployDuration(v int64) *BuildPipelineRun {
	s.DeployDuration = &v
	return s
}

func (s *BuildPipelineRun) SetEndTime(v int64) *BuildPipelineRun {
	s.EndTime = &v
	return s
}

func (s *BuildPipelineRun) SetImageConfig(v *BuildPipelineRunImageConfig) *BuildPipelineRun {
	s.ImageConfig = v
	return s
}

func (s *BuildPipelineRun) SetPackageConfig(v *BuildPipelineRunPackageConfig) *BuildPipelineRun {
	s.PackageConfig = v
	return s
}

func (s *BuildPipelineRun) SetPipelineId(v string) *BuildPipelineRun {
	s.PipelineId = &v
	return s
}

func (s *BuildPipelineRun) SetPipelineRunId(v string) *BuildPipelineRun {
	s.PipelineRunId = &v
	return s
}

func (s *BuildPipelineRun) SetStartTime(v int64) *BuildPipelineRun {
	s.StartTime = &v
	return s
}

func (s *BuildPipelineRun) SetStatus(v string) *BuildPipelineRun {
	s.Status = &v
	return s
}

func (s *BuildPipelineRun) SetSteps(v []*BuildPipelineRunSteps) *BuildPipelineRun {
	s.Steps = v
	return s
}

func (s *BuildPipelineRun) SetTriggerConfig(v *BuildPipelineRunTriggerConfig) *BuildPipelineRun {
	s.TriggerConfig = v
	return s
}

func (s *BuildPipelineRun) SetVersionId(v string) *BuildPipelineRun {
	s.VersionId = &v
	return s
}

func (s *BuildPipelineRun) SetWaitDuration(v int64) *BuildPipelineRun {
	s.WaitDuration = &v
	return s
}

func (s *BuildPipelineRun) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineRunBuildConfig struct {
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
	// This parameter is required.
	Trigger *BuildPipelineRunBuildConfigTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// example:
	//
	// code
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s BuildPipelineRunBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunBuildConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunBuildConfig) GetBeforeBuildCommand() *string {
	return s.BeforeBuildCommand
}

func (s *BuildPipelineRunBuildConfig) GetBuildType() *string {
	return s.BuildType
}

func (s *BuildPipelineRunBuildConfig) GetDockerfilePath() *string {
	return s.DockerfilePath
}

func (s *BuildPipelineRunBuildConfig) GetRunCommand() *string {
	return s.RunCommand
}

func (s *BuildPipelineRunBuildConfig) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *BuildPipelineRunBuildConfig) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *BuildPipelineRunBuildConfig) GetTomcatConfig() *TomcatConfig {
	return s.TomcatConfig
}

func (s *BuildPipelineRunBuildConfig) GetTrigger() *BuildPipelineRunBuildConfigTrigger {
	return s.Trigger
}

func (s *BuildPipelineRunBuildConfig) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *BuildPipelineRunBuildConfig) SetBeforeBuildCommand(v string) *BuildPipelineRunBuildConfig {
	s.BeforeBuildCommand = &v
	return s
}

func (s *BuildPipelineRunBuildConfig) SetBuildType(v string) *BuildPipelineRunBuildConfig {
	s.BuildType = &v
	return s
}

func (s *BuildPipelineRunBuildConfig) SetDockerfilePath(v string) *BuildPipelineRunBuildConfig {
	s.DockerfilePath = &v
	return s
}

func (s *BuildPipelineRunBuildConfig) SetRunCommand(v string) *BuildPipelineRunBuildConfig {
	s.RunCommand = &v
	return s
}

func (s *BuildPipelineRunBuildConfig) SetRuntimeType(v string) *BuildPipelineRunBuildConfig {
	s.RuntimeType = &v
	return s
}

func (s *BuildPipelineRunBuildConfig) SetRuntimeVersion(v string) *BuildPipelineRunBuildConfig {
	s.RuntimeVersion = &v
	return s
}

func (s *BuildPipelineRunBuildConfig) SetTomcatConfig(v *TomcatConfig) *BuildPipelineRunBuildConfig {
	s.TomcatConfig = v
	return s
}

func (s *BuildPipelineRunBuildConfig) SetTrigger(v *BuildPipelineRunBuildConfigTrigger) *BuildPipelineRunBuildConfig {
	s.Trigger = v
	return s
}

func (s *BuildPipelineRunBuildConfig) SetWorkingDir(v string) *BuildPipelineRunBuildConfig {
	s.WorkingDir = &v
	return s
}

func (s *BuildPipelineRunBuildConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineRunBuildConfigTrigger struct {
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

func (s BuildPipelineRunBuildConfigTrigger) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunBuildConfigTrigger) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunBuildConfigTrigger) GetBranchName() *string {
	return s.BranchName
}

func (s *BuildPipelineRunBuildConfigTrigger) GetTagName() *string {
	return s.TagName
}

func (s *BuildPipelineRunBuildConfigTrigger) GetType() *string {
	return s.Type
}

func (s *BuildPipelineRunBuildConfigTrigger) SetBranchName(v string) *BuildPipelineRunBuildConfigTrigger {
	s.BranchName = &v
	return s
}

func (s *BuildPipelineRunBuildConfigTrigger) SetTagName(v string) *BuildPipelineRunBuildConfigTrigger {
	s.TagName = &v
	return s
}

func (s *BuildPipelineRunBuildConfigTrigger) SetType(v string) *BuildPipelineRunBuildConfigTrigger {
	s.Type = &v
	return s
}

func (s *BuildPipelineRunBuildConfigTrigger) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineRunCodeConfig struct {
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

func (s BuildPipelineRunCodeConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunCodeConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunCodeConfig) GetAccountId() *string {
	return s.AccountId
}

func (s *BuildPipelineRunCodeConfig) GetBranchName() *string {
	return s.BranchName
}

func (s *BuildPipelineRunCodeConfig) GetCommitId() *string {
	return s.CommitId
}

func (s *BuildPipelineRunCodeConfig) GetCommitUrl() *string {
	return s.CommitUrl
}

func (s *BuildPipelineRunCodeConfig) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *BuildPipelineRunCodeConfig) GetProvider() *string {
	return s.Provider
}

func (s *BuildPipelineRunCodeConfig) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *BuildPipelineRunCodeConfig) GetRepoId() *string {
	return s.RepoId
}

func (s *BuildPipelineRunCodeConfig) SetAccountId(v string) *BuildPipelineRunCodeConfig {
	s.AccountId = &v
	return s
}

func (s *BuildPipelineRunCodeConfig) SetBranchName(v string) *BuildPipelineRunCodeConfig {
	s.BranchName = &v
	return s
}

func (s *BuildPipelineRunCodeConfig) SetCommitId(v string) *BuildPipelineRunCodeConfig {
	s.CommitId = &v
	return s
}

func (s *BuildPipelineRunCodeConfig) SetCommitUrl(v string) *BuildPipelineRunCodeConfig {
	s.CommitUrl = &v
	return s
}

func (s *BuildPipelineRunCodeConfig) SetOrganizationId(v string) *BuildPipelineRunCodeConfig {
	s.OrganizationId = &v
	return s
}

func (s *BuildPipelineRunCodeConfig) SetProvider(v string) *BuildPipelineRunCodeConfig {
	s.Provider = &v
	return s
}

func (s *BuildPipelineRunCodeConfig) SetRepoFullName(v string) *BuildPipelineRunCodeConfig {
	s.RepoFullName = &v
	return s
}

func (s *BuildPipelineRunCodeConfig) SetRepoId(v string) *BuildPipelineRunCodeConfig {
	s.RepoId = &v
	return s
}

func (s *BuildPipelineRunCodeConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineRunDeployConfig struct {
	AlwaysAllocateCPU      *bool   `json:"AlwaysAllocateCPU,omitempty" xml:"AlwaysAllocateCPU,omitempty"`
	MaximumInstanceCount   *int32  `json:"MaximumInstanceCount,omitempty" xml:"MaximumInstanceCount,omitempty"`
	MinimumInstanceCount   *int32  `json:"MinimumInstanceCount,omitempty" xml:"MinimumInstanceCount,omitempty"`
	UpdateApplicationInput *string `json:"UpdateApplicationInput,omitempty" xml:"UpdateApplicationInput,omitempty"`
	UpdateTraffic          *bool   `json:"UpdateTraffic,omitempty" xml:"UpdateTraffic,omitempty"`
}

func (s BuildPipelineRunDeployConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunDeployConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunDeployConfig) GetAlwaysAllocateCPU() *bool {
	return s.AlwaysAllocateCPU
}

func (s *BuildPipelineRunDeployConfig) GetMaximumInstanceCount() *int32 {
	return s.MaximumInstanceCount
}

func (s *BuildPipelineRunDeployConfig) GetMinimumInstanceCount() *int32 {
	return s.MinimumInstanceCount
}

func (s *BuildPipelineRunDeployConfig) GetUpdateApplicationInput() *string {
	return s.UpdateApplicationInput
}

func (s *BuildPipelineRunDeployConfig) GetUpdateTraffic() *bool {
	return s.UpdateTraffic
}

func (s *BuildPipelineRunDeployConfig) SetAlwaysAllocateCPU(v bool) *BuildPipelineRunDeployConfig {
	s.AlwaysAllocateCPU = &v
	return s
}

func (s *BuildPipelineRunDeployConfig) SetMaximumInstanceCount(v int32) *BuildPipelineRunDeployConfig {
	s.MaximumInstanceCount = &v
	return s
}

func (s *BuildPipelineRunDeployConfig) SetMinimumInstanceCount(v int32) *BuildPipelineRunDeployConfig {
	s.MinimumInstanceCount = &v
	return s
}

func (s *BuildPipelineRunDeployConfig) SetUpdateApplicationInput(v string) *BuildPipelineRunDeployConfig {
	s.UpdateApplicationInput = &v
	return s
}

func (s *BuildPipelineRunDeployConfig) SetUpdateTraffic(v bool) *BuildPipelineRunDeployConfig {
	s.UpdateTraffic = &v
	return s
}

func (s *BuildPipelineRunDeployConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineRunImageConfig struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Repository   *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
}

func (s BuildPipelineRunImageConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunImageConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunImageConfig) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BuildPipelineRunImageConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *BuildPipelineRunImageConfig) GetRepository() *string {
	return s.Repository
}

func (s *BuildPipelineRunImageConfig) SetInstanceType(v string) *BuildPipelineRunImageConfig {
	s.InstanceType = &v
	return s
}

func (s *BuildPipelineRunImageConfig) SetNamespace(v string) *BuildPipelineRunImageConfig {
	s.Namespace = &v
	return s
}

func (s *BuildPipelineRunImageConfig) SetRepository(v string) *BuildPipelineRunImageConfig {
	s.Repository = &v
	return s
}

func (s *BuildPipelineRunImageConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineRunPackageConfig struct {
	PackageName    *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	PackageType    *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PackageUrl     *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
}

func (s BuildPipelineRunPackageConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunPackageConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunPackageConfig) GetPackageName() *string {
	return s.PackageName
}

func (s *BuildPipelineRunPackageConfig) GetPackageType() *string {
	return s.PackageType
}

func (s *BuildPipelineRunPackageConfig) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *BuildPipelineRunPackageConfig) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *BuildPipelineRunPackageConfig) SetPackageName(v string) *BuildPipelineRunPackageConfig {
	s.PackageName = &v
	return s
}

func (s *BuildPipelineRunPackageConfig) SetPackageType(v string) *BuildPipelineRunPackageConfig {
	s.PackageType = &v
	return s
}

func (s *BuildPipelineRunPackageConfig) SetPackageUrl(v string) *BuildPipelineRunPackageConfig {
	s.PackageUrl = &v
	return s
}

func (s *BuildPipelineRunPackageConfig) SetPackageVersion(v string) *BuildPipelineRunPackageConfig {
	s.PackageVersion = &v
	return s
}

func (s *BuildPipelineRunPackageConfig) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineRunSteps struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Result      *string `json:"Result,omitempty" xml:"Result,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s BuildPipelineRunSteps) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunSteps) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunSteps) GetDescription() *string {
	return s.Description
}

func (s *BuildPipelineRunSteps) GetDuration() *int64 {
	return s.Duration
}

func (s *BuildPipelineRunSteps) GetEndTime() *int64 {
	return s.EndTime
}

func (s *BuildPipelineRunSteps) GetId() *string {
	return s.Id
}

func (s *BuildPipelineRunSteps) GetName() *string {
	return s.Name
}

func (s *BuildPipelineRunSteps) GetResult() *string {
	return s.Result
}

func (s *BuildPipelineRunSteps) GetStartTime() *int64 {
	return s.StartTime
}

func (s *BuildPipelineRunSteps) GetStatus() *string {
	return s.Status
}

func (s *BuildPipelineRunSteps) SetDescription(v string) *BuildPipelineRunSteps {
	s.Description = &v
	return s
}

func (s *BuildPipelineRunSteps) SetDuration(v int64) *BuildPipelineRunSteps {
	s.Duration = &v
	return s
}

func (s *BuildPipelineRunSteps) SetEndTime(v int64) *BuildPipelineRunSteps {
	s.EndTime = &v
	return s
}

func (s *BuildPipelineRunSteps) SetId(v string) *BuildPipelineRunSteps {
	s.Id = &v
	return s
}

func (s *BuildPipelineRunSteps) SetName(v string) *BuildPipelineRunSteps {
	s.Name = &v
	return s
}

func (s *BuildPipelineRunSteps) SetResult(v string) *BuildPipelineRunSteps {
	s.Result = &v
	return s
}

func (s *BuildPipelineRunSteps) SetStartTime(v int64) *BuildPipelineRunSteps {
	s.StartTime = &v
	return s
}

func (s *BuildPipelineRunSteps) SetStatus(v string) *BuildPipelineRunSteps {
	s.Status = &v
	return s
}

func (s *BuildPipelineRunSteps) Validate() error {
	return dara.Validate(s)
}

type BuildPipelineRunTriggerConfig struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	TagName    *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BuildPipelineRunTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunTriggerConfig) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunTriggerConfig) GetBranchName() *string {
	return s.BranchName
}

func (s *BuildPipelineRunTriggerConfig) GetTagName() *string {
	return s.TagName
}

func (s *BuildPipelineRunTriggerConfig) GetType() *string {
	return s.Type
}

func (s *BuildPipelineRunTriggerConfig) SetBranchName(v string) *BuildPipelineRunTriggerConfig {
	s.BranchName = &v
	return s
}

func (s *BuildPipelineRunTriggerConfig) SetTagName(v string) *BuildPipelineRunTriggerConfig {
	s.TagName = &v
	return s
}

func (s *BuildPipelineRunTriggerConfig) SetType(v string) *BuildPipelineRunTriggerConfig {
	s.Type = &v
	return s
}

func (s *BuildPipelineRunTriggerConfig) Validate() error {
	return dara.Validate(s)
}
