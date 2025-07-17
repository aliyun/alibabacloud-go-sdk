// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListDeploymentPackagesRequest
	GetCreator() *string
	SetEndCreateTime(v int64) *ListDeploymentPackagesRequest
	GetEndCreateTime() *int64
	SetEndExecuteTime(v int64) *ListDeploymentPackagesRequest
	GetEndExecuteTime() *int64
	SetExecutor(v string) *ListDeploymentPackagesRequest
	GetExecutor() *string
	SetKeyword(v string) *ListDeploymentPackagesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListDeploymentPackagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeploymentPackagesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDeploymentPackagesRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *ListDeploymentPackagesRequest
	GetProjectIdentifier() *string
	SetStatus(v int32) *ListDeploymentPackagesRequest
	GetStatus() *int32
}

type ListDeploymentPackagesRequest struct {
	// example:
	//
	// 110755000425****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1593877765000
	EndCreateTime *int64 `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	// example:
	//
	// 1593877765000
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// example:
	//
	// 2003****
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10003
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackagesRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListDeploymentPackagesRequest) GetEndCreateTime() *int64 {
	return s.EndCreateTime
}

func (s *ListDeploymentPackagesRequest) GetEndExecuteTime() *int64 {
	return s.EndExecuteTime
}

func (s *ListDeploymentPackagesRequest) GetExecutor() *string {
	return s.Executor
}

func (s *ListDeploymentPackagesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDeploymentPackagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeploymentPackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentPackagesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDeploymentPackagesRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *ListDeploymentPackagesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListDeploymentPackagesRequest) SetCreator(v string) *ListDeploymentPackagesRequest {
	s.Creator = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetEndCreateTime(v int64) *ListDeploymentPackagesRequest {
	s.EndCreateTime = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetEndExecuteTime(v int64) *ListDeploymentPackagesRequest {
	s.EndExecuteTime = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetExecutor(v string) *ListDeploymentPackagesRequest {
	s.Executor = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetKeyword(v string) *ListDeploymentPackagesRequest {
	s.Keyword = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetPageNumber(v int32) *ListDeploymentPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetPageSize(v int32) *ListDeploymentPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetProjectId(v int64) *ListDeploymentPackagesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetProjectIdentifier(v string) *ListDeploymentPackagesRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *ListDeploymentPackagesRequest) SetStatus(v int32) *ListDeploymentPackagesRequest {
	s.Status = &v
	return s
}

func (s *ListDeploymentPackagesRequest) Validate() error {
	return dara.Validate(s)
}
