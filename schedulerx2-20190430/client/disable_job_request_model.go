// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DisableJobRequest
	GetGroupId() *string
	SetJobId(v int64) *DisableJobRequest
	GetJobId() *int64
	SetNamespace(v string) *DisableJobRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *DisableJobRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *DisableJobRequest
	GetRegionId() *string
}

type DisableJobRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableJobRequest) GoString() string {
	return s.String()
}

func (s *DisableJobRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DisableJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DisableJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DisableJobRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *DisableJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableJobRequest) SetGroupId(v string) *DisableJobRequest {
	s.GroupId = &v
	return s
}

func (s *DisableJobRequest) SetJobId(v int64) *DisableJobRequest {
	s.JobId = &v
	return s
}

func (s *DisableJobRequest) SetNamespace(v string) *DisableJobRequest {
	s.Namespace = &v
	return s
}

func (s *DisableJobRequest) SetNamespaceSource(v string) *DisableJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DisableJobRequest) SetRegionId(v string) *DisableJobRequest {
	s.RegionId = &v
	return s
}

func (s *DisableJobRequest) Validate() error {
	return dara.Validate(s)
}
