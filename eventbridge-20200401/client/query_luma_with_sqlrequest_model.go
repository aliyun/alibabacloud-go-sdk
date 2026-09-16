// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLumaWithSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *QueryLumaWithSQLRequest
	GetAgentName() *string
	SetMaxRows(v int32) *QueryLumaWithSQLRequest
	GetMaxRows() *int32
	SetSql(v string) *QueryLumaWithSQLRequest
	GetSql() *string
}

type QueryLumaWithSQLRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// Rows exceeding this limit are truncated. The IsTruncated field in the response indicates whether truncation occurred.
	//
	// example:
	//
	// 100
	MaxRows *int32 `json:"MaxRows,omitempty" xml:"MaxRows,omitempty"`
	// Only query statements are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM my_table LIMIT 10
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s QueryLumaWithSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLumaWithSQLRequest) GoString() string {
	return s.String()
}

func (s *QueryLumaWithSQLRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryLumaWithSQLRequest) GetMaxRows() *int32 {
	return s.MaxRows
}

func (s *QueryLumaWithSQLRequest) GetSql() *string {
	return s.Sql
}

func (s *QueryLumaWithSQLRequest) SetAgentName(v string) *QueryLumaWithSQLRequest {
	s.AgentName = &v
	return s
}

func (s *QueryLumaWithSQLRequest) SetMaxRows(v int32) *QueryLumaWithSQLRequest {
	s.MaxRows = &v
	return s
}

func (s *QueryLumaWithSQLRequest) SetSql(v string) *QueryLumaWithSQLRequest {
	s.Sql = &v
	return s
}

func (s *QueryLumaWithSQLRequest) Validate() error {
	return dara.Validate(s)
}
