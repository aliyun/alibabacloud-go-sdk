// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEdgeClusterAddEdgeMachineResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EdgeClusterAddEdgeMachineResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EdgeClusterAddEdgeMachineResponse
  GetStatusCode() *int32 
  SetBody(v *EdgeClusterAddEdgeMachineResponseBody) *EdgeClusterAddEdgeMachineResponse
  GetBody() *EdgeClusterAddEdgeMachineResponseBody 
}

type EdgeClusterAddEdgeMachineResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EdgeClusterAddEdgeMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EdgeClusterAddEdgeMachineResponse) String() string {
  return dara.Prettify(s)
}

func (s EdgeClusterAddEdgeMachineResponse) GoString() string {
  return s.String()
}

func (s *EdgeClusterAddEdgeMachineResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EdgeClusterAddEdgeMachineResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EdgeClusterAddEdgeMachineResponse) GetBody() *EdgeClusterAddEdgeMachineResponseBody  {
  return s.Body
}

func (s *EdgeClusterAddEdgeMachineResponse) SetHeaders(v map[string]*string) *EdgeClusterAddEdgeMachineResponse {
  s.Headers = v
  return s
}

func (s *EdgeClusterAddEdgeMachineResponse) SetStatusCode(v int32) *EdgeClusterAddEdgeMachineResponse {
  s.StatusCode = &v
  return s
}

func (s *EdgeClusterAddEdgeMachineResponse) SetBody(v *EdgeClusterAddEdgeMachineResponseBody) *EdgeClusterAddEdgeMachineResponse {
  s.Body = v
  return s
}

func (s *EdgeClusterAddEdgeMachineResponse) Validate() error {
  return dara.Validate(s)
}

