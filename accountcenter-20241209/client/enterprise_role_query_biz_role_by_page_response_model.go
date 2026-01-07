// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryBizRoleByPageResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseRoleQueryBizRoleByPageResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseRoleQueryBizRoleByPageResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseRoleQueryBizRoleByPageResponseBody) *EnterpriseRoleQueryBizRoleByPageResponse
  GetBody() *EnterpriseRoleQueryBizRoleByPageResponseBody 
}

type EnterpriseRoleQueryBizRoleByPageResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseRoleQueryBizRoleByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleByPageResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleByPageResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) GetBody() *EnterpriseRoleQueryBizRoleByPageResponseBody  {
  return s.Body
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) SetHeaders(v map[string]*string) *EnterpriseRoleQueryBizRoleByPageResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) SetStatusCode(v int32) *EnterpriseRoleQueryBizRoleByPageResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) SetBody(v *EnterpriseRoleQueryBizRoleByPageResponseBody) *EnterpriseRoleQueryBizRoleByPageResponse {
  s.Body = v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

