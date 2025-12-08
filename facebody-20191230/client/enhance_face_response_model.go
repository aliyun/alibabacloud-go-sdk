// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceFaceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnhanceFaceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnhanceFaceResponse
  GetStatusCode() *int32 
  SetBody(v *EnhanceFaceResponseBody) *EnhanceFaceResponse
  GetBody() *EnhanceFaceResponseBody 
}

type EnhanceFaceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnhanceFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnhanceFaceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnhanceFaceResponse) GoString() string {
  return s.String()
}

func (s *EnhanceFaceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnhanceFaceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnhanceFaceResponse) GetBody() *EnhanceFaceResponseBody  {
  return s.Body
}

func (s *EnhanceFaceResponse) SetHeaders(v map[string]*string) *EnhanceFaceResponse {
  s.Headers = v
  return s
}

func (s *EnhanceFaceResponse) SetStatusCode(v int32) *EnhanceFaceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnhanceFaceResponse) SetBody(v *EnhanceFaceResponseBody) *EnhanceFaceResponse {
  s.Body = v
  return s
}

func (s *EnhanceFaceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

