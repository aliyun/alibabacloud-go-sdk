// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSupportAccountResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSupportAccountResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSupportAccountResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSupportAccountResponseBody) *EnableSupportAccountResponse
  GetBody() *EnableSupportAccountResponseBody 
}

type EnableSupportAccountResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSupportAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSupportAccountResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSupportAccountResponse) GoString() string {
  return s.String()
}

func (s *EnableSupportAccountResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSupportAccountResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSupportAccountResponse) GetBody() *EnableSupportAccountResponseBody  {
  return s.Body
}

func (s *EnableSupportAccountResponse) SetHeaders(v map[string]*string) *EnableSupportAccountResponse {
  s.Headers = v
  return s
}

func (s *EnableSupportAccountResponse) SetStatusCode(v int32) *EnableSupportAccountResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSupportAccountResponse) SetBody(v *EnableSupportAccountResponseBody) *EnableSupportAccountResponse {
  s.Body = v
  return s
}

func (s *EnableSupportAccountResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

