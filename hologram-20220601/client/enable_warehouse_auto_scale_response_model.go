// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWarehouseAutoScaleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableWarehouseAutoScaleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableWarehouseAutoScaleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableWarehouseAutoScaleResponseBody) *EnableWarehouseAutoScaleResponse
  GetBody() *EnableWarehouseAutoScaleResponseBody 
}

type EnableWarehouseAutoScaleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableWarehouseAutoScaleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableWarehouseAutoScaleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableWarehouseAutoScaleResponse) GoString() string {
  return s.String()
}

func (s *EnableWarehouseAutoScaleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableWarehouseAutoScaleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableWarehouseAutoScaleResponse) GetBody() *EnableWarehouseAutoScaleResponseBody  {
  return s.Body
}

func (s *EnableWarehouseAutoScaleResponse) SetHeaders(v map[string]*string) *EnableWarehouseAutoScaleResponse {
  s.Headers = v
  return s
}

func (s *EnableWarehouseAutoScaleResponse) SetStatusCode(v int32) *EnableWarehouseAutoScaleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableWarehouseAutoScaleResponse) SetBody(v *EnableWarehouseAutoScaleResponseBody) *EnableWarehouseAutoScaleResponse {
  s.Body = v
  return s
}

func (s *EnableWarehouseAutoScaleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

