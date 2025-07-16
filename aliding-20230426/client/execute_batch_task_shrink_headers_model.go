// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteBatchTaskShrinkHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExecuteBatchTaskShrinkHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContextShrink(v string) *ExecuteBatchTaskShrinkHeaders
  GetAccountContextShrink() *string 
}

type ExecuteBatchTaskShrinkHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContextShrink *string `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ExecuteBatchTaskShrinkHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBatchTaskShrinkHeaders) GoString() string {
  return s.String()
}

func (s *ExecuteBatchTaskShrinkHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExecuteBatchTaskShrinkHeaders) GetAccountContextShrink() *string  {
  return s.AccountContextShrink
}

func (s *ExecuteBatchTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *ExecuteBatchTaskShrinkHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExecuteBatchTaskShrinkHeaders) SetAccountContextShrink(v string) *ExecuteBatchTaskShrinkHeaders {
  s.AccountContextShrink = &v
  return s
}

func (s *ExecuteBatchTaskShrinkHeaders) Validate() error {
  return dara.Validate(s)
}

