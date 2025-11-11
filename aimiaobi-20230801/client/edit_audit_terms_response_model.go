// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditAuditTermsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditAuditTermsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditAuditTermsResponse
  GetStatusCode() *int32 
  SetBody(v *EditAuditTermsResponseBody) *EditAuditTermsResponse
  GetBody() *EditAuditTermsResponseBody 
}

type EditAuditTermsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditAuditTermsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditAuditTermsResponse) String() string {
  return dara.Prettify(s)
}

func (s EditAuditTermsResponse) GoString() string {
  return s.String()
}

func (s *EditAuditTermsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditAuditTermsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditAuditTermsResponse) GetBody() *EditAuditTermsResponseBody  {
  return s.Body
}

func (s *EditAuditTermsResponse) SetHeaders(v map[string]*string) *EditAuditTermsResponse {
  s.Headers = v
  return s
}

func (s *EditAuditTermsResponse) SetStatusCode(v int32) *EditAuditTermsResponse {
  s.StatusCode = &v
  return s
}

func (s *EditAuditTermsResponse) SetBody(v *EditAuditTermsResponseBody) *EditAuditTermsResponse {
  s.Body = v
  return s
}

func (s *EditAuditTermsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

