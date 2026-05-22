// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApiKeyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApiKeyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApiKeyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApiKeyResponseBody) *EnableApiKeyResponse
  GetBody() *EnableApiKeyResponseBody 
}

type EnableApiKeyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApiKeyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApiKeyResponse) GoString() string {
  return s.String()
}

func (s *EnableApiKeyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApiKeyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApiKeyResponse) GetBody() *EnableApiKeyResponseBody  {
  return s.Body
}

func (s *EnableApiKeyResponse) SetHeaders(v map[string]*string) *EnableApiKeyResponse {
  s.Headers = v
  return s
}

func (s *EnableApiKeyResponse) SetStatusCode(v int32) *EnableApiKeyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApiKeyResponse) SetBody(v *EnableApiKeyResponseBody) *EnableApiKeyResponse {
  s.Body = v
  return s
}

func (s *EnableApiKeyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

