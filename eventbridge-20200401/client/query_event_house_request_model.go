// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventHouseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *QueryEventHouseRequest
	GetLimit() *int32
	SetQuery(v string) *QueryEventHouseRequest
	GetQuery() *string
}

type QueryEventHouseRequest struct {
	// The maximum number of rows to return. If a LIMIT clause is specified in the SQL statement, the LIMIT clause in the SQL statement prevails.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The query statement. A three-part table name in the format of catalog.namespace.table is typically used.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM "test-es"."default"."product_info"
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s QueryEventHouseRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEventHouseRequest) GoString() string {
	return s.String()
}

func (s *QueryEventHouseRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryEventHouseRequest) GetQuery() *string {
	return s.Query
}

func (s *QueryEventHouseRequest) SetLimit(v int32) *QueryEventHouseRequest {
	s.Limit = &v
	return s
}

func (s *QueryEventHouseRequest) SetQuery(v string) *QueryEventHouseRequest {
	s.Query = &v
	return s
}

func (s *QueryEventHouseRequest) Validate() error {
	return dara.Validate(s)
}
