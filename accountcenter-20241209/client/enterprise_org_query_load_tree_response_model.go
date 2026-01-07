// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgQueryLoadTreeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseOrgQueryLoadTreeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseOrgQueryLoadTreeResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseOrgQueryLoadTreeResponseBody) *EnterpriseOrgQueryLoadTreeResponse
  GetBody() *EnterpriseOrgQueryLoadTreeResponseBody 
}

type EnterpriseOrgQueryLoadTreeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseOrgQueryLoadTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseOrgQueryLoadTreeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgQueryLoadTreeResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgQueryLoadTreeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseOrgQueryLoadTreeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseOrgQueryLoadTreeResponse) GetBody() *EnterpriseOrgQueryLoadTreeResponseBody  {
  return s.Body
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetHeaders(v map[string]*string) *EnterpriseOrgQueryLoadTreeResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetStatusCode(v int32) *EnterpriseOrgQueryLoadTreeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeResponse) SetBody(v *EnterpriseOrgQueryLoadTreeResponseBody) *EnterpriseOrgQueryLoadTreeResponse {
  s.Body = v
  return s
}

func (s *EnterpriseOrgQueryLoadTreeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

