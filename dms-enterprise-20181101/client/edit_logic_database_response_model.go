// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditLogicDatabaseResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditLogicDatabaseResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditLogicDatabaseResponse
  GetStatusCode() *int32 
  SetBody(v *EditLogicDatabaseResponseBody) *EditLogicDatabaseResponse
  GetBody() *EditLogicDatabaseResponseBody 
}

type EditLogicDatabaseResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditLogicDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditLogicDatabaseResponse) String() string {
  return dara.Prettify(s)
}

func (s EditLogicDatabaseResponse) GoString() string {
  return s.String()
}

func (s *EditLogicDatabaseResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditLogicDatabaseResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditLogicDatabaseResponse) GetBody() *EditLogicDatabaseResponseBody  {
  return s.Body
}

func (s *EditLogicDatabaseResponse) SetHeaders(v map[string]*string) *EditLogicDatabaseResponse {
  s.Headers = v
  return s
}

func (s *EditLogicDatabaseResponse) SetStatusCode(v int32) *EditLogicDatabaseResponse {
  s.StatusCode = &v
  return s
}

func (s *EditLogicDatabaseResponse) SetBody(v *EditLogicDatabaseResponseBody) *EditLogicDatabaseResponse {
  s.Body = v
  return s
}

func (s *EditLogicDatabaseResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

