// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDeploymentJobRequest
	GetCurrentPage() *int32
	SetJobType(v string) *ListDeploymentJobRequest
	GetJobType() *string
	SetShowSize(v int32) *ListDeploymentJobRequest
	GetShowSize() *int32
	SetStatus(v string) *ListDeploymentJobRequest
	GetStatus() *string
}

type ListDeploymentJobRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the deployment task.
	//
	// Valid values:
	//
	// 	- cloud: multi-cloud deployment task.
	//
	// 	- user: cloud service deployment task. This type of task does not support Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The number of certificates per page. Default value: **50**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The status of the deployment task.
	//
	// Valid values:
	//
	// 	- success
	//
	// 	- pending
	//
	// 	- scheduling
	//
	// 	- processing
	//
	// 	- error
	//
	// 	- editing
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDeploymentJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListDeploymentJobRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListDeploymentJobRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDeploymentJobRequest) SetCurrentPage(v int32) *ListDeploymentJobRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDeploymentJobRequest) SetJobType(v string) *ListDeploymentJobRequest {
	s.JobType = &v
	return s
}

func (s *ListDeploymentJobRequest) SetShowSize(v int32) *ListDeploymentJobRequest {
	s.ShowSize = &v
	return s
}

func (s *ListDeploymentJobRequest) SetStatus(v string) *ListDeploymentJobRequest {
	s.Status = &v
	return s
}

func (s *ListDeploymentJobRequest) Validate() error {
	return dara.Validate(s)
}
