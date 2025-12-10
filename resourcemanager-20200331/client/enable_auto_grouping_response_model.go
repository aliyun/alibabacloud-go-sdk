// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoGroupingResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAutoGroupingResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAutoGroupingResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAutoGroupingResponseBody) *EnableAutoGroupingResponse
  GetBody() *EnableAutoGroupingResponseBody 
}

type EnableAutoGroupingResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAutoGroupingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAutoGroupingResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoGroupingResponse) GoString() string {
  return s.String()
}

func (s *EnableAutoGroupingResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAutoGroupingResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAutoGroupingResponse) GetBody() *EnableAutoGroupingResponseBody  {
  return s.Body
}

func (s *EnableAutoGroupingResponse) SetHeaders(v map[string]*string) *EnableAutoGroupingResponse {
  s.Headers = v
  return s
}

func (s *EnableAutoGroupingResponse) SetStatusCode(v int32) *EnableAutoGroupingResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAutoGroupingResponse) SetBody(v *EnableAutoGroupingResponseBody) *EnableAutoGroupingResponse {
  s.Body = v
  return s
}

func (s *EnableAutoGroupingResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

