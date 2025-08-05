// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePolicyV2Response interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecutePolicyV2Response
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecutePolicyV2Response
  GetStatusCode() *int32 
  SetBody(v *ExecutePolicyV2ResponseBody) *ExecutePolicyV2Response
  GetBody() *ExecutePolicyV2ResponseBody 
}

type ExecutePolicyV2Response struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecutePolicyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecutePolicyV2Response) String() string {
  return dara.Prettify(s)
}

func (s ExecutePolicyV2Response) GoString() string {
  return s.String()
}

func (s *ExecutePolicyV2Response) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecutePolicyV2Response) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecutePolicyV2Response) GetBody() *ExecutePolicyV2ResponseBody  {
  return s.Body
}

func (s *ExecutePolicyV2Response) SetHeaders(v map[string]*string) *ExecutePolicyV2Response {
  s.Headers = v
  return s
}

func (s *ExecutePolicyV2Response) SetStatusCode(v int32) *ExecutePolicyV2Response {
  s.StatusCode = &v
  return s
}

func (s *ExecutePolicyV2Response) SetBody(v *ExecutePolicyV2ResponseBody) *ExecutePolicyV2Response {
  s.Body = v
  return s
}

func (s *ExecutePolicyV2Response) Validate() error {
  return dara.Validate(s)
}

