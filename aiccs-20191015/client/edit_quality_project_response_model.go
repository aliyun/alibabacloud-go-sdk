// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityProjectResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditQualityProjectResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditQualityProjectResponse
  GetStatusCode() *int32 
  SetBody(v *EditQualityProjectResponseBody) *EditQualityProjectResponse
  GetBody() *EditQualityProjectResponseBody 
}

type EditQualityProjectResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditQualityProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditQualityProjectResponse) String() string {
  return dara.Prettify(s)
}

func (s EditQualityProjectResponse) GoString() string {
  return s.String()
}

func (s *EditQualityProjectResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditQualityProjectResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditQualityProjectResponse) GetBody() *EditQualityProjectResponseBody  {
  return s.Body
}

func (s *EditQualityProjectResponse) SetHeaders(v map[string]*string) *EditQualityProjectResponse {
  s.Headers = v
  return s
}

func (s *EditQualityProjectResponse) SetStatusCode(v int32) *EditQualityProjectResponse {
  s.StatusCode = &v
  return s
}

func (s *EditQualityProjectResponse) SetBody(v *EditQualityProjectResponseBody) *EditQualityProjectResponse {
  s.Body = v
  return s
}

func (s *EditQualityProjectResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

