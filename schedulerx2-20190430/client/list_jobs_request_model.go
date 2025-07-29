// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListJobsRequest
	GetGroupId() *string
	SetJobName(v string) *ListJobsRequest
	GetJobName() *string
	SetNamespace(v string) *ListJobsRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *ListJobsRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *ListJobsRequest
	GetRegionId() *string
	SetStatus(v string) *ListJobsRequest
	GetStatus() *string
}

type ListJobsRequest struct {
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// DocTest.Group
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// helloword
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The ID of the namespace. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a72ecb1-b4cc-400a-a71b-20cdec9b****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable the job. Valid values:
	//
	// 	- **0**: disables the job.
	//
	// 	- **1**: enables the job.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListJobsRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListJobsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListJobsRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ListJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobsRequest) SetGroupId(v string) *ListJobsRequest {
	s.GroupId = &v
	return s
}

func (s *ListJobsRequest) SetJobName(v string) *ListJobsRequest {
	s.JobName = &v
	return s
}

func (s *ListJobsRequest) SetNamespace(v string) *ListJobsRequest {
	s.Namespace = &v
	return s
}

func (s *ListJobsRequest) SetNamespaceSource(v string) *ListJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListJobsRequest) SetRegionId(v string) *ListJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
