// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsDebugSampleLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetPtsDebugSampleLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetPtsDebugSampleLogsRequest
	GetPageSize() *int32
	SetPlanId(v string) *GetPtsDebugSampleLogsRequest
	GetPlanId() *string
}

type GetPtsDebugSampleLogsRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the debugging task.
	//
	// example:
	//
	// NJJBH8B
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s GetPtsDebugSampleLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPtsDebugSampleLogsRequest) GoString() string {
	return s.String()
}

func (s *GetPtsDebugSampleLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetPtsDebugSampleLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetPtsDebugSampleLogsRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *GetPtsDebugSampleLogsRequest) SetPageNumber(v int32) *GetPtsDebugSampleLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetPtsDebugSampleLogsRequest) SetPageSize(v int32) *GetPtsDebugSampleLogsRequest {
	s.PageSize = &v
	return s
}

func (s *GetPtsDebugSampleLogsRequest) SetPlanId(v string) *GetPtsDebugSampleLogsRequest {
	s.PlanId = &v
	return s
}

func (s *GetPtsDebugSampleLogsRequest) Validate() error {
	return dara.Validate(s)
}
