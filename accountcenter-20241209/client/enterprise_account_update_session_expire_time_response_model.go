// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateSessionExpireTimeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountUpdateSessionExpireTimeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountUpdateSessionExpireTimeResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountUpdateSessionExpireTimeResponseBody) *EnterpriseAccountUpdateSessionExpireTimeResponse
  GetBody() *EnterpriseAccountUpdateSessionExpireTimeResponseBody 
}

type EnterpriseAccountUpdateSessionExpireTimeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountUpdateSessionExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateSessionExpireTimeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateSessionExpireTimeResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) GetBody() *EnterpriseAccountUpdateSessionExpireTimeResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateSessionExpireTimeResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateSessionExpireTimeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) SetBody(v *EnterpriseAccountUpdateSessionExpireTimeResponseBody) *EnterpriseAccountUpdateSessionExpireTimeResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountUpdateSessionExpireTimeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

