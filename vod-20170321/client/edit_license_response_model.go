// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditLicenseResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditLicenseResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditLicenseResponse
  GetStatusCode() *int32 
  SetBody(v *EditLicenseResponseBody) *EditLicenseResponse
  GetBody() *EditLicenseResponseBody 
}

type EditLicenseResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditLicenseResponse) String() string {
  return dara.Prettify(s)
}

func (s EditLicenseResponse) GoString() string {
  return s.String()
}

func (s *EditLicenseResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditLicenseResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditLicenseResponse) GetBody() *EditLicenseResponseBody  {
  return s.Body
}

func (s *EditLicenseResponse) SetHeaders(v map[string]*string) *EditLicenseResponse {
  s.Headers = v
  return s
}

func (s *EditLicenseResponse) SetStatusCode(v int32) *EditLicenseResponse {
  s.StatusCode = &v
  return s
}

func (s *EditLicenseResponse) SetBody(v *EditLicenseResponseBody) *EditLicenseResponse {
  s.Body = v
  return s
}

func (s *EditLicenseResponse) Validate() error {
  return dara.Validate(s)
}

