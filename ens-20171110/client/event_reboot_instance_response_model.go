// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventRebootInstanceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EventRebootInstanceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EventRebootInstanceResponse
  GetStatusCode() *int32 
  SetBody(v *EventRebootInstanceResponseBody) *EventRebootInstanceResponse
  GetBody() *EventRebootInstanceResponseBody 
}

type EventRebootInstanceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EventRebootInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EventRebootInstanceResponse) String() string {
  return dara.Prettify(s)
}

func (s EventRebootInstanceResponse) GoString() string {
  return s.String()
}

func (s *EventRebootInstanceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EventRebootInstanceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EventRebootInstanceResponse) GetBody() *EventRebootInstanceResponseBody  {
  return s.Body
}

func (s *EventRebootInstanceResponse) SetHeaders(v map[string]*string) *EventRebootInstanceResponse {
  s.Headers = v
  return s
}

func (s *EventRebootInstanceResponse) SetStatusCode(v int32) *EventRebootInstanceResponse {
  s.StatusCode = &v
  return s
}

func (s *EventRebootInstanceResponse) SetBody(v *EventRebootInstanceResponseBody) *EventRebootInstanceResponse {
  s.Body = v
  return s
}

func (s *EventRebootInstanceResponse) Validate() error {
  return dara.Validate(s)
}

