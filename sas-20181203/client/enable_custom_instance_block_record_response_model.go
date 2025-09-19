// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomInstanceBlockRecordResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCustomInstanceBlockRecordResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCustomInstanceBlockRecordResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCustomInstanceBlockRecordResponseBody) *EnableCustomInstanceBlockRecordResponse
  GetBody() *EnableCustomInstanceBlockRecordResponseBody 
}

type EnableCustomInstanceBlockRecordResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCustomInstanceBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCustomInstanceBlockRecordResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomInstanceBlockRecordResponse) GoString() string {
  return s.String()
}

func (s *EnableCustomInstanceBlockRecordResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCustomInstanceBlockRecordResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCustomInstanceBlockRecordResponse) GetBody() *EnableCustomInstanceBlockRecordResponseBody  {
  return s.Body
}

func (s *EnableCustomInstanceBlockRecordResponse) SetHeaders(v map[string]*string) *EnableCustomInstanceBlockRecordResponse {
  s.Headers = v
  return s
}

func (s *EnableCustomInstanceBlockRecordResponse) SetStatusCode(v int32) *EnableCustomInstanceBlockRecordResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCustomInstanceBlockRecordResponse) SetBody(v *EnableCustomInstanceBlockRecordResponseBody) *EnableCustomInstanceBlockRecordResponse {
  s.Body = v
  return s
}

func (s *EnableCustomInstanceBlockRecordResponse) Validate() error {
  return dara.Validate(s)
}

