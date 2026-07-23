// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineBuildConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsDirectory(v string) *GetRoutineBuildConfigurationResponseBody
	GetAssetsDirectory() *string
	SetBuildBranches(v string) *GetRoutineBuildConfigurationResponseBody
	GetBuildBranches() *string
	SetBuildCommand(v string) *GetRoutineBuildConfigurationResponseBody
	GetBuildCommand() *string
	SetCreateTime(v string) *GetRoutineBuildConfigurationResponseBody
	GetCreateTime() *string
	SetEnvironmentVariables(v map[string]*string) *GetRoutineBuildConfigurationResponseBody
	GetEnvironmentVariables() map[string]*string
	SetGitAccountId(v int64) *GetRoutineBuildConfigurationResponseBody
	GetGitAccountId() *int64
	SetGitAccountType(v string) *GetRoutineBuildConfigurationResponseBody
	GetGitAccountType() *string
	SetGitPlatform(v string) *GetRoutineBuildConfigurationResponseBody
	GetGitPlatform() *string
	SetInstallCommand(v string) *GetRoutineBuildConfigurationResponseBody
	GetInstallCommand() *string
	SetIsPrivate(v bool) *GetRoutineBuildConfigurationResponseBody
	GetIsPrivate() *bool
	SetNodeVersion(v string) *GetRoutineBuildConfigurationResponseBody
	GetNodeVersion() *string
	SetProductionBranch(v string) *GetRoutineBuildConfigurationResponseBody
	GetProductionBranch() *string
	SetRepository(v string) *GetRoutineBuildConfigurationResponseBody
	GetRepository() *string
	SetRequestId(v string) *GetRoutineBuildConfigurationResponseBody
	GetRequestId() *string
	SetRootDirectory(v string) *GetRoutineBuildConfigurationResponseBody
	GetRootDirectory() *string
	SetRoutineBuildConfigurationId(v int64) *GetRoutineBuildConfigurationResponseBody
	GetRoutineBuildConfigurationId() *int64
	SetRoutineEntry(v string) *GetRoutineBuildConfigurationResponseBody
	GetRoutineEntry() *string
	SetRoutineName(v string) *GetRoutineBuildConfigurationResponseBody
	GetRoutineName() *string
	SetUpdateTime(v string) *GetRoutineBuildConfigurationResponseBody
	GetUpdateTime() *string
}

