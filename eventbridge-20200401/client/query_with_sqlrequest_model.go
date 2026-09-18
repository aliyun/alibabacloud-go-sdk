// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWithSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *QueryWithSQLRequest
	GetLimit() *int32
	SetQuery(v string) *QueryWithSQLRequest
	GetQuery() *string
}

type QueryWithSQLRequest struct {
	// The maximum number of rows to return. Default value: 20. Maximum value: 50.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The query statement. Typically uses a three-part table name in the format of catalog.namespace.table.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM "test-es"."default"."product_info"
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s QueryWithSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWithSQLRequest) GoString() string {
	return s.String()
}

func (s *QueryWithSQLRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryWithSQLRequest) GetQuery() *string {
	return s.Query
}

func (s *QueryWithSQLRequest) SetLimit(v int32) *QueryWithSQLRequest {
	s.Limit = &v
	return s
}

func (s *QueryWithSQLRequest) SetQuery(v string) *QueryWithSQLRequest {
	s.Query = &v
	return s
}

func (s *QueryWithSQLRequest) Validate() error {
	return dara.Validate(s)
}
