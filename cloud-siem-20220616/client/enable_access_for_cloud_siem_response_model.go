// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAccessForCloudSiemResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAccessForCloudSiemResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAccessForCloudSiemResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAccessForCloudSiemResponseBody) *EnableAccessForCloudSiemResponse
  GetBody() *EnableAccessForCloudSiemResponseBody 
}

type EnableAccessForCloudSiemResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAccessForCloudSiemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAccessForCloudSiemResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAccessForCloudSiemResponse) GoString() string {
  return s.String()
}

func (s *EnableAccessForCloudSiemResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAccessForCloudSiemResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAccessForCloudSiemResponse) GetBody() *EnableAccessForCloudSiemResponseBody  {
  return s.Body
}

func (s *EnableAccessForCloudSiemResponse) SetHeaders(v map[string]*string) *EnableAccessForCloudSiemResponse {
  s.Headers = v
  return s
}

func (s *EnableAccessForCloudSiemResponse) SetStatusCode(v int32) *EnableAccessForCloudSiemResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAccessForCloudSiemResponse) SetBody(v *EnableAccessForCloudSiemResponseBody) *EnableAccessForCloudSiemResponse {
  s.Body = v
  return s
}

func (s *EnableAccessForCloudSiemResponse) Validate() error {
  return dara.Validate(s)
}

