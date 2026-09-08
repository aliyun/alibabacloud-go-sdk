// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAttachment interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *QueryAttachment
	GetQuery() *string
	SetQueryId(v string) *QueryAttachment
	GetQueryId() *string
	SetResult(v *ExecutionResult) *QueryAttachment
	GetResult() *ExecutionResult
}

type QueryAttachment struct {
	// The query statement.
	//
	// example:
	//
	// "SELECT city, COUNT(*) AS cnt FROM events GROUP BY city"
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The stable identifier for the actual SQL tool execution, used for result tracking and interpretation.
	//
	// example:
	//
	// sqlx-0-a1b2c3d4
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The query execution result.
	Result *ExecutionResult `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s QueryAttachment) String() string {
	return dara.Prettify(s)
}

func (s QueryAttachment) GoString() string {
	return s.String()
}

func (s *QueryAttachment) GetQuery() *string {
	return s.Query
}

func (s *QueryAttachment) GetQueryId() *string {
	return s.QueryId
}

func (s *QueryAttachment) GetResult() *ExecutionResult {
	return s.Result
}

func (s *QueryAttachment) SetQuery(v string) *QueryAttachment {
	s.Query = &v
	return s
}

func (s *QueryAttachment) SetQueryId(v string) *QueryAttachment {
	s.QueryId = &v
	return s
}

func (s *QueryAttachment) SetResult(v *ExecutionResult) *QueryAttachment {
	s.Result = v
	return s
}

func (s *QueryAttachment) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
