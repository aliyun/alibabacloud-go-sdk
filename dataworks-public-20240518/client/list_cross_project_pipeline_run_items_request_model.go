// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectPipelineRunItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListCrossProjectPipelineRunItemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCrossProjectPipelineRunItemsRequest
	GetPageSize() *int32
	SetPipelineRunId(v string) *ListCrossProjectPipelineRunItemsRequest
	GetPipelineRunId() *string
	SetProjectId(v int64) *ListCrossProjectPipelineRunItemsRequest
	GetProjectId() *int64
}

type ListCrossProjectPipelineRunItemsRequest struct {
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
	// The ID of the cross-workspace publish pipeline.
	//
	// This parameter is required.
	//
	// example:
	//
	// fcfd4160-e2ff-4603-9719-09128fe733df
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListCrossProjectPipelineRunItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectPipelineRunItemsRequest) GoString() string {
	return s.String()
}

func (s *ListCrossProjectPipelineRunItemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrossProjectPipelineRunItemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrossProjectPipelineRunItemsRequest) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *ListCrossProjectPipelineRunItemsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCrossProjectPipelineRunItemsRequest) SetPageNumber(v int32) *ListCrossProjectPipelineRunItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsRequest) SetPageSize(v int32) *ListCrossProjectPipelineRunItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsRequest) SetPipelineRunId(v string) *ListCrossProjectPipelineRunItemsRequest {
	s.PipelineRunId = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsRequest) SetProjectId(v int64) *ListCrossProjectPipelineRunItemsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCrossProjectPipelineRunItemsRequest) Validate() error {
	return dara.Validate(s)
}
