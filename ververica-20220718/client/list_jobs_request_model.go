// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *ListJobsRequest
	GetDeploymentId() *string
	SetPageIndex(v int32) *ListJobsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
	SetSortName(v string) *ListJobsRequest
	GetSortName() *string
}

type ListJobsRequest struct {
	// The deployment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 58718c99-3b29-4c5e-93bb-c9fc4ec6****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The collation.
	//
	// Valid values:
	//
	// 	- gmt_create
	//
	// 	- job_id
	//
	// 	- status
	//
	// example:
	//
	// gmt_create
	SortName *string `json:"sortName,omitempty" xml:"sortName,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ListJobsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) GetSortName() *string {
	return s.SortName
}

func (s *ListJobsRequest) SetDeploymentId(v string) *ListJobsRequest {
	s.DeploymentId = &v
	return s
}

func (s *ListJobsRequest) SetPageIndex(v int32) *ListJobsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetSortName(v string) *ListJobsRequest {
	s.SortName = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
