// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateResourcesByAdvancedSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateResourcesByAdvancedSearchRequest
	GetAggregatorId() *string
	SetSql(v string) *ListAggregateResourcesByAdvancedSearchRequest
	GetSql() *string
}

type ListAggregateResourcesByAdvancedSearchRequest struct {
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-4b05626622af000c****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The SQL query statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT ResourceId, ResourceName WHERE Tags.Kvpair=\\"business:online\\"
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s ListAggregateResourcesByAdvancedSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourcesByAdvancedSearchRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateResourcesByAdvancedSearchRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateResourcesByAdvancedSearchRequest) GetSql() *string {
	return s.Sql
}

func (s *ListAggregateResourcesByAdvancedSearchRequest) SetAggregatorId(v string) *ListAggregateResourcesByAdvancedSearchRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchRequest) SetSql(v string) *ListAggregateResourcesByAdvancedSearchRequest {
	s.Sql = &v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchRequest) Validate() error {
	return dara.Validate(s)
}
