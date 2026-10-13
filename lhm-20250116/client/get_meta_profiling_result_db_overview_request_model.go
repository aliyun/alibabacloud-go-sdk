// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaProfilingResultDbOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsSummary(v string) *GetMetaProfilingResultDbOverviewRequest
	GetIsSummary() *string
	SetTaskId(v string) *GetMetaProfilingResultDbOverviewRequest
	GetTaskId() *string
}

type GetMetaProfilingResultDbOverviewRequest struct {
	// Specifies whether to return only summary data.
	//
	// example:
	//
	// false
	IsSummary *string `json:"isSummary,omitempty" xml:"isSummary,omitempty"`
	// The ID of the metadata profiling task. It is used to locate the profiling task for which you want to query the overview.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetMetaProfilingResultDbOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaProfilingResultDbOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetMetaProfilingResultDbOverviewRequest) GetIsSummary() *string {
	return s.IsSummary
}

func (s *GetMetaProfilingResultDbOverviewRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMetaProfilingResultDbOverviewRequest) SetIsSummary(v string) *GetMetaProfilingResultDbOverviewRequest {
	s.IsSummary = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewRequest) SetTaskId(v string) *GetMetaProfilingResultDbOverviewRequest {
	s.TaskId = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewRequest) Validate() error {
	return dara.Validate(s)
}
