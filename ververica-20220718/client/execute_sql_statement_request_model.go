// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSqlStatementRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBody(v *SqlStatementWithContext) *ExecuteSqlStatementRequest
  GetBody() *SqlStatementWithContext 
}

type ExecuteSqlStatementRequest struct {
  Body *SqlStatementWithContext `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSqlStatementRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSqlStatementRequest) GoString() string {
  return s.String()
}

func (s *ExecuteSqlStatementRequest) GetBody() *SqlStatementWithContext  {
  return s.Body
}

func (s *ExecuteSqlStatementRequest) SetBody(v *SqlStatementWithContext) *ExecuteSqlStatementRequest {
  s.Body = v
  return s
}

func (s *ExecuteSqlStatementRequest) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

