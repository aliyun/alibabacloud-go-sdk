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
	// The key word that you specify to query connections. Connections can be queried by prefixes.
	//
	// example:
	//
	// connection-name
	ConnectionNamePrefix *string `json:"ConnectionNamePrefix,omitempty" xml:"ConnectionNamePrefix,omitempty"`
	// The maximum number of entries to be returned in a single call. You can use this parameter and the NextToken parameter to implement paging.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If you set the Limit parameter and excess return values exist, this parameter is returned.
	//
	// 	- Default value: 0.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 按连接类型过滤查询结果。可选值：Http、MySQL、PostgreSQL、Elasticsearch。不传则返回所有类型
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
