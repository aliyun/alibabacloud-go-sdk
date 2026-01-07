// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryAccountsInfoResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountQueryAccountsInfoResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountQueryAccountsInfoResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountQueryAccountsInfoResponseBody) *EnterpriseAccountQueryAccountsInfoResponse
  GetBody() *EnterpriseAccountQueryAccountsInfoResponseBody 
}

type EnterpriseAccountQueryAccountsInfoResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountQueryAccountsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountQueryAccountsInfoResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryAccountsInfoResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) GetBody() *EnterpriseAccountQueryAccountsInfoResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) SetHeaders(v map[string]*string) *EnterpriseAccountQueryAccountsInfoResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) SetStatusCode(v int32) *EnterpriseAccountQueryAccountsInfoResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) SetBody(v *EnterpriseAccountQueryAccountsInfoResponseBody) *EnterpriseAccountQueryAccountsInfoResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

