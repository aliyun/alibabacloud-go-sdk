// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlExecutionFetchResult interface {
	dara.Model
	String() string
	GoString() string
	SetDqlResult(v *DqlResult) *SqlExecutionFetchResult
	GetDqlResult() *DqlResult
	SetSqlExecution(v *SqlExecution) *SqlExecutionFetchResult
	GetSqlExecution() *SqlExecution
}

type SqlExecutionFetchResult struct {
	DqlResult    *DqlResult    `json:"dqlResult,omitempty" xml:"dqlResult,omitempty"`
	SqlExecution *SqlExecution `json:"sqlExecution,omitempty" xml:"sqlExecution,omitempty"`
}

func (s SqlExecutionFetchResult) String() string {
	return dara.Prettify(s)
}

func (s SqlExecutionFetchResult) GoString() string {
	return s.String()
}

func (s *SqlExecutionFetchResult) GetDqlResult() *DqlResult {
	return s.DqlResult
}

func (s *SqlExecutionFetchResult) GetSqlExecution() *SqlExecution {
	return s.SqlExecution
}

func (s *SqlExecutionFetchResult) SetDqlResult(v *DqlResult) *SqlExecutionFetchResult {
	s.DqlResult = v
	return s
}

func (s *SqlExecutionFetchResult) SetSqlExecution(v *SqlExecution) *SqlExecutionFetchResult {
	s.SqlExecution = v
	return s
}

func (s *SqlExecutionFetchResult) Validate() error {
	if s.DqlResult != nil {
		if err := s.DqlResult.Validate(); err != nil {
			return err
		}
	}
	if s.SqlExecution != nil {
		if err := s.SqlExecution.Validate(); err != nil {
			return err
		}
	}
	return nil
}
