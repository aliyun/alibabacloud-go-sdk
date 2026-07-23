// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineBuildResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsDirectory(v string) *GetRoutineBuildResponseBody
	GetAssetsDirectory() *string
	SetBranch(v string) *GetRoutineBuildResponseBody
	GetBranch() *string
	SetBuildCommand(v string) *GetRoutineBuildResponseBody
	GetBuildCommand() *string
	SetCommitId(v string) *GetRoutineBuildResponseBody
	GetCommitId() *string
	SetCommitMessage(v string) *GetRoutineBuildResponseBody
	GetCommitMessage() *string
	SetCreateTime(v string) *GetRoutineBuildResponseBody
	GetCreateTime() *string
	SetEnvironmentVariables(v map[string]*string) *GetRoutineBuildResponseBody
	GetEnvironmentVariables() map[string]*string
	SetGitAccountId(v int64) *GetRoutineBuildResponseBody
	GetGitAccountId() *int64
	SetId(v int64) *GetRoutineBuildResponseBody
	GetId() *int64
	SetInstallCommand(v string) *GetRoutineBuildResponseBody
	GetInstallCommand() *string
	SetIsPrivate(v bool) *GetRoutineBuildResponseBody
	GetIsPrivate() *bool
	SetNodeVersion(v string) *GetRoutineBuildResponseBody
	GetNodeVersion() *string
	SetPipelineId(v int64) *GetRoutineBuildResponseBody
	GetPipelineId() *int64
	SetPipelineRunId(v int64) *GetRoutineBuildResponseBody
	GetPipelineRunId() *int64
	SetProductionBranch(v string) *GetRoutineBuildResponseBody
	GetProductionBranch() *string
	SetRepository(v string) *GetRoutineBuildResponseBody
	GetRepository() *string
	SetRequestId(v string) *GetRoutineBuildResponseBody
	GetRequestId() *string
	SetRootDirectory(v string) *GetRoutineBuildResponseBody
	GetRootDirectory() *string
	SetRoutineEntry(v string) *GetRoutineBuildResponseBody
	GetRoutineEntry() *string
	SetRoutineName(v string) *GetRoutineBuildResponseBody
	GetRoutineName() *string
	SetStatus(v string) *GetRoutineBuildResponseBody
	GetStatus() *string
	SetTemplateName(v string) *GetRoutineBuildResponseBody
	GetTemplateName() *string
	SetUpdateTime(v string) *GetRoutineBuildResponseBody
	GetUpdateTime() *string
}

type GetRoutineBuildResponseBody struct {
	// The static resource directory.
	//
	// example:
	//
	// /root/user
	AssetsDirectory *string `json:"AssetsDirectory,omitempty" xml:"AssetsDirectory,omitempty"`
	// The build branch.
	//
	// example:
	//
	// dev
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The build command.
	//
	// example:
	//
	// npm run build
	BuildCommand *string `json:"BuildCommand,omitempty" xml:"BuildCommand,omitempty"`
	// The commit ID.
	//
	// example:
	//
	// 9bf55641a1a608b9e7297d3fe51e39baa4b68ba0
	CommitId *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	// The commit message.
	//
	// example:
	//
	// Add configuration file.
	CommitMessage *string `json:"CommitMessage,omitempty" xml:"CommitMessage,omitempty"`
	// The creation time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-02-28T09:03:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The environment variables.
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The Git account ID.
	//
	// example:
	//
	// 4580717755793600
	GitAccountId *int64 `json:"GitAccountId,omitempty" xml:"GitAccountId,omitempty"`
	// The ID of the ER build task.
	//
	// example:
	//
	// 164557372123356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The install command.
	//
	// example:
	//
	// npm install
	InstallCommand *string `json:"InstallCommand,omitempty" xml:"InstallCommand,omitempty"`
	// Indicates whether the repository is private. Valid values:
	//
	// - true: The repository is private.
	//
	// - false: The repository is not private.
	//
	// example:
	//
	// false
	IsPrivate *bool `json:"IsPrivate,omitempty" xml:"IsPrivate,omitempty"`
	// The Node.js version. Valid values: `22.x`, `20.x`, `18.x`, `16.x`, `14.x`, `12.x`.
	//
	// example:
	//
	// 22.x
	NodeVersion *string `json:"NodeVersion,omitempty" xml:"NodeVersion,omitempty"`
	// The pipeline ID.
	//
	// example:
	//
	// 4371588
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The pipeline run ID.
	//
	// example:
	//
	// 70
	PipelineRunId *int64 `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	// The production branch name.
	//
	// example:
	//
	// main
	ProductionBranch *string `json:"ProductionBranch,omitempty" xml:"ProductionBranch,omitempty"`
	// The repository name.
	//
	// example:
	//
	// example-test
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C3CC8AF-7C4C-5841-BDAE-B295FD9AE913
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The root directory.
	//
	// example:
	//
	// /root/admin
	RootDirectory *string `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
	// The ER entry file path.
	//
	// example:
	//
	// /home
	RoutineEntry *string `json:"RoutineEntry,omitempty" xml:"RoutineEntry,omitempty"`
	// The ER name.
	//
	// example:
	//
	// test-routine
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
	// The status of the build task. Valid values:
	//
	// - int: init
	//
	// - pending: preparing
	//
	// - building: building
	//
	// - succeed: build succeeded
	//
	// - failed: build failed
	//
	// - canceled: canceled
	//
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template name.
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The modification time. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-07-20T09:59:28+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetRoutineBuildResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineBuildResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineBuildResponseBody) GetAssetsDirectory() *string {
	return s.AssetsDirectory
}

