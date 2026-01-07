// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactEditResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseContactEditResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseContactEditResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseContactEditResponseBody) *EnterpriseContactEditResponse
  GetBody() *EnterpriseContactEditResponseBody 
}

type EnterpriseContactEditResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseContactEditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseContactEditResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactEditResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseContactEditResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseContactEditResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseContactEditResponse) GetBody() *EnterpriseContactEditResponseBody  {
  return s.Body
}

func (s *EnterpriseContactEditResponse) SetHeaders(v map[string]*string) *EnterpriseContactEditResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseContactEditResponse) SetStatusCode(v int32) *EnterpriseContactEditResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseContactEditResponse) SetBody(v *EnterpriseContactEditResponseBody) *EnterpriseContactEditResponse {
  s.Body = v
  return s
}

func (s *EnterpriseContactEditResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

