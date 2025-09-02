// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCycleDagNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNoticeType(v string) *RunCycleDagNodesRequest
	GetAlertNoticeType() *string
	SetAlertType(v string) *RunCycleDagNodesRequest
	GetAlertType() *string
	SetBizBeginTime(v string) *RunCycleDagNodesRequest
	GetBizBeginTime() *string
	SetBizEndTime(v string) *RunCycleDagNodesRequest
	GetBizEndTime() *string
	SetConcurrentRuns(v int32) *RunCycleDagNodesRequest
	GetConcurrentRuns() *int32
	SetEndBizDate(v string) *RunCycleDagNodesRequest
	GetEndBizDate() *string
	SetExcludeNodeIds(v string) *RunCycleDagNodesRequest
	GetExcludeNodeIds() *string
	SetIncludeNodeIds(v string) *RunCycleDagNodesRequest
	GetIncludeNodeIds() *string
	SetName(v string) *RunCycleDagNodesRequest
	GetName() *string
	SetNodeParams(v string) *RunCycleDagNodesRequest
	GetNodeParams() *string
	SetParallelism(v bool) *RunCycleDagNodesRequest
	GetParallelism() *bool
	SetProjectEnv(v string) *RunCycleDagNodesRequest
	GetProjectEnv() *string
	SetRootNodeId(v int64) *RunCycleDagNodesRequest
	GetRootNodeId() *int64
	SetStartBizDate(v string) *RunCycleDagNodesRequest
	GetStartBizDate() *string
	SetStartFutureInstanceImmediately(v bool) *RunCycleDagNodesRequest
	GetStartFutureInstanceImmediately() *bool
}

type RunCycleDagNodesRequest struct {
	// The alert notification method. Valid values:
	//
	// 	- SMS
	//
	// 	- MAIL
	//
	// 	- SMS_MAIL
	//
	// example:
	//
	// SMS
	AlertNoticeType *string `json:"AlertNoticeType,omitempty" xml:"AlertNoticeType,omitempty"`
	// The alert type. Valid values:
	//
	// 	- SUCCESS: An alert is generated when data backfill succeeds.
	//
	// 	- FAILURE: An alert is generated when data backfill fails.
	//
	// 	- SUCCESS_FAILURE: An alert is generated regardless of whether data backfill succeeds or fails.
	//
	// example:
	//
	// FAILURE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The time when the node starts to run. This parameter is required only for auto triggered nodes that are scheduled by hour. Specify the value in the HH:mm:ss format. Valid values: 00:00:00 to 23:59:59.
	//
	// example:
	//
	// 00:00:00
	BizBeginTime *string `json:"BizBeginTime,omitempty" xml:"BizBeginTime,omitempty"`
	// The time when the node stops running. This parameter is required only for auto triggered nodes that are scheduled by hour. Specify the value in the HH:mm:ss format. Valid values: 00:00:00 to 23:59:59.
	//
	// example:
	//
	// 01:00:00
	BizEndTime *string `json:"BizEndTime,omitempty" xml:"BizEndTime,omitempty"`
	// The number of nodes that can run in parallel. Valid values: 2 to 10.
	//
	// example:
	//
	// 5
	ConcurrentRuns *int32 `json:"ConcurrentRuns,omitempty" xml:"ConcurrentRuns,omitempty"`
	// The data timestamp at which data is no longer backfilled. Specify the value in the yyyy-MM-dd 00:00:00 format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-21 00:00:00
	EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	// The IDs of the nodes for which no data needs to be backfilled. The system generates dry-run instances for all these nodes. After these dry-run instances are scheduled, the statuses of these instances are directly set to successful, but the script is not run.
	//
	// example:
	//
	// 1234,123465
	ExcludeNodeIds *string `json:"ExcludeNodeIds,omitempty" xml:"ExcludeNodeIds,omitempty"`
	// The ID of the node for which you want to backfill data. If you want to backfill data for multiple nodes, separate the IDs of the nodes with commas (,). You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to obtain the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74324,74325
	IncludeNodeIds *string `json:"IncludeNodeIds,omitempty" xml:"IncludeNodeIds,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// xm_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters that need to be configured for the node. Set this parameter to a JSON string. The key indicates the ID of the node, and the value indicates the actual values of the parameters.
	//
	// example:
	//
	// {74324:"a=123 b=456"}
	NodeParams *string `json:"NodeParams,omitempty" xml:"NodeParams,omitempty"`
	// Specifies whether data can be backfilled for multiple nodes at the same time.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Parallelism *bool `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV. The value PROD indicates the production environment, and the value DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The ID of the node for which data is first backfilled. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to obtain the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RootNodeId *int64 `json:"RootNodeId,omitempty" xml:"RootNodeId,omitempty"`
	// The data timestamp at which data starts to be backfilled. Specify the value in the yyyy-MM-dd 00:00:00 format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-20 00:00:00
	StartBizDate *string `json:"StartBizDate,omitempty" xml:"StartBizDate,omitempty"`
	// Specifies whether to immediately run an instance that is scheduled to run in the future. If you set this parameter to true, the instance that is scheduled to run in the future is run immediately. Otherwise, the instance is run as scheduled.
	//
	// example:
	//
	// false
	StartFutureInstanceImmediately *bool `json:"StartFutureInstanceImmediately,omitempty" xml:"StartFutureInstanceImmediately,omitempty"`
}

