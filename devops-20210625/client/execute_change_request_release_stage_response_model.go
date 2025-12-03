// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteChangeRequestReleaseStageResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteChangeRequestReleaseStageResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteChangeRequestReleaseStageResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteChangeRequestReleaseStageResponseBody) *ExecuteChangeRequestReleaseStageResponse
  GetBody() *ExecuteChangeRequestReleaseStageResponseBody 
}

type ExecuteChangeRequestReleaseStageResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteChangeRequestReleaseStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteChangeRequestReleaseStageResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteChangeRequestReleaseStageResponse) GoString() string {
  return s.String()
}

func (s *ExecuteChangeRequestReleaseStageResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteChangeRequestReleaseStageResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteChangeRequestReleaseStageResponse) GetBody() *ExecuteChangeRequestReleaseStageResponseBody  {
  return s.Body
}

func (s *ExecuteChangeRequestReleaseStageResponse) SetHeaders(v map[string]*string) *ExecuteChangeRequestReleaseStageResponse {
  s.Headers = v
  return s
}

func (s *ExecuteChangeRequestReleaseStageResponse) SetStatusCode(v int32) *ExecuteChangeRequestReleaseStageResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteChangeRequestReleaseStageResponse) SetBody(v *ExecuteChangeRequestReleaseStageResponseBody) *ExecuteChangeRequestReleaseStageResponse {
  s.Body = v
  return s
}

func (s *ExecuteChangeRequestReleaseStageResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

