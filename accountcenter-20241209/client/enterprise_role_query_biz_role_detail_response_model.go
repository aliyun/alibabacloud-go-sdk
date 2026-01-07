// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryBizRoleDetailResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseRoleQueryBizRoleDetailResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseRoleQueryBizRoleDetailResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseRoleQueryBizRoleDetailResponseBody) *EnterpriseRoleQueryBizRoleDetailResponse
  GetBody() *EnterpriseRoleQueryBizRoleDetailResponseBody 
}

type EnterpriseRoleQueryBizRoleDetailResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseRoleQueryBizRoleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) GetBody() *EnterpriseRoleQueryBizRoleDetailResponseBody  {
  return s.Body
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) SetHeaders(v map[string]*string) *EnterpriseRoleQueryBizRoleDetailResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) SetStatusCode(v int32) *EnterpriseRoleQueryBizRoleDetailResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) SetBody(v *EnterpriseRoleQueryBizRoleDetailResponseBody) *EnterpriseRoleQueryBizRoleDetailResponse {
  s.Body = v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

