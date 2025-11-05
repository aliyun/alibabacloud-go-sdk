// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditZeroCreditShutdownResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditZeroCreditShutdownResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditZeroCreditShutdownResponse
  GetStatusCode() *int32 
  SetBody(v *EditZeroCreditShutdownResponseBody) *EditZeroCreditShutdownResponse
  GetBody() *EditZeroCreditShutdownResponseBody 
}

type EditZeroCreditShutdownResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditZeroCreditShutdownResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditZeroCreditShutdownResponse) String() string {
  return dara.Prettify(s)
}

func (s EditZeroCreditShutdownResponse) GoString() string {
  return s.String()
}

func (s *EditZeroCreditShutdownResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditZeroCreditShutdownResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditZeroCreditShutdownResponse) GetBody() *EditZeroCreditShutdownResponseBody  {
  return s.Body
}

func (s *EditZeroCreditShutdownResponse) SetHeaders(v map[string]*string) *EditZeroCreditShutdownResponse {
  s.Headers = v
  return s
}

func (s *EditZeroCreditShutdownResponse) SetStatusCode(v int32) *EditZeroCreditShutdownResponse {
  s.StatusCode = &v
  return s
}

func (s *EditZeroCreditShutdownResponse) SetBody(v *EditZeroCreditShutdownResponseBody) *EditZeroCreditShutdownResponse {
  s.Body = v
  return s
}

func (s *EditZeroCreditShutdownResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

