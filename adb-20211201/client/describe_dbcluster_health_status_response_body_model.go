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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The access nodes of the queried cluster.
	CS *DescribeDBClusterHealthStatusResponseBodyCS `json:"CS,omitempty" xml:"CS,omitempty" type:"Struct"`
	// The compute node groups of the queried cluster.
	Executor *DescribeDBClusterHealthStatusResponseBodyExecutor `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	// The health state of the cluster. Valid values:
	//
	// 	- **RISK**
	//
	// 	- **NORMAL**
	//
	// 	- **UNAVAILABLE**
	//
	// >  When the states of the access nodes, compute node groups, and storage node groups of a cluster are all **NORMAL*	- and a connection to the cluster is established, the state of the cluster is **NORMAL**. When the state of the access nodes, compute node groups, or storage node groups of the cluster is **RISK**, the state of the cluster is **RISK**. When the state of the access nodes, compute node groups, or storage node groups of the cluster is **UNAVAILABLE**, the state of the cluster is **UNAVAILABLE**.
	//
	// example:
	//
	// NORMAL
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage node groups of the queried cluster.
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
	return dara.Validate(s)
}

type DescribeDBClusterHealthStatusResponseBodyCS struct {
	// The number of healthy access nodes.
	//
	// example:
	//
	// 2
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of access nodes.
	//
	// example:
	//
	// 2
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky nodes.
	//
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of access nodes. Valid values:
	//
	// 	- **RISK**
	//
	// 	- **NORMAL**
	//
	// 	- **UNAVAILABLE**
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable access nodes.
	//
	// example:
	//
	// 0
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
	// The number of healthy access nodes.
	//
	// example:
	//
	// 2
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of compute nodes.
	//
	// example:
	//
	// 2
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky nodes.
	//
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of compute node groups. Valid values:
	//
	// 	- **RISK**
	//
	// 	- **NORMAL**
	//
	// 	- **UNAVAILABLE**
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable access nodes.
	//
	// example:
	//
	// 0
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
	// The number of healthy storage node groups.
	//
	// example:
	//
	// 2
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The total number of storage node groups.
	//
	// example:
	//
	// 2
	ExpectedCount *int64 `json:"ExpectedCount,omitempty" xml:"ExpectedCount,omitempty"`
	// The number of risky storage node groups.
	//
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The health state of storage node groups. Valid values:
	//
	// 	- **RISK**
	//
	// 	- **NORMAL**
	//
	// 	- **UNAVAILABLE**
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of unavailable storage node groups.
	//
	// example:
	//
	// 0
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
