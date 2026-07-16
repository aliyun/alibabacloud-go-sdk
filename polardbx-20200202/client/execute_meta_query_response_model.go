// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteMetaQueryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteMetaQueryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteMetaQueryResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteMetaQueryResponseBody) *ExecuteMetaQueryResponse
  GetBody() *ExecuteMetaQueryResponseBody 
}

type ExecuteMetaQueryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteMetaQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteMetaQueryResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMetaQueryResponse) GoString() string {
  return s.String()
}

func (s *ExecuteMetaQueryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteMetaQueryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteMetaQueryResponse) GetBody() *ExecuteMetaQueryResponseBody  {
  return s.Body
}

func (s *ExecuteMetaQueryResponse) SetHeaders(v map[string]*string) *ExecuteMetaQueryResponse {
  s.Headers = v
  return s
}

func (s *ExecuteMetaQueryResponse) SetStatusCode(v int32) *ExecuteMetaQueryResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteMetaQueryResponse) SetBody(v *ExecuteMetaQueryResponseBody) *ExecuteMetaQueryResponse {
  s.Body = v
  return s
}

func (s *ExecuteMetaQueryResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

