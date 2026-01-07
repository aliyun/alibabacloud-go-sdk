// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactAddResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseContactAddResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseContactAddResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseContactAddResponseBody) *EnterpriseContactAddResponse
  GetBody() *EnterpriseContactAddResponseBody 
}

type EnterpriseContactAddResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseContactAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseContactAddResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactAddResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseContactAddResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseContactAddResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseContactAddResponse) GetBody() *EnterpriseContactAddResponseBody  {
  return s.Body
}

func (s *EnterpriseContactAddResponse) SetHeaders(v map[string]*string) *EnterpriseContactAddResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseContactAddResponse) SetStatusCode(v int32) *EnterpriseContactAddResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseContactAddResponse) SetBody(v *EnterpriseContactAddResponseBody) *EnterpriseContactAddResponse {
  s.Body = v
  return s
}

func (s *EnterpriseContactAddResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

