// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedQueryHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuerySql(v string) *CreateAdvancedQueryHistoryRequest
	GetQuerySql() *string
	SetSimpleQuery(v bool) *CreateAdvancedQueryHistoryRequest
	GetSimpleQuery() *bool
}

type CreateAdvancedQueryHistoryRequest struct {
	// example:
	//
	// event.userIdentity.accessKeyId: *
	QuerySql *string `json:"QuerySql,omitempty" xml:"QuerySql,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	SimpleQuery *bool `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
}

func (s CreateAdvancedQueryHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedQueryHistoryRequest) GoString() string {
	return s.String()
}

func (s *CreateAdvancedQueryHistoryRequest) GetQuerySql() *string {
	return s.QuerySql
}

func (s *CreateAdvancedQueryHistoryRequest) GetSimpleQuery() *bool {
	return s.SimpleQuery
}

func (s *CreateAdvancedQueryHistoryRequest) SetQuerySql(v string) *CreateAdvancedQueryHistoryRequest {
	s.QuerySql = &v
	return s
}

func (s *CreateAdvancedQueryHistoryRequest) SetSimpleQuery(v bool) *CreateAdvancedQueryHistoryRequest {
	s.SimpleQuery = &v
	return s
}

func (s *CreateAdvancedQueryHistoryRequest) Validate() error {
	return dara.Validate(s)
}
