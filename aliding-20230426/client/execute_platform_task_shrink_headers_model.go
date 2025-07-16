// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePlatformTaskShrinkHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExecutePlatformTaskShrinkHeaders
  GetCommonHeaders() map[string]*string 
  SetAccountContextShrink(v string) *ExecutePlatformTaskShrinkHeaders
  GetAccountContextShrink() *string 
}

type ExecutePlatformTaskShrinkHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  AccountContextShrink *string `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ExecutePlatformTaskShrinkHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExecutePlatformTaskShrinkHeaders) GoString() string {
  return s.String()
}

func (s *ExecutePlatformTaskShrinkHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExecutePlatformTaskShrinkHeaders) GetAccountContextShrink() *string  {
  return s.AccountContextShrink
}

func (s *ExecutePlatformTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *ExecutePlatformTaskShrinkHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExecutePlatformTaskShrinkHeaders) SetAccountContextShrink(v string) *ExecutePlatformTaskShrinkHeaders {
  s.AccountContextShrink = &v
  return s
}

func (s *ExecutePlatformTaskShrinkHeaders) Validate() error {
  return dara.Validate(s)
}

