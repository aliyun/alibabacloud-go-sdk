// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteBatchTaskHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExecuteBatchTaskHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContext(v *ExecuteBatchTaskHeadersAccountContext) *ExecuteBatchTaskHeaders
  GetAccountContext() *ExecuteBatchTaskHeadersAccountContext 
}

type ExecuteBatchTaskHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContext *ExecuteBatchTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ExecuteBatchTaskHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBatchTaskHeaders) GoString() string {
  return s.String()
}

func (s *ExecuteBatchTaskHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExecuteBatchTaskHeaders) GetAccountContext() *ExecuteBatchTaskHeadersAccountContext  {
  return s.AccountContext
}

func (s *ExecuteBatchTaskHeaders) SetCommonHeaders(v map[string]*string) *ExecuteBatchTaskHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExecuteBatchTaskHeaders) SetAccountContext(v *ExecuteBatchTaskHeadersAccountContext) *ExecuteBatchTaskHeaders {
  s.AccountContext = v
  return s
}

func (s *ExecuteBatchTaskHeaders) Validate() error {
  return dara.Validate(s)
}

type ExecuteBatchTaskHeadersAccountContext struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 012345
  AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ExecuteBatchTaskHeadersAccountContext) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBatchTaskHeadersAccountContext) GoString() string {
  return s.String()
}

func (s *ExecuteBatchTaskHeadersAccountContext) GetAccountId() *string  {
  return s.AccountId
}

func (s *ExecuteBatchTaskHeadersAccountContext) SetAccountId(v string) *ExecuteBatchTaskHeadersAccountContext {
  s.AccountId = &v
  return s
}

func (s *ExecuteBatchTaskHeadersAccountContext) Validate() error {
  return dara.Validate(s)
}

