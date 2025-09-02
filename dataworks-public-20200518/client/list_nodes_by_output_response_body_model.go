// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesByOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListNodesByOutputResponseBodyData) *ListNodesByOutputResponseBody
	GetData() []*ListNodesByOutputResponseBodyData
	SetErrorCode(v string) *ListNodesByOutputResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListNodesByOutputResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListNodesByOutputResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListNodesByOutputResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNodesByOutputResponseBody
	GetSuccess() *bool
}

type ListNodesByOutputResponseBody struct {
	// The nodes returned.
	Data []*ListNodesByOutputResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// SDFSDFSDF-asdfDFSDF-SDFSDf-SDfSFD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodesByOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByOutputResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesByOutputResponseBody) GetData() []*ListNodesByOutputResponseBodyData {
	return s.Data
}

func (s *ListNodesByOutputResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNodesByOutputResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListNodesByOutputResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNodesByOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesByOutputResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNodesByOutputResponseBody) SetData(v []*ListNodesByOutputResponseBodyData) *ListNodesByOutputResponseBody {
	s.Data = v
	return s
}

func (s *ListNodesByOutputResponseBody) SetErrorCode(v string) *ListNodesByOutputResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNodesByOutputResponseBody) SetErrorMessage(v string) *ListNodesByOutputResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListNodesByOutputResponseBody) SetHttpStatusCode(v int32) *ListNodesByOutputResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodesByOutputResponseBody) SetRequestId(v string) *ListNodesByOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesByOutputResponseBody) SetSuccess(v bool) *ListNodesByOutputResponseBody {
	s.Success = &v
	return s
}

