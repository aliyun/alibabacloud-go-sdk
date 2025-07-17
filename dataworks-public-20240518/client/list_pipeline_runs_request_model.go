// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListPipelineRunsRequest
	GetCreator() *string
	SetPageNumber(v int32) *ListPipelineRunsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPipelineRunsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListPipelineRunsRequest
	GetProjectId() *int64
	SetStatus(v string) *ListPipelineRunsRequest
	GetStatus() *string
}

type ListPipelineRunsRequest struct {
	// The ID of the user who creates the processes. This parameter specifies a filter condition.
	//
	// example:
	//
	// 110755000425****
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the processes. This parameter specifies a filter condition.
	//
	// Valid values:
	//
	// 	- Init
	//
	// 	- Running
	//
	// 	- Success
	//
	// 	- Fail
	//
	// 	- Termination
	//
	// 	- Cancel
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPipelineRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListPipelineRunsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPipelineRunsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPipelineRunsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListPipelineRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListPipelineRunsRequest) SetCreator(v string) *ListPipelineRunsRequest {
	s.Creator = &v
	return s
}

func (s *ListPipelineRunsRequest) SetPageNumber(v int32) *ListPipelineRunsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineRunsRequest) SetPageSize(v int32) *ListPipelineRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPipelineRunsRequest) SetProjectId(v int64) *ListPipelineRunsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListPipelineRunsRequest) SetStatus(v string) *ListPipelineRunsRequest {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsRequest) Validate() error {
	return dara.Validate(s)
}
