// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInternalSlbResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableInternalSlbResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableInternalSlbResponse
  GetStatusCode() *int32 
  SetBody(v *EnableInternalSlbResponseBody) *EnableInternalSlbResponse
  GetBody() *EnableInternalSlbResponseBody 
}

type EnableInternalSlbResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableInternalSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInternalSlbResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableInternalSlbResponse) GoString() string {
  return s.String()
}

func (s *EnableInternalSlbResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableInternalSlbResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableInternalSlbResponse) GetBody() *EnableInternalSlbResponseBody  {
  return s.Body
}

func (s *EnableInternalSlbResponse) SetHeaders(v map[string]*string) *EnableInternalSlbResponse {
  s.Headers = v
  return s
}

func (s *EnableInternalSlbResponse) SetStatusCode(v int32) *EnableInternalSlbResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableInternalSlbResponse) SetBody(v *EnableInternalSlbResponseBody) *EnableInternalSlbResponse {
  s.Body = v
  return s
}

func (s *EnableInternalSlbResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

