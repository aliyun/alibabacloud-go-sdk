// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTaskShrinkHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExecuteTaskShrinkHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContextShrink(v string) *ExecuteTaskShrinkHeaders
  GetAccountContextShrink() *string 
}

type ExecuteTaskShrinkHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContextShrink *string `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ExecuteTaskShrinkHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskShrinkHeaders) GoString() string {
  return s.String()
}

func (s *ExecuteTaskShrinkHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExecuteTaskShrinkHeaders) GetAccountContextShrink() *string  {
  return s.AccountContextShrink
}

func (s *ExecuteTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *ExecuteTaskShrinkHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExecuteTaskShrinkHeaders) SetAccountContextShrink(v string) *ExecuteTaskShrinkHeaders {
  s.AccountContextShrink = &v
  return s
}

func (s *ExecuteTaskShrinkHeaders) Validate() error {
  return dara.Validate(s)
}

