// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryJobInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *RetryJobInstanceRequest
	GetGroupId() *string
	SetJobId(v int64) *RetryJobInstanceRequest
	GetJobId() *int64
	SetJobInstanceId(v int64) *RetryJobInstanceRequest
	GetJobInstanceId() *int64
	SetNamespace(v string) *RetryJobInstanceRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *RetryJobInstanceRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *RetryJobInstanceRequest
	GetRegionId() *string
}

type RetryJobInstanceRequest struct {
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	JobInstanceId *int64 `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
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

func (s RetryJobInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryJobInstanceRequest) GoString() string {
	return s.String()
}

func (s *RetryJobInstanceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RetryJobInstanceRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *RetryJobInstanceRequest) GetJobInstanceId() *int64 {
	return s.JobInstanceId
}

func (s *RetryJobInstanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *RetryJobInstanceRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *RetryJobInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RetryJobInstanceRequest) SetGroupId(v string) *RetryJobInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *RetryJobInstanceRequest) SetJobId(v int64) *RetryJobInstanceRequest {
	s.JobId = &v
	return s
}

func (s *RetryJobInstanceRequest) SetJobInstanceId(v int64) *RetryJobInstanceRequest {
	s.JobInstanceId = &v
	return s
}

func (s *RetryJobInstanceRequest) SetNamespace(v string) *RetryJobInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *RetryJobInstanceRequest) SetNamespaceSource(v string) *RetryJobInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *RetryJobInstanceRequest) SetRegionId(v string) *RetryJobInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RetryJobInstanceRequest) Validate() error {
	return dara.Validate(s)
}
