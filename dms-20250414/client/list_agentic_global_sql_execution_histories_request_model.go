// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticGlobalSqlExecutionHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActorId(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetActorId() *string
	SetActorName(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetActorName() *string
	SetActorType(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetActorType() *string
	SetCatalogUuid(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetDatabaseUuid() *string
	SetEndTime(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetEndTime() *string
	SetPage(v int32) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetPage() *int32
	SetScenario(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetScenario() *string
	SetSize(v int32) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetSize() *int32
	SetSqlKeyword(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetSqlKeyword() *string
	SetSqlType(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetSqlType() *string
	SetStartTime(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest
	GetStartTime() *string
}

type ListAgenticGlobalSqlExecutionHistoriesRequest struct {
	// The ID of the actor. A user ID starts with usr_, and an agent ID starts with agt_. If you do not specify this parameter, all records within the current tenant are queried. If you specify a user ID, the records of the user and the agents under the user are included. If you specify an agent ID, only the records of the agent are matched. The actor must belong to the current tenant.
	//
	// example:
	//
	// usr_10001
	ActorId *string `json:"ActorId,omitempty" xml:"ActorId,omitempty"`
	// The name of the actor for exact match.
	//
	// example:
	//
	// demo-user
	ActorName *string `json:"ActorName,omitempty" xml:"ActorName,omitempty"`
	// The type of the actor for exact match. Valid values: HUMAN and AGENT.
	//
	// example:
	//
	// HUMAN
	ActorType *string `json:"ActorType,omitempty" xml:"ActorType,omitempty"`
	// The catalog UUID for exact match. You can obtain this value by calling ListAgenticCatalog or SearchAgenticCatalog.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The database UUID for exact match. You can obtain this value by calling ListAgenticDatabase or SearchAgenticDatabase.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The end time of the query. Specify the time in the yyyy-MM-dd HH:mm:ss format. If you do not specify this parameter, the default end time is the current time. If the format cannot be parsed or the end time is earlier than the start time, the request fails.
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
	// The execution scenario for exact match. Valid values:
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
	// The keyword for fuzzy match based on the original SQL statement.
	//
	// example:
	//
	// SELECT
	SqlKeyword *string `json:"SqlKeyword,omitempty" xml:"SqlKeyword,omitempty"`
	// The SQL type for exact match.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the query. Specify the time in the yyyy-MM-dd HH:mm:ss format. If you do not specify this parameter, the default time range is the last 7 days. If the specified time is earlier than 730 days ago or the format cannot be parsed, the query starts from 730 days ago.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAgenticGlobalSqlExecutionHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticGlobalSqlExecutionHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetActorId() *string {
	return s.ActorId
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetActorName() *string {
	return s.ActorName
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetActorType() *string {
	return s.ActorType
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetSqlKeyword() *string {
	return s.SqlKeyword
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetActorId(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.ActorId = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetActorName(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.ActorName = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetActorType(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.ActorType = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetCatalogUuid(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.CatalogUuid = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetDatabaseUuid(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetEndTime(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetPage(v int32) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.Page = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetScenario(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.Scenario = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetSize(v int32) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.Size = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetSqlKeyword(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.SqlKeyword = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetSqlType(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.SqlType = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) SetStartTime(v string) *ListAgenticGlobalSqlExecutionHistoriesRequest {
	s.StartTime = &v
	return s
}

func (s *ListAgenticGlobalSqlExecutionHistoriesRequest) Validate() error {
	return dara.Validate(s)
}
