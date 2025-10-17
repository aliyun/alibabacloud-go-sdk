// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistRunningSQLEngineResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExistRunningSQLEngineResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExistRunningSQLEngineResponse
  GetStatusCode() *int32 
  SetBody(v *ExistRunningSQLEngineResponseBody) *ExistRunningSQLEngineResponse
  GetBody() *ExistRunningSQLEngineResponseBody 
}

type ExistRunningSQLEngineResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExistRunningSQLEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExistRunningSQLEngineResponse) String() string {
  return dara.Prettify(s)
}

func (s ExistRunningSQLEngineResponse) GoString() string {
  return s.String()
}

func (s *ExistRunningSQLEngineResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExistRunningSQLEngineResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExistRunningSQLEngineResponse) GetBody() *ExistRunningSQLEngineResponseBody  {
  return s.Body
}

func (s *ExistRunningSQLEngineResponse) SetHeaders(v map[string]*string) *ExistRunningSQLEngineResponse {
  s.Headers = v
  return s
}

func (s *ExistRunningSQLEngineResponse) SetStatusCode(v int32) *ExistRunningSQLEngineResponse {
  s.StatusCode = &v
  return s
}

func (s *ExistRunningSQLEngineResponse) SetBody(v *ExistRunningSQLEngineResponseBody) *ExistRunningSQLEngineResponse {
  s.Body = v
  return s
}

func (s *ExistRunningSQLEngineResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

