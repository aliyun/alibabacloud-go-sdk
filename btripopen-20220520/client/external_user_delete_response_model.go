// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserDeleteResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExternalUserDeleteResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExternalUserDeleteResponse
  GetStatusCode() *int32 
  SetBody(v *ExternalUserDeleteResponseBody) *ExternalUserDeleteResponse
  GetBody() *ExternalUserDeleteResponseBody 
}

type ExternalUserDeleteResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExternalUserDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExternalUserDeleteResponse) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserDeleteResponse) GoString() string {
  return s.String()
}

func (s *ExternalUserDeleteResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExternalUserDeleteResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExternalUserDeleteResponse) GetBody() *ExternalUserDeleteResponseBody  {
  return s.Body
}

func (s *ExternalUserDeleteResponse) SetHeaders(v map[string]*string) *ExternalUserDeleteResponse {
  s.Headers = v
  return s
}

func (s *ExternalUserDeleteResponse) SetStatusCode(v int32) *ExternalUserDeleteResponse {
  s.StatusCode = &v
  return s
}

func (s *ExternalUserDeleteResponse) SetBody(v *ExternalUserDeleteResponseBody) *ExternalUserDeleteResponse {
  s.Body = v
  return s
}

func (s *ExternalUserDeleteResponse) Validate() error {
  return dara.Validate(s)
}

