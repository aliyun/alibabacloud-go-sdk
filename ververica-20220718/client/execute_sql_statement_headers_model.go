// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSqlStatementHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExecuteSqlStatementHeaders
  GetCommonHeaders() map[string]*string 
  SetWorkspace(v string) *ExecuteSqlStatementHeaders
  GetWorkspace() *string 
}

type ExecuteSqlStatementHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 710d6a64d8c34d
  Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ExecuteSqlStatementHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSqlStatementHeaders) GoString() string {
  return s.String()
}

func (s *ExecuteSqlStatementHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExecuteSqlStatementHeaders) GetWorkspace() *string  {
  return s.Workspace
}

func (s *ExecuteSqlStatementHeaders) SetCommonHeaders(v map[string]*string) *ExecuteSqlStatementHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExecuteSqlStatementHeaders) SetWorkspace(v string) *ExecuteSqlStatementHeaders {
  s.Workspace = &v
  return s
}

func (s *ExecuteSqlStatementHeaders) Validate() error {
  return dara.Validate(s)
}

