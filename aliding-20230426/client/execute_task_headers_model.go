// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTaskHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExecuteTaskHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContext(v *ExecuteTaskHeadersAccountContext) *ExecuteTaskHeaders
  GetAccountContext() *ExecuteTaskHeadersAccountContext 
}

type ExecuteTaskHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContext *ExecuteTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ExecuteTaskHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskHeaders) GoString() string {
  return s.String()
}

func (s *ExecuteTaskHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExecuteTaskHeaders) GetAccountContext() *ExecuteTaskHeadersAccountContext  {
  return s.AccountContext
}

func (s *ExecuteTaskHeaders) SetCommonHeaders(v map[string]*string) *ExecuteTaskHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExecuteTaskHeaders) SetAccountContext(v *ExecuteTaskHeadersAccountContext) *ExecuteTaskHeaders {
  s.AccountContext = v
  return s
}

func (s *ExecuteTaskHeaders) Validate() error {
  if s.AccountContext != nil {
    if err := s.AccountContext.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteTaskHeadersAccountContext struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 012345
  AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ExecuteTaskHeadersAccountContext) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskHeadersAccountContext) GoString() string {
  return s.String()
}

func (s *ExecuteTaskHeadersAccountContext) GetAccountId() *string  {
  return s.AccountId
}

func (s *ExecuteTaskHeadersAccountContext) SetAccountId(v string) *ExecuteTaskHeadersAccountContext {
  s.AccountId = &v
  return s
}

func (s *ExecuteTaskHeadersAccountContext) Validate() error {
  return dara.Validate(s)
}

