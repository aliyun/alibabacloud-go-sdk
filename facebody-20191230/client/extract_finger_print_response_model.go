// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtractFingerPrintResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExtractFingerPrintResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExtractFingerPrintResponse
  GetStatusCode() *int32 
  SetBody(v *ExtractFingerPrintResponseBody) *ExtractFingerPrintResponse
  GetBody() *ExtractFingerPrintResponseBody 
}

type ExtractFingerPrintResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExtractFingerPrintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExtractFingerPrintResponse) String() string {
  return dara.Prettify(s)
}

func (s ExtractFingerPrintResponse) GoString() string {
  return s.String()
}

func (s *ExtractFingerPrintResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExtractFingerPrintResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExtractFingerPrintResponse) GetBody() *ExtractFingerPrintResponseBody  {
  return s.Body
}

func (s *ExtractFingerPrintResponse) SetHeaders(v map[string]*string) *ExtractFingerPrintResponse {
  s.Headers = v
  return s
}

func (s *ExtractFingerPrintResponse) SetStatusCode(v int32) *ExtractFingerPrintResponse {
  s.StatusCode = &v
  return s
}

func (s *ExtractFingerPrintResponse) SetBody(v *ExtractFingerPrintResponseBody) *ExtractFingerPrintResponse {
  s.Body = v
  return s
}

func (s *ExtractFingerPrintResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