type GetRoutineBuildConfigurationResponseBody struct {
	// The static assets directory.
	//
	// example:
	//
	// /root/usr
	AssetsDirectory *string `json:"AssetsDirectory,omitempty" xml:"AssetsDirectory,omitempty"`
	// The branches that trigger builds. A value of 	- indicates all branches. Multiple specific branches are separated by commas.
	//
	// example:
	//
	// int,abc
	BuildBranches *string `json:"BuildBranches,omitempty" xml:"BuildBranches,omitempty"`
	// The build command.
	//
	// example:
	//
	// npm run build
	BuildCommand *string `json:"BuildCommand,omitempty" xml:"BuildCommand,omitempty"`
	// The creation time, in ISO 8601 format using UTC time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-03-11T01:23:21Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The environment variables.
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The Git account ID.
	//
	// example:
	//
	// 4695144764942144
	GitAccountId *int64 `json:"GitAccountId,omitempty" xml:"GitAccountId,omitempty"`
	// The Git account type. Valid values:
	//
	// - User: individual account.
	//
	// - Organization: organization account.
	//
	// example:
	//
	// User
	GitAccountType *string `json:"GitAccountType,omitempty" xml:"GitAccountType,omitempty"`
	// The Git platform.
	//
	// example:
	//
	// github
	GitPlatform *string `json:"GitPlatform,omitempty" xml:"GitPlatform,omitempty"`
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
	// The Node.js version.
	//
	// example:
	//
	// 22.x
	NodeVersion *string `json:"NodeVersion,omitempty" xml:"NodeVersion,omitempty"`
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
	// D1D7BBB5-9B5B-5A29-8848-398F3CA18A8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The root directory.
	//
	// example:
	//
	// /root/admin
	RootDirectory *string `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
	// The ER build configuration ID.
	//
	// example:
	//
	// 3472165674357056
	RoutineBuildConfigurationId *int64 `json:"RoutineBuildConfigurationId,omitempty" xml:"RoutineBuildConfigurationId,omitempty"`
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
	// The modification time, in ISO 8601 format using UTC time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-04-19T11:15:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetRoutineBuildConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineBuildConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineBuildConfigurationResponseBody) GetAssetsDirectory() *string {
	return s.AssetsDirectory
}

func (s *GetRoutineBuildConfigurationResponseBody) GetBuildBranches() *string {
	return s.BuildBranches
}

func (s *GetRoutineBuildConfigurationResponseBody) GetBuildCommand() *string {
	return s.BuildCommand
}

func (s *GetRoutineBuildConfigurationResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRoutineBuildConfigurationResponseBody) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *GetRoutineBuildConfigurationResponseBody) GetGitAccountId() *int64 {
	return s.GitAccountId
}

func (s *GetRoutineBuildConfigurationResponseBody) GetGitAccountType() *string {
	return s.GitAccountType
}

func (s *GetRoutineBuildConfigurationResponseBody) GetGitPlatform() *string {
	return s.GitPlatform
}

func (s *GetRoutineBuildConfigurationResponseBody) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *GetRoutineBuildConfigurationResponseBody) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *GetRoutineBuildConfigurationResponseBody) GetNodeVersion() *string {
	return s.NodeVersion
}

func (s *GetRoutineBuildConfigurationResponseBody) GetProductionBranch() *string {
	return s.ProductionBranch
}

func (s *GetRoutineBuildConfigurationResponseBody) GetRepository() *string {
	return s.Repository
}

func (s *GetRoutineBuildConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineBuildConfigurationResponseBody) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *GetRoutineBuildConfigurationResponseBody) GetRoutineBuildConfigurationId() *int64 {
	return s.RoutineBuildConfigurationId
}

func (s *GetRoutineBuildConfigurationResponseBody) GetRoutineEntry() *string {
	return s.RoutineEntry
}

func (s *GetRoutineBuildConfigurationResponseBody) GetRoutineName() *string {
	return s.RoutineName
}

func (s *GetRoutineBuildConfigurationResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetRoutineBuildConfigurationResponseBody) SetAssetsDirectory(v string) *GetRoutineBuildConfigurationResponseBody {
	s.AssetsDirectory = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetBuildBranches(v string) *GetRoutineBuildConfigurationResponseBody {
	s.BuildBranches = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetBuildCommand(v string) *GetRoutineBuildConfigurationResponseBody {
	s.BuildCommand = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetCreateTime(v string) *GetRoutineBuildConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetEnvironmentVariables(v map[string]*string) *GetRoutineBuildConfigurationResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetGitAccountId(v int64) *GetRoutineBuildConfigurationResponseBody {
	s.GitAccountId = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetGitAccountType(v string) *GetRoutineBuildConfigurationResponseBody {
	s.GitAccountType = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetGitPlatform(v string) *GetRoutineBuildConfigurationResponseBody {
	s.GitPlatform = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetInstallCommand(v string) *GetRoutineBuildConfigurationResponseBody {
	s.InstallCommand = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetIsPrivate(v bool) *GetRoutineBuildConfigurationResponseBody {
	s.IsPrivate = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetNodeVersion(v string) *GetRoutineBuildConfigurationResponseBody {
	s.NodeVersion = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetProductionBranch(v string) *GetRoutineBuildConfigurationResponseBody {
	s.ProductionBranch = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetRepository(v string) *GetRoutineBuildConfigurationResponseBody {
	s.Repository = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetRequestId(v string) *GetRoutineBuildConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetRootDirectory(v string) *GetRoutineBuildConfigurationResponseBody {
	s.RootDirectory = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetRoutineBuildConfigurationId(v int64) *GetRoutineBuildConfigurationResponseBody {
	s.RoutineBuildConfigurationId = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetRoutineEntry(v string) *GetRoutineBuildConfigurationResponseBody {
	s.RoutineEntry = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetRoutineName(v string) *GetRoutineBuildConfigurationResponseBody {
	s.RoutineName = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) SetUpdateTime(v string) *GetRoutineBuildConfigurationResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
