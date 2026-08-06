// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticJobRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobName(v string) *ListSemanticJobRunsRequest
	GetJobName() *string
	SetPageNumber(v int32) *ListSemanticJobRunsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSemanticJobRunsRequest
	GetPageSize() *int32
}

type ListSemanticJobRunsRequest struct {
	// The job name. Use the Data.Name value from the CreateSemanticJob response or the Name value from a ListSemanticJobs list item.
	//
	// This parameter is required.
	//
	// example:
	//
	// semantic-job-demo
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The page number, starting from 1. If this parameter is not specified or set to a value less than or equal to 0, page 1 is returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of run records to return per page. If this parameter is not specified or set to a value less than or equal to 0, the default value 50 is used. Maximum value: 200.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSemanticJobRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobRunsRequest) GoString() string {
	return s.String()
}

func (s *ListSemanticJobRunsRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListSemanticJobRunsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSemanticJobRunsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSemanticJobRunsRequest) SetJobName(v string) *ListSemanticJobRunsRequest {
	s.JobName = &v
	return s
}

func (s *ListSemanticJobRunsRequest) SetPageNumber(v int32) *ListSemanticJobRunsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSemanticJobRunsRequest) SetPageSize(v int32) *ListSemanticJobRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSemanticJobRunsRequest) Validate() error {
	return dara.Validate(s)
}
