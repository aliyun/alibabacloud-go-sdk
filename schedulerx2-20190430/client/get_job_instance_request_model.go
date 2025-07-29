// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetJobInstanceRequest
	GetGroupId() *string
	SetJobId(v int64) *GetJobInstanceRequest
	GetJobId() *int64
	SetJobInstanceId(v int64) *GetJobInstanceRequest
	GetJobInstanceId() *int64
	SetNamespace(v string) *GetJobInstanceRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GetJobInstanceRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *GetJobInstanceRequest
	GetRegionId() *string
}

type GetJobInstanceRequest struct {
	// The application ID. You can obtain the application ID on the Application Management page in the SchedulerX console.
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
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111111
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetJobInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetJobInstanceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetJobInstanceRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *GetJobInstanceRequest) GetJobInstanceId() *int64 {
	return s.JobInstanceId
}

func (s *GetJobInstanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetJobInstanceRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GetJobInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetJobInstanceRequest) SetGroupId(v string) *GetJobInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *GetJobInstanceRequest) SetJobId(v int64) *GetJobInstanceRequest {
	s.JobId = &v
	return s
}

func (s *GetJobInstanceRequest) SetJobInstanceId(v int64) *GetJobInstanceRequest {
	s.JobInstanceId = &v
	return s
}

func (s *GetJobInstanceRequest) SetNamespace(v string) *GetJobInstanceRequest {
	s.Namespace = &v
	return s
}

func (s *GetJobInstanceRequest) SetNamespaceSource(v string) *GetJobInstanceRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetJobInstanceRequest) SetRegionId(v string) *GetJobInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobInstanceRequest) Validate() error {
	return dara.Validate(s)
}
