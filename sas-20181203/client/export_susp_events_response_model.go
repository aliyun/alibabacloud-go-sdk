// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportSuspEventsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportSuspEventsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportSuspEventsResponse
  GetStatusCode() *int32 
  SetBody(v *ExportSuspEventsResponseBody) *ExportSuspEventsResponse
  GetBody() *ExportSuspEventsResponseBody 
}

type ExportSuspEventsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportSuspEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportSuspEventsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportSuspEventsResponse) GoString() string {
  return s.String()
}

func (s *ExportSuspEventsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportSuspEventsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportSuspEventsResponse) GetBody() *ExportSuspEventsResponseBody  {
  return s.Body
}

func (s *ExportSuspEventsResponse) SetHeaders(v map[string]*string) *ExportSuspEventsResponse {
  s.Headers = v
  return s
}

func (s *ExportSuspEventsResponse) SetStatusCode(v int32) *ExportSuspEventsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportSuspEventsResponse) SetBody(v *ExportSuspEventsResponseBody) *ExportSuspEventsResponse {
  s.Body = v
  return s
}

func (s *ExportSuspEventsResponse) Validate() error {
  return dara.Validate(s)
}

