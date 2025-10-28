// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServicesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableServicesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableServicesResponse
  GetStatusCode() *int32 
  SetBody(v *EnableServicesResponseBody) *EnableServicesResponse
  GetBody() *EnableServicesResponseBody 
}

type EnableServicesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableServicesResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableServicesResponse) GoString() string {
  return s.String()
}

func (s *EnableServicesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableServicesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableServicesResponse) GetBody() *EnableServicesResponseBody  {
  return s.Body
}

func (s *EnableServicesResponse) SetHeaders(v map[string]*string) *EnableServicesResponse {
  s.Headers = v
  return s
}

func (s *EnableServicesResponse) SetStatusCode(v int32) *EnableServicesResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableServicesResponse) SetBody(v *EnableServicesResponseBody) *EnableServicesResponse {
  s.Body = v
  return s
}

func (s *EnableServicesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

