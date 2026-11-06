// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCloudPermissionTemplateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCloudPermissionTemplateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCloudPermissionTemplateResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCloudPermissionTemplateResponseBody) *EnableCloudPermissionTemplateResponse
  GetBody() *EnableCloudPermissionTemplateResponseBody 
}

type EnableCloudPermissionTemplateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCloudPermissionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCloudPermissionTemplateResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCloudPermissionTemplateResponse) GoString() string {
  return s.String()
}

func (s *EnableCloudPermissionTemplateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCloudPermissionTemplateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCloudPermissionTemplateResponse) GetBody() *EnableCloudPermissionTemplateResponseBody  {
  return s.Body
}

func (s *EnableCloudPermissionTemplateResponse) SetHeaders(v map[string]*string) *EnableCloudPermissionTemplateResponse {
  s.Headers = v
  return s
}

func (s *EnableCloudPermissionTemplateResponse) SetStatusCode(v int32) *EnableCloudPermissionTemplateResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCloudPermissionTemplateResponse) SetBody(v *EnableCloudPermissionTemplateResponseBody) *EnableCloudPermissionTemplateResponse {
  s.Body = v
  return s
}

func (s *EnableCloudPermissionTemplateResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

