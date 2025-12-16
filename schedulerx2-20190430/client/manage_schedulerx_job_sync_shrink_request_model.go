// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxJobSyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIdListShrink(v string) *ManageSchedulerxJobSyncShrinkRequest
	GetJobIdListShrink() *string
	SetNamespaceSource(v string) *ManageSchedulerxJobSyncShrinkRequest
	GetNamespaceSource() *string
	SetOriginalGroupId(v string) *ManageSchedulerxJobSyncShrinkRequest
	GetOriginalGroupId() *string
	SetOriginalNamespace(v string) *ManageSchedulerxJobSyncShrinkRequest
	GetOriginalNamespace() *string
	SetRegionId(v string) *ManageSchedulerxJobSyncShrinkRequest
	GetRegionId() *string
	SetTargetGroupId(v string) *ManageSchedulerxJobSyncShrinkRequest
	GetTargetGroupId() *string
	SetTargetNamespace(v string) *ManageSchedulerxJobSyncShrinkRequest
	GetTargetNamespace() *string
}

type ManageSchedulerxJobSyncShrinkRequest struct {
	// The list of task IDs.
	//
	// This parameter is required.
	JobIdListShrink *string `json:"JobIdList,omitempty" xml:"JobIdList,omitempty"`
	// The source of the namespace. Required only for specific third-party cases.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The source application group to which the task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	OriginalGroupId *string `json:"OriginalGroupId,omitempty" xml:"OriginalGroupId,omitempty"`
	// The source namespace where the task resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	OriginalNamespace *string `json:"OriginalNamespace,omitempty" xml:"OriginalNamespace,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the destination application group to which the task will be synchronized.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSyncJobGroup
	TargetGroupId *string `json:"TargetGroupId,omitempty" xml:"TargetGroupId,omitempty"`
	// The destination namespace to which the task will be synchronized.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5f56ef65-b836-493d-b40b-c4db6425****
	TargetNamespace *string `json:"TargetNamespace,omitempty" xml:"TargetNamespace,omitempty"`
}

func (s ManageSchedulerxJobSyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxJobSyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxJobSyncShrinkRequest) GetJobIdListShrink() *string {
	return s.JobIdListShrink
}

func (s *ManageSchedulerxJobSyncShrinkRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ManageSchedulerxJobSyncShrinkRequest) GetOriginalGroupId() *string {
	return s.OriginalGroupId
}

func (s *ManageSchedulerxJobSyncShrinkRequest) GetOriginalNamespace() *string {
	return s.OriginalNamespace
}

func (s *ManageSchedulerxJobSyncShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ManageSchedulerxJobSyncShrinkRequest) GetTargetGroupId() *string {
	return s.TargetGroupId
}

func (s *ManageSchedulerxJobSyncShrinkRequest) GetTargetNamespace() *string {
	return s.TargetNamespace
}

func (s *ManageSchedulerxJobSyncShrinkRequest) SetJobIdListShrink(v string) *ManageSchedulerxJobSyncShrinkRequest {
	s.JobIdListShrink = &v
	return s
}

func (s *ManageSchedulerxJobSyncShrinkRequest) SetNamespaceSource(v string) *ManageSchedulerxJobSyncShrinkRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ManageSchedulerxJobSyncShrinkRequest) SetOriginalGroupId(v string) *ManageSchedulerxJobSyncShrinkRequest {
	s.OriginalGroupId = &v
	return s
}

func (s *ManageSchedulerxJobSyncShrinkRequest) SetOriginalNamespace(v string) *ManageSchedulerxJobSyncShrinkRequest {
	s.OriginalNamespace = &v
	return s
}

func (s *ManageSchedulerxJobSyncShrinkRequest) SetRegionId(v string) *ManageSchedulerxJobSyncShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ManageSchedulerxJobSyncShrinkRequest) SetTargetGroupId(v string) *ManageSchedulerxJobSyncShrinkRequest {
	s.TargetGroupId = &v
	return s
}

func (s *ManageSchedulerxJobSyncShrinkRequest) SetTargetNamespace(v string) *ManageSchedulerxJobSyncShrinkRequest {
	s.TargetNamespace = &v
	return s
}

func (s *ManageSchedulerxJobSyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
