// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineBuildConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsDirectory(v string) *UpdateRoutineBuildConfigurationRequest
	GetAssetsDirectory() *string
	SetBuildBranches(v string) *UpdateRoutineBuildConfigurationRequest
	GetBuildBranches() *string
	SetBuildCommand(v string) *UpdateRoutineBuildConfigurationRequest
	GetBuildCommand() *string
	SetEnvironmentVariables(v map[string]*string) *UpdateRoutineBuildConfigurationRequest
	GetEnvironmentVariables() map[string]*string
	SetGitAccountId(v int64) *UpdateRoutineBuildConfigurationRequest
	GetGitAccountId() *int64
	SetGitPlatform(v string) *UpdateRoutineBuildConfigurationRequest
	GetGitPlatform() *string
	SetInstallCommand(v string) *UpdateRoutineBuildConfigurationRequest
	GetInstallCommand() *string
	SetIsPrivate(v bool) *UpdateRoutineBuildConfigurationRequest
	GetIsPrivate() *bool
	SetNodeVersion(v string) *UpdateRoutineBuildConfigurationRequest
	GetNodeVersion() *string
	SetProductionBranch(v string) *UpdateRoutineBuildConfigurationRequest
	GetProductionBranch() *string
	SetRepository(v string) *UpdateRoutineBuildConfigurationRequest
	GetRepository() *string
	SetRootDirectory(v string) *UpdateRoutineBuildConfigurationRequest
	GetRootDirectory() *string
	SetRoutineEntry(v string) *UpdateRoutineBuildConfigurationRequest
	GetRoutineEntry() *string
	SetRoutineName(v string) *UpdateRoutineBuildConfigurationRequest
	GetRoutineName() *string
}

type UpdateRoutineBuildConfigurationRequest struct {
	// The static resource directory.
	//
	// example:
	//
	// /root/user
	AssetsDirectory *string `json:"AssetsDirectory,omitempty" xml:"AssetsDirectory,omitempty"`
	// The branches that trigger a build. Set this to 	- for all branches. To specify multiple branches, separate branch names with commas.
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
	// The environment variables.
	//
	// example:
	//
	// 100
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The Git account ID.
	//
	// example:
	//
	// 4580717755793600
	GitAccountId *int64 `json:"GitAccountId,omitempty" xml:"GitAccountId,omitempty"`
	// The Git platform. Valid values: github, gitee, and upload.
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
	// Specifies whether the repository is private. Valid values:
	//
	// - `true`: The repository is private.
	//
	// - `false`: The repository is not private.
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
	// This parameter is required.
	//
	// example:
	//
	// test-routine
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
}

func (s UpdateRoutineBuildConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineBuildConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoutineBuildConfigurationRequest) GetAssetsDirectory() *string {
	return s.AssetsDirectory
}

func (s *UpdateRoutineBuildConfigurationRequest) GetBuildBranches() *string {
	return s.BuildBranches
}

func (s *UpdateRoutineBuildConfigurationRequest) GetBuildCommand() *string {
	return s.BuildCommand
}

func (s *UpdateRoutineBuildConfigurationRequest) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateRoutineBuildConfigurationRequest) GetGitAccountId() *int64 {
	return s.GitAccountId
}

func (s *UpdateRoutineBuildConfigurationRequest) GetGitPlatform() *string {
	return s.GitPlatform
}

func (s *UpdateRoutineBuildConfigurationRequest) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *UpdateRoutineBuildConfigurationRequest) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *UpdateRoutineBuildConfigurationRequest) GetNodeVersion() *string {
	return s.NodeVersion
}

func (s *UpdateRoutineBuildConfigurationRequest) GetProductionBranch() *string {
	return s.ProductionBranch
}

func (s *UpdateRoutineBuildConfigurationRequest) GetRepository() *string {
	return s.Repository
}

func (s *UpdateRoutineBuildConfigurationRequest) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *UpdateRoutineBuildConfigurationRequest) GetRoutineEntry() *string {
	return s.RoutineEntry
}

func (s *UpdateRoutineBuildConfigurationRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *UpdateRoutineBuildConfigurationRequest) SetAssetsDirectory(v string) *UpdateRoutineBuildConfigurationRequest {
	s.AssetsDirectory = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetBuildBranches(v string) *UpdateRoutineBuildConfigurationRequest {
	s.BuildBranches = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetBuildCommand(v string) *UpdateRoutineBuildConfigurationRequest {
	s.BuildCommand = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetEnvironmentVariables(v map[string]*string) *UpdateRoutineBuildConfigurationRequest {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetGitAccountId(v int64) *UpdateRoutineBuildConfigurationRequest {
	s.GitAccountId = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetGitPlatform(v string) *UpdateRoutineBuildConfigurationRequest {
	s.GitPlatform = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetInstallCommand(v string) *UpdateRoutineBuildConfigurationRequest {
	s.InstallCommand = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetIsPrivate(v bool) *UpdateRoutineBuildConfigurationRequest {
	s.IsPrivate = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetNodeVersion(v string) *UpdateRoutineBuildConfigurationRequest {
	s.NodeVersion = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetProductionBranch(v string) *UpdateRoutineBuildConfigurationRequest {
	s.ProductionBranch = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetRepository(v string) *UpdateRoutineBuildConfigurationRequest {
	s.Repository = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetRootDirectory(v string) *UpdateRoutineBuildConfigurationRequest {
	s.RootDirectory = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetRoutineEntry(v string) *UpdateRoutineBuildConfigurationRequest {
	s.RoutineEntry = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) SetRoutineName(v string) *UpdateRoutineBuildConfigurationRequest {
	s.RoutineName = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
