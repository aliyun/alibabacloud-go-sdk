// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseOrgRenameNodeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseOrgRenameNodeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseOrgRenameNodeResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseOrgRenameNodeResponseBody) *EnterpriseOrgRenameNodeResponse
  GetBody() *EnterpriseOrgRenameNodeResponseBody 
}

type EnterpriseOrgRenameNodeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseOrgRenameNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseOrgRenameNodeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseOrgRenameNodeResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseOrgRenameNodeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseOrgRenameNodeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseOrgRenameNodeResponse) GetBody() *EnterpriseOrgRenameNodeResponseBody  {
  return s.Body
}

func (s *EnterpriseOrgRenameNodeResponse) SetHeaders(v map[string]*string) *EnterpriseOrgRenameNodeResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseOrgRenameNodeResponse) SetStatusCode(v int32) *EnterpriseOrgRenameNodeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseOrgRenameNodeResponse) SetBody(v *EnterpriseOrgRenameNodeResponseBody) *EnterpriseOrgRenameNodeResponse {
  s.Body = v
  return s
}

func (s *EnterpriseOrgRenameNodeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

