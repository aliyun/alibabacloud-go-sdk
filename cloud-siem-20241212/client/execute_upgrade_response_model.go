// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteUpgradeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteUpgradeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteUpgradeResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteUpgradeResponseBody) *ExecuteUpgradeResponse
  GetBody() *ExecuteUpgradeResponseBody 
}

type ExecuteUpgradeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteUpgradeResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteUpgradeResponse) GoString() string {
  return s.String()
}

func (s *ExecuteUpgradeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteUpgradeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteUpgradeResponse) GetBody() *ExecuteUpgradeResponseBody  {
  return s.Body
}

func (s *ExecuteUpgradeResponse) SetHeaders(v map[string]*string) *ExecuteUpgradeResponse {
  s.Headers = v
  return s
}

func (s *ExecuteUpgradeResponse) SetStatusCode(v int32) *ExecuteUpgradeResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteUpgradeResponse) SetBody(v *ExecuteUpgradeResponseBody) *ExecuteUpgradeResponse {
  s.Body = v
  return s
}

func (s *ExecuteUpgradeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

