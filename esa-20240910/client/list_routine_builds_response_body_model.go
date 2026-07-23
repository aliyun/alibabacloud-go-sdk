// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineBuildsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListRoutineBuildsResponseBodyData) *ListRoutineBuildsResponseBody
	GetData() []*ListRoutineBuildsResponseBodyData
	SetPageIndex(v int64) *ListRoutineBuildsResponseBody
	GetPageIndex() *int64
	SetPageSize(v int64) *ListRoutineBuildsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListRoutineBuildsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRoutineBuildsResponseBody
	GetTotalCount() *int64
	SetTotalPage(v int64) *ListRoutineBuildsResponseBody
	GetTotalPage() *int64
}

type ListRoutineBuildsResponseBody struct {
	// The list of ER build tasks.
	Data []*ListRoutineBuildsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number, same as the PageIndex request parameter.
	//
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A528000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListRoutineBuildsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildsResponseBody) GetData() []*ListRoutineBuildsResponseBodyData {
	return s.Data
}

func (s *ListRoutineBuildsResponseBody) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *ListRoutineBuildsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRoutineBuildsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutineBuildsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRoutineBuildsResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *ListRoutineBuildsResponseBody) SetData(v []*ListRoutineBuildsResponseBodyData) *ListRoutineBuildsResponseBody {
	s.Data = v
	return s
}

func (s *ListRoutineBuildsResponseBody) SetPageIndex(v int64) *ListRoutineBuildsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListRoutineBuildsResponseBody) SetPageSize(v int64) *ListRoutineBuildsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRoutineBuildsResponseBody) SetRequestId(v string) *ListRoutineBuildsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutineBuildsResponseBody) SetTotalCount(v int64) *ListRoutineBuildsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRoutineBuildsResponseBody) SetTotalPage(v int64) *ListRoutineBuildsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListRoutineBuildsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRoutineBuildsResponseBodyData struct {
	// The static assets directory.
	//
	// example:
	//
	// /root/usr
	AssetsDirectory *string `json:"AssetsDirectory,omitempty" xml:"AssetsDirectory,omitempty"`
	// The branch used for the build.
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
	// The ID of the commit.
	//
	// example:
	//
	// c08057f590f8d5be56fcae1e082128254a708f94
	CommitId *string `json:"CommitId,omitempty" xml:"CommitId,omitempty"`
	// The commit message.
	//
	// example:
	//
	// Add static files.
	CommitMessage *string `json:"CommitMessage,omitempty" xml:"CommitMessage,omitempty"`
	// The creation time, in ISO 8601 format using UTC time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-03-26T02:19:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The environment variables.
	EnvironmentVariables map[string]*string `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The Git account ID.
	//
	// example:
	//
	// 162124764031208
	GitAccountId *int64 `json:"GitAccountId,omitempty" xml:"GitAccountId,omitempty"`
	// The Git account name.
	//
	// example:
	//
	// rwa
	GitAccountName *string `json:"GitAccountName,omitempty" xml:"GitAccountName,omitempty"`
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
	// The Node.js version. Valid values: `22.x`, `20.x`, `18.x`, `16.x`, `14.x`, and `12.x`.
	//
	// example:
	//
	// 22.x
	NodeVersion *string `json:"NodeVersion,omitempty" xml:"NodeVersion,omitempty"`
	// The pipeline ID.
	//
	// example:
	//
	// 3850166
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The pipeline execution ID.
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
	// rwa-test
	Repository *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	// The root directory.
	//
	// example:
	//
	// /root/admin
	RootDirectory *string `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty"`
	// The ER build task ID.
	//
	// example:
	//
	// 4133325046294912
	RoutineBuildId *int64 `json:"RoutineBuildId,omitempty" xml:"RoutineBuildId,omitempty"`
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
	// - int: initialization
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
	// The modification time, in ISO 8601 format using UTC time. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-06-20T00:44:23Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1427812834792318
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRoutineBuildsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildsResponseBodyData) GetAssetsDirectory() *string {
	return s.AssetsDirectory
}

func (s *ListRoutineBuildsResponseBodyData) GetBranch() *string {
	return s.Branch
}

func (s *ListRoutineBuildsResponseBodyData) GetBuildCommand() *string {
	return s.BuildCommand
}

