// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteStructSyncResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteStructSyncResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteStructSyncResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteStructSyncResponseBody) *ExecuteStructSyncResponse
  GetBody() *ExecuteStructSyncResponseBody 
}

type ExecuteStructSyncResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteStructSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteStructSyncResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStructSyncResponse) GoString() string {
  return s.String()
}

func (s *ExecuteStructSyncResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteStructSyncResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteStructSyncResponse) GetBody() *ExecuteStructSyncResponseBody  {
  return s.Body
}

func (s *ExecuteStructSyncResponse) SetHeaders(v map[string]*string) *ExecuteStructSyncResponse {
  s.Headers = v
  return s
}

func (s *ExecuteStructSyncResponse) SetStatusCode(v int32) *ExecuteStructSyncResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteStructSyncResponse) SetBody(v *ExecuteStructSyncResponseBody) *ExecuteStructSyncResponse {
  s.Body = v
  return s
}

func (s *ExecuteStructSyncResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

