// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSqlExecutionResult interface {
	dara.Model
	String() string
	GoString() string
	SetNewlyCreated(v bool) *StartSqlExecutionResult
	GetNewlyCreated() *bool
	SetSqlExecutionId(v string) *StartSqlExecutionResult
	GetSqlExecutionId() *string
	SetSuccess(v bool) *StartSqlExecutionResult
	GetSuccess() *bool
}

type StartSqlExecutionResult struct {
	NewlyCreated   *bool   `json:"newlyCreated,omitempty" xml:"newlyCreated,omitempty"`
	SqlExecutionId *string `json:"sqlExecutionId,omitempty" xml:"sqlExecutionId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartSqlExecutionResult) String() string {
	return dara.Prettify(s)
}

func (s StartSqlExecutionResult) GoString() string {
	return s.String()
}

func (s *StartSqlExecutionResult) GetNewlyCreated() *bool {
	return s.NewlyCreated
}

func (s *StartSqlExecutionResult) GetSqlExecutionId() *string {
	return s.SqlExecutionId
}

func (s *StartSqlExecutionResult) GetSuccess() *bool {
	return s.Success
}

func (s *StartSqlExecutionResult) SetNewlyCreated(v bool) *StartSqlExecutionResult {
	s.NewlyCreated = &v
	return s
}

func (s *StartSqlExecutionResult) SetSqlExecutionId(v string) *StartSqlExecutionResult {
	s.SqlExecutionId = &v
	return s
}

func (s *StartSqlExecutionResult) SetSuccess(v bool) *StartSqlExecutionResult {
	s.Success = &v
	return s
}

func (s *StartSqlExecutionResult) Validate() error {
	return dara.Validate(s)
}
