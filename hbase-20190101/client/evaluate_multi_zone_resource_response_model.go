// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateMultiZoneResourceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EvaluateMultiZoneResourceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EvaluateMultiZoneResourceResponse
  GetStatusCode() *int32 
  SetBody(v *EvaluateMultiZoneResourceResponseBody) *EvaluateMultiZoneResourceResponse
  GetBody() *EvaluateMultiZoneResourceResponseBody 
}

type EvaluateMultiZoneResourceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EvaluateMultiZoneResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateMultiZoneResourceResponse) String() string {
  return dara.Prettify(s)
}

func (s EvaluateMultiZoneResourceResponse) GoString() string {
  return s.String()
}

func (s *EvaluateMultiZoneResourceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EvaluateMultiZoneResourceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EvaluateMultiZoneResourceResponse) GetBody() *EvaluateMultiZoneResourceResponseBody  {
  return s.Body
}

func (s *EvaluateMultiZoneResourceResponse) SetHeaders(v map[string]*string) *EvaluateMultiZoneResourceResponse {
  s.Headers = v
  return s
}

func (s *EvaluateMultiZoneResourceResponse) SetStatusCode(v int32) *EvaluateMultiZoneResourceResponse {
  s.StatusCode = &v
  return s
}

func (s *EvaluateMultiZoneResourceResponse) SetBody(v *EvaluateMultiZoneResourceResponseBody) *EvaluateMultiZoneResourceResponse {
  s.Body = v
  return s
}

func (s *EvaluateMultiZoneResourceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

