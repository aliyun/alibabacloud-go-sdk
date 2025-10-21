// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSavepointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *ListSavepointsRequest
	GetDeploymentId() *string
	SetJobId(v string) *ListSavepointsRequest
	GetJobId() *string
	SetPageIndex(v int32) *ListSavepointsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListSavepointsRequest
	GetPageSize() *int32
}

type ListSavepointsRequest struct {
	// The deployment ID. This parameter is optional.
	//
	// example:
	//
	// 88a8fc49-e090-430a-85d8-3ee8c79c****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The job ID. This parameter is optional.
	//
	// example:
	//
	// 99a8fc49-e090-430a-85d8-3ee8c79c****
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
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
}

func (s ListSavepointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSavepointsRequest) GoString() string {
	return s.String()
}

func (s *ListSavepointsRequest) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ListSavepointsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListSavepointsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListSavepointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSavepointsRequest) SetDeploymentId(v string) *ListSavepointsRequest {
	s.DeploymentId = &v
	return s
}

func (s *ListSavepointsRequest) SetJobId(v string) *ListSavepointsRequest {
	s.JobId = &v
	return s
}

func (s *ListSavepointsRequest) SetPageIndex(v int32) *ListSavepointsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListSavepointsRequest) SetPageSize(v int32) *ListSavepointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSavepointsRequest) Validate() error {
	return dara.Validate(s)
}
