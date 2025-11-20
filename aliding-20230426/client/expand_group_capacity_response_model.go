// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandGroupCapacityResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExpandGroupCapacityResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExpandGroupCapacityResponse
  GetStatusCode() *int32 
  SetBody(v *ExpandGroupCapacityResponseBody) *ExpandGroupCapacityResponse
  GetBody() *ExpandGroupCapacityResponseBody 
}

type ExpandGroupCapacityResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExpandGroupCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpandGroupCapacityResponse) String() string {
  return dara.Prettify(s)
}

func (s ExpandGroupCapacityResponse) GoString() string {
  return s.String()
}

func (s *ExpandGroupCapacityResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExpandGroupCapacityResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExpandGroupCapacityResponse) GetBody() *ExpandGroupCapacityResponseBody  {
  return s.Body
}

func (s *ExpandGroupCapacityResponse) SetHeaders(v map[string]*string) *ExpandGroupCapacityResponse {
  s.Headers = v
  return s
}

func (s *ExpandGroupCapacityResponse) SetStatusCode(v int32) *ExpandGroupCapacityResponse {
  s.StatusCode = &v
  return s
}

func (s *ExpandGroupCapacityResponse) SetBody(v *ExpandGroupCapacityResponseBody) *ExpandGroupCapacityResponse {
  s.Body = v
  return s
}

func (s *ExpandGroupCapacityResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

