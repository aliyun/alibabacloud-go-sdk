// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRightsSeparationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableRightsSeparationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableRightsSeparationResponse
  GetStatusCode() *int32 
  SetBody(v *EnableRightsSeparationResponseBody) *EnableRightsSeparationResponse
  GetBody() *EnableRightsSeparationResponseBody 
}

type EnableRightsSeparationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableRightsSeparationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableRightsSeparationResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableRightsSeparationResponse) GoString() string {
  return s.String()
}

func (s *EnableRightsSeparationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableRightsSeparationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableRightsSeparationResponse) GetBody() *EnableRightsSeparationResponseBody  {
  return s.Body
}

func (s *EnableRightsSeparationResponse) SetHeaders(v map[string]*string) *EnableRightsSeparationResponse {
  s.Headers = v
  return s
}

func (s *EnableRightsSeparationResponse) SetStatusCode(v int32) *EnableRightsSeparationResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableRightsSeparationResponse) SetBody(v *EnableRightsSeparationResponseBody) *EnableRightsSeparationResponse {
  s.Body = v
  return s
}

func (s *EnableRightsSeparationResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

