// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserAddResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExternalUserAddResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExternalUserAddResponse
  GetStatusCode() *int32 
  SetBody(v *ExternalUserAddResponseBody) *ExternalUserAddResponse
  GetBody() *ExternalUserAddResponseBody 
}

type ExternalUserAddResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExternalUserAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExternalUserAddResponse) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserAddResponse) GoString() string {
  return s.String()
}

func (s *ExternalUserAddResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExternalUserAddResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExternalUserAddResponse) GetBody() *ExternalUserAddResponseBody  {
  return s.Body
}

func (s *ExternalUserAddResponse) SetHeaders(v map[string]*string) *ExternalUserAddResponse {
  s.Headers = v
  return s
}

func (s *ExternalUserAddResponse) SetStatusCode(v int32) *ExternalUserAddResponse {
  s.StatusCode = &v
  return s
}

func (s *ExternalUserAddResponse) SetBody(v *ExternalUserAddResponseBody) *ExternalUserAddResponse {
  s.Body = v
  return s
}

func (s *ExternalUserAddResponse) Validate() error {
  return dara.Validate(s)
}

