// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInsightResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableInsightResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableInsightResponse
  GetStatusCode() *int32 
  SetBody(v *EnableInsightResponseBody) *EnableInsightResponse
  GetBody() *EnableInsightResponseBody 
}

type EnableInsightResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableInsightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInsightResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableInsightResponse) GoString() string {
  return s.String()
}

func (s *EnableInsightResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableInsightResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableInsightResponse) GetBody() *EnableInsightResponseBody  {
  return s.Body
}

func (s *EnableInsightResponse) SetHeaders(v map[string]*string) *EnableInsightResponse {
  s.Headers = v
  return s
}

func (s *EnableInsightResponse) SetStatusCode(v int32) *EnableInsightResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableInsightResponse) SetBody(v *EnableInsightResponseBody) *EnableInsightResponse {
  s.Body = v
  return s
}

func (s *EnableInsightResponse) Validate() error {
  return dara.Validate(s)
}

