// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHttp2Response interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableHttp2Response
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableHttp2Response
  GetStatusCode() *int32 
  SetBody(v *EnableHttp2ResponseBody) *EnableHttp2Response
  GetBody() *EnableHttp2ResponseBody 
}

type EnableHttp2Response struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableHttp2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHttp2Response) String() string {
  return dara.Prettify(s)
}

func (s EnableHttp2Response) GoString() string {
  return s.String()
}

func (s *EnableHttp2Response) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableHttp2Response) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableHttp2Response) GetBody() *EnableHttp2ResponseBody  {
  return s.Body
}

func (s *EnableHttp2Response) SetHeaders(v map[string]*string) *EnableHttp2Response {
  s.Headers = v
  return s
}

func (s *EnableHttp2Response) SetStatusCode(v int32) *EnableHttp2Response {
  s.StatusCode = &v
  return s
}

func (s *EnableHttp2Response) SetBody(v *EnableHttp2ResponseBody) *EnableHttp2Response {
  s.Body = v
  return s
}

func (s *EnableHttp2Response) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

