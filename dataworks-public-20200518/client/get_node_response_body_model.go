// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetNodeResponseBodyData) *GetNodeResponseBody
	GetData() *GetNodeResponseBodyData
	SetErrorCode(v string) *GetNodeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetNodeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetNodeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNodeResponseBody
	GetSuccess() *bool
}

type GetNodeResponseBody struct {
	// The details of the node.
	Data *GetNodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the node. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the node ID.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The connection string.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The operation that you want to perform. Set the value to **GetNode**.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Other parameters.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBody) GetData() *GetNodeResponseBodyData {
	return s.Data
}

func (s *GetNodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNodeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNodeResponseBody) SetData(v *GetNodeResponseBodyData) *GetNodeResponseBody {
	s.Data = v
	return s
}

func (s *GetNodeResponseBody) SetErrorCode(v string) *GetNodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeResponseBody) SetErrorMessage(v string) *GetNodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetNodeResponseBody) SetHttpStatusCode(v int32) *GetNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeResponseBody) SetRequestId(v string) *GetNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeResponseBody) SetSuccess(v bool) *GetNodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNodeResponseBodyData struct {
	// The description of the node.
	//
	// example:
	//
	// 123456
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The operation that you want to perform. Set the value to **GetNode**.
	//
	// example:
	//
	// 123
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV.
	//
	// example:
	//
	// odps_first_dev
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1727280000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The deployment date.
	//
	// example:
	//
	// 1727280000000
	DeployDate *int64 `json:"DeployDate,omitempty" xml:"DeployDate,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the node. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the node ID.
	//
	// example:
	//
	// [{"projectName":"test_0923001","tableName":"test_table_001","partition":"ds\\u003d$[yyyymmdd]"},{"projectName":"test_0923001","tableName":"test_table_002","partition":"NOTAPARTITIONTABLE"}]
	DqcDescription *string `json:"DqcDescription,omitempty" xml:"DqcDescription,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// 1
	DqcType *int32 `json:"DqcType,omitempty" xml:"DqcType,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file type. Different file types have different codes. For more information, see [DataWorks node collection](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The version of the file.
	//
	// example:
	//
	// 1
	FileVersion *int32 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1727280000000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The scheduling type of the node. Valid values:
	//
	// 	- NORMAL: The node is an auto triggered node.
	//
	// 	- MANUAL: The node is a manually triggered node. Manually triggered nodes cannot be automatically triggered.
	//
	// 	- PAUSE: The node is a paused node.
	//
	// 	- SKIP: The node is a dry-run node. Dry-run nodes are started as scheduled but the system sets the status of the nodes to successful when it starts to run them.
	//
	// example:
	//
	// The ID of the baseline.
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 17366294****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The CRON expression returned.
	//
	// example:
	//
	// a=b
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the owner of the node.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// Indicates whether the node can be rerun.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Indicates whether the node is associated with Data Quality. Valid values: 0 and 1. A value of 0 indicates that the node is associated with Data Quality. A value of 1 indicates that the node is not associated with Data Quality.
	//
	// example:
	//
	// 123
	RelatedFlowId *int64 `json:"RelatedFlowId,omitempty" xml:"RelatedFlowId,omitempty"`
	// The ID of the workflow to which the node belongs.
	//
	// example:
	//
	// 60
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// Rerun mode. 0 indicates that you can rerun only if you fail, 1 indicates that you can rerun in all cases, and 2 indicates that you cannot rerun in all cases.
	//
	// example:
	//
	// 1
	RepeatMode *int32 `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// true
	Repeatability *string `json:"Repeatability,omitempty" xml:"Repeatability,omitempty"`
	// The unique identifier of the resource group.
	//
	// example:
	//
	// group_123
	ResGroupIdentifier *string `json:"ResGroupIdentifier,omitempty" xml:"ResGroupIdentifier,omitempty"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// The table and partition filter expression in Data Quality that are associated with the node.
	ResGroupName *string `json:"ResGroupName,omitempty" xml:"ResGroupName,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s GetNodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyData) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetNodeResponseBodyData) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetNodeResponseBodyData) GetConnection() *string {
	return s.Connection
}

func (s *GetNodeResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNodeResponseBodyData) GetCronExpress() *string {
	return s.CronExpress
}

func (s *GetNodeResponseBodyData) GetDeployDate() *int64 {
	return s.DeployDate
}

func (s *GetNodeResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetNodeResponseBodyData) GetDqcDescription() *string {
	return s.DqcDescription
}

func (s *GetNodeResponseBodyData) GetDqcType() *int32 {
	return s.DqcType
}

func (s *GetNodeResponseBodyData) GetFileId() *int64 {
	return s.FileId
}

