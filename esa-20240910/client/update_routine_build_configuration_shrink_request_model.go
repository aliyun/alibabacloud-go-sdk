// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineBuildConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsDirectory(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetAssetsDirectory() *string
	SetBuildBranches(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetBuildBranches() *string
	SetBuildCommand(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetBuildCommand() *string
	SetEnvironmentVariablesShrink(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetEnvironmentVariablesShrink() *string
	SetGitAccountId(v int64) *UpdateRoutineBuildConfigurationShrinkRequest
	GetGitAccountId() *int64
	SetGitPlatform(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetGitPlatform() *string
	SetInstallCommand(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetInstallCommand() *string
	SetIsPrivate(v bool) *UpdateRoutineBuildConfigurationShrinkRequest
	GetIsPrivate() *bool
	SetNodeVersion(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetNodeVersion() *string
	SetProductionBranch(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetProductionBranch() *string
	SetRepository(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetRepository() *string
	SetRootDirectory(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetRootDirectory() *string
	SetRoutineEntry(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetRoutineEntry() *string
	SetRoutineName(v string) *UpdateRoutineBuildConfigurationShrinkRequest
	GetRoutineName() *string
}

type UpdateRoutineBuildConfigurationShrinkRequest struct {
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
	EnvironmentVariablesShrink *string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
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

func (s UpdateRoutineBuildConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineBuildConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetAssetsDirectory() *string {
	return s.AssetsDirectory
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetBuildBranches() *string {
	return s.BuildBranches
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetBuildCommand() *string {
	return s.BuildCommand
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetEnvironmentVariablesShrink() *string {
	return s.EnvironmentVariablesShrink
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetGitAccountId() *int64 {
	return s.GitAccountId
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetGitPlatform() *string {
	return s.GitPlatform
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetNodeVersion() *string {
	return s.NodeVersion
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetProductionBranch() *string {
	return s.ProductionBranch
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetRepository() *string {
	return s.Repository
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetRoutineEntry() *string {
	return s.RoutineEntry
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetAssetsDirectory(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.AssetsDirectory = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetBuildBranches(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.BuildBranches = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetBuildCommand(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.BuildCommand = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetEnvironmentVariablesShrink(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.EnvironmentVariablesShrink = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetGitAccountId(v int64) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.GitAccountId = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetGitPlatform(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.GitPlatform = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetInstallCommand(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.InstallCommand = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetIsPrivate(v bool) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.IsPrivate = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetNodeVersion(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.NodeVersion = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetProductionBranch(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.ProductionBranch = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetRepository(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.Repository = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetRootDirectory(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.RootDirectory = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetRoutineEntry(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.RoutineEntry = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) SetRoutineName(v string) *UpdateRoutineBuildConfigurationShrinkRequest {
	s.RoutineName = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
