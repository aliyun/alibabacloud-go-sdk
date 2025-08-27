// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedPriceQueryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EstimatedPriceQueryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EstimatedPriceQueryResponse
  GetStatusCode() *int32 
  SetBody(v *EstimatedPriceQueryResponseBody) *EstimatedPriceQueryResponse
  GetBody() *EstimatedPriceQueryResponseBody 
}

type EstimatedPriceQueryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EstimatedPriceQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EstimatedPriceQueryResponse) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryResponse) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EstimatedPriceQueryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EstimatedPriceQueryResponse) GetBody() *EstimatedPriceQueryResponseBody  {
  return s.Body
}

func (s *EstimatedPriceQueryResponse) SetHeaders(v map[string]*string) *EstimatedPriceQueryResponse {
  s.Headers = v
  return s
}

func (s *EstimatedPriceQueryResponse) SetStatusCode(v int32) *EstimatedPriceQueryResponse {
  s.StatusCode = &v
  return s
}

func (s *EstimatedPriceQueryResponse) SetBody(v *EstimatedPriceQueryResponseBody) *EstimatedPriceQueryResponse {
  s.Body = v
  return s
}

func (s *EstimatedPriceQueryResponse) Validate() error {
  return dara.Validate(s)
}

