// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvictTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvictTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvictTaskResponse
  GetStatusCode() *int32 
  SetBody(v *EvictTaskResponseBody) *EvictTaskResponse
  GetBody() *EvictTaskResponseBody 
}

type EvictTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvictTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvictTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s EvictTaskResponse) GoString() string {
  return s.String()
}

func (s *EvictTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvictTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvictTaskResponse) GetBody() *EvictTaskResponseBody  {
  return s.Body
}

func (s *EvictTaskResponse) SetHeaders(v map[string]*string) *EvictTaskResponse {
  s.Headers = v
  return s
}

func (s *EvictTaskResponse) SetStatusCode(v int32) *EvictTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *EvictTaskResponse) SetBody(v *EvictTaskResponseBody) *EvictTaskResponse {
  s.Body = v
  return s
}

func (s *EvictTaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

