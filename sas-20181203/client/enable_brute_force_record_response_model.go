// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBruteForceRecordResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableBruteForceRecordResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableBruteForceRecordResponse
  GetStatusCode() *int32 
  SetBody(v *EnableBruteForceRecordResponseBody) *EnableBruteForceRecordResponse
  GetBody() *EnableBruteForceRecordResponseBody 
}

type EnableBruteForceRecordResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableBruteForceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableBruteForceRecordResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableBruteForceRecordResponse) GoString() string {
  return s.String()
}

func (s *EnableBruteForceRecordResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableBruteForceRecordResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableBruteForceRecordResponse) GetBody() *EnableBruteForceRecordResponseBody  {
  return s.Body
}

func (s *EnableBruteForceRecordResponse) SetHeaders(v map[string]*string) *EnableBruteForceRecordResponse {
  s.Headers = v
  return s
}

func (s *EnableBruteForceRecordResponse) SetStatusCode(v int32) *EnableBruteForceRecordResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableBruteForceRecordResponse) SetBody(v *EnableBruteForceRecordResponseBody) *EnableBruteForceRecordResponse {
  s.Body = v
  return s
}

func (s *EnableBruteForceRecordResponse) Validate() error {
  return dara.Validate(s)
}

