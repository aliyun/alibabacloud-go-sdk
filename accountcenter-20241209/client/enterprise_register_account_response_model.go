// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRegisterAccountResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseRegisterAccountResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseRegisterAccountResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseRegisterAccountResponseBody) *EnterpriseRegisterAccountResponse
  GetBody() *EnterpriseRegisterAccountResponseBody 
}

type EnterpriseRegisterAccountResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseRegisterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRegisterAccountResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRegisterAccountResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseRegisterAccountResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseRegisterAccountResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseRegisterAccountResponse) GetBody() *EnterpriseRegisterAccountResponseBody  {
  return s.Body
}

func (s *EnterpriseRegisterAccountResponse) SetHeaders(v map[string]*string) *EnterpriseRegisterAccountResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseRegisterAccountResponse) SetStatusCode(v int32) *EnterpriseRegisterAccountResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseRegisterAccountResponse) SetBody(v *EnterpriseRegisterAccountResponseBody) *EnterpriseRegisterAccountResponse {
  s.Body = v
  return s
}

func (s *EnterpriseRegisterAccountResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

