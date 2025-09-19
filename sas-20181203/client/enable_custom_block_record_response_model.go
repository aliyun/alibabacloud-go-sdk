// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomBlockRecordResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCustomBlockRecordResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCustomBlockRecordResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCustomBlockRecordResponseBody) *EnableCustomBlockRecordResponse
  GetBody() *EnableCustomBlockRecordResponseBody 
}

type EnableCustomBlockRecordResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCustomBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCustomBlockRecordResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomBlockRecordResponse) GoString() string {
  return s.String()
}

func (s *EnableCustomBlockRecordResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCustomBlockRecordResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCustomBlockRecordResponse) GetBody() *EnableCustomBlockRecordResponseBody  {
  return s.Body
}

func (s *EnableCustomBlockRecordResponse) SetHeaders(v map[string]*string) *EnableCustomBlockRecordResponse {
  s.Headers = v
  return s
}

func (s *EnableCustomBlockRecordResponse) SetStatusCode(v int32) *EnableCustomBlockRecordResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCustomBlockRecordResponse) SetBody(v *EnableCustomBlockRecordResponseBody) *EnableCustomBlockRecordResponse {
  s.Body = v
  return s
}

func (s *EnableCustomBlockRecordResponse) Validate() error {
  return dara.Validate(s)
}

