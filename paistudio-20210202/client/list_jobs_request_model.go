// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListJobsRequest
	GetCreator() *string
	SetExperimentId(v string) *ListJobsRequest
	GetExperimentId() *string
	SetOrder(v string) *ListJobsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
}

type ListJobsRequest struct {
	// example:
	//
	// 13266*******76250
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// draft-8up80bg0k1q23stml6
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListJobsRequest) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListJobsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) SetCreator(v string) *ListJobsRequest {
	s.Creator = &v
	return s
}

func (s *ListJobsRequest) SetExperimentId(v string) *ListJobsRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListJobsRequest) SetOrder(v string) *ListJobsRequest {
	s.Order = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
