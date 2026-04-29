// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawChannelResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnablePolarClawChannelResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnablePolarClawChannelResponse
  GetStatusCode() *int32 
  SetBody(v *EnablePolarClawChannelResponseBody) *EnablePolarClawChannelResponse
  GetBody() *EnablePolarClawChannelResponseBody 
}

type EnablePolarClawChannelResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnablePolarClawChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnablePolarClawChannelResponse) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawChannelResponse) GoString() string {
  return s.String()
}

func (s *EnablePolarClawChannelResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnablePolarClawChannelResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnablePolarClawChannelResponse) GetBody() *EnablePolarClawChannelResponseBody  {
  return s.Body
}

func (s *EnablePolarClawChannelResponse) SetHeaders(v map[string]*string) *EnablePolarClawChannelResponse {
  s.Headers = v
  return s
}

func (s *EnablePolarClawChannelResponse) SetStatusCode(v int32) *EnablePolarClawChannelResponse {
  s.StatusCode = &v
  return s
}

func (s *EnablePolarClawChannelResponse) SetBody(v *EnablePolarClawChannelResponseBody) *EnablePolarClawChannelResponse {
  s.Body = v
  return s
}

func (s *EnablePolarClawChannelResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