func (s *GetRoutineBuildResponseBody) GetBranch() *string {
	return s.Branch
}

func (s *GetRoutineBuildResponseBody) GetBuildCommand() *string {
	return s.BuildCommand
}

func (s *GetRoutineBuildResponseBody) GetCommitId() *string {
	return s.CommitId
}

func (s *GetRoutineBuildResponseBody) GetCommitMessage() *string {
	return s.CommitMessage
}

func (s *GetRoutineBuildResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRoutineBuildResponseBody) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *GetRoutineBuildResponseBody) GetGitAccountId() *int64 {
	return s.GitAccountId
}

func (s *GetRoutineBuildResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetRoutineBuildResponseBody) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *GetRoutineBuildResponseBody) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *GetRoutineBuildResponseBody) GetNodeVersion() *string {
	return s.NodeVersion
}

func (s *GetRoutineBuildResponseBody) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *GetRoutineBuildResponseBody) GetPipelineRunId() *int64 {
	return s.PipelineRunId
}

func (s *GetRoutineBuildResponseBody) GetProductionBranch() *string {
	return s.ProductionBranch
}

func (s *GetRoutineBuildResponseBody) GetRepository() *string {
	return s.Repository
}

func (s *GetRoutineBuildResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineBuildResponseBody) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *GetRoutineBuildResponseBody) GetRoutineEntry() *string {
	return s.RoutineEntry
}

func (s *GetRoutineBuildResponseBody) GetRoutineName() *string {
	return s.RoutineName
}

func (s *GetRoutineBuildResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRoutineBuildResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetRoutineBuildResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetRoutineBuildResponseBody) SetAssetsDirectory(v string) *GetRoutineBuildResponseBody {
	s.AssetsDirectory = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetBranch(v string) *GetRoutineBuildResponseBody {
	s.Branch = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetBuildCommand(v string) *GetRoutineBuildResponseBody {
	s.BuildCommand = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetCommitId(v string) *GetRoutineBuildResponseBody {
	s.CommitId = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetCommitMessage(v string) *GetRoutineBuildResponseBody {
	s.CommitMessage = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetCreateTime(v string) *GetRoutineBuildResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetEnvironmentVariables(v map[string]*string) *GetRoutineBuildResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *GetRoutineBuildResponseBody) SetGitAccountId(v int64) *GetRoutineBuildResponseBody {
	s.GitAccountId = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetId(v int64) *GetRoutineBuildResponseBody {
	s.Id = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetInstallCommand(v string) *GetRoutineBuildResponseBody {
	s.InstallCommand = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetIsPrivate(v bool) *GetRoutineBuildResponseBody {
	s.IsPrivate = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetNodeVersion(v string) *GetRoutineBuildResponseBody {
	s.NodeVersion = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetPipelineId(v int64) *GetRoutineBuildResponseBody {
	s.PipelineId = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetPipelineRunId(v int64) *GetRoutineBuildResponseBody {
	s.PipelineRunId = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetProductionBranch(v string) *GetRoutineBuildResponseBody {
	s.ProductionBranch = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetRepository(v string) *GetRoutineBuildResponseBody {
	s.Repository = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetRequestId(v string) *GetRoutineBuildResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetRootDirectory(v string) *GetRoutineBuildResponseBody {
	s.RootDirectory = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetRoutineEntry(v string) *GetRoutineBuildResponseBody {
	s.RoutineEntry = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetRoutineName(v string) *GetRoutineBuildResponseBody {
	s.RoutineName = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetStatus(v string) *GetRoutineBuildResponseBody {
	s.Status = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetTemplateName(v string) *GetRoutineBuildResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetRoutineBuildResponseBody) SetUpdateTime(v string) *GetRoutineBuildResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetRoutineBuildResponseBody) Validate() error {
	return dara.Validate(s)
}
