// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInitDomainAutoRedirectResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableInitDomainAutoRedirectResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableInitDomainAutoRedirectResponse
  GetStatusCode() *int32 
  SetBody(v *EnableInitDomainAutoRedirectResponseBody) *EnableInitDomainAutoRedirectResponse
  GetBody() *EnableInitDomainAutoRedirectResponseBody 
}

type EnableInitDomainAutoRedirectResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableInitDomainAutoRedirectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInitDomainAutoRedirectResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableInitDomainAutoRedirectResponse) GoString() string {
  return s.String()
}

func (s *EnableInitDomainAutoRedirectResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableInitDomainAutoRedirectResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableInitDomainAutoRedirectResponse) GetBody() *EnableInitDomainAutoRedirectResponseBody  {
  return s.Body
}

func (s *EnableInitDomainAutoRedirectResponse) SetHeaders(v map[string]*string) *EnableInitDomainAutoRedirectResponse {
  s.Headers = v
  return s
}

func (s *EnableInitDomainAutoRedirectResponse) SetStatusCode(v int32) *EnableInitDomainAutoRedirectResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableInitDomainAutoRedirectResponse) SetBody(v *EnableInitDomainAutoRedirectResponseBody) *EnableInitDomainAutoRedirectResponse {
  s.Body = v
  return s
}

func (s *EnableInitDomainAutoRedirectResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

