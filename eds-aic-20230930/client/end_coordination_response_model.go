// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndCoordinationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndCoordinationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndCoordinationResponse
  GetStatusCode() *int32 
  SetBody(v *EndCoordinationResponseBody) *EndCoordinationResponse
  GetBody() *EndCoordinationResponseBody 
}

type EndCoordinationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndCoordinationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndCoordinationResponse) String() string {
  return dara.Prettify(s)
}

func (s EndCoordinationResponse) GoString() string {
  return s.String()
}

func (s *EndCoordinationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndCoordinationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndCoordinationResponse) GetBody() *EndCoordinationResponseBody  {
  return s.Body
}

func (s *EndCoordinationResponse) SetHeaders(v map[string]*string) *EndCoordinationResponse {
  s.Headers = v
  return s
}

func (s *EndCoordinationResponse) SetStatusCode(v int32) *EndCoordinationResponse {
  s.StatusCode = &v
  return s
}

func (s *EndCoordinationResponse) SetBody(v *EndCoordinationResponseBody) *EndCoordinationResponse {
  s.Body = v
  return s
}

func (s *EndCoordinationResponse) Validate() error {
  return dara.Validate(s)
}

