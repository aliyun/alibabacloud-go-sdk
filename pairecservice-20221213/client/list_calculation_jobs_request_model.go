// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalculationJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCalculationJobsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListCalculationJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCalculationJobsRequest
	GetPageSize() *int32
	SetSceneId(v string) *ListCalculationJobsRequest
	GetSceneId() *string
	SetStatus(v string) *ListCalculationJobsRequest
	GetStatus() *string
}

type ListCalculationJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCalculationJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCalculationJobsRequest) GoString() string {
	return s.String()
}

func (s *ListCalculationJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCalculationJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCalculationJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCalculationJobsRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListCalculationJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCalculationJobsRequest) SetInstanceId(v string) *ListCalculationJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCalculationJobsRequest) SetPageNumber(v int32) *ListCalculationJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCalculationJobsRequest) SetPageSize(v int32) *ListCalculationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCalculationJobsRequest) SetSceneId(v string) *ListCalculationJobsRequest {
	s.SceneId = &v
	return s
}

func (s *ListCalculationJobsRequest) SetStatus(v string) *ListCalculationJobsRequest {
	s.Status = &v
	return s
}

func (s *ListCalculationJobsRequest) Validate() error {
	return dara.Validate(s)
}
