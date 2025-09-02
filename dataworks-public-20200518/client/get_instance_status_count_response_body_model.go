// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStatusCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetInstanceStatusCountResponseBody
	GetRequestId() *string
	SetStatusCount(v *GetInstanceStatusCountResponseBodyStatusCount) *GetInstanceStatusCountResponseBody
	GetStatusCount() *GetInstanceStatusCountResponseBodyStatusCount
}

type GetInstanceStatusCountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of instances.
	StatusCount *GetInstanceStatusCountResponseBodyStatusCount `json:"StatusCount,omitempty" xml:"StatusCount,omitempty" type:"Struct"`
}

func (s GetInstanceStatusCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStatusCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceStatusCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceStatusCountResponseBody) GetStatusCount() *GetInstanceStatusCountResponseBodyStatusCount {
	return s.StatusCount
}

func (s *GetInstanceStatusCountResponseBody) SetRequestId(v string) *GetInstanceStatusCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceStatusCountResponseBody) SetStatusCount(v *GetInstanceStatusCountResponseBodyStatusCount) *GetInstanceStatusCountResponseBody {
	s.StatusCount = v
	return s
}

func (s *GetInstanceStatusCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceStatusCountResponseBodyStatusCount struct {
	// The number of instances that failed.
	//
	// example:
	//
	// 1
	FailureCount *int32 `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	// The number of instances that are not run.
	//
	// example:
	//
	// 1
	NotRunCount *int32 `json:"NotRunCount,omitempty" xml:"NotRunCount,omitempty"`
	// The number of instances that are running.
	//
	// example:
	//
	// 1
	RunningCount *int32 `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	// The number of instances that are successfully run.
	//
	// example:
	//
	// 1
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The total number of instances returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of instances that are waiting for resources.
	//
	// example:
	//
	// 1
	WaitResCount *int32 `json:"WaitResCount,omitempty" xml:"WaitResCount,omitempty"`
	// The number of instances that are waiting for their scheduling time to arrive.
	//
	// example:
	//
	// 1
	WaitTimeCount *int32 `json:"WaitTimeCount,omitempty" xml:"WaitTimeCount,omitempty"`
}

func (s GetInstanceStatusCountResponseBodyStatusCount) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStatusCountResponseBodyStatusCount) GoString() string {
	return s.String()
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) GetFailureCount() *int32 {
	return s.FailureCount
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) GetNotRunCount() *int32 {
	return s.NotRunCount
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) GetRunningCount() *int32 {
	return s.RunningCount
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) GetWaitResCount() *int32 {
	return s.WaitResCount
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) GetWaitTimeCount() *int32 {
	return s.WaitTimeCount
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) SetFailureCount(v int32) *GetInstanceStatusCountResponseBodyStatusCount {
	s.FailureCount = &v
	return s
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) SetNotRunCount(v int32) *GetInstanceStatusCountResponseBodyStatusCount {
	s.NotRunCount = &v
	return s
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) SetRunningCount(v int32) *GetInstanceStatusCountResponseBodyStatusCount {
	s.RunningCount = &v
	return s
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) SetSuccessCount(v int32) *GetInstanceStatusCountResponseBodyStatusCount {
	s.SuccessCount = &v
	return s
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) SetTotalCount(v int32) *GetInstanceStatusCountResponseBodyStatusCount {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) SetWaitResCount(v int32) *GetInstanceStatusCountResponseBodyStatusCount {
	s.WaitResCount = &v
	return s
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) SetWaitTimeCount(v int32) *GetInstanceStatusCountResponseBodyStatusCount {
	s.WaitTimeCount = &v
	return s
}

func (s *GetInstanceStatusCountResponseBodyStatusCount) Validate() error {
	return dara.Validate(s)
}
