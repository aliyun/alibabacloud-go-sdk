// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEngineSearchResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EngineSearchResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EngineSearchResponse
  GetStatusCode() *int32 
  SetBody(v *EngineSearchResponseBody) *EngineSearchResponse
  GetBody() *EngineSearchResponseBody 
}

type EngineSearchResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EngineSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EngineSearchResponse) String() string {
  return dara.Prettify(s)
}

func (s EngineSearchResponse) GoString() string {
  return s.String()
}

func (s *EngineSearchResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EngineSearchResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EngineSearchResponse) GetBody() *EngineSearchResponseBody  {
  return s.Body
}

func (s *EngineSearchResponse) SetHeaders(v map[string]*string) *EngineSearchResponse {
  s.Headers = v
  return s
}

func (s *EngineSearchResponse) SetStatusCode(v int32) *EngineSearchResponse {
  s.StatusCode = &v
  return s
}

func (s *EngineSearchResponse) SetBody(v *EngineSearchResponseBody) *EngineSearchResponse {
  s.Body = v
  return s
}

func (s *EngineSearchResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

