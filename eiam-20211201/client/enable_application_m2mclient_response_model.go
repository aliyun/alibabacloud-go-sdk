// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationM2MClientResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationM2MClientResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationM2MClientResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationM2MClientResponseBody) *EnableApplicationM2MClientResponse
  GetBody() *EnableApplicationM2MClientResponseBody 
}

type EnableApplicationM2MClientResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationM2MClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationM2MClientResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationM2MClientResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationM2MClientResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationM2MClientResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationM2MClientResponse) GetBody() *EnableApplicationM2MClientResponseBody  {
  return s.Body
}

func (s *EnableApplicationM2MClientResponse) SetHeaders(v map[string]*string) *EnableApplicationM2MClientResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationM2MClientResponse) SetStatusCode(v int32) *EnableApplicationM2MClientResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationM2MClientResponse) SetBody(v *EnableApplicationM2MClientResponseBody) *EnableApplicationM2MClientResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationM2MClientResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

