// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentId(v string) *ListScheduledPlanRequest
	GetDeploymentId() *string
	SetPageIndex(v int32) *ListScheduledPlanRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListScheduledPlanRequest
	GetPageSize() *int32
}

type ListScheduledPlanRequest struct {
	// The ID of the job.
	//
	// example:
	//
	// 737d0921-c5ac-47fc-9ba9-07a1e0b4****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The page number to return. The default value is 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries to return on each page. The default value is 10. The maximum value is 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListScheduledPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPlanRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanRequest) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ListScheduledPlanRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListScheduledPlanRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduledPlanRequest) SetDeploymentId(v string) *ListScheduledPlanRequest {
	s.DeploymentId = &v
	return s
}

func (s *ListScheduledPlanRequest) SetPageIndex(v int32) *ListScheduledPlanRequest {
	s.PageIndex = &v
	return s
}

func (s *ListScheduledPlanRequest) SetPageSize(v int32) *ListScheduledPlanRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledPlanRequest) Validate() error {
	return dara.Validate(s)
}
