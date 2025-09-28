// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventMigrateInstanceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EventMigrateInstanceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EventMigrateInstanceResponse
  GetStatusCode() *int32 
  SetBody(v *EventMigrateInstanceResponseBody) *EventMigrateInstanceResponse
  GetBody() *EventMigrateInstanceResponseBody 
}

type EventMigrateInstanceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EventMigrateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EventMigrateInstanceResponse) String() string {
  return dara.Prettify(s)
}

func (s EventMigrateInstanceResponse) GoString() string {
  return s.String()
}

func (s *EventMigrateInstanceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EventMigrateInstanceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EventMigrateInstanceResponse) GetBody() *EventMigrateInstanceResponseBody  {
  return s.Body
}

func (s *EventMigrateInstanceResponse) SetHeaders(v map[string]*string) *EventMigrateInstanceResponse {
  s.Headers = v
  return s
}

func (s *EventMigrateInstanceResponse) SetStatusCode(v int32) *EventMigrateInstanceResponse {
  s.StatusCode = &v
  return s
}

func (s *EventMigrateInstanceResponse) SetBody(v *EventMigrateInstanceResponseBody) *EventMigrateInstanceResponse {
  s.Body = v
  return s
}

func (s *EventMigrateInstanceResponse) Validate() error {
  return dara.Validate(s)
}

