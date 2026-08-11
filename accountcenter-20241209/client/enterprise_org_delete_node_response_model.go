// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgDeleteNodeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseOrgDeleteNodeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseOrgDeleteNodeResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseOrgDeleteNodeResponseBody) *EnterpriseOrgDeleteNodeResponse
  GetBody() *EnterpriseOrgDeleteNodeResponseBody 
}

type EnterpriseOrgDeleteNodeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseOrgDeleteNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseOrgDeleteNodeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgDeleteNodeResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgDeleteNodeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseOrgDeleteNodeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseOrgDeleteNodeResponse) GetBody() *EnterpriseOrgDeleteNodeResponseBody  {
  return s.Body
}

func (s *EnterpriseOrgDeleteNodeResponse) SetHeaders(v map[string]*string) *EnterpriseOrgDeleteNodeResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseOrgDeleteNodeResponse) SetStatusCode(v int32) *EnterpriseOrgDeleteNodeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseOrgDeleteNodeResponse) SetBody(v *EnterpriseOrgDeleteNodeResponseBody) *EnterpriseOrgDeleteNodeResponse {
  s.Body = v
  return s
}

func (s *EnterpriseOrgDeleteNodeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

