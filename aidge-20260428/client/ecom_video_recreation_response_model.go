// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcomVideoRecreationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EcomVideoRecreationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EcomVideoRecreationResponse
  GetStatusCode() *int32 
  SetBody(v *EcomVideoRecreationResponseBody) *EcomVideoRecreationResponse
  GetBody() *EcomVideoRecreationResponseBody 
}

type EcomVideoRecreationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EcomVideoRecreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EcomVideoRecreationResponse) String() string {
  return dara.Prettify(s)
}

func (s EcomVideoRecreationResponse) GoString() string {
  return s.String()
}

func (s *EcomVideoRecreationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EcomVideoRecreationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EcomVideoRecreationResponse) GetBody() *EcomVideoRecreationResponseBody  {
  return s.Body
}

func (s *EcomVideoRecreationResponse) SetHeaders(v map[string]*string) *EcomVideoRecreationResponse {
  s.Headers = v
  return s
}

func (s *EcomVideoRecreationResponse) SetStatusCode(v int32) *EcomVideoRecreationResponse {
  s.StatusCode = &v
  return s
}

func (s *EcomVideoRecreationResponse) SetBody(v *EcomVideoRecreationResponseBody) *EcomVideoRecreationResponse {
  s.Body = v
  return s
}

func (s *EcomVideoRecreationResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

