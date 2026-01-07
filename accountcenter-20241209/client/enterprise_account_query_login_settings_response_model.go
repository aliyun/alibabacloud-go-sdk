// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryLoginSettingsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountQueryLoginSettingsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountQueryLoginSettingsResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountQueryLoginSettingsResponseBody) *EnterpriseAccountQueryLoginSettingsResponse
  GetBody() *EnterpriseAccountQueryLoginSettingsResponseBody 
}

type EnterpriseAccountQueryLoginSettingsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountQueryLoginSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountQueryLoginSettingsResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryLoginSettingsResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) GetBody() *EnterpriseAccountQueryLoginSettingsResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) SetHeaders(v map[string]*string) *EnterpriseAccountQueryLoginSettingsResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) SetStatusCode(v int32) *EnterpriseAccountQueryLoginSettingsResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) SetBody(v *EnterpriseAccountQueryLoginSettingsResponseBody) *EnterpriseAccountQueryLoginSettingsResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountQueryLoginSettingsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

