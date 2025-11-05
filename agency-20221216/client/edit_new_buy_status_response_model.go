// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditNewBuyStatusResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditNewBuyStatusResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditNewBuyStatusResponse
  GetStatusCode() *int32 
  SetBody(v *EditNewBuyStatusResponseBody) *EditNewBuyStatusResponse
  GetBody() *EditNewBuyStatusResponseBody 
}

type EditNewBuyStatusResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditNewBuyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditNewBuyStatusResponse) String() string {
  return dara.Prettify(s)
}

func (s EditNewBuyStatusResponse) GoString() string {
  return s.String()
}

func (s *EditNewBuyStatusResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditNewBuyStatusResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditNewBuyStatusResponse) GetBody() *EditNewBuyStatusResponseBody  {
  return s.Body
}

func (s *EditNewBuyStatusResponse) SetHeaders(v map[string]*string) *EditNewBuyStatusResponse {
  s.Headers = v
  return s
}

func (s *EditNewBuyStatusResponse) SetStatusCode(v int32) *EditNewBuyStatusResponse {
  s.StatusCode = &v
  return s
}

func (s *EditNewBuyStatusResponse) SetBody(v *EditNewBuyStatusResponseBody) *EditNewBuyStatusResponse {
  s.Body = v
  return s
}

func (s *EditNewBuyStatusResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

