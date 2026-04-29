// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePolarClawCronJobResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnablePolarClawCronJobResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnablePolarClawCronJobResponse
  GetStatusCode() *int32 
  SetBody(v *EnablePolarClawCronJobResponseBody) *EnablePolarClawCronJobResponse
  GetBody() *EnablePolarClawCronJobResponseBody 
}

type EnablePolarClawCronJobResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnablePolarClawCronJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnablePolarClawCronJobResponse) String() string {
  return dara.Prettify(s)
}

func (s EnablePolarClawCronJobResponse) GoString() string {
  return s.String()
}

func (s *EnablePolarClawCronJobResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnablePolarClawCronJobResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnablePolarClawCronJobResponse) GetBody() *EnablePolarClawCronJobResponseBody  {
  return s.Body
}

func (s *EnablePolarClawCronJobResponse) SetHeaders(v map[string]*string) *EnablePolarClawCronJobResponse {
  s.Headers = v
  return s
}

func (s *EnablePolarClawCronJobResponse) SetStatusCode(v int32) *EnablePolarClawCronJobResponse {
  s.StatusCode = &v
  return s
}

func (s *EnablePolarClawCronJobResponse) SetBody(v *EnablePolarClawCronJobResponseBody) *EnablePolarClawCronJobResponse {
  s.Body = v
  return s
}

func (s *EnablePolarClawCronJobResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

