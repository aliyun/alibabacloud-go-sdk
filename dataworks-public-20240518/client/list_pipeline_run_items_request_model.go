// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPipelineRunItemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPipelineRunItemsRequest
	GetPageSize() *int32
	SetPipelineRunId(v string) *ListPipelineRunItemsRequest
	GetPipelineRunId() *string
	SetProjectId(v int64) *ListPipelineRunItemsRequest
	GetProjectId() *int64
}

type ListPipelineRunItemsRequest struct {
	// The page number, for pagination.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The number of entries per page. Default: 10. Maximum: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workflow task ID. To obtain the ID, see [ListPipelineRuns](https://help.aliyun.com/document_detail/438042.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 097c73fe-ed6e-4fb1-b109-a5d59e46cd58
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	// The ID of the DataWorks workspace. You can obtain the workspace ID from the workspace configuration page in the DataWorks console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListPipelineRunItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunItemsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunItemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPipelineRunItemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPipelineRunItemsRequest) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *ListPipelineRunItemsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListPipelineRunItemsRequest) SetPageNumber(v int32) *ListPipelineRunItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineRunItemsRequest) SetPageSize(v int32) *ListPipelineRunItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPipelineRunItemsRequest) SetPipelineRunId(v string) *ListPipelineRunItemsRequest {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineRunItemsRequest) SetProjectId(v int64) *ListPipelineRunItemsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListPipelineRunItemsRequest) Validate() error {
	return dara.Validate(s)
}
