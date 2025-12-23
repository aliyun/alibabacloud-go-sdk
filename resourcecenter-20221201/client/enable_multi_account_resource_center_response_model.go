// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableMultiAccountResourceCenterResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableMultiAccountResourceCenterResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableMultiAccountResourceCenterResponse
  GetStatusCode() *int32 
  SetBody(v *EnableMultiAccountResourceCenterResponseBody) *EnableMultiAccountResourceCenterResponse
  GetBody() *EnableMultiAccountResourceCenterResponseBody 
}

type EnableMultiAccountResourceCenterResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableMultiAccountResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableMultiAccountResourceCenterResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableMultiAccountResourceCenterResponse) GoString() string {
  return s.String()
}

func (s *EnableMultiAccountResourceCenterResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableMultiAccountResourceCenterResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableMultiAccountResourceCenterResponse) GetBody() *EnableMultiAccountResourceCenterResponseBody  {
  return s.Body
}

func (s *EnableMultiAccountResourceCenterResponse) SetHeaders(v map[string]*string) *EnableMultiAccountResourceCenterResponse {
  s.Headers = v
  return s
}

func (s *EnableMultiAccountResourceCenterResponse) SetStatusCode(v int32) *EnableMultiAccountResourceCenterResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableMultiAccountResourceCenterResponse) SetBody(v *EnableMultiAccountResourceCenterResponseBody) *EnableMultiAccountResourceCenterResponse {
  s.Body = v
  return s
}

func (s *EnableMultiAccountResourceCenterResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

