// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExchangeEntitlementResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExchangeEntitlementResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExchangeEntitlementResponse
  GetStatusCode() *int32 
  SetBody(v *ExchangeEntitlementResponseBody) *ExchangeEntitlementResponse
  GetBody() *ExchangeEntitlementResponseBody 
}

type ExchangeEntitlementResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExchangeEntitlementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExchangeEntitlementResponse) String() string {
  return dara.Prettify(s)
}

func (s ExchangeEntitlementResponse) GoString() string {
  return s.String()
}

func (s *ExchangeEntitlementResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExchangeEntitlementResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExchangeEntitlementResponse) GetBody() *ExchangeEntitlementResponseBody  {
  return s.Body
}

func (s *ExchangeEntitlementResponse) SetHeaders(v map[string]*string) *ExchangeEntitlementResponse {
  s.Headers = v
  return s
}

func (s *ExchangeEntitlementResponse) SetStatusCode(v int32) *ExchangeEntitlementResponse {
  s.StatusCode = &v
  return s
}

func (s *ExchangeEntitlementResponse) SetBody(v *ExchangeEntitlementResponseBody) *ExchangeEntitlementResponse {
  s.Body = v
  return s
}

func (s *ExchangeEntitlementResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

