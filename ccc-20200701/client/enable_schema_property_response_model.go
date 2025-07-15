// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSchemaPropertyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSchemaPropertyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSchemaPropertyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSchemaPropertyResponseBody) *EnableSchemaPropertyResponse
  GetBody() *EnableSchemaPropertyResponseBody 
}

type EnableSchemaPropertyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSchemaPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSchemaPropertyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSchemaPropertyResponse) GoString() string {
  return s.String()
}

func (s *EnableSchemaPropertyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSchemaPropertyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSchemaPropertyResponse) GetBody() *EnableSchemaPropertyResponseBody  {
  return s.Body
}

func (s *EnableSchemaPropertyResponse) SetHeaders(v map[string]*string) *EnableSchemaPropertyResponse {
  s.Headers = v
  return s
}

func (s *EnableSchemaPropertyResponse) SetStatusCode(v int32) *EnableSchemaPropertyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSchemaPropertyResponse) SetBody(v *EnableSchemaPropertyResponseBody) *EnableSchemaPropertyResponse {
  s.Body = v
  return s
}

func (s *EnableSchemaPropertyResponse) Validate() error {
  return dara.Validate(s)
}

