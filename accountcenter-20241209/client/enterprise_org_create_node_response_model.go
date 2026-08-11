// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgCreateNodeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseOrgCreateNodeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseOrgCreateNodeResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseOrgCreateNodeResponseBody) *EnterpriseOrgCreateNodeResponse
  GetBody() *EnterpriseOrgCreateNodeResponseBody 
}

type EnterpriseOrgCreateNodeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseOrgCreateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseOrgCreateNodeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgCreateNodeResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgCreateNodeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseOrgCreateNodeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseOrgCreateNodeResponse) GetBody() *EnterpriseOrgCreateNodeResponseBody  {
  return s.Body
}

func (s *EnterpriseOrgCreateNodeResponse) SetHeaders(v map[string]*string) *EnterpriseOrgCreateNodeResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseOrgCreateNodeResponse) SetStatusCode(v int32) *EnterpriseOrgCreateNodeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseOrgCreateNodeResponse) SetBody(v *EnterpriseOrgCreateNodeResponseBody) *EnterpriseOrgCreateNodeResponse {
  s.Body = v
  return s
}

func (s *EnterpriseOrgCreateNodeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

