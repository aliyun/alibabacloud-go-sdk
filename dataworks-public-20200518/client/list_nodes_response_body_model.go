// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListNodesResponseBodyData) *ListNodesResponseBody
	GetData() *ListNodesResponseBodyData
	SetErrorCode(v string) *ListNodesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListNodesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListNodesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNodesResponseBody
	GetSuccess() *bool
}

type ListNodesResponseBody struct {
	// The nodes.
	Data *ListNodesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The interval at which the node is rerun after the node fails to run.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The list of nodes.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the node can be rerun.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) GetData() *ListNodesResponseBodyData {
	return s.Data
}

func (s *ListNodesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListNodesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListNodesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNodesResponseBody) SetData(v *ListNodesResponseBodyData) *ListNodesResponseBody {
	s.Data = v
	return s
}

func (s *ListNodesResponseBody) SetErrorCode(v string) *ListNodesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNodesResponseBody) SetErrorMessage(v string) *ListNodesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListNodesResponseBody) SetHttpStatusCode(v int32) *ListNodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetSuccess(v bool) *ListNodesResponseBody {
	s.Success = &v
	return s
}

func (s *ListNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyData struct {
	// The information about the nodes.
	Nodes []*ListNodesResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The name of the node.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The cron expression returned.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// 66
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyData) GetNodes() []*ListNodesResponseBodyDataNodes {
	return s.Nodes
}

func (s *ListNodesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNodesResponseBodyData) SetNodes(v []*ListNodesResponseBodyDataNodes) *ListNodesResponseBodyData {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBodyData) SetPageNumber(v int32) *ListNodesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBodyData) SetPageSize(v int32) *ListNodesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBodyData) SetTotalCount(v int32) *ListNodesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListNodesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyDataNodes struct {
	// The number of the page to return. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 123456
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The operation that you want to perform. Set the value to **ListNodes**.
	//
	// example:
	//
	// 123
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// odps_first
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The timestamp when the node was created. Unit: milliseconds.
	//
	// example:
	//
	// 1593879116000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The timestamp when the node was deployed. Unit: milliseconds.
	//
	// example:
	//
	// 1734537600000
	DeployDate *int64 `json:"DeployDate,omitempty" xml:"DeployDate,omitempty"`
	// The priority for running the node. Valid values: 1, 3, 5, 7, and 8.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the owner.
	//
	// example:
	//
	// [{"projectName":"ztjy_dim","tableName":"dim_user_agent_manage_area_a","partition":"ds\\u003d$[yyyy-mm-dd-1]"}]
	DqcDescription *string `json:"DqcDescription,omitempty" xml:"DqcDescription,omitempty"`
	// The connection string.
	//
	// example:
	//
	// 1
	DqcType *int32 `json:"DqcType,omitempty" xml:"DqcType,omitempty"`
	// The file ID. You can call the ListFiles operation to query the ID.
	//
	// example:
	//
	// 20
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Different file types have different codes. For more information, see [DataWorks node collection](https://help.aliyun.com/document_detail/600169.html).
	//
	// You can also call the [ListFileType](https://help.aliyun.com/document_detail/212428.html) interface to query the code type of the file.
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The latest version number of the file.
	//
	// example:
	//
	// 3
	FileVersion *int32 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// The timestamp when the node was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1593879116000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The types of the nodes. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the type of the node.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The total number of nodes returned.
	//
	// example:
	//
	// liux_test_n****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The additional parameters.
	//
	// example:
	//
	// 19337906836551
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// a=b
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The ID of the owner.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// The information about the nodes.
	//
	// example:
	//
	// 33671
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The table and partition filter expression in Data Quality that are associated with the node.
	//
	// example:
	//
	// 1231123
	RelatedFlowId *int64 `json:"RelatedFlowId,omitempty" xml:"RelatedFlowId,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV.
	//
	// example:
	//
	// 60
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// The rerun mode. The value 0 indicates that rerun can be performed only if a failure occurs. The value 1 indicates that rerun can be performed in all cases. The value 2 indicates that rerun cannot be performed in all cases.
	//
	// example:
	//
	// 1
	RepeatMode *int32 `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// true
	Repeatability *bool `json:"Repeatability,omitempty" xml:"Repeatability,omitempty"`
	// The identifier of the resource group.
	//
	// example:
	//
	// group_123
	ResGroupIdentifier *string `json:"ResGroupIdentifier,omitempty" xml:"ResGroupIdentifier,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	ResGroupName *string `json:"ResGroupName,omitempty" xml:"ResGroupName,omitempty"`
	// The types of the nodes. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the type of the node.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s ListNodesResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyDataNodes) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListNodesResponseBodyDataNodes) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *ListNodesResponseBodyDataNodes) GetConnection() *string {
	return s.Connection
}

func (s *ListNodesResponseBodyDataNodes) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNodesResponseBodyDataNodes) GetCronExpress() *string {
	return s.CronExpress
}

func (s *ListNodesResponseBodyDataNodes) GetDeployDate() *int64 {
	return s.DeployDate
}

func (s *ListNodesResponseBodyDataNodes) GetDescription() *string {
	return s.Description
}

func (s *ListNodesResponseBodyDataNodes) GetDqcDescription() *string {
	return s.DqcDescription
}

func (s *ListNodesResponseBodyDataNodes) GetDqcType() *int32 {
	return s.DqcType
}

func (s *ListNodesResponseBodyDataNodes) GetFileId() *int64 {
	return s.FileId
}

func (s *ListNodesResponseBodyDataNodes) GetFileType() *int32 {
	return s.FileType
}

func (s *ListNodesResponseBodyDataNodes) GetFileVersion() *int32 {
	return s.FileVersion
}

func (s *ListNodesResponseBodyDataNodes) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListNodesResponseBodyDataNodes) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListNodesResponseBodyDataNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *ListNodesResponseBodyDataNodes) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListNodesResponseBodyDataNodes) GetParamValues() *string {
	return s.ParamValues
}

func (s *ListNodesResponseBodyDataNodes) GetPriority() *int32 {
	return s.Priority
}

func (s *ListNodesResponseBodyDataNodes) GetProgramType() *string {
	return s.ProgramType
}

func (s *ListNodesResponseBodyDataNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodesResponseBodyDataNodes) GetRelatedFlowId() *int64 {
	return s.RelatedFlowId
}

func (s *ListNodesResponseBodyDataNodes) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *ListNodesResponseBodyDataNodes) GetRepeatMode() *int32 {
	return s.RepeatMode
}

func (s *ListNodesResponseBodyDataNodes) GetRepeatability() *bool {
	return s.Repeatability
}

func (s *ListNodesResponseBodyDataNodes) GetResGroupIdentifier() *string {
	return s.ResGroupIdentifier
}

func (s *ListNodesResponseBodyDataNodes) GetResGroupName() *string {
	return s.ResGroupName
}

func (s *ListNodesResponseBodyDataNodes) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *ListNodesResponseBodyDataNodes) SetBaselineId(v int64) *ListNodesResponseBodyDataNodes {
	s.BaselineId = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetBusinessId(v int64) *ListNodesResponseBodyDataNodes {
	s.BusinessId = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetConnection(v string) *ListNodesResponseBodyDataNodes {
	s.Connection = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetCreateTime(v int64) *ListNodesResponseBodyDataNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetCronExpress(v string) *ListNodesResponseBodyDataNodes {
	s.CronExpress = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetDeployDate(v int64) *ListNodesResponseBodyDataNodes {
	s.DeployDate = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetDescription(v string) *ListNodesResponseBodyDataNodes {
	s.Description = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetDqcDescription(v string) *ListNodesResponseBodyDataNodes {
	s.DqcDescription = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetDqcType(v int32) *ListNodesResponseBodyDataNodes {
	s.DqcType = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetFileId(v int64) *ListNodesResponseBodyDataNodes {
	s.FileId = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetFileType(v int32) *ListNodesResponseBodyDataNodes {
	s.FileType = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetFileVersion(v int32) *ListNodesResponseBodyDataNodes {
	s.FileVersion = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetModifyTime(v int64) *ListNodesResponseBodyDataNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetNodeId(v int64) *ListNodesResponseBodyDataNodes {
	s.NodeId = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetNodeName(v string) *ListNodesResponseBodyDataNodes {
	s.NodeName = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetOwnerId(v string) *ListNodesResponseBodyDataNodes {
	s.OwnerId = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetParamValues(v string) *ListNodesResponseBodyDataNodes {
	s.ParamValues = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetPriority(v int32) *ListNodesResponseBodyDataNodes {
	s.Priority = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetProgramType(v string) *ListNodesResponseBodyDataNodes {
	s.ProgramType = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetProjectId(v int64) *ListNodesResponseBodyDataNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetRelatedFlowId(v int64) *ListNodesResponseBodyDataNodes {
	s.RelatedFlowId = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetRepeatInterval(v int64) *ListNodesResponseBodyDataNodes {
	s.RepeatInterval = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetRepeatMode(v int32) *ListNodesResponseBodyDataNodes {
	s.RepeatMode = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetRepeatability(v bool) *ListNodesResponseBodyDataNodes {
	s.Repeatability = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetResGroupIdentifier(v string) *ListNodesResponseBodyDataNodes {
	s.ResGroupIdentifier = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetResGroupName(v string) *ListNodesResponseBodyDataNodes {
	s.ResGroupName = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) SetSchedulerType(v string) *ListNodesResponseBodyDataNodes {
	s.SchedulerType = &v
	return s
}

func (s *ListNodesResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}
