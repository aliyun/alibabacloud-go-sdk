// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticSqlExecutionHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActorId(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetActorId() *string
	SetActorName(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetActorName() *string
	SetActorType(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetActorType() *string
	SetCatalogUuid(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetDatabaseUuid() *string
	SetEndTime(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetEndTime() *string
	SetPage(v int32) *ListAgenticSqlExecutionHistoriesRequest
	GetPage() *int32
	SetScenario(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetScenario() *string
	SetSize(v int32) *ListAgenticSqlExecutionHistoriesRequest
	GetSize() *int32
	SetSqlKeyword(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetSqlKeyword() *string
	SetSqlType(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetSqlType() *string
	SetStartTime(v string) *ListAgenticSqlExecutionHistoriesRequest
	GetStartTime() *string
}

type ListAgenticSqlExecutionHistoriesRequest struct {
	// The ID of the actor. User IDs start with usr_, and agent IDs start with agt_. If you do not specify this parameter, a user queries their own execution history and the execution history of their agents, while an agent queries only its own execution history. If you specify this parameter, exact matching is performed, and the value must be the ID of the current user or an agent owned by the current user.
	//
	// example:
	//
	// usr_10001
	ActorId *string `json:"ActorId,omitempty" xml:"ActorId,omitempty"`
	// The name of the actor for exact matching.
	//
	// example:
	//
	// demo-user
	ActorName *string `json:"ActorName,omitempty" xml:"ActorName,omitempty"`
	// The type of the actor for exact matching. Valid values: HUMAN and AGENT.
	//
	// example:
	//
	// HUMAN
	ActorType *string `json:"ActorType,omitempty" xml:"ActorType,omitempty"`
	// The catalog UUID for exact matching. You can obtain this value by calling ListAgenticCatalog or SearchAgenticCatalog.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The database UUID for exact matching. You can obtain this value by calling ListAgenticDatabase or SearchAgenticDatabase.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The end time of the query. Specify the time in the yyyy-MM-dd HH:mm:ss format. If you do not specify this parameter, the default value is the current time. If the format cannot be parsed or the end time is earlier than the start time, the request fails.
	//
	// example:
	//
	// 2026-06-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The execution scenario for exact matching. Valid values:
	//
	// - AGENT_QUERY: The agent directly executes the SQL statement.
	//
	// - WORKFLOW_QUERY: The SQL statement is asynchronously executed through a workflow.
	//
	// If you do not specify this parameter, both scenarios are queried.
	//
	// example:
	//
	// AGENT_QUERY
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The keyword for fuzzy matching against the original SQL statement.
	//
	// example:
	//
	// SELECT
	SqlKeyword *string `json:"SqlKeyword,omitempty" xml:"SqlKeyword,omitempty"`
	// The SQL type for exact matching.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the query. Specify the time in the yyyy-MM-dd HH:mm:ss format. If you do not specify this parameter, the default time range is the last 7 days. If the specified time is more than 730 days ago or the format cannot be parsed, the query is performed for the last 730 days.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAgenticSqlExecutionHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticSqlExecutionHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetActorId() *string {
	return s.ActorId
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetActorName() *string {
	return s.ActorName
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetActorType() *string {
	return s.ActorType
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetSqlKeyword() *string {
	return s.SqlKeyword
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *ListAgenticSqlExecutionHistoriesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetActorId(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.ActorId = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetActorName(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.ActorName = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetActorType(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.ActorType = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetCatalogUuid(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.CatalogUuid = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetDatabaseUuid(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetEndTime(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetPage(v int32) *ListAgenticSqlExecutionHistoriesRequest {
	s.Page = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetScenario(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.Scenario = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetSize(v int32) *ListAgenticSqlExecutionHistoriesRequest {
	s.Size = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetSqlKeyword(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.SqlKeyword = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetSqlType(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.SqlType = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) SetStartTime(v string) *ListAgenticSqlExecutionHistoriesRequest {
	s.StartTime = &v
	return s
}

func (s *ListAgenticSqlExecutionHistoriesRequest) Validate() error {
	return dara.Validate(s)
}
