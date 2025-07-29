// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEnableJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *BatchEnableJobsRequest
	GetGroupId() *string
	SetJobIdList(v []*int64) *BatchEnableJobsRequest
	GetJobIdList() []*int64
	SetNamespace(v string) *BatchEnableJobsRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *BatchEnableJobsRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *BatchEnableJobsRequest
	GetRegionId() *string
}

type BatchEnableJobsRequest struct {
	// The application ID. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job IDs. Multiple job IDs are separated with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 99341
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// The ID of the namespace to which the job belongs. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
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
	// The ID of the region to which the job belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchEnableJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchEnableJobsRequest) GoString() string {
	return s.String()
}

func (s *BatchEnableJobsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchEnableJobsRequest) GetJobIdList() []*int64 {
	return s.JobIdList
}

func (s *BatchEnableJobsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchEnableJobsRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *BatchEnableJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchEnableJobsRequest) SetGroupId(v string) *BatchEnableJobsRequest {
	s.GroupId = &v
	return s
}

func (s *BatchEnableJobsRequest) SetJobIdList(v []*int64) *BatchEnableJobsRequest {
	s.JobIdList = v
	return s
}

func (s *BatchEnableJobsRequest) SetNamespace(v string) *BatchEnableJobsRequest {
	s.Namespace = &v
	return s
}

func (s *BatchEnableJobsRequest) SetNamespaceSource(v string) *BatchEnableJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *BatchEnableJobsRequest) SetRegionId(v string) *BatchEnableJobsRequest {
	s.RegionId = &v
	return s
}

func (s *BatchEnableJobsRequest) Validate() error {
	return dara.Validate(s)
}
