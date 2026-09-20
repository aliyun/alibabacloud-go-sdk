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
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNodeResponseBodyData struct {
	// The baseline ID. The baseline ID configured for the node as a leaf node is returned. If no baseline is configured, a workspace default value is returned.
	//
	// example:
	//
	// 123456
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 123
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The connection string.
	//
	// example:
	//
	// odps_source_dev
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The creation time.
	//
	// The value is a 13-digit number, such as `1727280000000`.
	//
	// example:
	//
	// 1727280000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The CRON expression.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// The deployment date.
	//
	// The value is a 13-digit number, such as `1727280000000`.
	//
	// example:
	//
	// 1727280000000
	DeployDate *int64 `json:"DeployDate,omitempty" xml:"DeployDate,omitempty"`
	// The description of the node.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The DQC partitioning rule string.
	//
	// example:
	//
	// [{"projectName":"test_0923001","tableName":"test_table_001","partition":"ds\\u003d$[yyyymmdd]"},{"projectName":"test_0923001","tableName":"test_table_002","partition":"NOTAPARTITIONTABLE"}]
	DqcDescription *string `json:"DqcDescription,omitempty" xml:"DqcDescription,omitempty"`
	// The DQC type. A value of 0 indicates that no DQC rule is associated. A value of 1 indicates that a DQC rule is associated.
	//
	// example:
	//
	// 1
	DqcType *int32 `json:"DqcType,omitempty" xml:"DqcType,omitempty"`
	// The file ID. <warning>This field is deprecated.</warning>
	//
	// example:
	//
	// 123
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file type. Different file types have different codes. For more information, see [DataWorks nodes](https://help.aliyun.com/document_detail/600169.html).
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The file version.
	//
	// example:
	//
	// 1
	FileVersion *int32 `json:"FileVersion,omitempty" xml:"FileVersion,omitempty"`
	// The modification time.
	//
	// The value is a 13-digit number, such as `1727280000000`.
	//
	// example:
	//
	// 1727280000000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// sql_node
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the node owner.
	//
	// example:
	//
	// 17366294****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The additional parameters.
	//
	// example:
	//
	// a=b
	ParamValues *string `json:"ParamValues,omitempty" xml:"ParamValues,omitempty"`
	// The priority of the node. Valid values: 1, 3, 5, 7, and 8.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the associated workflow.
	//
	// example:
	//
	// 123
	RelatedFlowId *int64 `json:"RelatedFlowId,omitempty" xml:"RelatedFlowId,omitempty"`
	// The interval at which the node is rescheduled after a failure.
	//
	// example:
	//
	// 60
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// The rerun mode. A value of 0 indicates that the node can be rerun only upon failure. A value of 1 indicates that the node can be rerun in all cases. A value of 2 indicates that the node cannot be rerun in any case.
	//
	// example:
	//
	// 1
	RepeatMode *int32 `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// Indicates whether the node can be rerun.
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
	// The name of the resource group.
	//
	// example:
	//
	// Default Resource Group
	ResGroupName *string `json:"ResGroupName,omitempty" xml:"ResGroupName,omitempty"`
	// The scheduling type. Valid values:
	//
	// - NORMAL: normal scheduling node.
	//
	// - MANUAL: manual node that is not scheduled on a regular basis.
	//
	// - PAUSE: paused node.
	//
	// - SKIP: dry-run node that is scheduled on a regular basis but is directly set to successful when scheduling starts.
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
