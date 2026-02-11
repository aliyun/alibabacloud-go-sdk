// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAlertTemplateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAlertTemplateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAlertTemplateResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAlertTemplateResponseBody) *EnableAlertTemplateResponse
  GetBody() *EnableAlertTemplateResponseBody 
}

type EnableAlertTemplateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAlertTemplateResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAlertTemplateResponse) GoString() string {
  return s.String()
}

func (s *EnableAlertTemplateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAlertTemplateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAlertTemplateResponse) GetBody() *EnableAlertTemplateResponseBody  {
  return s.Body
}

func (s *EnableAlertTemplateResponse) SetHeaders(v map[string]*string) *EnableAlertTemplateResponse {
  s.Headers = v
  return s
}

func (s *EnableAlertTemplateResponse) SetStatusCode(v int32) *EnableAlertTemplateResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAlertTemplateResponse) SetBody(v *EnableAlertTemplateResponseBody) *EnableAlertTemplateResponse {
  s.Body = v
  return s
}

func (s *EnableAlertTemplateResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

