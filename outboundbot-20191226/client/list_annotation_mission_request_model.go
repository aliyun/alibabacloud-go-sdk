// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnnotationMissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotationMissionId(v string) *ListAnnotationMissionRequest
	GetAnnotationMissionId() *string
	SetAnnotationMissionName(v string) *ListAnnotationMissionRequest
	GetAnnotationMissionName() *string
	SetAnnotationStatusListFilter(v []*int32) *ListAnnotationMissionRequest
	GetAnnotationStatusListFilter() []*int32
	SetAnnotationStatusListStringFilter(v string) *ListAnnotationMissionRequest
	GetAnnotationStatusListStringFilter() *string
	SetCreateTimeEndFilter(v int64) *ListAnnotationMissionRequest
	GetCreateTimeEndFilter() *int64
	SetCreateTimeStartFilter(v int64) *ListAnnotationMissionRequest
	GetCreateTimeStartFilter() *int64
	SetInstanceId(v string) *ListAnnotationMissionRequest
	GetInstanceId() *string
	SetPageIndex(v int32) *ListAnnotationMissionRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListAnnotationMissionRequest
	GetPageSize() *int32
}

type ListAnnotationMissionRequest struct {
	// Annotation job ID
	//
	// example:
	//
	// 0943abcb-bd7d-4ace-8cf7-97d39d4dd0b9
	AnnotationMissionId *string `json:"AnnotationMissionId,omitempty" xml:"AnnotationMissionId,omitempty"`
	// Annotation job name
	//
	// example:
	//
	// 全景服务场景-标注任务-20230316-103253
	AnnotationMissionName *string `json:"AnnotationMissionName,omitempty" xml:"AnnotationMissionName,omitempty"`
	// Annotation status filter condition [Historical parameter, deprecated]
	AnnotationStatusListFilter []*int32 `json:"AnnotationStatusListFilter,omitempty" xml:"AnnotationStatusListFilter,omitempty" type:"Repeated"`
	// Filter condition for annotation status
	//
	// > Format: [1]. Multiple choice is supported. The meanings of specific array values are as follows:
	//
	// - 1: In progress
	//
	// - 2: Completed
	//
	// - 3: Shutdown
	//
	// example:
	//
	// [1]
	AnnotationStatusListStringFilter *string `json:"AnnotationStatusListStringFilter,omitempty" xml:"AnnotationStatusListStringFilter,omitempty"`
	// End time for filtering annotation jobs by creation time
	//
	// example:
	//
	// 1673280000000
	CreateTimeEndFilter *int64 `json:"CreateTimeEndFilter,omitempty" xml:"CreateTimeEndFilter,omitempty"`
	// Annotation job creation filter start time
	//
	// example:
	//
	// 1661961600000
	CreateTimeStartFilter *int64 `json:"CreateTimeStartFilter,omitempty" xml:"CreateTimeStartFilter,omitempty"`
	// Instance ID
	//
	// example:
	//
	// 191ef468-75a2-4004-9441-a5c31bf5cd9d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Number of entries displayed per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAnnotationMissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnnotationMissionRequest) GoString() string {
	return s.String()
}

func (s *ListAnnotationMissionRequest) GetAnnotationMissionId() *string {
	return s.AnnotationMissionId
}

func (s *ListAnnotationMissionRequest) GetAnnotationMissionName() *string {
	return s.AnnotationMissionName
}

func (s *ListAnnotationMissionRequest) GetAnnotationStatusListFilter() []*int32 {
	return s.AnnotationStatusListFilter
}

func (s *ListAnnotationMissionRequest) GetAnnotationStatusListStringFilter() *string {
	return s.AnnotationStatusListStringFilter
}

func (s *ListAnnotationMissionRequest) GetCreateTimeEndFilter() *int64 {
	return s.CreateTimeEndFilter
}

func (s *ListAnnotationMissionRequest) GetCreateTimeStartFilter() *int64 {
	return s.CreateTimeStartFilter
}

func (s *ListAnnotationMissionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnnotationMissionRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListAnnotationMissionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAnnotationMissionRequest) SetAnnotationMissionId(v string) *ListAnnotationMissionRequest {
	s.AnnotationMissionId = &v
	return s
}

func (s *ListAnnotationMissionRequest) SetAnnotationMissionName(v string) *ListAnnotationMissionRequest {
	s.AnnotationMissionName = &v
	return s
}

func (s *ListAnnotationMissionRequest) SetAnnotationStatusListFilter(v []*int32) *ListAnnotationMissionRequest {
	s.AnnotationStatusListFilter = v
	return s
}

func (s *ListAnnotationMissionRequest) SetAnnotationStatusListStringFilter(v string) *ListAnnotationMissionRequest {
	s.AnnotationStatusListStringFilter = &v
	return s
}

func (s *ListAnnotationMissionRequest) SetCreateTimeEndFilter(v int64) *ListAnnotationMissionRequest {
	s.CreateTimeEndFilter = &v
	return s
}

func (s *ListAnnotationMissionRequest) SetCreateTimeStartFilter(v int64) *ListAnnotationMissionRequest {
	s.CreateTimeStartFilter = &v
	return s
}

func (s *ListAnnotationMissionRequest) SetInstanceId(v string) *ListAnnotationMissionRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAnnotationMissionRequest) SetPageIndex(v int32) *ListAnnotationMissionRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAnnotationMissionRequest) SetPageSize(v int32) *ListAnnotationMissionRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnnotationMissionRequest) Validate() error {
	return dara.Validate(s)
}
