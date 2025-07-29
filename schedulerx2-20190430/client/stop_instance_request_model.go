// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *StopInstanceRequest
	GetGroupId() *string
	SetInstanceId(v int64) *StopInstanceRequest
	GetInstanceId() *int64
	SetJobId(v int64) *StopInstanceRequest
	GetJobId() *int64
	SetNamespace(v string) *StopInstanceRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *StopInstanceRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *StopInstanceRequest
	GetRegionId() *string
}

type StopInstanceRequest struct {
	// The ID of the application. You can obtain the application ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the job instance in the running state.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the job. You can obtain the ID of the job on the Task Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the namespace. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
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
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *StopInstanceRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *StopInstanceRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *StopInstanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *StopInstanceRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *StopInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopInstanceRequest) SetGroupId(v string) *StopInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v int64) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetJobId(v int64) *StopInstanceRequest {
	s.JobId = &v
	return s
}

func (s *StopInstanceRequest) SetNamespace(v string) *StopInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *StopInstanceRequest) SetNamespaceSource(v string) *StopInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *StopInstanceRequest) SetRegionId(v string) *StopInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopInstanceRequest) Validate() error {
	return dara.Validate(s)
}
