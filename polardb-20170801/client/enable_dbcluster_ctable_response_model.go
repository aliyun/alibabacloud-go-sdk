// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterCTableResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDBClusterCTableResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDBClusterCTableResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDBClusterCTableResponseBody) *EnableDBClusterCTableResponse
  GetBody() *EnableDBClusterCTableResponseBody 
}

type EnableDBClusterCTableResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDBClusterCTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDBClusterCTableResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterCTableResponse) GoString() string {
  return s.String()
}

func (s *EnableDBClusterCTableResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDBClusterCTableResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDBClusterCTableResponse) GetBody() *EnableDBClusterCTableResponseBody  {
  return s.Body
}

func (s *EnableDBClusterCTableResponse) SetHeaders(v map[string]*string) *EnableDBClusterCTableResponse {
  s.Headers = v
  return s
}

func (s *EnableDBClusterCTableResponse) SetStatusCode(v int32) *EnableDBClusterCTableResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDBClusterCTableResponse) SetBody(v *EnableDBClusterCTableResponseBody) *EnableDBClusterCTableResponse {
  s.Body = v
  return s
}

func (s *EnableDBClusterCTableResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

