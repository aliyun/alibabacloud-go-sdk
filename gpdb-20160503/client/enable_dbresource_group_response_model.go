// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBResourceGroupResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDBResourceGroupResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDBResourceGroupResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDBResourceGroupResponseBody) *EnableDBResourceGroupResponse
  GetBody() *EnableDBResourceGroupResponseBody 
}

type EnableDBResourceGroupResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDBResourceGroupResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDBResourceGroupResponse) GoString() string {
  return s.String()
}

func (s *EnableDBResourceGroupResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDBResourceGroupResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDBResourceGroupResponse) GetBody() *EnableDBResourceGroupResponseBody  {
  return s.Body
}

func (s *EnableDBResourceGroupResponse) SetHeaders(v map[string]*string) *EnableDBResourceGroupResponse {
  s.Headers = v
  return s
}

func (s *EnableDBResourceGroupResponse) SetStatusCode(v int32) *EnableDBResourceGroupResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDBResourceGroupResponse) SetBody(v *EnableDBResourceGroupResponseBody) *EnableDBResourceGroupResponse {
  s.Body = v
  return s
}

func (s *EnableDBResourceGroupResponse) Validate() error {
  return dara.Validate(s)
}

