// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesByAdvancedSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSql(v string) *ListResourcesByAdvancedSearchRequest
	GetSql() *string
}

type ListResourcesByAdvancedSearchRequest struct {
	// The SQL query statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT ResourceId, ResourceName WHERE Tags.Kvpair=\\"business:online\\"
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s ListResourcesByAdvancedSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByAdvancedSearchRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesByAdvancedSearchRequest) GetSql() *string {
	return s.Sql
}

func (s *ListResourcesByAdvancedSearchRequest) SetSql(v string) *ListResourcesByAdvancedSearchRequest {
	s.Sql = &v
	return s
}

func (s *ListResourcesByAdvancedSearchRequest) Validate() error {
	return dara.Validate(s)
}
