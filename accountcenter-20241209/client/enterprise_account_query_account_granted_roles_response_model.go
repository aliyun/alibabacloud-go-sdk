// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryAccountGrantedRolesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountQueryAccountGrantedRolesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountQueryAccountGrantedRolesResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountQueryAccountGrantedRolesResponseBody) *EnterpriseAccountQueryAccountGrantedRolesResponse
  GetBody() *EnterpriseAccountQueryAccountGrantedRolesResponseBody 
}

type EnterpriseAccountQueryAccountGrantedRolesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountQueryAccountGrantedRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) GetBody() *EnterpriseAccountQueryAccountGrantedRolesResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) SetHeaders(v map[string]*string) *EnterpriseAccountQueryAccountGrantedRolesResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) SetStatusCode(v int32) *EnterpriseAccountQueryAccountGrantedRolesResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) SetBody(v *EnterpriseAccountQueryAccountGrantedRolesResponseBody) *EnterpriseAccountQueryAccountGrantedRolesResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

