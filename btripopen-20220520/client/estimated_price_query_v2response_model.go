// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedPriceQueryV2Response interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EstimatedPriceQueryV2Response
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EstimatedPriceQueryV2Response
  GetStatusCode() *int32 
  SetBody(v *EstimatedPriceQueryV2ResponseBody) *EstimatedPriceQueryV2Response
  GetBody() *EstimatedPriceQueryV2ResponseBody 
}

type EstimatedPriceQueryV2Response struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EstimatedPriceQueryV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EstimatedPriceQueryV2Response) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryV2Response) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryV2Response) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EstimatedPriceQueryV2Response) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EstimatedPriceQueryV2Response) GetBody() *EstimatedPriceQueryV2ResponseBody  {
  return s.Body
}

func (s *EstimatedPriceQueryV2Response) SetHeaders(v map[string]*string) *EstimatedPriceQueryV2Response {
  s.Headers = v
  return s
}

func (s *EstimatedPriceQueryV2Response) SetStatusCode(v int32) *EstimatedPriceQueryV2Response {
  s.StatusCode = &v
  return s
}

func (s *EstimatedPriceQueryV2Response) SetBody(v *EstimatedPriceQueryV2ResponseBody) *EstimatedPriceQueryV2Response {
  s.Body = v
  return s
}

func (s *EstimatedPriceQueryV2Response) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

