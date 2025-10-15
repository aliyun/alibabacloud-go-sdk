// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationApiInvokeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationApiInvokeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationApiInvokeResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationApiInvokeResponseBody) *EnableApplicationApiInvokeResponse
  GetBody() *EnableApplicationApiInvokeResponseBody 
}

type EnableApplicationApiInvokeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationApiInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationApiInvokeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationApiInvokeResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationApiInvokeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationApiInvokeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationApiInvokeResponse) GetBody() *EnableApplicationApiInvokeResponseBody  {
  return s.Body
}

func (s *EnableApplicationApiInvokeResponse) SetHeaders(v map[string]*string) *EnableApplicationApiInvokeResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationApiInvokeResponse) SetStatusCode(v int32) *EnableApplicationApiInvokeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationApiInvokeResponse) SetBody(v *EnableApplicationApiInvokeResponseBody) *EnableApplicationApiInvokeResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationApiInvokeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