func (s RunCycleDagNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCycleDagNodesRequest) GoString() string {
	return s.String()
}

func (s *RunCycleDagNodesRequest) GetAlertNoticeType() *string {
	return s.AlertNoticeType
}

func (s *RunCycleDagNodesRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *RunCycleDagNodesRequest) GetBizBeginTime() *string {
	return s.BizBeginTime
}

func (s *RunCycleDagNodesRequest) GetBizEndTime() *string {
	return s.BizEndTime
}

func (s *RunCycleDagNodesRequest) GetConcurrentRuns() *int32 {
	return s.ConcurrentRuns
}

func (s *RunCycleDagNodesRequest) GetEndBizDate() *string {
	return s.EndBizDate
}

func (s *RunCycleDagNodesRequest) GetExcludeNodeIds() *string {
	return s.ExcludeNodeIds
}

func (s *RunCycleDagNodesRequest) GetIncludeNodeIds() *string {
	return s.IncludeNodeIds
}

func (s *RunCycleDagNodesRequest) GetName() *string {
	return s.Name
}

func (s *RunCycleDagNodesRequest) GetNodeParams() *string {
	return s.NodeParams
}

func (s *RunCycleDagNodesRequest) GetParallelism() *bool {
	return s.Parallelism
}

func (s *RunCycleDagNodesRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *RunCycleDagNodesRequest) GetRootNodeId() *int64 {
	return s.RootNodeId
}

func (s *RunCycleDagNodesRequest) GetStartBizDate() *string {
	return s.StartBizDate
}

func (s *RunCycleDagNodesRequest) GetStartFutureInstanceImmediately() *bool {
	return s.StartFutureInstanceImmediately
}

func (s *RunCycleDagNodesRequest) SetAlertNoticeType(v string) *RunCycleDagNodesRequest {
	s.AlertNoticeType = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetAlertType(v string) *RunCycleDagNodesRequest {
	s.AlertType = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetBizBeginTime(v string) *RunCycleDagNodesRequest {
	s.BizBeginTime = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetBizEndTime(v string) *RunCycleDagNodesRequest {
	s.BizEndTime = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetConcurrentRuns(v int32) *RunCycleDagNodesRequest {
	s.ConcurrentRuns = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetEndBizDate(v string) *RunCycleDagNodesRequest {
	s.EndBizDate = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetExcludeNodeIds(v string) *RunCycleDagNodesRequest {
	s.ExcludeNodeIds = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetIncludeNodeIds(v string) *RunCycleDagNodesRequest {
	s.IncludeNodeIds = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetName(v string) *RunCycleDagNodesRequest {
	s.Name = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetNodeParams(v string) *RunCycleDagNodesRequest {
	s.NodeParams = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetParallelism(v bool) *RunCycleDagNodesRequest {
	s.Parallelism = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetProjectEnv(v string) *RunCycleDagNodesRequest {
	s.ProjectEnv = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetRootNodeId(v int64) *RunCycleDagNodesRequest {
	s.RootNodeId = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetStartBizDate(v string) *RunCycleDagNodesRequest {
	s.StartBizDate = &v
	return s
}

func (s *RunCycleDagNodesRequest) SetStartFutureInstanceImmediately(v bool) *RunCycleDagNodesRequest {
	s.StartFutureInstanceImmediately = &v
	return s
}

func (s *RunCycleDagNodesRequest) Validate() error {
	return dara.Validate(s)
}
