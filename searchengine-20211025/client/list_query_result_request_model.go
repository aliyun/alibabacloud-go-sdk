// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *ListQueryResultRequest
	GetQuery() *string
	SetSql(v string) *ListQueryResultRequest
	GetSql() *string
}

type ListQueryResultRequest struct {
	// The query statement
	//
	// example:
	//
	// query%3D1%26%26config%3Dstart%3A0%2Chit%3A10%2Cformat%3Ajson%26%26cluster%3Dgeneral
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The SQL statement that is executed in the query
	//
	// example:
	//
	// query%3Dselect%20max(content_id)%20from%20generation
	Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
}

func (s ListQueryResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueryResultRequest) GoString() string {
	return s.String()
}

func (s *ListQueryResultRequest) GetQuery() *string {
	return s.Query
}

func (s *ListQueryResultRequest) GetSql() *string {
	return s.Sql
}

func (s *ListQueryResultRequest) SetQuery(v string) *ListQueryResultRequest {
	s.Query = &v
	return s
}

func (s *ListQueryResultRequest) SetSql(v string) *ListQueryResultRequest {
	s.Sql = &v
	return s
}

func (s *ListQueryResultRequest) Validate() error {
	return dara.Validate(s)
}
