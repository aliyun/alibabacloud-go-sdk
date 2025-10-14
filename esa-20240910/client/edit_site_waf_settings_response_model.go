// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditSiteWafSettingsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditSiteWafSettingsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditSiteWafSettingsResponse
  GetStatusCode() *int32 
  SetBody(v *EditSiteWafSettingsResponseBody) *EditSiteWafSettingsResponse
  GetBody() *EditSiteWafSettingsResponseBody 
}

type EditSiteWafSettingsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditSiteWafSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditSiteWafSettingsResponse) String() string {
  return dara.Prettify(s)
}

func (s EditSiteWafSettingsResponse) GoString() string {
  return s.String()
}

func (s *EditSiteWafSettingsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditSiteWafSettingsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditSiteWafSettingsResponse) GetBody() *EditSiteWafSettingsResponseBody  {
  return s.Body
}

func (s *EditSiteWafSettingsResponse) SetHeaders(v map[string]*string) *EditSiteWafSettingsResponse {
  s.Headers = v
  return s
}

func (s *EditSiteWafSettingsResponse) SetStatusCode(v int32) *EditSiteWafSettingsResponse {
  s.StatusCode = &v
  return s
}

func (s *EditSiteWafSettingsResponse) SetBody(v *EditSiteWafSettingsResponseBody) *EditSiteWafSettingsResponse {
  s.Body = v
  return s
}

func (s *EditSiteWafSettingsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

