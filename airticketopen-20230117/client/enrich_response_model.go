// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrichResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnrichResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnrichResponse
  GetStatusCode() *int32 
  SetBody(v *EnrichResponseBody) *EnrichResponse
  GetBody() *EnrichResponseBody 
}

type EnrichResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnrichResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnrichResponse) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponse) GoString() string {
  return s.String()
}

func (s *EnrichResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnrichResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnrichResponse) GetBody() *EnrichResponseBody  {
  return s.Body
}

func (s *EnrichResponse) SetHeaders(v map[string]*string) *EnrichResponse {
  s.Headers = v
  return s
}

func (s *EnrichResponse) SetStatusCode(v int32) *EnrichResponse {
  s.StatusCode = &v
  return s
}

func (s *EnrichResponse) SetBody(v *EnrichResponseBody) *EnrichResponse {
  s.Body = v
  return s
}

func (s *EnrichResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

