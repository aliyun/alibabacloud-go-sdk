// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectPipelineRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeFrom(v int64) *ListCrossProjectPipelineRunsRequest
	GetCreateTimeFrom() *int64
	SetCreateTimeTo(v int64) *ListCrossProjectPipelineRunsRequest
	GetCreateTimeTo() *int64
	SetCreator(v string) *ListCrossProjectPipelineRunsRequest
	GetCreator() *string
	SetDeploymentEnvironmentId(v int64) *ListCrossProjectPipelineRunsRequest
	GetDeploymentEnvironmentId() *int64
	SetExecutor(v string) *ListCrossProjectPipelineRunsRequest
	GetExecutor() *string
	SetObjectId(v string) *ListCrossProjectPipelineRunsRequest
	GetObjectId() *string
	SetObjectType(v string) *ListCrossProjectPipelineRunsRequest
	GetObjectType() *string
	SetPageNumber(v int32) *ListCrossProjectPipelineRunsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCrossProjectPipelineRunsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCrossProjectPipelineRunsRequest
	GetProjectId() *int64
	SetStatus(v string) *ListCrossProjectPipelineRunsRequest
	GetStatus() *string
}

type ListCrossProjectPipelineRunsRequest struct {
	// The start of the creation time range. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788739200000
	CreateTimeFrom *int64 `json:"CreateTimeFrom,omitempty" xml:"CreateTimeFrom,omitempty"`
	// The end of the creation time range. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788825599999
	CreateTimeTo *int64 `json:"CreateTimeTo,omitempty" xml:"CreateTimeTo,omitempty"`
	// The creator.
	//
	// example:
	//
	// creator
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The cross-workspace publish environment ID.
	//
	// example:
	//
	// 101
	DeploymentEnvironmentId *int64 `json:"DeploymentEnvironmentId,omitempty" xml:"DeploymentEnvironmentId,omitempty"`
	// The executor.
	//
	// example:
	//
	// executor
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The publish object ID.
	//
	// example:
	//
	// 1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The publish object type.
	//
	// example:
	//
	// ODPS_SQL
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The publish flow status. Valid values:
	//
	// - Building: Building.
	//
	// - Ready: Ready and waiting for execution.
	//
	// - Running: Running.
	//
	// - Termination: Terminated.
	//
	// - Success: Succeeded.
	//
	// - Fail: Failed.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCrossProjectPipelineRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunsRequest) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunsRequest) GetCreateTimeFrom() *int64 {
	return s.CreateTimeFrom
}

func (s *ListCrossProjectPipelineRunsRequest) GetCreateTimeTo() *int64 {
	return s.CreateTimeTo
}

func (s *ListCrossProjectPipelineRunsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListCrossProjectPipelineRunsRequest) GetDeploymentEnvironmentId() *int64 {
	return s.DeploymentEnvironmentId
}

func (s *ListCrossProjectPipelineRunsRequest) GetExecutor() *string {
	return s.Executor
}

func (s *ListCrossProjectPipelineRunsRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListCrossProjectPipelineRunsRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListCrossProjectPipelineRunsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrossProjectPipelineRunsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrossProjectPipelineRunsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCrossProjectPipelineRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCrossProjectPipelineRunsRequest) SetCreateTimeFrom(v int64) *ListCrossProjectPipelineRunsRequest {
	s.CreateTimeFrom = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetCreateTimeTo(v int64) *ListCrossProjectPipelineRunsRequest {
	s.CreateTimeTo = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetCreator(v string) *ListCrossProjectPipelineRunsRequest {
	s.Creator = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetDeploymentEnvironmentId(v int64) *ListCrossProjectPipelineRunsRequest {
	s.DeploymentEnvironmentId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetExecutor(v string) *ListCrossProjectPipelineRunsRequest {
	s.Executor = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetObjectId(v string) *ListCrossProjectPipelineRunsRequest {
	s.ObjectId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetObjectType(v string) *ListCrossProjectPipelineRunsRequest {
	s.ObjectType = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetPageNumber(v int32) *ListCrossProjectPipelineRunsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetPageSize(v int32) *ListCrossProjectPipelineRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetProjectId(v int64) *ListCrossProjectPipelineRunsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) SetStatus(v string) *ListCrossProjectPipelineRunsRequest {
	s.Status = &v
	return s
}

func (s *ListCrossProjectPipelineRunsRequest) Validate() error {
	return dara.Validate(s)
}
