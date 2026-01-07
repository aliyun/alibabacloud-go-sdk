// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateIpMaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountUpdateIpMaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountUpdateIpMaskResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountUpdateIpMaskResponseBody) *EnterpriseAccountUpdateIpMaskResponse
  GetBody() *EnterpriseAccountUpdateIpMaskResponseBody 
}

type EnterpriseAccountUpdateIpMaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountUpdateIpMaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateIpMaskResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateIpMaskResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateIpMaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountUpdateIpMaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountUpdateIpMaskResponse) GetBody() *EnterpriseAccountUpdateIpMaskResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountUpdateIpMaskResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateIpMaskResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateIpMaskResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskResponse) SetBody(v *EnterpriseAccountUpdateIpMaskResponseBody) *EnterpriseAccountUpdateIpMaskResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountUpdateIpMaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

