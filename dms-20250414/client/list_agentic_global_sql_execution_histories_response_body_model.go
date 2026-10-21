// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticGlobalSqlExecutionHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) *ListAgenticGlobalSqlExecutionHistoriesResponseBody
	GetData() []*ListAgenticGlobalSqlExecutionHistoriesResponseBodyData
	SetErrorCode(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgenticGlobalSqlExecutionHistoriesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgenticGlobalSqlExecutionHistoriesResponseBody
	GetTotalCount() *int64
}

type ListAgenticGlobalSqlExecutionHistoriesResponseBody struct {
	// The list of SQL execution history.
	Data []*ListAgenticGlobalSqlExecutionHistoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that meet the conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgenticGlobalSqlExecutionHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticGlobalSqlExecutionHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) GetData() []*ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	return s.Data
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) SetData(v []*ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) *ListAgenticGlobalSqlExecutionHistoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) SetErrorCode(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) SetErrorMessage(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) SetRequestId(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) SetSuccess(v bool) *ListAgenticGlobalSqlExecutionHistoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) SetTotalCount(v int64) *ListAgenticGlobalSqlExecutionHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBody) Validate() error {
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

type ListAgenticGlobalSqlExecutionHistoriesResponseBodyData struct {
	// The ID of the actor. A user ID starts with usr_, and an agent ID starts with agt_.
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
	// The type of the actor. A value of HUMAN indicates a user, and a value of AGENT indicates an agent.
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
	// The IP address of the client that initiates the execution, excluding the port. This parameter is empty in workflow scenarios or when the IP address is not recorded.
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
	// The actual SQL statement that is executed.
	//
	// example:
	//
	// SELECT 1
	ExecutedSql *string `json:"ExecutedSql,omitempty" xml:"ExecutedSql,omitempty"`
	// The execution status. Valid values:
	//
	// - SUCCESS: The execution is successful.
	//
	// - FAILED: The execution fails.
	//
	// - BLOCKED: The execution is blocked or not actually executed.
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

func (s ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetActorId() *string {
	return s.ActorId
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetActorName() *string {
	return s.ActorName
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetActorType() *string {
	return s.ActorType
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetAffectedRows() *int64 {
	return s.AffectedRows
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetAgentOwnerId() *string {
	return s.AgentOwnerId
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetCost() *int64 {
	return s.Cost
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetDatabaseSearchName() *string {
	return s.DatabaseSearchName
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetError() *string {
	return s.Error
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetExecutedSql() *string {
	return s.ExecutedSql
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetOriginalSql() *string {
	return s.OriginalSql
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetQueryId() *string {
	return s.QueryId
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetReturnedRows() *int64 {
	return s.ReturnedRows
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetScenario() *string {
	return s.Scenario
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetSqlType() *string {
	return s.SqlType
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetActorId(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.ActorId = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetActorName(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.ActorName = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetActorType(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.ActorType = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetAffectedRows(v int64) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.AffectedRows = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetAgentOwnerId(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.AgentOwnerId = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetCatalogUuid(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.CatalogUuid = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetClientIp(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.ClientIp = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetCost(v int64) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.Cost = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetDatabaseSearchName(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.DatabaseSearchName = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetDatabaseUuid(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.DatabaseUuid = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetError(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.Error = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetExecutedSql(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.ExecutedSql = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetExecutionStatus(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.ExecutionStatus = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetOriginalSql(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.OriginalSql = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetQueryId(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.QueryId = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetReturnedRows(v int64) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.ReturnedRows = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetScenario(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.Scenario = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetSqlType(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.SqlType = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetStartTime(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) SetWorkflowId(v string) *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
