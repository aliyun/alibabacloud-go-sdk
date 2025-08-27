// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserQueryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExternalUserQueryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExternalUserQueryResponse
  GetStatusCode() *int32 
  SetBody(v *ExternalUserQueryResponseBody) *ExternalUserQueryResponse
  GetBody() *ExternalUserQueryResponseBody 
}

type ExternalUserQueryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExternalUserQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExternalUserQueryResponse) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserQueryResponse) GoString() string {
  return s.String()
}

func (s *ExternalUserQueryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExternalUserQueryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExternalUserQueryResponse) GetBody() *ExternalUserQueryResponseBody  {
  return s.Body
}

func (s *ExternalUserQueryResponse) SetHeaders(v map[string]*string) *ExternalUserQueryResponse {
  s.Headers = v
  return s
}

func (s *ExternalUserQueryResponse) SetStatusCode(v int32) *ExternalUserQueryResponse {
  s.StatusCode = &v
  return s
}

func (s *ExternalUserQueryResponse) SetBody(v *ExternalUserQueryResponseBody) *ExternalUserQueryResponse {
  s.Body = v
  return s
}

func (s *ExternalUserQueryResponse) Validate() error {
  return dara.Validate(s)
}

