// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteJobRequest
	GetGroupId() *string
	SetJobId(v int64) *DeleteJobRequest
	GetJobId() *int64
	SetNamespace(v string) *DeleteJobRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *DeleteJobRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *DeleteJobRequest
	GetRegionId() *string
}

type DeleteJobRequest struct {
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the job. You can obtain the ID on the **Task Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the namespace. You can obtain the ID of the namespace on the **Namespace*	- page in the SchedulerX console.
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

func (s DeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DeleteJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteJobRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *DeleteJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteJobRequest) SetGroupId(v string) *DeleteJobRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteJobRequest) SetJobId(v int64) *DeleteJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteJobRequest) SetNamespace(v string) *DeleteJobRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteJobRequest) SetNamespaceSource(v string) *DeleteJobRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DeleteJobRequest) SetRegionId(v string) *DeleteJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
