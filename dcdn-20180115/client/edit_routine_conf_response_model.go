// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditRoutineConfResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditRoutineConfResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditRoutineConfResponse
  GetStatusCode() *int32 
  SetBody(v *EditRoutineConfResponseBody) *EditRoutineConfResponse
  GetBody() *EditRoutineConfResponseBody 
}

type EditRoutineConfResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditRoutineConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditRoutineConfResponse) String() string {
  return dara.Prettify(s)
}

func (s EditRoutineConfResponse) GoString() string {
  return s.String()
}

func (s *EditRoutineConfResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditRoutineConfResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditRoutineConfResponse) GetBody() *EditRoutineConfResponseBody  {
  return s.Body
}

func (s *EditRoutineConfResponse) SetHeaders(v map[string]*string) *EditRoutineConfResponse {
  s.Headers = v
  return s
}

func (s *EditRoutineConfResponse) SetStatusCode(v int32) *EditRoutineConfResponse {
  s.StatusCode = &v
  return s
}

func (s *EditRoutineConfResponse) SetBody(v *EditRoutineConfResponseBody) *EditRoutineConfResponse {
  s.Body = v
  return s
}

func (s *EditRoutineConfResponse) Validate() error {
  return dara.Validate(s)
}

