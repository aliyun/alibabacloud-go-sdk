// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportClientEventsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportClientEventsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportClientEventsResponse
  GetStatusCode() *int32 
  SetBody(v *ExportClientEventsResponseBody) *ExportClientEventsResponse
  GetBody() *ExportClientEventsResponseBody 
}

type ExportClientEventsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportClientEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportClientEventsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportClientEventsResponse) GoString() string {
  return s.String()
}

func (s *ExportClientEventsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportClientEventsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportClientEventsResponse) GetBody() *ExportClientEventsResponseBody  {
  return s.Body
}

func (s *ExportClientEventsResponse) SetHeaders(v map[string]*string) *ExportClientEventsResponse {
  s.Headers = v
  return s
}

func (s *ExportClientEventsResponse) SetStatusCode(v int32) *ExportClientEventsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportClientEventsResponse) SetBody(v *ExportClientEventsResponseBody) *ExportClientEventsResponse {
  s.Body = v
  return s
}

func (s *ExportClientEventsResponse) Validate() error {
  return dara.Validate(s)
}