func (s *ListRoutineBuildsResponseBodyData) GetCommitId() *string {
	return s.CommitId
}

func (s *ListRoutineBuildsResponseBodyData) GetCommitMessage() *string {
	return s.CommitMessage
}

func (s *ListRoutineBuildsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRoutineBuildsResponseBodyData) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *ListRoutineBuildsResponseBodyData) GetGitAccountId() *int64 {
	return s.GitAccountId
}

func (s *ListRoutineBuildsResponseBodyData) GetGitAccountName() *string {
	return s.GitAccountName
}

func (s *ListRoutineBuildsResponseBodyData) GetInstallCommand() *string {
	return s.InstallCommand
}

func (s *ListRoutineBuildsResponseBodyData) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *ListRoutineBuildsResponseBodyData) GetNodeVersion() *string {
	return s.NodeVersion
}

func (s *ListRoutineBuildsResponseBodyData) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *ListRoutineBuildsResponseBodyData) GetPipelineRunId() *int64 {
	return s.PipelineRunId
}

func (s *ListRoutineBuildsResponseBodyData) GetProductionBranch() *string {
	return s.ProductionBranch
}

func (s *ListRoutineBuildsResponseBodyData) GetRepository() *string {
	return s.Repository
}

func (s *ListRoutineBuildsResponseBodyData) GetRootDirectory() *string {
	return s.RootDirectory
}

func (s *ListRoutineBuildsResponseBodyData) GetRoutineBuildId() *int64 {
	return s.RoutineBuildId
}

func (s *ListRoutineBuildsResponseBodyData) GetRoutineEntry() *string {
	return s.RoutineEntry
}

func (s *ListRoutineBuildsResponseBodyData) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListRoutineBuildsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListRoutineBuildsResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListRoutineBuildsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRoutineBuildsResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *ListRoutineBuildsResponseBodyData) SetAssetsDirectory(v string) *ListRoutineBuildsResponseBodyData {
	s.AssetsDirectory = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetBranch(v string) *ListRoutineBuildsResponseBodyData {
	s.Branch = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetBuildCommand(v string) *ListRoutineBuildsResponseBodyData {
	s.BuildCommand = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetCommitId(v string) *ListRoutineBuildsResponseBodyData {
	s.CommitId = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetCommitMessage(v string) *ListRoutineBuildsResponseBodyData {
	s.CommitMessage = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetCreateTime(v string) *ListRoutineBuildsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetEnvironmentVariables(v map[string]*string) *ListRoutineBuildsResponseBodyData {
	s.EnvironmentVariables = v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetGitAccountId(v int64) *ListRoutineBuildsResponseBodyData {
	s.GitAccountId = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetGitAccountName(v string) *ListRoutineBuildsResponseBodyData {
	s.GitAccountName = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetInstallCommand(v string) *ListRoutineBuildsResponseBodyData {
	s.InstallCommand = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetIsPrivate(v bool) *ListRoutineBuildsResponseBodyData {
	s.IsPrivate = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetNodeVersion(v string) *ListRoutineBuildsResponseBodyData {
	s.NodeVersion = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetPipelineId(v int64) *ListRoutineBuildsResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetPipelineRunId(v int64) *ListRoutineBuildsResponseBodyData {
	s.PipelineRunId = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetProductionBranch(v string) *ListRoutineBuildsResponseBodyData {
	s.ProductionBranch = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetRepository(v string) *ListRoutineBuildsResponseBodyData {
	s.Repository = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetRootDirectory(v string) *ListRoutineBuildsResponseBodyData {
	s.RootDirectory = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetRoutineBuildId(v int64) *ListRoutineBuildsResponseBodyData {
	s.RoutineBuildId = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetRoutineEntry(v string) *ListRoutineBuildsResponseBodyData {
	s.RoutineEntry = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetRoutineName(v string) *ListRoutineBuildsResponseBodyData {
	s.RoutineName = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetStatus(v string) *ListRoutineBuildsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetTemplateName(v string) *ListRoutineBuildsResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetUpdateTime(v string) *ListRoutineBuildsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) SetUserId(v int64) *ListRoutineBuildsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListRoutineBuildsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
