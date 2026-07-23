// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineBuildConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsDirectory(v string) *CreateRoutineBuildConfigurationRequest
	GetAssetsDirectory() *string
	SetBuildBranches(v string) *CreateRoutineBuildConfigurationRequest
	GetBuildBranches() *string
	SetBuildCommand(v string) *CreateRoutineBuildConfigurationRequest
	GetBuildCommand() *string
	SetEnvironmentVariables(v map[string]*string) *CreateRoutineBuildConfigurationRequest
	GetEnvironmentVariables() map[string]*string
	SetGitAccountId(v int64) *CreateRoutineBuildConfigurationRequest
	GetGitAccountId() *int64
	SetGitPlatform(v string) *CreateRoutineBuildConfigurationRequest
	GetGitPlatform() *string
	SetInstallCommand(v string) *CreateRoutineBuildConfigurationRequest
	GetInstallCommand() *string
	SetIsPrivate(v bool) *CreateRoutineBuildConfigurationRequest
	GetIsPrivate() *bool
	SetNodeVersion(v string) *CreateRoutineBuildConfigurationRequest
	GetNodeVersion() *string
	SetProductionBranch(v string) *CreateRoutineBuildConfigurationRequest
	GetProductionBranch() *string
	SetRepository(v string) *CreateRoutineBuildConfigurationRequest
	GetRepository() *string
	SetRootDirectory(v string) *CreateRoutineBuildConfigurationRequest
	GetRootDirectory() *string
	SetRoutineEntry(v string) *CreateRoutineBuildConfigurationRequest
	GetRoutineEntry() *string
	SetRoutineName(v string) *CreateRoutineBuildConfigurationRequest
	GetRoutineName() *string
	SetTemplateName(v string) *CreateRoutineBuildConfigurationRequest
	GetTemplateName() *string
}

type CreateRoutineBuildConfigurationRequest struct {
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
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
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

func (s CreateRoutineBuildConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineBuildConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineBuildConfigurationRequest) GetAssetsDirectory() *string {
	return s.AssetsDirectory
}

func (s *CreateRoutineBuildConfigurationRequest) GetBuildBranches() *string {
	return s.BuildBranches
}

func (s *CreateRoutineBuildConfigurationRequest) GetBuildCommand() *string {
	return s.BuildCommand
}

func (s *CreateRoutineBuildConfigurationRequest) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateRoutineBuildConfigurationRequest) GetGitAccountId() *int64 {
	return s.GitAccountId
}

func (s *CreateRoutineBuildConfigurationRequest) GetGitPlatform() *string {
	return s.GitPlatform
}

func (s *CreateRoutineBuildConfigurationRequest) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *CreateRoutineBuildConfigurationRequest) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *CreateRoutineBuildConfigurationRequest) GetNodeVersion() *string {
	return s.NodeVersion
}

func (s *CreateRoutineBuildConfigurationRequest) GetProductionBranch() *string {
	return s.ProductionBranch
}

func (s *CreateRoutineBuildConfigurationRequest) GetRepository() *string {
	return s.Repository
}

func (s *CreateRoutineBuildConfigurationRequest) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *CreateRoutineBuildConfigurationRequest) GetRoutineEntry() *string {
	return s.RoutineEntry
}

func (s *CreateRoutineBuildConfigurationRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *CreateRoutineBuildConfigurationRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateRoutineBuildConfigurationRequest) SetAssetsDirectory(v string) *CreateRoutineBuildConfigurationRequest {
	s.AssetsDirectory = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetBuildBranches(v string) *CreateRoutineBuildConfigurationRequest {
	s.BuildBranches = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetBuildCommand(v string) *CreateRoutineBuildConfigurationRequest {
	s.BuildCommand = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetEnvironmentVariables(v map[string]*string) *CreateRoutineBuildConfigurationRequest {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetGitAccountId(v int64) *CreateRoutineBuildConfigurationRequest {
	s.GitAccountId = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetGitPlatform(v string) *CreateRoutineBuildConfigurationRequest {
	s.GitPlatform = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetInstallCommand(v string) *CreateRoutineBuildConfigurationRequest {
	s.InstallCommand = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetIsPrivate(v bool) *CreateRoutineBuildConfigurationRequest {
	s.IsPrivate = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetNodeVersion(v string) *CreateRoutineBuildConfigurationRequest {
	s.NodeVersion = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetProductionBranch(v string) *CreateRoutineBuildConfigurationRequest {
	s.ProductionBranch = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetRepository(v string) *CreateRoutineBuildConfigurationRequest {
	s.Repository = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetRootDirectory(v string) *CreateRoutineBuildConfigurationRequest {
	s.RootDirectory = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetRoutineEntry(v string) *CreateRoutineBuildConfigurationRequest {
	s.RoutineEntry = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetRoutineName(v string) *CreateRoutineBuildConfigurationRequest {
	s.RoutineName = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) SetTemplateName(v string) *CreateRoutineBuildConfigurationRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateRoutineBuildConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
