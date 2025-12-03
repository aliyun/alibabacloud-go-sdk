// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndAllCoordinationByOwnerResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndAllCoordinationByOwnerResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndAllCoordinationByOwnerResponse
  GetStatusCode() *int32 
  SetBody(v *EndAllCoordinationByOwnerResponseBody) *EndAllCoordinationByOwnerResponse
  GetBody() *EndAllCoordinationByOwnerResponseBody 
}

type EndAllCoordinationByOwnerResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndAllCoordinationByOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndAllCoordinationByOwnerResponse) String() string {
  return dara.Prettify(s)
}

func (s EndAllCoordinationByOwnerResponse) GoString() string {
  return s.String()
}

func (s *EndAllCoordinationByOwnerResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndAllCoordinationByOwnerResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndAllCoordinationByOwnerResponse) GetBody() *EndAllCoordinationByOwnerResponseBody  {
  return s.Body
}

func (s *EndAllCoordinationByOwnerResponse) SetHeaders(v map[string]*string) *EndAllCoordinationByOwnerResponse {
  s.Headers = v
  return s
}

func (s *EndAllCoordinationByOwnerResponse) SetStatusCode(v int32) *EndAllCoordinationByOwnerResponse {
  s.StatusCode = &v
  return s
}

func (s *EndAllCoordinationByOwnerResponse) SetBody(v *EndAllCoordinationByOwnerResponseBody) *EndAllCoordinationByOwnerResponse {
  s.Body = v
  return s
}

func (s *EndAllCoordinationByOwnerResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

