// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventRedeployInstanceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EventRedeployInstanceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EventRedeployInstanceResponse
  GetStatusCode() *int32 
  SetBody(v *EventRedeployInstanceResponseBody) *EventRedeployInstanceResponse
  GetBody() *EventRedeployInstanceResponseBody 
}

type EventRedeployInstanceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EventRedeployInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EventRedeployInstanceResponse) String() string {
  return dara.Prettify(s)
}

func (s EventRedeployInstanceResponse) GoString() string {
  return s.String()
}

func (s *EventRedeployInstanceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EventRedeployInstanceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EventRedeployInstanceResponse) GetBody() *EventRedeployInstanceResponseBody  {
  return s.Body
}

func (s *EventRedeployInstanceResponse) SetHeaders(v map[string]*string) *EventRedeployInstanceResponse {
  s.Headers = v
  return s
}

func (s *EventRedeployInstanceResponse) SetStatusCode(v int32) *EventRedeployInstanceResponse {
  s.StatusCode = &v
  return s
}

func (s *EventRedeployInstanceResponse) SetBody(v *EventRedeployInstanceResponseBody) *EventRedeployInstanceResponse {
  s.Body = v
  return s
}

func (s *EventRedeployInstanceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

