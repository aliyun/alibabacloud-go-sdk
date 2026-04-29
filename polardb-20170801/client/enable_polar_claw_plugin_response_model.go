// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawPluginResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnablePolarClawPluginResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnablePolarClawPluginResponse
  GetStatusCode() *int32 
  SetBody(v *EnablePolarClawPluginResponseBody) *EnablePolarClawPluginResponse
  GetBody() *EnablePolarClawPluginResponseBody 
}

type EnablePolarClawPluginResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnablePolarClawPluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnablePolarClawPluginResponse) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawPluginResponse) GoString() string {
  return s.String()
}

func (s *EnablePolarClawPluginResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnablePolarClawPluginResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnablePolarClawPluginResponse) GetBody() *EnablePolarClawPluginResponseBody  {
  return s.Body
}

func (s *EnablePolarClawPluginResponse) SetHeaders(v map[string]*string) *EnablePolarClawPluginResponse {
  s.Headers = v
  return s
}

func (s *EnablePolarClawPluginResponse) SetStatusCode(v int32) *EnablePolarClawPluginResponse {
  s.StatusCode = &v
  return s
}

func (s *EnablePolarClawPluginResponse) SetBody(v *EnablePolarClawPluginResponseBody) *EnablePolarClawPluginResponse {
  s.Body = v
  return s
}

func (s *EnablePolarClawPluginResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