func (s *ListNodesByOutputResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodesByOutputResponseBodyData struct {
	// The information about the nodes returned.
	NodeList []*ListNodesByOutputResponseBodyDataNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
	// The output name of the current node.
	//
	// example:
	//
	// test_0709_1.630003556_out
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesByOutputResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByOutputResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodesByOutputResponseBodyData) GetNodeList() []*ListNodesByOutputResponseBodyDataNodeList {
	return s.NodeList
}

func (s *ListNodesByOutputResponseBodyData) GetOutput() *string {
	return s.Output
}

func (s *ListNodesByOutputResponseBodyData) SetNodeList(v []*ListNodesByOutputResponseBodyDataNodeList) *ListNodesByOutputResponseBodyData {
	s.NodeList = v
	return s
}

func (s *ListNodesByOutputResponseBodyData) SetOutput(v string) *ListNodesByOutputResponseBodyData {
	s.Output = &v
	return s
}

func (s *ListNodesByOutputResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListNodesByOutputResponseBodyDataNodeList struct {
	// The baseline ID.
	//
	// example:
	//
	// 1235667
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the data source.
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
	// The description of the node.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The table and partition filter expression in Data Quality that are associated with the node.
	//
	// example:
	//
	// [{"projectName":"ztjy_dim","tableName":"dim_user_agent_manage_area_a","partition":"ds\\u003d$[yyyy-mm-dd-1]"}]
	DqcDescription *string `json:"DqcDescription,omitempty" xml:"DqcDescription,omitempty"`
	// Indicates whether the node is associated with a monitoring rule in Data Quality. Valid values: 0 and 1. The value 0 indicates that the node is associated with a monitoring rule in Data Quality. The value 1 indicates that the node is not associated with a monitoring rule in Data Quality.
	//
	// example:
	//
	// 1
	DqcType *int32 `json:"DqcType,omitempty" xml:"DqcType,omitempty"`
	// The node type. Valid values:
	//
	// 6 (Shell), 10 (ODPS SQL), 11 (ODPS MR), 23 (Data Integration), 24 (ODPS Script), 99 (zero load), 221 (PyODPS 2), 225 (ODPS Spark), 227 (EMR Hive), 228 (EMR Spark), 229 (EMR Spark SQL), 230 (EMR MR), 239 (OSS object inspection), 257 (EMR Shell), 258 (EMR Spark Shell), 259 (EMR Presto), 260 (EMR Impala), 900 (real-time synchronization), 1089 (cross-tenant collaboration), 1091 (Hologres development), 1093 (Hologres SQL), 1100 (assignment), and 1221 (PyODPS 3)
	//
	// example:
	//
	// ODPS_SQL
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 125677
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// liux_test_n****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 19337906836551
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The scheduling parameters of the node.
	//
	// example:
	//
	// a=b
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The priority of the node. Valid values: 1, 3, 5, 7, and 8. A greater value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The node type. This parameter is deprecated. For more information about node types, see valid values of the FileType parameter.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 33671
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the workflow to which the node belongs.
	//
	// example:
	//
	// 1235655464
	RelatedFlowId *int64 `json:"RelatedFlowId,omitempty" xml:"RelatedFlowId,omitempty"`
	// The interval at which the node is rerun after the node fails to run.
	//
	// example:
	//
	// 60
	RepeatInterval *int32 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// Indicates whether the node can be rerun.
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
	// The scheduling type of the node. Valid values:
	//
	// 	- NORMAL: The node is an auto triggered node. The scheduling system regularly runs the node.
	//
	// 	- MANUAL: The node is a manually triggered node. The scheduling system does not regularly run the node.
	//
	// 	- PAUSE: The node is a frozen node. The scheduling system regularly runs the node but sets the status of the node to failed when the scheduling system starts to run the node.
	//
	// 	- SKIP: The node is a dry-run node. The scheduling system regularly runs the node but sets the status of the node to successful when the scheduling system starts to run the node.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s ListNodesByOutputResponseBodyDataNodeList) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByOutputResponseBodyDataNodeList) GoString() string {
	return s.String()
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetConnection() *string {
	return s.Connection
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetCronExpress() *string {
	return s.CronExpress
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetDescription() *string {
	return s.Description
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetDqcDescription() *string {
	return s.DqcDescription
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetDqcType() *int32 {
	return s.DqcType
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetFileType() *string {
	return s.FileType
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetNodeName() *string {
	return s.NodeName
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetParamValues() *string {
	return s.ParamValues
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetProgramType() *string {
	return s.ProgramType
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetRelatedFlowId() *int64 {
	return s.RelatedFlowId
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetRepeatInterval() *int32 {
	return s.RepeatInterval
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetRepeatability() *bool {
	return s.Repeatability
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetResGroupName() *string {
	return s.ResGroupName
}

func (s *ListNodesByOutputResponseBodyDataNodeList) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetBaselineId(v int64) *ListNodesByOutputResponseBodyDataNodeList {
	s.BaselineId = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetConnection(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.Connection = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetCronExpress(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.CronExpress = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetDescription(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.Description = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetDqcDescription(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.DqcDescription = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetDqcType(v int32) *ListNodesByOutputResponseBodyDataNodeList {
	s.DqcType = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetFileType(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.FileType = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetNodeId(v int64) *ListNodesByOutputResponseBodyDataNodeList {
	s.NodeId = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetNodeName(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.NodeName = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetOwnerId(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.OwnerId = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetParamValues(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.ParamValues = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetPriority(v int32) *ListNodesByOutputResponseBodyDataNodeList {
	s.Priority = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetProgramType(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.ProgramType = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetProjectId(v int64) *ListNodesByOutputResponseBodyDataNodeList {
	s.ProjectId = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetRelatedFlowId(v int64) *ListNodesByOutputResponseBodyDataNodeList {
	s.RelatedFlowId = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetRepeatInterval(v int32) *ListNodesByOutputResponseBodyDataNodeList {
	s.RepeatInterval = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetRepeatability(v bool) *ListNodesByOutputResponseBodyDataNodeList {
	s.Repeatability = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetResGroupName(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.ResGroupName = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) SetSchedulerType(v string) *ListNodesByOutputResponseBodyDataNodeList {
	s.SchedulerType = &v
	return s
}

func (s *ListNodesByOutputResponseBodyDataNodeList) Validate() error {
	return dara.Validate(s)
}
