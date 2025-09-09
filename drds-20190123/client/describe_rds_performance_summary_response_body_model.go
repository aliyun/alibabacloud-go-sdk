// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsPerformanceSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRdsPerformanceInfos(v []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) *DescribeRdsPerformanceSummaryResponseBody
	GetRdsPerformanceInfos() []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos
	SetRequestId(v string) *DescribeRdsPerformanceSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRdsPerformanceSummaryResponseBody
	GetSuccess() *bool
}

type DescribeRdsPerformanceSummaryResponseBody struct {
	// A collection of objects.
	RdsPerformanceInfos []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos `json:"RdsPerformanceInfos,omitempty" xml:"RdsPerformanceInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B6876277-ECFD-4658-AC1E-A7FAD8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponseBody) GetRdsPerformanceInfos() []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	return s.RdsPerformanceInfos
}

func (s *DescribeRdsPerformanceSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdsPerformanceSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetRdsPerformanceInfos(v []*DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) *DescribeRdsPerformanceSummaryResponseBody {
	s.RdsPerformanceInfos = v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetRequestId(v string) *DescribeRdsPerformanceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) SetSuccess(v bool) *DescribeRdsPerformanceSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos struct {
	// The number of active sessions of the RDS instance.
	//
	// example:
	//
	// 0
	ActiveSessions *int32 `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	// The CPU utilization of an RDS instance.
	//
	// example:
	//
	// 0.26
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The IOPS of the RDS instance.
	//
	// example:
	//
	// 17.62
	Iops *float32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The ID of an RDS instance.
	//
	// example:
	//
	// rm-**************
	RdsId *string `json:"RdsId,omitempty" xml:"RdsId,omitempty"`
	// The disk usage of apsaradb for RDS. Unit: MB.
	//
	// example:
	//
	// 4145144777
	SpaceUsage *int64 `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
	// The total number of current RDS sessions.
	//
	// example:
	//
	// 162
	TotalSessions *int32 `json:"TotalSessions,omitempty" xml:"TotalSessions,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GetActiveSessions() *int32 {
	return s.ActiveSessions
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GetIops() *float32 {
	return s.Iops
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GetRdsId() *string {
	return s.RdsId
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GetSpaceUsage() *int64 {
	return s.SpaceUsage
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) GetTotalSessions() *int32 {
	return s.TotalSessions
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetActiveSessions(v int32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.ActiveSessions = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetCpu(v float32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.Cpu = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetIops(v float32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.Iops = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetRdsId(v string) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.RdsId = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetSpaceUsage(v int64) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.SpaceUsage = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) SetTotalSessions(v int32) *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos {
	s.TotalSessions = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponseBodyRdsPerformanceInfos) Validate() error {
	return dara.Validate(s)
}
