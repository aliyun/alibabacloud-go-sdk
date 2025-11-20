// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePlatformTaskHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExecutePlatformTaskHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContext(v *ExecutePlatformTaskHeadersAccountContext) *ExecutePlatformTaskHeaders
  GetAccountContext() *ExecutePlatformTaskHeadersAccountContext 
}

type ExecutePlatformTaskHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContext *ExecutePlatformTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ExecutePlatformTaskHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExecutePlatformTaskHeaders) GoString() string {
  return s.String()
}

func (s *ExecutePlatformTaskHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExecutePlatformTaskHeaders) GetAccountContext() *ExecutePlatformTaskHeadersAccountContext  {
  return s.AccountContext
}

func (s *ExecutePlatformTaskHeaders) SetCommonHeaders(v map[string]*string) *ExecutePlatformTaskHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExecutePlatformTaskHeaders) SetAccountContext(v *ExecutePlatformTaskHeadersAccountContext) *ExecutePlatformTaskHeaders {
  s.AccountContext = v
  return s
}

func (s *ExecutePlatformTaskHeaders) Validate() error {
  if s.AccountContext != nil {
    if err := s.AccountContext.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecutePlatformTaskHeadersAccountContext struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 012345
  AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ExecutePlatformTaskHeadersAccountContext) String() string {
  return dara.Prettify(s)
}

func (s ExecutePlatformTaskHeadersAccountContext) GoString() string {
  return s.String()
}

func (s *ExecutePlatformTaskHeadersAccountContext) GetAccountId() *string  {
  return s.AccountId
}

func (s *ExecutePlatformTaskHeadersAccountContext) SetAccountId(v string) *ExecutePlatformTaskHeadersAccountContext {
  s.AccountId = &v
  return s
}

func (s *ExecutePlatformTaskHeadersAccountContext) Validate() error {
  return dara.Validate(s)
}

