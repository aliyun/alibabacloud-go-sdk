// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionNamePrefix(v string) *ListConnectionsRequest
	GetConnectionNamePrefix() *string
	SetMaxResults(v int64) *ListConnectionsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListConnectionsRequest
	GetNextToken() *string
	SetType(v string) *ListConnectionsRequest
	GetType() *string
}

type ListConnectionsRequest struct {
	// The prefix of the connection configuration name. Supports fuzzy match by prefix.
	//
	// example:
	//
	// connection-name
	ConnectionNamePrefix *string `json:"ConnectionNamePrefix,omitempty" xml:"ConnectionNamePrefix,omitempty"`
	// The maximum number of entries to return per request. You can use this parameter together with NextToken to implement paging.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. If the number of results exceeds the value of MaxResults, a NextToken value is returned.
	//
	// - The NextToken value starts from 0. Default value: 0.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Filters query results by connection type. Valid values: Http, MySQL, PostgreSQL, Elasticsearch, OSS_TABLES, SLS, OTS, MaxCompute, MongoDB, Redis, SQLServer, ClickHouse, Oracle, Hive, Iceberg, and lakehouse. If this parameter is not specified, all types are returned.
	//
	// example:
	//
	// Http
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsRequest) GetConnectionNamePrefix() *string {
	return s.ConnectionNamePrefix
}

func (s *ListConnectionsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListConnectionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectionsRequest) GetType() *string {
	return s.Type
}

func (s *ListConnectionsRequest) SetConnectionNamePrefix(v string) *ListConnectionsRequest {
	s.ConnectionNamePrefix = &v
	return s
}

func (s *ListConnectionsRequest) SetMaxResults(v int64) *ListConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectionsRequest) SetNextToken(v string) *ListConnectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectionsRequest) SetType(v string) *ListConnectionsRequest {
	s.Type = &v
	return s
}

func (s *ListConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
