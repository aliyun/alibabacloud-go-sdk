// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditModelResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditModelResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditModelResponse
  GetStatusCode() *int32 
  SetBody(v *EditModelResponseBody) *EditModelResponse
  GetBody() *EditModelResponseBody 
}

type EditModelResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditModelResponse) String() string {
  return dara.Prettify(s)
}

func (s EditModelResponse) GoString() string {
  return s.String()
}

func (s *EditModelResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditModelResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditModelResponse) GetBody() *EditModelResponseBody  {
  return s.Body
}

func (s *EditModelResponse) SetHeaders(v map[string]*string) *EditModelResponse {
  s.Headers = v
  return s
}

func (s *EditModelResponse) SetStatusCode(v int32) *EditModelResponse {
  s.StatusCode = &v
  return s
}

func (s *EditModelResponse) SetBody(v *EditModelResponseBody) *EditModelResponse {
  s.Body = v
  return s
}

func (s *EditModelResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