func (s *GetNodeResponseBodyData) GetFileType() *int32 {
	return s.FileType
}

func (s *GetNodeResponseBodyData) GetFileVersion() *int32 {
	return s.FileVersion
}

func (s *GetNodeResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetNodeResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetNodeResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *GetNodeResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetNodeResponseBodyData) GetParamValues() *string {
	return s.ParamValues
}

func (s *GetNodeResponseBodyData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetNodeResponseBodyData) GetProgramType() *string {
	return s.ProgramType
}

func (s *GetNodeResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeResponseBodyData) GetRelatedFlowId() *int64 {
	return s.RelatedFlowId
}

func (s *GetNodeResponseBodyData) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *GetNodeResponseBodyData) GetRepeatMode() *int32 {
	return s.RepeatMode
}

func (s *GetNodeResponseBodyData) GetRepeatability() *string {
	return s.Repeatability
}

func (s *GetNodeResponseBodyData) GetResGroupIdentifier() *string {
	return s.ResGroupIdentifier
}

func (s *GetNodeResponseBodyData) GetResGroupName() *string {
	return s.ResGroupName
}

func (s *GetNodeResponseBodyData) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *GetNodeResponseBodyData) SetBaselineId(v int64) *GetNodeResponseBodyData {
	s.BaselineId = &v
	return s
}

func (s *GetNodeResponseBodyData) SetBusinessId(v int64) *GetNodeResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *GetNodeResponseBodyData) SetConnection(v string) *GetNodeResponseBodyData {
	s.Connection = &v
	return s
}

func (s *GetNodeResponseBodyData) SetCreateTime(v int64) *GetNodeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetNodeResponseBodyData) SetCronExpress(v string) *GetNodeResponseBodyData {
	s.CronExpress = &v
	return s
}

func (s *GetNodeResponseBodyData) SetDeployDate(v int64) *GetNodeResponseBodyData {
	s.DeployDate = &v
	return s
}

func (s *GetNodeResponseBodyData) SetDescription(v string) *GetNodeResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetNodeResponseBodyData) SetDqcDescription(v string) *GetNodeResponseBodyData {
	s.DqcDescription = &v
	return s
}

func (s *GetNodeResponseBodyData) SetDqcType(v int32) *GetNodeResponseBodyData {
	s.DqcType = &v
	return s
}

func (s *GetNodeResponseBodyData) SetFileId(v int64) *GetNodeResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetNodeResponseBodyData) SetFileType(v int32) *GetNodeResponseBodyData {
	s.FileType = &v
	return s
}

func (s *GetNodeResponseBodyData) SetFileVersion(v int32) *GetNodeResponseBodyData {
	s.FileVersion = &v
	return s
}

func (s *GetNodeResponseBodyData) SetModifyTime(v int64) *GetNodeResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetNodeResponseBodyData) SetNodeId(v int64) *GetNodeResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetNodeResponseBodyData) SetNodeName(v string) *GetNodeResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *GetNodeResponseBodyData) SetOwnerId(v string) *GetNodeResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetNodeResponseBodyData) SetParamValues(v string) *GetNodeResponseBodyData {
	s.ParamValues = &v
	return s
}

func (s *GetNodeResponseBodyData) SetPriority(v int32) *GetNodeResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetNodeResponseBodyData) SetProgramType(v string) *GetNodeResponseBodyData {
	s.ProgramType = &v
	return s
}

func (s *GetNodeResponseBodyData) SetProjectId(v int64) *GetNodeResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetNodeResponseBodyData) SetRelatedFlowId(v int64) *GetNodeResponseBodyData {
	s.RelatedFlowId = &v
	return s
}

func (s *GetNodeResponseBodyData) SetRepeatInterval(v int64) *GetNodeResponseBodyData {
	s.RepeatInterval = &v
	return s
}

func (s *GetNodeResponseBodyData) SetRepeatMode(v int32) *GetNodeResponseBodyData {
	s.RepeatMode = &v
	return s
}

func (s *GetNodeResponseBodyData) SetRepeatability(v string) *GetNodeResponseBodyData {
	s.Repeatability = &v
	return s
}

func (s *GetNodeResponseBodyData) SetResGroupIdentifier(v string) *GetNodeResponseBodyData {
	s.ResGroupIdentifier = &v
	return s
}

func (s *GetNodeResponseBodyData) SetResGroupName(v string) *GetNodeResponseBodyData {
	s.ResGroupName = &v
	return s
}

func (s *GetNodeResponseBodyData) SetSchedulerType(v string) *GetNodeResponseBodyData {
	s.SchedulerType = &v
	return s
}

func (s *GetNodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
