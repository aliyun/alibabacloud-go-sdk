// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *ListAgenticDataSourcesRequest
	GetDbType() *string
	SetHost(v string) *ListAgenticDataSourcesRequest
	GetHost() *string
	SetInstanceSource(v string) *ListAgenticDataSourcesRequest
	GetInstanceSource() *string
	SetMaxResults(v int32) *ListAgenticDataSourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgenticDataSourcesRequest
	GetNextToken() *string
	SetPort(v int32) *ListAgenticDataSourcesRequest
	GetPort() *int32
	SetSchema(v string) *ListAgenticDataSourcesRequest
	GetSchema() *string
	SetSearchKey(v string) *ListAgenticDataSourcesRequest
	GetSearchKey() *string
	SetState(v int32) *ListAgenticDataSourcesRequest
	GetState() *int32
}

type ListAgenticDataSourcesRequest struct {
	// Filters data sources by type. If you leave this parameter empty, no restriction is applied. Valid values: MYSQL, POSTGRESQL, ORACLE, SQLSERVER, MARIADB, POLARDB, POLARDB_POSTGRESQL, POLARDB_ORACLE, OCEANBASEX, DRDS, ADB_MYSQL, ADB_POSTGRESQL, CLICKHOUSE, SELECTDB, STARROCKS, HOLOGRES, MAXCOMPUTE, ES_REST, and TRINO.
	//
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// Filters data sources by endpoint. If you leave this parameter empty, no restriction is applied.
	//
	// example:
	//
	// mysql.example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Filters data sources by source. If you leave this parameter empty, no restriction is applied. Valid values:
	//
	// - RDS: ApsaraDB instance
	//
	// - PUBLIC_OWN: self-managed database on the public network
	//
	// - ECS_OWN: self-managed database on an ECS instance
	//
	// - GATEWAY: connected through Database Gateway
	//
	// - VPC_IDC: connected over a VPC or IDC
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The number of entries to return per page. Default value: 20. Maximum value: 200.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination cursor. Leave this parameter empty when you query the first page. When you query subsequent pages, pass the NextToken value returned from the previous request.
	//
	// example:
	//
	// page-2-token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Filters data sources by connection port. If you leave this parameter empty, no restriction is applied.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Filters data sources by connection target name. The meaning of this parameter varies based on the data source type. If you leave this parameter empty, no restriction is applied.
	//
	// example:
	//
	// postgres
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Performs a fuzzy search by data source name or endpoint. If you leave this parameter empty, no fuzzy search is performed.
	//
	// example:
	//
	// demo
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// Filters data sources by status. If you leave this parameter empty, no restriction is applied. Valid values:
	//
	// - 0: normal
	//
	// - -1: deleted
	//
	// - -2: disabled
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAgenticDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListAgenticDataSourcesRequest) GetDbType() *string {
	return s.DbType
}

func (s *ListAgenticDataSourcesRequest) GetHost() *string {
	return s.Host
}

func (s *ListAgenticDataSourcesRequest) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *ListAgenticDataSourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgenticDataSourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgenticDataSourcesRequest) GetPort() *int32 {
	return s.Port
}

func (s *ListAgenticDataSourcesRequest) GetSchema() *string {
	return s.Schema
}

func (s *ListAgenticDataSourcesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListAgenticDataSourcesRequest) GetState() *int32 {
	return s.State
}

func (s *ListAgenticDataSourcesRequest) SetDbType(v string) *ListAgenticDataSourcesRequest {
	s.DbType = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) SetHost(v string) *ListAgenticDataSourcesRequest {
	s.Host = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) SetInstanceSource(v string) *ListAgenticDataSourcesRequest {
	s.InstanceSource = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) SetMaxResults(v int32) *ListAgenticDataSourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) SetNextToken(v string) *ListAgenticDataSourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) SetPort(v int32) *ListAgenticDataSourcesRequest {
	s.Port = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) SetSchema(v string) *ListAgenticDataSourcesRequest {
	s.Schema = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) SetSearchKey(v string) *ListAgenticDataSourcesRequest {
	s.SearchKey = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) SetState(v int32) *ListAgenticDataSourcesRequest {
	s.State = &v
	return s
}

func (s *ListAgenticDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
