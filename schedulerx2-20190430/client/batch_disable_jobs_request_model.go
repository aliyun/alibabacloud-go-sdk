// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDisableJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *BatchDisableJobsRequest
	GetGroupId() *string
	SetJobIdList(v []*int64) *BatchDisableJobsRequest
	GetJobIdList() []*int64
	SetNamespace(v string) *BatchDisableJobsRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *BatchDisableJobsRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *BatchDisableJobsRequest
	GetRegionId() *string
}

type BatchDisableJobsRequest struct {
	// The ID of the application group. You can find this ID on the **Application Management*	- page in the console.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The IDs of the jobs to disable.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99341
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// The ID of the namespace that contains the jobs. You can find this ID on the **Namespaces*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is used only by specific third-party services.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region where the jobs are located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchDisableJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDisableJobsRequest) GoString() string {
	return s.String()
}

func (s *BatchDisableJobsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchDisableJobsRequest) GetJobIdList() []*int64 {
	return s.JobIdList
}

func (s *BatchDisableJobsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDisableJobsRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *BatchDisableJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchDisableJobsRequest) SetGroupId(v string) *BatchDisableJobsRequest {
	s.GroupId = &v
	return s
}

func (s *BatchDisableJobsRequest) SetJobIdList(v []*int64) *BatchDisableJobsRequest {
	s.JobIdList = v
	return s
}

func (s *BatchDisableJobsRequest) SetNamespace(v string) *BatchDisableJobsRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDisableJobsRequest) SetNamespaceSource(v string) *BatchDisableJobsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *BatchDisableJobsRequest) SetRegionId(v string) *BatchDisableJobsRequest {
	s.RegionId = &v
	return s
}

func (s *BatchDisableJobsRequest) Validate() error {
	return dara.Validate(s)
}
