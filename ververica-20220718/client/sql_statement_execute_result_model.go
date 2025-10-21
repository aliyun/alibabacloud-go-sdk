// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSqlStatementExecuteResult interface {
	dara.Model
	String() string
	GoString() string
	SetErrorDetails(v *ErrorDetails) *SqlStatementExecuteResult
	GetErrorDetails() *ErrorDetails
	SetExecuteSuccess(v bool) *SqlStatementExecuteResult
	GetExecuteSuccess() *bool
	SetStatement(v string) *SqlStatementExecuteResult
	GetStatement() *string
}

type SqlStatementExecuteResult struct {
	ErrorDetails   *ErrorDetails `json:"errorDetails,omitempty" xml:"errorDetails,omitempty"`
	ExecuteSuccess *bool         `json:"executeSuccess,omitempty" xml:"executeSuccess,omitempty"`
	Statement      *string       `json:"statement,omitempty" xml:"statement,omitempty"`
}

func (s SqlStatementExecuteResult) String() string {
	return dara.Prettify(s)
}

func (s SqlStatementExecuteResult) GoString() string {
	return s.String()
}

func (s *SqlStatementExecuteResult) GetErrorDetails() *ErrorDetails {
	return s.ErrorDetails
}

func (s *SqlStatementExecuteResult) GetExecuteSuccess() *bool {
	return s.ExecuteSuccess
}

func (s *SqlStatementExecuteResult) GetStatement() *string {
	return s.Statement
}

func (s *SqlStatementExecuteResult) SetErrorDetails(v *ErrorDetails) *SqlStatementExecuteResult {
	s.ErrorDetails = v
	return s
}

func (s *SqlStatementExecuteResult) SetExecuteSuccess(v bool) *SqlStatementExecuteResult {
	s.ExecuteSuccess = &v
	return s
}

func (s *SqlStatementExecuteResult) SetStatement(v string) *SqlStatementExecuteResult {
	s.Statement = &v
	return s
}

func (s *SqlStatementExecuteResult) Validate() error {
	if s.ErrorDetails != nil {
		if err := s.ErrorDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}
