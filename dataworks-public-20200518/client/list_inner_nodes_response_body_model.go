// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInnerNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPaging(v *ListInnerNodesResponseBodyPaging) *ListInnerNodesResponseBody
	GetPaging() *ListInnerNodesResponseBodyPaging
	SetRequestId(v string) *ListInnerNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInnerNodesResponseBody
	GetSuccess() *bool
}

type ListInnerNodesResponseBody struct {
	// The pagination information.
	Paging *ListInnerNodesResponseBodyPaging `json:"Paging,omitempty" xml:"Paging,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInnerNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInnerNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInnerNodesResponseBody) GetPaging() *ListInnerNodesResponseBodyPaging {
	return s.Paging
}

func (s *ListInnerNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInnerNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInnerNodesResponseBody) SetPaging(v *ListInnerNodesResponseBodyPaging) *ListInnerNodesResponseBody {
	s.Paging = v
	return s
}

func (s *ListInnerNodesResponseBody) SetRequestId(v string) *ListInnerNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInnerNodesResponseBody) SetSuccess(v bool) *ListInnerNodesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInnerNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInnerNodesResponseBodyPaging struct {
	// The list of inner nodes.
	Nodes []*ListInnerNodesResponseBodyPagingNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of inner nodes returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInnerNodesResponseBodyPaging) String() string {
	return dara.Prettify(s)
}

func (s ListInnerNodesResponseBodyPaging) GoString() string {
	return s.String()
}

func (s *ListInnerNodesResponseBodyPaging) GetNodes() []*ListInnerNodesResponseBodyPagingNodes {
	return s.Nodes
}

func (s *ListInnerNodesResponseBodyPaging) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInnerNodesResponseBodyPaging) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInnerNodesResponseBodyPaging) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInnerNodesResponseBodyPaging) SetNodes(v []*ListInnerNodesResponseBodyPagingNodes) *ListInnerNodesResponseBodyPaging {
	s.Nodes = v
	return s
}

func (s *ListInnerNodesResponseBodyPaging) SetPageNumber(v int32) *ListInnerNodesResponseBodyPaging {
	s.PageNumber = &v
	return s
}

func (s *ListInnerNodesResponseBodyPaging) SetPageSize(v int32) *ListInnerNodesResponseBodyPaging {
	s.PageSize = &v
	return s
}

func (s *ListInnerNodesResponseBodyPaging) SetTotalCount(v int32) *ListInnerNodesResponseBodyPaging {
	s.TotalCount = &v
	return s
}

func (s *ListInnerNodesResponseBodyPaging) Validate() error {
	return dara.Validate(s)
}

type ListInnerNodesResponseBodyPagingNodes struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 123
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The connection string.
	//
	// example:
	//
	// odps_first
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The CRON expression.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The description of the inner node.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The table and partition filter expression in Data Quality that are associated with the inner node.
	//
	// example:
	//
	// [{"projectName":"ztjy_dim","tableName":"dim_user_agent_manage_area_a","partition":"ds\\u003d$[yyyy-mm-dd-1]"}]
	DqcDescription *string `json:"DqcDescription,omitempty" xml:"DqcDescription,omitempty"`
	// Indicates whether the inner node is associated with a monitoring rule in Data Quality. Valid values: 0 and 1. The value 0 indicates that the inner node is associated with a monitoring rule in Data Quality. The value 1 indicates that the inner node is not associated with a monitoring rule in Data Quality.
	//
	// example:
	//
	// 1
	DqcType *string `json:"DqcType,omitempty" xml:"DqcType,omitempty"`
	// The inner node ID.
	//
	// example:
	//
	// 12
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the inner node.
	//
	// example:
	//
	// liux_test_n****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 1933****36551
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The additional parameters.
	//
	// example:
	//
	// a=b
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The priority of the inner node. Valid values: 1, 3, 5, 7, and 8.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the inner node.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The interval at which the inner node is rerun after the inner node fails to run.
	//
	// example:
	//
	// 60
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// Indicates whether the inner node can be rerun.
	//
	// example:
	//
	// true
	Repeatability *bool `json:"Repeatability,omitempty" xml:"Repeatability,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// Default Resource Group
	ResGroupName *string `json:"ResGroupName,omitempty" xml:"ResGroupName,omitempty"`
	// The scheduling type of the inner node. Valid values:
	//
	// 	- NORMAL: The inner node is an auto triggered node.
	//
	// 	- MANUAL: The inner node is a manually triggered node. The scheduling system does not run the node on a regular basis.
	//
	// 	- PAUSE: The inner node is a paused node.
	//
	// 	- SKIP: The inner node is a dry-run node. Dry-run nodes are started as scheduled, but the scheduling system sets the status of the nodes to successful when it starts to run them.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s ListInnerNodesResponseBodyPagingNodes) String() string {
	return dara.Prettify(s)
}

func (s ListInnerNodesResponseBodyPagingNodes) GoString() string {
	return s.String()
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetConnection() *string {
	return s.Connection
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetCronExpress() *string {
	return s.CronExpress
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetDescription() *string {
	return s.Description
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetDqcDescription() *string {
	return s.DqcDescription
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetDqcType() *string {
	return s.DqcType
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetParamValues() *string {
	return s.ParamValues
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetPriority() *int32 {
	return s.Priority
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetProgramType() *string {
	return s.ProgramType
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetRepeatability() *bool {
	return s.Repeatability
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetResGroupName() *string {
	return s.ResGroupName
}

func (s *ListInnerNodesResponseBodyPagingNodes) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetBaselineId(v int64) *ListInnerNodesResponseBodyPagingNodes {
	s.BaselineId = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetBusinessId(v int64) *ListInnerNodesResponseBodyPagingNodes {
	s.BusinessId = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetConnection(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.Connection = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetCronExpress(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.CronExpress = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetDescription(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.Description = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetDqcDescription(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.DqcDescription = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetDqcType(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.DqcType = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetNodeId(v int64) *ListInnerNodesResponseBodyPagingNodes {
	s.NodeId = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetNodeName(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.NodeName = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetOwnerId(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.OwnerId = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetParamValues(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.ParamValues = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetPriority(v int32) *ListInnerNodesResponseBodyPagingNodes {
	s.Priority = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetProgramType(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.ProgramType = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetProjectId(v int64) *ListInnerNodesResponseBodyPagingNodes {
	s.ProjectId = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetRepeatInterval(v int64) *ListInnerNodesResponseBodyPagingNodes {
	s.RepeatInterval = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetRepeatability(v bool) *ListInnerNodesResponseBodyPagingNodes {
	s.Repeatability = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetResGroupName(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.ResGroupName = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) SetSchedulerType(v string) *ListInnerNodesResponseBodyPagingNodes {
	s.SchedulerType = &v
	return s
}

func (s *ListInnerNodesResponseBodyPagingNodes) Validate() error {
	return dara.Validate(s)
}
