// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEkycVerifyV2Response interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EkycVerifyV2Response
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EkycVerifyV2Response
  GetStatusCode() *int32 
  SetBody(v *EkycVerifyV2ResponseBody) *EkycVerifyV2Response
  GetBody() *EkycVerifyV2ResponseBody 
}

type EkycVerifyV2Response struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EkycVerifyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EkycVerifyV2Response) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyV2Response) GoString() string {
  return s.String()
}

func (s *EkycVerifyV2Response) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EkycVerifyV2Response) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EkycVerifyV2Response) GetBody() *EkycVerifyV2ResponseBody  {
  return s.Body
}

func (s *EkycVerifyV2Response) SetHeaders(v map[string]*string) *EkycVerifyV2Response {
  s.Headers = v
  return s
}

func (s *EkycVerifyV2Response) SetStatusCode(v int32) *EkycVerifyV2Response {
  s.StatusCode = &v
  return s
}

func (s *EkycVerifyV2Response) SetBody(v *EkycVerifyV2ResponseBody) *EkycVerifyV2Response {
  s.Body = v
  return s
}

func (s *EkycVerifyV2Response) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

