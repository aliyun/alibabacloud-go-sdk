// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableJobResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableJobResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableJobResponse
  GetStatusCode() *int32 
  SetBody(v *EnableJobResponseBody) *EnableJobResponse
  GetBody() *EnableJobResponseBody 
}

type EnableJobResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableJobResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableJobResponse) GoString() string {
  return s.String()
}

func (s *EnableJobResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableJobResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableJobResponse) GetBody() *EnableJobResponseBody  {
  return s.Body
}

func (s *EnableJobResponse) SetHeaders(v map[string]*string) *EnableJobResponse {
  s.Headers = v
  return s
}

func (s *EnableJobResponse) SetStatusCode(v int32) *EnableJobResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableJobResponse) SetBody(v *EnableJobResponseBody) *EnableJobResponse {
  s.Body = v
  return s
}

func (s *EnableJobResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

