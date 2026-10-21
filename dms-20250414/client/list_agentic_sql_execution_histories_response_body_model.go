// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticSqlExecutionHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAgenticSqlExecutionHistoriesResponseBodyData) *ListAgenticSqlExecutionHistoriesResponseBody
	GetData() []*ListAgenticSqlExecutionHistoriesResponseBodyData
	SetErrorCode(v string) *ListAgenticSqlExecutionHistoriesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAgenticSqlExecutionHistoriesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListAgenticSqlExecutionHistoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgenticSqlExecutionHistoriesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgenticSqlExecutionHistoriesResponseBody
	GetTotalCount() *int64
}

type ListAgenticSqlExecutionHistoriesResponseBody struct {
	// The list of SQL execution history.
	Data []*ListAgenticSqlExecutionHistoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. This parameter is empty when the request is successful.
	//
	// example:
	//
	// INVALID_PARAMETER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Invalid endTime format, expected yyyy-MM-dd HH:mm:ss
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34E01EDD-6A16-4CF0-9541-C644D1BE01AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgenticSqlExecutionHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticSqlExecutionHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) GetData() []*ListAgenticSqlExecutionHistoriesResponseBodyData {
	return s.Data
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) SetData(v []*ListAgenticSqlExecutionHistoriesResponseBodyData) *ListAgenticSqlExecutionHistoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) SetErrorCode(v string) *ListAgenticSqlExecutionHistoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) SetErrorMessage(v string) *ListAgenticSqlExecutionHistoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) SetRequestId(v string) *ListAgenticSqlExecutionHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) SetSuccess(v bool) *ListAgenticSqlExecutionHistoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) SetTotalCount(v int64) *ListAgenticSqlExecutionHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgenticSqlExecutionHistoriesResponseBodyData struct {
	// The ID of the actor. User IDs start with usr_, and agent IDs start with agt_.
	//
	// example:
	//
	// usr_10001
	ActorId *string `json:"ActorId,omitempty" xml:"ActorId,omitempty"`
	// The name of the actor.
	//
	// example:
	//
	// demo-user
	ActorName *string `json:"ActorName,omitempty" xml:"ActorName,omitempty"`
	// The type of the actor. HUMAN indicates a user, and AGENT indicates an agent.
	//
	// example:
	//
	// HUMAN
	ActorType *string `json:"ActorType,omitempty" xml:"ActorType,omitempty"`
	// The number of affected rows.
	//
	// example:
	//
	// 0
	AffectedRows *int64 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	// The ID of the user to which the agent belongs. The ID starts with usr_.
	//
	// example:
	//
	// usr_10001
	AgentOwnerId *string `json:"AgentOwnerId,omitempty" xml:"AgentOwnerId,omitempty"`
	// The catalog UUID to which the execution belongs.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The IP address of the client that initiated the execution, excluding the port. This parameter is empty in workflow scenarios or when the IP address is not recorded.
	//
	// example:
	//
	// 203.0.113.10
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The execution duration of the SQL statement.
	//
	// example:
	//
	// 12
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The search name of the database.
	//
	// example:
	//
	// demo_db
	DatabaseSearchName *string `json:"DatabaseSearchName,omitempty" xml:"DatabaseSearchName,omitempty"`
	// The database UUID to which the execution belongs.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The error message returned when the execution fails.
	//
	// example:
	//
	// syntax error
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The actual SQL statement that was executed.
	//
	// example:
	//
	// SELECT 1
	ExecutedSql *string `json:"ExecutedSql,omitempty" xml:"ExecutedSql,omitempty"`
	// The execution status. Valid values:
	//
	// - SUCCESS: The execution succeeded.
	//
	// - FAILED: The execution failed.
	//
	// - BLOCKED: The execution was blocked or not actually executed.
	//
	// example:
	//
	// SUCCESS
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	// The original SQL statement.
	//
	// example:
	//
	// SELECT 1
	OriginalSql *string `json:"OriginalSql,omitempty" xml:"OriginalSql,omitempty"`
	// The query ID of the execution.
	//
	// example:
	//
	// trace-10001
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The number of returned rows.
	//
	// example:
	//
	// 1
	ReturnedRows *int64 `json:"ReturnedRows,omitempty" xml:"ReturnedRows,omitempty"`
	// The execution scenario. Valid values:
	//
	// - AGENT_QUERY: The agent directly executes the SQL statement.
	//
	// - WORKFLOW_QUERY: The SQL statement is asynchronously executed through a workflow.
	//
	// example:
	//
	// AGENT_QUERY
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The SQL type.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the execution. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The workflow ID. This parameter may be empty in non-workflow scenarios.
	//
	// example:
	//
	// wf-10001
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListAgenticSqlExecutionHistoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticSqlExecutionHistoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetActorId() *string {
	return s.ActorId
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetActorName() *string {
	return s.ActorName
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetActorType() *string {
	return s.ActorType
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetAffectedRows() *int64 {
	return s.AffectedRows
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetAgentOwnerId() *string {
	return s.AgentOwnerId
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetCost() *int64 {
	return s.Cost
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetDatabaseSearchName() *string {
	return s.DatabaseSearchName
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetError() *string {
	return s.Error
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetExecutedSql() *string {
	return s.ExecutedSql
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetOriginalSql() *string {
	return s.OriginalSql
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetQueryId() *string {
	return s.QueryId
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetReturnedRows() *int64 {
	return s.ReturnedRows
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetScenario() *string {
	return s.Scenario
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetSqlType() *string {
	return s.SqlType
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetActorId(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.ActorId = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetActorName(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.ActorName = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetActorType(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.ActorType = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetAffectedRows(v int64) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.AffectedRows = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetAgentOwnerId(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.AgentOwnerId = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetCatalogUuid(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.CatalogUuid = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetClientIp(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.ClientIp = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetCost(v int64) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.Cost = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetDatabaseSearchName(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.DatabaseSearchName = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetDatabaseUuid(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.DatabaseUuid = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetError(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.Error = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetExecutedSql(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.ExecutedSql = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetExecutionStatus(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.ExecutionStatus = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetOriginalSql(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.OriginalSql = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetQueryId(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.QueryId = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetReturnedRows(v int64) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.ReturnedRows = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetScenario(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.Scenario = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetSqlType(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.SqlType = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetStartTime(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) SetWorkflowId(v string) *ListAgenticSqlExecutionHistoriesResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
