// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendImageStyleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExtendImageStyleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExtendImageStyleResponse
  GetStatusCode() *int32 
  SetBody(v *ExtendImageStyleResponseBody) *ExtendImageStyleResponse
  GetBody() *ExtendImageStyleResponseBody 
}

type ExtendImageStyleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExtendImageStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExtendImageStyleResponse) String() string {
  return dara.Prettify(s)
}

func (s ExtendImageStyleResponse) GoString() string {
  return s.String()
}

func (s *ExtendImageStyleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExtendImageStyleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExtendImageStyleResponse) GetBody() *ExtendImageStyleResponseBody  {
  return s.Body
}

func (s *ExtendImageStyleResponse) SetHeaders(v map[string]*string) *ExtendImageStyleResponse {
  s.Headers = v
  return s
}

func (s *ExtendImageStyleResponse) SetStatusCode(v int32) *ExtendImageStyleResponse {
  s.StatusCode = &v
  return s
}

func (s *ExtendImageStyleResponse) SetBody(v *ExtendImageStyleResponseBody) *ExtendImageStyleResponse {
  s.Body = v
  return s
}

func (s *ExtendImageStyleResponse) Validate() error {
  return dara.Validate(s)
}

