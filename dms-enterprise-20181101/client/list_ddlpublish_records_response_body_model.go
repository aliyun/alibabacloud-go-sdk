// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDDLPublishRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDDLPublishRecordList(v []*ListDDLPublishRecordsResponseBodyDDLPublishRecordList) *ListDDLPublishRecordsResponseBody
	GetDDLPublishRecordList() []*ListDDLPublishRecordsResponseBodyDDLPublishRecordList
	SetErrorCode(v string) *ListDDLPublishRecordsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDDLPublishRecordsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDDLPublishRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDDLPublishRecordsResponseBody
	GetSuccess() *bool
}

type ListDDLPublishRecordsResponseBody struct {
	// The details of the publishing records.
	DDLPublishRecordList []*ListDDLPublishRecordsResponseBodyDDLPublishRecordList `json:"DDLPublishRecordList,omitempty" xml:"DDLPublishRecordList,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A1549FB0-D4B8-4140-919F-17322C1072B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDDLPublishRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDDLPublishRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponseBody) GetDDLPublishRecordList() []*ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	return s.DDLPublishRecordList
}

func (s *ListDDLPublishRecordsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDDLPublishRecordsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDDLPublishRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDDLPublishRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDDLPublishRecordsResponseBody) SetDDLPublishRecordList(v []*ListDDLPublishRecordsResponseBodyDDLPublishRecordList) *ListDDLPublishRecordsResponseBody {
	s.DDLPublishRecordList = v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) SetErrorCode(v string) *ListDDLPublishRecordsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) SetErrorMessage(v string) *ListDDLPublishRecordsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) SetRequestId(v string) *ListDDLPublishRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) SetSuccess(v bool) *ListDDLPublishRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDDLPublishRecordsResponseBodyDDLPublishRecordList struct {
	// The time when the approval expires.
	//
	// example:
	//
	// 2020-12-14 20:52:38
	AuditExpireTime *string `json:"AuditExpireTime,omitempty" xml:"AuditExpireTime,omitempty"`
	// The approval state of the ticket. Valid values:
	//
	// 	- **EXEMPT_PASS**: The ticket passes without approval.
	//
	// 	- **TO_AUDIT**: The ticket is pending for approval.
	//
	// 	- **CANCEL**: The ticket is canceled.
	//
	// 	- **SUCCESS**: The ticket is approved.
	//
	// 	- **FAIL**: The ticket fails to pass the approval.
	//
	// example:
	//
	// CANCEL
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Release remarks.
	//
	// example:
	//
	// Release remarks
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the user who creates the ticket. You can obtain the user ID by calling the [GetUser](https://help.aliyun.com/document_detail/147098.html) operation and querying the value of the UserId parameter. The value is not the unique ID (UID) of the Alibaba Cloud account.
	//
	// example:
	//
	// 1423
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// Indicates whether the approval is terminated. Valid values:
	//
	// 	- **true**: The approval is terminated.
	//
	// 	- **false**: The approval is not terminated.
	//
	// > Multiple reasons can terminate the approval. For example, you withdraw the application or your ticket is not approved before the specified time.
	//
	// example:
	//
	// true
	Finality *bool `json:"Finality,omitempty" xml:"Finality,omitempty"`
	// The reason for the termination.
	//
	// example:
	//
	// CANCEL
	FinalityReason *string `json:"FinalityReason,omitempty" xml:"FinalityReason,omitempty"`
	// The publishing state of the ticket. Valid values:
	//
	// 	- **START**: The ticket is created.
	//
	// 	- **ANALYZE**: The ticket is under analysis.
	//
	// 	- **AUDIT**: The ticket is under approval.
	//
	// 	- **DISPATCH**: A task is generated for the ticket.
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// example:
	//
	// AUDIT
	PublishStatus *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The list of publishing tasks.
	PublishTaskInfoList []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList `json:"PublishTaskInfoList,omitempty" xml:"PublishTaskInfoList,omitempty" type:"Repeated"`
	// The risk level of the operation. Valid values:
	//
	// 	- **NONE_RISK**: The operation does not have risks.
	//
	// 	- **LOW_RISK**: The operation is at low risk.
	//
	// 	- **MIDDLE_RISK**: The operation is at medium risk.
	//
	// 	- **HIGH_RISK**: The operation is at high risk.
	//
	// example:
	//
	// LOW_RISK
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The description of the publishing state.
	//
	// example:
	//
	// CANCEL
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The ID of the approval process.
	//
	// example:
	//
	// 432153
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordList) String() string {
	return dara.Prettify(s)
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetAuditExpireTime() *string {
	return s.AuditExpireTime
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetComment() *string {
	return s.Comment
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetFinality() *bool {
	return s.Finality
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetFinalityReason() *string {
	return s.FinalityReason
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetPublishTaskInfoList() []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	return s.PublishTaskInfoList
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetAuditExpireTime(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.AuditExpireTime = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetAuditStatus(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.AuditStatus = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetComment(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.Comment = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetCreatorId(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.CreatorId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetFinality(v bool) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.Finality = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetFinalityReason(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.FinalityReason = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetPublishStatus(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.PublishStatus = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetPublishTaskInfoList(v []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.PublishTaskInfoList = v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetRiskLevel(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.RiskLevel = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetStatusDesc(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.StatusDesc = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) SetWorkflowInstanceId(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordList {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordList) Validate() error {
	return dara.Validate(s)
}

type ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList struct {
	// The ID of the database.
	//
	// example:
	//
	// 4325
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: the database is not a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The time to publish the ticket.
	//
	// example:
	//
	// 2020-12-14 20:52:38
	PlanTime *string `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
	// The list of the publishing tasks.
	PublishJobList []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList `json:"PublishJobList,omitempty" xml:"PublishJobList,omitempty" type:"Repeated"`
	// The publishing policy. Valid values:
	//
	// 	- **IMMEDIATELY**: immediately publishes the ticket.
	//
	// 	- **REGULARLY**: publishes the ticket at a scheduled time.
	//
	// example:
	//
	// IMMEDIATELY
	PublishStrategy *string `json:"PublishStrategy,omitempty" xml:"PublishStrategy,omitempty"`
	// The description of the state.
	//
	// example:
	//
	// NONE
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The state of the task.
	//
	// example:
	//
	// NONE
	TaskJobStatus *string `json:"TaskJobStatus,omitempty" xml:"TaskJobStatus,omitempty"`
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GetLogic() *bool {
	return s.Logic
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GetPlanTime() *string {
	return s.PlanTime
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GetPublishJobList() []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	return s.PublishJobList
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GetPublishStrategy() *string {
	return s.PublishStrategy
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) GetTaskJobStatus() *string {
	return s.TaskJobStatus
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetDbId(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.DbId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetLogic(v bool) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.Logic = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetPlanTime(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.PlanTime = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetPublishJobList(v []*ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.PublishJobList = v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetPublishStrategy(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.PublishStrategy = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetStatusDesc(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.StatusDesc = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) SetTaskJobStatus(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList {
	s.TaskJobStatus = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoList) Validate() error {
	return dara.Validate(s)
}

type ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList struct {
	// The ID of the SQL task group.
	//
	// example:
	//
	// 423515
	DBTaskGroupId *int64 `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
	// The number of SQL statements that are executed.
	//
	// example:
	//
	// 0
	ExecuteCount *int64 `json:"ExecuteCount,omitempty" xml:"ExecuteCount,omitempty"`
	// The script for data changes.
	//
	// example:
	//
	// ALTER TABLE test_toolkit_rename_table_after_rename MODIFY COLUMN gmt_modified datetime NOT NULL
	Scripts *string `json:"Scripts,omitempty" xml:"Scripts,omitempty"`
	// The description of the state.
	//
	// example:
	//
	// NONE
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The name of the table after the change.
	//
	// example:
	//
	// test_toolkit_rename_table_after_rename
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The state of the publishing task. Valid values:
	//
	// 	- **NONE**: The state of the task is unknown.
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **FAIL**: The task fails.
	//
	// example:
	//
	// NONE
	TaskJobStatus *string `json:"TaskJobStatus,omitempty" xml:"TaskJobStatus,omitempty"`
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) String() string {
	return dara.Prettify(s)
}

func (s ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) GoString() string {
	return s.String()
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) GetDBTaskGroupId() *int64 {
	return s.DBTaskGroupId
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) GetExecuteCount() *int64 {
	return s.ExecuteCount
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) GetScripts() *string {
	return s.Scripts
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) GetTableName() *string {
	return s.TableName
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) GetTaskJobStatus() *string {
	return s.TaskJobStatus
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetDBTaskGroupId(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.DBTaskGroupId = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetExecuteCount(v int64) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.ExecuteCount = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetScripts(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.Scripts = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetStatusDesc(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.StatusDesc = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetTableName(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.TableName = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) SetTaskJobStatus(v string) *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList {
	s.TaskJobStatus = &v
	return s
}

func (s *ListDDLPublishRecordsResponseBodyDDLPublishRecordListPublishTaskInfoListPublishJobList) Validate() error {
	return dara.Validate(s)
}
