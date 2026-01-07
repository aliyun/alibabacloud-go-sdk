// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactQueryDetailResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseContactQueryDetailResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseContactQueryDetailResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseContactQueryDetailResponseBody) *EnterpriseContactQueryDetailResponse
  GetBody() *EnterpriseContactQueryDetailResponseBody 
}

type EnterpriseContactQueryDetailResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseContactQueryDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseContactQueryDetailResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactQueryDetailResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseContactQueryDetailResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseContactQueryDetailResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseContactQueryDetailResponse) GetBody() *EnterpriseContactQueryDetailResponseBody  {
  return s.Body
}

func (s *EnterpriseContactQueryDetailResponse) SetHeaders(v map[string]*string) *EnterpriseContactQueryDetailResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseContactQueryDetailResponse) SetStatusCode(v int32) *EnterpriseContactQueryDetailResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponse) SetBody(v *EnterpriseContactQueryDetailResponseBody) *EnterpriseContactQueryDetailResponse {
  s.Body = v
  return s
}

func (s *EnterpriseContactQueryDetailResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

