// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcClassicLinkResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableVpcClassicLinkResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableVpcClassicLinkResponse
  GetStatusCode() *int32 
  SetBody(v *EnableVpcClassicLinkResponseBody) *EnableVpcClassicLinkResponse
  GetBody() *EnableVpcClassicLinkResponseBody 
}

type EnableVpcClassicLinkResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableVpcClassicLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableVpcClassicLinkResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcClassicLinkResponse) GoString() string {
  return s.String()
}

func (s *EnableVpcClassicLinkResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableVpcClassicLinkResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableVpcClassicLinkResponse) GetBody() *EnableVpcClassicLinkResponseBody  {
  return s.Body
}

func (s *EnableVpcClassicLinkResponse) SetHeaders(v map[string]*string) *EnableVpcClassicLinkResponse {
  s.Headers = v
  return s
}

func (s *EnableVpcClassicLinkResponse) SetStatusCode(v int32) *EnableVpcClassicLinkResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableVpcClassicLinkResponse) SetBody(v *EnableVpcClassicLinkResponseBody) *EnableVpcClassicLinkResponse {
  s.Body = v
  return s
}

func (s *EnableVpcClassicLinkResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

