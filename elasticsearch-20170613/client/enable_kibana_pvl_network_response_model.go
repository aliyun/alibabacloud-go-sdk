// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableKibanaPvlNetworkResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableKibanaPvlNetworkResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableKibanaPvlNetworkResponse
  GetStatusCode() *int32 
  SetBody(v *EnableKibanaPvlNetworkResponseBody) *EnableKibanaPvlNetworkResponse
  GetBody() *EnableKibanaPvlNetworkResponseBody 
}

type EnableKibanaPvlNetworkResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableKibanaPvlNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableKibanaPvlNetworkResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableKibanaPvlNetworkResponse) GoString() string {
  return s.String()
}

func (s *EnableKibanaPvlNetworkResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableKibanaPvlNetworkResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableKibanaPvlNetworkResponse) GetBody() *EnableKibanaPvlNetworkResponseBody  {
  return s.Body
}

func (s *EnableKibanaPvlNetworkResponse) SetHeaders(v map[string]*string) *EnableKibanaPvlNetworkResponse {
  s.Headers = v
  return s
}

func (s *EnableKibanaPvlNetworkResponse) SetStatusCode(v int32) *EnableKibanaPvlNetworkResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableKibanaPvlNetworkResponse) SetBody(v *EnableKibanaPvlNetworkResponseBody) *EnableKibanaPvlNetworkResponse {
  s.Body = v
  return s
}

func (s *EnableKibanaPvlNetworkResponse) Validate() error {
  return dara.Validate(s)
}

