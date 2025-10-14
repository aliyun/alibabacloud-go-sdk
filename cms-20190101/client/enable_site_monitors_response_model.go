// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSiteMonitorsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSiteMonitorsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSiteMonitorsResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSiteMonitorsResponseBody) *EnableSiteMonitorsResponse
  GetBody() *EnableSiteMonitorsResponseBody 
}

type EnableSiteMonitorsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSiteMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSiteMonitorsResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSiteMonitorsResponse) GoString() string {
  return s.String()
}

func (s *EnableSiteMonitorsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSiteMonitorsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSiteMonitorsResponse) GetBody() *EnableSiteMonitorsResponseBody  {
  return s.Body
}

func (s *EnableSiteMonitorsResponse) SetHeaders(v map[string]*string) *EnableSiteMonitorsResponse {
  s.Headers = v
  return s
}

func (s *EnableSiteMonitorsResponse) SetStatusCode(v int32) *EnableSiteMonitorsResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSiteMonitorsResponse) SetBody(v *EnableSiteMonitorsResponseBody) *EnableSiteMonitorsResponse {
  s.Body = v
  return s
}

func (s *EnableSiteMonitorsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

