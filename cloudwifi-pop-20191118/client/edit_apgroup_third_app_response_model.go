// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditApgroupThirdAppResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditApgroupThirdAppResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditApgroupThirdAppResponse
  GetStatusCode() *int32 
  SetBody(v *EditApgroupThirdAppResponseBody) *EditApgroupThirdAppResponse
  GetBody() *EditApgroupThirdAppResponseBody 
}

type EditApgroupThirdAppResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditApgroupThirdAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditApgroupThirdAppResponse) String() string {
  return dara.Prettify(s)
}

func (s EditApgroupThirdAppResponse) GoString() string {
  return s.String()
}

func (s *EditApgroupThirdAppResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditApgroupThirdAppResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditApgroupThirdAppResponse) GetBody() *EditApgroupThirdAppResponseBody  {
  return s.Body
}

func (s *EditApgroupThirdAppResponse) SetHeaders(v map[string]*string) *EditApgroupThirdAppResponse {
  s.Headers = v
  return s
}

func (s *EditApgroupThirdAppResponse) SetStatusCode(v int32) *EditApgroupThirdAppResponse {
  s.StatusCode = &v
  return s
}

func (s *EditApgroupThirdAppResponse) SetBody(v *EditApgroupThirdAppResponseBody) *EditApgroupThirdAppResponse {
  s.Body = v
  return s
}

func (s *EditApgroupThirdAppResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

