// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTriggerNodeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTriggerNodeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTriggerNodeResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTriggerNodeResponseBody) *ExecuteTriggerNodeResponse
  GetBody() *ExecuteTriggerNodeResponseBody 
}

type ExecuteTriggerNodeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTriggerNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTriggerNodeResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTriggerNodeResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTriggerNodeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTriggerNodeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTriggerNodeResponse) GetBody() *ExecuteTriggerNodeResponseBody  {
  return s.Body
}

func (s *ExecuteTriggerNodeResponse) SetHeaders(v map[string]*string) *ExecuteTriggerNodeResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTriggerNodeResponse) SetStatusCode(v int32) *ExecuteTriggerNodeResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTriggerNodeResponse) SetBody(v *ExecuteTriggerNodeResponseBody) *ExecuteTriggerNodeResponse {
  s.Body = v
  return s
}

func (s *ExecuteTriggerNodeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

