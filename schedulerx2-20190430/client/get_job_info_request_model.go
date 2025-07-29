// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetJobInfoRequest
	GetGroupId() *string
	SetJobId(v int64) *GetJobInfoRequest
	GetJobId() *int64
	SetJobName(v string) *GetJobInfoRequest
	GetJobName() *string
	SetNamespace(v string) *GetJobInfoRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GetJobInfoRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *GetJobInfoRequest
	GetRegionId() *string
}

type GetJobInfoRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
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
	// The job name.
	//
	// example:
	//
	// simpleJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The namespace source. This parameter is required only for a special third party.
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

func (s GetJobInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoRequest) GoString() string {
	return s.String()
}

func (s *GetJobInfoRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetJobInfoRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *GetJobInfoRequest) GetJobName() *string {
	return s.JobName
}

func (s *GetJobInfoRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetJobInfoRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GetJobInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetJobInfoRequest) SetGroupId(v string) *GetJobInfoRequest {
	s.GroupId = &v
	return s
}

func (s *GetJobInfoRequest) SetJobId(v int64) *GetJobInfoRequest {
	s.JobId = &v
	return s
}

func (s *GetJobInfoRequest) SetJobName(v string) *GetJobInfoRequest {
	s.JobName = &v
	return s
}

func (s *GetJobInfoRequest) SetNamespace(v string) *GetJobInfoRequest {
	s.Namespace = &v
	return s
}

func (s *GetJobInfoRequest) SetNamespaceSource(v string) *GetJobInfoRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetJobInfoRequest) SetRegionId(v string) *GetJobInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobInfoRequest) Validate() error {
	return dara.Validate(s)
}
