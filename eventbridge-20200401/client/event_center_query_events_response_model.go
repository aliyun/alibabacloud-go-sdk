// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventCenterQueryEventsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EventCenterQueryEventsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EventCenterQueryEventsResponse
  GetStatusCode() *int32 
  SetBody(v *EventCenterQueryEventsResponseBody) *EventCenterQueryEventsResponse
  GetBody() *EventCenterQueryEventsResponseBody 
}

type EventCenterQueryEventsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EventCenterQueryEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EventCenterQueryEventsResponse) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsResponse) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EventCenterQueryEventsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EventCenterQueryEventsResponse) GetBody() *EventCenterQueryEventsResponseBody  {
  return s.Body
}

func (s *EventCenterQueryEventsResponse) SetHeaders(v map[string]*string) *EventCenterQueryEventsResponse {
  s.Headers = v
  return s
}

func (s *EventCenterQueryEventsResponse) SetStatusCode(v int32) *EventCenterQueryEventsResponse {
  s.StatusCode = &v
  return s
}

func (s *EventCenterQueryEventsResponse) SetBody(v *EventCenterQueryEventsResponseBody) *EventCenterQueryEventsResponse {
  s.Body = v
  return s
}

func (s *EventCenterQueryEventsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

