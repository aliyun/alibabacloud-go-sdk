// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateOperateRiskControlResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountUpdateOperateRiskControlResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountUpdateOperateRiskControlResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountUpdateOperateRiskControlResponseBody) *EnterpriseAccountUpdateOperateRiskControlResponse
  GetBody() *EnterpriseAccountUpdateOperateRiskControlResponseBody 
}

type EnterpriseAccountUpdateOperateRiskControlResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountUpdateOperateRiskControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateOperateRiskControlResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateOperateRiskControlResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) GetBody() *EnterpriseAccountUpdateOperateRiskControlResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateOperateRiskControlResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateOperateRiskControlResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) SetBody(v *EnterpriseAccountUpdateOperateRiskControlResponseBody) *EnterpriseAccountUpdateOperateRiskControlResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountUpdateOperateRiskControlResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

