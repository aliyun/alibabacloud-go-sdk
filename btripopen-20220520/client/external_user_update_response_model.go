// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserUpdateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExternalUserUpdateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExternalUserUpdateResponse
  GetStatusCode() *int32 
  SetBody(v *ExternalUserUpdateResponseBody) *ExternalUserUpdateResponse
  GetBody() *ExternalUserUpdateResponseBody 
}

type ExternalUserUpdateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExternalUserUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExternalUserUpdateResponse) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserUpdateResponse) GoString() string {
  return s.String()
}

func (s *ExternalUserUpdateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExternalUserUpdateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExternalUserUpdateResponse) GetBody() *ExternalUserUpdateResponseBody  {
  return s.Body
}

func (s *ExternalUserUpdateResponse) SetHeaders(v map[string]*string) *ExternalUserUpdateResponse {
  s.Headers = v
  return s
}

func (s *ExternalUserUpdateResponse) SetStatusCode(v int32) *ExternalUserUpdateResponse {
  s.StatusCode = &v
  return s
}

func (s *ExternalUserUpdateResponse) SetBody(v *ExternalUserUpdateResponseBody) *ExternalUserUpdateResponse {
  s.Body = v
  return s
}

func (s *ExternalUserUpdateResponse) Validate() error {
  return dara.Validate(s)
}

