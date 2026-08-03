// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedQueryHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRunResult(v string) *CreateAdvancedQueryHistoryResponseBody
	GetDryRunResult() *string
	SetQueryId(v string) *CreateAdvancedQueryHistoryResponseBody
	GetQueryId() *string
	SetQuerySql(v string) *CreateAdvancedQueryHistoryResponseBody
	GetQuerySql() *string
	SetRequestId(v string) *CreateAdvancedQueryHistoryResponseBody
	GetRequestId() *string
	SetSimpleQuery(v bool) *CreateAdvancedQueryHistoryResponseBody
	GetSimpleQuery() *bool
}

type CreateAdvancedQueryHistoryResponseBody struct {
	DryRunResult *string `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty"`
	// The ID of the advanced event query record.
	//
	// example:
	//
	// query-uIkIvLiVSuCKqg0yoa****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The advanced event query statement.
	//
	// example:
	//
	// event.userIdentity.accessKeyId: *
	QuerySql *string `json:"QuerySql,omitempty" xml:"QuerySql,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0227506-AA8C-5998-8A62-74769106****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether to enable the simple query mode.
	//
	// example:
	//
	// false
	SimpleQuery *bool `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
}

func (s CreateAdvancedQueryHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedQueryHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdvancedQueryHistoryResponseBody) GetDryRunResult() *string {
	return s.DryRunResult
}

func (s *CreateAdvancedQueryHistoryResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *CreateAdvancedQueryHistoryResponseBody) GetQuerySql() *string {
	return s.QuerySql
}

func (s *CreateAdvancedQueryHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAdvancedQueryHistoryResponseBody) GetSimpleQuery() *bool {
	return s.SimpleQuery
}

func (s *CreateAdvancedQueryHistoryResponseBody) SetDryRunResult(v string) *CreateAdvancedQueryHistoryResponseBody {
	s.DryRunResult = &v
	return s
}

func (s *CreateAdvancedQueryHistoryResponseBody) SetQueryId(v string) *CreateAdvancedQueryHistoryResponseBody {
	s.QueryId = &v
	return s
}

func (s *CreateAdvancedQueryHistoryResponseBody) SetQuerySql(v string) *CreateAdvancedQueryHistoryResponseBody {
	s.QuerySql = &v
	return s
}

func (s *CreateAdvancedQueryHistoryResponseBody) SetRequestId(v string) *CreateAdvancedQueryHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdvancedQueryHistoryResponseBody) SetSimpleQuery(v bool) *CreateAdvancedQueryHistoryResponseBody {
	s.SimpleQuery = &v
	return s
}

func (s *CreateAdvancedQueryHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}
