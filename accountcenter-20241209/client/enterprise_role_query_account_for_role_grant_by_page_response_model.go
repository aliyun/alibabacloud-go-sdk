// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryAccountForRoleGrantByPageResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse
  GetBody() *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody 
}

type EnterpriseRoleQueryAccountForRoleGrantByPageResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) GetBody() *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody  {
  return s.Body
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) SetHeaders(v map[string]*string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) SetStatusCode(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) SetBody(v *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) *EnterpriseRoleQueryAccountForRoleGrantByPageResponse {
  s.Body = v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

