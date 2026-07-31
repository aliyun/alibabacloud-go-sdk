// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterHealthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeDBClusterHealthStatusResponseBody
	GetAccessDeniedDetail() *string
	SetCS(v *DescribeDBClusterHealthStatusResponseBodyCS) *DescribeDBClusterHealthStatusResponseBody
	GetCS() *DescribeDBClusterHealthStatusResponseBodyCS
	SetExecutor(v *DescribeDBClusterHealthStatusResponseBodyExecutor) *DescribeDBClusterHealthStatusResponseBody
	GetExecutor() *DescribeDBClusterHealthStatusResponseBodyExecutor
	SetInstanceStatus(v string) *DescribeDBClusterHealthStatusResponseBody
	GetInstanceStatus() *string
	SetRequestId(v string) *DescribeDBClusterHealthStatusResponseBody
	GetRequestId() *string
	SetWorker(v *DescribeDBClusterHealthStatusResponseBodyWorker) *DescribeDBClusterHealthStatusResponseBody
	GetWorker() *DescribeDBClusterHealthStatusResponseBodyWorker
}

type DescribeDBClusterHealthStatusResponseBody struct {
	// Details of the authentication failure.
	//
	// example:
	//
	// Authentication failed.
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The health status of the instance access nodes.
	CS *DescribeDBClusterHealthStatusResponseBodyCS `json:"CS,omitempty" xml:"CS,omitempty" type:"Struct"`
	// The health status of the executor groups.
	Executor *DescribeDBClusterHealthStatusResponseBodyExecutor `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	// The health status of the cluster. Valid values:
	//
	// - **RISK**: The cluster is at risk.
	//
	// - **NORMAL**: The cluster is healthy.
	//
	// - **UNAVAILABLE**: The cluster is unavailable.
	//
	// > The cluster health status is considered **NORMAL*	- only if the instance access nodes, executor groups, and worker node groups are all **NORMAL**, and the instance is responsive. If any of these components has a **RISK*	- status, the cluster status is **RISK**. If any component has an **UNAVAILABLE*	- status, the cluster status is **UNAVAILABLE**.
	//
	// example:
	//
	// NORMAL
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CAV
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The health status of the worker node groups.
	Worker *DescribeDBClusterHealthStatusResponseBodyWorker `json:"Worker,omitempty" xml:"Worker,omitempty" type:"Struct"`
}

func (s DescribeDBClusterHealthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeDBClusterHealthStatusResponseBody) GetCS() *DescribeDBClusterHealthStatusResponseBodyCS {
	return s.CS
}

func (s *DescribeDBClusterHealthStatusResponseBody) GetExecutor() *DescribeDBClusterHealthStatusResponseBodyExecutor {
	return s.Executor
}

func (s *DescribeDBClusterHealthStatusResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeDBClusterHealthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterHealthStatusResponseBody) GetWorker() *DescribeDBClusterHealthStatusResponseBodyWorker {
	return s.Worker
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetAccessDeniedDetail(v string) *DescribeDBClusterHealthStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetCS(v *DescribeDBClusterHealthStatusResponseBodyCS) *DescribeDBClusterHealthStatusResponseBody {
	s.CS = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetExecutor(v *DescribeDBClusterHealthStatusResponseBodyExecutor) *DescribeDBClusterHealthStatusResponseBody {
	s.Executor = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetInstanceStatus(v string) *DescribeDBClusterHealthStatusResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetRequestId(v string) *DescribeDBClusterHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) SetWorker(v *DescribeDBClusterHealthStatusResponseBodyWorker) *DescribeDBClusterHealthStatusResponseBody {
	s.Worker = v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBody) Validate() error {
	if s.CS != nil {
		if err := s.CS.Validate(); err != nil {
			return err
		}
	}
	if s.Executor != nil {
		if err := s.Executor.Validate(); err != nil {
			return err
		}
	}
	if s.Worker != nil {
		if err := s.Worker.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterHealthStatusResponseBodyCS struct {
	// The number of healthy instance access nodes.
	//
	// example:
	//
	// 0
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of instance access nodes.
	//
	// example:
	//
	// 1
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of instance access nodes at risk.
	//
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health status of the instance access nodes. Valid values:
	//
	// - **RISK**: The instance access nodes are at risk.
	//
	// - **NORMAL**: The instance access nodes are healthy.
	//
	// - **UNAVAILABLE**: The instance access nodes are unavailable.
	//
	// example:
	//
	// UNAVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable instance access nodes.
	//
	// example:
	//
	// 1
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyCS) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyCS) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) GetExpectedCount() *int64 {
	return s.ExpectedCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) GetUnavailableCount() *int64 {
	return s.UnavailableCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyCS {
	s.UnavailableCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyCS) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterHealthStatusResponseBodyExecutor struct {
	// The number of healthy executor nodes.
	//
	// example:
	//
	// 0
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of executor nodes.
	//
	// example:
	//
	// 1
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of executor nodes at risk.
	//
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health status of the executor groups. Valid values:
	//
	// - **RISK**: The executor groups are at risk.
	//
	// - **NORMAL**: The executor groups are healthy.
	//
	// - **UNAVAILABLE**: The executor groups are unavailable.
	//
	// example:
	//
	// UNAVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable executor nodes.
	//
	// example:
	//
	// 1
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyExecutor) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyExecutor) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) GetExpectedCount() *int64 {
	return s.ExpectedCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) GetUnavailableCount() *int64 {
	return s.UnavailableCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyExecutor {
	s.UnavailableCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyExecutor) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterHealthStatusResponseBodyWorker struct {
	// The number of healthy worker node groups.
	//
	// example:
	//
	// 0
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of worker node groups.
	//
	// example:
	//
	// 1
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of worker node groups at risk.
	//
	// example:
	//
	// 1
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health status of the worker node groups. Valid values:
	//
	// - **RISK**: The worker node groups are at risk.
	//
	// - **NORMAL**: The worker node groups are healthy.
	//
	// - **UNAVAILABLE**: The worker node groups are unavailable.
	//
	// example:
	//
	// UNAVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable worker node groups.
	//
	// example:
	//
	// 1
	UnavailableCount *int64 `json:"UnavailableCount,omitempty" xml:"UnavailableCount,omitempty"`
}

func (s DescribeDBClusterHealthStatusResponseBodyWorker) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterHealthStatusResponseBodyWorker) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) GetExpectedCount() *int64 {
	return s.ExpectedCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) GetUnavailableCount() *int64 {
	return s.UnavailableCount
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetActiveCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetExpectedCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.ExpectedCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetRiskCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.RiskCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetStatus(v string) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.Status = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) SetUnavailableCount(v int64) *DescribeDBClusterHealthStatusResponseBodyWorker {
	s.UnavailableCount = &v
	return s
}

func (s *DescribeDBClusterHealthStatusResponseBodyWorker) Validate() error {
	return dara.Validate(s)
}
