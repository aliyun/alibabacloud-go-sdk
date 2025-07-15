// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTicketTemplateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableTicketTemplateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableTicketTemplateResponse
  GetStatusCode() *int32 
  SetBody(v *EnableTicketTemplateResponseBody) *EnableTicketTemplateResponse
  GetBody() *EnableTicketTemplateResponseBody 
}

type EnableTicketTemplateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableTicketTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableTicketTemplateResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableTicketTemplateResponse) GoString() string {
  return s.String()
}

func (s *EnableTicketTemplateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableTicketTemplateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableTicketTemplateResponse) GetBody() *EnableTicketTemplateResponseBody  {
  return s.Body
}

func (s *EnableTicketTemplateResponse) SetHeaders(v map[string]*string) *EnableTicketTemplateResponse {
  s.Headers = v
  return s
}

func (s *EnableTicketTemplateResponse) SetStatusCode(v int32) *EnableTicketTemplateResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableTicketTemplateResponse) SetBody(v *EnableTicketTemplateResponseBody) *EnableTicketTemplateResponse {
  s.Body = v
  return s
}

func (s *EnableTicketTemplateResponse) Validate() error {
  return dara.Validate(s)
}

