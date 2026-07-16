// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneGroupTemplateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSceneGroupTemplateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSceneGroupTemplateResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSceneGroupTemplateResponseBody) *EnableSceneGroupTemplateResponse
  GetBody() *EnableSceneGroupTemplateResponseBody 
}

type EnableSceneGroupTemplateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSceneGroupTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSceneGroupTemplateResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneGroupTemplateResponse) GoString() string {
  return s.String()
}

func (s *EnableSceneGroupTemplateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSceneGroupTemplateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSceneGroupTemplateResponse) GetBody() *EnableSceneGroupTemplateResponseBody  {
  return s.Body
}

func (s *EnableSceneGroupTemplateResponse) SetHeaders(v map[string]*string) *EnableSceneGroupTemplateResponse {
  s.Headers = v
  return s
}

func (s *EnableSceneGroupTemplateResponse) SetStatusCode(v int32) *EnableSceneGroupTemplateResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSceneGroupTemplateResponse) SetBody(v *EnableSceneGroupTemplateResponseBody) *EnableSceneGroupTemplateResponse {
  s.Body = v
  return s
}

func (s *EnableSceneGroupTemplateResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

