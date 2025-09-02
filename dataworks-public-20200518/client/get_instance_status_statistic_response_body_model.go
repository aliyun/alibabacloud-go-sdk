// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStatusStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetInstanceStatusStatisticResponseBody
	GetRequestId() *string
	SetStatusCount(v *GetInstanceStatusStatisticResponseBodyStatusCount) *GetInstanceStatusStatisticResponseBody
	GetStatusCount() *GetInstanceStatusStatisticResponseBodyStatusCount
}

type GetInstanceStatusStatisticResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The numbers of instances in different states.
	StatusCount *GetInstanceStatusStatisticResponseBodyStatusCount `json:"StatusCount,omitempty" xml:"StatusCount,omitempty" type:"Struct"`
}

func (s GetInstanceStatusStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStatusStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceStatusStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceStatusStatisticResponseBody) GetStatusCount() *GetInstanceStatusStatisticResponseBodyStatusCount {
	return s.StatusCount
}

func (s *GetInstanceStatusStatisticResponseBody) SetRequestId(v string) *GetInstanceStatusStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceStatusStatisticResponseBody) SetStatusCount(v *GetInstanceStatusStatisticResponseBodyStatusCount) *GetInstanceStatusStatisticResponseBody {
	s.StatusCount = v
	return s
}

func (s *GetInstanceStatusStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceStatusStatisticResponseBodyStatusCount struct {
	// The number of instances that failed to run.
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
	// The number of instances that are waiting to run.
	//
	// example:
	//
	// 1
	WaitTimeCount *int32 `json:"WaitTimeCount,omitempty" xml:"WaitTimeCount,omitempty"`
}

func (s GetInstanceStatusStatisticResponseBodyStatusCount) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStatusStatisticResponseBodyStatusCount) GoString() string {
	return s.String()
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) GetFailureCount() *int32 {
	return s.FailureCount
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) GetNotRunCount() *int32 {
	return s.NotRunCount
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) GetRunningCount() *int32 {
	return s.RunningCount
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) GetWaitResCount() *int32 {
	return s.WaitResCount
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) GetWaitTimeCount() *int32 {
	return s.WaitTimeCount
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) SetFailureCount(v int32) *GetInstanceStatusStatisticResponseBodyStatusCount {
	s.FailureCount = &v
	return s
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) SetNotRunCount(v int32) *GetInstanceStatusStatisticResponseBodyStatusCount {
	s.NotRunCount = &v
	return s
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) SetRunningCount(v int32) *GetInstanceStatusStatisticResponseBodyStatusCount {
	s.RunningCount = &v
	return s
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) SetSuccessCount(v int32) *GetInstanceStatusStatisticResponseBodyStatusCount {
	s.SuccessCount = &v
	return s
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) SetTotalCount(v int32) *GetInstanceStatusStatisticResponseBodyStatusCount {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) SetWaitResCount(v int32) *GetInstanceStatusStatisticResponseBodyStatusCount {
	s.WaitResCount = &v
	return s
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) SetWaitTimeCount(v int32) *GetInstanceStatusStatisticResponseBodyStatusCount {
	s.WaitTimeCount = &v
	return s
}

func (s *GetInstanceStatusStatisticResponseBodyStatusCount) Validate() error {
	return dara.Validate(s)
}
