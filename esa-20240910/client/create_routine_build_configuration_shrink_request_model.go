// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineBuildConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsDirectory(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetAssetsDirectory() *string
	SetBuildBranches(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetBuildBranches() *string
	SetBuildCommand(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetBuildCommand() *string
	SetEnvironmentVariablesShrink(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetEnvironmentVariablesShrink() *string
	SetGitAccountId(v int64) *CreateRoutineBuildConfigurationShrinkRequest
	GetGitAccountId() *int64
	SetGitPlatform(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetGitPlatform() *string
	SetInstallCommand(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetInstallCommand() *string
	SetIsPrivate(v bool) *CreateRoutineBuildConfigurationShrinkRequest
	GetIsPrivate() *bool
	SetNodeVersion(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetNodeVersion() *string
	SetProductionBranch(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetProductionBranch() *string
	SetRepository(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetRepository() *string
	SetRootDirectory(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetRootDirectory() *string
	SetRoutineEntry(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetRoutineEntry() *string
	SetRoutineName(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetRoutineName() *string
	SetTemplateName(v string) *CreateRoutineBuildConfigurationShrinkRequest
	GetTemplateName() *string
}

type CreateRoutineBuildConfigurationShrinkRequest struct {
	// The static resource directory.
	//
	// example:
	//
	// /root/user
	AssetsDirectory *string `json:"AssetsDirectory,omitempty" xml:"AssetsDirectory,omitempty"`
	// The branches that trigger a build. Set this parameter to 	- for all branches. To specify multiple branches, separate branch names with commas.
	//
	// example:
	//
	// feature/test,hotfix/test
	BuildBranches *string `json:"BuildBranches,omitempty" xml:"BuildBranches,omitempty"`
	// The build command.
	//
	// example:
	//
	// npm run build
	BuildCommand *string `json:"BuildCommand,omitempty" xml:"BuildCommand,omitempty"`
	// The environment variables.
	EnvironmentVariablesShrink *string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The Git account ID.
	//
	// example:
	//
	// 3472021274759488
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
	// This parameter is required.
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
	// webdeck
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
	// rwa-test
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
	// The build template name.
	//
	// example:
	//
	// react-router
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateRoutineBuildConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineBuildConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetAssetsDirectory() *string {
	return s.AssetsDirectory
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetBuildBranches() *string {
	return s.BuildBranches
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetBuildCommand() *string {
	return s.BuildCommand
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetEnvironmentVariablesShrink() *string {
	return s.EnvironmentVariablesShrink
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetGitAccountId() *int64 {
	return s.GitAccountId
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetGitPlatform() *string {
	return s.GitPlatform
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetNodeVersion() *string {
	return s.NodeVersion
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetProductionBranch() *string {
	return s.ProductionBranch
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetRepository() *string {
	return s.Repository
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetRoutineEntry() *string {
	return s.RoutineEntry
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetAssetsDirectory(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.AssetsDirectory = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetBuildBranches(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.BuildBranches = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetBuildCommand(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.BuildCommand = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetEnvironmentVariablesShrink(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.EnvironmentVariablesShrink = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetGitAccountId(v int64) *CreateRoutineBuildConfigurationShrinkRequest {
	s.GitAccountId = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetGitPlatform(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.GitPlatform = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetInstallCommand(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.InstallCommand = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetIsPrivate(v bool) *CreateRoutineBuildConfigurationShrinkRequest {
	s.IsPrivate = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetNodeVersion(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.NodeVersion = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetProductionBranch(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.ProductionBranch = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetRepository(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.Repository = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetRootDirectory(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.RootDirectory = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetRoutineEntry(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.RoutineEntry = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetRoutineName(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.RoutineName = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) SetTemplateName(v string) *CreateRoutineBuildConfigurationShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateRoutineBuildConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
