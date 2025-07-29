// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *BatchDeleteJobsRequest
	GetGroupId() *string
	SetJobIdList(v []*int64) *BatchDeleteJobsRequest
	GetJobIdList() []*int64
	SetNamespace(v string) *BatchDeleteJobsRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *BatchDeleteJobsRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *BatchDeleteJobsRequest
	GetRegionId() *string
}

type BatchDeleteJobsRequest struct {
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job IDs. Separate multiple job IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 99341
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// The ID of the namespace to which the job belongs. You can obtain the ID of the namespace on the **Namespace*	- page in the SchedulerX console.
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
	// Schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region to which the job belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchDeleteJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteJobsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchDeleteJobsRequest) GetJobIdList() []*int64 {
	return s.JobIdList
}

func (s *BatchDeleteJobsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteJobsRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *BatchDeleteJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchDeleteJobsRequest) SetGroupId(v string) *BatchDeleteJobsRequest {
	s.GroupId = &v
	return s
}

func (s *BatchDeleteJobsRequest) SetJobIdList(v []*int64) *BatchDeleteJobsRequest {
	s.JobIdList = v
	return s
}

func (s *BatchDeleteJobsRequest) SetNamespace(v string) *BatchDeleteJobsRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteJobsRequest) SetNamespaceSource(v string) *BatchDeleteJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *BatchDeleteJobsRequest) SetRegionId(v string) *BatchDeleteJobsRequest {
	s.RegionId = &v
	return s
}

func (s *BatchDeleteJobsRequest) Validate() error {
	return dara.Validate(s)
}
