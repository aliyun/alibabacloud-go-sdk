// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterServerlessResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDBClusterServerlessResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDBClusterServerlessResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDBClusterServerlessResponseBody) *EnableDBClusterServerlessResponse
  GetBody() *EnableDBClusterServerlessResponseBody 
}

type EnableDBClusterServerlessResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDBClusterServerlessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDBClusterServerlessResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterServerlessResponse) GoString() string {
  return s.String()
}

func (s *EnableDBClusterServerlessResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDBClusterServerlessResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDBClusterServerlessResponse) GetBody() *EnableDBClusterServerlessResponseBody  {
  return s.Body
}

func (s *EnableDBClusterServerlessResponse) SetHeaders(v map[string]*string) *EnableDBClusterServerlessResponse {
  s.Headers = v
  return s
}

func (s *EnableDBClusterServerlessResponse) SetStatusCode(v int32) *EnableDBClusterServerlessResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDBClusterServerlessResponse) SetBody(v *EnableDBClusterServerlessResponseBody) *EnableDBClusterServerlessResponse {
  s.Body = v
  return s
}

func (s *EnableDBClusterServerlessResponse) Validate() error {
  return dara.Validate(s)
}

