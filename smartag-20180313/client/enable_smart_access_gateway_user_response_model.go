// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmartAccessGatewayUserResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSmartAccessGatewayUserResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSmartAccessGatewayUserResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSmartAccessGatewayUserResponseBody) *EnableSmartAccessGatewayUserResponse
  GetBody() *EnableSmartAccessGatewayUserResponseBody 
}

type EnableSmartAccessGatewayUserResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSmartAccessGatewayUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSmartAccessGatewayUserResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSmartAccessGatewayUserResponse) GoString() string {
  return s.String()
}

func (s *EnableSmartAccessGatewayUserResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSmartAccessGatewayUserResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSmartAccessGatewayUserResponse) GetBody() *EnableSmartAccessGatewayUserResponseBody  {
  return s.Body
}

func (s *EnableSmartAccessGatewayUserResponse) SetHeaders(v map[string]*string) *EnableSmartAccessGatewayUserResponse {
  s.Headers = v
  return s
}

func (s *EnableSmartAccessGatewayUserResponse) SetStatusCode(v int32) *EnableSmartAccessGatewayUserResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSmartAccessGatewayUserResponse) SetBody(v *EnableSmartAccessGatewayUserResponseBody) *EnableSmartAccessGatewayUserResponse {
  s.Body = v
  return s
}

func (s *EnableSmartAccessGatewayUserResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

