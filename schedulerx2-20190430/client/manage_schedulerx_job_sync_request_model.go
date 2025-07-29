// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxJobSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIdList(v []*int64) *ManageSchedulerxJobSyncRequest
	GetJobIdList() []*int64
	SetNamespaceSource(v string) *ManageSchedulerxJobSyncRequest
	GetNamespaceSource() *string
	SetOriginalGroupId(v string) *ManageSchedulerxJobSyncRequest
	GetOriginalGroupId() *string
	SetOriginalNamespace(v string) *ManageSchedulerxJobSyncRequest
	GetOriginalNamespace() *string
	SetRegionId(v string) *ManageSchedulerxJobSyncRequest
	GetRegionId() *string
	SetTargetGroupId(v string) *ManageSchedulerxJobSyncRequest
	GetTargetGroupId() *string
	SetTargetNamespace(v string) *ManageSchedulerxJobSyncRequest
	GetTargetNamespace() *string
}

type ManageSchedulerxJobSyncRequest struct {
	// This parameter is required.
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	OriginalGroupId *string `json:"OriginalGroupId,omitempty" xml:"OriginalGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	OriginalNamespace *string `json:"OriginalNamespace,omitempty" xml:"OriginalNamespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testSyncJobGroup
	TargetGroupId *string `json:"TargetGroupId,omitempty" xml:"TargetGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5f56ef65-b836-493d-b40b-c4db6425****
	TargetNamespace *string `json:"TargetNamespace,omitempty" xml:"TargetNamespace,omitempty"`
}

func (s ManageSchedulerxJobSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxJobSyncRequest) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxJobSyncRequest) GetJobIdList() []*int64 {
	return s.JobIdList
}

func (s *ManageSchedulerxJobSyncRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ManageSchedulerxJobSyncRequest) GetOriginalGroupId() *string {
	return s.OriginalGroupId
}

func (s *ManageSchedulerxJobSyncRequest) GetOriginalNamespace() *string {
	return s.OriginalNamespace
}

func (s *ManageSchedulerxJobSyncRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ManageSchedulerxJobSyncRequest) GetTargetGroupId() *string {
	return s.TargetGroupId
}

func (s *ManageSchedulerxJobSyncRequest) GetTargetNamespace() *string {
	return s.TargetNamespace
}

func (s *ManageSchedulerxJobSyncRequest) SetJobIdList(v []*int64) *ManageSchedulerxJobSyncRequest {
	s.JobIdList = v
	return s
}

func (s *ManageSchedulerxJobSyncRequest) SetNamespaceSource(v string) *ManageSchedulerxJobSyncRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ManageSchedulerxJobSyncRequest) SetOriginalGroupId(v string) *ManageSchedulerxJobSyncRequest {
	s.OriginalGroupId = &v
	return s
}

func (s *ManageSchedulerxJobSyncRequest) SetOriginalNamespace(v string) *ManageSchedulerxJobSyncRequest {
	s.OriginalNamespace = &v
	return s
}

func (s *ManageSchedulerxJobSyncRequest) SetRegionId(v string) *ManageSchedulerxJobSyncRequest {
	s.RegionId = &v
	return s
}

func (s *ManageSchedulerxJobSyncRequest) SetTargetGroupId(v string) *ManageSchedulerxJobSyncRequest {
	s.TargetGroupId = &v
	return s
}

func (s *ManageSchedulerxJobSyncRequest) SetTargetNamespace(v string) *ManageSchedulerxJobSyncRequest {
	s.TargetNamespace = &v
	return s
}

func (s *ManageSchedulerxJobSyncRequest) Validate() error {
	return dara.Validate(s)
}
