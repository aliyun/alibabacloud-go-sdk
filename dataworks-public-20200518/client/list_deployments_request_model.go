// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListDeploymentsRequest
	GetCreator() *string
	SetEndCreateTime(v int64) *ListDeploymentsRequest
	GetEndCreateTime() *int64
	SetEndExecuteTime(v int64) *ListDeploymentsRequest
	GetEndExecuteTime() *int64
	SetExecutor(v string) *ListDeploymentsRequest
	GetExecutor() *string
	SetKeyword(v string) *ListDeploymentsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListDeploymentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeploymentsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDeploymentsRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *ListDeploymentsRequest
	GetProjectIdentifier() *string
	SetStatus(v int32) *ListDeploymentsRequest
	GetStatus() *int32
}

type ListDeploymentsRequest struct {
	// The ID of the Alibaba Cloud account used by the user who creates the deployment packages.
	//
	// example:
	//
	// 20030****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the deployment packages to be queried are created. This value must be a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1593877765000
	EndCreateTime *int64 `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	// The time when the deployment packages are run. This value must be a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1593877765000
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who runs the deployment packages.
	//
	// example:
	//
	// 2003****
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The keyword that is contained in the names of the deployment packages. A fuzzy search is supported. After you enter a keyword, all deployment packages whose names contain the keyword are displayed.
	//
	// example:
	//
	// hello
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace name.
	//
	// You must configure either this parameter or the ProjectId parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The status of the deployment packages. Valid values:
	//
	// 	- 0: The deployment packages are ready.
	//
	// 	- 1: The deployment packages are deployed.
	//
	// 	- 2: The deployment packages fail to be deployed.
	//
	// 	- 6: The deployment packages are rejected.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListDeploymentsRequest) GetEndCreateTime() *int64 {
	return s.EndCreateTime
}

func (s *ListDeploymentsRequest) GetEndExecuteTime() *int64 {
	return s.EndExecuteTime
}

func (s *ListDeploymentsRequest) GetExecutor() *string {
	return s.Executor
}

func (s *ListDeploymentsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDeploymentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeploymentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDeploymentsRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *ListDeploymentsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListDeploymentsRequest) SetCreator(v string) *ListDeploymentsRequest {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsRequest) SetEndCreateTime(v int64) *ListDeploymentsRequest {
	s.EndCreateTime = &v
	return s
}

func (s *ListDeploymentsRequest) SetEndExecuteTime(v int64) *ListDeploymentsRequest {
	s.EndExecuteTime = &v
	return s
}

func (s *ListDeploymentsRequest) SetExecutor(v string) *ListDeploymentsRequest {
	s.Executor = &v
	return s
}

func (s *ListDeploymentsRequest) SetKeyword(v string) *ListDeploymentsRequest {
	s.Keyword = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageNumber(v int32) *ListDeploymentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageSize(v int32) *ListDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsRequest) SetProjectId(v int64) *ListDeploymentsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentsRequest) SetProjectIdentifier(v string) *ListDeploymentsRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *ListDeploymentsRequest) SetStatus(v int32) *ListDeploymentsRequest {
	s.Status = &v
	return s
}

func (s *ListDeploymentsRequest) Validate() error {
	return dara.Validate(s)
}
